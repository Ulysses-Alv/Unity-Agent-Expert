# Adding profiler counters to your code

To add a **profiler counter**, create **scripts** to do the following:

* [Create and define a new counter](#creating-and-defining-a-counter).
* [Update the counter’s value](#updating-a-counters-value).

The code examples in these sections add a Profiler counter to track the total number of **particles** that Unity created for every instance of a **GameObject**’s trail effects. In these examples, the GameObject’s name is “Tank”.

## Prerequisites

Some examples use [the Profiling Core package](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest), which you must install before you start. The Unity Profiling Core package isn’t discoverable in the Package Manager UI because it’s a core package. To install the package [add it by its name](upm-ui-quick.md), which is `com.unity.profiling.core`.

## Creating and defining a counter

To create a new counter:

1. Write a script that defines the value type of the new counter.
2. Assign a name and a unit to this type.

**Important:** When you create a counter you must specify which **Profiler category** your new counter belongs to. To do this, use an existing Unity category from the [`ProfilerCategory`](../ScriptReference/Unity.Profiling.ProfilerCategory.md) class. The script in the following example uses the existing `ProfilerCategory.Scripts` category.

The Profiler counters API supports push and pull operations. You can push the value of the counter to the Profiler, or the Profiler can pull the value at the end of the frame.

If your data changes infrequently, for example once per frame, use the [`ProfilerCounter`](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest?subfolder=/api/Unity.Profiling.ProfilerCounter-1.html) API to push the counter data to the Profiler. If your data changes multiple times per frame, use the `ProfilerCounterValue` API. This makes the Profiler automatically pick up the last value at the end of the frame.

The following example script defines the [`ProfilerCounterValue`](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest?subfolder=/api/Unity.Profiling.ProfilerCounter-1.html) `TankTrailParticleCount`, with the name `Tank Trail Particles` and a unit of [`ProfilerMarkerDataUnit.Count`](../ScriptReference/Unity.Profiling.ProfilerMarkerDataUnit.Count.md):

```
public static class GameStats
{
   public static readonly ProfilerCategory TanksCategory = ProfilerCategory.Scripts;

   public const string TankTrailParticleCountName = "Tank Trail Particles";
   public static readonly ProfilerCounterValue<int> TankTrailParticleCount =
       new ProfilerCounterValue<int>(TanksCategory, TankTrailParticleCountName, ProfilerMarkerDataUnit.Count,
           ProfilerCounterOptions.FlushOnEndOfFrame | ProfilerCounterOptions.ResetToZeroOnFlush);
}
```

The options `FlushOnEndOfFrame` and `ResetToZeroOnFlush` automatically send the counter to the Profiler data stream and reset the Count value to zero at the end of the frame.

## Updating a counter’s value

To update the value of a counter, create a MonoBehaviour script that sets the value of a counter you have defined.

The following MonoBehaviour script counts the number of trail particles that belong to an assigned GameObject every frame in the Update function. To do this, it uses the counter called `TankTrailParticleCount`.

The following example script also creates a public property called Trail **Particle System** (`m_TrailParticleSystem`) in the **Inspector**:

```
 using UnityEngine;

 class TankMovement : MonoBehaviour
 {
    public ParticleSystem m_TrailParticleSystem;

    void Update()
    {
        GameStats.TankTrailParticleCount.Value += m_TrailParticleSystem.particleCount;
    }
 }
```

### Adding Profiler counters to a custom Profiler module

To display a Profiler counter in a custom Profiler module via code, you must create a new `ProfilerModule` script and define the module’s properties including the counters it displays, its name, and its icon. For information on how to do this refer to [Creating Profiler modules in code](profiler-create-modules.md).

## Additional resources

* [Adding profiling information to your code introduction](profiler-adding-information-code-intro.md)
* [Adding profiler markers to your code](profiler-add-markers-code.md)
* [Customizing Profiler modules](profiler-customizing.md)
* [Profiling Core package](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest)