# Creating models for animation

This page contains guidance on creating a model for use with Unity’s [Animation System](AnimationOverview.md).

For general best practice guidance on creating all types of models, see [Creating models for optimal performance](ModelingOptimizedCharacters.md). For information on importing a model with animations into Unity, see [Importing a model with humanoid animations](ConfiguringtheAvatar.md) or [Importing a model with non-humanoid (generic) animations](GenericAnimations.md).

![Stages for preparing a character (modeling, rigging, and skinning)](../uploads/Main/Char200.png)

Stages for preparing a character (modeling, rigging, and skinning)

There are two types of models that you can use with Unity’s Animation System:

* A **Humanoid** model is a specific structure, containing at least 15 bones organized in a way that loosely conforms to an actual human skeleton.
* A **Generic** model is everything else. This might be anything from a teakettle to a dragon.

Both require a rigged, skinned model. The stages of creating such a model are:

1. **Modeling**: Creating the model. The model contains the **mesh**, textures, and so on.
2. ****Rigging****: Creating the skeleton, also called the rig or **joint** hierarchy. This defines the bones inside the mesh, and their movement in relation to one another.
3. ****Skinning****: Joining the mesh to the skeleton. This defines which parts of the character mesh move when a given joint is animated.

You usually perform all of these steps in 3D modeling software (such as Autodesk® 3ds Max®, Autodesk® Maya®, or Blender), export the resulting model, and import it into Unity.

## Modeling

There are a few guidelines you can follow to ensure a model works well with animation in a Unity Project:

* Use a topology with a **well-formed structure**. Generally, you should consider how the vertices and triangles of the model distort as it animates. A good topology allows the model to move without distorting the mesh. Study existing character meshes to see how the topology is arranged and why.
* Check the **scale** of your mesh. Do a test import and compare the size of your imported Model with a “meter cube”. The standard Unity cube primitive has a side length of one unit, so it can be taken as a 1m cube for most purposes. Check the units your 3D modeling software uses and adjust the export settings so that the size of the Model is in correct proportion to the cube.
* While you are building, **clean up your model**. Where possible, cap holes, weld vertices, and remove hidden faces. This helps with skinning, especially automated skinning processes.
* If your character has feet, arrange the mesh so that they stand on the local origin or “anchor point” of the model. It makes walking much easier to handle if a character’s anchor point (that is, its transform position) is directly on the floor.
* If your character is a biped, model in a **T-pose**. This gives you space to refine polygon detail where you need it (such as the underarms). This also makes it easier to position your rig inside the Mesh.

![Skin Mesh, textured and triangulated](../uploads/Main/SkinMesh256.png)

Skin Mesh, textured and triangulated

## Rigging

3D modeling software provides a number of ways to create joints for your rig.

For a **Humanoid** model, there are specific requirements for the skeleton so that Unity can match it to a [Humanoid Avatar](class-Avatar.md):

* A minimum of 15 bones. Your skeleton needs to have at least the required bones in place.
* In order to improve your chances for finding a match to the **Avatar**, name your bones in a way that reflects the body parts they represent. For example, “LeftArm” and “RightForearm” make it clear what these bones control.
* The joint/bone hierachy should follow a natural structure for the character you are creating. Given that arms and legs come in pairs, you should use a consistent convention for naming them (for example, “arm\_L” for the left arm, “arm\_R” for the right arm). Possible structures for the hierarchy include:
  + HIPS - spine - chest - shoulders - arm - forearm - hand
  + HIPS - spine - chest - neck - head
  + HIPS - UpLeg - Leg - foot - toe - toe\_end

![Biped Skeleton, positioned in T-pose](../uploads/Main/Skeleton256.png)

Biped Skeleton, positioned in T-pose

For a **Generic** model, the only requirement is that the skeleton contains a bone that you can designate as the **Root node** when you import the model. This effectively defines the model’s **center of mass**.

## Skinning

The initial setup is typically automated. Skinning usually requires a fair amount of work and testing with animations in order to ensure satisfactory results for the skin deformation.

![Interactive Skin Bind, one of many skinning methods](../uploads/Main/Skinning256.png)

Interactive Skin Bind, one of many skinning methods

Some general guidelines for this process include:

* Use an automated process initially to set up some of the skinning (see the skinning tutorials available for your 3D modeling software).
* Create a simple animation for your rig or import some animation data to act as a test for the skinning. This should give you a quick way to evaluate whether or not the skinning looks good in motion.
* Incrementally edit and refine your skinning solution.
* For performance reasons, consider limiting the number of influences when using a soft bind to a maximum of four. This is the maximum number that Unity supports by default. If your Mesh uses more than four influences, you can choose either of these options:
  + Increase the maximum number of influences using the [Skin Weights property](FBXImporter-Rig.md#SkinWeights).
  + Bake the deformation joints before exporting your Model from Maya, using the *Bake Deformer* tool (from Maya’s top menu: **Skin** > **Bake Deformers to Skin Weights**). For more information, see [Bake deformers on a character](http://help.autodesk.com/view/MAYAUL/2018/ENU/?guid=GUID-8B68FF4A-E9BC-44E9-B0D9-53CFE0C008FC) in the Autodesk documentation.

## Exporting and verifying your model

Unity imports a number of different generic and native 3D file formats, but the recommended file format is .fbx. For more information, see [Model file formats](3D-formats.md).

When you use .fbx, you can:

* Export the Mesh with the skeleton hierarchy, normals, textures and animation.
* Re-import the Mesh into your 3D modeling software to verify your animated Model looks as expected.
* Export animations without the Mesh.