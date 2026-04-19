# Native audio plug-in SDK

The Unity native audio **plug-in** SDK lets you develop custom audio plug-ins for Unity. You can use this SDK to extend the audio capabilities of Unity and create advanced audio processing solutions tailored to your project’s needs. Examples of custom audio plug-ins you can create include **audio effects** and **audio spatializers**.

The native audio plug-in system consists of two parts:

* The native audio Digital Signal Processing (DSP) plug-in
* The managed graphical user interface (GUI)

Refer to the following pages to learn more about how to create an audio plug-in, customize the plug-in’s GUI, and access useful examples.

| **Topic** | **Description** |
| --- | --- |
| **[Develop a native audio DSP plug-in for Unity](AudioNativeDSPPlugin.md)** | Learn how to create your own native DSP plug-in. |
| **[Customize the Unity GUI for your audio plug-in](AudioNativeCustomGUI.md)** | Learn how to customize the GUI of your audio plug-in. |
| **[Import your audio plug-in and GUI to Unity](AudioNativePluginImport.md)** | Learn how to prepare your plug-in and GUI for Unity and import them. |
| **[Example plug-ins](AudioNativePluginExamples.md)** | Example DSP plug-ins with and without GUI customization. |

## Important files your DSP and GUI code uses

The `AudioPluginInterface.h` file has the necessary structures, types, and function declarations required to create a custom audio plug-in.

Both native DSP and GUI DLLs can contain multiple plug-ins. To add multiple plug-in effects within the same DLL, Unity provides additional code to handle the effect definition and parameter registration in a unified manner:

* `AudioPluginUtil.h`
* `AudioPluginUtil.cpp`

If you want your DLLs to contain multiple effects, include `AudioPluginUtil.h` in your code.

## Additional resources

* [Audio Mixer](AudioMixer.md)
* [Audio Spatializer SDK](AudioSpatializerSDK.md)
* [Ambisonic Audio](AmbisonicAudio.md)

---