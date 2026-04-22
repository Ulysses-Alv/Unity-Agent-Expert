# Optimize font files with font subsetting

Font subsetting is the process of reducing a font file’s size by including only the specific characters or glyphs needed. This was previously handled through [static font asset](UIE-get-started-with-text.md#create-a-static-font-asset). However, since font files are now required by the [Advanced Text Generator](UIE-advanced-text-generator.md), static font assets are no longer supported.

To subset fonts, use external tools, such as the command-line [FontTools](https://fonttools.readthedocs.io/en/latest/) library.

## Prerequisites

1. Install Python 3.x.
2. Install FontTools library via pip: `pip install fonttools`.

## Subset fonts with pyftsubset

pyftsubset is a command-line tool included in the [FontTools](https://github.com/fonttools/fonttools) library, used for font subsetting. It allows you to create a smaller font file by including only the specific characters or glyphs you need, which helps reduce file size and improve performance.

Assume you have an example font file named `example.ttf`. The following examples show how to use pyftsubset to subset the font file by Unicode ranges, text, or glyphs.

### Subset by Unicode range

To subset for Basic Latin and Latin–1 Supplement, use the following command:

`pyftsubset example.ttf --unicodes=U+0000-007F,U+00A0-00FF --output-file=subsetted.ttf`

### Subset by text

To subset font files by text, use the following command:

`pyftsubset example.ttf --text="Hello, World! test" --output-file=subsetted.ttf`

### Subset by glyph names

To subset by glyph names, use the following command:

`pyftsubset example.ttf --glyphs=A,B,C --output-file=subsetted.ttf`

### Retain OpenType layout features

By default, subsetting might remove some OpenType layout features, such as kerning or ligatures. You can use the `--layout-features` option to specify the features you want to keep. For example `--layout-features='kern,liga'`. To retain all features, use: `--layout-features=*`