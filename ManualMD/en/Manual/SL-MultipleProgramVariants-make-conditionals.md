# Make shader behavior conditional on keywords

To mark parts of your **shader** code conditional based on whether you enable or disable a shader keyword, use [an HLSL if statement](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-if).

For example:

```
#pragma multi_compile QUALITY_LOW QUALITY_MED QUALITY_HIGH

if (QUALITY_LOW)
{
    // code for low quality setting
}
else if (QUALITY_MED)
{
    // code for medium quality setting
}
else if (QUALITY_HIGH)
{
    // code for high quality setting
}
```

You can now enable and disable keywords [using the Inspector](shader-keywords-material-inspector.md) or [C# scripting](shader-keywords-scripts.md).

Refer to [How Unity compiles branching shaders](shader-conditionals-choose-a-type.md) for more information about when to use which shader directive.

## Branch if a keyword exists

For each keyword in a set, Unity automatically adds a `_KEYWORD_DECLARED` keyword. For example, if you declare a `QUALITY_LOW` keyword, Unity adds a `QUALITY_LOW_KEYWORD_DECLARED` keyword.

You can use this to check if a keyword exists, regardless of whether it’s enabled or disabled.

For example:

```
#pragma multi_compile QUALITY_LOW QUALITY_MED QUALITY_HIGH

#if defined(QUALITY_LOW_KEYWORD_DECLARED)
{
    // The QUALITY_LOW keyword exists
}
```

## Branch when all keywords in a set are disabled

To execute code when all keywords in a `shader_feature` or `multi_compile` set are disabled, Unity must create an additional shader variant for that state.

If you use `shader_feature` to create a single keyword, Unity automatically creates an additional shader variant. For example:

```
// Creates a variant for when FEATURE_1_ON is enabled, and another for when FEATURE_1_ON is disabled. 
#pragma shader_feature FEATURE_1_ON
```

If you use `shader_feature` or `multi_compile` to create a set of two or more keywords, add `_` when you declare a keyword set. For example:

```
// Creates 5 shader variants, including one for when RED, GREEN, BLUE, and WHITE are all disabled.
#pragma shader_feature _ RED GREEN BLUE WHITE

// Creates 2 shader variants, including one for when FEATURE_3_ON is disabled.
#pragma multi_compile _ FEATURE_3_ON
```

If you use `dynamic_branch`, Unity doesn’t need an additional uniform integer for when all keywords are disabled. But to indicate that you use that state in conditional code, you should add `_` when you declare a keyword set. For example:

```
// Creates 3 uniform integers.
// The underscore doesn't create an additional uniform integer, but it indicates you use the additional state in conditional code.
#pragma dynamic_branch _ QUALITY_LOW QUALITY_MED QUALITY_HIGH
```

## Use other statements to make shader behavior conditional

You can also use the following HLSL pre-processor directives to create conditional code with `shader_feature` or `multi_compile`:

* [`#if, #elif, #else and #endif`](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-if).
* [`#ifdef and #ifndef`](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl-appendix-pre-ifdef).

Using these instead of `if` makes it more difficult to change the `#pragma` keyword directive later. For example, if you need to reduce the number of shader variants, it’s more difficult to change to `dynamic_branch`.

## Additional resources

* [Default keyword sets](shader-keywords-default.md)
* [Add built-in keyword sets](SL-MultipleProgramVariants-shortcuts.md)