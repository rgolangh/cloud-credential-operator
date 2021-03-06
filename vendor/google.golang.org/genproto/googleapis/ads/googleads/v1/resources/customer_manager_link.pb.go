// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_manager_link.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents customer-manager link relationship.
type CustomerManagerLink struct {
	// Name of the resource.
	// CustomerManagerLink resource names have the form:
	//
	// `customers/{customer_id}/customerManagerLinks/{manager_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The manager customer linked to the customer.
	ManagerCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=manager_customer,json=managerCustomer,proto3" json:"manager_customer,omitempty"`
	// ID of the customer-manager link. This field is read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// Status of the link between the customer and the manager.
	Status               enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CustomerManagerLink) Reset()         { *m = CustomerManagerLink{} }
func (m *CustomerManagerLink) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLink) ProtoMessage()    {}
func (*CustomerManagerLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3318e064b05720, []int{0}
}

func (m *CustomerManagerLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLink.Unmarshal(m, b)
}
func (m *CustomerManagerLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLink.Marshal(b, m, deterministic)
}
func (m *CustomerManagerLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLink.Merge(m, src)
}
func (m *CustomerManagerLink) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLink.Size(m)
}
func (m *CustomerManagerLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLink proto.InternalMessageInfo

func (m *CustomerManagerLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerManagerLink) GetManagerCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ManagerCustomer
	}
	return nil
}

func (m *CustomerManagerLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerManagerLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerManagerLink)(nil), "google.ads.googleads.v1.resources.CustomerManagerLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_manager_link.proto", fileDescriptor_7e3318e064b05720)
}

var fileDescriptor_7e3318e064b05720 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xd1, 0xca, 0xd3, 0x30,
	0x14, 0xa6, 0xfd, 0xf5, 0x07, 0xab, 0x73, 0x52, 0x6f, 0xca, 0x1c, 0xb2, 0x29, 0x83, 0x5d, 0x25,
	0x74, 0x8a, 0x42, 0xc4, 0x8b, 0x6e, 0xe8, 0x98, 0xa8, 0x8c, 0x0e, 0x7a, 0x21, 0xc5, 0x92, 0xad,
	0x31, 0x94, 0x35, 0x49, 0x49, 0xd2, 0xf9, 0x04, 0xbe, 0x88, 0x57, 0xe2, 0xa3, 0xf8, 0x28, 0x3e,
	0x85, 0xac, 0x69, 0xea, 0x64, 0x4e, 0xef, 0x4e, 0x73, 0xbe, 0xef, 0x3b, 0xdf, 0x77, 0x4e, 0xbd,
	0x97, 0x54, 0x08, 0x5a, 0x12, 0x88, 0x73, 0x05, 0x4d, 0x79, 0xac, 0x0e, 0x21, 0x94, 0x44, 0x89,
	0x5a, 0xee, 0x88, 0x82, 0xbb, 0x5a, 0x69, 0xc1, 0x88, 0xcc, 0x18, 0xe6, 0x98, 0x12, 0x99, 0x95,
	0x05, 0xdf, 0x83, 0x4a, 0x0a, 0x2d, 0xfc, 0xb1, 0xe1, 0x00, 0x9c, 0x2b, 0xd0, 0xd1, 0xc1, 0x21,
	0x04, 0x1d, 0x7d, 0xf0, 0xfc, 0xd2, 0x04, 0xc2, 0x6b, 0xa6, 0xe0, 0xa9, 0x68, 0xa6, 0x34, 0xd6,
	0xb5, 0x32, 0xda, 0x83, 0xa1, 0x25, 0x56, 0x05, 0xc4, 0x9c, 0x0b, 0x8d, 0x75, 0x21, 0xb8, 0xed,
	0x3e, 0x6c, 0xbb, 0xcd, 0xd7, 0xb6, 0xfe, 0x04, 0x3f, 0x4b, 0x5c, 0x55, 0x44, 0xb6, 0xfd, 0x47,
	0xdf, 0x5c, 0xef, 0xfe, 0xa2, 0x75, 0xfe, 0xce, 0xcc, 0x78, 0x5b, 0xf0, 0xbd, 0xff, 0xd8, 0xeb,
	0x59, 0x6f, 0x19, 0xc7, 0x8c, 0x04, 0xce, 0xc8, 0x99, 0xde, 0x8a, 0xef, 0xd8, 0xc7, 0xf7, 0x98,
	0x11, 0x7f, 0xe9, 0xdd, 0xb3, 0xbe, 0x6c, 0xfa, 0xe0, 0x6a, 0xe4, 0x4c, 0x6f, 0xcf, 0x86, 0x6d,
	0x4c, 0x60, 0xe7, 0x82, 0x8d, 0x96, 0x05, 0xa7, 0x09, 0x2e, 0x6b, 0x12, 0xf7, 0x5b, 0x96, 0x1d,
	0xec, 0x2f, 0xbc, 0xfe, 0x1f, 0x01, 0x8b, 0x3c, 0xb8, 0xd1, 0xe8, 0x3c, 0x38, 0xd3, 0x59, 0x71,
	0xfd, 0xec, 0xa9, 0x91, 0xe9, 0xb1, 0xdf, 0x86, 0x57, 0xb9, 0xff, 0xd1, 0xbb, 0x36, 0x8b, 0x09,
	0x6e, 0x8e, 0x9c, 0xe9, 0xdd, 0xd9, 0x6b, 0x70, 0x69, 0xeb, 0xcd, 0x4a, 0xc1, 0x49, 0xdc, 0x4d,
	0xc3, 0x7b, 0xc5, 0x6b, 0x76, 0xfe, 0x1a, 0xb7, 0xaa, 0xf3, 0x2f, 0xae, 0x37, 0xd9, 0x09, 0x06,
	0xfe, 0x7b, 0xcb, 0x79, 0xf0, 0x97, 0x8d, 0xae, 0x8f, 0x01, 0xd6, 0xce, 0x87, 0x37, 0x2d, 0x9d,
	0x8a, 0x12, 0x73, 0x0a, 0x84, 0xa4, 0x90, 0x12, 0xde, 0xc4, 0xb3, 0x77, 0xaf, 0x0a, 0xf5, 0x8f,
	0x1f, 0xed, 0x45, 0x57, 0x7d, 0x75, 0xaf, 0x96, 0x51, 0xf4, 0xdd, 0x1d, 0x2f, 0x8d, 0x64, 0x94,
	0x2b, 0x60, 0xca, 0x63, 0x95, 0x84, 0x20, 0xb6, 0xc8, 0x1f, 0x16, 0x93, 0x46, 0xb9, 0x4a, 0x3b,
	0x4c, 0x9a, 0x84, 0x69, 0x87, 0xf9, 0xe9, 0x4e, 0x4c, 0x03, 0xa1, 0x28, 0x57, 0x08, 0x75, 0x28,
	0x84, 0x92, 0x10, 0xa1, 0x0e, 0xb7, 0xbd, 0x6e, 0xcc, 0x3e, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x5f, 0x95, 0x56, 0x14, 0x03, 0x00, 0x00,
}
