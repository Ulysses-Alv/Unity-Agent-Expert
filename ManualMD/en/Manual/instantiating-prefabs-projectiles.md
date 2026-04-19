# Instantiate projectiles and explosions

You can instantiate **prefabs** to use as projectiles and destroy them with explosion effects in your application.

The following example instantiates a projectile prefab when the user presses the fire button. You can attach it to a **GameObject** which acts as a launcher for the prefab.

## Create a projectile prefab asset

1. In the Hierarchy, right-click and select **3D Object** > **Sphere**.
2. Select the sphere, and in its Inspector, select **Add Component** > **Rigidbody**. The sphere needs a Rigidbody so that it can fly through the air and detect when a collision happens.
3. Rename the sphere to `Projectile` and then drag it into the `Assets` folder of your project to [create a prefab asset](CreatingPrefabs.md#create-a-prefab-asset).
4. You can then delete the sphere from the Hierarchy.

You can optionally add a [texture](Textures.md) to the prefab, change its dimensions, or import a different [model](models.md) to act as the projectile.

## Add an explosion script to the projectile prefab asset

To add an explosion to the projectile prefab, you must have a prefab asset that represents an explosion. You can use the [particle system](ParticleSystems.md) to [create a prefab asset](CreatingPrefabs.md#create-a-prefab-asset), or find an explosion effect on the [Asset Store](https://assetstore.unity.com/vfx/particles/fire-explosions) and add it to the `Assets` folder of your project.

Then create a script called `Projectile` as follows:

```
using UnityEngine;

public class Projectile : MonoBehaviour
{

    public GameObject explosion;
    void Start()
    {
         // Deletes the projectile after 10 seconds, regardless
         // of whether it collided with anything. This prevents
         // instances from staying in the scene indefinitely.
         Destroy(gameObject,10);
    }
    void OnCollisionEnter()
    {
        // When the projectile hits something, create an explosion
        // and remove the projectile.
        Instantiate(explosion,transform.position,transform.rotation);
        Destroy(gameObject);
    }
}
```

The script instantiates the explosion at the projectile’s current position and removes the projectile GameObject when the projectile collides with something.

To use the script, attach it to the projectile prefab asset:

1. Select the `Projectile` prefab asset and open it in [prefab editing mode](EditingInPrefabMode.md).
2. In the **Inspector**, drag the `Projectile` script onto it.
3. Drag the explosion prefab asset into the **Explosion** field.

## Create a script to launch the projectiles

To launch the projectiles, you need to create a script that instantiates projectiles when the fire key is pressed, and add that script to a GameObject.

Create a script called `FireProjectile` and add the following contents to it:

```
using UnityEngine;
using UnityEngine.InputSystem;

public class FireProjectile : MonoBehaviour
{
    // Only GameObjects with a Rigidbody can be assigned as the projectile.
    public Rigidbody projectile;
    // Speed of the projectile when fired.
    // This is a public variable so it can be adjusted in the Unity Editor.
    public float speed = 4;
    // Update is called once per frame
    // This method checks for input and fires a projectile if the attack action is pressed.
    void Update()
    {
        // Check if the "Attack" action was pressed this frame
        // If it was, instantiate a projectile at the player's position and set its velocity.
        if (InputSystem.actions.FindAction("Attack").WasPressedThisFrame())
        {
            Rigidbody p = Instantiate(projectile, transform.position, transform.rotation);
            p.linearVelocity = transform.forward * speed;
        }
    }
}
```

This script uses `Instantiate` to launch a projectile. When making a public prefab variable, the variable type can be a GameObject, or it can be any valid component type (either a built-in Unity component or one of your own MonoBehaviour scripts).

For component type variables (such as Rigidbody, Collider, and Light), you can only assign GameObjects of that component type to the variable, and the `Instantiate` function returns a reference to that specific component on the new GameObject instance.

## Attach the launcher script to a GameObject

You must attach the script to a GameObject to use it. To do so:

1. Make sure that your scene has a ground GameObject for the projectile to collide with. If you’re using an empty project, create a Plane GameObject: right-click on the Hierarchy and select **3D Object** > **Plane**.
2. Create a Cube GameObject: right-click on the Hierarchy and select **3D Object** > **Cube** and position it over the plane.
3. Select the cube, and in its Inspector, delete the **Box **Collider**** component. If you don’t delete this component, the projectiles collide with the cube before hitting the ground.
4. Drag the `FireProjectile` script onto the cube.
5. Drag the `Projectile` prefab asset into the **Projectile** field of the Fire Projectile script.
6. Enter Play mode, and then click your mouse button.

The cube fires the sphere projectiles and the explosion happens when they collide with the ground.

![Projectile and explosion prefabs being instantiated and destroyed.](../uploads/Main/prefab-projectile-instantiate.png)

Projectile and explosion prefabs being instantiated and destroyed.

Note that any instantiated objects appear in the Hierarchy with `(Clone)` appended to the name.

## Additional resources

* [Build a structure with prefabs](instantiating-prefabs-structure.md)
* [`Instantiate` API reference](../ScriptReference/Object.Instantiate.md)