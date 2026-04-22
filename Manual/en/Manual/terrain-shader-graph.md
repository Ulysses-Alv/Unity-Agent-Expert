# Render Terrain with Shader Graph

Terrain uses a default **shader** for its material, but you can create your own material with Shader Graph to fit specific goals and improve performance.

Possibilities include high quality tile repetition break-up solutions, detail mapping, smooth blending with the distant material, parallax mapping, auto materials, tri-planar projection, and more.

## Create and assign a Terrain material from Shader Graph

To create a new **Terrain** material with Shader Graph:

1. In the **Project** window, right-click and select **Create** > **Shader Graph** > **URP** or **HDRP** > **Terrain Lit Shader Graph**.

   Unity creates a Shader Graph asset. This asset is configured to use the default **Terrain Lit** material type, and the graph is empty, apart from the Master Stack.

   **Note**: If you want to start from the **Terrain Simple** template instead of an empty graph, refer to [Create a shader graph from a template](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?subfolder=/manual/Create-Shader-Graph.html).
2. Use the Shader Graph window to [build a graph according to your needs](#use-render-options-in-shader-graph-to-improve-terrain) with the [available nodes](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?subfolder=/manual/ui-nodes.html), and save the Shader Graph asset.
3. In your project, right click on the Shader Graph asset and select **Create** > **Material**.

   Unity creates a material based on your shader graph. From now, this material automatically updates if you modify and save your shader graph again.
4. Open your Terrain **GameObject** in the ****Inspector**** window.
5. Go to [**Terrain Settings** > **Basic Terrain**](terrain-OtherSettings.md).
6. Replace the default material with the material created from your shader graph.

## Use render options in Shader Graph to improve Terrain

Use Shader Graph to improve Terrain performance and visual fidelity.

### Improve performance

To improve the performance of Terrain rendering, you need to sample the textures less. The default Terrain shader does three samples per layer, and another sample of the control map. To reduce sampling, do one of the following:

* Pack the Terrain materials into fewer textures. For example, use tangent-space reconstruction to store normals in fewer channels.
* Pack the Terrain material into texture arrays, and sample only one entry for the array, based on the data in the control map.
* Don’t sample at all. Instead, construct the Terrain materials procedurally in the shader.

### Improve visual fidelity

To improve visual fidelity, do one of the following:

* Break up the Terrain grid with UV sampling.
* Remove Terrain stretching on steep surfaces with triplanar mapping.
* Add depth to the Terrain with parallax mapping.
* Use height based blending for Terrain materials.

**Note:** Many of these techniques increase the number of texture samples, which can reduce performance.

## Additional resources

To get started with Shader Graph Terrain rendering, refer to the [Shader Graph Terrain Samples](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?subfolder=/manual/Shader-Graph-Sample-Terrain.html) documentation.

For more information on specific pipeline support, refer to:

* URP: [Terrain Lit Shader Graph reference for URP](urp/prebuilt-shader-graphs-urp-terrain-lit.md)
* HDRP: [Terrain Master Stack reference](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/terrain-master-stack-reference.html)