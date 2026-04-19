# Upgrade to Unity 6.1

This page lists changes in Unity 6.1 that can affect existing projects when you upgrade them from Unity 6.0 to Unity 6.1.

Review changes for Unity 6.1 in these areas:

* [Editor and workflow](#editor)
* [Graphics](#graphics)
* [Physics](#physics)
* [Platforms](#platforms)

## Editor and workflow

This section outlines recent updates to the Editor and its general workflows that can affect your upgrade experience.

### Window menu changes

Package Manager, **Asset Store**, Services, My Assets, and **Version Control** now [appear in new submenus](WhatsNewUnity61.md#editor-and-workflow). This change breaks any custom keyboard shortcuts to open these windows.

After upgrading, recreate any shortcuts to these windows.

## Graphics

This section outlines recent updates to Unity’s graphics systems that can affect your upgrade experience.

### Render pipelines

For upgrade guides for specific **render pipelines** in Unity, refer to one of the following pages:

* [Upgrade to URP 17.1 (Unity 6.1)](urp/upgrade-guide-unity-6-1.md)
* [Upgrade to HDRP 17.1 (Unity 6.1)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.1/manual/upgrade-guides.html)

To upgrade other packages, refer to [the documentation for the packages you’re using](upm-docs.md).

### The `_CLUSTER_LIGHT_LOOP` shader keyword replaces `_FORWARD_PLUS`

Unity 6.1 replaces the `_FORWARD_PLUS` **shader** keyword with the `_CLUSTER_LIGHT_LOOP` keyword. If your custom shaders use the `_FORWARD_PLUS` keyword, replace it with `_CLUSTER_LIGHT_LOOP`.

For more information about using the keywords, refer to [Shader keywords and macros reference in URP](urp/urp-shaders/shader-keywords-macros.md).

### Support for the PVRTC format is deprecated

Support for PVRTC **compression** is deprecated and will be fully removed in the future. It’s recommended that you transition to other compression formats, such as ASTC or ETC, to avoid issues in future upgrades.

## Physics

This section outlines recent updates to Physics-specific tools and settings that can affect your upgrade experience.

### Rigidbody.SetDensity is deprecated

The `SetDensity` method was deprecated for [`Rigidbody`](../ScriptReference/Rigidbody.md). Use [`Rigidbody.mass`](../ScriptReference/Rigidbody-mass.md) instead.

The `SetDensity` method was effectively used as a mass multiplier for the body by the physics integration instead of a per-Collider value used during mass distribution to determine each **Collider**’s mass.

## Platforms

This section outlines recent updates to platform-specific tools and settings that can affect your upgrade experience.

### Android

#### The default Android tools versions have changed

Unity has updated the default versions of the following tools used by Android. The default versions of SDK build tools, SDK platform tools, and JDK remain unchanged. The updated versions are as follows:

| **Tool** | **Version** |
| --- | --- |
| **Gradle** | 8.11 |
| Android Gradle Plugin | 8.7.2 |
| SDK Command-line tools | 16.0 |
| NDK | r27c |

If your project uses any [custom Gradle templates](gradle-templates.md), consider recreating those templates to avoid any build issues with the updated Android Gradle Plugin version. For more information, refer to [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md).

### Windows

#### New default Auto Graphics API

The default Auto Graphics API in Unity 6.1 is DirectX12. Projects upgrading from older versions of Unity maintain the default setting from their version.

To change to the default for 6.1, in ****Player Settings****, enable **Auto Graphics API**.