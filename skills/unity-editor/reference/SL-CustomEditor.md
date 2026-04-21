# Custom Editor block in ShaderLab reference

This page contains information on using a `CustomEditor` or `CustomEditorForRenderPipeline` block in your **ShaderLab** code to assign [custom editors](editor-CustomEditors.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: CustomEditor block** | Yes | Yes | Yes | Yes |
| **ShaderLab: CustomEditorForRenderPipeline block** | Yes | Yes | Yes | No |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| `CustomEditor "[custom editor class name]"` | Unity uses the custom editor defined in the named class, unless this is overridden by a `CustomEditorForRenderPipeline` block. |
| `CustomEditorForRenderPipeline "[custom editor class name]" "[render pipeline asset class name]"` | When the active Render Pipeline Asset is the named type, Unity uses the custom editor defined in the named class. |

## Additional resources

* [Adding material properties to shaders](writing-shader-change-properties.md)