# Feature support by provider

Discover which Adaptive Performance features you can use on your target platforms.

Adaptive Performance relies on platform-specific [providers](providers.md) to gather data. The following tables outline the data each provider supplies to Adaptive Performance and the features each provider supports.

## Provider data

This table outlines what device data each provider supplies to Adaptive Performance.

| **Data** | **Android provider** | **Basic provider** | **Samsung provider (deprecated)** | **Device Simulator **plug-in**** |
| --- | --- | --- | --- | --- |
| **Temperature warning** | Yes | No | Yes | Yes |
| **Temperature level** | Yes | No | Yes | Yes |
| **Performance status** | Yes | Yes | Yes | Yes |
| **CPU time** | Yes | Yes | Yes | Yes |
| **GPU time** | Yes | Yes | Yes | Yes |
| **Game Mode** | Yes | No | No | Yes |
| **Game state** | Yes | No | No | Yes |

## Feature support

The following tables demonstrate the features available to each provider. Support for these features directly depends on the underlying data each provider supplies.

Adaptive Performance features offer a scripting API in C# as well as components with controls and settings in the Unity Editor.

### Thermal features

| **Feature** | **Implementation** | **Android provider** | **Basic provider** | **Samsung Provider (deprecated)** | **Device Simulator plug-in** |
| --- | --- | --- | --- | --- | --- |
| **Temperature trend** | API | Yes | No | Yes | Yes |
| **Temperature warning** | API | Yes | No | Yes | Yes |
| **Temperature level** | API | Yes | No | Yes | Yes |

### Performance features

| **Feature** | **Implementation** | **Android provider** | **Basic provider** | **Apple provider** | **Samsung provider (deprecated)** | **Device Simulator plug-in** |
| --- | --- | --- | --- | --- | --- | --- |
| **CPU and GPU level**  **Note**: Only the deprecated Samsung provider implements the CPU level and GPU level features. These features are available for testing and development with the Device Simulator but won’t work on other providers. | API | No | No | No | Yes | Yes |
| **Bottleneck detection** | API | Yes | Yes | Yes | Yes | Yes |
| **Auto Performance Control** | API, Component | Yes | No | No | Yes | Yes |
| **Scalers** | API, Component | Yes | Yes | Yes | Yes | Yes |
| **Variable refresh rate (VRR)** | API | No | No | No | Yes | Yes |

### General and platform features

| **Feature** | **Implementation** | **Android provider** | **Basic provider** | **Samsung Provider (deprecated)** | **Device Simulator plug-in** |
| --- | --- | --- | --- | --- | --- |
| **Game Mode** | API | Yes | No | No | Yes |
| **Game state** | API, Component | Yes | No | No | Yes |
| **Indexer** | API, Component | Yes | Yes | Yes | Yes |
| **Device simulator plug-in** | Component | Yes | No | Yes | Yes |

## Additional resources

* [Check for feature support](check-feature-support.md)
* [Scripting with Adaptive Performance](scripting-api.md)
* [Introduction to providers](providers-introduction.md)