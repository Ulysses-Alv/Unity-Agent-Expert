# HD Sample Buffer Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/HD-Sample-Buffer-Node.html)

# HD Sample Buffer Node

## Description

The HD Sample Buffer Node samples a buffer directly from the Camera.

## Render pipeline compatibility

| **Node** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** |
| --- | --- | --- |
| **HD Sample Buffer** | No | Yes |

## Ports

| **Name** | **Direction** | **Type** | **Binding** | **Description** |
| --- | --- | --- | --- | --- |
| **UV** | Input | Vector 2 | UV | Input UV value. |
| **Sampler** | Input | SamplerState | None | Determines the sampler that Unity uses to sample the buffer. |
| **Layer Mask** | Input | Float | None | Set the number of the Layer to sample. This port appears when you select **Thickness** in the **Source Buffer** dropdown. |
| **Output** | Output | Changes to one of the following depending on the Source Buffer you select: • Float • Vector 2 • Vector 3 • Vector 4 | None | Output value. |
| **Thickness** | Output | Float | None | Sample the Worldspace value, in meters, between the near and the far plane of the camera.This port appears when you select **Thickness** in the **Source Buffer** dropdown. |
| **Overlap Count** | Output | Float | None | Count the number of triangles for a given pixel. This is useful for vegetation or flat surfaces. This port appears when you select **Thickness** in the **Source Buffer** dropdown. |

## Controls

| **Name** | **Type** | **Options** | **Description** |
| --- | --- | --- | --- |
| **Source Buffer** | Dropdown | • NormalWorldSpace • Smoothness • MotionVectors • IsSky • PostProcessInput • RenderingLayerMask • Thickness | Determines which buffer to sample. |

## Generated Code Example

The following example code represents one possible outcome of this node:

```
float4 Unity_HDRP_SampleBuffer_float(float2 uv, SamplerState samplerState)
{
    return SAMPLE_TEXTURE2D_X_LOD(_CustomPostProcessInput, samplerState, uv * _RTHandlePostProcessScale.xy, 0);
}
```