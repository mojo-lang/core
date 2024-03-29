// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/code_module.proto

package org.mojolang.mojo.core;

public interface CodeModuleOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.CodeModule)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int64 base_address = 1;</code>
   * @return The baseAddress.
   */
  long getBaseAddress();

  /**
   * <code>int64 size = 2;</code>
   * @return The size.
   */
  long getSize();

  /**
   * <code>string code_file = 3;</code>
   * @return The codeFile.
   */
  java.lang.String getCodeFile();
  /**
   * <code>string code_file = 3;</code>
   * @return The bytes for codeFile.
   */
  com.google.protobuf.ByteString
      getCodeFileBytes();

  /**
   * <code>string code_identifier = 4;</code>
   * @return The codeIdentifier.
   */
  java.lang.String getCodeIdentifier();
  /**
   * <code>string code_identifier = 4;</code>
   * @return The bytes for codeIdentifier.
   */
  com.google.protobuf.ByteString
      getCodeIdentifierBytes();

  /**
   * <code>string debug_file = 5;</code>
   * @return The debugFile.
   */
  java.lang.String getDebugFile();
  /**
   * <code>string debug_file = 5;</code>
   * @return The bytes for debugFile.
   */
  com.google.protobuf.ByteString
      getDebugFileBytes();

  /**
   * <code>string debug_identifier = 6;</code>
   * @return The debugIdentifier.
   */
  java.lang.String getDebugIdentifier();
  /**
   * <code>string debug_identifier = 6;</code>
   * @return The bytes for debugIdentifier.
   */
  com.google.protobuf.ByteString
      getDebugIdentifierBytes();

  /**
   * <code>string version = 7;</code>
   * @return The version.
   */
  java.lang.String getVersion();
  /**
   * <code>string version = 7;</code>
   * @return The bytes for version.
   */
  com.google.protobuf.ByteString
      getVersionBytes();
}
