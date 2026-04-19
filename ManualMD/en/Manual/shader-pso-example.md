# PSO tracing and warming example

Get started with PSO tracing and warm up using the Universal **Render Pipeline** (URP) sample.

**Important**: `GraphicsStateCollection` is an experimental API and might be changed or removed in the future.

The Universal 3D sample uses the `GraphicsStateCollection` API. You can use this sample to help you get started with PSO tracing and warming methods. For more information about creating a project from the Universal 3D sample, refer to [Create a new project that uses URP](urp/creating-a-new-project-with-urp.md).

## Sample script reference

The following table outlines the key PSO tracing and warming **scripts** used in this sample and their function:

| **Script** | **Description** |
| --- | --- |
| `Assets/SharedAssets/Scripts/Runtime/GraphicsStateCollectionManager.cs` | Performs PSO tracing and warmup at runtime. Holds a list of recorded collections. |
| `Assets/SharedAssets/Scripts/Editor/GraphicsStateCollectionStripper.cs` | Ensures the build only includes the relevant collections for the target platform. The sample project targets multiple platforms and quality levels, so it contains multiple collections under the `Assets/SharedAssets/GraphicsStateCollections` folder. The filtering script removes unneeded collections for the target platform to reduce build size. |
| `Assets/SharedAssets/Scripts/Editor/GraphicsStateCollectionCombiner.cs` | Utility script that combines multiple selected collection files into one. Unity uses the first collection you select as the result collection. **Note**: This script combines collections that match the platform, graphics API, and quality level. |

### Trace and warm up with the URP sample project

To trace and warm up PSOs with the sample project:

1. Open the main terminal **scene** from `Assets/Scenes/Terminal/TerminalScene.unity`.
2. Add an empty **GameObject** to the scene.
3. Open your empty GameObject in the ****Inspector****, select **Add Component** and add the **Graphics State Collection Manager** component.
4. Set the **Mode** property of the **Graphics State Collection Manager** to **Warm Up**.
5. Open the **More** (⋮) menu next to the **Graphics State Collection Manager**. Select **Update collection list** to populate the collection list with the collections from the `GraphicsStateCollections` folder.
6. Save the scene and project.
7. Build and run the Player, targeting either Windows and the DirectX 12 or Vulkan graphics API, or OSX and the Metal graphics API.

When you load the scene in the Player, Unity will now perform PSO warmup automatically.

## Additional resources

* [Introduction to PSO tracing and warming](shader-pso-introduction.md)