// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Empty
	Podcast
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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Podcast struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Length int32  `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
}

func (m *Podcast) Reset()                    { *m = Podcast{} }
func (m *Podcast) String() string            { return proto.CompactTextString(m) }
func (*Podcast) ProtoMessage()               {}
func (*Podcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*Podcast)(nil), "api.Podcast")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Podcasts service

type PodcastsClient interface {
	GetPodcasts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Podcast, error)
}

type podcastsClient struct {
	cc *grpc.ClientConn
}

func NewPodcastsClient(cc *grpc.ClientConn) PodcastsClient {
	return &podcastsClient{cc}
}

func (c *podcastsClient) GetPodcasts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Podcast, error) {
	out := new(Podcast)
	err := grpc.Invoke(ctx, "/api.Podcasts/GetPodcasts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Podcasts service

type PodcastsServer interface {
	GetPodcasts(context.Context, *Empty) (*Podcast, error)
}

func RegisterPodcastsServer(s *grpc.Server, srv PodcastsServer) {
	s.RegisterService(&_Podcasts_serviceDesc, srv)
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
			MethodName: "GetPodcasts",
			Handler:    _Podcasts_GetPodcasts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x62, 0xe7, 0x62, 0x75, 0xcd,
	0x2d, 0x28, 0xa9, 0x54, 0xf2, 0xe5, 0x62, 0x0f, 0xc8, 0x4f, 0x49, 0x4e, 0x2c, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85,
	0xc4, 0xb8, 0xd8, 0x12, 0x4b, 0x4b, 0x32, 0xf2, 0x8b, 0x24, 0x98, 0xc0, 0xa2, 0x50, 0x1e, 0x48,
	0x3c, 0x27, 0x35, 0x2f, 0xbd, 0x24, 0x43, 0x82, 0x59, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x33,
	0x32, 0xe6, 0xe2, 0x80, 0x1a, 0x57, 0x2c, 0xa4, 0xce, 0xc5, 0xed, 0x9e, 0x5a, 0x02, 0xe7, 0x72,
	0xe9, 0x81, 0xdc, 0x00, 0xb6, 0x55, 0x8a, 0x07, 0xcc, 0x86, 0x4a, 0x25, 0xb1, 0x81, 0x1d, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x25, 0xa4, 0x63, 0xa5, 0x00, 0x00, 0x00,
}
