# Material Quality Keyword node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Material-Quality-Keyword-Node.html)

# Material Quality Keyword node

The Material Quality Keyword node adds and uses global keywords MATERIAL\_QUALITY\_HIGH, MATERIAL\_QUALITY\_MEDIUM and MATERIAL\_QUALITY\_LOW and enables different behaviors for each one of the available quality types.

**Note**:

* Adding this keyword increases the amount of shader variants.
* These quality keywords are only available in URP and HDRP, they are not available at the Built-in Render Pipeline.

To manually set the quality level from a C# script, use the `UnityEngine.Rendering.MaterialQualityUtilities.SetGlobalShaderKeywords(...)` function.

## Additional resources:

* [Keywords](Keywords.html)