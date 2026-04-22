# Render a camera view to a Render Texture

To create a live **scene** **camera** in your game, follow these steps:

1. To create a Render Texture, go to **Assets** > **Create** > **Rendering** > **Render Texture**.
2. Select **GameObject** > **Camera** to create a second camera.
3. Assign the **Render Texture** to the **Target Texture** property in the second camera’s ****Inspector**** window.
4. Select **GameObject** > **3D Object** > **Quad** to create a quad.
5. Drag the Render Texture onto the **quad**. Unity automatically creates and assigns a material that uses the Render Texture.
6. Enter Play mode. The quad displays the camera view, which updates in real time.

![The result of the previous steps. The quad in the sky displays the view from the Secondary Camera.](../uploads/Main/RenderTextureLiveCam.jpg)  
The result of the previous steps. The quad in the sky displays the live view from the **Secondary Camera**.

## Additional resources

* [Render Texture Inspector window reference](class-RenderTexture.md)
* [Render a camera’s output to a Render Texture in URP](urp/rendering-to-a-render-texture.md)