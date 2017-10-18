// Code generated by go-bindata.
// sources:
// msg.partial.tmpl
// msg.tmpl
// srv.tmpl
// DO NOT EDIT!

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

var _msgPartialTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xdf\x6f\xdb\x36\x10\xc7\x9f\xc5\xbf\xe2\x3b\x03\x0d\xa4\x54\x93\x8b\x0d\x7d\x71\xa6\x01\x03\x9a\x0d\x03\x96\x6c\x68\x02\xec\x21\x08\x5a\xce\x3a\xc9\x44\x25\x2a\xa0\xa8\x3a\xae\xc0\xff\x7d\x38\x4a\x96\x2d\xd7\x4e\xb2\xa1\x7e\x32\x4f\xf7\xe3\x7b\xbc\xcf\xd1\x6e\x1e\x08\x1f\xae\x9a\xa2\xeb\x90\x5c\xcb\x8a\xe0\x1c\x1a\x6b\xda\xa5\x45\x27\x02\x4b\x8f\x96\x8f\x4a\x17\x22\xd0\xfc\x79\x7b\xa8\xb2\xb7\x4d\x5b\x6d\x8f\x4e\x88\xbc\xd5\x4b\x84\x16\xe7\x07\xd9\x22\xdc\xd2\xa3\x0d\xa3\xc1\x95\xb3\x1a\xb2\xad\xd1\xb0\x09\xa7\x7f\x3a\x96\xff\x1c\x8f\x65\x35\x4f\xc7\x5e\xbd\x7b\x7b\xd3\x56\xc7\xa3\x7b\xf9\xcf\xd4\xa6\xf5\x15\x35\x8d\x2c\x58\x81\xa9\x9b\x64\x38\x71\xa2\x0a\x8b\x14\x9a\xd6\xe1\x7e\x84\x08\xba\x6e\x7e\x8e\xeb\x3f\x6f\x2f\x17\xb8\xae\xa1\x89\x32\xd8\x1a\x4a\x2b\xab\x64\xa9\xbe\x10\x72\x45\x65\xd6\x40\x36\xb0\x2b\xda\x40\x1a\x82\x2c\x4b\x7c\x21\x53\xe3\xb3\x2c\x5b\x8a\xb1\x5e\xa9\xe5\x0a\xaa\x41\x46\xb9\x6c\x4b\x0b\xa5\x51\xd4\x38\x9f\x3b\x37\x36\xe0\x95\x7f\x96\x06\xa1\x08\x0e\x86\x97\xe2\xec\x70\x9e\x9d\x08\x82\x8f\x6c\x78\x2f\xd7\x70\xee\x63\x2c\x82\x60\xc6\xe7\xbf\xe4\xf2\x93\x2c\x68\xf0\x9b\xef\xc5\xcc\x46\x9f\xfe\x16\x07\x93\x13\x91\x10\x9e\x99\xe3\xbc\x74\xdd\xf7\x30\x52\x17\x84\xe4\xd7\xbe\x53\x16\x3d\xd1\xd2\x41\xe5\x48\x7e\x6f\x7e\x31\x46\x6e\xe0\xdc\x5d\x6f\x29\x2c\x12\x6f\xba\xe1\x6b\x7a\x03\xe7\x38\x6a\x67\xf1\x67\xd2\x19\x9c\xbb\x1f\xff\xb1\xcb\x6f\xf5\xed\xe6\x61\xdb\x42\xaf\xa0\xff\xb8\x1b\x6e\x85\xf3\xc9\x60\x6f\xc8\xf4\xe3\x08\xd7\x50\x75\xf2\xb7\x51\x96\x4c\x84\x90\x8c\x01\x19\x53\x9b\xe8\x64\x2f\x00\x9b\xa7\x1d\x88\x60\x3e\x87\x4f\x82\x86\xa5\x96\xca\xda\x92\x58\x85\x92\x5a\x80\x53\x22\xc5\x3f\x4a\x4b\xb3\xe9\x8b\x85\xeb\x78\x7b\xfe\xc3\x3b\x5f\x7a\xdf\x18\xad\xd2\xf6\xc7\x1f\xc2\x92\x74\x58\x25\xfb\x9a\xa3\x48\x04\x2a\xf7\xa9\xbe\x4b\xa1\x55\xe9\x87\x3a\xd0\x90\x57\x36\xb9\x64\xdd\x79\x38\x5b\xd6\x6d\x99\x41\xd7\x16\x6b\xaf\x48\x7a\x91\x25\xe9\xc2\xae\x16\x78\xd5\xcc\x62\x4e\x12\xf1\x28\x81\xbc\x36\xf8\x10\x83\x4a\xf2\x3c\xf7\xdd\x4e\x0a\xa3\x13\x00\x30\x94\x4e\xfd\x1a\x8c\xd7\x37\xec\x83\xbf\x1d\x6e\xc9\xf3\xb2\x37\x8d\x59\x8c\x33\xce\x1d\x5d\x4c\x85\xc3\xff\x06\xf1\x64\x8c\x37\xb8\x5e\x92\x9f\x5f\xd9\x50\x7f\xd9\xff\xbf\xee\xf4\xfa\x8e\x08\x98\x94\x1f\x2b\xf7\xe4\x04\x23\x61\xdb\x8d\x3b\x0d\xd3\x3b\x6a\x46\x9c\x0c\xe3\xf4\x9e\x64\xf6\x52\x9c\x98\x9c\xbd\x6c\xbd\xd7\x01\x5d\xf0\x82\xe7\x73\x70\xe2\x63\x84\x05\xfc\x10\x78\x7b\x8f\xcf\x48\xca\x08\x1d\x47\x86\xe6\x04\x73\x67\x1c\xfa\xc4\x0d\x4d\xe8\x92\x9a\xd1\x32\xac\xa4\x27\xcb\xd7\x65\x8e\xf6\xda\x38\xe0\x6c\xdb\xd4\xd7\x2b\xee\x95\xfa\x0c\x3f\xe3\x70\xdd\x4f\x01\x7e\xba\x2c\x6c\x5d\xa3\x94\xa6\xa0\x05\xe8\xf1\x81\x96\x96\xb2\xf4\x30\x6d\x8c\xa2\xb6\xe9\xab\x6c\x16\xfb\x1c\x3b\x81\x5b\xe8\x82\x29\xff\x29\x2a\xf9\x89\xc2\xbb\xfb\xaf\xde\x9a\x18\x4a\xdb\xd0\x27\x89\x26\xef\x4e\xc0\xba\x14\x2f\xd4\x9b\x0b\x28\xfc\xb4\xf3\xbb\x80\x7a\xfd\xfa\xd8\x4a\xed\x41\x34\x81\xdb\xbc\x00\xee\x3b\x75\xff\x5f\x16\xcc\x89\x69\xbf\xdf\x50\xc9\x0b\xd6\x2c\x78\x6e\xcb\xfe\x0d\x00\x00\xff\xff\xb8\xb1\x60\x8c\x91\x08\x00\x00")

