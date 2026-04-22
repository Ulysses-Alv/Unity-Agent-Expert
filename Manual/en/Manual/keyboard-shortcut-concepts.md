# Keyboard shortcut concepts

Understand the concepts behind keyboard shortcuts in the Unity Editor.

This page explains the difference between global and contextual commands. It describes the three types of commands and provides information about shortcut profiles.

For information about how to view, assign, and manage shortcuts, refer to [Shortcuts window reference](shortcuts-view.md) and [Manage keyboard shortcuts](shortcuts-manage.md).

## Global commands, contextual commands, and conflicts

Editor shortcuts can be global or contextual.

* Global shortcuts are always available. For example, by default the command to Undo an action uses the shortcut **Ctrl/Cmd + Z**. Using that shortcut always undoes the most recent action, regardless of which windows are open and which tools are active.
* Contextual shortcuts only work when you use a particular view or set of tools. For example, by default the square bracket keys **[** and **]** have different uses in the **Terrain** tool and the Grid Painting tool. When you use the bracket shortcut, Unity executes the command for whichever tool is active.

You can assign any keyboard shortcut to either:

* One global command.
* One or more contextual commands from different contexts.

### Context-based conflicts

You get a shortcut conflict if you assign the same keyboard shortcut to:

* More than one global command.
* More than one contextual command from the same context.
* A global command and a contextual command.

When you create a conflict, the Editor displays a warning icon in the **Shortcuts** window. You can resolve the conflict by changing one of the conflicting shortcuts.

### Exceptions to context-based conflicts

The Editor has a few global commands that can share shortcuts without creating conflict, because something other than context differentiates them. For example, the **Q**, **W**, and **E** keys:

* In the ****Scene**** and **Game** views, use these keys to toggle the View, Move, and Rotate tools.
* In the **Flythrough** mode, use these keys to move down, forward, and up.

## Command types

There are three types of commands in Unity:

* Action: Triggers once, when you press the shortcut.
* Clutch: Triggers twice, when you press the shortcut and when you release it.
* Menu: Activates a main menu option.

  **Tip:** You can also access menu items from the **Search** window. For more information, refer to [Search the Unity Editor main menu](search-menus.md).

## Profiles

The Editor saves keyboard shortcuts to profiles. You can create multiple profiles and move between them without restarting the Editor. With this, you can customize your shortcuts for different workflows, package sets, or custom tools. For more information, refer to [Manage shortcut profiles](shortcut-profiles.md).

## Additional sources

* [Shortcuts window reference](shortcuts-view.md)
* [Manage shortcuts](shortcuts-manage.md)
* [Manage shortcut profiles](shortcut-profiles.md)