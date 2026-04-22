# Use managed UWP plug-ins

Managed **plug-ins** are managed .NET assemblies that are managed outside of Unity and compiled into dynamically-linked libraries (DLLs). For information about how to create and use managed plug-ins, refer to [Managed plug-ins](plug-ins-managed.md).

The **IL2CPP** **scripting backend** exposes the same .NET API surface as the Unity Editor and standalone Player. As a result, you can use the same plug-ins and you don’t have to compile separate versions to target different .NET APIs for Universal Windows Platform (UWP).

## Additional resources

* [Managed plug-ins](plug-ins-managed.md)
* [Import and configure plug-ins](plug-in-inspector.md)
* [Microsoft documentation on Managed code](https://learn.microsoft.com/en-us/dotnet/standard/managed-code)