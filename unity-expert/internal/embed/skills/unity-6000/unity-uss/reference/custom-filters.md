# Create and apply custom filters

You can create custom filters to apply any effect you want to a **visual element**.

**Important**: You can’t use UI **Shader** Graph to create custom filters. Custom filters are **post-processing** shaders that apply effects to render targets containing subtrees of visual elements. They are fundamentally different from shaders created with [UI Shader Graph](ui-shader-graph.md), which render UI **mesh** elements directly.

## Create a custom filter

To create a custom filter, follow these steps:

1. Create a `FilterFunctionDefinition` asset. This asset is a simple `ScriptableObject`.
2. Define filter parameters to set input values for the filter. The supported filter parameter types are `float` and `color`. When specified in USS, filter parameters such as numeric, **pixels**, percentages, durations, and angles are converted into float values. Parameters can be set in style or directly in the UI Builder. For example, the `blur(5px)` filter function has only one parameter: the blur radius. This [custom swirl filter function](create-custom-swirl-filter.md) has two parameters, the `angle` and the `radius` of the swirl.
3. Provide an `interpolation default value` (optional). This value is used for smooth transitions between filter states.
4. Add effect passes. Filters often need multiple passes to apply an effect. For example, a blur effect typically has separate horizontal and vertical passes. Each effect pass in the filter includes:

   * The material to execute: In Unity, shaders and materials work together to define how surfaces look. A shader is a program that tells the GPU how to render graphics and determines the effects on surfaces. Materials store the settings that feed into the shader.
   * The pass index for that material.
