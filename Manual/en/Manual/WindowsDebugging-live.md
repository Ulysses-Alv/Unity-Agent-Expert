# Set up live debugging for Unity

Learn how to set up Visual Studio or WinDbg to debug your application or the Unity Editor as it runs. This type of debugging is called live debugging.

Live debugging is an effective way to identify and resolve issues as they occur during application runtime.

If you prefer to use crash dump files to debug your application, refer to [Set up forensic debugging for Unity](WindowsDebugging-forensic.md).

## Prerequisites

If you plan to debug unmanaged code (code written in low-level languages such as C, C++, or Assembly), you need to configure your debugger to resolve Unity symbols. For instructions, refer to [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md).

## Set up automatic exception and crash debugging

Follow these steps to enable your debugger to debug your application or the Editor while it runs.

1. Open a text editor (like Notepad) and save the empty page as a .reg file.
2. Paste the following code into the .reg file:

   ```
   Windows Registry Editor Version 5.00
   [HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\AeDebug]
    “Auto”=“1”
   [HKEY_LOCAL_MACHINE\SOFTWARE\Wow6432Node\Microsoft\Windows NT\CurrentVersion\AeDebug]
    “Auto”=“1”
   ```
3. Save the file to a known location.
4. Double-click the .reg file and select **Merge**.
5. In the dialog that appears, select **Yes** to change the Windows Registry.

As a result, your registry is now ready to accommodate automatic exception debugging.

## Live debug the Unity Editor

You can use Visual Studio or WinDbg to debug processes within the Unity Editor itself, rather than your game or application. This is useful if you want to investigate when the Editor behaves unexpectedly.

Follow these instructions to attach a debugger to the Editor on launch:

1. Navigate to the folder containing your Unity Editor application file.
2. In the address bar, enter `cmd` and press **Enter**. This opens the command prompt window and uses your directory as the working directory.
3. Enter the following command:

   ```
   Unity.exe -dbgbreak
   ```

This prompt launches Unity and suggests a debugger to attach to the session.

## Additional resources

* [Unity debugging guide for Windows](WindowsDebugging.md)
* [Introduction to debugging Unity in Windows](WindowsDebugging-types.md)
* [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md)
* [Set up forensic debugging for Unity](WindowsDebugging-forensic.md)