// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/field_mask.proto

package org.mojolang.mojo.core;

public final class FieldMaskProto {
  private FieldMaskProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_FieldMask_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_FieldMask_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\032mojo/core/field_mask.proto\022\tmojo.core\"" +
      "\032\n\tFieldMask\022\r\n\005paths\030\001 \003(\tB[\n\026org.mojol" +
      "ang.mojo.coreB\016FieldMaskProtoP\001Z/github." +
      "com/mojo-lang/core/go/pkg/mojo/core;core" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_FieldMask_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_FieldMask_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_FieldMask_descriptor,
        new java.lang.String[] { "Paths", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
