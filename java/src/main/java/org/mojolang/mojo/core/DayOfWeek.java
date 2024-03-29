// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/time.proto

package org.mojolang.mojo.core;

/**
 * Protobuf enum {@code mojo.core.DayOfWeek}
 */
public enum DayOfWeek
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <code>DAY_OF_WEEK_UNSPECIFIED = 0;</code>
   */
  DAY_OF_WEEK_UNSPECIFIED(0),
  /**
   * <code>DAY_OF_WEEK_MONDAY = 1;</code>
   */
  DAY_OF_WEEK_MONDAY(1),
  /**
   * <code>DAY_OF_WEEK_TUESDAY = 2;</code>
   */
  DAY_OF_WEEK_TUESDAY(2),
  /**
   * <code>DAY_OF_WEEK_WEDNESDAY = 3;</code>
   */
  DAY_OF_WEEK_WEDNESDAY(3),
  /**
   * <code>DAY_OF_WEEK_THURSDAY = 4;</code>
   */
  DAY_OF_WEEK_THURSDAY(4),
  /**
   * <code>DAY_OF_WEEK_FRIDAY = 5;</code>
   */
  DAY_OF_WEEK_FRIDAY(5),
  /**
   * <code>DAY_OF_WEEK_SATURDAY = 6;</code>
   */
  DAY_OF_WEEK_SATURDAY(6),
  /**
   * <code>DAY_OF_WEEK_SUNDAY = 7;</code>
   */
  DAY_OF_WEEK_SUNDAY(7),
  UNRECOGNIZED(-1),
  ;

  /**
   * <code>DAY_OF_WEEK_UNSPECIFIED = 0;</code>
   */
  public static final int DAY_OF_WEEK_UNSPECIFIED_VALUE = 0;
  /**
   * <code>DAY_OF_WEEK_MONDAY = 1;</code>
   */
  public static final int DAY_OF_WEEK_MONDAY_VALUE = 1;
  /**
   * <code>DAY_OF_WEEK_TUESDAY = 2;</code>
   */
  public static final int DAY_OF_WEEK_TUESDAY_VALUE = 2;
  /**
   * <code>DAY_OF_WEEK_WEDNESDAY = 3;</code>
   */
  public static final int DAY_OF_WEEK_WEDNESDAY_VALUE = 3;
  /**
   * <code>DAY_OF_WEEK_THURSDAY = 4;</code>
   */
  public static final int DAY_OF_WEEK_THURSDAY_VALUE = 4;
  /**
   * <code>DAY_OF_WEEK_FRIDAY = 5;</code>
   */
  public static final int DAY_OF_WEEK_FRIDAY_VALUE = 5;
  /**
   * <code>DAY_OF_WEEK_SATURDAY = 6;</code>
   */
  public static final int DAY_OF_WEEK_SATURDAY_VALUE = 6;
  /**
   * <code>DAY_OF_WEEK_SUNDAY = 7;</code>
   */
  public static final int DAY_OF_WEEK_SUNDAY_VALUE = 7;


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
  public static DayOfWeek valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static DayOfWeek forNumber(int value) {
    switch (value) {
      case 0: return DAY_OF_WEEK_UNSPECIFIED;
      case 1: return DAY_OF_WEEK_MONDAY;
      case 2: return DAY_OF_WEEK_TUESDAY;
      case 3: return DAY_OF_WEEK_WEDNESDAY;
      case 4: return DAY_OF_WEEK_THURSDAY;
      case 5: return DAY_OF_WEEK_FRIDAY;
      case 6: return DAY_OF_WEEK_SATURDAY;
      case 7: return DAY_OF_WEEK_SUNDAY;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<DayOfWeek>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      DayOfWeek> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<DayOfWeek>() {
          public DayOfWeek findValueByNumber(int number) {
            return DayOfWeek.forNumber(number);
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
    return org.mojolang.mojo.core.TimeProto.getDescriptor().getEnumTypes().get(1);
  }

  private static final DayOfWeek[] VALUES = values();

  public static DayOfWeek valueOf(
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

  private DayOfWeek(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:mojo.core.DayOfWeek)
}

