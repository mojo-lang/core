// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/version_range.proto

package org.mojolang.mojo.core;

public final class VersionRangeProto {
  private VersionRangeProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_VersionRange_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_VersionRange_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\035mojo/core/version_range.proto\022\tmojo.co" +
      "re\032\027mojo/core/version.proto\"\200\001\n\014VersionR" +
      "ange\022!\n\005start\030\001 \001(\0132\022.mojo.core.Version\022" +
      "\037\n\003end\030\002 \001(\0132\022.mojo.core.Version\022\026\n\016star" +
      "t_excluded\030\t \001(\010\022\024\n\014end_included\030\n \001(\010B^" +
      "\n\026org.mojolang.mojo.coreB\021VersionRangePr" +
      "otoP\001Z/github.com/mojo-lang/core/go/pkg/" +
      "mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.VersionProto.getDescriptor(),
        });
    internal_static_mojo_core_VersionRange_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_VersionRange_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_VersionRange_descriptor,
        new java.lang.String[] { "Start", "End", "StartExcluded", "EndIncluded", });
    org.mojolang.mojo.core.VersionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
