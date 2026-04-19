# Toggle shader keywords in a script

Use a C# script to enable or disable keywords, and check whether a keyword has local or global scope.

## Toggle a keyword

To change the state of a keyword, use the `Material` API. For example, `Material.EnableKeyword()`.

## Toggle keywords in multiple shaders

By default, HLSL **shader** keywords have global scope. As a result, to toggle keywords with the same name across multiple shaders, you can use a C# global keyword object.

Follow these steps:

1. Create a [`GlobalKeyword`](../ScriptReference/Rendering.GlobalKeyword.md) object with the same name as the keywords you want to enable or disable.
2. Use the [`Shader`](../ScriptReference/Shader.md) API to enable or disable the global keyword, which adjusts the HLSL keyword with the same name in all shaders. For example, [`Shader.EnableKeyword`](../ScriptReference/Shader.EnableKeyword.md).

**Note:** When you create a new `GlobalKeyword`, Unity updates its internal mapping between global and local keyword space for all loaded shaders. This can be a CPU-intensive operation. To reduce the impact of this operation, try to create all global keywords after application startup, while your application is loading.

### Prevent global keyword objects affecting shaders

To prevent global keyword objects affecting shader keywords, give the shader keyword local scope instead of global scope.

Add `_local` to the keyword directive. For example:

```
#pragma shader_feature_local RED
```

Now the shader isn’t affected by global `Shader` APIs. You can only enable or disable the keyword using `Material` APIs like `Material.EnableKeyword()`.

**Note:** If you use [Fallback](SL-Fallback.md) or [UsePass](SL-UsePass.md) in a shader, its keyword scope overrides the keyword scopes of the referenced shaders.

## Check the scope of a keyword

To check the scope of a keyword in a shader, select the shader in the **Project** window, then in the ****Inspector**** window open the **Keywords** dropdown. Unity displays keywords with global scope under **Overridable**, and keywords with local scope under **Not Overridable**.

You can also use the [LocalKeyword.isOverridable](../ScriptReference/Rendering.LocalKeyword-isOverridable.md) API.

**Note:** The term `Local` in `LocalKeyword` means the API represents an HLSL keyword rather than a C# keyword object. It doesn’t relate to the scope of the keyword.

To check if a keyword is enabled, do the following:

* If the keyword is overridable, check whether a C# global keyword with the same name exists. If so, check the state of the keyword using the `Shader` API. For example, `Shader.IsKeywordEnabled()`.
* If the keyword isn’t overridable, check the state of a keyword using the `Material` API. For example, `Material.IsKeywordEnabled()`.

This following code demonstrates how to check whether Unity considers a keyword enabled or disabled for a material:

```
using UnityEngine;
using UnityEngine.Rendering;

public class KeywordExample : MonoBehaviour
{
    public Material material;

    void Start()
    {
        CheckShaderKeywordState();
    }

    void CheckShaderKeywordState()
    {
        // Get the instance of the Shader class that the material uses
        var shader = material.shader;

        // Get all the local keywords that affect the Shader
        var keywordSpace = shader.keywordSpace;

        // Iterate over the local keywords
        foreach (var localKeyword in keywordSpace.keywords)
        {
            // If the local keyword is overridable (i.e., it was declared with a global scope),
            // and a global keyword with the same name exists and is enabled,
            // then Unity uses the global keyword state
            if (localKeyword.isOverridable && Shader.IsKeywordEnabled(localKeyword.name))
            {
                Debug.Log("Local keyword with name of " + localKeyword.name + " is overridden by a global keyword, and is enabled");
            }
            // Otherwise, Unity uses the local keyword state
            else
            {
                var state = material.IsKeywordEnabled(localKeyword) ? "enabled" : "disabled";
                Debug.Log("Local keyword with name of " + localKeyword.name + " is " + state);
            }            
        }
    }
}
```

## Change a keyword with a command buffer

To enable or disable a local or global keyword with a Command Buffer, use [`CommandBuffer.EnableKeyword`](../ScriptReference/Rendering.CommandBuffer.EnableKeyword.md) or [`CommandBuffer.DisableKeyword`](../ScriptReference/Rendering.CommandBuffer.DisableKeyword.md).

## Additional resources

* [Shader Graph Keywords](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?subfolder=/manual/Keywords.html)
* [ComputeShader](../ScriptReference/ComputeShader.md) API