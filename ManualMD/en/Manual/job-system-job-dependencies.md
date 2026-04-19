# Job dependencies

Often, one job depends on the results of another job. For example, Job A might write to a `NativeArray` that job B uses as input. You must tell the job system about such a dependency when you schedule a dependent job. The job system won’t run the dependent job until the job it depends upon is finished. One job can depend on more than one job.

You can also have a chain of jobs in which each job depends on the previous one. However, dependencies delay job execution because you must wait for any dependencies of a job to complete before it can run. Completing a dependent job must first complete any job it depends on, and any jobs those jobs depend on.

When you call the [`Schedule`](../ScriptReference/Unity.Jobs.IJobExtensions.Schedule.md) method of a job it returns a [`JobHandle`](../ScriptReference/Unity.Jobs.JobHandle.md). You can use a `JobHandle` as a dependency for other jobs. If a job depends on the results of another job, you can pass the first job’s `JobHandle` as a parameter to the second job’s `Schedule` method, like so:

```
JobHandle firstJobHandle = firstJob.Schedule();
secondJob.Schedule(firstJobHandle);
```

## Combining dependencies

If a job has a lot of dependencies, you can use the method [`JobHandle.CombineDependencies`](../ScriptReference/Unity.Jobs.JobHandle.CombineDependencies.md) to merge them. `CombineDependencies` allows you to pass dependencies onto the `Schedule` method.

```
NativeArray<JobHandle> handles = new NativeArray<JobHandle>(numJobs, Allocator.TempJob);

// Populate `handles` with `JobHandles` from multiple scheduled jobs...

JobHandle jh = JobHandle.CombineDependencies(handles);
```

## An example of multiple jobs and dependencies

The following is an example of multiple jobs that have multiple dependencies. It’s best practice to put the job code (`MyJob` and `AddOneJob`) in a separate file to the `Update` and `LateUpdate` code, but for the purposes of clarity, this example is one file:

```
using UnityEngine;
using Unity.Collections;
using Unity.Jobs;

public class MyDependentJob : MonoBehaviour
{
    // Create a native array of a single float to store the result. This example waits for the job to complete.
    NativeArray<float> result;
    // Create a JobHandle to access the results
    JobHandle secondHandle;

    // Set up the first job
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

    // Set up the second job, which adds one to a value
    public struct AddOneJob : IJob
    {
        public NativeArray<float> result;

        public void Execute()
        {
            result[0] = result[0] + 1;
        }
    }

    // Update is called once per frame
    void Update()
    {
        // Set up the job data for the first job
        result = new NativeArray<float>(1, Allocator.TempJob);

        MyJob jobData = new MyJob
        {
            a = 10,
            b = 10,
            result = result
        };

        // Schedule the first job
        JobHandle firstHandle = jobData.Schedule();

        // Setup the data for the second job
        AddOneJob incJobData = new AddOneJob
        {
            result = result
        };

        // Schedule the second job
        secondHandle = incJobData.Schedule(firstHandle);
    }

    private void LateUpdate()
    {
        // Sometime later in the frame, wait for the job to complete before accessing the results.
        secondHandle.Complete();

        // All copies of the NativeArray point to the same memory, you can access the result in "your" copy of the NativeArray
        // float aPlusBPlusOne = result[0];

        // Free the memory allocated by the result array
        result.Dispose();
    }

}
```

## Additional resources

* [Parallel jobs](job-system-parallel-for-jobs.md)