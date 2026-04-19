# Lens Flare component reference

[Switch to Scripting](../ScriptReference/LensFlare.md "Go to LensFlare page in the Scripting Reference")

Explore the properties in the **Lens Flare** component to create a lens flare with a [Flare asset](class-Flare.md).

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Flare** | The [Flare asset](class-Flare.md) to render.  **Note**: The other properties in the Lens Flare Inspector are stored in this Flare asset. They are not stored in the Lens Flare component. |
| **Color** | The color to tint the Flare to. You can use this to make the lens flare better fit the **Scene**’s mood. |
| **Brightness** | The size and brightness the lens flare appears. If you use a very bright lens flare, make sure its direction fits with your Scene’s primary light source. Can only affect properties that Flare has enabled. |
| **Fade Speed** | The speed at which Unity fades the lens flare in and out. |
| **Ignore Layers** | A mask that determins which Layers can see the lens flare. |
| **Directional** | Indicates whether Unity orients the Flare along the positive z-axis of the **GameObject**. It appears as if it was infinitely far away, and does not track the GameObject’s position, only the direction of its z-axis. |

LensFlare