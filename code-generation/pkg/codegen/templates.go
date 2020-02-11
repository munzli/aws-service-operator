package codegen

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

var _aws_service_operator_yaml_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xdf\x8f\xe2\x36\x10\x7e\xcf\x5f\x31\xda\x3d\xa9\x52\xd5\x64\x0f\xa9\x0f\xa7\x08\x56\xa2\x40\xaf\x68\xb7\xec\x0a\xb8\xbe\xde\x0d\xce\x10\x5c\x1c\xdb\xb2\x1d\x38\x0e\xf1\xbf\x57\x71\x7e\xf0\x9b\x3d\xb5\xd7\x3c\x65\x26\xdf\xcc\xf7\xcd\x78\x3c\xb9\x87\xc7\xf2\x81\xfe\x0b\x8c\x5e\xa6\x30\xe8\x0f\xa7\x30\xfd\x63\x38\x81\xdf\x87\xcf\x03\x68\x37\x4f\x70\x0f\xd3\x05\xb7\x30\xe7\x82\x80\x5b\xc0\xdc\xa9\x94\x24\x19\x74\x94\xc0\x8a\x23\x7c\xc1\xb5\x0d\x95\x2e\x3c\xca\x40\xfd\xed\x4b\x70\x0f\xc3\x39\x6c\x54\xfe\x53\x02\x82\x2f\x09\xdc\x82\x80\x2d\x50\xa6\x04\x28\x37\x6e\xc1\x65\x0a\x38\x53\xb9\x03\xd7\x10\x64\xb8\x24\xa0\x84\x3b\x0b\x4e\xf9\x88\xc8\x51\xa6\x45\x70\x5f\x09\x90\xde\xa9\x97\xe9\x03\x53\x09\xa5\x24\x1f\xd0\x5a\x72\x16\x12\x6e\x88\x39\x65\x36\x51\x10\xa0\xe6\x7f\x91\xb1\x5c\xc9\x18\x56\xad\x60\xc9\x65\x12\xc3\x33\xb7\x2e\xe0\x8e\x32\x1b\x07\x21\x94\xbe\x11\x66\x64\x35\x32\x0a\x00\x4e\x82\x00\x32\x72\x98\xa0\xc3\x38\x00\x00\x90\x98\x51\x0c\x45\xa9\x96\xcc\x8a\x33\x6a\x4a\x0e\xb6\x5b\xe3\xab\x7a\xc7\x65\x42\x5f\x7f\x81\x77\x24\x28\x23\xe9\x20\xee\x40\x34\x2c\x18\x77\xbb\x20\x3c\x22\x40\xcd\xe9\xab\x23\x59\x58\x36\x5a\x7e\xb0\x11\x57\x0f\xab\xd6\x8c\x1c\x16\xd4\xa5\xba\x5e\x6e\x9d\xca\xc6\x64\x55\x6e\x18\xf5\x69\xce\x25\x77\x5c\xc9\x2b\xda\xb6\xdb\x9a\x38\x9a\x68\x62\x51\x1d\x18\xbd\x8a\xdc\xa0\xd8\xed\xa2\x53\xe5\x11\xae\x6d\x00\x60\x35\xb1\x32\x51\x6a\x54\xae\x63\xb8\x82\x2b\x89\x6c\x09\xad\x45\x9e\xb2\x3e\x71\x99\xec\x76\x15\x44\x70\xeb\x9e\x6e\xc0\xfc\x99\x94\x50\xed\x45\x7e\x47\x15\x15\xde\x2e\x94\x71\xa3\x43\x3d\xdb\x6d\x08\x17\x0e\xa2\x38\x85\x2b\x29\x27\x45\x0e\x5f\x53\x93\x36\xac\x05\x44\x45\xee\xc6\x5d\xe4\xa6\x83\xc2\x2c\x97\x69\x2e\xd0\xdc\xd0\x7b\x10\x6f\x99\xd2\xb7\x0e\x68\x52\x7c\xaf\xb0\xab\xfd\x0c\xa2\xd0\x0b\x6c\x05\xdb\x6d\xc9\x5c\x0f\x6d\x4f\xe4\xd6\x91\x19\x2b\x71\x3a\xb6\x66\x86\x2c\xc2\xdc\x2d\x94\xe1\xdf\xb0\x18\x95\xf3\xd1\xfa\xee\xa9\x06\x30\xb9\x28\xbb\xeb\x67\xf7\x63\x31\x1b\x55\xb3\x43\xb8\xbb\xf3\x2f\xa6\x2a\xa1\xf1\x5b\x62\x86\x9c\xad\x2c\xad\x92\xfa\x95\x29\x39\xe7\x69\x86\xda\x36\x48\x4f\x58\x9b\xb4\x22\x59\xc5\xad\xc8\xcc\x9a\x84\x29\xb9\xea\x4d\xd4\xc3\x12\xc2\x1a\x1d\x5b\xd4\x89\x0d\xa1\xa3\xca\x48\x48\x50\x63\xe4\x3a\xd9\x7f\xd1\x55\xc8\x85\x5a\x2e\x5d\xc6\xcb\xd5\x31\x7f\x23\x6b\x77\xd2\xdc\xc8\x1f\xa7\xfb\x82\xbc\xab\xb7\xf1\x4c\xde\xdd\xcf\x77\xe7\x42\x0a\x67\x33\x3b\x93\x32\x57\x97\x31\x95\x4b\xf7\x9f\xb6\x5e\xb3\x0f\x8a\xfd\x79\x05\x73\x69\x66\x7f\xe3\x32\xe1\x32\xfd\x9f\x47\x57\x09\x1a\xd3\xbc\x44\xd6\x1d\xbd\xc1\x12\xec\x17\xda\xf1\xf5\x7a\x83\xc7\xe6\xb3\xbf\x89\xb9\xea\x96\x5c\xe9\xf1\x8f\x6b\x63\x9f\xb4\x50\x9b\x62\x83\x9c\xb4\x0f\xb5\xb6\xff\xae\x53\x6f\xb3\x1f\xfe\x20\x0c\x69\xc1\x19\xda\x18\x5a\xde\xf6\x3f\x67\x74\x54\xef\xe0\x63\x62\xdf\x7c\x29\x95\xf3\xad\xb6\x7b\x27\x00\xc7\x2c\xc2\x0c\xbf\x29\x89\x6b\x1b\x31\x95\x3d\x14\x47\x76\x43\xa5\xff\x9d\xe0\x8c\xc4\x51\x1a\xd4\xfa\x66\xcc\x5e\xb9\xb7\x8e\x4e\x66\x74\xbb\x2d\xc5\xc3\x94\x74\xc8\x25\x99\x03\xd2\xf0\xad\x7e\x56\x05\x66\x98\x96\xa8\x0a\x54\x63\x1e\x2e\x05\xc6\xab\xf7\xd1\xfb\xa8\x15\xfa\x95\xff\xeb\x69\x9a\xd7\x5c\x88\x57\x25\x38\xdb\xc4\xd0\x15\x6b\xdc\xd8\xc3\x16\x98\xf4\xa8\x25\xf5\xca\x20\x73\xe2\x0c\x43\x56\x8e\x76\x58\x14\xd0\x69\xf7\x9e\x3f\x4d\xa6\x83\xf1\xe7\x51\xf7\xcf\xc1\xe3\x19\xd6\x50\xca\x95\xec\xb4\xc7\x83\x8f\xc3\x97\xd1\xf9\x77\x2c\x9b\x18\xf2\xa4\xd3\xee\xf6\x7a\x2f\x9f\x46\xd3\xcf\xc3\xfe\x39\x6e\xf9\xc1\x86\xcd\x84\x75\xda\x4f\x1f\x26\x9e\x70\xf2\xda\xed\x0d\x1e\x83\x7f\x02\x00\x00\xff\xff\xad\x8d\x07\x5d\x88\x0a\x00\x00")

func aws_service_operator_yaml_templ() ([]byte, error) {
	return bindata_read(
		_aws_service_operator_yaml_templ,
		"aws-service-operator.yaml.templ",
	)
}

