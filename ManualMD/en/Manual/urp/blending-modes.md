# Blending Modes in URP

In a material in the Universal **Render Pipeline** (URP), the **Blending Mode** property determines how Unity calculates the color of each **pixel** of a transparent Material when it blends the Material with the background.

In the context of Blending Modes, Source refers to the transparent Material where the Blending Mode is set and Destination refers to anything that Material overlaps with.

## Alpha

![Alpha blending mode example](../../uploads/urp/blend-modes/blend-mode-alpha.png)  
*Alpha blending mode.*

**Alpha** uses the Material’s alpha value to change how transparent an object is. 0 is fully transparent. 255 is fully opaque, this is translated to a value of 1 when used with the blending equations. The Material is always rendered in the Transparent render pass regardless of its alpha value.
This mode lets you use the Preserve Specular Lighting property.

Alpha equation:

OutputRGB = (SourceRGB \* SourceAlpha) + DestinationRGB \* (1 - SourceAlpha)

## Premultiply

![Premultiply blending mode example](../../uploads/urp/blend-modes/blend-mode-premultiply.png)  
*Premultiply blending mode.*

**Premultiply** helps reduce artifacts that might appear at the edge of the overlap between opaque and transparent pixels, by allowing opaque areas of the material to blend with the background.

**Premultiply** uses SourceRGB values that have already been multiplied by the alpha value. Otherwise, the blending equation is the same as **Alpha** mode.

Premultiply equation:

OutputRGB = SourceRGBPremultiplied + DestinationRGB \* (1 - SourceAlpha)

## Additive

![Additive blending mode example](../../uploads/urp/blend-modes/blend-mode-additive.png)  
*Additive blending mode.*

**Additive** adds the color values of the Materials together to create the blend effect. The alpha value determines the strength of the source Material’s color before the blend is calculated.
This mode lets you use the Preserve Specular Lighting property.

Additive equation:

OutputRGB = (SourceRGB \* SourceAlpha) + DestinationRGB

## Multiply

![Multiply blending mode example](../../uploads/urp/blend-modes/blend-mode-multiply.png)  
*Multiply blending mode.*

**Multiply** multiplies the color of the Material with the color behind the surface. This creates a darker effect, like when you look through colored glass.
This mode uses the Material’s alpha value to adjust how much the colors blend. An alpha value of 1 results in unadjusted multiplication of the colors while lower values blend the colors towards white.

Multiply equation:

OutputRGB = SourceRGB \* DestinationRGB

## Performance of blending modes

The approximate order of blending modes from fastest to slowest is the following:

* Additive. The calculation is simple.
* Multiply. The calculation is simple, but includes SourceAlpha.
* Premultiply. The calculation is more complicated than Additive and Multiply, but the GPU doesn’t need to multiply the SourceRGB value by the alpha value per pixel.
* Alpha. The calculation is the most complicated, especially if multiple Alpha Mode materials overlap.

The performance of your project might depend more on overdraw than the blending mode, especially if you target a mobile platform. For more information, refer to [Reduce rendering work on the CPU or GPU](../OptimizingGraphicsPerformance.md).