// Code generated by protoc-gen-go.
// source: github.com/codelingo/rpc/flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	github.com/codelingo/rpc/flow/flow.proto

It has these top-level messages:
	Request
	Reply
	ReviewRequest
	Issue
	IssueRange
	Position
	SearchRequest
	SearchReply
	PropertyValue
*/
package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

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

// Request contains the information to start a Flow
type Request struct {
	// flow is the name of the Flow
	Flow string `protobuf:"bytes,1,opt,name=flow" json:"flow,omitempty"`
	// payload contains any data required to configure the Flow.
	Payload *google_protobuf.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetFlow() string {
	if m != nil {
		return m.Flow
	}
	return ""
}

func (m *Request) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Reply is one reply of an executed Flow.
type Reply struct {
	// payload contains any data returned by the executed Flow.
	Payload *google_protobuf.Any `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	// error is any error found during the execution of the Flow.
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Reply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// The request message containing the files or directories to review.
type ReviewRequest struct {
	Host     string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	// Types that are valid to be assigned to OwnerOrDepot:
	//	*ReviewRequest_Owner
	//	*ReviewRequest_Depot
	OwnerOrDepot  isReviewRequest_OwnerOrDepot `protobuf_oneof:"ownerOrDepot"`
	Repo          string                       `protobuf:"bytes,4,opt,name=repo" json:"repo,omitempty"`
	Sha           string                       `protobuf:"bytes,5,opt,name=sha" json:"sha,omitempty"`
	Patches       []string                     `protobuf:"bytes,6,rep,name=Patches" json:"Patches,omitempty"`
	IsPullRequest bool                         `protobuf:"varint,7,opt,name=isPullRequest" json:"isPullRequest,omitempty"`
	PullRequestID int64                        `protobuf:"varint,8,opt,name=pullRequestID" json:"pullRequestID,omitempty"`
	Vcs           string                       `protobuf:"bytes,9,opt,name=vcs" json:"vcs,omitempty"`
	Dotlingo      string                       `protobuf:"bytes,10,opt,name=dotlingo" json:"dotlingo,omitempty"`
	Dir           string                       `protobuf:"bytes,11,opt,name=dir" json:"dir,omitempty"`
}

func (m *ReviewRequest) Reset()                    { *m = ReviewRequest{} }
func (m *ReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*ReviewRequest) ProtoMessage()               {}
func (*ReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isReviewRequest_OwnerOrDepot interface {
	isReviewRequest_OwnerOrDepot()
}

type ReviewRequest_Owner struct {
	Owner string `protobuf:"bytes,12,opt,name=owner,oneof"`
}
type ReviewRequest_Depot struct {
	Depot string `protobuf:"bytes,13,opt,name=depot,oneof"`
}

func (*ReviewRequest_Owner) isReviewRequest_OwnerOrDepot() {}
func (*ReviewRequest_Depot) isReviewRequest_OwnerOrDepot() {}

func (m *ReviewRequest) GetOwnerOrDepot() isReviewRequest_OwnerOrDepot {
	if m != nil {
		return m.OwnerOrDepot
	}
	return nil
}

func (m *ReviewRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ReviewRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ReviewRequest) GetOwner() string {
	if x, ok := m.GetOwnerOrDepot().(*ReviewRequest_Owner); ok {
		return x.Owner
	}
	return ""
}

func (m *ReviewRequest) GetDepot() string {
	if x, ok := m.GetOwnerOrDepot().(*ReviewRequest_Depot); ok {
		return x.Depot
	}
	return ""
}

func (m *ReviewRequest) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *ReviewRequest) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *ReviewRequest) GetPatches() []string {
	if m != nil {
		return m.Patches
	}
	return nil
}

func (m *ReviewRequest) GetIsPullRequest() bool {
	if m != nil {
		return m.IsPullRequest
	}
	return false
}

func (m *ReviewRequest) GetPullRequestID() int64 {
	if m != nil {
		return m.PullRequestID
	}
	return 0
}

func (m *ReviewRequest) GetVcs() string {
	if m != nil {
		return m.Vcs
	}
	return ""
}

func (m *ReviewRequest) GetDotlingo() string {
	if m != nil {
		return m.Dotlingo
	}
	return ""
}

func (m *ReviewRequest) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReviewRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReviewRequest_OneofMarshaler, _ReviewRequest_OneofUnmarshaler, _ReviewRequest_OneofSizer, []interface{}{
		(*ReviewRequest_Owner)(nil),
		(*ReviewRequest_Depot)(nil),
	}
}

func _ReviewRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReviewRequest)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *ReviewRequest_Owner:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Owner)
	case *ReviewRequest_Depot:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Depot)
	case nil:
	default:
		return fmt.Errorf("ReviewRequest.OwnerOrDepot has unexpected type %T", x)
	}
	return nil
}

func _ReviewRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReviewRequest)
	switch tag {
	case 12: // ownerOrDepot.owner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &ReviewRequest_Owner{x}
		return true, err
	case 13: // ownerOrDepot.depot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &ReviewRequest_Depot{x}
		return true, err
	default:
		return false, nil
	}
}

func _ReviewRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReviewRequest)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *ReviewRequest_Owner:
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Owner)))
		n += len(x.Owner)
	case *ReviewRequest_Depot:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Depot)))
		n += len(x.Depot)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Issue returned from a review.
type Issue struct {
	// The name of the issue.
	Name          string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Position      *IssueRange       `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Comment       string            `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	CtxBefore     string            `protobuf:"bytes,4,opt,name=ctxBefore" json:"ctxBefore,omitempty"`
	LineText      string            `protobuf:"bytes,5,opt,name=lineText" json:"lineText,omitempty"`
	CtxAfter      string            `protobuf:"bytes,6,opt,name=ctxAfter" json:"ctxAfter,omitempty"`
	Metrics       map[string]string `protobuf:"bytes,7,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tags          []string          `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
	Link          string            `protobuf:"bytes,9,opt,name=link" json:"link,omitempty"`
	NewCode       bool              `protobuf:"varint,10,opt,name=newCode" json:"newCode,omitempty"`
	Patch         string            `protobuf:"bytes,11,opt,name=patch" json:"patch,omitempty"`
	Err           string            `protobuf:"bytes,12,opt,name=err" json:"err,omitempty"`
	Discard       bool              `protobuf:"varint,13,opt,name=discard" json:"discard,omitempty"`
	DiscardReason string            `protobuf:"bytes,14,opt,name=discardReason" json:"discardReason,omitempty"`
	// isHeartbeat is used to periodically send pings from the server to keep long running connections alive.
	// clients can safely ignore this field as no other data will be sent with it.
	IsHeartbeat bool `protobuf:"varint,15,opt,name=isHeartbeat" json:"isHeartbeat,omitempty"`
}

func (m *Issue) Reset()                    { *m = Issue{} }
func (m *Issue) String() string            { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()               {}
func (*Issue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Issue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Issue) GetPosition() *IssueRange {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Issue) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Issue) GetCtxBefore() string {
	if m != nil {
		return m.CtxBefore
	}
	return ""
}

func (m *Issue) GetLineText() string {
	if m != nil {
		return m.LineText
	}
	return ""
}

func (m *Issue) GetCtxAfter() string {
	if m != nil {
		return m.CtxAfter
	}
	return ""
}

func (m *Issue) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Issue) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Issue) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Issue) GetNewCode() bool {
	if m != nil {
		return m.NewCode
	}
	return false
}

