# Deferred and Deferred+ rendering paths in URP

![Scene rendered with the Deferred Rendering Path](../../../uploads/urp/rendering-deferred/deferred-intro-image.png)

Scene rendered with the Deferred Rendering Path

Resources for using the Deferred and Deferred+ rendering paths. The Deferred **rendering path** has no limit on the number of lights that can affect an opaque **GameObject**, however the Deferred+ rendering path uses [Forward+](forward-rendering-paths.md) instead of Forward for the transparent pass and for the forward only opaque pass, and has the same light limit as Forward+.

| **Page** | **Description** |
| --- | --- |
| [Introduction to the Deferred rendering path](../rendering/deferred-rendering-path-introduction.md) | Learn about how the Deferred rendering path works, and its limitations. |
| [Render passes in the Deferred and Deferred+ rendering path](../rendering/render-passes-deferred.md) | Learn about the sequence of render pass events in the Deferred rendering path. |
| [G-buffer layout in the Deferred and Deferred+ rendering path](../rendering/g-buffer-layout.md) | Understand how Unity stores material attributes in the geometry buffer (G-buffer) in the Deferred rendering path. |
| [Enable accurate G-buffer normals in the Deferred and Deferred+ rendering path in URP](../rendering/accurate-g-buffer-normals.md) | Configure how Unity encodes normals when it stores them in the G-buffer. |
| [Make a shader compatible with the Deferred or Deferred+ rendering paths in URP](../rendering/make-shader-compatible-with-deferred.md) | Use the `LightMode` tag in a **shader** to make the shader compatible with the Deferred rendering path. |