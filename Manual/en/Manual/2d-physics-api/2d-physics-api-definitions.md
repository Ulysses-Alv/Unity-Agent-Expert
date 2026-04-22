# Configure LowLevelPhysics2D API objects with definitions

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To configure the properties of a physics world and its objects, such as gravity, friction, position, and shape, use definitions. Definitions are objects that store physics properties and values. You create a definition, then pass it into the world, body, or shape you create.

Definitions optimize physics performance by avoiding you setting properties after you create a world, body, or shape, which can make the CPU do extra work. Definitions also allow you to pass around and reuse sets of values.

For the full list of definition objects, refer to the [`LowLevelPhysics2D` API](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.md) documentation.

## Create and use a definition

Follow these steps:

1. Create a public definition object in your script, for example a `PhysicsBodyDefinition` object to configure a `PhysicsBody`.

   ```
   // Create a new body definition with default properties and values
   public PhysicsBodyDefinition bodyDefinition = new PhysicsBodyDefinition();
   ```

   A new definition has a set of default values. To change the default values you receive, refer to the [Change the default definition values](#change-default-definitions) section.
2. Pass in the definition when you create the physics object. For example:

   ```
   PhysicsBody myObject = world.CreateBody(bodyDefinition);
   ```

If your script is in a `MonoBehaviour` class attached to a **GameObject**, Unity displays the properties of the definition in the ****Inspector**** window so you can configure them. For the full list of default properties, refer to [Definitions reference for the LowLevelPhysics2D API](2d-physics-api-reference.md).

If you change the definition after you create an object with it, the changes don’t affect the existing object.

**Note:** Definitions are large `struct` objects. To pass definitions into other methods, the recommended best practice is to use the `ref` or `in` keyword to pass them by reference. Passing by reference avoids using extra memory to copy the structure.

## Set a definition property in a script

To configure the definition in your script, set the properties of the definition before you create the object with it.

For the full list of properties, refer to the [`LowLevelPhysics2D` API](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.md) documentation.

For example:

```
public PhysicsBodyDefinition bodyDefinition = new PhysicsBodyDefinition
{
    // Set the linear velocity
    linearVelocity = Vector2.right * 4f,

    // Set the gravity scale
    gravityScale = 0f
};

// Create the body
PhysicsBody myObject = world.CreateBody(bodyDefinition);
```

**Note**: Changing the `position` of a physics body doesn’t change its **center of mass**. Unity calculates the center of mass based on the mass of the shapes attached to the body.

## Create multiple objects with the same definition

To create multiple physics bodies or physics shapes with the same configuration, call the `CreateBodyBatch` method of your world object. Pass in the definition and the number of objects you want to create. For example:

```
// Create 500 bodies using the single definition
NativeArray<PhysicsBody> fiveHundredBodies = world.CreateBodyBatch(bodyDefinition, 500);
```

**Important**: `CreateBodyBatch` returns a `NativeArray` type that points to native memory, so you must dispose the memory when you finish using it. For more information, refer to [Destroy LowLevelPhysics2D API objects and manage memory](2d-physics-api-destroy.md).

## Change the default definitions

To change the default values you receive when you use `new` or `.defaultDefinition` to create a definition object, follow these steps:

1. Create a Physics Low Level Settings 2D asset. For more information, refer to [Configure global 2D physics settings](2d-physics-api-global-settings.md).
2. Select the asset in the **Project** window.
3. In the **Inspector** window, configure the properties in the **Default Definitions** section.

For a full list of properties, refer to [Physics Low Level Settings window reference](2d-physics-api-reference-low-level-settings).

### Example

The following example uses a body definition to set the position of a physics body, then a shape definition to set the density of a shape.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class CreateObjectsWithDefinitions : MonoBehaviour
{
    
    PhysicsWorld world;
    
    // Declare definitions with default properties for the body and shape
    public PhysicsBodyDefinition bodyDefinition = new PhysicsBodyDefinition();
    public PhysicsShapeDefinition  shapeDefinition = new PhysicsShapeDefinition();
 
    void Start()
    {
        // Get the default world
        world = PhysicsWorld.defaultWorld;

        // Set the position of the body using the body definition
        bodyDefinition.position = new Vector2(0f, 5f);

        // Create the physics body with the body definition
        PhysicsBody myObject = world.CreateBody(bodyDefinition);

        // Set the density of the shape to 5 kg/m²
        shapeDefinition.density = 5f;

        // Create the circle geometry
        CircleGeometry circleGeometry = new CircleGeometry { radius = 1.5f };

        // Create the shape with both the geometry and the shape definition
        myObject.CreateShape(circleGeometry, shapeDefinition);
    }
}
```

## Additional resources

* [Create a world with the LowLevelPhysics2D API](2d-physics-api-world.md)
* [Create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)
* [Physics Low Level Settings 2D asset Inspector window reference](../class-PhysicsLowLevelSettings2D.md)