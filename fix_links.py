#!/usr/bin/env python3
"""
Post-process Markdown files to fix internal links.
Changes .html extensions to .md for internal references.
"""

import os
import re
from pathlib import Path
from multiprocessing import Pool

MANUALMD_DIR = Path("ManualMD")

def fix_links_in_file(filepath: Path) -> tuple[str, bool, str]:
    """Fix .html links to .md in a single file."""
    try:
        content = filepath.read_text(encoding="utf-8")

        # Replace internal links: change .html to .md
        # Pattern matches:](path/file.html) or ](path/file.html#anchor)
        # But NOT external links (http, https, mailto, etc.)

        original = content

        # Fix markdown links: [text](file.html) -> [text](file.md)
        # But preserve external URLs (http, https, mailto, ftp)
        # Pattern: ](non-http-path) - we extract and fix the .html inside
        content = re.sub(
            r'(\]\()((?!http|https|mailto|ftp)[^)]+)(\))',
            lambda m: m.group(1) + m.group(2).replace('.html', '.md') + m.group(3),
            content
        )

        # Also fix reference-style links if any
        content = re.sub(
            r'(\]\: )([^hhtps])([^\s]*\.html)',
            lambda m: m.group(1) + m.group(2) + m.group(3).replace('.html', '.md'),
            content
        )

        if content != original:
            filepath.write_text(content, encoding="utf-8")
            return str(filepath), True, "fixed"
        return str(filepath), True, "no change"

    except Exception as e:
        return str(filepath), False, str(e)


def process_wrapper(filepath):
    return fix_links_in_file(Path(filepath))


def main():
    md_files = list(MANUALMD_DIR.rglob("*.md"))
    total = len(md_files)
    print(f"Processing {total} Markdown files...")

    fixed_count = 0
    with Pool(4) as pool:
        results = pool.map(process_wrapper, [str(f) for f in md_files])

    for path, ok, status in results:
        if status == "fixed":
            fixed_count += 1

    print(f"Done. Fixed links in {fixed_count} files.")

    # Sample: show a before/after from a physics file
    sample_file = MANUALMD_DIR / "en" / "Manual" / "2d-physics" / "joints" / "hinge-joint-2d-fundamentals.md"
    if sample_file.exists():
        content = sample_file.read_text()
        # Count .html references remaining
        html_refs = re.findall(r'\.html', content)
        print(f"\nSample check ({sample_file.name}): {len(html_refs)} .html references remaining")
        # Show lines with .html
        for i, line in enumerate(content.split('\n'), 1):
            if '.html' in line:
                print(f"  Line {i}: {line.strip()[:80]}...")

    return 0


if __name__ == "__main__":
    exit(main())