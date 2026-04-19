#!/usr/bin/env python3
"""Quality check for ManualMD dataset."""

from pathlib import Path
from collections import Counter

MANUALMD = Path("ManualMD")

def check_file(filepath: Path) -> dict:
    """Check a single file for quality issues."""
    try:
        content = filepath.read_text(encoding="utf-8")
        lines = content.split('\n')

        issues = []

        # Check for empty or near-empty files
        if len(content.strip()) < 50:
            issues.append("VERY_SHORT")

        # Check for files with only nav items (should have been cleaned)
        if content.strip() and all(line.strip().startswith('*') or line.strip() == '' for line in lines if line.strip()):
            issues.append("ONLY_NAV_ITEMS")

        # Check for too many newlines (poor formatting)
        if '\n\n\n\n' in content:
            issues.append("EXCESSIVE_NEWLINES")

        # Check for .html references that weren't fixed
        if '.html' in content:
            issues.append("UNFIXED_HTML_LINKS")

        return {
            "path": str(filepath.relative_to(MANUALMD)),
            "lines": len(lines),
            "chars": len(content),
            "issues": issues
        }
    except Exception as e:
        return {
            "path": str(filepath),
            "error": str(e)
        }

def main():
    md_files = list(MANUALMD.rglob("*.md"))
    total = len(md_files)
    print(f"Checking {total} Markdown files...")

    results = []
    for f in md_files:
        results.append(check_file(f))

    # Aggregate stats
    issue_counts = Counter()
    short_files = []
    issue_files = []

    for r in results:
        if "error" in r:
            print(f"ERROR reading {r['path']}: {r['error']}")
            continue

        for issue in r.get("issues", []):
            issue_counts[issue] += 1
            if issue != "EXCESSIVE_NEWLINES":  # Too common to list individually
                issue_files.append((r["path"], issue))

        if r["lines"] < 5:
            short_files.append((r["path"], r["lines"], r["chars"]))

    print("\n=== QUALITY REPORT ===")
    print(f"Total files: {total}")

    print(f"\n--- Issue Summary ---")
    for issue, count in sorted(issue_counts.items(), key=lambda x: -x[1]):
        pct = (count / total) * 100
        print(f"  {issue}: {count} ({pct:.1f}%)")

    print(f"\n--- Short Files (<5 lines) ---")
    if short_files:
        for path, lines, chars in sorted(short_files)[:20]:
            print(f"  [{lines} lines, {chars} chars] {path}")
        if len(short_files) > 20:
            print(f"  ... and {len(short_files) - 20} more")
    else:
        print("  None!")

    print(f"\n--- Files with unfixed issues ---")
    if issue_files:
        for path, issue in issue_files[:10]:
            print(f"  {issue}: {path}")
        if len(issue_files) > 10:
            print(f"  ... and {len(issue_files) - 10} more")
    else:
        print("  None!")

    # Sample content from different sections
    print("\n--- Sample Files ---")
    samples = [
        "en/Manual/2d-physics/2d-physics.md",
        "en/Manual/sprite/sprite-landing.md",
        "en/Manual/urp/introduction-landing.md",
        "en/Manual/ui-systems/introduction-ui-toolkit.md",
    ]
    for sample in samples:
        fp = MANUALMD / sample
        if fp.exists():
            content = fp.read_text()
            print(f"\n{sample} ({len(content)} chars):")
            print(f"  First 200 chars: {content[:200].replace(chr(10), ' ')}...")
        else:
            print(f"\n{sample}: FILE NOT FOUND")

    return 0

if __name__ == "__main__":
    exit(main())