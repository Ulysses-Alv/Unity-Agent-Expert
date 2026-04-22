# Create a Video Player component

Create a Video Player component to play videos in your **scene**.

There are a few ways to create a Video Player component in Unity. Choose from one of the following methods:

* [Create a Video Player from the menu or hierarchy](#menu)
* [Add the component to an existing GameObject](#existing)
* [Drag a video clip into the scene](#drag)
* [Use C# scripting to add the component at runtime](#script)

After you set up your component, you can configure the settings to your liking. For information about the settings, refer to [Video Player component reference](class-VideoPlayer.md).

## Create a Video Player from the menu or hierarchy

To create the Video Player component from the menu, select **GameObject** > **Video** > **Video Player**.

Alternatively, in the **Hierarchy**, select the **Add (+)** button and then select **Video** > **Video Player**.

Either method creates a new **GameObject** with a Video Player component attached.

## Add the component to an existing GameObject

You can also add a Video Player component to an existing GameObject. To do this, either drag a video file onto the GameObject, or:

1. Click on the GameObject.
2. In the **Inspector** window, select **Add Component**.
3. Select **Video** > **Video Player**.

As a result, your GameObject contains a Video Player component.

## Drag a video clip into the scene

The quickest way to create a Video Player component is to drag your [video clip](video-clips-use.md) into the scene:

1. Locate the video file in your `Asset` folder in Unity.
2. Click and drag your file into an empty spot in the Scene Hierarchy, in the **Scene view**, or in the Inspector.

As a result, Unity creates a GameObject that contains a Video Player component and automatically assigns your video clip to the component.

## Use C# scripting to add the component at runtime

You can also use scripting to add a VideoPlayer component to a GameObject and configure its settings using C# **scripts**. Here’s an example code snippet:

```
 // Add a Video Player component to the GameObject.
    VideoPlayer videoPlayer = gameObject.AddComponent<VideoPlayer>();
```

This code dynamically adds a Video Player component to the GameObject that contains this script at runtime.

You can then use the `videoPlayer` variable to access and modify the Video Player’s properties.

## Additional resources

* [VideoPlayer API](../ScriptReference/Video.VideoPlayer.md)
* [Video Player](VideoPlayer.md)
* [Video Player component targets](VideoPlayer-intro.md)
* [Video Player component reference](class-VideoPlayer.md)
* [Set up your Render Texture to display video](VideoPlayer-rendertexture.md)
* [Create a Video Player component](VideoPlayer-instructions.md)
* [Video Clip Importer reference](class-VideoClip.md)