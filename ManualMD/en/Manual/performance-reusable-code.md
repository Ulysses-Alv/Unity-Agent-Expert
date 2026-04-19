# Pooling and reusing objects

Object pooling and reuse is a programming pattern that returns instances of frequently-used objects to a pool, from which they can be retrieved and used again. Object pooling reduces the overhead caused by repeatedly instantiating and destroying new object instances and limits the overall number of allocations and deallocations. This helps minimize garbage collection (GC) overhead and reduce load on the CPU.

You can pool **GameObjects**, but you can also pool other types you use frequently, including collection types such as arrays, lists, and dictionaries. You can write your own pooling logic, but for convenience Unity provides a set of APIs in the `UnityEngine.Pool` namespace with implementations for pooling GameObjects and several collection types.

**Note**: Like many other core Unity APIs, the `UnityEngine.Pool` APIs are not thread-safe, and can only be called safely from the main thread.

## Pooling and reusing GameObjects

Some types of objects in games, such as projectiles, might appear repeatedly and often, even though only a few of them are ever in a **scene** at the same time. Instead of instantiating a new projectile object from a **prefab** every time one is fired, you can create a pool of projectile objects at the start of the game, and then reuse them as needed.

### Creating an object pool

Unity’s [`ObjectPool<T>`](../ScriptReference/Pool.ObjectPool_1.md) class provides an implementation of the reusable object pool pattern, and is the simplest way to implement an object pool in your application. An instance of `ObjectPool<T>` manages a pool of objects of type `T`. When you create an `ObjectPool<T>` through the [`ObjectPool<T>` constructor](../ScriptReference/Pool.ObjectPool_1-ctor.md), you provide parameters to define the following key aspects of the pool:

* The starting capacity of the pool.
* The maximum capacity the pool can grow to.
* Whether collection safety checks are performed to prevent double-releasing the same object instance back to the pool.
* A series of C# [`Action<T>`](https://learn.microsoft.com/en-us/dotnet/api/system.action-1?view=net-10.0) delegates, which are callbacks Unity invokes when performing the following key actions on the pool:
  + A callback to invoke on retrieving an object from the pool.
  + A callback to invoke on returning an object to the pool.
  + A callback to invoke on creating a new object instance when one is requested and the pool is empty.
  + A callback to invoke on destroying an object instance when one is returned to the pool but the pool is at maximum capacity.

### Retrieving objects from the pool

You call [`ObjectPool<T>.Get`](../ScriptReference/ScriptRef:Pool.ObjectPool_1.Get.md) to retrieve an object from the pool. This action also invokes the method you provided as the `actionOnGet` parameter to the `ObjectPool` constructor, which gives you the opportunity to prepare and activate the object for use.

If you request an object from the pool when the pool is empty, a new object instance is created. The creation process invokes the method you provided as the `createFunc` parameter to the `ObjectPool` constructor.

`ObjectPool<T>.Get` has an overload `Get(out PooledObject<T> pooledObject)`. When the `PooledObject` is disposed, it releases the instance automatically. This is useful for short-lived scopes.

For more information and code examples, refer to the [`PooledObject<T>.Get` API reference](../ScriptReference/Pool.PooledObject_1.Get.md).

### Returning objects to the pool

When you’re finished with an object, you return it to the pool by calling [`ObjectPool<T>.Release`](../ScriptReference/Pool.ObjectPool_1.Release.md) with the object you’re returning as the argument. This action also invokes the method you provided as the `actionOnRelease` parameter to the `ObjectPool` constructor, which gives you the opportunity to deactivate and reset state on the returned object.

If you return an object to the pool when the pool is full, the object is destroyed. The destruction process invokes the method you provided as the `actionOnDestroy` parameter to the `ObjectPool` constructor.

Passing `true` as the `collectionCheck` parameter to the constructor enables collection safety checking. If collection checks are enabled and you try to return an object that’s already in the pool, Unity throws an exception. Even when enabled, collection checking is only available in the Editor and not in Player builds. Enabling collection checking adds overhead but catches double-release errors and is recommended during development.

### Resetting used objects

Objects you retrieve from the pool can change their state during use. For example, a pooled enemy GameObject might change its position, health, animation state, and physics state while active in the scene. It’s important to reset this state when the object is returned to the pool, so that when it’s retrieved again later, it starts in the same state as on its first use.

Depending on the state you’re resetting, you might choose to reset the object either when you get it from the pool in the `actionOnGet` delegate, or when you release it back to the pool in the `actionOnRelease` delegate. Another option is to implement a custom deactivation method on the pooled object itself and call those methods from the pool’s get or release actions.

