// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/any.proto

package org.mojolang.mojo.core;

public final class AnyProto {
  private AnyProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Any_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Any_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\023mojo/core/any.proto\022\tmojo.core\032\017mojo/m" +
      "ojo.proto\"+\n\003Any\022\027\n\004type\030\001 \001(\tB\t\312\317$\005@typ" +
      "e\022\013\n\003val\030\002 \001(\014BU\n\026org.mojolang.mojo.core" +
      "B\010AnyProtoP\001Z/github.com/mojo-lang/core/" +
      "go/pkg/mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.MojoProtos.getDescriptor(),
        });
    internal_static_mojo_core_Any_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Any_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Any_descriptor,
        new java.lang.String[] { "Type", "Val", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.protobuf.MojoProtos.alias);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.protobuf.MojoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
