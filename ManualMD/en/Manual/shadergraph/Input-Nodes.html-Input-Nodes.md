# Input nodes | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Input-Nodes.html)

# Input nodes

Supply shaders with essential data such as constants, mesh attributes, gradients, matrices, deformation, PBR parameters, scene details, and texture sampling options.

## Basic

| **Topic** | **Description** |
| --- | --- |
| [Boolean](Boolean-Node.html) | Defines a constant Boolean value in the shader. |
| [Color](Color-Node.html) | Defines a constant Vector 4 value in the shader using a Color field. |
| [Constant](Constant-Node.html) | Defines a Float of a mathematical constant value in the shader. |
| [Integer](Integer-Node.html) | Defines a constant Float value in the shader using an Integer field. |
| [Slider](Slider-Node.html) | Defines a constant Float value in the shader using a Slider field. |
| [Time](Time-Node.html) | Provides access to various Time parameters in the shader. |
| [Float](Float-Node.html) | Defines a Float value in the shader. |
| [Vector 2](Vector-2-Node.html) | Defines a Vector 2 value in the shader. |
| [Vector 3](Vector-3-Node.html) | Defines a Vector 3 value in the shader. |
| [Vector 4](Vector-4-Node.html) | Defines a Vector 4 value in the shader. |

## Geometry

| **Topic** | **Description** |
| --- | --- |
| [Bitangent Vector](Bitangent-Vector-Node.html) | Provides access to the mesh vertex or fragment's Bitangent Vector. |
| [Normal Vector](Normal-Vector-Node.html) | Provides access to the mesh vertex or fragment's Normal Vector. |
| [Position](Position-Node.html) | Provides access to the mesh vertex or fragment's Position. |
| [Screen Position](Screen-Position-Node.html) | Provides access to the mesh vertex or fragment's Screen Position. |
| [Tangent Vector](Tangent-Vector-Node.html) | Provides access to the mesh vertex or fragment's Tangent Vector. |
| [UV](UV-Node.html) | Provides access to the mesh vertex or fragment's UV coordinates. |
| [Vertex Color](Vertex-Color-Node.html) | Provides access to the mesh vertex or fragment's Vertex Color value. |
| [View Direction](View-Direction-Node.html) | Provides access to the mesh vertex or fragment's View Direction vector. |
| [Vertex ID](Vertex-ID-Node.html) | Provides access to the mesh vertex or fragment's Vertex ID value. |

## Gradient

| **Topic** | **Description** |
| --- | --- |
| [Blackbody](Blackbody-Node.html) | Samples a radiation based gradient from temperature input (in Kelvin). |
| [Gradient](Gradient-Node.html) | Defines a constant Gradient in the shader. |
| [Sample Gradient](Sample-Gradient-Node.html) | Samples a Gradient given the input of Time. |

## Lighting

| **Topic** | **Description** |
| --- | --- |
| [Ambient](Ambient-Node.html) | Provides access to the Scene's Ambient color values. |
| [Baked GI](Baked-GI-Node.html) | Provides access to the Baked GI values at the vertex or fragment's position. |
| [Main Light Direction](Main-Light-Direction-Node.html) | Provides access to the direction of the main directional light in the scene. |
| [Reflection Probe](Reflection-Probe-Node.html) | Provides access to the nearest Reflection Probe to the object. |

## Matrix

| **Topic** | **Description** |
| --- | --- |
| [Matrix 2x2](Matrix-2x2-Node.html) | Defines a constant Matrix 2x2 value in the shader. |
| [Matrix 3x3](Matrix-3x3-Node.html) | Defines a constant Matrix 3x3 value in the shader. |
| [Matrix 4x4](Matrix-4x4-Node.html) | Defines a constant Matrix 4x4 value in the shader. |
| [Transformation Matrix](Transformation-Matrix-Node.html) | Defines a constant Matrix 4x4 value for a default Unity Transformation Matrix in the shader. |

## Mesh Deformation

| **Topic** | **Description** |
| --- | --- |
| [Compute Deformation Node](Compute-Deformation-Node.html) | Passes compute deformed vertex data to a vertex shader. Only works with the [Entities Graphics package](https://docs.unity3d.com/Packages/com.unity.entities.graphics@latest/). |
| [Linear Blend Skinning Node](Linear-Blend-Skinning-Node.html) | Applies Linear Blend Vertex Skinning. Only works with the [Entities Graphics package](https://docs.unity3d.com/Packages/com.unity.entities.graphics@latest/). |

## Sprite Deformation

| **Topic** | **Description** |
| --- | --- |
| [Sprite Skinning Node](Sprite-Skinning-Node.html) | Applies Vertex Skinning on Sprites. Only works with the [2D Animation](https://docs.unity3d.com/Packages/com.unity.2d.animation@latest/). |

## PBR

| **Topic** | **Description** |
| --- | --- |
| [Dielectric Specular](Dielectric-Specular-Node.html) | Returns a Dielectric Specular F0 value for a physically based material. |
| [Metal Reflectance](Metal-Reflectance-Node.html) | Returns a Metal Reflectance value for a physically based material. |

## Scene

| **Topic** | **Description** |
| --- | --- |
| [Camera](Camera-Node.html) | Provides access to various parameters of the current Camera. |
| [Fog](Fog-Node.html) | Provides access to the Scene's Fog parameters. |
| [Object](Object-Node.html) | Provides access to various parameters of the Object. |
| [Scene Color](Scene-Color-Node.html) | Provides access to the current Camera's color buffer. |
| [Scene Depth](Scene-Depth-Node.html) | Provides access to the current Camera's depth buffer. |
| [Screen](Screen-Node.html) | Provides access to parameters of the screen. |
| [Eye Index](Eye-Index-Node.html) | Provides access to the Eye Index when stereo rendering. |

## Texture

| **Topic** | **Description** |
| --- | --- |
| [Cubemap Asset](Cubemap-Asset-Node.html) | Defines a constant Cubemap Asset for use in the shader. |
| [Sample Cubemap](Sample-Cubemap-Node.html) | Samples a Cubemap and returns a Vector 4 color value for use in the shader. |
| [Sample Reflected Cubemap Node](Sample-Reflected-Cubemap-Node.html) | Samples a Cubemap with reflected vector and returns a Vector 4 color value for use in the shader. |
| [Sample Texture 2D](Sample-Texture-2D-Node.html) | Samples a Texture 2D and returns a color value for use in the shader. |
| [Sample Texture 2D Array](Sample-Texture-2D-Array-Node.html) | Samples a Texture 2D Array at an Index and returns a color value for use in the shader. |
| [Sample Texture 2D LOD](Sample-Texture-2D-LOD-Node.html) | Samples a Texture 2D at a specific LOD and returns a color value for use in the shader. |
| [Sample Texture 3D](Sample-Texture-3D-Node.html) | Samples a Texture 3D and returns a color value for use in the shader. |
| [Sample Virtual Texture](Sample-Virtual-Texture-Node.html) | Samples a Virtual Texture and returns color values for use in the shader. |
| [Sampler State](Sampler-State-Node.html) | Defines a Sampler State for sampling textures. |
| [Texture Size](Texture-Size-Node.html) | Returns the Width and Height of the texel size of Texture 2D input. |
| [Texture 2D Array Asset](Texture-2D-Array-Asset-Node.html) | Defines a constant Texture 2D Array Asset for use in the shader. |
| [Texture 2D Asset](Texture-2D-Asset-Node.html) | Defines a constant Texture 2D Asset for use in the shader. |
| [Texture 3D Asset](Texture-3D-Asset-Node.html) | Defines a constant Texture 3D Asset for use in the shader. |

## UI

| **Topic** | **Description** |
| --- | --- |
| [Element Texture UV](element-texture-uv-node.html) | Provides the texture coordinates (UV) typically used to sample the texture assigned to a UI element. |
| [Element Layout UV](element-layout-uv-node.html) | Provides the layout UV coordinates within a UI element's layout rectangle. |
| [Element Texture Size](element-texture-size-node.html) | Provides the size of the texture assigned to a UI element. |