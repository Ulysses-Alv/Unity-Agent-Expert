# Destroy LowLevelPhysics2D API objects and manage memory

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Destroy worlds, bodies, and shapes when you no longer need them. If you get physics objects in a `NativeArray` or `ReadOnlySpan`, you must also free up their memory to avoid memory leaks.

## Destroy physics objects

To destroy a world, body, or shape, call the `Destroy` method of the object. For example:

```
// Create a physics body
PhysicsBody myBody = world.CreateBody();

// Destroy the body when you no longer need it
myBody.Destroy();
```

To destroy objects quickly, use the following approaches:

* Destroy a physics body to automatically destroy all its shapes and chains, and any **joints** it’s connected to.
* Destroy a world to automatically destroy all its bodies and shapes.

Unity logs an error in the **Console** window if you try to destroy an object that no longer exists. To check if an object is already destroyed, get the `isValid` property of the object.

Unity might destroy your world automatically when you enter or exit Play mode.

## Destroy a batch of objects

To destroy a batch of objects, use the following methods instead:

* `DestroyBodyBatch` if you created the batch using `CreateBodyBatch`.
* `DestroyShapeBatch` if you created the batch using `CreateShapeBatch`.

### Protect objects from deletion

To protect an object from deletion, call the `SetOwner` method of the `PhysicsBody` object. `SetOwner` returns a unique key integer. If you try to destroy the body, you must pass in the key or the destruction fails. For more information, refer to [`PhysicsBody`](../../ScriptReference/LowLevelPhysics2D.PhysicsBody.md).

**Warning**: The key is only a deterrent and doesn’t meet cryptographic standards.

## Avoid memory leaks

Most of the `LowLevelPhysics2D` API avoids creating memory that needs disposing or garbage collection, and most methods return simple `struct` types.

However if you use an API that stores and returns physics objects or query results as a [`NativeArray`](../../ScriptReference/Unity.Collections.NativeArray_1.md) or `ReadOnlySpan` type, the recommended best practice is to dispose of the native memory when you no longer need it. Disposing avoids memory leaks.

For example if you create an array of body definitions, use the `Dispose` method to free the memory when you no longer need it. For example:

```
void Start(){
    // Create a native array to hold body definitions
    // The allocator property determines the lifetime of the array
    NativeArray<PhysicsBodyDefinition> bodyDefinitions = new NativeArray<PhysicsBodyDefinition>(bodyCount, Allocator.Temp);
}

void OnDestroy()
{
    // Dispose of the native array to free up memory
    bodyDefinitions.Dispose();
}
```

You can also create native arrays with `using` statements to automatically dispose of them when they go out of scope. For example:

```
using NativeArray<PhysicsQuery.WorldCastResult> results = world.CastGeometry(objectGeometry, translation, filter, PhysicsQuery.WorldCastMode.Closest, Allocator.Temp);
```

To make `NativeArray` or `ReadOnlySpan` results persist beyond their scope, for example if you need data to persist for multiple frames, pass `Allocator.Persistent` into the method. This approach can decrease performance, and you must still dispose of the native array or read-only span when you no longer need it.

For more information about native memory and allocators, refer to [Native memory](../performance-native-memory.md).

## Additional resources

* [Optimize the LowLevelPhysics2D API with multithreading](2d-physics-api-multithreading.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub