# GPU instancing shader reference for the Built-In Render Pipeline

This section contains information about **shader** additions that relate to GPU instancing.

| Addition | Description |
| --- | --- |
| `#pragma multi_compile_instancing` | Generates instancing variants. This is required for fragment and **vertex shaders**. It is optional for **Surface Shaders**. |
| `#pragma instancing_options` | Specifies options that Unity uses for instances. For information on the option switches available, see [`#pragma instancing_options`](#instancing_options-switches). |
| `UNITY_VERTEX_INPUT_INSTANCE_ID` | Defines an instance ID in the vertex shader input/output structure. To use this macro, enable the INSTANCING\_ON shader keyword. Otherwise, Unity doesn’t set up the instance ID. To access the instance ID, use `vertexInput.instanceID` inside an #ifdef INSTANCING\_ON block. If you don’t use this block, variants fail to compile. |
| `UNITY_INSTANCING_BUFFER_START(bufferName)` | Declares the start of a per-instance constant buffer named `bufferName`. Use this macro with `UNITY_INSTANCING_BUFFER_END` to wrap declarations of the properties that you want to be unique to each instance. Declare properties inside the buffer using`UNITY_DEFINE_INSTANCED_PROP`. |
| `UNITY_INSTANCING_BUFFER_END(bufferName)` | Declares the end of a per-instance constant buffer named `bufferName`. Use this macro with `UNITY_INSTANCING_BUFFER_START` to wrap declarations of the properties that you want to be unique to each instance. Declare properties inside the buffer using`UNITY_DEFINE_INSTANCED_PROP`. |
| `UNITY_DEFINE_INSTANCED_PROP(type, propertyName)` | Defines a per-instance shader property with the specified type and name. In the examples below, the `_Color` property is unique. |
| `UNITY_SETUP_INSTANCE_ID(v);` | Allows shader functions to access the instance ID. For vertex shaders, this macro is required at the beginning. For fragment shaders, this addition is optional. For an example, see [Vertex and fragment shader](gpu-instancing-vertex-fragment-shader-example.md). |
| `UNITY_TRANSFER_INSTANCE_ID(v, o);` | Copies the instance ID from the input structure to the output structure in the vertex shader. Use this macro if you need to access per-instance data in the fragment shader. |
| `UNITY_ACCESS_INSTANCED_PROP(bufferName, propertyName)` | Accesses a per-instance shader property in an instancing constant buffer. Unity uses the instance ID to index into the instance data array.`bufferName` must match the name of the constant buffer that contains the specified property.This macro compiles differently for INSTANCING\_ON and non-instancing variants. |

## Instancing\_options switches

The [#pragma instancing\_options](#pragma-instancing\_options) directive can use the following switches:

| **Switch** | **Description** |
| --- | --- |
| `forcemaxcount:batchSize` and `maxcount:batchSize` | On most platforms, Unity automatically calculates the instancing data array size. It divides the maximum constant buffer size on the target device with the size of the structure that contains all per-instance properties. Generally, you don’t need to worry about the batch size. However, some platforms require a fixed array size. To specify the batch size for those platforms, use the `maxcount` option. Other platforms ignore this option. If you want to force a batch size for all platforms, use `forcemaxcount`. This is useful when, for example, your project uses RenderMeshInstanced to issue draw calls with 256 instanced **sprites**. The default value for the two options is 500. |
| `assumeuniformscaling` | Instructs Unity to assume that all the instances have uniform scalings (the same scale for all X, Y, and Z axes). |
| `nolodfade` | Makes Unity not apply GPU instancing to [LOD](LevelOfDetail.md) fade values. |
| `nolightprobe` | Prevents Unity from applying GPU instancing to [Light Probe](LightProbes.md) values and their occlusion data. Setting this option to `ON` can improve performance if your project doesn’t contain **GameObjects** that use both GPU instancing and Light Probes. |
| `nolightmap` | Prevents Unity from applying GPU instancing to **lightmap** atlas information values. Setting this option to `ON` can improve performance if your project doesn’t contain GameObjects that use both GPU instancing and lightmaps. |
| `procedural:FunctionName` | Generates an additional variant for use with [Graphics.RenderMeshIndirect](../ScriptReference/Graphics.RenderMeshIndirect.md). At the beginning of the vertex shader stage, Unity calls the function specified after the colon. To set up the instance data manually, add per-instance data to this function in the same way you would normally add per-instance data to a shader. Unity also calls this function at the beginning of a fragment shader if any of the fetched instance properties are included in the fragment shader. |