# Scene view navigation

By default, you look through and control the [Scene Camera](SceneViewCamera.md) when you navigate through the **Scene** view. To look through and control a **GameObject** that has a **camera** component attached to it, use the [Cameras overlay](cameras-overlay.md).

You can use the following navigation controls to move the Scene Camera or a [GameObject that has a camera component attached to it](control-camera.md) around the **Scene view**:

* The [Move, Orbit and Zoom](#tools) tools
* The [Center](#center) tool

### Mac trackpad gestures

On a Mac with a trackpad, you can drag with two fingers to zoom the view.

You can also use three fingers to simulate the effect of clicking the arms of the [Orientation overlay](overlay-orientation.md): drag up, left, right or down to snap the Scene Camera to the corresponding direction. You must enable three-finger swiping in the MacOS trackpad gesture settings to use this feature.

## Move, orbit and zoom in the Scene view

Moving, orbiting, and zooming are key operations in Scene view navigation. Unity provides several ways to perform them for maximum accessibility:

* [Use the arrow keys](#arrow)
* [Use the View tool](#handtool)
* [Use flythrough mode](#flythrough)
* [Change the move speed of the camera](#speed)
* [Use movement shortcuts](#shortcuts)

### Use the arrow keys

You can use the **Arrow Keys** to move around the Scene as though “walking” through it. The Up and Down arrow keys move the Camera forward and backward in the direction it faces. The Left and Right arrow keys pan the view sideways. Hold down the **Shift** and an arrow key to move faster.

### Use the View tool

When the View tool is selected (shortcut: **Q**), the following mouse controls are available:

| **Control** | **Description** |
| --- | --- |
| **Pan** | In View tool mode, the Move icon appears  Click and drag to pan the Camera around. |
| **Orbit** | When you hold down the Alt button in View tool mode, the Orbit icon appears   Hold **Alt** (Windows) or **Option** (macOS), and left-click and drag to orbit the Camera around the current pivot point.   This option isn’t available in 2D mode, because the view is orthographic.  **Note**: If you can’t orbit the Camera, make sure that the padlock on the top right of the [Orientation overlay](overlay-orientation.md) is disabled. |
| **Zoom** | Hold **Alt** (Windows) or **Option** (macOS), and right-click and drag to zoom the Scene view.   On macOS, you can also hold **Control**, and left-click and drag instead. |

Hold down **Shift** to increase the rate of movement and zooming.

### Use flythrough mode

Use the **Flythrough mode** to fly around the Scene view in first-person, similar to how you would navigate in many games.

Flythrough mode is designed for **Perspective Mode**. In **Orthographic Mode**, if you click and hold the right mouse button and move your mouse, your view orbits the Camera instead.

Flythrough mode isn’t available in 2D mode. In 2D mode, if you click and hold the right mouse button and move your mouse, your view pans around the Scene view.

To enter **Flythrough mode** and navigate through the Scene view in **Flythrough mode**:

1. Click and hold the right mouse button.
2. Do the following to navigate through the Scene view:
   * Use your mouse to move the view.
   * To move forward or backward, press **W** or **S**.
   * To move left or right, press **A** or **D**.
   * To move up or down, press **Q** or **E**.
   * To move faster, press and hold **Shift**.

### Change the move speed of the camera

The [Scene Camera](SceneViewCamera.md) displays the Scene view in the Editor. By default, the Scene Camera is what you control and look through when you navigate through the Scene view. To learn how to control a GameObject that has a camera component attached, refer to [Control a camera in first person](control-camera.md).

To change the speed that a Camera moves at in the Scene view, select the Camera icon in the View Options [overlay](overlays.md) then adjust the value of the **Camera Speed** property to the speed you want.

**Note**: To find the View Options overlay, press **`** to open the Overlays menu. In the Overlays menu, hover over **View Options** to highlight the View Options overlay in the Scene view.

In **Flythrough mode**, use the scroll wheel on your mouse or drag two fingers on a trackpad to change the speed that the Scene Camera moves at through the Scene.

### Movement shortcuts

For extra efficiency, these controls can also be used regardless of which transform tool is selected.
The most convenient controls depend on which mouse or track-pad you are using.

| **Action** | **3-button mouse** | **2-button mouse or track-pad** | **Mac with only one mouse button or track-pad** |
| --- | --- | --- | --- |
| **Pan** | Hold the middle-mouse button then drag | Hold **Alt**+**Control**+left-click, then drag | Hold **Option**+**Command**+left-click, then drag |
| **Orbit** (Not available in 2D mode) | Hold **Alt**+left-click, then drag | Hold **Alt**+left-click, then drag | Hold **Option**+left-click, then drag |
| **Zoom** | Use the scroll wheel, or hold **Alt**+right-click, then drag | Hold **Alt**+right-click, then drag | Use the two-finger swipe method to scroll in and out, or hold **Option**+**Control**+left-click, then drag |
| **Change speed** (only available in Flythrough mode) | Use the scroll wheel while moving. | Drag with two fingers while moving | Drag with two fingers while moving |

## Center the view on a GameObject

To center the Scene view on a GameObject, select the GameObject in the Hierarchy, then move the mouse over the Scene view and press **F**. If the GameObject is already selected, **F** zooms in to the pivot point. This feature can also be found in the menu bar under **Edit** > **Frame Selected**.

To lock the view to the GameObject even when the GameObject is moving, press **Shift+F**. This feature is also in the menu bar under **Edit** > **Lock View to Selected**.