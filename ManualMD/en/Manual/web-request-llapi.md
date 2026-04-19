# Web request low-level API reference

While the [high-level web request API](web-request-hlapi.md) minimizes the need for boilerplate code, the low-level web request API offers more flexibility and control.

This section lists key members of the low-level API and the scenarios they’re intended for. In general, using the low-level API involves creating a [`UnityWebRequest`](../ScriptReference/Networking.UnityWebRequest.md) to which you attach specialized [`DownloadHandler`](../ScriptReference/Networking.DownloadHandler.md) or [`UploadHandler`](../ScriptReference/Networking.UploadHandlert.md) objects for particular download and upload cases.

The high and low-level APIs aren’t mutually exclusive. You can always customize `UnityWebRequest` objects created via the high-level API if you need more control over a common operation.

## Instantiating UnityWebRequest objects

`UnityWebRequest` can be instantiated through its [constructors](../ScriptReference/Networking.UnityWebRequest-ctor.md) like any other object. Several constructor signatures are available for creating a `UnityWebRequest` with or without upload or download handlers. `UnityWebRequest` also has properties for configuring the request, tracking its status, and checking the result.

For more information, refer to the [`UnityWebRequest`](../ScriptReference/Networking.UnityWebRequest.md) API reference.

## Download handlers

Specialist download handlers allow you to handle specific download cases and control the processing of incoming data. The following download handlers are available, which all specialize the base [`DownloadHandler`](../ScriptReference/Networking.DownloadHandler.md) class:

| **Handler type** | **Description** |
| --- | --- |
| [`DownloadHandlerBuffer`](../ScriptReference/Networking.DownloadHandlerBuffer.md) | Used for simple data storage. This download handler is the simplest, and handles the majority of use cases. It stores received data in a native code buffer. When the download is complete, you can access the buffered data either as an array of bytes or as a text string. |
| [`DownloadHandlerFile`](../ScriptReference/Networking.DownloadHandlerFile.md) | Specialized download handler for downloading and saving large files to disk with low memory footprint. It writes downloaded bytes directly to file, so the memory usage is low regardless of the size of the file being downloaded. The difference from other download handlers is that you retrieve data from the saved file rather than from the handler itself. |
| [`DownloadHandlerTexture`](../ScriptReference/Networking.DownloadHandlerTexture.md) | Specialized download handler optimized for downloading images for use as Unity [Texture](../ScriptReference/Texture.md) objects. This is a more efficient alternative to using a `DownloadHandlerBuffer` to download an image file and then creating a texture from the raw bytes using `Texture.LoadImage`. |
| [`DownloadHandlerAssetBundle`](../ScriptReference/Networking.DownloadHandlerAssetBundle.md) | Specialized download handler optimized for downloading Unity AssetBundles. It can stream data directly to Unity’s asset system and produce an [`AssetBundle`](../ScriptReference/AssetBundle.md) object as a result. |
| [`DownloadHandlerAudioClip`](../ScriptReference/Networking.DownloadHandlerAudioClip.md) | Specialized download handler optimized for downloading audio files. This is a more efficient alternative to downloading raw bytes using `DownloadHandlerBuffer` and then creating an [`AudioClip`](../ScriptReference/AudioClip.md) from them. |
| [`DownloadHandlerScript`](../ScriptReference/Networking.DownloadHandlerScript.md) | An abstract base class for custom, scriptable `DownloadHandler` implementations. Inherit from this class to receive callbacks from the `UnityWebRequest` system, which you can use to perform custom handling of data as it arrives from the network. |

`UnityWebRequest` has a [`UnityWebRequest.disposeDownloadHandlerOnDispose`](../ScriptReference/Networking.UnityWebRequest-disposeDownloadHandlerOnDispose.md) property, which defaults to true. If this property is true, then calls to [`UnityWebRequest.Dispose`](../ScriptReference/Networking.UnityWebRequest.Dispose.md) also invoke [`DownloadHandler.Dispose`](../ScriptReference/Networking.DownloadHandler.Dispose.md) on the attached download handler. If you keep a reference to the download handler longer than the reference to the associated `UnityWebRequest`, set `disposeDownloadHandlerOnDispose` to false.

## Upload handlers

Upload handlers allow you more detailed control of aspects of the upload process. The following upload handlers are available, which all specialize the base `UploadHandler` class:

| **Handler type** | **Description** |
| --- | --- |
| [`UploadHandlerRaw`](../ScriptReference/Networking.UploadHandlerRaw.md) | Accepts a data buffer at construction time. When this buffer is an array of bytes, it’s copied internally into native code memory. `UnityWebRequest` system uses this buffer when the remote server is ready to receive the request body data. When the buffer is provided as a `NativeArray`, no copying is performed. |
| [`UploadHandlerFile`](../ScriptReference/Networking.UploadHandlerFile.md) | Allows you to send the contents of a file as the request body. Using this handler, you can send a large file to a server without using a lot of memory. Because the handler continually reads data from the file and then sends it, only a small fraction of the file is kept in memory at any given time. |

The value of the [`UploadHandler.contentType`](../ScriptReference/Networking.UploadHandler-contentType.md) property is used for the value of the web request’s `Content-Type` header if you don’t set one directly on the `UnityWebRequest` itself. If you do set a `Content-Type` header on the `UnityWebRequest` object, the `Content-Type` on the upload handler object is ignored. If you don’t set a `Content-Type` on either the `UnityWebRequest` or the `UploadHandler`, the system defaults to setting a `Content-Type` of `application/octet-stream`.

`UnityWebRequest` has a [`disposeUploadHandlerOnDispose`](../ScriptReference/Networking.UnityWebRequest-disposeUploadHandlerOnDispose.md), which defaults to true. If this property is true, then calls to [`UnityWebRequest.Dispose`](../ScriptReference/Networking.UnityWebRequest.Dispose.md) also invoke [`UploadHandler.Dispose`](../ScriptReference/Networking.UploadHandler.Dispose.md) on the attached upload handler. If you keep a reference to upload handler longer than the reference to `UnityWebRequest`, set `disposeUploadHandlerOnDispose` to false.

## Additional resources

* [Introduction to the Unity web request APIs](web-request-intro.md)
* [Web request high-level API reference](web-request-hlapi.md)