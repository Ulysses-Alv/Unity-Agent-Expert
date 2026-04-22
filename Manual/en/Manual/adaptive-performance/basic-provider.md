# Basic provider

The Basic provider is a built-in generic provider that works with performance metrics collected using the `FrameTimingManager` API.

To install the Basic provider, enable **Basic Provider** when you [set up Adaptive Performance](initial-setup.md).

## Platform support

The following platforms support the Basic provider:

* Android
* iOS
* Linux
* macOS
* tvOS
* visionOS
* UWP

## Limitations

The Basic provider doesn’t support the following APIs and features:

* The thermal hint API or related thermal settings (for example, [`IThermalStatus`](../../ScriptReference/AdaptivePerformance.IThermalStatus.md))
* Graphics APIs that lack GPU frame timing support

## Additional resources

* [FrameTimingManager](../../ScriptReference/FrameTimingManager.md)
* [Android provider package](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance.google.android@latest?subfolder=/manual/index.html)
* [Introduction to providers](providers-introduction.md)
* [Provider settings reference](provider-settings-reference.md)
* [Create a custom provider](create-custom-provider.md)