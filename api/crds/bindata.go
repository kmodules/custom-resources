// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// appcatalog.appscode.com_appbindings.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _appcatalogAppscodeCom_appbindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xdb\x6f\xdb\xb8\xd2\x7f\xcf\x5f\x31\xf0\x3e\xa4\x05\x7c\x41\xbf\x7d\xf9\xe0\xef\x61\xbf\x34\x4d\x81\xee\xb6\x69\x11\x67\x7b\xb0\x38\x3d\x38\xa1\xc4\x91\xc5\x8d\x44\x6a\x49\xca\x8e\x4f\xd1\xff\xfd\x60\x48\xea\x62\x9b\x92\x9d\xee\xae\x1e\xda\x98\x97\xe1\xcc\x6f\xee\x94\x58\x25\x3e\xa3\x36\x42\xc9\x25\xb0\x4a\xe0\x93\x45\x49\xbf\xcc\xfc\xf1\x7f\xcd\x5c\xa8\xc5\xe6\x55\x82\x96\xbd\xba\x78\x14\x92\x2f\xe1\xba\x36\x56\x95\x77\x68\x54\xad\x53\x7c\x83\x99\x90\xc2\x0a\x25\x2f\x4a\xb4\x8c\x33\xcb\x96\x17\x00\xa9\x46\x46\x83\xf7\xa2\x44\x63\x59\x59\x2d\x41\xd6\x45\x71\x01\x50\xb0\x04\x0b\x43\x6b\x00\x58\x55\x2d\x21\x65\x96\x15\x6a\x7d\x01\x20\x59\x89\xc4\x42\x95\x08\xc9\x85\x5c\x9b\x39\xab\xaa\x30\x4d\x7f\x9a\x54\x71\x9c\xa7\xaa\xbc\x30\x15\xa6\x44\x82\x71\xee\xce\x66\xc5\x27\x2d\xa4\x45\x7d\xad\x8a\xba\x94\x8e\xfc\x0c\x7e\x5e\x7d\xbc\xfd\xc4\x6c\xbe\x84\x39\x6d\x98\xdb\x5d\x85\xee\x5c\x7f\xd0\x7d\xf3\x93\xc6\x97\x60\xac\x16\x72\x1d\xdd\xb8\xf1\xf0\xf4\xf6\x7e\xee\x8d\x8c\x6d\x6f\x30\x99\x1f\x01\xd2\x23\x76\xb5\xee\xf3\xc1\x99\xa5\x9f\x6b\xad\xea\xca\xa1\x11\x45\xc0\xef\x0d\x38\xa6\xcc\xe2\x5a\x69\xd1\xfc\x9e\xf5\x40\xa5\x5f\xcd\xce\xe6\xa7\x53\x04\x80\xd7\xe7\x55\x55\xbd\xf6\x78\xbb\xc1\x42\x18\xfb\xcb\xc1\xc4\x7b\x61\xac\x9b\xac\x8a\x5a\xb3\x62\x4f\x47\x6e\xdc\x08\xb9\xae\x0b\xa6\xfb\x33\x17\x00\x95\x46\x83\x7a\x83\xbf\xca\x47\xa9\xb6\xf2\xad\xc0\x82\x9b\x25\x64\xac\x30\xc4\x8b\x49\x15\x09\x7c\x4b\x82\x54\x2c\x45\x4e\x63\x75\xa2\x83\x69\x99\x25\x7c\xfd\x76\x01\xb0\x61\x85\xe0\x0e\x3c\x2f\x9d\xaa\x50\x5e\x7d\x7a\xf7\xf9\xc7\x55\x9a\x63\xc9\xfc\x20\x1d\xa6\x2a\xd4\xb6\x05\xc1\x1b\x58\x6b\xda\xed\x18\x00\x47\x93\x6a\x51\x39\x8a\x70\x49\xa4\xfc\x1a\xe0\x64\xcc\x68\xc0\xe6\x08\x41\xe7\xc8\xc1\xb8\x63\x40\x65\x60\x73\x61\x40\xa3\x13\x4b\x5a\xc7\x52\x8f\x2c\xd0\x12\x26\x41\x25\xbf\x63\x6a\xe7\xb0\x22\xd1\xb5\x01\x93\xab\xba\xe0\x90\x2a\xb9\x41\x6d\x41\x63\xaa\xd6\x52\xfc\xa7\xa5\x6c\xc0\x2a\x77\x64\xc1\x2c\x06\xa0\x9b\xc7\x19\xb5\x64\x05\x81\x50\xe3\x14\x98\xe4\x50\xb2\x1d\x68\xa4\x33\xa0\x96\x3d\x6a\x6e\x89\x99\xc3\x07\xa5\x11\x84\xcc\xd4\x12\x72\x6b\x2b\xb3\x5c\x2c\xd6\xc2\x36\xce\x9c\xaa\xb2\xac\xa5\xb0\xbb\x45\xaa\xa4\xd5\x22\xa9\xad\xd2\x66\xc1\x71\x83\xc5\xc2\x88\xf5\x8c\xe9\x34\x17\x16\x53\x5b\x6b\x5c\xb0\x4a\xcc\x1c\xe3\xd2\xba\x88\x50\xf2\x1f\x5a\xf5\x5c\xf6\x38\x3d\xf0\x01\xff\x38\xfb\x1a\xc4\x9d\x8c\x0c\x84\x01\x16\xb6\x79\xfe\x3b\x78\x69\x88\x50\xb9\xbb\x59\xdd\x43\x73\xa8\x53\xc1\x3e\xe6\x0e\xed\x6e\x9b\xe9\x80\x27\xa0\x84\xcc\x50\x7b\xc5\x65\x5a\x95\x8e\x22\x4a\x5e\x29\x21\xad\xfb\x91\x16\x02\xe5\x3e\xe8\xa6\x4e\x4a\x61\x49\xd3\x7f\xd4\x68\x2c\xe9\x67\x0e\xd7\x4c\x4a\x65\x21\x41\xa8\x2b\x72\x51\x3e\x87\x77\x12\xae\x59\x89\xc5\x35\x33\xf8\xb7\xc3\x4e\x08\x9b\x19\x41\x7a\x1a\xf8\x7e\x24\xde\x5f\xe8\xd1\x6a\x87\x9b\x38\x1a\xd5\x50\xe7\xff\xab\x0a\x53\x52\x15\xe1\x45\x5b\x20\x53\x9a\x3c\xbd\xb7\x33\xe6\x7d\x2e\x34\x39\x78\xaf\x95\xcc\xc4\x7a\x7f\xe6\xe0\xb4\xeb\xde\xc2\xd6\x11\x73\xb5\x25\xe7\x08\xe0\x51\x98\x83\xad\xb0\xb9\x63\x84\x55\xd5\x1c\xee\xf0\x8f\x5a\x68\x17\x39\xfa\xcf\x10\x37\x8e\x23\xf6\xba\x96\xbc\xc0\xe3\x99\x43\x8e\xae\xfc\x42\x6f\xa4\x9f\x6e\x3e\x00\x4a\x8a\xa2\x1c\xae\xaf\x20\xf1\x53\xdb\x5c\xa4\x39\x6c\x45\x51\x38\xcb\x30\x47\x9c\x04\xf0\x55\x13\xc5\xd0\x83\x88\x7a\x43\xf6\x9d\x12\x93\x99\x17\xac\x89\x2f\x24\x57\x84\x48\xa6\x74\xc9\xec\x12\x92\x9d\xc5\xc8\x74\xd4\x0e\x9a\x47\x48\x83\x69\xad\x71\xf5\x28\xaa\xfb\xf7\xab\xcf\xa8\x45\xb6\x3b\x29\xff\xbb\xd8\x2e\xe0\xc2\xb0\xa4\x40\x03\xf7\xef\x57\x7b\xfc\x6f\x68\x9e\xfe\x3c\x8c\x8a\xcd\xb3\xcd\x51\xf6\x54\x49\xf2\x07\x65\x06\xa9\xe1\x9e\xfe\x12\x86\xc4\x50\x72\x5d\xb8\xc3\x52\x55\x6b\xb6\x26\x77\x83\xdf\x54\x1d\x25\x1c\x02\x6c\x6d\x3c\xb8\x9d\xde\xa4\xb1\xc8\x78\x0c\x4d\x0f\x57\xa2\x54\x81\xec\x98\x5b\xa7\x9e\xf4\xb4\x85\x4c\x1e\xc2\xca\x07\x6f\x23\x1a\x33\xd4\x28\x29\x4c\xa9\x4e\xcf\x29\x3a\x7f\x19\x53\x2e\xc0\x8d\xb0\x39\x6a\xe8\x08\x2a\x0d\x0f\xb5\x2e\x1e\xa0\xac\x8d\x0b\x3b\xe4\x78\x22\x13\x84\xc4\x17\x09\xef\x32\x77\xc0\x16\x93\x5c\xa9\xc7\x28\x49\xca\x55\xb5\x94\x0d\xce\x42\x86\x78\x57\x1b\x8b\x7a\x4a\x3f\x24\xec\x54\xdd\x87\xaf\x3d\x7e\x3e\x89\x90\x1c\xf3\x2a\x68\xaa\x99\xe8\xcc\x61\xec\x7f\xa0\xa5\x0f\x4d\x48\xa1\x1f\xde\xfc\x5b\xc4\x3a\xcf\xbe\x1c\x20\x38\x6a\xf0\x8e\x5b\xaa\xc0\xce\xe3\x86\x96\x7a\x15\x4a\x50\x95\x2f\x28\xe1\xd7\xbb\xf7\x8e\xc6\x81\x8f\x9b\xc3\x6c\xb1\x07\xb9\x04\x26\x77\x4d\xe2\xf0\x56\x40\xf6\x1c\x84\xfa\x7e\x59\x94\xb6\x67\xc9\x72\x9f\xa3\x5b\x0c\x36\x67\xb6\xe5\x19\x9f\x2a\x65\x90\x43\xb2\x3b\x61\x85\x5d\x98\x11\xd2\xfe\xf8\x3f\xa3\xec\x52\x69\xb2\x46\x1d\x5d\xf3\x47\x8d\x3a\x1a\x60\x8e\x18\xbe\x7c\x70\x6b\x1d\xfa\x2d\xf4\x4d\x9c\x75\x53\x01\x97\xa9\x33\x62\x55\x0f\x83\x7f\x79\xf9\xd3\xe5\x65\x44\x5b\x7f\x9b\x56\x5c\xf9\x76\xa6\xc5\xaf\x82\xf7\x9a\xc0\xa0\xdf\x4b\xbc\xd4\x06\xa7\x2e\x40\xe0\x13\x2b\xab\x02\x7d\xf9\x30\x1d\x14\xd3\x15\x17\xe4\xff\x6d\x40\x08\xbe\x2c\x82\xc2\x59\x55\x15\x02\x39\x30\x43\x05\x78\x26\x9e\xc0\xb9\xfe\x41\xdd\xd4\x7f\x1a\xa5\x07\x81\x16\x0b\x22\x4f\xd5\xce\xe1\x11\x52\x51\x1c\x59\xb7\xf0\x7a\xfa\xdf\xed\xa4\x3a\xf8\x78\x0c\xc2\x99\x0b\x0b\xd1\x09\x32\xf0\xe8\x84\xe7\x7f\x30\xdc\x1f\x14\x3f\xcd\x53\xeb\xe2\x8c\x48\xef\x62\xf1\x5a\x6c\x42\x7b\x50\x28\x9f\xe9\x9a\xb8\xc5\xaa\x6a\x4a\x38\x1b\xcb\x24\x67\x9a\x53\xf8\x88\x82\x42\x58\xc3\x8b\x87\x7f\xb6\x58\xff\x2b\x57\xc6\x2e\x49\xa6\x85\x8b\x43\x2f\xe7\x70\xf3\xc4\x52\x5b\xec\x40\x49\x17\x17\xfd\xd9\xaa\x97\x1d\xa2\x94\xe3\x89\x82\x22\xc2\x03\x1d\xf1\xd0\x04\x7a\x52\xac\xcb\x54\x64\x7d\xac\x71\x83\x28\xc9\x26\x7f\xec\xe7\x8e\xff\x6b\x53\x6d\x97\xae\x32\xea\xed\xda\x8c\xeb\x4e\xa5\x43\xe3\x8c\x8a\x75\xee\x38\xa5\xaa\xbe\xd8\x50\xeb\x22\x18\xe0\x53\x68\x75\xde\xdc\xae\x1c\x92\xaa\x24\x58\x85\x09\xd5\xfc\x0b\x9c\xaf\xe7\x53\x78\x78\xac\x13\x9c\xb5\xe3\x71\x28\x52\x5f\xac\x07\xfa\x20\xe4\x2c\xb0\xee\x88\x53\xc7\xe5\xc2\xa3\x83\x23\x41\x60\x50\xb0\x1d\xfa\x26\x44\xa8\xc2\x29\xf6\x65\x3c\x42\x06\x28\xa9\xb5\x60\x85\x51\x6e\xb7\x84\x77\x9f\x80\x71\xae\xd1\x18\x87\xf9\x95\x4f\x1c\xbd\x90\xe6\x3b\x37\x91\xc5\xa3\xbb\xef\x5c\x1c\x51\x47\xaf\x89\x79\x50\xa1\x2e\x85\x31\x22\x71\xd5\x0c\x30\xb2\xaa\x39\xd5\x41\x6e\x6d\xd0\xc2\x60\xf6\x23\xfd\x56\xcc\xb8\xb4\xc6\x74\x22\xac\x66\x6d\x38\x6d\x2a\x14\x67\xb7\xbd\xe8\x33\x05\xd6\xa8\x39\x5e\x54\x70\xea\x49\x32\x81\xda\x4b\x6a\x2d\x96\x95\x0d\x04\x89\x21\x46\xff\x6a\xb2\xd6\x84\x19\x91\x02\xab\x6d\x0e\xa4\x3a\xf8\x32\xa1\x99\x25\x71\xb4\x55\x9a\xff\xff\x97\x58\x8d\xe1\xca\x16\xd2\x1d\x2b\x0a\xb5\x25\x1b\x7e\xab\xd9\xba\xa4\xc6\x0e\x5e\x7c\x99\xfc\x30\x9f\xcf\xbf\x4c\x5e\x3a\x34\x7d\x76\xa8\x98\x66\x25\x5a\x67\x21\x5f\x26\x3f\xf9\xf9\x28\x61\xa6\xb1\x4f\x79\x0a\xe8\x6a\xae\x68\xa9\x33\x12\xb8\x06\x63\x49\xc7\xc9\x68\xa3\x33\xf9\xd4\x71\xec\xdb\x5f\xb4\x4d\x14\xe9\x09\x63\x55\xd3\x51\xf8\x0e\x48\xca\x58\xec\xea\xb4\xe8\x7d\x4e\xc8\x42\x48\x84\xdf\xae\x3e\xbc\x5f\xfc\xbc\xfa\x78\x0b\x15\xdb\x15\x8a\xf1\x40\xce\x6a\x26\x4d\x41\xdd\x2b\xa5\x6f\x05\x14\x7f\x37\xac\x88\x95\x34\x6e\x77\x73\x95\x11\xe2\x48\x8f\xf3\xe0\xef\x06\x6e\x3f\xde\x83\xc1\x54\x93\x10\x1a\x7c\xc7\xc0\x43\xca\x3d\x22\xba\x25\xb7\x91\xbc\x89\x44\xb7\x37\x9f\x6f\xee\xfa\x62\xe6\xaa\xe0\x94\xb3\x8d\xb0\x62\xe3\xbb\x69\xca\x4c\x42\xc9\x39\xdc\x2b\x42\xea\x88\x64\x1f\x32\x72\x6a\x6a\xaf\x19\x85\x0f\xcf\x53\x8f\xc4\xb4\x5f\xed\x5e\xbd\xff\xc7\xd5\x6f\x2b\x30\x56\xe9\x63\x07\x72\x84\x7a\x3b\xbd\xef\xad\x1c\xc5\x23\x73\x19\xb4\x07\xcf\xc1\xa8\x2d\xac\x02\x93\xb1\x3a\xd8\xcd\x90\xea\x35\x52\x7b\x15\x42\x70\xd7\x94\x5f\x1e\xe7\x73\xd9\xdc\xa4\xf5\xca\x3f\x07\xa9\x0b\xde\x1a\x9d\xef\xb2\xc2\x00\x33\x46\xa5\xc2\xd9\x41\xdb\x4f\x77\x94\x0f\x23\xdf\x58\xdd\x3f\x54\xf3\xef\x57\x3f\xb7\x3d\xc9\x42\x93\x64\xa3\x37\x26\x14\xdc\xb5\x44\x8b\xee\xd2\x84\xab\xd4\x2c\x52\x25\x53\xac\xac\x59\xa8\x0d\x25\x1b\xdc\x2e\xb6\x4a\x3f\x0a\xb9\x9e\x11\xeb\x33\x0f\xbc\x59\x38\xd9\x17\x3f\xb8\xff\xa2\xee\x7f\xff\xf1\xcd\xc7\x25\x5c\x71\x0e\xca\xb5\x5a\xb5\xc1\xac\x2e\xbc\x21\x9b\x79\xef\xaa\x70\xea\x2e\xae\xa6\x50\x0b\xfe\x53\xac\xb0\xf9\x9e\xd8\xe0\xd5\x79\x4f\xee\x47\x56\x35\x1e\x21\xde\x0b\xe3\x23\x42\xb3\xdc\x19\x61\xb0\xef\x60\xbf\x09\xb6\x75\x5e\x88\x01\x3d\xfd\x1e\x31\x1d\xd3\xf7\xca\xa7\xee\xa0\x73\x48\x30\x23\x75\xd8\x1c\x77\x2e\x52\x0a\x69\x50\xb7\x81\x22\x96\x66\x82\x3f\x1c\x8c\x0b\x8b\xc7\xe2\x1d\x55\xc3\xfb\x70\x84\x38\x28\xe4\xba\xc0\x03\xa9\x83\x2f\x9a\x46\xda\x98\x3e\x8e\xe4\x07\x8d\xb6\xd6\x12\x79\x77\xe7\x97\x68\xf5\x88\x7a\x50\xca\x08\xd9\x46\xee\xc6\x49\x4f\x63\x38\x87\xd7\x98\x32\x4a\x82\x5c\x64\xde\xc8\x23\x74\x3d\x27\x54\x9b\xab\x8d\xe0\xcd\x2d\xa7\x21\x0f\x21\xf3\x21\xc5\x37\xd7\x06\x94\xe4\x91\xa5\x79\x90\x07\xd8\x28\xe1\x3e\x00\xc6\xea\xda\x5d\x25\x4e\x5d\x3a\x36\x54\x11\x85\xc2\x70\xe7\xce\x8b\xd9\x56\x84\xe6\xa0\xb5\xad\xda\xf8\xc4\x38\xab\x2c\x08\x6b\x00\xa5\xd5\xd4\xe1\x58\x05\xdb\x9c\x59\xdc\x44\x6b\x88\xfe\xbd\x48\xaa\xa4\xa9\x4b\xa4\xea\xa3\x22\x2f\x9e\xc3\xdb\x7e\x29\x32\xa4\xd6\x18\xaa\xbb\xbe\x9a\xfd\xcd\x6f\x5a\xd4\xdc\xd7\xa9\x8f\xb8\x83\xc9\xaf\xab\x9b\xbb\xdb\xab\x0f\x37\x93\x29\x24\x75\xb8\xfc\x6d\xce\x0f\x9d\x48\x2c\x72\xd0\x3a\xc2\xd0\x45\x67\x9f\x46\x9b\x7e\xba\x96\xdc\x5d\x2e\x87\x03\xde\xbc\xfe\x37\x9d\x31\xe9\x95\xc1\x0a\x72\xb6\x89\x76\x24\x9d\xf5\xc0\xb5\x7f\x59\xd3\xe9\xa4\x87\xb0\x07\x21\x53\x54\xb3\x90\xad\x1c\x78\x4e\x84\xf2\x51\x1b\x40\xa9\xe3\xc0\x50\xdd\x5b\xad\x83\x98\xb4\x84\x19\x7c\x9d\x68\x24\x39\x7f\xc1\xdd\x24\x16\xd5\xbf\x4e\xc8\xa1\x26\xcb\x3d\x30\x27\x56\xd1\x48\x23\xfd\xb7\x6f\xf0\x51\x76\xcd\x4b\x27\x4a\x7b\xd2\x65\x24\x75\x01\x94\x58\x26\xcd\xad\xfd\x5e\x17\x73\x1c\x83\xc7\x2f\xa2\x18\xe7\xbf\xe0\xc0\xed\xc3\xfe\x05\xb7\x5b\xd8\x7b\x75\x00\x6c\x4f\x07\xcc\x12\x2d\x5f\x3e\xb7\x2f\x1a\x07\x3a\x5d\x32\x80\x48\x20\xf2\x92\x0f\x54\xfd\xe3\xb7\x69\x00\xbf\x1b\x25\x3f\x31\x9b\xdf\x3c\x11\x83\x87\x6f\xb1\x46\x04\xbb\xa4\x62\xad\x79\x0b\x49\xd6\x1a\xb6\x4f\x43\x02\x36\x75\xe1\x12\x8c\xbb\x8a\x18\x24\x09\x6d\x9b\xcf\x38\x3f\xf2\xfc\xce\xf6\x3b\x43\x7b\xc4\x9d\xf3\xe0\x11\x92\xad\x6f\x53\x2b\x2d\x0f\x8c\xbb\x87\xdc\x12\xbe\xc2\x24\x53\x64\x59\x5f\x61\x92\x30\x1d\xb5\xc7\xe6\xa1\x95\xb4\x06\xbe\xc1\x37\x57\x60\x12\xdd\x63\xf8\x60\xf2\x75\x9e\x29\x35\x4f\x98\xfe\x36\x99\x0e\xf6\x4e\xfe\xf1\xaf\xa3\x5a\xca\xed\x8d\x12\x55\x8c\xbc\xed\x88\xcf\xd3\xb5\x7f\x86\x20\x1b\xba\x3d\x39\xe3\x0a\xca\x19\xde\x99\x56\x71\x7f\x50\x62\x06\x93\x65\x3c\xfe\xce\xe2\xcc\xf3\xfd\xe4\x67\x02\xeb\x19\x7c\x84\x26\xf4\x85\x54\x72\x96\x08\xc9\xf4\xee\x65\x00\xdc\x73\xb4\x6f\x6c\xdf\x81\xe9\x9f\x11\x69\xf3\x4c\x61\xbc\x00\x81\xff\x17\x95\x72\x0d\xfb\x0e\x48\x36\x7f\xce\xcb\x53\x38\xc3\xb9\xce\xf5\x2e\x83\x44\xd9\x3c\x9c\xc5\xe4\x18\xc9\x9e\x66\x5c\xad\x73\x78\x99\xe7\x69\x08\x03\x62\x2d\x9d\x4d\xbb\xce\xac\xdb\x34\x42\xda\xbd\xa3\xa1\x3d\xc3\x38\x9f\x78\x5f\x15\xa4\x3e\xa5\x8c\xb1\x1b\x42\x80\x59\xc4\xc7\x07\x16\x3e\xe2\x71\x37\xea\x67\x4e\x49\x3c\xf3\x40\xc5\xdf\xeb\x8d\xdd\x2b\x36\x99\xc8\xbc\xd5\xaa\x3c\x3b\x1d\xb9\xd5\xa3\x39\xa9\x44\xbd\x46\xd3\x7e\xc5\x11\xe1\xca\xbd\xe3\xf6\xb5\x98\xff\x24\x01\x9f\x84\xb1\x5d\xf9\xd0\xd5\xb6\x7f\x59\xae\xf2\xc5\xc4\x1d\x66\xcf\xf0\x9b\xa3\xd7\x65\x4d\x51\xb9\xdf\xe7\x38\x79\xc7\x0c\x7d\x44\x9a\x61\xfb\x3c\x2d\x12\x9c\x78\xa5\x15\x91\x2a\xda\xe7\x8e\x6e\x3f\x23\x1c\x41\xbf\xb3\x7f\x26\x33\xfe\x36\xe0\xaf\xe7\xe8\x84\xe1\x9f\x5c\xa0\xb1\x54\x1b\x3c\xaf\x4c\xbb\x6b\xd6\x8e\x7a\x85\xa7\x48\x13\x63\x8d\xb0\x7f\x82\x9d\x91\x8f\xc4\xa3\xc2\x69\xe3\x78\x5e\xda\x0d\xa9\xd6\xf3\xd8\x35\xa5\x27\xb2\xdb\x9f\x8e\x8e\x43\x41\xef\xa4\x72\x42\x0f\x70\x96\x72\xc2\xda\x13\xca\x71\x06\xfc\x6c\xe5\x5c\x9a\x41\x19\x4e\xab\x28\x1b\x0c\xbb\x47\x52\x0c\x94\x46\x9e\xed\x3f\x53\x4a\x58\xf5\x1c\x0e\x70\xeb\xb9\xf0\x9f\x0b\xe0\xa0\xec\x67\x1d\x7e\xca\x38\x08\x9e\x81\x29\xab\x9e\x6f\x36\x23\x93\x7e\x8a\x69\xcd\xf6\xc5\x71\xe3\x63\x17\x61\xf7\xbb\xaa\xbb\x1b\xce\x58\x2a\x0a\x61\x99\x45\xd2\xfd\x5a\xb3\xb2\x64\x56\xa4\x90\x33\xc9\x0b\xca\x6d\x94\xea\xaa\xaa\x08\x9f\xa0\x1c\x06\xb9\x41\xbc\x36\xc7\x1f\x08\x1e\x31\xd2\x7c\x20\xf8\xf7\xf2\x12\xd3\xd8\x6c\xef\x0b\xaa\x8b\x51\xbc\x0f\x86\x1a\xc1\x60\xf3\x8a\x15\x55\xce\x5e\x75\x63\xe1\x03\x59\xff\xf9\x69\x6f\xda\x7f\xfa\x82\x7c\x09\x56\x87\x92\x87\xea\x3c\xb6\xc6\x30\xf2\xdf\x00\x00\x00\xff\xff\xc3\x60\x6e\x88\x2c\x2c\x00\x00")

func appcatalogAppscodeCom_appbindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_appcatalogAppscodeCom_appbindingsYaml,
		"appcatalog.appscode.com_appbindings.yaml",
	)
}

func appcatalogAppscodeCom_appbindingsYaml() (*asset, error) {
	bytes, err := appcatalogAppscodeCom_appbindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appcatalog.appscode.com_appbindings.yaml", size: 11308, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"appcatalog.appscode.com_appbindings.yaml": appcatalogAppscodeCom_appbindingsYaml,
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
	"appcatalog.appscode.com_appbindings.yaml": {appcatalogAppscodeCom_appbindingsYaml, map[string]*bintree{}},
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
