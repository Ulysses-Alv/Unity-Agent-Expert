# Focus an Inspector window

By default, an **Inspector** window displays properties for the current selection, and its focus changes whenever the selection changes. To keep the focus on a single item or group of items, regardless of the current selection, do one of the following:

* Lock the **Inspector** window to the current selection. When you lock an **Inspector** window, it no longer updates if you change the selection.
* Open a focused **Inspector**, which only displays the properties of the item you opened it for. In other words, it’s an inherently locked **Inspector** window.

## Lock the Inspector

To lock an Inspector window to an item, click the lock icon in the **Inspector** window. The lock icon changes to show that the **Inspector** is now locked to the selection.

This is useful when you’re editing one item’s properties frequently or want to compare two items. For example, if you’re fine-tuning a **camera**’s position, you can lock an **Inspector** window to the camera, and open a second, unlocked **Inspector** that continues to match your selections.

## Use a focused Inspector

A focused **Inspector** window is dedicated to a single **GameObject**, component, or asset. It always displays the properties of the item you opened it for, even if you select something else in the **scene** or project.

You can open as many focused **Inspector** windows as you want. Focused **Inspector** windows open in floating windows that you can [reposition, dock, and resize](CustomizingYourWorkspace.md) like any other window.

**Tip:** The Unity Editor saves any open focused **Inspectors** when you close a project, and restores them when you reopen it.

### Open a focused Inspector window

How you open a focused Inspector window depends on what you want to inspect. You can open a focused Inspector for:

* GameObjects and project assets.
* Components and reference properties.
* Hovered over GameObjects in the **Hierarchy** window.

#### Open a focused Inspector window for a GameObject or an asset in your project

To open a focused **Inspector** for one or more GameObjects or assets in the **Hierarchy** window or the **Project** window, select the item or items you want to inspect and do one of the following:

* From the main menu, select **Assets** > **Properties**.
* Right-click the item. From the context menu, select **Properties**.
* Press **Alt+P** (macOS: **Option+Shift+P**).
* In the **Inspector** window, click the **More items** (**⋮**) menu and select **Properties**.

#### Open a focused Inspector window for a component or reference

A focused **Inspector** window is helpful for properties or references you edit frequently, especially if the GameObject has a lot of components. For example, if you often move a GameObject but rarely change its other properties, you can open its [Transform](class-Transform.md) component in a focused **Inspector** window.

To open a focused **Inspector** for a component or reference attached to a GameObject:

1. Inspect the GameObject and locate the component.
2. From the component’s **More items** (**⋮**) menu, select **Properties**.

#### Set a shortcut to open a focused Inspector window for a hovered-over GameObject

You can set up a shortcut to open a focused **Inspector** for the item your mouse hovers over in the **Hierarchy** window. This lets you open a focused **Inspector** for an item without affecting the current selection.

To assign a keyboard shortcut to the **PropertyEditor/OpenMouseOver** command, use the [Shortcuts Manager](ShortcutsManager.md).

## Locate a focused Inspector’s source

To locate the item whose properties are displayed in a focused **Inspector**, do one of the following:

* From the focused **Inspector** window’s **More Items** (**⋮**) menu, select **Ping** to highlight the item in the **Hierarchy** window.
* Hover over the focused **Inspector** window’s tab header to display a tooltip of the item’s full path.

## Additional resources

* [Customizing your workspace](CustomizingYourWorkspace.md)
* [Keyboard shortcuts](ShortcutsManager.md)