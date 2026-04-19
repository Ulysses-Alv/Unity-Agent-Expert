# Accessibility for Unity applications

Create accessible Unity applications for mobile and desktop users.

Creating accessible applications makes your digital experiences available to a wider audience that includes people with disabilities. Use the [Accessibility module APIs](../ScriptReference/UnityEngine.AccessibilityModule.md) to add screen reader support and configure UI based on the system accessibility settings of the user’s device.

## Screen reader support

Most major mobile and desktop operating systems have built-in screen readers. Screen readers describe audibly what is displayed on the screen. They are designed to help users who are blind or have low vision navigate and interact with their devices without needing to see the screen.

The following table lists the built-in screen readers that the Accessibility module supports and their associated platforms:

| **Platform** | **Screen reader** |
| --- | --- |
| **Android** | [TalkBack](https://support.google.com/accessibility/android/topic/3529932?ref_topic=9078845) |
| **iOS** | [VoiceOver](https://support.apple.com/en-us/guide/iphone/iph3e2e415f/ios) |
| **Windows** | [Narrator](https://support.microsoft.com/en-us/windows/complete-guide-to-narrator-e4397a0d-ef4f-b386-d8ae-c172f109bdb1) |
| **macOS** | [VoiceOver](https://support.apple.com/en-us/guide/voiceover/welcome/mac) |

Use Unity’s [Assistive Support API](../ScriptReference/Accessibility.AssistiveSupport.md) to enable screen reader capabilities for your application. The `AssistiveSupport` class stores the active accessibility hierarchy that you create, allows your application to notify the screen reader of changes in your UI, and notifies your application of events based on user actions. Use the **Accessibility Hierarchy Viewer** (menu: **Window** > **Accessibility** > **Accessibility Hierarchy Viewer**) during Play mode to view your active accessibility hierarchy and its nodes in real-time.

The Assistive Support APIs are available for the following platforms:

| **Platform** | **Operating system version** |
| --- | --- |
| **Android** | 8.0 (API 26) and newer |
| **iOS** | 13 and newer |
| **Windows** | 10 version 21H1 (build 19043) and newer |
| **macOS** | Big Sur 11 and newer |

## System settings

Use Unity’s [Accessibility Settings API](../ScriptReference/Accessibility.AccessibilitySettings.md) to configure your UI to interact with native font scaling, bold text, and closed captions on mobile devices. This improves the readability and visibility of your application’s UI for all your users.

The Accessibility Settings APIs are available for the following platforms:

| **Platform** | **Operating system version** |
| --- | --- |
| **Android** | 7.1 (API 25) and newer |
| **iOS** | 13 and newer |

## Additional resources

* [Sample project using the accessibility APIs](https://github.com/Unity-Technologies/a11y-public-sample)
* [Accessibility module](../ScriptReference/UnityEngine.AccessibilityModule.md) API documentation
* [Mobile screen reader support in Unity](https://unity.com/blog/engine-platform/mobile-screen-reader-support-in-unity)
* [Unity Discussions: Accessibility](https://discussions.unity.com/tag/Accessibility-Features)
* [User Interfaces roadmap](https://unity.com/roadmap#unity-platform-ui)