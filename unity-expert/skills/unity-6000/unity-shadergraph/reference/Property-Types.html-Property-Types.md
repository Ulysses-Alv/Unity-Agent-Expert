# Property Types | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Property-Types.html)

# Property Types

## Description

**Property Types** are the types of [Property](https://docs.unity3d.com/Manual/SL-Properties.html) than can be defined on the [Blackboard](Blackboard.html) for use in the **Graph**. These [Properties](https://docs.unity3d.com/Manual/SL-Properties.html) are exposed to the [Inspector](https://docs.unity3d.com/Manual/UsingTheInspector.html) for [Materials](https://docs.unity3d.com/Manual/class-Material.html) that use the shader.

Each property has an associated **Data Type**. See [Data Types](Data-Types.html) for more information.

## Common Parameters

All properties have the following common parameters in addition to those specific to their [Data Types](Data-Types.html).

| Parameter | Description |
| --- | --- |
| **Name** | Displays the user-facing name of the property in the UI. |
| **Reference** | Defines the internal identifier used by the shader for this property; use this `Reference` instead of the display `Name` when accessing the property from a script.  If you overwrite this parameter, be aware of the following:  * If the string doesn't begin with an underscore, Unity automatically adds one. * If the string contains any characters that HLSL does not support, Unity removes them. * You can revert to the default value: right-click on the **Reference** field label, and select **Reset Reference**. |
| **Promote to final Shader** | Makes the property accessible across the final shader as a material property, not as an input port on the Subgraph Node. |
| **Precision** | Sets the numeric precision for the property’s data type.  The options are:  * **Inherit**: Uses the precision defined by the graph or parent context. * **Single**: Uses single-precision (float) for maximum accuracy. * **Half**: Uses half-precision to reduce memory and improve performance. * **Use Graph Precision**: Uses the precision mode set in the graph settings.   For more details, refer to [Precision Modes](Precision-Modes.html). |
| **Scope** | Specifies where and how the property is edited across materials. The options are:  * **Global**: Makes the property editable at a global level, through a C# script only, for all materials that use it. Selecting this option hides or grays out all parameters that relate to the Inspector UI display. * **Per Material**: Makes the property independently editable per material, either through a C# script, or in the Inspector UI if you enable **Show In Inspector**. * **Hybrid Per Instance**: Has the same effect as **Per Material**, unless you're using [DOTS instancing](https://docs.unity3d.com/Packages/com.unity.entities.graphics@latest/index.html?subfolder=/manual/dots-instancing-shader.html). |
| **Default Value** | Sets the property's initial value to be serialized and used when new material instances are created. The value depends on the property type. |
| **Preview Value** | Sets a value to use for preview in the Shader Graph window, only when you set **Scope** to **Global**. |
| **Show In Inspector** | Displays the property in the material Inspector when enabled. If you disable this option, it includes an `[HideInInspector]` attribute to the material property (refer to [Properties block reference in ShaderLab](https://docs.unity3d.com/Manual/SL-Properties.html#material-property-attributes) for more details). |
| **Read Only** | Marks the property as non-editable in the material Inspector by adding the [`PerRendererData`](https://docs.unity3d.com/ScriptReference/Rendering.ShaderPropertyFlags.html) attribute. |
| **Custom Attributes** | Enables attachment of custom scripted drawers or decorators to extend the material property UI, such as adding static headers or complex controls. The [Custom Material Property Drawers](Shader-Graph-Sample-Custom-Material-Property-Drawers.html) sample shows how to display a Vector2 as a min/max slider, for example.  **Note**: When you declare the custom functions in the script, make sure to suffix their names with `Drawer` or `Decorator`.  In the list, use **+** or **-** to add or remove entries. Each entry corresponds to a function call which requires the following parameters:  * **Name**: A shortened version of the function name, without its `Drawer` or `Decorator` suffix. * **Value**: The input values for the function as the script expects them.  **Note**: A property can only have one drawer at any given time. |
| **Use Custom Binding** | Turns the property into a bound input port for connection to the [**Branch On Input Connection**](Branch-On-Input-Connection-Node.html) node. In the **Label** field, enter the label for the default value that displays on your Subgraph node's port binding in its parent Shader Graph. This property is available only in sub graphs. |

## Float

Defines a **Float** value.

Parameters specific to Float properties in addition to the [common parameters](#common-parameters):

| Parameter | Description |
| --- | --- |
| **Mode** | Selects the UI mode used to display and edit the property value in the material Inspector, requiring a specific subset of parameters depending on the chosen option.  The options are:  * **Default**: Displays a scalar input field in the material Inspector; only requires a **Default Value**. * **Slider**: Defines the Float property in [`Range`](https://docs.unity3d.com/Manual/SL-Properties.html#material-property-declaration-syntax-by-type) mode to display a slider; use [additional parameters](#slider) to define the slider type. * **Integer**: Displays an integer input field in the material Inspector; only requires a **Default Value**. * **Enum**: Adds an [`Enum`](https://docs.unity3d.com/ScriptReference/MaterialPropertyDrawer.html) attribute to the Float property to display a drop-down of specific values; use [additional parameters](#enum) to define the enum type. |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). The value is either a float or an integer depending on the selected **Mode** and its options. Not available when **Scope** is set to **Global**. |
| **Requires Literal Input** | Requires the input to be a constant value. When enabled, if the user connects a variable, the shader compilation fails with an error. |

### Slider

Additional parameters available when you set the Float property **Mode** to **Slider**.

| Parameter | Description |
| --- | --- |
| **Slider Type** | Selects the slider response type applied when adjusting the value in the material Inspector.  The options are:  * **Default**: Displays a slider with a linear response; the value changes linearly within the slider range. * **Power**: Adds a [`PowerSlider`](https://docs.unity3d.com/ScriptReference/MaterialPropertyDrawer.html) attribute to display a slider with a non-linear response; the value changes exponentially within the range according to the specified **Power**. * **Integer**: Adds an [`IntRange`](https://docs.unity3d.com/ScriptReference/MaterialPropertyDrawer.html) attribute to display a slider with integer steps; the value changes in whole-number increments within the range. |
| **Min** | Sets the minimum value of the slider range. |
| **Max** | Sets the maximum value of the slider range. |
| **Power** | Defines the exponent used for non-linear response between **Min** and **Max** when **Slider Type** is set to **Power**. |

### Enum

Additional parameters available when you set the Float property **Mode** to **Enum**.

| Parameter | Description |
| --- | --- |
| **Enum Type** | Selects the source used to populate the dropdown entries in the material Inspector.  The options are:  * **Explicit Values**: Uses a list of **Entries** you specify directly in this interface. * **Type Reference**: Uses a **C# Enum Type** reference that contains predefined entries. |
| **Entries** | Defines the list of dropdown entries when **Enum Type** is set to **Explicit Values**.  Use **+** or **-** to add or remove entries. Each entry requires the following parameters:  * **Name**: Sets the label displayed in the dropdown in the material Inspector. * **Value**: Sets the numeric value applied to the Float property when its **Name** is selected.  **Note**: Supports up to 7 entries. If you need more, use **Type Reference**. |
| **C# Enum Type** | Specifies the existing enum type to reference when **Enum Type** is set to **Type Reference**. Enter the full type path with namespace, for example: `UnityEngine.Rendering.BlendMode`. |

## Vector 2

Defines a **Vector 2** value. Displays a **Vector 4** input field in the material inspector, where the z and w components are not used.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 2D vector value (Vector2). |

## Vector 3

Defines a **Vector 3** value. Displays a **Vector 4** input field in the material inspector, where the w component is not used.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 3D vector value (Vector3). |

## Vector 4

Defines a **Vector 4** value. Displays a **Vector 4** input field in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 4D vector value (Vector4). |

## Color

Defines a **Color** value. If the Property Inspector displays **Main Color**, this is the [Main Color](https://docs.unity3d.com/Manual/SL-Properties.html) for the shader. To select or deselect this node as the **Main Color**, right-click it in the graph or Blackboard and select **Set as Main Color** or **Clear Main Color**. Corresponds to the [`MainColor`](https://docs.unity3d.com/Manual/SL-Properties.html) ShaderLab Properties attribute.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). |
| **Mode** | Selects the color input mode.  The options are:  * **Default**: Allows to select a standard sRGB color. * **HDR**: Allows to select an HDR color and sets its intensity from -10 to 10 exposure stops. |

**Note:** In versions prior to 10.0, Shader Graph didn't correct HDR colors for the project colorspace. Version 10.0 corrected this behavior. HDR color properties that you created with older versions maintain the old behavior, but you can use the [Graph Inspector](Internal-Inspector.html) to upgrade them. To mimic the old behavior in a gamma space project, you can use the [Colorspace Conversion Node](Colorspace-Conversion-Node.html) to convert a new HDR **Color** property from **RGB** to **Linear** space.

## Boolean

Defines a **Boolean** value. Displays a **ToggleUI** field in the material inspector. Note that internally to the shader this value is a **Float**. The **Boolean** type in Shader Graph is merely for usability.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A boolean value. |

## Gradient

Defines a constant **Gradient**.

Parameters specific to Gradient properties in addition to the [common parameters](#common-parameters):

| Parameter | Description |
| --- | --- |
| **Default Value** | Displays the **HDR Gradient Editor** with selectable modes. The options are:  * **Blend (Classic)**: Creates transitions based on traditional HDR handling. * **Blend (Perceptual)**: Creates human vision–based gradient transitions in HDR. * **Fixed**: Creates a discrete gradient with fixed steps. |

**Note:** The **Promote to final Shader** parameter is not available for this property.

## Texture 2D

Defines a [Texture 2D](https://docs.unity3d.com/Manual/class-TextureImporter.html) value. Displays an object field of type [Texture](https://docs.unity3d.com/Manual/class-TextureImporter.html) in the material inspector. If the Property Inspector displays **Main Texture**, this is the `Main Texture` for the shader. To select or deselect this node as the `Main Texture`, right-click on it in the graph or Blackboard and select **Set as Main Texture** or **Clear Main Texture**. Corresponds to the [`MainTexture`](https://docs.unity3d.com/Manual/SL-Properties.html) ShaderLab Properties attribute.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A texture asset reference. |
| **Mode** | Defines the fallback texture Unity uses when none is provided.  The options are:  * **White**: Sets a solid white (1,1,1) texture to ensure full-intensity sampling. * **Black**: Sets a solid black (0,0,0) texture to yield zero contribution. * **Grey**: Sets a mid-grey in sRGB (~0.5) as a neutral fallback. * **Normal Map**: Sets a flat normal value to keep surfaces flat without a normal texture. * **Linear Grey**: Sets a mid-grey in linear color space. * **Red**: Sets a solid red (1,0,0) texture, useful for data expected in the red channel. |
| **Use Tiling and Offset** | Toggles the property `NoScaleOffset` to enable manipulating scale and offset separately from other texture properties; see [SplitTextureTransformNode](Split-Texture-Transform-Node.html). A boolean value. |
| **Use TexelSize** | Uses the size of texels expressed in UV space. |

## Texture 2D Array

Defines a [Texture 2D Array](https://docs.unity3d.com/Manual/class-TextureImporter.html) value. Displays an object field of type [Texture 2D Array](https://docs.unity3d.com/Manual/class-TextureImporter.html) in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A texture asset reference. |

## Texture 3D

Defines a [Texture 3D](https://docs.unity3d.com/Manual/class-TextureImporter.html) value. Displays an object field of type [Texture 3D](https://docs.unity3d.com/Manual/class-TextureImporter.html) in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A texture asset reference. |

## Cubemap

Defines a [Cubemap](https://docs.unity3d.com/Manual/class-Cubemap.html) value. Displays an object field of type [Texture](https://docs.unity3d.com/Manual/class-TextureImporter.html) in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A cubemap asset reference. |

## Virtual Texture

Defines a [Texture Stack](https://docs.unity3d.com/Manual/svt-use-in-shader-graph.html), which appears as object fields of type [Texture](https://docs.unity3d.com/Manual/class-TextureImporter.html) in the Material Inspector. The number of fields corresponds to the number of layers in the property.

| Parameter | Description |
| --- | --- |
| **Layers** | Manages the collection of layers in the stack.  The options are:  * **Add (+)**: Adds a new layer. * **Remove (−)**: Removes the selected layer.   Select the active layer to edit its parameters. |
| **Layer Name** | Displays the user-defined name for the selected layer. |
| **Layer Reference** | Defines the internal identifier used to reference the selected layer. |
| **Layer Texture** | Assigns the default texture asset for the selected layer. |
| **Layer Texture Type** | Specifies the expected data type for the selected layer’s texture, which determines import settings and sampling behavior, such as sRGB vs Linear and normal map decoding.  The options are:  * **Normal tangent space**: Encodes per-texel normals relative to the mesh’s tangent basis so surface detail follows UVs and local orientation. * **Normal object space**: Preserves per-texel normals in object coordinates. |

**Note:** The **Use Custom Binding** parameter isn't available for this property.

**Note:** The **Promote to final Shader** parameter is not available for this property.

## Matrix 2

Defines a Matrix 2. Matrices do not display in the **Inspector** window of the material.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 2×2 matrix value (Matrix2). |

## Matrix 3

Defines a Matrix 3 value. Can't be displayed in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 3×3 matrix value (Matrix3). |

## Matrix 4

Defines a Matrix 4 value. Can't be displayed in the material inspector.

| Parameter | Description |
| --- | --- |
| **Default Value** | Sets the initial value of the [Property](https://docs.unity3d.com/Manual/SL-Properties.html). A 4×4 matrix value (Matrix4). |

## SamplerState

Defines a **SamplerState**.

Parameters specific to Float properties in addition to the [common parameters](#common-parameters):

| Parameter | Description |
| --- | --- |
| **Filter** | Specifies the texture filtering mode used when sampling. The options are:  * **Linear**: Sets bilinear filtering within mip levels for smoother results, at the cost of potential blur. * **Point**: Sets nearest-neighbor sampling for a crisp, pixelated look. * **Trilinear**: Sets bilinear filtering with interpolation between mip levels for smoother transitions. |
| **Wrap** | Specifies how UVs outside the [0–1] range are handled. The options are:  * **Repeat**: Tiles the texture infinitely. * **Clamp**: Clamps to edge texels with no tiling. * **Mirror**: Tiles by mirroring each repeat. * **MirrorOnce**: Mirrors once, then clamps. |
| **Aniso** | Specifies the anisotropic filtering level to improve texture clarity at grazing angles. The options are:  * **None**: Disables anisotropic filtering. * **x2**: Applies a low level for higher performance. * **x4**: Applies a moderate level. * **x8**: Applies a high level. * **x16**: Applies the maximum level for best quality at lower performance. |

## Dropdown

Defines a **Dropdown**. This property is available only in sub graphs.

Parameters specific to Dropdown properties in addition to the [common parameters](#common-parameters):

| Parameter | Description |
| --- | --- |
| **Default Value** | Selects the default Entry that you want Shader Graph to select on your property. |
| **Entries** | Adds a corresponding input port to the node for each entry.  The options are:  * **Add to the list (+)**: Adds a new option to your dropdown. * **Remove selection from the list (-)**: Removes the selected entry from the list. |

**Note:** The **Promote to final Shader** parameter is not available for this property.