# Custom Render Texture Inspector window reference

[Switch to Scripting](../ScriptReference/CustomRenderTexture.md "Go to CustomRenderTexture page in the Scripting Reference")

The **Custom Render Textures** **Inspector** window displays many of the same properties as the **Render Texture** Inspector, and some properties specific to **Custom Render Textures**.

### Render Texture:

| Property: | Function: |
| --- | --- |
| **Dimension** | Dimension of the Render Texture. |
| 2D | Makes the Render Texture two-dimensional. |
| Cube | Makes the Render Texture a cube map. |
| 3D | Makes the Render Texture three-dimensional. |
| **Size** | The size of the Render Texture in **pixels**. |
| **Color Format** | The format of the Render Texture. |
| **sRGB (Color Render Texture)** | Enable to allow the Render Texture to use sRGB read/write conversions (Read Only). |
| **Mipmap** | Enable this property to allow the Render Texture to use a mipmap. |
| **Auto-generate** | Enable this property to automatically generate the mipmap to use. |
| **Wrap Mode** | Defines how the Texture behaves when tiled. |
| Repeat | The Texture repeats itself as a tile. |
| Clamp | Unity stretches the edges of the Texture. |
| **Filter Mode** | Defines how Unity filters the Texture when it is stretched by 3D transformations. |
| Point | The Texture becomes blocky up close. |
| Bilinear | The Texture becomes blurry up close. |
| Trilinear | Like Bilinear, but the Texture also blurs between different mipmap levels. |
| **Aniso Level** | Increases Texture quality when viewing the texture at a steep angle. Good for floor and ground textures. |

### Custom Texture:

These properties are exclusive to Custom Render Textures. Custom Texture parameters are separated into three categories:

* **Material**: Defines what **shader** is used to update the texture.
* **Initialization**: Controls how the texture is initialized before the shader performs any updates.
* **Update**: Controls how the shader updates the texture.

| Property: | Function: |
| --- | --- |
| **Material** | The Material that Unity uses to update the Custom Render Texture. |
| Shader Pass | The Shader Pass that Unity uses to update the Custom Render Texture. The drop-down shows all passes available in your Material. |
| **Initialization Mode** | The rate at which Unity initializes the texture. |
| OnLoad | Unity initializes the texture once upon creation. |
| Realtime | Unity initializes the texture every frame. |
| OnDemand | Unity initializes the texture on demand from the script. |
| **Source** | How Unity texture initializes the texture. |
| Texture and Color | Unity uses a texture multiplied by a color to initialize the texture. |
| Initialization Color | Defines the color that Unity uses to initialize the Custom Render Texture. If you also provide an **Initialization Texture**, Unity uses the multiplication of the color and the texture to initialize the Custom Render Texture. |
| Initialization Texture | Defines the texture that Unity uses to initialize the Custom Render Texture. If you also provide an **Initialization Color**, Unity uses the multiplication of the color and the texture to initialize the Custom Render Texture. |
| Material | Unity uses a Material to initialize the texture. |
| Initialization Material | Defines the Material that Unity uses to initialize the Custom Render Texture. |
| **Update Mode** | The rate at which the shader updates the Custom Render Texture. |
| OnLoad | The shader updates the texture once upon creation. |
| Realtime | The shader updates the texture every frame. |
| OnDemand | The shader updates the texture on demand from script. |
| **Period** | The amount of time in seconds over which Unity updates a real-time texture. A value of 0.0 updates every frame. This property is only available when the **Update Mode** to property is set to Realtime. |
| **Double Buffered** | Double-buffers the texture. Each update swaps the two buffers. This allows you to read the result of the preceding update in the shader. |
| **Wrap Update Zones** | Enable to allow partial update zones to wrap around the border of the texture. |
| **Cubemap Faces** | (Cubemap only) Series of toggle allowing user to enable/disable update on each of the cube map faces. |
| **Update Zone Space** | The coordinate system where Unity defines the update zones. |
| Normalized | All coordinates and sizes are between 0 and 1 with the top-left corner starting at (0, 0). |
| Pixel | All coordinates and sizes are expressed in pixels limited by the width and height of the texture. Top-left corner starting at (0, 0). |
| **Update Zone List** | List of update zones for the texture. For more information, see [Update Zones](class-CustomRenderTexture-configure.md). |

## Exporting Custom Render Texture to a file

You can export Custom Render Textures to a PNG or EXR file (depending on the texture format) in the **Export** menu.

## Update Zones

| Property: | Function: |
| --- | --- |
| Center | The coordinates of the center of the update zone.. |
| Size | The size of the update zone. |
| Rotation | The orientation of the update zone in degrees (unavailable for 3D textures). |
| Shader Pass | Defines the Shader Pass to use for this update zone. If you set this property as default, this update zone uses the Shader Pass that you defined in the main part of the inspector. Otherwise it will use the Shader Pass you provide. |
| Swap (Double Buffer) | (Only for Double Buffered textures) When you enable this property, Unity swaps the buffers before processing this update zone. |