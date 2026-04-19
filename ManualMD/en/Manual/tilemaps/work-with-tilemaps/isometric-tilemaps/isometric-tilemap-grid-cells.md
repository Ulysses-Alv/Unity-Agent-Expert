# Isometric tilemap grid cells

The **Cell Size** and **Cell Layout** properties of the Grid object the **Tilemap** is the child of determine some of the characteristics of the tilemap’s appearance. For more information on these properties, refer to the [Grid component reference](../../grid-reference.md).

When the Cell Layout property is set to Isometric Z as Y, tiles on the tilemap have their Z-axis value is added to its Y-axis value. This causes the Tile to be visually offset along the Y-axis by that value, creating the illusion of height as the Tiles are placed on the Tilemap.

Isometric Tilemaps use either **dimetric projection** or true **isometric projection** parallel projection angles. For more information about the two forms of projection, please refer to this [article](https://en.wikipedia.org/wiki/Isometric_computer_graphics) for further details.

Changing the Cell Size of the Grid component changes the size of angles that make up each Cell, which affects the type of projection being simulated. By default, Cell Size of the Isometric Cell Layout is **(1, 0.5, 1)** which simulates **dimetric projection** angles. True **isometric projection** instead uses a Y value of **0.57735**.