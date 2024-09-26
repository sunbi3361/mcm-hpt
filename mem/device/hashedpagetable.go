package device

import (
	"container/list"
	"sync"

	"gitlab.com/akita/util/ca"
)

// // A Page is an entry in the page table, maintaining the information about how
// // to translate a virtual address to a physical address
// type Page struct {
// 	PID         ca.PID
// 	PAddr       uint64
// 	VAddr       uint64
// 	PageSize    uint64
// 	Valid       bool
// 	DeviceID    uint64
// 	Unified     bool
// 	IsMigrating bool
// 	IsPinned    bool
// }

// A PageTable holds the a list of pages.
// type PageTable interface {
// 	Insert(page Page)
// 	Remove(pid ca.PID, vAddr uint64)
// 	Find(pid ca.PID, Addr uint64) (Page, bool)
// 	Update(page Page)
// 	FindAddr(pid ca.PID, vAddr uint64, level uint64) uint64
// 	PageTablePagesAsBytes(pid ca.PID) []uint64
// 	// GetPageTableAsBuffer(pid ca.PID)
// }

// NewPageTable creates a new page table
func NewHashedPageTable(log2PageSize uint64) *HashedPageTableImpl {
	return &HashedPageTableImpl{
		log2PageSize: log2PageSize,
		tables:       make(map[ca.PID]*processHashTableImpl),
		tableSize:    12, // sbin: 4096 hpt entries (covers 8GB region)
	}
}

// PageTableImpl is the default implementation of a Page Table
type HashedPageTableImpl struct {
	sync.Mutex
	log2PageSize uint64
	tables       map[ca.PID]*processHashTableImpl
	memAllocator MemoryAllocator
	entries      *list.List
	entriesTable map[uint64]*list.Element
	tableSize    uint64 // sbin: HPT size in bits
}

func (pt *HashedPageTableImpl) getTable(pid ca.PID) *processHashTableImpl {
	pt.Lock()
	defer pt.Unlock()
	table, found := pt.tables[pid]
	if !found {
		table = newProcessHashTable(pt.log2PageSize, pt.tableSize, pt.memAllocator)
		pt.tables[pid] = table
	}

	return table
}

func (pt *HashedPageTableImpl) FindAddr(pid ca.PID, vAddr uint64, level uint64) uint64 {
	table := pt.getTable(pid)
	return table.findAddr(vAddr, level)
}

func (pt *HashedPageTableImpl) AlignToPage(addr uint64) uint64 {
	return (addr >> pt.log2PageSize) << pt.log2PageSize
}

// Insert put a new page into the PageTable
func (pt *HashedPageTableImpl) Insert(page Page) {
	table := pt.getTable(page.PID)
	// fmt.Println(page.VAddr)
	table.insert(page)
}

// Remove removes the entry in the page table that contains the target
// address.
func (pt *HashedPageTableImpl) Remove(pid ca.PID, vAddr uint64) {
	table := pt.getTable(pid)
	table.remove(vAddr)
}

// Find returns the page that contains the given virtual address. The bool
// return value invicates if the page is found or not.
func (pt *HashedPageTableImpl) Find(pid ca.PID, vAddr uint64) (Page, bool) {
	table := pt.getTable(pid)
	vAddr = pt.AlignToPage(vAddr)
	return table.find(vAddr)
}

// Update changes the field of an existing page. The PID and the VAddr field
// will be used to locate the page to update.
func (pt *HashedPageTableImpl) Update(page Page) {
	table := pt.getTable(page.PID)
	table.update(page)
}

func (pt *HashedPageTableImpl) SetMemoryAllocator(a MemoryAllocator) {
	pt.memAllocator = a
}

func (pt *HashedPageTableImpl) PageTablePagesAsBytes(pid ca.PID) []uint64 {
	table := pt.getTable(pid)
	return table.PageTablePagesAsBytes()
}

func (pt *HashedPageTableImpl) Rearrange(vAddr uint64) uint64 {
	return pt.getTable(0).rearrange(vAddr)
}

func (pt *HashedPageTableImpl) GetRoot(pid ca.PID, chiplet uint64) uint64 {
	return pt.getTable(pid).root[chiplet][0].pAddr
}

func (pt *HashedPageTableImpl) AddOffset(root, vAddr uint64) uint64 {
	return pt.getTable(0).addOffset(root, vAddr)
}

// func (pt *HashedPageTableImpl) NextLevel(vAddr uint64) uint64 {
// 	return pt.getTable(0).nextLevel(vAddr)
// }

// func (pt *HashedPageTableImpl) MoveToLevel(vAddr uint64, level int) uint64 {
// 	for i := 0; i < level; i++ {
// 		vAddr = pt.getTable(0).nextLevel(vAddr)
// 	}
// 	return vAddr
// }

