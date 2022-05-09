// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/value.proto

package org.mojolang.mojo.core;

/**
 * Protobuf enum {@code mojo.core.ValueKind}
 */
public enum ValueKind
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <code>VALUE_KIND_UNSPECIFIED = 0;</code>
   */
  VALUE_KIND_UNSPECIFIED(0),
  /**
   * <code>VALUE_KIND_NULL = 1;</code>
   */
  VALUE_KIND_NULL(1),
  /**
   * <code>VALUE_KIND_BOOLEAN = 2;</code>
   */
  VALUE_KIND_BOOLEAN(2),
  /**
   * <code>VALUE_KIND_INTEGER = 3;</code>
   */
  VALUE_KIND_INTEGER(3),
  /**
   * <code>VALUE_KIND_NUMBER = 4;</code>
   */
  VALUE_KIND_NUMBER(4),
  /**
   * <code>VALUE_KIND_STRING = 5;</code>
   */
  VALUE_KIND_STRING(5),
  /**
   * <code>VALUE_KIND_BYTES = 6;</code>
   */
  VALUE_KIND_BYTES(6),
  /**
   * <code>VALUE_KIND_ARRAY = 10;</code>
   */
  VALUE_KIND_ARRAY(10),
  /**
   * <code>VALUE_KIND_OBJECT = 11;</code>
   */
  VALUE_KIND_OBJECT(11),
  UNRECOGNIZED(-1),
  ;

  /**
   * <code>VALUE_KIND_UNSPECIFIED = 0;</code>
   */
  public static final int VALUE_KIND_UNSPECIFIED_VALUE = 0;
  /**
   * <code>VALUE_KIND_NULL = 1;</code>
   */
  public static final int VALUE_KIND_NULL_VALUE = 1;
  /**
   * <code>VALUE_KIND_BOOLEAN = 2;</code>
   */
  public static final int VALUE_KIND_BOOLEAN_VALUE = 2;
  /**
   * <code>VALUE_KIND_INTEGER = 3;</code>
   */
  public static final int VALUE_KIND_INTEGER_VALUE = 3;
  /**
   * <code>VALUE_KIND_NUMBER = 4;</code>
   */
  public static final int VALUE_KIND_NUMBER_VALUE = 4;
  /**
   * <code>VALUE_KIND_STRING = 5;</code>
   */
  public static final int VALUE_KIND_STRING_VALUE = 5;
  /**
   * <code>VALUE_KIND_BYTES = 6;</code>
   */
  public static final int VALUE_KIND_BYTES_VALUE = 6;
  /**
   * <code>VALUE_KIND_ARRAY = 10;</code>
   */
  public static final int VALUE_KIND_ARRAY_VALUE = 10;
  /**
   * <code>VALUE_KIND_OBJECT = 11;</code>
   */
  public static final int VALUE_KIND_OBJECT_VALUE = 11;


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
  public static ValueKind valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static ValueKind forNumber(int value) {
    switch (value) {
      case 0: return VALUE_KIND_UNSPECIFIED;
      case 1: return VALUE_KIND_NULL;
      case 2: return VALUE_KIND_BOOLEAN;
      case 3: return VALUE_KIND_INTEGER;
      case 4: return VALUE_KIND_NUMBER;
      case 5: return VALUE_KIND_STRING;
      case 6: return VALUE_KIND_BYTES;
      case 10: return VALUE_KIND_ARRAY;
      case 11: return VALUE_KIND_OBJECT;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<ValueKind>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      ValueKind> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<ValueKind>() {
          public ValueKind findValueByNumber(int number) {
            return ValueKind.forNumber(number);
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
    return org.mojolang.mojo.core.ValueProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final ValueKind[] VALUES = values();

  public static ValueKind valueOf(
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

  private ValueKind(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:mojo.core.ValueKind)
}

