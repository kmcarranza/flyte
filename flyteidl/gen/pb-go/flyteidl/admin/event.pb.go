// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/event.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import event "github.com/lyft/flyteidl/gen/pb-go/flyteidl/event"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventErrorAlreadyInTerminalState struct {
	CurrentPhase         string   `protobuf:"bytes,1,opt,name=current_phase,json=currentPhase,proto3" json:"current_phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventErrorAlreadyInTerminalState) Reset()         { *m = EventErrorAlreadyInTerminalState{} }
func (m *EventErrorAlreadyInTerminalState) String() string { return proto.CompactTextString(m) }
func (*EventErrorAlreadyInTerminalState) ProtoMessage()    {}
func (*EventErrorAlreadyInTerminalState) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{0}
}
func (m *EventErrorAlreadyInTerminalState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Unmarshal(m, b)
}
func (m *EventErrorAlreadyInTerminalState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Marshal(b, m, deterministic)
}
func (dst *EventErrorAlreadyInTerminalState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventErrorAlreadyInTerminalState.Merge(dst, src)
}
func (m *EventErrorAlreadyInTerminalState) XXX_Size() int {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Size(m)
}
func (m *EventErrorAlreadyInTerminalState) XXX_DiscardUnknown() {
	xxx_messageInfo_EventErrorAlreadyInTerminalState.DiscardUnknown(m)
}

var xxx_messageInfo_EventErrorAlreadyInTerminalState proto.InternalMessageInfo

func (m *EventErrorAlreadyInTerminalState) GetCurrentPhase() string {
	if m != nil {
		return m.CurrentPhase
	}
	return ""
}

type EventFailureReason struct {
	// Types that are valid to be assigned to Reason:
	//	*EventFailureReason_AlreadyInTerminalState
	Reason               isEventFailureReason_Reason `protobuf_oneof:"reason"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventFailureReason) Reset()         { *m = EventFailureReason{} }
func (m *EventFailureReason) String() string { return proto.CompactTextString(m) }
func (*EventFailureReason) ProtoMessage()    {}
func (*EventFailureReason) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{1}
}
func (m *EventFailureReason) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFailureReason.Unmarshal(m, b)
}
func (m *EventFailureReason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFailureReason.Marshal(b, m, deterministic)
}
func (dst *EventFailureReason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFailureReason.Merge(dst, src)
}
func (m *EventFailureReason) XXX_Size() int {
	return xxx_messageInfo_EventFailureReason.Size(m)
}
func (m *EventFailureReason) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFailureReason.DiscardUnknown(m)
}

var xxx_messageInfo_EventFailureReason proto.InternalMessageInfo

type isEventFailureReason_Reason interface {
	isEventFailureReason_Reason()
}

type EventFailureReason_AlreadyInTerminalState struct {
	AlreadyInTerminalState *EventErrorAlreadyInTerminalState `protobuf:"bytes,1,opt,name=already_in_terminal_state,json=alreadyInTerminalState,proto3,oneof"`
}

func (*EventFailureReason_AlreadyInTerminalState) isEventFailureReason_Reason() {}

func (m *EventFailureReason) GetReason() isEventFailureReason_Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (m *EventFailureReason) GetAlreadyInTerminalState() *EventErrorAlreadyInTerminalState {
	if x, ok := m.GetReason().(*EventFailureReason_AlreadyInTerminalState); ok {
		return x.AlreadyInTerminalState
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventFailureReason) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventFailureReason_OneofMarshaler, _EventFailureReason_OneofUnmarshaler, _EventFailureReason_OneofSizer, []interface{}{
		(*EventFailureReason_AlreadyInTerminalState)(nil),
	}
}

func _EventFailureReason_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventFailureReason)
	// reason
	switch x := m.Reason.(type) {
	case *EventFailureReason_AlreadyInTerminalState:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AlreadyInTerminalState); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventFailureReason.Reason has unexpected type %T", x)
	}
	return nil
}

