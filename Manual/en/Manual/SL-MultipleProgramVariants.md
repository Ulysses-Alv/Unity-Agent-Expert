# Changing how shaders work using keywords

Resources and techniques for adding **shader** keywords, using them to create branches and shader variants, and toggling them in the Unity Editor or in a script.

| **Page** | **Description** |
| --- | --- |
| [Shader keywords workflow](shader-keywords.md) | Learn about defining shader keywords to create shaders that share some common code, but have different functionality when you enable or disable a keyword. |
| [How Unity compiles branching shaders](shader-conditionals-choose-a-type.md) | Choose dynamic branching or shader variants. |
| [Declare shader keywords](SL-MultipleProgramVariants-declare.md) | Use a `#pragma` directive in HLSL code to declare keywords. |
| [Make shader behavior conditional on keywords](SL-MultipleProgramVariants-make-conditionals.md) | Use an `if` statement in HLSL code to mark parts of your shader code conditional based on whether you enable or disable a shader keyword. |
| [Toggle shader keywords in the Editor](shader-keywords-material-inspector.md) | Add material properties for shader keywords. |
| [Toggle shader keywords in a script](shader-keywords-scripts.md) | Check if a keyword is enabled or disabled, and enable and disable it. |
| [Add built-in keyword sets](SL-MultipleProgramVariants-shortcuts.md) | Use Unity shader directive shortcuts to create sets of shader keywords and variants. |

## Additional resources

* [Limit shader variants when you declare shader keywords](shader-variant-stripping.md#limit-shader-variants-when-you-declare-shader-keywords)
* [Unity Shader Variants Optimization & Troubleshooting Tips](https://unity.com/blog/engine-platform/shader-variants-optimization-troubleshooting-tips)
* [HLSL `pragma` directives reference](SL-PragmaDirectives.md)