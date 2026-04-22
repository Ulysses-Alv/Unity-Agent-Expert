# Enabling Streaming Virtual Texturing in your project

To enable Streaming Virtual Texturing, you need to enable **Virtual Texturing** in your project. To do this, go to **Edit** > **Project Settings** > **Player** > **Other Settings** and enable the **Virtual Texturing** setting.

Virtual Texturing is a project-wide setting that is shared for all platforms. You cannot build a player for platforms and graphics APIs that don’t support Virtual Texturing. Virtual Texturing can allocate resources such as buffers, even if you don’t use the feature in your project, so don’t enable **Virtual Texturing** if you don’t plan to use it.

When you enable **Virtual Texturing** in your project, Unity adds the following compiler directives:

* `ENABLE_VIRTUALTEXTURES`: C# define that evaluates to `True` if your project supports Virtual Texturing.
* `ENABLE_VIRTUALTEXTURING`: C++ and C# editor define that evaluates to `True` if Virtual Texturing is possible on the current build target.
* `UNITY_VIRTUAL_TEXTURING`: **Shader** version of `ENABLE_VIRTUALTEXTURES`.

Note that these might be renamed in future versions of Unity.