func _EventFailureReason_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventFailureReason)
	switch tag {
	case 1: // reason.already_in_terminal_state
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventErrorAlreadyInTerminalState)
		err := b.DecodeMessage(msg)
		m.Reason = &EventFailureReason_AlreadyInTerminalState{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventFailureReason_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventFailureReason)
	// reason
	switch x := m.Reason.(type) {
	case *EventFailureReason_AlreadyInTerminalState:
		s := proto.Size(x.AlreadyInTerminalState)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request to send a notification that a workflow execution event has occurred.
type WorkflowExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.WorkflowExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *WorkflowExecutionEventRequest) Reset()         { *m = WorkflowExecutionEventRequest{} }
func (m *WorkflowExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventRequest) ProtoMessage()    {}
func (*WorkflowExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{2}
}
func (m *WorkflowExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventRequest.Merge(dst, src)
}
func (m *WorkflowExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Size(m)
}
func (m *WorkflowExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventRequest proto.InternalMessageInfo

func (m *WorkflowExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetEvent() *event.WorkflowExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type WorkflowExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionEventResponse) Reset()         { *m = WorkflowExecutionEventResponse{} }
func (m *WorkflowExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventResponse) ProtoMessage()    {}
func (*WorkflowExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{3}
}
func (m *WorkflowExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventResponse.Merge(dst, src)
}
func (m *WorkflowExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Size(m)
}
func (m *WorkflowExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventResponse proto.InternalMessageInfo

// Request to send a notification that a node execution event has occurred.
type NodeExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.NodeExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *NodeExecutionEventRequest) Reset()         { *m = NodeExecutionEventRequest{} }
func (m *NodeExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventRequest) ProtoMessage()    {}
func (*NodeExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{4}
}
func (m *NodeExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventRequest.Unmarshal(m, b)
}
func (m *NodeExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventRequest.Merge(dst, src)
}
func (m *NodeExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventRequest.Size(m)
}
func (m *NodeExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventRequest proto.InternalMessageInfo

func (m *NodeExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetEvent() *event.NodeExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type NodeExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEventResponse) Reset()         { *m = NodeExecutionEventResponse{} }
func (m *NodeExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventResponse) ProtoMessage()    {}
func (*NodeExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{5}
}
func (m *NodeExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventResponse.Unmarshal(m, b)
}
func (m *NodeExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventResponse.Merge(dst, src)
}
func (m *NodeExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventResponse.Size(m)
}
func (m *NodeExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventResponse proto.InternalMessageInfo

// Request to send a notification that a task execution event has occurred.
type TaskExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.TaskExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TaskExecutionEventRequest) Reset()         { *m = TaskExecutionEventRequest{} }
func (m *TaskExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventRequest) ProtoMessage()    {}
func (*TaskExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{6}
}
func (m *TaskExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventRequest.Unmarshal(m, b)
}
func (m *TaskExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventRequest.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventRequest.Merge(dst, src)
}
func (m *TaskExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventRequest.Size(m)
}
func (m *TaskExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventRequest proto.InternalMessageInfo

func (m *TaskExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetEvent() *event.TaskExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type TaskExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEventResponse) Reset()         { *m = TaskExecutionEventResponse{} }
func (m *TaskExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventResponse) ProtoMessage()    {}
func (*TaskExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_424b5552574ecbeb, []int{7}
}
func (m *TaskExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventResponse.Unmarshal(m, b)
}
func (m *TaskExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventResponse.Marshal(b, m, deterministic)
}
func (dst *TaskExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventResponse.Merge(dst, src)
}
func (m *TaskExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventResponse.Size(m)
}
func (m *TaskExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventErrorAlreadyInTerminalState)(nil), "flyteidl.admin.EventErrorAlreadyInTerminalState")
	proto.RegisterType((*EventFailureReason)(nil), "flyteidl.admin.EventFailureReason")
	proto.RegisterType((*WorkflowExecutionEventRequest)(nil), "flyteidl.admin.WorkflowExecutionEventRequest")
	proto.RegisterType((*WorkflowExecutionEventResponse)(nil), "flyteidl.admin.WorkflowExecutionEventResponse")
	proto.RegisterType((*NodeExecutionEventRequest)(nil), "flyteidl.admin.NodeExecutionEventRequest")
	proto.RegisterType((*NodeExecutionEventResponse)(nil), "flyteidl.admin.NodeExecutionEventResponse")
	proto.RegisterType((*TaskExecutionEventRequest)(nil), "flyteidl.admin.TaskExecutionEventRequest")
	proto.RegisterType((*TaskExecutionEventResponse)(nil), "flyteidl.admin.TaskExecutionEventResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/event.proto", fileDescriptor_event_424b5552574ecbeb) }

var fileDescriptor_event_424b5552574ecbeb = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0xdf, 0xbe, 0x60, 0xb1, 0xe3, 0x9f, 0x43, 0x0e, 0xd2, 0x16, 0x2b, 0x25, 0x82, 0x78,
	0x31, 0x51, 0x7b, 0xf1, 0xe0, 0xc5, 0x42, 0xd5, 0x5e, 0x44, 0x62, 0x41, 0xf0, 0x12, 0xb6, 0xcd,
	0xb4, 0x5d, 0xba, 0xd9, 0x8d, 0xbb, 0x13, 0xb5, 0xe0, 0x57, 0xf0, 0x3b, 0x4b, 0x37, 0xd1, 0x58,
	0x49, 0x11, 0x04, 0x6f, 0xc9, 0xcc, 0xe4, 0xf9, 0x3d, 0x59, 0x66, 0xa1, 0x39, 0x16, 0x73, 0x42,
	0x1e, 0x09, 0x9f, 0x45, 0x31, 0x97, 0x3e, 0x3e, 0xa1, 0x24, 0x2f, 0xd1, 0x8a, 0x94, 0xb3, 0xfd,
	0xd1, 0xf3, 0x6c, 0xaf, 0x59, 0xcc, 0xda, 0xa9, 0xaf, 0xb3, 0xee, 0x15, 0xb4, 0x7b, 0x8b, 0xd7,
	0x9e, 0xd6, 0x4a, 0x5f, 0x08, 0x8d, 0x2c, 0x9a, 0xf7, 0xe5, 0x00, 0x75, 0xcc, 0x25, 0x13, 0x77,
	0xc4, 0x08, 0x9d, 0x7d, 0xd8, 0x1a, 0xa5, 0x5a, 0xa3, 0xa4, 0x30, 0x99, 0x32, 0x83, 0xf5, 0x4a,
	0xbb, 0x72, 0x58, 0x0b, 0x36, 0xf3, 0xe2, 0xed, 0xa2, 0xe6, 0xbe, 0x55, 0xc0, 0xb1, 0xa4, 0x4b,
	0xc6, 0x45, 0xaa, 0x31, 0x40, 0x66, 0x94, 0x74, 0x62, 0x68, 0xb0, 0x8c, 0x1a, 0x72, 0x19, 0x52,
	0xce, 0x0d, 0xcd, 0x02, 0x6c, 0x39, 0x1b, 0xa7, 0xc7, 0xde, 0xb2, 0xaf, 0xf7, 0x93, 0xd0, 0xf5,
	0xbf, 0x60, 0x87, 0x95, 0x76, 0xba, 0xeb, 0x50, 0xd5, 0x36, 0xd8, 0x7d, 0x85, 0xd6, 0xbd, 0xd2,
	0xb3, 0xb1, 0x50, 0xcf, 0xbd, 0x17, 0x1c, 0xa5, 0xc4, 0x95, 0xb4, 0xe0, 0x00, 0x1f, 0x53, 0x34,
	0xe4, 0xb4, 0x00, 0x74, 0xf6, 0x18, 0xf2, 0x28, 0xff, 0xa5, 0x5a, 0x5e, 0xe9, 0x47, 0xce, 0x39,
	0xac, 0xd9, 0x73, 0xaa, 0xff, 0xb7, 0x92, 0x07, 0x85, 0x64, 0x76, 0x7c, 0x2b, 0xe0, 0xd9, 0x47,
	0x6e, 0x1b, 0xf6, 0x56, 0xa5, 0x9b, 0x44, 0x49, 0x83, 0x2e, 0x41, 0xe3, 0x46, 0x45, 0xf8, 0x2b,
	0xb7, 0xb3, 0x65, 0x37, 0xf7, 0xbb, 0x5b, 0x09, 0x38, 0xf7, 0xda, 0x85, 0x66, 0x59, 0x6a, 0xe1,
	0x34, 0x60, 0x66, 0xf6, 0x27, 0x4e, 0x25, 0xe0, 0xc2, 0xa9, 0x2c, 0x35, 0x73, 0xea, 0x76, 0x1e,
	0x4e, 0x26, 0x9c, 0xa6, 0xe9, 0xd0, 0x1b, 0xa9, 0xd8, 0x17, 0xf3, 0x31, 0xf9, 0x9f, 0xeb, 0x3c,
	0x41, 0xe9, 0x27, 0xc3, 0xa3, 0x89, 0xf2, 0x97, 0x6f, 0xc3, 0xb0, 0x6a, 0x97, 0xbb, 0xf3, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x4d, 0x05, 0x6c, 0x2c, 0x26, 0x03, 0x00, 0x00,
}
