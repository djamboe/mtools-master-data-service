// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master-data-service.proto

package master_data_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Customer struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Customername         string   `protobuf:"bytes,2,opt,name=customername,proto3" json:"customername,omitempty"`
	Customeremail        string   `protobuf:"bytes,3,opt,name=customeremail,proto3" json:"customeremail,omitempty"`
	Customerphone        string   `protobuf:"bytes,4,opt,name=customerphone,proto3" json:"customerphone,omitempty"`
	Customeraddress      string   `protobuf:"bytes,5,opt,name=customeraddress,proto3" json:"customeraddress,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Customer) GetCustomername() string {
	if m != nil {
		return m.Customername
	}
	return ""
}

func (m *Customer) GetCustomeremail() string {
	if m != nil {
		return m.Customeremail
	}
	return ""
}

func (m *Customer) GetCustomerphone() string {
	if m != nil {
		return m.Customerphone
	}
	return ""
}

func (m *Customer) GetCustomeraddress() string {
	if m != nil {
		return m.Customeraddress
	}
	return ""
}

func (m *Customer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetCustomerDataRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Collection           string   `protobuf:"bytes,4,opt,name=collection,proto3" json:"collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerDataRequest) Reset()         { *m = GetCustomerDataRequest{} }
func (m *GetCustomerDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerDataRequest) ProtoMessage()    {}
func (*GetCustomerDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{1}
}

