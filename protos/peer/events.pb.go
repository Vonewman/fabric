// Code generated by protoc-gen-go.
// source: peer/events.proto
// DO NOT EDIT!

package peer

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

type EventType int32

const (
	EventType_REGISTER  EventType = 0
	EventType_BLOCK     EventType = 1
	EventType_CHAINCODE EventType = 2
	EventType_REJECTION EventType = 3
)

var EventType_name = map[int32]string{
	0: "REGISTER",
	1: "BLOCK",
	2: "CHAINCODE",
	3: "REJECTION",
}
var EventType_value = map[string]int32{
	"REGISTER":  0,
	"BLOCK":     1,
	"CHAINCODE": 2,
	"REJECTION": 3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// ChaincodeReg is used for registering chaincode Interests
// when EventType is CHAINCODE
type ChaincodeReg struct {
	ChaincodeID string `protobuf:"bytes,1,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	EventName   string `protobuf:"bytes,2,opt,name=eventName" json:"eventName,omitempty"`
}

func (m *ChaincodeReg) Reset()                    { *m = ChaincodeReg{} }
func (m *ChaincodeReg) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeReg) ProtoMessage()               {}
func (*ChaincodeReg) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Interest struct {
	EventType EventType `protobuf:"varint,1,opt,name=eventType,enum=protos.EventType" json:"eventType,omitempty"`
	// Ideally we should just have the following oneof for different
	// Reg types and get rid of EventType. But this is an API change
	// Additional Reg types may add messages specific to their type
	// to the oneof.
	//
	// Types that are valid to be assigned to RegInfo:
	//	*Interest_ChaincodeRegInfo
	RegInfo isInterest_RegInfo `protobuf_oneof:"RegInfo"`
}

func (m *Interest) Reset()                    { *m = Interest{} }
func (m *Interest) String() string            { return proto.CompactTextString(m) }
func (*Interest) ProtoMessage()               {}
func (*Interest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type isInterest_RegInfo interface {
	isInterest_RegInfo()
}

type Interest_ChaincodeRegInfo struct {
	ChaincodeRegInfo *ChaincodeReg `protobuf:"bytes,2,opt,name=chaincodeRegInfo,oneof"`
}

func (*Interest_ChaincodeRegInfo) isInterest_RegInfo() {}

func (m *Interest) GetRegInfo() isInterest_RegInfo {
	if m != nil {
		return m.RegInfo
	}
	return nil
}

func (m *Interest) GetChaincodeRegInfo() *ChaincodeReg {
	if x, ok := m.GetRegInfo().(*Interest_ChaincodeRegInfo); ok {
		return x.ChaincodeRegInfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Interest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Interest_OneofMarshaler, _Interest_OneofUnmarshaler, _Interest_OneofSizer, []interface{}{
		(*Interest_ChaincodeRegInfo)(nil),
	}
}

func _Interest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Interest)
	// RegInfo
	switch x := m.RegInfo.(type) {
	case *Interest_ChaincodeRegInfo:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChaincodeRegInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Interest.RegInfo has unexpected type %T", x)
	}
	return nil
}

func _Interest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Interest)
	switch tag {
	case 2: // RegInfo.chaincodeRegInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChaincodeReg)
		err := b.DecodeMessage(msg)
		m.RegInfo = &Interest_ChaincodeRegInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Interest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Interest)
	// RegInfo
	switch x := m.RegInfo.(type) {
	case *Interest_ChaincodeRegInfo:
		s := proto.Size(x.ChaincodeRegInfo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ---------- consumer events ---------
// Register is sent by consumers for registering events
// string type - "register"
type Register struct {
	Events []*Interest `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *Register) Reset()                    { *m = Register{} }
func (m *Register) String() string            { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()               {}
func (*Register) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Register) GetEvents() []*Interest {
	if m != nil {
		return m.Events
	}
	return nil
}

// Rejection is sent by consumers for erroneous transaction rejection events
// string type - "rejection"
type Rejection struct {
	Tx       *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	ErrorMsg string       `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *Rejection) Reset()                    { *m = Rejection{} }
func (m *Rejection) String() string            { return proto.CompactTextString(m) }
func (*Rejection) ProtoMessage()               {}
func (*Rejection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Rejection) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// ---------- producer events ---------
type Unregister struct {
	Events []*Interest `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *Unregister) Reset()                    { *m = Unregister{} }
func (m *Unregister) String() string            { return proto.CompactTextString(m) }
func (*Unregister) ProtoMessage()               {}
func (*Unregister) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Unregister) GetEvents() []*Interest {
	if m != nil {
		return m.Events
	}
	return nil
}

