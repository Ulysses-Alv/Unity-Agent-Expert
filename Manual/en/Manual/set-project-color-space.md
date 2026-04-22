# Set a project’s color space

The Unity Editor offers both linear and gamma workflows. The linear workflow has a color space crossover where [Textures](Textures.md) that were authored in gamma color space can be correctly and precisely rendered in linear color space. See documentation on [Linear rendering overview](color-spaces-landing.md) for more information about gamma and linear color space.

Textures tend to be saved in gamma color space, while Shaders expect linear color space. As such, when Textures are sampled in Shaders, the gamma-based values lead to inaccurate results. To overcome this, you can set Unity to use an sRGB sampler to cross over from gamma to linear sampling. This ensures a linear workflow with all inputs and outputs of a **Shader** in the correct color space, resulting in a correct outcome.

## Select the color space

Select the color space for your project with the following steps:

1. Go to **Edit** > **Project Settings**, then select the **Player** category.
2. Navigate to the **Other Settings**, open the **Rendering** section, and set the **Color Space** property to **Linear** or **Gamma**, depending on your preference.