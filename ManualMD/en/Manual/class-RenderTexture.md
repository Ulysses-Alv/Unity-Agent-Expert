# Render Texture Inspector window reference

[Switch to Scripting](../ScriptReference/RenderTexture.md "Go to RenderTexture page in the Scripting Reference")

The **Render Texture** **inspector** is similar to the [Texture Inspector](class-TextureImporter.md).

The Render Texture inspector displays the current contents of Render Texture in real-time and can be an invaluable debugging tool for effects that use render textures.

| **Property** | **Description** |
| --- | --- |
| **Dimension** | Set the number of dimensions of the Render Texture. For more information, refer to [Dimension dropdown](#dimension-dropdown). |
| **Size** | Set the size of the Render Texture in **pixels**. |
| **Anti-Aliasing** | Set the number of **antialiasing** samples Unity applies. The options are **None**, **2 samples**, **4 samples**, or **8 samples**. If you select **None**, Unity doesn’t apply antialiasing. |
| **Enable Compatible Format** | If you enable this property, Unity converts the **Color Format** and **Depth Stencil Format** to compatible formats, if the platform doesn’t support the formats you select. |
| **Color Format** | Set the [GraphicsFormat](../ScriptReference/Experimental.Rendering.GraphicsFormat.md) of the color buffer of the Render Texture. If you select **None**, Unity doesn’t allocate a color buffer for this Render Texture. |
| **Depth Stencil Format** | Set the [GraphicsFormat](../ScriptReference/Experimental.Rendering.GraphicsFormat.md) of the depth/**stencil buffer** of the Render Texture. If you select **None**, Unity doesn’t allocate a depth/stencil buffer for this Render Texture. The stencil buffer is a general purpose buffer that allows you to store an additional unsigned 8-bit integer (0 to 255) for each pixel Unity draws to the screen. |
| **Mipmap** | Allocate a [mipmap](texture-mipmaps-introduction.md) for the Render Texture. |
| **Auto-generate** | Automatically generate a mipmap. If you disable this property, use the [`RenderTexture.GenerateMips`](../ScriptReference/RenderTexture.GenerateMips.md) API to generate a mipmap manually. Alternatively, choose which mipmap level to render into when you use [`Graphics.SetRenderTarget`](../ScriptReference/Graphics.SetRenderTarget.md) and [`Rendering.CommandBuffer.SetRenderTarget`](../ScriptReference/Rendering.CommandBuffer.SetRenderTarget.md). This property is visible only if you enable **Enable Mip Maps**. |
| **Dynamic Scaling** | Use [dynamic resolution](DynamicResolution-landing.md) to resize the Render Texture. If you disable this property, the Render Texture stays the same size regardless of the **camera**’s **Dynamic Resolution** setting. |
| **Random Write** | Allow **shaders** to write into arbitrary locations of the color buffer of the Render Texture. Refer to [`RenderTexture.enableRandomWrite`](../ScriptReference/RenderTexture-enableRandomWrite.md) for more information. |
| **Wrap Mode** | Control how Unity wraps the texture. For more information, refer to [Wrap Mode dropdown](#wrap-mode-dropdown). |
| **Filter Mode** | Control how the sampling of the texture uses nearby pixels. For more information, refer to [Filter Mode dropdown](#filter-mode-dropdown). |
| ****Aniso Level**** | Set the anisotropic filtering level of the texture. This increases texture quality when you view the texture at a steep angle. Use anisotropic filtering for floor, ground, and road textures. This property is visible only if you set **Depth Stencil Format** to **None**. |
| **Shadow Sampling Mode** | Select the sampling technique the GPU uses if the render texture is used as a [shadow map](shadow-mapping.md). This property is only visible if **Depth Stencil Format** is not set to **None**, and only has an effect if you set **Color Format** to **None**. |

## Dimension dropdown

| **Value** | **Description** |
| --- | --- |
| **2D** | Set the Render Texture as a 2D texture. |
| **2D Array** | Set the Render Texture as a 2D texture array. |
| **Cube** | Set the Render Texture as a **cubemap**. |
| **3D** | Set the Render Texture as 3D texture. |

## Wrap Mode dropdown

| **Value** | **Description** |
| --- | --- |
| **Repeat** | Tile the texture to create a repeating pattern. |
| **Clamp** | Stretch the edges of the texture. This is useful for preventing wrapping artifacts when you map an image onto a **GameObject**, and you don’t want the texture to tile. |
| **Mirror** | Tile the texture to create a repeating pattern that mirrors the texture at every integer boundary. |
| **Mirror Once** | Mirror the texture once, then fall back to clamping. |
| **Per-axis** | Set different **Wrap Mode** values for the u-axis and v-axis. For example, when you use latitude-longitude environment maps for **reflection probes**, you can use **Repeat** for the u-axis and **Clamp** for the v-axis. |

## Filter Mode dropdown

| **Value** | **Description** |
| --- | --- |
| **Point** | Use the nearest pixel. This makes the texture appear pixelated. |
| **Bilinear** | Use a weighted average of the four nearest texels. This makes the texture appear blurry when you magnify it. |
| **Trilinear** | Use a weighted average of the two nearest mipmap levels, which Unity filters using bilinear filtering. This creates a soft transition between mipmap levels, but the texture is slightly more blurry. |

RenderTexture