#!/usr/bin/env python3
"""Final quality check for ManualMD dataset - focused on internal links."""

from pathlib import Path

MANUALMD = Path("ManualMD")

def check_internal_html_refs(filepath: Path) -> list:
    """Find internal .html links that weren't fixed."""
    try:
        content = filepath.read_text(encoding="utf-8")
        lines = content.split('\n')
        issues = []

        for i, line in enumerate(lines, 1):
            # Look for markdown links to .html that are NOT http/https
            import re
            matches = re.findall(r'\[([^\]]+)\]\((?!http|https|mailto|ftp)([^)]+\.html)', line)
            for text, url in matches:
                issues.append((i, text, url))

        return issues
    except Exception as e:
        return [(-1, "ERROR", str(e))]

def main():
    md_files = list(MANUALMD.rglob("*.md"))
    total = len(md_files)

    print("=== FINAL QUALITY CHECK ===")
    print(f"Total MD files: {total}")

    # Check for unfixed internal .html links
    files_with_issues = []
    total_internal_html = 0

    for f in md_files:
        issues = check_internal_html_refs(f)
        if issues:
            files_with_issues.append((str(f.relative_to(MANUALMD)), issues))
            total_internal_html += len(issues)

    print(f"\n--- Internal .html links (should be 0) ---")
    print(f"Unfixed internal links: {total_internal_html}")
    print(f"Files affected: {len(files_with_issues)}")

    if files_with_issues:
        print("\nFiles with unfixed internal .html links:")
        for path, issues in files_with_issues[:10]:
            print(f"  {path}")
            for line_num, text, url in issues[:3]:
                print(f"    Line {line_num}: [{text}]({url})")

    # Verify some sample files from different sections
    print("\n--- Sample content verification ---")
    samples = {
        "2d-physics": "en/Manual/2d-physics/2d-physics.md",
        "sprite": "en/Manual/sprite/sprite-landing.md",
        "urp": "en/Manual/urp/introduction-landing.md",
        "ui-systems": "en/Manual/ui-systems/introduction-ui-toolkit.md",
        "graphics": "en/Manual/Graphics.md",
    }

    for section, path in samples.items():
        fp = MANUALMD / path
        if fp.exists():
            content = fp.read_text()
            # Check for internal .md links
            import re
            internal_links = re.findall(r'\[([^\]]+)\]\((?!http|https|mailto|ftp)([^\)]+\.md)', content)
            # Check for .html refs (should only be external)
            html_refs = re.findall(r'\.html', content)

            print(f"\n{section}: {path}")
            print(f"  Size: {len(content)} chars, {len(content.split(chr(10)))} lines")
            print(f"  Internal .md links: {len(internal_links)}")
            print(f"  .html refs (external only): {len(html_refs)}")
            if internal_links:
                print(f"  Sample links: {[l[0] for l in internal_links[:3]]}...")
        else:
            print(f"\n{section}: NOT FOUND ({path})")

    print("\n=== RESULT ===")
    if total_internal_html == 0:
        print("PASS: No unfixed internal .html links")
    else:
        print(f"FAIL: {total_internal_html} unfixed internal .html links in {len(files_with_issues)} files")

    return 0 if total_internal_html == 0 else 1

if __name__ == "__main__":
    exit(main())