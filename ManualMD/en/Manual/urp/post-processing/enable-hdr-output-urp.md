# Enable HDR Output in URP

HDR Output in URP requires you to enable multiple properties to work. These properties allow URP to render **scenes** correctly with **HDR** Rendering and then output the HDR values to a display.

To activate HDR output, follow these steps.

1. Locate the [URP Asset](./../universalrp-asset.md) in the Project window under **Assets** > **Settings**.
2. Navigate to **Quality** > **HDR** and enable the checkbox to enable **HDR**.
3. Navigate to **Edit** > **Project Settings** > **Player** > **Other Settings** and enable the following settings:

   * **Allow HDR Display Output**
   * **Use HDR Display Output**

**Note**: Only enable **Use HDR Display Output** if you need the main display to use HDR Output.

If you switch to a URP Asset that does not have HDR enabled, URP disables HDR Output until you change to a URP Asset with HDR enabled.

**Note**: When HDR Output is active the color grading mode is HDR by default, even if a different Color Grading Mode is active in the URP Asset.