# Mesh vertex data

The elements of vertex data are called **vertex attributes**.

Every vertex can have the following attributes:

* [Position](#position)
* [Normal](#normal)
* [Tangent](#tangent)
* [Color](#color)
* Up to 8 [texture coordinates](#uvs)
* [Bone weights and blend indices](#bone-weight) (skinned meshes only)

Internally, all vertex data is stored in separate arrays of the same size. If your **mesh** contains an array with 10 vertex positions, it also has arrays with 10 elements for each other vertex attribute that it uses.

In C#, Unity describes the available vertex attributes with the [VertexAttribute](../ScriptReference/Rendering.VertexAttribute.md) enum. You can check whether an instance of the `Mesh` class has a given vertex attribute with the [Mesh.HasVertexAttribute](../ScriptReference/Mesh.HasVertexAttribute.md) function.

## Position

The **vertex position** represents the position of the vertex in **object space**.

Unity uses this value to determine the surface of the mesh.

This vertex attribute is required for all meshes.

In the `Mesh` class, the simplest way to access this data is with [Mesh.GetVertices](../ScriptReference/Mesh.GetVertices.md) and [Mesh.SetVertices](../ScriptReference/Mesh.SetVertices.md). Unity also stores this data in [Mesh.vertices](../ScriptReference/Mesh-vertices.md), but this older property is less efficient and user-friendly.

## Normal

The **vertex normal** represents the direction that points directly “out” from the surface at the position of the vertex.

Unity uses this value to calculate the way that light reflects off the surface of a mesh.

This vertex attribute is optional.

In the `Mesh` class, the simplest way to access this data is with [Mesh.GetNormals](../ScriptReference/Mesh.GetNormals.md) and [Mesh.SetNormals](../ScriptReference/Mesh.SetNormals.md). Unity also stores this data in [Mesh.normals](../ScriptReference/Mesh-normals.md), but this older property is less efficient and user-friendly.

## Tangent

The **vertex tangent** represents the direction that points along the “u” (horizontal texture) axis of the surface at the position of the vertex.

Unity stores the vertex tangent with an additional piece of data, in a four-component vector. The x,y,z components of the vector describe the tangent, and the w component of the vector describes its [orientation](https://en.wikipedia.org/wiki/Orientation_(vector_space)). Unity uses the w value to compute the **binormal**, which is the cross product of the tangent and normal.

Unity uses the tangent and binormal values in normal mapping.

This vertex attribute is optional.

In the `Mesh` class, the simplest way to access this data is with [Mesh.GetTangents](../ScriptReference/Mesh.GetTangents.md) and [Mesh.SetTangents](../ScriptReference/Mesh.SetTangents.md). Unity also stores this data in [Mesh.tangents](../ScriptReference/Mesh-tangents.md), but this older property is less efficient and user-friendly.

## Texture coordinates (UVs)

A mesh can contain up to eight sets of **texture coordinates**. Texture coordinates are commonly called **UVs**, and the sets are called **channels**.

Unity uses texture coordinates when it “wraps” a texture around the mesh. The UVs indicate which part of the texture aligns with the mesh surface at the vertex position.

UV channels are commonly called “UV0” for the first channel, “UV1” for the second channel, and so on up to “UV7”. The channels respectively map to the **shader** semantics `TEXCOORD0`, `TEXCOORD1`, and so on up to `TEXCOORD7`.

By default, Unity uses the first channel (UV0) to store UVs for regular textures such as diffuse maps and specular maps. Unity can use the second channel (UV1) to store baked **lightmap** UVs, and the third channel (UV2) to store input data for real-time lightmap UVs. For more information on lightmap UVs and how Unity uses these channels, see [Lightmap UVs](LightingGiUvs.md).

All eight texture coordinate attributes are optional.

In the `Mesh` class, the simplest way to access this data is with [Mesh.GetUVs](../ScriptReference/Mesh.GetUVs.md) and [Mesh.SetUVs](../ScriptReference/Mesh.SetUVs.md). Unity also stores this data in the following properties: [Mesh.uv](../ScriptReference/Mesh-uv.md), [Mesh.uv2](../ScriptReference/Mesh-uv2.md), [Mesh.uv3](../ScriptReference/Mesh-uv3.md) and so on, up to [Mesh.uv8](../ScriptReference/Mesh-uv8.md). Note that these older properties are less efficient and user-friendly.

## Color

The **vertex color** represents the base color of a vertex, if any.

This color exists independently of any textures that the mesh may use.

This vertex attribute is optional.

In the `Mesh` class, the simplest way to access this data is with [Mesh.GetColors](../ScriptReference/Mesh.GetColors.md) and [Mesh.SetColors](../ScriptReference/Mesh.SetColors.md). Unity also stores this data in [Mesh.colors](../ScriptReference/Mesh-colors.md), but this older property is less efficient and user-friendly.

## Blend indices and bone weights

In a skinned mesh, **blend indices** indicate which bones affects a vertex, and **bone weights** describe how much influence those bones have on the vertex.

In Unity, these vertex attributes are stored together.

Unity uses blend indices and bone weights to deform a skinned mesh based on the movement of its skeleton. For more information, see [Skinned Mesh Renderer](class-SkinnedMeshRenderer.md).

These vertex attributes are required for skinned meshes.

In the past, Unity only allowed up to 4 bones to influence a vertex. It stored this data in a [BoneWeight](../ScriptReference/BoneWeight.md) struct, in the [Mesh.boneWeights](../ScriptReference/Mesh-boneWeights.md) array. Now, Unity allows up to 256 bones to influence a vertex. It stores this data in a [BoneWeight1](../ScriptReference/BoneWeight1.md) struct, and you can access it with [Mesh.GetAllBoneWeights](../ScriptReference/Mesh.GetAllBoneWeights.md) and [Mesh.SetBoneWeights](../ScriptReference/Mesh.SetBoneWeights.md). For more information, read the linked API documentation.