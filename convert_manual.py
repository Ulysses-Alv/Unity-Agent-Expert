#!/usr/bin/env python3
"""
Unity 6000 Manual HTML to Markdown Converter
Converts all HTML files in Manual/ to clean Markdown in ManualMD/
"""

import os
import sys
from multiprocessing import Pool, cpu_count
from pathlib import Path

# Install dependencies if needed
try:
    from bs4 import BeautifulSoup
    from markdownify import markdownify as md
except ImportError:
    import subprocess
    subprocess.check_call([sys.executable, "-m", "pip", "install", "beautifulsoup4", "markdownify"])
    from bs4 import BeautifulSoup
    from markdownify import markdownify as md

INPUT_DIR = Path("Manual")
OUTPUT_DIR = Path("ManualMD")

def clean_html(html: str) -> str:
    """Remove script, style, nav, footer, header tags and extract main content."""
    soup = BeautifulSoup(html, "html.parser")

    # Remove noise tags
    for tag in soup(["script", "style", "nav", "footer", "header", "noscript"]):
        tag.decompose()

    # Find main content - look for div.section inside div.content-wrap
    # The actual content is inside <div class="content"> inside content-wrap, NOT the sidebar
    main = (
        soup.find("div", {"class": "section"}) or  # main article content
        soup.find("main") or
        soup.find("article") or
        # Find content inside content-wrap specifically
        soup.find("div", {"id": "content-wrap"}) or
        soup.body
    )

    if main and hasattr(main, 'find_all'):
        # Remove specific noise elements
        for noise in main.find_all(["div", "span"], class_=["version-number", "lang-switcher", "breadcrumbs",
                                                             "nextprev", "sidebar", "header-wrapper",
                                                             "footer-wrapper", "toolbar", "mb20"]):
            noise.decompose()

        # Remove feedback and empty elements
        for fb in main.find_all(["div", "span"], id=lambda x: x and (x.startswith("_") or "feedback" in x.lower())):
            fb.decompose()

        # Remove tooltips (tooltiptext divs)
        for tooltip in main.find_all(["div", "span"], class_="tooltiptext"):
            tooltip.decompose()

        # Remove empty divs that just have IDs
        for empty in main.find_all("div", id=lambda x: x and x.startswith("_")):
            empty.decompose()

    return str(main) if main else str(soup.body)


def convert_file(input_path: str) -> tuple[str, bool, str]:
    """
    Convert a single HTML file to Markdown.
    Returns: (output_path, success, error_message)
    """
    input_p = Path(input_path)

    # Compute output path
    try:
        relative = input_p.relative_to(INPUT_DIR)
    except ValueError:
        # File is outside INPUT_DIR structure, use filename only
        relative = Path(input_p.name)

    output_path = (OUTPUT_DIR / relative).with_suffix(".md")
    output_path.parent.mkdir(parents=True, exist_ok=True)

    try:
        with open(input_p, "r", encoding="utf-8", errors="ignore") as f:
            html = f.read()

        cleaned = clean_html(html)
        markdown = md(cleaned, heading_style="ATX")

        # Clean up excessive newlines
        while "\n\n\n" in markdown:
            markdown = markdown.replace("\n\n\n", "\n\n")

        with open(output_path, "w", encoding="utf-8") as f:
            f.write(markdown)

        return str(output_path), True, ""
    except Exception as e:
        return str(output_path), False, str(e)


def process_wrapper(args):
    """Wrapper for multiprocessing."""
    return convert_file(args)


def main():
    # Find all HTML files
    html_files = []
    for root, _, files in os.walk(INPUT_DIR):
        for file in files:
            if file.endswith(".html"):
                html_files.append(os.path.join(root, file))

    total = len(html_files)
    print(f"Found {total} HTML files to convert")
    print(f"Using {min(cpu_count(), 4)} worker processes")

    # Process files in parallel
    results = []
    with Pool(min(cpu_count(), 4)) as pool:
        results = pool.map(process_wrapper, html_files)

    # Report results
    success = sum(1 for _, ok, _ in results if ok)
    failed = [(path, err) for path, ok, err in results if not ok]

    print(f"\nConversion complete:")
    print(f"  [OK] {success} files converted")
    print(f"  [FAIL] {len(failed)} files failed")

    if failed:
        print("\nFailed files:")
        for path, err in failed[:20]:
            print(f"  - {path}: {err}")
        if len(failed) > 20:
            print(f"  ... and {len(failed) - 20} more")

    # Verify output
    md_count = sum(1 for _, _, _, in [(os.walk(OUTPUT_DIR))] for _ in _)
    print(f"\nOutput directory: {OUTPUT_DIR}")
    print(f"Files created: {success}")

    return 0 if not failed else 1


if __name__ == "__main__":
    sys.exit(main())