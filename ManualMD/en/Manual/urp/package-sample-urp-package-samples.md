# Package samples reference for URP

The package samples that URP provides are:

* [**URP Package Samples**](#urp-package-samples): A set of example **shaders**, C# **scripts**, and other assets to help you build or enhance your application.
* [**URP RenderGraph Samples**](#rendergraph-samples): A set of example custom render passes that use the [render graph system](render-graph.md).

For information about installing package samples, refer to [Import a package sample in URP](package-samples.md).

## URP Package Samples

URP Package Samples is a [package sample](package-samples.md) for the Universal **Render Pipeline** (URP). It contains example shaders, C# scripts, and other assets you can build upon, use to learn how to use a feature, or use directly in your application. For information on how to import URP Package Samples into your project, refer to [Importing package samples](package-samples.md#importing-package-samples).

Each example uses its own [URP Asset](universalrp-asset.md) so, if you want to build an example **scene**, [add the example’s URP Asset to your Graphics settings](InstallURPIntoAProject.md#set-urp-active). If you don’t do this, Unity might strip shaders or render passes that the example uses.

### Camera Stacking

The `URP Package Samples/CameraStacking` folder contains examples for [Camera Stacking](camera-stacking.md). The following table describes each **Camera** Stacking example in this folder.

| **Example** | **Description** |
| --- | --- |
| **Mixed field of view** | The example in `CameraStacking/MixedFOV` demonstrates how to use Camera Stacking in a first-person application to prevent the character’s equipped items from clipping into the environment. This setup also makes it possible to have different fields of view for the environment camera and the equipped items camera. |
| **Split screen** | The example in `CameraStacking/SplitScreenPPUI` demonstrates how to create a split-screen camera setup where each screen has its own Camera Stack. It also demonstrates how to apply **post-processing** on world-space and screen-space camera UI. |
| **3D **skybox**** | The example in `CameraStacking/3D Skybox` uses Camera Stacking to transform a miniature environment into a skybox. One overlay camera renders a miniature city and another renders miniature planets. The overlay cameras render to **pixels** that the main camera did not draw to. With some additional scripted translation, this makes the miniature environment appear full size in the background of the main camera’s view. |

### Decals

The `URP Package Samples/Decals` folder contains examples for [decals](renderer-feature-decal.md). The following table describes each decal example in this folder.

| **Example** | **Description** |
| --- | --- |
| **Blob shadows** | The example in `Decals/BlobShadow` uses the [Decal Projector component](renderer-feature-decal.md) to cast a shadow under a character. This method of shadow rendering is less resource-intensive than shadow maps and is suitable for use on low-end devices. |
| **Paint splat** | The example in `Decals/PaintSplat` uses a WorldSpaceUV Sub Graph and the [Simple Noise](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Simple-Noise-Node.html) Shader Graph node to create procedural decals. The noise in each paint splat uses the world position of the Decal Projector component. |
| **Proxy lighting** | The example in `Decals/ProxyLighting` builds on the **Blob shadows** example and uses Decal Projectors to add proxy spotlights. These decals modify the emission of surfaces inside the projector’s volume.  **Note**: To demonstrate the extent of its lighting simulation, this example disables normal real-time lighting. |

### Lens Flares

The `URP Package Samples/LensFlares` folder contains **lens flare** examples. The following table describes each lens flare example in this folder.

| **Example** | **Description** |
| --- | --- |
| **Sun flare** | The `LensFlares/SunFlare` example demonstrates how to use the [Lens Flare component](shared/lens-flare/lens-flare-component.md) to add a lens flare effect to the main directional light in the scene. |
| **Lens flare showroom** | The `LensFlares/LensFlareShowroom` example helps you to author lens flares. To use it: 1. In the Hierarchy window, select the **Lens Flare** GameObject. 2. In the Lens Flare component, assign a [LensFlareDataSRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@12.0/api/UnityEngine.Rendering.LensFlareDataSRP.html) asset to the **Lens Flare Data** property. 3. Change the Lens Flare component and data properties and view the lens flare in the Game View. **Note**: If the text box is in the way, disable the Canvas in the scene. |

### Lighting

The `URP Package Samples/Lighting` folder contains examples for [lighting](lighting-landing.md). The following table describes each lighting example in this folder.

| **Example** | **Description** |
| --- | --- |
| **Reflection probes** | The example in `Lighting/Reflection Probes` uses [reflection probes](lighting/reflection-probes.md) to create reflection maps for a reflective sphere **GameObject**. This sample shows how the **Probe Blending** and **Box Projection** settings can change the reflection within a scene that uses reflection probes. |

### Renderer Features

The `URP Package Samples/RendererFeatures` folder contains examples for [Renderer Features](urp-renderer-feature.md). The following table describes each Renderer Feature example in this folder.

| **Example** | **Description** |
| --- | --- |
| ****Ambient occlusion**** | The example in `RendererFeatures/AmbientOcclusion` uses a Renderer Feature to add [screen space ambient occlusion (SSAO)](post-processing-ssao.md) to URP. For an example of how to set up this effect, refer to the `SSAO_Renderer` asset. |
| **Glitch effect** | The example in `RendererFeatures/GlitchEffect` uses the [Render Objects](renderer-features/renderer-feature-render-objects.md) Render Feature and the [Scene Color](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Scene-Color-Node.html) Shader Graph node to draw some GameObjects with a glitchy effect. For an example of how to set up this effect, refer to the `Glitch_Renderer` asset. |
| **Keep frame** | The example in `RendererFeatures/KeepFrame` uses a custom Renderer Feature to preserve frame color between frames. The example uses this to create a swirl effect from a simple particle system. **Note**: The effect is only visible in Play Mode. |
| **Occlusion effect** | The example in `RendererFeatures/OcclusionEffect` uses the Render Objects Renderer Feature to draw occluded geometry. The example achieves this effect without any code and sets everything up in the `OcclusionEffect_Renderer` asset. |
| **Trail effect** | The example in `RendererFeatures/TrailEffect` uses the Renderer Feature from the **Keep frame** example on an additional camera to create a trail map. To do this, the additional camera draws depth to a RenderTexture. The `Sand_Graph` shader samples the map and displaces vertices on the ground. |

### Shaders

The `URP Package Samples/Shaders` folder contains examples for shaders. The following table describes each shader example in this folder.

| **Example** | **Description** |
| --- | --- |
| **Lit** | The example in `Shaders/Lit` demonstrates how different properties of the [Lit shader](lit-shader.md) affect the surface of some geometry. You can use the materials and textures as guidelines on how to set up materials in URP. |

## RenderGraph Samples

After you install the **Render Graph Samples**, the `URPRenderGraphSamples` folder contains examples of [Scriptable Renderer Features](renderer-features/scriptable-renderer-features/inject-a-pass-using-a-scriptable-renderer-feature.md) that use the render graph system.

To use an example, follow these steps:

1. In the **Project** window, select the URP Renderer your project uses.
2. In the ****Inspector**** window, select **Add Renderer Feature**.
3. Select the example you want to use.

For more information, refer to [Add a Renderer Feature to a URP Renderer](urp-renderer-feature.md).

To get more information about each sample, use the [Render Graph Viewer](render-graph-viewer-reference.md) to visualize the sequence of render passes. Some examples are for API demonstrative purposes and don’t generate a visible output in the Game or **Scene view**, but you can use the [Frame Debugger](../FrameDebugger.md) to see their output.

| **Example** | **Description** |
| --- | --- |
| ****Blit**** | Copies the active color texture to a new texture. |
| **BlitWithFrameData** | Demonstrates how Blit operations can be handled using frameData with multiple ScriptableRenderPasses. |
| **BlitWithMaterial** | Blits the active CameraColor to a new texture. Shows how to perform a blit with material and use ResourceData to avoid another blit back to the active color target. This example is for API demonstrative purposes. |
| **Compute** | Shows how a compute shader can be used together with RenderGraph. |
| **Culling** | Renders the scene geometry associated with a specific layer using the culling results, for API demonstrative purposes. |
| **FramebufferFetch** | Copies the target of the previous pass to a new texture using a custom material and framebuffer fetch. This example is for API demonstrative purposes. |
| **GBufferVisualization** | Uses the gBuffer components in a RenderPass when they are not global. |
| **GlobalGBuffers** | Sets the gBuffer components as globals (it renders nothing itself). Adding this feature to the scriptable renderer allows subsequent passes to access the gBuffers as globals. |
| **MRT** | Demonstrates how to use Multiple Render Targets (MRT) in RenderGraph with URP. Useful when more than 4 channels of data (a single RGBA texture) need to be written by a pass. |
| **OutputTexture** | Uses RenderGraph to output a specific texture in URP, shows how to attach a texture by name to a material, and demonstrates how two render passes can be merged if executed in the correct order. |
| **RendererList** | Clears the current active color texture, then renders the scene geometry associated with a **layer mask**. |
| **TextureReferenceWithFrameData** | Creates a texture reference ContextItem in frameData to hold a reference used by future passes. This avoids additional blit operations copying back and forth to the camera’s color attachment. |
| **UnsafePass** | Copies the active color texture to a new texture and then downsamples the source texture twice. This example is for API demonstrative purposes. |

## Additional resources

* [Introduction to the Universal Render Pipeline](urp-introduction.md)