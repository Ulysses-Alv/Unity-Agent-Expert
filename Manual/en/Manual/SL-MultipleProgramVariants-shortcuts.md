# Add built-in keyword sets

You can use Unity **shader** directive shortcuts to create sets of shader keywords and variants. For example, the following example adds `SHADOWS_DEPTH` and `SHADOWS_CUBE` keywords and variants:

```
#pragma multi_compile_shadowcaster
```

## Shader directive shortcuts

| **Shortcut** | **Shader keywords Unity adds** | **Unity adds a shader variant with all the keywords off** | **Shader pass that uses the shader variants** |
| --- | --- | --- | --- |
| `multi_compile_fog` | `FOG_LINEAR`, `FOG_EXP`, `FOG_EXP2` | Yes | Fog |
| `multi_compile_fwdadd` | `POINT` `DIRECTIONAL` `SPOT` `POINT_COOKIE` `DIRECTIONAL_COOKIE` | No | [`PassType.ForwardAdd`](../ScriptReference/Rendering.PassType.ForwardAdd.md) |
| `multi_compile_fwdadd_fullshadows` | `POINT`, `DIRECTIONAL`, `SPOT`, `POINT_COOKIE`, `DIRECTIONAL_COOKIE`, `SHADOWS_DEPTH SHADOWS_SCREEN` `SHADOWS_CUBE SHADOWS_SOFT SHADOWS_SHADOWMASK`, `LIGHTMAP_SHADOW_MIXING` | No | [`PassType.ForwardAdd`](../ScriptReference/Rendering.PassType.ForwardAdd.md). Adds the ability for the lights to have real-time shadows. |
| `multi_compile_fwdbase` | `DIRECTIONAL`, `LIGHTMAP_ON`, `DIRLIGHTMAP_COMBINED`, `DYNAMICLIGHTMAP_ON`, `SHADOWS_SCREEN`, `SHADOWS_SHADOWMASK`, `LIGHTMAP_SHADOW_MIXING`, `LIGHTPROBE_SH` | No | [`PassType.ForwardBase`](../ScriptReference/Rendering.PassType.ForwardBase.md). |
| `multi_compile_fwdbasealpha` | `DIRECTIONAL`, `LIGHTMAP_ON`, `DIRLIGHTMAP_COMBINED`, `DYNAMICLIGHTMAP_ON`, `LIGHTMAP_SHADOW_MIXING`, `VERTEXLIGHT_ON`, `LIGHTPROBE_SH` | No | [`PassType.ForwardBase`](../ScriptReference/Rendering.PassType.ForwardBase.md) |
| `multi_compile_instancing` | `INSTANCING_ON`. Also adds `PROCEDURAL_ON` if your project uses procedural instancing. | Yes | Instancing |
| `multi_compile_lightpass` | `POINT`, `DIRECTIONAL`, `SPOT`, `POINT_COOKIE`, `DIRECTIONAL_COOKIE`, `SHADOWS_DEPTH`, `SHADOWS_SCREEN`, `SHADOWS_CUBE`, `SHADOWS_SOFT`, `SHADOWS_SHADOWMASK`, `LIGHTMAP_SHADOW_MIXING` | No | All passes that draw real-time light and shadows, except **Light Probes**. |
| `multi_compile_particles` | `SOFTPARTICLES_ON` | Yes | [Particle System](class-ParticleSystem.md) passes. |
| `multi_compile_prepassfinal` | `LIGHTMAP_ON`, `DIRLIGHTMAP_COMBINED`, `DYNAMICLIGHTMAP_ON`, `UNITY_HDR_ON`, `SHADOWS_SHADOWMASK`, `LIGHTPROBE_SH` | Yes | [`PassType.Deferred`](../ScriptReference/Rendering.PassType.Deferred.md) |
| `multi_compile_shadowcaster` | `SHADOWS_DEPTH`, `SHADOWS_CUBE` | No | [`PassType.ShadowCaster`](../ScriptReference/Rendering.PassType.ShadowCaster.md) |
| `multi_compile_shadowcollector` | `SHADOWS_SPLIT_SPHERES`, `SHADOWS_SINGLE_CASCADE` | Yes | Screen-space shadows. |

## Remove variants

You can use the `skip_variants` directive to remove keywords you don’t use.

For example, the following example adds the keywords for the `multi_compile_fwdadd` set, but removes the `POINT` and `POINT_COOKIES` variants.

```
#pragma multi_compile_fwdadd
#pragma skip_variants POINT POINT_COOKIE
```

Refer to [Strip shader variants](shader-variant-stripping.md) for more information about removing variants.

## Additional resources

* [Changing how shaders work using keywords](SL-MultipleProgramVariants.md)