// Code generated by protoc-gen-gogo.
// source: http.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// / Type of peer handling request.
type PeerType int32

const (
	PeerType_Client PeerType = 1
	PeerType_Server PeerType = 2
)

var PeerType_name = map[int32]string{
	1: "Client",
	2: "Server",
}
var PeerType_value = map[string]int32{
	"Client": 1,
	"Server": 2,
}

func (x PeerType) Enum() *PeerType {
	p := new(PeerType)
	*p = x
	return p
}
func (x PeerType) String() string {
	return proto.EnumName(PeerType_name, int32(x))
}
func (x *PeerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PeerType_value, data, "PeerType")
	if err != nil {
		return err
	}
	*x = PeerType(value)
	return nil
}
func (PeerType) EnumDescriptor() ([]byte, []int) { return fileDescriptorHttp, []int{0} }

// / HTTP method.
type Method int32

const (
	Method_GET               Method = 1
	Method_POST              Method = 2
	Method_PUT               Method = 3
	Method_DELETE            Method = 4
	Method_HEAD              Method = 5
	Method_ACL               Method = 6
	Method_BASELINE_CONTROL  Method = 7
	Method_BIND              Method = 8
	Method_CHECKIN           Method = 9
	Method_CHECKOUT          Method = 10
	Method_CONNECT           Method = 11
	Method_COPY              Method = 12
	Method_DEBUG             Method = 13
	Method_LABEL             Method = 14
	Method_LINK              Method = 15
	Method_LOCK              Method = 16
	Method_MERGE             Method = 17
	Method_MKACTIVITY        Method = 18
	Method_MKCALENDAR        Method = 19
	Method_MKCOL             Method = 20
	Method_MKREDIRECTREF     Method = 21
	Method_MKWORKSPACE       Method = 22
	Method_MOVE              Method = 23
	Method_OPTIONS           Method = 24
	Method_ORDERPATCH        Method = 25
	Method_PATCH             Method = 26
	Method_PRI               Method = 27
	Method_PROPFIND          Method = 28
	Method_PROPPATCH         Method = 29
	Method_REBIND            Method = 30
	Method_REPORT            Method = 31
	Method_SEARCH            Method = 32
	Method_SHOWMETHOD        Method = 33
	Method_SPACEJUMP         Method = 34
	Method_TEXTSEARCH        Method = 35
	Method_TRACE             Method = 36
	Method_TRACK             Method = 37
	Method_UNBIND            Method = 38
	Method_UNCHECKOUT        Method = 39
	Method_UNLINK            Method = 40
	Method_UNLOCK            Method = 41
	Method_UPDATE            Method = 42
	Method_UPDATEREDIRECTREF Method = 43
	Method_VERSION_CONTROL   Method = 44
)

var Method_name = map[int32]string{
	1:  "GET",
	2:  "POST",
	3:  "PUT",
	4:  "DELETE",
	5:  "HEAD",
	6:  "ACL",
	7:  "BASELINE_CONTROL",
	8:  "BIND",
	9:  "CHECKIN",
	10: "CHECKOUT",
	11: "CONNECT",
	12: "COPY",
	13: "DEBUG",
	14: "LABEL",
	15: "LINK",
	16: "LOCK",
	17: "MERGE",
	18: "MKACTIVITY",
	19: "MKCALENDAR",
	20: "MKCOL",
	21: "MKREDIRECTREF",
	22: "MKWORKSPACE",
	23: "MOVE",
	24: "OPTIONS",
	25: "ORDERPATCH",
	26: "PATCH",
	27: "PRI",
	28: "PROPFIND",
	29: "PROPPATCH",
	30: "REBIND",
	31: "REPORT",
	32: "SEARCH",
	33: "SHOWMETHOD",
	34: "SPACEJUMP",
	35: "TEXTSEARCH",
	36: "TRACE",
	37: "TRACK",
	38: "UNBIND",
	39: "UNCHECKOUT",
	40: "UNLINK",
	41: "UNLOCK",
	42: "UPDATE",
	43: "UPDATEREDIRECTREF",
	44: "VERSION_CONTROL",
}
var Method_value = map[string]int32{
	"GET":               1,
	"POST":              2,
	"PUT":               3,
	"DELETE":            4,
	"HEAD":              5,
	"ACL":               6,
	"BASELINE_CONTROL":  7,
	"BIND":              8,
	"CHECKIN":           9,
	"CHECKOUT":          10,
	"CONNECT":           11,
	"COPY":              12,
	"DEBUG":             13,
	"LABEL":             14,
	"LINK":              15,
	"LOCK":              16,
	"MERGE":             17,
	"MKACTIVITY":        18,
	"MKCALENDAR":        19,
	"MKCOL":             20,
	"MKREDIRECTREF":     21,
	"MKWORKSPACE":       22,
	"MOVE":              23,
	"OPTIONS":           24,
	"ORDERPATCH":        25,
	"PATCH":             26,
	"PRI":               27,
	"PROPFIND":          28,
	"PROPPATCH":         29,
	"REBIND":            30,
	"REPORT":            31,
	"SEARCH":            32,
	"SHOWMETHOD":        33,
	"SPACEJUMP":         34,
	"TEXTSEARCH":        35,
	"TRACE":             36,
	"TRACK":             37,
	"UNBIND":            38,
	"UNCHECKOUT":        39,
	"UNLINK":            40,
	"UNLOCK":            41,
	"UPDATE":            42,
	"UPDATEREDIRECTREF": 43,
	"VERSION_CONTROL":   44,
}

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}
func (x Method) String() string {
	return proto.EnumName(Method_name, int32(x))
}
func (x *Method) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Method_value, data, "Method")
	if err != nil {
		return err
	}
	*x = Method(value)
	return nil
}
func (Method) EnumDescriptor() ([]byte, []int) { return fileDescriptorHttp, []int{1} }

