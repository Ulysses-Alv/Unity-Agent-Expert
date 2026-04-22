# Create a new Animation Clip

To create a new **Animation Clip**, do the following:

1. Select a **GameObject** in your **Scene**.
2. Go to **Window** > **Animation** > **Animation** to open the **Animation Window**,

If the GameObject is not assigned an animation clip, the Create button displays in the centre of the Animation Window.

![Use the Create button to create a new Animation Clip and assign it to the selected GameObject](../uploads/Main/AnimationEditorNewClip.png)

Use the Create button to create a new Animation Clip and assign it to the selected GameObject

Click the **Create** button. Unity prompts you to save your new empty Animation Clip in your **Assets** folder. When you save this new empty Animation Clip, Unity does the following:

* Creates a new Animator Controller Asset.
* Adds the new clip into the Animator Controller as the default state.
* Adds an **Animator Component** to the selected GameObject.
* Assigns the new Animator Controller to the Animator Component.

The required elements of the animation system are set up. You can begin animating the GameObject.

## Create another Animation Clip

If the selected GameObject already has one or more Animation Clips assigned and you open the Animation window, the Create button is not displayed. Instead, one of the animation clips assigned to the selected GameObject is displayed.

To switch between Animation Clips, use the menu in the top-left of the Animation window, under the playback controls.

To create a new Animation Clip on a GameObject that has existing animations, select **Create New Clip** from this menu. Unity prompts you to save your new empty Animation Clip.

![Option to create a new animation clip](../uploads/Main/AnimationEditorNewClipMenu.png)

Option to create a new animation clip

## How it fits together

The above steps automatically set up the components and assets needed to animate a GameObject. It is useful to understand how these components, clips, and assets connect together:

* A GameObject must have an **Animator** component
* The Animator component must have an **Animator Controller** Asset assigned
* The Animator Controller asset must have one or more Animation Clips assigned

The diagram below shows how Unity assigns these components and assets, starting from a new Animation Clip.

![How animation clips, components, and assets are linked together](../uploads/Main/AnimationNewClipAutoSetup.png)

How animation clips, components, and assets are linked together

After you create a new Animation Clip, the follow happens:

* The Animation Window displays a timeline with a white playback line. The `Cube Animation Clip` clip is selected in the clip menu.
* In the **Inspector** window, the Cube GameObject has an Animator Component. The **Controller** field is assigned the `Cube` Animator Controller Asset.
* The **Project Window** has two new Assets: the `Cube` Animator Controller asset and the `Cube Animation Clip` animation clip asset.
* The **Animator Window** displays the contents of the `Cube` Animator Controller asset. The `Cube Animation Clip` is set as the default state, as indicated by the orange color.

![New Animation Clip in the Project Window](../uploads/Main/AnimationEditorNewAnimationAdded.png)

New Animation Clip in the Project Window