var _base_go_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x3f\xcf\x9b\x30\x10\x87\xe7\xf8\x53\x9c\xa2\x57\x15\x44\x60\x76\xd4\x4e\xed\xd2\xa1\xa9\xda\x77\xe8\x50\x75\x30\xce\xc5\x41\x09\x36\xb5\x8f\x86\x0a\xf9\xbb\x57\x36\xd0\x90\xe1\x15\xc9\x74\xf1\xff\xdf\xf3\xe4\x68\x85\x3c\x0b\x85\x50\x09\x87\x8c\xd5\x4d\x6b\x2c\x41\xc2\x00\xb6\xd2\x68\xc2\x9e\xb6\xe1\xb7\xaa\xe9\xd4\x55\x5c\x9a\xa6\x10\x57\x77\x11\x95\x0b\x35\x77\x68\xff\xd4\x12\x73\xd3\xa2\x15\x64\x6c\xd1\x9e\x55\x21\x8d\x3e\xd6\xea\xe9\x63\xbf\x3b\xec\xb0\x11\x5a\x28\xb4\xe1\xf0\x30\xe4\x60\x85\x56\x08\x2f\xb5\x3e\x60\x9f\xc1\x0b\x5e\xb0\x41\x4d\x50\x7e\x00\xfe\x99\xb0\x71\xde\x3f\xfb\xca\x3c\x70\xc5\x30\xcc\xf7\xf1\xd7\x16\x25\xff\x8e\xce\x74\x56\x22\xdf\x8b\x06\xbd\x9f\x23\xa0\x3e\x78\xcf\x52\xc6\xe8\x6f\x3b\x5a\x02\x47\xb6\x93\x04\x03\x03\x18\x59\xa7\xc2\x3f\xc6\xc2\x00\x22\xcb\x97\x91\x05\x76\x4b\x32\xfe\x6d\xb1\xf4\x0c\xe5\x4a\x5c\xd8\xad\x6c\xe0\x5f\x27\xf2\x3b\x2e\xcf\xd8\xb1\xd3\x12\xf6\x78\x4d\xde\xc0\xc9\x1e\xe7\xc9\x58\x0a\xbb\xa8\x28\xb8\xb1\x48\x9d\xd5\xf0\x2e\x4c\x84\xf1\x7c\x7d\x39\xd5\x2c\xce\x2d\xaf\x2e\xef\x46\xe3\xfa\xc3\x82\x56\x15\x95\x6b\x1b\xf8\x1e\xaf\xb3\xa5\x64\xca\x78\x97\x28\xbd\x45\x1a\xf5\x01\xdc\x14\x26\xd5\xc8\x9e\xc2\x0f\x41\xf2\x94\x48\xea\x61\xfa\x84\x82\xc9\x50\x33\xd0\xa2\x41\xd7\x0a\x19\x9b\xa8\xd6\x2a\x8d\xa6\x1e\x66\xac\x8f\x10\xfa\x3c\xfe\x3b\x73\x78\xf7\x73\xbb\xd6\xca\xbf\xe2\x2b\x1b\x65\xa0\xe2\x6b\x0e\x5e\x49\x58\xfa\x4f\xb0\x48\x9c\x46\xda\xcd\x0d\x9e\x6d\xde\xe7\x92\x7a\xfe\xc9\x68\x4c\x52\xe6\xff\x05\x00\x00\xff\xff\x59\xfb\x37\x1f\x4a\x04\x00\x00")

func base_go_templ() ([]byte, error) {
	return bindata_read(
		_base_go_templ,
		"base.go.templ",
	)
}

