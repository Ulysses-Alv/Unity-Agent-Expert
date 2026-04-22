# Configure shadow cascades

When you enable shadow cascades in your project, you can configure the number of cascades and the distances where each cascade ends.

In URP, configure shadow cascades using the [**Cascade Count**](urp/universalrp-asset.md#shadows) property, and then configure the cascade split distances. To visualize shadow cascades in URP, use the [**Shadow Cascades**](urp/shadow-cascades-visualize.md) view in the [**Rendering Debugger**](urp/features/rendering-debugger.md).

In the Built-In **Render Pipeline**, configure shadow cascades per quality level property in your project’s [Quality Settings](class-QualitySettings.md).

The number of shadow cascades affects performance. For more information, refer to [Performance impact of shadow cascades](shadow-cascades-performance.md).

In the High Definition Render Pipeline (HDRP), configure shadow cascades for each [Volume](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/understand-volumes.html).

## Additional resources

* [Performance impact of shadow cascades](shadow-cascades-performance.md)