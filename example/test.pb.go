// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,5,opt,name=OptionalGroup,json=optionalgroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "example.Test.OptionalGroup")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55,
	0x3a, 0xc4, 0xc8, 0xc5, 0x12, 0x92, 0x5a, 0x5c, 0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x93, 0x98, 0x94,
	0x9a, 0x23, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0x94, 0x54,
	0x16, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x5a, 0x31, 0x99, 0x9b, 0x07, 0x81, 0xf9, 0x42,
	0x42, 0x5c, 0x2c, 0x45, 0xa9, 0x05, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xcc, 0x41, 0x60, 0xb6,
	0x90, 0x23, 0x17, 0x6f, 0x7e, 0x41, 0x49, 0x66, 0x7e, 0x5e, 0x62, 0x4e, 0x7a, 0x51, 0x7e, 0x69,
	0x81, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x97, 0x91, 0xb4, 0x1e, 0xd4, 0x2e, 0x3d, 0x90, 0x3d, 0x7a,
	0xfe, 0x50, 0x25, 0xee, 0x20, 0x25, 0x41, 0xa8, 0x3a, 0xa4, 0x4c, 0xb9, 0x78, 0x51, 0xe4, 0x85,
	0x54, 0xb8, 0x78, 0x83, 0x52, 0x0b, 0x4b, 0x33, 0x8b, 0x52, 0x53, 0xdc, 0x32, 0x53, 0x73, 0x52,
	0x24, 0x58, 0xc1, 0xae, 0x43, 0x15, 0xd4, 0xe2, 0xe1, 0x62, 0x76, 0xf3, 0xf7, 0x17, 0x62, 0xe5,
	0x62, 0x8c, 0x10, 0x10, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x3a, 0x34, 0x51, 0xe8, 0x00,
	0x00, 0x00,
}
