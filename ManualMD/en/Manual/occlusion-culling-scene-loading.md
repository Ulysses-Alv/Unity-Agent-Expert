# Set up multiple scenes for occlusion culling

At runtime, Unity loads only one **occlusion culling** data asset at a time, no matter how many Scenes are open. You must therefore prepare your occlusion culling data differently depending on whether you plan to load one **Scene** at a time, or multiple Scenes at a time.

## Loading one Scene at a time

If you load one Scene at a time with [LoadSceneMode.Single](../ScriptReference/SceneManagement.LoadSceneMode.Single.md), you must bake the occlusion culling data for each Scene separately, like this:

1. In the Unity Editor, open a single Scene that you want to bake occlusion culling data for.
2. Bake the data for the Scene. Unity generates the data and saves it as *Assets/[Scene name]/OcclusionCullingData.asset*. Unity adds a reference to this asset to the open Scene.
3. Save the Scene.
4. Repeat steps 1, 2, and 3 for each Scene.

At runtime, load your Scenes like this:

1. Load a Scene as the default Scene of your Project, or using LoadSceneMode.Single. Unity unloads the active Scene if there is one, along with its occlusion data asset if it has one. Unity then loads your Scene along with its occlusion data asset.

## Loading more than one Scene at a time

If you load multiple Scenes at a time with [LoadSceneMode.Additive](../ScriptReference/SceneManagement.LoadSceneMode.Additive.md), you must bake the data for those Scenes together, like this:

1. In the Unity Editor, open the first Scene of the group that Unity will load at runtime. This becomes the active Scene.
2. Additively open the other Scenes in the group, so that all of them are open in the Unity Editor at the same time.
3. Bake the data for all Scenes. Unity generates the data for all of the open Scenes, and saves it as *Assets/[active Scene name]/OcclusionCullingData.asset*. Unity adds a reference to this asset to all of the open Scenes.
4. Save the Scenes.

At runtime, load your Scenes like this:

1. Load the first Scene of the group as the default Scene of your Project, or using `LoadSceneMode.Single`. Unity unloads the active Scene if there is one, along with its occlusion data asset if it has one. Unity then loads your Scene along with the shared occlusion data asset.
2. Load other Scenes as required with `LoadSceneMode.Additive`. If Unity finds that the the occlusion data of an additively loaded Scene is the same as that of the active Scene, occlusion culling works as intended.

Note that the shared occlusion data asset has a larger file size. There is no additional runtime CPU impact to using a larger occlusion data asset.

For further information on working with multiple Scenes in the Unity Editor, see the documentation on [multi-Scene editing](MultiSceneEditing.md).