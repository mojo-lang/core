// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/ordering.proto

package org.mojolang.mojo.core;

public final class OrderingProto {
  private OrderingProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Ordering_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Ordering_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Ordering_Value_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Ordering_Value_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\030mojo/core/ordering.proto\022\tmojo.core\"\256\001" +
      "\n\010Ordering\022\'\n\004vals\030\001 \003(\0132\031.mojo.core.Ord" +
      "ering.Value\032>\n\005Value\022\r\n\005field\030\001 \001(\t\022&\n\004s" +
      "ort\030\002 \001(\0162\030.mojo.core.Ordering.Sort\"9\n\004S" +
      "ort\022\024\n\020SORT_UNSPECIFIED\020\000\022\014\n\010SORT_ASC\020\001\022" +
      "\r\n\tSORT_DESC\020\002BZ\n\026org.mojolang.mojo.core" +
      "B\rOrderingProtoP\001Z/github.com/mojo-lang/" +
      "core/go/pkg/mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_Ordering_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Ordering_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Ordering_descriptor,
        new java.lang.String[] { "Vals", });
    internal_static_mojo_core_Ordering_Value_descriptor =
      internal_static_mojo_core_Ordering_descriptor.getNestedTypes().get(0);
    internal_static_mojo_core_Ordering_Value_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Ordering_Value_descriptor,
        new java.lang.String[] { "Field", "Sort", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}