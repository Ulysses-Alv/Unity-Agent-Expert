# Parallel jobs

When you [schedule a job](job-system-creating-jobs.md#schedule-a-job), its `Execute` method runs a single time. However, there might be times when you need to perform the same operation on a lot of objects. To do this, use a ParallelFor job type, which inherits from [`IJobParallelFor`](../ScriptReference/Unity.Jobs.IJobParallelFor.md).

A ParallelFor job uses a [`NativeArray`](../ScriptReference/Unity.Collections.NativeArray_1.md) of data to act on as its data source. ParallelFor jobs run across multiple CPU cores. There’s one job per core, with each handling a subset of the workload.

`IJobParallelFor` behaves like [`IJob`](job-system-jobs.md), but instead of a single [`Execute`](../ScriptReference/Unity.Jobs.IJob.Execute.md) method, it invokes the `Execute` method once per item in the data source. There’s also an integer parameter index in the `Execute` method, which you can use to access and operate on a single element of the data source within the job implementation.

The following is an example of a ParallelFor job definition:

```
struct IncrementByDeltaTimeJob: IJobParallelFor
{
    public NativeArray<float> values;
    public float deltaTime;

    public void Execute (int index)
    {
        float temp = values[index];
        temp += deltaTime;
        values[index] = temp;
    }
}
```

## Schedule a ParallelFor job

To schedule a `ParallelFor` job, you must specify the length of the `NativeArray` data source that you want to split. The job system doesn’t know which `NativeArray` you want to use as the data source if there are several in the struct. The length also tells the job system how many `Execute` methods to expect.

In Unity’s native code, the scheduling of `ParallelFor` jobs is more complicated. When Unity schedules a `ParallelFor` job, the job system divides the work into batches to distribute between cores. Each batch contains a subset of `Execute` methods. The job system then schedules one job in Unity’s native job system per CPU core and passes that native job to the batches to complete.

![A ParallelFor job dividing batches across cores](../uploads/Main/jobsystem_parallelfor_job_batches.svg)

A ParallelFor job dividing batches across cores

When a native job completes its batches before others, it [steals](job-system-overview.md#work-stealing) remaining batches from the other native jobs. It only steals half of a native job’s remaining batches at a time, to ensure [cache locality](https://en.wikipedia.org/wiki/Locality_of_reference).

To optimize the process, you need to specify a batch count. The batch count controls how many jobs you get, and how fine-grained the redistribution of work between threads is. Having a low batch count, such as 1, gives you an even distribution of work between threads. However, it comes with some overhead, so sometimes it’s better to increase the batch count. Starting at 1 and increasing the batch count until there are negligible performance gains is a good strategy.

The following is an example of scheduling a ParallelFor job

Job code:

```
// Job adding two floating point values together
public struct MyParallelJob : IJobParallelFor
{
    [ReadOnly]
    public NativeArray<float> a;
    [ReadOnly]
    public NativeArray<float> b;
    public NativeArray<float> result;

    public void Execute(int i)
    {
        result[i] = a[i] + b[i];
    }
}
```

Main thread code:

```
NativeArray<float> a = new NativeArray<float>(2, Allocator.TempJob);

NativeArray<float> b = new NativeArray<float>(2, Allocator.TempJob);

NativeArray<float> result = new NativeArray<float>(2, Allocator.TempJob);

a[0] = 1.1f;
b[0] = 2.2f;
a[1] = 3.3f;
b[1] = 4.4f;

MyParallelJob jobData = new MyParallelJob();
jobData.a = a;
jobData.b = b;
jobData.result = result;

// Schedule the job with one Execute per index in the results array and only 1 item per processing batch
JobHandle handle = jobData.Schedule(result.Length, 1);

// Wait for the job to complete
handle.Complete();

// Free the memory allocated by the arrays
a.Dispose();
b.Dispose();
result.Dispose();
```

## ParallelForTransform jobs

A `ParallelForTransform` job is another type of `ParallelFor` job designed specifically for operating on [Transforms](class-Transform.md). It’s useful for working on Transform operations from jobs efficiently.

## Additional resources

* [ParallelForTransform API reference](../ScriptReference/Jobs.IJobParallelForTransform.md).
* [Create and run a job](job-system-creating-jobs.md)