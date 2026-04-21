# Migrate static font assets to Advanced Text Generator

Advanced Text Generator doesn’t support static font assets. If your project uses static font assets, you must migrate them after you [enable Advanced Text Generator](enable-and-use-atg.md) so that your UI text elements can render correctly.

Unity projects commonly use [Static font assets](../UIE-font-asset.md#atlas-population-modes) for optimization. Choose your migration approach based on your original optimization goals for the static assets:

* [Reduce memory footprint](#goal-reduce-memory-footprint) through character set limitation.
* [Optimize startup and first-frame performance](#goal-optimize-startup-behavior) through a font atlas pre-populated with commonly used characters.
* [Remove static font assets](#remove-static-font-assets) without font-asset-specific optimization preservation.

## Reduce memory footprint

If you created the static font asset mainly to ship a smaller character set, subset the font file and then [create a dynamic font asset](../UIE-get-started-with-text.md#create-a-dynamic-font-asset) from the subset font.

This reduces the font file size and can further lower memory usage because it eliminates the need for a pre-populated atlas texture. For reference, an atlas texture in Alpha 8 is approximately:

* 1024 × 1024 → ~1 MB
* 2048 × 2048 (common for CJK) → ~4 MB

To migrate, do the following:

1. Identify the characters or **scripts** required by your UI.
2. [Subset the font file with an external tooling](../UIE-font-subsetting.md).
3. Import the subset font into Unity.
4. [Create a dynamic font asset](../UIE-get-started-with-text.md#create-a-dynamic-font-asset) from the subset font.
5. Replace all references to the static font asset (USS/UXML/C#).

## Optimize startup or first-frame behavior

If you used a static font asset mainly to pre-populate an atlas (to reduce runtime glyph generation during startup), the equivalent workflow in Advanced Text Generator is to use a dynamic font asset and pre-populate its atlas with a chosen character set.

**Note:** With SDFAA atlas rendering mode, many projects no longer need to pre-populate an atlas for acceptable first-frame performance. However, if you still require control or warm-up of glyph generation, Advanced Text Generator supports the workflow through a dynamic font asset.

To migrate, do the following:

1. [Create or select a dynamic font asset](../UIE-get-started-with-text.md#create-a-dynamic-font-asset)
2. In the **Font Asset Creator** window, load the dynamic font asset.
3. Populate the atlas with the desired character set.
4. Replace all references to the static font asset (USS/UXML/C#).

## Remove static font assets without preserving font-asset-specific optimizations

If you don’t need memory reduction nor startup optimization, you can replace static font references directly with the Unity font asset imported from the typeface file, such as `.ttf` or `.otf`.

This requires no font-asset creation step; however, you lose font-asset-specific customization options.

To migrate, do the following:

1. Locate the corresponding Unity font asset (the imported typeface) in the project.
2. Replace references to the static font asset with the Unity font asset (USS/UXML/C#).

Your UI text elements can now render correctly, and you can take advantage of Advanced Text Generator’s enhanced text rendering features going forward.

## Additional resources

* [Font assets](../UIE-font-assets)