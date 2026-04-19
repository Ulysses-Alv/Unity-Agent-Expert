# Attach custom data to LowLevelPhysics2D API objects

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To attach custom data to an object you create using the `LowLevelPhysics2D` API, use the `PhysicsUserData` type.

You can attach data to the following types:

* `World`
* `Body`
* `Shape`
* `Chain`
* `Joint`

Follow these steps:

1. Create a `PhysicsUserData` instance and populate it with your custom fields. For example:

   ```
   PhysicsUserData physicsUserData = new PhysicsUserData
       {
           customObject = this,
           customBool = true,
           customFloat = 123.4f,
           customInt = 567,
           customPhysicsMask = PhysicsMask.All
       };
   ```
2. Assign a `PhysicsUserData` instance to the `userData` property of the physics object. For example, to assign the custom data to a shape:

   ```
   myShape.userData = physicsUserData;
   ```
3. To fetch the data, for example when you [detect a collision](2d-physics-api-collision-handle.md), get the `userData` property of the physics object. For example:

   ```
   Debug.Log(myShape.userData.intValue);
   ```

## Additional resources

* [Collisions and interactions in the LowLevelPhysics2D API](2d-physics-api-interactions-landing.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub