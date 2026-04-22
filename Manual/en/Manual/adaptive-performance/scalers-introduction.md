# Modifying asset quality with scalers

Understand how scalers dynamically adjust asset quality levels to optimize performance based on system conditions.

For practical examples of scaler use, refer to the [Adaptive Performance samples](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html).

## Scalers and the Indexer

Scalers are components that represent individual features, such as graphics or physics, and control the quality of these features in response to real-time system conditions.

The Indexer captures these system conditions and tracks thermal and performance state to produce a quantified quality index.

### How scalers work

Each scaler has default settings that specify the minimum scale, maximum scale, and maximum scale level of the associated feature. As the scale level increases, the associated feature’s quality (as measured by level of detail) decreases. The scale step is calculated as (highest scale - lowest scale) / maximum scale level. The default scale level always starts at zero.

Scalers respond to priorities that the Indexer supplies, using the following targets:

* Current bottleneck
* Lowest quality level
* Lowest visual impact

### Enable scalers

Scalers only work when the Indexer is active. When you enable scalers, they’re automatically added to your application.

To enable scalers:

1. Go to **Edit** > **Project Settings** > **Adaptive Performance** > **Provider** (for example, **Samsung (Android)**).
2. In the **Runtime Settings** section under **Indexer Settings**, enable **Active**.
3. Under **Scaler Profiles**, enable the scalers you want to use.

## Built-in and custom scalers

Adaptive Performance provides a set of [built-in scalers](scalers-reference.md) with default settings.

You can also [create custom scalers](create-custom-scalers.md) to control additional features in your application.

## Managing scalers with profiles

Once you have several scalers active, managing their settings individually can be complex. To simplify this, you can group scaler settings into scaler profiles.

A profile lets you create and change between different performance configurations (like high quality or power saving) with a single command.

To create and load scaler profiles at runtime, refer to [Use scaler profiles](scaler-profiles.md).

## Additional resources

* [Apply performance optimizations](performance-optimization-strategies.md)
* [Provider settings reference](provider-settings-reference.md)