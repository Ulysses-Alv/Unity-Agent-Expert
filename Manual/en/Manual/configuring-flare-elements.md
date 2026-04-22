# Configuring Flare elements

Learn how Unity manages elements on a Flare asset, and compare texture layout options.

A Flare consists of multiple **Elements**, arranged along a line. The line is calculated by comparing the position of the **GameObject** containing the Lens Flare to the center of the screen. The line extends beyond the containing GameObject and the screen center. All Flare **Elements** are strung out on this line.

For performance reasons, all **Elements** of one Flare must share the same Texture. This Texture contains a collection of the different images that are available as Elements in a single Flare. The **Texture Layout** defines how the **Elements** are laid out in the **Flare Texture**. If you use many different Flare assets, using a shared single **Flare Texture** that contains all the **Elements** will give you best rendering performance.

Lens Flares are blocked by **Colliders**. A Collider in-between the Flare GameObject and the **Camera** will hide the Flare, even if the Collider does not have a **Mesh Renderer**. If the in-between Collider is marked as Trigger it will block the flare if and only if **Physics.queriesHitTriggers** is true.

To override the **shader** used for Flares, open the [Graphics](class-GraphicsSettings.md) window and set **Lens Flares** to the shader that you would like to use as the override.

## Understand texture layouts

These are the options you have for different Flare **Texture Layouts**. The numbers in the images correspond to the **Image Index** property for each **Element**.

### 1 Large 4 Small

Designed for large sun-style Flares where you need one of the **Elements** to have a higher fidelity than the others. This is designed to be used with Textures that are twice as high as they are wide.

![The large particle (element 0) is the top half of the texture. The 4 small particles are in a 2 x 2 layout in the bottom half of the texture. The element values increment from left to right in each row.](../uploads/Main/FlaresLayout0.png)

The large particle (element 0) is the top half of the texture. The 4 small particles are in a 2 x 2 layout in the bottom half of the texture. The element values increment from left to right in each row.

### 1 Large 2 Medium 8 Small

Designed for complex flares that require 1 high-definition, 2 medium and 8 small images. This is used in the standard assets “50mm Zoom Flare” where the two medium Elements are the rainbow-colored circles. This is designed to be used with textures that are twice as high as they are wide.

![The large particle (element 0) is the top half of the texture. In the bottom half, the 2 medium particles (1 and 2) are on the left side, with element 1 on top. The 8 small particles (starting with element 3) are in a 2 x 4 pattern on the right side. For the small particles, the element values increment from left to right in each row.](../uploads/Main/FlaresLayout1.png)

The large particle (element 0) is the top half of the texture. In the bottom half, the 2 medium particles (1 and 2) are on the left side, with element 1 on top. The 8 small particles (starting with element 3) are in a 2 x 4 pattern on the right side. For the small particles, the element values increment from left to right in each row.

### 1 Texture

A single image.

![A single particle texture.](../uploads/Main/FlaresLayout2.png)

A single particle texture.

### 2x2 grid

A simple 2x2 grid.

![A 2 x 2 grid of particles. The element values increment from left to right in each row.](../uploads/Main/FlaresLayout3.png)

A 2 x 2 grid of particles. The element values increment from left to right in each row.

### 3x3 grid

A simple 3x3 grid.

![A 3 x 3 grid of particles](../uploads/Main/FlaresLayout4.png)

A 3 x 3 grid of particles

### 4x4 grid

A simple 4x4 grid.

![A 4 x 4 grid of particles.](../uploads/Main/FlaresLayout5.png)

A 4 x 4 grid of particles.

Flare