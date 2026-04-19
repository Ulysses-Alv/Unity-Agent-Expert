# Check whether a mesh uses the asynchronous upload pipeline

To identify which upload pipeline Unity is using for a **mesh**, you can use the [Profiler](Profiler.md) or another profiling tool and observe thread activity and **profiler markers**.

The following indicate that Unity is using the asynchronous upload pipeline to upload textures or meshes:

* The `AsyncUploadManager.ScheduleAsyncRead`, `AsyncReadManager.ReadFile`, and `Async.DirectTextureLoadBegin` profiler markers.
* Activity on the `AsyncRead` thread.

If you do not see this activity, then Unity is not using the asynchronous upload pipeline.

Note that the following profiler markers do not indicate that Unity is using the asynchronous upload pipeline; Unity calls them to check whether any asynchronous upload work needs to occur:

* `Initialization.AsyncUploadTimeSlicedUpdate`
* `AsyncUploadManager.AsyncResourceUpload`
* `AsyncUploadManager.ScheduleAsyncCommands`