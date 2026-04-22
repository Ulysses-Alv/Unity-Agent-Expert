# Branch based on shader pass or shader stage

## Shader stage being compiled

Preprocessor macros `SHADER_STAGE_VERTEX`, `SHADER_STAGE_FRAGMENT`, `SHADER_STAGE_DOMAIN`, `SHADER_STAGE_HULL`, `SHADER_STAGE_GEOMETRY`, `SHADER_STAGE_COMPUTE` are defined when compiling each **Shader** stage. Typically they are useful when sharing Shader code between **pixel** Shaders and compute Shaders, to handle cases where some things have to be done slightly differently.

## Surface Shader pass

When [Surface Shaders](SL-SurfaceShaders.md) are compiled, they generate a lot of code for various passes to do lighting. When compiling each pass, one of the following macros is defined:

| **Macro:** | **Use:** |
| --- | --- |
| `UNITY_PASS_FORWARDBASE` | **Forward rendering** base pass (main directional light, lightmaps, SH). |
| `UNITY_PASS_FORWARDADD` | Forward rendering additive pass (one light per pass). |
| `UNITY_PASS_DEFERRED` | **Deferred shading** pass (renders G-buffer). |
| `UNITY_PASS_SHADOWCASTER` | Shadow caster and depth Texture rendering pass. |

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)