var _cft_go_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\x9b\xc8\x13\x7f\x0d\x9f\x62\xfe\x56\xd5\x3f\x44\x0e\xbe\xbe\xb5\x2e\x27\xf5\x9a\x3e\x44\xad\x92\x2a\x76\xdb\x17\x51\xd4\xae\x61\x8c\xb7\xc1\x2c\xc7\x2e\xf1\xf9\x2c\xbe\xfb\x69\x1f\x30\x0b\x36\x76\xdc\xf6\xda\xeb\xc9\xa8\x95\x61\x99\xf9\xcd\xe3\xce\xec\x84\xc1\x00\x7e\xd3\x17\x9c\x5f\xc1\xe5\xd5\x18\x9e\x9f\x5f\x8c\x61\xfc\xea\x62\x04\x2f\x2e\xde\x3c\x87\x5f\xd7\x97\x3b\x18\xc0\x78\x46\x39\x4c\x69\x82\x40\x39\x90\x42\xb0\x18\x53\xcc\x89\xc0\x08\xee\x29\x81\x4f\x64\xc1\x4f\x59\x26\x57\x58\x7e\x1a\xb2\x08\x63\x4c\x21\xcb\x59\x88\x9c\x7f\x92\x00\x17\x53\x58\xb2\xe2\xff\x11\x24\xf4\x0e\x41\xcc\x10\xc2\x19\x49\x63\x04\x92\x2e\xc5\x8c\xa6\x31\x90\x09\x2b\x04\x88\xb5\xa0\x39\xb9\x43\xc0\x88\x0a\x0e\x82\x29\x8e\x40\xe0\x3c\x4b\x24\x9a\xd6\x24\x55\xab\xd9\x5d\x3c\x30\x12\x07\x84\x73\x14\x1c\x22\x9a\x63\x28\x58\xbe\x0c\x5c\x37\x23\xe1\x1d\x89\x11\x56\xab\x60\x94\x61\x18\x5c\x23\x67\x45\x1e\x62\x70\x49\xe6\x58\x96\xae\x4b\xe7\x19\xcb\x05\x78\xae\xd3\xc3\x3c\x67\x39\xef\xb9\x4e\x2f\xa6\x62\x56\x4c\x82\x90\xcd\x07\x64\xc1\xe5\xff\x53\x1e\xdd\x9d\xc6\x4c\xde\xee\x26\xe0\x98\xdf\xd3\x10\x07\x61\xc2\x8a\x68\xca\xf2\x39\x11\x94\xa5\x3d\xd7\x21\x0b\xfe\xfe\x09\x49\xb2\x19\x79\x02\x2d\xfe\x84\x4c\x0c\x86\x66\x5e\xfb\x72\x20\xad\x23\x19\xe5\x83\xf6\x9b\x40\x8a\xbd\x37\x78\x9b\x1a\xed\x46\x0c\x59\x3a\xa5\xf1\xa1\x5c\x33\x4c\x32\x94\xfe\xf1\x5d\x19\x84\x4b\x5c\x40\x95\x06\x1c\x08\xa4\xb8\x00\x36\xf9\x8c\xa1\x70\xa7\x45\x1a\xca\xf7\x9e\x16\x04\xfa\x27\x78\xa6\x7e\xfa\x5d\xb1\x80\x13\xcb\x45\x41\x45\xf4\x9a\xa6\x51\x59\xf6\x41\xb0\x8c\x86\x4f\xaf\x2f\x81\x8b\x9c\xa6\xb1\x0f\x27\xcf\x1a\x1e\x86\x95\xeb\xe4\x28\x8a\x3c\x85\xc7\xcd\x37\x2b\xd7\x71\x9a\x68\xc3\x2e\x15\xfa\xae\xe3\x68\x65\x87\x8e\xbc\xf4\x7d\xdf\x05\x80\xb5\x02\x43\xd0\x57\xf5\xdc\x77\x9d\xd2\x2d\x95\x4b\x5a\x1a\x45\x38\xa5\x29\x72\x95\xa6\x5d\x36\x87\x53\xc1\x5d\xb1\xcc\xb0\xcd\xcc\x45\x5e\x84\x42\x5a\x65\x9c\x58\x5d\x0d\x67\xba\x2d\xcb\x76\xf9\xd0\xad\x75\x36\x50\xda\x95\x46\xf9\x91\x20\xe1\x9d\xd4\x0a\xb4\x1b\xb5\xde\xa9\x5c\x60\x53\x75\xcf\x25\x05\x4c\x08\xc7\x08\x98\xde\x7d\xad\x8d\xaf\x54\x52\xd1\xf7\x78\x3b\x40\x7e\x2d\xc1\xf3\x8d\x68\x2b\x68\x26\xbb\x82\x9a\x88\x07\x95\xa9\x49\xc1\x05\xe6\x72\xb1\x0f\xbd\x0e\x4f\xf6\xfa\xc0\x5b\x06\x07\x9a\x63\xeb\x32\xcf\x48\x88\xbe\x31\xfd\x25\x8a\xab\x42\x64\x85\xe0\xc6\x76\xcb\x5c\x66\x5e\x4c\x73\x36\x57\xcb\xe7\xc8\xc3\x9c\x4e\x50\x29\xca\x21\x24\x49\xd2\x6d\x72\x8d\xec\xf9\xe0\xcd\x49\x76\xa3\x0d\xbf\xd5\x3f\x7d\x50\x35\xc7\x97\x7e\xa8\x04\x0d\xcf\x60\x83\x6e\x55\xba\x0e\x47\xae\x5e\xae\xdd\xf2\xf4\xc3\x68\x84\x9c\x53\x96\xba\x0e\xbf\x0f\xe5\xbb\x66\xd1\x09\xe4\x16\x94\x6c\xbe\xeb\x3a\xca\x9a\x8b\xb4\x12\xd1\xa2\x6c\x1a\xa5\xc8\xe4\xb6\x59\x07\x43\x26\x3d\x59\xc8\xe8\x48\x75\x3c\x3b\x4c\xbe\xaf\x76\x40\x65\x80\x32\x49\xe9\x79\x1f\xb6\x60\xbd\xc7\x96\x12\xbe\xeb\xd0\xa9\xa2\xfd\xdf\x19\xa4\x34\x91\x2e\xa8\x72\x21\xa5\x89\x82\x91\xb8\x8e\x2c\x35\x4c\x00\x2f\x72\x04\x3a\xd5\x1d\x82\x72\xc0\x7b\xd9\x62\x18\xe7\x74\x92\xa0\xc2\x4a\x30\xf5\xb4\x0e\x5a\x3b\xee\x4b\xe8\x27\xdb\x80\x59\xce\x95\x73\x7a\x29\xd3\x61\xae\x02\x8f\x11\x2c\xa8\x98\x81\x98\x11\x61\x12\x40\xee\x80\x9e\xaf\x4d\x9c\xb2\x1c\x3e\xf6\x65\x4e\x48\x0b\x73\xd5\xc1\x1a\x22\x6f\x7e\xb9\x0d\xaa\x54\x92\x72\x4d\x50\x6f\x4e\x58\x21\xcc\x8b\xd7\xb8\xbc\x85\x33\xb0\x56\xde\x93\xa4\x40\x2d\xc0\xe8\x69\xd8\xb4\x13\x4c\x6d\xc9\x91\x08\xed\x48\x58\xd0\x24\x81\x50\x2d\x58\x99\x6a\x14\x47\xe0\x45\x96\x25\x14\x23\xc8\x48\x4e\xe6\xbc\x3b\x3b\x2d\x4c\x99\x9e\x5a\x2a\x9c\xb4\x72\xc3\xa2\xba\xb2\x42\x5c\x67\xee\x57\x66\x66\x38\x55\x8d\x5d\x1a\x33\x3c\x5b\xd7\x81\x97\x28\x94\xba\x2f\x2a\x9e\xb1\x21\x5a\xd7\x85\xc3\x6a\x41\xf5\x38\x5e\x66\x58\x96\xc1\x76\xec\x8e\x82\xf1\x60\x5e\x53\x55\xf6\xed\x36\xcb\xa3\x87\x6f\x35\xa7\x12\xf8\xee\xfa\xcd\xd0\xa6\xac\xfd\xa8\xc8\x2e\x99\xa0\x53\x1a\x2a\x89\x4f\xaf\x2f\xf9\x10\x6e\x6e\x4f\x4c\x39\x71\x1d\xc7\x69\xc8\xa8\x1a\x83\xe2\x2c\xfb\x55\x32\x6a\xaf\xaa\xa6\x60\x85\x46\xab\xff\x56\x26\x97\xd7\xbb\xb6\x88\xba\x6a\xb0\x5f\x63\xbd\xc7\x5c\xe6\xc5\x5e\x38\x43\xb7\x0d\xb1\x45\xe2\xbb\x4e\x5a\xb9\xbe\x13\x76\x1d\x9c\x2e\x15\x4d\xe4\x9c\xb0\xee\x35\x9d\x60\x56\x3f\x52\x70\x9b\x6d\xca\x77\x57\xab\x53\x53\x1f\x1e\x09\x96\x49\x28\x2d\xf2\x77\x16\x2d\x83\x51\x38\xc3\x39\x09\xde\xe6\xb2\x75\x0a\x8a\xbc\x2c\x65\x23\x3f\x95\xf5\x6d\xcb\xaa\xc1\x49\x91\xcb\x03\xb7\x84\x6a\xd0\xd4\xac\x26\x33\xe8\x5f\x18\x69\x5e\x29\x3b\x78\x8d\xcb\xb2\x5c\xad\x0c\xbf\x7e\x94\x94\x12\xa9\x57\xaf\x57\x69\x55\x96\x3d\x8d\x88\x09\xc7\x87\xc3\xc8\x7f\x65\x19\x5c\x4d\x3e\x07\xab\xd5\xa3\xc6\x7e\x31\x00\x23\x75\xa2\x51\x7c\x2a\x02\xf5\xa3\x62\xed\xd5\x82\xd3\x7d\xea\xab\x8a\xb9\xee\x34\x55\x90\x6a\xf3\xbd\x3d\x3a\xf7\xd7\x3c\xe7\x44\x90\xd5\xd5\xe4\xf3\x50\xa5\xc5\xa3\xe6\x99\x53\x9f\xb1\x86\x50\x57\x9c\x57\x9a\x6d\xb8\xe6\x97\x85\xcc\x2f\x77\x35\x33\xab\x2d\xca\x6d\x05\xb0\x43\xb7\xce\x94\xdb\x0c\x93\xa2\xef\xf5\xad\x73\x93\xdc\xca\x74\xba\xdc\x65\xbb\xf2\x9b\xef\xb7\xdc\x6c\x87\x5a\x25\x92\x87\x7f\xa8\xb4\x55\x01\x84\x9e\x5d\x07\x7a\x7e\x3b\x34\xbb\x74\x56\x18\x7b\x15\xde\xf0\xbc\x2e\x1a\x7e\x43\xb7\x07\x65\xb9\xc9\x47\xa7\x2d\x7b\x67\x4e\x5b\x4c\x07\x25\xf1\xfe\xac\x3d\x28\x4d\x7f\x70\x5e\x7e\xe3\x98\x6e\xb8\x61\x5b\xd6\xa9\xdb\xfa\xce\x75\xd4\x81\x05\x05\xe6\xaa\x6b\xde\xdc\xb6\x8f\x22\x6f\xab\xf7\xf2\x40\x6c\x11\x9f\x01\xc9\x32\x4c\x23\xaf\x5e\xeb\x83\x9d\xb6\xfe\xc3\xc9\xeb\xb6\xb2\x8f\x23\xb5\xfa\xc6\x3e\xda\xf0\x07\x35\x87\xbd\x8a\xed\xa8\x16\xdd\x75\xe2\x10\x50\xbf\x2b\xd8\xd6\x19\x29\x18\xa1\x58\xc7\x96\x5b\x50\x7e\xeb\x18\x32\x26\xf1\x66\x6e\x8e\x49\xfc\xa5\xe7\x90\xfd\x78\x5f\x76\x10\xe9\xc6\xfd\x92\x93\x48\x37\xda\x43\x8e\x22\xae\x23\x48\xdc\xb1\xa1\xc6\x44\xcd\x96\x8a\x60\x1d\x49\xf9\xd4\xdc\x3e\x63\x12\xfb\xbb\xa9\x6a\x7f\x76\x10\xda\xae\xe9\x20\x69\x5a\xec\x6f\xa6\xc8\x98\xc4\x5c\xd1\xfa\xae\x0b\x76\x01\x03\x3d\x6e\xda\xf3\x4c\x6b\xd6\xd4\x45\xcf\xcc\x52\xef\xb2\xa8\x39\x4b\x15\x6a\x41\x0d\x4f\xf8\x27\xe5\x82\xa6\xb1\x1e\xaa\xba\x67\x27\x0b\xc3\xd3\xec\xd1\xae\x3f\xbe\x74\x8f\x57\x16\xd0\xcf\x33\x5e\x19\x8b\x0f\x99\xaa\x0e\x66\x79\xe0\x30\x65\xf9\xef\x38\x4c\x1d\x87\xa9\xe3\x30\xf5\x6f\x18\xa6\xcc\x66\xff\x0e\x47\xd5\xe3\x04\xf5\x4f\x4c\x50\xf0\x5f\x9a\xa0\xbe\x5f\x32\x1e\xe7\xa6\xe3\xdc\x74\x9c\x9b\x8e\x73\xd3\x4f\x36\x37\xd9\xb3\xcc\xae\xb9\xe9\x1c\x13\x6c\xcc\x4d\x91\x5a\xa8\xbf\x41\x75\x8f\x4b\x16\xab\xe7\x83\xf7\x4d\x27\x9c\x7d\x9f\x36\xd7\x92\xf5\x7c\x50\x6e\xf8\xc7\xfe\xe6\x6c\x0f\x08\xd2\x5b\x1f\x6d\x47\xd9\x56\xec\x72\xd4\x07\x42\xc5\xbb\x54\xd0\x44\x91\x6a\xae\xe8\x50\x97\x6d\x05\xf9\xee\xce\xfb\xea\xef\xc2\x60\xf9\x6f\x9b\x49\xcf\xd8\x3c\x93\xbf\x9d\x0e\xfd\x3b\x00\x00\xff\xff\xd3\x02\x8f\xa7\xac\x24\x00\x00")

