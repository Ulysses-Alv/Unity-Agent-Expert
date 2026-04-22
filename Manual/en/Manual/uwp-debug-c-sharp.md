# Debug C# code

Before you can debug C# code with **IL2CPP** for your Universal Windows Platform (UWP) application, you need to do the following:

1. Go to **File** > **Build Profiles**. In the **Platform Settings** section, enable **Script Debugging**.
2. In **Player** settings, enable the **InternetClient**, **InternetClientServer**, and **PrivateNetworkClientServer** capabilities.
     
   **Note:** You can also enable these capabilities in the manifest. You will need to change these capabilities from the Visual Studio manifest editor because the manifest isn’t overwritten when you build on top of an earlier build.

The debugging procedure for UWP applications is the same as for other platforms. For more information on how to debug C# code, refer to [Debug C# code in Unity](managed-code-debugging.md).

## Additional resources

* [Debug C# code in Unity](managed-code-debugging.md)
* [Microsoft documentation on how to Debug C# code using Visual Studio](https://learn.microsoft.com/en-us/visualstudio/get-started/csharp/tutorial-debugger?toc=%2Fvisualstudio%2Fdebugger%2Ftoc.json&view=vs-2022)
* [Microsoft documentation on Unity debugging with Visual Studio](https://learn.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/using-visual-studio-tools-for-unity?pivots=windows#unity-debugging)
* [IL2CPP overview](./scripting-backends-il2cpp.md)
* [IL2CPP scripting backend for UWP](uwp-il2cpp-scripting.md)