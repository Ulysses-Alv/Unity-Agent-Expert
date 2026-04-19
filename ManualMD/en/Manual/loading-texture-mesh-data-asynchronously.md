# Loading texture and mesh data asynchronously

Strategies and resources for uploading texture and **mesh** data to the GPU asynchronously, to improve performance.

| **Topic** | **Description** |
| --- | --- |
| [Uploading texture and mesh data to the GPU](LoadingTextureandMeshData.md) | Understand the difference between the synchronous and asynchronous upload pipelines, and what makes a mesh eligible for asynchronous upload. |
| [Check whether a mesh uses the asynchronous upload pipeline](identify-mesh-upload-pipeline.md) | Identify whether Unity is using the synchronous or asynchronous upload pipeline for a specific mesh. |
| [Configure the asynchronous upload pipeline](configure-asynchronous-upload-pipeline.md) | Configure the streaming buffer and upload time Unity allocates for asynchronously uploading mesh data to the GPU. |