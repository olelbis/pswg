# pswg User Manual

`pswg` generates passwords from a small set of explicit composition rules. It prints only the generated password to stdout, which makes it suitable for shell pipelines and command substitution.

## Commands

Generate a password:

```sh
pswg [generation flags]
```

Show version information:

```sh
pswg -version
```

Show help:

```sh
pswg -h
```

## Generation Flags

Generation flags can be used in any order and any subset.

| Flag | Meaning | Default | Valid range |
| --- | --- | ---: | --- |
| `-l int` | Total password length | `12` | `12` to `128` |
| `-u int` | Number of uppercase letters | `1` | `0` to `-l` |
| `-s int` | Number of special characters | `1` | `0` to `-l` |
| `-n int` | Number of numeric characters | `1` | `0` to `-l` |
| `-safe` | Use shell-safe special characters | `false` | boolean |

The generated password length is exactly `-l`. After `-u`, `-s`, and `-n` are applied, all remaining characters are lowercase letters.

The effective default policy is:

```text
-l 12 -u 1 -s 1 -n 1
```

`-safe` changes only the special-character pool. It does not change length or character counts.

## Valid Combinations

Use only defaults:

```sh
pswg
```

Change only length:

```sh
pswg -l 24
```

Disable numeric and special-character requirements:

```sh
pswg -n 0 -s 0
```

Require more uppercase and special characters:

```sh
pswg -u 4 -s 4
```

Set a full custom policy:

```sh
pswg -l 32 -n 8 -u 4 -s 4
```

Use shell-safe special characters:

```sh
pswg -safe
pswg -l 24 -u 4 -s 4 -n 4 -safe
```

Generate a password made only of uppercase letters:

```sh
pswg -l 12 -u 12 -s 0 -n 0
```

Generate a password made only of numbers:

```sh
pswg -l 12 -u 0 -s 0 -n 12
```

## Validation Rules

Generation succeeds only when all rules are true:

- `12 <= -l <= 128`
- `-u >= 0`
- `-s >= 0`
- `-n >= 0`
- `-u + -s + -n <= -l`

`-version` is informational and cannot be combined with generation flags or positional arguments.

`pswg` does not accept positional arguments.

## Shell Usage

By default, `pswg` can generate shell metacharacters such as `*`, `?`, `[`, `]`, `$`, `&`, `;`, `<`, and `>`.

That is fine when the password is quoted. In scripts, prefer this pattern:

```sh
password="$(pswg)"
some-command --password "$password"
```

Avoid unquoted command substitution:

```sh
some-command --password $(pswg)
```

Unquoted command substitution can trigger shell behavior such as pathname expansion for characters like `*`, `?`, and `[`. The shell parses operators before expansion, so not every metacharacter becomes syntax, but unquoted passwords are still fragile and should be avoided.

If you need a password that is safer for unquoted shell usage, use `-safe`:

```sh
password="$(pswg -safe)"
some-command --password "$password"
```

`-safe` uses this special-character pool:

```text
@_:,.
```

Even with `-safe`, quoting variables is still the recommended shell practice.

## Security Model

`pswg` uses Go's `crypto/rand` package for random selection and shuffling. It is intended to reduce predictable password generation and to make composition rules explicit.

`pswg` is not a password manager. It does not store, sync, rotate, encrypt, check reuse, or validate destination-specific password policies. The project has not had an external security audit.

In scope:

- generating passwords with operating-system cryptographic randomness
- enforcing the requested length and character counts
- avoiding shell-fragile special characters when `-safe` is used

Out of scope:

- protecting passwords after they are printed to stdout
- compromised terminals, shells, hosts, logs, process inspection, or clipboard managers
- unsafe script usage such as unquoted command substitution
- validating whether a generated password is accepted by a specific service

For scripts, capture and pass passwords through quoted variables:

```sh
password="$(pswg -safe)"
some-command --password "$password"
```

See also:

- [Security policy](../SECURITY.md)
- [Security review checklist](security-review.md)

## Invalid Examples

Length below the minimum:

```sh
pswg -l 11
```

Length above the maximum:

```sh
pswg -l 129
```

Requested character counts exceed total length:

```sh
pswg -l 12 -u 13
```

Negative character counts:

```sh
pswg -u -1
```

Combining `-version` with generation flags:

```sh
pswg -version -l 16
```

Extra positional arguments:

```sh
pswg extra
pswg -version extra
```

## Output And Exit Codes

Successful password generation writes the password to stdout and exits with code `0`.

```sh
pswg
```

Errors are written to stderr and exit with code `2`.

| Exit code | Meaning |
| ---: | --- |
| `0` | Password, help, or version printed successfully |
| `2` | Invalid flags, arguments, or password policy |

## Character Sets

`pswg` uses these character pools:

| Category | Pool |
| --- | --- |
| Lowercase | `abcdefghijklmnopqrstuvwxyz` |
| Uppercase | `ABCDEFGHIJKLMNOPQRSTUVWXYZ` |
| Numeric | `1234567890` |
| Special | `!&%$=?^+*][{}-_.:,;()><` |
| Shell-safe special | `@_:,.` |

Special characters are printable ASCII for broad compatibility.

When possible, `pswg` shuffles the generated password so it does not start with a number.

## Build From Source

Run checks:

```sh
make check
```

Build a local binary:

```sh
make build VERSION=vX.Y.Z
./build/pswg -version
```

Build release archives, Linux packages, and checksums:

```sh
make release VERSION=vX.Y.Z
```

`make release` produces:

- `.tar.gz` archives for Darwin arm64/amd64, Linux arm64/amd64, and Windows arm64/amd64
- `.zip` archives for Windows arm64/amd64
- `.deb` packages for Linux
- `.rpm` packages for Linux
- `SHA256SUMS`

Package builds require `nfpm`.

The release workflow installs `nfpm` automatically. For local package builds, install it with:

```sh
go install github.com/goreleaser/nfpm/v2/cmd/nfpm@v2.46.3
```

## Release Verification

GitHub releases include SHA-256 checksums, an SPDX SBOM, and Sigstore-backed GitHub artifact attestations.

Verify checksums:

```sh
shasum -a 256 -c SHA256SUMS
```

Verify artifact provenance:

```sh
gh attestation verify pswg_VERSION_OS_ARCH.tar.gz -R olelbis/pswg
```

Read [distribution notes](distribution.md) for Homebrew, Scoop, WinGet, and AUR packaging templates.

## Installed Files In Linux Packages

The `.deb` and `.rpm` packages install:

| Source | Destination |
| --- | --- |
| `pswg` | `/usr/bin/pswg` |
| `docs/pswg.1` | `/usr/share/man/man1/pswg.1` |
| `LICENSE` | `/usr/share/doc/pswg/LICENSE` |
