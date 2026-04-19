# Call native plug-ins for iOS

Your app can only call iOS native **plug-ins** when deployed on an actual device. Wrap all native code methods with an additional C# code layer to only call native methods when the app is running on a device. Store this C# file in your Project’s `Assets` folder.

This C# layer can use [platform conditional compilation](platform-dependent-compilation.md) or check [`Application.platform`](../ScriptReference/Application-platform.md). For code running in the Unity Editor, return placeholder values.

Refer to the following sections for simple implementations of these methods. For a more detailed implementation, download the [Bonjour Browser sample](ios-native-plugin-bonjour-sample.md).

## Use conditional compilation

Platform dependent compilation is faster than `Application.platform` as it’s evaluated at compile time, rather than runtime.

Use the following to implement conditional compilation:

```
void MyMethod()
{
#if UNITY_IOS && !UNITY_EDITOR
    CallNativeMethodImplementation();
#else
    CallEditorMethodImplementation();
#endif
}
```

## Check Application.platform

Use the following to implement [`Application.platform`](../ScriptReference/Application-platform.md) and return placeholder values in the Editor:

```
void MyMethod()
 {
    if (Application.platform != RuntimePlatform.OSXEditor)
    {
        return _GetLookupStatus();
    }
    else
    {
        return "Done";
    }
}
```

## Additional resources

* [Callback from native code](ios-native-plugin-call-back.md)
* [Bonjour Browser sample](ios-native-plugin-bonjour-sample.md).