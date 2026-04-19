# Generate lightmap UVs

Unity can calculate the UVs for baked **lightmaps** when you import a model, or you can provide your own data.

This page contains the following information:

* [How to provide your own lightmap UVs](#providing-own-uvs)
* [How to automatically generate lightmap UVs](#automatically-generate-uvs)
  + [Automatic lightmap UV generation settings](LightingGiUvs-Reference.md)

## How to provide your own lightmap UVs

You can author your own lightmap UVs in the content creation software of your choice. Unity uses these UVs as input for its calculations.

Where you put this data depends on whether you are providing UVs for baked lightmaps, real-time lightmaps, or both:

* For baked lightmaps, you must place lightmap UVs in the [Mesh.uv2](../ScriptReference/Mesh-uv2.md). This channel is also called “UV1”.
* For real-time lightmaps:
  + If you already have baked lightmap UVs in the `Mesh.uv2` of your **mesh**, and you want to use the same UVs as input for the real-time lightmaps, you do not need to do anything. Unity falls back to sharing the baked lightmap UVs.
  + If you already have baked lightmap UVs in `Mesh.uv2`, and you want to provide different UVs as input for your real-time lightmaps, place the real-time lightmap UVs in [Mesh.uv3](../ScriptReference/Mesh-uv3.md), also called “UV2”.
  + If you do not already have baked lightmap UVs in the second channel of your mesh, it’s your choice whether you use `Mesh.uv2` or `Mesh.uv3` for real-time lightmap UVs.

A good UV set for lightmaps should adhere to the following rules:

* It should be within the [0,1] x [0,1] UV space.
* It should have a wide enough margin between individual charts. For more information, see [UV overlap feedback](troubleshooting-lightmapping-artifacts.md#ProgressiveLightmapper-UVOverlap).
* It should not have any overlapping faces.
* There should be a low difference between the angles in the UV and the angles in the original geometry.
* There should be a low difference between the relative scale of triangles in the UV and the relative scale of the triangles in the original geometry), unless you want some areas to have a higher lightmap resolution.

## How to automatically generate lightmap UVs

You can tell Unity to automatically generate lightmap UVs for a Model, using the [Model Import Settings](class-FBXImporter.md).

1. Select the Model in your Project view. Unity opens the Model Import Settings in the **Inspector**.
2. In the Model Import Settings, navigate to the **Model** tab, and then the **Geometry** section.
3. Tick the **Generate Lightmap UVs** checkbox. The **Lightmap UVs settings** section appears below the Generate Lightmap UVs checkbox.
4. Optional: Configure the settings in the **Lightmap UVs settings** section. See [Settings](LightingGiUvs-Reference.md) for more information.
5. Click the **Apply** button. Unity generates lightmap UVs into the [Mesh.uv2](../ScriptReference/Mesh-uv2.md) channel.