# Automated plug-in integration

Unity supports automated **plug-in** integration and copies files with the following extensions to the generated Xcode project: `.a`, `.m` , `.mm`, `.c`, `.cpp`, `.h`, `.swift`.

To enable automated plug-in integration, enable iOS plug-ins in the Plug-in **Inspector**.

If you place files with these extensions in the `Assets/Plugins/iOS` folder, Unity only enables them for the iOS platform.

**Note:** Files copied to the generated Xcode project are no longer linked to their counterparts in your Unity Project. If you change these files in Xcode, you must copy them back into your Unity Project. Otherwise, Unity will overwrite them the next time you build your Project.

## Additional resources

* [Bonjour browser code sample](ios-native-plugin-bonjour-sample.md)