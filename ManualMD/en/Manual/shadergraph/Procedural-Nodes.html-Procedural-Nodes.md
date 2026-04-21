# Procedural nodes | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Procedural-Nodes.html)

# Procedural nodes

Generate patterns, noise textures, and customizable geometric shapes procedurally using UV input.

| **Topic** | **Description** |
| --- | --- |
| [Checkerboard](Checkerboard-Node.html) | Generates a checkerboard of alternating colors between inputs Color A and Color B based on input UV. |

## Noise

| **Topic** | **Description** |
| --- | --- |
| [Gradient Noise](Gradient-Noise-Node.html) | Generates a gradient, or Perlin, noise based on input UV. |
| [Simple Noise](Simple-Noise-Node.html) | Generates a simple, or Value, noise based on input UV. |
| [Voronoi](Voronoi-Node.html) | Generates a Voronoi, or Worley, noise based on input UV. |

## Shape

| **Topic** | **Description** |
| --- | --- |
| [Ellipse](Ellipse-Node.html) | Generates an ellipse shape based on input UV at the size specified by inputs Width and Height. |
| [Polygon](Polygon-Node.html) | Generates a regular polygon shape based on input UV at the size specified by inputs Width and Height. The polygon's amount of sides is determined by input Sides. |
| [Rectangle](Rectangle-Node.html) | Generates a rectangle shape based on input UV at the size specified by inputs Width and Height. |
| [Rounded Polygon](Rounded-Polygon-Node.html) | Generates a rounded polygon shape based on input UV at the size specified by inputs Width and Height. The input Sides specifies the number of sides, and the input Roundness defines the roundness of each corner. |
| [Rounded Rectangle](Rounded-Rectangle-Node.html) | Generates a rounded rectangle shape based on input UV at the size specified by inputs Width and Height. The input Radius defines the radius of each corner. |