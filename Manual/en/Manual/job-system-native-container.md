# Introduction to NativeContainer

`NativeContainer` is a thread-safe C# wrapper for native memory. `NativeContainer` objects also allow a job to access data shared with the [main thread](job-system-overview.md#multithreading) rather than working with a copy.

## Types of NativeContainers

The `Unity.Collections` namespace contains the following built-in `NativeContainer` objects:

* [`NativeArray`](../ScriptReference/Unity.Collections.NativeArray_1.md): An unmanaged array which exposes a buffer of native memory to managed code
* [`NativeSlice`](../ScriptReference/Unity.Collections.NativeSlice_1.md): Gets a subset of a `NativeArray` from a particular position to a certain length.

**Note:** The [Collections package](https://docs.unity3d.com/Packages/com.unity.collections@latest/) contains additional `NativeContainer`s. For a full list of the additional types, see the Collections documentation on [Collection types](https://docs.unity3d.com/Packages/com.unity.collections@latest/index.html?subfolder=/manual/collection-types.html).

## Read and write access

By default, when a job has access to a `NativeContainer` instance, it has both read and write access. This configuration can slow performance. This is because the job system doesn’t allow you to schedule a job that has write access to a `NativeContainer` instance at the same time as another job that’s writing to it.

However, If a job doesn’t need to write to a `NativeContainer` instance, you can mark the `NativeContainer` with the `[ReadOnly]` attribute, like so:

```
[ReadOnly]
public NativeArray<int> input;
```

In the above example, you can execute the job at the same time as other jobs that also have read-only access to the first `NativeArray`.

## Memory allocators

When you create a `NativeContainer` instance, you must specify the memory allocation type that you need. The allocation type you use depends on how long you would like to keep the native container available for. This way you can tailor the allocation to get the best performance possible in each situation.

There are three [Allocator](../ScriptReference/Unity.Collections.Allocator.md) types for `NativeContainer` memory allocation and release. You must specify the appropriate one when instantiating a `NativeContainer` instance:

* `Allocator.Temp`: The fastest allocation. Use it for allocations with a lifespan of one frame or fewer. You can’t use `Temp` to pass allocations to `NativeContainer` instances stored in a job’s member field.
* `Allocator.TempJob`: A slower allocation than `Temp` but faster than `Persistent`. Use it for thread-safe allocations within a lifespan of four frames. **Important**: You must `Dispose` of this allocation type within four frames, or the console prints a warning, generated from the native code. Most small jobs use this allocation type.
* `Allocator.Persistent`: The slowest allocation but can last as long as you need it to, and if necessary, throughout the application’s lifetime. It’s a wrapper for a direct call to [`malloc`](http://www.cplusplus.com/reference/cstdlib/malloc/). Longer jobs can use this NativeContainer allocation type. Don’t use `Persistent` where performance is essential.

For example:

```
NativeArray<float> result = new NativeArray<float>(1, Allocator.TempJob);
```

**Note:** The number 1 in the example above indicates the size of the NativeArray. In this case, it has only one array element because it only stores one piece of data in its result.

## NativeContainer safety system

The [safety system](job-system-overview.md#safety-system) is built into all `NativeContainer` instances. It tracks what reads or writes to any `NativeContainer` instance, and uses that information to enforce certain rules on the use of `NativeContainer`s that makes them behave in a deterministic way across multiple jobs and threads.

For example, if two independent scheduled jobs write to the same `NativeArray`, this is unsafe because you can’t predict which job executes first. This means that you won’t know which of the jobs will overwrite data from the other. The safety system throws an exception with a clear error message that explains why and how to solve the problem, when you schedule the second job.

If you want to schedule two jobs that write to the same `NativeContainer` instance, you can schedule the [jobs with a dependency](job-system-job-dependencies.md). The first job writes to the `NativeContainer`, and once it has finished executing, the next job safely reads and writes to that same `NativeContainer`. Introducing the dependency guarantees that the jobs always execute in a consistent order and that the resulting data in the `NativeContainer` is deterministic.

The safety system allows multiple jobs to read from the same data in parallel.

These read and write restrictions also apply when accessing data from the main thread. For example, if you try to read the contents of a NativeContainer before the job that writes to it has completed, the safety system throws an error. Likewise, if you try to write to a NativeContainer while there are still pending jobs that read or write to it, then the safety system also throws an error.

Also, because NativeContainers don’t implement [`ref return`](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/ref-returns), you can’t directly change the contents of a `NativeContainer`. For example, `nativeArray[0]++;` is the same as writing `var temp = nativeArray[0]; temp++;` which doesn’t update the value in `nativeArray`.

Instead, you must copy the data from the index into a local temporary copy, modify that copy, and save it back. For example:

```
MyStruct temp = myNativeArray[i];
temp.memberVariable = 0;
myNativeArray[i] = temp;
```

## Additional resources

* [Implement a custom `NativeContainer`](job-system-custom-nativecontainer.md)