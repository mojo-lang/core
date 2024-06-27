// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/color.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Color}
 */
public final class Color extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Color)
    ColorOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Color.newBuilder() to construct.
  private Color(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Color() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Color();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Color.class, org.mojolang.mojo.core.Color.Builder.class);
  }

  public static final int RED_FIELD_NUMBER = 1;
  private int red_;
  /**
   * <code>uint32 red = 1;</code>
   * @return The red.
   */
  @java.lang.Override
  public int getRed() {
    return red_;
  }

  public static final int GREEN_FIELD_NUMBER = 2;
  private int green_;
  /**
   * <code>uint32 green = 2;</code>
   * @return The green.
   */
  @java.lang.Override
  public int getGreen() {
    return green_;
  }

  public static final int BLUE_FIELD_NUMBER = 3;
  private int blue_;
  /**
   * <code>uint32 blue = 3;</code>
   * @return The blue.
   */
  @java.lang.Override
  public int getBlue() {
    return blue_;
  }

  public static final int ALPHA_FIELD_NUMBER = 4;
  private org.mojolang.mojo.core.Float32Value alpha_;
  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   * @return Whether the alpha field is set.
   */
  @java.lang.Override
  public boolean hasAlpha() {
    return alpha_ != null;
  }
  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   * @return The alpha.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Float32Value getAlpha() {
    return alpha_ == null ? org.mojolang.mojo.core.Float32Value.getDefaultInstance() : alpha_;
  }
  /**
   * <code>.mojo.core.Float32Value alpha = 4;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Float32ValueOrBuilder getAlphaOrBuilder() {
    return getAlpha();
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (red_ != 0) {
      output.writeUInt32(1, red_);
    }
    if (green_ != 0) {
      output.writeUInt32(2, green_);
    }
    if (blue_ != 0) {
      output.writeUInt32(3, blue_);
    }
    if (alpha_ != null) {
      output.writeMessage(4, getAlpha());
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (red_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(1, red_);
    }
    if (green_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(2, green_);
    }
    if (blue_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(3, blue_);
    }
    if (alpha_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(4, getAlpha());
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.core.Color)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Color other = (org.mojolang.mojo.core.Color) obj;

    if (getRed()
        != other.getRed()) return false;
    if (getGreen()
        != other.getGreen()) return false;
    if (getBlue()
        != other.getBlue()) return false;
    if (hasAlpha() != other.hasAlpha()) return false;
    if (hasAlpha()) {
      if (!getAlpha()
          .equals(other.getAlpha())) return false;
    }
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + RED_FIELD_NUMBER;
    hash = (53 * hash) + getRed();
    hash = (37 * hash) + GREEN_FIELD_NUMBER;
    hash = (53 * hash) + getGreen();
    hash = (37 * hash) + BLUE_FIELD_NUMBER;
    hash = (53 * hash) + getBlue();
    if (hasAlpha()) {
      hash = (37 * hash) + ALPHA_FIELD_NUMBER;
      hash = (53 * hash) + getAlpha().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Color parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.core.Color prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.core.Color}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Color)
      org.mojolang.mojo.core.ColorOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Color.class, org.mojolang.mojo.core.Color.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Color.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      red_ = 0;

      green_ = 0;

      blue_ = 0;

      if (alphaBuilder_ == null) {
        alpha_ = null;
      } else {
        alpha_ = null;
        alphaBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Color.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color build() {
      org.mojolang.mojo.core.Color result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color buildPartial() {
      org.mojolang.mojo.core.Color result = new org.mojolang.mojo.core.Color(this);
      result.red_ = red_;
      result.green_ = green_;
      result.blue_ = blue_;
      if (alphaBuilder_ == null) {
        result.alpha_ = alpha_;
      } else {
        result.alpha_ = alphaBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.core.Color) {
        return mergeFrom((org.mojolang.mojo.core.Color)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Color other) {
      if (other == org.mojolang.mojo.core.Color.getDefaultInstance()) return this;
      if (other.getRed() != 0) {
        setRed(other.getRed());
      }
      if (other.getGreen() != 0) {
        setGreen(other.getGreen());
      }
      if (other.getBlue() != 0) {
        setBlue(other.getBlue());
      }
      if (other.hasAlpha()) {
        mergeAlpha(other.getAlpha());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {
              red_ = input.readUInt32();

              break;
            } // case 8
            case 16: {
              green_ = input.readUInt32();

              break;
            } // case 16
            case 24: {
              blue_ = input.readUInt32();

              break;
            } // case 24
            case 34: {
              input.readMessage(
                  getAlphaFieldBuilder().getBuilder(),
                  extensionRegistry);

              break;
            } // case 34
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private int red_ ;
    /**
     * <code>uint32 red = 1;</code>
     * @return The red.
     */
    @java.lang.Override
    public int getRed() {
      return red_;
    }
    /**
     * <code>uint32 red = 1;</code>
     * @param value The red to set.
     * @return This builder for chaining.
     */
    public Builder setRed(int value) {
      
      red_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>uint32 red = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearRed() {
      
      red_ = 0;
      onChanged();
      return this;
    }

    private int green_ ;
    /**
     * <code>uint32 green = 2;</code>
     * @return The green.
     */
    @java.lang.Override
    public int getGreen() {
      return green_;
    }
    /**
     * <code>uint32 green = 2;</code>
     * @param value The green to set.
     * @return This builder for chaining.
     */
    public Builder setGreen(int value) {
      
      green_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>uint32 green = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearGreen() {
      
      green_ = 0;
      onChanged();
      return this;
    }

    private int blue_ ;
    /**
     * <code>uint32 blue = 3;</code>
     * @return The blue.
     */
    @java.lang.Override
    public int getBlue() {
      return blue_;
    }
    /**
     * <code>uint32 blue = 3;</code>
     * @param value The blue to set.
     * @return This builder for chaining.
     */
    public Builder setBlue(int value) {
      
      blue_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>uint32 blue = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlue() {
      
      blue_ = 0;
      onChanged();
      return this;
    }

    private org.mojolang.mojo.core.Float32Value alpha_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Float32Value, org.mojolang.mojo.core.Float32Value.Builder, org.mojolang.mojo.core.Float32ValueOrBuilder> alphaBuilder_;
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     * @return Whether the alpha field is set.
     */
    public boolean hasAlpha() {
      return alphaBuilder_ != null || alpha_ != null;
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     * @return The alpha.
     */
    public org.mojolang.mojo.core.Float32Value getAlpha() {
      if (alphaBuilder_ == null) {
        return alpha_ == null ? org.mojolang.mojo.core.Float32Value.getDefaultInstance() : alpha_;
      } else {
        return alphaBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public Builder setAlpha(org.mojolang.mojo.core.Float32Value value) {
      if (alphaBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        alpha_ = value;
        onChanged();
      } else {
        alphaBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public Builder setAlpha(
        org.mojolang.mojo.core.Float32Value.Builder builderForValue) {
      if (alphaBuilder_ == null) {
        alpha_ = builderForValue.build();
        onChanged();
      } else {
        alphaBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public Builder mergeAlpha(org.mojolang.mojo.core.Float32Value value) {
      if (alphaBuilder_ == null) {
        if (alpha_ != null) {
          alpha_ =
            org.mojolang.mojo.core.Float32Value.newBuilder(alpha_).mergeFrom(value).buildPartial();
        } else {
          alpha_ = value;
        }
        onChanged();
      } else {
        alphaBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public Builder clearAlpha() {
      if (alphaBuilder_ == null) {
        alpha_ = null;
        onChanged();
      } else {
        alpha_ = null;
        alphaBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public org.mojolang.mojo.core.Float32Value.Builder getAlphaBuilder() {
      
      onChanged();
      return getAlphaFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    public org.mojolang.mojo.core.Float32ValueOrBuilder getAlphaOrBuilder() {
      if (alphaBuilder_ != null) {
        return alphaBuilder_.getMessageOrBuilder();
      } else {
        return alpha_ == null ?
            org.mojolang.mojo.core.Float32Value.getDefaultInstance() : alpha_;
      }
    }
    /**
     * <code>.mojo.core.Float32Value alpha = 4;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Float32Value, org.mojolang.mojo.core.Float32Value.Builder, org.mojolang.mojo.core.Float32ValueOrBuilder> 
        getAlphaFieldBuilder() {
      if (alphaBuilder_ == null) {
        alphaBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Float32Value, org.mojolang.mojo.core.Float32Value.Builder, org.mojolang.mojo.core.Float32ValueOrBuilder>(
                getAlpha(),
                getParentForChildren(),
                isClean());
        alpha_ = null;
      }
      return alphaBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.core.Color)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Color)
  private static final org.mojolang.mojo.core.Color DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Color();
  }

  public static org.mojolang.mojo.core.Color getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Color>
      PARSER = new com.google.protobuf.AbstractParser<Color>() {
    @java.lang.Override
    public Color parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<Color> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Color> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Color getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

