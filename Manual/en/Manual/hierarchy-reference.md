# Hierarchy window reference

Explore the settings and functions in the Hierarchy window to view and organize objects in your **scene**.

The Hierarchy window displays every [GameObject](GameObjects.md) in a scene, including models, cameras, and [prefabs](Prefabs.md). You can use the Hierarchy window to manage and group the GameObjects you use in a scene. When you add or remove GameObjects in the Scene view, you also add or remove them from the Hierarchy window.

For more information about creating and organizing items in the Hierarchy window, refer to [Manage GameObjects in the Hierarchy window](Hierachy).

## Icons in the Hierarchy window

Icons in the Hierarchy window indicate the type, visibility, and pickability of items in your scene.

### GameObject icons in the Hierarchy

The icons next to the GameObject name indicate whether the item is a basic GameObject or a prefab, and whether a prefab has been modified from its original values.

You can identify the following types of items using the icons in the Hierarchy window.

| Item | Icon | Description |
| --- | --- | --- |
| GameObject | The icon for a GameObject is an outline of a cube. | The foundational building blocks of all items in a scene. GameObjects are containers for components, which provide functionality.   For more information, refer to [Introduction to GameObjects](GameObjects.md) |
| Prefab | The icon for a prefab is a blue cube. | Prefabs are assets that act as a template for specific items in your scene. When you edit a prefab asset, updates are applied to all instances of that asset that appear in your scene.   For more information, refer to [Introduction to prefabs](Prefabs.md). |
| Prefab variant | The icon for a prefab variant is a blue cube with one side shaded gray. | A prefab variant is a prefab asset that inherits some properties from a different prefab asset, while others properties have new values. Essentially, it’s a new prefab that’s based on an original, base prefab with some changes, or overrides. Editing the base prefab also edits the variant, except for properties that are already changed in the variant.   For more information, refer to [Create variations of prefabs](PrefabVariants.md). |
| Prefab model | The icon for a prefab model is a blue cube with one side gray and a line cut into the top. | Prefab models are prefabs that are imported from other digital content creation (DCC) software. Prefab models can contain skeletons, materials, meshes, animations, and more. |

### Scene visibility icons

The scene visibility status controls whether the GameObject is hidden or displayed in the **Scene view**. You can toggle visibility on and off using the status indicator at the edge of the Hierarchy window. For more information, refer to [Scene visibility](SceneVisibility.md).

The following visibility states are available.

| State | Icon | Description |
| --- | --- | --- |
| Visible | The icon for a visible GameObject is an eye. | The GameObject and all its children are visible in the scene. This is the default status for a new GameObject. This status icon is visible only when you hover over the GameObject. |
| Hidden | The icon for a hidden GameObject is an eye with a strikethrough. | The GameObject and all its children are hidden in the scene. |
| Visible parent | The icon for a visible parent GameObject is an eye with a dot underneath the right corner. | The GameObject is visible in the scene, but one of more of its children are hidden. |
| Hidden parent | The icon for a hidden parent GameObject is an eye with a dot underneath the right corner and a strikethrough. | The GameObject is hidden in the scene, but one of more of its children are visible. |

### Scene pickability icons

The scene pickability status controls whether you can select the GameObject in the Scene view. You can toggle pickability on and off using the status indicator at the edge of the Hierarchy window. For more information, refer to [Scene picking controls](ScenePickingControls.md).

The following pickability states are available.

| State | Icon | Description |
| --- | --- | --- |
| Pickable | The icon for a pickable GameObject is a hand with pointing finger. | You can select the GameObject, and any of its children, in the Scene view. This is the default status for a new GameObject. This status icon is visible only when you hover over the GameObject. |
| Not pickable | The icon for a not pickable GameObject is hand with pointing finger with a strikethrough. | You can’t select the GameObject, or any of its children, in the Scene view. |
| Pickable parent | The icon for a pickable parent GameObject is hand with pointing finger with a dot underneath the right corner. | You can select the GameObject in the Scene view, but you can’t select one or more of its children. |
| Not pickable parent | The icon for a not pickable parent GameObject is a hand with pointing finger with a dot underneath the right corner and a strikethrough. | You can’t select the GameObject in the Scene view, but you can select one or more of its children. |

## New Hierarchy

Enable the [**New Hierarchy** setting](Preferences.md#hierarchy-window) in the Preferences window (**Edit** > **Preferences** (macOS: **Unity > Settings**) > **General**) to change the view of the Hierarchy window. This view displays extra information about the GameObjects in your scene such as their [layers](Layers.md) and [tags](Tags.md), and optionally highlights each alternate row for easier navigation of items.

**Important**: The **New Hierarchy** setting is in preview and might change in future versions of Unity.

![New Hierarchy window setting enabled. The Hierarchy window has columns with additional information about the objects in the scene.](../uploads/Main/hierarchy-window-new.png)

New Hierarchy window setting enabled. The Hierarchy window has columns with additional information about the objects in the scene.

Right-click on the header row of the Hierarchy window to customize the information displayed:

| **Option** | **Description** |
| --- | --- |
| **Resize to Fit** | Resizes the columns to fit the width of the Hierarchy window. |
| **Visibility** | Displays a column to control the [scene visibility status](#scene-visibility-icons) of the GameObject. |
| **Picking** | Displays a column to control the [scene pickability status](#scene-pickability-icons) of the GameObject. |
| **Active** | Displays a column to control the [active status](class-GameObject.md#active-status) of the GameObject. |
| **Static** | Displays a column to control the [static status](class-GameObject.md#static-status) of the GameObject. |
| **Layer** | Displays a column to view and control the [layer](Layers.md) assigned to the GameObject. Use the dropdown menu to assign a new layer to the GameObject. |
| **Tag** | Displays a column to view and control the [tags](Tags.md) assigned to the GameObject. Use the dropdown menu to assign a new tag to the GameObject. |
| **Reset Columns** | Resets the Hierarchy window to the default columns, which are **Picking**, **Active**, and **Name**. |

Unity displays only the values that are different from the default in each column.

## Override indicators

When you edit a prefab instance in a scene, Unity displays an indicator next to the parent GameObject in the Hierarchy window. This indicator highlights any prefab that has non-default override values in any of its child GameObjects. To open the [Overrides dropdown](prefab-instance-inspector-reference.md) directly from the Hierarchy window, select the override indicator. The override indicator appears as a blue line in the left side of the margin and is identical to the instance override in the **Inspector** window. For more information, refer to [Override prefab instances](PrefabInstanceOverrides.md).

![The override indicator is displayed next to prefab A when its child, GameObject C, has a value in a non-default state.](../uploads/Main/hierarchy-override-indicator.png)

The override indicator is displayed next to prefab A when its child, GameObject C, has a value in a non-default state.

Any objects that are added to the instance have a plus badge on their icons.

![The Hierarchy window showing a Prefab instance with a child prefab instance called Banana added as an override.](../uploads/Main/prefab-instance-hierarchy-override.png)

The Hierarchy window showing a Prefab instance with a child prefab instance called Banana added as an override.

## Additional resources

* [Manage GameObjects in the Hierarchy window](Hierarchy.md)
* [Scene visibility](SceneVisibility.md)
* [Scene picking controls](ScenePickingControls.md)
* [Override prefab instances](PrefabInstanceOverrides.md)