# Web browser access to device features

The Unity Web platform supports [WebCam access](../ScriptReference/WebCamDevice.md). To allow a Web application to access the webcam on a device, the browser must request its user to provide access to the **camera**. Without the permission to access the camera, the browser returns incomplete or inaccurate information.

**Note:** Currently, the Web platform supports WebCam devices only.

To request browser permission to access the webcam, use the `Application.RequestUserAuthorization` API:

```
using UnityEngine;
using UnityEngine.iOS;
using System.Collections;

// Get WebCam information from the browser
public class ExampleClass : MonoBehaviour
{
    private WebCamDevice[] devices;
    
    // Use this for initialization
    IEnumerator Start()
    {
        yield return Application.RequestUserAuthorization(UserAuthorization.WebCam);
        if (Application.HasUserAuthorization(UserAuthorization.WebCam))
        {
            Debug.Log("webcam found");
            devices = WebCamTexture.devices;
            for (int cameraIndex = 0; cameraIndex < devices.Length; ++cameraIndex)
            {
                Debug.Log("devices[cameraIndex].name: ");
                Debug.Log(devices[cameraIndex].name);
                Debug.Log("devices[cameraIndex].isFrontFacing");
                Debug.Log(devices[cameraIndex].isFrontFacing);
            }
        }
        else
        {
            Debug.Log("no webcams found");
        }
    }
}
```

**Note:** Unity recommends to use the `MediaDevices.getUserMedia()` API to request user permission for accessing the device. This feature is available only in secure contexts (HTTPS).