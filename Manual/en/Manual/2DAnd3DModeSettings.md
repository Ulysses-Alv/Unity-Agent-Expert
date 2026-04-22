# 2D and 3D mode settings

When creating a new Project, you can specify whether to start the Unity Editor in 2D mode or 3D mode. However, you also have the option of switching the Editor between 2D mode and 3D mode at any time. You can read more about [the difference between 2D and 3D Projects here](2Dor3D.md). This page provides information about how to switch modes, and what exactly changes within the editor when you do.

## Switching between 3D and 2D modes

To change modes between 2D or 3D mode:

1. Open the [Editor](class-EditorManager.md) settings (top menu: **Edit > Project Settings**, then select the **Editor** category).
2. Then set **Default Behavior Mode** to either **2D** or **3D**.

![Use the Default Behavior Mode setting in the Editor settings to set the Project to 2D or 3D](../uploads/Main/BehaviorMode.png)

Use the Default Behavior Mode setting in the Editor settings to set the Project to 2D or 3D

## 2D vs 3D mode settings

2D or 3D mode determines some settings for the Unity Editor. These are listed below.

### In 2D Project mode:

* Any images you import are assumed to be 2D images (**Sprites**) and set to **Sprite** mode.
* The **Scene View** is set to 2D.
* The default game objects do not have real time, directional light.
* The **camera**’s default position is at 0,0,–10. (It is 0,1,–10 in 3D Mode.)
* The camera is set to be **Orthographic**. (In 3D Mode it is **Perspective**.)
* In the Lighting Window:
  + **Skybox** is disabled for new **scenes**.
  + **Ambient Source** is set to **Color**. (With the color set as a dark grey: RGB: 54, 58, 66.)
  + **Realtime Global Illumination** (Enlighten) is set to off.
  + **Baked Global Illumination** is enabled.
  + **Auto-Building** set to off.

### In 3D Project mode:

* Any images you import are NOT assumed to be 2D images (**Sprites**).
* The **Scene View** is set to 3D.
* The default game objects have real time, directional light.
* The camera’s default position is at 0,1,–10. (It is 0,0,–10. in 2D Mode.)
* The camera is set to be **Perspective**. (In 2D Mode it is **Orthographic**.)
* In the Lighting Window:
  + **Skybox** is the built-in default **Skybox Material**.
  + **Ambient Source** is set to **Skybox**.
  + **Realtime Global Illumination** (Enlighten) is set to on.
  + **Baked Global Illumination** is set to on.
  + **Auto-Building** is set to on.