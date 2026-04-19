# Importing a model with non-humanoid (generic) animations

This page contains guidance on importing a model for use with Unity’s [Animation System](AnimationOverview.md). For information on creating a model for use with the Animation System, see [Creating models for animation](UsingHumanoidChars.md).

The Animation System works with two types of models:

* A **Humanoid** model is a specific structure, containing at least 15 bones organized in a way that loosely conforms to an actual human skeleton. For information on importing this type of model, see [Importing a model with humanoid animations](ConfiguringtheAvatar.md).
* A **Generic** model is everything else. This might be anything from a teakettle to a dragon. This page contains guidance on importing this type of model.

For general importing guidance that is applicable to all types of models, see [Importing a model](ImportingModelFiles.md).

## Overview

When Unity imports a **Generic** model, you must tell it which bone is the **Root node**. This effectively defines the model’s **center of mass**.

Since there is only one bone to map, Generic setups do not use the [Humanoid Avatar window](class-Avatar.md). As a result, preparing to import your non-Humanoid **model file** into Unity requires fewer steps than for Humanoid models.

1. [Set up your Rig](#RigSetup) as **Generic**.
2. You can optionally limit the animation that gets imported on certain bones by [defining an Avatar Mask](#AvatarMask).
3. From the **Animation** tab, enable the **Import Animation** option and then set the other Asset-specific properties, .
4. If the file consists of multiple animations or actions, you can [define specific frame ranges as Animation Clips](Splittinganimations.md).
5. For each **Animation Clip** defined in the file, you can:
   * [Set the pose and root transform](class-AnimationClip.md#ClipProperties)
   * [Optimize looping](LoopingAnimationClips.md)
   * [Add curves to the clip](AnimationCurvesOnImportedClips.md) in order to animate the timings of other items
   * [Add events to the clip](AnimationEventsOnImportedClips.md) in order to trigger certain actions in time with the animation
   * [Discard part of the animation](AnimationMaskOnImportedClips.md) similar to using a runtime [Avatar Mask](class-AvatarMask.md) but applied at import time
   * [Select a different Root Motion Node](class-AnimationClip.md#AllClipProperties) to drive the action from
   * [Read any messages from Unity about importing the clip](class-AnimationClip.md#ImportMessages)
   * [Watch a preview of the animation clip](class-AnimationClip.md#AnimationPreview)
6. To save your changes, click the **Apply** button at the bottom of the **Import Settings** window or **Revert** to discard your changes.

## Setting up the Rig

From the [Rig tab of the Inspector window](FBXImporter-Rig.md), set the **Avatar** (animation) type to **Generic**. By default, the **Avatar Definition** property is set to **Create From This Model** and the **Root node** option is set to **None**.

In some cases, you can change the **Avatar Definition** option to **Copy From Other Avatar** to use an Avatar you already defined for another Model file. For example, if you create a **Mesh** (skin) in your 3D modeling application with several distinct animations, you can export the Mesh to one FBX file, and each animation to its own FBX file. When you import these files into Unity, you only need to create a single Avatar for the first file you import (usually the Mesh). As long as all the files use the same bone structure, you can re-use that Avatar for the rest of the files (for example, all the animations).

![Generic Rig](../uploads/Main/Rig-1.png)

Generic Rig

If you keep the **Create From This Model** option, you must then choose a bone from the **Root node** property.

If you decide to change the **Avatar Definition** option to **Copy From Other Avatar**, you need to specify which Avatar you want to use by setting the **Source** property.

You can also change the maximum number of bones that can influence a given vertex with the **Skin Weights** property. By default, this property limits influence to four bones, but you can specify more or fewer.

When you click the **Apply** button, Unity creates a **Generic Avatar** and adds an Avatar sub-Asset under the Model Asset, which you can find in the Project view.

![Avatar appears as a sub-Asset of the imported Model](../uploads/Main/MecanimFBXNoAvatar2.png)

Avatar appears as a sub-Asset of the imported Model

**Note:** The Generic Avatar is not the same thing as the Humanoid Avatar, but it does appear in the Project view, and it does hold the Root node mapping. However, if you click on the Avatar icon in the Project view to display its properties in the **Inspector**, only its name appears and there is no **Configure Avatar** button.

## Creating an Avatar Mask

You can apply masking to animation clips either during import time, or at runtime. Masking during import time is preferable, because it allows the discarded animation data to be omitted from your build, making the files smaller and therefore using less memory. It also makes for faster processing because there is less animation data to be blended at runtime. In some cases, import masking may not be suitable for your purposes. In that case, you can apply a mask at runtime by creating an **Avatar Mask** Asset, and [using it in the layer settings](AnimationLayers.md) of your **Animator Controller**.

To create an empty Avatar Mask Asset, you can either:

* Choose **Create** > **Avatar Mask** from the **Assets** menu.
* Click the Model object you want to define the mask on in the **Project** view, and then right-click and choose **Create** > **Avatar Mask**.

The new Asset appears in the **Project** view:

![The Avatar Mask window](../uploads/Main/ConfiguringtheAvatar-Mask.png)

The Avatar Mask window

You can now choose which bones to include or exclude from a [Transform hierarchy](class-AvatarMask.md#Transform) and then add the mask to either an [Animation Layer](AnimationLayers.md) or add a reference to it under the [Mask](AnimationMaskOnImportedClips.md) section of the [Animation tab](class-AnimationClip.md).