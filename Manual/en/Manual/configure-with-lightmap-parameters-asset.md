# Save and load lighting settings with Lightmap Parameters Assets

A **Lightmap** Parameters Asset contains a set of values for the parameters that control Unity’s lighting features. These Assets allow you to define and save different sets of values for lighting, for use in different situations.

Lightmap Parameters Assets allow you to quickly create presets optimized for different types of **GameObjects**, or for different platforms and different Scene types (for example, indoor or outdoor Scenes).

## Creating a Lightmap Parameters Asset

To create a new Lightmap Parameters Asset, right-click in the Project window and go to **Create** > **Lightmap Parameters**. Unity stores this in your Project folder.

## Assigning Lightmap Parameters Assets

### Scenes

To assign a Lightmap Parameters Asset to the whole Scene:

1. Open the [Lighting window](lighting-window.md) (**Window** > **Rendering** > **Lighting**)
2. Click the **Scene** tab
3. Navigate to the **Lightingmapping Settings**.
4. Use the **Lightmap Parameters** drop-down to assign a default Lightmap Parameters Asset. This drop-down lists all available Lightmap Parameters Assets.

### GameObjects

To assign a Lightmap Parameters Asset to a single GameObject, ensure the GameObject has a **Mesh** Renderer or **Terrain** component attached.

To assign a Lightmap Parameters Asset to a **Mesh Renderer** component:

1. In the Inspector, go to **Mesh Renderer** > **Lighting**
2. Enable **Contribute Global Illumination**
3. In the mesh Renderer component, go to **Lightmapping** > **Lightmap Parameters**.
4. Select an option from the menu. Select **Scene Default Parameter** to use the same Lightmap Parameters Asset that’s assigned to the whole Scene.

To assign a Lightmap Parameters Asset to a Terrain component:

1. In the Inspector, go to **Terrain** > **Terrain Settings** > **Lighting**
2. Enable **Contribute Global Illumination**
3. In **Terrain Settings**, go to **Lightmapping** > **Lightmap Parameters**.
4. Select an option from the menu. Select **Scene Default Parameter** to use the same Lightmap Parameters Asset that’s assigned to the whole Scene.