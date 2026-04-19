# Profiler markers reference

Unity’s code is instrumented with [profiler markers](profiler-adding-information-code-intro.md#profiler-markers) that give you insight into what takes up time in your application. The following tables contain a list of some of the in-built markers.

* [Main thread base markers](#main-thread)
* [Script update markers](#script-update)
* [Rendering and VSync markers](#rendering)
* [Back end scripting markers](#backend)
* [Multithreading markers](#multithreading)
* [Physics markers](#physics)
* [Animation markers](#animation)
* [UI Toolkit markers](UIE-profiler-markers.md)
* [Performance warnings](#warnings)

## Main thread base markers

The main thread base markers provide a clear separation between the time spent on your application and time spent on Unity Editor and Profiler activities. You can also use these markers with the [`ProfilerRecorder`](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) API to get the timing of a frame on the main thread, or the [`FrameTimingManager`](../ScriptReference/FrameTimingManager.md) API for high level frame stats at runtime.

You can also use the rendering **profiler counters** CPU Main Thread Frame Time, CPU Render Thread Frame Time, and CPU Total Frame Time to get high-level timings of the CPU usage of your application. For more information, refer to [Rendering profiler counters reference](profiler-counters-reference.md#rendering).

| **Marker** | **Description** |
| --- | --- |
| **PlayerLoop** | Contains any samples that originate from your application’s main loop. If you [target the Editor](profiling-edit-mode.md) instead of Play mode while the Player is running within the Editor in active Play mode, PlayerLoop samples nest under the EditorLoop. |
| **EditorLoop** *(Editor only marker)* | Contains any samples that originate from the Editor’s main loop. This is only present while you [profile a player in the Editor](profiling-edit-mode.md). When you [target Play mode with the Profiler](profiling-play-mode.md), EditorLoop samples indicate the amount of time spent rendering and running the Editor that contains the Player.  For more information, refer to [Play mode and Edit mode samples](profiler-play-edit-samples.md). |
| **Profiler.CollectEditorStats** *(Editor only marker)* | Contains any samples that relate to collecting statistics for different active Profiler modules.  Samples nested under the `Profiler.CollectGlobalStats` marker indicate how much overhead the Player has when it collects the statistics of a particular module. All other child samples only reflect their effect in the Editor.  To remove the overhead that a particular module has, close the module’s chart, or call `Profiler.SetAreaEnabled`.  **Note:** [Custom Profiler modules](profiler-customizing.md) that use built-in counters enable the built-in counter’s area, even if the module it belongs is disabled. To prevent the Profiler from collecting these statistics and creating collection overhead, make sure that both the built-in Profiler module and your custom Profiler module are disabled. |

## Script update markers

Unless you’re using the [job system](job-system.md), most of your scripting code is nested underneath the following markers. For information on job system samples, refer to the [Multi threading markers](#multithreading) section of this page.

For further details on Unity’s Update Loop, refer to the documentation on [Order of execution of event functions](execution-order.md). You can insert your own callbacks into the [Player Loop](../ScriptReference/LowLevel.PlayerLoop.md) using [`PlayerLoop.SetPlayerLoop`](../ScriptReference/LowLevel.PlayerLoop.SetPlayerLoop.md). If you insert callbacks directly into the main loop, your script update samples appear on their own. If you insert callbacks as a subsystem, the samples appear under the respective main PlayerLoopSystem marker.

| **Marker** | **Description** |
| --- | --- |
| **BehaviourUpdate** | Contains all samples of [`MonoBehaviour.Update`](../ScriptReference/MonoBehaviour.Update.md) methods. |
| **CoroutinesDelayedCalls** | Contains all samples of [coroutines](../ScriptReference/Coroutine..md) after their first yield. |
| **FixedBehaviourUpdate** | Contains all samples of [`Monobehaviour.FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md) methods. |
| **PreLateUpdate.ScriptRunBehaviourLateUpdate** | Contains all samples of [`Monobehaviour.LateUpdate`](../ScriptReference/MonoBehaviour.LateUpdate.md) methods. |
| **Update.ScriptRunBehaviourUpdate** | Contains all samples of [`MonoBehaviour.Update`](../ScriptReference/MonoBehaviour.Update.md) and [coroutines](../ScriptReference/Coroutine.md). |

## Rendering and VSync markers

These markers contain the samples where the CPU spends time processing data for the GPU, or where it might be waiting for the GPU to finish. If the [GPU Profiler module](ProfilerGPU.md) isn’t available, or it adds too much overhead, the **toolbar** in the Profiler module details pane doesn’t display this information. The samples under these markers can give you a good idea of whether your application is CPU-bound or GPU-bound.

| **Marker** | **Description** |
| --- | --- |
| **WaitForTargetFPS** | Indicates how much time your application spent waiting for the targeted FPS that [`Application.targetFrameRate`](../ScriptReference/Application-targetFrameRate.md) specifies.  If the sample is a subsample of Gfx.WaitForPresentOnGfxThread, it represents the amount of time that your application spent waiting for the GPU. For example, this could be time that the GPU spent waiting for the next VSync, if that is configured in [QualitySettings.vSyncCount](../ScriptReference/QualitySettings-vSyncCount.md), or if vSync is enforced on your target platform. However, samples with this marker are also emitted if the GPU hasn’t finished computing the frame.  To determine what is causing samples with this marker to use a lot of time, switch to the [Timeline view in the CPU Profiler module](ProfilerCPU.md). In this view, you can check what happened on the render thread and how much time passed between this sample ending in the current frame and the same sample ending in surrounding frames.  If the duration is larger than your application’s frame time should be (based on the targeted frame rate or vSync) your frames are taking too long to render or compute. If that’s the case, investigate the render thread and see how much time it spent on Gfx.PresentFrame over other work it did to prepare and issue commands to the GPU. If the render thread spent a large amount of time in Gfx.PresentFrame, your rendering work is GPU-bound. If the render thread’s time was spent preparing commands, your application is CPU-bound.  To find out what to focus on, if your application is GPU bound, profile the GPU work with the Unity Profiler or a platform profiler. For more information, see the User manual documentation on how to [optimize graphics performance](OptimizingGraphicsPerformance.md).  **Note:** The Editor doesn’t VSync on the GPU and instead uses WaitForTargetFPS to simulate the delay for VSync. Some platforms, in particular Android and iOS, enforce VSync or have a default frame rate cap of 30 or 60. |
| **Gfx.PresentFrame** | Represents the time your application spent waiting for the GPU to render and present the frame, which includes waiting for VSync.  Samples with the WaitForTargetFPS marker on the main thread show how much time is spent waiting for VSync. |
| **Gfx.ProcessCommands** | Contains all processing of the rendering commands on the render thread. Your application might have spent some of this processing time waiting for **VSync** or new commands from the main thread, which you can see from its child sample Gfx.WaitForPresentOnGfxThread. |
| **Gfx.WaitForCommands** | Indicates that the render thread was ready for new commands. If you see this marker, it might indicate a bottleneck on the main thread. |
| `<GraphicsAPIName>.WaitForLastPresent` for example: **GfxDeviceD3D11.WaitForLastPresent** **GfxDeviceD3D12.WaitForLastPresent** **GfxDeviceMetal.WaitForLastPresent** | Samples with this marker appear when the main thread waited for the GPU to flip the frame number to the screen (`Time.frameCount - QualitySettings.maxQueuedFrames + 1`). This means that if QualitySettings.maxQueuedFrames is greater than one, this time is spent waiting for the GPU to flip a frame that your application requested to render in a previous main thread frame.  For more details on this sample and an overview of Unity’s Frame Pipeline, see [Unity’s blog post on fixing Delta Time](https://unity.com/blog/engine-platform/fixing-time-deltatime-in-unity-2020-2-for-smoother-gameplay). |
| **Gfx.WaitForPresentOnGfxThread** | Indicates that the main thread was ready to start rendering the next frame, but the render thread didn’t finish waiting for the GPU to present the frame. This might indicate that your application is GPU-bound. To check what the render thread is simultaneously spending time on, check the CPU Profiler module’s [Timeline view](profiler-cpu-introduction.md#timeline-view).   If the render thread spends time in Camera.Render, your application is CPU-bound and might be spending too much time sending draw calls or textures to the GPU.  If the render thread spends time in `Gfx.PresentFrame`, your application is GPU-bound, or it might be waiting for VSync on the GPU. A `WaitForTargetFPS` sub-sample of `Gfx.WaitForPresentOnGfxThread` represents the portion of the Present phase that your application spent waiting for VSync. The Present phase is the portion of time between Unity instructing the graphics API to swap the buffers, to the time that this operation completed. |
| **Gfx.WaitForRenderThread** | Indicates that the main thread was waiting for the render thread to process all the commands in its command stream. Samples with this marker only appear in multithreaded rendering. |

## Back end scripting markers

The samples highlight Mono or **IL2CPP** **scripting backend** activities and are useful for troubleshooting issues with garbage collection and allocation.

| **Marker** | **Description** |
| --- | --- |
| **GC.Alloc** | Represents an allocation in the managed heap that contains managed allocations that are subject to automatic garbage collection. To reduce the time your application spends on automatic garbage collection, you should minimize these types of samples. |
| **GC.Collect** | Represents samples that relate to garbage collection. Whenever Unity needs to perform garbage collection, it stops running your program code and only resumes normal execution when the garbage collector has finished all its work. **Note:** If you have enabled [Incremental Garbage Collection](performance-incremental-garbage-collection.md) the garbage collector might not finish its work in a single frame.  This interruption might cause delays in the execution of your application that last anywhere from less than one millisecond to hundreds of milliseconds. This depends on how much memory the garbage collector needs to process and the platform your application is running on. For more information, see the documentation on [Understanding automatic memory management](performance-managed-memory.md). |
| **Mono.JIT** *Mono-only* | Contains samples that relate to just-in-time compilation of a scripting method. When a function is executed for the first time, Mono compiles it and Mono.JIT represents this compilation overhead. |
| **UnsafeUtility.Malloc** | Contains samples that call UnsafeUtility.Malloc to allocate unmanaged memory. While the Garbage Collector does not track this memory, allocating memory might have a significant performance impact which is shown with this sample. To investigate the source of this call, you can enable [Call Stack](ProfilerCPU.md) recording for this marker in the Profiler window. |

## Multithreading markers

These markers contain samples that don’t measure the CPU cycles consumed, but instead highlight information that relates to thread synchronization and the job system. When you see these samples, use the CPU Profiler module’s [Timeline view](ProfilerCPU.md) to check what’s happening on other threads at the same time.

| **Marker** | **Description** |
| --- | --- |
| **Idle** | Contains samples that indicate the length of time that a Worker Thread is inactive for. A worker thread is inactive any time that the job system doesn’t use it, and it goes into wait mode, where it waits on the semaphore.  Small gaps between Idle samples usually happen when the Job System wakes them up, for example, to schedule new jobs. Longer gaps might indicate that a native job that hasn’t been instrumented is running on the thread. |
| **JobHandle.Complete** | Contains samples that indicate when a sync point on a job happened. Sync points might have a performance impact on your application and might interfere with the execution of multi-threaded job code. To make it easier to find where exactly the sync point happened, enable [Call Stack](ProfilerCPU.md) recording for this sample. In the CPU Profiler module’s **Timeline** view you can enable [Flow Events](ProfilerCPU.md) to identify which jobs finished at this point. |
| **Semaphore.WaitForSignal** | Contains a sample that depicts a synchronization point on a thread. To find the thread it’s waiting for, check the Timeline view for samples that ended shortly before this one. |
| **WaitForJobGroupID** | A Sync Fence on a JobHandle was triggered. This might lead to [work stealing](../ScriptReference/Unity.Jobs.IJobParallelFor.md), which happens when a worker finishes its work and then looks at other workers’ jobs to complete. These display as job samples executed under this marker. Jobs that were stolen aren’t necessarily the jobs that were being waited for. |

## Physics markers

The following table outlines some high-level physics Profiler markers. [`FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md) calls all these measurements.

| **Marker** | **Description** |
| --- | --- |
| **Physics.FetchResults** | Contains samples that collect the results of the physics simulation from the **physics engine**, such as contact streams, trigger overlaps, and **joint** breakage events. |
| **Physics.Interpolation** | Contains samples that measure the execution time of the `Physics.Interpolation` method. This method manages the interpolation of positions and rotations for all the physics objects in your application. |
| **Physics.Processing** | Contains samples that spent time waiting on the main thread until the physics simulation completed across all threads. If your application spends a lot of time in `Physics.Processing` but only has a few physics related **GameObjects** in the **Scene**, it might indicate that worker threads picked up other systems tasks due to job stealing and reported as physics. This is because while waiting, the main thread picks up jobs from the high priority queue. |
| **Physics.ProcessingCloth** | Contains samples that measure the execution time of the `Physics.ProcessingCloth` method. This method processes all cloth physics jobs. Expand this sample to display the low-level detail of the work done internally in the physics system. |
| **Physics.ProcessReports** | Contains samples that correspond to time spent forwarding physics data to scripts via callbacks such as [`OnTriggerEnter`](../ScriptReference/MonoBehaviour.OnTriggerEnter.md). **Note:** These samples don’t compute the data required because they have already been prepared during `FetchResults`.  There are four distinct stages:   * **Physics.Contacts**: Contains samples that measure the execution time of `Physics.Contacts`. This processes [`OnCollisionEnter`](../ScriptReference/MonoBehaviour.OnCollisionEnter.md), [`OnCollisionExit`](../ScriptReference/MonoBehaviour.OnCollisionExit.md), and [`OnCollisionStay`](../ScriptReference/MonoBehaviour.OnCollisionStay.md) events. * **Physics.JointBreaks**: Contains samples that measure the execution time of `Physics.JointBreaks`. This processes updates and messages related to broken joints. * **Physics.TriggerEnterExits**: Contains samples that measure the execution time of `Physics.TriggerEnterExits`. This processes [`OnTriggerEnter`](../ScriptReference/MonoBehaviour.OnTriggerEnter.md) and [`OnTriggerExit`](../ScriptReference/MonoBehaviour.OnTriggerExit.md) events. * **Physics.TriggerStays**: Contains samples that measure the execution time of `Physics.TriggerStays`. This processes [`OnTriggerStay`](../ScriptReference/MonoBehaviour.OnTriggerStay.md) events. |
| **Physics.Simulate** | Contains samples that measure the amount of time spent working on the pre-requisites for the [`Physics.Simulate`](../ScriptReference/Physics.Simulate.md) method. This method instructs the physics system to run its simulation, which updates the state of the current physics. |
| **Physics.UpdateBodies** | Contains samples that update all the physics bodies’ positions and rotations. For each GameObject that has a **Rigidbody** component, samples with this marker read the pose from the physics system and write it to the Transform. |
| **Physics.UpdateCloth** | Contains samples that measure the execution time of the `Physics.UpdateCloth` method. This method processes updates that relate to cloth and their Skinned Meshes. |

For more information about script lifecycles and general samples within a script lifecycle, refer to [Order of execution for event functions](execution-order.md#UpdateOrder).

## Animation markers

The following tables contain markers from the [Animation system](class-Animator.md). For general information on Animation system performance, refer to the [Mecanim Performance and and Optimization](MecanimPeformanceandOptimization.md) page.

### PreLateUpdate.DirectorUpdateAnimationBegin stage

In this stage, every Animator that’s active and enabled is processed. Active Animators are processed regardless of Culling Mode and visibility.

Markers that start with `Director.` can also cover **Playables** and Timeline.

| **Marker** | **Description** |
| --- | --- |
| **Director.PrepareFrame** | Set up, schedule, and await `Director.PrepareFrameJob` jobs. These jobs evaluate the **state machines** for all active **Animator components**. |
| **Director.PrepareFrameJob *(job)*** | Evaluate the state machine of a single active Animator. The amount of time taken to process an Animator grows with the complexity of its state machine (i.e. the number of states, transitions, properties, and layers).  Evaluation of state machines with [`StateMachineBehaviours`](StateMachineBehaviours.md) that implement [`OnStateMachineEnter`](../ScriptReference/StateMachineBehaviour.OnStateMachineEnter.md) or [`OnStateMachineExit`](../ScriptReference/StateMachineBehaviour.OnStateMachineExit.md) listeners will be constrained to the main thread. |
| **Animators.PrepareFirstPass** | Prepare for the next processing steps. |
| **Animators.ProcessGraphJob** | Build, schedule, and await the results of `Animator.ProcessGraph` jobs. |
| **Animators.ProcessGraph *(job)*** | Evaluate all properties across all connected AnimationClips.  Calculate the root motion displacements by blending the root motion properties from all clips together.  Collect new animation events for the current [`deltaTime`](../ScriptReference/Time-deltaTime.md).  When evaluating AnimationClips’ properties, the curve segments are cached locally between frames to avoid reading from the AnimationClips’ memory more than necessary. |
| **Animators.FireAnimationEventsAndBehaviours** | Fire all **animation events** messages on [`StateMachineBehaviours`](StateMachineBehaviours.md) and MonoBehaviours, except [`OnStateMachineEnter`](../ScriptReference/StateMachineBehaviour.OnStateMachineEnter.md) or [`OnStateMachineExit`](../ScriptReference/StateMachineBehaviour.OnStateMachineExit.md) listeners which have already been fired in `Director.PrepareFrameJob`. |
| **Animators.ApplyOnAnimatorMove** | Apply **root motion** and trigger the [`OnAnimatorMove`](../ScriptReference/MonoBehaviour.OnAnimatorMove.md) message on MonoBehaviours. |

### PreLateUpdate.DirectorUpdateAnimationEnd

In this stage, only Animators with the Culling Mode `Always Update` and visible Animators with the Culling Mode `UpdateTransform` or `Cull Completely` get processed. Animators with `Cull Completely` that were moved out of frame by the Root Motion don’t participate in this phase. Neither do Animators that were disabled by user code triggered by StateMachineBehaviours and AnimationEvents

| **Marker** | **Function** |
| --- | --- |
| **Animators.PrepareSecondPass** | Set up the next processing steps. This includes filtering out Animators based on their Culling Mode and visibility status. |
| **Animators.SortWriteJob** | The Transform system only allows a single thread to modify a Transform Hierarchy at a time. To accommodate this constraint, `Animators.SortWriteJob` groups `Animators.WriteTransforms` jobs by their [Transform.root](../ScriptReference/Transform-root.md).  For the best performance, avoid having multiple Animators in the same transform hierarchy to increase the granularity at which jobs can be scheduled and parallelized. |
| **Animators.ProcessAnimationsJob** | Build, schedule, and await the results of `Animators.ProcessAnimations` jobs. |
| **Animator.ProcessAnimations (*job*)** | Blend all properties on current active AnimationClips for a single Animator, except for root motion. Apply them to the internal avatar representation.  The length of this marker scales with animation and blending complexity. |
| **OnAnimatorIK** | Sets up animation IK. This is called once for each **Animator Controller** layer with IK pass enabled. |
| **Animators.WriteJob** | Build, schedule, and await the results of `Animator.WriteTransform` jobs. These jobs write transform poses from Animation **avatars** to the related GameObject transform hierarchy. |
| **Animators.WriteTransforms *(job)*** | Writes all animated transforms to the scene from worker threads. |
| **Animator.WriteProperties** | Write any animated property that is not a transform pose to the target object. |

## Performance warnings

The CPU Profiler detects some common performance issues and warns you about them. These appear in the Warning column of the [CPU Profiler module’s Hierarchy](profiler-cpu-introduction.md#hierarchy-views) view in the module details pane.

The Profiler detects some specific calls to avoid in performance-critical contexts. It displays the warnings with the reasons the operations are affecting performance as follows:

| **Warning** | **Description** |
| --- | --- |
| **Animation.DestroyAnimationClip** **Animation.AddClip** **Animation.RemoveClip** **Animation.Clone** **Animation.Deactivate** | Indicates that RebuildInternalState has been triggered. RebuildInternalState is an operation that goes through the list of curves for each clip in the Animation component, and then rebinds each curve to a value on a component, on a GameObject.  This is a resource-intensive operation, so you should avoid calling these methods at runtime as much as possible. |
| **AssetBundle.asset/allAssets** | Indicates that Unity called the AssetBundleRequest.assets/allAssets API while the AssetBundle loading was not complete (AssetBundleRequest.isDone is false). This causes a stall on the main thread and waits for the loading operation to complete. |
| **AsyncUploadManager.AsyncBufferResized** **AsyncUploadManager.AsyncBufferDelete** | Indicates that the internal buffer for uploading data to the GPU is resized because it’s not big enough. This resizing is slow and causes spikes in CPU activity.  You can avoid this warning if you can spare the memory to allocate a larger size up front.  You can use **Async Upload Buffer Size** setting in [Quality Settings](class-QualitySettings.md) to set the default size.  The `AsyncUploadManager.AsyncBufferResized` marker indicates the newly allocated size which you can use as a guide for the default buffer size. |
| **Rigidbody.SetKinematic** | Recreate non-convex MeshCollider for Rigidbody. |

## Additional resources

* [UI Toolkit markers](UIE-profiler-markers.md)
* [Play mode and Editor samples](profiler-play-edit-samples.md)
* [CPU Usage Profiler module](profiler-cpu.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)