func (m *Issue) GetPatch() string {
	if m != nil {
		return m.Patch
	}
	return ""
}

func (m *Issue) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Issue) GetDiscard() bool {
	if m != nil {
		return m.Discard
	}
	return false
}

func (m *Issue) GetDiscardReason() string {
	if m != nil {
		return m.DiscardReason
	}
	return ""
}

func (m *Issue) GetIsHeartbeat() bool {
	if m != nil {
		return m.IsHeartbeat
	}
	return false
}

type IssueRange struct {
	Start *Position `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *Position `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *IssueRange) Reset()                    { *m = IssueRange{} }
func (m *IssueRange) String() string            { return proto.CompactTextString(m) }
func (*IssueRange) ProtoMessage()               {}
func (*IssueRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IssueRange) GetStart() *Position {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IssueRange) GetEnd() *Position {
	if m != nil {
		return m.End
	}
	return nil
}

type Position struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Line     int64  `protobuf:"varint,3,opt,name=Line" json:"Line,omitempty"`
	Column   int64  `protobuf:"varint,4,opt,name=Column" json:"Column,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Position) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Position) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Position) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Position) GetColumn() int64 {
	if m != nil {
		return m.Column
	}
	return 0
}

type SearchRequest struct {
	Dotlingo string `protobuf:"bytes,1,opt,name=dotlingo" json:"dotlingo,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SearchRequest) GetDotlingo() string {
	if m != nil {
		return m.Dotlingo
	}
	return ""
}

type SearchReply struct {
	Properties map[string]*PropertyValue `protobuf:"bytes,1,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error      string                    `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Depth      int64                     `protobuf:"varint,3,opt,name=depth" json:"depth,omitempty"`
	FactName   string                    `protobuf:"bytes,4,opt,name=factName" json:"factName,omitempty"`
	Decorators []string                  `protobuf:"bytes,5,rep,name=decorators" json:"decorators,omitempty"`
	Children   []*SearchReply            `protobuf:"bytes,6,rep,name=children" json:"children,omitempty"`
}

func (m *SearchReply) Reset()                    { *m = SearchReply{} }
func (m *SearchReply) String() string            { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()               {}
func (*SearchReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SearchReply) GetProperties() map[string]*PropertyValue {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *SearchReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SearchReply) GetDepth() int64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *SearchReply) GetFactName() string {
	if m != nil {
		return m.FactName
	}
	return ""
}

func (m *SearchReply) GetDecorators() []string {
	if m != nil {
		return m.Decorators
	}
	return nil
}

func (m *SearchReply) GetChildren() []*SearchReply {
	if m != nil {
		return m.Children
	}
	return nil
}

type PropertyValue struct {
	// Types that are valid to be assigned to Prop:
	//	*PropertyValue_StringProp
	//	*PropertyValue_BoolProp
	//	*PropertyValue_Int64Prop
	//	*PropertyValue_FloatProp
	Prop       isPropertyValue_Prop `protobuf_oneof:"prop"`
	Decorators []string             `protobuf:"bytes,5,rep,name=decorators" json:"decorators,omitempty"`
}

func (m *PropertyValue) Reset()                    { *m = PropertyValue{} }
func (m *PropertyValue) String() string            { return proto.CompactTextString(m) }
func (*PropertyValue) ProtoMessage()               {}
func (*PropertyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isPropertyValue_Prop interface {
	isPropertyValue_Prop()
}

type PropertyValue_StringProp struct {
	StringProp string `protobuf:"bytes,1,opt,name=stringProp,oneof"`
}
type PropertyValue_BoolProp struct {
	BoolProp bool `protobuf:"varint,2,opt,name=boolProp,oneof"`
}
type PropertyValue_Int64Prop struct {
	Int64Prop int64 `protobuf:"varint,3,opt,name=int64Prop,oneof"`
}
type PropertyValue_FloatProp struct {
	FloatProp float32 `protobuf:"fixed32,4,opt,name=floatProp,oneof"`
}

func (*PropertyValue_StringProp) isPropertyValue_Prop() {}
func (*PropertyValue_BoolProp) isPropertyValue_Prop()   {}
func (*PropertyValue_Int64Prop) isPropertyValue_Prop()  {}
func (*PropertyValue_FloatProp) isPropertyValue_Prop()  {}

func (m *PropertyValue) GetProp() isPropertyValue_Prop {
	if m != nil {
		return m.Prop
	}
	return nil
}

func (m *PropertyValue) GetStringProp() string {
	if x, ok := m.GetProp().(*PropertyValue_StringProp); ok {
		return x.StringProp
	}
	return ""
}

func (m *PropertyValue) GetBoolProp() bool {
	if x, ok := m.GetProp().(*PropertyValue_BoolProp); ok {
		return x.BoolProp
	}
	return false
}

func (m *PropertyValue) GetInt64Prop() int64 {
	if x, ok := m.GetProp().(*PropertyValue_Int64Prop); ok {
		return x.Int64Prop
	}
	return 0
}

func (m *PropertyValue) GetFloatProp() float32 {
	if x, ok := m.GetProp().(*PropertyValue_FloatProp); ok {
		return x.FloatProp
	}
	return 0
}

func (m *PropertyValue) GetDecorators() []string {
	if m != nil {
		return m.Decorators
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PropertyValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PropertyValue_OneofMarshaler, _PropertyValue_OneofUnmarshaler, _PropertyValue_OneofSizer, []interface{}{
		(*PropertyValue_StringProp)(nil),
		(*PropertyValue_BoolProp)(nil),
		(*PropertyValue_Int64Prop)(nil),
		(*PropertyValue_FloatProp)(nil),
	}
}

func _PropertyValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PropertyValue)
	// prop
	switch x := m.Prop.(type) {
	case *PropertyValue_StringProp:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringProp)
	case *PropertyValue_BoolProp:
		t := uint64(0)
		if x.BoolProp {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *PropertyValue_Int64Prop:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Prop))
	case *PropertyValue_FloatProp:
		b.EncodeVarint(4<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatProp)))
	case nil:
	default:
		return fmt.Errorf("PropertyValue.Prop has unexpected type %T", x)
	}
	return nil
}

