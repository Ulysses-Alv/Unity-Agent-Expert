# Check LowLevelPhysics2D API intersections

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To check whether 2D physics objects overlap or will collide, also known as intersections, use `LowLevelPhysics2D` API query methods.

Use one of the following methods:

* The [`PhysicsQuery`](../../ScriptReference/LowLevelPhysics2D.PhysicsQuery.md) API to check for overlaps. For example, `PhysicsQuery.CapsuleAndCircle` to check whether a capsule overlaps with a circle.
* The `Intersect` method of a geometry object to check if it overlaps with another geometry object. For example, [`CapsuleGeometry.Intersect`](../../ScriptReference/LowLevelPhysics2D.CapsuleGeometry.Intersect.md).

The world object also has the following methods:

* Methods that cast rays or shapes. For example, [`PhysicsWorld.CastRay`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.CastRay.md) or [`PhysicsWorld.CastGeometry`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.CastGeometry.md).
* Overlap test methods that return true if a point or shape overlaps another object. For example, [`PhysicsWorld.TestOverlapAABB`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.TestOverlapAABB.md).
* Overlap methods that return the objects that overlap a point or shape. For example, [`PhysicsWorld.OverlapCircle`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.OverlapCircle.md).

All methods are thread-safe. For more information, refer to [Optimize the LowLevelPhysics2D API with multithreading](2d-physics-api-multithreading.md).

**Important**: These methods return `NativeArray` types, which you must dispose of to free the memory they use. For more information, refer to [Destroy objects and manage memory](2d-physics-api-destroy.md).

## Example

To cast rays to check whether a shape will collide with other shapes, use the `CastGeometry` API.

For example, the following script checks if a small falling circle will collide with objects underneath it. Attach the script to a **GameObject** and enter Play mode, then use the **Big Circle** checkbox to toggle an object underneath the falling circle.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;
using Unity.Collections;

public class CastGeometryExample : MonoBehaviour
{
    private PhysicsWorld world;
    public bool BigCircle;

    void Start()
    {
        world = PhysicsWorld.defaultWorld;

        // Create a small falling circle
        CircleGeometry object1Geometry = CircleGeometry.Create(0.5f, new Vector2(0.5f, 8f));
        PhysicsBody object1 = world.CreateBody(new PhysicsBodyDefinition
        { 
            position = object1Geometry.center,
            type = PhysicsBody.BodyType.Dynamic
        });
        object1.CreateShape(object1Geometry);

        if (BigCircle)
        {
            // Create a larger static circle below
            CircleGeometry object2Geometry = CircleGeometry.Create(3f, new Vector2(0f, 0f));
            PhysicsBody object2 = world.CreateBody(new PhysicsBodyDefinition
            { 
                position = object2Geometry.center,
                type = PhysicsBody.BodyType.Static
            });
            object2.CreateShape(object2Geometry);
        }

        // Set translation and filter for the cast
        Vector2 translation = new Vector2(0, -10f); // Move downwards
        PhysicsQuery.QueryFilter filter = new PhysicsQuery.QueryFilter();

        // Cast the circle geometry through the world
        NativeArray<PhysicsQuery.WorldCastResult> results = world.CastGeometry(object1Geometry, translation, filter, PhysicsQuery.WorldCastMode.Closest, Allocator.Temp);

        if (results.Length > 0)
        {
            Debug.Log("Collision will occur!");
        }

        // Dispose of the NativeArray to free memory
        results.Dispose();
    }
}
```

## Additional resources

* [Introduction to collision in the LowLevelPhysics2D API](2d-physics-api-interactions-introduction.md)
* [`PhysicsAABB`](../../ScriptReference/LowLevelPhysics2D.PhysicsAABB.md)