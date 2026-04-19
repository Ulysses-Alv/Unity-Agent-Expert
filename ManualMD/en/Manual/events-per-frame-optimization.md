# Using a custom update manager

Unity’s built-in per-frame [event function](event-functions.md) updates such as [`Update`](../ScriptReference/MonoBehaviour.Update.md), [`FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md) and [`LateUpdate`](../ScriptReference/MonoBehaviour.LateUpdate.md) can impact performance at scale. Although the corresponding callbacks are invoked on your C# MonoBehaviour **scripts**, the function calls originate from Unity’s native code. Unity has to maintain internal lists to track which objects to call these update functions on. MonoBehaviour script instances are added to or removed from these lists when they are enabled or disabled, respectively.

While it’s convenient to add the appropriate callbacks to every MonoBehaviour instance in your project that requires them, this becomes more inefficient as the number of callbacks grows. There is a small but significant overhead to invoking managed-code callbacks from native code, which leads to the following consequences:

* Degraded frame times when invoking large numbers of `Update` callbacks.
* Degraded instantiation times when [instantiating prefabs](instantiating-prefabs.md) that contain large numbers of MonoBehaviours, due to the performance overhead of invoking `Awake` and `OnEnable` callbacks on each component in a **prefab**.

To avoid these issues, instead of relying on built-in callbacks you can create a global custom update manager singleton instance and have MonoBehaviour scripts, or even standard C# objects, subscribe to it. This way, the update manager singleton can distribute `Update`, `LateUpdate`, and other callbacks to all objects that have subscribed to them, and all update code stays in the managed layer. This has the additional benefit of allowing code to unsubscribe from callbacks when they have no operation to perform, which reduces the number of functions that must be called each frame.

## When to use a custom update manager

A custom update manager can be beneficial when the number of MonoBehaviour instances with per-frame update callbacks reaches the hundreds or thousands.

You can improve performance significantly by eliminating callbacks that rarely execute. Consider the following example:

```
void Update() {
    if(!someVeryRareCondition) { return; }
// … some operation …
}
```

If your project has many MonoBehaviours with `Update` callbacks similar to this, then a significant amount of the time consumed running `Update` callbacks is spent switching between native and managed code domains for MonoBehaviour execution that then exits immediately. If these classes instead subscribe to a global update manager only while `someVeryRareCondition` is true, and unsubscribe thereafter, then less time is spent both on switching code domains and evaluating the rare condition.

**Important**: A custom update manager is not a one-size-fits-all solution. It’s important to [profile](Profiler.md) your project to determine its specific performance issues and whether a custom update manager is appropriate. Depending on the specific performance bottlenecks in your project, other ways to optimize performance include converting your project to use the Entity Component System (ECS) architecture, or [customizing the Player loop](player-loop-customizing.md).

## Example custom update manager

To implement a custom update manager, first create a C# script to define the interface as follows:

```
public interface IUpdatable
{
    void CustomUpdate(float deltaTime);
}
```

You can then create a MonoBehaviour script for the update manager singleton. The update manager implements the built-in `Update` callback and then other MonoBehaviour script components can subscribe to this update manager rather than to `Update` directly:

```
// Singleton update manager. Attach to a GameObject in your scene.

using System.Collections.Generic;
using UnityEngine;

public class UpdateManager : MonoBehaviour
{
    private static UpdateManager _instance;
    public static UpdateManager Instance => _instance;

    private readonly List<IUpdatable> updatables = new List<IUpdatable>();

    void Awake()
    {
        if (_instance == null)
            _instance = this;
        else
            Destroy(gameObject);
    }

    public void Register(IUpdatable updatable)
    {
        if (!updatables.Contains(updatable))
            updatables.Add(updatable);
    }

    public void Unregister(IUpdatable updatable)
    {
        updatables.Remove(updatable);
    }

    void Update()
    {
        float deltaTime = Time.deltaTime;
        // Optionally: for performance, iterate backwards if you allow removal during iteration
        foreach (var updatable in updatables)
        {
            updatable.CustomUpdate(deltaTime);
        }
    }
}
```

Finally, create a MonoBehaviour script component that registers itself with the update manager instance on enable and de-registers itself on disable. The following example uses the custom update callback to move the parent **GameObject**:

```
// Script component. Attach to a GameObject in your scene to move it on each custom update.

using UnityEngine;

public class MyMovingObject : MonoBehaviour, IUpdatable
{
    void OnEnable()
    {
        UpdateManager.Instance.Register(this);
    }

    void OnDisable()
    {
        if (UpdateManager.Instance != null)
            UpdateManager.Instance.Unregister(this);
    }

    public void CustomUpdate(float deltaTime)
    {
        // Your update logic here
        transform.position += Vector3.right * deltaTime;
    }
}
```

## Additional resources

* 📚 **Documentation**: [Event function execution order](execution-order.md)
* 📚 **Documentation**: [Event functions](event-functions.md)
* 📚 **Documentation**: [10000 Update calls](https://unity.com/blog/engine-platform/10000-update-calls)