# Edit prefab assets

To edit **prefab** assets, you can open them in an isolated editing mode which allows you to view and edit the contents of the prefab asset separately from any other **GameObjects** in the **scene**. Changes that you make in prefab editing mode affect all instances of that prefab.

## Open a prefab asset in prefab editing mode

You can open a prefab asset in prefab editing mode in the following ways:

* **In isolation:** Unity hides the rest of the current working scene, and displays only the GameObjects that relate to the prefab asset in the Hierarchy.
* **In context:** The rest of the current working scene remains visible, but locked for editing. The Hierarchy displays only the GameObjects that relate to the prefab asset.

When you open prefab editing mode either in context or in isolation, the **Scene view** displays only the contents of that prefab. At the top of the Scene view, a breadcrumb bar displays the name of the open prefab asset. You can use this bar to navigate back to the main scene or other prefab assets you might have open.

The Hierarchy window displays a bar at the top which shows the currently open prefab. You can use the back arrow in the header bar to navigate back one step, which is equivalent to clicking the previous breadcrumb in the bar in the Scene view.

![The Scene view and Hierarchy, with prefab editing mode in isolation. The prefab breadcrumb bar is highlighted.](../uploads/Main/prefabs-isolation-mode.png)

The Scene view and Hierarchy, with prefab editing mode in isolation. The prefab breadcrumb bar is highlighted.

### Open in isolation

To open a prefab asset and edit it in isolation, perform one of the following actions:

* Double-click the prefab asset in the **Project window**.
* Select a prefab asset in the Project window. Select **Open** in the top right of the **Inspector** window.

Unity then opens the prefab asset in isolation in prefab editing mode.

### Open in context

To open a prefab asset and edit it in context, perform one of the following actions:

* Select a prefab instance in the Hierarchy window. Select **Open** in the top right of the Inspector window.
* Select a prefab instance in the Hierarchy window and press **P** on the keyboard. This is the default [keyboard binding](ShortcutsManager.md).
* Select the arrow next to the prefab instance in the Hierarchy window.

Unity then opens the prefab asset in context in prefab editing mode.

By default, Unity opens prefab instances in context when you select **Open** from the Inspector. To open the prefab instance in isolation, hold down **Alt** (macOS: **Option**) and select **Open**.

To change the default behavior, perform the following steps:

1. Open the Preferences window: **Edit** > **Preferences** (macOS: **Unity** > **Settings**).
2. Go to **General** > **Hierarchy window** > **Default Prefab Mode**, and change the setting to **In Isolation**.

To open the prefab asset in context after adjusting the default behavior, select the prefab instance and press **P**. You can also set up custom shortcuts for editing prefabs in the [Shortcuts window](ShortcutsManager.md).

## Edit a prefab asset in isolation

By default, when you edit a prefab asset in isolation mode, Unity displays it in the default **skybox**, out of context of the scenes in your project. You can assign a scene as an editing environment if you prefer to edit prefab assets against a background related to your project. To change the default environment, perform the following steps:

1. Open the Editor Project Settings: **Edit** > **Project Settings** > **Editor**.
2. Under **Prefab Mode** > **Editing Environments**, select a scene for regular prefabs and UI prefabs. UI prefabs are defined as prefabs which have a [Rect Transform](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/class-RectTransform.html) component on the root, rather than a regular Transform component.

Unity only uses this editing environment when you open prefab editing mode in isolation.

## Edit a prefab asset in context

When you [open prefab editing mode in context](#open-in-context), by default Unity grays out the rest of the context of the scene.

![The Scene view and Hierarchy with prefab editing mode in context, and the Context set to Gray. The prefab options are highlighted.](../uploads/Main/prefab-mode-context-grey.png)

The Scene view and Hierarchy with prefab editing mode in context, and the **Context** set to **Gray**. The prefab options are highlighted.

To change this view, select one of the **Context** options in the bar at the top of the Scene view:

* **Normal**: Displays the prefab in context with the rest of the scene in its regular colors.
* **Gray**: Displays the prefab in context with the rest of the scene in gray.
* **Hidden**: Displays the prefab in context with the rest of the scene hidden.

You can’t select any other GameObjects that are in the rest of the scene, so you can only focus on editing the prefab. When **Context** is set to anything other than **Hidden**, you can continue to use Unity’s snapping features to snap any GameObjects that are part of the prefab to the scene.

### Editing prefab properties

Unity displays the prefab contents at a position that matches that of the prefab instance it was opened through. This means that you might preview the **root transform** of the prefab contents with different position and rotation values than the prefab asset actually has.

**Important**: You can’t edit the **Transform** values while prefab editing mode is open in context. You must [open the prefab in isolation](#open-in-isolation) to edit these values.

To visualize the [overrides on the main prefab asset](PrefabInstanceOverrides.md), select the **Show Overrides** setting in the top bar of the Scene view. Disable this option to view the default properties of the prefab asset.

## Automatically save edits to the prefab asset

To automatically save any changes you make to the prefab asset, select **Auto Save** in the top right of the Scene view, while in prefab editing mode. You can disable this option if editing a complex prefab is slow. Unity then displays a dialog asking if you want to save the prefab asset when you exit prefab editing mode.

You can disable **Auto Save** completely and remove the option from prefab editing mode. To do this:

1. Open the Editor Project Settings: **Edit** > **Project Settings** > **Editor**.
2. Under **Prefab Mode**, disable the **Allow Auto Save** setting.

## Additional resources

* [Create prefabs](CreatingPrefabs.md)
* [Create variations of prefabs](PrefabVariants.md)
* [Override prefab instances](PrefabInstanceOverrides.md)