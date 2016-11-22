// Code generated by protoc-gen-go.
// source: peer/chaincodeevent.proto
// DO NOT EDIT!

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	peer/chaincodeevent.proto
	peer/chaincode_proposal.proto
	peer/chaincode.proto
	peer/chaincode_transaction.proto
	peer/events.proto
	peer/fabric_block.proto
	peer/fabric_message.proto
	peer/fabric_proposal.proto
	peer/fabric_proposal_response.proto
	peer/fabric.proto
	peer/fabric_service.proto
	peer/fabric_transaction.proto
	peer/server_admin.proto

It has these top-level messages:
	ChaincodeEvent
	ChaincodeHeaderExtension
	ChaincodeProposalPayload
	ChaincodeAction
	ChaincodeID
	ChaincodeInput
	ChaincodeSpec
	ChaincodeDeploymentSpec
	ChaincodeInvocationSpec
	ChaincodeSecurityContext
	ChaincodeMessage
	PutStateInfo
	RangeQueryState
	RangeQueryStateNext
	RangeQueryStateClose
	RangeQueryStateKeyValue
	RangeQueryStateResponse
	ChaincodeActionPayload
	ChaincodeEndorsedAction
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	Event
	Block2
	Message2
	SignedProposal
	Proposal
	ProposalResponse
	Response2
	ProposalResponsePayload
	Endorsement
	Transaction
	TransactionBlock
	TransactionResult
	Block
	BlockchainInfo
	NonHashData
	PeerAddress
	PeerID
	PeerEndpoint
	PeersMessage
	PeersAddresses
	SignedTransaction
	InvalidTransaction
	Transaction2
	TransactionAction
	ServerStatus
	LogLevelRequest
	LogLevelResponse
*/
package peer

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

// ChaincodeEvent is used for events and registrations that are specific to chaincode
// string type - "chaincode"
type ChaincodeEvent struct {
	ChaincodeID string `protobuf:"bytes,1,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	TxID        string `protobuf:"bytes,2,opt,name=txID" json:"txID,omitempty"`
	EventName   string `protobuf:"bytes,3,opt,name=eventName" json:"eventName,omitempty"`
	Payload     []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ChaincodeEvent) Reset()                    { *m = ChaincodeEvent{} }
func (m *ChaincodeEvent) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeEvent) ProtoMessage()               {}
func (*ChaincodeEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*ChaincodeEvent)(nil), "protos.ChaincodeEvent")
}

func init() { proto.RegisterFile("peer/chaincodeevent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0x4d, 0x2d,
	0xd2, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0x4b, 0xce, 0x4f, 0x49, 0x4d, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x75, 0x5c, 0x7c, 0xce, 0x30,
	0x79, 0x57, 0x90, 0xbc, 0x90, 0x02, 0x17, 0x37, 0x5c, 0x87, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0xb2, 0x90, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x85, 0xa7, 0x8b, 0x04, 0x13, 0x58,
	0x0a, 0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x04, 0x1b, 0xef, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x0c, 0x96,
	0x40, 0x08, 0x08, 0x49, 0x70, 0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48, 0xb0, 0x28,
	0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x4e, 0xce, 0x5c, 0x62, 0xf9, 0x45, 0xe9, 0x7a, 0x19, 0x95,
	0x05, 0xa9, 0x45, 0x39, 0xa9, 0x29, 0xe9, 0xa9, 0x45, 0x10, 0x07, 0x16, 0x47, 0x69, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x49, 0xeb, 0xa7, 0x25, 0x26, 0x15,
	0x65, 0x26, 0xeb, 0x43, 0x54, 0xe9, 0x83, 0x7c, 0x98, 0x04, 0xf1, 0x8c, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x78, 0x83, 0xed, 0x60, 0xf0, 0x00, 0x00, 0x00,
}
