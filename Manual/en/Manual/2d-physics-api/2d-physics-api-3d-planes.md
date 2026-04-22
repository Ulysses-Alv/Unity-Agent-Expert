# Use 2D physics in 3D space using the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Create a 2D physics world at any angle in 3D space, instead of the traditional x, y plane that faces the **camera**. This allows you to use 2D physics in a 3D game.

Follow these steps:

1. [Create a physics world](2d-physics-api-world.md).
2. Set the `transformPlane` property of the world to either `XZ` or `ZY`.

The physics system continues to use 2D `Vector2` coordinates and vectors, but Unity converts the final positions to 3D space according to the `transformPlane`:

* `XY` means the x-axis remains horizontal and the y-axis remains vertical, so that the world faces the camera in 3D space. This is the default.
* `XZ` means the x-axis remains horizontal but the y-axis becomes depth, so that the world lies flat in 3D space.
* `ZY` means the x-axis becomes depth but the y-axis remains vertical, so that the world stands upright in 3D space.

To convert coordinates manually between 2D `Vector2` and 3D `Vector3` types, use the helper methods in the [`PhysicsMath`](../../ScriptReference/LowLevelPhysics2D.PhysicsMath.md) API.

The following example uses the `XZ` plane, which creates a 2D physics world that lies flat on the ground in 3D space.

To check the effect, you might need to set the **Scene** view to **3D** mode. For more information, refer to [Gizmos menu](../GizmosMenu.md).

```
using Unity.Collections;
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class CreateXZPlane : MonoBehaviour
{
    public PhysicsWorld world;

    private void Start()
    {
        // Create a world with a transform plane of XZ, which is the floor in 3D space.
        PhysicsWorldDefinition worldProperties = new PhysicsWorldDefinition
        {
            transformPlane = PhysicsWorld.TransformPlane.XZ,
            gravity = Vector2.zero
        };
        world = PhysicsWorld.Create(worldProperties);

        // Create a square boundary
        PhysicsBody boundary = world.CreateBody();
        using NativeList<Vector2> extentPoints = new NativeList<Vector2>(Allocator.Temp)
        {
            new(-4f, 4f),
            new(4f, 4f),
            new(4f, -4f),
            new(-4f, -4f)
        };
        ChainGeometry boundaryWalls = new ChainGeometry(extentPoints.AsArray());
        boundary.CreateChain(boundaryWalls, PhysicsChainDefinition.defaultDefinition);

        // Create a body and set it moving
        PhysicsBodyDefinition bodyDefinition = new PhysicsBodyDefinition();
        bodyDefinition.type = PhysicsBody.BodyType.Dynamic;
        bodyDefinition.linearVelocity = new Vector2(7.3f, 5.7f);
        bodyDefinition.angularVelocity = 0f;
        PhysicsBody body = world.CreateBody(bodyDefinition);

        // Add a shape with a bouncy material
        body.transformObject = transform;
        PhysicsShapeDefinition shapeDefinition = new PhysicsShapeDefinition();
        shapeDefinition.surfaceMaterial = new PhysicsShape.SurfaceMaterial{bounciness = 1f, friction = 0f};
        body.CreateShape(new CircleGeometry { radius = 1f }, shapeDefinition);
    }

    private void OnDisable()
    {
        world.Destroy();
    }
}
```

## Additional resources

* [3D GameObjects in 2D URP scenes](../urp/2d-renderer-urp-features-landing.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub.
* [World definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-world.md)