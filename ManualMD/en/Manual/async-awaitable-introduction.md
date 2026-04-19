# Introduction to asynchronous programming with Awaitable

The [`Awaitable`](../ScriptReference/Awaitable.md) class is a custom Unity type that can be awaited and used as an async return type in the C# asynchronous programming model. Most of Unity’s asynchronous APIs support the `async` and `await` pattern, including:

* Unity coroutines: [`NextFrameAsync`](../ScriptReference/Awaitable.NextFrameAsync.md), [`WaitForSecondsAsync`](../ScriptReference/Awaitable.WaitForSecondsAsync.md), [`EndOfFrameAsync`](../ScriptReference/Awaitable.EndOfFrameAsync.md), [`FixedUpdateAsync`](../ScriptReference/Awaitable.FixedUpdateAsync.md)
* Switching to [Background Thread](../ScriptReference/Awaitable.BackgroundThreadAsync.md) or [Main Thread](../ScriptReference/Awaitable.MainThreadAsync.md)
* All types inheriting from [`AsyncOperation`](../ScriptReference/AsyncOperation.md)
* [Unity Events](unity-events.md)
* [Async GPU Readback](../ScriptReference/Rendering.AsyncGPUReadback.md)

You can use the `Awaitable` class with both the `await` operator and as an `async` return type in your own code, as follows:

```
async Awaitable<List<Achievement>> GetAchievementsAsync()
{
    var apiResult = await SomeMethodReturningATask(); // or any await-compatible type
    List<Achievement> achievements = JsonConvert.DeserializeObject<List<Achievement>>(apiResult);
    return achievements;
}

async Awaitable ShowAchievementsView()
{
    ShowLoadingOverlay();
    List<Achievement> achievements = await GetAchievementsAsync();
    HideLoadingOverlay();
    ShowAchivementsList(achievements);
}
```

## Awaitable compared to .NET Task

