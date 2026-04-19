# Hydraulic erosion

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

In hydraulic erosion, water carries sediment as it flows along the **terrain**. The movement of sediment creates channels and redistributes the sediment along those channels.

You can use the tool to simulate either temporary flow from rain or snow melt, or permanent and semi-permanent bodies of water. Remember that the tool deposits the sediment along the channel, it doesn’t remove it from the terrain.

## Considerations when using hydraulic erosion

When using the Hydraulic erosion tool, consider the following.

### Natural impacts on erosion

Hydraulic erosion takes into account:

* The way the environment impacts the water’s ability to erode it. For example, higher evaporation rates limits the erosion because as water loses volume it slows down, lift less sediment, and deposits more sediment.
* Water’s ability to lift and carry different sediment types.
* The way different sediment types settle. The thermal smoothing property accounts for the steepness of the slope the sediment can sustain.
* Sediment deposition on vertical or near vertical surfaces (the riverbank) is different to deposition on horizontal or near horizontal surfaces (the riverbed).

### Riverbeds and riverbanks

Hydraulic erosion builds river banks and river beds according to how the sediment flows with the water.

### Erosion around 3D objects

The Hydraulic erosion tool doesn’t build up sediment against a **3D object**, or maintain sediment below it. For example, if you have a 3D rock in a river bed, the tool erodes around and under it, but it doesn’t build up sediment against the rock.

### Water speed

When working with Hydraulic erosion, consider the following:

* The river’s width and depth impact the speed of the water, and therefore the erosion. A wider river moves more slowly than a narrow river moving the same volume of water, and therefore erodes less.
* To simulate the slower movement of water along the banks, use a brush with weaker edges, rather than a solid brush.

## Access the tool

To erode the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Erosion** > **Hydraulic**.

## Tool options

The Erosion tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

**Tip:** To invert the erosion effect, hold **Ctrl** (macOS: **Cmd**). For example, use inversion to add detail to a flat surface.

## Hydraulic Erosion Controls

If you change property values but the sculpting results don’t change, try:

* Increasing the **Simulation Scale** value.
* Decreasing the **Iterations** value.
* Decreasing the **Perception** value.
* Increasing the **Evaporation** value.

These four properties can have a significant impact on the erosion simulation, which hides the impact of all other properties.

### General properties

| **Property** | **Function** |
| --- | --- |
| **Simulation Scale** | How detailed the terrain’s height and features are after they’re eroded. Lower values create a smooth terrain and reduce a lot of its height; higher values keep more of the original finer details and height. This is a simple way to change the erosion simulation to work for most needs. You can use the other controls in this section for more precise simulation changes, but you can’t disable the **Simluation Scale** control - all other controls work within its context. |
| **Advanced** > **Time Interval (dt)** | How much time passes during each iteration of the simulation. A low value gives each iteration a small impact, because the erosion doesn’t have a lot of time to act. A higher value gives each iteration greater impact, because the erosion runs for longer. When you change **Time Interval**, you need to think of changing **Iterations**, too. If both values are low, each brush stroke has a low visual impact. If the values are high, the effect can be so strong that all the other parameters have little impact. |
| **Advanced** > **Iterations** | How many times the simulation runs for each brush stroke. The higher the value, the more each stroke impacts the terrain, because the simulation runs more times. |

### Advanced > Thermal Smoothing

Smooths the terrain where the soil type is steeper than its resting angle.

| **Property** | **Function** |
| --- | --- |
| **Physical Material Presets** | Choose a soil type to use realistic resting angles. Refer to Physical Material Presets above for more information. |
| **Iterations** | The number of times the erosion is applied for each brush stroke. A low values gives each iteration a small impact. |
| **Time Interval (dt)** | How much time passes during each iteration of the simulation. A low value gives each iteration a small impact. |
| **Resting Angle** | The maximum slope angle for the material. Lower resting angles limit the details the erosion can create, because it can’t create a lot of variation when the slopes are low. |

### Advanced > Water Transport

How much water each iteration of the simulation moves across the terrain, and how quickly it flows.

| **Property** | **Function** | **Real world considertions** |
| --- | --- | --- |
| **Precipitation Rate** | How much rainfall to apply during each iteration. | High precipitation with a strong flow (see **Flow Rate** below) has the strongest impact on the terrain. |
| **Evaporation Rate** | How much water to remove during each iteration. Evaporation means the amount of water eroding the terrain goes down as the brush moves. | Evaporation rates change along the course of a water channel. For example, shade and a deep, narrow channel lower precipitation compared to sunlight and shallow, wide channels. |
| **Flow Rate** | How quickly the water flows. Faster water lifts more sediment, so higher **Flow Rate** values have a larger impact on the terrain. Note, however, that `0` has the biggest impact. | Water channel flow rates aren’t steady. For example, the flow rate is higher where the angle is steeper, or where there is lower friction because of a smoother bed. |

### Advanced > Sediment Transport

How the sediment moves with the water.

| **Property** | **Function** | **Real world considertions** |
| --- | --- | --- |
| **Seidment Capacity** | The amount of sediment the water can carry, as a percentage of the available sediment. | Sediment capacity changes based on soil type along the water’s path. |
| **Seidment Deposition Rate** | The rate at which the simulation adds sediment back to the terrain. | Deposition rates change along the path of the water. For example, in turbulent areas, sediment doesn’t settle well. |
| **Seidment Dissolution Rate** | The rate at which the simulation removes sediment from the terrain. | Whereas **Sediment Capacity** changes how much the water can carry away, the **Sediment Dissolution Rate** is local damage to the terrain. |

### Advanced > Riverbank

How the water carves channels in the terrain. The result of carving depends building up or tearing down the banks, and how much sediment is deposited in the channel to create its bed.

| **Property** | **Function** | **Real world considertions** |
| --- | --- | --- |
| **Riverbank Despoit Rate** | The rate at which the simulation builds up the sides of the channels. | Higher values raise the banks and create deeper channels. |
| **Riverbank Dissolution Rate** | The rate at which the simulation breaks down the sides of the channels. | Higher values lower the banks and create shallow channels. |
| **Riverbed Deposit Rate** | The rate at which the simulation fills the channel by depositing water in the channel bed. | Higher values built up the bed and create shallow channel. |
| **Riverbed Dissolution Rate** | The rate at which the simulation carries sediment away from the channel. | Lower values protect the river bed from erosion and keep the channel shallow. |

## Additional resources

* [Considerations for working with the Erosion tool](terrain-Erosion-Considerations.md)
* [Thermal erosion](terrain-Erosion-Thermal.md)
* [Wind erosion](terrain-Erosion-Wind.md)