# Artistic nodes | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Artistic-Nodes.html)

# Artistic nodes

Adjust colors, blend layers, filter images, mask regions, manipulate normal maps, and convert color spaces.

## Adjustment

| **Topic** | **Description** |
| --- | --- |
| [Channel Mixer](Channel-Mixer-Node.html) | Controls the amount each of the channels of input In contribute to each of the output channels. |
| [Contrast](Contrast-Node.html) | Adjusts the contrast of input In by the amount of input Contrast. |
| [Hue](Hue-Node.html) | Offsets the hue of input In by the amount of input Offset. |
| [Invert Colors](Invert-Colors-Node.html) | Inverts the colors of input In on a per channel basis. |
| [Replace Color](Replace-Color-Node.html) | Replaces values in input In equal to input From to the value of input To. |
| [Saturation](Saturation-Node.html) | Adjusts the saturation of input In by the amount of input Saturation. |
| [White Balance](White-Balance-Node.html) | Adjusts the temperature and tint of input In by the amount of inputs Temperature and Tint respectively. |

## Blend

| **Topic** | **Description** |
| --- | --- |
| [Blend](Blend-Node.html) | Blends the value of input Blend onto input Base using the blending mode defined by parameter Mode. |

## Filter

| **Topic** | **Description** |
| --- | --- |
| [Dither](Dither-Node.html) | Dither is an intentional form of noise used to randomize quantization error. It is used to prevent large-scale patterns such as color banding in images.. |

## Mask

| **Topic** | **Description** |
| --- | --- |
| [Channel Mask](Channel-Mask-Node.html) | Masks values of input In on channels selected in dropdown Channels. |
| [Color Mask](Color-Mask-Node.html) | Creates a mask from values in input In equal to input Mask Color. |

## Normal

| **Topic** | **Description** |
| --- | --- |
| [Normal Blend](Normal-Blend-Node.html) | Blends two normal maps defined by inputs A and B together. |
| [Normal From Height](Normal-From-Height-Node.html) | Creates a normal map from a height map defined by input Texture. |
| [Normal Strength](Normal-Strength-Node.html) | Adjusts the strength of the normal map defined by input In by the amount of input Strength. |
| [Normal Unpack](Normal-Unpack-Node.html) | Unpacks a normal map defined by input In. |

## Utility

| **Topic** | **Description** |
| --- | --- |
| [Colorspace Conversion](Colorspace-Conversion-Node.html) | Returns the result of converting the value of input In from one colorspace space to another. |