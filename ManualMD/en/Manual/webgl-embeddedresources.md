# Embedded resources in Web

.NET assemblies can contain embedded resources. An embedded resource is a collection of binary data that’s part of the .NET assembly. You can access this binary data in code via a file-like API.

By default, Web builds don’t include embedded resources. This helps to keep the size of the final binary lower, because embedded resources can be rather large. However, some user code and .NET class library APIs depend on the presence of embedded resources to work properly. For example, string comparisons that use the `StringComparison.InvariantCultureIgnoreCase` value use embedded resources when comparing non-ASCII characters.

You can use the following script in the Unity Editor to enable embedded resources for a Web build:

```
using UnityEditor;

public class WebGLEditorScript
{
    [MenuItem("WebGL/Enable Embedded Resources")]
    public static void EnableEmbeddedResources()
    {
        PlayerSettings.WebGL.useEmbeddedResources = true;
    }
}
```

When this Player setting is enabled, the Web build includes any embedded resources in any .NET assemblies the Project uses.