// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../avl.go
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _AvlGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4b\x6f\xdb\x46\x10\x3e\x93\xbf\x62\x6a\xa0\x01\xe9\x50\x0f\x2a\x97\x42\x00\x73\x4a\x9a\x06\x79\x14\x48\xd0\x5e\x0c\x01\x61\xc4\xa1\xb5\xd0\x72\xc9\x2e\x57\x7e\x54\xd0\x7f\x2f\x66\x1f\xe4\x52\x22\x9d\xa4\x3d\xd4\x07\x5b\xda\x9d\xf9\x66\xe6\xfb\x66\x5f\x6e\xf2\xed\x3e\xbf\x45\xa8\x65\x51\xe5\x4d\x18\xb2\xaa\xa9\xa5\x82\xab\xb2\x52\x57\x61\xa8\x1e\x1b\x84\xd7\x42\xc9\x47\x68\x95\x3c\x6c\x15\x1c\xc3\xe0\x1d\x3e\x02\xc0\x3b\x7c\x0c\x83\x3f\x73\x7e\x40\xd0\xbf\xc3\x93\x35\xff\x5d\x16\x1f\xf2\xc6\xb3\x37\xfe\x00\x06\x28\x0c\x76\x60\x7f\x98\x50\x61\xb0\xdd\x31\x5e\x48\x14\x70\xb3\xda\x5c\x1b\x5f\x82\x2a\x0f\x62\x0b\x91\x00\x3b\x14\xc3\x6f\xc8\x6e\x77\x2a\x8a\xc9\x8b\x50\x59\x09\x02\xb2\x0c\x04\xe3\xf4\x35\x90\xa8\x0e\x52\xc0\x32\x0c\x4e\xa1\xfb\x22\xe6\xbb\x0e\x6b\x5b\x57\x5f\x99\xc0\xe2\x15\x36\x6a\x17\x89\x34\x01\xb1\xea\xd1\x2d\x68\x91\xc2\x3a\x03\x91\xce\x5d\xb4\x30\x28\x56\x7a\x68\xe5\x0d\xdd\xe5\x12\x0a\x93\x3d\x2b\xa1\x48\xe1\x25\x14\x2b\x9d\x44\x01\x19\x14\x69\x18\x9c\x00\x79\x8b\xfd\xd0\xca\xcf\xaa\x80\xe7\x90\x76\x79\x55\xfb\x8f\x75\x81\x11\x6a\x92\x34\x43\x09\x70\x2c\x95\xcb\x2d\x01\x49\x71\xfb\x54\xed\x07\x02\xb7\x80\xcf\xcc\x08\x45\xd3\x00\x6b\x22\x57\x03\x26\x61\x10\xec\xd6\x8e\xef\x21\x05\x14\xc4\x82\xc7\x64\xe7\x84\x58\x7b\x4a\x1c\x3d\xa3\x53\x42\x45\xf4\xd2\xd4\x05\xf6\x49\xbd\x41\x15\xed\xf1\x91\xba\x22\x86\xe8\xae\x6f\x8b\x04\xea\x3d\x7c\xad\x6b\x1e\x3b\xd1\xc8\xcf\xd3\xad\xde\x43\x06\x65\xce\x5b\xec\x45\x5c\x2c\xe0\xd0\x32\x71\x0b\x22\xaf\xb0\x00\x33\xda\x42\x5b\xc3\x3d\xc2\x1e\xb1\x01\xb5\x43\xf8\x1b\x65\x0d\x26\x54\x59\x4b\xf8\xa2\x3f\x7e\xd1\x4c\xb3\x12\xf6\xf8\x38\x7f\x8f\x6d\xab\x13\x9d\x6b\x5e\xe6\x3a\x3b\xaf\x59\xf4\x94\xab\xfb\x66\xb9\x99\xdb\x32\x62\x07\x32\xf4\x35\x78\xfb\x27\x41\xd2\x33\x90\xc5\x02\xf0\xaf\x43\xce\xc3\x81\xb9\x81\xb4\x04\x29\x69\x56\xcf\x18\xaf\x6f\x45\x8b\xb2\xa3\x36\x01\x8f\xd9\x41\x27\x8c\x10\x6b\xe3\xd9\xfe\xd2\x11\x8f\xfb\x0e\xe3\x94\x90\xa1\xfe\x65\x12\xc5\xbe\xf5\x5c\xcb\x51\xdf\x77\xd9\x26\x17\x6c\x9d\x8f\xa4\x9b\xef\xe0\x4c\x03\x67\x26\xc0\xbc\xaf\xce\x66\x15\x77\x4b\xe7\x5b\x0a\xea\x25\x92\xe9\x74\x9f\x82\x39\x82\x13\xa0\x0d\x03\x53\x23\x64\x70\x41\x86\xbf\x3c\x65\xad\x72\x65\x17\xe4\x80\x8f\x78\x4a\xa4\x4f\x58\xd5\x77\xd8\xf7\xff\x77\xe9\x22\x18\xff\xdf\x78\xef\x13\xfe\x77\x84\x8f\xf9\x0f\x99\xae\xf2\x07\x2a\x43\x5b\x7f\xc8\x1f\x68\xd7\xa4\x1c\x69\xd8\xa3\xa2\xa3\x9c\xb2\x0a\x03\x7f\xdb\x1c\x8b\x57\xe5\x0f\x3a\x1f\x9a\x76\x52\x5e\x57\xf9\x03\x79\xfe\xa0\x82\xbe\xc9\x8f\xee\xba\xac\xb4\x34\xba\x13\x61\xa6\x93\xec\x0e\xa8\x97\x90\x1a\x36\x58\xd5\x70\x86\xad\x85\xfa\x29\x33\x9a\xd3\x96\x40\x9b\x1b\x47\x1d\x91\xc4\xe1\xc4\x95\xc1\xf4\x74\xa6\x19\x39\x32\x43\x7a\x53\x12\xa3\xdb\x39\x8f\x67\x52\x9e\xe5\x42\x7c\x2d\x16\x50\xd4\x87\xaf\x1c\x4d\xe9\xac\x16\x1e\xff\x76\x8f\xa0\x91\x40\x72\xdb\x77\xfa\x9b\x7f\x3a\x75\x54\x72\x3f\xcb\x78\x60\x68\x52\xb5\x84\xfa\x86\xe9\x26\x01\x29\x8d\x71\x6c\x04\x3b\x8b\x3e\x70\x1d\x8f\x1b\x6b\x0c\xb7\x39\x0f\x58\x9f\x0d\x25\xb9\x14\x41\xab\x3b\xa2\x81\xeb\x3d\xce\xbb\x86\x1d\x6a\xc0\xe5\xe5\xc4\xb8\x04\x1a\x2a\x01\x2e\xe3\x19\xe7\xff\x45\x03\x2e\x47\x34\xd0\x19\xb8\x56\xe5\x14\x66\x5a\x06\xc7\x9b\x3c\x13\xa0\x3b\xec\x27\x34\x38\x8f\x71\x0e\xe7\x20\x62\x7f\xb9\x8d\x6a\xf5\xe4\x86\xf9\x1e\xc5\xf0\x22\x37\xbe\x45\x0e\xee\x72\x29\x3c\xbf\x3c\xaf\x0d\xd0\xf9\x44\x6a\x27\xa6\xc2\x53\x81\x0c\xdb\x28\x86\x9b\x8d\xb9\x9c\x1e\xc3\x00\x39\x56\x2d\x29\x5d\xe5\x7b\x8c\xec\x44\x02\x4b\x7b\xdf\x6b\x15\x36\x40\x60\xde\x9d\x34\x0c\xf4\x68\x76\x3e\xae\x6b\x38\xbf\x9f\xda\x3a\x2c\xed\xe4\x18\x89\x81\x80\x74\x3e\xe9\x1c\x32\xc8\x9b\x06\x45\x11\xe9\xaf\x09\x08\xa3\x48\x3c\xe2\x96\x6e\x8c\x12\x66\xbc\x2e\xe8\xe4\xb3\x7c\x69\xe7\x29\x06\xf0\x41\x49\xac\x30\x2a\x98\x24\x19\x62\xb8\xee\x78\xf8\xf6\x89\x55\x32\x71\x8b\xd2\x9d\x52\x61\x40\x97\x2f\x33\xd6\x27\x56\x30\xb9\xb1\xab\x4d\x43\x58\x9f\x6c\xd4\xd0\x97\xf9\x99\x35\x30\x6f\x85\x89\xf4\x3f\x30\x92\xbd\xcf\xd9\xbf\x5a\xb9\xd2\x96\x93\xf2\xeb\xf3\xe8\x1b\xde\xe9\xa4\xf7\x5b\x85\x92\x8e\x0e\xf7\xa9\x96\x3e\x06\xde\xbb\x51\xed\xa6\xfb\xe7\x69\xa0\x4f\x78\x87\xb2\xfd\x6e\x3c\x93\x98\x7e\x68\x31\x3b\xf3\x59\xe5\xdb\xfd\xaf\x32\xaf\xd0\x7b\x74\xe9\x58\x2e\x18\x35\x48\xae\x90\xa4\xfe\xa5\x73\xef\xc2\xf5\x4e\x05\x93\xb8\xa5\x7d\xa9\x7b\x9d\xb5\x84\xed\xde\x0f\x37\x9b\xcb\x90\x61\xb0\x3d\x48\x89\x42\xbd\xee\x8f\xd2\xae\xe2\xf3\xf4\xfb\xd3\xb5\x8f\xa4\xdb\xcf\x2f\xdd\x84\xec\x17\xe2\x65\xcc\x04\x52\x7b\x07\x72\x7b\x6c\x6c\xdd\x6e\x96\x1b\xc8\x46\x88\x39\x92\xf5\x1a\x0c\x85\x9a\x8b\x35\x2c\xe9\x14\x51\xa6\x91\x5d\xfc\x63\x97\xd7\xba\x4f\x51\x7b\x6c\xf7\x6b\xf3\xc7\x7a\xcd\x3f\xe2\x83\x7e\x0c\x5a\xa5\x68\xac\x57\x9a\xc1\xb5\x83\x8c\xe1\x55\x2d\x48\x5e\x7a\x04\x79\xd2\x72\x14\x11\x9b\x6b\xc8\x98\x16\xdb\x72\xc2\xfb\x0d\xaa\x77\xf8\x18\xc5\x74\xb5\xf4\xdc\xd9\xdc\xe7\x9d\xae\x45\xd3\xfe\xfa\xb9\x10\xc5\xe6\xd9\x30\x8d\xd1\x3d\xe3\x47\x50\x4c\xb5\xe4\x4b\xab\x7d\x90\xfb\x4b\x58\x9a\x25\xae\x3b\x70\x9d\xc1\x33\x3b\x75\xe3\x9b\xcd\xf4\x91\xd9\xde\x33\xb5\xdd\x81\x36\x9d\x9b\x9e\x24\xd7\x6d\xde\x22\x2c\xd7\xb4\x4b\xb2\xd2\xce\x9e\x6f\x42\x41\xc0\xf3\x56\x99\xa3\xd8\x8b\x3e\x83\x54\x4f\x76\x31\xf3\x56\x4d\xf4\xc0\x89\xae\x03\xfa\xd1\x58\x1f\x94\xef\x44\xe6\xd6\x7d\x6d\xfc\xfd\x9f\xc5\x02\x9a\xba\x21\x7b\xff\x7a\x1a\xf8\x35\x64\x26\x89\x93\x2b\x25\x35\xa5\x74\xe8\x76\x3f\xb7\x03\xc9\x74\x83\xf6\xa5\xf7\x3b\x24\x9b\x77\x9d\xb8\xf1\x9a\x57\xdf\x82\x87\x49\xac\x5c\xfc\x95\x8d\x3f\x58\x9a\x99\x8f\x6e\xff\x15\x73\x8e\xf0\x62\x70\x50\x69\xac\x17\x6b\x7b\x75\xa9\xef\x50\x4a\x56\xa0\x81\x81\x19\xa8\x9c\x71\xd8\xe6\x9c\x43\xdd\x28\x56\xb1\xb6\xbb\xd2\x4c\x34\xc0\x53\x4b\x73\xac\xf2\x74\x36\x51\x7b\x18\x04\x05\x96\xf9\x81\x2b\x9d\x5c\x93\x0b\xb6\x8d\xca\x4a\xcd\x3f\x37\x92\x09\x55\x46\x57\x7f\x88\xbd\xa8\xef\x85\x71\x81\x9f\xef\xae\x12\xbf\xe9\x62\x73\xfb\x31\xff\xcf\xf8\x27\x00\x00\xff\xff\x88\x9e\xc0\x1f\x02\x13\x00\x00")

func AvlGoBytes() ([]byte, error) {
	return bindataRead(
		_AvlGo,
		"../../avl.go",
	)
}

func AvlGo() (*asset, error) {
	bytes, err := AvlGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../avl.go", size: 4866, mode: os.FileMode(420), modTime: time.Unix(1573405293, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../../avl.go": AvlGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"avl.go": &bintree{AvlGo, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
