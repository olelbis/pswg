# pswg

[![Go Version](https://img.shields.io/github/go-mod/go-version/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/go.mod)
[![CI](https://github.com/olelbis/pswg/actions/workflows/ci.yml/badge.svg)](https://github.com/olelbis/pswg/actions/workflows/ci.yml)
[![Version](https://img.shields.io/github/v/tag/olelbis/pswg?include_prereleases&label=version)](https://github.com/olelbis/pswg/tags)
[![License: MIT](https://img.shields.io/github/license/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/LICENSE)
[![Status: experimental](https://img.shields.io/badge/status-experimental-orange)](https://github.com/olelbis/pswg)

`pswg` is a small experimental password generator written while learning Go.

It uses `crypto/rand` for character selection and shuffling, keeps the generated password within a configurable length, and can require a minimum number of uppercase, special, and numeric characters.

> Experimental software: use it as a learning project. It has not been audited as a security tool.

## Install

```sh
go install github.com/olelbis/pswg@v0.5.0-alpha
```

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
--silent
    Exit without output.
-version
    Print the current version.
```

## Build

Run the full local check:

```sh
make check
```

Build a local binary with version metadata:

```sh
make build
./bin/pswg -version
```

Build release-style binaries for supported targets:

```sh
make dist
```

## Notes

- Minimum password length is 12 characters.
- Maximum password length is 128 characters.
- Requested character counts cannot be negative or exceed the password length.
- When possible, the generated password is shuffled so it does not start with a number.

## License

MIT. See [LICENSE](LICENSE).
