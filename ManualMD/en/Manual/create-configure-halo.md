# Create and configure a halo light effect

**Note:** This workflow is compatible only with the Built-in **Render Pipeline**. For similar functionality in other render pipelines, see [Lens flares and halos](visual-effects-lens-flares.md).

Add a Halo component to a **GameObject** with a [Light component](class-Light.md). Then set its size and color properties to give the desired glowing effect.

**Note:** A Light can also display a halo without a separate Halo component by enabling its *Draw Halo* property, but you cannot configure its size and color.

To see Halos in the **scene** view, check **Fx** button in the **Scene View** **Toolbar**.

To override the **shader** used for Halos, open the [Graphics](class-GraphicsSettings.md) window and set **Light Halo** to the shader that you would like to use as the override.