func cft_go_templ() ([]byte, error) {
	return bindata_read(
		_cft_go_templ,
		"cft.go.templ",
	)
}

var _operator_go_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5d\x73\xdb\x36\xb3\xbe\x26\x7f\x05\xca\xd3\xa6\x62\xaa\x90\xcd\x5d\xc7\x73\x7c\x3a\xae\xad\xb6\x9e\x38\xb6\x8f\xed\x34\x17\x99\x4c\x0a\x93\x2b\x89\x31\x09\xb0\x00\x64\xc5\x55\xf5\xdf\xcf\x2c\x16\xa0\x48\x99\x92\xa5\x36\xd3\xf3\xce\xfb\xd6\x37\xe2\x07\xb0\xd8\xdd\xe7\xd9\x0f\x80\x4e\x53\xf6\x3f\xf4\xc7\x4e\x2e\xd8\xf9\xc5\x0d\x1b\x9d\x9c\xde\xb0\x9b\x9f\x4f\xaf\xd9\x8f\xa7\x67\x23\xf6\xdf\xcd\x5f\x98\xa6\xec\x66\x5a\x68\x36\x2e\x4a\x60\x85\x66\x7c\x66\xe4\x04\x04\x28\x6e\x20\x67\xf7\x05\x67\xbf\xf2\xb9\x7e\x21\x6b\x7c\x22\x15\xf3\xef\x7e\xc5\x99\xa7\x63\xf6\x20\x67\x5f\xe7\xac\x2c\xee\x80\x99\x29\xb0\x6c\xca\xc5\x04\x18\x17\x0f\x66\x5a\x88\x09\xe3\xb7\x72\x66\x98\x69\x56\xa8\xf8\x1d\x30\xc8\x0b\xa3\x99\x91\x76\x46\x62\xa0\xaa\x4b\x94\x46\x2a\x08\xfb\xb4\xbe\x9b\xa4\x99\xcc\x61\x02\x22\xe5\x5a\x83\xd1\x2c\x2f\x14\x64\x46\xaa\x87\x24\x0c\x6b\x9e\xdd\xf1\x09\xb0\xc5\x22\xb9\xae\x21\x4b\xae\x40\xcb\x99\xca\x20\x39\xe7\x15\x2c\x97\x61\x58\x54\xb5\x54\x86\x0d\xc2\x20\xca\xa4\x30\xf0\xc9\x44\x61\xb0\x58\xbc\x60\xc5\x98\xd1\x94\x53\x7d\x3c\xd3\x46\x56\xc5\xef\x90\x2f\x97\x61\xc0\x2a\x30\xfc\xfe\x25\x8b\xee\xbe\xd3\x49\x21\x53\x5e\x17\x15\xcf\xa6\x85\x00\xf5\x90\xa2\x3a\xbc\x2e\x74\x8a\x83\xd2\xfb\x97\x51\x18\x44\x93\xc2\x4c\x67\xb7\x49\x26\xab\x94\xcf\x75\xc9\x6f\x35\xfe\xbe\xd0\xa0\xee\x8b\x0c\x1a\x8f\xd9\xb9\x53\x28\x6b\x50\x1a\xa7\x29\x18\x97\x90\x79\x75\x40\xe0\xe2\x7b\x4a\xcb\xa4\x18\x17\x93\x28\x64\x6c\xaf\x69\xbf\xcd\x60\x06\x15\x17\x7c\x02\x0a\x27\x77\xdc\xf1\xbf\xf8\x72\xb9\x0c\x19\xcb\xa4\x82\xae\x1f\x52\x7c\x64\xad\xee\xae\x58\x70\x91\xc9\x12\x45\xa6\xda\xa8\x8c\x6b\xb0\x23\xb4\x51\x85\x98\xa0\xb1\x7c\xae\xb3\xb2\x00\x61\xf6\x53\x94\xe6\xb8\x1f\x0d\x26\xbd\x07\xa5\x0b\x29\x20\x4f\xcd\x43\x0d\x79\xba\x3e\x2b\xe1\x73\x9d\xde\xbf\xe4\x65\x3d\xe5\x7b\x63\x63\xdd\xd2\x81\x63\xbf\xf9\xfe\x06\xd7\x75\x3e\x23\xcd\x5f\x4c\x64\x6a\xa4\x2c\x75\x9a\xf1\x6c\x0a\x51\x68\x1d\xf2\x8b\x53\x73\x3f\x97\x58\xf2\x3d\x65\x76\x07\x50\xcf\x6e\x6e\x0a\x29\x74\x72\x49\x11\x63\x11\x8e\x7c\xd8\x6c\x1a\xd3\xf1\x46\x1c\x62\x70\x5e\xf8\x04\xa0\xa0\x56\xa0\x11\x18\xc6\x19\xc6\x96\x92\x65\x09\x8a\xc9\xdb\x8f\x90\x19\x36\x96\xcd\xa5\x36\x52\x01\xcb\xec\x1a\x4c\xb9\x08\xd5\x21\x42\xb8\x12\xa7\x8d\x9a\x65\x86\x2d\xc2\x80\x58\xcd\xe8\x8f\x6e\x92\x63\xfb\x13\x06\x46\xd6\x45\x76\x74\x75\x6e\xdf\x11\xbf\xc2\xc0\x02\xf7\x9a\xf8\xcc\x9e\xb7\xd9\x4d\x6c\x76\xaf\xc2\xa5\x35\xe0\x1c\xe6\xcd\xa2\x99\x02\x6e\xa0\xad\x3d\xaa\x3d\xe7\x26\xb3\x49\x6b\xab\xfe\x6e\x72\x1e\x8e\x67\x22\x6b\x4b\x1d\x38\x03\x3a\xaa\x0f\xd9\x8e\x5a\xc6\xec\x79\xa3\xde\x62\x0d\x4a\x1f\x9b\x64\x71\x66\x54\xc9\x0e\x0e\x49\x70\x72\x0e\x73\xb7\xf0\xd0\xaf\x7c\xf4\xf6\xfa\xd8\x87\xce\x90\xbd\xfc\x36\x5e\xf9\x6f\xc8\x3e\x34\x53\x51\x4e\x72\x05\x93\x42\x1b\x50\x83\x68\x43\x2a\x8d\xe2\xae\xa3\x93\xa3\x3c\x1f\xac\xc4\x75\xec\xf9\x99\x8b\xbc\x04\xf5\xe3\x4c\x64\x03\xab\xf3\x9b\x3a\xe7\x06\x54\x1c\x77\x93\x9d\x02\x33\x53\x82\x3d\xf3\x06\x2f\xc2\xc0\xa1\x7f\xd0\x81\x7f\x18\x06\x1b\x1c\xd1\x18\x74\xc0\x1a\x5d\xdc\x60\x17\xc2\x1d\xa5\x0f\x3a\x28\x0c\xc3\x60\xe9\x48\x71\x6d\xb8\x32\x6f\x11\x77\x42\x1f\xb4\x65\x42\x21\xb4\xe1\x02\xc1\x96\x63\x76\x41\x6c\xb8\xee\x67\x03\x17\x39\xe3\x99\xd1\x4c\xda\xc2\x55\x11\x2d\x06\xd9\x0a\xce\xb8\xb5\xca\x20\x33\x9f\x98\x2b\x48\xc8\x10\xfc\x1d\x32\xc1\x2b\xd0\x35\xcf\xc0\x51\x3b\x46\x06\xf8\x25\x9c\x53\x35\xe2\x66\xf3\x48\x83\xd0\xe8\x1e\x84\x69\xf9\x5c\xa3\x1f\x8f\xf2\x1c\xaf\xad\x23\xb3\x44\x8a\xa3\x3c\x47\xcf\x10\x12\xf4\x06\x1f\xd3\x3d\xbe\x39\x81\x12\xda\x6f\xe8\xde\xfa\x28\x0c\x30\xd3\xe0\xc2\x4d\xc6\x41\xba\x3d\xa6\xca\x65\x39\x53\xbc\x5c\x2e\xa3\x96\x2d\x43\xb6\x6e\xc1\x90\x65\x49\x0f\x45\x93\xab\xd1\xf5\x0d\xdd\x0d\x90\x29\xb8\x56\x42\xde\x7a\xd6\xca\x97\x89\x5f\xf5\x55\x81\x08\x2f\x96\x43\x96\x99\x4f\xc9\x89\x14\x80\xd3\x96\x61\xd8\x4b\x95\x34\x65\x6d\x2a\xb2\x79\x51\x96\xcc\x70\xd7\xae\x54\xa0\x35\x9f\x20\xea\x4a\x56\xf6\x89\x25\x8a\x45\xb5\x56\x32\x03\xad\x5b\xa8\xb6\x05\x6d\x88\xf6\x4a\x4f\xd6\x82\xfc\x35\x2d\xf1\x83\xcc\x1f\x62\x06\x4a\x51\x7c\x97\x72\x32\x21\xd7\x3a\x01\x67\xf6\x41\x18\xdc\x73\x65\x7d\xf8\x98\x15\x61\x50\x8c\x51\x7e\x62\x35\xe0\xb7\x25\xa0\xa0\x00\x87\xb1\x43\xfb\xc2\xe3\x81\x91\xeb\xde\x90\x00\x7a\x7d\xee\xef\xc3\x60\xc9\xa0\xd4\x34\x9f\x6a\xd5\x35\xa6\x0a\x9b\x1b\x9a\x9a\x8d\x58\xff\x28\x15\x59\xe6\xcc\x25\xac\xec\x65\x1c\x86\x41\x43\x52\x3d\x44\xd3\xac\x39\x5e\x5c\x83\x17\x91\x83\xb2\x89\x17\xf3\xea\x3b\xdd\x68\x13\x27\x67\x85\x36\x03\x6a\xbf\xec\xf5\x45\x6d\xeb\xd1\x62\x19\x87\x01\x1a\x8d\xa2\xbf\x38\x64\xa2\x28\xad\xc6\xce\x79\xc9\xdb\xc2\x4c\x47\xe8\xd0\x01\x28\x15\x27\x74\x19\x91\x8b\x27\x60\x0c\xa6\xf2\xcd\x54\x45\xd9\x3e\x13\x81\x52\x61\x10\x60\xd6\xc0\xe8\xff\xb0\xa2\x2e\x5a\xa4\x6c\x47\xdb\x58\x9a\x9c\x1a\xa8\x34\x29\x52\x8c\x9b\xe7\xc9\xb5\xe1\x66\xa6\xf1\x27\xbb\x3b\x3d\x61\x87\xe4\xf3\x4b\xae\x34\xe4\x8e\x03\xef\x22\x7a\x9b\x47\xef\x69\xbe\x07\x4f\xb5\x73\x6e\xf3\xc2\x63\xd7\x79\xeb\x00\x24\x6d\x97\x14\xa5\xc5\xd8\x92\x05\x7d\x14\x45\xec\xd9\xb3\x16\x75\xe8\x11\x2e\xc6\x85\x90\x86\x0a\x3d\x5a\x55\xf1\xfa\x1d\xf1\xea\x3d\xfd\x58\x85\x9c\x82\x27\xd1\xc1\x56\xed\x87\xab\xb1\xa8\xd3\x96\xd1\xf6\xb5\x1b\xef\x31\xb8\x79\xa8\x37\x4c\xe9\x8c\xb0\xb3\x96\xc4\x80\x2d\x63\xc9\xef\xd1\x7b\xf4\x78\x74\x75\x71\x76\xf6\xc3\xd1\xf1\xab\x0f\xc7\x17\xaf\x2f\xcf\x46\x37\x23\xb2\x3d\x90\xb7\x1f\x1b\x86\xe6\x36\xc1\x59\xed\x9a\xca\xb9\x16\x71\xc3\xad\xd6\xc7\x0e\xfa\x75\x5a\x76\xd9\x64\x55\x0f\x7c\xd8\x40\x26\x55\x8e\xd5\x93\x50\x80\xdc\x26\xef\xf1\xc0\x2a\xd6\x82\x66\xe8\xda\xef\xc4\xbe\x47\x3f\xbc\xe5\x4a\x14\x62\x32\x64\xae\xc5\x4e\x6e\xe4\x31\xaf\xa0\x1c\xb8\x36\x3b\xb9\x91\x67\x72\x0e\x6a\xb0\x83\x8f\xe2\xb8\xd7\xb2\xee\xb0\x2b\xe0\x5a\x0a\x32\xd3\xe5\x89\x3d\x10\x38\x19\xa1\xdb\xb7\xf9\x7f\x66\x13\x28\x4d\xfa\x73\x00\xec\x60\xc4\x4e\x83\xda\x96\xfe\x87\x00\x1a\xa0\x8d\x87\xac\x10\x99\x82\x0a\x84\xb9\x92\x65\x79\xcb\xb3\xbb\x63\x39\x13\x66\x13\x1a\x7b\xf8\xa7\x55\x58\xfe\x01\xbd\x05\xfa\xb9\x54\x15\x2f\xff\x9f\x82\x38\xa4\x32\xe1\x0c\x13\x45\x19\x2e\xc3\x56\x3f\xde\xd3\xb1\xda\xbe\x11\xcd\x62\x85\x30\xa0\xc6\x3c\x83\xc5\xd2\xb6\xa5\xb6\x78\xc8\xdb\x8f\xc9\xe0\xf9\xe6\xe6\x2c\x4e\x4e\x00\xea\x63\x59\x3f\x0c\xb0\x4d\x70\x8d\x99\x90\xa6\x7f\x6f\x7a\x94\xe7\x76\x5f\x5a\x8c\x99\xf6\x75\xb4\x6b\x8e\x4d\x2d\x11\xfb\xe3\x8f\xad\x03\x7a\x72\x8f\x6d\x83\xc7\x06\x95\xb6\x3b\xa5\xc4\x73\xcf\x76\xa3\x7e\xef\x10\xdb\x71\x72\x66\xea\x99\x59\x75\x32\x63\x93\x1c\xdb\xbd\x1e\xd5\x0a\x1a\xb4\xce\x18\xb7\x5f\x49\x3a\x6d\x5c\x6f\x6b\x32\xf6\xbd\x89\xdd\x40\xf6\x36\x27\xd4\x24\xb1\xaf\xbf\xd2\x5f\x47\x43\xa6\xed\x7d\xec\x96\x20\xf4\xec\xcd\x32\xec\x5b\xf3\x54\x8c\xe5\x78\x10\xf1\x3c\x87\x7c\xab\x68\x36\x2f\xcc\x94\x69\xd7\xa8\x74\x16\x1b\xba\x5e\x73\xf0\x9c\x9c\xe1\xda\x99\x3c\x8e\xb7\x2d\x79\x5f\xc0\x9c\x71\xc3\xa6\xc6\xd4\xfa\x20\x4d\x33\x29\xb4\x2c\x21\xe1\x73\x9d\xf0\x8a\xff\x2e\x85\x3d\xe4\xc8\x4a\x39\xcb\xc7\x18\x08\x08\x7b\x3a\x95\x15\x7c\xff\x5f\xa9\xd5\x23\xcd\xc1\xf0\xa2\xfc\x9e\x94\xca\x0f\xbf\xd2\xd1\x16\x55\xc2\x20\xf8\x40\x30\xad\x27\x96\x15\xbc\xde\x9c\x55\xd3\xb4\x51\xe0\x90\x45\xc7\x57\xa3\xa3\x9b\xd1\x87\xd3\xf3\x0f\x97\x57\x17\x3f\x5d\x8d\xae\xaf\xa3\x21\x8b\xa2\x0d\xfd\xe7\x2e\x68\x7b\xb0\xad\x7e\x08\xb6\xa6\xf0\xa5\x80\x44\x0c\x69\xb7\x5a\x6a\xbb\xa1\xdd\x70\x22\x63\x23\xa3\x6d\x56\x67\x27\xbd\x29\x76\x69\x8b\x32\x90\x65\x7e\x81\x89\x49\xc0\xfc\xa2\x27\x92\xa5\xb4\xa1\x6c\x07\xed\x1e\xcd\x81\xb0\xd3\x48\xe6\x3e\x49\xc0\x26\x80\xfd\x82\x17\x3d\x4d\x99\x48\xc8\xd8\xfa\xec\xe9\x5c\x42\xb6\xfb\x74\xe2\xce\x5d\x93\x53\x6d\xc1\x3e\x96\x55\x8d\xcd\xdf\x40\x6e\xd0\x64\xc8\xc6\xbc\xd4\x10\x63\x13\xfd\x85\x3b\xa9\xb5\x66\x8c\x7e\x9b\xf1\x12\xa7\x79\x1b\x31\xa9\x2f\x97\x43\xb4\xa9\xfb\x28\xde\x96\x75\xa4\xdc\x29\xed\xbc\xf1\xa4\xce\xee\x9c\xe9\x9f\x25\xf1\x34\x5c\x7c\x3a\x3b\x08\x98\xb3\x9a\x2b\x5e\x69\xf6\xd5\x37\xf7\x76\x3f\x2c\xcb\x1c\xaf\x23\x6b\x33\x45\x97\x90\x68\xd2\xbe\x19\x8a\x42\x76\x97\x1c\xe5\x34\x70\x29\xaa\x59\xf6\xdf\x24\x47\x49\x6f\x8f\xbb\x78\x32\x4b\xbd\xb9\x3c\xf9\xdb\xb3\x14\x7d\x26\x78\x22\x4d\xf9\xa0\x5b\x63\x3a\x52\x77\x87\x74\x45\x27\x4e\x7f\x5b\xb7\x41\xcb\xa1\x35\xbb\x34\x06\x41\x2b\x2c\x4f\x5a\x3b\xc7\x38\xec\xf1\xf9\x5e\xd1\x68\xf7\xa1\x7b\xb4\x01\xbe\x33\xa5\x66\x6e\x03\xcd\x69\x73\xbb\x3d\xb6\xda\x32\x77\x00\xd7\xfb\x6b\x5b\x19\x0a\xb7\x7d\x5a\xb3\xa0\x6f\xdd\x71\xac\x9f\xa2\xd9\xd3\x0c\x8a\x83\xbe\xd3\xd1\xa7\x4f\xd1\xfe\xfc\xa1\x56\x73\xa4\xb5\xeb\x89\xd6\x6a\x97\x94\xfc\x04\x66\x40\x7b\x1a\x77\x98\xf5\x13\x74\xce\xb2\x1e\x13\xe6\xb3\x9d\x64\xb5\x37\x2d\xd4\xeb\xd3\x28\x8c\x0a\x7b\x88\xe5\x67\xb5\x4b\x79\x7b\x90\x13\xde\x06\xa7\x75\xf4\xb4\x61\xc4\x37\x2f\xc3\x70\x95\xe6\x76\x77\x94\x6b\x50\xda\xd2\xff\xac\x7b\x9a\xcc\xe5\x85\xf5\xb8\xa3\xbb\xf3\x21\x3e\xf6\x6c\x48\xf7\x62\xe1\xb0\x69\xa0\x5b\xf7\xd8\xcb\xf8\x5b\x65\x77\x60\x0d\x67\xb7\x25\xb0\x21\x31\x3a\xfe\x87\xd2\x3d\x94\x16\x45\x39\xfc\x4b\xbc\xee\xef\x38\x1d\x5c\xbb\x0c\xa6\xbd\xb4\x8d\x06\xbc\xe8\x9f\xd2\x9c\xfb\x7a\x5e\x50\xcb\xbb\xa9\x01\xd5\xdd\x76\xd3\x16\x8f\x56\x35\x72\x0c\x6c\x2f\xe4\x2b\x3d\xb5\x05\xba\xd3\x2f\x22\x28\xf4\x78\xf0\xd7\x8f\xcc\xdd\x02\x91\x3f\x3a\xb0\x99\x9e\x0e\xc1\xbf\xac\x95\xac\x71\x51\xc2\xed\xc2\x75\x28\xd9\x14\x2a\x9e\x5c\x2a\x59\x83\x32\x05\x68\x66\x3f\xc0\xb9\x92\x60\xe7\x24\x37\x50\xd5\x25\x37\xfe\x9f\x2d\x82\xc5\x82\x9e\xbf\x82\x07\xe4\xbf\x65\xb4\xf7\xd5\x6a\xec\x20\xf2\xe3\xdc\x33\xb0\xdf\x98\xfc\xc0\x13\x6e\xf8\xe2\xe2\xf6\xe3\xc1\x9a\xa3\x8e\xdd\xb7\x44\x72\xe3\x32\x6e\x7d\xa8\xb0\x90\x39\xbd\xbd\xec\x6b\xfb\x09\xda\x6a\xc2\x0e\x59\x47\x33\xff\x4d\xd1\x55\xc8\x3d\xc4\x38\x2f\xbe\x7b\x64\x81\x1d\x10\xbd\x5f\xfb\x5a\xd9\xba\x5e\xfe\x6b\xe7\xd4\x6e\x3c\xee\x47\x71\xb2\x49\x3f\x88\xec\x28\xcf\x0b\xcc\x22\xbc\xf4\xb1\xa6\x7b\x59\xbf\x17\x9d\xb1\x07\xf2\x36\xe0\x22\x6d\x13\x3c\x9f\x5b\xe5\xa0\x4b\x9a\x76\x71\x78\xfc\x89\xe0\x2f\xd5\x86\x7f\xb2\xff\x67\xca\xfe\x5b\x52\xa4\x4b\x8f\x2e\x6a\x76\x69\xd6\x7b\x57\x58\xcd\x7f\xcb\x0b\xf3\x46\x98\xa2\xb4\x22\x48\x5a\x4e\xd5\xa5\xc3\x1e\x9a\xec\x99\xb3\x9d\xdb\xeb\x2c\xd2\x6c\xdb\xb6\x86\xa1\xf3\x5a\xe4\xf8\x1b\x39\xd0\x3a\x36\x23\x16\xf8\xf3\xb4\x5d\x79\xf0\xa8\x01\x73\xdf\x53\xfb\x2b\x36\xee\x26\x56\xdb\x87\x1e\xff\x25\xd7\xf4\xaf\x4f\x58\x5a\xb4\xbf\x3c\x38\x64\xef\xfc\xc7\xcb\xe5\x62\xe1\xd2\x67\xab\x5a\xb5\xbf\xe2\xee\x24\xba\xd9\x3c\xc2\x6f\xab\xd9\xf6\x5c\x87\x45\xa3\x4f\x06\x94\x20\x0f\x45\x34\xf4\x4b\xb5\xb6\xbb\x6a\x15\x31\x3a\x39\x6e\x4f\x72\xeb\x0c\x5a\x1b\xcd\xde\x53\xca\xa8\x2d\xb8\x2d\xc0\xd6\xbe\x68\xb1\x18\x14\x22\x87\x4f\x2d\x05\x2f\xa5\x32\x9a\x7d\x1b\xdb\x0b\x8a\x9f\xc6\x49\x87\x8c\xd7\x35\x88\x7c\xe0\x9f\x0c\xd9\x63\xc5\x3b\xdb\xb9\xee\xe5\xae\xa8\x34\x12\xb7\x8e\x3b\x64\x5e\x0d\x8f\xd7\xd3\xd0\x13\x89\x5f\xf3\x1a\x97\x21\xe7\x55\x78\xf3\x99\xe0\xef\x88\xef\x43\x15\xbb\x8c\x8d\x5f\xcc\x5b\xcb\x15\x06\x2a\x1c\xb7\x9a\x6f\x67\xda\xf2\x8e\x98\xe2\x7b\xea\x2a\xa2\x03\xd6\x3c\xf8\x85\x97\x33\x8b\xec\x7a\x13\xb0\x13\xbf\x1a\xe5\x9f\x64\x55\xbf\x5d\xb1\xff\xb7\x3b\xeb\xd1\x86\x2b\xab\x67\xbb\xb0\x65\x77\xf4\xb6\x92\xa4\x35\xd2\x97\x3f\xd4\xa0\x21\xca\xee\x5d\x51\x27\x7b\xad\xf5\x45\x3b\x64\x2b\x7f\x92\xf9\xe8\x4b\xd9\xff\x05\x00\x00\xff\xff\x2f\x58\x1a\x10\x58\x2d\x00\x00")

