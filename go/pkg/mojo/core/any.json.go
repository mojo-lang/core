package core

import (
    "github.com/golang/protobuf/proto"
    jsoniter "github.com/json-iterator/go"
    "github.com/modern-go/reflect2"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/reflect/protoregistry"
    "reflect"
    "sort"
    "sync"
    "unsafe"
)

var anyFieldEncoders map[string]jsoniter.ValEncoder

func RegisterAnyFieldEncoder(typ string, field string, encoder jsoniter.ValEncoder) {
    (&sync.Once{}).Do(func() {
        anyFieldEncoders = make(map[string]jsoniter.ValEncoder)
    })

    anyFieldEncoders[typ+"."+field] = encoder
}

func init() {
    RegisterJSONTypeDecoder("core.Any", &AnyCodec{})
    RegisterJSONTypeEncoder("core.Any", &AnyCodec{})
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

type Fields []protoreflect.FieldDescriptor

func (f Fields) Len() int {
    return len(f)
}
func (f Fields) Swap(i, j int) {
    f[i], f[j] = f[j], f[i]
}
func (f Fields) Less(i, j int) bool {
    return f[i].JSONName() < f[j].JSONName()
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
            fullName := string(reflectMsg.Descriptor().FullName())

            stream.WriteObjectStart()
            stream.WriteObjectField("@type")
            stream.WriteVal(fullName)

            t := reflect2.TypeOf(any.typeVal)
            if t.Kind() == reflect.Ptr {
                t = t.(*reflect2.UnsafePtrType).Elem()
            }
            if obj, ok := t.(*reflect2.UnsafeStructType); ok {
                var fields Fields
                reflectMsg.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
                    fields = append(fields, descriptor)
                    return true
                })

                sort.Sort(fields)
                for _, descriptor := range fields {
                    stream.WriteMore()
                    stream.WriteObjectField(descriptor.JSONName())

                    fieldName := strcase.ToCamel(string(descriptor.Name()))
                    field := obj.FieldByName(fieldName)
                    if encoder, ok := anyFieldEncoders[t.String()+"."+fieldName]; ok {
                        f := field.UnsafeGet(reflect2.PtrOf(any.typeVal))
                        if !encoder.IsEmpty(f) {
                            encoder.Encode(f, stream)
                        }
                    } else {
                        stream.WriteVal(field.Get(any.typeVal))
                    }
                }
            } else {
                logs.Errorw("invalid the struct type for any object", "type", fullName)
            }

            stream.WriteObjectEnd()
        }
    default:
        kvWriter(GetMojoTypeFullName(any.typeVal), v)
    }
}
