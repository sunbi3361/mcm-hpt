// Code generated by "esc -o esc.go -pkg fdtd2d -private kernels.hsaco"; DO NOT EDIT.

package fdtd2d

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/kernels.hsaco": {
		name:    "kernels.hsaco",
		local:   "kernels.hsaco",
		size:    17968,
		modtime: 1601813687,
		compressed: `
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
coj9Xwvs/u7h9v8OAAD//wrpstkwRgAA
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}