// / An HttpStartStop event represents the whole lifecycle of an HTTP request.
type HttpStartStop struct {
	StartTimestamp   *int64    `protobuf:"varint,1,req,name=startTimestamp" json:"startTimestamp,omitempty"`
	StopTimestamp    *int64    `protobuf:"varint,2,req,name=stopTimestamp" json:"stopTimestamp,omitempty"`
	RequestId        *UUID     `protobuf:"bytes,3,req,name=requestId" json:"requestId,omitempty"`
	PeerType         *PeerType `protobuf:"varint,4,req,name=peerType,enum=events.PeerType" json:"peerType,omitempty"`
	Method           *Method   `protobuf:"varint,5,req,name=method,enum=events.Method" json:"method,omitempty"`
	Uri              *string   `protobuf:"bytes,6,req,name=uri" json:"uri,omitempty"`
	RemoteAddress    *string   `protobuf:"bytes,7,req,name=remoteAddress" json:"remoteAddress,omitempty"`
	UserAgent        *string   `protobuf:"bytes,8,req,name=userAgent" json:"userAgent,omitempty"`
	StatusCode       *int32    `protobuf:"varint,9,req,name=statusCode" json:"statusCode,omitempty"`
	ContentLength    *int64    `protobuf:"varint,10,req,name=contentLength" json:"contentLength,omitempty"`
	ApplicationId    *UUID     `protobuf:"bytes,12,opt,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32    `protobuf:"varint,13,opt,name=instanceIndex" json:"instanceIndex,omitempty"`
	InstanceId       *string   `protobuf:"bytes,14,opt,name=instanceId" json:"instanceId,omitempty"`
	Forwarded        []string  `protobuf:"bytes,15,rep,name=forwarded" json:"forwarded,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *HttpStartStop) Reset()                    { *m = HttpStartStop{} }
func (m *HttpStartStop) String() string            { return proto.CompactTextString(m) }
func (*HttpStartStop) ProtoMessage()               {}
func (*HttpStartStop) Descriptor() ([]byte, []int) { return fileDescriptorHttp, []int{0} }

func (m *HttpStartStop) GetStartTimestamp() int64 {
	if m != nil && m.StartTimestamp != nil {
		return *m.StartTimestamp
	}
	return 0
}

func (m *HttpStartStop) GetStopTimestamp() int64 {
	if m != nil && m.StopTimestamp != nil {
		return *m.StopTimestamp
	}
	return 0
}

func (m *HttpStartStop) GetRequestId() *UUID {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *HttpStartStop) GetPeerType() PeerType {
	if m != nil && m.PeerType != nil {
		return *m.PeerType
	}
	return PeerType_Client
}

func (m *HttpStartStop) GetMethod() Method {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return Method_GET
}

func (m *HttpStartStop) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *HttpStartStop) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *HttpStartStop) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *HttpStartStop) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *HttpStartStop) GetContentLength() int64 {
	if m != nil && m.ContentLength != nil {
		return *m.ContentLength
	}
	return 0
}

