# Set up forensic debugging for Unity

Learn how to set up Visual Studio or WinDbg to debug your application or Unity Editor after it discovers an issue or crashes. This type of debugging is called forensic debugging.

Occasionally, an application doesn’t crash despite having the debugger attached, or it crashes on a remote device where the debugger isn’t available. In such cases, you can get useful information from the dump file the debugger creates. This information covers how to set up the debugger to create a dump file.

If you want to debug as the application runs instead, refer to [Set up live debugging for Unity](WindowsDebugging-live.md).

## Force applications to create a dump file

**Note:** These instructions apply to Windows Standalone and Universal Windows platforms when running on desktop.

To make your applications produce a dump file when they run into an issue:

1. Open the Windows registry.
2. Navigate to `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\Windows Error Reporting`.
3. Create a `LocalDumps` key (folder) if one doesn’t exist.
4. Create a new **String Value** called `DumpFolder` whose value is the fully qualified path where you want to store your dump files. For example, `C:\Temp`.
5. Create a new DWORD called `DumpCount` and set its value to `00000010`.
6. Create a new DWORD called `DumpType` and set its value to `00000002`.
7. Launch the application via the Windows Standalone or Universal Windows executable.
8. Reproduce the crash.

The application creates the dump file in the folder you specified earlier. You can open the dump file with your preferred debugging tool, such as Visual Studio or WinDbg.

## Additional resources

* [Unity debugging guide for Windows](WindowsDebugging.md)
* [Introduction to debugging Unity in Windows](WindowsDebugging-types.md)
* [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md)
* [Set up live debugging for Unity](WindowsDebugging-live.md)