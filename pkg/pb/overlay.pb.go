// Code generated by protoc-gen-go. DO NOT EDIT.
// source: overlay.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

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

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}
var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{0}
}

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{1}
}

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}
var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}
func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{13, 0}
}

type Restriction_Operand int32

const (
	Restriction_freeBandwidth Restriction_Operand = 0
	Restriction_freeDisk      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "freeBandwidth",
	1: "freeDisk",
}
var Restriction_Operand_value = map[string]int32{
	"freeBandwidth": 0,
	"freeDisk":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}
func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{13, 1}
}

// LookupRequest is is request message for the lookup rpc call
type LookupRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (dst *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(dst, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// LookupResponse is is response message for the lookup rpc call
type LookupResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{1}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (dst *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(dst, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// LookupRequests is a list of LookupRequest
type LookupRequests struct {
	Lookuprequest        []*LookupRequest `protobuf:"bytes,1,rep,name=lookuprequest,proto3" json:"lookuprequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LookupRequests) Reset()         { *m = LookupRequests{} }
func (m *LookupRequests) String() string { return proto.CompactTextString(m) }
func (*LookupRequests) ProtoMessage()    {}
func (*LookupRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{2}
}
func (m *LookupRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequests.Unmarshal(m, b)
}
func (m *LookupRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequests.Marshal(b, m, deterministic)
}
func (dst *LookupRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequests.Merge(dst, src)
}
func (m *LookupRequests) XXX_Size() int {
	return xxx_messageInfo_LookupRequests.Size(m)
}
func (m *LookupRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequests.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequests proto.InternalMessageInfo

func (m *LookupRequests) GetLookuprequest() []*LookupRequest {
	if m != nil {
		return m.Lookuprequest
	}
	return nil
}

// LookupResponse is a list of LookupResponse
type LookupResponses struct {
	Lookupresponse       []*LookupResponse `protobuf:"bytes,1,rep,name=lookupresponse,proto3" json:"lookupresponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LookupResponses) Reset()         { *m = LookupResponses{} }
func (m *LookupResponses) String() string { return proto.CompactTextString(m) }
func (*LookupResponses) ProtoMessage()    {}
func (*LookupResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{3}
}
func (m *LookupResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponses.Unmarshal(m, b)
}
func (m *LookupResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponses.Marshal(b, m, deterministic)
}
func (dst *LookupResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponses.Merge(dst, src)
}
func (m *LookupResponses) XXX_Size() int {
	return xxx_messageInfo_LookupResponses.Size(m)
}
func (m *LookupResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponses.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponses proto.InternalMessageInfo

func (m *LookupResponses) GetLookupresponse() []*LookupResponse {
	if m != nil {
		return m.Lookupresponse
	}
	return nil
}

// FindStorageNodesResponse is is response message for the FindStorageNodes rpc call
type FindStorageNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStorageNodesResponse) Reset()         { *m = FindStorageNodesResponse{} }
func (m *FindStorageNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesResponse) ProtoMessage()    {}
func (*FindStorageNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{4}
}
func (m *FindStorageNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesResponse.Unmarshal(m, b)
}
func (m *FindStorageNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesResponse.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesResponse.Merge(dst, src)
}
func (m *FindStorageNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesResponse.Size(m)
}
func (m *FindStorageNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesResponse proto.InternalMessageInfo

func (m *FindStorageNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// FindStorageNodesRequest is is request message for the FindStorageNodes rpc call
type FindStorageNodesRequest struct {
	ObjectSize           int64              `protobuf:"varint,1,opt,name=objectSize,proto3" json:"objectSize,omitempty"`
	ContractLength       *duration.Duration `protobuf:"bytes,2,opt,name=contractLength,proto3" json:"contractLength,omitempty"`
	Opts                 *OverlayOptions    `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindStorageNodesRequest) Reset()         { *m = FindStorageNodesRequest{} }
func (m *FindStorageNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesRequest) ProtoMessage()    {}
func (*FindStorageNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{5}
}
func (m *FindStorageNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesRequest.Unmarshal(m, b)
}
func (m *FindStorageNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesRequest.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesRequest.Merge(dst, src)
}
func (m *FindStorageNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesRequest.Size(m)
}
func (m *FindStorageNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesRequest proto.InternalMessageInfo

func (m *FindStorageNodesRequest) GetObjectSize() int64 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *FindStorageNodesRequest) GetContractLength() *duration.Duration {
	if m != nil {
		return m.ContractLength
	}
	return nil
}

func (m *FindStorageNodesRequest) GetOpts() *OverlayOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=overlay.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{6}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (dst *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(dst, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// OverlayOptions is a set of criteria that a node must meet to be considered for a storage opportunity
type OverlayOptions struct {
	MaxLatency           *duration.Duration `protobuf:"bytes,1,opt,name=maxLatency,proto3" json:"maxLatency,omitempty"`
	MinReputation        *NodeRep           `protobuf:"bytes,2,opt,name=minReputation,proto3" json:"minReputation,omitempty"`
	MinSpeedKbps         int64              `protobuf:"varint,3,opt,name=minSpeedKbps,proto3" json:"minSpeedKbps,omitempty"`
	Amount               int64              `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Restrictions         *NodeRestrictions  `protobuf:"bytes,5,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	ExcludedNodes        []string           `protobuf:"bytes,6,rep,name=excluded_nodes,json=excludedNodes,proto3" json:"excluded_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverlayOptions) Reset()         { *m = OverlayOptions{} }
func (m *OverlayOptions) String() string { return proto.CompactTextString(m) }
func (*OverlayOptions) ProtoMessage()    {}
func (*OverlayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{7}
}
func (m *OverlayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverlayOptions.Unmarshal(m, b)
}
func (m *OverlayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverlayOptions.Marshal(b, m, deterministic)
}
func (dst *OverlayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverlayOptions.Merge(dst, src)
}
func (m *OverlayOptions) XXX_Size() int {
	return xxx_messageInfo_OverlayOptions.Size(m)
}
func (m *OverlayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_OverlayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_OverlayOptions proto.InternalMessageInfo

func (m *OverlayOptions) GetMaxLatency() *duration.Duration {
	if m != nil {
		return m.MaxLatency
	}
	return nil
}

func (m *OverlayOptions) GetMinReputation() *NodeRep {
	if m != nil {
		return m.MinReputation
	}
	return nil
}

func (m *OverlayOptions) GetMinSpeedKbps() int64 {
	if m != nil {
		return m.MinSpeedKbps
	}
	return 0
}

func (m *OverlayOptions) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OverlayOptions) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func (m *OverlayOptions) GetExcludedNodes() []string {
	if m != nil {
		return m.ExcludedNodes
	}
	return nil
}

// NodeRep is the reputation characteristics of a node
type NodeRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRep) Reset()         { *m = NodeRep{} }
func (m *NodeRep) String() string { return proto.CompactTextString(m) }
func (*NodeRep) ProtoMessage()    {}
func (*NodeRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{8}
}
func (m *NodeRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRep.Unmarshal(m, b)
}
func (m *NodeRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRep.Marshal(b, m, deterministic)
}
func (dst *NodeRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRep.Merge(dst, src)
}
func (m *NodeRep) XXX_Size() int {
	return xxx_messageInfo_NodeRep.Size(m)
}
func (m *NodeRep) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRep.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRep proto.InternalMessageInfo

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=freeBandwidth,proto3" json:"freeBandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=freeDisk,proto3" json:"freeDisk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{9}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (dst *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(dst, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// Node represents a node in the overlay network
type Node struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=overlay.NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{10}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

type QueryRequest struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Target               *Node    `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Pingback             bool     `protobuf:"varint,4,opt,name=pingback,proto3" json:"pingback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{11}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetPingback() bool {
	if m != nil {
		return m.Pingback
	}
	return false
}

type QueryResponse struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Response             []*Node  `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{12}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (dst *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(dst, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryResponse) GetResponse() []*Node {
	if m != nil {
		return m.Response
	}
	return nil
}

type Restriction struct {
	Operator             Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand              Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value                int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Restriction) Reset()         { *m = Restriction{} }
func (m *Restriction) String() string { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()    {}
func (*Restriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8b1c97981f3c6aaa, []int{13}
}
func (m *Restriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restriction.Unmarshal(m, b)
}
func (m *Restriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restriction.Marshal(b, m, deterministic)
}
func (dst *Restriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restriction.Merge(dst, src)
}
func (m *Restriction) XXX_Size() int {
	return xxx_messageInfo_Restriction.Size(m)
}
func (m *Restriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Restriction.DiscardUnknown(m)
}

var xxx_messageInfo_Restriction proto.InternalMessageInfo

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_freeBandwidth
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "overlay.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "overlay.LookupResponse")
	proto.RegisterType((*LookupRequests)(nil), "overlay.LookupRequests")
	proto.RegisterType((*LookupResponses)(nil), "overlay.LookupResponses")
	proto.RegisterType((*FindStorageNodesResponse)(nil), "overlay.FindStorageNodesResponse")
	proto.RegisterType((*FindStorageNodesRequest)(nil), "overlay.FindStorageNodesRequest")
	proto.RegisterType((*NodeAddress)(nil), "overlay.NodeAddress")
	proto.RegisterType((*OverlayOptions)(nil), "overlay.OverlayOptions")
	proto.RegisterType((*NodeRep)(nil), "overlay.NodeRep")
	proto.RegisterType((*NodeRestrictions)(nil), "overlay.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "overlay.Node")
	proto.RegisterType((*QueryRequest)(nil), "overlay.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "overlay.QueryResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
	proto.RegisterEnum("overlay.NodeTransport", NodeTransport_name, NodeTransport_value)
	proto.RegisterEnum("overlay.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OverlayClient is the client API for Overlay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OverlayClient interface {
	// Lookup finds a nodes address from the network
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error)
}

type overlayClient struct {
	cc *grpc.ClientConn
}

func NewOverlayClient(cc *grpc.ClientConn) OverlayClient {
	return &overlayClient{cc}
}

func (c *overlayClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error) {
	out := new(LookupResponses)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/BulkLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error) {
	out := new(FindStorageNodesResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/FindStorageNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverlayServer is the server API for Overlay service.
type OverlayServer interface {
	// Lookup finds a nodes address from the network
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(context.Context, *LookupRequests) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(context.Context, *FindStorageNodesRequest) (*FindStorageNodesResponse, error)
}

func RegisterOverlayServer(s *grpc.Server, srv OverlayServer) {
	s.RegisterService(&_Overlay_serviceDesc, srv)
}

func _Overlay_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_BulkLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).BulkLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/BulkLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).BulkLookup(ctx, req.(*LookupRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_FindStorageNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).FindStorageNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/FindStorageNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).FindStorageNodes(ctx, req.(*FindStorageNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Overlay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Overlay",
	HandlerType: (*OverlayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Overlay_Lookup_Handler,
		},
		{
			MethodName: "BulkLookup",
			Handler:    _Overlay_BulkLookup_Handler,
		},
		{
			MethodName: "FindStorageNodes",
			Handler:    _Overlay_FindStorageNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

// NodesClient is the client API for Nodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodesClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesServer is the server API for Nodes service.
type NodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Nodes_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_overlay_8b1c97981f3c6aaa) }

var fileDescriptor_overlay_8b1c97981f3c6aaa = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0xaf, 0xf3, 0x3f, 0x93, 0xd8, 0xf8, 0x46, 0x47, 0x6b, 0x22, 0x38, 0xb5, 0x0b, 0x15, 0xa5,
	0x48, 0x39, 0x29, 0x77, 0xaa, 0x54, 0x09, 0x54, 0xb5, 0xd7, 0x52, 0x9d, 0x08, 0xed, 0xdd, 0xc6,
	0x12, 0x12, 0x12, 0xaa, 0x9c, 0x78, 0x2f, 0x67, 0x9a, 0x78, 0x8d, 0xbd, 0x3e, 0x2e, 0x3c, 0x04,
	0x6f, 0xc1, 0x07, 0x24, 0x9e, 0x89, 0xe7, 0xe0, 0x23, 0xf2, 0xee, 0xda, 0x89, 0xd3, 0xa6, 0xe2,
	0x3e, 0xd9, 0x33, 0xf3, 0x9b, 0xd9, 0x99, 0xdf, 0xcc, 0xce, 0x82, 0xc9, 0xdf, 0xb1, 0x78, 0xe6,
	0x2d, 0xfa, 0x51, 0xcc, 0x05, 0xc7, 0xa6, 0x16, 0x7b, 0x4f, 0xa6, 0x9c, 0x4f, 0x67, 0xec, 0xa9,
	0x54, 0x8f, 0xd3, 0x37, 0x4f, 0xfd, 0x34, 0xf6, 0x44, 0xc0, 0x43, 0x05, 0x24, 0x5f, 0x82, 0x39,
	0xe4, 0xfc, 0x36, 0x8d, 0x28, 0xfb, 0x35, 0x65, 0x89, 0xc0, 0x6d, 0x68, 0x84, 0xdc, 0x67, 0x2f,
	0xcf, 0x1d, 0x63, 0xd7, 0x38, 0x68, 0x53, 0x2d, 0x91, 0x67, 0x60, 0xe5, 0xc0, 0x24, 0xe2, 0x61,
	0xc2, 0x70, 0x0f, 0x6a, 0x99, 0x4d, 0xe2, 0x3a, 0x03, 0xb3, 0x9f, 0x67, 0x70, 0xc5, 0x7d, 0x46,
	0xa5, 0x89, 0x5c, 0x2d, 0x9d, 0x64, 0xf4, 0x04, 0xbf, 0x01, 0x73, 0x26, 0x35, 0xb1, 0xd2, 0x38,
	0xc6, 0x6e, 0xf5, 0xa0, 0x33, 0xd8, 0x2e, 0xbc, 0x4b, 0x78, 0x5a, 0x06, 0x13, 0x0a, 0x1f, 0x95,
	0x93, 0x48, 0xf0, 0x04, 0xac, 0x1c, 0xa3, 0x54, 0x3a, 0xe2, 0xce, 0x9d, 0x88, 0xca, 0x4c, 0xd7,
	0xe0, 0xe4, 0x04, 0x9c, 0xef, 0x82, 0xd0, 0x1f, 0x09, 0x1e, 0x7b, 0x53, 0x96, 0x25, 0x9f, 0x14,
	0x25, 0x7e, 0x0e, 0xf5, 0xac, 0x8e, 0x44, 0xc7, 0x5c, 0xab, 0x51, 0xd9, 0xc8, 0x5f, 0x06, 0xec,
	0xdc, 0x8d, 0xa0, 0xd8, 0x7c, 0x02, 0xc0, 0xc7, 0xbf, 0xb0, 0x89, 0x18, 0x05, 0xbf, 0x2b, 0xa6,
	0xaa, 0x74, 0x45, 0x83, 0xa7, 0x60, 0x4d, 0x78, 0x28, 0x62, 0x6f, 0x22, 0x86, 0x2c, 0x9c, 0x8a,
	0xb7, 0x4e, 0x45, 0xb2, 0xf9, 0x49, 0x5f, 0xf5, 0xad, 0x9f, 0xf7, 0xad, 0x7f, 0xae, 0xfb, 0x46,
	0xd7, 0x1c, 0xf0, 0x6b, 0xa8, 0xf1, 0x48, 0x24, 0x4e, 0x55, 0x3a, 0x2e, 0xcb, 0xbe, 0x56, 0xdf,
	0xeb, 0x28, 0xf3, 0x4a, 0xa8, 0x04, 0x91, 0x9f, 0xa1, 0x93, 0xe5, 0x77, 0xea, 0xfb, 0x31, 0x4b,
	0x12, 0x7c, 0x0e, 0x6d, 0x11, 0x7b, 0x61, 0x12, 0xf1, 0x58, 0xc8, 0xec, 0xac, 0x95, 0x4e, 0x64,
	0x40, 0x37, 0xb7, 0xd2, 0x25, 0x10, 0x1d, 0x68, 0x7a, 0x2a, 0x80, 0xcc, 0xb6, 0x4d, 0x73, 0x91,
	0xfc, 0x59, 0x01, 0xab, 0x7c, 0x2e, 0x1e, 0x03, 0xcc, 0xbd, 0xf7, 0x43, 0x4f, 0xb0, 0x70, 0xb2,
	0xd0, 0xb3, 0xf2, 0x40, 0x75, 0x2b, 0x60, 0x3c, 0x02, 0x73, 0x1e, 0x84, 0x94, 0x45, 0xa9, 0x90,
	0x46, 0xcd, 0x8d, 0x5d, 0xee, 0x02, 0x8b, 0x68, 0x19, 0x86, 0x04, 0xba, 0xf3, 0x20, 0x1c, 0x45,
	0x8c, 0xf9, 0xdf, 0x8f, 0x23, 0xc5, 0x4c, 0x95, 0x96, 0x74, 0xd9, 0x98, 0x7b, 0x73, 0x9e, 0x86,
	0xc2, 0xa9, 0x49, 0xab, 0x96, 0xf0, 0x5b, 0xe8, 0xc6, 0x2c, 0x11, 0x71, 0x30, 0x91, 0xe9, 0x3b,
	0x75, 0x9d, 0x70, 0xf9, 0xc8, 0x25, 0x80, 0x96, 0xe0, 0xb8, 0x0f, 0x16, 0x7b, 0x3f, 0x99, 0xa5,
	0x3e, 0xf3, 0x6f, 0xd4, 0xe4, 0x34, 0x76, 0xab, 0x07, 0x6d, 0x6a, 0xe6, 0x5a, 0x39, 0x1d, 0xa4,
	0x0d, 0x4d, 0x9d, 0x3b, 0x71, 0xc1, 0x5e, 0x8f, 0x89, 0x5f, 0x80, 0xf9, 0x26, 0x66, 0xec, 0xcc,
	0x0b, 0xfd, 0xdf, 0x02, 0x5f, 0xbc, 0xd5, 0x83, 0x53, 0x56, 0x62, 0x0f, 0x5a, 0x99, 0xe2, 0x3c,
	0x48, 0x6e, 0x25, 0x33, 0x55, 0x5a, 0xc8, 0xe4, 0x6f, 0x03, 0x6a, 0x59, 0x58, 0xb4, 0xa0, 0x12,
	0xf8, 0xfa, 0x2a, 0x57, 0x02, 0x1f, 0xfb, 0xe5, 0xde, 0x75, 0x06, 0x8f, 0x4b, 0xa5, 0xe9, 0xc1,
	0x28, 0x3a, 0x8a, 0xfb, 0x50, 0x13, 0x8b, 0x88, 0x49, 0x0e, 0xad, 0xc1, 0xa3, 0xf2, 0x70, 0x2c,
	0x22, 0x46, 0xa5, 0xf9, 0x0e, 0x6d, 0xb5, 0x0f, 0xa2, 0x8d, 0xfc, 0x61, 0x40, 0xf7, 0x75, 0xca,
	0xe2, 0x45, 0x7e, 0x6f, 0xf6, 0xa1, 0x91, 0xb0, 0xd0, 0x67, 0xf1, 0xfd, 0xdb, 0x45, 0x1b, 0x33,
	0x98, 0xf0, 0xe2, 0x29, 0x13, 0xba, 0x98, 0x75, 0x98, 0x32, 0xe2, 0x63, 0xa8, 0xcf, 0x82, 0x79,
	0x20, 0xf4, 0x24, 0x28, 0x21, 0xe3, 0x2f, 0x0a, 0xc2, 0xe9, 0xd8, 0x9b, 0xdc, 0xca, 0x7c, 0x5b,
	0xb4, 0x90, 0x89, 0x07, 0xa6, 0xce, 0x47, 0x6f, 0x82, 0xff, 0x99, 0xd0, 0x57, 0xd0, 0x2a, 0xf6,
	0x50, 0xe5, 0xbe, 0x9d, 0x51, 0x98, 0xc9, 0xbf, 0x06, 0x74, 0x56, 0x28, 0xc1, 0x63, 0x68, 0xf1,
	0x88, 0xc5, 0x9e, 0xe0, 0xb1, 0xbe, 0x8a, 0x9f, 0x15, 0xae, 0x2b, 0xb8, 0xfe, 0xb5, 0x06, 0xd1,
	0x02, 0x8e, 0x47, 0xd0, 0x94, 0xff, 0xa1, 0x2f, 0x79, 0xb0, 0x06, 0x9f, 0x6e, 0xf6, 0x0c, 0x7d,
	0x9a, 0x83, 0x33, 0x5e, 0xde, 0x79, 0xb3, 0x94, 0xe5, 0xbc, 0x48, 0x81, 0x3c, 0x87, 0x56, 0x7e,
	0x06, 0x36, 0xa0, 0x32, 0x74, 0xed, 0xad, 0xec, 0x7b, 0xf1, 0xda, 0x36, 0xb2, 0xef, 0xa5, 0x6b,
	0x57, 0xb0, 0x09, 0xd5, 0xa1, 0x7b, 0x61, 0x57, 0xb3, 0x9f, 0x4b, 0xf7, 0xc2, 0xae, 0x91, 0x43,
	0x68, 0xea, 0xf8, 0xf8, 0x68, 0x6d, 0x7c, 0xed, 0x2d, 0xec, 0x2e, 0x67, 0xd5, 0x36, 0x0e, 0xf7,
	0xc0, 0x2c, 0x2d, 0x17, 0xb4, 0xa1, 0xeb, 0xbe, 0x78, 0x75, 0xe3, 0x0e, 0x47, 0x37, 0x97, 0xf4,
	0xd5, 0x0b, 0x7b, 0xeb, 0x90, 0x40, 0x2b, 0x1f, 0x31, 0x6c, 0x43, 0xfd, 0xf4, 0xfc, 0x87, 0x97,
	0x57, 0xf6, 0x16, 0x76, 0xa0, 0x39, 0x72, 0xaf, 0xe9, 0xe9, 0xe5, 0x85, 0x6d, 0x0c, 0xfe, 0x31,
	0xa0, 0xa9, 0xb7, 0x0d, 0x1e, 0x43, 0x43, 0xed, 0x79, 0xdc, 0xf0, 0x94, 0xf4, 0x36, 0x3d, 0x08,
	0x78, 0x02, 0x70, 0x96, 0xce, 0x6e, 0xb5, 0xfb, 0xce, 0xfd, 0xee, 0x49, 0xcf, 0xd9, 0xe0, 0x9f,
	0xe0, 0x8f, 0x60, 0xaf, 0xef, 0x7f, 0xdc, 0x2d, 0xd0, 0x1b, 0x9e, 0x86, 0xde, 0xde, 0x03, 0x08,
	0x15, 0x79, 0x70, 0x02, 0x75, 0x15, 0xed, 0x08, 0xea, 0x72, 0x1c, 0xf1, 0xe3, 0xc2, 0x69, 0xf5,
	0xba, 0xf4, 0xb6, 0xd7, 0xd5, 0x2a, 0xc0, 0x59, 0xed, 0xa7, 0x4a, 0x34, 0x1e, 0x37, 0xe4, 0x9a,
	0x7d, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x3d, 0x3c, 0xf6, 0x24, 0x08, 0x00, 0x00,
}
