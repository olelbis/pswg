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
- generated package-manager manifests archive

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

Download `pswg_VERSION_package_manifests.tar.gz` from the release and use `homebrew/Formula/pswg.rb` as the source for a Homebrew tap formula.

## Scoop

Download `pswg_VERSION_package_manifests.tar.gz` from the release and use `scoop/pswg.json` as the source for a Scoop bucket manifest.

## AUR

Download `pswg_VERSION_package_manifests.tar.gz` from the release and use `aur/PKGBUILD` as the source for an AUR package.

## WinGet

Download `pswg_VERSION_package_manifests.tar.gz` from the release and use the files in `winget/` as the source for a `microsoft/winget-pkgs` submission.