`Awaitable` is designed to offer a more efficient alternative to .NET [`Task`](https://learn.microsoft.com/en-us/dotnet/api/system.threading.tasks.task?view=net-8.0) for asynchronous code in Unity projects. The efficiency of `Awaitable` comes with some important limitations compared to `Task`.

The most significant limitation is that `Awaitable` instances are pooled to limit allocations. Consider the following example:

```
class SomeMonoBehaviorWithAwaitable : MonoBehavior
{
    public async void Start()
    {
        while(true)
        {
            // do some work on each frame
            await Awaitable.NextFrameAsync();
        }
    }
}
```

Without pooling, each instance of the [`MonoBehavior`](class-monobehaviour.md) in this example would allocate an `Awaitable` object each frame, increasing [garbage collector workload](performance-garbage-collector.md) and degrading performance. To mitigate this, Unity returns the `Awaitable` object to the internal `Awaitable` pool once it’s been awaited.

**Important:** The pooling of `Awaitable` instances means it’s never safe to `await` more than once on an `Awaitable` instance. Doing so can result in undefined behavior such as an exception or a deadlock.

## Awaitable compared to .NET ValueTask

The .NET [`ValueTask<TResult>`](https://learn.microsoft.com/en-us/dotnet/api/system.threading.tasks.valuetask-1?view=net-8.0) offers some of the same key benefits and limitations of `Awaitable`. The typical recommended use for `ValueTask` is for asynchronous workloads that are expected to complete synchronously most of the time. For more information, refer to [Understanding the Whys, Whats, and Whens of ValueTask](https://devblogs.microsoft.com/dotnet/understanding-the-whys-whats-and-whens-of-valuetask/).

## Awaitable, Task, and ValueTask summary

The following table summarizes the feature comparison between Unity’s `Awaitable` class and .NET `Task` and `ValueTask`:

| **Feature** | `Task` | `ValueTask` | `UnityEngine.Awaitable` |
| --- | --- | --- | --- |
| **Required allocations** | **Many**.  Allocates on every call to a `Task`-returning method, increasing memory use and garbage collector workload. | **As-needed**.  Can be optimized with pooling. | **Minimal as-needed**.  Calling an `Awaitable`-returning method usually doesn’t allocate memory, since `Awaitable` instances are pooled by default. |
| **Safe to await multiple times** | **Yes**. | **No**.  Must convert to a `Task` with `ValueTask.AsTask`. | **No**.  Must convert to a `Task` with custom `AsTask` extension methods, refer to [Awaiting multiple times in the same method](async-awaitable-examples.md#await-multiple-times) in the code examples reference. |
| **Continuations run asynchronously** | **Yes**.  Using the synchronization context by default, otherwise using the `ThreadPool`. This increases latency when completing on the main thread in Unity because code must wait until the next frame `Update` to resume. | **Yes**.  Optimized for the case where awaited tasks complete synchronously. If they complete asynchronously, the continuation behavior is equivalent to `Task`. | **No**.  Continuation runs synchronously when completion is triggered, meaning code resumes immediately in the same frame in which completion is triggered. Refer to [Awaitable completion and continuation](async-awaitable-continuations.md) for more information. |
| **Completion can be triggered by code** | **Yes**.  Using [`TaskCompletionSource`](https://learn.microsoft.com/en-us/dotnet/api/system.threading.tasks.taskcompletionsource?view=net-8.0). | Not applicable in the typical use case, which is for tasks that mostly complete synchronously. | **Yes**.  Using [`AwaitableCompletionSource`](../ScriptReference/AwaitableCompletionSource.md). |
| **Can return a value** | **Yes**.  Using [`Task<TResult>`](https://learn.microsoft.com/en-us/dotnet/api/system.threading.tasks.task-1?view=net-8.0). | **Yes**.  Using [`ValueTask<TResult>`](https://learn.microsoft.com/en-us/dotnet/api/system.threading.tasks.valuetask-1?view=net-8.0). | **Yes**.  Using [`UnityEngine.Awaitable<T>`](../ScriptReference/Awaitable_1.md). |
| **Built-in support for `WaitAll` and `WaitAny`** | **Yes**. | **No**.  Must convert to `Task` with `ValueTask.AsTask`. | **No**.  Must convert to a `Task` with custom `AsTask` extension methods, refer to [Wrapping Awaitable in .NET Task](async-awaitable-examples.md#awaitable-as-task) in the code examples reference. |
| **Unity thread and update loop-aware execution scheduling** | **No**. | **No**. | **Yes**.  You can specify which thread an `Awaitable` resumes on with [`Awaitable.BackgroundThreadAsync`](../ScriptReference/Awaitable.BackgroundThreadAsync.md) and [`Awaitable.MainThreadAsync`](../ScriptReference/Awaitable.MainThreadAsync.md). You can also schedule work relative to the `Update` or `FixedUpdate` loops with [`Awaitable.NextFrameAsync`](../ScriptReference/Awaitable.NextFrameAsync.md) and [`Awaitable.FixedUpdateAsync`](../ScriptReference/Awaitable.FixedUpdateAsync.md). For more information, refer to [Awaitable completion and continuation](async-awaitable-continuations.md). |

## When to use Awaitable over Task or ValueTask

The choice of API depends on the performance profile of your asynchronous code, but in general:

* `Task` is the only choice when you need to await multiple times or from several consumers concurrently.
* `ValueTask` is a good choice if you have high-throughput asynchronous code that completes synchronously most of the time.
* `Awaitable` is a good choice when:
  + You don’t need to await your methods multiple times and expect them to mostly complete asynchronously.
  + You want your asynchronous tasks to have built-in support for Unity-specific concepts like the main thread and the [Update](time-per-frame-updates.md) and [FixedUpdate](fixed-updates.md) loops.

## Awaitable compared to iterator-based coroutines

`Awaitable` coroutines are usually more efficient than iterator-based coroutines, especially for cases where the iterator returns non-null values, such as [`WaitForFixedUpdate`](../ScriptReference/WaitForFixedUpdate.md).

However, the performance advantage of `Awaitable` coroutines reduces when you run many of them concurrently. For example, a MonoBehaviour such as the one in the previous code example, which awaits [`Awaitable.NextFrameAsync`](../ScriptReference/Awaitable.NextFrameAsync.md) in a `while` loop, is likely to cause performance problems if attached to every **GameObject** in a large project.

**Note**: You can safely `yield return` an [`Awaitable`](../ScriptReference/Awaitable.md) from a traditional iterator-based coroutine, but you can’t `yield return` an [`Awaitable<T0>`](../ScriptReference/Awaitable_1.md).

## Additional resources

* [Awaitable completion and continuation](async-awaitable-continuations.md)
* [Awaitable code example reference](async-awaitable-examples.md)