package filter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _filter_o = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x4d\x4c\x13\x5b\x14\xc7\xcf\x4c\x29\x94\x07\x8f\x07\xa5\xaf\xaf\xe1\x21\x21\x68\x10\x8b\x54\x5a\x0b\xd4\xc4\x05\xc1\x28\x2c\x48\xc0\x0f\x12\x76\x93\xa1\x19\xa0\x49\x4b\x4a\xa7\x1a\x34\x46\x13\x13\x5d\xb8\x70\x67\xc2\xd2\xb8\x82\x95\x4b\x58\x59\x63\x42\xe2\xc6\xc4\x85\x0b\x17\x9a\xb8\xd0\x84\x05\x0b\x16\x26\xb2\x91\x9a\xb9\xf7\xdf\xce\x70\x66\x6a\x43\xc2\x4d\xe8\x99\xdf\xff\x7e\x9c\x33\xe7\x9c\xb9\x3c\xbc\x3a\x7d\x4d\x55\x14\xaa\x0c\x85\x7e\x92\x4d\xf6\xd8\x08\xd8\xcf\xe3\xf8\x6d\x25\x85\x4a\x61\xa9\xe9\x0b\xb3\xd2\xea\xd3\xc2\x96\x3a\xa5\xde\xe4\x23\x6a\x23\xa2\x2d\x22\x52\x89\x68\xe8\xf4\xbc\xd0\x57\x43\xad\xd2\x76\xfe\x2d\xec\x92\x8f\xc8\x72\x31\x73\x45\xee\xfb\xd7\xf7\x3f\x3d\xf9\x22\xf7\x49\xbe\x4e\x81\x46\xa2\x39\xdf\x8c\x58\x57\x39\xcf\xe9\xa7\x4f\x9c\x2f\xfd\xeb\xc1\x1e\x61\x23\x3e\xa2\xcd\xc3\xcd\xc3\x4a\xec\xa7\x12\x53\x88\x37\x22\x6c\x3a\xba\x5b\x16\xdc\xd1\x25\xb9\x6b\x4f\x70\xe9\x25\xce\x55\x89\x76\xcb\xe5\x72\x84\x25\xe5\xb1\xc8\x15\x51\x09\x7a\x35\x4e\xe5\xb2\xb0\x99\x0e\xc9\x93\x8a\xcc\x6a\x3a\x2c\xfd\xfa\x29\x55\x5d\xcf\xe3\x1f\x15\xf1\x27\x65\x5e\x82\xfd\xc2\xce\xa9\x3d\xf4\x4b\xf0\x59\x70\x37\xe6\x07\xc0\x5d\x22\xbf\xab\xc1\x73\xe0\x08\xbd\x15\x1c\x05\x87\xb1\x7e\x10\x1c\x02\x9f\x07\x07\xc1\x43\xe0\x76\x70\x0c\xdc\x06\xbe\x00\x46\xdd\x82\xc3\xe0\xbf\xc0\x71\x70\x00\x9c\x00\x37\x82\x2f\x82\x1b\xc0\x49\xb0\x2a\xf2\xb0\x1a\x1c\x91\xf9\x53\x43\x74\xcb\xab\x3e\xc1\x3e\xf0\x41\xd9\x59\xdf\x74\xf4\x07\x18\xf5\x8b\xee\xa3\x9e\xe1\x13\xaa\xa7\x8a\xba\xed\x94\x9d\x75\x7b\x8e\xf9\xe3\xf7\x3f\xea\x5b\xb7\xff\xff\x63\xfd\x3f\x84\xfe\x1f\xac\xd3\xff\x03\x88\x87\xf7\x77\xcf\x09\xe5\x43\xf6\x5f\xed\xfe\x8e\x10\xfd\xb1\xbf\x43\xde\xf1\xa1\x5f\xec\xfa\xc6\x58\x7d\xa3\xac\xbe\xfd\x27\x5c\xdf\xef\x9e\xf5\xb5\x96\x07\xa0\xd1\x80\x7d\x9e\xea\x38\xdb\x5a\x13\xaa\xb3\x66\x72\x56\xf6\x45\x7b\x65\xcf\xbd\x1b\x14\xb8\xdf\xa2\x58\x5d\x10\xc1\x5f\x75\xf8\xec\x47\x2b\xdb\x49\xc7\xd4\x57\xaf\xcb\xd9\x63\x7c\xc4\x41\xf3\xea\x51\x7d\x07\xfa\x3e\xd3\x5f\x41\x3f\x60\xfa\x6b\xe8\x3c\x9f\x1b\xd0\x53\x4c\xff\x26\x7e\xfd\xf4\x81\xe9\x9f\xa1\x7f\x62\xfa\x7b\xe8\x07\x2c\xfe\x2d\xe8\xdc\xaf\xf4\xea\x73\x8b\x42\xf7\xbb\xb4\x07\x22\xe7\x01\x97\x7e\x53\xe8\xee\x73\x2e\x09\xdd\x7d\xce\x6d\xa1\x37\xb9\xf4\x2c\xf4\x33\x4c\xdf\x85\xb5\xc2\xff\x87\x88\xb6\x55\x9b\xad\x1e\x78\xc7\xe6\xd7\x95\xa3\xf3\xdd\x8e\x79\xeb\x86\x4d\x39\xb8\xc5\x7a\x88\x15\x8d\xb5\x22\xc5\x0a\x46\x36\x9d\xd5\x4d\x33\xb3\x98\x31\x0a\x5a\x66\x65\xa9\x60\x98\x26\x97\x0d\xa9\xe6\xf4\xbc\x49\xee\xd5\xda\x62\x26\x5b\x34\x0a\xe4\xda\x50\x99\xc8\xe4\xef\x8c\x6a\x39\x3d\x6f\x3d\x24\xc5\x83\xa6\x69\x5a\x36\x93\x36\x56\x4c\x43\xf8\x8a\x19\xcb\xda\x62\x41\xcf\x19\x14\x33\x8b\x85\xa2\xbe\x40\x31\xf3\x6e\xce\xb2\xd3\x13\x13\xc3\x5a\x4a\x9a\x84\xb0\x71\x6d\x4c\x22\x4c\x62\x4c\xaa\xa3\xc0\x11\x89\xf1\x38\xec\xb0\x67\xad\x8f\x3b\xd6\x8f\x7e\x56\xd5\xb1\x8d\x92\xee\x31\x9d\xb7\x9d\x82\xbf\x46\xa6\x8f\xd7\xf0\xd7\xc0\xb8\xaf\xde\x7e\x16\x1c\xef\x58\xeb\xf6\x6e\xf6\xf0\x33\x8f\xf8\x7b\xc1\x2d\x38\xaa\xb2\xbf\xdd\xa1\x7b\xfa\x87\x5f\x7e\xaf\x70\xff\x4d\x35\xfc\xaf\x79\xf8\xf7\x7b\xf8\x4f\xc2\x3f\xaf\x41\x2f\x3e\xb5\x29\xa6\xf3\xfc\x3d\xaa\xb1\x7f\xd9\xef\xbd\x9e\xd7\xef\x19\x34\x76\xbd\x51\x1e\xfb\x67\x99\xce\xdf\xff\x69\x8d\xf7\x5f\xf7\x78\xff\x66\x8f\xf7\x7f\xe1\xe1\xdb\x1a\x6f\xe0\xdf\xf9\xfd\xb7\x3a\xf6\x57\xfe\x2f\xfc\x0e\x00\x00\xff\xff\xaf\xc0\xe5\xde\xa8\x0b\x00\x00")

func filter_o() ([]byte, error) {
	return bindata_read(
		_filter_o,
		"filter.o",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"filter.o": filter_o,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"filter.o": &_bintree_t{filter_o, map[string]*_bintree_t{}},
}}
