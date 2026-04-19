# Customizing new scene creation

To run custom code when Unity instantiates a new **scene** from a template, create a Scene Template Pipeline script and connect it to the template. Each time you create a new scene from the template, Unity creates a new instance of the pipeline script as well.

To connect the script to a template:

1. Inspect the template to [edit its properties](scene-templates-editing.md).
2. Set the **Scene Template Pipeline** property to point to your Scene Template Pipeline script.

You can also use the [`SceneTemplateAsset.templatePipeline`](../ScriptReference/SceneTemplate.SceneTemplateAsset-templatePipeline.md) method to connect the script to the template via C#.

A Scene Template Pipeline script must derive from the [`ISceneTemplatePipeline`](../ScriptReference/SceneTemplate.ISceneTemplatePipeline.md) interface or [`SceneTemplatePipelineAdapter`](../ScriptReference/SceneTemplate.SceneTemplatePipelineAdapter.md). It should implement the events you want to react to; for example, `BeforeTemplateInstantiation` or `AfterTemplateInstantiation` in the code below.

**Example:**

```
using UnityEditor.SceneTemplate;
using UnityEngine;
using UnityEngine.SceneManagement;
public class DummySceneTemplatePipeline : ISceneTemplatePipeline
{
    public void BeforeTemplateInstantiation(SceneTemplateAsset sceneTemplateAsset, bool isAdditive, string sceneName)
    {
        if (sceneTemplateAsset)
        {
            Debug.Log($"Before Template Pipeline {sceneTemplateAsset.name} isAdditive: {isAdditive} sceneName: {sceneName}");
        }
    }

    public void AfterTemplateInstantiation(SceneTemplateAsset sceneTemplateAsset, Scene scene, bool isAdditive, string sceneName)
    {
        if (sceneTemplateAsset)
        {
            Debug.Log($"After Template Pipeline {sceneTemplateAsset.name} scene: {scene} isAdditive: {isAdditive} sceneName: {sceneName}");
        }
    }
}
```

## Scene Template instantiation sequence

When you create a [new scene from a template](CreatingScenes.md#new-scene-dialog) with cloneable dependencies, Unity performs several file operations. Most of these operations trigger Unity events that you can listen for, and react to, in **scripts**.

After you create a new scene, Unity performs the following actions:

1. Creates a scene template asset, a template scene associated with the template, and a new instance of the template scene.
2. Triggers the `ISceneTemplatePipeline.BeforeTemplateInstantiation` event for the template asset, and binds the asset to a `ISceneTemplatePipeline` script that it triggers.
3. Triggers the [`SceneTemplate.NewTemplateInstantiating`](../ScriptReference/SceneTemplate.SceneTemplateService-newSceneTemplateInstantiating.md) event.
4. Creates a new scene that’s a copy of the template scene.
5. Creates a folder with the same name as the new scene, and copies all cloneable dependencies into that folder.
6. Opens the new scene in memory, and triggers the following events:
   * [`EditorSceneManager.sceneOpening`](../ScriptReference/SceneManagement.EditorSceneManager-sceneOpening.md)
   * [`MonoBehavior.OnValidate`](../ScriptReference/MonoBehaviour.OnValidate.md) (on all GameObjects that implement it)
   * [`EditorSceneManager.sceneOpened`](../ScriptReference/SceneManagement.EditorSceneManager-sceneOpened.md)
7. Remaps references to all cloneable assets, so the new scene points to the clones.
8. Saves the new scene, and triggers the following events:
   * [`EditorSceneManager.sceneSaving`](../ScriptReference/SceneManagement.EditorSceneManager-sceneSaving.md)
   * [`EditorSceneManager.sceneSaved`](../ScriptReference/SceneManagement.EditorSceneManager-sceneSaved.md)
9. Triggers the `ISceneTemplatePipeline.AfterTemplateInstantiation` for the template asset, and binds the asset to a `ISceneTemplatePipeline` script that it triggers.
10. Triggers the [`SceneTemplate.NewTemplateInstantiated`](../ScriptReference/SceneTemplate.SceneTemplateService-newSceneTemplateInstantiated.md) event.

## Additional resources

* [Creating Scene templates](scene-templates-creating.md)
* [Editing Scene templates](scene-templates-editing.md)