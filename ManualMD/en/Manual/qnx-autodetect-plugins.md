# Autodetect plug-ins for QNX

Unity automatically detects **plug-ins** for QNX. When you import plug-ins, Unity creates metadata files for each of those plug-ins. For example, `.so` for shared object or shared libraries and `.a` for archive files in QNX. These metadata files contain the plug-in information, such as the target architecture and platform. The Unity build system refers to these metadata files for tracking what files to copy over during the build process.

You can have several shared libraries with the same name in a project. For example, `libFoo.so` for x86\_64 and `libFoo.so` for arm64 in the same project and Unity detects the correct `libFoo.so` and copies it across to the Player build depending on the target you are building for.

You can edit these files manually in the Unity Editor. However, you don’t need to manually add plug-ins to the `Plugins` folder in your project. If you place them in special folders located under the project’s `Assets/Plugins/QNX` folder in the project directory, Unity automatically detects and sets their platform and architecture for you when importing.

## Autodetection rules

Unity automatically detects plug-ins for QNX based on the following rules:

* **Architecture-specific folders** - Place plug-ins under `Assets/Plugins/QNX/<arch>`, where `<arch>` is x86, x86\_64, armeabi-v7a, or arm64-v8a. Unity copies them only when building for the respective target architecture. For example, if you place a plug-in under `Assets/Plugins/QNX/x86_64`, Unity copies it to the player build only when building for x86\_64.
* **SDK-specific folders** - If the plug-in targets a specific SDK (QNX 7.0 or QNX 7.1), you can place it under `Assets/Plugins/QNX/<sdk>/<arch>` where `<sdk>` can be Neutrino70 or Neutrino71. Unity copies it only when building for the respective SDK version and target architecture. **Note:** QNX 7.1 doesn’t support x86, therefore this combination is logged as an error.
* **Architecture-specific plug-ins** - Place plug-ins under `Assets/Plugins/QNX` and ensure that they’re checked for the target architecture via its ELF headers and the appropriate architecture is assigned.

## Additional resources

* [Build and deliver a QNX project](qnx-build-and-deliver.md)