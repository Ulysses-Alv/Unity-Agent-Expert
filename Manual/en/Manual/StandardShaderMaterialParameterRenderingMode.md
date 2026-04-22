# Set the Rendering Mode in the Standard Shader using a script

When you change the **Rendering Mode**, Unity applies a number of changes to the Material. There is no single C# API to change the Rendering Mode of a Material, but you can make the same changes in your code.

To see the changes that Unity makes when you change the **Rendering Mode**:

1. Download the source code for Unity’s built-in **shaders**. See [Make your own shader](https://docs.unity3d.com/Manual/StandardShaderMakeYourOwn.html) for instructions.
2. Open the file *StandardShaderGUI.cs*.
3. Locate the area of the file that looks like this, and observe the changes for each **Rendering Mode**.

```
switch (blendMode)
        {
            case BlendMode.Opaque:
               // Changes associated with Opaque Rendering Mode are here
                break;
            case BlendMode.Cutout:
                // Changes associated with Cutout Rendering Mode are here
                break;
            case BlendMode.Fade:
                // Changes associated with Fade Rendering Mode are here
                break;
            case BlendMode.Transparent:
                // Changes associated with Transparent Rendering Mode are here
                break;
        }
```

![The helmet visor in this image is rendered using the Transparent mode, because it is supposed to represent a real physical object that has transparent properties. Here the visor is reflecting the skybox in the scene. ](../uploads/Main/StandardShaderTransparencySkyBoxReflection.jpg)

The helmet visor in this image is rendered using the Transparent mode, because it is supposed to represent a real physical object that has transparent properties. Here the visor is reflecting the skybox in the scene.

![These windows use Transparent mode, but have some fully opaque areas defined in the texture (the window frames). The Specular reflection from the light sources reflects off the transparent areas and the opaque areas.](../uploads/Main/StandardShaderTransparentWindow.jpg)

These windows use Transparent mode, but have some fully opaque areas defined in the texture (the window frames). The Specular reflection from the light sources reflects off the transparent areas and the opaque areas.

![The hologram in this image is rendered using the Fade mode, because it is supposed to represent an opaque object that is partially faded out.](../uploads/Main/StandardShaderFadeHologram.jpg)

The hologram in this image is rendered using the Fade mode, because it is supposed to represent an opaque object that is partially faded out.

![The grass in this image is rendered using the Cutout mode. This gives clear sharp edges to objects which is defined by specifying a cut-off threshold. All parts of the image with the alpha value above this threshold are 100% opaque, and all parts below the threshold are invisible. To the right in the image you can see the material settings and the alpha channel of the texture used.](../uploads/Main/StandardShaderCutoutGrassExample.jpg)

The grass in this image is rendered using the Cutout mode. This gives clear sharp edges to objects which is defined by specifying a cut-off threshold. All parts of the image with the alpha value above this threshold are 100% opaque, and all parts below the threshold are invisible. To the right in the image you can see the material settings and the alpha channel of the texture used.