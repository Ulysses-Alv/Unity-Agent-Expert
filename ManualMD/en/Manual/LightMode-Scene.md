# Set the Lighting Mode of a scene

1. Select the [Lighting Settings Asset](class-LightingSettings.md) for the **Scene**.
2. In the **Inspector**, navigate to **Mixed Lighting**.
3. Use the drop-down menu to set the **Lighting Mode**.

If your project uses the High-Definition **Render Pipeline** (HDRP), follow the instructions in the [HDRP Shadowmask Lighting Mode](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/Shadows-in-HDRP.html%23shadowmasks) documentation instead.

## Match shadow colors in Subtractive mode

When you set a Scene’s Lighting Mode to Subtractive, Unity displays the **Realtime Shadow Color** property in the Lighting window. Unity uses this color when it combines real-time shadows with baked shadows. Change this value to approximate the color of the indirect lighting in your Scene, to better match the real-time shadows with the baked shadows.