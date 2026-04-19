# Set shadow distance in a scene

Use the **Shadow Distance** property to determine the distance from the **Camera** up to which Unity renders real-time shadows.

Shadows from **GameObjects** become less noticeable the farther the GameObjects are from the Camera. This is both because the shadows appear smaller on the screen, and because distant GameObjects are usually not the focus of attention. You can take advantage of this effect by disabling real-time shadow rendering for distant GameObjects. This saves on wasted rendering operations, and can improve runtime performance. Additionally, the **Scene** often looks better without distant shadows.

If the current Camera Far Plane is closer than the Shadow Distance, Unity uses the Camera Far Plane instead of the Shadow Distance.

To disguise missing shadows beyond the Shadow Distance, you can use visual effects such as fog.

## Setting the Shadow Distance

In the Built-in **Render Pipeline**, set the Shadow Distance property in your Project’s [Quality Settings](class-QualitySettings.md).

In the Universal Render Pipeline (URP), set the Shadow Distance property in the [Universal Render Pipeline Asset](urp/universalrp-asset.md).

In the High Definition Render Pipeline (HDRP), set the Shadow Distance property for each [Volume](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Volumes.html).

## Shadow Distance and Shadowmask Lighting Mode

If your Scene uses the [Shadowmask Lighting Mode](lighting-mode.md#baked-indirect), Unity renders shadows from [Mixed Lights](lighting-mode-landing.md) beyond the Shadow Distance, using either [Light Probes](LightProbes.md) or a shadow mask Texture. You can configure how Unity renders shadows beyond the Shadow Distance.