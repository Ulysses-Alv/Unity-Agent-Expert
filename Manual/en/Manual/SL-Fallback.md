# Fallback block in ShaderLab reference

This page contains information on using a `Fallback` block in your **ShaderLab** code to assign a fallback **Shader** object. For information on how a **Shader object** works, and how Unity chooses when to use a fallback, see [Shader objects introduction](shader-objects.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: Fallback block** | Yes | Yes | Yes | Yes |

## Syntax

To assign a fallback, you place a `Fallback` block inside a `Shader` block.

| **Signature** | **Function** |
| --- | --- |
| `Fallback "<name>"` | If no compatible SubShaders are found, use the named Shader object. |
| `Fallback Off` | Do not use a fallback Shader object in place of this one. If no compatible SubShaders are found, display the [error shader](shader-error.md). |

## Example

This example code demonstrates the syntax for creating a Shader object that has a named fallback.

```
Shader "Examples/ExampleFallback"
{
    SubShader
    {
        // Code that defines the SubShader goes here.

        Pass
        {                
              // Code that defines the Pass goes here.
        }
    }

    Fallback "ExampleOtherShader"
}
```