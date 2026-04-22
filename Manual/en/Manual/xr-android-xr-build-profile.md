# Android XR build platform and build profile

Understand how to configure the Android **XR** build platform and build profile.

Unity supports development for headsets running Android XR. In Unity 6.1 and later, you can use the Android XR platform and build profile in the [Build Profiles](build-profiles.md) window.

The following sections outline Unity’s resources and packages to develop for Android XR devices.

## Prerequisites

* [Android Build Support](android-install-dependencies.md) module for the version of the Editor you’re using.

## Select the Android XR build platform

You can use this build platform to configure settings for Android XR devices.

To select the Android XR build platform:

1. Go to **File** > **Build Profiles** to open the [Build Profiles](build-profiles-reference.md) window.
2. In the **Platforms** list, select **Android XR**.
3. Click the **Switch Platform** button.

## Android XR build profile

By default, the Android XR platform shares **Player settings** with the [Android build platform](android-build-settings.md). You can use the Android XR build profile to override **project settings** specifically for Android XR.

### Add the Android XR build profile

To add an Android XR build profile:

1. Open the **Build Profiles** window (**File** > **Build Profiles**), and select **Add Build Profile** to open the [Platform Browser](platform-browser-reference.md) window.
2. Select **Android XR** from the list of platforms.
3. Select **Add Build Profile**, which creates an Android XR build profile and installs package dependencies.

### Customize build profile settings

In the **Android XR Build Profile** tab, you can configure the following settings derived from the Android XR build platform:

* **Platform Settings**: customizable build settings specific for the Android XR platform. These settings aren’t shared by the Android build platform. Refer to [Android build platform](android-build-settings.md) for details on these options.
* **Player Settings**, **Graphics Settings**, and **Quality Settings**: shared with the Android build platform by default. You can customize and override these settings for the Android XR platform. To override these settings, select the relevant override option from the **Player Settings Overrides** section to reveal the override options. Refer to [Overrride settings with build profiles](build-profiles-override-settings.md) to learn more.

### Android XR platform settings

* The Android XR build profile adds a new **Android XR (Build Profile)** [Quality level](class-QualitySettings.md) (**Project Settings** > **Quality Settings**). You can modify the quality settings.
* When you first switch to the Android XR build target or build profile, Unity [enables Android XR features](https://docs.unity3d.com/Packages/com.unity.xr.androidxr-openxr@1.0/manual/get-started/project-setup.html#enable-unity-openxr-android-xr) and [Foveated Rendering](https://docs.unity3d.com/Packages/com.unity.xr.androidxr-openxr@1.0/manual/features/foveated-rendering.html). These features remain enabled for the Android XR build target or build profile when you swap to a different build target.

## Additional resources

* [XR Plug-in Management settings](xr-plugin-management.md)
* [Optimize for untethered XR devices](xr-untethered-device-optimization.md)