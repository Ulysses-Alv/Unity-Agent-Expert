#!/usr/bin/env python3
"""
Fix reference docs: change .html file references to .md equivalents.
All skills' references/docs.md files point to old HTML manual - fix to .md
"""

import re
from pathlib import Path
from multiprocessing import Pool

SKILLS_DIR = Path("D:/Projects/Skill creation/skills")

def fix_file(filepath: Path) -> tuple[str, bool, str]:
    """Fix .html refs to .md in a single reference doc."""
    try:
        content = filepath.read_text(encoding="utf-8")
        original = content

        # Pattern 1: | topic | `filename.html` | → | topic | `filename.md` |
        # Pattern 2: | topic | `dir/filename.html` | → | topic | `dir/filename.md` |
        content = re.sub(r'`([^`]*\.html)`', lambda m: '`' + m.group(1).replace('.html', '.md') + '`', content)

        # Pattern 3: - filename.html → - filename.md (list items)
        content = re.sub(r'(?<!\w)([A-Za-z0-9_-]+\.html)', lambda m: m.group(0).replace('.html', '.md'), content)

        # Also fix any inline HTML references like (PhysicsOverview.html)
        content = re.sub(r'\(([A-Za-z0-9_/-]+\.html)\)', lambda m: '(' + m.group(1).replace('.html', '.md') + ')', content)

        if content != original:
            filepath.write_text(content, encoding="utf-8")
            return str(filepath), True, "fixed"
        return str(filepath), True, "no_change"
    except Exception as e:
        return str(filepath), False, str(e)

def main():
    # Find all references/docs.md files
    ref_docs = list(SKILLS_DIR.rglob("references/docs.md"))
    print(f"Found {len(ref_docs)} reference doc files")

    results = []
    for f in ref_docs:
        results.append(fix_file(f))

    fixed = sum(1 for _, ok, s in results if s == "fixed")
    print(f"\nFixed {fixed} files")

    # Show before/after for first file
    if ref_docs:
        first = ref_docs[0]
        content = first.read_text()
        # Count remaining .html
        remaining = re.findall(r'\.html', content)
        print(f"\nSample: {first.parent.parent.name}")
        print(f"  Remaining .html refs: {len(remaining)}")
        if remaining[:3]:
            print(f"  First few: {remaining[:3]}")

    return 0

if __name__ == "__main__":
    exit(main())