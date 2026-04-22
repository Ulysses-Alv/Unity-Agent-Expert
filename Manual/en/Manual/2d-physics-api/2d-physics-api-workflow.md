# LowLevelPhysics2D API workflow

Create a 2D **scene** that simulates physics using the [`LowLevelPhysics2D` API](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.md).

Follow these steps:

1. [Include the LowLevelPhysics2D library](#script) in a C# script.
2. [Create a 2D physics world](#world).
3. [Create physics objects](#physics-objects) and add them to the world.
4. [Configure objects with definitions](#configure-objects).
5. [Attach the script to a GameObject](#attach-script), then enter Play mode.

## Include the LowLevelPhysics2D library

In a [C# script](../scripting-get-started.md), add `using UnityEngine.LowLevelPhysics2D;` at the top to include the `LowLevelPhysics2D` library.

You can use the `LowLevelPhysics2D` API in any C# script, not only [MonoBehaviour script files](../class-MonoBehaviour.md). But if you use a MonoBehaviour script file, you can configure physics properties in the ****Inspector**** window of a **GameObject**, move GameObjects, and attach **sprites**.

## Create a 2D physics world

To add 2D physics objects, you first need to create a `PhysicsWorld` object to add them to. Do either of the following:

* Use the world Unity automatically creates.
* Create a new world.

For more information, refer to [Create a physics world](2d-physics-api-world.md).

## Create physics objects

To add physics objects to the world, do the following:

1. Create a `PhysicsBody` object, which defines the position, rotation, and velocity of an object. It doesn’t define an area.
2. Attach one or more `PhysicsShape` objects to the `PhysicsBody`, which define the area that interacts with other shapes. Unity also draws the shapes as a debug visualization.
3. Add the body to the world you created.

For more information, refer to [Create a physics object](2d-physics-api-physics-object.md).

## Configure objects with definitions

To configure the properties of the world and its physics objects, create definition objects, for example `PhysicsBodyDefinition`. Definition objects contain properties like `position` and `gravity` that you can adjust in the Unity Editor, or pass into the object when you create it.

To configure the properties in the **Inspector** window, make the definition object a `public` field.

For more information, refer to [Configure objects using definitions](2d-physics-api-definitions.md).

## Attach the script to a GameObject

To start the simulation, you usually attach the script to a GameObject in your scene, then enter Play mode.

By default, physics objects don’t move a GameObject. To make a physics object update the **Transform component** of the GameObject, refer to [Move a GameObject with the LowLevelPhysics2D API](2d-physics-api-move-gameobject.md).

To add a sprite, refer to [Add a sprite to an object](2d-physics-api-add-sprite.md).

## Example

The following example creates a small circle that falls under gravity onto a large circle. Attach the script to a GameObject in your scene, then enter Play mode.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class Example2DPhysics : MonoBehaviour
{
    void Start()
    {
        // Create the world
        PhysicsWorld world = PhysicsWorld.defaultWorld;

        // Create a body in the world, with a definition object that sets the position and enables physics
        PhysicsBody object1 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0.5f, 8f),
            type = PhysicsBody.BodyType.Dynamic
        });

        // Attach a circle shape to the first body 
        PhysicsShape object1shape = object1.CreateShape(new CircleGeometry());

        // Create a second body in the world, with a definition object that sets the position and disables physics
        PhysicsBody object2 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = new Vector2(0f, 0f),
            type = PhysicsBody.BodyType.Static
        });

        // Attach a circle shape to the second body 
        object2.CreateShape(new CircleGeometry { 
            radius = 3f,
        });
    }
}
```

## Additional resources

* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub