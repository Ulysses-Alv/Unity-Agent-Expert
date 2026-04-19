# Customize the GUI for your audio plug-in

You can customize how Unity displays the parameter descriptions you create in your [native DSP plug-in](AudioNativeDSPPlugin.md). By default, Unity displays these parameter descriptions as sliders. However, some **plug-ins** need more complex user interaction that allows the user more control.

You can use C# to customize the GUI and replace the default sliders. To access examples of custom GUIs:

1. [Download the latest audio plug-in SDK](https://github.com/Unity-Technologies/NativeAudioPlugins).
2. In the SDK, navigate to `NativeAudioPlugins-master\GUICode\AudioPluginDemoGUI`.
3. Open the `AudioPluginDemoGUI.sln` project in Visual Studio, or open the GUI files in this folder in another IDE.

This page gives an overview of steps to take to create a GUI for your native audio DSP plug-in.

## 1. Link your GUI to your audio plug-in

After Unity loads the **native plug-in** DLL files and registers the contained audio plug-ins, it searches for corresponding GUI files that match the names of the registered plug-ins.

To make sure Unity links your custom GUI to your plug-in:

1. Open your custom GUI class.
2. Make sure your GUI class inherits from `IAudioEffectPluginGUI`.
3. Override the **Name** property of your class so it is the same as your plug-in. This is what Unity uses to link your files, not the DLL file name. For example:

   ```
   public override string Name
    {
               get { return "Demo Equalizer"; }
    }
   ```

## 2. Add more complex user interaction to your GUI

The example audio plug-in SDK has a few examples of different ways you can present your parameters, including filter curves. You can use any of these examples in your project.

In your custom GUI class, create variables that you want to control via user interaction. For example:

```
private float masterGain;
private float lowGain, midGain, highGain;
```

Refer to the example projects for ideas for UI elements and
create your preferred UI elements.

## 3. Bind your GUI parameters to your audio plug-in parameters

You need to define how the parameters in your plug-in relate to your GUI customizations. The `OnGUI(IAudioEffectPlugin plugin)` boolean function receives a handle to the native plug-in and then uses it to read and write parameters the native plug-in has defined. To bind your parameters:

1. Find the `OnGUI(IAudioEffectPlugin plugin)` boolean function in your GUI class.
2. Use `plugin.GetFloatParameter` to read each parameter from your native plug-in and map it to your GUI variable.

   ```
   plugin.GetFloatParameter("MasterGain", out masterGain);
   ```

   This returns true if a parameter is found.
3. Use `plugin.SetFloatParameter` to set the value of your plug-in parameter to the current value of your GUI variable.

   ```
   plugin.SetFloatParameter("MasterGain", masterGain);
   ```

   This returns true if the parameter is found.
4. To avoid duplicate or contradicting parameter definitions in the native DSP and GUI code, use the `plugin.GetFloatParameterInfo` function.

   ```
   plugin.GetFloatParameterInfo("NAME", out minVal, out maxVal, out defVal);
   ```

   If your OnGUI function returns true, the **Inspector** shows the default UI sliders below the custom GUI. This is useful for GUI development because:
5. All the parameters are available while you develop your custom GUI.
6. You can check that the right actions result in the expected parameter changes.

## 4. Compile and import your GUI for your audio plug-in

To import your GUI for your plug-ins into Unity, refer to [Use your native DSP plug-in and GUI in Unity](AudioNativePluginImport.md).

## Additional resources

* [Audio Mixer](AudioMixer.md)
* [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md)
* [Develop an audio spatializer plug-in](AudioSpatializerSDK.md)
* [Develop an ambisonic decoder plug-in](AudioDevelopAmbisonicDecoder.md)

---