# Distribution

Official release assets are published on GitHub Releases:

https://github.com/olelbis/pswg/releases/latest

Each stable release includes:

- Darwin amd64/arm64 tarballs
- Linux amd64/arm64 tarballs
- Windows amd64/arm64 tarballs and zip archives
- Linux amd64/arm64 `.deb` packages
- Linux amd64/arm64 `.rpm` packages
- `SHA256SUMS`
- SPDX SBOM
- GitHub artifact attestations

## Verify A Download

Download `SHA256SUMS` from the same release as the artifact and run:

```sh
shasum -a 256 -c SHA256SUMS
```

Verify provenance with GitHub artifact attestations:

```sh
gh attestation verify pswg_VERSION_OS_ARCH.tar.gz -R olelbis/pswg
```

## Homebrew

Use `packaging/homebrew/Formula/pswg.rb.in` as the source for a Homebrew tap formula. Replace the template values with the release version and the SHA-256 digest of the Darwin tarball.

## Scoop

Use `packaging/scoop/pswg.json.in` as the source for a Scoop bucket manifest. Replace the template values with the release version and Windows zip SHA-256 digests.

## AUR

Use `packaging/aur/PKGBUILD.in` as the source for an AUR package. Replace the template values with the release version and Linux amd64/arm64 SHA-256 digests.

## WinGet

Use the templates in `packaging/winget/` as the source for a `microsoft/winget-pkgs` submission. Replace the template values with the release version, release date, and Windows zip SHA-256 digests from `SHA256SUMS`.