type processHashTableImpl struct {
	sync.Mutex
	root         [][]*hptEntry
	log2PageSize uint64
	tableSize    uint64
	hptEntrySize uint64
	offsetMask   uint64
	memAllocator MemoryAllocator
	entries      *list.List
	entriesTable map[uint64]*list.Element
}

type hptEntry struct {
	sync.Mutex
	pAddr    uint64
	page     Page
	vpnTag   uint64
	valid    bool
	deleted  bool
	children []*treeNode
}

func newProcessHashTable(log2PageSize uint64, tableSize uint64, a MemoryAllocator) *processHashTableImpl {
	t := new(processHashTableImpl)
	t.root = make([][]*hptEntry, 4) // sbin: per chiplet hashed page table
	for i := range t.root {
		t.root[i] = make([]*hptEntry, uint64(1)<<tableSize)
		for j := range t.root[i] {
			e := new(hptEntry)
			e.pAddr = a.allocatePageTablePage(0, 0, 0, 0)
			e.vpnTag = ^uint64(0)
			e.valid = false
			e.deleted = false
			t.root[i][j] = e
		}
	}
	t.log2PageSize = log2PageSize
	t.tableSize = tableSize
	t.hptEntrySize = uint64(9)
	t.offsetMask = ^(^uint64(0) << t.hptEntrySize)
	t.memAllocator = a
	t.entries = list.New()
	t.entriesTable = make(map[uint64]*list.Element)
	return t
}

func (t *processHashTableImpl) rearrange(vAddr uint64) uint64 {
	rearrangedVpn := vAddr / (uint64(1) << (t.log2PageSize + t.hptEntrySize))
	return rearrangedVpn
}

// func (t *processHashTableImpl) newEntry(PID ca.PID, GPUID uint64, vAddr, pAddr uint64) *hptEntry {
// 	n := new(treeNode)
// 	isLeaf := level == t.numLevels-1
// 	if !isLeaf {
// 		vAddr = vAddr >> (t.bitsPerLevel*(t.numLevels-level-1) + t.log2PageSize)
// 		n.pAddr = (t.memAllocator).allocatePageTablePage(PID, int(GPUID), vAddr, pAddr)
// 		fmt.Println("allocated physical page:", n.pAddr, (((n.pAddr-4096)>>12)%32)/8, level) // this arithmetic is off !!
// 		n.children = make([]*treeNode, t.numChildren)
// 	}
// 	return n
// }

// func (t *processHashTableImpl) newTreeNode(PID ca.PID, GPUID uint64, vAddr, pAddr uint64, level uint64) *treeNode {
// 	n := new(treeNode)
// 	isLeaf := level == t.numLevels-1
// 	if !isLeaf {
// 		vAddr = vAddr >> (t.bitsPerLevel*(t.numLevels-level-1) + t.log2PageSize)
// 		n.pAddr = (t.memAllocator).allocatePageTablePage(PID, int(GPUID), vAddr, pAddr)
// 		fmt.Println("allocated physical page:", n.pAddr, (((n.pAddr-4096)>>12)%32)/8, level) // this arithmetic is off !!
// 		n.children = make([]*treeNode, t.numChildren)
// 	}
// 	return n
// }

func (t *processHashTableImpl) insert(page Page) {
	t.Lock()
	defer t.Unlock()
	// if t.root == nil {
	// 	t.root = t.newTreeNode(page.PID, page.DeviceID, 0, 0, 0)
	// }
	// n := t.root
	// vAddr := t.rearrange(page.VAddr)
	// var i uint64 = 0
	// for ; i < t.numLevels-1; i++ {
	// 	indexOfChild := vAddr & t.bitsPerLevelMask
	// 	if n.children[indexOfChild] == nil {
	// 		newTreeNode := t.newTreeNode(page.PID, page.DeviceID, page.VAddr, page.PAddr, i)
	// 		n.children[indexOfChild] = newTreeNode
	// 	}
	// 	n = n.children[indexOfChild]
	// 	vAddr = vAddr >> t.bitsPerLevel
	// }
	// indexOfChild := vAddr & t.bitsPerLevelMask
	// if n.children[indexOfChild] != nil {
	// 	panic("page already present")
	// }
	// n.children[indexOfChild] = t.newTreeNode(page.PID, page.DeviceID, page.VAddr, page.PAddr, i)
	// n.children[indexOfChild].pAddr = page.PAddr
	// n.children[indexOfChild].page = page

	vpnTag := t.rearrange(page.VAddr)
	localHashVal := vpnTag % (uint64(1) << t.tableSize)
	globalHashVal := vpnTag % 4
	e := t.root[globalHashVal][localHashVal]

	topDelEntry := uint64(0)
	for i := 0; i < 16; i++ {
		if e.valid {
			if e.vpnTag != vpnTag { // sbin: tag mismatch
				localHashVal += 1 // sbin: open-addressing
				e = t.root[globalHashVal][localHashVal]
			} else { // sbin: tag match
				break
			}
		} else if e.deleted {
			if topDelEntry == 0 {
				topDelEntry = localHashVal + 1
			}
			localHashVal += 1 // sbin: open-addressing
			e = t.root[globalHashVal][localHashVal]
			continue
		} else if !e.valid {
			if topDelEntry != uint64(0) {
				localHashVal = topDelEntry
				e = t.root[globalHashVal][localHashVal]
			}
			e.valid = true
			e.vpnTag = vpnTag
		}
	}

	indexOfChild := (page.VAddr >> t.log2PageSize) & t.offsetMask
	if e.children[indexOfChild] != nil {
		panic("page already present")
	}
	pte := new(treeNode)
	e.children[indexOfChild] = pte
	e.children[indexOfChild].pAddr = page.PAddr
	e.children[indexOfChild].page = page

	t.pageMustNotExist(page.VAddr)
	elem := t.entries.PushBack(page)
	t.entriesTable[page.VAddr] = elem
}

