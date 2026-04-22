# Texture properties

For each texture that is setup as a **shader**/material property, Unity also sets up some extra information in additional vector properties.

## Texture tiling & offset

[Materials](class-Material.md) often have Tiling and Offset fields for their texture properties. This information is passed into shaders in a float4 *{TextureName}*`_ST` property:

* `x` contains X tiling value
* `y` contains Y tiling value
* `z` contains X offset value
* `w` contains Y offset value

For example, if a shader contains texture named `_MainTex`, the tiling information will be in a `_MainTex_ST` vector.

## Texture size

*{TextureName}*`_TexelSize` - a float4 property contains texture size information:

* `x` contains 1.0/width
* `y` contains 1.0/height
* `z` contains width
* `w` contains height

## Texture HDR parameters

*{TextureName}*`_HDR` - a float4 property with information on how to decode a potentially **HDR** (e.g. RGBM-encoded) texture depending on the [color space](color-spaces-landing.md) used. See `DecodeHDR` function in [UnityCG.cginc](SL-BuiltinIncludes.md) shader include file.