func (m *HttpStartStop) GetApplicationId() *UUID {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *HttpStartStop) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *HttpStartStop) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *HttpStartStop) GetForwarded() []string {
	if m != nil {
		return m.Forwarded
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpStartStop)(nil), "events.HttpStartStop")
	proto.RegisterEnum("events.PeerType", PeerType_name, PeerType_value)
	proto.RegisterEnum("events.Method", Method_name, Method_value)
}
func (m *HttpStartStop) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *HttpStartStop) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StartTimestamp == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("startTimestamp")
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintHttp(data, i, uint64(*m.StartTimestamp))
	}
	if m.StopTimestamp == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("stopTimestamp")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintHttp(data, i, uint64(*m.StopTimestamp))
	}
	if m.RequestId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("requestId")
	} else {
		data[i] = 0x1a
		i++
		i = encodeVarintHttp(data, i, uint64(m.RequestId.Size()))
		n1, err := m.RequestId.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PeerType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("peerType")
	} else {
		data[i] = 0x20
		i++
		i = encodeVarintHttp(data, i, uint64(*m.PeerType))
	}
	if m.Method == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("method")
	} else {
		data[i] = 0x28
		i++
		i = encodeVarintHttp(data, i, uint64(*m.Method))
	}
	if m.Uri == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("uri")
	} else {
		data[i] = 0x32
		i++
		i = encodeVarintHttp(data, i, uint64(len(*m.Uri)))
		i += copy(data[i:], *m.Uri)
	}
	if m.RemoteAddress == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("remoteAddress")
	} else {
		data[i] = 0x3a
		i++
		i = encodeVarintHttp(data, i, uint64(len(*m.RemoteAddress)))
		i += copy(data[i:], *m.RemoteAddress)
	}
	if m.UserAgent == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("userAgent")
	} else {
		data[i] = 0x42
		i++
		i = encodeVarintHttp(data, i, uint64(len(*m.UserAgent)))
		i += copy(data[i:], *m.UserAgent)
	}
	if m.StatusCode == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("statusCode")
	} else {
		data[i] = 0x48
		i++
		i = encodeVarintHttp(data, i, uint64(*m.StatusCode))
	}
	if m.ContentLength == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("contentLength")
	} else {
		data[i] = 0x50
		i++
		i = encodeVarintHttp(data, i, uint64(*m.ContentLength))
	}
	if m.ApplicationId != nil {
		data[i] = 0x62
		i++
		i = encodeVarintHttp(data, i, uint64(m.ApplicationId.Size()))
		n2, err := m.ApplicationId.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.InstanceIndex != nil {
		data[i] = 0x68
		i++
		i = encodeVarintHttp(data, i, uint64(*m.InstanceIndex))
	}
	if m.InstanceId != nil {
		data[i] = 0x72
		i++
		i = encodeVarintHttp(data, i, uint64(len(*m.InstanceId)))
		i += copy(data[i:], *m.InstanceId)
	}
	if len(m.Forwarded) > 0 {
		for _, s := range m.Forwarded {
			data[i] = 0x7a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Http(data []byte, offset int, v uint64) int {
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
func encodeFixed32Http(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHttp(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *HttpStartStop) Size() (n int) {
	var l int
	_ = l
	if m.StartTimestamp != nil {
		n += 1 + sovHttp(uint64(*m.StartTimestamp))
	}
	if m.StopTimestamp != nil {
		n += 1 + sovHttp(uint64(*m.StopTimestamp))
	}
	if m.RequestId != nil {
		l = m.RequestId.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.PeerType != nil {
		n += 1 + sovHttp(uint64(*m.PeerType))
	}
	if m.Method != nil {
		n += 1 + sovHttp(uint64(*m.Method))
	}
	if m.Uri != nil {
		l = len(*m.Uri)
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.RemoteAddress != nil {
		l = len(*m.RemoteAddress)
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.UserAgent != nil {
		l = len(*m.UserAgent)
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.StatusCode != nil {
		n += 1 + sovHttp(uint64(*m.StatusCode))
	}
	if m.ContentLength != nil {
		n += 1 + sovHttp(uint64(*m.ContentLength))
	}
	if m.ApplicationId != nil {
		l = m.ApplicationId.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.InstanceIndex != nil {
		n += 1 + sovHttp(uint64(*m.InstanceIndex))
	}
	if m.InstanceId != nil {
		l = len(*m.InstanceId)
		n += 1 + l + sovHttp(uint64(l))
	}
	if len(m.Forwarded) > 0 {
		for _, s := range m.Forwarded {
			l = len(s)
			n += 1 + l + sovHttp(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttp(x uint64) (n int) {
	return sovHttp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpStartStop) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
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
			return fmt.Errorf("proto: HttpStartStop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpStartStop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestamp", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StartTimestamp = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StopTimestamp", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StopTimestamp = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestId == nil {
				m.RequestId = &UUID{}
			}
			if err := m.RequestId.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			var v PeerType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (PeerType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PeerType = &v
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var v Method
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (Method(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Method = &v
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Uri = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000020)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.RemoteAddress = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000040)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAgent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.UserAgent = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000080)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StatusCode = &v
			hasFields[0] |= uint64(0x00000100)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentLength", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ContentLength = &v
			hasFields[0] |= uint64(0x00000200)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApplicationId == nil {
				m.ApplicationId = &UUID{}
			}
			if err := m.ApplicationId.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceIndex", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InstanceIndex = &v
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.InstanceId = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forwarded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Forwarded = append(m.Forwarded, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("startTimestamp")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("stopTimestamp")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("requestId")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("peerType")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("method")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("uri")
	}
	if hasFields[0]&uint64(0x00000040) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("remoteAddress")
	}
	if hasFields[0]&uint64(0x00000080) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("userAgent")
	}
	if hasFields[0]&uint64(0x00000100) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("statusCode")
	}
	if hasFields[0]&uint64(0x00000200) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("contentLength")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttp(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttp
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
					return 0, ErrIntOverflowHttp
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
					return 0, ErrIntOverflowHttp
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
				return 0, ErrInvalidLengthHttp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttp
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
				next, err := skipHttp(data[start:])
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
	ErrInvalidLengthHttp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("http.proto", fileDescriptorHttp) }

var fileDescriptorHttp = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x54, 0x5d, 0x73, 0xdb, 0x44,
	0x14, 0x9d, 0xc4, 0x1f, 0xb1, 0xd7, 0xb1, 0x73, 0xb3, 0x6d, 0x41, 0x84, 0x12, 0xd2, 0x50, 0x4a,
	0x09, 0xc5, 0x9d, 0xc9, 0x33, 0x2f, 0xb2, 0xb4, 0x89, 0x85, 0x6c, 0x49, 0xb3, 0x5a, 0xa7, 0xf4,
	0x89, 0x71, 0xad, 0x8d, 0xe3, 0x99, 0xd8, 0x2b, 0xf4, 0x51, 0xc8, 0xef, 0xe0, 0x3f, 0x31, 0x3c,
	0xf2, 0x13, 0x18, 0x7e, 0x09, 0x7b, 0x57, 0x8e, 0x13, 0xc3, 0x83, 0x66, 0xce, 0x3d, 0xe7, 0xec,
	0xbd, 0x77, 0xcf, 0x26, 0x26, 0xe4, 0xa6, 0x28, 0xd2, 0x7e, 0x9a, 0xa9, 0x42, 0xd1, 0xa6, 0xfc,
	0x28, 0x57, 0x45, 0x7e, 0xf4, 0xfd, 0x7c, 0x51, 0xdc, 0x94, 0x1f, 0xfa, 0x33, 0xb5, 0x7c, 0x3b,
	0x57, 0x73, 0xf5, 0xd6, 0xc8, 0x1f, 0xca, 0x6b, 0x53, 0x99, 0xc2, 0xa0, 0xea, 0xd8, 0x11, 0x29,
	0xcb, 0x45, 0x52, 0xe1, 0xd3, 0xdf, 0xeb, 0xa4, 0x3b, 0xd4, 0x1d, 0xe3, 0x62, 0x9a, 0x15, 0x71,
	0xa1, 0x52, 0xfa, 0x8a, 0xf4, 0x72, 0x2c, 0xc4, 0x62, 0x29, 0x35, 0x58, 0xa6, 0xd6, 0xce, 0xc9,
	0xee, 0xeb, 0x1a, 0xff, 0x0f, 0x4b, 0x5f, 0x92, 0x6e, 0xae, 0xfd, 0x0f, 0xb6, 0x5d, 0x63, 0xdb,
	0x26, 0xe9, 0x19, 0x69, 0x67, 0xf2, 0x97, 0x52, 0x57, 0x5e, 0x62, 0xd5, 0xb4, 0xa3, 0x73, 0xbe,
	0xdf, 0xaf, 0xd6, 0xee, 0x4f, 0x26, 0x9e, 0xcb, 0x1f, 0x64, 0xfa, 0x86, 0xb4, 0x52, 0x29, 0x33,
	0x71, 0x97, 0x4a, 0xab, 0xae, 0xad, 0xbd, 0x73, 0xb8, 0xb7, 0x46, 0x6b, 0x9e, 0x6f, 0x1c, 0x7a,
	0xcf, 0xe6, 0x52, 0x16, 0x37, 0x2a, 0xb1, 0x1a, 0xc6, 0xdb, 0xbb, 0xf7, 0x8e, 0x0d, 0xcb, 0xd7,
	0x2a, 0x05, 0x52, 0x2b, 0xb3, 0x85, 0xd5, 0xd4, 0xa6, 0x36, 0x47, 0x88, 0x9b, 0x67, 0x72, 0xa9,
	0x0a, 0x69, 0x27, 0x49, 0x26, 0xf3, 0xdc, 0xda, 0x33, 0xda, 0x36, 0x49, 0x9f, 0x93, 0x76, 0x99,
	0xcb, 0xcc, 0x9e, 0xeb, 0xa6, 0x56, 0xcb, 0x38, 0x1e, 0x08, 0x7a, 0x4c, 0x88, 0xbe, 0x60, 0x51,
	0xe6, 0x8e, 0x4a, 0xa4, 0xd5, 0xd6, 0x72, 0x83, 0x3f, 0x62, 0x70, 0xc6, 0x4c, 0xad, 0x0a, 0x6d,
	0x1d, 0xc9, 0xd5, 0xbc, 0xb8, 0xb1, 0x48, 0x95, 0xce, 0x16, 0x49, 0xcf, 0x49, 0x77, 0x9a, 0xa6,
	0xb7, 0x8b, 0xd9, 0xb4, 0x58, 0xa8, 0x95, 0x4e, 0x68, 0xff, 0x64, 0xe7, 0x7f, 0x09, 0x6d, 0x5b,
	0xb0, 0xf3, 0x62, 0xa5, 0x27, 0xad, 0x66, 0xd2, 0x5b, 0x25, 0xf2, 0x37, 0xab, 0xab, 0xcf, 0x34,
	0xf8, 0x36, 0x89, 0xfb, 0x6d, 0x88, 0xc4, 0xea, 0x69, 0x4b, 0x9b, 0x3f, 0x62, 0xf0, 0x76, 0xd7,
	0x2a, 0xfb, 0x75, 0x9a, 0x25, 0x32, 0xb1, 0x0e, 0x4e, 0x6a, 0x78, 0xbb, 0x0d, 0x71, 0x76, 0x4a,
	0x5a, 0xf7, 0x89, 0x53, 0x42, 0x9a, 0xce, 0xed, 0x42, 0xaf, 0x03, 0x3b, 0x88, 0x63, 0x99, 0x7d,
	0x94, 0x19, 0xec, 0x9e, 0xfd, 0x51, 0x27, 0xcd, 0x2a, 0x6a, 0xba, 0x47, 0x6a, 0x97, 0x4c, 0x68,
	0xbd, 0x45, 0xea, 0x51, 0x18, 0x0b, 0xd8, 0x45, 0x2a, 0x9a, 0x08, 0xa8, 0xe1, 0x11, 0x97, 0x8d,
	0x98, 0x60, 0x50, 0x47, 0x79, 0xc8, 0x6c, 0x17, 0x1a, 0x28, 0xdb, 0xce, 0x08, 0x9a, 0xf4, 0x29,
	0x81, 0x81, 0x1d, 0xb3, 0x91, 0x17, 0xb0, 0x9f, 0x9d, 0x30, 0x10, 0x3c, 0x1c, 0xc1, 0x1e, 0x1a,
	0x07, 0x5e, 0xe0, 0x42, 0x8b, 0x76, 0xc8, 0x9e, 0x33, 0x64, 0x8e, 0xef, 0x05, 0xd0, 0xa6, 0xfb,
	0xa4, 0x65, 0x8a, 0x50, 0x77, 0x26, 0x46, 0x0a, 0x83, 0x80, 0x39, 0x02, 0x3a, 0x78, 0xc2, 0x09,
	0xa3, 0xf7, 0xb0, 0x4f, 0xdb, 0xa4, 0xe1, 0xb2, 0xc1, 0xe4, 0x12, 0xba, 0x08, 0x47, 0xf6, 0x80,
	0x8d, 0xa0, 0x87, 0xba, 0x9e, 0xe1, 0xc3, 0x81, 0x41, 0xa1, 0xe3, 0x03, 0xa0, 0x3c, 0x66, 0xfc,
	0x92, 0xc1, 0x21, 0xed, 0x11, 0x32, 0xf6, 0x6d, 0x47, 0x78, 0x57, 0x9e, 0x78, 0x0f, 0xb4, 0xaa,
	0x1d, 0x7b, 0xc4, 0x02, 0xd7, 0xe6, 0xf0, 0xc4, 0x58, 0x7d, 0x47, 0xef, 0xf6, 0x94, 0x1e, 0x92,
	0xee, 0xd8, 0xe7, 0xcc, 0xf5, 0xb8, 0x9e, 0xcc, 0xd9, 0x05, 0x3c, 0xa3, 0x07, 0xa4, 0x33, 0xf6,
	0xdf, 0x85, 0xdc, 0x8f, 0x23, 0xdb, 0x61, 0xf0, 0x09, 0xce, 0x18, 0x87, 0x57, 0x0c, 0x3e, 0xc5,
	0x25, 0xc3, 0x48, 0x78, 0x61, 0x10, 0x83, 0x85, 0x5d, 0x43, 0xee, 0x32, 0x1e, 0xd9, 0xc2, 0x19,
	0xc2, 0x67, 0xd8, 0xb5, 0x82, 0x47, 0x26, 0x2f, 0xee, 0xc1, 0xe7, 0x78, 0xc7, 0x88, 0x87, 0xd1,
	0x05, 0x5e, 0xff, 0x39, 0xed, 0x92, 0x36, 0x56, 0x95, 0xeb, 0x0b, 0x0c, 0x93, 0x33, 0x93, 0xcc,
	0x71, 0x85, 0xa3, 0x90, 0x0b, 0xf8, 0xd2, 0xbc, 0x0b, 0xb3, 0xb9, 0xf6, 0x9c, 0xe0, 0x90, 0x78,
	0x18, 0xbe, 0x1b, 0x33, 0x31, 0x0c, 0x5d, 0x78, 0x81, 0x2d, 0xcc, 0x5a, 0x3f, 0x4e, 0xc6, 0x11,
	0x9c, 0xa2, 0x2c, 0xd8, 0x4f, 0x62, 0x6d, 0xff, 0x0a, 0x77, 0x10, 0x1c, 0xb7, 0x7e, 0x79, 0x0f,
	0x7d, 0xf8, 0x1a, 0x1b, 0x4e, 0x02, 0x33, 0xe8, 0x15, 0x9e, 0x98, 0x04, 0x9b, 0xdc, 0xbf, 0xa9,
	0x34, 0x13, 0xe6, 0xeb, 0x35, 0xc6, 0x38, 0xbf, 0x35, 0x38, 0x72, 0x6d, 0xfd, 0xd2, 0x67, 0xf4,
	0x19, 0x39, 0xac, 0xf0, 0xe3, 0xa0, 0xbe, 0xa3, 0x4f, 0xc8, 0xc1, 0x15, 0xe3, 0xb1, 0x8e, 0x63,
	0xf3, 0xd8, 0x6f, 0x06, 0x3f, 0xfc, 0xf9, 0xcf, 0xf1, 0xce, 0x5f, 0xfa, 0xfb, 0x5b, 0x7f, 0xe4,
	0x85, 0xca, 0xe6, 0xfd, 0xd9, 0xad, 0x2a, 0x93, 0x6b, 0x55, 0xae, 0x92, 0xec, 0xae, 0x9f, 0x64,
	0x2a, 0xcd, 0x95, 0xfe, 0xab, 0x5e, 0xff, 0x57, 0x0c, 0x3a, 0xf8, 0x83, 0x75, 0x31, 0x9d, 0x15,
	0x2a, 0xbb, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xdc, 0x82, 0x50, 0x10, 0x05, 0x00, 0x00,
}
