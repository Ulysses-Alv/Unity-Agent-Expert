# Declare shader keywords

To declare **shader** keywords, use a `#pragma` directive in the HLSL code. For example:

```
#pragma shader_feature REFLECTION_TYPE1 REFLECTION_TYPE2 REFLECTION_TYPE3
```

Use one of the following shader directives:

* `dynamic_branch` to keep branching code in one compiled shader program.
* `multi_compile` or `shader_feature` to compile multiple shader variants, which means there’s a separate shader program for each branch.

For information about when to use which shader directive, refer to [How Unity compiles branching shaders](shader-conditionals-choose-a-type.md).

If a shader uses more than 128 keywords in total, your project might be slower. As a result, it’s best to keep the number of keywords low. The total number of keywords include 4 keywords per shader that Unity always reserves.

## Declare sets of keywords

The keywords in a single `#pragma` statement are together called a set.

For example, to declare a set of three keywords:

```
#pragma shader_feature REFLECTION_TYPE1 REFLECTION_TYPE2 REFLECTION_TYPE3
```

You can declare multiple sets of keywords in a single shader. For example, to create two sets:

```
#pragma shader_feature REFLECTION_TYPE1 REFLECTION_TYPE2 REFLECTION_TYPE3
#pragma shader_feature RED GREEN BLUE WHITE
```

You can’t do the following:

* Include two keywords with the same name in one set.
* Include duplicate keyword sets in one shader.
* Declare a keyword as both `dynamic_branch` and `shader_feature` or `multi_compile`. Unity uses `dynamic_branch` if you do this.

### Enable multiple keywords in a set

At runtime, Unity has no concept of sets. You can enable or disable any keyword independently. As a result, it’s possible to enable two or more keywords from the same set at the same time, or disable all the keywords.

If you use `shader_feature` and enable multiple keywords from the same set, Unity tries to find the shader variant that most closely matches. This can lead to unintended results.

## Additional resources

* [Targeting shader models and GPU features in HLSL](SL-ShaderCompileTargets.md)
* [Shader Graph: Keyword Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest?subfolder=/manual/Keyword-Node.html)