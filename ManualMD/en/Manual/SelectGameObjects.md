# Select GameObjects

You can select a single **GameObject** in the [Scene view](UsingTheSceneView.md) or from the [Hierarchy window](Hierarchy.md). You can also select more than one GameObject at a time.

The Editor highlights selected GameObjects and their children in the Scene view. By default, the selection outline color is orange, and the child outline color is blue. You can also choose to highlight selected the wireframes of selected GameObjects in a different color. You can change all of these outline highlight colors from the [Preferences window](Preferences.md). To open the **Preferences** window, go to **Edit > Preferences** (macOS: **Unity > Settings**) in the main menu.

If you are working with a large **Scene** that contains a lot of Scene items, such as GameObjects, **Terrain** objects, **Cameras**, and Lights, selecting multiple objects can be difficult. To help you select only the items you want, use the Scene picking controls to block some objects from being picked or use the [selection piercing menu](SelectionPiercingMenu.md).

For more information about the outline and wireframe selection visualizations, refer to [Gizmos menu](GizmosMenu.md).

To select a single GameObject:

* Click on it in the [Scene view](UsingTheSceneView.md). If you repeatedly click on the shared space between overlapping GameObjects, the selection cycles between them. You can also use the [selection piercing menu](SelectionPiercingMenu.md) to select a specific GameObject when multiple GameObjects are in the same location.
* Click on the name of the GameObject in the [Hierarchy window](Hierarchy.md).

To select or deselect multiple GameObjects:

* Drag a rectangle around multiple GameObjects to selects anything that is inside this bounding box.
* Hold the **Shift** key and select GameObjects in the Scene. You can also use hold **Ctrl** (macOS: **Command**) and select a GameObject to add or remove GameObjects from the selection.

**Note:** When a feature in the Editor needs a single selected GameObject to perform an action, it looks for an “active” object. For example, when the Editor has to decide which GameObject to use as the pivot for transform tools while in Pivot mode. By default, Unity considers the last GameObject you select to be the “active” object. When you hold **Shift** and select one of several selected GameObjects, you change which one of them is active. If a GameObject is active in the Scene view, Unity doesn’t display any visible cues that it’s active. However, you can see which GameObject is active in the Scene view when you repeatedly hold **Shift** and click in [Pivot mode](PositioningGameObjects.md#GizmoHandlePositions) with multiple objects selected.

## Additional resources

* [Pick and select GameObjects](ScenePicking.md)
* [Hierarchy window](Hierarchy.md)
* [Select a GameObject that overlaps with other GameObjects](SelectionPiercingMenu.md)\*\*
* [Gizmos menu](GizmosMenu.md)]