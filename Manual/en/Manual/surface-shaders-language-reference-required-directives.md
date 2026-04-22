# Surface Shader required directives reference for the Built-In Render Pipeline

* `surfaceFunction` - which Cg function has surface **shader** code. The function should have the form of `void surf (Input IN, inout SurfaceOutput o)`, where Input is a structure you have defined. Input should contain any texture coordinates and extra automatic variables needed by surface function.
* `lightModel` - lighting model to use. Built-in ones are physically based `Standard` and `StandardSpecular`, as well as simple non-physically based `Lambert` (diffuse) and `BlinnPhong` (specular). See [Custom Lighting Models](SL-SurfaceShaderLighting.md) page for how to write your own.
  + `Standard` lighting model uses `SurfaceOutputStandard` output struct, and matches the Standard (metallic workflow) shader in Unity.
  + `StandardSpecular` lighting model uses `SurfaceOutputStandardSpecular` output struct, and matches the Standard (specular setup) shader in Unity.
  + `Lambert` and `BlinnPhong` lighting models are not physically based (coming from Unity 4.x), but the shaders using them can be faster to render on low-end hardware.