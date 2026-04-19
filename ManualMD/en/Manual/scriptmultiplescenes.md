## Use scripts to edit multiple scenes

You can edit multiple **scenes** using **scripts** within the Editor or at runtime.

## Use scripts within the Editor

When using (or running) scripts within the Editor, use:

* [`Scene`](../ScriptReference/SceneManagement.Scene.md) struct: Available both in the Editor and at runtime, Scene struct contains read-only properties that relate to the scene itself, such as name and asset path.
* [`EditorSceneManager`](../ScriptReference/SceneManagement.EditorSceneManager.md) API: This class is only available only in the Editor and has several functions to implement all the Multi-Scene editing features described in the pages [Setup multiple scenes](setupmultiplescenes.md) and [Bake data in multiple scenes](bakemultiplescenes.md).
* [`SceneSetup`](../ScriptReference/SceneManagement.SceneSetup.md) utility class: A utility class that you can use to store information about a scene that is in the Hierarchy window.

## Use runtime scripts

When using scripts at runtime to edit multiple scenes, use the functions in the [`SceneManager`](../ScriptReference/SceneManagement.SceneManager.md) class such as `LoadScene` and `UnloadScene`.

**Tips**:

* To instantiate a [prefab](Prefabs.md) in a scene, use [`PrefabUtility.InstantiatePrefab`](../ScriptReference/PrefabUtility.InstantiatePrefab.md).
* To move objects to the root of a scene, use [`Undo.MoveGameObjectToScene`](../ScriptReference/Undo.MoveGameObjectToScene.md).
* To avoid having to setup your Hierarchy window every time you restart the Editor, or to store different setups, use [`EditorSceneManager.GetSceneManagerSetup`](../ScriptReference/SceneManagement.EditorSceneManager.GetSceneManagerSetup.md), which also gets a list of SceneSetup objects that describes the current setup. You can serialize the objects into a `ScriptableObject` along with any other information you want to store about your scene setup. To restore your Hierarchy window, recreate the list of `SceneSetups` and use [`EditorSceneManager.RestoreSceneManagerSetup`](../ScriptReference/SceneManagement.EditorSceneManager.RestoreSceneManagerSetup.md).
* To get a list of your loaded scenes at runtime, get the `sceneCount` and iterate over the scenes with [`GetSceneAt`](../ScriptReference/SceneManagement.SceneManager.GetSceneAt.md).