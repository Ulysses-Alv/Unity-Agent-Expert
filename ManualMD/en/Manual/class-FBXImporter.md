# Model Import Settings reference

To open a model’s Import Settings, select the model from the [Project window](ProjectView.md). The [Inspector](UsingTheInspector.md) then displays the Import Settings for the selected model.

![The Import Settings window](../uploads/Main/model-import-settings.png)

The Import Settings window

**Note:** These settings are for importing models and animations created in most 3D modeling applications. However, models created in SketchUp and SpeedTree use specialized settings. For more information, refer to [SketchUp Import Settings](class-SketchUpImporter.md), and [SpeedTree Import Settings](class-SpeedTreeImporter.md).

You can customize how Unity imports the selected file by setting the properties in the following tabs:

## Model

Contains settings for 3D models, which can represent a character, a building, or a piece of furniture. In these cases, Unity creates multiple assets from a single **model file**. In the Project window, the main imported object is a model [prefab](Prefabs.md). Usually there are also several **mesh** objects that the model prefab references.

For more information, refer to [Model Import Settings reference](FBXImporter-Model.md).

## Rig

Contains settings for rigs (sometimes called skeletons). A rig includes a set of deformers arranged in a hierarchy that animate a mesh (sometimes called skin) on one or more models created in a 3D modeling application such as Autodesk 3ds Max or Autodesk Maya. For humanoid and generic (non-humanoid) models, Unity creates an **avatar** to reconcile the imported rig with the Unity **GameObject**.

For more information, refer to [Rig Import Settings reference](FBXImporter-Rig.md).

## Animation

Contains settings for animation clips. You can define any series of different poses that happen over a set of frames, such as walking, running, or idling as an **animation clip**. You can reuse clips for any model that has an identical rig. Often a single file contains several different actions, each of which you can define as a specific animation clip.

For more information, refer to [Animation Import Settings reference](class-AnimationClip.md).

## Materials

Contains settings for the materials and textures in your model. You can extract materials and textures or leave them embedded within the model. You can also adjust how materials are mapped in the model.

For more information, refer to [Materials Import Settings reference](FBXImporter-Materials.md).

## Additional resources

* [Model import workflows](ImportingModelFiles.md)
* [Model file formats](3D-formats.md)