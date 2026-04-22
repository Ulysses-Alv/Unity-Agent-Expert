# Render 3D GameObjects in 2D scenes

To render a 3D **GameObject** in a 2D Universal **Render Pipeline** (URP) **scene**, set its **Mesh** Renderer component to be compatible with 2D scenes. This makes the 3D GameObject use 2D features like [2D lighting in URP](2d-index.md).

Use either of the following approaches:

* Create a 3D GameObject in a 2D URP project.
* Create a custom **shader** in Shader Graph.

## Create a 3D GameObject in a 2D URP project

When you create a 3D GameObject in a 2D URP project, Unity automatically assigns the **Mesh2D-Lit-Default** material. This material makes the GameObject compatible with 2D Lights.

## Add the default material to a 3D GameObject

To add the default **Mesh2D-Lit-Default** material to an existing 3D GameObject, follow these steps:

1. Select the 3D GameObject.
2. In the ****Inspector**** window, go to the ****Mesh Renderer**** or **Skinned Mesh Renderer** component.
3. Open the **Materials** dropdown, then select the Material picker (**⊙**).
4. Select **Mesh2D-Lit-Default**.

## Create a compatible custom shader in Shader Graph

To create a compatible shader in Shader Graph for a Mesh Renderer or Skinned Mesh Renderer, follow these steps:

1. From the main menu, select **Assets** > **Create** > **Shader Graph** > **URP**, then either **Sprite Lit Shader Graph**, **Sprite Unlit Shader Graph**, or **Sprite Custom Lit Shader Graph**.
2. Open the [Graph Inspector](https://docs.unity3d.com/Packages/com.unity.shadergraph@laetst?subfolder=/manual/Internal-Inspector.html), then in the **Graph Settings** enable **Sort 3D As 2D Compatible**.
3. Complete and save your shader. Refer to the [Shader Graph package documentation](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?) for more information.
4. In the **Inspector** window of a 3D GameObject, apply the new custom material to the Mesh Renderer component or Skinned Mesh Renderer component.

## Additional resources

* [Shader Graph package documentation](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest)
* [2D lighting in URP](2d-index.md)
* [2D Renderer asset component reference for URP](2DRendererData-overview.md)