// Copyright 2021 Mojo-lang.org
//
// Copyright (C) 2009 The Libphonenumber Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// [source](https://github.com/google/libphonenumber)

/// Definition for representing international telephone numbers.
///
/// Examples:
///
/// Google MTV, +1 650-253-0000, (650) 253-0000
/// country_code: 1
/// national_number: 6502530000
///
/// Google Paris, +33 (0)1 42 68 53 00, 01 42 68 53 00
/// country_code: 33
/// national_number: 142685300
///
/// Google Beijing, +86-10-62503000, (010) 62503000
/// country_code: 86
/// national_number: 1062503000
///
/// Google Italy, +39 02-36618 300, 02-36618 300
/// country_code: 39
/// national_number: 236618300
/// italian_leading_zero: true
type PhoneNumber {
  /// The country calling code for this number, as defined by the International
  /// Telecommunication Union (ITU). For example, this would be 1 for NANPA
  /// countries, and 33 for France.
  country_code: Int32 @1

  /// The National (significant) Number, as defined in International
  /// Telecommunication Union (ITU) Recommendation E.164, without any leading
  /// zero. The leading-zero is stored separately if required, since this is an
  /// uint64 and hence cannot store such information. Do not use this field
  /// directly: if you want the national significant number, call the
  /// getNationalSignificantNumber method of PhoneNumberUtil.
  ///
  /// For countries which have the concept of an "area code" or "national
  /// destination code", this is included in the National (significant) Number.
  /// Although the ITU says the maximum length should be 15, we have found longer
  /// numbers in some countries e.g. Germany.
  /// Note that the National (significant) Number does not contain the National
  /// (trunk) prefix. Obviously, as a uint64, it will never contain any
  /// formatting (hyphens, spaces, parentheses), nor any alphanumeric spellings.
  national_number: UInt64 @2 @required

  /// Extension is not standardized in ITU recommendations, except for being
  /// defined as a series of numbers with a maximum length of 40 digits. It is
  /// defined as a string here to accommodate for the possible use of a leading
  /// zero in the extension (organizations have complete freedom to do so, as
  /// there is no standard defined). Other than digits, some other dialling
  /// characters such as "," (indicating a wait) may be stored here.
  extension: String @3

  /// In some countries, the national (significant) number starts with one or
  /// more "0"s without this being a national prefix or trunk code of some kind.
  /// For example, the leading zero in the national (significant) number of an
  /// Italian phone number indicates the number is a fixed-line number.  There
  /// have been plans to migrate fixed-line numbers to start with the digit two
  /// since December 2000, but it has not happened yet. See
  /// http://en.wikipedia.org/wiki/%2B39 for more details.
  ///
  /// These fields can be safely ignored (there is no need to set them) for most
  /// countries. Some limited number of countries behave like Italy - for these
  /// cases, if the leading zero(s) of a number would be retained even when
  /// dialling internationally, set this flag to true, and also set the number of
  /// leading zeros.
  ///
  /// Clients who use the parsing functionality of the i18n phone
  /// number libraries will have these fields set if necessary automatically.
  italian_leading_zero: Bool @4
  number_of_leading_zeros: Int32 @8 = 1

  /// The next few fields are non-essential fields for a phone number. They
  /// retain extra information about the form the phone number was in when it was
  /// provided to us to parse. They can be safely ignored by most clients. To
  /// populate them, call parseAndKeepRawInput on PhoneNumberUtil.
  ///
  /// This field is used to store the raw input string containing phone numbers
  /// before it was canonicalized by the library. For example, it could be used
  /// to store alphanumerical numbers such as "1-800-GOOG-411".
  raw_input: String @5

  /// The source from which the country_code is derived. This is not set in the
  /// general parsing method, but in the method that parses and keeps raw_input.
  /// New fields could be added upon request.
  enum CountryCodeSource {
    /// Default value returned if this is not set, because the phone number was
    /// created using parse, not parseAndKeepRawInput. hasCountryCodeSource will
    /// return false if this is the case.
    unspecified @0

    /// The country_code is derived based on a phone number with a leading "+",
    /// e.g. the French number "+33 1 42 68 53 00".
    from_number_with_plus_sign @1

    /// The country_code is derived based on a phone number with a leading IDD,
    /// e.g. the French number "011 33 1 42 68 53 00", as it is dialled from US.
    from_number_with_idd @5

    /// The country_code is derived based on a phone number without a leading
    /// "+", e.g. the French number "33 1 42 68 53 00" when defaultCountry is
    /// supplied as France.
    from_number_without_plus_sign @10

    /// The country_code is derived NOT based on the phone number itself, but
    /// from the defaultCountry parameter provided in the parsing function by the
    /// clients. This happens mostly for numbers written in the national format
    /// (without country code). For example, this would be set when parsing the
    /// French number "01 42 68 53 00", when defaultCountry is supplied as
    /// France.
    from_default_country @20
  }

  /// The source from which the country_code is derived.
  country_code_source: CountryCodeSource @6

  /// The carrier selection code that is preferred when calling this phone number
  /// domestically. This also includes codes that need to be dialed in some
  /// countries when calling from landlines to mobiles or vice versa. For
  /// example, in Columbia, a "3" needs to be dialed before the phone number
  /// itself when calling from a mobile phone to a domestic landline phone and
  /// vice versa.
  ///
  /// Note this is the "preferred" code, which means other codes may work as
  /// well.
  preferred_domestic_carrier_code: String @7
}