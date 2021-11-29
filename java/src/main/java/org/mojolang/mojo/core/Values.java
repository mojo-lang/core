// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/value.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Values}
 */
public final class Values extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Values)
    ValuesOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Values.newBuilder() to construct.
  private Values(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Values() {
    vals_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Values();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Values(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              vals_ = new java.util.ArrayList<org.mojolang.mojo.core.Value>();
              mutable_bitField0_ |= 0x00000001;
            }
            vals_.add(
                input.readMessage(org.mojolang.mojo.core.Value.parser(), extensionRegistry));
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000001) != 0)) {
        vals_ = java.util.Collections.unmodifiableList(vals_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.ValueProto.internal_static_mojo_core_Values_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ValueProto.internal_static_mojo_core_Values_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Values.class, org.mojolang.mojo.core.Values.Builder.class);
  }

  public static final int VALS_FIELD_NUMBER = 1;
  private java.util.List<org.mojolang.mojo.core.Value> vals_;
  /**
   * <code>repeated .mojo.core.Value vals = 1;</code>
   */
  @java.lang.Override
  public java.util.List<org.mojolang.mojo.core.Value> getValsList() {
    return vals_;
  }
  /**
   * <code>repeated .mojo.core.Value vals = 1;</code>
   */
  @java.lang.Override
  public java.util.List<? extends org.mojolang.mojo.core.ValueOrBuilder> 
      getValsOrBuilderList() {
    return vals_;
  }
  /**
   * <code>repeated .mojo.core.Value vals = 1;</code>
   */
  @java.lang.Override
  public int getValsCount() {
    return vals_.size();
  }
  /**
   * <code>repeated .mojo.core.Value vals = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Value getVals(int index) {
    return vals_.get(index);
  }
  /**
   * <code>repeated .mojo.core.Value vals = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.ValueOrBuilder getValsOrBuilder(
      int index) {
    return vals_.get(index);
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
      output.writeMessage(1, vals_.get(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    for (int i = 0; i < vals_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, vals_.get(i));
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.core.Values)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Values other = (org.mojolang.mojo.core.Values) obj;

    if (!getValsList()
        .equals(other.getValsList())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
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
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Values parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Values parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Values parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Values parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Values parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Values parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Values prototype) {
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
   * Protobuf type {@code mojo.core.Values}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Values)
      org.mojolang.mojo.core.ValuesOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ValueProto.internal_static_mojo_core_Values_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ValueProto.internal_static_mojo_core_Values_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Values.class, org.mojolang.mojo.core.Values.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Values.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
        getValsFieldBuilder();
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (valsBuilder_ == null) {
        vals_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
      } else {
        valsBuilder_.clear();
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ValueProto.internal_static_mojo_core_Values_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Values getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Values.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Values build() {
      org.mojolang.mojo.core.Values result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Values buildPartial() {
      org.mojolang.mojo.core.Values result = new org.mojolang.mojo.core.Values(this);
      int from_bitField0_ = bitField0_;
      if (valsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          vals_ = java.util.Collections.unmodifiableList(vals_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.vals_ = vals_;
      } else {
        result.vals_ = valsBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.core.Values) {
        return mergeFrom((org.mojolang.mojo.core.Values)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Values other) {
      if (other == org.mojolang.mojo.core.Values.getDefaultInstance()) return this;
      if (valsBuilder_ == null) {
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
      } else {
        if (!other.vals_.isEmpty()) {
          if (valsBuilder_.isEmpty()) {
            valsBuilder_.dispose();
            valsBuilder_ = null;
            vals_ = other.vals_;
            bitField0_ = (bitField0_ & ~0x00000001);
            valsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getValsFieldBuilder() : null;
          } else {
            valsBuilder_.addAllMessages(other.vals_);
          }
        }
      }
      this.mergeUnknownFields(other.unknownFields);
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
      org.mojolang.mojo.core.Values parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.Values) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.util.List<org.mojolang.mojo.core.Value> vals_ =
      java.util.Collections.emptyList();
    private void ensureValsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        vals_ = new java.util.ArrayList<org.mojolang.mojo.core.Value>(vals_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Value, org.mojolang.mojo.core.Value.Builder, org.mojolang.mojo.core.ValueOrBuilder> valsBuilder_;

    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Value> getValsList() {
      if (valsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(vals_);
      } else {
        return valsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public int getValsCount() {
      if (valsBuilder_ == null) {
        return vals_.size();
      } else {
        return valsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public org.mojolang.mojo.core.Value getVals(int index) {
      if (valsBuilder_ == null) {
        return vals_.get(index);
      } else {
        return valsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder setVals(
        int index, org.mojolang.mojo.core.Value value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.set(index, value);
        onChanged();
      } else {
        valsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder setVals(
        int index, org.mojolang.mojo.core.Value.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.set(index, builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder addVals(org.mojolang.mojo.core.Value value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.add(value);
        onChanged();
      } else {
        valsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder addVals(
        int index, org.mojolang.mojo.core.Value value) {
      if (valsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureValsIsMutable();
        vals_.add(index, value);
        onChanged();
      } else {
        valsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder addVals(
        org.mojolang.mojo.core.Value.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.add(builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder addVals(
        int index, org.mojolang.mojo.core.Value.Builder builderForValue) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.add(index, builderForValue.build());
        onChanged();
      } else {
        valsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder addAllVals(
        java.lang.Iterable<? extends org.mojolang.mojo.core.Value> values) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, vals_);
        onChanged();
      } else {
        valsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder clearVals() {
      if (valsBuilder_ == null) {
        vals_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        valsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public Builder removeVals(int index) {
      if (valsBuilder_ == null) {
        ensureValsIsMutable();
        vals_.remove(index);
        onChanged();
      } else {
        valsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public org.mojolang.mojo.core.Value.Builder getValsBuilder(
        int index) {
      return getValsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public org.mojolang.mojo.core.ValueOrBuilder getValsOrBuilder(
        int index) {
      if (valsBuilder_ == null) {
        return vals_.get(index);  } else {
        return valsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public java.util.List<? extends org.mojolang.mojo.core.ValueOrBuilder> 
         getValsOrBuilderList() {
      if (valsBuilder_ != null) {
        return valsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(vals_);
      }
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public org.mojolang.mojo.core.Value.Builder addValsBuilder() {
      return getValsFieldBuilder().addBuilder(
          org.mojolang.mojo.core.Value.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public org.mojolang.mojo.core.Value.Builder addValsBuilder(
        int index) {
      return getValsFieldBuilder().addBuilder(
          index, org.mojolang.mojo.core.Value.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Value vals = 1;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Value.Builder> 
         getValsBuilderList() {
      return getValsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Value, org.mojolang.mojo.core.Value.Builder, org.mojolang.mojo.core.ValueOrBuilder> 
        getValsFieldBuilder() {
      if (valsBuilder_ == null) {
        valsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            org.mojolang.mojo.core.Value, org.mojolang.mojo.core.Value.Builder, org.mojolang.mojo.core.ValueOrBuilder>(
                vals_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        vals_ = null;
      }
      return valsBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Values)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Values)
  private static final org.mojolang.mojo.core.Values DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Values();
  }

  public static org.mojolang.mojo.core.Values getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Values>
      PARSER = new com.google.protobuf.AbstractParser<Values>() {
    @java.lang.Override
    public Values parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Values(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Values> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Values> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Values getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