func operator_go_templ() ([]byte, error) {
	return bindata_read(
		_operator_go_templ,
		"operator.go.templ",
	)
}

var _template_functions_go_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x8f\xdb\x36\x10\x3d\x8b\xbf\x62\x2a\x04\xa8\x05\x78\x45\xec\xad\x08\x9a\x1e\xda\x6c\x12\x23\xc1\x6e\x91\x35\x90\x63\x43\x4b\x63\x69\x60\x8a\x24\xc8\x91\x1d\x43\xd0\x7f\x2f\x48\x7d\x04\xd8\x22\xdd\x8d\x2e\x94\x46\x9c\xf7\xe6\x3d\xce\x50\x4a\xf8\x63\x7a\xe0\xed\x03\xdc\x3f\xec\xe1\xee\xed\x6e\x0f\xfb\x0f\xbb\x47\x78\xb7\xfb\x74\x07\xbf\xaf\x8f\x90\x12\xf6\x2d\x05\x38\x92\x46\xa0\x00\xaa\x67\xdb\xa0\x41\xaf\x18\x6b\x38\x93\x82\xaf\xea\x12\x6e\xac\x8b\x11\xeb\x6f\x2a\x5b\x63\x83\x06\x9c\xb7\x15\x86\xf0\x35\x02\xec\x8e\x70\xb5\xfd\xaf\x35\x68\x3a\x21\x70\x8b\x50\xb5\xca\x34\x08\xca\x5c\xb9\x25\xd3\x80\x3a\xd8\x9e\x81\x57\xa2\x4e\x9d\x10\xb0\x26\x0e\xc0\x36\x65\x94\x8c\x9d\xd3\x11\x6d\xaa\xc4\xa4\xa8\x3b\x35\x72\x66\x94\x2a\x04\xe4\x00\x35\x79\xac\xd8\xfa\x6b\x29\x84\x53\xd5\x49\x35\x08\x2d\x6a\x87\x3e\x08\x41\x9d\xb3\x9e\x61\x23\x32\x75\x09\x95\x26\x34\x0c\x79\x43\xdc\xf6\x87\xb2\xb2\x9d\x54\x97\xa0\xd5\x21\xc4\xf5\x26\xa0\x3f\x53\x85\xab\x34\x99\xc8\x52\xce\xbc\x04\x64\x79\x46\x1f\xc8\x1a\xac\x25\x5f\x1d\xd6\xf2\x69\x56\xa9\x2e\x41\x9e\x6f\x95\x76\xad\xba\xcd\x45\xf6\x73\x6c\xd6\x1c\xa9\xc9\x45\xd6\x21\xab\xf3\x2d\xe4\xa7\xdf\x42\x49\x56\x2a\x47\x9d\xaa\x5a\x32\xe8\xaf\x69\xa3\x72\x14\x64\xdc\x24\xcf\xb7\xb9\x28\x44\xf4\xe9\x1e\x2f\xf0\x85\xb4\x06\x8f\xdc\x7b\xb3\x98\x10\x0d\x3d\x20\xb8\x68\x57\x0d\x64\x66\x7f\x93\xbd\x8a\x31\x88\x63\x6f\xaa\x98\xbc\x29\xe0\xc3\x9c\x32\x88\x6c\x06\x99\x23\x83\xc8\x00\x3e\xf6\x07\xf4\x06\x19\xc3\x67\x0c\xb6\xf7\x15\xde\xab\x0e\x5f\xff\x20\xbe\x15\x59\x36\x0c\x37\xe0\xd3\xc9\xbf\x22\x53\xe3\xb7\x2d\xbc\x42\x8d\x5d\x3c\x85\xd7\x6f\xa0\xdc\x31\x76\x01\xc6\x51\x64\xd9\x7b\xe4\x61\x58\x7e\x96\x8f\x0e\xab\xf2\x23\x99\x7a\x1c\xff\xbc\x4e\x24\xcf\x6c\x58\xd8\x30\x86\x44\x36\x8a\x31\x79\xb2\x08\xaa\xf1\x48\x06\x03\x28\xad\x93\xfa\x29\x0e\x51\x3a\x93\x35\x01\xb8\x55\x0c\xca\x23\xe0\x37\x67\xa3\x51\xff\xb1\x29\x1e\xf7\x8a\x17\xd8\xf7\x15\xc3\x20\x7e\xe4\x4a\x82\xde\x04\xf6\x64\x9a\x02\x96\x17\xf1\x72\x47\x9e\xd1\x3b\xe1\x4f\xfd\x52\xfe\x95\x96\x2d\x4c\x2c\xcb\x5a\xc0\x86\x0c\xa3\x3f\xaa\x0a\x87\x71\x0b\xe8\xbd\xf5\x73\x09\x93\x4d\xa3\x10\x2f\xad\x47\xca\xe7\x4e\x00\x2e\xb1\xf7\x8e\x64\xea\x64\x9c\x9f\xcd\x80\xc3\x15\x8c\xea\x70\x6a\xb3\x67\x30\x66\x41\xf0\x44\x57\xcc\x5f\xc5\xc5\x8f\xe0\x54\x85\xff\x2b\x33\xb6\xb0\xb6\x4d\x83\x3e\xea\x98\xf1\x3e\xa5\x80\xc8\xa6\x69\x7e\x44\xde\xc2\x3f\xf1\xf7\x7a\x37\x94\xf7\x78\x79\x67\xfd\xc4\xbb\xb8\xfb\xf9\xee\x71\x3f\x45\x8a\x38\x16\x93\xac\xc4\x93\xa0\x17\xac\xf2\xa9\xb0\xbf\x75\xef\x95\x8e\xb2\xc6\x71\xb3\x56\x5d\x94\xef\x91\xd3\xe7\x16\xa6\x29\x8f\x81\x07\x97\xda\x70\x18\x0b\x91\xd1\x31\x61\xff\xf2\x06\x0c\xe9\xa8\x63\x16\x52\x7e\x21\x6e\xef\xa2\xba\x0d\x7a\x5f\x94\xd3\x6b\x9e\xf4\x42\x83\xcc\xf1\x62\x0d\xb1\x97\xad\xa3\x2a\x2f\x44\xb6\x0c\x71\x9e\xa7\x72\xe3\x58\xac\x83\xfd\x5d\x88\x21\x2d\x46\xf1\xbd\x2b\xfe\x0d\x00\x00\xff\xff\x2d\x5b\xdf\x7e\x2e\x06\x00\x00")

