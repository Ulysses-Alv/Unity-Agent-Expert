# Manage keyboard shortcuts

Assign, remove, and reset keyboard shortcuts; manage conflicts; and create shortcuts for custom tools, **scripts**, and packages in the Unity Editor.

## Assign shortcuts to commands

Assign shortcuts to commands with the key map or the table view.

To assign a shortcut with the key map:

1. Find the command you want to assign a shortcut to in the table. To make this easier, you can search or select a category from the list.
2. If you want to add a modifier key (for example, **Shift**), in the key map, select that key.
3. Drag the command from the table onto the shortcut key, for example, **A**.

To assign a shortcut with the table:

1. Find the command you want to assign a shortcut to in the table. To make this easier, you can search or select a category from the list.
2. Double-click the command in the table.
3. With the cursor in the shortcut field, press the shortcut keys, including any modifier key you want to include (for example, **Shift** or **Ctrl/Cmd**). The shortcut field updates to show the keys you pressed.

   ![The Shortcut field shows Shift + 5 for the selected command](../uploads/Main/ShortcutManagerListAssign.png)

   The Shortcut field shows **Shift** + **5** for the selected command

If the shortcut is already assigned to a command, the new assignment creates a conflict. For information about managing conflicts, refer to [Manage shortcut conflicts](#manage-shortcut-conflicts).

## Remove or reset shortcuts

Use the key map or the table to remove a command’s assigned shortcut or reset a modified shortcut to its default mapping.

**Note:** If you want to reset most or all of your shortcuts, create a new profile from the default profile. For more information, refer to [Manage shortcut profiles](#shortcut-profiles).

To remove or reset a shortcut:

1. Right-click the shortcut you want to remove or reset, in either the key map or the table. To remove or reset a key combination with the key map, select any relevant modifier keys (**Shift** or **Ctrl/Cmd**, for example).

   A context menu opens, with all the commands assigned to that key. Each command has its own options to remove or reset the shortcut.
2. From the context menu of the command, select one of the following options:

   * To replace your shortcut with the default shortcut, select **Reset to default**. If the command has no default shortcut, this option leaves the shortcut field empty.
   * To remove the shortcut without returning to the default shortcut, select **Remove shortcut**.

## Manage shortcut conflicts

When you assign a shortcut to a command, it can conflict with other commands that are already assigned to the same shortcut. This can also happen if you install a new package or custom tool with its own shortcuts.

When you create a new shortcut that conflicts with an existing one, the Editor displays the following options:

* **Reassign:** Reassign the shortcut to the new command, and remove it from conflicting commands.
* **Create Conflict:** Assign the shortcut to the new command, and don’t remove it from existing commands. All the commands using this shortcut are listed in the **Binding Conflicts** category in the **Shortcuts** window and have a caution icon in their table entry.
* **Cancel:** Don’t assign the shortcut to the new command, and exit the dialog.

**Tip:** When you use a shortcut with conflicting commands, the Editor displays a similar dialog but with the added option to select which command to execute.

![The Caution icon at the end of a command line](../uploads/Main/ShortcutManagerConflictCat.png)

The Caution icon at the end of a command line

To view all conflicting commands, from the **Category** list, select **Binding Conflicts**. You can then resolve the conflicts by reassigning or removing the conflicting shortcuts.

## Add shortcuts for custom tools, scripts, and packages

When you create custom tools, you can add commands to the main menu and assign shortcuts to them. You can also add categories to the **Category** list in the **Shortcuts** window.

How you create these shortcuts depends on the system you use to create the interface for your custom tools.

* For the [IMGUI](GUIScriptingGuide.md) system, refer to [Event.current](../ScriptReference/Event-current.md).
* For the [UI Elements](UIElements.md) system, refer to:
  + [KeyDownEvent](../ScriptReference/UIElements.KeyDownEvent.md)
  + [CallbackEventHandler.RegisterCallback](../ScriptReference/UIElements.CallbackEventHandler.RegisterCallback.md)
  + [CallbackEventHandler.HandleEvent](../ScriptReference/UIElements.CallbackEventHandler.HandleEvent.md)

To define custom shortcuts for actions that aren’t from custom tools, such as scripts and packages, use the classes in the `UnityEditor.ShortcutManagement` namespace, such as [ShortcutAttribute](../ScriptReference/ShortcutManagement.ShortcutAttribute.md) and [ClutchShortcutAttribute](../ScriptReference/ShortcutManagement.ClutchShortcutAttribute.md).

When the Editor loads the commands and shortcuts from another source, it:

* Adds any new commands in the main menu to the **All Unity Commands** category in the commands table.
* Adds custom command categories to the **Category** list.
* Adds custom keyboard shortcuts to the shortcuts database, so they appear in all shortcut profiles.
* Shows any conflicts between the custom shortcuts and existing shortcuts in the **Shortcut Conflicts** category.

## Additional sources

* [Keyboard shortcut concepts](keyboard-shortcut-concepts.md)
* [Shortcuts window reference](shortcuts-view.md)
* [Manage shortcut profiles](shortcut-profiles.md)
* [Apple touchpad gestures](SceneViewNavigation.md)