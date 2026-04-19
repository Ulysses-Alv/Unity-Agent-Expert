# Troubleshooting the Unity Editor

The following sections explain how to troubleshoot and prevent problems with the Unity Editor in different situations. In general, make sure your computer meets all the [system requirements](http://unity3d.com/unity/system-requirements), is up-to-date, and that you have the required user permissions in your system. Also, make backups regularly to protect your projects.

## Versions

You can install different versions of the Editor in different folders. However, make sure you backup your projects as these might be upgraded by a newer version, and you won’t be able to open them in an older version of Unity. Refer to [installing Unity](InstallingUnity.md) for further information.

Licenses of add-ons are valid only for the Unity versions that share the same major number, for example 3.x and 4.x. If you upgrade to a minor version of Unity, for example 4.0 to 4.1, the add-ons will be kept.

## Activation

The Unity Hub is the primary method for all license management operations. For information about all license activation methods, refer to [License activation methods](LicenseActivationMethods.md).

If you experience issues activating your license, refer to [License Troubleshooting](ActivationFAQ.md).

If you still have problems activating your license or logging in to your user account, contact [Unity Support](https://support.unity.com/hc/en-us).

## Failure to Start

If Unity crashes when starting then firstly make sure that your computer meets the minimal [system requirements](http://unity3d.com/unity/system-requirements). Also update to the latest graphic and sound drivers.

If you get disk write errors, check your user account restrictions. When in macOS, note the “root user” is not recommended and Unity hasn’t been tested in this mode. Unity must always have write permissions for its folders, but if you are granting them manually check these folders:

On Windows:

* Unity’s installation folder
* `%AllUsersProfile%\Unity` (typically C:\ProgramData\Unity)
* `C:\Documents and Settings\<user>\Local Settings\Application Data\Unity`
* `C:\Users\<user>\AppData\Local\Unity`

On macOS:

* Package contents of Unity.app
* `~Users/<USER>/Library/Application Support/Unity`
* `~Users/<USER>/Library/Logs/Unity`

Some users have experienced difficulties when using hard disks formatted with non-native partitions, and using certain software to translate data between storage devices.

### Fonts

Corrupt fonts can crash Unity, you can find damaged files following these steps:

On Windows:

1. Open the fonts folder on your computer, located in the “Windows” folder.
2. Select “Details” from the “View” menu.
3. Check the “Size” column for fonts with a “0” size, which indicates a problematic file.
4. Delete corrupt fonts and reinstall them.

On macOS:

1. Launch your Font Book application.
2. Select all the fonts.
3. Open the “File” menu and choose “Validate Fonts” -> problematic fonts will be shown as invalid.
4. Delete corrupt fonts and reinstall them.

The system might have resources constrained, for example running in a virtual machine. Use the Task Manager to find processes consuming lots of memory.

### Corrupt Project or Installation

Unity might try to open a project that is corrupt, and this might include the default sample project. If this is the case, rename or move the folder of the project. After Unity starts correctly, you can restore the project’s folder if wished.

In the event of a corrupt installation, you might need to reinstall Unity. Refer to the instructions below.

In Windows, there might be problems like installation errors, registry corruption, conflicts, and other issues. For example, error 0xC0000005 means the program has attempted to access memory that it shouldn’t. If you added new hardware or drivers recently, remove and replace the hardware to determine if it’s causing the problem. Run diagnostics software and check information on troubleshooting the operating system.

## Performance and Crashes

If the Editor runs slowly or crashes, particularly on builds, this might be caused by all of the available system resources being consumed. Close all other applications when you build the project. Clean up the system using its utilities, and consult the Task Manager (Windows) or Activity Monitor (MacOS) to find out if there are processes using lots of resources, for example memory. Sometimes virus protection software can slow down or even block the file system with its scanning process.

## Project Loss

There are many factors that can destroy a project. The recommended best practice is that you **constantly backup your projects** to prevent unfortunate accidents. When in macOS, activate Time Machine using an external hard disk reserved for this sole purpose. After a loss you can try any of the file recovery utilities that exist, but sometimes this is irreversible.

## Re-installation

Follow these steps to reinstall the Editor:

1. Uninstall Unity. When in macOS, drag the Unity app to trash.
2. Delete these files if present:

   * Windows:
     + `%AllUsersProfile%\Unity\` (typically C:\ProgramData\Unity)
   * MacOS:
     + `/Library/Application Support/Unity/`
3. Restart the computer.
4. Download the latest version from our website, since your original install might be corrupt: http://unity3d.com/unity/download/archive
5. Reinstall Unity.

## Additional resources

* [Install Unity](GettingStartedInstallingUnity.md)
* [License activation methods](LicenseActivationMethods.md)
* [License troubleshooting](ActivationFAQ.md)