func _PropertyValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PropertyValue)
	switch tag {
	case 1: // prop.stringProp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Prop = &PropertyValue_StringProp{x}
		return true, err
	case 2: // prop.boolProp
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Prop = &PropertyValue_BoolProp{x != 0}
		return true, err
	case 3: // prop.int64Prop
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Prop = &PropertyValue_Int64Prop{int64(x)}
		return true, err
	case 4: // prop.floatProp
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Prop = &PropertyValue_FloatProp{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _PropertyValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PropertyValue)
	// prop
	switch x := m.Prop.(type) {
	case *PropertyValue_StringProp:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringProp)))
		n += len(x.StringProp)
	case *PropertyValue_BoolProp:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *PropertyValue_Int64Prop:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64Prop))
	case *PropertyValue_FloatProp:
		n += proto.SizeVarint(4<<3 | proto.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Request)(nil), "flow.Request")
	proto.RegisterType((*Reply)(nil), "flow.Reply")
	proto.RegisterType((*ReviewRequest)(nil), "flow.ReviewRequest")
	proto.RegisterType((*Issue)(nil), "flow.Issue")
	proto.RegisterType((*IssueRange)(nil), "flow.IssueRange")
	proto.RegisterType((*Position)(nil), "flow.Position")
	proto.RegisterType((*SearchRequest)(nil), "flow.SearchRequest")
	proto.RegisterType((*SearchReply)(nil), "flow.SearchReply")
	proto.RegisterType((*PropertyValue)(nil), "flow.PropertyValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Flow service

type FlowClient interface {
	// Run sends a Request to run a Flow and returns a stream of the Flow's replies.
	Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (Flow_RunClient, error)
}

type flowClient struct {
	cc *grpc.ClientConn
}

func NewFlowClient(cc *grpc.ClientConn) FlowClient {
	return &flowClient{cc}
}

func (c *flowClient) Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (Flow_RunClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Flow_serviceDesc.Streams[0], c.cc, "/flow.Flow/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Flow_RunClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type flowRunClient struct {
	grpc.ClientStream
}

func (x *flowRunClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Flow service

type FlowServer interface {
	// Run sends a Request to run a Flow and returns a stream of the Flow's replies.
	Run(*Request, Flow_RunServer) error
}

func RegisterFlowServer(s *grpc.Server, srv FlowServer) {
	s.RegisterService(&_Flow_serviceDesc, srv)
}

func _Flow_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowServer).Run(m, &flowRunServer{stream})
}

type Flow_RunServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type flowRunServer struct {
	grpc.ServerStream
}