func (m *GetCustomerDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerDataRequest.Unmarshal(m, b)
}
func (m *GetCustomerDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerDataRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerDataRequest.Merge(m, src)
}
func (m *GetCustomerDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerDataRequest.Size(m)
}
func (m *GetCustomerDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerDataRequest proto.InternalMessageInfo

func (m *GetCustomerDataRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetCustomerDataRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetCustomerDataRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GetCustomerDataRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetCustomerDataResponse struct {
	Api                  string      `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                bool        `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	Customer             []*Customer `protobuf:"bytes,4,rep,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCustomerDataResponse) Reset()         { *m = GetCustomerDataResponse{} }
func (m *GetCustomerDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetCustomerDataResponse) ProtoMessage()    {}
func (*GetCustomerDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{2}
}

func (m *GetCustomerDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerDataResponse.Unmarshal(m, b)
}
func (m *GetCustomerDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerDataResponse.Marshal(b, m, deterministic)
}
func (m *GetCustomerDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerDataResponse.Merge(m, src)
}
func (m *GetCustomerDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetCustomerDataResponse.Size(m)
}
func (m *GetCustomerDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerDataResponse proto.InternalMessageInfo

func (m *GetCustomerDataResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetCustomerDataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetCustomerDataResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *GetCustomerDataResponse) GetCustomer() []*Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type Product struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Productname          string   `protobuf:"bytes,2,opt,name=productname,proto3" json:"productname,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{3}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Product) GetProductname() string {
	if m != nil {
		return m.Productname
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetProductDataRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Collection           string   `protobuf:"bytes,4,opt,name=collection,proto3" json:"collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductDataRequest) Reset()         { *m = GetProductDataRequest{} }
func (m *GetProductDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductDataRequest) ProtoMessage()    {}
func (*GetProductDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{4}
}

func (m *GetProductDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductDataRequest.Unmarshal(m, b)
}
func (m *GetProductDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductDataRequest.Marshal(b, m, deterministic)
}
func (m *GetProductDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductDataRequest.Merge(m, src)
}
func (m *GetProductDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductDataRequest.Size(m)
}
func (m *GetProductDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductDataRequest proto.InternalMessageInfo

func (m *GetProductDataRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetProductDataRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetProductDataRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GetProductDataRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetProductDataResponse struct {
	Api                  string     `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                bool       `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	Product              []*Product `protobuf:"bytes,4,rep,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProductDataResponse) Reset()         { *m = GetProductDataResponse{} }
func (m *GetProductDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductDataResponse) ProtoMessage()    {}
func (*GetProductDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e48557cb742acd00, []int{5}
}

func (m *GetProductDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductDataResponse.Unmarshal(m, b)
}
func (m *GetProductDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductDataResponse.Marshal(b, m, deterministic)
}
func (m *GetProductDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductDataResponse.Merge(m, src)
}
func (m *GetProductDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductDataResponse.Size(m)
}
func (m *GetProductDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductDataResponse proto.InternalMessageInfo

func (m *GetProductDataResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetProductDataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetProductDataResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *GetProductDataResponse) GetProduct() []*Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "Customer")
	proto.RegisterType((*GetCustomerDataRequest)(nil), "GetCustomerDataRequest")
	proto.RegisterType((*GetCustomerDataResponse)(nil), "GetCustomerDataResponse")
	proto.RegisterType((*Product)(nil), "Product")
	proto.RegisterType((*GetProductDataRequest)(nil), "GetProductDataRequest")
	proto.RegisterType((*GetProductDataResponse)(nil), "GetProductDataResponse")
}

func init() { proto.RegisterFile("master-data-service.proto", fileDescriptor_e48557cb742acd00) }

var fileDescriptor_e48557cb742acd00 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x14, 0x54, 0xda, 0x6e, 0xdb, 0x7d, 0x0b, 0x14, 0x2c, 0x68, 0x4d, 0x0f, 0xa8, 0x58, 0x2c, 0xaa,
	0x80, 0x26, 0x62, 0xb9, 0x71, 0x05, 0x69, 0xc5, 0x01, 0x09, 0x85, 0x0f, 0x58, 0x99, 0xe4, 0xa9,
	0x1b, 0x29, 0x89, 0xbd, 0xb6, 0x13, 0x69, 0x2f, 0x1c, 0xf6, 0x17, 0xf8, 0x2e, 0x4e, 0xfc, 0x02,
	0xff, 0xc0, 0x15, 0xc5, 0xb1, 0x51, 0xb2, 0x0d, 0x37, 0xc4, 0xad, 0x33, 0x6f, 0xa4, 0xe9, 0xcc,
	0xf3, 0x0b, 0x3c, 0x2e, 0xb8, 0x36, 0xa8, 0x76, 0x29, 0x37, 0x7c, 0xa7, 0x51, 0xd5, 0x59, 0x82,
	0xa1, 0x54, 0xc2, 0x88, 0xf5, 0xd6, 0x5c, 0x66, 0x2a, 0xbd, 0x90, 0x5c, 0x99, 0xeb, 0x68, 0x2f,
	0xc4, 0x3e, 0xc7, 0x88, 0xcb, 0x2c, 0xe2, 0x65, 0x29, 0x0c, 0x37, 0x99, 0x28, 0x75, 0xab, 0x64,
	0xdf, 0x03, 0x98, 0xbf, 0xab, 0xb4, 0x11, 0x05, 0x2a, 0xb2, 0x80, 0xf1, 0x45, 0x96, 0xd2, 0x60,
	0x13, 0x6c, 0x8f, 0xe3, 0xd1, 0x87, 0x94, 0x30, 0xb8, 0x93, 0xb8, 0x61, 0xc9, 0x0b, 0xa4, 0x23,
	0x3b, 0xe9, 0x71, 0xe4, 0x19, 0xdc, 0xf5, 0x18, 0x0b, 0x9e, 0xe5, 0x74, 0x6c, 0x45, 0x7d, 0xb2,
	0xab, 0x92, 0x97, 0xa2, 0x44, 0x3a, 0xe9, 0xab, 0x2c, 0x49, 0xb6, 0xb0, 0xf0, 0x04, 0x4f, 0x53,
	0x85, 0x5a, 0xd3, 0x23, 0xab, 0xbb, 0x4d, 0x93, 0x25, 0x4c, 0xb5, 0xe1, 0xa6, 0xd2, 0x74, 0x6a,
	0x05, 0x0e, 0xb1, 0x1a, 0x96, 0xe7, 0x68, 0x7c, 0xa2, 0xf7, 0xdc, 0xf0, 0x18, 0xaf, 0x2a, 0xd4,
	0x86, 0xdc, 0x87, 0x31, 0x97, 0x99, 0x0b, 0xd7, 0xfc, 0x24, 0x0f, 0xe1, 0xe8, 0xaa, 0x42, 0x75,
	0xed, 0x62, 0xb5, 0xa0, 0x61, 0x6b, 0x9e, 0x57, 0xe8, 0x72, 0xb4, 0x80, 0x3c, 0x01, 0x48, 0x44,
	0x9e, 0x63, 0xd2, 0x94, 0xe7, 0xfe, 0x7c, 0x87, 0x61, 0x37, 0x01, 0xac, 0x0e, 0x8c, 0xb5, 0x14,
	0xa5, 0xc6, 0x01, 0x67, 0x0a, 0xb3, 0x02, 0xb5, 0xe6, 0x7b, 0x5f, 0xa9, 0x87, 0x8d, 0x3b, 0x2a,
	0x25, 0x94, 0x75, 0x9f, 0xc7, 0x2d, 0x20, 0xa7, 0x30, 0xf7, 0x05, 0xd0, 0xc9, 0x66, 0xbc, 0x3d,
	0x39, 0x3b, 0x0e, 0xbd, 0x55, 0xfc, 0x67, 0xc4, 0x6a, 0x98, 0x7d, 0x52, 0x22, 0xad, 0x12, 0x73,
	0xb8, 0xca, 0x0d, 0x9c, 0xc8, 0x76, 0xd6, 0xd9, 0x64, 0x97, 0x6a, 0x14, 0x29, 0xea, 0x44, 0x65,
	0xd2, 0x66, 0x6c, 0xe3, 0x77, 0xa9, 0x4e, 0xe9, 0x93, 0x5e, 0xe9, 0x15, 0x3c, 0x3a, 0x47, 0xe3,
	0xac, 0xff, 0x5f, 0xe7, 0x5f, 0xed, 0xae, 0x7b, 0xb6, 0xff, 0xac, 0x71, 0x06, 0x33, 0xd7, 0x8d,
	0x2b, 0x7c, 0x1e, 0x3a, 0xa3, 0xd8, 0x0f, 0xce, 0x7e, 0x05, 0xf0, 0xe0, 0xa3, 0xbd, 0xc1, 0xc6,
	0xfc, 0x73, 0x7b, 0x81, 0x44, 0xc1, 0xe2, 0xd6, 0x43, 0x20, 0xab, 0x70, 0xf8, 0x4d, 0xae, 0x69,
	0xf8, 0x97, 0x37, 0xc3, 0x5e, 0xdd, 0xfc, 0xf8, 0xf9, 0x6d, 0xf4, 0x9c, 0x3d, 0x8d, 0xea, 0xd7,
	0x51, 0xe7, 0xd0, 0xa3, 0x3d, 0x9a, 0x9d, 0x5f, 0xb8, 0x65, 0xde, 0x06, 0x2f, 0x48, 0x09, 0xf7,
	0xfa, 0x4d, 0x90, 0x65, 0x38, 0xb8, 0x91, 0xf5, 0x2a, 0x1c, 0xae, 0x8c, 0xbd, 0xb4, 0x86, 0xa7,
	0x6c, 0x33, 0x64, 0xe8, 0x12, 0x7b, 0xbf, 0x2f, 0x53, 0xfb, 0xf1, 0x78, 0xf3, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x31, 0x67, 0x0d, 0x83, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterDataServiceClient is the client API for MasterDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterDataServiceClient interface {
	GetCustomerData(ctx context.Context, in *GetCustomerDataRequest, opts ...grpc.CallOption) (*GetCustomerDataResponse, error)
	GetProductData(ctx context.Context, in *GetProductDataRequest, opts ...grpc.CallOption) (*GetProductDataResponse, error)
}

type masterDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewMasterDataServiceClient(cc *grpc.ClientConn) MasterDataServiceClient {
	return &masterDataServiceClient{cc}
}

func (c *masterDataServiceClient) GetCustomerData(ctx context.Context, in *GetCustomerDataRequest, opts ...grpc.CallOption) (*GetCustomerDataResponse, error) {
	out := new(GetCustomerDataResponse)
	err := c.cc.Invoke(ctx, "/MasterDataService/GetCustomerData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDataServiceClient) GetProductData(ctx context.Context, in *GetProductDataRequest, opts ...grpc.CallOption) (*GetProductDataResponse, error) {
	out := new(GetProductDataResponse)
	err := c.cc.Invoke(ctx, "/MasterDataService/GetProductData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterDataServiceServer is the server API for MasterDataService service.
type MasterDataServiceServer interface {
	GetCustomerData(context.Context, *GetCustomerDataRequest) (*GetCustomerDataResponse, error)
	GetProductData(context.Context, *GetProductDataRequest) (*GetProductDataResponse, error)
}

// UnimplementedMasterDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMasterDataServiceServer struct {
}

func (*UnimplementedMasterDataServiceServer) GetCustomerData(ctx context.Context, req *GetCustomerDataRequest) (*GetCustomerDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerData not implemented")
}
func (*UnimplementedMasterDataServiceServer) GetProductData(ctx context.Context, req *GetProductDataRequest) (*GetProductDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductData not implemented")
}

func RegisterMasterDataServiceServer(s *grpc.Server, srv MasterDataServiceServer) {
	s.RegisterService(&_MasterDataService_serviceDesc, srv)
}

func _MasterDataService_GetCustomerData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDataServiceServer).GetCustomerData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MasterDataService/GetCustomerData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDataServiceServer).GetCustomerData(ctx, req.(*GetCustomerDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDataService_GetProductData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDataServiceServer).GetProductData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MasterDataService/GetProductData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDataServiceServer).GetProductData(ctx, req.(*GetProductDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MasterDataService",
	HandlerType: (*MasterDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerData",
			Handler:    _MasterDataService_GetCustomerData_Handler,
		},
		{
			MethodName: "GetProductData",
			Handler:    _MasterDataService_GetProductData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master-data-service.proto",
}
