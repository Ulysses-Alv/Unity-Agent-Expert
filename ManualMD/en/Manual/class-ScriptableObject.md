# ScriptableObject

[Switch to Scripting](../ScriptReference/ScriptableObject.md "Go to ScriptableObject page in the Scripting Reference")

ScriptableObject is a serializable Unity type derived from [`UnityEngine.Object`](class-Object.md). As with [MonoBehaviour](class-MonoBehaviour.md), you don’t instantiate the ScriptableObject class directly but create your own custom C# classes that derive from it, and then create instances of those custom classes, usually through the **Assets** menu in the Unity Editor.

All instances of classes derived from ScriptableObject are commonly referred to as ScriptableObjects. Unlike MonoBehaviours, ScriptableObjects are not attached to **GameObjects** as components but exist in the project as [assets](AssetWorkflow.md), independent of GameObjects. Because ScriptableObjects inherit from `UnityEngine.Object`, you can drag or pick instances of them into fields in the **Inspector**.

The main value of a ScriptableObject is as a data store, but they can also define behavior. A common use for ScriptableObjects is as a container for shared data used by multiple objects at runtime, which can reduce a project’s memory usage by avoiding copies of values.

For example, if your project has a **prefab** that stores unchanging data in attached MonoBehaviour **scripts**, then every new instance of the prefab gets its own copy of the data. Instead of duplicating data like this, you can use a ScriptableObject to store the data and then access it by reference from all the prefabs. This means that there is one copy of the data in memory.

The main use cases for ScriptableObjects are:

* Saving and storing data during an Editor session. This is why many authoring tools in Unity, such as [`EditorTool`](../ScriptReference/EditorTools.EditorTool.md) and [`EditorWindow`](../ScriptReference/EditorWindow.md), derive from `ScriptableObject`.
* Saving data as an asset in your project to use at runtime.

For a complete reference of every member of the ScriptableObject class, refer to the [ScriptableObject script reference](../ScriptReference/ScriptableObject.md).

## Create a ScriptableObject

To create a new ScriptableObject script, the quickest way is to use the predefined ScriptableObject script template from the **Assets** menu in one of the following ways:

* In the main menu, go to **Assets** > **Create** > **Scripting** > and select **ScriptableObject Script**.
* In the [Project window toolbar](ProjectView.md), right-click to open the Project window context menu, then select **Create** > **Scripting** > **ScriptableObject Script**. You can also click the plus sign in the Project window to open the **Create** menu directly.

This gives you a custom base class that inherits from `UnityEngine.ScriptableObject`. You can then use the [CreateAssetMenu](../ScriptReference/CreateAssetMenuAttribute.md) attribute to create instances of this class, each of which becomes an asset in your project.

## Example: Instantiate prefabs with a ScriptableObject

The following example uses a ScriptableObject to store data defined at authoring time that is later used to determine where to instantiate prefabs at runtime. First, create the following base ScriptableObject class in your `Assets` folder:

```
using UnityEngine;
// Use the CreateAssetMenu attribute to allow creating instances of this ScriptableObject from the Unity Editor.
[CreateAssetMenu(fileName = "Data", menuName = "ScriptableObjects/SpawnManagerScriptableObject", order = 1)]
public class SpawnManagerScriptableObject : ScriptableObject
{
    public string prefabName;

    public int numberOfPrefabsToCreate;
    public Vector3[] spawnPoints;
}
```

With the previous script in your `Assets` folder, create an instance of your new ScriptableObject by navigating to **Assets > Create > ScriptableObjects > SpawnManagerScriptableObject**. Give your new ScriptableObject instance a meaningful name and alter the values. To use these values at runtime, you need to create a new script that references your ScriptableObject, in this case, a `SpawnManagerScriptableObject` as follows:

```
using UnityEngine;

public class ScriptableObjectManagedSpawner : MonoBehaviour
{
    // The GameObject to instantiate.
    public GameObject entityToSpawn;

    // An instance of the ScriptableObject defined above.
    public SpawnManagerScriptableObject spawnManagerValues;

    // This will be appended to the name of the created entities and increment when each is created.
    int instanceNumber = 1;

    void Start()
    {
        SpawnEntities();
    }

    void SpawnEntities()
    {
        int currentSpawnPointIndex = 0;

        for (int i = 0; i < spawnManagerValues.numberOfPrefabsToCreate; i++)
        {
            // Creates an instance of the prefab at the current spawn point.
            GameObject currentEntity = Instantiate(entityToSpawn, spawnManagerValues.spawnPoints[currentSpawnPointIndex], Quaternion.identity);

            // Sets the name of the instantiated entity to be the string defined in the ScriptableObject and then appends it with a unique number. 
            currentEntity.name = spawnManagerValues.prefabName + instanceNumber;

            // Moves to the next spawn point index. If it goes out of range, it wraps back to the start.
            currentSpawnPointIndex = (currentSpawnPointIndex + 1) % spawnManagerValues.spawnPoints.Length;

            instanceNumber++;
        }
    }
}
```

**Note:** The script file must have the same name as the class.

Attach the previous script to a GameObject in your [Scene](CreatingScenes.md). Then, in the Inspector, populate the **Spawn Manager Values** field with the new `.asset` instance of `SpawnManagerScriptableObject` that you set up.

Set the **Entity To Spawn** field to any prefab in your `Assets` folder, then enter Play mode. The prefab you referenced in the `ScriptableObjectManagedSpawner` instantiates using the values you set in the `SpawnManagerScriptableObject` instance.

If you’re working with ScriptableObject references in the Inspector, you can double click the reference field to open the Inspector for your ScriptableObject. You can also create a custom Inspector for your type to help manage the data that it represents.

## Saving changes to ScriptableObject data

In the Unity Editor, you can save data to ScriptableObjects in Edit mode and Play mode. In a standalone Player at runtime, you can only read saved data from the ScriptableObject assets. When you use Editor authoring tools or the Inspector to modify a ScriptableObject asset, Unity automatically writes the data to disk and it persists between Editor sessions.

However, Unity doesn’t automatically save changes to a `ScriptableObject` made via script in Edit mode. In these cases, you must call [`EditorUtility.SetDirty`](../ScriptReference/EditorUtility.SetDirty.md) on the ScriptableObject to ensure Unity’s serialization system recognizes it as changed and saves the changes to disk. Without this, changes may not persist between Editor sessions.

The following is a simple ScriptableObject for storing game settings:

```
using UnityEngine;

[CreateAssetMenu(fileName = "GameSettings", menuName = "ScriptableObjects/GameSettingsScriptableObject", order = 2)]
public class GameSettingsScriptableObject : ScriptableObject
{
    public int highScore;
}
```

Create a new instance of the `GameSettingsScriptableObject` in your project via **Assets** > **Create** > **ScriptableObjects** > **GameSettingsScriptableObject**. Then, in the Inspector, set the `highScore` value.

The following Editor script adds a simple window with a button for increasing the high score at **Window** > **Game Settings Editor**.

```
using UnityEditor;
using UnityEngine;

public class GameSettingsEditor : EditorWindow
{
    GameSettingsScriptableObject settings;

    // Create a simple Editor Window to modify the high score in the ScriptableObject
    // Accessible at Window > Game Settings Editor
    [MenuItem("Window/Game Settings Editor")]
    public static void ShowWindow()
    {
        GetWindow<GameSettingsEditor>("Game Settings Editor");
    }

    void OnGUI()
    {
        settings = EditorGUILayout.ObjectField("Settings", settings, typeof(GameSettingsScriptableObject), false) as GameSettingsScriptableObject;
        if (settings == null) return;

        EditorGUILayout.LabelField("High Score", settings.highScore.ToString());

        // Click the Increase High Score button to increase the high score by 10
        if (GUILayout.Button("Increase High Score"))
        {
            settings.highScore += 10;

            // Call SetDirty to ensure the change is saved
            EditorUtility.SetDirty(settings);
            // Optional: AssetDatabase.SaveAssets(); // To save immediately
        }
    }
}
```

Without the call to [`EditorUtility.SetDirty`](../ScriptReference/EditorUtility.SetDirty.md) in this example, the change to `highScore` appears in memory, but if you close and reopen the Editor the value reverts to its previous value.

## Additional resources

* [`ScriptableObject` API reference](../ScriptReference/ScriptableObject.md)
* [Introduction to ScriptableObjects](https://learn.unity.com/tutorial/introduction-to-scriptable-objects)