func (t *processHashTableImpl) remove(vAddr uint64) {
	t.Lock()
	defer t.Unlock()
	vpnTag := t.rearrange(vAddr)
	localHashVal := vpnTag % (uint64(1) << t.tableSize)
	globalHashVal := vpnTag % 4
	e := t.root[globalHashVal][localHashVal]

	for i := 0; i < 16; i++ {
		if e.valid {
			if e.vpnTag != vpnTag { // sbin: tag mismatch
				localHashVal += 1 // sbin: open-addressing
				e = t.root[globalHashVal][localHashVal]
			} else {
				e.valid = false
				e.deleted = true
			}
		} else {
			if e.deleted {
				localHashVal += 1 // sbin: open-addressing
				e = t.root[globalHashVal][localHashVal]
			} else {
				panic("can not remove empty page!")
			}
		}
	}

	t.pageMustExist(vAddr)
	elem := t.entriesTable[vAddr]
	t.entries.Remove(elem)
	delete(t.entriesTable, vAddr)

}

func (t *processHashTableImpl) update(page Page) {
	t.Lock()
	defer t.Unlock()
	vpnTag := t.rearrange(page.VAddr)
	localHashVal := vpnTag % (uint64(1) << t.tableSize)
	globalHashVal := vpnTag % 4
	e := t.root[globalHashVal][localHashVal]

	for i := 0; i < 16; i++ {
		if e.valid {
			if e.vpnTag != vpnTag { // sbin: tag mismatch
				localHashVal += 1 // sbin: open-addressing
				e = t.root[globalHashVal][localHashVal]
			} else { // sbin: tag match
				break
			}
		} else if e.deleted {
			localHashVal += 1 // sbin: open-addressing
			e = t.root[globalHashVal][localHashVal]
			continue
		} else if !e.valid {
			panic("this can not be happend!")
		}
	}

	indexOfChild := (page.VAddr >> t.log2PageSize) & t.offsetMask
	if e.children[indexOfChild] == nil {
		panic("page does not exist!")
	}

	e.children[indexOfChild].pAddr = page.PAddr
	e.children[indexOfChild].page = page

	t.pageMustExist(page.VAddr)
	elem := t.entriesTable[page.VAddr]
	elem.Value = page

}

func (t *processHashTableImpl) find(vAddr uint64) (Page, bool) {
	t.Lock()
	defer t.Unlock()

	elem, found := t.entriesTable[vAddr]
	if found {
		return elem.Value.(Page), true
	}

	return Page{}, false
}

func (t *processHashTableImpl) findAddr(vAddr, level uint64) uint64 {
	return 0
}

func (t *processHashTableImpl) addOffset(root, vAddr uint64) uint64 {
	// fmt.Println(root, vAddr&t.bitsPerLevelMask, (vAddr&t.bitsPerLevelMask)*8)
	return root + (vAddr&t.offsetMask)*8
}

// func (t *processTableImpl) nextLevel(vAddr uint64) uint64 {
// 	return vAddr >> t.bitsPerLevel
// }

func (t *processHashTableImpl) pageMustExist(vAddr uint64) {
	_, found := t.entriesTable[vAddr]
	if !found {
		panic("page does not exist")
	}
}

func (t *processHashTableImpl) pageMustNotExist(vAddr uint64) {
	_, found := t.entriesTable[vAddr]
	if found {
		panic("page exist")
	}
}

func (t *processHashTableImpl) PageTablePagesAsBytes() []uint64 {
	pagesAsInts := make([]uint64, 0)
	queue := make([][]*hptEntry, 0)
	for i := 0; i < 4; i++ {
		queue = append(queue, t.root[i])
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < (1 << t.tableSize); j++ {
			pagesAsInts = append(pagesAsInts, queue[i][j].pAddr)
			for k := 0; k < 512; k++ {
				pagesAsInts = append(pagesAsInts, queue[i][j].children[k].pAddr)
			}
		}
	}

	// fmt.Println(pagesAsInts)
	return pagesAsInts
}
