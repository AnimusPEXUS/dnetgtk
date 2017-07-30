// Code generated by go-bindata.
// sources:
// ui/progress-window.glade
// DO NOT EDIT!

package progress_window

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _uiProgressWindowGlade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x97\x41\x6f\xdb\x3c\x0c\x86\xef\xfd\x15\xfa\x74\xfd\x90\x26\xed\x36\x60\x07\xdb\x05\x3a\xa0\xbd\xec\xb6\x0e\x3b\x1a\x8c\xc4\xd8\x5c\x15\xc9\x93\xe8\xa4\xd9\xaf\x1f\x52\xa7\x6b\x5a\xcb\xb5\x1d\x64\x68\x6e\x4a\xc4\x27\xd4\xcb\x57\x22\x91\xe4\xea\x61\x69\xc4\x0a\x7d\x20\x67\x53\x79\x71\x3e\x93\x02\xad\x72\x9a\x6c\x91\xca\xef\x77\x37\x93\xcf\xf2\x2a\x3b\x4b\xfe\x9b\x4c\xc4\x2d\x5a\xf4\xc0\xa8\xc5\x9a\xb8\x14\x85\x01\x8d\xe2\xc3\xf9\xe5\xec\x7c\x26\x26\x93\xec\x2c\x21\xcb\xe8\x17\xa0\x30\x3b\x13\x22\xf1\xf8\xab\x26\x8f\x41\x18\x9a\xa7\xb2\xe0\xfb\xff\xe5\x73\xa2\x2d\x26\xa7\x8f\x71\x6e\xfe\x13\x15\x0b\x65\x20\x84\x54\xde\xf2\xfd\x0f\xb2\xda\xad\xa5\x20\x9d\xca\x75\xb3\xde\x06\x0a\x91\x54\xde\x55\xe8\x79\x23\x2c\x2c\x31\x95\x0a\x6c\xbe\x70\xaa\x0e\x32\xbb\x01\x13\x30\x99\x3e\x05\xc4\xe3\x3d\x06\xfa\x0d\x73\x83\x03\xe3\x9b\xe4\x79\xe5\x02\x31\x39\x2b\x33\x85\x5b\x85\x7d\x98\xc6\x05\xd4\x86\xf3\x35\x69\x2e\x65\xf6\x71\x36\x1b\x4a\x94\x48\x45\xc9\x32\xbb\x8c\x20\xaa\x24\xa3\x9b\x75\xac\x68\xb7\x9e\xb4\x7c\xda\x6e\x67\x58\x51\xa0\x47\xe1\x77\xbe\x6e\xe9\x3e\xa4\xb6\x31\x66\x09\xbe\x20\x9b\x1b\x5c\xb0\xcc\x3e\x8d\x20\x7c\x23\x7b\x0c\xc2\xae\x1a\x07\xcc\x1d\xb3\x5b\x0e\x64\xbc\x5b\xe7\xa1\x02\x45\xb6\x18\x48\x28\x67\xea\xa5\xed\x83\x5e\x98\x18\x37\xf2\xba\x66\x76\xf6\xda\x3d\xc8\xfd\xb8\x03\x1c\x3d\xd4\xd5\x68\x32\x7c\xa8\xc0\xea\x71\xc9\x0c\x6c\x5c\xcd\x79\xe0\xcd\xf6\x98\x68\x75\x27\xd8\x2a\xcb\x5b\xa5\x69\x1a\xc3\xbc\x59\xbf\x86\x62\xa7\x98\xa3\x91\x82\x3d\xd8\x60\x80\xb7\xcf\x3f\x95\x1b\x0c\x32\xfb\x02\x56\xa1\xe9\x3a\xd4\xc1\x35\xef\xa9\xfb\x58\xd4\xa3\x42\x5a\x61\xc8\x77\x4d\xa2\xef\x17\x92\x69\x53\xb7\xd6\xf7\x15\xa8\x7b\xb2\x45\x7f\xc6\x21\x5e\xc7\xb8\x05\x19\x33\x9e\x7a\xee\xae\xad\x9e\xb7\x2f\x2a\x7a\xfa\x64\xda\x7e\x50\x11\xf9\x71\xe9\xad\x8b\x82\x0b\xce\x81\x19\x54\xf9\xc6\x59\x5e\x53\xec\xaa\xbf\xd0\xe5\x50\x68\x37\x19\x3a\xe2\x23\x62\x5b\x42\x87\x74\x92\x6f\x15\x59\x8b\xbe\x79\x2f\x61\xf7\xe1\x84\x9a\x0a\x28\xa6\xd5\x9b\xb9\xde\xd3\xcc\x8b\x7f\x6b\xce\x57\x5c\xa1\xb9\x86\x9d\x3b\x95\x77\x85\xc7\x10\x4e\xc9\x9e\xb2\xbf\x0f\x1c\xc9\x9f\x8e\x52\xbf\xaf\x3f\xcd\xd8\xd8\x9a\xd3\x4c\x90\x53\x72\x06\x0c\x15\x56\x66\x81\xc1\xf3\x11\xed\x3c\xda\xdc\xef\x9a\xb8\x8f\x3b\xa7\x77\x95\x3a\xfa\xc3\x81\x57\xa9\x32\xa0\xb0\x74\x46\xa3\x9f\x76\x92\x2f\xe5\xee\x6d\x3e\x6f\x24\xd3\xbd\xff\x53\x7f\x02\x00\x00\xff\xff\x44\x53\xcc\x60\xa8\x0d\x00\x00")

func uiProgressWindowGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiProgressWindowGlade,
		"ui/progress-window.glade",
	)
}

func uiProgressWindowGlade() (*asset, error) {
	bytes, err := uiProgressWindowGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/progress-window.glade", size: 3496, mode: os.FileMode(420), modTime: time.Unix(1490480157, 0)}
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
	"ui/progress-window.glade": uiProgressWindowGlade,
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
	"ui": &bintree{nil, map[string]*bintree{
		"progress-window.glade": &bintree{uiProgressWindowGlade, map[string]*bintree{}},
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

