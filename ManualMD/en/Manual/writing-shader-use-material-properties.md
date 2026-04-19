# Set shader variables with material property values

## Using material properties to set variables in ShaderLab code

To set the value of a variable in your **ShaderLab** code from a material property, put the material property name in square brackets in your ShaderLab code.

This example code demonstrates the syntax for using a material property to set the `units` value of the ShaderLab `Offset` command.

```
Shader "Examples/MaterialPropertyShaderLab"
{
    Properties
    {
        // Change this value in the Material Inspector to affect the value of the Offset command
        _OffsetUnitScale ("Offset unit scale", Integer) = 1
    }
    SubShader
    {
        // The code that defines the rest of the SubShader goes here

        Pass
        {
            Offset 0, [_OffsetUnitScale]

           // The code that defines the rest of the Pass goes here
        }
    }
}
```

## Using material properties to set variables in HLSL code

To set the value of a variable in HLSL code using a material property, give the material property the same name as the **shader** property.

You can see this technique in the following articles, which include working code examples:

* [Writing vertex and fragment shaders for the Built-in Render Pipeline](built-in-shader-examples.md)
* [Writing custom shaders for the Universal Render Pipeline](urp/writing-custom-shaders-urp.md)