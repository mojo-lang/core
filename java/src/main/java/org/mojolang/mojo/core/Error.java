// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Error}
 */
public final class Error extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Error)
    ErrorOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Error.newBuilder() to construct.
  private Error(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Error() {
    message_ = "";
    details_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Error();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.ErrorProto.internal_static_mojo_core_Error_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ErrorProto.internal_static_mojo_core_Error_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Error.class, org.mojolang.mojo.core.Error.Builder.class);
  }

  public static final int CODE_FIELD_NUMBER = 1;
  private org.mojolang.mojo.core.ErrorCode code_;
  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   * @return Whether the code field is set.
   */
  @java.lang.Override
  public boolean hasCode() {
    return code_ != null;
  }
  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   * @return The code.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.ErrorCode getCode() {
    return code_ == null ? org.mojolang.mojo.core.ErrorCode.getDefaultInstance() : code_;
  }
  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.ErrorCodeOrBuilder getCodeOrBuilder() {
    return getCode();
  }

  public static final int MESSAGE_FIELD_NUMBER = 4;
  private volatile java.lang.Object message_;
  /**
   * <code>string message = 4;</code>
   * @return The message.
   */
  @java.lang.Override
  public java.lang.String getMessage() {
    java.lang.Object ref = message_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      message_ = s;
      return s;
    }
  }
  /**
   * <code>string message = 4;</code>
   * @return The bytes for message.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getMessageBytes() {
    java.lang.Object ref = message_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      message_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int DETAILS_FIELD_NUMBER = 10;
  private java.util.List<org.mojolang.mojo.core.Any> details_;
  /**
   * <code>repeated .mojo.core.Any details = 10;</code>
   */
  @java.lang.Override
  public java.util.List<org.mojolang.mojo.core.Any> getDetailsList() {
    return details_;
  }
  /**
   * <code>repeated .mojo.core.Any details = 10;</code>
   */
  @java.lang.Override
  public java.util.List<? extends org.mojolang.mojo.core.AnyOrBuilder> 
      getDetailsOrBuilderList() {
    return details_;
  }
  /**
   * <code>repeated .mojo.core.Any details = 10;</code>
   */
  @java.lang.Override
  public int getDetailsCount() {
    return details_.size();
  }
  /**
   * <code>repeated .mojo.core.Any details = 10;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Any getDetails(int index) {
    return details_.get(index);
  }
  /**
   * <code>repeated .mojo.core.Any details = 10;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.AnyOrBuilder getDetailsOrBuilder(
      int index) {
    return details_.get(index);
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
    if (code_ != null) {
      output.writeMessage(1, getCode());
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(message_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, message_);
    }
    for (int i = 0; i < details_.size(); i++) {
      output.writeMessage(10, details_.get(i));
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (code_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getCode());
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(message_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, message_);
    }
    for (int i = 0; i < details_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(10, details_.get(i));
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
    if (!(obj instanceof org.mojolang.mojo.core.Error)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Error other = (org.mojolang.mojo.core.Error) obj;

    if (hasCode() != other.hasCode()) return false;
    if (hasCode()) {
      if (!getCode()
          .equals(other.getCode())) return false;
    }
    if (!getMessage()
        .equals(other.getMessage())) return false;
    if (!getDetailsList()
        .equals(other.getDetailsList())) return false;
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
    if (hasCode()) {
      hash = (37 * hash) + CODE_FIELD_NUMBER;
      hash = (53 * hash) + getCode().hashCode();
    }
    hash = (37 * hash) + MESSAGE_FIELD_NUMBER;
    hash = (53 * hash) + getMessage().hashCode();
    if (getDetailsCount() > 0) {
      hash = (37 * hash) + DETAILS_FIELD_NUMBER;
      hash = (53 * hash) + getDetailsList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Error parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Error parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Error parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Error parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Error parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Error parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Error prototype) {
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
   * Protobuf type {@code mojo.core.Error}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Error)
      org.mojolang.mojo.core.ErrorOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ErrorProto.internal_static_mojo_core_Error_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ErrorProto.internal_static_mojo_core_Error_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Error.class, org.mojolang.mojo.core.Error.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Error.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (codeBuilder_ == null) {
        code_ = null;
      } else {
        code_ = null;
        codeBuilder_ = null;
      }
      message_ = "";

      if (detailsBuilder_ == null) {
        details_ = java.util.Collections.emptyList();
      } else {
        details_ = null;
        detailsBuilder_.clear();
      }
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ErrorProto.internal_static_mojo_core_Error_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Error getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Error.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Error build() {
      org.mojolang.mojo.core.Error result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Error buildPartial() {
      org.mojolang.mojo.core.Error result = new org.mojolang.mojo.core.Error(this);
      int from_bitField0_ = bitField0_;
      if (codeBuilder_ == null) {
        result.code_ = code_;
      } else {
        result.code_ = codeBuilder_.build();
      }
      result.message_ = message_;
      if (detailsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          details_ = java.util.Collections.unmodifiableList(details_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.details_ = details_;
      } else {
        result.details_ = detailsBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.core.Error) {
        return mergeFrom((org.mojolang.mojo.core.Error)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Error other) {
      if (other == org.mojolang.mojo.core.Error.getDefaultInstance()) return this;
      if (other.hasCode()) {
        mergeCode(other.getCode());
      }
      if (!other.getMessage().isEmpty()) {
        message_ = other.message_;
        onChanged();
      }
      if (detailsBuilder_ == null) {
        if (!other.details_.isEmpty()) {
          if (details_.isEmpty()) {
            details_ = other.details_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureDetailsIsMutable();
            details_.addAll(other.details_);
          }
          onChanged();
        }
      } else {
        if (!other.details_.isEmpty()) {
          if (detailsBuilder_.isEmpty()) {
            detailsBuilder_.dispose();
            detailsBuilder_ = null;
            details_ = other.details_;
            bitField0_ = (bitField0_ & ~0x00000001);
            detailsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getDetailsFieldBuilder() : null;
          } else {
            detailsBuilder_.addAllMessages(other.details_);
          }
        }
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
              input.readMessage(
                  getCodeFieldBuilder().getBuilder(),
                  extensionRegistry);

              break;
            } // case 10
            case 34: {
              message_ = input.readStringRequireUtf8();

              break;
            } // case 34
            case 82: {
              org.mojolang.mojo.core.Any m =
                  input.readMessage(
                      org.mojolang.mojo.core.Any.parser(),
                      extensionRegistry);
              if (detailsBuilder_ == null) {
                ensureDetailsIsMutable();
                details_.add(m);
              } else {
                detailsBuilder_.addMessage(m);
              }
              break;
            } // case 82
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

    private org.mojolang.mojo.core.ErrorCode code_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.ErrorCode, org.mojolang.mojo.core.ErrorCode.Builder, org.mojolang.mojo.core.ErrorCodeOrBuilder> codeBuilder_;
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     * @return Whether the code field is set.
     */
    public boolean hasCode() {
      return codeBuilder_ != null || code_ != null;
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     * @return The code.
     */
    public org.mojolang.mojo.core.ErrorCode getCode() {
      if (codeBuilder_ == null) {
        return code_ == null ? org.mojolang.mojo.core.ErrorCode.getDefaultInstance() : code_;
      } else {
        return codeBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public Builder setCode(org.mojolang.mojo.core.ErrorCode value) {
      if (codeBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        code_ = value;
        onChanged();
      } else {
        codeBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public Builder setCode(
        org.mojolang.mojo.core.ErrorCode.Builder builderForValue) {
      if (codeBuilder_ == null) {
        code_ = builderForValue.build();
        onChanged();
      } else {
        codeBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public Builder mergeCode(org.mojolang.mojo.core.ErrorCode value) {
      if (codeBuilder_ == null) {
        if (code_ != null) {
          code_ =
            org.mojolang.mojo.core.ErrorCode.newBuilder(code_).mergeFrom(value).buildPartial();
        } else {
          code_ = value;
        }
        onChanged();
      } else {
        codeBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public Builder clearCode() {
      if (codeBuilder_ == null) {
        code_ = null;
        onChanged();
      } else {
        code_ = null;
        codeBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public org.mojolang.mojo.core.ErrorCode.Builder getCodeBuilder() {
      
      onChanged();
      return getCodeFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    public org.mojolang.mojo.core.ErrorCodeOrBuilder getCodeOrBuilder() {
      if (codeBuilder_ != null) {
        return codeBuilder_.getMessageOrBuilder();
      } else {
        return code_ == null ?
            org.mojolang.mojo.core.ErrorCode.getDefaultInstance() : code_;
      }
    }
    /**
     * <code>.mojo.core.ErrorCode code = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.ErrorCode, org.mojolang.mojo.core.ErrorCode.Builder, org.mojolang.mojo.core.ErrorCodeOrBuilder> 
        getCodeFieldBuilder() {
      if (codeBuilder_ == null) {
        codeBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.ErrorCode, org.mojolang.mojo.core.ErrorCode.Builder, org.mojolang.mojo.core.ErrorCodeOrBuilder>(
                getCode(),
                getParentForChildren(),
                isClean());
        code_ = null;
      }
      return codeBuilder_;
    }

    private java.lang.Object message_ = "";
    /**
     * <code>string message = 4;</code>
     * @return The message.
     */
    public java.lang.String getMessage() {
      java.lang.Object ref = message_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        message_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string message = 4;</code>
     * @return The bytes for message.
     */
    public com.google.protobuf.ByteString
        getMessageBytes() {
      java.lang.Object ref = message_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        message_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string message = 4;</code>
     * @param value The message to set.
     * @return This builder for chaining.
     */
    public Builder setMessage(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      message_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string message = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearMessage() {
      
      message_ = getDefaultInstance().getMessage();
      onChanged();
      return this;
    }
    /**
     * <code>string message = 4;</code>
     * @param value The bytes for message to set.
     * @return This builder for chaining.
     */
    public Builder setMessageBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      message_ = value;
      onChanged();
      return this;
    }

    private java.util.List<org.mojolang.mojo.core.Any> details_ =
      java.util.Collections.emptyList();
    private void ensureDetailsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        details_ = new java.util.ArrayList<org.mojolang.mojo.core.Any>(details_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Any, org.mojolang.mojo.core.Any.Builder, org.mojolang.mojo.core.AnyOrBuilder> detailsBuilder_;

    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Any> getDetailsList() {
      if (detailsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(details_);
      } else {
        return detailsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public int getDetailsCount() {
      if (detailsBuilder_ == null) {
        return details_.size();
      } else {
        return detailsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public org.mojolang.mojo.core.Any getDetails(int index) {
      if (detailsBuilder_ == null) {
        return details_.get(index);
      } else {
        return detailsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder setDetails(
        int index, org.mojolang.mojo.core.Any value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.set(index, value);
        onChanged();
      } else {
        detailsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder setDetails(
        int index, org.mojolang.mojo.core.Any.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.set(index, builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder addDetails(org.mojolang.mojo.core.Any value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.add(value);
        onChanged();
      } else {
        detailsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder addDetails(
        int index, org.mojolang.mojo.core.Any value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.add(index, value);
        onChanged();
      } else {
        detailsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder addDetails(
        org.mojolang.mojo.core.Any.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.add(builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder addDetails(
        int index, org.mojolang.mojo.core.Any.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.add(index, builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder addAllDetails(
        java.lang.Iterable<? extends org.mojolang.mojo.core.Any> values) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, details_);
        onChanged();
      } else {
        detailsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder clearDetails() {
      if (detailsBuilder_ == null) {
        details_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        detailsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public Builder removeDetails(int index) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.remove(index);
        onChanged();
      } else {
        detailsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public org.mojolang.mojo.core.Any.Builder getDetailsBuilder(
        int index) {
      return getDetailsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public org.mojolang.mojo.core.AnyOrBuilder getDetailsOrBuilder(
        int index) {
      if (detailsBuilder_ == null) {
        return details_.get(index);  } else {
        return detailsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public java.util.List<? extends org.mojolang.mojo.core.AnyOrBuilder> 
         getDetailsOrBuilderList() {
      if (detailsBuilder_ != null) {
        return detailsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(details_);
      }
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public org.mojolang.mojo.core.Any.Builder addDetailsBuilder() {
      return getDetailsFieldBuilder().addBuilder(
          org.mojolang.mojo.core.Any.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public org.mojolang.mojo.core.Any.Builder addDetailsBuilder(
        int index) {
      return getDetailsFieldBuilder().addBuilder(
          index, org.mojolang.mojo.core.Any.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Any details = 10;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Any.Builder> 
         getDetailsBuilderList() {
      return getDetailsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Any, org.mojolang.mojo.core.Any.Builder, org.mojolang.mojo.core.AnyOrBuilder> 
        getDetailsFieldBuilder() {
      if (detailsBuilder_ == null) {
        detailsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            org.mojolang.mojo.core.Any, org.mojolang.mojo.core.Any.Builder, org.mojolang.mojo.core.AnyOrBuilder>(
                details_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        details_ = null;
      }
      return detailsBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Error)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Error)
  private static final org.mojolang.mojo.core.Error DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Error();
  }

  public static org.mojolang.mojo.core.Error getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Error>
      PARSER = new com.google.protobuf.AbstractParser<Error>() {
    @java.lang.Override
    public Error parsePartialFrom(
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

  public static com.google.protobuf.Parser<Error> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Error> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Error getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

