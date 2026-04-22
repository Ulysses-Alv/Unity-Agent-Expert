# Web request high-level API reference

You can use the high-level API to send simple web requests to backends with common HTTP verbs such as `GET`, `POST`, `PUT`. The following table lists some of the common operations supported by the high-level API. Refer to the linked API documentation for details and example usage for each operation:

| **Method** | **Description** |
| --- | --- |
| [`UnityWebRequest.Get`](../ScriptReference/Networking.UnityWebRequest.Get.md) | Retrieve text or binary data from an HTTP Server (`GET`) |
| [`UnityWebRequest.Post`](../ScriptReference/Networking.UnityWebRequest.Post.md) | Send a form to an HTTP server (`POST`). Overloads are available that support form data submission in both the legacy [`WWWForm`](../ScriptReference/WWWForm.md) and the recommended [`IMultipartFormSection`](../ScriptReference/Networking.IMultipartFormSection.md) formats. |
| [`UnityWebRequest.Put`](../ScriptReference/Networking.UnityWebRequest.Put.md) | Upload raw data to an HTTP server (`PUT`) |
| [`UnityWebRequestTexture.GetTexture`](../ScriptReference/Networking.UnityWebRequestTexture.GetTexture.md) | Use a specialized type of web request optimized for downloading Unity [`Texture`](../ScriptReference/Texture.md) images from an HTTP server (`GET`) |
| [`UnityWebRequestAssetBundle.GetAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.GetAssetBundle.md) | Use a specialized type of web request optimized for downloading Unity [`AssetBundle`](../ScriptReference/AssetBundle.md) objects from an HTTP server (`GET`). |

These high-level requests automatically attach the appropriate upload or download handlers to the underlying `UnityWebRequest`. These handlers provide the opportunity for customizing data before sending it and processing data after receiving it. For more information on upload and download handlers, refer to the [low-level API reference](web-request-llapi.md).

## Additional resources

* [Web request low-level API reference](web-request-llapi.md)
* [`UnityWebRequest` class](../ScriptReference/Networking.UnityWebRequest.md)