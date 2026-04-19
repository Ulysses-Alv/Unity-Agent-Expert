# WinRT API in C# scripts for UWP

Unity includes Windows Runtime (WinRT) support when building for Universal Windows Platform with the [IL2CPP](scripting-backends-il2cpp.md) scripting backend. Use Windows Runtime support to call into both native system Windows Runtime APIs as well as custom .winmd files directly from managed (C#) **scripts** and plugins.

Unity automatically references Windows Runtime APIs such as `Windows.winmd` on Universal Windows Platform. To use custom .winmd files, import them into your Unity project folder together with any accompanying DLLs and configure them. For more information, refer to [Import and configure plug-ins](plug-in-inspector.md).

To use WinRT API in your Unity scripts:

* Your scripts must be written in C#.
* All the code that uses WinRT API must be under the `ENABLE_WINMD_SUPPORT` [directive](platform-dependent-compilation.md). This is necessary because the Editor uses [Mono](scripting-backends-mono.md), which doesn’t support WinRT APIs.

The following code example gets an advertising ID using WinRT API directly:

```
using UnityEngine;
public class WinRTAPI : MonoBehaviour
{
    void Update()
    {
        auto adId = GetAdvertisingId();
        // ...
    }

    string GetAdvertisingId()
    {
        #if ENABLE_WINMD_SUPPORT
            return Windows.System.UserProfile.AdvertisingManager.AdvertisingId;
        #else
            return "";
        #endif
    }
}
```

## Additional resources

* [Windows build settings reference](WindowsStandaloneBinaries.md)
* [Scripting backends](scripting-backends.md)
* [Microsoft documentation on WinRT APIs with Unity for HoloLens](https://learn.microsoft.com/en-us/windows/mixed-reality/develop/unity/using-the-windows-namespace-with-unity-apps-for-hololens)