// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/color.proto

package org.mojolang.mojo.core;

public interface ColorOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.Color)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint32 red = 1;</code>
   * @return The red.
   */
  int getRed();

  /**
   * <code>uint32 green = 2;</code>
   * @return The green.
   */
  int getGreen();

  /**
   * <code>uint32 blue = 3;</code>
   * @return The blue.
   */
  int getBlue();

  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   * @return Whether the alpha field is set.
   */
  boolean hasAlpha();
  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   * @return The alpha.
   */
  org.mojolang.mojo.core.Float32Value getAlpha();
  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   */
  org.mojolang.mojo.core.Float32ValueOrBuilder getAlphaOrBuilder();
}