// Event is used by
//  - consumers (adapters) to send Register
//  - producer to advertise supported types and events
type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_Register
	//	*Event_Block
	//	*Event_ChaincodeEvent
	//	*Event_Rejection
	//	*Event_Unregister
	Event isEvent_Event `protobuf_oneof:"Event"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Register struct {
	Register *Register `protobuf:"bytes,1,opt,name=register,oneof"`
}
type Event_Block struct {
	Block *Block `protobuf:"bytes,2,opt,name=block,oneof"`
}
type Event_ChaincodeEvent struct {
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,3,opt,name=chaincodeEvent,oneof"`
}
type Event_Rejection struct {
	Rejection *Rejection `protobuf:"bytes,4,opt,name=rejection,oneof"`
}
type Event_Unregister struct {
	Unregister *Unregister `protobuf:"bytes,5,opt,name=unregister,oneof"`
}

func (*Event_Register) isEvent_Event()       {}
func (*Event_Block) isEvent_Event()          {}
func (*Event_ChaincodeEvent) isEvent_Event() {}
func (*Event_Rejection) isEvent_Event()      {}
func (*Event_Unregister) isEvent_Event()     {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetRegister() *Register {
	if x, ok := m.GetEvent().(*Event_Register); ok {
		return x.Register
	}
	return nil
}

func (m *Event) GetBlock() *Block {
	if x, ok := m.GetEvent().(*Event_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Event) GetChaincodeEvent() *ChaincodeEvent {
	if x, ok := m.GetEvent().(*Event_ChaincodeEvent); ok {
		return x.ChaincodeEvent
	}
	return nil
}

func (m *Event) GetRejection() *Rejection {
	if x, ok := m.GetEvent().(*Event_Rejection); ok {
		return x.Rejection
	}
	return nil
}

func (m *Event) GetUnregister() *Unregister {
	if x, ok := m.GetEvent().(*Event_Unregister); ok {
		return x.Unregister
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Register)(nil),
		(*Event_Block)(nil),
		(*Event_ChaincodeEvent)(nil),
		(*Event_Rejection)(nil),
		(*Event_Unregister)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// Event
	switch x := m.Event.(type) {
	case *Event_Register:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Register); err != nil {
			return err
		}
	case *Event_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *Event_ChaincodeEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChaincodeEvent); err != nil {
			return err
		}
	case *Event_Rejection:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rejection); err != nil {
			return err
		}
	case *Event_Unregister:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Unregister); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Event has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // Event.register
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Register)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Register{msg}
		return true, err
	case 2: // Event.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Block{msg}
		return true, err
	case 3: // Event.chaincodeEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChaincodeEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_ChaincodeEvent{msg}
		return true, err
	case 4: // Event.rejection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rejection)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Rejection{msg}
		return true, err
	case 5: // Event.unregister
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Unregister)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Unregister{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// Event
	switch x := m.Event.(type) {
	case *Event_Register:
		s := proto.Size(x.Register)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ChaincodeEvent:
		s := proto.Size(x.ChaincodeEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Rejection:
		s := proto.Size(x.Rejection)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Unregister:
		s := proto.Size(x.Unregister)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ChaincodeReg)(nil), "protos.ChaincodeReg")
	proto.RegisterType((*Interest)(nil), "protos.Interest")
	proto.RegisterType((*Register)(nil), "protos.Register")
	proto.RegisterType((*Rejection)(nil), "protos.Rejection")
	proto.RegisterType((*Unregister)(nil), "protos.Unregister")
	proto.RegisterType((*Event)(nil), "protos.Event")
	proto.RegisterEnum("protos.EventType", EventType_name, EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Events service

type EventsClient interface {
	// event chatting using Event
	Chat(ctx context.Context, opts ...grpc.CallOption) (Events_ChatClient, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Events_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Events_serviceDesc.Streams[0], c.cc, "/protos.Events/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsChatClient{stream}
	return x, nil
}

type Events_ChatClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventsChatClient struct {
	grpc.ClientStream
}

func (x *eventsChatClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventsChatClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Events service

type EventsServer interface {
	// event chatting using Event
	Chat(Events_ChatServer) error
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventsServer).Chat(&eventsChatServer{stream})
}

type Events_ChatServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventsChatServer struct {
	grpc.ServerStream
}

func (x *eventsChatServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventsChatServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Events_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor4,
}

func init() { proto.RegisterFile("peer/events.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8f, 0x93, 0x40,
	0x18, 0xc5, 0x81, 0x6e, 0xbb, 0xf0, 0x75, 0xbb, 0xa1, 0x9f, 0xc6, 0x60, 0xe3, 0xa1, 0xc1, 0x98,
	0xd4, 0x35, 0x29, 0x8a, 0x8d, 0x67, 0x85, 0x25, 0x82, 0xae, 0x6d, 0x32, 0xd6, 0x8b, 0x37, 0xca,
	0xce, 0x52, 0x74, 0x17, 0x9a, 0x61, 0xd6, 0xec, 0xfe, 0x0b, 0x1e, 0xfd, 0x8b, 0x4d, 0x07, 0x06,
	0x58, 0x7b, 0xf2, 0xd4, 0xce, 0xbc, 0xdf, 0xfb, 0xe6, 0xf1, 0x32, 0x03, 0xe3, 0x1d, 0xa5, 0xcc,
	0xa1, 0xbf, 0x68, 0xce, 0xcb, 0xf9, 0x8e, 0x15, 0xbc, 0xc0, 0x81, 0xf8, 0x29, 0x27, 0x4f, 0x85,
	0x94, 0x6c, 0xe3, 0x2c, 0x4f, 0x8a, 0x4b, 0x2a, 0x98, 0x0a, 0x99, 0x54, 0xae, 0xab, 0x78, 0xc3,
	0xb2, 0xa4, 0xda, 0xb2, 0x97, 0x70, 0xe2, 0x4b, 0x94, 0xd0, 0x14, 0xa7, 0x30, 0x6c, 0xac, 0xd1,
	0xb9, 0xa5, 0x4e, 0xd5, 0x99, 0x41, 0xba, 0x5b, 0xf8, 0x0c, 0x0c, 0x31, 0x73, 0x19, 0xdf, 0x50,
	0x4b, 0x13, 0x7a, 0xbb, 0x61, 0xff, 0x56, 0x41, 0x8f, 0x72, 0x4e, 0x19, 0x2d, 0x39, 0x3a, 0x35,
	0xba, 0xbe, 0xdf, 0x51, 0x31, 0xea, 0xd4, 0x1d, 0x57, 0xe7, 0x96, 0xf3, 0x40, 0x0a, 0xa4, 0x65,
	0xd0, 0x03, 0x33, 0xe9, 0xa4, 0x89, 0xf2, 0xab, 0x42, 0x1c, 0x31, 0x74, 0x1f, 0x4b, 0x5f, 0x37,
	0x6d, 0xa8, 0x90, 0x03, 0xde, 0x33, 0xe0, 0xb8, 0xfe, 0x6b, 0x2f, 0x40, 0x27, 0x34, 0xcd, 0x4a,
	0x4e, 0x19, 0xce, 0x60, 0x50, 0xd5, 0x65, 0xa9, 0xd3, 0xde, 0x6c, 0xe8, 0x9a, 0x72, 0xa0, 0x4c,
	0x4b, 0x6a, 0xdd, 0xbe, 0x00, 0x83, 0xd0, 0x1f, 0x34, 0xe1, 0x59, 0x91, 0xe3, 0x73, 0xd0, 0xf8,
	0x9d, 0xc8, 0x3e, 0x74, 0x1f, 0x49, 0xcb, 0x9a, 0xc5, 0x79, 0x19, 0x0b, 0x80, 0x68, 0xfc, 0x0e,
	0x27, 0xa0, 0x53, 0xc6, 0x0a, 0xf6, 0xa5, 0x4c, 0xeb, 0x46, 0x9a, 0xb5, 0xfd, 0x0e, 0xe0, 0x5b,
	0xce, 0xfe, 0x3f, 0xc5, 0x1f, 0x0d, 0xfa, 0xa2, 0x23, 0x9c, 0x83, 0x2e, 0xfd, 0x75, 0x90, 0xc6,
	0x25, 0xbf, 0x2e, 0x54, 0x48, 0xc3, 0xe0, 0x0b, 0xe8, 0x6f, 0xae, 0x8b, 0xe4, 0x67, 0xdd, 0xdc,
	0x48, 0xc2, 0xde, 0x7e, 0x33, 0x54, 0x48, 0xa5, 0xe2, 0x7b, 0x38, 0x6d, 0xba, 0x13, 0x07, 0x59,
	0x3d, 0xc1, 0x3f, 0x39, 0x68, 0x5a, 0xa8, 0xa1, 0x42, 0xfe, 0xe1, 0xf1, 0x0d, 0x18, 0x4c, 0x16,
	0x65, 0x1d, 0x09, 0xf3, 0xb8, 0x4d, 0x56, 0x0b, 0xa1, 0x42, 0x5a, 0x0a, 0x17, 0x00, 0xb7, 0x4d,
	0x1b, 0x56, 0x5f, 0x78, 0x50, 0x7a, 0xda, 0x9e, 0x42, 0x85, 0x74, 0x38, 0xef, 0xb8, 0xae, 0xe2,
	0xcc, 0x03, 0xa3, 0xb9, 0x37, 0x78, 0x02, 0x3a, 0x09, 0x3e, 0x46, 0x5f, 0xd7, 0x01, 0x31, 0x15,
	0x34, 0xa0, 0xef, 0x5d, 0xac, 0xfc, 0xcf, 0xa6, 0x8a, 0x23, 0x30, 0xfc, 0xf0, 0x43, 0xb4, 0xf4,
	0x57, 0xe7, 0x81, 0xa9, 0xed, 0x97, 0x24, 0xf8, 0x14, 0xf8, 0xeb, 0x68, 0xb5, 0x34, 0x7b, 0xee,
	0x02, 0x06, 0x62, 0x46, 0x89, 0x67, 0x70, 0xe4, 0x6f, 0x63, 0x8e, 0xa3, 0x07, 0x77, 0x72, 0xf2,
	0x70, 0x69, 0x2b, 0x33, 0xf5, 0xb5, 0xea, 0xbd, 0xfa, 0xfe, 0x32, 0xcd, 0xf8, 0xf6, 0x76, 0x33,
	0x4f, 0x8a, 0x1b, 0x67, 0x7b, 0xbf, 0xa3, 0xec, 0x9a, 0x5e, 0xa6, 0xcd, 0x73, 0x72, 0x2a, 0x8f,
	0xb3, 0x7f, 0x61, 0x9b, 0xea, 0x29, 0xbe, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xea, 0x3a, 0x18,
	0x18, 0xa6, 0x03, 0x00, 0x00,
}
