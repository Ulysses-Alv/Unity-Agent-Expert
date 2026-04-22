# ShaderLab: legacy vertex data channel mapping

**Note**: The **ShaderLab** functionality on this page is legacy, and is documented for backwards compatibility only. If your **shader** source file includes HLSL code, Unity ignores these commands completely. If your shader source file does not include HLSL code, Unity compiles these commands into regular shader programs on import.

## Render pipeline compatibility

| **Feature name** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** |
| --- | --- | --- | --- | --- |
| **Legacy vertex data channel mapping** | Yes | No | No | No |

## Overview

The **BindChannels** command allows you to specify how vertex data maps to the graphics hardware. By default, Unity figures out the bindings for you, but in some cases you want custom ones to be used.

For example you could map the primary UV set to be used in the first texture stage and the secondary UV set to be used in the second texture stage; or tell the hardware that vertex colors should be taken into account.

## Syntax

```
BindChannels { Bind "source", target }
```

Specifies that vertex data *source* maps to hardware *target*.

**Source** can be one of:

* **Vertex**: vertex position
* **Normal**: vertex normal
* **Tangent**: vertex tangent
* **Texcoord**: primary UV coordinate
* **Texcoord1**: secondary UV coordinate
* **Color**: per-vertex color

**Target** can be one of:

* **Vertex**: vertex position
* **Normal**: vertex normal
* **Tangent**: vertex tangent
* **Texcoord0**, **Texcoord1**, …: texture coordinates for corresponding texture stage
* **Texcoord**: texture coordinates for all texture stages
* **Color**: vertex color

## Details

Unity places some restrictions on which sources can be mapped to which targets. Source and target must match for **Vertex**, **Normal**, **Tangent** and **Color**. Texture coordinates from the **mesh** (**Texcoord** and **Texcoord1**) can be mapped into texture coordinate targets (**Texcoord** for all texture stages, or **TexcoordN** for a specific stage).

There are two typical use cases for BindChannels:

* Shaders that take vertex colors into account.
* Shaders that use two UV sets.

## Examples

```
// Maps the first UV set to the first texture stage
// and the second UV set to the second texture stage
BindChannels {
   Bind "Vertex", vertex
   Bind "texcoord", texcoord0
   Bind "texcoord1", texcoord1
}
```

```
// Maps the first UV set to all texture stages
// and uses vertex colors
BindChannels {
   Bind "Vertex", vertex
   Bind "texcoord", texcoord
   Bind "Color", color
}
```