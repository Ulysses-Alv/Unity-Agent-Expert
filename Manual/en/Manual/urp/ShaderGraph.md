# Create 2D lit shader with Shader Graph

Create a **shader** graph that enables 2D sprites to react to lighting, and provides support for lighting features like normal mapping and **sprite** masking. Without this shader, light sources do not affect sprites.

To create and set up a 2D sprite lit shader graph, follow these steps:

1. In the **Assets** menu, go to **Create** > **Shader Graph** > **URP** > **Sprite Lit Shader Graph**.
2. Open the newly created shader graph asset.
3. Add three **Sample Texture 2D** nodes to the shader graph.
4. Set the **Type** of one of the new nodes to **Normal**.
5. Connect the outputs of the nodes to the following **Fragment** context input slots:

   * First **Default** type node: Connect **RGBA** to **Base Color**, then connect **A** to **Alpha**.
   * Second **Default** type node: Connect **RGBA** to ****Sprite Mask****.
   * **Normal** type node: Connect **RGBA** to **Normal (Tangent Space)**.
6. Create three **Texture 2D** properties by selecting the **+** on the [Blackboard](http://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Blackboard.html), then choosing **Texture 2D**:

   * **Main Texture** `MainTex`: Defines the base color and transparency (alpha) of the sprite.
   * **Mask Texture** `MaskTex`: Determines areas of the sprite that should be visible or hidden using a custom mask.
   * ****Normal Map**** `NormalMap`: Adds surface details by simulating bumps and grooves, enhancing lighting effects on the sprite.
7. Drag each property into the shader graph workspace and connect them to the **Texture** input slots in the corresponding **Sample Texture 2D** nodes.
8. Drag the **NormalMap** property into the shader graph workspace and attach it to the **Texture** input slot of the **Normal** type **Sample Texture 2D** node.
9. Select **Save Asset** to save the shader.

   ![A shader graph that enables 2D sprites to react to lighting. Three textures are connected to three Sample Texture 2D nodes. The third node has a type of Normal. The nodes are connected to the inputs of the Fragment context.](../../uploads/urp/2D/ShaderGraph.png)

   A shader graph that enables 2D sprites to react to lighting. Three textures are connected to three Sample Texture 2D nodes. The third node has a type of Normal. The nodes are connected to the inputs of the Fragment context.

You can now apply the shader graph to a material and use it on sprites in a **scene**, allowing sprites to interact with 2D lights.

## Additional resources

* [Create and assign a material](../create-material.md)
* [Unity Shader Graph documentation](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html)
* [Blackboard in Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Blackboard.html)