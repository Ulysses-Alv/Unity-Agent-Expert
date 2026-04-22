# Create a world with the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To create physics objects using the `LowLevelPhysics2D` API, you first need to create a physics world.

## Prerequisites

Before you create a physics world, follow these steps:

1. Create a [MonoBehaviour script file](../class-MonoBehaviour.md): from the main menu, select **Assets** > **Create** > **C# Script**.
2. To import the `LowLevelPhysics2D` API namespace, add `using UnityEngine.LowLevelPhysics2D;` at the top of the script.

**Note:** You can use the API in any C# script, not only MonoBehaviour script files. But if you use a MonoBehaviour script file, you can configure physics properties in the ****Inspector**** window of a **GameObject**, move GameObjects, and attach **sprites**.

## Fetch the default world

Unity automatically creates a default world. To fetch it, follow these steps:

1. Get the `defaultWorld` property from the `PhysicsWorld` class.

   ```
   PhysicsWorld world = PhysicsWorld.defaultWorld;
   ```
2. Attach the script to a GameObject in your **scene**.
3. Enter Play mode to run the script.

To adjust the properties of the world, refer to [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md).

Unity creates or recreates the default world at the following times:

* When the Editor starts.
* When you enter or exit Play mode.
* When your built application starts.

## Create your own world

To create your own world, follow these steps:

1. Create a public `PhysicsWorldDefinition` object that holds the world properties and displays them in the **Inspector** window. For example:

   ```
   public PhysicsWorldDefinition worldDefinition = new PhysicsWorldDefinition();
   ```

   A new definition has a set of default values. For example, gravity is set to –9.81f. For more information about definitions and changing the default values, refer to [Configure objects using definitions](2d-physics-api-definitions.md).

   You can also use `PhysicsWorld.defaultDefinition` to get a definition object with the default values.
2. Create a `PhysicsWorld` object with the definition. For example:

   ```
   PhysicsWorld world = PhysicsWorld.Create(worldDefinition);
   ```
3. Attach the script to a GameObject in your scene.
4. To adjust the properties of the world, modify the values in the **Inspector** window.

   To configure the world in your script instead, set the properties of the definition object before you create the world. For more information, refer to [Configure objects using definitions](2d-physics-api-definitions.md).
5. Enter Play mode to run the script and create the world.

## Pause a world

A world starts running as soon as you create it. To pause the simulation, set the [`paused`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld-paused.md) property of the world object to `true`.

## Example

The following example fetches the default world and logs the gravity value to the **Console** window.

```
using UnityEngine;
using UnityEngine.LowLevelPhysics2D;

public class GetDefaultWorld : MonoBehaviour
{
    void Awake()
    {
        // Fetch the default physics world
        PhysicsWorld world = PhysicsWorld.defaultWorld;

        // Log the gravity value to check the world is created
        Debug.Log("The gravity in this world is " + world.gravity);
    }    
}
```

## Additional resources

* [Create a physics object](2d-physics-api-physics-object.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub
* [World definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-world.md)