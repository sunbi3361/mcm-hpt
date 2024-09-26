// Code generated by "esc -private -o esc.go -pkg bfs ./kernels.hsaco"; DO NOT EDIT.

package bfs

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
		local:   "./kernels.hsaco",
		size:    9608,
		modtime: 1592039480,
		compressed: `
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
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
