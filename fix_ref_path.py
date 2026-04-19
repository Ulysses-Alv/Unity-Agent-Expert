#!/usr/bin/env python3
"""
Update reference docs to point to actual ManualMD location.
Also ensure paths are correctly formatted as relative from ManualMD.
"""

import re
from pathlib import Path

SKILLS_DIR = Path("D:/Projects/Skill creation/skills")
MANUALMD_BASE = Path("D:/Projects/Skill creation/ManualMD/en/Manual")

def update_file(filepath: Path) -> tuple[str, bool, str]:
    """Update a reference doc to point to ManualMD."""
    try:
        content = filepath.read_text(encoding="utf-8")
        original = content

        # Update the base path at the top of the file
        old_path = r"C:\\Program Files\\Unity\\Hub\\Editor\\6000\.3\.9f1\\Editor\\Data\\Documentation\\en\\Manual\\"
        new_path = "D:/Projects/Skill creation/ManualMD/en/Manual/"
        content = re.sub(old_path, new_path, content)
        content = re.sub(r"Manual\\en\\Manual\\", "ManualMD/en/Manual/", content)

        # Fix: Some entries just have filename.md, need to ensure they exist
        # The convention seems to be: relative path from ManualMD/en/Manual/
        # e.g., "ui-systems/ui-toolkit-live-reload.md" exists at ManualMD/en/Manual/ui-systems/

        # For entries that are just filenames (no dir prefix), check if they exist
        # If they don't exist but do exist in a subdir, the path is correct as-is

        if content != original:
            filepath.write_text(content, encoding="utf-8")
            return str(filepath), True, "updated"
        return str(filepath), True, "no_change"
    except Exception as e:
        return str(filepath), False, str(e)

def verify_paths():
    """Verify that reference paths actually exist in ManualMD."""
    skill_refs = [
        ("unity-uitk-csharp", "UIE-get-started-with-runtime-ui.md", "UIE-get-started-with-runtime-ui.md"),
        ("unity-uitk-csharp", "ui-systems/ui-toolkit-live-reload.md", "ui-systems/ui-toolkit-live-reload.md"),
        ("unity-uxml", "UIE-uxml.html", "UIE-uxml.md"),  # should already be .md
    ]

    print("Verifying sample paths exist in ManualMD:")
    for skill, ref_path, expected in skill_refs:
        # Check in subdirs of ManualMD/en/Manual/
        found = False
        for md_file in MANUALMD_BASE.rglob(expected):
            found = True
            print(f"  [OK] {skill}: {expected} exists at {md_file.relative_to(MANUALMD_BASE)}")
            break
        if not found:
            print(f"  [FAIL] {skill}: {expected} NOT FOUND")

def main():
    ref_docs = list(SKILLS_DIR.rglob("references/docs.md"))
    print(f"Found {len(ref_docs)} reference doc files")

    updated = 0
    for f in ref_docs:
        path, ok, status = update_file(f)
        if status == "updated":
            updated += 1
            print(f"Updated: {f.parent.parent.name}/references/docs.md")

    print(f"\nTotal updated: {updated}")

    verify_paths()

    return 0

if __name__ == "__main__":
    exit(main())