# Include additional files in a build

`StreamingAssets` is a [reserved folder](SpecialFolders.md) that you can use to make files available to a Player build directly and bypass the standard project build process. The standard build process serializes **scenes** and assets into binary files inside the generated Player. `StreamingAssets` allows you to add files that are already in the correct format for the target platform and that your application can load directly.

Example usages include:

* Configuration files in JSON, XML, SQLite, or other formats.
* Video files, for example movie files on iOS devices.
* Files required by **plug-ins**.
* AssetBundles (described below).

The `StreamingAssets` folder must be in the root of the `Assets` folder at `Assets/StreamingAssets`. The name is case sensitive and must be `StreamingAssets` exactly. Any files in this directory are copied without modification into the Player output. When deployed to a target device the files are copied to a location appropriate for the platform.

**Tip**: To avoid the overhead of Unity automatically importing every file from `StreamingAssets`, you can add content from other directories during the build process by calling [`AddAdditionalPathToStreamingAssets`](../ScriptReference/Build.BuildPlayerContext.AddAdditionalPathToStreamingAssets.md) from a callback that implements [`BuildPlayerProcessor.PrepareForBuild`](../ScriptReference/Build.BuildPlayerProcessor.PrepareForBuild.md). This is useful if, for example, your content is large and generated.

## Accessing streaming assets

The location of the `StreamingAssets` folder in your deployed application varies between platforms. To retrieve the path to the `StreamingAssets` folder reliably, use the [`Application.streamingAssetsPath`](../ScriptReference/Application-streamingAssetsPath.md) property because it always points to the correct location on the current host platform.

On most platforms [`Application.streamingAssetsPath`](../ScriptReference/Application-streamingAssetsPath.md) is a directory accessible using regular file system APIs.

On Android and the Web platform, it’s impossible to access the streaming asset files directly via file system APIs because these platforms return a URL. Use the [`UnityWebRequest`](../ScriptReference/Networking.UnityWebRequest.md) class to access the content instead.

The following example loads a file from the `StreamingAssets` folder:

```
using UnityEngine;
using UnityEngine.Networking;
using System.Collections;
using System;

public class LoadStreamingAssets : MonoBehaviour
{
    void Start()
    {
        // Pass the specific filename and the callback method
        StartCoroutine(LoadFile("config.json", OnFileLoaded));
    }

    // This callback receives the raw string data once the coroutine finishes
    void OnFileLoaded(string loadedText)
    {
        if (string.IsNullOrEmpty(loadedText))
        {
            Debug.LogError("File data was empty or failed to load.");
            return;
        }

        // The data is now safely out of the coroutine.
        // You can parse it here (e.g., JsonUtility, XML, custom parsing)
        Debug.Log("File loaded successfully: " + loadedText);
    }

    IEnumerator LoadFile(string fileName, Action<string> callback)
    {
        string fileContent = "";
        string filePath = System.IO.Path.Combine(Application.streamingAssetsPath, fileName);

        if (filePath.StartsWith("jar") || filePath.StartsWith("http"))
        {
            // Special case to access StreamingAsset content on Android and Web
            UnityWebRequest request = UnityWebRequest.Get(filePath);
            yield return request.SendWebRequest();

            if (request.result == UnityWebRequest.Result.Success)
            {
                fileContent = request.downloadHandler.text;
            }
            else
            {
                Debug.LogError($"Failed to load StreamingAsset: {request.error}");
            }
        }
        else
        {
            // Regular file path on most platforms and in Editor
            if (System.IO.File.Exists(filePath))
            {
                fileContent = System.IO.File.ReadAllText(filePath);
            }
            else
            {
                Debug.LogError($"File not found at: {filePath}");
            }
        }

        // Send the loaded content back to the caller
        callback?.Invoke(fileContent);
    }
}
```

**Note:** If you want to load the data from a JSON file, use the methods in the [`JsonUtility` class](../ScriptReference/JsonUtility.md).

### StreamingAssets folder limitations

* At runtime the `StreamingAssets` location is read-only and you can’t modify or write new files to it. To write files, you can use [`Application.persistentDataPath`](../ScriptReference/Application.persistentDataPath.md) instead.
* Don’t put Unity files (such as `.unity`, `.prefab`, and `.asset`) into the `StreamingAssets` folder. These file types must be processed by a Player or [AssetBundle](AssetBundles-Building.md) build to be loadable at runtime.
* `.dll` and script files located in the `StreamingAssets` folder aren’t included during script compilation.

## AssetBundles and the StreamingAssets folder

The `StreamingAssets` folder is useful if you intend to distribute AssetBundles directly in the Player installation, rather than downloading them on-demand.

To do this:

1. Build the [AssetBundles](AssetBundles-Building.md) to an output path inside `Assets/StreamingAssets`. **Note**: The `.manifest` files generated alongside the AssetBundle files aren’t required in the runtime and can be removed. For more information, refer to [AssetBundles file format](assetbundles-file-format.md).
2. [Build the Player](BuildSettings.md). The AssetBundles are included in the platform-specific output.
3. Write loading code that uses the `Application.streamingAssetsPath` to determine the path to the AssetBundles. Depending on the platform, you can load them as a local file path, or with [`UnityWebRequestAssetBundle`](ScriptingRef:Networking.UnityWebRequestAssetBundle.GetAssetBundle). **Note**: On Android it’s not necessary to cache the AssetBundle or to perform CRC checks, because the file is already present on local storage.

The [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest) package uses this mechanism automatically for local content.

Alternatively, you can host AssetBundles and Addressables on a remote server and downloaded them on-demand rather than using the `StreamingAssets` folder. This is preferred for situations where the content is large, or frequently updated, or where you want to avoid the overhead of releasing new Player builds when you want to change or add new content.

## Additional resources

* [Loading resources at runtime](LoadingResourcesatRuntime.md)
* [Modifying source assets from code](ModifyingSourceAssetsThroughScripting.md)