func msgPartialTmplBytes() ([]byte, error) {
	return bindataRead(
		_msgPartialTmpl,
		"msg.partial.tmpl",
	)
}

func msgPartialTmpl() (*asset, error) {
	bytes, err := msgPartialTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "msg.partial.tmpl", size: 2193, mode: os.FileMode(436), modTime: time.Unix(1508163007, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _msgTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbb\x4e\xec\x30\x14\x45\xeb\xf8\x2b\xf6\x75\x75\x29\xe2\xf4\x74\xc0\x80\x48\x33\x83\x34\xf3\x03\x67\x92\x13\xcb\xc2\x2f\xd9\xa6\x88\xac\xfc\x3b\xf2\x84\x87\xe8\xec\xa3\xb5\xb6\xd6\x30\xe0\x29\xcc\x0c\xcd\x9e\x13\x15\x9e\x71\x5d\x91\x42\xee\x35\xfb\x5e\x07\x25\x86\x01\x39\x7c\xa4\x89\xef\x51\x2b\xd4\x8b\xb1\x3c\xfa\x25\xa8\xd1\xb7\xe7\x23\x65\xc6\xb6\x35\xea\x70\xc2\xf1\x74\xc1\xf3\x61\xbc\xfc\x13\x91\xa6\x77\xd2\xfc\x57\x79\xdb\x8f\x47\x72\x37\x47\x18\x17\x43\x2a\xf8\x2f\xba\x5a\x7b\x98\x05\x21\x41\x9d\x23\x4f\xea\x95\xf2\xd9\x9a\x89\x7f\xbf\x0f\x29\xd1\xda\xac\x4e\xb2\x9f\xc2\x6c\xbc\x1e\xae\xc6\x53\x5a\xa5\xe8\xe4\xe2\x8a\xdc\x57\xd8\xcf\x3b\x65\x82\x14\x02\x2d\x20\x91\xd7\xdf\x4b\x5f\x09\xb9\x31\x80\x6c\x79\xd8\x36\x79\x03\x7f\xe4\x3b\x21\x6a\x45\x61\x17\x2d\x15\x86\x74\x59\xab\x48\xa9\x18\xb2\xaa\xb8\x68\xe5\x3e\xd6\xd0\xcf\x00\x00\x00\xff\xff\x3a\xb2\xfc\x15\x3f\x01\x00\x00")

func msgTmplBytes() ([]byte, error) {
	return bindataRead(
		_msgTmpl,
		"msg.tmpl",
	)
}

func msgTmpl() (*asset, error) {
	bytes, err := msgTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "msg.tmpl", size: 319, mode: os.FileMode(436), modTime: time.Unix(1508172091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _srvTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcf\x6f\xda\x30\x14\xc7\xcf\xf8\xaf\xf8\xce\x87\x2a\x54\x6b\x72\xea\x65\xd2\x0e\xdb\xba\x69\x1c\xa0\x53\xe1\xbe\xba\xc9\x23\x8a\x46\x9c\xd4\x76\xa0\x28\xca\xff\x3e\xd9\x71\xa0\x80\x0b\xdc\x9c\xf7\xde\xf7\x93\xf7\xcb\x4e\x12\xfc\xa8\x32\x42\x4e\x92\x94\x30\x94\xe1\x65\x0b\x55\xe9\xbb\x9c\xe4\x5d\x5e\xc5\x2c\x49\xa0\xab\x46\xa5\xf4\x05\x6d\x8b\xf8\x57\xb1\xa2\x89\x5c\x56\xf1\x44\xda\xe3\x77\xa1\x09\x5d\x67\xa3\x1e\x1e\x31\x7b\x5c\xe0\xe7\xc3\x64\xf1\x89\xd5\x22\xfd\x27\x72\x3a\x94\xfc\xe9\x8d\x33\x51\x3a\x0d\x2b\xca\xba\x52\x06\x11\x1b\xb5\xed\x1d\x8a\x25\x2a\x85\x78\x5e\x53\x1a\x3f\xd1\x6b\x43\xda\xb8\xf3\x6f\xa1\xe7\xab\x22\xa5\xb0\xeb\x9b\x52\x62\xbb\x73\xe9\xba\x92\x9a\xc2\xb2\x43\x5f\xaf\xeb\x3a\x36\xe2\x24\xd3\x2a\x2b\x64\x9e\xbc\x14\x52\xa8\x2d\x67\x23\xbe\x2c\x0d\xef\xb3\x22\x99\xf5\x51\x45\xc5\x19\x83\x2d\x48\x09\x99\x0f\x58\x5f\x92\xb6\x31\x00\xb7\xe5\xa2\xeb\xb8\x0b\xdc\x89\xc7\xcc\xf6\x67\x4e\x6a\x6d\xf3\x31\xdb\x9a\x50\x92\x11\x99\x30\x82\xb9\xaf\xbf\x73\xb5\xb6\x52\x87\xf4\xed\x81\x36\xaa\x49\x0d\x5a\x06\x48\x6b\x02\xac\xa9\x90\x39\x03\xca\xec\x5e\x37\xe5\x3b\x83\xa1\x37\x73\x10\xa1\xe8\x75\x61\xd1\xaa\xd2\xf1\x94\xb4\x16\x39\xd9\x6f\xe7\xd1\x41\x4f\xc7\xd8\xb2\x91\x29\x22\x83\xdb\x50\x42\x63\xd8\x43\x34\xf6\xff\x40\x0b\x45\xa6\x51\x12\x26\x76\xf9\x75\x97\xe4\xd3\x87\xfb\x79\x53\x06\x01\xbe\x9e\x8b\x88\x05\xbd\x99\x20\xc0\xd5\x7f\x51\xee\x77\xc7\x96\x1b\x8d\x8f\x1b\xf0\x1e\x37\x74\xef\x0a\x62\xbf\x56\xd7\x20\xf5\x75\xc8\x19\x6d\xfc\xaa\x78\xe0\xb0\x38\x76\x13\x30\x00\x25\x6d\xa2\x13\xad\x9d\xe1\x5a\x28\x44\x0c\x08\xad\xd4\x57\xdc\x04\x57\xad\x65\xa3\x11\x3f\x73\x55\x93\x63\x05\xff\xec\x72\xe1\x3b\x7b\x3f\xdb\xbd\xe7\x79\xe7\x79\x12\x1b\x74\xdd\x73\x6f\x9e\xea\xfc\x18\xe5\x67\x72\xc6\xdf\x77\xd8\x06\xb8\xab\xe4\x6e\xcc\xb9\xdb\x82\x61\xce\x27\x51\xde\xee\x63\x7a\x6e\x20\xa8\x77\xec\x2f\x84\xc6\x6d\x70\x97\xfc\xa4\x0f\xe7\xbe\x9f\xf9\x8d\x1e\x1e\xab\xdd\xd0\x3f\x20\xe9\x6b\x48\x3e\xdd\xce\x3d\x26\x1f\x55\xc6\xda\x16\x86\xca\x7a\x25\x0c\x81\x97\x3a\x8f\x6b\xa1\x4c\x21\x56\xb1\x29\xeb\x15\x3f\x7d\x43\xdd\x43\x1c\x04\xfa\x2e\x5c\x4b\xdc\x3f\xaf\x16\xf9\x3f\x00\x00\xff\xff\xfe\x6b\x1b\x69\x57\x06\x00\x00")

func srvTmplBytes() ([]byte, error) {
	return bindataRead(
		_srvTmpl,
		"srv.tmpl",
	)
}

func srvTmpl() (*asset, error) {
	bytes, err := srvTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "srv.tmpl", size: 1623, mode: os.FileMode(436), modTime: time.Unix(1507566040, 0)}
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
	"msg.partial.tmpl": msgPartialTmpl,
	"msg.tmpl": msgTmpl,
	"srv.tmpl": srvTmpl,
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
	"msg.partial.tmpl": &bintree{msgPartialTmpl, map[string]*bintree{}},
	"msg.tmpl": &bintree{msgTmpl, map[string]*bintree{}},
	"srv.tmpl": &bintree{srvTmpl, map[string]*bintree{}},
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

