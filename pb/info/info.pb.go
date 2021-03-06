// Code generated by protoc-gen-go. DO NOT EDIT.
// source: info.proto

package info

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_cc5eb80ce96b5cb4, []int{0}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(dst, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type GetInfoResponse struct {
	Version              string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Webservices          []*WebServiceStatus `protobuf:"bytes,2,rep,name=webservices,proto3" json:"webservices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_cc5eb80ce96b5cb4, []int{1}
}
func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(dst, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetInfoResponse) GetWebservices() []*WebServiceStatus {
	if m != nil {
		return m.Webservices
	}
	return nil
}

type WebServiceStatus struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebServiceStatus) Reset()         { *m = WebServiceStatus{} }
func (m *WebServiceStatus) String() string { return proto.CompactTextString(m) }
func (*WebServiceStatus) ProtoMessage()    {}
func (*WebServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_cc5eb80ce96b5cb4, []int{2}
}
func (m *WebServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebServiceStatus.Unmarshal(m, b)
}
func (m *WebServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebServiceStatus.Marshal(b, m, deterministic)
}
func (dst *WebServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebServiceStatus.Merge(dst, src)
}
func (m *WebServiceStatus) XXX_Size() int {
	return xxx_messageInfo_WebServiceStatus.Size(m)
}
func (m *WebServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_WebServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_WebServiceStatus proto.InternalMessageInfo

func (m *WebServiceStatus) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *WebServiceStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebServiceStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *WebServiceStatus) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "info.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "info.GetInfoResponse")
	proto.RegisterType((*WebServiceStatus)(nil), "info.WebServiceStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoServiceClient interface {
	Get(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type infoServiceClient struct {
	cc *grpc.ClientConn
}

func NewInfoServiceClient(cc *grpc.ClientConn) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) Get(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/info.InfoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
type InfoServiceServer interface {
	Get(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Get(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InfoService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}

func init() { proto.RegisterFile("info.proto", fileDescriptor_info_cc5eb80ce96b5cb4) }

var fileDescriptor_info_cc5eb80ce96b5cb4 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0xd5, 0x24, 0xb4, 0xe0, 0x08, 0xa8, 0x2c, 0xa8, 0x4e, 0x11, 0x43, 0x95, 0xa9, 0x53,
	0x23, 0x95, 0x85, 0x95, 0xa9, 0x62, 0x43, 0xe9, 0xc0, 0x9c, 0x80, 0x5b, 0x9d, 0x04, 0x76, 0xc8,
	0xb9, 0x85, 0x99, 0x57, 0xe0, 0xd1, 0x78, 0x05, 0x1e, 0x04, 0xe5, 0x2e, 0x54, 0xd0, 0xed, 0xff,
	0x7f, 0x7f, 0x27, 0xfb, 0x7e, 0x00, 0xcb, 0x6b, 0x99, 0x37, 0xad, 0xa8, 0x60, 0xd2, 0xe9, 0xec,
	0x6a, 0x23, 0xb2, 0x79, 0xa6, 0xa2, 0x6a, 0x6c, 0x51, 0x31, 0x8b, 0x56, 0x6a, 0x85, 0x5d, 0x60,
	0xf2, 0x31, 0x9c, 0x2d, 0x49, 0xef, 0x78, 0x2d, 0x25, 0xbd, 0x6e, 0xc9, 0x69, 0x4e, 0x70, 0xbe,
	0x4f, 0x5c, 0x23, 0xec, 0x08, 0x0d, 0x8c, 0x76, 0xd4, 0x3a, 0x2b, 0x6c, 0x06, 0xd3, 0xc1, 0xec,
	0xa4, 0xfc, 0xb5, 0x78, 0x03, 0xe9, 0x1b, 0xd5, 0x8e, 0xda, 0x9d, 0x7d, 0x24, 0x67, 0xa2, 0x69,
	0x3c, 0x4b, 0x17, 0x93, 0xb9, 0x3f, 0xe2, 0x81, 0xea, 0x55, 0x18, 0xac, 0xb4, 0xd2, 0xad, 0x2b,
	0xff, 0xa2, 0x39, 0xc3, 0xf8, 0x10, 0xc0, 0x0c, 0x8e, 0x89, 0x9f, 0x1a, 0xb1, 0xac, 0xfd, 0xa2,
	0xbd, 0x47, 0x84, 0x84, 0xab, 0x17, 0x32, 0x91, 0xcf, 0xbd, 0xc6, 0x09, 0x0c, 0x9d, 0x7f, 0x69,
	0x62, 0x9f, 0xf6, 0xae, 0x63, 0x95, 0xde, 0xd5, 0x24, 0x81, 0xed, 0xf4, 0xe2, 0x1e, 0xd2, 0xee,
	0x4f, 0xfd, 0x42, 0xbc, 0x85, 0x78, 0x49, 0x8a, 0x17, 0xe1, 0xd4, 0xff, 0x15, 0x64, 0x97, 0x07,
	0x69, 0xa8, 0x21, 0x3f, 0xfd, 0xf8, 0xfa, 0xfe, 0x8c, 0x46, 0x78, 0x54, 0x74, 0xe3, 0x7a, 0xe8,
	0x1b, 0xbc, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xec, 0xa0, 0x97, 0x73, 0x01, 0x00, 0x00,
}
