# Create a shader variant collection

A **shader** variant collection is effectively a list of [shader variants](shader-variants.md). Use shader variant collections to [prewarm shader variants](shader-loading.md), or to ensure that shader variants that are required at runtime but not referenced in a **scene** are not excluded (“stripped”) from your build.

## Create a shader variant collection asset

You can create a shader variant collection asset in the following ways:

* In the Create Asset menu, choose **Shader** > **Shader Variant Collection**.
* The Unity Editor can track which shader variants your application uses when it runs, and automatically create a shader variant collection asset that contains them. For more information, see [Graphics Settings: Shader loading](class-GraphicsSettings.md#shader-loading).

## View and edit a shader variant collection

![Shader variant collection inspector](../uploads/Main/ShaderVariantCollection.png)

Shader variant collection inspector

When you select a shader variant collection asset in your project, you can view and edit it in the **Inspector**.

Use the controls to build a list of [Pass types](../ScriptReference/Rendering.PassType.md) and [shader keyword](shader-keywords.md) combinations to load in advance.

You can also configure a shader variant collection asset using the [ShaderVariantCollection](../ScriptReference/ShaderVariantCollection.md) API.

## Prewarm a shader variant collection

To avoid visible stalls at performance-intensive times, Unity can ask the graphics driver to create GPU representations of shader variants before they are first needed. This is called **prewarming**. For more information on prewarming the shader variants in a shader variant collection, see [Shader loading: Prewarming shader variants](shader-prewarm.md).