// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/media_type.proto

package org.mojolang.mojo.core;

public final class MediaTypeProto {
  private MediaTypeProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_MediaType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_MediaType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_MediaType_Parameter_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_MediaType_Parameter_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\032mojo/core/media_type.proto\022\tmojo.core\032" +
      "\025mojo/core/value.proto\"\230\001\n\tMediaType\022\014\n\004" +
      "type\030\001 \001(\t\022\017\n\007subtype\030\002 \001(\t\0221\n\tparameter" +
      "\030\003 \001(\0132\036.mojo.core.MediaType.Parameter\0329" +
      "\n\tParameter\022\013\n\003key\030\001 \001(\t\022\037\n\005value\030\002 \001(\0132" +
      "\020.mojo.core.ValueB[\n\026org.mojolang.mojo.c" +
      "oreB\016MediaTypeProtoP\001Z/github.com/mojo-l" +
      "ang/core/go/pkg/mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.ValueProto.getDescriptor(),
        });
    internal_static_mojo_core_MediaType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_MediaType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_MediaType_descriptor,
        new java.lang.String[] { "Type", "Subtype", "Parameter", });
    internal_static_mojo_core_MediaType_Parameter_descriptor =
      internal_static_mojo_core_MediaType_descriptor.getNestedTypes().get(0);
    internal_static_mojo_core_MediaType_Parameter_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_MediaType_Parameter_descriptor,
        new java.lang.String[] { "Key", "Value", });
    org.mojolang.mojo.core.ValueProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
