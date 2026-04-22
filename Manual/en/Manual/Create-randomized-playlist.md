# Create a randomized playlist with the Audio Random Container

To create a randomized playlist, you must create an Audio Random Container and then play it in the game.

You can use a randomized playlist to randomize and play almost any sound, such as footsteps, impacts, and AI dialog.

## Prerequisites

You can create and use a randomized playlist if you know how to do the following:

* Create a [GameObject](GameObjects.md)
* Add a button to the scene

## Create an Audio Random Container

You can create an Audio Random Container from more than one place.

**Note**: If you drag an existing Audio Random Container to the Hierarchy, you create a GameObject at the Transform location (0,0,0) with an Audio Source component that references the dragged Audio Random Container.

### From the menu

In the Unity Editor, do the following:

1. Select **Assets** > **Create** > **Audio** > **Audio Random Container**.

An empty Audio Random Container asset appears in the **Project window** under **Assets**.

### From the Project view

In the Unity Editor, do the following:

1. Right click in the Project window to bring up the asset menu that contains a working menu item.
2. Select **Create**.
3. Select **Audio**.
4. Select **Audio Random Container**.

An empty Audio Random Container asset appears in the Project window under **Assets**.

### From the window

In the Unity Editor, do the following:

1. Select **Window**.
2. Select **Audio**.
3. Select **Audio Random Container**.

An empty Audio Random Container asset appears in the Project window under **Assets**.

## Play an Audio Random Container in the game

To play an Audio Random Container, you must integrate it with the Audio Source component and then play it, such as with one of the game trigger options.

### Integrate with the Audio Source component

You can only use the Audio Random Container in conjunction with the Audio Source component. It’s the emitter for the Audio Random Container and determines whether the Audio Random Container has a position in the world or sounds like it doesn’t have a position. For more information, refer to [Audio Source](class-AudioSource.md).

To integrate, do the following:

1. Create an Audio Random Container from the menu, the Project view, or the window.
2. Create a GameObject.
3. In the **Inspector** window, select **Add Component**.
4. Select **Audio** > **Audio Source**.
5. Next to **AudioResource**, choose the Audio Random Container you created.

You have integrated the Audio Random Container with the Audio Source component. You can play it with the PlayOnAwake bool, the API, or the different game trigger options.

### Use game triggers

Sounds in games are often triggered by a game event or by a trigger. You can use game triggers, such as buttons, to play the Audio Random Container\:

1. Add a button to a **scene**.
2. In the Inspector window, select **+** to add an [On Click event](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Button.html#events).
3. Select the Audio Random Container.

## Additional resources

* [GameObject](GameObjects.md)
* [Audio Source](class-AudioSource.md)