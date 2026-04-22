# Render queues and sorting behaviours

The order in which Unity renders objects is based on two things: which render queue the object is in, and how Unity sorts objects within that render queue.

## Render queues

Unity sorts objects into groups called render queues. Unity renders the contents of one render queue, and then renders the contents of another render queue, and so on.

Unity has the following named render queues, which it renders in the following order:

| **Name** | **Index** | **Description** |
| --- | --- | --- |
| **Background** | 1000 | Use this queue for anything that should be drawn in the background of your **Scene**. |
| **Geometry** | 2000 | Use this queue for opaque geometry. This is the default queue. |
| **AlphaTest** | 2450 | Use this queue for alpha tested geometry. This is after the **Geometry** queue because it’s more efficient to render alpha-tested objects after all solid ones are drawn. |
| **Transparent** | 3000 | Use this queue for anything alpha-blended (shaders that don’t write to the depth buffer). Examples include glass, or **particle** effects. |
| **Overlay** | 4000 | Use this queue for effects that are rendered on top of everything else, such as **lens flares**. |

Note that **Skybox** materials are a special case. Unity draws Skybox materials after all opaque geometry (after queue index 2500), but before all transparent geometry (before queue index 2501).

## Sorting behaviors within render queues

Within each render queue, Unity sorts and draws objects based on their distance from the **camera**. The sorting behavior depends on the index of the render queue:

* For queues with an index up to and including 2500 (Geometry + 500), Unity sorts Renderers in these queues using the behavior defined in [OpaqueSortMode.FrontToBack](../ScriptReference/OpaqueSortMode.FrontToBack.md) by default.
* For queues with an index of 2501 or above, Unity sorts Renderers in these queues using the behavior defined in [TransparencySortMode.Default](../ScriptReference/TransparencySortMode.Default.md) by default.

## Set how a camera sorts materials

How you change the sorting behavior within a render queue depends on the index of the render queue:

* For queues with an index up to and including 2500 (Geometry + 500), you can change the opaque sort mode for a Camera by using the [Camera.opaqueSortMode](../ScriptReference/Camera-opaqueSortMode.md) API.
* For queues with an index of 2501 or above, you can change the default transparent sort mode by using the [Rendering.GraphicsSettings-transparencySortMode](../ScriptReference/Rendering.GraphicsSettings-transparencySortMode.md) API. You can change the transparent sort mode for a Camera by using the [Camera.transparencySortMode](../ScriptReference/Camera-transparencySortMode.md) API.