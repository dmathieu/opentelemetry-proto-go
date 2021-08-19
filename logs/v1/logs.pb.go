// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/logs/v1/logs.proto

package v1

import (
	fmt "fmt"
	v11 "github.com/dmathieu/opentelemetry-proto-go/common/v1"
	v1 "github.com/dmathieu/opentelemetry-proto-go/resource/v1"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "SEVERITY_NUMBER_UNSPECIFIED",
	1:  "SEVERITY_NUMBER_TRACE",
	2:  "SEVERITY_NUMBER_TRACE2",
	3:  "SEVERITY_NUMBER_TRACE3",
	4:  "SEVERITY_NUMBER_TRACE4",
	5:  "SEVERITY_NUMBER_DEBUG",
	6:  "SEVERITY_NUMBER_DEBUG2",
	7:  "SEVERITY_NUMBER_DEBUG3",
	8:  "SEVERITY_NUMBER_DEBUG4",
	9:  "SEVERITY_NUMBER_INFO",
	10: "SEVERITY_NUMBER_INFO2",
	11: "SEVERITY_NUMBER_INFO3",
	12: "SEVERITY_NUMBER_INFO4",
	13: "SEVERITY_NUMBER_WARN",
	14: "SEVERITY_NUMBER_WARN2",
	15: "SEVERITY_NUMBER_WARN3",
	16: "SEVERITY_NUMBER_WARN4",
	17: "SEVERITY_NUMBER_ERROR",
	18: "SEVERITY_NUMBER_ERROR2",
	19: "SEVERITY_NUMBER_ERROR3",
	20: "SEVERITY_NUMBER_ERROR4",
	21: "SEVERITY_NUMBER_FATAL",
	22: "SEVERITY_NUMBER_FATAL2",
	23: "SEVERITY_NUMBER_FATAL3",
	24: "SEVERITY_NUMBER_FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"SEVERITY_NUMBER_UNSPECIFIED": 0,
	"SEVERITY_NUMBER_TRACE":       1,
	"SEVERITY_NUMBER_TRACE2":      2,
	"SEVERITY_NUMBER_TRACE3":      3,
	"SEVERITY_NUMBER_TRACE4":      4,
	"SEVERITY_NUMBER_DEBUG":       5,
	"SEVERITY_NUMBER_DEBUG2":      6,
	"SEVERITY_NUMBER_DEBUG3":      7,
	"SEVERITY_NUMBER_DEBUG4":      8,
	"SEVERITY_NUMBER_INFO":        9,
	"SEVERITY_NUMBER_INFO2":       10,
	"SEVERITY_NUMBER_INFO3":       11,
	"SEVERITY_NUMBER_INFO4":       12,
	"SEVERITY_NUMBER_WARN":        13,
	"SEVERITY_NUMBER_WARN2":       14,
	"SEVERITY_NUMBER_WARN3":       15,
	"SEVERITY_NUMBER_WARN4":       16,
	"SEVERITY_NUMBER_ERROR":       17,
	"SEVERITY_NUMBER_ERROR2":      18,
	"SEVERITY_NUMBER_ERROR3":      19,
	"SEVERITY_NUMBER_ERROR4":      20,
	"SEVERITY_NUMBER_FATAL":       21,
	"SEVERITY_NUMBER_FATAL2":      22,
	"SEVERITY_NUMBER_FATAL3":      23,
	"SEVERITY_NUMBER_FATAL4":      24,
}

func (x SeverityNumber) String() string {
	return proto.EnumName(SeverityNumber_name, int32(x))
}

func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_LOG_RECORD_FLAG_UNSPECIFIED      LogRecordFlags = 0
	LogRecordFlags_LOG_RECORD_FLAG_TRACE_FLAGS_MASK LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "LOG_RECORD_FLAG_UNSPECIFIED",
	255: "LOG_RECORD_FLAG_TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"LOG_RECORD_FLAG_UNSPECIFIED":      0,
	"LOG_RECORD_FLAG_TRACE_FLAGS_MASK": 255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}

