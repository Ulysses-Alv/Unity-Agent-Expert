# Light a VFX Graph asset with 2D lights in URP

Create a Visual Effect Graph asset and then light it with a 2D light by using [Shader Graphs](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest).

[Visual Effect Graph assets](https://docs.unity3d.com/Packages/com.unity.visualeffectgraph@latest?subfolder=/manual/VisualEffectGraphAsset.html) are compatible with the 2D Renderer by using [Shader Graphs](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest). Follow the steps below to first [create a Visual Effect Graph asset](#create-a-visual-effect-graph-asset) and then [light it with a 2D light](#light-a-visual-effect-with-2d-lights).

## Prerequisites

Refer to the Visual Effect Graph’s [requirements and compatibility](https://docs.unity3d.com/Packages/com.unity.visualeffectgraph@latest?subfolder=/manual/System-Requirements.html) for the required versions of packages for your Project.

## Create a Visual Effect Graph asset

To create a Visual Effect Graph asset (VFX Asset):

1. Create a new VFX asset by selecting **Assets > Create > Visual Effects > Visual Effect Graph**.
2. In the **Create new VFX Asset** window, select **Simple Loop**. Unity creates the VFX Asset in the **Asset** folder in the **Project** window.
3. Double-click the asset to open the VFX Graph.
4. Replace the **Output **Particle** Unlit** node with an **Output Particle **Shader** Graph **Quad**** node.
5. In the **Output Particle Shader Graph Quad** node, select the **Shader Graph** picker (⊙).
6. In the **Select Shader Graph Vfx Asset** window, select the eye icon to show hidden packages.
7. Select **VFXSpriteLit** so you can light the visual effect.

## Light a Visual Effect with 2D lights

To light a Visual Effect:

1. To create a Visual Effect GameObject, select **GameObject** > **Visual Effects** > **Visual Effect**.
2. In the **Visual Effect** properties, locate **Asset Template** and select the asset picker (circle). In the **Select VisualEffectAsset** window, select the VFX asset created earlier.
3. To light the Visual Effect, add [2D light](Lights-2D-intro.md) to the **scene**.

![An example of a VFX Graph asset with 2D lights.](../../uploads/urp/2D/visual-effect-3.png)

An example of a VFX Graph asset with 2D lights.

## Additional resources

* [Using a Shader Graph in a visual effect](https://docs.unity3d.com/Packages/com.unity.visualeffectgraph@latest?subfolder=/manual/sg-working-with.html#using-a-shader-graph-in-a-visual-effect)