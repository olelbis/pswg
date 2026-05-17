# pswg

[![Go Version](https://img.shields.io/github/go-mod/go-version/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/go.mod)
[![CI](https://github.com/olelbis/pswg/actions/workflows/ci.yml/badge.svg)](https://github.com/olelbis/pswg/actions/workflows/ci.yml)
[![Version](https://img.shields.io/github/v/tag/olelbis/pswg?label=version)](https://github.com/olelbis/pswg/tags)
[![License: MIT](https://img.shields.io/github/license/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/LICENSE)

`pswg` is a small password generator CLI written in Go.

It uses `crypto/rand` for character selection and shuffling, keeps the generated password within a configurable length, and can require a minimum number of uppercase, special, and numeric characters.

## Usage

Generate a password with the default policy. The password is printed by itself on stdout, so it can be used in scripts:

```sh
pswg
```

Generate a 16-character password with 2 uppercase letters, 2 special characters, and 2 numbers:

```sh
pswg -l 16 -u 2 -s 2 -n 2
```

Print the current version:

```sh
pswg -version
```

## Options

```text
-l int
    Password length. Default: 12. Maximum: 128.
-u int
    Number of uppercase characters. Default: 1.
-s int
    Number of special characters. Default: 1.
-n int
    Number of numeric characters. Default: 1.
-version
    Print the current version.
```

## Build

Run the full local check:

```sh
make check
```

Build a local binary:

```sh
make build VERSION=v1.0.0
./build/pswg -version
```

Build release archives with checksums:

```sh
make release VERSION=v1.0.0
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
