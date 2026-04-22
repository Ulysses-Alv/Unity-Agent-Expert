# 2D renderer sorting

## Overview

Unity sorts Renderers according to a priority order that depends on their types and usages. You can specify the render order of Renderers through their [Render Queue](SL-SubShaderTags.md). In general, there are two main queues: the [Opaque queue](../ScriptReference/Rendering.RenderQueue.Geometry.md) and the [Transparent queue](../ScriptReference/Rendering.RenderQueue.Transparent.md). 2D Renderers are mainly within the Transparent queue, and include the [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md), [Tilemap Renderer](tilemaps/work-with-tilemaps/tilemap-renderer-reference.md), and [Sprite Shape Renderer](sprite/shape-renderer/shape-renderer-landing) types.

## Transparent Queue Sorting Order by Priority

2D Renderers within the Transparent Queue generally follow the priority order below:

1. [Sorting Layer and Order in Layer](#sortlayer)
2. [Specify Render Queue](#renderqueue)
3. [Distance to Camera](#cameradistance)

   * [Perspective](#persp)
   * [Orthographic](#ortho)
   * [Custom Axis sort mode](#customaxis)
   * [Sprite Sort Point](#sortpoint)
4. [Sorting Group](#sortgroup)
5. [Material/Shader](#material)
6. [Tiebreaker](#tiebreak)

There are other factors which can cause the sorting order to differ from the regular priority order. These factors vary from project to project.

## Sorting Layer and Order in Layer

The [Sorting Layer](https://docs.unity3d.com/Manual/class-TagManager.html#SortingLayers) and **Order in Layer** (in the Renderer’s **Property** settings) are available to all 2D Renderers through the **Inspector** window or via the Unity Scripting API. Set the Renderer to an existing **Sorting Layer** or create a new one to determine its priority in the rendering queue. Change the value of the **Order in Layer** to set the Renderer’s priority among other Renderers within the same **Sorting Layer**.

## Specify Render Queue

You can specify the Render Queue type of the Renderer in its Material settings or in the **Shader** settings of its Material. This is useful for grouping and sorting Renderers which are using different Materials. Refer to documentation on [ShaderLab: SubShader Tags](SL-SubShaderTags.md) for more details.

## Distance to Camera

The [Camera component](class-Camera.md) sorts Renderers based on its **Projection** setting. The two options are **Perspective** and **Orthographic**.

### Perspective

In this mode, the sorting distance of a Renderer is the direct distance of the Renderer from the **Camera**’s position.

### Orthographic

The sorting distance of a Renderer is the distance between the position of the Renderer and the Camera along the Camera’s view direction. For the default 2D setting, this is along the (0, 0, 1) axis.

When you set the Camera component to **Perspective** or **Orthographic,** Unity automatically sets the Camera’s [TransparencySortMode](../ScriptReference/TransparencySortMode.md) to match the selected mode. You can set the Transparency Sort Mode manually in two ways:

* Open the ****Project Settings**** and go to [Graphics](class-GraphicsSettings.md#Camera), then under **Camera Settings** use Transparent Sort Mode
* Set the Camera’s [TransparencySortMode](../ScriptReference/TransparencySortMode.md) via the Scripting API.

The Camera **Transparency Sort Mode** settings are under the **Graphics** category in the **Project Settings** (main menu: **Edit** > **Project Settings** > **Graphics**). When this is set to **Default**, a Camera component’s Projection setting take priority. When this is set to an option other than **Default**, the Camera component’s Projection setting remains the same, but the Camera’s **Transparency Sort Mode** changes to that option.

An additional option available through the Project settings and via the Scripting API is the [Custom Axis sort mode](../ScriptReference/TransparencySortMode.CustomAxis.md).

### Custom Axis sort mode

Select this mode to sort Renderers based on their distance along the custom axis you set in the Project settings (main menu: **Edit** > **Project Settings** > **Graphics** > **Transparency Sort Axis**). This is commonly used in projects with [Isometric Tilemaps](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md) to sort and render the Tile Sprites correctly on the Tilemap. Refer to [Creating an Isometric Tilemap](tilemaps/work-with-tilemaps/isometric-tilemaps/create-isometric-tilemap.md) for further information.

**Note:** If your project is a 2D Universal **Render Pipeline** (URP) project, the **Transparency Sort Mode** isn’t available in the **Project Settings**, and is instead configured in the 2D Renderer asset properties. Refer to the [Configure the 2D Renderer Asset documentation](urp/2DRendererData-overview.md) for more information.

### Sprite Sort Point

By default, a **Sprite**’s **[Sort Point](sprite/renderer/set-sort-point-sprite.md)** is set to its **Center**, and Unity measures the distance between the camera’s Transform position and the Center of the Sprite to determine their render order during sorting. An alternate option is to set a Sprite’s **Sort Point** to its **Pivot** position in World Space. Select the **Pivot** option in the Sprite’s [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md) property settings and edit the Sprite’s Pivot position in the [Sprite Editor tab of the Sprite Editor window](sprite/sprite-editor/sprite-editor-window-reference.md).

## Sorting Group

The [Sorting Group](sprite/sorting-group/sorting-group-landing.md) is a component that groups Renderers which share a common root together for sorting purposes. All Renderers within the same Sorting Group share the same **Sorting Layer, Order in Layer,** and **Distance to Camera**. Refer to documentation on the [Sorting Group](sprite/sorting-group/sorting-group-landing.md) component and its settings for more details.

## Material/Shader

Unity sorts Renderers with the same [Material settings](Shaders.md) together for more efficient rendering performance, such as with [dynamic batching](DrawCallBatching.md).

## Tiebreaker

When multiple Renderers have identical sorting priority, the tiebreaker is the order that Unity places the Renderers in the Render Queue. Because this is an internal process that you have no control over, you should use the sorting options (such as **Sorting Layers** and **Sorting Groups**) to make sure all Renderers have distinct sorting priorities.