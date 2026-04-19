# Connect LowLevelPhysics2D API objects with joints

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To create a connection or constraint between two `LowLevelPhysics2D` API objects, create a **joint** that connects the two physics bodies.

For example, create a **fixed joint** that welds two bodies together, or a distance joint that keeps them a set distance apart. Use joints to simulate real-world mechanical behaviours like a spring or a hinge.

All joints allow you to set the anchor points on the bodies they connect, and limits on the force and torque they apply. For more information, refer to the [`PhysicsJoint`](../../ScriptReference/LowLevelPhysics2D.PhysicsJoint.md) API.

Follow these steps:

1. Create two `PhysicsBody` objects. For more information, refer to [Create a physics object](2d-physics-api-physics-object.md).
2. Create a definition object for the type of joint you want. For example, a `PhysicsDistanceJointDefinition` object. Make it a `public` field so it displays its properties in the ****Inspector**** window.

   For example:

   ```
   public PhysicsDistanceJointDefinition jointDefinition = new PhysicsDistanceJointDefinition
   {
       // Specify the two bodies to connect
       bodyA = body1,
       bodyB = body2,

       // Set the distance
       distance = 2f
   };
   ```

   A new definition has a set of default values. For more information about definitions and changing the default values, refer to [Configure 2D physics properties using a definition](2d-physics-api-definitions.md).

   For the full list of joint types, refer to [`PhysicsJoint.JointType`](../../ScriptReference/LowLevelPhysics2D.PhysicsJoint.JointType.md).
3. To create the joint, use the `CreateJoint` method of your world object, and pass in the definition. For example:

   ```
   PhysicsJoint distanceJoint = world.CreateJoint(jointDefinition);
   ```
4. If your script is in a `MonoBehaviour` class attached to a **GameObject**, adjust the properties in the **Inspector** window.

   To configure the properties in your script instead, set the properties of the definition object before you create the joint. For more information, refer to [Configure 2D physics properties using a definition](2d-physics-api-definitions.md).

If you destroy a body, any joints connected to that body are also destroyed. For more information, refer to [Destroy physics objects and manage memory](2d-physics-api-destroy.md).

## Example

The following example uses a fixed joint to create a fixed point that a large circle swings from.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D; 

public class CreateJoint : MonoBehaviour
{
    public PhysicsDistanceJointDefinition jointDefinition = new PhysicsDistanceJointDefinition();

    void Awake()
    {
        PhysicsWorld world = PhysicsWorld.defaultWorld;

        // Create a static fixed body
        PhysicsBodyDefinition bodyDefinition1 = new PhysicsBodyDefinition{
            type = PhysicsBody.BodyType.Static,
            position = new Vector2(-5f, 0f),
        };
        PhysicsBody object1 = world.CreateBody(bodyDefinition1);

        // Create a large circle below the fixed body
        PhysicsBodyDefinition bodyDefinition2 = new PhysicsBodyDefinition{
            type = PhysicsBody.BodyType.Dynamic,
            position = new Vector2(5f, 0f),
        };
        PhysicsBody object2 = world.CreateBody(bodyDefinition2);
        object2.CreateShape(new CircleGeometry { radius = 3f });

        // Add the bodies to the joint definition
        jointDefinition.bodyA = object1;
        jointDefinition.bodyB = object2;
        
        // Add the joint to the world
        world.CreateJoint(jointDefinition);        
    }
}
```

## Additional resources

* [Collisions and interactions in the LowLevelPhysics2D API](2d-physics-api-interactions-landing.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub