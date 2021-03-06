// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/kafka.proto

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// KafkaMessage is a wrapper type for the messages
// that the Kafka-based orderer deals with.
type KafkaMessage struct {
	// Types that are valid to be assigned to Type:
	//	*KafkaMessage_Regular
	//	*KafkaMessage_TimeToCut
	//	*KafkaMessage_Connect
	Type isKafkaMessage_Type `protobuf_oneof:"Type"`
}

func (m *KafkaMessage) Reset()                    { *m = KafkaMessage{} }
func (m *KafkaMessage) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessage) ProtoMessage()               {}
func (*KafkaMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isKafkaMessage_Type interface {
	isKafkaMessage_Type()
}

type KafkaMessage_Regular struct {
	Regular *KafkaMessageRegular `protobuf:"bytes,1,opt,name=regular,oneof"`
}
type KafkaMessage_TimeToCut struct {
	TimeToCut *KafkaMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,oneof"`
}
type KafkaMessage_Connect struct {
	Connect *KafkaMessageConnect `protobuf:"bytes,3,opt,name=connect,oneof"`
}

func (*KafkaMessage_Regular) isKafkaMessage_Type()   {}
func (*KafkaMessage_TimeToCut) isKafkaMessage_Type() {}
func (*KafkaMessage_Connect) isKafkaMessage_Type()   {}

func (m *KafkaMessage) GetType() isKafkaMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *KafkaMessage) GetRegular() *KafkaMessageRegular {
	if x, ok := m.GetType().(*KafkaMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (m *KafkaMessage) GetTimeToCut() *KafkaMessageTimeToCut {
	if x, ok := m.GetType().(*KafkaMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (m *KafkaMessage) GetConnect() *KafkaMessageConnect {
	if x, ok := m.GetType().(*KafkaMessage_Connect); ok {
		return x.Connect
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KafkaMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KafkaMessage_OneofMarshaler, _KafkaMessage_OneofUnmarshaler, _KafkaMessage_OneofSizer, []interface{}{
		(*KafkaMessage_Regular)(nil),
		(*KafkaMessage_TimeToCut)(nil),
		(*KafkaMessage_Connect)(nil),
	}
}

func _KafkaMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regular); err != nil {
			return err
		}
	case *KafkaMessage_TimeToCut:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeToCut); err != nil {
			return err
		}
	case *KafkaMessage_Connect:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Connect); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KafkaMessage.Type has unexpected type %T", x)
	}
	return nil
}

func _KafkaMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KafkaMessage)
	switch tag {
	case 1: // Type.regular
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageRegular)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Regular{msg}
		return true, err
	case 2: // Type.time_to_cut
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageTimeToCut)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_TimeToCut{msg}
		return true, err
	case 3: // Type.connect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageConnect)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Connect{msg}
		return true, err
	default:
		return false, nil
	}
}

func _KafkaMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		s := proto.Size(x.Regular)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_TimeToCut:
		s := proto.Size(x.TimeToCut)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_Connect:
		s := proto.Size(x.Connect)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KafkaMessageRegular wraps a marshalled envelope.
type KafkaMessageRegular struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *KafkaMessageRegular) Reset()                    { *m = KafkaMessageRegular{} }
func (m *KafkaMessageRegular) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageRegular) ProtoMessage()               {}
func (*KafkaMessageRegular) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *KafkaMessageRegular) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// KafkaMessageTimeToCut is used to signal to the orderers
// that it is time to cut block <block_number>.
type KafkaMessageTimeToCut struct {
	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber" json:"block_number,omitempty"`
}

func (m *KafkaMessageTimeToCut) Reset()                    { *m = KafkaMessageTimeToCut{} }
func (m *KafkaMessageTimeToCut) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageTimeToCut) ProtoMessage()               {}
func (*KafkaMessageTimeToCut) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *KafkaMessageTimeToCut) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// KafkaMessageConnect is posted by an orderer upon booting up.
// It is used to prevent the panic that would be caused if we
// were to consume an empty partition. It is ignored by all
// orderers when processing the partition.
type KafkaMessageConnect struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *KafkaMessageConnect) Reset()                    { *m = KafkaMessageConnect{} }
func (m *KafkaMessageConnect) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageConnect) ProtoMessage()               {}
func (*KafkaMessageConnect) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *KafkaMessageConnect) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// LastOffsetPersisted is the encoded value for the Metadata message
// which is encoded in the ORDERER block metadata index for the case
// of the Kafka-based orderer.
type KafkaMetadata struct {
	LastOffsetPersisted int64 `protobuf:"varint,1,opt,name=last_offset_persisted,json=lastOffsetPersisted" json:"last_offset_persisted,omitempty"`
}

func (m *KafkaMetadata) Reset()                    { *m = KafkaMetadata{} }
func (m *KafkaMetadata) String() string            { return proto.CompactTextString(m) }
func (*KafkaMetadata) ProtoMessage()               {}
func (*KafkaMetadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *KafkaMetadata) GetLastOffsetPersisted() int64 {
	if m != nil {
		return m.LastOffsetPersisted
	}
	return 0
}

func init() {
	proto.RegisterType((*KafkaMessage)(nil), "orderer.KafkaMessage")
	proto.RegisterType((*KafkaMessageRegular)(nil), "orderer.KafkaMessageRegular")
	proto.RegisterType((*KafkaMessageTimeToCut)(nil), "orderer.KafkaMessageTimeToCut")
	proto.RegisterType((*KafkaMessageConnect)(nil), "orderer.KafkaMessageConnect")
	proto.RegisterType((*KafkaMetadata)(nil), "orderer.KafkaMetadata")
}

func init() { proto.RegisterFile("orderer/kafka.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3d, 0x6b, 0xf3, 0x40,
	0x10, 0x84, 0xed, 0xd7, 0xc6, 0xe6, 0x3d, 0x3b, 0x8d, 0x8c, 0x41, 0x45, 0x08, 0x89, 0xab, 0x34,
	0x91, 0xc0, 0x49, 0x11, 0x52, 0x05, 0xbb, 0x31, 0x84, 0x7c, 0x20, 0x5c, 0xa5, 0x11, 0x2b, 0x69,
	0x25, 0x1f, 0x92, 0x6e, 0xc5, 0xdd, 0xaa, 0xf0, 0x7f, 0xcc, 0x8f, 0x0a, 0xfa, 0x38, 0x08, 0x89,
	0x71, 0xb9, 0x33, 0xf3, 0x30, 0x7b, 0x7b, 0x62, 0x41, 0x3a, 0x41, 0x8d, 0xda, 0xcf, 0x21, 0xcd,
	0xc1, 0xab, 0x34, 0x31, 0x39, 0xd3, 0x5e, 0x5c, 0x7d, 0x0d, 0xc5, 0xfc, 0xa5, 0x31, 0x5e, 0xd1,
	0x18, 0xc8, 0xd0, 0x79, 0x14, 0x53, 0x8d, 0x59, 0x5d, 0x80, 0x76, 0x87, 0xd7, 0xc3, 0xdb, 0xd9,
	0xfa, 0xd2, 0xeb, 0xb3, 0xde, 0xcf, 0x5c, 0xd0, 0x65, 0x76, 0x83, 0xc0, 0xc6, 0x9d, 0x67, 0x31,
	0x63, 0x59, 0x62, 0xc8, 0x14, 0xc6, 0x35, 0xbb, 0xff, 0x5a, 0xfa, 0xea, 0x24, 0xbd, 0x97, 0x25,
	0xee, 0x69, 0x5b, 0xf3, 0x6e, 0x10, 0xfc, 0x67, 0x3b, 0x34, 0xdd, 0x31, 0x29, 0x85, 0x31, 0xbb,
	0xa3, 0x33, 0xdd, 0xdb, 0x2e, 0xd3, 0x74, 0xf7, 0xf1, 0xcd, 0x44, 0x8c, 0xf7, 0xc7, 0x0a, 0x57,
	0xbe, 0x58, 0x9c, 0xd8, 0xd2, 0x71, 0xc5, 0xb4, 0x82, 0x63, 0x41, 0x90, 0xb4, 0x8f, 0x9a, 0x07,
	0x76, 0x5c, 0x3d, 0x89, 0xe5, 0xc9, 0xc5, 0x9c, 0x1b, 0x31, 0x8f, 0x0a, 0x8a, 0xf3, 0x50, 0xd5,
	0x65, 0x84, 0xdd, 0x31, 0xc6, 0xc1, 0xac, 0xd5, 0xde, 0x5a, 0xe9, 0x77, 0x59, 0xbf, 0xd6, 0x99,
	0xb2, 0xad, 0xb8, 0xe8, 0x01, 0x86, 0x04, 0x18, 0x9c, 0xb5, 0x58, 0x16, 0x60, 0x38, 0xa4, 0x34,
	0x35, 0xc8, 0x61, 0x85, 0xda, 0x48, 0xc3, 0xd8, 0x81, 0xa3, 0x60, 0xd1, 0x98, 0xef, 0xad, 0xf7,
	0x61, 0xad, 0x0d, 0x8a, 0x3b, 0xd2, 0x99, 0x27, 0x55, 0x5e, 0x40, 0x64, 0x52, 0xaa, 0x55, 0x02,
	0x2c, 0x49, 0x35, 0x4a, 0x7c, 0x00, 0xa9, 0xba, 0x2f, 0x36, 0xf6, 0x74, 0x9f, 0x0f, 0x99, 0xe4,
	0x43, 0x1d, 0x79, 0x31, 0x95, 0xfe, 0x1f, 0xca, 0xb7, 0x94, 0xdf, 0x51, 0x7e, 0x4f, 0x45, 0x93,
	0x76, 0xbe, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x93, 0x92, 0x93, 0x3f, 0x02, 0x00, 0x00,
}
