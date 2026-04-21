# Shader Graph Window | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Shader-Graph-Window.html)

# Shader Graph Window

The **Shader Graph Window** contains the workspace to edit your shader graphs.

To access the **Shader Graph Window**, you must first create a [Shader Graph Asset](index.html). If the Shader Graph window doesn't open automatically after you [create a new shader graph asset](Create-Shader-Graph.html), you have to double-click on the created asset.

## Shader Graph window layout

![The Shader Graph window with its main elements labeled from A to F.](images/ShaderGraphWindow.png)

| Label | Name | Description |
| --- | --- | --- |
| **A** | [Toolbar](#toolbar) | A set of tools to manage the shader graph asset, display elements in the window, and more. |
| **B** | [Workspace](#workspace) | The area where you create your graph. |
| **C** | [Master Stack](Master-Stack.html) | The final connection that determines your shader output. It consists of two separate contexts: **Vertex** and **Fragment**. |
| **D** | [Main Preview](Main-Preview.html) | Previews the current shader output. Use this to rotate the object, and zoom in and out. You can also change the basic mesh on which the shader is previewed. |
| **E** | [Blackboard](Blackboard.html) | Contains all of the shader's properties and keywords in a single, collected view. Use the Blackboard to add, remove, rename, and reorder properties and keywords. |
| **F** | [Graph Inspector](Internal-Inspector.html) | Displays the properties of the currently selected component. Use this to modify properties, node options, and the graph settings. This window is hidden by default and only appears when something is selected that can be edited by the user. |

## Toolbar

Use the **Shader Graph Window** toolbar to manage the shader graph asset, display elements in the window, and more.

| Icon | Item | Description |
| --- | --- | --- |
| The Save Asset icon | **Save Asset** | Save the graph to update the [Shader Graph Asset](index.html). |
| The asset file management icon | **Asset file management** | Use additional options to manage the graph asset file. The options are:  * **Save As**: Save the [Shader Graph Asset](index.html) under a new name. * **Show In Project**: Highlight the [Shader Graph Asset](index.html) in the [Project Window](https://docs.unity3d.com/Manual/ProjectView.html). * **Check Out**: If version control is enabled, check out the [Shader Graph Asset](index.html) from the source control provider. |
| The Color Mode Selector icon | **Color Mode Selector** | Select a [Color Mode](Color-Modes.html) for the graph. |
| The Blackboard icon | **Blackboard** | Toggle the visibility of the [Blackboard](Blackboard.html). |
| The Graph Inspector icon | **Graph Inspector** | Toggle the visibility of the [Graph Inspector](Internal-Inspector.html). |
| The Main Preview icon | **Main Preview** | Toggle the visibility of the [Main Preview](Main-Preview.html). |
| The Help icon | **Help** | Open the Shader Graph documentation in the browser. |
| The Resources icon | **Resources** | Contains links to Shader Graph resources (like samples and User forums). |

## Workspace

Use the **Shader Graph Window** workspace to create [Node](Node.html) networks and connect them to the **Master Stack**.

To navigate the workspace, do the following:

* Press and hold the Alt key and drag with the left mouse button to pan.
* Use the mouse scroll wheel to zoom in and out.

You can hold the left mouse button and drag to select multiple [Nodes](Node.html) with a marquee. There are also various [shortcut keys](Keyboard-shortcuts.html) you can use for better workflow.

### Context Menu

Right-click in the workspace area, on a node, or on a selection of nodes to open a context menu.

| Item | Description |
| --- | --- |
| **Create Node** | Opens the [Create Node Menu](Create-Node-Menu.html). |
| **Create Sticky Note** | Creates a new [Sticky Note](Sticky-Notes.html) on the Graph. |
| **Collapse All Previews** | Collapses previews on all nodes. |
| **Cut** | Removes the selected nodes from the graph and places them in the clipboard. |
| **Copy** | Copies the selected nodes to the clipboard. |
| **Paste** | Pastes the nodes from the clipboard. |
| **Delete** | Deletes the selected nodes. |
| **Duplicate** | Duplicates the selected nodes. |
| **Select** > **Unused Nodes** | Selects all nodes on the graph that are not contributing to the final shader output from the [Master Stack](Master-Stack.html). |
| **View** > **Collapse Ports** | Collapses unused ports on all selected nodes. |
| **View** > **Expand Ports** | Expands unused ports on all selected nodes. |
| **View** > **Collapse Previews** | Collapses previews on all selected nodes. |
| **View** > **Expand Previews** | Expands previews on all selected nodes. |
| **Precision** > **Inherit** | Sets the precision of all selected nodes to Inherit. |
| **Precision** > **Float** | Sets the precision of all selected nodes to Float. |
| **Precision** > **Half** | Sets the precision of all selected nodes to Half. |

## Additional resources

* [Color Modes](Color-Modes.html)
* [Create Node Menu](Create-Node-Menu.html)
* [Keyboard shortcuts](Keyboard-shortcuts.html)
* [Master Stack](Master-Stack.html)
* [Nodes](Node.html)
* [Sticky Notes](Sticky-Notes.html)