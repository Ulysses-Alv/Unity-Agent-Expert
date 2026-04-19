# Introduction to Scriptable Render Passes in URP

Scriptable Render Passes are a way to alter how Unity renders a **scene** or the objects within a scene. They allow you to fine tune how Unity renders each scene in your project on a scene-by-scene basis.

You inject a Scriptable Render Pass into the **render pipeline** to achieve a custom visual effect. For more information, refer to [Adding a Scriptable Render Pass to the frame rendering loop](../inject-a-render-pass.md).

A Scriptable Render Pass lets you to do the following:

* Change the properties of materials in your scene.
* Change the order that Unity renders **GameObjects** in.
* Lets Unity read **camera** buffers and use them in **shaders**.

For example, you can use a Scriptable Render Pass to blur a camera’s view when showing the in-game menu.

Unity injects Scriptable Render Passes at certain points during the URP render loop. These points are called injection points. You can change the injection point Unity inserts your pass at to control how the Scriptable Render Pass affects the appearance of your scene. For more information on injection points, refer to [Injection Points](../customize/custom-pass-injection-points.md).

## Additional resources

* [Adding pre-built effects with Renderer Features in URP](../urp-renderer-feature-landing.md)
* [Inject a render pass using a Scriptable Renderer Feature](../customize/inject-render-pass-via-script.md)