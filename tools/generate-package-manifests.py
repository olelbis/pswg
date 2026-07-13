#!/usr/bin/env python3
"""Generate package-manager manifests from release checksums."""

from __future__ import annotations

import argparse
import datetime as dt
import hashlib
import pathlib
import shutil
import tarfile


ROOT = pathlib.Path(__file__).resolve().parents[1]
TEMPLATE_DIR = ROOT / "packaging"


def read_checksums(path: pathlib.Path) -> dict[str, str]:
    checksums: dict[str, str] = {}
    for line in path.read_text().splitlines():
        digest, filename = line.split(maxsplit=1)
        checksums[filename] = digest
    return checksums


def checksum_for(checksums: dict[str, str], filename: str) -> str:
    try:
        return checksums[filename]
    except KeyError as exc:
        raise SystemExit(f"missing checksum for {filename}") from exc


def render_template(path: pathlib.Path, replacements: dict[str, str]) -> str:
    rendered = path.read_text()
    for key, value in replacements.items():
        rendered = rendered.replace(f"{{{{{key}}}}}", value)
    return rendered


def write_rendered_tree(output_dir: pathlib.Path, replacements: dict[str, str]) -> None:
    if output_dir.exists():
        shutil.rmtree(output_dir)
    output_dir.mkdir(parents=True)

    for template in sorted(TEMPLATE_DIR.rglob("*.in")):
        relative = template.relative_to(TEMPLATE_DIR)
        target = output_dir / relative.with_suffix("")
        target.parent.mkdir(parents=True, exist_ok=True)
        target.write_text(render_template(template, replacements))


def create_archive(output_dir: pathlib.Path, archive_path: pathlib.Path) -> None:
    if archive_path.exists():
        archive_path.unlink()
    with tarfile.open(archive_path, "w:gz") as archive:
        archive.add(output_dir, arcname=output_dir.name)


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--version", required=True, help="release version, with or without leading v")
    parser.add_argument("--dist", default="dist")
    parser.add_argument("--output", default="build/package-manifests")
    parser.add_argument("--release-date", default=dt.date.today().isoformat())
    args = parser.parse_args()

    version = args.version.removeprefix("v")
    tag = f"v{version}"
    dist = ROOT / args.dist
    output_dir = ROOT / args.output
    checksums = read_checksums(dist / "SHA256SUMS")

    replacements = {
        "VERSION": version,
        "RELEASE_DATE": args.release_date,
        "LICENSE_SHA256": hashlib.sha256((ROOT / "LICENSE").read_bytes()).hexdigest(),
        "DARWIN_AMD64_SHA256": checksum_for(checksums, f"pswg_{tag}_darwin_amd64.tar.gz"),
        "DARWIN_ARM64_SHA256": checksum_for(checksums, f"pswg_{tag}_darwin_arm64.tar.gz"),
        "LINUX_AMD64_SHA256": checksum_for(checksums, f"pswg_{tag}_linux_amd64.tar.gz"),
        "LINUX_ARM64_SHA256": checksum_for(checksums, f"pswg_{tag}_linux_arm64.tar.gz"),
        "WINDOWS_AMD64_SHA256": checksum_for(checksums, f"pswg_{tag}_windows_amd64.zip"),
        "WINDOWS_ARM64_SHA256": checksum_for(checksums, f"pswg_{tag}_windows_arm64.zip"),
    }

    write_rendered_tree(output_dir, replacements)
    create_archive(output_dir, dist / f"pswg_{tag}_package_manifests.tar.gz")
    print(f"Generated package manifests for {tag}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
