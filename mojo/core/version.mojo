/// Semantic Versioning
///
/// (spec)[http://semver.org/spec/v2.0.0.html]
///
@template("{major}.{minor}.{path}[-{pre_release}][+{build_metadata}]")
type Version {
    type Identifier = String @format(r"[0-9A-Za-z-]+")

    ///
    major: UInt @1
    
    ///
    minor: UInt @2
    
    ///
    patch: UInt @3
    
    ///
    pre_release:    [Identifier] @5
    
    ///
    build_metadata: [Identifier] @6
}

/// Precedence refers to how versions are compared to each other when ordered. Precedence MUST be calculated by separating the version into major, minor, patch and pre-release identifiers in that order (Build metadata does not figure into precedence). Precedence is determined by the first difference when comparing each of these identifiers from left to right as follows: Major, minor, and patch versions are always compared numerically. Example: 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1. When major, minor, and patch are equal, a pre-release version has lower precedence than a normal version. Example: 1.0.0-alpha < 1.0.0. Precedence for two pre-release versions with the same major, minor, and patch version MUST be determined by comparing each dot separated identifier from left to right until a difference is found as follows: identifiers consisting of only digits are compared numerically and identifiers with letters or hyphens are compared lexically in ASCII sort order. Numeric identifiers always have lower precedence than non-numeric identifiers. A larger set of pre-release fields has a higher precedence than a smaller set, if all of the preceding identifiers are equal. Example: 1.0.0-alpha < 1.0.0-alpha.1 < 1.0.0-alpha.beta < 1.0.0-beta < 1.0.0-beta.2 < 1.0.0-beta.11 < 1.0.0-rc.1 < 1.0.0.
func less(left:Version, right:Version) -> Bool

/// Major version zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable.
func is_initial_development(version: Version) -> Bool {
    return version.major == 0
}