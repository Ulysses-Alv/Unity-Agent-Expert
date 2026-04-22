# Awaitable code example reference

The examples in this reference demonstrate `Awaitable` solutions to common scenarios encountered when writing asynchronous code.

## Asynchronous tests

Unity’s [Test Framework](test-framework/test-framework-introduction.md) doesn’t recognize `Awaitable` as a valid test return type. However, the following example shows how you can use the `Awaitable` implementation of `IEnumerator` to write async tests:

```
[UnityTest]
public IEnumerator SomeAsyncTest(){
    async Awaitable TestImplementation(){
        // test something with async / await support here
    };
    return TestImplementation();
}
```

## Frame coroutines

You can use the frame-related async methods in the `Awaitable` class to create asynchronous Unity coroutines as an alternative to iterator-based coroutines:

```
async Awaitable SampleSchedulingJobsForNextFrame()
{
    // Wait until end of frame to avoid competing over resources with other Unity subsystems
    await Awaitable.EndOfFrameAsync(); 
    var jobHandle = ScheduleSomethingWithJobSystem();
    // Let the job execute while the next frame starts
    await Awaitable.NextFrameAsync();
    jobHandle.Complete();
    // Use results of computation
}

JobHandle ScheduleSomethingWithJobSystem()
{
    ...
}
```

### Conditional wait

In iterator-based coroutines, [WaitUntil](../ScriptReference/WaitUntil.md) suspends a coroutine execution until a delegate evaluates `true`. You can create equivalent behavior for an `Awaitable`-returning asynchronous method by making it wait until a condition changes using a cancellation token:

```
public static async Awaitable AwaitableUntil(Func<bool> condition, CancellationToken cancellationToken)
{
   while(!condition()){
     cancellationToken.ThrowIfCancellationRequested();
     await Awaitable.NextFrameAsync();
  }
}
```

You can then pass in a cancellation token as follows:

```
cancellationTokenSource = new CancellationTokenSource();
currentTask = AwaitableUntil(myCondition, cancellationTokenSource.Token);
```

## Loading resources asynchronously

You can `await` an asynchronous resource loading operation so that it doesn’t block the main thread:

```
public async Awaitable LoadResourcesAsync()
{
    // Load texture resource asynchronously
    var operation = Resources.LoadAsync("my-texture");
    // Return control to the main thread while the resource loads
    await operation;
    var texture = operation.asset as Texture2D;
}
```

## Composition

You can `await` multiple different `await`-compatible types in the same method:

```
public async Awaitable Bar()
{
    await CallSomeThirdPartyAPIReturningDotnetTask();
    await Awaitable.NextFrameAsync();
    await SceneManager.LoadSceneAsync("my-scene");
    await SomeUserCodeReturningAwaitable();
    ...
}
```

## Wrapping Awaitable in .NET Task

To work around some of the limitations of `Awaitable` you can wrap it in a .NET `Task`. This incurs the cost of an allocation but gives you access to methods such as `WhenAll` and `WhenAny` from the `Task` API. To do this you can write your own custom `AsTask` extension methods as follows:

```
// Implement custom AsTask extension methods to wrap Awaitable in Task
public static class AwaitableExtensions
    {
        public static async Task AsTask(this Awaitable a)
        {
            await a;
        }

        public static async Task<T> AsTask<T>(this Awaitable<T> a)
        {
            return await a;
        }
    }
```

## Await a result multiple times

A major difference between `Awaitable` and `Task` is that `Awaitable` objects are pooled to reduce allocations. You can’t safely `await` an `Awaitable`-returning method that completes with a result multiple times because once returned the original `Awaitable` object is returned to the pool.

### Unsafe version

The following code is **not safe** and will lead to exceptions and deadlocks:

```
async Awaitable Bar(){
  var taskWithResult = SomeAwaitableReturningFunction();
  var awaitOnce = await taskWithResult;
  // Do something
  // The following will cause errors because at this point taskWithResult has already been pooled back
  var awaitTwice = await taskWithResult;
 }
```

### Safe version

This is one of the scenarios where you can [wrap Awaitable in a Task](#awaitable-as-task) at the cost of an allocation. The `Task` can then be safely awaited multiple times:

```
// Implement custom AsTask extension methods to wrap Awaitable in Task
public static class AwaitableExtensions
    {
        public static async Task AsTask(this Awaitable a)
        {
            await a;
        }

        public static async Task<T> AsTask<T>(this Awaitable<T> a)
        {
            return await a;
        }
    }

async Awaitable Bar(){
  var taskWithResult = SomeAwaitableReturningFunction();
  // Wrap the returned Awaitable in a Task
  var taskWithResultAsTask = taskWithResult.AsTask();
  // The task can now be safely awaited multiple times, at the cost of allocating
  var awaitOnce = await taskWithResultAsTask;
  // Do something
  var awaitTwice = await taskWithResultAsTask;
 }
```

## Additional resources

* [Comparing Unity Awaitable and .NET Task](async-awaitable-introduction.md)
* [Async continuation and completion](async-awaitable-continuations.md)
* [Awaitable API reference](../ScriptReference/Awaitable.md)