# Render to a 2D texture array

To render to a 2D texture array, create a [render texture](render-texture-landing.md) and set the **Dimension** property to **2D Array**.

If you use the [`Graphics.SetRenderTarget`](../ScriptReference/Graphics.SetRenderTarget.md) API, set the `depthSlice` parameter to the slice you want to render to.

If the platform supports geometry shaders, use a geometry **shader** to render to individual slices, or set `depthSlice` to `-1` to render to all the slices.