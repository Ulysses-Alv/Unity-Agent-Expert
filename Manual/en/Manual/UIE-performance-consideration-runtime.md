# Performance consideration for runtime UI

This section describes how you can improve the performance for runtime UI.

| **Topic** | **Description** |
| --- | --- |
| [Use usage hints to reduce draw calls and geometry regeneration](UIE-use-usage-hints-to-reduce-draw-calls-and-geometry-regeneration.md) | Use usage hints to set how an element is used at runtime. Usage hints can reduce draw calls and avoid geometry regeneration. |
| [Control textures of the dynamic atlas](UIE-control-textures-of-the-dynamic-atlas.md) | Use an atlas to reduce the number of batches broken by texture changes and achieve good performance. |
| [Platform and mesh consideration](UIE-platform-and-mesh.md) | Consider device capabilities and avoid **mesh** tessellation. |

## Additional resources

* [Panel Settings asset](UIE-Runtime-Panel-Settings.md)
* [`usageHints`](../ScriptReference/UIElements.VisualElement-usageHints.md)