The state you need to reset depends on the specific behavior and components of your pooled object. Common reset tasks on release include stopping coroutines, unsubscribing from events, resetting physics states, clearing animations, and stopping **particle** systems. It’s also important to deactivate GameObjects on release to prevent them receiving [`Update`](../ScriptReference/MonoBehaviour.Update.md) calls while inactive in the pool.

### Example: pooling projectiles

A typical implementation of an `ObjectPool<T>` for projectiles has the following key stages:

1. Create a pool of inactive projectile objects at the start of the game. Provide the constructor with the required delegates to invoke as callbacks when creating new a pooled object, retrieving an object from the pool, releasing an object back to the pool, and destroying an object because the pool is full.
2. When a projectile is needed (for example, when the user fires a weapon), retrieve an inactive projectile from the pool with `ObjectPool<T>.Get`. From the corresponding `actionOnGet` callback, activate the retrieved object, and initialize any relevant properties such as its position and direction.
3. When the projectile is no longer needed (for example, when it hits a target or goes out of bounds), return it to the pool with `ObjectPool<T>.Release`. It’s common to give the projectile a reference to its own pool so it can release itself when done. From the corresponding `actionOnRelease` callback, deactivate the returned object, and reset any relevant properties such as its velocity and animation state.

The following is the script component code for a projectile that’s pooled in an `ObjectPool<T>`. Note that the projectile has a public property to give it a reference to its own pool, which it uses to release itself back to the pool after a time delay:

```
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Pool;
public class RevisedProjectile : MonoBehaviour
{
    // Deactivate after delay
    [SerializeField] private float timeoutDelay = 3f;

    private IObjectPool<RevisedProjectile> objectPool;

    // Public property to give the projectile a reference to its ObjectPool
    public IObjectPool<RevisedProjectile> ObjectPool { set => objectPool = value; }

    public void Deactivate()
    {
        StartCoroutine(DeactivateRoutine(timeoutDelay));
    }

    IEnumerator DeactivateRoutine(float delay)
    {
        yield return new WaitForSeconds(delay);

        // Reset the moving Rigidbody
        Rigidbody rBody = GetComponent<Rigidbody>();
        rBody.linearVelocity = new Vector3(0f, 0f, 0f);
        rBody.angularVelocity = new Vector3(0f, 0f, 0f);

        // Release the projectile back to the pool
        objectPool.Release(this);
    }
}
```

The following is the code for a component attached to the gun turret that fires the projectile and manages the pool of projectile objects. The gun calls `Get` on the object pool to retrieve a pooled projectile when the user shoots:

```
using UnityEngine;
using UnityEngine.Pool;
using UnityEngine.Events;

public class RevisedGun : MonoBehaviour
{
    [Tooltip("Prefab to shoot")]
    [SerializeField] private RevisedProjectile projectilePrefab;
    [Tooltip("Projectile force")]
    [SerializeField] private float muzzleVelocity = 1500f;
    [Tooltip("End point of gun where shots appear")]
    [SerializeField] private Transform muzzlePosition;
    [Tooltip("Time between shots / smaller = higher rate of fire")]
    [SerializeField] private float cooldownWindow = 0.1f;

    [SerializeField] private UnityEvent m_GunFired;

    // Stack-based ObjectPool available with Unity 2021 and above
    private IObjectPool<RevisedProjectile> objectPool;

    // Throw an exception if we try to return an existing item, already in the pool
    [SerializeField] private bool collectionCheck = true;

    // Extra options to control the pool capacity and maximum size
    [SerializeField] private int defaultCapacity = 20;
    [SerializeField] private int maxSize = 100;

    private float nextTimeToShoot;

    private void Awake()
    {
        objectPool = new ObjectPool<RevisedProjectile>(CreateProjectile,
            OnGetFromPool, OnReleaseToPool, OnDestroyPooledObject,
            collectionCheck, defaultCapacity, maxSize);
    }

    // Invoked when creating an item to populate the object pool
    private RevisedProjectile CreateProjectile()
    {
        RevisedProjectile projectileInstance = Instantiate(projectilePrefab);
        projectileInstance.ObjectPool = objectPool;
        return projectileInstance;
    }

    // Invoked when returning an item to the object pool
    private void OnReleaseToPool(RevisedProjectile pooledObject)
    {
        pooledObject.gameObject.SetActive(false);
    }

    // Invoked when retrieving the next item from the object pool
    private void OnGetFromPool(RevisedProjectile pooledObject)
    {
        pooledObject.gameObject.SetActive(true);
    }

    // Invoked when the maximum number of pooled items is exceeded (i.e. destroy the pooled object)
    private void OnDestroyPooledObject(RevisedProjectile pooledObject)
    {
        Destroy(pooledObject.gameObject);
    }

    private void FixedUpdate()
    {
        // Invokes the Shoot method if the delay time is exceeded
        if (Input.GetButton("Fire1") && Time.time > nextTimeToShoot && objectPool != null)
        {
            Shoot();
        }
    }

    private void Shoot()
    {
        // Get a pooled object instead of instantiating
        RevisedProjectile bulletObject = objectPool.Get();

        if (bulletObject == null)
            return;

        // Align to gun barrel/muzzle position
        bulletObject.transform.SetPositionAndRotation(muzzlePosition.position, muzzlePosition.rotation);

        // Move projectile forward
        bulletObject.GetComponent<Rigidbody>().AddForce(bulletObject.transform.forward * muzzleVelocity, ForceMode.Acceleration);

        // Turn off after a few seconds
        bulletObject.Deactivate();

        // Set cooldown delay
        nextTimeToShoot = Time.time + cooldownWindow;

        m_GunFired.Invoke();
    }
}
```

