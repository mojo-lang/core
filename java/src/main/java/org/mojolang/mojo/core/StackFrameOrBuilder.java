// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/stack_frame.proto

package org.mojolang.mojo.core;

public interface StackFrameOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.StackFrame)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int64 instruction = 1;</code>
   * @return The instruction.
   */
  long getInstruction();

  /**
   * <code>.mojo.core.CodeModule module = 2;</code>
   * @return Whether the module field is set.
   */
  boolean hasModule();
  /**
   * <code>.mojo.core.CodeModule module = 2;</code>
   * @return The module.
   */
  org.mojolang.mojo.core.CodeModule getModule();
  /**
   * <code>.mojo.core.CodeModule module = 2;</code>
   */
  org.mojolang.mojo.core.CodeModuleOrBuilder getModuleOrBuilder();

  /**
   * <code>string function_name = 3;</code>
   * @return The functionName.
   */
  java.lang.String getFunctionName();
  /**
   * <code>string function_name = 3;</code>
   * @return The bytes for functionName.
   */
  com.google.protobuf.ByteString
      getFunctionNameBytes();

  /**
   * <code>int64 function_base = 4;</code>
   * @return The functionBase.
   */
  long getFunctionBase();

  /**
   * <code>string source_file_name = 5;</code>
   * @return The sourceFileName.
   */
  java.lang.String getSourceFileName();
  /**
   * <code>string source_file_name = 5;</code>
   * @return The bytes for sourceFileName.
   */
  com.google.protobuf.ByteString
      getSourceFileNameBytes();

  /**
   * <code>int32 source_line = 6;</code>
   * @return The sourceLine.
   */
  int getSourceLine();

  /**
   * <code>int64 source_line_base = 7;</code>
   * @return The sourceLineBase.
   */
  long getSourceLineBase();

  /**
   * <code>.mojo.core.StackFrame.Trust trust = 8;</code>
   * @return The enum numeric value on the wire for trust.
   */
  int getTrustValue();
  /**
   * <code>.mojo.core.StackFrame.Trust trust = 8;</code>
   * @return The trust.
   */
  org.mojolang.mojo.core.StackFrame.Trust getTrust();
}
