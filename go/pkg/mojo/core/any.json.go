package core

import (
    "google.golang.org/protobuf/reflect/protoregistry"
    "reflect"
    "sort"
    "sync"
    "unsafe"

    jsoniter "github.com/json-iterator/go"
    "github.com/modern-go/reflect2"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
)

const (
    typeFieldName   = "@type"
    valueFieldName  = "value"
    valuesFieldName = "values"
)

func (x *Any) MarshalJSON() ([]byte, error) {
    return jsoniter.ConfigFastest.Marshal(x)
}

func (x *Any) UnmarshalJSON(err []byte) error {
    return jsoniter.ConfigFastest.Unmarshal(err, x)
}

var once = &sync.Once{}
var anyFieldEncoders map[string]jsoniter.ValEncoder

func RegisterAnyFieldEncoder(typ string, field string, encoder jsoniter.ValEncoder) {
    once.Do(func() {
        anyFieldEncoders = make(map[string]jsoniter.ValEncoder)
    })

    anyFieldEncoders[typ+"."+field] = encoder
}

func init() {
    RegisterJSONTypeDecoder("core.Any", &AnyCodec{})
    RegisterJSONTypeEncoder("core.Any", &AnyCodec{})
}

// AnyCodec
// FIXME should implement all the support types codec
type AnyCodec struct {
}

func (codec *AnyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    anyVal := (*Any)(ptr)
    if any.ValueType() == jsoniter.ObjectValue {
        anyVal.Type = any.Get(typeFieldName).ToString()

        switch anyVal.Type {
        case Int64TypeName:
            anyVal.typedVal = any.Get(valueFieldName).ToInt64()
        case StringTypeName:
            anyVal.typedVal = any.Get(valueFieldName).ToString()
        default:
            if msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(anyVal.Type)); err != nil {
                iter.ReportError("Any Decode", err.Error())
                return
            } else {
                msg := msgType.New().Interface()
                any.ToVal(msg)
                anyVal.typedVal = msg
            }
        }
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
        stream.WriteObjectField(typeFieldName)
        stream.WriteVal(t)
        stream.WriteMore()
        if _, ok := v.(ArrayType); ok {
            stream.WriteObjectField(valuesFieldName)
        } else {
            stream.WriteObjectField(valueFieldName)
        }
        stream.WriteVal(v)
        stream.WriteObjectEnd()
    }

    any := (*Any)(ptr)
    switch v := any.typedVal.(type) {
    case proto.Message:
        reflectMsg := v.ProtoReflect()

        if value, ok := any.typedVal.(ScalarType); ok {
            kvWriter(string(reflectMsg.Descriptor().FullName()), value.AsScalarType())
        } else {
            fullName := string(reflectMsg.Descriptor().FullName())

            stream.WriteObjectStart()
            stream.WriteObjectField(typeFieldName)
            stream.WriteVal(fullName)

            t := reflect2.TypeOf(any.typedVal)
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
                        f := field.UnsafeGet(reflect2.PtrOf(any.typedVal))
                        if !encoder.IsEmpty(f) {
                            encoder.Encode(f, stream)
                        }
                    } else {
                        stream.WriteVal(field.Get(any.typedVal))
                    }
                }
            } else {
                logs.Errorw("invalid the struct type for any object", "type", fullName)
            }

            stream.WriteObjectEnd()
        }
    default:
        kvWriter(GetMojoTypeName(any.typedVal), v)
    }
}
