# 2D game creation workflow

Before you create a 2D game, you need to decide on a [game perspective](./2d-game-perspective-reference.md) and an [art style](./2d-game-art-syle-reference.md).

To create a 2D game, [set up your Unity project](./setup-project-2d-game.md) and then familiarize yourself with the relevant concepts in the following order:

1. [Fundamentals](#Fundamentals)
2. [Scripting](#Scripting)
3. [Sprites](#Sprites)
4. [Building in-game environments](#Environment)
5. [Character animation](#Animation)
6. [Graphics](#Graphics)
7. [Physics 2D](#Physics)
8. [Audio](#Audio)
9. [User interface](#UI)
10. [Profiling, optimizing and testing](#Profiling)
11. [Publishing](#Publishing)

## Fundamentals

[GameObjects](GameObjects.md) are fundamental objects in Unity that represent characters, props, scenery, and more. Every object in your game is a GameObject.

GameObjects represent the items in your game; the space in which you place them to build your level is called a [scene](CreatingScenes.md). Scenes in Unity are always 3D; when you make a 2D game in Unity, you typically choose to ignore the third dimension (the z-axis) but you can also use it in special cases, for example when making 2.5D games.

The behavior of GameObjects is defined by blocks of functionality called components. The following components are fundamental for 2D games:

* [Transform](class-Transform.md): the **Transform component** determines the Position, Rotation, and Scale of each GameObject in the scene. Every GameObject has a Transform component.
* [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md): the Sprite Renderer component renders the Sprite and controls how it looks in a scene.
* [Cameras](class-Camera.md): devices that capture and display the world to the player. Marking a Camera as **Orthographic** removes all perspective from the Camera’s view. This is mostly useful for making isometric or 2D games.
* [Collider 2D](2d-physics/collider/collider-2d-landing.md): this component defines the shape of a 2D GameObject for the purposes of physical **collisions**. See [2D Physics](class-Physics2DManager).

Components are UI representations of C# classes; you can use scripts to change and interact with components, or create new ones. See the [Scripting](#Scripting) section for more details.

## Scripting

All 2D games need [scripts](scripting.md). Scripts respond to input from the player and arrange for events in the gameplay to happen when they should.

For details on how to use scripts in Unity see [Scripting Overview](object-oriented-development.md). Also see the Unity Learn [Beginner Scripting](https://learn.unity.com/course/beginner-scripting) course.

Scripts are attached to GameObjects, and any script you create inherits from the [MonoBehaviour](../ScriptReference/MonoBehaviour.md) class.

## Sprites

[Sprites](sprite/sprite-landing.md) are 2D graphic objects. You use Sprites for all types of 2D games. For example, you can import an image of your main character as a Sprite.

![A character Sprite](../uploads/Main/quickstart-sprite-2d.png)

A character Sprite

You can also use a collection of Sprites to build a character. This allows you greater control over the movement and animation of your characters.

![Multiple Sprites that make up the parts of a character, displayed in the Sprite Editor](../uploads/Main/quickstart-sprite-editor-2d.png)

Multiple Sprites that make up the parts of a character, displayed in the Sprite Editor

### Importing and setting up Sprites

Import your Sprites with Unity’s recommended settings; see [Importing and Setting Up Sprites](sprite/sprite-landing.md).

### Rendering Sprites

Use the [Sprite Renderer](sprite/renderer/sprite-renderer-reference.md) component to render your Sprites. For example, you can use the Sprite Renderer to change the color and opacity of a Sprite.

![Adjusting the color of a Sprite with the Sprite Renderer](../uploads/Main/quickstart-sprite-renderer-2d.png)

Adjusting the color of a Sprite with the Sprite Renderer

See the [Introduction to the Sprite Renderer Learn tutorial](https://learn.unity.com/tutorial/introduction-to-the-sprite-renderer#). Sorting SpritesBy organizing Sprites in layers, you can create an illusion of depth. You can sort Sprites according to many strategies. See [Sorting Sprites](sprite/sprite-landing.md) for full details. For example, you might sort Sprites along the y-axis, so that Sprites that are higher up are sorted behind Sprites that are lower, to make the Sprites that are higher appear further away than the Sprites that are lower.

![Sprites sorted along the y-axis](../uploads/Main/quickstart-sort-sprites-2d.png)

Sprites sorted along the y-axis

To set the overlay order of Sprites, use [Sorting Layers](class-TagManager.md).

To group GameObjects with [Sprite Renderers](sprite/renderer/sprite-renderer-reference.md), and control the order in which they render their Sprites, use [Sorting Groups](sprite/sorting-group/sorting-group-landing.md).

### Sprite Atlas

You can use a [Sprite Atlas](sprite/atlas/atlas-landing.md) to consolidate several Textures into a single combined Texture. This optimizes your game and saves memory. For example, you can add all your Sprites associated with a particular character or purpose to a Sprite Atlas.

![A Sprite Atlas](../uploads/Main/quickstart-sprite-atlas-2d.png)

A Sprite Atlas

See the [Introduction to the Sprite Atlas Learn tutorial](https://learn.unity.com/tutorial/introduction-to-the-sprite-atlas).

## Building in-game environments

Environment design refers to the process of building your game’s levels and environments. You can combine the environment design tools in this section in whichever way makes the most sense for your game; for example, you can make a top-down game using only 9-slice, or you can make a side on platformer with Tilemap and SpriteShape.

### 9-slicing

9-slicing is a 2D technique that allows you to reuse an image at various sizes without needing to prepare multiple assets. Unity can dynamically stretch and tile designated parts of a Sprite to allow one Sprite to serve as the border or background for UI elements of many sizes. See [9-slicing Sprites](sprite/9-slice/9-slice-landing.md).

For example, you could use 9-slicing to stretch a Sprite to shape when you build a 2D level.

![A 9-sliced Sprite, split into nine sections](../uploads/Main/quickstart-9-slice-2d.png)

A 9-sliced Sprite, split into nine sections

See the [Using 9-Slicing for Scalable Sprites Learn tutorial](https://learn.unity.com/tutorial/using-9-slicing-for-scalable-sprites-2019-3).

### Tilemap

The [Tilemap](tilemaps/work-with-tilemaps/tilemap-reference.md) component is a system that stores and handles Tile assets for creating 2D levels. Use the [2D Tilemap Editor](tilemaps/work-with-tilemaps/tilemap-reference.md) package (installed by default) to use Tilemaps.

For example, you can use Tilemaps to paint levels using Tiles and brush tools and define rules for how Tiles behave.

![The Tile Palette window, used to edit Tilemaps](../uploads/Main/quickstart-tilemap-2d.png)

The Tile Palette window, used to edit Tilemaps

See the [Introduction to Tilemaps Learn tutorial](https://learn.unity.com/tutorial/introduction-to-tilemaps).

#### 2D Tilemap Extras

To add some extra Tilemap assets to your Project, install the [2D Tilemap Extras](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@1.6/manual/index.html) package. This package contains reusable 2D and Tilemap **Editor scripts** that you can use for your own Projects. You can customize the behavior of the scripts to create new Brushes that suit different scenarios.

#### Isometric Tilemaps

For games with isometric perspective, you can create [Isometric Tilemaps](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md).

### SpriteShape

In a similar way to a vector drawing tool, SpriteShape provides a more flexible way to create larger Sprites, such as organic-looking landscapes and paths. See the [Sprite Shape Profile](https://docs.unity3d.com/Packages/com.unity.2d.spriteshape@3.0/manual/SSProfile.html).

![A path created in SpriteShape](../uploads/Main/quickstart-spriteshape-path-2d.png)

A path created in SpriteShape

See the [Working with SpriteShape tutorial](https://learn.unity.com/tutorial/working-with-spriteshape).

## Character animation

There are three different ways you can animate 2D characters:

| 2D animation type | Used for |
| --- | --- |
| Frame-by-frame | Artistic reasons, if you want your game to have a classic animation art style. Frame-by-frame animation is relatively resource-intensive, both to make and to run. |
| Cutout | Smooth skeletal animation, when the characters don’t require realistic articulation. |
| Skeletal | Smooth skeletal animation where Sprites bend according to the bone structure. Use this when the characters need a more organic feel. |

### Frame-by-frame

Frame-by-frame animation is based on the traditional cel animation technique of drawing each moment of an animation as individual images, which are played in fast sequence, like flipping pages on a flipbook.

For information, refer to [Sprite Swap](https://docs.unity3d.com/Packages/com.unity.2d.animation@latest?subfolder=/manual/SpriteSwapLanding.html) in the 2D Animation package manual.

![Frame-by-frame animation in the Sprite Editor](../uploads/Main/quickstart-frame-by-frame-2d.png)

Frame-by-frame animation in the Sprite Editor

See the [Introduction to Sprite Animations Learn tutorial](https://learn.unity.com/tutorial/introduction-to-sprite-animations).

### Cutout

In cutout animation, multiple Sprites make up the body of a character, and each piece moves to give the visual effect of the whole character moving. This animation style is similar to skeletal animation (see below), except that the Sprites don’t bend.

![Cutout animation in the Sprite Editor](../uploads/Main/quickstart-cutout-2d.png)

Cutout animation in the Sprite Editor

### Skeletal

With skeletal animation, you map a Sprite or a group of Sprites onto an animation skeleton. You can create and define animation bones for characters and objects, that define how they should bend and move. This approach allows the bones to bend and deform the Sprites, for a more natural movement style. To use skeletal animation, you need to use the [2D Animation](com.unity.2d.animation.md) package (installed by default).

For a 2D Animation workflow, including a guide to working with the Bone Editor, see the [2D Animation documentation](https://docs.unity3d.com/Packages/com.unity.2d.animation@1.0/manual/index.html).

![A character with bones in the Bone Editor](../uploads/Main/quickstart-skeletal-2d.png)

A character with bones in the Bone Editor

## Graphics

This section describes the [2D URP graphic features](2d-urp-landing.md) available when using Universal **Render Pipeline** (URP).

### Lighting

Because you’re using URP with the 2D Renderer, you can use the Light 2D component to apply optimized 2D lighting to Sprites. For details, see [Introduction to Lights 2D](urp/Lights-2D-intro.md).

![These two images show the same scene; in the image on the left, 2D Lights are disabled, and in the image on the right, 2D lights are enabled. With 2D Lights, you can use the same Sprites to create different weather conditions or moods.](../uploads/Main/quickstart-lighting-2d.png)

These two images show the same scene; in the image on the left, 2D Lights are disabled, and in the image on the right, 2D lights are enabled. With 2D Lights, you can use the same Sprites to create different weather conditions or moods.

See the [Lighting in URP Learn tutorial](https://learn.unity.com/tutorial/editing-lighting-in-the-lightweight-render-pipeline).

### Shadows

To define the shape and properties that a Light uses to determine the shadows it casts, use the [Shadow Caster 2D component](urp/2DShadows.md).

![A shadow intensity of 0.5 in the Shadow Caster 2D component](../uploads/Main/quickstart-shadows-2d.png)

A shadow intensity of 0.5 in the Shadow Caster 2D component

### Enhanced look and feel

Particle systems and **post-processing** are optional tools that you can use to add polish to your game.

#### Particle systems

You can use **particle** systems to create dynamic objects like fire, smoke or liquids, as an alternative to using a Sprite. Sprites are more suited to physical objects. See [Particle systems](ParticleSystems.md).

![A fire effect, created with the Particle System and Shader Graph for 2D](../uploads/Main/quickstart-particle-system-2d.png)

A fire effect, created with the Particle System and Shader Graph for 2D

#### Post-processing

You can use post-processing effects and full-screen effects to significantly improve the appearance of your game. For example, you can use these effects to simulate physical camera or film properties, or to create stylized visuals.

URP has its own post-processing implementation. See [Post-processing in the Universal Render Pipeline](urp/integration-with-post-processing.md).

![The Lost Crypt demo uses the bloom and vignette post-processing effects](../uploads/Main/quickstart-post-processing-2d.png)

The Lost Crypt demo uses the bloom and vignette post-processing effects

## Physics 2D

The Physics 2D settings define limits on the accuracy of the physical simulation in your 2D game. See [2D Physics](class-Physics2DManager).

This [video](https://www.youtube.com/watch?v=Xxbs9x2qB7Y&feature=youtu.be) provides an overview of 2D physics features in Unity 2020.1.

To learn how to use Unity’s 2D **physics engine**, see the [2D Physics Learn tutorial](https://learn.unity.com/tutorial/2d-physics).

The following 2D physics tools are useful for 2D games.

* [Rigidbody 2D](#Rigidbody)
* [Collider 2D](#Collider)
* [Triggers](#Triggers)
* [2D Joints](#Joints)
* [2D Effectors](#Effectors)

### Rigidbody 2D

A **Rigidbody** 2D component places a GameObject under the control of the physics engine. See [Rigidbody 2D](2d-physics/rigidbody/introduction-to-rigidbody-2d.md).

![The Rigidbody 2D component](../uploads/Main/quickstart-rigidbody-component-2d.png)

The Rigidbody 2D component

### Collider 2D

Collider 2D components define the shape of a 2D GameObject for the purposes of physical collisions. You can also use **Collider** 2D components for input detection. For example, in mobile games you can use them to make Sprites selectable.

The Collider 2D types that you can use with Rigidbody 2D are:

* [Circle Collider 2D](2d-physics/collider/circle-collider-2d-reference.md)
* [Box Collider 2D](2d-physics/collider/box-collider-2d-reference.md)
* [Polygon Collider 2D](2d-physics/collider/polygon-collider-2d-reference.md)
* [Edge Collider 2D](2d-physics/collider/edge-collider-2d-reference.md)
* [Capsule Collider 2D](2d-physics/collider/capsule-collider/capsule-collider-2d-reference.md)
* [Composite Collider 2D](2d-physics/collider/composite-collider/composite-collider-2d-reference.md)

![The Circle Collider 2D component](../uploads/Main/quickstart-circle-collider-component.png)

The Circle Collider 2D component

### Triggers

When you set a Collider 2D as a Trigger (by enabling its **Is Trigger** property), it no longer behaves as a physical object, and it can intersect with other Colliders without causing a collision. Instead, when a Collider enters its space, Unity calls the `OnTriggerEnter` function on the Trigger GameObject’s scripts.

![The Circle Collider 2D component with Is Trigger selected](../uploads/Main/quickstart-is-trigger-2d.png)

The Circle Collider 2D component with **Is Trigger** selected

### 2D Joints

Joints attach GameObjects together. You can only attach 2D **joints** to GameObjects that have a Rigidbody 2D component attached, or to a fixed position in world space. See [2D Joints](2d-physics/joints/2d-joints-landing.md).

### 2D Effectors

Use Effector 2D [components](UsingComponents.md) with [Collider 2D](2d-physics/collider/collider-2d-landing.md) components to direct the forces of [physics](PhysicsSection.md) in your scene when [GameObject](GameObjects.md) Colliders come into contact with each other. See [2D Effectors](2d-physics/effectors/effectors-2d-landing.md).

## Audio

You can add background music and sound effects to your game in Unity; see [Audio Overview](AudioOverview.md). Use third-party software to create your audio and import it into Unity with the recommended settings.

## User interface

If you want to add a menu or help to your game, you need to set up a [user interface](UIToolkits.md). To set up a user interface, use [Unity UI](https://docs.unity3d.com/Manual/UIToolkits.html).

## Profiling, optimizing and testing a build

### Profiling

Profiling allows you to see how resource-intensive the different parts of your game are. You should always profile your game on its target release platform; see [Profiling your application](profiler-profiling-applications.md).

### Testing

Test your game and your code with the Unity Test Framework; see [Unity Test Framework](https://docs.unity3d.com/Packages/com.unity.test-framework@1.1/manual/index.html).

![The Test Runner window](../uploads/Main/quickstart-testing-2d.png)

The Test Runner window

## Publishing

When you’ve finished your game, you’re ready to publish it. See [Publishing Builds](PublishingBuilds).