# Test conditional compilation

The following example shows how to test your conditionally compiled code. It also prints a message based on the platform selected for the target build.

## Sample code

```
  using UnityEngine;
  using System.Collections;
  public class PlatformDefines : MonoBehaviour {
  void Start () {

    #if UNITY_EDITOR
      Debug.Log("Unity Editor");
    #endif

    #if UNITY_IOS
      Debug.Log("Unity iOS");
    #endif

    #if UNITY_STANDALONE_OSX
        Debug.Log("Standalone OSX");
    #endif

    #if UNITY_STANDALONE_WIN
      Debug.Log("Standalone Windows");
    #endif

  }          
  }
```

## Test instructions

1. Open the **Build Profiles** window (menu: **File** > **Build Profiles**).
2. Check that the platform you want to test your code on is the Active platform profile. If it isn’t, select your preferred platform and click **Switch Profile**.
3. Create a [script](creating-scripts.md) and copy and paste the [sample code](#Sample).
4. In the [Game view](GameView.md) **toolbar**, click the Play button to enter Play mode. Confirm that the code works by checking for messages relevant to the platform selected in the Unity console. For example, if you choose iOS, the messages `Unity Editor` and `Unity iOS` appear in the console.

## Additional resources

* [Unity scripting symbol reference](scripting-symbol-reference.md)
* [Custom scripting symbols](custom-scripting-symbols.md)