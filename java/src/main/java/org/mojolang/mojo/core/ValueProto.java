// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/value.proto

package org.mojolang.mojo.core;

public final class ValueProto {
  private ValueProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Object_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Object_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Object_ValsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Object_ValsEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Values_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Values_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Value_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Value_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\025mojo/core/value.proto\022\tmojo.core\032\024mojo" +
      "/core/null.proto\"r\n\006Object\022)\n\004vals\030\001 \003(\013" +
      "2\033.mojo.core.Object.ValsEntry\032=\n\tValsEnt" +
      "ry\022\013\n\003key\030\001 \001(\t\022\037\n\005value\030\002 \001(\0132\020.mojo.co" +
      "re.Value:\0028\001\"(\n\006Values\022\036\n\004vals\030\001 \003(\0132\020.m" +
      "ojo.core.Value\"\212\002\n\005Value\022\'\n\nvalues_val\030\001" +
      " \001(\0132\021.mojo.core.ValuesH\000\022\'\n\nobject_val\030" +
      "\002 \001(\0132\021.mojo.core.ObjectH\000\022#\n\010null_val\030\003" +
      " \001(\0132\017.mojo.core.NullH\000\022\022\n\010bool_val\030\004 \001(" +
      "\010H\000\022\026\n\014positive_val\030\005 \001(\004H\000\022\026\n\014negative_" +
      "val\030\006 \001(\004H\000\022\024\n\ndouble_val\030\007 \001(\001H\000\022\024\n\nstr" +
      "ing_val\030\010 \001(\tH\000\022\023\n\tbytes_val\030\t \001(\014H\000B\005\n\003" +
      "val*\335\001\n\tValueKind\022\032\n\026VALUE_KIND_UNSPECIF" +
      "IED\020\000\022\023\n\017VALUE_KIND_NULL\020\001\022\026\n\022VALUE_KIND" +
      "_BOOLEAN\020\002\022\026\n\022VALUE_KIND_INTEGER\020\003\022\025\n\021VA" +
      "LUE_KIND_NUMBER\020\004\022\025\n\021VALUE_KIND_STRING\020\005" +
      "\022\024\n\020VALUE_KIND_BYTES\020\006\022\024\n\020VALUE_KIND_ARR" +
      "AY\020\n\022\025\n\021VALUE_KIND_OBJECT\020\013BW\n\026org.mojol" +
      "ang.mojo.coreB\nValueProtoP\001Z/github.com/" +
      "mojo-lang/core/go/pkg/mojo/core;coreb\006pr" +
      "oto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.NullProto.getDescriptor(),
        });
    internal_static_mojo_core_Object_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Object_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Object_descriptor,
        new java.lang.String[] { "Vals", });
    internal_static_mojo_core_Object_ValsEntry_descriptor =
      internal_static_mojo_core_Object_descriptor.getNestedTypes().get(0);
    internal_static_mojo_core_Object_ValsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Object_ValsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_mojo_core_Values_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_mojo_core_Values_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Values_descriptor,
        new java.lang.String[] { "Vals", });
    internal_static_mojo_core_Value_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_mojo_core_Value_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Value_descriptor,
        new java.lang.String[] { "ValuesVal", "ObjectVal", "NullVal", "BoolVal", "PositiveVal", "NegativeVal", "DoubleVal", "StringVal", "BytesVal", "Val", });
    org.mojolang.mojo.core.NullProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
