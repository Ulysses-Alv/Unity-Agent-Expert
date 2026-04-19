# Introduction to rendering paths in URP

You can select one of the following **rendering paths** in the Universal **Render Pipeline** (URP):

* Forward
* Forward+
* Deferred
* Deferred+

Each rendering path affects how Unity draws and lights objects, which affects lighting results and rendering time. The effects depend on the platform you build for.

For more information about choosing a rendering path, refer to [Choose a rendering path in URP](rendering-paths-comparison.md).

## Rendering path requirements in URP

| **Feature** | **Forward and Forward+** | **Deferred and Deferred+** |
| --- | --- | --- |
| Minimum **shader** model | 2.0 | 4.5 |
| OpenGL and OpenGL ES support | Yes | No |

## Additional resources

* [Understanding Rendering Paths](https://learn.unity.com/tutorial/understanding-rendering-paths-2019-3) on the Unity Learn site
* [Unity LTS 2022 Release Live!](https://www.youtube.com/watch?v=oUQapNQgpRI&t=8183s) - a Unity YouTube video that demonstrates the Forward+ rendering path