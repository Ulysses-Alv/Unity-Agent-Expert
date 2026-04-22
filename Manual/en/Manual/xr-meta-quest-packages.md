# Packages and templates for Meta Quest development

Learn about the packages and templates you can use to develop for Meta Quest with Unity.

Unity provides [packages](#unity-packages) for Meta Quest development via the Package Manager, and templates that are accessible from the Unity Hub to help you get started. You can also use partner packages from [Meta](#meta-packages) to develop for Meta Quest.

Refer to the following sections to understand the packages and templates you can use to develop for Meta Quest devices.

## Unity packages

You can install the following packages from the [Package Manager](upm-ui-install.md):

* The [Unity OpenXR plug-in](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/index.html) integrates core OpenXR features for all [XR runtimes](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/index.html#runtimes), including Meta Quest.
* The [Unity OpenXR: Meta](https://docs.unity3d.com/Packages/com.unity.xr.meta-openxr@2.2/manual/index.html) (OpenXR Meta) package integrates Meta-specific extensions to provide additional Meta Quest-specific features.

**Tip:** You can use Unity’s OpenXR **plug-in** and OpenXR Meta package together to develop a cross-platform application that has additional features tailored for Meta Quest devices.

To learn more about the packages Unity provides for **XR** development, refer to [XR packages](xr-support-packages.md).

## Meta packages

Meta also provides packages to develop for Meta Quest in Unity. Refer to Meta’s [Import XR packages](https://developers.meta.com/horizon/documentation/unity/unity-package-manager/#meta-xr-packages-overview) documentation to learn about the packages Meta provides.

**Note:** These packages and related documentation are maintained by Meta.

In Unity 6.3 and later, you can install Meta packages when you add the Meta Quest **build profile** using the **Platform Browser**, as described in [Choose packages from the platform browser](xr-meta-quest-build-profile.md#choose-packages). When you install these packages with the platform browser, you don’t need to install them from the **Asset Store**.

The following Meta packages are available in the platform browser:

* [Meta XR All-in-one SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-all-in-one-sdk-269657?srsltid=AfmBOopMC8pEIn7dKPSvxnFPwVnxBoRjMuEekzsNeytiOjNnlAQ3X0n_) (com.meta.xr.sdk.all)
* [Meta XR Core SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-core-sdk-269169) (com.meta.xr.sdk.core)
* [Meta XR Audio SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-audio-sdk-264557) (com.meta.xr.sdk.audio)
* [Meta XR Haptics SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-haptics-sdk-272446) (com.meta.xr.sdk.haptics)
* [Meta XR Interaction SDK Essentials](https://assetstore.unity.com/packages/tools/integration/meta-xr-interaction-sdk-essentials-264559) (com.meta.xr.sdk.interaction)
* [Meta XR Interaction SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-interaction-sdk-265014) (com.meta.xr.sdk.interaction.ovr)
* [Meta XR Platform SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-platform-sdk-262366) (com.meta.xr.sdk.platform)
* [Meta XR Voice SDK](https://assetstore.unity.com/packages/tools/integration/meta-voice-sdk-immersive-voice-commands-264555) (com.meta.xr.sdk.voice)
* [Meta XR Simulator](https://assetstore.unity.com/packages/tools/integration/meta-xr-simulator-266732) (com.meta.xr.sdk.simulator)
* [Meta XR Mixed Reality Utility Kit](https://assetstore.unity.com/packages/tools/integration/meta-mr-utility-kit-272450) (com.meta.xr.mrutilitykit)

## Templates

You can use the following templates as a starting point for your Meta Quest project:

* [Mixed Reality template](https://docs.unity3d.com/Packages/com.unity.template.mixed-reality@2.0)
* [VR Template](https://docs.unity3d.com/Packages/com.unity.template.vr@9.0)
* [VR Multiplayer template](https://docs.unity3d.com/Packages/com.unity.template.vr-multiplayer@2.0)
* [Mixed Reality multiplayer tabletop template](https://docs.unity3d.com/Packages/com.unity.template.mr-multiplayer@1.0/manual/index.html)

Refer to [Create an XR project](xr-create-projects.md) to understand how to use templates as a starting point for your project.

## Additional resources

* [XR Plug-in Management settings](xr-plugin-management.md)