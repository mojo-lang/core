// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.StringValues}
 */
public final class StringValues extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.StringValues)
    StringValuesOrBuilder {
private static final long serialVersionUID = 0L;
  // Use StringValues.newBuilder() to construct.
  private StringValues(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private StringValues() {
    vals_ = com.google.protobuf.LazyStringArrayList.EMPTY;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new StringValues();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringValues_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringValues_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.StringValues.class, org.mojolang.mojo.core.StringValues.Builder.class);
  }

  public static final int VALS_FIELD_NUMBER = 1;
  private com.google.protobuf.LazyStringList vals_;
  /**
   * <code>repeated string vals = 1;</code>
   * @return A list containing the vals.
   */
  public com.google.protobuf.ProtocolStringList
      getValsList() {
    return vals_;
  }
  /**
   * <code>repeated string vals = 1;</code>
   * @return The count of vals.
   */
  public int getValsCount() {
    return vals_.size();
  }
  /**
   * <code>repeated string vals = 1;</code>
   * @param index The index of the element to return.
   * @return The vals at the given index.
   */
  public java.lang.String getVals(int index) {
    return vals_.get(index);
  }
  /**
   * <code>repeated string vals = 1;</code>
   * @param index The index of the value to return.
   * @return The bytes of the vals at the given index.
   */
  public com.google.protobuf.ByteString
      getValsBytes(int index) {
    return vals_.getByteString(index);
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
    for (int i = 0; i < vals_.size(); i++) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, vals_.getRaw(i));
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
      for (int i = 0; i < vals_.size(); i++) {
        dataSize += computeStringSizeNoTag(vals_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getValsList().size();
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
    if (!(obj instanceof org.mojolang.mojo.core.StringValues)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.StringValues other = (org.mojolang.mojo.core.StringValues) obj;

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

  public static org.mojolang.mojo.core.StringValues parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringValues parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringValues parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringValues parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.StringValues prototype) {
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
   * Protobuf type {@code mojo.core.StringValues}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.StringValues)
      org.mojolang.mojo.core.StringValuesOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringValues_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringValues_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.StringValues.class, org.mojolang.mojo.core.StringValues.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.StringValues.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      vals_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringValues_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringValues getDefaultInstanceForType() {
      return org.mojolang.mojo.core.StringValues.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringValues build() {
      org.mojolang.mojo.core.StringValues result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringValues buildPartial() {
      org.mojolang.mojo.core.StringValues result = new org.mojolang.mojo.core.StringValues(this);
      int from_bitField0_ = bitField0_;
      if (((bitField0_ & 0x00000001) != 0)) {
        vals_ = vals_.getUnmodifiableView();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.vals_ = vals_;
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
      if (other instanceof org.mojolang.mojo.core.StringValues) {
        return mergeFrom((org.mojolang.mojo.core.StringValues)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.StringValues other) {
      if (other == org.mojolang.mojo.core.StringValues.getDefaultInstance()) return this;
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
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();
              ensureValsIsMutable();
              vals_.add(s);
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

    private com.google.protobuf.LazyStringList vals_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    private void ensureValsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        vals_ = new com.google.protobuf.LazyStringArrayList(vals_);
        bitField0_ |= 0x00000001;
       }
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @return A list containing the vals.
     */
    public com.google.protobuf.ProtocolStringList
        getValsList() {
      return vals_.getUnmodifiableView();
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @return The count of vals.
     */
    public int getValsCount() {
      return vals_.size();
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param index The index of the element to return.
     * @return The vals at the given index.
     */
    public java.lang.String getVals(int index) {
      return vals_.get(index);
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param index The index of the value to return.
     * @return The bytes of the vals at the given index.
     */
    public com.google.protobuf.ByteString
        getValsBytes(int index) {
      return vals_.getByteString(index);
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param index The index to set the value at.
     * @param value The vals to set.
     * @return This builder for chaining.
     */
    public Builder setVals(
        int index, java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureValsIsMutable();
      vals_.set(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param value The vals to add.
     * @return This builder for chaining.
     */
    public Builder addVals(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureValsIsMutable();
      vals_.add(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param values The vals to add.
     * @return This builder for chaining.
     */
    public Builder addAllVals(
        java.lang.Iterable<java.lang.String> values) {
      ensureValsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, vals_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearVals() {
      vals_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string vals = 1;</code>
     * @param value The bytes of the vals to add.
     * @return This builder for chaining.
     */
    public Builder addValsBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      ensureValsIsMutable();
      vals_.add(value);
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


    // @@protoc_insertion_point(builder_scope:mojo.core.StringValues)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.StringValues)
  private static final org.mojolang.mojo.core.StringValues DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.StringValues();
  }

  public static org.mojolang.mojo.core.StringValues getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<StringValues>
      PARSER = new com.google.protobuf.AbstractParser<StringValues>() {
    @java.lang.Override
    public StringValues parsePartialFrom(
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

  public static com.google.protobuf.Parser<StringValues> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<StringValues> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.StringValues getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

