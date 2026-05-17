# pswg

[![Go Version](https://img.shields.io/github/go-mod/go-version/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/go.mod)
[![CI](https://github.com/olelbis/pswg/actions/workflows/ci.yml/badge.svg)](https://github.com/olelbis/pswg/actions/workflows/ci.yml)
[![Version](https://img.shields.io/github/v/tag/olelbis/pswg?label=version)](https://github.com/olelbis/pswg/tags)
[![License: MIT](https://img.shields.io/github/license/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/LICENSE)

`pswg` is a small password generator CLI written in Go.

It uses `crypto/rand` for character selection and shuffling, keeps the generated password within a configurable length, and can require a minimum number of uppercase, special, and numeric characters.

## Usage

Generate a password with the default policy:

```sh
pswg
```

The password is printed by itself on stdout, so it can be used in scripts.

Generate a 16-character password with 2 uppercase letters, 2 special characters, and 2 numbers:

```sh
pswg -l 16 -u 2 -s 2 -n 2
```

Print the current version:

```sh
pswg -version
```

Print help:

```sh
pswg -h
```

## Options

```text
-l int
    Password length. Default: 12. Valid range: 12-128.
-u int
    Number of uppercase characters. Default: 1. Valid range: 0-length.
-s int
    Number of special characters. Default: 1. Valid range: 0-length.
-n int
    Number of numeric characters. Default: 1. Valid range: 0-length.
-version
    Print the current version. Cannot be combined with generation flags.
```

## Flag Combinations

Generation flags can be used in any order and any subset:

```sh
pswg -l 24
pswg -u 4 -s 4
pswg -n 0 -s 0
pswg -l 32 -n 8 -u 4 -s 4
```

The generated password length is exactly `-l`. The remaining characters after uppercase, special, and numeric requirements are lowercase letters.

Valid generation rules:

- `12 <= -l <= 128`
- `-u >= 0`
- `-s >= 0`
- `-n >= 0`
- `-u + -s + -n <= -l`

Invalid combinations fail with exit code `2`:

```sh
pswg -l 11
pswg -l 129
pswg -l 12 -u 13
pswg -u -1
pswg -version -l 16
pswg -version extra
pswg extra
```

Exit codes:

- `0`: password, help, or version printed successfully.
- `2`: invalid flags, arguments, or password policy.

## Build

Run the full local check:

```sh
make check
```

Build a local binary:

```sh
make build VERSION=v1.0.1
./build/pswg -version
```

Build release archives with checksums:

```sh
make release VERSION=v1.0.1
```

Pushing a `v*` tag runs the release workflow and publishes the generated archives plus `SHA256SUMS` to GitHub Releases.

## Notes

- Minimum password length is 12 characters.
- Maximum password length is 128 characters.
- Requested character counts cannot be negative or exceed the password length.
- Special characters are printable ASCII for broad compatibility.
- When possible, the generated password is shuffled so it does not start with a number.

## License

MIT. See [LICENSE](LICENSE).
