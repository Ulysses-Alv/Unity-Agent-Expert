# Scene view context menu

Use the **Scene** view context menu to access common **Scene view** actions directly in the scene rather than from the menu **toolbar**.

The menu options that display in the Scene view context menu depend on your current selection and the tool context enabled in the Tools overlay. If you’ve selected a **GameObject**, the Scene view context menu displays options for that GameObject and any relevant attached components.

To display the context menu, right-click in the Scene view.

**Note**: You can select a different shortcut to display the Scene view context menu in the [Shortcuts Manager](ShortcutsManager.md).

The default tool context in the Scene view is **GameObject**. When the **GameObject** tool context is enabled, the Scene view context menu displays the following menu items by default when you select a GameObject:

* Clipboard actions such as cut, copy, paste, delete, and duplicate.
* GameObject view options to help you visualize the selected GameObject in the Scene view.
* The option to [Isolate selected GameObjects](SceneVisibility.md#isolate-go) in the Scene view so that only the selected GameObject is visible.
* The option to a add a [component](UsingComponents.md) to the selected GameObject.
* The option to open the [properties](InspectorManageComponents.md) of the selected in GameObject in a new window.
* The [Prefab](Prefabs.md) menu.
* The [Transform](class-Transform.md) menu.

If the GameObject has additional components attached to it, actions related to those components display at the end of the Scene view context menu.

If your project contains multiple tool contexts, you can use the first button in the Tools overlay to select a tool context. If you enable a tool context that isn’t **GameObject**, the Scene view context menu displays options relevant to your selection in that tool context. For example, if your project contains the [Splines](https://docs.unity3d.com/Packages/com.unity.splines@latest) package and you enable the **Splines** tool context, then the Scene view context menu displays options to create and edit splines.

## Additional resources

* [Using the Scene view](UsingTheSceneView.md)
* [Scene view navigation](SceneViewNavigation.md)
* [Keyboard shortcuts](ShortcutsManager.md)