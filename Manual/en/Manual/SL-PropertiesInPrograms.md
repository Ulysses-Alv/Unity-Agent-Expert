# Add material properties

To assign material properties to a **Shader** object in **ShaderLab**, you place a `Properties` block inside a `Shader` block.

If you want to access some of those properties in a [shader program](writing-shader-writing-shader-programs-hlsl.md), you need to declare a Cg/HLSL variable with the same name and a matching type.

For example these shader properties:

```
_MyColor ("Some Color", Color) = (1,1,1,1) 
_MyVector ("Some Vector", Vector) = (0,0,0,0) 
_MyFloat ("My float", Float) = 0.5 
_MyTexture ("Texture", 2D) = "white" {} 
_MyCubemap ("Cubemap", CUBE) = "" {}
```

would be declared for access in Cg/HLSL code as:

```
fixed4 _MyColor; // low precision type is usually enough for colors
float4 _MyVector;
float _MyFloat; 
sampler2D _MyTexture;
samplerCUBE _MyCubemap;
```

Cg/HLSL can also accept **uniform** keyword, but it is not necessary:

```
uniform float4 _MyColor;
```

Property types in ShaderLab map to Cg/HLSL variable types this way:

* Color and Vector properties map to **float4**, **half4** or **fixed4** variables.
* Range and Float properties map to **float**, **half** or **fixed** variables.
* Texture properties map to **sampler2D** variables for regular (2D) textures; **Cubemaps** map to **samplerCUBE**; and 3D textures map to **sampler3D**.

## Color spaces and color/vector shader data

When using [Linear color space](color-spaces-landing.md), all material color properties are supplied as sRGB colors, but are converted into linear values when passed into shaders.

For example, if your [Properties](SL-Properties.md) shader block contains a `Color` property called “*MyColor“, then the corresponding ”*MyColor” HLSL variable will get the linear color value.

For properties that are marked as `Float` or `Vector` type, no color space conversions are done by default; it is assumed that they contain non-color data. It is possible to add `[Gamma]` attribute for float/vector properties to indicate that they are specified in sRGB space, just like colors (see [Properties](SL-Properties.md)).

## Additional resources

* [Properties block reference in ShaderLab](SL-Properties.md)
* [Writing Shader Programs](writing-shader-writing-shader-programs-hlsl.md)