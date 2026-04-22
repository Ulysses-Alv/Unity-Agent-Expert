# Bonjour browser sample

For a simple example of how to use a native **plug-in**, download the [Bonjour Browser Sample](../uploads/Examples/iPhoneNativeCodeSample.zip).

This example demonstrates how you can invoke Objective-C code from a Unity iOS application. This application implements a simple Bonjour client and consists of:

* A Unity iOS Project where `Plugins\Bonjour.cs` is the C# interface to the native code, and `BonjourTest.cs` is the script that implements the application logic.
* Native code (located in `Assets/Plugins/iOS`) that will be added to the built Xcode project as described in [Automated plug-in integration](ios-native-plugin-automated-integration.md).