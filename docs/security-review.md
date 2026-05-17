# Security Review

This document is the maintainer review checklist for `pswg`. It is not a third-party audit report.

## Current Assessment

`pswg` is a small local CLI with no network access, no persistent storage, and no password database. Its security-sensitive behavior is concentrated in:

- random character selection
- password shuffling
- validation of user-requested composition rules
- shell-safe output guidance
- release artifact integrity

## Review Checklist

### Randomness

- Password character selection uses Go `crypto/rand`.
- Shuffling uses Go `crypto/rand`.
- The code does not use `math/rand`.
- Random helper functions propagate entropy-source errors to callers.

### Composition Rules

- Password length is bounded by `MinPasswordLength` and `MaxPasswordLength`.
- Uppercase, special, and numeric counts reject negative values.
- Character requirements cannot exceed total length.
- `-safe` changes only the special-character pool.

### Shell Safety

- Default special characters may contain shell metacharacters.
- The user manual documents quoted command substitution.
- The `-safe` mode uses a smaller special-character pool for shell-heavy workflows.
- Documentation still recommends quoting variables even with `-safe`.

### Release Integrity

- Release artifacts are built in GitHub Actions from a pushed tag.
- Release assets include SHA-256 checksums.
- Release assets include an SPDX SBOM.
- Release artifacts and SBOMs are covered by GitHub artifact attestations.
- The release workflow uses OIDC-backed Sigstore signing through GitHub artifact attestations.

### Open Risks

- The project has not had an independent external security audit.
- Generated passwords are visible to whatever can read stdout, terminal output, shell variables, process arguments of downstream commands, logs, or clipboard content.
- Shell-safe output reduces risk from common shell metacharacters but cannot make unquoted shell usage universally safe.
- Reproducible builds are not guaranteed because release metadata includes build time.

## External Audit Readiness

An external reviewer should receive:

- this review checklist
- `docs/user-manual.md`
- `SECURITY.md`
- release workflow logs for the audited version
- the `SHA256SUMS` and attestation verification commands for the audited version
