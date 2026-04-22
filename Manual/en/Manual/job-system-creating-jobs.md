# Create and run a job

To create and successfully run a job, you must do the following:

* Create a job: Implement the [`IJob`](../ScriptReference/Unity.Jobs.IJob.md) interface.
* Schedule the job: Call the [`Schedule`](../ScriptReference/Unity.Jobs.IJobExtensions.Schedule.md) method on the job.
* Wait for the job to complete: It returns immediately if the job is already complete, and you can call the [`Complete`](../ScriptReference/Unity.Jobs.JobHandle.Complete.md) method on the job when you want to access the data.

## Create a job

To create a job in Unity, implement the [`IJob`](../ScriptReference/Unity.Jobs.IJob.md) interface. You can use your `IJob` implementation to schedule a single job that runs in parallel to any other jobs that are running.

`IJob` has one required method: [`Execute`](../ScriptReference/Unity.Jobs.IJob.Execute.md), which Unity invokes whenever a [worker thread](job-system-overview.md#multithreading) runs the job.

When you create a job, you can also create a [`JobHandle`](../ScriptReference/Unity.Jobs.JobHandle.md) for it, which other methods need to use to reference the job.

**Important:** There’s no protection against accessing non-readonly or **mutable** static data from within a job. Accessing this kind of data circumvents all safety systems and might crash your application or the Unity Editor.

When Unity runs, the job system makes a copy of scheduled job data, which prevents more than one thread from reading or writing the same data. Only data written to a `NativeContainer` can be accessed after the job finishes. This is because both the copy of the `NativeContainer` that the job uses and the original `NativeContainer` object point to the same memory. For more information, see the documentation on [Thread safe types](job-system-native-container.md).

When the job system picks up a job from its job queue, it runs the `Execute` method once on a single thread. Typically, the job system runs jobs on background threads, but it can choose the main thread if it becomes idle. For this reason, you should design your jobs to complete in under a frame.

## Schedule a job

To schedule a job, call `Schedule`. This puts the job into the job queue, and the job system begins executing the job once all its [dependencies](job-system-job-dependencies.md), if any, complete. Once scheduled, you can’t interrupt a job. You can only call `Schedule` from the main thread.

**Tip:** Jobs have a [`Run`](../ScriptReference/Unity.Jobs.IJobExtensions.Run.md) method that you can use in place of `Schedule` to immediately execute the job on the main thread. You can use this for debugging purposes.

## Complete the job

Once you call `Schedule` and the job system has executed a job, you can call the [`Complete`](../ScriptReference/Unity.Jobs.JobHandle.Complete.md) method on the `JobHandle` to access the data in the job. It’s best practice to call `Complete` as late as possible in your code. When you call `Complete`, the main thread can safely access the [`NativeContainer`](job-system-native-container.md) instance that the job was using. Calling `Complete` also cleans up the state in the safety system. Not doing so introduces a memory leak.

## Job examples

The following is an example of a job that adds two floating point values together. It implements [`IJob`](../ScriptReference/Unity.Jobs.IJob.md), uses a `NativeArray` to get the results of the job, and uses the [`Execute`](../ScriptReference/Unity.Jobs.IJob.Execute.md) method with the implementation of the job inside it:

```
using UnityEngine;
using Unity.Collections;
using Unity.Jobs;

// Job adding two floating point values together
public struct MyJob : IJob
{
    public float a;
    public float b;
    public NativeArray<float> result;

    public void Execute()
    {
        result[0] = a + b;
    }
}
```

The following example builds on the `MyJob` job to schedule a job on the main thread:

```
using UnityEngine;
using Unity.Collections;
using Unity.Jobs;

public class MyScheduledJob : MonoBehaviour
{
    // Create a native array of a single float to store the result. Using a
    // NativeArray is the only way you can get the results of the job, whether
    // you're getting one value or an array of values.
    NativeArray<float> result;
    // Create a JobHandle for the job
    JobHandle handle;

    // Set up the job
    public struct MyJob : IJob
    {
        public float a;
        public float b;
        public NativeArray<float> result;

        public void Execute()
        {
            result[0] = a + b;
        }
    }

    // Update is called once per frame
    void Update()
    {
        // Set up the job data
        result = new NativeArray<float>(1, Allocator.TempJob);

        MyJob jobData = new MyJob
        {
            a = 10,
            b = 10,
            result = result
        };

        // Schedule the job
        handle = jobData.Schedule();
    }

    private void LateUpdate()
    {
        // Sometime later in the frame, wait for the job to complete before accessing the results.
        handle.Complete();

        // All copies of the NativeArray point to the same memory, you can access the result in "your" copy of the NativeArray
        // float aPlusB = result[0];

        // Free the memory allocated by the result array
        result.Dispose();
    }

}
```

## Schedule and Complete best practices

It’s best practice to call `Schedule` on a job as soon as you have the data it needs, and don’t call `Complete` on it until you need the results.

You can schedule less important jobs in a part of the frame where they aren’t competing with more important jobs.

For example, if there is a period between the end of one frame and the beginning of the next frame where no jobs are running, and a one frame latency is acceptable, you can schedule the job towards the end of a frame and use its results in the following frame. Alternatively, if your application saturates that changeover period with other jobs, and there’s an under-utilized period somewhere else in the frame, it’s more efficient to schedule your job there instead.

You can also use the [Profiler](Profiler.md) to see where Unity is waiting for jobs to complete. The marker `WaitForJobGroupID` on the main thread indicates this. This marker might mean that you’ve introduced a data dependency somewhere that you should resolve. Look for `JobHandle.Complete` to track down where you have data dependencies that are forcing the main thread to wait.

### Avoid using long running jobs

Unlike threads, jobs don’t yield execution. Once a job starts, that job worker thread commits to completing the job before running any other job. As such, it’s best practice to break up long running jobs into smaller jobs that [depend on one another](job-system-job-dependencies.md), instead of submitting jobs that take a long time to complete relative to other jobs in the system.

The job system usually runs multiple chains of job dependencies, so if you break up long running tasks into multiple pieces there is a chance for multiple job chains to progress. If instead the job system is filled with long running jobs, they might completely consume all worker threads and block independent jobs from executing. This might push out the completion time of important jobs that the main thread explicitly waits for, resulting in stalls on the main thread that otherwise wouldn’t exist.

In particular, long running [`IJobParallelFor`](../ScriptReference/Unity.Jobs.IJobParallelFor.md) jobs impact negatively on the job system because these job types intentionally try to run on as many worker threads as possible for the job batch size. If you can’t break up long parallel jobs, consider increasing the batch size of your job when scheduling it to limit how many workers pick up the long running job.

```
MyParallelJob jobData = new MyParallelJob();
jobData.Data = someData;  
jobData.Result = someArray;  

// Use half the available worker threads, clamped to a minimum of 1 worker thread
const int numBatches = Math.Max(1, JobsUtility.JobWorkerCount / 2); 
const int totalItems = someArray.Length;
const int batchSize = totalItems / numBatches;

// Schedule the job with one Execute per index in the results array and batchSize items per processing batch
JobHandle handle = jobData.Schedule(result.Length, totalItems, batchSize);
```

## Additional resources

* [Job dependencies](job-system-job-dependencies.md)
* [Thread safe types](job-system-native-container.md)