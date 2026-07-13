#!/usr/bin/env python3
"""Check local links in Markdown and site HTML files."""

from __future__ import annotations

import html.parser
import pathlib
import re
import sys
from urllib.parse import unquote, urlparse

ROOT = pathlib.Path(__file__).resolve().parents[1]
MARKDOWN_FILES = [ROOT / "README.md", ROOT / "SECURITY.md", *sorted((ROOT / "docs").glob("*.md"))]
HTML_FILES = sorted((ROOT / "site").glob("*.html"))

MD_LINK_RE = re.compile(r"(?<!!)\[[^\]]+\]\(([^)]+)\)")


class HTMLLinkParser(html.parser.HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.links: list[str] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        for name, value in attrs:
            if name in {"href", "src"} and value:
                self.links.append(value)


def is_external(target: str) -> bool:
    parsed = urlparse(target)
    return parsed.scheme in {"http", "https", "mailto"}


def normalize_local_target(source: pathlib.Path, target: str) -> pathlib.Path | None:
    target = target.strip()
    if not target or target.startswith("#") or is_external(target):
        return None

    parsed = urlparse(target)
    if parsed.scheme:
        return None

    path = unquote(parsed.path)
    if not path:
        return None
    return (source.parent / path).resolve()


def markdown_links(path: pathlib.Path) -> list[str]:
    return MD_LINK_RE.findall(path.read_text())


def html_links(path: pathlib.Path) -> list[str]:
    parser = HTMLLinkParser()
    parser.feed(path.read_text())
    return parser.links


def main() -> int:
    failures: list[str] = []

    for path in [*MARKDOWN_FILES, *HTML_FILES]:
        links = markdown_links(path) if path.suffix == ".md" else html_links(path)
        for link in links:
            target = normalize_local_target(path, link)
            if target is None:
                continue
            if not target.exists():
                failures.append(f"{path.relative_to(ROOT)}: broken local link: {link}")

    if failures:
        for failure in failures:
            print(failure, file=sys.stderr)
        return 1

    print("Local links OK")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
