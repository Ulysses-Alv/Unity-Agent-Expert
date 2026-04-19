# Jobs overview

A job is a small unit of work that does one specific task. A job receives parameters and operates on data, similar to how a method call behaves. Jobs can be self-contained, or they can depend on other jobs to complete before they can run. In Unity, a job refers to any struct that implements [the `IJob` interface.](../ScriptReference/Unity.Jobs.IJob.md)

Only the main thread can schedule and complete jobs. It can’t access the content of any running jobs, and two jobs can’t access the contents of a job at the same time. To ensure efficient running of jobs, you can make them dependent on each other. Unity’s job system allows you to create complex dependency chains to ensure that your jobs complete in the correct order.

## Job types

* [IJob](../ScriptReference/Unity.Jobs.IJob.md): Runs a single task on a job thread.
* [IJobParallelFor](../ScriptReference/Unity.Jobs.IJobParallelFor.md): Runs a task in parallel. Each worker thread that runs in parallel has an exclusive index to access shared data between worker threads safely.
* [IJobParallelForTransform](../ScriptReference/Jobs.IJobParallelForTransform.md): Runs a task in parallel. Each worker thread running in parallel has an exclusive Transform from the transform hierarchy to operate on.
* [IJobFor](../ScriptReference/Unity.Jobs.IJobFor.md): The same as `IJobParallelFor`, but allows you to schedule the job so that it doesn’t run in parallel.

## Additional resources

* [Create a job](job-system-creating-jobs.md)
* [Job dependencies](job-system-job-dependencies.md)
* [Parallel jobs](job-system-parallel-for-jobs.md)