func (x *flowRunServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

var _Flow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flow.Flow",
	HandlerType: (*FlowServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Flow_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/codelingo/rpc/flow/flow.proto",
}

func init() { proto.RegisterFile("github.com/codelingo/rpc/flow/flow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 939 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0xce, 0xf9, 0xe2, 0xf8, 0x3c, 0x8e, 0xd3, 0xb2, 0x44, 0xd5, 0x2a, 0xaa, 0x2a, 0x63, 0x15,
	0xc9, 0x08, 0x7a, 0x46, 0x06, 0x21, 0xd4, 0xb7, 0x24, 0x05, 0xa5, 0x52, 0x42, 0xad, 0xa5, 0x82,
	0xe7, 0xcb, 0xdd, 0xd8, 0x5e, 0x7a, 0xde, 0x3d, 0x76, 0xd7, 0x49, 0xfc, 0x77, 0xf8, 0x13, 0x3c,
	0xf0, 0x13, 0xf8, 0x3d, 0xbc, 0xa3, 0xd9, 0xdb, 0x73, 0x2e, 0xa5, 0x12, 0x7d, 0x48, 0x3c, 0xdf,
	0x37, 0xb3, 0x73, 0x33, 0xdf, 0xcc, 0xde, 0xc1, 0x64, 0x29, 0xdd, 0x6a, 0x73, 0x9d, 0xe6, 0x7a,
	0x3d, 0xcd, 0x75, 0x81, 0xa5, 0x54, 0x4b, 0x3d, 0x35, 0x55, 0x3e, 0x5d, 0x94, 0xfa, 0xd6, 0xff,
	0x4b, 0x2b, 0xa3, 0x9d, 0x66, 0xfb, 0x64, 0x9f, 0x4c, 0x5b, 0xf1, 0x4b, 0x5d, 0x66, 0x6a, 0x39,
	0xf5, 0xee, 0xeb, 0xcd, 0x62, 0x5a, 0xb9, 0x6d, 0x85, 0x76, 0x9a, 0xa9, 0x2d, 0xfd, 0xd5, 0xc7,
	0xc6, 0x57, 0xd0, 0x13, 0xf8, 0xfb, 0x06, 0xad, 0x63, 0x0c, 0x7c, 0x0e, 0x1e, 0x8d, 0xa2, 0x49,
	0x5f, 0x78, 0x9b, 0xa5, 0xd0, 0xab, 0xb2, 0x6d, 0xa9, 0xb3, 0x82, 0x77, 0x46, 0xd1, 0x64, 0x30,
	0x3b, 0x4e, 0x97, 0x5a, 0x2f, 0x4b, 0x4c, 0x9b, 0xb4, 0xe9, 0xa9, 0xda, 0x8a, 0x26, 0x68, 0x7c,
	0x05, 0x5d, 0x81, 0x55, 0xb9, 0x6d, 0x1f, 0x8c, 0x3e, 0xe2, 0x20, 0x3b, 0x86, 0x2e, 0x1a, 0xa3,
	0x8d, 0x7f, 0x4c, 0x5f, 0xd4, 0x60, 0xfc, 0x77, 0x07, 0x86, 0x02, 0x6f, 0x24, 0xde, 0xb6, 0x8a,
	0x5c, 0x69, 0xeb, 0x9a, 0x22, 0xc9, 0x66, 0x27, 0x90, 0xd0, 0xaf, 0xca, 0xd6, 0x18, 0x8e, 0xef,
	0x30, 0x7b, 0x02, 0x5d, 0x7d, 0xab, 0xd0, 0xf0, 0x43, 0x72, 0x5c, 0xec, 0x89, 0x1a, 0x12, 0x5f,
	0x60, 0xa5, 0x1d, 0x1f, 0x36, 0xbc, 0x87, 0x94, 0xdf, 0x60, 0xa5, 0xf9, 0x7e, 0x9d, 0x9f, 0x6c,
	0xf6, 0x18, 0x62, 0xbb, 0xca, 0x78, 0xd7, 0x53, 0x64, 0x32, 0x0e, 0xbd, 0x79, 0xe6, 0xf2, 0x15,
	0x5a, 0x7e, 0x30, 0x8a, 0x27, 0x7d, 0xd1, 0x40, 0xf6, 0x1c, 0x86, 0xd2, 0xce, 0x37, 0x65, 0x19,
	0x0a, 0xe6, 0xbd, 0x51, 0x34, 0x49, 0xc4, 0x43, 0x92, 0xa2, 0xaa, 0x7b, 0xf8, 0xfa, 0x15, 0x4f,
	0x46, 0xd1, 0x24, 0x16, 0x0f, 0x49, 0x7a, 0xee, 0x4d, 0x6e, 0x79, 0xbf, 0x7e, 0xee, 0x4d, 0x6e,
	0xa9, 0xd3, 0x42, 0x3b, 0xbf, 0x04, 0x1c, 0xea, 0x4e, 0x1b, 0x4c, 0xd1, 0x85, 0x34, 0x7c, 0x50,
	0x47, 0x17, 0xd2, 0x9c, 0x1d, 0xc1, 0xa1, 0x6f, 0xf6, 0x8d, 0x79, 0x45, 0xbd, 0x8d, 0xff, 0x89,
	0xa1, 0xfb, 0xda, 0xda, 0x0d, 0x52, 0x97, 0x5e, 0xad, 0xa0, 0xa2, 0x57, 0xea, 0x2b, 0x48, 0x2a,
	0x6d, 0xa5, 0x93, 0x5a, 0x85, 0x59, 0x3f, 0x4e, 0xfd, 0x7e, 0xf9, 0x23, 0x22, 0x53, 0x4b, 0x14,
	0xbb, 0x08, 0x52, 0x20, 0xd7, 0xeb, 0x35, 0x2a, 0xc7, 0x63, 0x9f, 0xa4, 0x81, 0xec, 0x29, 0xf4,
	0x73, 0x77, 0x77, 0x86, 0x0b, 0x6d, 0x30, 0xc8, 0x78, 0x4f, 0x50, 0x07, 0xa5, 0x54, 0xf8, 0x16,
	0xef, 0x5c, 0x10, 0x74, 0x87, 0xc9, 0x97, 0xbb, 0xbb, 0xd3, 0x85, 0x43, 0xc3, 0x0f, 0x6a, 0x5f,
	0x83, 0xd9, 0x0c, 0x7a, 0x6b, 0x74, 0x46, 0xe6, 0x96, 0xf7, 0x46, 0xf1, 0x64, 0x30, 0xe3, 0xad,
	0xe2, 0xd2, 0xab, 0xda, 0xf5, 0x83, 0x72, 0x66, 0x2b, 0x9a, 0x40, 0xea, 0xd2, 0x65, 0x4b, 0xcb,
	0x13, 0x3f, 0x22, 0x6f, 0x13, 0x57, 0x4a, 0xf5, 0x2e, 0x88, 0xea, 0x6d, 0xea, 0x45, 0xe1, 0xed,
	0xb9, 0x2e, 0xd0, 0x8b, 0x9a, 0x88, 0x06, 0xd2, 0x56, 0x56, 0x34, 0xd8, 0xa0, 0x6a, 0x0d, 0x48,
	0x69, 0x34, 0x61, 0xa3, 0x04, 0x99, 0x94, 0xa1, 0x90, 0x36, 0xcf, 0x4c, 0xe1, 0xf7, 0x29, 0x11,
	0x0d, 0xa4, 0x49, 0x07, 0x53, 0x60, 0x66, 0xb5, 0xe2, 0x47, 0xfe, 0xd4, 0x43, 0x92, 0x8d, 0x60,
	0x20, 0xed, 0x05, 0x66, 0xc6, 0x5d, 0x63, 0xe6, 0xf8, 0x23, 0x9f, 0xa3, 0x4d, 0x9d, 0xbc, 0x84,
	0xc3, 0x76, 0x93, 0x54, 0xc3, 0x3b, 0xdc, 0x86, 0x01, 0x92, 0x49, 0xb5, 0xde, 0x64, 0xe5, 0xa6,
	0xb9, 0x02, 0x35, 0x78, 0xd9, 0xf9, 0x3e, 0x1a, 0xbf, 0x05, 0xb8, 0x9f, 0x21, 0x7b, 0x0e, 0x5d,
	0xeb, 0x32, 0xe3, 0xc2, 0xbd, 0x3c, 0xaa, 0x75, 0x9c, 0x87, 0xc1, 0x8a, 0xda, 0xc9, 0x46, 0x10,
	0xa3, 0x6a, 0x2e, 0xfd, 0xfb, 0x31, 0xe4, 0x1a, 0xff, 0x06, 0x49, 0x43, 0xd0, 0xe4, 0x16, 0xb2,
	0xc4, 0xd6, 0x4e, 0xed, 0x30, 0x7b, 0x02, 0x07, 0x6f, 0x16, 0x0b, 0x8b, 0xce, 0x27, 0x8b, 0x45,
	0x40, 0x34, 0x89, 0x4b, 0xa9, 0xd0, 0xaf, 0x4f, 0x2c, 0xbc, 0x4d, 0xb1, 0xe7, 0xba, 0xdc, 0xac,
	0x95, 0x5f, 0x9c, 0x58, 0x04, 0x34, 0xfe, 0x12, 0x86, 0x3f, 0x63, 0x66, 0xf2, 0x55, 0x73, 0x81,
	0xda, 0x17, 0x21, 0x7a, 0x78, 0x11, 0xc6, 0x7f, 0x75, 0x60, 0xd0, 0x44, 0xd3, 0xab, 0xe8, 0x14,
	0xa0, 0x32, 0xba, 0x42, 0xe3, 0x24, 0x5a, 0x1e, 0xf9, 0xed, 0xf9, 0xac, 0xee, 0xa8, 0x15, 0x96,
	0xce, 0x77, 0x31, 0xf5, 0x1a, 0xb5, 0x0e, 0x7d, 0xf8, 0xed, 0x44, 0x6c, 0x81, 0x95, 0x5b, 0x85,
	0x16, 0x6a, 0xe0, 0xb5, 0xc8, 0x72, 0xf7, 0x13, 0x69, 0xb1, 0x1f, 0xb4, 0x08, 0x98, 0x3d, 0x03,
	0x28, 0x30, 0xd7, 0x26, 0x73, 0xda, 0x58, 0xde, 0xf5, 0x7b, 0xd9, 0x62, 0xd8, 0x0b, 0x48, 0xf2,
	0x95, 0x2c, 0x0b, 0x83, 0xca, 0xbf, 0x58, 0x06, 0xb3, 0x4f, 0xfe, 0x53, 0xa8, 0xd8, 0x85, 0x9c,
	0x08, 0x78, 0xf4, 0x5e, 0xd5, 0x1f, 0xd8, 0x8b, 0x2f, 0xda, 0x7b, 0x31, 0x98, 0x7d, 0x1a, 0x66,
	0x59, 0x9f, 0xdb, 0xfe, 0x42, 0xae, 0xf6, 0xb2, 0xfc, 0x19, 0xc1, 0xf0, 0x81, 0x93, 0x8d, 0x00,
	0xac, 0x33, 0x52, 0x2d, 0x89, 0xae, 0x33, 0x5f, 0xec, 0x89, 0x16, 0xc7, 0x9e, 0x42, 0x72, 0xad,
	0x75, 0xe9, 0xfd, 0xf4, 0x94, 0xe4, 0x62, 0x4f, 0xec, 0x18, 0xf6, 0x0c, 0xfa, 0x52, 0xb9, 0xef,
	0xbe, 0xf5, 0x6e, 0x2f, 0xd5, 0xc5, 0x9e, 0xb8, 0xa7, 0xc8, 0xbf, 0x28, 0x75, 0xe6, 0xbc, 0x9f,
	0x14, 0xeb, 0x90, 0x7f, 0x47, 0xfd, 0x9f, 0x68, 0x67, 0x07, 0xb0, 0x4f, 0xa3, 0x9a, 0xbd, 0x80,
	0xfd, 0x1f, 0xe9, 0x9b, 0xf5, 0x39, 0xc4, 0x62, 0xa3, 0xd8, 0xb0, 0x6e, 0x34, 0x6c, 0xcc, 0xc9,
	0xa0, 0x81, 0x55, 0xb9, 0x1d, 0xef, 0x7d, 0x1d, 0x9d, 0xa5, 0x70, 0x2c, 0x75, 0xba, 0xfb, 0xa8,
	0xa6, 0x4b, 0x53, 0xe5, 0x69, 0x5e, 0x9e, 0x1d, 0xd1, 0xcd, 0xbf, 0x24, 0x6a, 0x4e, 0x5f, 0xaa,
	0x79, 0xf4, 0x47, 0x27, 0x3e, 0xbf, 0xfc, 0xf5, 0xfa, 0xc0, 0x7f, 0xb8, 0xbe, 0xf9, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x91, 0x89, 0x69, 0x43, 0x93, 0x07, 0x00, 0x00,
}
