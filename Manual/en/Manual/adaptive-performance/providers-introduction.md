# Introduction to providers

Use providers to supply Adaptive Performance with performance data from a device.

Adaptive Performance requires a data provider subsystem to function. You can install and enable a provider for each platform from the **Adaptive Performance** section of the **Project Settings** window.

## Adaptive Performance providers

The following providers are officially built and supported by Unity:

| **Provider** | **Description** |
| --- | --- |
| [**Basic provider**](basic-provider.md) | A built-in generic provider that works with performance metrics collected using the [`FrameTimingManager`](../../ScriptReference/FrameTimingManager.md) API. |
| [**Android provider**](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance.google.android@latest?subfolder=/manual/index.html) | A package that extends Adaptive Performance to Android devices. |

## Custom providers

You can also implement your own custom providers. To do so, refer to [Create a custom provider](create-custom-provider.md).

## Default provider

The default provider is the one Unity automatically selects when you enable Adaptive Performance. This selection is based on a priority system: Unity loads the installed provider with the lowest priority number. You can configure these priority values through [IAdaptivePerformanceLoaderMetadata](../../ScriptReference/AdaptivePerformance.Editor.Metadata.IAdaptivePerformanceLoaderMetadata.md).

## Additional resources

* [Provider settings reference](provider-settings-reference.md)
* [Provider lifecycle management](provider-lifecycle.md)