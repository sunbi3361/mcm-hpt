// Package relu implements the relu algorithm as a benchmark.
package relu

import (
	"fmt"
	"log"

	// embed hsaco files
	_ "embed"

	"gitlab.com/akita/mgpusim/driver"
	"gitlab.com/akita/mgpusim/insts"
	"gitlab.com/akita/mgpusim/kernels"
)

// KernelArgs defines kernel arguments
type KernelArgs struct {
	Count               uint32
	Padding             uint32
	Input               driver.GPUPtr
	Output              driver.GPUPtr
	HiddenGlobalOffsetX int64
	HiddenGlobalOffsetY int64
	HiddenGlobalOffsetZ int64
}

// Benchmark defines a benchmark
type Benchmark struct {
	driver  *driver.Driver
	context *driver.Context
	gpus    []int
	hsaco   *insts.HsaCo

	Length      int
	inputData   []float32
	outputData  []float32
	gInputData  driver.GPUPtr
	gOutputData driver.GPUPtr

	useUnifiedMemory      bool
	useLASPMemoryAlloc    bool
	useLASPHSLMemoryAlloc bool
}

//go:embed kernels.hsaco
var hsacoBytes []byte

// NewBenchmark returns a benchmark
func NewBenchmark(driver *driver.Driver) *Benchmark {
	b := new(Benchmark)

	b.driver = driver
	b.context = driver.Init()
	b.hsaco = kernels.LoadProgramFromMemory(hsacoBytes, "ReLUForward")

	return b
}

// SelectGPU selects GPU
func (b *Benchmark) SelectGPU(gpus []int) {
	b.gpus = gpus
}

// SetUnifiedMemory uses Unified Memory
func (b *Benchmark) SetUnifiedMemory() {
	b.useUnifiedMemory = true
}

// SetLASPMemoryAlloc use Unified Memory
func (b *Benchmark) SetLASPMemoryAlloc() {
	b.useLASPMemoryAlloc = true
}

// SetLASPMemoryAlloc use Unified Memory
func (b *Benchmark) SetLASPHSLMemoryAlloc() {
	b.useLASPHSLMemoryAlloc = true
}

// Run runs
func (b *Benchmark) Run() {
	b.driver.SelectGPU(b.context, b.gpus[0])
	b.initMem()
	b.exec()
}

func (b *Benchmark) initMem() {
	if b.useUnifiedMemory {
		b.gInputData = b.driver.AllocateUnifiedMemory(b.context,
			uint64(b.Length*4))
		b.gOutputData = b.driver.AllocateUnifiedMemory(b.context,
			uint64(b.Length*4))
	} else if b.useLASPMemoryAlloc {
		b.gInputData = b.driver.AllocateMemoryLASP(b.context, uint64(b.Length*4), "div4")
		b.gOutputData = b.driver.AllocateMemoryLASP(b.context, uint64(b.Length*4), "div4")
	} else {
		b.gInputData = b.driver.AllocateMemory(b.context, uint64(b.Length*4))
		b.driver.Distribute(b.context, b.gInputData, uint64(b.Length*4), b.gpus)

		b.gOutputData = b.driver.AllocateMemory(b.context, uint64(b.Length*4))
		b.driver.Distribute(b.context, b.gOutputData,
			uint64(b.Length*4), b.gpus)
	}
	fmt.Printf("Footprint: %.3f MB\n",
		float64(uint64(b.Length*4)*2)/1024.0/1024.0)

	b.inputData = make([]float32, b.Length)
	b.outputData = make([]float32, b.Length)
	for i := 0; i < b.Length; i++ {
		b.inputData[i] = float32(i) - 0.5
	}

	b.driver.MemCopyH2D(b.context, b.gInputData, b.inputData)
}

func (b *Benchmark) exec() {
	queues := make([]*driver.CommandQueue, len(b.gpus))

	for i, gpu := range b.gpus {
		b.driver.SelectGPU(b.context, gpu)
		q := b.driver.CreateCommandQueue(b.context)
		queues[i] = q

		numWI := b.Length / len(b.gpus)

		kernArg := KernelArgs{
			uint32(b.Length), 0,
			b.gInputData, b.gOutputData,
			int64(numWI * i), 0, 0,
		}

		log.Printf("Launching kernel: relu")
		b.driver.EnqueueLaunchKernel(
			q,
			b.hsaco,
			[3]uint32{uint32(numWI), 1, 1},
			[3]uint16{64, 1, 1},
			&kernArg,
		)
	}

	for _, q := range queues {
		b.driver.DrainCommandQueue(q)
	}

	b.driver.MemCopyD2H(b.context, b.outputData, b.gOutputData)
}

// Verify verifies
func (b *Benchmark) Verify() {
	for i := 0; i < b.Length; i++ {
		if b.inputData[i] > 0 && b.outputData[i] != b.inputData[i] {
			log.Panicf("mismatch at %d, input %f, output %f", i,
				b.inputData[i], b.outputData[i])
		}

		if b.inputData[i] <= 0 && b.outputData[i] != 0 {
			log.Panicf("mismatch at %d, input %f, output %f", i,
				b.inputData[i], b.outputData[i])
		}
	}

	log.Printf("Passed!\n")
}
