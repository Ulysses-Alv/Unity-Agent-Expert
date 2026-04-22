# Support for proprietary model file formats

Unity supports a number of proprietary **model file** formats. You should not use these file formats in production; instead, export to the .fbx format wherever possible. For more information, see [Model file formats](3D-formats.md).

**Note:** This page contains guidance on using proprietary file formats that use FBX conversion. However, there are two file formats that do not use FBX as an intermediary: [SketchUp](https://www.sketchup.com/) and [SpeedTree](https://store.speedtree.com/unity/). For more information about limitations with these file formats, see [SketchUp Settings](class-SketchUpImporter.md) and [SpeedTree](class-SpeedTreeImporter.md).

When Unity imports a proprietary file, it launches the 3D modeling software in the background. Unity then communicates with that proprietary software to convert the native file into a format Unity can read.

The first time you import a proprietary file into Unity, the 3D modeling software has to launch in a command-line process. This can take a while, but subsequent imports are very quick.

**Note:** As of Unity 2019.3, Unity no longer provides built-in support for Cinema4D files. To continue using Cinema4D files in Unity for versions 2019.3 and later, install [Maxon’s Cinema4D importer](https://assetstore.unity.com/packages/tools/integration/cineware-by-maxon-158381). Alternatively, you can export an FBX file from Cinema4D instead.

## Requirements

You need to have the 3D modeling software installed to import proprietary files directly into Unity. If you don’t have the software installed, use the FBX format instead.

For more information about importing FBX files, see [Model Import Settings window](class-FBXImporter.md).

## Application-specific issues

You [import files](ImportingModelFiles.md) in the same way, regardless of whether they are generic or proprietary files. However, there are some differences between which features are supported. For more information on the limitations with a specific 3D application, see:

* [Importing objects from Autodesk® Maya®](#Maya)
* [Importing objects from Autodesk® 3ds Max®](#Max)
* [Importing objects from Cheetah3D](#Cheetah3D)
* [Importing objects from Modo](#Modo)
* [Importing objects from LightWave](#Lightwave)
* [Importing objects from Blender](#Blender)

## Importing objects from Autodesk® Maya®

Unity imports Autodesk® Maya® files (*.mb* and *.ma*) through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with vertex colors, normals and up to 2 UV sets
* Materials with texture and diffuse color; multiple materials per mesh
* Animation
* Joints
* Blend shapes
* Lights and Cameras
* Visibilty
* Custom property animation

### Limitations

Unity does not support Autodesk® Maya®’s *Rotate Axis* (pre-rotation).

Joint limitations include:

* **Joint** Orient (joint only post-rotation)
* Segment Scale Compensate (joint only option)

Unity imports and supports any Rotate Order you specify in Autodesk® Maya®; however, once imported, you cannot change that order again inside Unity.

If you import a Model that uses a different rotation order from Unity’s, Unity [displays that rotation order](AnimationEulerCurveImport.md#RotationOrder) in the **Inspector** beside the **Rotation** property.

### Tips and troubleshooting

* Keep your **scene** simple: only export the objects you need in Unity when exporting.
* Unity only supports polygons, so convert any patches or NURBS surfaces into polygons before exporting; see [Autodesk® Maya® documentation](https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2015/ENU/Maya/files/Polygon-selection-and-creation-Convert-NURBS-surfaces-to-a-polygon-mesh-htm.html) for instructions.
* If your model did not export correctly, the node history in Autodesk® Maya® might be causing a problem. In Autodesk® Maya®, select **Edit** > **Delete by Type** > **Non-Deformer History** and then re-export the model.
* The Autodesk® Maya® FBX Exporter bakes un-supported complex animations constraints, such as Set Driven Keys, in order to import the animation into Unity properly. If you are using Set Driven Keys in Autodesk® Maya®, make sure to set keys on your drivers in order for the animation to be baked properly. For more information, see the Autodesk® Maya® documentation on **Keyframe** Animation.
* In Autodesk® Maya®, the visibility value is present on each shape but can’t be animated and is not exported to the FBX file. Always set the visibility value on a node and not on a shape.

## Importing Objects From Autodesk® 3ds Max®

Unity imports Autodesk® 3ds Max® files (*.max*) through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with vertex colors, normals and one or more UV sets
* Materials with diffuse texture and color. Multiple materials per mesh
* Animations
* Bone-based animations
* Morph targets (blend shapes)
* Visibility

## Importing objects from Cheetah3D

Unity imports Cheetah3D files (*.jas*) through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with vertices, polygons, triangles, UVs, and normals
* Animations
* Materials with diffuse color and textures

## Importing objects from Modo

Unity imports Modo files (*.lxo*) through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with vertices, normals and UVs.
* Materials with Texture and diffuse color; multiple Materials per mesh
* Animations

To get started, save your **.lxo** file in your Project’s Assets folder. In Unity, the file appears in the Project View.

Unity re-imports the Asset when it detects a change in the .lxo file.

## Importing objects from Lightwave

Unity imports Lightwave files through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with up to 2 UV channels
* Normals
* Materials with Texture and diffuse color; multiple materials per mesh
* Animations
* Bone-based animations

You can also configure the Lightwave AppLink **plug-in** which automatically saves the FBX export settings you use the first time you import your Lightwave scene file into Unity.
For more information, refer to the [Lightwave Unity Interchange documentation](https://docs.lightwave3d.com/2025/unity.html).

### Limitations

Bake your Lightwave-specific materials as textures so that Unity can read them. For information on doing this using a non-destructive pipeline, refer to Lightwave’s [Node Editor documentation](https://docs.lightwave3d.com/2025/node-editor.html).

Unity does not support splines or patches. Convert all splines and patches to polygons before saving and exporting to Unity. For more information, refer to the [Lightwave documentation](https://docs.lightwave3d.com/2025/2025.html).

## Importing objects from Blender

Unity imports [Blender](https://docs.blender.org/) (*.blend*) files through the FBX format, supporting the following:

* All nodes with position, rotation and scale; pivot points and names are also imported
* Meshes with vertices, polygons, triangles, UVs, and normals
* Bones
* Skinned Meshes
* Animations

### Limitations

Textures and diffuse color are not assigned automatically. You can manually assign them by dragging the texture onto the **mesh** in the **Scene View** in Unity.

Blender does not export the visibility value inside animations in the FBX file.