// A collection of InstrumentationLibraryLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then resource info is unknown.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibraryLogs that originate from a resource.
	InstrumentationLibraryLogs []*InstrumentationLibraryLogs `protobuf:"bytes,2,rep,name=instrumentation_library_logs,json=instrumentationLibraryLogs,proto3" json:"instrumentation_library_logs,omitempty"`
	// This schema_url applies to the data in the "resource" field. It does not apply
	// to the data in the "instrumentation_library_logs" field which have their own
	// schema_url field.
	SchemaUrl            string   `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceLogs) Reset()         { *m = ResourceLogs{} }
func (m *ResourceLogs) String() string { return proto.CompactTextString(m) }
func (*ResourceLogs) ProtoMessage()    {}
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}
func (m *ResourceLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLogs.Unmarshal(m, b)
}
func (m *ResourceLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLogs.Marshal(b, m, deterministic)
}
func (m *ResourceLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLogs.Merge(m, src)
}
func (m *ResourceLogs) XXX_Size() int {
	return xxx_messageInfo_ResourceLogs.Size(m)
}
func (m *ResourceLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLogs proto.InternalMessageInfo

func (m *ResourceLogs) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceLogs) GetInstrumentationLibraryLogs() []*InstrumentationLibraryLogs {
	if m != nil {
		return m.InstrumentationLibraryLogs
	}
	return nil
}

func (m *ResourceLogs) GetSchemaUrl() string {
	if m != nil {
		return m.SchemaUrl
	}
	return ""
}

// A collection of Logs produced by an InstrumentationLibrary.
type InstrumentationLibraryLogs struct {
	// The instrumentation library information for the logs in this message.
	// Semantically when InstrumentationLibrary isn't set, it is equivalent with
	// an empty instrumentation library name (unknown).
	InstrumentationLibrary *v11.InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of log records.
	Logs []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	// This schema_url applies to all logs in the "logs" field.
	SchemaUrl            string   `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentationLibraryLogs) Reset()         { *m = InstrumentationLibraryLogs{} }
func (m *InstrumentationLibraryLogs) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibraryLogs) ProtoMessage()    {}
func (*InstrumentationLibraryLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}
func (m *InstrumentationLibraryLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibraryLogs.Unmarshal(m, b)
}
func (m *InstrumentationLibraryLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibraryLogs.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibraryLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibraryLogs.Merge(m, src)
}
func (m *InstrumentationLibraryLogs) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibraryLogs.Size(m)
}
func (m *InstrumentationLibraryLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibraryLogs.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibraryLogs proto.InternalMessageInfo

func (m *InstrumentationLibraryLogs) GetInstrumentationLibrary() *v11.InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetSchemaUrl() string {
	if m != nil {
		return m.SchemaUrl
	}
	return ""
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=opentelemetry.proto.logs.v1.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. Name describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. [Optional].
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *v11.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	Attributes             []*v11.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32          `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. [Optional].
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. [Optional].
	SpanId               []byte   `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{2}
}
func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetTimeUnixNano() uint64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogRecord) GetBody() *v11.AnyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *LogRecord) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.logs.v1.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("opentelemetry.proto.logs.v1.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "opentelemetry.proto.logs.v1.ResourceLogs")
	proto.RegisterType((*InstrumentationLibraryLogs)(nil), "opentelemetry.proto.logs.v1.InstrumentationLibraryLogs")
	proto.RegisterType((*LogRecord)(nil), "opentelemetry.proto.logs.v1.LogRecord")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/logs/v1/logs.proto", fileDescriptor_d1c030a3ec7e961e)
}

