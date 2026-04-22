# Use bar sliders

A bar slider lets you allocate sections of a whole. For example, the [level of detail (LOD) Group component](class-LODGroup.md) uses a bar slider to define transitions between the levels. All levels together make up 100%.

To change allocations, drag the divider between two sections left or right. This changes both sections at the same time, but doesn’t change any other section.

Each section in a bar slider shows how far into the whole it reaches, indicated as a percentage. That percentage is the end point for the section, not the total allocation of that section. The last section is always 100%.

![The LOD Group bar, showing the percentage at the point where the divider last stopped.](../uploads/Main/InspectorLODBarSlider.png)

The LOD Group bar, showing the percentage at the point where the divider last stopped.

## Additional resources

* [LOD Group component](class-LODGroup.md)
* [LODGroup](../ScriptReference/LODGroup.md)