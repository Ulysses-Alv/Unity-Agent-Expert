# Branch in a shader via built-in macros

Resources for adding a condition in a **shader** based on the platform, shader model, Unity Version, or shader pass.

| **Page** | **Description** |
| --- | --- |
| [Branch based on platform or graphics API](shader-branching-api.md) | Use `SHADER_API` macros to make shader behaviour conditional on the platform or graphics API. |
| [Branch based on shader model](shader-branching-shader-model.md) | Use `SHADER_TARGET` to make shader behaviour conditional on the shader model. |
| [Branch based on platform features](shader-branching-platform.md) | Use platform difference helpers to make shader behaviour conditional on platform functionality. |
| [Branch based on Unity version](shader-branching-unity-version.md) | Use `UNITY_VERSION` to make shader behaviour conditional on the Unity version. |
| [Branch based on shader pass or shader stage](shader-branching-pass.md) | Use `SHADER_STAGE` or `UNITY_PASS` to make shader behaviour conditional on the shader stage or shader pass. |

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)