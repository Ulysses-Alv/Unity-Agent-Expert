# Rendering paths in Unity

A **rendering path** is the series of operations that draws and lights **GameObjects** the **camera** sees. Different rendering paths have different capabilities and performance characteristics.

Unity supports forward and deferred rendering paths.

## Forward

The **Forward rendering** path is the default rendering path in Unity projects. It works in the following way:

* Unity lights each GameObject in turn.
* Lighting has limits, for example how often or how well Unity can light each GameObject. These limits are different for each **render pipeline**.

The Universal Render Pipeline (URP) also has a [Forward+ rendering path](urp/rendering/forward-rendering-paths.md), which is similar to the Forward rendering path but has no limit on lights per GameObject.

**Note:** The Built-In Render Pipeline also has a [Legacy Vertex Lit](RenderTech-VertexLit.md) rendering path, which is a subset of the Forward rendering path.

## Deferred

The Deferred rendering path works in the following way:

* Unity first creates a geometry buffer (G-buffer), which is a set of textures that stores data about the geometry and materials the camera sees.
* Unity uses the data from the G-buffer to light all the GameObjects at once.
* There are fewer limits on lighting, so GameObjects and shadows are more detailed. For example, **normal maps** and cookies work with all lights.

Deferred rendering paths can’t render transparent objects, so at the end of the rendering path Unity uses a forward render pass to render transparent objects.

## Choose a rendering path

Each rendering path has advantages and disadvantages. For more information, refer to the following:

* [Introduction to rendering paths in URP](urp/rendering-paths-introduction-urp.md)
* [Introduction to rendering paths in the Built-In Render Pipeline](built-in-rendering-paths.md)

## Additional resources

* [Understanding rendering paths](https://learn.unity.com/tutorial/understanding-rendering-paths) on the Unity Learn site
* [Per-pixel and per-vertex lights](PerPixelLights.md)