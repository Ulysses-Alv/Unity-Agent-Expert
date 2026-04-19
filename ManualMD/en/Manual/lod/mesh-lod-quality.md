# Mesh LOD runtime quality

Learn about settings that let you configure how Unity selects **Mesh** **LODs** at runtime.

For meshes that have LODs, Unity selects a LOD index automatically based on the size of the object on the screen.

The following settings let you adjust the LOD selection and make Unity select more or less detailed LODs.

## Project-wide quality setting

The following settings affect how Unity selects LODs on a project level:

* **Quality** > **Level of Detail** > **Mesh LOD Threshold**

Unity uses the **Mesh LOD Threshold** parameter when evaluating the size of the mesh on the screen and selecting a specific LOD index. Increasing the **Mesh LOD Threshold** setting makes Unity favor less detailed LODs in the evaluation process.

To choose a value appropriate for your application:

1. Place a variety of assets from your project in a **scene**.
2. Change the **Mesh LOD Threshold** value and inspect the assets while moving the **camera** around the scene.
3. Repeat step 2 to find a value that provides a good balance between performance, visual quality of assets, and smoothness of LOD transitions.

This setting can have different values for different quality levels.

## Configure per-object LOD settings

To adjust how Unity selects LODs for a particular object:

1. Select a **GameObject** with the ****Mesh Renderer**** or **Skinned Mesh Renderer** component.
2. In the Inspector window, go to **Mesh Renderer** > **Mesh LOD**.
3. Use one of the following properties to adjust the behavior:

   * [**LOD Override**](../class-MeshRenderer.md#mesh-lod)
   * [**LOD Selection Bias**](../class-MeshRenderer.md#mesh-lod)

For information on the settings in the **Mesh LOD** section, refer to [Mesh LOD reference](../class-MeshRenderer.md#mesh-lod).

## Additional resources

* [Introduction to level of detail](../LevelOfDetail.md)