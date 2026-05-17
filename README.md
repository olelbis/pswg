<p align="center">
  <img src="assets/logo.svg" alt="pswg pixel art logo" width="128" height="128">
</p>

# pswg

[![Go Version](https://img.shields.io/github/go-mod/go-version/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/go.mod)
[![CI](https://github.com/olelbis/pswg/actions/workflows/ci.yml/badge.svg)](https://github.com/olelbis/pswg/actions/workflows/ci.yml)
[![Pages](https://github.com/olelbis/pswg/actions/workflows/pages.yml/badge.svg)](https://github.com/olelbis/pswg/actions/workflows/pages.yml)
[![Version](https://img.shields.io/github/v/tag/olelbis/pswg?label=version)](https://github.com/olelbis/pswg/tags)
[![License: MIT](https://img.shields.io/github/license/olelbis/pswg)](https://github.com/olelbis/pswg/blob/main/LICENSE)

`pswg` is a small password generator CLI written in Go.

It uses `crypto/rand`, prints only the generated password to stdout, and lets you set password length plus uppercase, special-character, and numeric requirements.

## Quick Start

Generate a password with the default policy:

```sh
pswg
```

Generate a 16-character password with 2 uppercase letters, 2 special characters, and 2 numbers:

```sh
pswg -l 16 -u 2 -s 2 -n 2
```

Generate a password with shell-safe special characters:

```sh
pswg -safe
```

Show version and help:

```sh
pswg -version
pswg -h
```

## Documentation

Project page: [olelbis.github.io/pswg](https://olelbis.github.io/pswg/)

Read the [user manual](docs/user-manual.md) for all flags, valid combinations, shell usage, security model, exit codes, character sets, packaging, and build commands.

Additional project docs:

- [Security policy](SECURITY.md)
- [Security review checklist](docs/security-review.md)
- [Distribution notes](docs/distribution.md)

## License

MIT. See [LICENSE](LICENSE).
