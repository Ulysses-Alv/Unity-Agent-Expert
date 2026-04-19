# Scene templates introduction

To create new scenes, Unity copies a **scene** template, which is a pre-configured scene that contains all the content you want to start with. For example, the default Basic template usually contains a **Camera** and a light.

You can create your own scene templates to customize the types of new scene you can create in a project. For example, you might create templates for each level in a game so that everyone working on the project can start their scenes with the correct assets and configuration.

You can create a template from any Unity scene. After you create a template, you can create any number of new scenes from it. Like scenes, most scene templates are assets that are stored in the project.

## Types of scene templates

Most scene templates are user-defined, meaning you create them from your own scenes. User-defined scene templates are assets that Unity stores in the project.

Unity also ships with built-in templates for each project type. For example, some project types include a Basic template that creates a scene with a Camera and a light, and an Empty template that creates an empty scene.

Built-in templates are different from other templates because they’re not assets stored in the project, and you can’t modify them. **Note:** Some Unity packages might also include scene templates that they install when you install the package.

## Templates and scene dependencies

When you create a scene template, you can specify whether Unity clones or references its dependencies (the assets it includes) when you create a new scene from it.

To specify which assets Unity clones for a specific template, [edit the template’s properties](scene-templates-editing.md).

A typical template might contain a mix of cloned and referenced assets. Unity sets several asset types to clone by default.

To change whether Unity clones or references a given asset type by default, edit the [Scene template project settings](scene-templates-settings.md).

### Cloning template assets

Cloned assets are copies of the original assets that the template scene uses. When Unity creates the new scene from the template, it automatically modifies the new scene to use any cloned assets. If you modify the cloned assets, it doesn’t affect the template scene. If you modify the original assets in the template scene, it doesn’t affect the new scene.

Cloning template assets is useful when you want new scenes to contain a starting set of assets that users might modify.

### Referencing template assets

Referenced assets are the original assets that the template scene uses. When Unity creates the new scene from the template, the new scene points to the same assets as the template scene. If you modify those assets, it affects both the new scene and the template scene.

Referencing template assets is useful when you want new scenes to contain a default set of assets that users build on top of, but don’t modify.

## Additional resources

* [Creating scene templates](scene-templates-creating.md)
* [Editing scene templates](scene-templates-editing.md)
* [Customizing new scene creation](scene-templates-customizing-scene-instantiation.md)