# pswg User Manual

`pswg` generates passwords from a small set of explicit composition rules. It prints only the generated password to stdout, which makes it suitable for shell pipelines and command substitution.

## Commands

Generate a password:

```sh
pswg [flags]
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

The generated password length is exactly `-l`. After `-u`, `-s`, and `-n` are applied, all remaining characters are lowercase letters.

The effective default policy is:

```text
-l 12 -u 1 -s 1 -n 1
```

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

Special characters are printable ASCII for broad compatibility.

When possible, `pswg` shuffles the generated password so it does not start with a number.

## Build From Source

Run checks:

```sh
make check
```

Build a local binary:

```sh
make build VERSION=v1.0.3
./build/pswg -version
```

Build release archives with checksums:

```sh
make release VERSION=v1.0.3
```
