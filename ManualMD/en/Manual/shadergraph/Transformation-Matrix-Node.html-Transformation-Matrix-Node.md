# Transformation Matrix Node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Transformation-Matrix-Node.html)

# Transformation Matrix Node

## Description

Defines a constant **Matrix 4x4** value for a common **Transformation Matrix** in the shader. The **Transformation Matrix** can be selected from the dropdown parameter.

Two output value options for this node, **Inverse Projection** and **Inverse View Projection**, are not compatible with the Built-In Render Pipeline target. When you choose either of these options and target the Built-In Render Pipeline, this node produces an entirely black result.

##### Important

The Built-In Render Pipeline is deprecated and will be made obsolete in a future release.  
It remains supported, including bug fixes and maintenance, through the full Unity 6.7 LTS lifecycle.  
For more information on migration, refer to [Migrating from the Built-In Render Pipeline to URP](https://docs.unity3d.com/6000.5/Documentation/Manual/urp/upgrading-from-birp.html) and [Render pipeline feature comparison](https://docs.unity3d.com/6000.5/Documentation/Manual/render-pipelines-feature-comparison.html).

## Ports

| Name | Direction | Type | Binding | Description |
| --- | --- | --- | --- | --- |
| Out | Output | Matrix 4 | None | Output value |

## Controls

| Control | Description |
| --- | --- |
| (Dropdown) | Sets the output value. The options are:  * **Model** * **InverseModel** * **View** * **InverseView** * **Projection** * **InverseProjection** * **ViewProjection** * **InverseViewProjection** |

## Generated Code Example

The following example code represents one possible outcome of this node per mode.

**Model**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_M;
```

**InverseModel**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_I_M;
```

**View**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_V;
```

**InverseView**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_I_V;
```

**Projection**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_P;
```

**InverseProjection**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_I_P;
```

**ViewProjection**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_VP;
```

**InverseViewProjection**

```
float4x4 _TransformationMatrix_Out = UNITY_MATRIX_I_VP;
```