# Example Custom Render Texture with Shader Graph | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Custom-Render-Texture-Example.html)

# Example Custom Render Texture with Shader Graph

This example demonstrates how to set up a Shader Graph for Custom Render Texture shaders to create a self-healing deformation effect that could be used for snow, sand, etc.

[](images/CustomRenderTexture-Animation.mp4)

For this effect, we need a Render Texture that contains the pixels we want to displace. This Render Texture is displayed at the top left corner of the preceding image. It was directly assigned to a camera that renders the objects in red above the UV plane.

1. Create a **Custom Render Texture** Shader Graph as follows:

   * Right-click in the Project window.
   * Select **Create** > **Shader Graph** > **Custom Render Texture** and name your shader.
2. Create a new **Custom Render Texture** as follows:

   * Select **Create** > **Rendering** > **Custom Render Texture**.
   * Configure the parameters to match the following settings:
     + **Dimension**: 2D
     + **Size**: 1024 x 1024
     + **Anti-aliasing**: None
     + **Enable Compatible Format**: selected
     + **Color Format**: R32\_SFLOAT
     + **Depth Stencil Format**: None
     + **Mipmap**: not selected
     + **Dynamic Scaling**: not selected
     + **Random Write**: not selected
     + **Wrap Mode**: Clamp
     + **Filter Mode**: Bilinear
     + **Aniso Level**: 0
     + **Material**: SnowHeightmapUpdate
       **Note**: The material in the **Material** field was created.
     + **Shader Pass**: <Unnamed Pass 0>
     + **Initialization Mode**: OnDemand
     + **Source**: Texture and Color
     + **Color**: select a color
     + **Texture**: None (Texture)
     + **Update Mode**: Realtime
     + **Period**: 0
     + **Double Buffered**: selected
     + **Wrap Update Zones**: not selected
     + **Update Zone Space**: Normalized
     + **Update Zones**: list empty

The following Shader Graph reads the output of the Camera Render Texture, as well as the Self texture, then adds them and lerps the result so that it tends towards 0 over time.

![The graph that is described in the preceding text.](images/CustomRenderTextureShaderGraph.png)

3. Finally you need to assign the **Custom Render Texture** to a material that can deform the geometry (tessellation or pixel displacement).

## New ShaderGraph Nodes for Custom Render Textures

Three new ShaderGraph nodes have been introduced to facilitate working with Custom Render Textures:

1. [Custom Render Texture Self](Custom-Render-Texture-Self-Node.html)

   * Outputs the double-buffered texture from the previous update.
   * It provides three outputs for different types of textures: 2D, 3D, and Cube textures.
2. [Custom Render Texture Size](Custom-Render-Texture-Size-Node.html)

   * Outputs the width, height, and volume depth of the current Custom Render Texture.
   * Use this node to determine the size of the texture being worked with.
3. [Custom Render Texture Slice](Slice-Index-Cubemap-Face-Node.html)

   * Outputs the current slice for 2D Texture Arrays or the Cubemap face index being updated.

## Known Limitations

When working with Custom Render Texture shaders, be aware of the following limitations:

* Rendering Order
  + The Custom Render Texture update loop runs before any rendering operations. This means:
    - You cannot access any render pipeline-specific information in the shader, such as depth, normal, or scene color.
    - Despite this, you can rely on the content of the Custom Render Texture being ready for use in the scene when the camera renders.

## Additional resources

* [Custom Render Textures](https://docs.unity3d.com/Manual/class-CustomRenderTexture.html)
* [Custom Render Texture Nodes](Custom-Render-Texture-Nodes.html)