# Configure the asynchronous upload pipeline

Unity provides optimization settings specific to the asynchronous upload pipeline. To access these settings, open the Project’s [Quality settings](class-QualitySettings.md).

Note that you cannot configure settings for the synchronous upload pipeline.

![The Async Upload settings in the Quality settings](../uploads/Main/AsyncUploadQualitySettings.png)

The Async Upload settings in the Quality settings

### Configure the streaming buffer for mesh data

Unity re-uses a single ring buffer to stream texture and **mesh** data to the GPU. This reduces the number of memory allocations required.

The **Async Upload Buffer** determines the size of this ring buffer. It has a minimum size of 2 megabytes, and a maximum size of 2047 megabytes.

Unity automatically resizes the buffer to fit the largest texture or mesh that is currently loading. This can be a slow operation, especially if Unity has to perform it more than once; for example, if you are loading many textures that are larger than the default buffer size. To reduce the number of times that Unity must resize the buffer, set this value to fit the largest value that you expect to load. This is usually the largest texture in the **Scene**.

You can configure this value in the Quality settings window, or using the [QualitySettings.asyncUploadBufferSize](../ScriptReference/QualitySettings-asyncUploadBufferSize.md) API.

### Configure the CPU’s upload time

To control the amount of time the CPU spends uploading texture or mesh data to the GPU, use **Async Upload Time Slice**. The value is in milliseconds per frame.

A larger value means the data will be ready on the GPU sooner, but the CPU will spend more time on upload operations during those frames. Note that Unity only uses this time for uploading if there is data waiting in the buffer to be uploaded to the GPU; if there is no data waiting, Unity can use this time for other operations.

You can configure this value in the Quality settings window, or using the [QualitySettings.asyncUploadTimeSlice](../ScriptReference/QualitySettings-asyncUploadTimeSlice.md) API.