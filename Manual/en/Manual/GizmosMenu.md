# Gizmos menu reference

The [Scene view](UsingTheSceneView.md) and the [Game view](GameView.md) both have a **Gizmos** menu. Click the **Gizmos** button in the **toolbar** of the **Scene** view or the Game view to access the **Gizmos** menu.

![The Gizmos button in the Scene view](../uploads/Main/gizmo-button-scene-view.png)

The **Gizmos** button in the Scene view

![The Gizmos button in the Game view](../uploads/Main/gizmo-button-game-view.png)

The **Gizmos** button in the Game view

## General options

| **Property** | **Function** |
| --- | --- |
| **3D Icons** | The **3D Icons** checkbox controls whether component icons (such as those for Lights and Cameras) are drawn by the Editor in 3D in the Scene view.  When the **3D Icons** checkbox is ticked, component icons are scaled by the Editor according to their distance from the Camera, and obscured by GameObjects in the Scene. Use the slider to control their apparent overall size. When the **3D Icons** checkbox is not ticked, component icons are drawn at a fixed size, and are always drawn on top of any GameObjects in the Scene view.  Refer to [Gizmos introduction](gizmos-introduction.md) for more information. |
| **Fade Gizmos** | Fade out and stop rendering gizmos that are small on screen. |
| **Selection Outline** | Check **Selection Outline** to display selected GameObjects with a colored outline, and their children with a different colored outline. By default, Unity highlights the selected GameObject in orange, and child GameObjects in blue.  **NOTE:** This option is only available in the Scene view Gizmos menu; you cannot enable it in the Game view Gizmos menu.   Refer to [Selection Outline and Selection Wire](gizmos-introduction.md#SelectionOutlineWire), for more information. |
| **Selection Wire** | Check **Selection Wire** to show the selected GameObjects with their wireframe Mesh visible. To change the colour of the Selection Wire, go to **Edit > Preferences** (macOS: **Unity > Settings**) in the main menu, select **Colors**, and alter the **Selected Wireframe** setting.  This option is only available in the Scene view Gizmos menu; you cannot enable it in the Game view Gizmos menu.  Refer to [Selection Outline and Selection Wire](gizmos-introduction.md#SelectionOutlineWire), for more information. |
| **Display Weights** | When enabled, Unity draws a line from the Light Probe used for the active selection to the positions on the tetrahedra used for interpolation. This is a way to debug probe interpolation and placement problems. |
| **Display Occlusion** | When enabled, Unity displays occlusion data for Light Probes if the **Lighting Mode** is set to **Shadowmask**. |
| **Highlight Invalid Cells** | Enable to highlight tetrahedrons that Unity can’t use for indirect lighting. For example, if the Light Probes are very close together. |
| **Recently Changed** | Controls the visibility of the icons and gizmos for components and scripts that you modified recently.  This section appears the first time you change one or more items, and updates after subsequent changes.  For more information, refer to [Built-in components, scripts, and recently changed items](gizmos-introduction.md#Components). |
| **Scripts** | Controls the visibility of the icons and gizmos for scripts in the Scene.  This section appears only when your Scene uses scripts that meet specific criteria.  For more information, refer to [Built-in components, scripts, and recently changed items](gizmos-introduction.md#Components) below. |
| **Built-in Components** | Controls the visibility of the icons and gizmos for all component types that have an icon or gizmo.  For more information, refer to [Built-in components, scripts, and recently changed items](gizmos-introduction.md#Components). |

## Light Probe Visualization options

Select which [Light Probes](LightProbes.md) to display in the Scene view. The default value is **Only Probes Used By Selection**.

| **Property** | **Function** |
| --- | --- |
| **Only Probes Used By Selection** | Display only Light Probes that affect the current selection. |
| **All Probes No Cells** | Display all Light Probes. |
| **All Probes With Cells** | Display all Light Probes, and the tetrahedrons that Unity uses for interpolation of Light Probe data. |
| **None** | Display no Light Probes. |

## Additional resources

* [Gizmos introduction](gizmos-introduction.md)
* [Programming with gizmos and handles](gizmos-handles-programming.md)