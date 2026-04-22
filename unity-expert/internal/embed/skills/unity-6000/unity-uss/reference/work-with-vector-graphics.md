# Work with vector graphics

Vector graphics are scalable, resolution-independent graphics that you can use in your UI. They are defined by mathematical equations rather than **pixels**, which allows them to maintain their quality at any size.

## Import and use vector graphics

You can [import](../UIE-image-import-settings.md) vector graphics to use in your UI.

When you import a vector image, set its **Generated Asset Type** to **UI Toolkit Vector Image** in the ****Inspector**** to get the best results. This creates a specialized asset designed exclusively for use within UI Toolkit, such as in the UI Builder or for runtime UI.

Alternately, you can set its **Generated Asset Type** to [**Texture2D**](../../ScriptReference/Texture2D.md) for use in other contexts.

For more information, refer to [Vector images import settings](../UIE-image-import-settings.md#vector-images) and [Set background image](../UIB-styling-ui-backgrounds.md).

## Limitations

When you work with vector graphics, be aware of the following limitations:

* Avoid starting Unity with the `-no-graphics` command-line argument, as this might prevent SVG vector images from importing correctly.
* The SVG importer supports a subset of the [SVG 1.1 specification](https://www.w3.org/TR/SVG11/). The following features aren’t currently supported:

  + [Text elements](https://www.w3.org/TR/SVG11/text.html)
  + [Per-pixel masking](https://www.w3.org/TR/SVG11/masking.html#Masking)
  + [Filter effects](https://www.w3.org/TR/SVG11/filters.html)
  + [Interactivity features](https://www.w3.org/TR/SVG11/interact.html)
  + [Animations](https://www.w3.org/TR/SVG11/animate.html)

## Create vector graphics with C# scripts

You can create and manipulate vector graphics with C# **scripts** in your UI. Unity provides a set of classes and methods that let you:

* Create vector graphics elements, such as paths and shapes.
* Apply fills and gradients to shapes.
* Clip content using other shapes as clippers.
* Tessellate and render vector graphics to meshes, **sprites**, or textures for display in your UI.

There are two main kind of drawable instances: paths and shapes.

### Paths

Paths are vector shapes defined by a sequence of cubic Bézier curves, grouped into a [`BezierContour`](../../ScriptReference/Unity.VectorGraphics.BezierContour.md). Each contour consists of an array of path segments and a flag indicating whether the contour is closed. You can use these to create complex, continuous curves for your UI graphics.

Each path segment specifies a start point and two control points. The next segment’s start point continues the curve, allowing you to chain multiple segments for smooth, continuous shapes. You need at least two segments for a valid contour.

For example, consider this path:

![Contour](../../uploads/Main/uss/contour.png)

Contour

To create a path, define a list of segments and specify whether to close the contour. You can then assign stroke and fill properties to style the path as needed.

```
var a = new Vector2(0, 0);
var b = new Vector2(100, 0);
var c = new Vector2(100, 100);
var d = new Vector2(0, 100);
var e = new Vector2(50, 50);
var f = new Vector2(150, 50);
var g = new Vector2(200, 0);

var segments = new BezierPathSegment[] {
    new BezierPathSegment() { P0 = a, P1 = b, P2 = c },
    new BezierPathSegment() { P0 = d, P1 = e, P2 = f },
    new BezierPathSegment() { P0 = g }
};
var path = new Shape()
{
    Contours = new BezierContour[]
    {
        new BezierContour()
        {
            Segments = segments,
            Closed = true
        }
    },
    PathProps = new PathProperties()
    {
        Stroke = new Stroke() { Color = Color.red, HalfThickness = 1.0f }
    }
};
```

When you define a `BezierContour` with `Closed = true`, the contour connects the last path segment to the first path segment. The last path segment uses its `P1` and `P2` values as control points.

### Shapes

Shapes are similar to paths, but can contain multiple contours and support filling. You can fill shapes with solid colors, textures, or gradients, and apply transformations to the fill. This allows you to create complex, filled vector graphics for your UI.

You can use different fill types:

* [SolidFill](../../ScriptReference/Unity.VectorGraphics.SolidFill.md)
* [TextureFill](../../ScriptReference/Unity.VectorGraphics.TextureFill.md)
* [GradientFill](../../ScriptReference/Unity.VectorGraphics.GradientFill.md)

### Gradients

Gradient fills let you blend multiple colors together, either in a straight line (linear) or radiating from a point (radial). You define a gradient by specifying a type and a series of color stops, each with a color and a percentage along the gradient.

For example, you can create a linear gradient that transitions from blue to red to yellow by defining stops at 0%, 50%, and 100%:

```
var gradient = new GradientFill()
{
    Type = GradientFillType.Linear,
    Stops = new GradientStop[] {
        new GradientStop() { Color = Color.blue, StopPercentage = 0.0f },
        new GradientStop() { Color = Color.red, StopPercentage = 0.5f },
        new GradientStop() { Color = Color.yellow, StopPercentage = 1.0f }
    }
};
```

![Linear fill](../../uploads/Main/uss/addressing.png)

Linear fill

The gradient [addressing modes](../../ScriptReference/Unity.VectorGraphics.AddressMode.md) define how Unity displays the color when the gradient coordinates fall outside of the range, as illustrated here:

![Addressing modes](../../uploads/Main/uss/addressing.png)

Addressing modes

### Fill Mode

The fill mode determines how the interior of complex shapes and holes are calculated. The main fill modes are:

* [`NonZero`](../../ScriptReference/Unity.VectorGraphics.FillMode.NonZero.md): Determines whether a point is inside a shape by counting how many times the path contours wind around the point. If the result is nonzero, you consider the point inside. The direction of each contour (clockwise or counterclockwise) affects the result.

  ![NonZero fill](../../uploads/Main/uss/fill_nonzero.png)

  NonZero fill
* [`OddEven`](../../ScriptReference/Unity.VectorGraphics.FillMode.OddEven.md): Counts how many times a horizontal line from the point crosses the shape’s contours. If the number of crossings is odd, the point is inside; if even, it’s outside. This mode is often used for shapes with holes.

  ![EvenOdd fill](../../uploads/Main/uss/fill_evenodd.png)

  EvenOdd fill

Choose the fill mode that best matches the desired appearance of your vector shapes, especially when working with overlapping paths or holes.

### Clipping

The [`SceneNode`](../../ScriptReference/Unity.VectorGraphics.SceneNode.md) class includes a [`Clipper`](../../ScriptReference/Unity.VectorGraphics.SceneNode.Clipper.md) property that allows you to clip the contents of a node with a shape.

![Clipper](../../uploads/Main/uss/clipper.png)

Clipper

The following code example creates an ellipse that clips a pattern of repeating squares:

```
// Create an ellipse shape to use as the clipper
var ellipseShape = new Shape();
VectorUtils.MakeEllipseShape(ellipseShape, new Vector2(0, 0), 50, 100);

var ellipseClipper = new SceneNode
{
    Shapes = new List<Shape> { ellipseShape }
};

// Create a pattern for clipping, for example, repeating squares
var squaresPattern = new SceneNode
{
    Shapes = new List<Shape>() // Add your fill pattern here
};

// Create a SceneNode that applies the clipper to its children
var clippedNode = new SceneNode
{
    Children = new List<SceneNode> { squaresPattern },
    Clipper = ellipseClipper
};
```

**Note**: Only use shapes as clippers. The clipping process ignores any strokes that you define on the clipper shape. You can clip any combination of shapes and strokes as content.

**Warning:** Clipping can be computationally expensive. While clipping simple shapes with a simple clipper is usually efficient, using complex shapes or clippers might significantly impact performance.

### Render vector graphics

To render vector graphics elements on screen, first get a tessellated (triangulated) version of the **scene**. When you have a VectorScene instance set up, you can tessellate it with the following [`VectorUtils`](../../ScriptReference/Unity.VectorGraphics.VectorUtils.md) method:

```
var geometry = VectorUtils.TessellateScene(
    scene,
    new VectorUtils.TessellationOptions()
    {
        StepDistance = 10.0f,
        MaxCordDeviation = 0.1f,
        MaxTanAngleDeviation= 0.1f
    });
```

**Note**: You must specify `maxTanAngleDeviation` in radians.

To disable the `maxCordDeviation` constraint, set it to `float.MaxValue`.
To disable the `maxTanAngleDeviation` constraint, set it to `Mathf.PI/2.0f`.

Disabling the constraints makes the tessellation faster, but might generate more vertices.

The resulting list of `Geometry` objects contains all the vertices and accompanying information required to render the scene properly.

### Textures and gradients atlases

If the scene has any textures or gradients, you must generate a texture atlas and fill the UVs of the geometry. These methods are part of the `VectorUtils` class:

```
var atlas = VectorUtils.GenerateAtlasAndFillUVs(geoms, 16);
```

The [`GenerateAtlas`](../../ScriptReference/Unity.VectorGraphics.VectorUtils.GenerateAtlas.md) method is an expensive operation, so cache the resulting Texture2D object whenever possible. You only need to regenerate the atlas when a texture or gradient changes inside the scene.

When vertices change inside the geometry, call the [`FillUVs`](../../ScriptReference/Unity.VectorGraphics.VectorUtils.FillUVs.md) method, which is cheap.

### Draw a tessellated scene

You can render the geometry in several ways. For example, you can render the geometry by:

* Filling a `Mesh` asset
* Building a `Sprite` asset
* Using Unity’s low-level graphics library

For any of these methods, use the provided materials to draw the tessellated vector graphics content.

To fill a **mesh** asset, use the following `VectorUtils` method:

```
var mesh = new Mesh();
VectorUtils.FillMesh(
    mesh, 
    geoms,
    100.0f, // svgPixelsPerUnit
    false); // flipYAxis
```

To build a sprite asset, use the following `VectorUtils` method:

```
var sprite = VectorUtils.BuildSprite(
    geoms,
    100.0f,                       // svgPixelsPerUnit
    VectorUtils.Alignment.Center, // alignment
    Vector2.zero,                 // pivot
    16,                           // gradient resolution
    false);                       // flipYAxis
```

To render a sprite to a [`Texture2D`](../../ScriptReference/Texture2D.md), use the following `VectorUtils` method:

```
var tex = VectorUtils.RenderSpriteToTexture2D(
    sprite,          // The sprite to draw
    256, 256,       // The texture dimensions
    null,           // The material to use ("Unlit_Vector" or "Unlit_VectorGradient")
    1);             // The number of samples per pixel
```

To render the generated sprite using immediate mode `GL` commands, use the [`RenderSprite`](../../ScriptReference/Unity.VectorGraphics.VectorUtils.RenderSprite.md) method in the `VectorUtils` class to draw the sprite into a unit square (a box between 0 and 1 in both X and Y directions):

```
// Or "Hidden/VectorGraphics/VectorGradient" for gradients
var mat = new Material(Shader.Find("Hidden/VectorGraphics/Vector"));

VectorUtils.RenderSprite(
    sprite,  // The sprite to draw
    mat);    // The material to use ("Unlit_Vector" or "Unlit_VectorGradient")
```

## Additional resources

* 📺 **Video**: [UI Toolkit Tutorial Series: SVG Support](https://www.youtube.com/watch?v=cDN-VW3McfQ)
* [Image import settings](../UIE-image-import-settings.md)
* [`SVGImporter`](../../ScriptReference/Unity.VectorGraphics.Editor.SVGImporter.md)
* [Vector Graphics package documentation](https://docs.unity3d.com/Packages/com.unity.vectorgraphics@latest)