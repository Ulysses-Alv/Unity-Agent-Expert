# Root Outputs

A root output lets you implement custom audio systems or sound generators, and integrate third party middleware (FMOD, Wise) that write directly to the main output, bypassing the Unity’s built-in mixer entirely.

## Setting up a root output

You can set up root outputs from code using the following steps:

1. **Implement the interfaces**: Provide implementations of `RootOutputInstance.IRealtime` and `RootOutputInstance.IControl`.
2. **Allocate the root output**: Call `ControlContext.AllocateRootOutput`, passing your real-time and control implementations.
3. **Destroy when finished**: When the root output is no longer needed, call `ControlContext.Destroy` to release it.

When you use `ControlContext.AllocateRootOutput` with the built-in control context, it connects the new root output to Unity’s main audio output and uses the real-time part of your instance to generate audio.

## Root output processing stages

Implementations of `RootOutputInstance.IRealtime` processes audio in three stages:

1. **Early processing**: Perform work that must complete before later stages and before you use any resources managed by the root output. For example, a hardware input might sample its data here once so that later stages read a stable snapshot.
2. **Processing**: Schedule the main body of work. This stage runs in parallel with other root outputs. If you use jobs, you are responsible for managing dependencies and ensuring they complete in order. When consuming other processors (for example, generators), either schedule your jobs with the provided input dependency or complete that work first.
3. **End processing**: Mix the final results into the main audio output.

At the end of early processing, each root output might return an optional `JobHandle`. Before the processing stage starts, Unity combines all returned handles into a single dependency and passes this dependency to every root output during processing. If you schedule jobs in the processing stage, use this dependency to make sure any work started in early processing completes first.

## Threading & real-time constraints

For each invocation of a root output, the early processing, processing, and end processing stages all execute sequentially on the same thread. Depending on the control context, this thread might be the audio thread (for the built-in context) or a job worker thread (for manual contexts).

The following real-time constraints apply to all three stages:

* No memory allocations neither managed nor native.
* No system calls (such as file access, networking, logging to disk/console, or sleeping).
* No locks or blocking operations such as mutexes, waiting on events/semaphores, or blocking of file access.

## Best practices

* Keep the processing stages real‑time safe. Avoid allocations, locks, blocking I/O, logging, and throwing exceptions on the audio thread. Prefer stack-allocated or pooled buffers and struct-based state.
* Beware of threading. Don’t call `UnityEngine` APIs from within `Process`. Use thread-safe and lock-free patterns for synchronizing shared state.
* Use Burst-compiled, struct-based code with simple value types and tight loops. Choose `Unity.Mathematics` and lay out data for Burst auto-vectorization.
* Return a valid `JobHandle` from the early stage if you schedule work there.
* Chain your processing-stage jobs to the combined dependency so the audio pipeline doesn’t stall waiting for them to finish.

## Additional resources

* [Scriptable processors concepts](audio-scriptable-processors-concepts.md)
* [Example: Creating a simple root output](audio-scriptable-processors-example-creating-a-root-output.md)