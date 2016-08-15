// Code generated by protoc-gen-gogo.
// source: cockroach/storage/api.proto
// DO NOT EDIT!

/*
	Package storage is a generated protocol buffer package.

	It is generated from these files:
		cockroach/storage/api.proto
		cockroach/storage/raft.proto

	It has these top-level messages:
		StoreRequestHeader
		PollFrozenRequest
		PollFrozenResponse
		ReservationRequest
		ReservationResponse
		RaftMessageRequest
		RaftMessageResponseUnion
		RaftMessageResponse
		ConfChangeContext
*/
package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_cockroachdb_cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// StoreRequestHeader locates a Store on a Node.
type StoreRequestHeader struct {
	NodeID  github_com_cockroachdb_cockroach_roachpb.NodeID  `protobuf:"varint,1,opt,name=node_id,json=nodeId,casttype=github.com/cockroachdb/cockroach/roachpb.NodeID" json:"node_id"`
	StoreID github_com_cockroachdb_cockroach_roachpb.StoreID `protobuf:"varint,2,opt,name=store_id,json=storeId,casttype=github.com/cockroachdb/cockroach/roachpb.StoreID" json:"store_id"`
}

func (m *StoreRequestHeader) Reset()                    { *m = StoreRequestHeader{} }
func (m *StoreRequestHeader) String() string            { return proto.CompactTextString(m) }
func (*StoreRequestHeader) ProtoMessage()               {}
func (*StoreRequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

// A PollFrozenRequest asks the addressed Store for its frozen or thawed
// Replicas.
type PollFrozenRequest struct {
	StoreRequestHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	// When true, collect the frozen Replicas, and the thawed ones otherwise.
	CollectFrozen bool `protobuf:"varint,2,opt,name=collect_frozen,json=collectFrozen" json:"collect_frozen"`
}

func (m *PollFrozenRequest) Reset()                    { *m = PollFrozenRequest{} }
func (m *PollFrozenRequest) String() string            { return proto.CompactTextString(m) }
func (*PollFrozenRequest) ProtoMessage()               {}
func (*PollFrozenRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

// A PollFrozenResponse is the response returned from a PollFrozenRequest.
type PollFrozenResponse struct {
	Results []cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,1,rep,name=results" json:"results"`
}

func (m *PollFrozenResponse) Reset()                    { *m = PollFrozenResponse{} }
func (m *PollFrozenResponse) String() string            { return proto.CompactTextString(m) }
func (*PollFrozenResponse) ProtoMessage()               {}
func (*PollFrozenResponse) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

// A ReservationRequest asks the addressed Store to reserve the space for a new
// replica for the Range of RangeID reserving RangeSize bytes.
type ReservationRequest struct {
	StoreRequestHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	FromNodeID         github_com_cockroachdb_cockroach_roachpb.NodeID  `protobuf:"varint,2,opt,name=from_node_id,json=fromNodeId,casttype=github.com/cockroachdb/cockroach/roachpb.NodeID" json:"from_node_id"`
	FromStoreID        github_com_cockroachdb_cockroach_roachpb.StoreID `protobuf:"varint,3,opt,name=from_store_id,json=fromStoreId,casttype=github.com/cockroachdb/cockroach/roachpb.StoreID" json:"from_store_id"`
	RangeID            github_com_cockroachdb_cockroach_roachpb.RangeID `protobuf:"varint,4,opt,name=range_id,json=rangeId,casttype=github.com/cockroachdb/cockroach/roachpb.RangeID" json:"range_id"`
	RangeSize          int64                                            `protobuf:"varint,5,opt,name=range_size,json=rangeSize" json:"range_size"`
}

func (m *ReservationRequest) Reset()                    { *m = ReservationRequest{} }
func (m *ReservationRequest) String() string            { return proto.CompactTextString(m) }
func (*ReservationRequest) ProtoMessage()               {}
func (*ReservationRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{3} }

// A ReservationResponse is the repoonse returned from a ReservationRequest.
type ReservationResponse struct {
	Reserved bool `protobuf:"varint,1,opt,name=reserved" json:"reserved"`
	// The current number of ranges and reservations on the target
	// store. This is returned even if the reservation isn't filled.
	RangeCount int32 `protobuf:"varint,2,opt,name=range_count,json=rangeCount" json:"range_count"`
}

func (m *ReservationResponse) Reset()                    { *m = ReservationResponse{} }
func (m *ReservationResponse) String() string            { return proto.CompactTextString(m) }
func (*ReservationResponse) ProtoMessage()               {}
func (*ReservationResponse) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{4} }

func init() {
	proto.RegisterType((*StoreRequestHeader)(nil), "cockroach.storage.StoreRequestHeader")
	proto.RegisterType((*PollFrozenRequest)(nil), "cockroach.storage.PollFrozenRequest")
	proto.RegisterType((*PollFrozenResponse)(nil), "cockroach.storage.PollFrozenResponse")
	proto.RegisterType((*ReservationRequest)(nil), "cockroach.storage.ReservationRequest")
	proto.RegisterType((*ReservationResponse)(nil), "cockroach.storage.ReservationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Stores service

type StoresClient interface {
	PollFrozen(ctx context.Context, in *PollFrozenRequest, opts ...grpc.CallOption) (*PollFrozenResponse, error)
	Reserve(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
}

type storesClient struct {
	cc *grpc.ClientConn
}

func NewStoresClient(cc *grpc.ClientConn) StoresClient {
	return &storesClient{cc}
}

func (c *storesClient) PollFrozen(ctx context.Context, in *PollFrozenRequest, opts ...grpc.CallOption) (*PollFrozenResponse, error) {
	out := new(PollFrozenResponse)
	err := grpc.Invoke(ctx, "/cockroach.storage.Stores/PollFrozen", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesClient) Reserve(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := grpc.Invoke(ctx, "/cockroach.storage.Stores/Reserve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stores service

type StoresServer interface {
	PollFrozen(context.Context, *PollFrozenRequest) (*PollFrozenResponse, error)
	Reserve(context.Context, *ReservationRequest) (*ReservationResponse, error)
}

func RegisterStoresServer(s *grpc.Server, srv StoresServer) {
	s.RegisterService(&_Stores_serviceDesc, srv)
}

func _Stores_PollFrozen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollFrozenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).PollFrozen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.storage.Stores/PollFrozen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).PollFrozen(ctx, req.(*PollFrozenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stores_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.storage.Stores/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).Reserve(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stores_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.storage.Stores",
	HandlerType: (*StoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PollFrozen",
			Handler:    _Stores_PollFrozen_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _Stores_Reserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorApi,
}

func (m *StoreRequestHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreRequestHeader) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintApi(data, i, uint64(m.NodeID))
	data[i] = 0x10
	i++
	i = encodeVarintApi(data, i, uint64(m.StoreID))
	return i, nil
}

func (m *PollFrozenRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PollFrozenRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintApi(data, i, uint64(m.StoreRequestHeader.Size()))
	n1, err := m.StoreRequestHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x10
	i++
	if m.CollectFrozen {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *PollFrozenResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PollFrozenResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0xa
			i++
			i = encodeVarintApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ReservationRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReservationRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintApi(data, i, uint64(m.StoreRequestHeader.Size()))
	n2, err := m.StoreRequestHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x10
	i++
	i = encodeVarintApi(data, i, uint64(m.FromNodeID))
	data[i] = 0x18
	i++
	i = encodeVarintApi(data, i, uint64(m.FromStoreID))
	data[i] = 0x20
	i++
	i = encodeVarintApi(data, i, uint64(m.RangeID))
	data[i] = 0x28
	i++
	i = encodeVarintApi(data, i, uint64(m.RangeSize))
	return i, nil
}

func (m *ReservationResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReservationResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	if m.Reserved {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x10
	i++
	i = encodeVarintApi(data, i, uint64(m.RangeCount))
	return i, nil
}

func encodeFixed64Api(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Api(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintApi(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StoreRequestHeader) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovApi(uint64(m.NodeID))
	n += 1 + sovApi(uint64(m.StoreID))
	return n
}

func (m *PollFrozenRequest) Size() (n int) {
	var l int
	_ = l
	l = m.StoreRequestHeader.Size()
	n += 1 + l + sovApi(uint64(l))
	n += 2
	return n
}

func (m *PollFrozenResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *ReservationRequest) Size() (n int) {
	var l int
	_ = l
	l = m.StoreRequestHeader.Size()
	n += 1 + l + sovApi(uint64(l))
	n += 1 + sovApi(uint64(m.FromNodeID))
	n += 1 + sovApi(uint64(m.FromStoreID))
	n += 1 + sovApi(uint64(m.RangeID))
	n += 1 + sovApi(uint64(m.RangeSize))
	return n
}

func (m *ReservationResponse) Size() (n int) {
	var l int
	_ = l
	n += 2
	n += 1 + sovApi(uint64(m.RangeCount))
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreRequestHeader) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreRequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreRequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StoreID |= (github_com_cockroachdb_cockroach_roachpb.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PollFrozenRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PollFrozenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollFrozenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreRequestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoreRequestHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectFrozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CollectFrozen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PollFrozenResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PollFrozenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollFrozenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, cockroach_roachpb.ReplicaDescriptor{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReservationRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReservationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReservationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreRequestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoreRequestHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromNodeID", wireType)
			}
			m.FromNodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FromNodeID |= (github_com_cockroachdb_cockroach_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromStoreID", wireType)
			}
			m.FromStoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FromStoreID |= (github_com_cockroachdb_cockroach_roachpb.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeID |= (github_com_cockroachdb_cockroach_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeSize", wireType)
			}
			m.RangeSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReservationResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReservationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReservationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reserved = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeCount", wireType)
			}
			m.RangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/storage/api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x65, 0x3e, 0x12, 0xcc, 0x77, 0x69, 0x2a, 0x65, 0xda, 0x05, 0xa2, 0x92, 0xa1, 0x6e, 0xa8,
	0x90, 0x2a, 0xd9, 0x15, 0x6f, 0x50, 0x42, 0x93, 0xb2, 0xa9, 0x2a, 0x67, 0x53, 0x51, 0x29, 0xc8,
	0xd8, 0x17, 0x63, 0xc5, 0x30, 0xee, 0x78, 0xe8, 0x82, 0x37, 0xe8, 0xae, 0x8f, 0xd3, 0x6d, 0x77,
	0x2c, 0x59, 0x76, 0x85, 0x5a, 0xfa, 0x16, 0x5d, 0x55, 0x9e, 0x19, 0x43, 0x2a, 0x47, 0x4a, 0xff,
	0x36, 0xc8, 0xba, 0x73, 0xcf, 0x39, 0xf7, 0xcc, 0x3d, 0x03, 0x3c, 0xf0, 0x99, 0x7f, 0xc5, 0x99,
	0xe7, 0x4f, 0x9d, 0x54, 0x30, 0xee, 0x85, 0xe8, 0x78, 0x49, 0x64, 0x27, 0x9c, 0x09, 0x46, 0x8f,
	0x77, 0x87, 0xb6, 0x3e, 0x6c, 0xb4, 0xf6, 0xfd, 0xf2, 0x37, 0x19, 0x3b, 0x33, 0x14, 0x5e, 0xe0,
	0x09, 0x4f, 0x81, 0x1a, 0xf7, 0x43, 0x16, 0x32, 0xf9, 0xe9, 0x64, 0x5f, 0xaa, 0x6a, 0xad, 0x09,
	0xd0, 0x0b, 0xc1, 0x38, 0xba, 0xf8, 0x76, 0x81, 0xa9, 0x78, 0x81, 0x5e, 0x80, 0x9c, 0x0e, 0xc1,
	0x98, 0xb3, 0x00, 0x47, 0x51, 0x50, 0x27, 0x2d, 0xd2, 0x39, 0xec, 0x3d, 0x5b, 0x6d, 0x9a, 0xa5,
	0xed, 0xa6, 0x59, 0x79, 0xc9, 0x02, 0x1c, 0xf4, 0xbf, 0x6f, 0x9a, 0x4e, 0x18, 0x89, 0xe9, 0x62,
	0x6c, 0xfb, 0x6c, 0xe6, 0xec, 0xc4, 0x83, 0xb1, 0x53, 0x18, 0xc4, 0x56, 0x10, 0xb7, 0x92, 0x31,
	0x0e, 0x02, 0x7a, 0x09, 0xd5, 0x6c, 0x6a, 0x49, 0xfe, 0x9f, 0x24, 0x3f, 0xd5, 0xe4, 0x86, 0x9c,
	0x44, 0xb2, 0x3f, 0xfd, 0x65, 0x76, 0x8d, 0x71, 0x0d, 0x49, 0x3a, 0x08, 0xac, 0xf7, 0x04, 0x8e,
	0x5f, 0xb1, 0x38, 0x3e, 0xe3, 0x6c, 0x89, 0x73, 0xed, 0x8b, 0x9e, 0x43, 0x65, 0x2a, 0xbd, 0x49,
	0x43, 0xb5, 0x6e, 0xdb, 0x2e, 0x5c, 0xa2, 0x5d, 0xbc, 0x88, 0x5e, 0x35, 0x1b, 0x6d, 0xbd, 0x69,
	0x12, 0x57, 0xc3, 0xe9, 0x13, 0xb8, 0xeb, 0xb3, 0x38, 0x46, 0x5f, 0x8c, 0x26, 0x52, 0x41, 0x9a,
	0xa8, 0xf6, 0x0e, 0xb2, 0x4e, 0xf7, 0x48, 0x9f, 0x29, 0x71, 0x6b, 0x08, 0xf4, 0xfa, 0x28, 0x69,
	0xc2, 0xe6, 0x29, 0xd2, 0x3e, 0x18, 0x1c, 0xd3, 0x45, 0x2c, 0xd2, 0x3a, 0x69, 0x95, 0x3b, 0xb5,
	0xee, 0xc9, 0xb5, 0x61, 0x72, 0x5f, 0x2e, 0x26, 0x71, 0xe4, 0x7b, 0x7d, 0x4c, 0x7d, 0x1e, 0x25,
	0x82, 0x71, 0xad, 0x90, 0x43, 0xad, 0x8f, 0x65, 0xa0, 0x2e, 0xa6, 0xc8, 0xdf, 0x79, 0x22, 0x62,
	0xff, 0xde, 0x68, 0x08, 0x77, 0x26, 0x9c, 0xcd, 0x46, 0x79, 0x10, 0xd4, 0xae, 0x9e, 0xeb, 0x5d,
	0xc1, 0x19, 0x67, 0xb3, 0x3f, 0x0f, 0x03, 0x4c, 0x72, 0x78, 0x40, 0xaf, 0xe0, 0x48, 0x0a, 0xed,
	0x52, 0x51, 0x96, 0x4a, 0xe7, 0x5a, 0xa9, 0x96, 0x29, 0xfd, 0x4d, 0x32, 0x6a, 0x93, 0x1d, 0x81,
	0x4c, 0x1f, 0xf7, 0xe6, 0xa1, 0xd4, 0x39, 0x68, 0x91, 0x4e, 0x79, 0x9f, 0x3e, 0x37, 0xab, 0xff,
	0xa6, 0x86, 0xc6, 0xb8, 0x86, 0x24, 0x1d, 0x04, 0xf4, 0x11, 0x80, 0xe2, 0x4f, 0xa3, 0x25, 0xd6,
	0x0f, 0xa5, 0x82, 0x5a, 0xdc, 0xff, 0xb2, 0x7e, 0x11, 0x2d, 0xd1, 0xba, 0x84, 0x7b, 0x3f, 0x6d,
	0x4e, 0xe7, 0xa2, 0x05, 0x55, 0x2e, 0xcb, 0xa8, 0x9e, 0x5d, 0x1e, 0xaa, 0x5d, 0x95, 0xb6, 0xa1,
	0xa6, 0xd8, 0x7d, 0xb6, 0x98, 0x0b, 0xbd, 0x12, 0xd5, 0xa4, 0x64, 0x4f, 0xb3, 0x7a, 0xf7, 0x13,
	0x81, 0x8a, 0x34, 0x9c, 0xd2, 0x37, 0x00, 0xfb, 0x04, 0xd2, 0x93, 0x1b, 0xc2, 0x50, 0x78, 0x2b,
	0x8d, 0xf6, 0x2d, 0x5d, 0x6a, 0x5c, 0xab, 0x94, 0xfd, 0x4d, 0x28, 0x1f, 0x48, 0x6f, 0xc2, 0x14,
	0xd3, 0xd9, 0x78, 0x7c, 0x5b, 0x5b, 0xce, 0xdd, 0x7b, 0xb8, 0xfa, 0x6a, 0x96, 0x56, 0x5b, 0x93,
	0xac, 0xb7, 0x26, 0xf9, 0xbc, 0x35, 0xc9, 0x97, 0xad, 0x49, 0x3e, 0x7c, 0x33, 0x4b, 0x43, 0x43,
	0x03, 0x5f, 0x93, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0x2c, 0x3f, 0xaf, 0x27, 0x05, 0x00,
	0x00,
}
