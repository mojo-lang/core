// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/percentage.proto

package org.mojolang.mojo.core;

public final class PercentageProto {
  private PercentageProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Percentage_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Percentage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\032mojo/core/percentage.proto\022\tmojo.core\"" +
      "\031\n\nPercentage\022\013\n\003val\030\001 \001(\rB\\\n\026org.mojola" +
      "ng.mojo.coreB\017PercentageProtoP\001Z/github." +
      "com/mojo-lang/core/go/pkg/mojo/core;core" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_Percentage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Percentage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Percentage_descriptor,
        new java.lang.String[] { "Val", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
