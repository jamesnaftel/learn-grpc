// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Podcast
	PodcastRequest
	PodcastResponse
	PodcastsResponse
	Empty
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Podcast struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Length int32  `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
}

func (m *Podcast) Reset()                    { *m = Podcast{} }
func (m *Podcast) String() string            { return proto.CompactTextString(m) }
func (*Podcast) ProtoMessage()               {}
func (*Podcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Podcast) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Podcast) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Podcast) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type PodcastRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PodcastRequest) Reset()                    { *m = PodcastRequest{} }
func (m *PodcastRequest) String() string            { return proto.CompactTextString(m) }
func (*PodcastRequest) ProtoMessage()               {}
func (*PodcastRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PodcastRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PodcastResponse struct {
	Podcast *Podcast `protobuf:"bytes,1,opt,name=podcast" json:"podcast,omitempty"`
}

func (m *PodcastResponse) Reset()                    { *m = PodcastResponse{} }
func (m *PodcastResponse) String() string            { return proto.CompactTextString(m) }
func (*PodcastResponse) ProtoMessage()               {}
func (*PodcastResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PodcastResponse) GetPodcast() *Podcast {
	if m != nil {
		return m.Podcast
	}
	return nil
}

type PodcastsResponse struct {
	Podcasts []*Podcast `protobuf:"bytes,1,rep,name=podcasts" json:"podcasts,omitempty"`
}

func (m *PodcastsResponse) Reset()                    { *m = PodcastsResponse{} }
func (m *PodcastsResponse) String() string            { return proto.CompactTextString(m) }
func (*PodcastsResponse) ProtoMessage()               {}
func (*PodcastsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PodcastsResponse) GetPodcasts() []*Podcast {
	if m != nil {
		return m.Podcasts
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Podcast)(nil), "api.Podcast")
	proto.RegisterType((*PodcastRequest)(nil), "api.PodcastRequest")
	proto.RegisterType((*PodcastResponse)(nil), "api.PodcastResponse")
	proto.RegisterType((*PodcastsResponse)(nil), "api.PodcastsResponse")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Podcasts service

type PodcastsClient interface {
	GetPodcast(ctx context.Context, in *PodcastRequest, opts ...grpc.CallOption) (*PodcastResponse, error)
	GetPodcasts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodcastsResponse, error)
}

type podcastsClient struct {
	cc *grpc.ClientConn
}

func NewPodcastsClient(cc *grpc.ClientConn) PodcastsClient {
	return &podcastsClient{cc}
}

func (c *podcastsClient) GetPodcast(ctx context.Context, in *PodcastRequest, opts ...grpc.CallOption) (*PodcastResponse, error) {
	out := new(PodcastResponse)
	err := grpc.Invoke(ctx, "/api.Podcasts/GetPodcast", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podcastsClient) GetPodcasts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodcastsResponse, error) {
	out := new(PodcastsResponse)
	err := grpc.Invoke(ctx, "/api.Podcasts/GetPodcasts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Podcasts service

type PodcastsServer interface {
	GetPodcast(context.Context, *PodcastRequest) (*PodcastResponse, error)
	GetPodcasts(context.Context, *Empty) (*PodcastsResponse, error)
}

func RegisterPodcastsServer(s *grpc.Server, srv PodcastsServer) {
	s.RegisterService(&_Podcasts_serviceDesc, srv)
}

func _Podcasts_GetPodcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastsServer).GetPodcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Podcasts/GetPodcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastsServer).GetPodcast(ctx, req.(*PodcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Podcasts_GetPodcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodcastsServer).GetPodcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Podcasts/GetPodcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodcastsServer).GetPodcasts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Podcasts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Podcasts",
	HandlerType: (*PodcastsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPodcast",
			Handler:    _Podcasts_GetPodcast_Handler,
		},
		{
			MethodName: "GetPodcasts",
			Handler:    _Podcasts_GetPodcasts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xf2, 0xe5, 0x62, 0x0f, 0xc8,
	0x4f, 0x49, 0x4e, 0x2c, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x12, 0x4b, 0x4b, 0x32, 0xf2, 0x8b,
	0x24, 0x98, 0xc0, 0xa2, 0x50, 0x1e, 0x48, 0x3c, 0x27, 0x35, 0x2f, 0xbd, 0x24, 0x43, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0x52, 0xe1, 0xe2, 0x83, 0x1a, 0x17, 0x94, 0x5a, 0x58,
	0x9a, 0x8a, 0xdd, 0x54, 0x25, 0x4b, 0x2e, 0x7e, 0xb8, 0xaa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x35, 0x2e, 0xf6, 0x02, 0x88, 0x10, 0x58, 0x25, 0xb7, 0x11, 0x8f, 0x1e, 0xc8, 0xa5, 0x30,
	0x65, 0x30, 0x49, 0x25, 0x1b, 0x2e, 0x01, 0xa8, 0x58, 0x31, 0x5c, 0xaf, 0x06, 0x17, 0x07, 0x54,
	0xba, 0x58, 0x82, 0x51, 0x81, 0x19, 0x43, 0x33, 0x5c, 0x56, 0x89, 0x9d, 0x8b, 0xd5, 0x35, 0xb7,
	0xa0, 0xa4, 0xd2, 0xa8, 0x94, 0x8b, 0x03, 0x66, 0x8c, 0x90, 0x39, 0x17, 0x97, 0x7b, 0x6a, 0x09,
	0x2c, 0x14, 0x84, 0x51, 0xb4, 0x42, 0x3c, 0x21, 0x25, 0x82, 0x2a, 0x08, 0xb5, 0xd7, 0x80, 0x8b,
	0x1b, 0xa1, 0xb1, 0x58, 0x88, 0x0b, 0xac, 0x08, 0x6c, 0xbe, 0x94, 0x28, 0xb2, 0x06, 0xb8, 0x4b,
	0x93, 0xd8, 0xc0, 0x21, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x04, 0x80, 0x38, 0x86,
	0x01, 0x00, 0x00,
}
