# Install Shader Graph | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/install-shader-graph.html)

# Install Shader Graph

Learn about the Shader Graph installation requirements and follow the installation instructions.

## Requirements

Use Shader Graph with either of the Scriptable Render Pipelines (SRPs) available in Unity version 2018.1 and later:

* The [High Definition Render Pipeline (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest)
* The [Universal Render Pipeline (URP)](https://docs.unity3d.com/Manual/urp/urp-introduction.html)

As of Unity version 2021.2, you can also use Shader Graph with the [Built-In Render Pipeline](https://docs.unity3d.com/Documentation/Manual/built-in-render-pipeline.html).

##### Important

The Built-In Render Pipeline is deprecated and will be made obsolete in a future release.  
It remains supported, including bug fixes and maintenance, through the full Unity 6.7 LTS lifecycle.  
For more information on migration, refer to [Migrating from the Built-In Render Pipeline to URP](https://docs.unity3d.com/6000.5/Documentation/Manual/urp/upgrading-from-birp.html) and [Render pipeline feature comparison](https://docs.unity3d.com/6000.5/Documentation/Manual/render-pipelines-feature-comparison.html).

##### Note

* Shader Graph support for the Built-In Render Pipeline is for compatibility purposes only. Shader Graph doesn't receive updates for Built-In Render Pipeline support, aside from bug fixes for existing features. It's recommended to use Shader Graph with the Scriptable Render Pipelines.
* In the Built-In Render Pipeline, Shader Graph doesn't support XR.
* In the Built-In Render Pipeline, GPU Instancing doesn't work with Shader Graph shaders.
* In URP and HDRP, GPU instancing works with custom shaders only if you disable the [Scriptable Render Pipeline (SRP) Batcher](SRPBatcher) or [make a shader incompatible with the SRP Batcher](SRPBatcher-Incompatible).

## Installation

When you install HDRP or URP into your project, Unity also installs the Shader Graph package automatically. You can manually install Shader Graph for use with the Built-In Render Pipeline on Unity version 2021.2 and later with the Package Manager.

* For more information about how to install a package, see [Adding and removing packages](https://docs.unity3d.com/Manual/upm-ui-actions.html) in the Unity User Manual.
* For more information about how to set up a Scriptable Render Pipeline, see [Getting started with HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/getting-started-in-hdrp.html) or [Getting started with URP](https://docs.unity3d.com/Manual/urp/InstallingAndConfiguringURP).