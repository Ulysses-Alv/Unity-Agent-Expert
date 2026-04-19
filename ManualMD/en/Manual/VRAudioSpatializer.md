# Audio spatializers in XR

Audio spatializers use physical characteristics of a **scene**, such as the distance and angle between an [AudioListener](class-AudioListener.md) and an [AudioSource](class-AudioSource.md), to modify the properties of sound transmitted to the user. Spatialization can improve the perception that a sound originates from a specific location in a scene.

The Unity audio system supports spatialization through **plug-ins** built with the Unity [Audio Spatializer SDK](AudioSpatializerSDK.md). Unity doesn’t provide any built-in spatializer plug-ins itself, but several plug-in implementations are available in third-party 3D audio SDKs. These audio SDKs typically provide additional Unity components and tools for 3D audio. Refer to the following sections to understand the available third-party audio SDKs and how to enable them in Unity.

## Audio spatializer reference

The following is a non-comprehensive list of third-party audio SDKs that provide Unity audio spatialization plug-ins:

| Maker | Name | Platforms | Documentation |
| --- | --- | --- | --- |
| Microsoft | [Microsoft Spatializer](https://github.com/microsoft/spatialaudio-unity) | Windows, Android | [Spatial sound in Unity](https://learn.microsoft.com/en-us/windows/mixed-reality/develop/unity/spatial-sound-in-unity) (Microsoft learn) |
| Meta | [Meta XR Audio SDK](https://assetstore.unity.com/packages/tools/integration/meta-xr-audio-sdk-264557) | Windows, Android | [Meta XR Audio SDK Unity Plugin](https://developers.meta.com/horizon/documentation/unity/meta-xr-audio-sdk-unity-intro) (Meta developer documentation) |
| Qualcomm | [Qualcomm 3D Audio Plugin for Unity](https://assetstore.unity.com/packages/tools/audio/3d-audio-plugin-for-unity-148379) | Windows, Android | [3D Audio Plugin for Unity](https://www.qualcomm.com/developer/software/3d-audio-plugin-for-unity) (Qualcomm documentation) |
| Steam | [Steam Audio](https://valvesoftware.github.io/steam-audio/downloads.html) | Windows, macOS, Linux, Android | [Steam Audio Unity Integration](https://valvesoftware.github.io/steam-audio/doc/unity/index.html) (Valve GitHub) |
| Vive | [3DSP Audio SDK](https://developer.vive.com/resources/vive-sense/3dsp-audio-sdk/download/latest/) | Windows, Android | [VIVE 3DSP Audio SDK Developer Guide](https://hub.vive.com/storage/3dsp/) (Vive) |
| Google (now open source) | [Resonance Audio](https://resonance-audio.github.io/resonance-audio/) | Android, iOS, Web | [Developer Guide for Resonance Audio for Unity](https://resonance-audio.github.io/resonance-audio/develop/unity/developer-guide) (Resonance audio GitHub) |
| Apple | [PHASE](https://github.com/apple/unityplugins/tree/main/plug-ins/Apple.PHASE/) | iOS, macOS | [PHASE](https://developer.apple.com/documentation/phase) (Apple developer documentation) |

Some [XR provider plug-in](xr-support-packages.md) include spatializer plug-ins for an associated **XR** device. For example, the Windows **Mixed Reality** feature group for OpenXR includes the [MS HRTF Spatializer](https://github.com/microsoft/spatialaudio-unity) plug-in. Note that these provider plug-ins don’t include any additional components that might be available in the full SDK packages from their maker.

**Note:** Although many spatializer plug-ins were developed for use with **VR** devices, their use isn’t limited to VR games or applications.

## Enable an audio spatializer plug-in

After you have added the package containing an **audio spatializer** plug-in to your project, you can enable the plug-in in the project audio settings.

To enable a plug-in:

1. Open the Project Settings window (menu: **Edit > Project Settings**).
2. Select the **Audio** category.
3. Choose the plug-in from the **Spatializer Plugin** dropdown menu.

## Additional resources

* [Audio overview](https://docs.unity3d.com/Manual/AudioOverview.html)