// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/query.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type QueryTicketsRequest struct {
	// The Pool representing the set of Filters to be queried.
	Pool                 *Pool    `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTicketsRequest) Reset()         { *m = QueryTicketsRequest{} }
func (m *QueryTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTicketsRequest) ProtoMessage()    {}
func (*QueryTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{0}
}

func (m *QueryTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketsRequest.Unmarshal(m, b)
}
func (m *QueryTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketsRequest.Marshal(b, m, deterministic)
}
func (m *QueryTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketsRequest.Merge(m, src)
}
func (m *QueryTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryTicketsRequest.Size(m)
}
func (m *QueryTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketsRequest proto.InternalMessageInfo

func (m *QueryTicketsRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type QueryTicketsResponse struct {
	// Tickets that meet all the filtering criteria requested by the pool.
	Tickets              []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryTicketsResponse) Reset()         { *m = QueryTicketsResponse{} }
func (m *QueryTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTicketsResponse) ProtoMessage()    {}
func (*QueryTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{1}
}

func (m *QueryTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketsResponse.Unmarshal(m, b)
}
func (m *QueryTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketsResponse.Marshal(b, m, deterministic)
}
func (m *QueryTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketsResponse.Merge(m, src)
}
func (m *QueryTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryTicketsResponse.Size(m)
}
func (m *QueryTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketsResponse proto.InternalMessageInfo

func (m *QueryTicketsResponse) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type QueryTicketIdsRequest struct {
	// The Pool representing the set of Filters to be queried.
	Pool                 *Pool    `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTicketIdsRequest) Reset()         { *m = QueryTicketIdsRequest{} }
func (m *QueryTicketIdsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTicketIdsRequest) ProtoMessage()    {}
func (*QueryTicketIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{2}
}

