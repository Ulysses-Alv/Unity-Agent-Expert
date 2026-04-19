# Animate trees with Wind Zones

[Switch to Scripting](../ScriptReference/WindZone.md "Go to WindZone page in the Scripting Reference")

To create the effect of wind on your trees and **particle** systems, you can add one or more **GameObjects** with **Wind Zone** components. Trees within a wind zone bend in a realistic animated fashion, and the wind itself moves in pulses to create natural patterns of movement among the trees.

## Wind for Terrain

To create a Wind Zone GameObject, from the main menu, select **GameObject** > **3D Object** > **Wind Zone**.

### Wind modes

There are two modes for wind:

* In **Directional** mode, the wind affects the whole **Terrain** at once. This is useful for creating effects like the natural movement of the trees.
* In **Spherical** mode, the wind blows outwards within a sphere defined by the **Radius** property. This is useful for creating special effects like explosions (for a particle system) or a local eddy (for trees).

The **Main** property determines the overall strength of the wind. Use **Turbulence** for randomised variation.

The wind blows over the trees in pulses to create a more natural effect, where the wind speed isn’t a constant but rather falls and rises with gusts. You can control the strength of the pulses with **Pulse Magnitude**, and the time interval between them with **Pulse Frequency**.

### Tree-specific properties

Wind Zones work on an area of the terrain or an entire tile. However, different trees have different behaviors in the same wind conditions. Use the [Tree Editor](class-Tree.md) to control how the branches and leaves of each tree model react to wind.

You can’t control SpeedTree animation behavior in Unity; use SpeedTree to change wind animation and import your trees again.

### Using multiple Wind Zones

You can use more than one Wind Zone GameObject in your terrain. This creates different wind conditions in different areas of the terrain. For example, you can create a strong wind in one area and a gentle breeze in another.

You can use a Directional mode on the whole terrain and a Spherical mode to create local conditions (these are added to the Directional wind), or use only Spherical mode winds that together cover the whole terrain. Don’t use multiple Directional mode winds; a Directional wind covers the entire terrain, so using two Directional winds doesn’t create local conditions.

## Wind for particle systems

The main use of wind is to animate trees, but it can also affect particles in a particle system that uses the [External Forces module](PartSysExtForceModule.md). Refer to [Particle System](class-ParticleSystem.md) for more information.

To add a Wind Zone to [an existing particle system](class-ParticleSystem.md):

1. In the ****Inspector**** window, select the particle system GameObject.
2. From the main menu, select **Component** > **Effects** > **Wind Zone**.

## Wind for SpeedTree trees

Trees from SpeedTree rely on their own wind behavior, which is imported with them. Refer to [Games wind](https://docs9.speedtree.com/modeler/doku.php?id=windgames) for more information.

## Wind for grass

Wind Zones don’t apply to grass on the terrain. To animate grass, refer to [Terrain Settings reference](terrain-OtherSettings.md).

## Additional resources

* [Wind reference](terrain-Wind-Reference.md)
* [Terrain Settings reference](terrain-OtherSettings.md)
* [Create trees with Tree Editor](class-Tree.md)
* [SpeedTree Games wind mode](https://docs9.speedtree.com/modeler/doku.php?id=windgames)
* [Particle System](class-ParticleSystem.md)