# Considerations when using erosion tools

This section introduces considerations for:

* The purpose of the Erosion tools.
* How the Erosion tools interact with the **terrain**.
* Getting realistic results.

**Note:** The Erosion tools are only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

## The purpose of the Erosion tools

Erosion creates varied, realistic terrains. For example, use it to:

* Add variety to a terrain that’s too smooth and uniform.
* To create complex riverbeds and riverbanks.
* To prevent slopes too steep for the material they’re made of.
* Make civilizations look old and worn down, even abandoned.

[Noise](terrain-Sculpt-Noise-Landing.md) is another way to add variety to the terrain. The difference between the two tools is that erosion is directional: it moves sediment from point to point. Noise stamps the terrain with height variation.

## Interaction with the terrain

Erosion tools interact with the terrain in the following ways.

### Simulation scale

The simulation scale controls the detail level of the simulation. Higher values produce a more complex terrain, while lower values produce a smooth terrain. In other words, a lower value simulates sustained, heavy erosion that has smoothed the surface.

The **Simulation Scale** property is always active, and is often the only property you need to change to get the terrain look you want. When you change other properties, be aware that they always work within the context of the **Simulation Scale** property. In other words, if your **Simulation Scale** value is low, which means the terrain loses a lot of its details and height with every brush stroke, the other properties have little impact.

The **Simulation Scale** property isn’t available for Thermal Erosion. Instead, Thermal Erosion relies on the **Resting Angle** property to control its impact.

### Erosion maximum depth

Like all sculpting tools, Erosion cannot go below the terrain’s base height. To create dry river beds or other features that are below the rest of the terrain, use the Set Height tool to set a starting height for sculpting. For more information, refer to [Set the height of an area or tile](terrain-SetHeight.md).

### Resolution recommendation

Erosion adds fine details to the terrain, and looks best when the **heightmap** resolution is 1025 or greater. For more information, refer to [Terrain Settings reference](terrain-OtherSettings.md).

### Mouse speed

The speed at which you move the mouse affects the erosion. The erosion is stronger - changes the terrain more - with a slow mouse movement.

## Considerations for realistic erosion

To get realistic erosion results, consider the following.

### Erosion intervals and iterations

The **Iteration** property of each Erosion tool controls how many times the simulation runs for each brush stroke. Within each iteration, the **Interval** property specifies the duration the simulation covers. The combination of these two properties controls the impact of each brush stroke. Use lower impact brush strokes to sculpt a particular look, such as an exact shape for a riverbank. Use a higher impact to create variation across the terrain when you’re not worried about specific details.

### Wind and hydraulic erosion together

In nature, wind and hydraulic erosion interact not just with the terrain, but with each other. For example, think of how the prevaling wind impacts waves, rainfall, and how trees grow (trees reduce erosion).

### Erosion and vegetation

Wind and hydraulic erosion have the greatest impact where vegetation is sparse, and hydraulic erosion also has the greatest impact in areas of steep slopes. However, well covered, shallow slopes still have soil creep, which makes trees leans. Trees also grow leaning away from strong, prevalent winds.

The Erosion tool moves trees and other objects to match the terrain height, but only detail meshes and grass adjust to the angle of the terrain. To have trees that lean as if erosion pulled them after they grew, adjust the trees themselves.

### Grass and detail meshes

Grass and detail meshes always stay on top of the terrain surface, but at the same location. In other words, if you erode the terrain, the grass and details will match its new elevation and angle, but they aren’t carried forward in the direction of the erosion.

### Textures

The erosion tools don’t move the textures with the terrain. To get the best results, use the Erosion tools first, then paint the textures.

### Sediment buildup against objects

The erosion tools don’t build up sediment against **3D objects**. For example, when eroding a river, the sediment doesn’t build up against rocks in the riverbed. To simulate this, use other height tools to create the drift against the object.

### Thermal smoothing for wind and hydraulic erosion

The Wind and Hydraulic erosion tools have a Thermal Smoothing control. This control stops the erosion from creating unnaturally steep slopes.

## Additional resources

* [Hydraulic erosion](terrain-Erosion-Hydraulic.md)
* [Thermal erosion](terrain-Erosion-Thermal.md)
* [Wind erosion](terrain-Erosion-Wind.md)