# Combine LowLevelPhysics2D API shapes

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To combine physics shapes into a compound shape, create a composer using a `PhysicsComposer` object. For example, create a car shape by combining a car body and two wheels.

Follow these steps:

1. At the top of your script, import the `Unity.Collections` library that contains the allocator APIs the composer uses:

   ```
   using Unity.Collections;
   ```
2. Create the `Geometry` shapes you want to combine. For more information, refer to [Create a physics object](2d-physics-api-physics-object.md).
3. Create a new `PhysicsComposer` object. Pass in an allocator, which creates temporary memory for building up the combined object. For example:

   ```
   PhysicsComposer physicsComposer = PhysicsComposer.Create(Allocator.Temp);
   ```

   For more information about allocators, refer to [Allocator](../../ScriptReference/Unity.Collections.Allocator.md).
4. Add each object to the composer as a layer, specifying its geometry, its position, and how to combine the shape with the existing layers. For example:

   ```
   physicsComposer.AddLayer(circleGeometry, PhysicsTransform.identity, PhysicsComposer.Operation.OR);
   ```

   For more information about `Operation`, refer to the [Layer operations](#layer-operations) section.
5. Create the combined shape using `CreatePolygonGeometry`. Pass in a scale, and an allocator that creates the memory to store the combined shape. For example:

   ```
   using NativeArray<PolygonGeometry> combinedShape = composer.CreatePolygonGeometry(new Vector2(1f, 1f), Allocator.Temp);
   ```

   **Note**: `using` automatically disposes of the memory allocated for the `NativeArray` when it goes out of scope. For more information, refer to [Destroy LowLevelPhysics2D API objects and manage memory](2d-physics-api-destroy.md).
6. Create a body with the combined shape using the `CreateShapeBatch` method. For example:

   ```
   body.CreateShapeBatch(combinedShape, PhysicsShapeDefinition.defaultDefinition);
   ```

The `PhysicsComposer` API combines `CircleGeometry`, `CapsuleGeometry`, `PolygonGeometry`, `PhysicsShape`, or any closed region defined by a set of `Vector2` points.

## Layer operations

The following `PhysicsComposer.Operation` operations are available:

* `OR`: Adds the shape.
* `AND`: Keeps only the overlapping areas between the new shape and the existing shape.
* `NOT`: Subtracts the new shape from the existing shapes.
* `XOR`: Removes overlapping areas, but keeps non-overlapping areas.

## Example

The following example combines a circle and a capsule shape. Attach this script to a **GameObject** in a **scene**, then enter Play mode to check the combined shape.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;
using Unity.Collections;

public class CombineShapes : MonoBehaviour
{
    // Create the definition for a circle object
    public CircleGeometry MyCircleGeometry = CircleGeometry.defaultGeometry;
    public PhysicsTransform CircleTransform = new Vector2(0f, 0f);

    // Create the definition for a capsule object
    public CapsuleGeometry MyCapsuleGeometry = CapsuleGeometry.defaultGeometry;
    public PhysicsTransform CapsuleTransform = new Vector2(0.75f, 0f);

    private void Awake()
    {
        PhysicsWorld m_PhysicsWorld = PhysicsWorld.defaultWorld;
        
        // Create a composer to combine shapes
        PhysicsComposer composer = PhysicsComposer.Create(Allocator.Temp);
        
        // Add both shapes to the composer
        composer.AddLayer(MyCircleGeometry, CircleTransform);
        composer.AddLayer(MyCapsuleGeometry, CapsuleTransform, PhysicsComposer.Operation.OR);

        // Combine the shapes
        using NativeArray<PolygonGeometry> combinedShape = composer.CreatePolygonGeometry(new Vector2(1f, 1f), Allocator.Temp);
        
        // Create a body with the combined shapes
        PhysicsBody body = m_PhysicsWorld.CreateBody();       
        body.CreateShapeBatch(combinedShape, PhysicsShapeDefinition.defaultDefinition);

        // Destroy the composer to also destroy the temporary allocator.
        composer.Destroy();
    }   
}
```

## Additional resources

* [Connect LowLevelPhysics2D API objects with joints](2d-physics-api-joints.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub.