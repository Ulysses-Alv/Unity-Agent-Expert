# Replace shaders at runtime in the Built-In Render Pipeline

In the Built-in **Render Pipeline**, you can tell a **Camera** to change the shader that it uses to render certain geometry at runtime. You might do this to achieve a visual effect such as edge detection.

Shader replacement is done from **scripts** with either the [Camera.RenderWithShader](../ScriptReference/Camera.RenderWithShader.md) or [Camera.SetReplacementShader](../ScriptReference/Camera.SetReplacementShader.md) function. Both functions take a **shader** and a **replacementTag**.

It works like this: the camera renders the **scene** as it normally would. The objects still use their materials, but the shader they use changes:

* If **replacementTag** is empty (must be an empty string not null), then all objects in the scene are rendered with the given replacement shader.
* If **replacementTag** is not empty, then for each object that would be rendered:
  + The real object’s shader is queried for the [tag value](SL-SubShaderTags.md).
  + If it does not have that tag, object is **not rendered**.
  + A [subshader](SL-SubShader.md) is found in the replacement shader that has a given tag with the found value. If no such subshader is found, object is **not rendered**.
  + Now that subshader is used to render the object.

So if all shaders would have, for example, a “RenderType” tag with values like “Opaque”, “Transparent”, “Background”, “Overlay”, you could write a replacement shader that only renders solid objects by using one subshader with RenderType=Solid [tag](SL-SubShaderTags.md). The other tag types would not be found in the replacement shader, so the objects would be not rendered. Or you could write several subshaders for different “RenderType” tag values. Incidentally, all built-in **Shader objects** have a “RenderType” tag set.

## Lit shader replacement

When using shader replacement the scene is rendered using the render path that is configured on the camera. This means that the shader used for replacement can contain shadow and lighting passes (you can use surface shaders for shader replacement). This can be useful for doing rendering of special effects and scene debugging.

## Shader replacement tags in built-in shaders

All built-in shaders have a “**RenderType**” tag set that can be used when rendering with replaced shaders. Tag values are the following:

* **Opaque**: most of the shaders ([Normal](shader-NormalFamily.md), [Self Illuminated](shader-SelfIllumFamily.md), [Reflective](shader-ReflectiveFamily.md), terrain shaders).
* **Transparent**: most semitransparent shaders ([Transparent](shader-TransparentFamily.md), Particle, Font, terrain additive pass shaders).
* **TransparentCutout**: masked transparency shaders ([Transparent Cutout](shader-TransparentCutoutFamily.md), two pass vegetation shaders).
* **Background**: **Skybox** shaders.
* **Overlay**: Halo, Flare shaders.
* **TreeOpaque**: **terrain** engine tree bark.
* **TreeTransparentCutout**: terrain engine tree leaves.
* **TreeBillboard**: terrain engine billboarded trees.
* **Grass**: terrain engine grass.
* **GrassBillboard**: terrain engine billboarded grass.

## Built-in scene depth/normals texture

A Camera has a built-in capability to render depth or depth+normals texture, if you need that in some of your effects. See [Camera Depth Texture](SL-CameraDepthTexture.md) page. Note that in some cases (depending on the hardware), the depth and depth+normals textures can internally be rendered using shader replacement. So it is important to have the correct “**RenderType**” tag in your shaders.

## Code Example

Your Start() function specifies the replacement shaders:

```
void Start() {
    camera.SetReplacementShader (EffectShader, "RenderType");
}
```

This requests that the EffectShader will use the RenderType key. The EffectShader will have key-value tags for each RenderType that you want. The Shader will look something like:

```
Shader "EffectShader" {
     SubShader {
         Tags { "RenderType"="Opaque" }
         Pass {
             ...
         }
     }
     SubShader {
         Tags { "RenderType"="SomethingElse" }
         Pass {
             ...
         }
     }
 ...
 }
```

SetReplacementShader will look through all the objects in the scene and, instead of using their normal shader, use the first subshader which has a matching value for the specified key. In this example, any objects whose shader has Rendertype=“Opaque” tag will be replaced by first subshader in EffectShader, any objects with RenderType=“SomethingElse” shader will use second replacement subshader and so one. Any objects whose shader does not have a matching tag value for the specified key in the replacement shader will not be rendered.