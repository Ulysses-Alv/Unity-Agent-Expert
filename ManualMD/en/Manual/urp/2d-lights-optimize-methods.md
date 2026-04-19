# Optimize 2D lights

In addition to the standard optimization techniques such as reducing draw calls, culling and optimizing **shaders**, there are several techniques and considerations that are unique to the 2D Lighting graphics pipeline.

## Number of blend styles

The easiest way to increase rendering performance is to reduce the number of blend styles used. Each blend style is a **Render Texture** that needs to be rendered and subsequently uploaded.

Reducing the number of blend styles has a direct impact on the performance. For simple scenes a single blend style could suffice. It is also common to use up to 2 blend styles in a **scene**.

## Light Render Texture scale

The 2D Lighting system relies on screen space Light Render Texture to capture light contribution. This means there are a lot of Render Texture drawing subsequent uploading. Choosing the right Render Texture size directly impacts the performance.

By default it is set at 0.5x of screen resolution. Smaller Light Render Texture size will give better performance at the cost of visual artifact. Half screen size resolution provides a good performance with almost no noticeable artifact in most situations.

Experiment and find a scale suitable for your project.

## Layer Batching

To further reduce the number of Light Render Textures, it is crucial to make the Sorting Layer batchable. Layers that are batched together share the same set of Light Render Textures. Uniquely lit layers will have its own set thus increasing the amount of work needed.

Layers can be batch together if they share the same set of lights.

## Pre-rendering of Light Render Texture

Multiple sets of Light Render Textures can be rendered ahead of drawing the Renderers. In an ideal situation, all the Light Render Textures will be rendered upfront and only then will the pipeline proceed with drawing the Renderers onto the final color output. This reduces the need to load/unload/reload of final color output.

In a very complex setup with many distinctly lit Layers, it may not be practical to pre-render all Light Render Textures. The limit can be configured in the 2D Renderer Data **inspector**.

## Normal Maps

Using **normal maps** to simulate depth is currently a very expensive operation. If it is enabled, a full size Render Texture is created during a depth pre-pass and the Renderers are drawn onto it. This is done for each Layer batch.

If the normal mapping effect to simulate depth perception is not needed, ensure that all lights have the normal map option disabled.