// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Float32Values}
 */
public final class Float32Values extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Float32Values)
    Float32ValuesOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Float32Values.newBuilder() to construct.
  private Float32Values(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Float32Values() {
    vals_ = emptyFloatList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Float32Values();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Values_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Values_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Float32Values.class, org.mojolang.mojo.core.Float32Values.Builder.class);
  }

  public static final int VALS_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private com.google.protobuf.Internal.FloatList vals_;
  /**
   * <code>repeated float vals = 1;</code>
   * @return A list containing the vals.
   */
  @java.lang.Override
  public java.util.List<java.lang.Float>
      getValsList() {
    return vals_;
  }
  /**
   * <code>repeated float vals = 1;</code>
   * @return The count of vals.
   */
  public int getValsCount() {
    return vals_.size();
  }
  /**
   * <code>repeated float vals = 1;</code>
   * @param index The index of the element to return.
   * @return The vals at the given index.
   */
  public float getVals(int index) {
    return vals_.getFloat(index);
  }
  private int valsMemoizedSerializedSize = -1;

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
    getSerializedSize();
    if (getValsList().size() > 0) {
      output.writeUInt32NoTag(10);
      output.writeUInt32NoTag(valsMemoizedSerializedSize);
    }
    for (int i = 0; i < vals_.size(); i++) {
      output.writeFloatNoTag(vals_.getFloat(i));
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    {
      int dataSize = 0;
      dataSize = 4 * getValsList().size();
      size += dataSize;
      if (!getValsList().isEmpty()) {
        size += 1;
        size += com.google.protobuf.CodedOutputStream
            .computeInt32SizeNoTag(dataSize);
      }
      valsMemoizedSerializedSize = dataSize;
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
    if (!(obj instanceof org.mojolang.mojo.core.Float32Values)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Float32Values other = (org.mojolang.mojo.core.Float32Values) obj;

    if (!getValsList()
        .equals(other.getValsList())) return false;
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
    if (getValsCount() > 0) {
      hash = (37 * hash) + VALS_FIELD_NUMBER;
      hash = (53 * hash) + getValsList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Float32Values parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Values parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Values parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Float32Values parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Float32Values prototype) {
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
   * Protobuf type {@code mojo.core.Float32Values}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Float32Values)
      org.mojolang.mojo.core.Float32ValuesOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Values_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Values_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Float32Values.class, org.mojolang.mojo.core.Float32Values.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Float32Values.newBuilder()
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
      vals_ = emptyFloatList();
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_Float32Values_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Values getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Float32Values.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Values build() {
      org.mojolang.mojo.core.Float32Values result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Float32Values buildPartial() {
      org.mojolang.mojo.core.Float32Values result = new org.mojolang.mojo.core.Float32Values(this);
      buildPartialRepeatedFields(result);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartialRepeatedFields(org.mojolang.mojo.core.Float32Values result) {
      if (((bitField0_ & 0x00000001) != 0)) {
        vals_.makeImmutable();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.vals_ = vals_;
    }

    private void buildPartial0(org.mojolang.mojo.core.Float32Values result) {
      int from_bitField0_ = bitField0_;
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
      if (other instanceof org.mojolang.mojo.core.Float32Values) {
        return mergeFrom((org.mojolang.mojo.core.Float32Values)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Float32Values other) {
      if (other == org.mojolang.mojo.core.Float32Values.getDefaultInstance()) return this;
      if (!other.vals_.isEmpty()) {
        if (vals_.isEmpty()) {
          vals_ = other.vals_;
          bitField0_ = (bitField0_ & ~0x00000001);
        } else {
          ensureValsIsMutable();
          vals_.addAll(other.vals_);
        }
        onChanged();
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
              float v = input.readFloat();
              ensureValsIsMutable();
              vals_.addFloat(v);
              break;
            } // case 13
            case 10: {
              int length = input.readRawVarint32();
              int limit = input.pushLimit(length);
              ensureValsIsMutable();
              while (input.getBytesUntilLimit() > 0) {
                vals_.addFloat(input.readFloat());
              }
              input.popLimit(limit);
              break;
            } // case 10
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

    private com.google.protobuf.Internal.FloatList vals_ = emptyFloatList();
    private void ensureValsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        vals_ = mutableCopy(vals_);
        bitField0_ |= 0x00000001;
      }
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @return A list containing the vals.
     */
    public java.util.List<java.lang.Float>
        getValsList() {
      return ((bitField0_ & 0x00000001) != 0) ?
               java.util.Collections.unmodifiableList(vals_) : vals_;
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @return The count of vals.
     */
    public int getValsCount() {
      return vals_.size();
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @param index The index of the element to return.
     * @return The vals at the given index.
     */
    public float getVals(int index) {
      return vals_.getFloat(index);
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @param index The index to set the value at.
     * @param value The vals to set.
     * @return This builder for chaining.
     */
    public Builder setVals(
        int index, float value) {
      
      ensureValsIsMutable();
      vals_.setFloat(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @param value The vals to add.
     * @return This builder for chaining.
     */
    public Builder addVals(float value) {
      
      ensureValsIsMutable();
      vals_.addFloat(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @param values The vals to add.
     * @return This builder for chaining.
     */
    public Builder addAllVals(
        java.lang.Iterable<? extends java.lang.Float> values) {
      ensureValsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, vals_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated float vals = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearVals() {
      vals_ = emptyFloatList();
      bitField0_ = (bitField0_ & ~0x00000001);
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Float32Values)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Float32Values)
  private static final org.mojolang.mojo.core.Float32Values DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Float32Values();
  }

  public static org.mojolang.mojo.core.Float32Values getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Float32Values>
      PARSER = new com.google.protobuf.AbstractParser<Float32Values>() {
    @java.lang.Override
    public Float32Values parsePartialFrom(
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

  public static com.google.protobuf.Parser<Float32Values> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Float32Values> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Float32Values getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

