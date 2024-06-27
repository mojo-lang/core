// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/checksum.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Checksum}
 */
public final class Checksum extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Checksum)
    ChecksumOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Checksum.newBuilder() to construct.
  private Checksum(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Checksum() {
    algorithm_ = 0;
    val_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Checksum();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.ChecksumProto.internal_static_mojo_core_Checksum_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ChecksumProto.internal_static_mojo_core_Checksum_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Checksum.class, org.mojolang.mojo.core.Checksum.Builder.class);
  }

  /**
   * Protobuf enum {@code mojo.core.Checksum.Algorithm}
   */
  public enum Algorithm
      implements com.google.protobuf.ProtocolMessageEnum {
    /**
     * <code>ALGORITHM_UNSPECIFIED = 0;</code>
     */
    ALGORITHM_UNSPECIFIED(0),
    /**
     * <code>ALGORITHM_MD5 = 1;</code>
     */
    ALGORITHM_MD5(1),
    /**
     * <code>ALGORITHM_SHA1 = 2;</code>
     */
    ALGORITHM_SHA1(2),
    /**
     * <code>ALGORITHM_SHA256 = 3;</code>
     */
    ALGORITHM_SHA256(3),
    /**
     * <code>ALGORITHM_SHA512 = 4;</code>
     */
    ALGORITHM_SHA512(4),
    UNRECOGNIZED(-1),
    ;

    /**
     * <code>ALGORITHM_UNSPECIFIED = 0;</code>
     */
    public static final int ALGORITHM_UNSPECIFIED_VALUE = 0;
    /**
     * <code>ALGORITHM_MD5 = 1;</code>
     */
    public static final int ALGORITHM_MD5_VALUE = 1;
    /**
     * <code>ALGORITHM_SHA1 = 2;</code>
     */
    public static final int ALGORITHM_SHA1_VALUE = 2;
    /**
     * <code>ALGORITHM_SHA256 = 3;</code>
     */
    public static final int ALGORITHM_SHA256_VALUE = 3;
    /**
     * <code>ALGORITHM_SHA512 = 4;</code>
     */
    public static final int ALGORITHM_SHA512_VALUE = 4;


    public final int getNumber() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalArgumentException(
            "Can't get the number of an unknown enum value.");
      }
      return value;
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     * @deprecated Use {@link #forNumber(int)} instead.
     */
    @java.lang.Deprecated
    public static Algorithm valueOf(int value) {
      return forNumber(value);
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     */
    public static Algorithm forNumber(int value) {
      switch (value) {
        case 0: return ALGORITHM_UNSPECIFIED;
        case 1: return ALGORITHM_MD5;
        case 2: return ALGORITHM_SHA1;
        case 3: return ALGORITHM_SHA256;
        case 4: return ALGORITHM_SHA512;
        default: return null;
      }
    }

    public static com.google.protobuf.Internal.EnumLiteMap<Algorithm>
        internalGetValueMap() {
      return internalValueMap;
    }
    private static final com.google.protobuf.Internal.EnumLiteMap<
        Algorithm> internalValueMap =
          new com.google.protobuf.Internal.EnumLiteMap<Algorithm>() {
            public Algorithm findValueByNumber(int number) {
              return Algorithm.forNumber(number);
            }
          };

    public final com.google.protobuf.Descriptors.EnumValueDescriptor
        getValueDescriptor() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalStateException(
            "Can't get the descriptor of an unrecognized enum value.");
      }
      return getDescriptor().getValues().get(ordinal());
    }
    public final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptorForType() {
      return getDescriptor();
    }
    public static final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptor() {
      return org.mojolang.mojo.core.Checksum.getDescriptor().getEnumTypes().get(0);
    }

    private static final Algorithm[] VALUES = values();

    public static Algorithm valueOf(
        com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
      if (desc.getType() != getDescriptor()) {
        throw new java.lang.IllegalArgumentException(
          "EnumValueDescriptor is not for this type.");
      }
      if (desc.getIndex() == -1) {
        return UNRECOGNIZED;
      }
      return VALUES[desc.getIndex()];
    }

    private final int value;

    private Algorithm(int value) {
      this.value = value;
    }

    // @@protoc_insertion_point(enum_scope:mojo.core.Checksum.Algorithm)
  }

  public static final int ALGORITHM_FIELD_NUMBER = 1;
  private int algorithm_;
  /**
   * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
   * @return The enum numeric value on the wire for algorithm.
   */
  @java.lang.Override public int getAlgorithmValue() {
    return algorithm_;
  }
  /**
   * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
   * @return The algorithm.
   */
  @java.lang.Override public org.mojolang.mojo.core.Checksum.Algorithm getAlgorithm() {
    @SuppressWarnings("deprecation")
    org.mojolang.mojo.core.Checksum.Algorithm result = org.mojolang.mojo.core.Checksum.Algorithm.valueOf(algorithm_);
    return result == null ? org.mojolang.mojo.core.Checksum.Algorithm.UNRECOGNIZED : result;
  }

  public static final int VAL_FIELD_NUMBER = 2;
  private volatile java.lang.Object val_;
  /**
   * <code>string val = 2;</code>
   * @return The val.
   */
  @java.lang.Override
  public java.lang.String getVal() {
    java.lang.Object ref = val_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      val_ = s;
      return s;
    }
  }
  /**
   * <code>string val = 2;</code>
   * @return The bytes for val.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getValBytes() {
    java.lang.Object ref = val_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      val_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
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
    if (algorithm_ != org.mojolang.mojo.core.Checksum.Algorithm.ALGORITHM_UNSPECIFIED.getNumber()) {
      output.writeEnum(1, algorithm_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(val_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, val_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (algorithm_ != org.mojolang.mojo.core.Checksum.Algorithm.ALGORITHM_UNSPECIFIED.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, algorithm_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(val_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, val_);
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
    if (!(obj instanceof org.mojolang.mojo.core.Checksum)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Checksum other = (org.mojolang.mojo.core.Checksum) obj;

    if (algorithm_ != other.algorithm_) return false;
    if (!getVal()
        .equals(other.getVal())) return false;
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
    hash = (37 * hash) + ALGORITHM_FIELD_NUMBER;
    hash = (53 * hash) + algorithm_;
    hash = (37 * hash) + VAL_FIELD_NUMBER;
    hash = (53 * hash) + getVal().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Checksum parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Checksum parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Checksum parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Checksum parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Checksum prototype) {
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
   * Protobuf type {@code mojo.core.Checksum}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Checksum)
      org.mojolang.mojo.core.ChecksumOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ChecksumProto.internal_static_mojo_core_Checksum_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ChecksumProto.internal_static_mojo_core_Checksum_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Checksum.class, org.mojolang.mojo.core.Checksum.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Checksum.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      algorithm_ = 0;

      val_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ChecksumProto.internal_static_mojo_core_Checksum_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Checksum getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Checksum.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Checksum build() {
      org.mojolang.mojo.core.Checksum result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Checksum buildPartial() {
      org.mojolang.mojo.core.Checksum result = new org.mojolang.mojo.core.Checksum(this);
      result.algorithm_ = algorithm_;
      result.val_ = val_;
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
      if (other instanceof org.mojolang.mojo.core.Checksum) {
        return mergeFrom((org.mojolang.mojo.core.Checksum)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Checksum other) {
      if (other == org.mojolang.mojo.core.Checksum.getDefaultInstance()) return this;
      if (other.algorithm_ != 0) {
        setAlgorithmValue(other.getAlgorithmValue());
      }
      if (!other.getVal().isEmpty()) {
        val_ = other.val_;
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
            case 8: {
              algorithm_ = input.readEnum();

              break;
            } // case 8
            case 18: {
              val_ = input.readStringRequireUtf8();

              break;
            } // case 18
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

    private int algorithm_ = 0;
    /**
     * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
     * @return The enum numeric value on the wire for algorithm.
     */
    @java.lang.Override public int getAlgorithmValue() {
      return algorithm_;
    }
    /**
     * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
     * @param value The enum numeric value on the wire for algorithm to set.
     * @return This builder for chaining.
     */
    public Builder setAlgorithmValue(int value) {
      
      algorithm_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
     * @return The algorithm.
     */
    @java.lang.Override
    public org.mojolang.mojo.core.Checksum.Algorithm getAlgorithm() {
      @SuppressWarnings("deprecation")
      org.mojolang.mojo.core.Checksum.Algorithm result = org.mojolang.mojo.core.Checksum.Algorithm.valueOf(algorithm_);
      return result == null ? org.mojolang.mojo.core.Checksum.Algorithm.UNRECOGNIZED : result;
    }
    /**
     * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
     * @param value The algorithm to set.
     * @return This builder for chaining.
     */
    public Builder setAlgorithm(org.mojolang.mojo.core.Checksum.Algorithm value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      algorithm_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.Checksum.Algorithm algorithm = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearAlgorithm() {
      
      algorithm_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object val_ = "";
    /**
     * <code>string val = 2;</code>
     * @return The val.
     */
    public java.lang.String getVal() {
      java.lang.Object ref = val_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        val_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string val = 2;</code>
     * @return The bytes for val.
     */
    public com.google.protobuf.ByteString
        getValBytes() {
      java.lang.Object ref = val_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        val_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string val = 2;</code>
     * @param value The val to set.
     * @return This builder for chaining.
     */
    public Builder setVal(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      val_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string val = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearVal() {
      
      val_ = getDefaultInstance().getVal();
      onChanged();
      return this;
    }
    /**
     * <code>string val = 2;</code>
     * @param value The bytes for val to set.
     * @return This builder for chaining.
     */
    public Builder setValBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      val_ = value;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Checksum)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Checksum)
  private static final org.mojolang.mojo.core.Checksum DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Checksum();
  }

  public static org.mojolang.mojo.core.Checksum getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Checksum>
      PARSER = new com.google.protobuf.AbstractParser<Checksum>() {
    @java.lang.Override
    public Checksum parsePartialFrom(
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

  public static com.google.protobuf.Parser<Checksum> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Checksum> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Checksum getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

