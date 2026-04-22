# Strip shader variants in URP

By default, the Universal **Render Pipeline** (URP) compiles both **shader** variants where a feature is enabled, and shader variants where a feature is disabled. This can increase the number of shader variants in your build exponentially, depending on how many features you use.

To remove a feature and its shader variants, do either of the following:

* Disable the feature in the Editor settings.
* Disable shader keywords by using the Shader Build Settings in [Graphics settings](urp-global-settings.md).

If you need a feature for your project, you can instead make sure you enable it across your project in the Editor settings, so Unity can strip the variants where the feature is disabled.

**Important**: When you reduce shader variants, Unity might strip shader variants your built project needs. To check how many shader variants your project has, enable strict shader variant matching, otherwise Unity might replace missing shader variants with others. For more information, refer to [Check how many shader variants you have](../shader-how-many-variants.md#highlight-missing-shaders).

## Strip unused shader variants

To make sure Unity strips unused shader variants from your build, follow these steps:

1. From the main menu, go to **Edit** > **Project Settings**.
2. Select the **Graphics** tab.
3. In the **Additional Shader Stripping Settings** section, enable **Strip Unused Variants**.

If you disable the **Strip Unused Variants setting**, URP can’t strip variants where the feature is disabled. This might increase the number of variants.

## Disable a feature in settings

To remove the shader variants for a feature, disable the feature in all the URP Assets or Universal Renderer assets in your project, or elsewhere in the Editor settings.

For a list of settings, refer to [Settings and keywords reference for stripping keywords in URP](shader-stripping-features-and-keywords.md).

If a setting is in the URP asset, make sure you disable the feature in all the URP Assets included in your build. Unity includes the following URP Assets in your build:

* The URP Asset you set as the default render pipeline asset in [Graphics settings](https://docs.unity3d.com/Manual/class-GraphicsSettings.html).
* Any URP Asset you set as a **Render Pipeline Asset** in a [Quality settings level](https://docs.unity3d.com/Manual/class-QualitySettings.html) you enable for the current build target.

**Note:** Avoid including URP Assets in your build that use different [rendering paths](rendering-paths-introduction-urp.md). This can cause Unity to create two sets of variants for each keyword.

## Disable shader keywords

For a list of keywords in prebuilt shaders in URP, refer to [Settings and keywords reference for stripping keywords in URP](shader-stripping-features-and-keywords.md).

Follow these steps:

1. From the main menu, go to **Edit** > **Project Settings**.
2. Select the **Graphics** tab.
3. In the **Shader Build Settings** section, select the **Add** (**+**) button.
4. Add the keyword set, for example `_ _ADDITIONAL_LIGHTS_VERTEX _ADDITIONAL_LIGHTS`.
5. Use the foldout (triangle) to expand the keyword set, then disable the individual keywords.
6. Select **Apply**. Unity recompiles the shaders.

To keep a feature but still reduce the number of shader variants, refer to [Enable dynamic branching in prebuilt shaders](shader-stripping-fog.md).

## Strip post-processing shader variants

To strip shader variants for [Volume Overrides](VolumeOverrides.md) you don’t use, follow these steps:

1. From the main menu, go to **Edit** > **Project Settings**.
2. Select the **Graphics** tab.
3. Enable **Strip Unused Post Processing Variants**. Unity now keeps only the shader variants for Volume Overrides you use in your project.
4. Enable **Strip Screen Coord Override Variants**, unless you override screen coordinates to support **post-processing** on large numbers of multiple displays, also known as cluster displays.

To disable the shader keywords for a post-processing effect instead, refer to [Settings and keywords reference for stripping keywords in URP](shader-stripping-features-and-keywords.md).

Unity checks for Volume Overrides in all scenes, so you can’t strip variants by removing a **scene** from your build but keeping it in your project. You must remove the scene from the project.

## Strip XR and VR shader variants

If you don’t use [XR](https://docs.unity3d.com/Manual/XR.html) or [VR](https://docs.unity3d.com/Manual/VROverview.html), [disable the XR and VR modules](https://docs.unity3d.com/Documentation/Manual/upm-ui.html). This allows URP to strip XR and VR-related shader variants from its standard shaders.

## Open the URP stripping code

To examine the code that strips shaders in URP, check the `Editor/ShaderPreprocessor.cs` file. The file uses the [IPreprocessShaders](https://docs.unity3d.com/ScriptReference/Build.IPreprocessShaders.html) API.

## Additional resources

* [Reduce shader variants](../shader-variant-stripping.md)