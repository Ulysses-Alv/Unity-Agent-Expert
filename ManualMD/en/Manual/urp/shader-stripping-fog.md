# Enable dynamic branching in shaders in URP

To reduce the number of **shader** variants in the Universal **Render Pipeline** (URP), enable Unity using dynamic branching in shaders instead of keywords and shader variants.

For more information about keywords, refer to [Changing how shaders work using keywords](../SL-MultipleProgramVariants.md)

**Note**: If you target an older mobile device, dynamic branching can cause performance issues. Test on your target device.

Follow these steps:

1. From the main menu, go to **Edit** > **Project Settings**.
2. Select the **Graphics** tab.
3. In the **Shader Build Settings** section, select the **Add** (**+**) button.
4. Input the keyword set, including the single underscore (`_`) or multiple underscores (`__`) if they’re part of the set. For example `_ FOG_LINEAR FOG_EXP FOG_EXP2`.
5. Set **Type Override** to **dynamic\_branch**.
6. Select **Apply**. Unity recompiles the shaders.

## Compatible keyword sets in prebuilt shaders

If you use [prebuilt shaders in URP](shaders-in-universalrp-reference.md), you can only enable dynamic branching for the keyword sets in the following table.

| **Keyword set** | **Feature dynamic branching is enabled for** |
| --- | --- |
| `_ FOG_LINEAR FOG_EXP FOG_EXP2` | Fog |
| `_REFLECTION_PROBE_BLENDING` | **Reflection Probe** blending. For more information about Reflection Probes, refer to [Reflection Probes in URP](lighting/reflection-probes-introduction.md). |
| `_REFLECTION_PROBE_BOX_PROJECTION` | Reflection Probe box projection. For more information about box projection, refer to [Troubleshooting reflections](../AdvancedRefProbe.md). |

Other keyword sets in prebuilt shaders aren’t compatible with dynamic branching.

## Check compatibility of custom shaders

To enable dynamic branching for a keyword set in a [custom shader](../urp/writing-custom-shaders-urp.md), use HLSL `if` statements for branches, not preprocessor directives such as `#if`.

## Additional resources

* [Reduce shader variants](../shader-variant-stripping.md)
* [How Unity compiles branching shaders](../shader-conditionals-choose-a-type.md)