# UWP requirements and compatibility

This page contains system requirements and compatibility information for Universal Windows Platform (UWP) development in Unity.

## UWP application support

Unity supports the following CPU architectures when targeting UWP:

* x86
* x64
* ARM64

Unity supports the following operating system versions:

* Windows 10+
* HoloLens

Unity supports GPUs capable of running:

* DX10
* DX11
* DX12

You need the following additional software to develop with UWP:

* Windows 10+ SDK
* Visual Studio 2019 (or later)

## Activate Windows Developer Mode

You need to activate Windows Developer Mode on your PC and any other testing devices to be able to develop an application for UWP. When you open a UWP project in Visual Studio for the first time, a dialog will open that instructs you on how to select Developer Mode in your settings page. For more information, refer to the Microsoft documentation on how to [enable your device for development](https://learn.microsoft.com/en-us/windows/apps/get-started/enable-your-device-for-development).

## Player log

You can locate the player log under `<user>\AppData\Local\Packages\<productname>\TempState`. For more information, refer to [log files](log-files.md).