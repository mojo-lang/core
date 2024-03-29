// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/range.proto

package org.mojolang.mojo.core;

public interface IntRangeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.IntRange)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int64 start = 1;</code>
   * @return The start.
   */
  long getStart();

  /**
   * <code>int64 end = 2;</code>
   * @return The end.
   */
  long getEnd();

  /**
   * <code>bool start_excluded = 9;</code>
   * @return The startExcluded.
   */
  boolean getStartExcluded();

  /**
   * <code>bool end_included = 10;</code>
   * @return The endIncluded.
   */
  boolean getEndIncluded();
}