func template_functions_go_templ() ([]byte, error) {
	return bindata_read(
		_template_functions_go_templ,
		"template_functions.go.templ",
	)
}

var _types_go_templ = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x60\x04\x85\xd5\x26\x36\x7c\x5b\x18\xd8\xc3\x6e\x16\x2d\xb6\x69\x36\x45\x9c\xf6\xb2\x58\x20\x0c\x35\x92\xb9\x96\x48\x96\xa4\xd2\x1a\x86\xfe\x7b\xc1\x0f\xc9\xa6\x25\x6d\xd3\x43\x0e\xab\x93\x34\xf3\xe6\x91\x7c\xf3\x21\x4a\x42\x77\xa4\x44\x78\x5e\x91\x4a\x6e\xc9\x2a\x4d\x59\x2d\x85\x32\x30\x4f\x93\x1a\x0d\x79\x5e\xc1\x6c\xf7\x46\x2f\x98\x58\x12\xc9\x6a\x42\xb7\x8c\xa3\xda\x2f\xe5\xae\xb4\x06\xbd\xb4\xa0\xe5\xf3\x6a\x96\x26\x93\x38\xd5\x70\xc3\x6a\x9c\xa5\x59\x9a\x2e\x97\xf0\x53\x89\x9c\x56\x0c\xb9\x89\xbf\xd6\x5c\x6c\x0c\x31\x8d\x76\xe6\xdd\x1b\xbd\xce\x11\x25\x15\x72\x7f\x55\x22\x5f\x33\x6e\x50\x15\x84\xa2\x7e\xfb\x1f\x0b\x2d\xee\x9e\xbe\x22\x35\x6e\xb1\xc3\x61\xb1\x91\x48\x17\x37\x8c\xe7\x6d\x0b\x39\x16\x8c\xa3\x06\xb3\x45\x78\x22\x1a\x41\xa1\x16\x8d\xa2\x98\x9a\xbd\xc4\x73\xb4\x36\xaa\xa1\x06\x0e\x9d\x14\x8b\x87\xbd\xc4\x5b\x34\x04\x00\x1e\xbf\x6a\xc1\xd7\xb3\x4b\xc6\x2b\xc6\x71\xf6\xd8\x63\xfc\xe2\x0e\x15\x30\xd6\x91\x13\x43\x2c\xa8\x5b\xc1\x32\xb5\x2d\x00\x9c\xad\x39\xf0\x07\x8e\x53\xfb\xaf\x9b\xbb\x4f\x6d\x6b\xd9\xbc\x5e\x70\xf2\xc4\x6c\x03\x7f\x60\xd3\xce\x6e\x19\xee\x1a\x23\x1b\x33\xcd\x30\xf0\x07\x06\xe1\xec\x96\xe1\x5d\x9e\x33\xc3\x04\x27\xd5\x7d\xd0\x52\x9f\x71\x8c\x21\x02\x0d\x19\xba\x66\x8f\x69\x9b\xa6\x87\xc3\x15\xb0\x02\xe6\x25\xc2\xbc\x42\x0e\x9e\xef\xbd\xc8\xf7\x8b\x0d\xdd\x62\x4d\x16\x37\xb8\xbf\x25\x52\x32\x5e\x66\xb0\xca\xda\xd6\x85\x28\xc2\x4b\x84\x0b\xc6\x73\xfc\xe7\x12\x2e\xb0\xc2\x1a\xb9\x81\xf5\xdb\x6f\x12\xb4\xad\x2f\x95\x8b\x38\x13\x9e\xe6\xac\x6a\x4e\xcc\x5d\xed\x40\x21\xd4\x79\x74\x57\x4f\x53\x94\xc7\xd2\xea\x4e\x2a\x14\xcc\xf1\xaf\xb0\x77\x98\xd9\xb0\x59\x16\x99\x3e\xd8\x1a\xca\xec\x51\x93\xeb\x4a\x34\xf9\xcf\x42\xd5\xc4\xaa\xf7\x80\xb5\xac\x88\xc1\x4f\xa4\x46\xcb\xcc\x78\xd9\x09\x4c\x27\x81\x36\x77\xd3\x34\x5a\x12\xfa\x72\x2e\x87\xb6\x84\xf7\xa2\xaa\x9e\x08\xdd\x5d\x8b\x86\x1b\x60\xdc\x74\xb1\xea\xd4\xe1\x1b\xe1\x0a\xd0\x29\x65\x4b\xae\xcf\x9c\x54\x42\xda\x74\x75\xa9\xb3\x87\xed\x14\xe2\xc1\xef\x9a\x00\x66\x9d\xfc\xee\x30\x59\x20\xf2\x80\x8d\x53\xf7\x06\xf7\x6d\xdb\xdb\x6c\xd0\x5d\xe1\xbb\x13\xa2\xb4\x1c\x7b\xcc\x23\x5d\xdc\xd8\x1e\xfd\xbb\xad\x34\xff\xe6\x6c\x95\xc6\xc9\x74\x47\xbd\x7c\x4c\xf9\x77\x9f\xbd\x70\x74\x27\xc2\x70\xca\x86\x99\x71\xda\x35\x7e\x5c\x9c\x77\xcc\x58\xc3\x8c\x10\xf5\xc2\x8d\x56\x8a\x0f\xf0\xd0\xae\xb5\x7f\x57\x42\xa2\x32\x0c\x35\xbc\x42\x5d\xc4\xc5\x30\x3c\x7f\x98\xba\x5b\x51\xe5\xfe\xf4\x7e\xdc\x82\x28\xdc\x97\x4b\x5b\xd1\x25\x02\x4c\xc8\xc4\x98\x00\x81\xe9\x44\x80\x6e\x4c\x46\x93\x3d\x4e\xb4\x8a\x20\x6e\xbf\x71\xd4\x3d\x12\x2d\xf8\x37\xa3\x3c\xa4\x3f\x2b\x2b\x82\xcc\x7f\x68\x8c\xab\xce\xc9\xbb\x31\x84\xee\x3e\x7e\x80\xe8\x89\xe9\xb5\x87\xbc\x40\xbd\xb1\xbf\xc5\x51\xca\xe3\x0f\xa3\x2f\x26\x3d\xa6\xdc\x18\xcb\x60\xe6\xfa\x80\x11\xe8\x62\x83\xea\x99\x51\xd4\x76\xfc\x74\xef\x49\xf2\xf9\x8b\x3f\x54\xd2\x1d\x2a\x78\xe2\x59\xf1\x12\x76\xaa\xd0\x04\x72\xf7\x9a\x8c\x91\x3b\xcf\xff\xe5\xbe\x16\xbc\x60\xe5\x2d\x91\x8e\xfe\xf8\x35\xe0\xa7\xbd\x2b\x5e\xc2\xe7\xe4\x75\xee\x5f\xbf\x31\x1d\xcf\x85\xca\x1a\x88\x31\x8a\x3d\x35\xc6\x8f\x05\xff\x93\x8d\xfa\xd0\xe6\x77\x2c\xc9\x8e\x6e\xfa\x92\x36\x79\x45\xb3\x71\x53\x17\xb4\x8f\x06\xeb\xd3\x1b\xd3\xe7\x2f\x67\x9b\x09\x31\xcc\xe2\xfc\x5d\xa5\x68\x38\x05\xc6\x99\x99\x67\x76\x1b\x95\xa0\xa4\x72\x83\x08\xdf\x37\xac\xca\x51\x2d\xee\xb1\x64\xda\xa0\x9a\x93\x3c\x8f\xe9\xec\x66\x75\xd6\xb3\x8c\xfb\xe7\xda\xb1\xc1\x8f\x9d\xc0\x9e\x3d\x03\x54\xca\x0e\xd2\x34\xf1\x00\x5b\x10\x37\x5c\xfc\xcd\x7d\x94\x47\xfd\xa2\x44\x23\xff\x44\xa5\x99\xe0\x97\x69\x92\xfc\x70\x76\xe1\x6c\x47\x8c\x56\x20\xe7\xc8\x7a\xcd\xde\xe5\xf9\x83\x38\xe5\x0a\x9b\xba\x84\xe1\x32\x59\x9a\x28\x34\x8d\xe2\xc0\x59\x95\xb6\xe9\xbf\x01\x00\x00\xff\xff\x10\x11\xf5\x38\x65\x0c\x00\x00")

func types_go_templ() ([]byte, error) {
	return bindata_read(
		_types_go_templ,
		"types.go.templ",
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
	"aws-service-operator.yaml.templ": aws_service_operator_yaml_templ,
	"base.go.templ":                   base_go_templ,
	"cft.go.templ":                    cft_go_templ,
	"operator.go.templ":               operator_go_templ,
	"template_functions.go.templ":     template_functions_go_templ,
	"types.go.templ":                  types_go_templ,
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
	"aws-service-operator.yaml.templ": &_bintree_t{aws_service_operator_yaml_templ, map[string]*_bintree_t{}},
	"base.go.templ":                   &_bintree_t{base_go_templ, map[string]*_bintree_t{}},
	"cft.go.templ":                    &_bintree_t{cft_go_templ, map[string]*_bintree_t{}},
	"operator.go.templ":               &_bintree_t{operator_go_templ, map[string]*_bintree_t{}},
	"template_functions.go.templ":     &_bintree_t{template_functions_go_templ, map[string]*_bintree_t{}},
	"types.go.templ":                  &_bintree_t{types_go_templ, map[string]*_bintree_t{}},
}}
