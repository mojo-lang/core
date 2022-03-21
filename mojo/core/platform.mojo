
@case("camel")
enum Arch {
    unspecified @0

    x86 @1
    amd64 @2

    arm @5
    arm64 @6

    wasm @10
}

@case("camel")
enum OS {
    unspecified @0

    android @1

    darwin @2

    ios @6 //@alias("iOS")

    linux @10

    windows @20
}

/// the platform information
type Platform {
    arch: Arch @1
    os: OS @2
    os_version: String @3
}
