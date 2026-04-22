# Toggle shader keywords in the Editor

To toggle **shader** keywords in the Editor, add a material property for the keyword. This allows you to set different keyword states for different materials without needing to use code.

Follow these steps:

1. Add `_ON` to the end of the keyword you declare. For example, change `_RED` to `_RED_ON`. Make sure you update your branching code to use the new name.
2. Make sure you or Unity declare state for when the keyword is off. For more information, refer to [Branch when all keywords in a set are disabled](SL-MultipleProgramVariants-make-conditionals.md#branch-when-all-keywords-in-a-set-are-disabled).

   For example, if you use `shader_feature` to declare only one keyword, Unity creates the ‘off’ state automatically.
3. Add an integer [material property](SL-Properties.md) with the attribute `[Toggle]`, and the name of the keyword without `_ON`. For example:

   ```
       Properties
       {
           [Toggle] _RED ("Make red", Integer) = 0
       }
   ```

Now when you create a material that uses the shader, the [Inspector window of the material](class-Material.md) has a checkbox that enables or disables the keyword.

To disable the property by default, use `[ToggleOff]` instead of `[Toggle]`.

### Toggle keyword sets

To toggle keywords in a set so they’re mutually exclusive, add a property with the attribute `KeywordEnum`. Unity creates a dropdown with all the keywords. When you enable one keyword in the set, Unity disables the others.

Follow these steps:

1. Declare the keyword set in the format `SETNAME_KEYWORDNAME`. For example:

   ```
   #pragma shader_feature _COLOR_RED _COLOR_GREEN _COLOR_BLUE
   ```
2. Add a material property with the attribute `[KeywordEnum]` and the name of the keyword set. Pass in the names of the keywords. For example:

   ```
   Properties
   {
       [KeywordEnum(RED, GREEN, BLUE)] _COLOR ("Color", Integer) = 0
   }
   ```

For more information and examples, refer to the [MaterialPropertyDrawer](../ScriptReference/MaterialPropertyDrawer.md) API.

## Example

The following example adds a property called `_RED`, which enables and disables the keyword `_RED_ON`.

```
Shader "Unlit/ToggleRed"
{
    Properties
    {
        [Toggle] _RED ("Make red", Integer) = 0
    }
    SubShader
    {
        Tags { "RenderType"="Opaque" }

        Pass
        {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag

            // Add the keyword RED_ON, for when the toggle is on
            // Unity automatically adds a keyword for when the toggle is off 
            #pragma shader_feature _RED_ON

            #include "UnityCG.cginc"

            struct appdata
            {
                float4 vertex : POSITION;
            };

            struct v2f
            {
                float4 vertex : SV_POSITION;
            };

            v2f vert (appdata v)
            {
                v2f o;
                o.vertex = UnityObjectToClipPos(v.vertex);
                return o;
            }

            fixed4 frag (v2f i) : SV_Target
            {
                // Return red if the toggle is on
                #if _RED_ON
                    return fixed4(1, 0, 0, 1);
                #else
                    return fixed4(0, 0, 0, 1);
                #endif
            }
            ENDHLSL
        }
    }
}
```

## Additional resources

* [Adding material properties to shaders](writing-shader-change-properties.md)
* [Properties block reference in ShaderLab](SL-Properties.md)