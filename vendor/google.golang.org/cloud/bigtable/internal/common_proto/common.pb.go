// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/common_proto/common.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/common_proto/common.proto

It has these top-level messages:
*/
package google_bigtable_admin_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StorageType int32

const (
	StorageType_STORAGE_TYPE_UNSPECIFIED StorageType = 0
	StorageType_SSD                      StorageType = 1
	StorageType_HDD                      StorageType = 2
)

var StorageType_name = map[int32]string{
	0: "STORAGE_TYPE_UNSPECIFIED",
	1: "SSD",
	2: "HDD",
}
var StorageType_value = map[string]int32{
	"STORAGE_TYPE_UNSPECIFIED": 0,
	"SSD": 1,
	"HDD": 2,
}

func (x StorageType) String() string {
	return proto.EnumName(StorageType_name, int32(x))
}
func (StorageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("google.bigtable.admin.v2.StorageType", StorageType_name, StorageType_value)
}

func init() {
	proto.RegisterFile("google.golang.org/cloud/bigtable/internal/common_proto/common.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8e, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0xc6, 0x7d, 0x0a, 0x0a, 0x79, 0x4b, 0xe9, 0xf4, 0x90, 0x82, 0xbb, 0x60, 0x02, 0x3a, 0x3b,
	0xd8, 0x26, 0x6a, 0x17, 0x0d, 0xa6, 0x0e, 0x4e, 0x25, 0x69, 0x63, 0x08, 0x24, 0xb9, 0x12, 0x53,
	0xc1, 0xff, 0x5e, 0x9b, 0xd6, 0xf1, 0x6d, 0xf7, 0x71, 0xf7, 0xdd, 0xef, 0x87, 0x1a, 0x03, 0x60,
	0x9c, 0xc6, 0x06, 0x9c, 0x0c, 0x06, 0x43, 0x34, 0x64, 0x70, 0x30, 0x8f, 0x44, 0x59, 0x93, 0xa4,
	0x72, 0x9a, 0xd8, 0x90, 0x74, 0x0c, 0xd2, 0x91, 0x01, 0xbc, 0x87, 0xd0, 0x4f, 0x11, 0x12, 0x6c,
	0x01, 0xe7, 0x50, 0x1e, 0xb6, 0x27, 0xff, 0x1d, 0x2c, 0x47, 0x6f, 0x03, 0xfe, 0xbe, 0xbd, 0xbc,
	0x5a, 0x37, 0x24, 0xdf, 0xa9, 0xf9, 0x93, 0x24, 0xeb, 0xf5, 0x57, 0x92, 0x7e, 0x5a, 0xab, 0xd7,
	0xf7, 0x68, 0x2f, 0x12, 0x44, 0x69, 0x74, 0xf7, 0x33, 0xe9, 0xb2, 0x42, 0x07, 0xd1, 0xbd, 0xbe,
	0x3d, 0x3c, 0xb1, 0xbe, 0xfb, 0xe0, 0xac, 0x7f, 0x7f, 0x11, 0x9c, 0x35, 0xed, 0x63, 0xcb, 0x68,
	0x71, 0x52, 0x5e, 0xa0, 0x33, 0x21, 0x68, 0xb1, 0x5b, 0x86, 0x67, 0x4a, 0x8b, 0xd3, 0xfa, 0x06,
	0x55, 0x7f, 0x26, 0xf8, 0x18, 0xbf, 0xde, 0x37, 0xd9, 0x93, 0x2f, 0x2c, 0xbe, 0x53, 0xe7, 0x19,
	0x7a, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x91, 0xdb, 0x20, 0x91, 0xf6, 0x00, 0x00, 0x00,
}
