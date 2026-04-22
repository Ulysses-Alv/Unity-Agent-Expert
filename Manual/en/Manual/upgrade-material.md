# Upgrade material assets to URP or HDRP

When you upgrade your project from the Built-In **Render Pipeline** (BiRP) to either the Universal Render Pipeline (URP) or the High Definition Render Pipeline (HDRP), you need to upgrade your materials, otherwise, the materials appear bright pink in **Scene** view.

![A bright pink cube in Scene view.](../uploads/Main/shader-error.png)

A bright pink cube in Scene view.

**Notes:**

* Make sure there are no shader-related errors in the console, or in the ****Inspector**** window when you select a material.
* If your assets use custom **shaders**, refer to [Upgrade custom shaders for URP compatibility](urp/urp-shaders/birp-urp-custom-shader-upgrade-guide.md).

## Upgrade BiRP materials to URP

To upgrade all material assets from BiRP to URP:

1. Back up your Built-in Render Pipeline material assets.
2. In your URP project, go to **Window** > **Rendering** > **Render Pipeline Converter**.
3. In the **Render Pipeline Converter** window, select **Built-in to URP** from the drop-down, enable **Material Upgrade**, then select **Initialize and Convert**.

To upgrade only some material assets from BiRP to URP:

1. Back up your Built-in Render Pipeline material assets.
2. In your URP project, select your Built-in Render Pipeline material assets.
3. Go to **Edit** > **Rendering** > **Materials** > **Convert Selected Built-in Materials to URP**.

**Note:** If the console or the **Inspector** window displays [error messages](shader-error.md) when you select a material, there’s an issue with a shader that an automatic converter can’t solve.

For more information, refer to [Convert assets using the Render Pipeline Converter](urp/features/rp-converter.md).

## Upgrade BiRP materials to HDRP

To upgrade all material assets from BiRP to HDRP:

1. Back up your Built-in Render Pipeline material assets.
2. In your HDRP project, go to **Edit** > **Rendering** > **Materials** > **Convert All Built-in Materials to HDRP**.

To upgrade only some material assets from BiRP to HDRP:

1. Back up your Built-in Render Pipeline material assets.
2. In your HDRP project, select your Built-in Render Pipeline material assets in the **Project window**.
3. Go to **Edit** > **Rendering** > **Materials** > **Convert Selected Built-in Materials to HDRP**.

For more information, refer to [Convert materials and shaders (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/convert-from-built-in-convert-materials-and-shaders.html).

## Additional resources

* [Upgrading from the Built-In Render Pipeline to URP](urp/upgrading-from-birp.md)
* [Upgrading to HDRP from the built-in render pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/convert-project-from-built-in-render-pipeline.html)