5. Use data binding to link the filter’s parameters to the material properties. For example, to apply a tint effect, bind a color parameter directly to the material’s `_Tint` property. For more complex parameter handling, [use a C# callback](#complex-parameter).
6. Set margins (optional). Some effects, such as blur, require extra space around the texture for reading and writing. Define static margins if needed, or [use a C# callback to compute margins dynamically](#dynamic-margin).

For an example, refer to [Create a custom swirl filter](create-custom-swirl-filter.md).

## Create a custom filter in C# code

To create a custom filter in C#, define a class that inherits from [`FilterFunction`](../../ScriptReference/UIElements.FilterFunction.md) and add the filter’s parameters and effects. The following creates a custom filter in C# for the [swirl filter example](create-custom-swirl-filter.md):

```
var customSwirl = ScriptableObject.CreateInstance<FilterFunctionDefinition>();

customSwirl.parameters = new[] {
    new FilterParameterDeclaration() {
        interpolationDefaultValue = new FilterParameter(0.0f)
    },
    new FilterParameterDeclaration() {
        interpolationDefaultValue = new FilterParameter(0.5f)
    }
};

customSwirl.passes = new[]
{
    new PostProcessingPass()
    {
        material = SwirlMaterial,
        passIndex = 0,
        parameterBindings = new[] {
            new ParameterBinding() { index = 0, name = "_Angle" },
            new ParameterBinding() { index = 1, name = "_Radius" }
        }
    }
};
```

## Apply the custom filter in stylesheets

To apply your custom filter to a visual element, use the filter property in a USS file. Reference the path to the custom filter asset, followed by the parameters needed for the filter. For example:

```
.custom-effect {
    filter: filter("/Assets/Filters/CustomFilter.asset" 10px);
}
```

You can combine custom filters with built-in filters in the same filter property. This allows you to create complex visual effects by stacking filters. For example:

```
.custom-effect-blur {
    filter: filter("/Assets/Filters/CustomFilter.asset" 10px) blur(25px);
}
```

## Apply custom filters in C# code

The following example shows how to apply a blur filter and a custom filter to a visual element in C#:

```
// Create a blur filter using the built-in blur filter function
var blur = new FilterFunction(FilterFunctionType.BuiltinBlur);
blur.AddParameter(new FilterParameter(20.0f));

// Create a custom filter function definition
var custom = new FilterFunction(functionDef);
// Add parameters to the custom filter
custom.AddParameter(new FilterParameter(Color.black));

// Apply both the blur and custom filter to the visual element
element.style.filter = new List<FilterFunction> { blur, custom };
```

## Dynamic margin

You can use a C# callback to compute margins dynamically for a filter.

The following C# code snippet sets dynamic margins for a blur function. Since the extra margins depend on the blur radius, static margins aren’t sufficient in this case.

Before calculating the **render texture** size, call this callback. This allows you to specify margins for the given `FilterFunction`. In this example, the blur radius is read as the function’s first parameter, which is then used to set the margins.

```
effect.computeRequiredReadMarginsCallback = ComputeMargins;
effect.computeRequiredWriteMarginsCallback = ComputeMargins;

// ...

static PostProcessingMargins ComputeMargins(FilterFunction func)
{
    float radius = func.GetParameter(0).floatValue; // Blur radius value
    return new PostProcessingMargins() {
        left = radius, top = radius, right = radius, bottom = radius
    };
}
```

## Complex parameter

For complex parameters, use a C# callback. The following example demonstrates how to use a callback to handle a filter parameter. It converts a single sepia amount value into a full color matrix. By implementing a callback, you can manually set material properties instead of relying on automatic parameter binding. To achieve this, provide a `MaterialPropertyBlock` to the callback, which then sets the material properties directly, bypassing the automatic binding process.

```
effect.prepareMaterialPropertyBlockCallback = PrepareSepiaMatrix;

// ...

static void PrepareSepiaMatrix(MaterialPropertyBlock mpb, FilterFunction func)
{
    float s = func.parameters[0].floatValue; // Sepia value

    var colorMatrix = new Matrix4x4(
        new Vector4(0.393f + 0.607f * (1 - s), 0.349f - 0.349f * (1 - s), 0.272f - 0.272f * (1 - s), 0),
        new Vector4(0.769f - 0.769f * (1 - s), 0.686f + 0.314f * (1 - s), 0.534f - 0.534f * (1 - s), 0),
        new Vector4(0.189f - 0.189f * (1 - s), 0.168f - 0.168f * (1 - s), 0.131f + 0.869f * (1 - s), 0),
        new Vector4(0, 0, 0, 1));

    mpb.SetMatrix("_ColorMatrix", colorMatrix);
}
```

## Custom filter UV region

UI Toolkit renders filtered elements into a texture atlas before applying filter shaders. As a result, the corresponding UV coordinates cover only a subregion of the full `0–1` range. When drawing the filters, the UVs provided to the fragment shader are adjusted to fit the UV region. In simple cases, it is sufficient to sample the texture with the provided UV coordinate.

However, if your custom filter requires UV manipulations, you need to convert the UVs to the `0-1` range and then back to the atlas UV rectangle. To do the conversion, first extract the filter UV rect index from the vertex data with the `GetFilterRectIndex` function:

```
#include "UnityUIEFilter.cginc"

v2f vert (FilterVertexInput v)
{
    v2f o;
    o.vertex = UnityObjectToClipPos(v.vertex);
    o.uv = TRANSFORM_TEX(v.uv, _MainTex);
    o.rectIndex = GetFilterRectIndex(v); // Pass filter UV rect index to fragment shader
    return o;
}
```

In the vertex or fragment shader, you can extract the UV rect with the `GetFilterUVRect` function, and then remap the UVs to and from the `0-1` range:

```
float2 NormalizeUVs(float2 uv, float4 uvRect)
{
    // Normalize UV coordinates based on the atlas rect
    return float2(
        (uv.x - uvRect.x) / uvRect.z,
        (uv.y - uvRect.y) / uvRect.w
    );
}

float2 MapToUVRect(float2 uv, float4 uvRect)
{
    // Map UV coordinates to the atlas rect
    return float2(
        uv.x * uvRect.z + uvRect.x,
        uv.y * uvRect.w + uvRect.y
    );
}

half4 frag (v2f i) : SV_Target
{
    // Get the UV rect from the index
    float4 uvRect = GetFilterUVRect(i.rectIndex);

    // Convert the atlas UVs to the 0-1 range
    float2 uv = NormalizeUVs(i.uv, uvRect);

    // Do UVs manipulations, (e.g. flip Y axis)
    uv.y = 1.0f - uv.y;

    // Convert the UVs back to the atlas region
    uv = MapToUVRect(uv, uvRect);

    return tex2D(_MainTex, uv);
}
```

**Note**: The texture passed into `_MainTex` when you create a custom filter is pre-multiplied. This means that the RGB values have already been multiplied by the alpha channel, as the texture is the direct output from the previous pass.

## Custom filter and force gamma considerations

When working in a UI Toolkit force gamma workflow in a [linear color space](../linear-color-space-landing.md) (when you enable **Force Gamma Rendering** in the [Panel Settings](../UIE-Runtime-Panel-Settings.md) asset), you need to take extra considerations with the color output of custom filters. In this workflow, the last filter applied must linearize the color before output.

To enable the `UIE_OUTPUT_LINEAR` keyword functionality, include a multi-compile instruction in your shader:

```
#pragma multi_compile _ _UIE_OUTPUT_LINEAR
```

UI Toolkit automatically enables the `UIE_OUTPUT_LINEAR` keyword on the material when the last filter is applied, allowing you to check for this keyword to determine when to linearize the color. Note that this is only necessary with the force gamma workflow. For Editor UI in linear projects, force gamma is implicitly enabled.

The following example shows how to handle force gamma considerations in a custom filter shader:

```
half4 frag(v2f IN) : SV_Target
{
    half4 col = tex2D(_MainTex, IN.uv);

    // Apply filter effect.

    // Linearize the color for the force gamma workflow.
    #if UIE_OUTPUT_LINEAR
    col.rgb = GammaToLinearSpace(col.rgb);
    #endif

    return col;
}
```

The following table summarizes how custom filters behave in different project color space configurations when reading from and writing to render targets:

| **Project configuration** | **Filter position** | **Input color space** | **Output color space** | **Special handling required** |
| --- | --- | --- | --- | --- |
| Gamma project | Any | Gamma | Gamma | None. Shaders work in gamma space throughout |
| Linear project without force gamma | Any | Linear | Linear | None. Shaders work in linear space throughout |
| Linear project with force gamma | Non-final passes | Gamma | Gamma | None. Intermediate passes maintain gamma |
| Linear project with force gamma | Final pass | Gamma | Linear | Convert output to linear using `GammaToLinearSpace` |

**Note**: When using `MaterialPropertyBlock.SetColor` to set color parameters in custom filters within a linear project that has force gamma enabled, the colors are provided in linear space while your filter shader operates in gamma space (as shown in the previous table). To maintain color accuracy, convert these color parameters to gamma space in your shader:

```
// Convert color parameter from linear to gamma space in shader.
fixed4 colorParam = _ColorParameter;
colorParam.rgb = LinearToGammaSpace(colorParam.rgb);
// Use the converted colorParam in your gamma-space filter calculations.
```

This conversion ensures color parameters match the gamma color space your filter operates in, preventing color accuracy issues.

## Additional resources

* [Create a custom swirl filter](create-custom-swirl-filter.md)