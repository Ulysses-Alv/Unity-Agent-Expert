# Create a native plug-in for iOS

Create a native **plug-in** for iOS and import it into your Unity project.

## Define an extern method

For each native function you want to call, define an extern method in the C# file as follows:

```
[DllImport ("__Internal")] 
    
private static extern float FooPluginFunction();
```

## Use C linkage to prevent name mangling

If using C++ (.cpp) or Objective-C++ (.mm) to implement your plug-in, declare functions with C linkage to avoid issues with name mangling:

```
extern "C" {
  float FooPluginFunction();
}
```

Plug-ins written in C or Objective-C don’t need this, because these languages don’t use name mangling.

## Import your native plug-in into your Unity Project

Add your native code source files to the Unity Project’s `Assets` Folder.

## Configure plug-in settings

To configure plug-in settings for iOS:

1. Select the plug-in and view it in the ****Inspector****.
2. In the **Select platforms for plugin** section, enable **iOS**.
3. In the **Platform settings** section, set **CPU** to the CPU architecture for your plug-in.
4. Select **Apply**.

## Additional resources

* [Use your native plug-in for iOS](ios-native-plugin-use.md)