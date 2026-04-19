# HLSL pragma directives reference

## Standard pragma directives

Unity supports all `#pragma` directives that are part of standard HLSL, as long as these directives are in regular include files. For more information on these directives, see the HLSL documentation: [pragma Directive](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-pragma).

## Surface Shaders

If you are writing a [Surface Shader](SL-SurfaceShaders.md), use this directive to tell the compiler which function to use as the **surface function**, and pass data to that function.

| **Statement** | **Function** |
| --- | --- |
| `#pragma surface <surface function> <lighting model> <optional parameters>` | Compile the function with the given name as the surface shader, so that it works with the given lighting model. For more information, see [Surface Shaders](SL-SurfaceShaders.md). |

## Shader stages

If you are writing a regular graphics **shader**, use these directives to tell the compiler which functions to use for different shader stages. The `#pragma vertex` and `#pragma fragment` directives are required, but other stages are optional.

| **Statement** | **Function** |
| --- | --- |
| `#pragma vertex <name>` | Compile the function with the given name as the vertex shader. Replace <name> with the function name. This directive is required in regular graphics shaders. |
| `#pragma fragment <name>` | Compile the function with the given name as the fragment shader. Replace <name> with the function name. This directive is required in regular graphics shaders. |
| `#pragma geometry <name>` | Compile the function with the given name as the geometry shader. Replace <name> with the function name. This option automatically turns on `#pragma require geometry`; for more information, see [Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md).  **Note**: Metal does not support geometry shaders. |
| `#pragma hull <name>` | Compile the function with the given name as the DirectX 11 hull shader. Replace <name> with the function name. This automatically adds `#pragma require tessellation`; for more information, see [Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md). |
| `#pragma domain <name>` | Compile the function with the given name as the DirectX 11 domain shader. Replace <name> with the function name. This option automatically turns on `#pragma require tessellation`; for more information, see [Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md). |

## Shader variants and keywords

Use these directives to tell the shader compiler how to handle shader variants and keywords. For more information, see [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md).

| **Directive** | **Description** |
| --- | --- |
| `#pragma multi_compile <keywords>` | Declares a collection of keywords. The compiler includes all of the keywords in the build.  You can use suffixes such as `_local` to set additional options.  For more information and a list of supported suffixes, see [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md). |
| `#pragma shader_feature <keywords>` | Declares a collection of keywords. The compiler excludes unused keywords from the build.  You can use suffixes such as `_local` to set additional options.  For more information and a list of supported suffixes, see [Declaring and using shader keywords in HLSL](SL-MultipleProgramVariants.md). |
| `#pragma hardware_tier_variants <values>` | Built-in Render Pipeline only: Add keywords for **graphics tiers** when compiling for a given graphics API. For more information, see [Graphics tiers](graphics-tiers.md). |
| `#pragma skip_variants <list of keywords>` | Strip specified keywords. |

## GPU requirements and shader model support

Use these directives to tell the compiler that your shader requires specific GPU features.

| **Statement** | **Function** |
| --- | --- |
| `#pragma target <value>` | The minimum shader model that this shader program is compatible with. Replace <value> with a valid value. For a list of valid values, see [Shader compilation: Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md). |
| `#pragma require <value>` | The minimum GPU features that this shader is compatible with. Replace <value> with a valid value, or multiple valid values separated by a space. For a list of valid values, see [Shader compilation: Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md). |

## Graphics APIs

Use these directives to tell Unity to include or exclude code for a given graphics API.

| **Statement** | **Function** |
| --- | --- |
| `#pragma only_renderers <value>` | Compile this shader program only for given graphics APIs. Replace <values> with a space-delimited list of valid values. For more information and a list of valid values, refer to [Targeting graphics APIs and platforms in HLSL](SL-ShaderCompilationAPIs.md).  For example, `#pragma only_renderers glcore` only compiles for the desktop OpenGL. Like the ES 3 target, this also scales up to contain all desktop OpenGL versions, where basic shaders will support OpenGL 2.x while shaders requiring shader model 5.0 features require OpenGL 4.2+. |
| `#pragma exclude_renderers <value>` | Do not compile this shader program for given graphics APIs. Replace <value> with a space-delimited list of valid values. For more information and a list of valid values, refer to [Targeting graphics APIs and platforms in HLSL](SL-ShaderCompilationAPIs.md). |

## Other pragma directives

| **Statement** | **Function** |
| --- | --- |
| `#pragma instancing_options <options>` | Enable GPU instancing in this shader, with given options. For more information, see [GPU instancing](GPUInstancing.md) |
| `#pragma once` | Put this directive in a file to ensure that the compiler includes the file only once in a shader program.  **Note:** Unity only supports this directive when the [Caching Shader Preprocessor](shader-compilation.md#preprocessor) is enabled. |
| `#pragma enable_d3d11_debug_symbols` | Generates shader debug symbols for supported graphics APIs, and disables optimizations for all graphics APIs. Use this for debugging shader code in an external tool.  Unity generates debug symbols for Vulkan, DirectX 11 and 12, and supported console platforms.  **Warning:** Using this results in an increased file size and reduced shader performance. When you have finished debugging your shaders and you are ready to make a final build of your application, remove this line from your shader source code and recompile the shaders. |
| `#pragma skip_optimizations <value>` | Forces optimizations off for given graphics APIs. Replace <values> with a space-delimited list of valid values. For a list of valid values, see [Targeting graphics APIs and platforms in HLSL](SL-ShaderCompilationAPIs.md) |
| `#pragma hlslcc_bytecode_disassembly` | Embed disassembled HLSLcc bytecode into a translated shader. |
| `#pragma disable_fastmath` | Enable precise IEEE 754 rules involving NaN handling. This currently only affects the Metal platform. |
| `#pragma editor_sync_compilation` | Force synchronous compilation. This affects the Unity Editor only. For more information, see [Asynchronous Shader compilation](AsynchronousShaderCompilation.md). |
| `#pragma enable_cbuffer` | Emit `cbuffer(name)` when using `CBUFFER_START(name)` and `CBUFFER_END` macros from HLSLSupport even if the current platform does not support constant buffers. |