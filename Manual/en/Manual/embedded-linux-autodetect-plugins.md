# Autodetect plug-ins for Embedded Linux

Unity automatically detects plug-ins for Embedded Linux. When you import plug-ins, Unity creates metadata files for each **plug-in**. For example, a `.so` for shared objects and shared libraries and an `.a` for archive files. These metadata files contain the plug-in information, such as the target architecture and platform. The Unity build system refers to these metadata files for tracking which files to copy over during the build process.

You can have several shared libraries with the same name in a project. For example, you can have a `libFoo.so` for x86\_64 and a `libFoo.so` for ARM64 in the same project as Unity detects the correct `libFoo.so` for your build target and copies it across to the Player build.

You can edit these files manually in the Unity Editor. However, you don’t need to manually add plug-ins to the `Plugins` folder in your project. Instead, you can place them in special folders located under the project’s `Assets/Plugins/EmbeddedLinux` folder in the project directory so Unity automatically detects and sets their platform and architecture for you when importing.

## Auto detection rules

Unity automatically detects plug-ins for Embedded Linux based on the following rules:

* **Architecture-specific folders**: Place plug-ins under `Assets/Plugins/EmbeddedLinux/<arch>`, where `<arch>` is x86, x86\_64, ARM, or ARM64. Unity copies them only when building an app for the respective target architecture. For example, if you place a plug-in under `Assets/Plugins/EmbeddedLinux/x86_64`, Unity copies it to the Player build only when building for x86\_64.
* **Plugins with no architecture**: Place plug-ins with no architecture under `Assets/Plugins/EmbeddedLinux` and ensure that they’re checked for the target architecture through their respective ELF headers and that the appropriate architecture is assigned.