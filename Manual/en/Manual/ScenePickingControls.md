# Scene picking controls

You can toggle the **scene** picking status of a **GameObject** to choose which GameObjects are selectable in the **Scene view**. By default, all GameObjects are pickable, which means you can select them in the Scene view. You can disable the pickability of a GameObject to prevent that GameObject from being added to selections in the Scene view. For example, if you are working in a large scene with many GameObjects, you can temporarily block specific GameObjects from being selected to prevent them from being edited accidentally.

A GameObject that has its scene picking disabled is still rendered in the Scene view, but you can’t select or modify it in the Scene view. The pickability state persists only in the Editor and only for the user in the project where the pickability was set. Changing the pickability state of a GameObject is not recorded as a modification to the scene.

Scene picking controls are similar to the [Scene visibility](SceneVisibility.md) controls.

## Toggle the pickability of a GameObject

Control scene picking status for individual GameObjects from the [**Hierarchy** window](Hierarchy.md).

![Every GameObject has a Scene pickability icon/toggle](../uploads/Main/scene-picking-icons-hierarchy.png)

Every GameObject has a Scene pickability icon/toggle

To toggle the scene picking status of a GameObject:

* Select a GameObject’s pickability icon in the [**Hierarchy** window](Hierarchy.md) to toggle between enabling and disabling picking the GameObject and its children. Toggling pickability for a GameObject and its children affects all child GameObjects, from the target GameObject to the bottom of the hierarchy.
* For multiple GameObjects, select them all in the **Hierarchy** window, then press **L** to toggle their pickability.
* Select a GameObject in the Scene view, and press **L** to toggle the pickability of the GameObject and its children.
* Hold **Alt** (macOS: **Option**) and select a GameObject’s pickability icon in the [**Hierarchy** window](Hierarchy.md) to toggle between enabling and disabling picking the GameObject only. Toggling pickability for a single GameObject does not affect its children. They retain whatever pickability status they had previously.

## Pickability icons

You can toggle pickability for a whole branch of GameObjects or a single GameObject, so you can have GameObjects that are selectable but have children or parents that you can’t select. To help you track what is pickable or not, the pickability icon changes to show the status of each GameObject.

Any scene picking changes you make in the **Hierarchy** window are persistent. Scene picking changes are reapplied whenever you use **Select All**, **Deselect All**, or close and re-open the Scene.

![Unity only selects pickable items when you draw a selection bounding box in the Scene view](../uploads/Main/scene-picking-icons-ovw.png)

Unity only selects pickable items when you draw a selection bounding box in the Scene view

| Screenshot annotation | Pickability icon | Description |
| --- | --- | --- |
| **A** | Pick | You can pick the GameObject, but you can’t pick some of its children. |
| **B** | Cant pick | You can’t pick the GameObject, but you can pick some of its children. |
| **C** | Cant pick | You can’t pick the GameObject or its children. |
| **D** | Pick | You can pick the GameObject and its children. This icon only appears when you hover over the GameObject. |

## Additional resources

* [Pick and select GameObjects](ScenePicking.md)
* [Scene visibility](SceneVisibility.md)
* [Select GameObjects](SelectGameObjects.md)
* [Select a GameObject that shares space with other GameObjects](SelectionPiercingMenu.md)