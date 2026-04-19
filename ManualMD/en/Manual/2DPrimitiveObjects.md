# Types of 2D primitive GameObjects

Unity has import support for different image formats, such as .png or Adobe’s .psd, which gives you additional options when creating and preparing your 2D assets for your project. But if you want to create a quick prototype, Unity provides 2D primitive **GameObjects** to help you build your project, without needing to prepare and import your own assets.

This page details the dimensions and common uses of the available 2D primitive options.

**Important:** Your project must have the [2D Sprite package](https://docs.unity3d.com/Packages/com.unity.2d.sprite@1.0/manual/index.html) installed to enable the following 2D primitive options. The 2D Sprite package is automatically installed when you create a project with the [2D project template](https://docs.unity3d.com/hub/manual/project-create.html) selected, or you can install the 2D Sprite package via the [Package Manager](Packages.md).

## Create a 2D primitive GameObject

To create one of these preset 2D primitive GameObjects, go to **GameObject** > **2D Object** > **Sprites** or **Create** > **2D** > **Sprites** and select one of the available options:

* [Triangle](#triangle)
* [Square](#square)
* [Circle](#circle)
* [Capsule](#capsule)
* [Isometric Diamond](#iso-diamond)
* [Hexagon Flat-Top](#hex-flat-top)
* [Hexagon Point-Top](#hex-point-top)
* [9-Sliced](#9-sliced)

## Default sprite dimensions

The default sprite size of the 2D primitives is 256 X 256 **pixels**, with a pixels-per-unit (PPU) size of 256. This combination of dimensions and PPU value makes the sprite’s size equal to one **Unity unit** in the **scene**. There are two exceptions - the [Capsule](#capsule) primitive is 256 X 512 pixels (1:2 units) in size, and the [Isometric Diamond](#iso-diamond) primitive is 256 X 128 pixels (1:0.5 units).

## Triangle

![The Triangle 2D primitive GameObject.](../uploads/Main/2D-primitive-triangle.png)

The Triangle 2D primitive GameObject.

The **Triangle** 2D primitive is white isosceles triangle with a base measuring one Unity unit in size. You can use it as a placeholder for various elements in your scene, such as obstacles or parts of the user interface. You can add the [Polygon Collider 2D](2d-physics/collider/polygon-collider-2d-reference.md) component to the primitive to have it interact with other GameObjects and the 2D physics system.

## Square

![The Square 2D primitive GameObject.](../uploads/Main/2D-primitive-square.png)

The Square 2D primitive GameObject.

The **Square** 2D primitive is a white square that’s 1 X 1 Unity units in size. You can use it as a placeholder for various elements, such as obstacles or platforms. You can add the [Box Collider 2D](2d-physics/collider/box-collider-2d-reference.md) component to the primitive to have it interact with other GameObjects and the 2D physics system. You can select the [9-Sliced](#9-sliced) option for a more scalable sprite that resizes dynamically.

## Circle

![The Circle 2D primitive GameObject.](../uploads/Main/2D-primitive-circle.png)

The Circle 2D primitive GameObject.

The **Circle** 2D primitive is a white circle that’s one Unity unit in diameter. You can use it as a placeholder for various elements in your scene, such as obstacles or props. You can add the [Circle Collider 2D](2d-physics/collider/circle-collider-2d-reference.md) to the primitive to have it interact with other GameObjects and the 2D physics system.

## Capsule

![The Capsule 2D primitive GameObject.](../uploads/Main/2D-primitive-capsule.png)

The Capsule 2D primitive GameObject.

The **Capsule** 2D primitive is a white capsule that’s 1 X 2 units in size. You can use it as a placeholder for various elements in your scene, such as an obstacle, prop, or a stand-in for a character. You can add a [Capsule Collider 2D](2d-physics/collider/capsule-collider/capsule-collider-2d-reference.md) to the primitive to have it interact with other GameObjects and the 2D physics system.

## Isometric Diamond

![The Isometric Diamond 2D primitive GameObject.](../uploads/Main/2D-primitive-isodiamond.png)

The Isometric Diamond 2D primitive GameObject.

The **Isometric Diamond** 2D primitive is a white diamond-shaped sprite that’s 1 X 0.5 units in size. This sprite is often used as a placeholder for Tiles placed on [Isometric Tilemaps](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md). The pixels at the top and bottom of this sprite are slightly cropped to improve tiling.

## Hexagon Flat-Top

![The Hexagon Flat-Top 2D primitive GameObject.](../uploads/Main/2D-primitive-hex-flattop.png)

The Hexagon Flat-Top 2D primitive GameObject.

The **Hexagon Flat-Top** 2D primitive is a regular hexagon that’s one unit wide, with its flat sides oriented towards the top and bottom. This sprite is often used as a placeholder for Tiles placed on [Hexagonal Flat-Top Tilemaps](tilemaps/work-with-tilemaps/hexagonal-tilemaps.md). The pixels to the left and right of this sprite are slightly cropped to improve tiling.

## Hexagon Point-Top

![The Hexagon Point-Top 2D primitive GameObject.](../uploads/Main/2D-primitive-hex-pointtop.png)

The Hexagon Point-Top 2D primitive GameObject.

The **Hexagon Point-Top** 2D primitive is a regular hexagon that’s one unit tall, with its points oriented towards the top and bottom. This sprite is often used as a sprite placeholder for Tiles placed on a [Hexagonal Pointed-Top Tilemaps](tilemaps/work-with-tilemaps/hexagonal-tilemaps.md). The pixels to the top and bottom of this sprite are slightly cropped to improve tiling.

## 9-Sliced

![The 9-Sliced 2D primitive GameObject.](../uploads/Main/2D-primitive-9-sliced.png)

The 9-Sliced 2D primitive GameObject.

The **9-Sliced** 2D primitive is a white square with rounded corners that’s 1 X 1 units in size. This sprite has been [9-sliced](sprite/9-slice/9-slice-landing.md) with borders measuring 64 pixels on each side. This sprite is primarily used with the [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md) component’s **Sliced** and **Tiled** Draw Modes. You can use the 9-sliced sprite as a flexible placeholder for various elements in your scene and project (see [9-slicing sprites](sprite/9-slice/9-slice-sprite.md) for more information). You can add a [Box Collider 2D](2d-physics/collider/box-collider-2d-reference.md) with **Auto Tiling** enabled to have the sprite interact with other objects and the 2D physics system.

## Additional resources

* [2D game development quickstart guide](./2d-game-development-landing.md)
* [Sprites](sprite/sprite-landing.md)
* [Tilemap](tilemaps/work-with-tilemaps/tilemap-reference.md)
* [Physics Reference 2D](2d-physics/2d-physics.md)