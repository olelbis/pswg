#!/usr/bin/env python3
"""Extract release notes for a version from CHANGELOG.md."""

from __future__ import annotations

import argparse
import pathlib
import re


ROOT = pathlib.Path(__file__).resolve().parents[1]


def release_notes(version: str) -> str:
    version = version.removeprefix("v")
    changelog = (ROOT / "CHANGELOG.md").read_text()
    pattern = re.compile(
        rf"^### \[{re.escape(version)}\] - .+?\n(?P<body>.*?)(?=^---\n\n### |\Z)",
        re.MULTILINE | re.DOTALL,
    )
    match = pattern.search(changelog)
    if not match:
        raise SystemExit(f"no CHANGELOG.md entry found for {version}")

    body = match.group("body").strip()
    return f"See SECURITY.md for verification guidance.\n\n{body}\n"


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--version", required=True)
    args = parser.parse_args()
    print(release_notes(args.version))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
