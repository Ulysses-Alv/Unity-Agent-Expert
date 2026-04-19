# Move a GameObject with the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

By default, when you [create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md), the object isn’t connected to GameObjects in the **scene**.

To move a GameObject, configure the physics object to update the **Transform component** of the GameObject.

Follow these steps:

1. To enable physics bodies updating Transform components globally, set the `TransformWriteMode` property of the world to [`Fast2D`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.TransformWriteMode.Fast2D.md) or [`Slow2D`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.TransformWriteMode.Slow2D.md).
2. To enable the physics body updating Transform components, set its `TransformWriteMode` to `Current`.
3. To attach the Transform component of the GameObject to the physics body, set the `transformObject` property to the Transform component.

   **Note**: The `transformObject` property is available only in a C# script, not in the ****Inspector**** window of a public `PhysicsBodyDefinition` object.
4. Make sure the `type` property of the physics body is set to `PhysicsBody.BodyType.Dynamic`.

For example, attach the following script to a GameObject, then enter Play mode. The physics body falls under gravity, and updates the position in the Transform component in the **Inspector** window.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class MoveGameObject : MonoBehaviour
{
    public PhysicsWorld world;
    public PhysicsWorldDefinition worldDefinition = PhysicsWorldDefinition.defaultDefinition;

    void Awake()
    {
        // Enable physics bodies updating transforms
        worldDefinition.transformWriteMode = PhysicsWorld.TransformWriteMode.Fast2D;
        world = PhysicsWorld.Create(worldDefinition);
    }    

    void Start()
    {
        // Create a body and shape 
        PhysicsBody circleBody = world.CreateBody();
        CircleGeometry circleGeometry = new CircleGeometry { radius = 2f };
        circleBody.CreateShape(circleGeometry);
        
        // Set body to dynamic so it falls under gravity
        circleBody.type = PhysicsBody.BodyType.Dynamic; 

        // Enable this physics body updating transforms
        circleBody.transformWriteMode = PhysicsBody.TransformWriteMode.Current;

        // Set the updated transform as the transform of this GameObject
        circleBody.transformObject = transform;
    }
}
```

## Rotate objects more quickly

By default, Unity sets a limit on how fast a physics body rotates, to avoid forces becoming too large and objects passing through each other incorrectly. To remove the limit, create a [Physics Low Level Settings 2D asset](2d-physics-api-global-settings.md) and set its **Fast Rotation Allowed** property to `true`. This property is recommended only for circular objects.

For more information, refer to [Configure global 2D physics settings](2d-physics-api-global-settings.md).

## Interpolate positions

By default, the physics body updates the Transform component after Unity finishes calculating the physics simulation. To interpolate positions between simulation steps instead, set the `TransformWriteMode` of the physics body to `Interpolate` or `Extrapolate`.

## Additional resources

* [GameObjects](../working-with-gameobjects.md)
* [Add a sprite to a LowLevelPhysics2D API object](2d-physics-api-add-sprite.md)
* [World definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-world.md)
* [Body definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-body.md)