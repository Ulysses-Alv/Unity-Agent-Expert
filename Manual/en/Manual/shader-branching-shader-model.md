# Branch based on shader model

`SHADER_TARGET` is defined to a numeric value that matches the **Shader** target compilation model (that is, matching `#pragma target` directive). For example, `SHADER_TARGET` is `30` when compiling into Shader model 3.0. You can use it in Shader code to do conditional checks. For example:

```
#if SHADER_TARGET < 30
    // less than Shader model 3.0:
    // very limited Shader capabilities, do some approximation
#else
    // decent capabilities, do a better thing
#endif
```

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)