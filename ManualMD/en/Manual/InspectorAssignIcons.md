# Assign icons to inspected items

Use the ****Inspector**** window to assign an icon to the item you’re inspecting. The icon appears in the ****Scene**** view, making it easier to identify the item.

The icon’s behavior in the [Scene view](UsingTheSceneView.md) depends on the item type:

* **GameObject**: The icon appears over that GameObject, and any duplicates.
* **Prefab**: The icon appears over any instances of that prefab.
* Script: The icon appears over any GameObject that has the script attached.

To control how Unity draws custom icons in the **Scene** view, use the [Gizmos menu](GizmosMenu.md).

**Note:** When you change an asset’s icon, Unity marks the asset as modified, and **version control** systems register the change.

### Icon types

The Unity Editor supports three icons types.

| **Icon type** | **Description** |
| --- | --- |
| **Label icons** | Colored capsules that show the item’s name. |
| **Image only icons** | Colored circles. They’re useful for objects that don’t have a visual representation such as waypoints. |
| **Custom icons** (other) | An icon based on an asset in the project. For example, use a skull and crossbones icon to indicate danger areas in a game level. |

## Manage icons in the **Inspector** window

To add icons, in the **Inspector** window, select **Select icon**. Then:

* To assign an icon to a GameObject or a script, do one of the following:
  + To assign a label or image icon, click any icon from the list.
  + To assign a custom icon, select **Other…** and select a texture.
* To assign an icon to a prefab, open the prefab in [prefab editing mode](EditingInPrefabMode.md).

To replace or remove icons, in the **Inspector** window, select **Select icon**. Then:

* To replace the icon, select any other icon from the list.
* To remove the icon, select **None**. The item’s icon reverts to the default Unity icon.

## Additional resources

* [Using the Scene view](UsingTheSceneView.md)
* [Gizmos menu](GizmosMenu.md)
* [Edit prefab assets](EditingInPrefabMode.md)