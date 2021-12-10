package core

import (
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Any", &AnyCodec{})
	jsoniter.RegisterTypeEncoder("core.Any", &AnyCodec{})
}

type AnyCodec struct {
}

func (codec *AnyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	anyVal := (*Any)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		//keys := any.Keys()
		anyVal.Type = any.Get("@type").ToString()

		msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(anyVal.Type))
		if err != nil {
			iter.ReportError("Any Decode", err.Error())
			return
		}
		msg := proto.MessageV1(msgType.New())
		any.ToVal(msg)
		anyVal.typeVal = msg
	}
}

func (codec *AnyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*Any)(ptr).Empty()
}

func (codec *AnyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	kvWriter := func(t string, v interface{}) {
		stream.WriteObjectStart()
		stream.WriteObjectField("@type")
		stream.WriteVal(t)
		stream.WriteMore()
		if _, ok := v.(ArrayType); ok {
			stream.WriteObjectField("values")
		} else {
			stream.WriteObjectField("value")
		}
		stream.WriteVal(v)
		stream.WriteObjectEnd()
	}

	any := (*Any)(ptr)
	switch v := any.typeVal.(type) {
	case proto.Message:
		reflectMsg := proto.MessageReflect(v)

		if value, ok := any.typeVal.(ScalarType); ok {
			kvWriter(string(reflectMsg.Descriptor().FullName()), value.AsScalarType())
		} else {
			stream.WriteObjectStart()
			stream.WriteObjectField("@type")
			stream.WriteVal(reflectMsg.Descriptor().FullName())

			reflectMsg.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
				stream.WriteMore()
				stream.WriteObjectField(descriptor.JSONName())
				stream.WriteVal(value.Interface())
				return true
			})

			stream.WriteObjectEnd()
		}
	default:
		kvWriter(GetMojoTypeName(any.typeVal), v)
	}
}
