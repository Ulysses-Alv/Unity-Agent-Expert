# Rendering trails for moving GameObjects

The **Trail Renderer** component renders a trail of polygons behind a moving **GameObject**, over time. This can be used to give an emphasized feeling of motion to a moving object, or to highlight the path or position of moving objects.

The Trail Renderer uses the same algorithm for trail rendering as the [Line Renderer](class-LineRenderer.md).

## Distance between trail points

The **Min Vertex Distance** value determines how far in world units the GameObject to which the trail applies must travel before a new segment is added to the trail. Low values like 0.1 create trail segments more often, creating smoother trails. Higher values like 1.5 create segments that are more jagged in appearance. Additionally, wide trails may exhibit visual artifacts when the vertices are very close together and the trail changes direction significantly over a short distance.

For performance reasons, it is best to use the largest possible value that achieves the effect you are trying to create.