func (m *QueryTicketIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketIdsRequest.Unmarshal(m, b)
}
func (m *QueryTicketIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketIdsRequest.Marshal(b, m, deterministic)
}
func (m *QueryTicketIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketIdsRequest.Merge(m, src)
}
func (m *QueryTicketIdsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryTicketIdsRequest.Size(m)
}
func (m *QueryTicketIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketIdsRequest proto.InternalMessageInfo

func (m *QueryTicketIdsRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type QueryTicketIdsResponse struct {
	// TicketIDs that meet all the filtering criteria requested by the pool.
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryTicketIdsResponse) Reset()         { *m = QueryTicketIdsResponse{} }
func (m *QueryTicketIdsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTicketIdsResponse) ProtoMessage()    {}
func (*QueryTicketIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{3}
}

func (m *QueryTicketIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryTicketIdsResponse.Unmarshal(m, b)
}
func (m *QueryTicketIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryTicketIdsResponse.Marshal(b, m, deterministic)
}
func (m *QueryTicketIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTicketIdsResponse.Merge(m, src)
}
func (m *QueryTicketIdsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryTicketIdsResponse.Size(m)
}
func (m *QueryTicketIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTicketIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTicketIdsResponse proto.InternalMessageInfo

func (m *QueryTicketIdsResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

// BETA FEATURE WARNING:  This Request messages are not finalized and
// still subject to possible change or removal.
type QueryBackfillsRequest struct {
	// The Pool representing the set of Filters to be queried.
	Pool                 *Pool    `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryBackfillsRequest) Reset()         { *m = QueryBackfillsRequest{} }
func (m *QueryBackfillsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBackfillsRequest) ProtoMessage()    {}
func (*QueryBackfillsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{4}
}

func (m *QueryBackfillsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryBackfillsRequest.Unmarshal(m, b)
}
func (m *QueryBackfillsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryBackfillsRequest.Marshal(b, m, deterministic)
}
func (m *QueryBackfillsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBackfillsRequest.Merge(m, src)
}
func (m *QueryBackfillsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryBackfillsRequest.Size(m)
}
func (m *QueryBackfillsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBackfillsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBackfillsRequest proto.InternalMessageInfo

func (m *QueryBackfillsRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

// BETA FEATURE WARNING:  This Request messages are not finalized and
// still subject to possible change or removal.
type QueryBackfillsResponse struct {
	// Backfills that meet all the filtering criteria requested by the pool.
	Backfills            []*Backfill `protobuf:"bytes,1,rep,name=backfills,proto3" json:"backfills,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QueryBackfillsResponse) Reset()         { *m = QueryBackfillsResponse{} }
func (m *QueryBackfillsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBackfillsResponse) ProtoMessage()    {}
func (*QueryBackfillsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec7651f31a90698, []int{5}
}

func (m *QueryBackfillsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryBackfillsResponse.Unmarshal(m, b)
}
func (m *QueryBackfillsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryBackfillsResponse.Marshal(b, m, deterministic)
}
func (m *QueryBackfillsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBackfillsResponse.Merge(m, src)
}
func (m *QueryBackfillsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryBackfillsResponse.Size(m)
}
func (m *QueryBackfillsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBackfillsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBackfillsResponse proto.InternalMessageInfo

func (m *QueryBackfillsResponse) GetBackfills() []*Backfill {
	if m != nil {
		return m.Backfills
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryTicketsRequest)(nil), "openmatch.QueryTicketsRequest")
	proto.RegisterType((*QueryTicketsResponse)(nil), "openmatch.QueryTicketsResponse")
	proto.RegisterType((*QueryTicketIdsRequest)(nil), "openmatch.QueryTicketIdsRequest")
	proto.RegisterType((*QueryTicketIdsResponse)(nil), "openmatch.QueryTicketIdsResponse")
	proto.RegisterType((*QueryBackfillsRequest)(nil), "openmatch.QueryBackfillsRequest")
	proto.RegisterType((*QueryBackfillsResponse)(nil), "openmatch.QueryBackfillsResponse")
}

func init() { proto.RegisterFile("api/query.proto", fileDescriptor_5ec7651f31a90698) }

var fileDescriptor_5ec7651f31a90698 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x4e, 0x13, 0x41,
	0x18, 0xc7, 0xb3, 0x2d, 0x81, 0x74, 0x30, 0x82, 0x83, 0x10, 0xd2, 0x18, 0x1c, 0x96, 0x98, 0x40,
	0xb1, 0x3b, 0xa5, 0x72, 0x55, 0x35, 0xe1, 0x78, 0x41, 0x2c, 0x1e, 0x16, 0xe3, 0x85, 0x77, 0xd3,
	0xd9, 0x8f, 0xed, 0xc8, 0x76, 0x67, 0xd8, 0x99, 0x05, 0x49, 0xbc, 0x32, 0x3e, 0x81, 0xde, 0x18,
	0x1e, 0xc1, 0x97, 0xf0, 0x21, 0x7c, 0x05, 0xe3, 0x73, 0x98, 0x3d, 0xb4, 0x5d, 0x68, 0x21, 0xe1,
	0x6a, 0x77, 0xbe, 0xd3, 0xff, 0x37, 0xdf, 0x37, 0x33, 0x68, 0x86, 0x29, 0x41, 0x4f, 0x63, 0x88,
	0x2e, 0x1c, 0x15, 0x49, 0x23, 0x71, 0x45, 0x2a, 0x08, 0x7b, 0xcc, 0xf0, 0x6e, 0x15, 0x27, 0xbe,
	0x1e, 0x68, 0xcd, 0x7c, 0xd0, 0x99, 0xbb, 0xfa, 0xc8, 0x97, 0xd2, 0x0f, 0x80, 0x26, 0x2e, 0x16,
	0x86, 0xd2, 0x30, 0x23, 0x64, 0xd8, 0xf7, 0x3e, 0x4d, 0x3f, 0xbc, 0xee, 0x43, 0x58, 0xd7, 0xe7,
	0xcc, 0xf7, 0x21, 0xa2, 0x52, 0xa5, 0x11, 0xa3, 0xd1, 0x76, 0x0b, 0xcd, 0xbd, 0x4b, 0x94, 0xdf,
	0x0b, 0x7e, 0x02, 0x46, 0xbb, 0x70, 0x1a, 0x83, 0x36, 0x78, 0x05, 0x4d, 0x28, 0x29, 0x83, 0x45,
	0x8b, 0x58, 0xab, 0xd3, 0xcd, 0x19, 0x67, 0x00, 0xe4, 0xbc, 0x95, 0x32, 0x70, 0x53, 0xa7, 0xbd,
	0x8b, 0x1e, 0x5e, 0xcd, 0xd5, 0x4a, 0x86, 0x1a, 0xf0, 0x3a, 0x9a, 0x32, 0x99, 0x69, 0xd1, 0x22,
	0xe5, 0xd5, 0xe9, 0xe6, 0x83, 0x42, 0x7e, 0x16, 0xec, 0xf6, 0x23, 0xec, 0x17, 0x68, 0xbe, 0x50,
	0xe4, 0xc0, 0xbb, 0x1b, 0x42, 0x0d, 0x2d, 0x5c, 0xcf, 0xce, 0x21, 0x66, 0x51, 0x59, 0x78, 0x19,
	0x40, 0xc5, 0x4d, 0x7e, 0x07, 0x4a, 0x3b, 0x8c, 0x9f, 0x1c, 0x8b, 0x20, 0xb8, 0x9b, 0xd2, 0xab,
	0x5c, 0xa9, 0x90, 0x9d, 0x2b, 0x6d, 0xa0, 0x4a, 0xa7, 0x6f, 0xcc, 0x37, 0x3c, 0x57, 0xa8, 0xd1,
	0x4f, 0x70, 0x87, 0x51, 0xcd, 0xcb, 0x32, 0xba, 0x97, 0x56, 0x3b, 0x82, 0xe8, 0x4c, 0x70, 0xc0,
	0x5f, 0xf2, 0x75, 0xde, 0x4a, 0xbc, 0x54, 0x28, 0x30, 0x66, 0x3e, 0xd5, 0xc7, 0x37, 0xfa, 0x33,
	0x28, 0x7b, 0xed, 0xeb, 0x9f, 0xbf, 0x3f, 0x4a, 0x2b, 0xf6, 0x12, 0x3d, 0xdb, 0xc8, 0xce, 0x96,
	0xce, 0xa4, 0x68, 0xde, 0xf8, 0x56, 0x6a, 0x6c, 0x59, 0xb5, 0x86, 0x85, 0xbf, 0x59, 0xe8, 0xfe,
	0xd5, 0x36, 0x62, 0x32, 0x5e, 0x60, 0x38, 0x9f, 0xea, 0xf2, 0x2d, 0x11, 0x39, 0xc4, 0x7a, 0x0a,
	0xf1, 0xc4, 0x26, 0x37, 0x40, 0x08, 0x6f, 0x3c, 0xc6, 0xa0, 0xc7, 0xa3, 0x18, 0xd7, 0x87, 0x37,
	0x8a, 0x31, 0x32, 0xa0, 0x5b, 0x30, 0x06, 0x13, 0x29, 0x60, 0xec, 0xfc, 0x2c, 0x7f, 0xdf, 0xfe,
	0x57, 0xc2, 0xbf, 0x2d, 0x34, 0x7f, 0x78, 0x48, 0xda, 0xd2, 0x17, 0x9c, 0xac, 0xee, 0x31, 0xc3,
	0x48, 0x9b, 0x5d, 0x40, 0xb4, 0x66, 0x1f, 0x20, 0xf4, 0x46, 0x41, 0x48, 0x0e, 0x13, 0x51, 0xbc,
	0xd0, 0x35, 0x46, 0xe9, 0x16, 0xa5, 0x09, 0x47, 0x3d, 0x03, 0xf1, 0xe0, 0xac, 0xba, 0x32, 0x5c,
	0xd7, 0x3d, 0xa1, 0x79, 0xac, 0xf5, 0x56, 0x76, 0x71, 0xfd, 0x48, 0xc6, 0x4a, 0x3b, 0x5c, 0xf6,
	0x6a, 0x1f, 0x10, 0xde, 0x56, 0x8c, 0x77, 0x81, 0x34, 0x9d, 0x06, 0x69, 0x0b, 0x0e, 0xc9, 0x89,
	0xda, 0xea, 0x97, 0xf4, 0x85, 0xe9, 0xc6, 0x9d, 0x24, 0x92, 0x66, 0xa9, 0xc7, 0x32, 0xf2, 0x59,
	0x0f, 0x74, 0x41, 0x8c, 0x76, 0x02, 0xd9, 0xa1, 0x3d, 0xa6, 0x0d, 0x44, 0xb4, 0x7d, 0xb0, 0xbb,
	0xff, 0xfa, 0x68, 0xbf, 0x59, 0xde, 0x70, 0x1a, 0xb5, 0x92, 0x55, 0x6a, 0xce, 0x32, 0xa5, 0x02,
	0xc1, 0xd3, 0x3b, 0x4f, 0x3f, 0x69, 0x19, 0xb6, 0x46, 0x2c, 0xee, 0x73, 0x54, 0xde, 0x6c, 0x6c,
	0xe2, 0x4d, 0x54, 0x73, 0xc1, 0xc4, 0x51, 0x08, 0x1e, 0x39, 0xef, 0x42, 0x48, 0x4c, 0x17, 0x48,
	0x04, 0x5a, 0xc6, 0x11, 0x07, 0xe2, 0x49, 0xd0, 0x24, 0x94, 0x86, 0xc0, 0x67, 0xa1, 0x8d, 0x83,
	0x27, 0xd1, 0xc4, 0x65, 0xc9, 0x9a, 0x8a, 0x5e, 0xa2, 0xc5, 0x61, 0x33, 0xc8, 0x9e, 0xe4, 0x71,
	0x0f, 0xc2, 0xec, 0x8d, 0xc1, 0xcb, 0xe3, 0x5b, 0x43, 0xb5, 0x30, 0x40, 0x3d, 0xc9, 0x35, 0xfd,
	0x48, 0xae, 0xb9, 0x0a, 0xfb, 0x52, 0x27, 0x3e, 0x55, 0x9d, 0x5f, 0xa5, 0x4a, 0x52, 0x3f, 0x2d,
	0xdf, 0x99, 0x4c, 0x1f, 0xad, 0x67, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x99, 0x47, 0x45,
	0x32, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	// QueryTickets pages the Tickets by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error)
	// QueryTicketIds gets the list of TicketIDs that meet all the filtering criteria requested by the pool.
	//   - If the Pool contains no Filters, QueryTicketIds will return all TicketIDs in the state storage.
	// QueryTicketIds pages the TicketIDs by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTicketIds(ctx context.Context, in *QueryTicketIdsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketIdsClient, error)
	// QueryBackfills gets a list of Backfills.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	QueryBackfills(ctx context.Context, in *QueryBackfillsRequest, opts ...grpc.CallOption) (QueryService_QueryBackfillsClient, error)
}

type queryServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryServiceClient(cc *grpc.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/openmatch.QueryService/QueryTickets", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryTicketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryTicketsClient interface {
	Recv() (*QueryTicketsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryTicketsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryTicketsClient) Recv() (*QueryTicketsResponse, error) {
	m := new(QueryTicketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) QueryTicketIds(ctx context.Context, in *QueryTicketIdsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[1], "/openmatch.QueryService/QueryTicketIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryTicketIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryTicketIdsClient interface {
	Recv() (*QueryTicketIdsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryTicketIdsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryTicketIdsClient) Recv() (*QueryTicketIdsResponse, error) {
	m := new(QueryTicketIdsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) QueryBackfills(ctx context.Context, in *QueryBackfillsRequest, opts ...grpc.CallOption) (QueryService_QueryBackfillsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[2], "/openmatch.QueryService/QueryBackfills", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryBackfillsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryBackfillsClient interface {
	Recv() (*QueryBackfillsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryBackfillsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryBackfillsClient) Recv() (*QueryBackfillsResponse, error) {
	m := new(QueryBackfillsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	// QueryTickets pages the Tickets by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTickets(*QueryTicketsRequest, QueryService_QueryTicketsServer) error
	// QueryTicketIds gets the list of TicketIDs that meet all the filtering criteria requested by the pool.
	//   - If the Pool contains no Filters, QueryTicketIds will return all TicketIDs in the state storage.
	// QueryTicketIds pages the TicketIDs by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTicketIds(*QueryTicketIdsRequest, QueryService_QueryTicketIdsServer) error
	// QueryBackfills gets a list of Backfills.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	QueryBackfills(*QueryBackfillsRequest, QueryService_QueryBackfillsServer) error
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryTickets(req *QueryTicketsRequest, srv QueryService_QueryTicketsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTickets not implemented")
}
func (*UnimplementedQueryServiceServer) QueryTicketIds(req *QueryTicketIdsRequest, srv QueryService_QueryTicketIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTicketIds not implemented")
}
func (*UnimplementedQueryServiceServer) QueryBackfills(req *QueryBackfillsRequest, srv QueryService_QueryBackfillsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryBackfills not implemented")
}

func RegisterQueryServiceServer(s *grpc.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryTickets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryTicketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryTickets(m, &queryServiceQueryTicketsServer{stream})
}

type QueryService_QueryTicketsServer interface {
	Send(*QueryTicketsResponse) error
	grpc.ServerStream
}

type queryServiceQueryTicketsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryTicketsServer) Send(m *QueryTicketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_QueryTicketIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryTicketIdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryTicketIds(m, &queryServiceQueryTicketIdsServer{stream})
}

type QueryService_QueryTicketIdsServer interface {
	Send(*QueryTicketIdsResponse) error
	grpc.ServerStream
}

type queryServiceQueryTicketIdsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryTicketIdsServer) Send(m *QueryTicketIdsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_QueryBackfills_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryBackfillsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryBackfills(m, &queryServiceQueryBackfillsServer{stream})
}

type QueryService_QueryBackfillsServer interface {
	Send(*QueryBackfillsResponse) error
	grpc.ServerStream
}

type queryServiceQueryBackfillsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryBackfillsServer) Send(m *QueryBackfillsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryTickets",
			Handler:       _QueryService_QueryTickets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryTicketIds",
			Handler:       _QueryService_QueryTicketIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryBackfills",
			Handler:       _QueryService_QueryBackfills_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/query.proto",
}
