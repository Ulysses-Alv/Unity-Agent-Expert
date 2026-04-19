# Import a file from the shader library in the Built-In Render Pipeline

Unity contains several files that can be used by your [shader programs](writing-shader-writing-shader-programs-hlsl.md) to bring in predefined variables and helper functions. This is done by the standard `#include` directive, e.g.:

```
SubShader {
    Pass {
        HLSLPROGRAM

        #include "UnityCG.cginc"

        ENDHLSL
    }
}
```

Shader include files in Unity are with `.cginc` extension, and the built-in ones are:

* `HLSLSupport.cginc`: Declares various [preprocessor macros](shader-branching-built-in-macros.md) to aid in multi-platform **shader** development. Unity automatically includes this file if you use `CGPROGRAM`. For more information, refer to [Shader code blocks in ShaderLab reference](shader-shaderlab-code-blocks.md).
* `UnityShaderVariables.cginc`: Declares various [built-in global variables](SL-UnityShaderVariables.md) that are commonly used in shaders. Unity automatically includes this file if you use `CGPROGRAM`. For more information, refer to [Shader code blocks in ShaderLab reference](shader-shaderlab-code-blocks.md).
* `UnityCG.cginc`: Commonly used [built-in helper functions](SL-BuiltinFunctions.md) and data structures.
* `AutoLight.cginc`: Lighting and shadowing functionality, for example [surface shaders](SL-SurfaceShaders.md) use this file internally.
* `Lighting.cginc`: Standard [surface shader](SL-SurfaceShaders.md) lighting models. Unity automatically includes this file if your write surface shaders.
* `TerrainEngine.cginc`: Helper functions for **terrain** and vegetation shaders.

**Note**: Although shader library files have a .cginc file extension, they’re written in HLSL rather than CG.

These files are found inside Unity application (**{unity install path}/Data/Resources/CGIncludes/UnityCG.cginc** on Windows, **/Applications/Unity/Unity.app/Contents/Resources/CGIncludes/UnityCG.cginc** on Mac), if you want to take a look at what exactly is done in any of the helper code.