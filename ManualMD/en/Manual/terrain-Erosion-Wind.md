# Wind erosion

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

The Wind erosion tool simulates the wind lifting and carrying sediment. The sediment can travel long distances in any direction, though it settles gradually along its course, not all at once at the end.

The Wind Erosion tool takes into account:

* The way the environment itself impacts the wind’s ability to erode it. For example, drag impacts the wind’s decay, so that it slows down as it moves through the **terrain**. A slower wind carries less sediment.
* The wind’s ability to lift and carry different sediment types: Heavy, wet soil isn’t lifted as easily or carried as far, and a slow wind doesn’t lift much soil and doesn’t carry it far.
* The way different sediment types settle. The thermal smoothing property accounts for the steepness of the slope the sediment can sustain.

## Access the tool

To erode the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Erosion** > **Wind**.

## Tool options

The Erosion tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

When you rotate the brush, the Wind Erosion Tool uses a smooth rotation jitter control, which is different from the standard rotation control. Instead of computing a random rotation value for each frame, the Wind Erosion Tool uses a smooth noise function to simulate gradual changes in wind direction.

## Wind Erosion Controls

If you change property values but the sculpting results don’t change, try:

* Increasing the **Simulation Scale** value.
* Decreasing the **Wind Speed** value.
* Decreasing the **Iterations** value.

These properties can have a significant impact on the erosion simulation, which hides the impact of all other properties.

### General properties

| **Property** | **Function** |
| --- | --- |
| **Simulation Scale** | How detailed the terrain’s height and features are after they’re eroded. Lower values create a smooth terrain and reduce a lot of its height; higher values keep more of the original finer details and height. This is a simple way to change the erosion simulation to work for most needs. You can use the other controls in this section for more precise simulation changes, but you can’t disable the **Simluation Scale** control - all other controls work within its context. |
| **Wind Speed** | The wind speed at the start of each brush stroke. Higher wind speeds create more erosion. Wind speeds drop as the wind moves across the terrain, based on the **Drag** and **Surface Reflection** properties. |
| **Advanced** > **Time Interval (dt)** | How much time passes during each iteration of the simulation. A low value gives each iteration a small impact, because the erosion doesn’t have a lot of time to act. A higher value gives each iteration greater impact, because the erosion runs for longer. When you change **Time Interval**, you need to think of changing **Iterations**, too. If both values are low, each brush stroke has a low visual impact. If the values are high, the effect can be so strong that all the other propeties have very little impact. |
| **Advanced** > **Iterations** | How many times the simulation runs for each brush stroke. The higher the value, the more each stroke impacts the terrain, because the simulation runs more times. |

### Advanced > Sediment and wind behavior

| **Property** | **Function** |
| --- | --- |
| **Particulate Suspension Rate** | How much sediment to lift into the wind at the start of the brush stroke. |
| **Particulate Deposit Rate** | How much sediment to lay down as the brush moves. Wind can carry sediment over long distances in any direction. To simulate this behavior, you can increase the **Particulate Suspension Rate**, which controls the amount of sediment that the wind can carry, and decrease the **Particulate Deposit Rate**, which controls the amount of sediment that the wind can deposit. This combination removes more sediment than it deposits, which creates a net loss of sediment, simulating a situation where the sediment is carried far away, perhaps past the edge of the terrain. |
| **Slope Factor** | How much the Terrain slope affects sediment suspension. Low values allow the wind to lift sediment from both sharp and flat features. High values protect the flat features but still erode the sharp features. |
| **Flow Rate** | How fast sediment moves. Higher values shorten the sediment’s flight time, which allows more of it to reach the end of the brush stroke. Lower values deposit more sediment along the path. |
| **Drag** | How much the wind slows down from its initial speed as the brush moves over the terrain. Higher values slow the wind more, depositing more sediment along the path. Lower values maintain more of the wind’s initial speed, lifting more sediment and carrying it farther. |
| **Surface Reflection** | How much speed the wind maintains after it hits a steep slope. Higher values maintain the wind speed so erosion continues for longer. |
| **Diffusion Rate** | The rate at which suspended sediment mixes in the air, moving from areas of low concentration to areas of high concentration. This property only has an effect where different where different areas of the brush pick up different amounts of sediment. For example, a large brush with a high **Slope Factor** going along the join between a flat area and a hill has different results for different diffusion rates, because it doesn’t lift sediment from the flat portion. High diffusion rates spread the sediment throughout the brush. Low diffusion rates keep the sediment in the part of the brush that lifted it. |
| **Abrasiveness** | How much sediment erodes the terrain as it flows over it. A high value for **Simulation Scale** (see above) can hide the differences in abrasiveness. |
| **Viscosity** | How dense the air is. Dense air slows the wind down more. This is different to **Drag**, where the wind slows because of contact with the terrain. |

### Advanced > Thermal Smoothing

Smooths the terrain where the soil type is steeper than its resting angle.

| **Property** | **Function** |
| --- | --- |
| **Physical Material Presets** | Choose a soil type to use realistic resting angles. Refer to Physical Material Presets above for more information. |
| **Iterations** | The number of times the erosion is applied for each brush stroke. A low values gives each iteration a small impact. |
| **Time Interval (dt)** | How much time passes during each iteration of the simulation. A low value gives each iteration a small impact. |
| **Resting Angle** | The maximum slope angle for the material. Lower resting angles limit the details the erosion can create, because it can’t create a lot of variation when the slopes are low. |

## Additional resources

* [Considerations for working with the Erosion tool](terrain-Erosion-Considerations.md)
* [Thermal erosion](terrain-Erosion-Thermal.md)
* [Hydraulic erosion](terrain-Erosion-Hydraulic.md)