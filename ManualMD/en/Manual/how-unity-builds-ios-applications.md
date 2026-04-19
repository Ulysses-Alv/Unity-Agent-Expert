# How Unity builds iOS applications

Unity uses [Xcode](https://developer.apple.com/xcode/) to build iOS applications. You can use [iOS Player settings](class-PlayerSettingsiOS.md) to configure most aspects of the final build. However, for more granular control, building an Xcode project allows you to directly modify Xcode project files.

## The build process

1. Unity collects project resources, code libraries, and **plug-ins** from your Unity project and uses them to create a valid [Xcode project](StructureOfXcodeProject.md).
2. Unity updates the Xcode project based on the Unity project’s **Player settings** and build settings. Depending on whether you use [replace or append mode](iphone-BuildProcess.md#replace-and-append-mode), Unity replaces or preserves previous changes you made. Append mode preserves changes you previously made and only overwrites certain values. Replace mode creates a new project which overwrites any changes you previously made.
3. Unity generates C++ source files based on your C# **scripts** and places them in the generated Xcode project. Xcode then invokes the [IL2CPP](./scripting-backends-il2cpp.md) program which compiles the C++ source files into libraries called `libGameAssembly.a` and `il2cpp.a`.
4. Xcode builds the project into a standalone application and deploys and launches it on a connected device or the [Xcode simulator](https://developer.apple.com/documentation/xcode/running-your-app-in-simulator-or-on-a-device).

## Incremental build pipeline

Unity uses the [incremental build pipeline](building-introduction.md#incremental-build-pipeline) when it generates the Xcode project for iOS. This means that Unity incrementally builds and generates files such as [Information Property List](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AboutInformationPropertyListFiles.html) (plist) files and [Entitlement](https://developer.apple.com/documentation/bundleresources/entitlements) files. If you implement callbacks that modify or move any iOS file or asset that the incremental build pipeline uses, refer to [Creating clean builds](build-clean-build.md).

## Additional resources

* [Build an iOS application](iphone-BuildProcess.md)
* [iOS build settings reference](BuildSettingsiOS.md)