A full working example project using this code, including all necessary prefabs and a working turret that fires the pooled projectiles, is available for download as a free package from the [Asset Store](https://assetstore.unity.com/packages/essentials/tutorial-projects/level-up-your-code-with-design-patterns-and-solid-289616). For more information, refer to [Use object pooling to boost performance of C# scripts in Unity](https://learn.unity.com/course/design-patterns-unity-6/tutorial/use-object-pooling-to-boost-performance-of-c-scripts-in-unity#67a0d722edbc2a14cfcb7400).

### Handling scene unload for pooled objects

By default, an object pool is tied to the scene it was created in. When that scene is unloaded, Unity destroys all GameObjects in that scene, including the pool and any pooled instances.

You can make an object pool persist after scene reload by placing the pool manager under a root GameObject on which you call [`object.DontDestroyOnLoad`](../ScriptReference/Object.DontDestroyOnLoad.md). You can call [`transform.SetParent`](../ScriptReference/Transform.SetParent.md) to reparent pooled items back under the pool root on release so they’re not destroyed with a scene.

If you might have active pooled instances in the scene when unloading, ensure they release themselves in [`OnDisable`](../ScriptReference/MonoBehaviour.OnDisable.md) and [`OnDestroy`](../ScriptReference/MonoBehaviour.OnDestroy.md) or register a [`SceneManager.activeSceneChanged`](../ScriptReference/SceneManagement.SceneManager-activeSceneChanged.md) callback to recall or destroy items.

Avoid long-lived references to pooled instances when inactive. Use [`ObjectPool<T>.Clear`](../ScriptReference/Pool.ObjectPool_1.Clear.md) to destroy all inactive instances, for example, between levels.

## Pooling and reusing collection types

You can also pool and reuse collection objects from the [`System.Collection`](https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic?view=net-6.0) namespace, such as Lists and Dictionaries.

Collection classes expose a `Clear` method, which removes a collection’s values but doesn’t release the memory allocated to the collection. This is useful if you want to allocate temporary helper collections for complex computations.

The following example code allocates the `nearestNeighbors` `List` once per frame to collect a set of data points, which is an inefficient bad practice:

```
// Bad C# script example. This Update method allocates a new `List` every frame.
void Update() {

    List<float> nearestNeighbors = new List<float>();

    findDistancesToNearestNeighbors(nearestNeighbors);

    nearestNeighbors.Sort();

    // … use the sorted list somehow …
}
```

Instead, you can hoist this `List` out of the method and into the containing class and clear and reuse it each time, so that your code doesn’t need to allocate a new `List` each frame:

```
// Good C# script example. This method re-uses the same List every frame.
List<float> m_NearestNeighbors = new List<float>();

void Update() {

    m_NearestNeighbors.Clear();

    findDistancesToNearestNeighbors(NearestNeighbors);

    m_NearestNeighbors.Sort();

    // … use the sorted list somehow …
}
```

This example code retains and reuses the `List` instances memory across multiple frames. The code only allocates new memory when the `List` needs to expand.

The `UnityEngine.Pool` namespace also provides ready-made pool implementations for several collection types, including `List<T>`, `HashSet<T>`, and `Dictionary<TKey, TValue>`. They operate on the same principle as `ObjectPool<T>`, allowing you to get and release pooled collection object instances as needed. For more information and examples, refer to the API reference documentation for [`CollectionPool<T>`](../ScriptReference/Pool.CollectionPool_1.md) and the other collection type-specific pool classes under the same namespace.

## Additional resources

* [Learn Tutorial: Use object pooling to boost performance of C# scripts in Unity](https://learn.unity.com/course/design-patterns-unity-6/tutorial/use-object-pooling-to-boost-performance-of-c-scripts-in-unity#67a0d722edbc2a14cfcb7400)
* [Managed memory introduction](performance-managed-memory-introduction.md)
* [Reference type management](performance-reference-types.md)
* [Optimizing arrays](performance-optimizing-arrays.md)
* [Instantiate projectiles and explosions](instantiating-prefabs-projectiles.md).