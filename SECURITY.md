# Security Policy

## Supported Versions

Security fixes are released for the latest stable version of `pswg`.

## Reporting A Vulnerability

Please report suspected vulnerabilities privately through GitHub Security Advisories:

https://github.com/olelbis/pswg/security/advisories/new

If GitHub Security Advisories are not available to you, open a minimal public issue that says you need a private security contact. Do not include exploit details, generated secrets, or sensitive logs in the public issue.

## Scope

In scope:

- predictable or biased password generation
- incorrect validation of length or character-count rules
- unsafe release, package, checksum, SBOM, or attestation behavior
- documentation that could lead users to unsafe shell usage

Out of scope:

- passwords exposed after `pswg` prints them to stdout
- compromised terminals, shells, hosts, logs, process inspection, or clipboard managers
- password policies enforced by third-party services
- denial-of-service reports that require local shell access only

## Release Integrity

Stable releases publish:

- SHA-256 checksums
- Linux `.deb` and `.rpm` packages
- platform tarballs
- SPDX SBOMs
- GitHub artifact attestations signed through Sigstore

After downloading an artifact, verify provenance with:

```sh
gh attestation verify path/to/artifact -R olelbis/pswg
```

Verify checksums with:

```sh
shasum -a 256 -c SHA256SUMS
```
