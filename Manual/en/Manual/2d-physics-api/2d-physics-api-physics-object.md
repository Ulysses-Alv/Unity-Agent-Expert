# Create an object with the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

After you [create a physics world](2d-physics-api-world.md), you can add physics objects to the world.

An object is usually made up of two parts:

* A physics body, which defines the position, rotation, and velocity of the object. It doesn’t define an area.
* One or more physics shapes attached to the body, which define the area that interacts with other shapes. Unity also automatically draws the shapes as a debug visualization.

You can add any number of shapes to a physics body. You can create shapes with the following geometries:

* [Circle](../../ScriptReference/LowLevelPhysics2D.CircleGeometry.md)
* [Capsule](../../ScriptReference/LowLevelPhysics2D.CapsuleGeometry.md)
* [Polygon](../../ScriptReference/LowLevelPhysics2D.PolygonGeometry.md)
* [Segment](../../ScriptReference/LowLevelPhysics2D.SegmentGeometry.md), which creates a line
* [Chain segment](../../ScriptReference/LowLevelPhysics2D.ChainGeometry.md), which you use to create a series of connected line segments

## Create a physics body and physics shape

Follow these steps:

1. Create objects that hold the properties of the body and shapes. Make them `public` fields so they display their properties in the ****Inspector**** window. For example:

   ```
   public PhysicsBodyDefinition bodyDefinition = new PhysicsBodyDefinition();
   public PhysicsShapeDefinition shapeDefinition = new PhysicsShapeDefinition();
   ```

   A new definition has a set of default values. For more information about definitions and changing the default values, refer to [Configure 2D physics properties using a definition](2d-physics-api-definitions.md).

   You can also use `PhysicsBody.defaultDefinition` to get a definition object with the default values.
2. Use the `CreateBody` method of the world object to add a body to the world, and pass in the body definition.

   ```
   PhysicsBody myObject = world.CreateBody(bodyDefinition);
   ```
3. Create a shape using a geometry object, for example a `CircleGeometry` object. For example, the following creates a circle with a radius of 2 meters.

   ```
   CircleGeometry circleShape = new CircleGeometry { radius = 2f };
   ```
4. Attach the shape to the body using the `CreateShape` method of the body, and pass in the shape definition.

   ```
   myObject.CreateShape(circleShape, shapeDefinition);
   ```
5. If your script is in a `MonoBehaviour` class attached to a **GameObject**, adjust the properties of the body, the geometry, and the shape in the **Inspector** window.

   To configure the objects in your script instead, set the properties of the definition object before you create the objects. For more information, refer to [Configure 2D physics properties using a definition](2d-physics-api-definitions.md).
6. Enter Play mode to run the script and create the objects. Unity displays the shape in the **Scene** view and Game view as a debug visualization. The line on the shape points towards the rotation direction of the shape.

**Note**: By default, a physics body is static and isn’t affected by physics. To make it dynamic, in the **Inspector** window of the GameObject, set ****Body Type**** to **Dynamic**. For more information, refer to [Configure objects with definitions](2d-physics-api-definitions.md).

After you create a shape from geometry, the shape doesn’t change if you change the geometry object.

## Example

The following example creates a circle. Attach the script to a GameObject in your scene then enter Play mode to check the shape.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class CreateWorldAndObjects : MonoBehaviour
{
    // Declare definitions that contain default properties for the body and shape
    public PhysicsBodyDefinition bodyDefinition = PhysicsBodyDefinition.defaultDefinition;
    public PhysicsShapeDefinition shapeDefinition = PhysicsShapeDefinition.defaultDefinition;
 
    void Start()
    {
        // Get the default world
        PhysicsWorld world = PhysicsWorld.defaultWorld;

        // Create the physics body with the body definition
        PhysicsBody myObject = world.CreateBody(bodyDefinition);

        // Create the circle geometry
        CircleGeometry circleGeometry = new CircleGeometry { radius = 1.5f };

        // Create the shape with both the geometry and the shape definition
        myObject.CreateShape(circleGeometry, shapeDefinition);
    }
}
```

## Additional resources

* [Configure LowLevelPhysics2D API objects with definitions](2d-physics-api-definitions.md)
* [`PhysicsBody.PhysicsTransform`](../../ScriptReference/LowLevelPhysics2D.PhysicsBody.PhysicsTransform.md)