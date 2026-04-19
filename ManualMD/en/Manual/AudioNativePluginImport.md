# Use your native audio DSP plug-in and GUI in Unity

Once you create your [audio DSP plug-in](AudioNativeDSPPlugin.md) and [customize your GUI to your liking](AudioNativeCustomGUI.md), you can build your file and import it into Unity.

## 1. Compile and build your plug-ins

You must implement your **plug-ins** as a dynamic library for your preferred platform. Unlike with **scripts**, any platform that you want to support must be able to compile this file, along with platform-specific optimizations.

For information about how to build your plug-in into a library file for each platform, refer to [Building plug-ins for desktop platforms](plug-ins-for-desktop.md).

## 2. Rename your audio plug-in DLL file

* Add the prefix “audioplugin” to the name of your DLL file (case insensitive). For example `audioplugin-myplugin.dll`

Unlike other **native plug-ins**, Unity needs to load audio plug-ins before it creates any mixer assets that might need effects from a plug-in. But Unity doesn’t do this by default.

Add the prefix “audioplugin” to the name of your DLL file so that the Unity Editor detects and adds your plug-in to a list of plug-ins that Unity automatically loads as the build starts.

## 3. Import your audio plug-in into Unity

Click and drag your plug-in library file into the Asset folder of your Unity project. Unity automatically installs your plug-in and it is ready to use.

## 4. Link the plug-in to a platform

Native audio plug-ins use the same scheme as other native or **managed plug-ins**: You must use the plug-in importer **Inspector** to associate your plug-in with your preferred platforms. For more information, refer to [Plugin Inspector](plug-in-inspector.md).

## Additional resources

* [Audio Mixer](AudioMixer.md)
* [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md)
* [Audio Spatializer](AudioSpatializerSDK.md)
* [Building plug-ins for desktop platforms](plug-ins-for-desktop.md)