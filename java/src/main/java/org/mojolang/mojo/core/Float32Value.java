// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Float32Value}
 */
public final class Float32Value extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Float32Value)
    Float32ValueOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Float32Value.newBuilder() to construct.
  private Float32Value(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Float32Value() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Float32Value();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Value_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Value_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Float32Value.class, org.mojolang.mojo.core.Float32Value.Builder.class);
  }

  public static final int VAL_FIELD_NUMBER = 1;
  private float val_ = 0F;
  /**
   * <code>float val = 1;</code>
   * @return The val.
   */
  @java.lang.Override
  public float getVal() {
    return val_;
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
    if (java.lang.Float.floatToRawIntBits(val_) != 0) {
      output.writeFloat(1, val_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (java.lang.Float.floatToRawIntBits(val_) != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(1, val_);
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
    if (!(obj instanceof org.mojolang.mojo.core.Float32Value)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Float32Value other = (org.mojolang.mojo.core.Float32Value) obj;

    if (java.lang.Float.floatToIntBits(getVal())
        != java.lang.Float.floatToIntBits(
            other.getVal())) return false;
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
    hash = (37 * hash) + VAL_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getVal());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Float32Value parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Value parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Value parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Value parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Float32Value prototype) {
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
   * Protobuf type {@code mojo.core.Float32Value}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Float32Value)
      org.mojolang.mojo.core.Float32ValueOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Value_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Value_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Float32Value.class, org.mojolang.mojo.core.Float32Value.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Float32Value.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      val_ = 0F;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Value_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Value getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Float32Value.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Value build() {
      org.mojolang.mojo.core.Float32Value result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Value buildPartial() {
      org.mojolang.mojo.core.Float32Value result = new org.mojolang.mojo.core.Float32Value(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(org.mojolang.mojo.core.Float32Value result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.val_ = val_;
      }
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
      if (other instanceof org.mojolang.mojo.core.Float32Value) {
        return mergeFrom((org.mojolang.mojo.core.Float32Value)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Float32Value other) {
      if (other == org.mojolang.mojo.core.Float32Value.getDefaultInstance()) return this;
      if (other.getVal() != 0F) {
        setVal(other.getVal());
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
            case 13: {
              val_ = input.readFloat();
              bitField0_ |= 0x00000001;
              break;
            } // case 13
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
    private int bitField0_;

    private float val_ ;
    /**
     * <code>float val = 1;</code>
     * @return The val.
     */
    @java.lang.Override
    public float getVal() {
      return val_;
    }
    /**
     * <code>float val = 1;</code>
     * @param value The val to set.
     * @return This builder for chaining.
     */
    public Builder setVal(float value) {
      
      val_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>float val = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearVal() {
      bitField0_ = (bitField0_ & ~0x00000001);
      val_ = 0F;
      onChanged();
      return this;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Float32Value)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Float32Value)
  private static final org.mojolang.mojo.core.Float32Value DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Float32Value();
  }

  public static org.mojolang.mojo.core.Float32Value getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Float32Value>
      PARSER = new com.google.protobuf.AbstractParser<Float32Value>() {
    @java.lang.Override
    public Float32Value parsePartialFrom(
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

  public static com.google.protobuf.Parser<Float32Value> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Float32Value> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Float32Value getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

