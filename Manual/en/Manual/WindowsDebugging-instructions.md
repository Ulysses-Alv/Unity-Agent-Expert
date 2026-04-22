# Configure your debugging tool to debug Unity in Windows

Set up Windows Debugger (WinDbg) or Visual Studio to resolve Unity symbols so you can debug your Unity applications or the Unity Editor.

If your debugging tool is already set up to resolve Unity symbols and you want to know how to set up live or forensic debugging, refer to:

* [Set up live debugging for Unity](WindowsDebugging-live.md)
* [Set up forensic debugging for Unity](WindowsDebugging-forensic.md)

## Set up Windows Debugger (WinDbg) to debug Unity

Follow these instructions to configure Windows Debugger (WinDbg) to automatically download and resolve Unity store symbols. When WinDbg has access to these symbols, you can use it to debug your application or the Editor.

1. Open WinDbg as administrator.
2. Go to **File** > **Attach to process**. A list of applications shows in the window.
3. In the list, choose the application you want to debug:
   * To debug the Editor, use `Unity.exe`.
   * To debug your built game, use the .exe file of your game.
4. Select **Attach**. WinDbg connects to your application, which might take a few minutes.
5. Enter the following into the **Command** window in WinDbg to add Unity symbols to your debugger:

   ```
   .sympath+ SRV*c:\symbols-cache*http://symbolserver.unity3d.com/
   ```

As a result, WinDbg is ready to debug your Unity Editor or application.

### Detailed breakdown of the code

`.sympath+ SRV*c:\symbols-cache*http://symbolserver.unity3d.com/`

#### .sympath+

`.sympath` on its own sets a new symbol path, which overwrites the previous path. The `+` addition ignores the existing symbol path, and appends the symbol store lookup.

#### SRV\*c:\symbols-cache

The `SRV` indicates a remote server to fetch from. The `c:\symbols-cache` is a local path to cache the downloaded symbols. Before WinDbg retrieves symbols, it checks that path first so it doesn’t download the same symbol multiple times. You can change this path to any directory on your system.

#### \*https://symbolserver.unity3d.com/

This is the URL of the symbol store which your debugging software will fetch from.

## Set up Visual Studio to debug Unity

Follow these instructions to configure Visual Studio (2019 or later) to automatically download and resolve Unity store symbols. When Visual Studio has access to these symbols, you can use it to debug applications or the Editor.

1. Open Visual Studio.
2. Go to **Tools > Options**.
3. Expand the **Debugging** section and then select **Symbols**.
4. Specify a cache directory if not already specified.
5. Click **Add (+)**. This adds a new symbol location to the **Symbol file (.pdb) locations** list.
6. Enter `https://symbolserver.unity3d.com/` into the new list entry.

Visual Studio is ready to debug your Unity code.

## Additional resources

* [Unity debugging guide for Windows](WindowsDebugging.md)
* [Introduction to debugging Unity in Windows](WindowsDebugging-types.md)
* [Set up live debugging for Unity](WindowsDebugging-live.md)
* [Set up forensic debugging for Unity](WindowsDebugging-forensic.md)