var fileDescriptor_d1c030a3ec7e961e = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0xeb, 0x84, 0x40, 0x38, 0x21, 0xec, 0x74, 0x9a, 0x4d, 0xbc, 0xa4, 0xed, 0x5a, 0xdb,
	0x76, 0x4b, 0x53, 0x2d, 0x28, 0x86, 0xaa, 0x55, 0x7b, 0xe5, 0x10, 0x83, 0xd0, 0xb2, 0x10, 0x0d,
	0xb0, 0xfd, 0xb8, 0xb1, 0x0c, 0x9e, 0x12, 0x4b, 0x78, 0x06, 0xd9, 0x63, 0x04, 0xcf, 0xd7, 0xbb,
	0xbe, 0x48, 0x2f, 0xfb, 0x08, 0xad, 0x3c, 0x7c, 0xec, 0x82, 0x3c, 0xe4, 0x8a, 0x99, 0xf3, 0x3b,
	0xff, 0x33, 0xff, 0x73, 0xb0, 0xc7, 0xf0, 0x9a, 0xcf, 0x28, 0x13, 0x74, 0x4a, 0x03, 0x2a, 0xc2,
	0x65, 0x75, 0x16, 0x72, 0xc1, 0xab, 0x53, 0x3e, 0x89, 0xaa, 0xf3, 0x5b, 0xf9, 0x5b, 0x91, 0x21,
	0x7c, 0xbd, 0x93, 0xb7, 0x0a, 0x56, 0x24, 0x9f, 0xdf, 0x96, 0x6e, 0xd2, 0x8a, 0x8c, 0x79, 0x10,
	0x70, 0x96, 0x94, 0x59, 0xad, 0x56, 0x9a, 0x52, 0x25, 0x2d, 0x37, 0xa4, 0x11, 0x8f, 0xc3, 0x31,
	0x4d, 0xb2, 0x37, 0xeb, 0x55, 0xfe, 0xab, 0x7f, 0x35, 0x28, 0x90, 0x75, 0xa8, 0xc3, 0x27, 0x11,
	0xb6, 0xe1, 0x74, 0x93, 0xa2, 0x6b, 0x86, 0x56, 0x3e, 0x33, 0xbf, 0xab, 0xa4, 0x99, 0xdb, 0xd6,
	0x99, 0xdf, 0x56, 0x36, 0x05, 0xc8, 0x56, 0x8a, 0x97, 0xf0, 0xb9, 0xcf, 0x22, 0x11, 0xc6, 0x01,
	0x65, 0xc2, 0x15, 0x3e, 0x67, 0xce, 0xd4, 0x1f, 0x85, 0x6e, 0xb8, 0x74, 0x92, 0xb6, 0xf4, 0x23,
	0xe3, 0xb8, 0x7c, 0x66, 0xfe, 0x58, 0x39, 0xd0, 0x77, 0xa5, 0xbd, 0x5b, 0xa0, 0xb3, 0xd2, 0x27,
	0x2e, 0x49, 0xc9, 0x57, 0x32, 0xfc, 0x05, 0x40, 0x34, 0x7e, 0xa4, 0x81, 0xeb, 0xc4, 0xe1, 0x54,
	0x3f, 0x36, 0xb4, 0x72, 0x9e, 0xe4, 0x57, 0x91, 0x61, 0x38, 0x7d, 0xf5, 0x8f, 0x06, 0x25, 0x75,
	0x65, 0xcc, 0xe0, 0x4a, 0x61, 0x7c, 0x3d, 0x8e, 0x1f, 0x52, 0x3d, 0xaf, 0xff, 0x04, 0xa5, 0x6b,
	0x72, 0x99, 0xee, 0x18, 0xff, 0x0c, 0x99, 0x8f, 0x06, 0xf2, 0xfa, 0xe0, 0x40, 0x3a, 0x7c, 0x42,
	0xe8, 0x98, 0x87, 0x1e, 0x91, 0x9a, 0xa7, 0x3a, 0xfd, 0xfb, 0x18, 0xf2, 0x5b, 0x09, 0xfe, 0x1a,
	0x8a, 0xc2, 0x0f, 0xa8, 0x13, 0x33, 0x7f, 0xe1, 0x30, 0x97, 0x71, 0xd9, 0x4f, 0x96, 0x14, 0x92,
	0xe8, 0x90, 0xf9, 0x8b, 0xae, 0xcb, 0x38, 0x1e, 0xc0, 0xb3, 0x88, 0xce, 0x69, 0xe8, 0x8b, 0xa5,
	0xc3, 0xe2, 0x60, 0x44, 0x43, 0xfd, 0xc8, 0xd0, 0xca, 0x45, 0xf3, 0xfb, 0x83, 0xce, 0xfa, 0x6b,
	0x4d, 0x57, 0x4a, 0x48, 0x31, 0xda, 0xd9, 0xe3, 0xaf, 0xe0, 0x7c, 0x5b, 0x55, 0xd0, 0x85, 0x58,
	0x7b, 0x2d, 0x6c, 0x82, 0x03, 0xba, 0x10, 0x18, 0x43, 0x86, 0xb9, 0x01, 0xd5, 0x33, 0x92, 0xc9,
	0x35, 0xfe, 0x05, 0x32, 0x23, 0xee, 0x2d, 0xf5, 0x13, 0x39, 0xfa, 0x6f, 0x9f, 0x18, 0xbd, 0xc5,
	0x96, 0xef, 0xdd, 0x69, 0x4c, 0x89, 0x14, 0xe1, 0x16, 0x80, 0x2b, 0x44, 0xe8, 0x8f, 0x62, 0x41,
	0x23, 0x3d, 0x2b, 0x07, 0xfc, 0x54, 0x89, 0xb7, 0x74, 0x5d, 0xe2, 0x23, 0x29, 0xfe, 0x09, 0x74,
	0x2f, 0xe4, 0xb3, 0x19, 0xf5, 0x9c, 0x0f, 0x51, 0x67, 0xcc, 0x63, 0x26, 0xf4, 0x9c, 0xa1, 0x95,
	0xcf, 0xc9, 0xe5, 0x9a, 0x5b, 0x5b, 0xdc, 0x48, 0x28, 0xbe, 0x80, 0x93, 0x3f, 0xa7, 0xee, 0x24,
	0xd2, 0x4f, 0x0d, 0xad, 0x9c, 0x23, 0xab, 0x0d, 0x7e, 0x01, 0xa7, 0x22, 0x74, 0xc7, 0xd4, 0xf1,
	0x3d, 0x3d, 0x6f, 0x68, 0xe5, 0x02, 0xc9, 0xc9, 0x7d, 0xdb, 0xc3, 0x57, 0x90, 0x8b, 0x66, 0x2e,
	0x4b, 0x08, 0x48, 0x92, 0x4d, 0xb6, 0x6d, 0xef, 0xe6, 0xaf, 0x13, 0x28, 0xee, 0x4e, 0x19, 0xbf,
	0x84, 0xeb, 0xbe, 0xfd, 0xde, 0x26, 0xed, 0xc1, 0xef, 0x4e, 0x77, 0xf8, 0xee, 0xce, 0x26, 0xce,
	0xb0, 0xdb, 0x7f, 0xb0, 0x1b, 0xed, 0x66, 0xdb, 0xbe, 0x47, 0x9f, 0xe0, 0x17, 0xf0, 0x7c, 0x3f,
	0x61, 0x40, 0xac, 0x86, 0x8d, 0x34, 0x5c, 0x82, 0xcb, 0x54, 0x64, 0xa2, 0x23, 0x25, 0xab, 0xa1,
	0x63, 0x25, 0xab, 0xa3, 0x4c, 0xda, 0x71, 0xf7, 0xf6, 0xdd, 0xb0, 0x85, 0x4e, 0xd2, 0x64, 0x12,
	0x99, 0x28, 0xab, 0x64, 0x35, 0x94, 0x53, 0xb2, 0x3a, 0x3a, 0xc5, 0x3a, 0x5c, 0xec, 0xb3, 0x76,
	0xb7, 0xd9, 0x43, 0xf9, 0x34, 0x23, 0x09, 0x31, 0x11, 0xa8, 0x50, 0x0d, 0x9d, 0xa9, 0x50, 0x1d,
	0x15, 0xd2, 0x8e, 0xfa, 0xd5, 0x22, 0x5d, 0x74, 0x9e, 0x26, 0x4a, 0x88, 0x89, 0x8a, 0x2a, 0x54,
	0x43, 0xcf, 0x54, 0xa8, 0x8e, 0x50, 0x1a, 0xb2, 0x09, 0xe9, 0x11, 0xf4, 0x69, 0xda, 0x30, 0x24,
	0x32, 0x11, 0x56, 0xb2, 0x1a, 0xfa, 0x4c, 0xc9, 0xea, 0xe8, 0x22, 0xed, 0xb8, 0xa6, 0x35, 0xb0,
	0x3a, 0xe8, 0x79, 0x9a, 0x4c, 0x22, 0x13, 0x5d, 0x2a, 0x59, 0x0d, 0x5d, 0x29, 0x59, 0x1d, 0xe9,
	0x37, 0xbf, 0x41, 0x71, 0x7b, 0x23, 0x35, 0xe5, 0xbb, 0xf0, 0x12, 0xae, 0x3b, 0xbd, 0x96, 0x43,
	0xec, 0x46, 0x8f, 0xdc, 0x3b, 0xcd, 0x8e, 0xd5, 0xda, 0x7b, 0x88, 0xbf, 0x01, 0x63, 0x3f, 0x41,
	0x3e, 0x71, 0x72, 0xd9, 0x77, 0xde, 0x59, 0xfd, 0xb7, 0xe8, 0x3f, 0xed, 0xce, 0x85, 0x2f, 0x7d,
	0x7e, 0xe8, 0x8e, 0xba, 0x4b, 0xee, 0xc2, 0xe8, 0x21, 0x09, 0x3d, 0x68, 0x7f, 0x98, 0x13, 0x5f,
	0x3c, 0xc6, 0xa3, 0xe4, 0xcd, 0xaf, 0x7a, 0x81, 0x2b, 0x1e, 0x7d, 0x1a, 0x57, 0x77, 0xd4, 0x6f,
	0xa4, 0xfa, 0xcd, 0x64, 0xfb, 0xbd, 0x1e, 0x65, 0x65, 0xa4, 0xf6, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xf5, 0x75, 0xef, 0xd5, 0x07, 0x00, 0x00,
}