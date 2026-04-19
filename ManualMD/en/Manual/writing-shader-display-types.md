# Control material properties in the Inspector window

In the Unity Editor, you can control how material properties appear in the **Inspector** window. The easiest way to do this is using a [MaterialPropertyDrawer](../ScriptReference/MaterialPropertyDrawer.md).

For more complex needs, you can use the [MaterialEditor](../ScriptReference/MaterialEditor.md), [MaterialProperty](../ScriptReference/MaterialProperty.md), and [ShaderGUI](../ScriptReference/ShaderGUI.md) classes.

## Use a custom editor

Use custom editors to display data types that Unity can’t display using its default material Inspector, or to define custom controls or data validation.

In **ShaderLab**, you can assign a custom editor for all **render pipelines**. To do this, you can place a `CustomEditor` block inside a `Shader` block. You can also assign different custom editors for render pipelines based on the Scriptable Render Pipeline by placing a `CustomEditorForRenderPipeline` block inside a `Shader` block. If your code contains both a `CustomEditor` and `CustomEditorForRenderPipeline` block, the render pipeline specific one overrides the `CustomEditor` one.

### Create a custom editor class for a shader asset

To define a custom editor for **shader** assets that represent a given **Shader object**, you create a script that inherits from the [ShaderGUI](../ScriptReference/ShaderGUI.md) class. Place your script in a folder named *Editor*, in your *Assets* folder.

The script should follow this format:

```
using UnityEditor;

public class ExampleShaderGUI : ShaderGUI 
{
    public override void OnGUI (MaterialEditor materialEditor, MaterialProperty[] properties)
    {
        // Custom code that controls the appearance of the Inspector goes here

        base.OnGUI (materialEditor, properties);
    }
}
```

### Enable the default custom editor

This example code demonstrates the syntax for specifying a default custom editor for a shader asset using the `CustomEditor` block, and then specifying two additional custom editors for specific Render Pipeline Assets using the `CustomEditorForRenderPipeline` block.

```
Shader "Examples/UsesCustomEditor"
{
    // The Unity Editor uses the class ExampleCustomEditor to configure the Inspector for this shader asset
    CustomEditor "ExampleShaderGUI"
    CustomEditorForRenderPipeline "ExampleRenderPipelineShaderGUI" "ExampleRenderPipelineAsset"
    CustomEditorForRenderPipeline "OtherExampleRenderPipelineShaderGUI" "OtherExampleRenderPipelineAsset"

    SubShader
    {
        // Code that defines the SubShader goes here.

        Pass
        {                
              // Code that defines the Pass goes here.
        }
    }
}
```

## Additional resources

* [Custom Editor block in ShaderLab reference](SL-CustomEditor.md)