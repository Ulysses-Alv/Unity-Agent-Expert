# Example native audio plug-ins included in the SDK

The [audio plug-in SDK](https://github.com/Unity-Technologies/NativeAudioPlugins) contains a variety of different audio **plug-in** types. The plug-in code in the SDK is available in the public domain so you can use this code as a template. The project contains examples of native DSP plug-ins with and without a customized GUI.

If you want to use an IDE to edit and compile the example plug-in code, the SDK includes projects for multiple IDEs. The following table shows you where to find the projects for each of the IDEs:

| **IDE** | **Location** |
| --- | --- |
| **Visual Studio** | `NativeAudioPlugins-master\NativeCode\VisualStudio\AudioPluginDemo.sln` (Plug-in project). |
| **Visual Studio** | `NativeAudioPlugins-master\GUICode\AudioPluginDemoGUI\AudioPluginDemoGUI.sln` (GUI project). |
| **XCode** | `NativeAudioPlugins-master\NativeCode\Xcode\AudioPluginDemo.xcodeproj` |
| **UWP** | `NativeAudioPlugins-master\NativeCode\UWP\AudioPluginDemo.sln` |

To access the example plug-ins in Unity:

1. Open the `demos` Unity project (folder: **NativeAudioPlugin** > **Assets**).
2. In the Unity project, go to **Assets** > **mixers**.
3. Click on the arrow next to some of the mixers to expand them.
4. Click on any Audio Mixer Group Controller. In the **Inspector**, it shows any effects attached to the Audio Mixer Group Controller.
5. To add an effect, click the **Add Effect** button. This lists all of the example plug-in effects that you can add.
6. Select the effect you want to add. Unity displays your effect in the Inspector.

If you enter Play mode with the Inspector open, some of the effects update as the sounds change.

## Example plug-ins without custom GUIs

This section lists a few example plug-ins included in the `AudioNativePlugin` project that have the default Unity sliders instead of a custom GUI.

To access the plug-in code files directly, in the SDK folder, go to `NativeAudioPlugins/NativeCode`.

![Simple plug-ins without custom GUIs.](../uploads/Main/AudioMixerStereoWidener.png)

Simple plug-ins without custom GUIs.

### The NoiseBox plug-in

The NoiseBox plug-in is a plug-in that adds and multiplies the input signal by white noise at variable frequencies. You could use this in combination with the Lofinator plug-in and animated parameters to simulate effects such as mobile phone sounds, poor radio reception, and broken loudspeakers.

### The Ring Modulator example plug-in

The Ring Modulator example plug-in is an example of a simple modulator that multiplies any incoming signal by a sine wave. This gives a radio-noise or lost-signal effect. To increase the effect, chain together multiple ring modulation effects with different frequencies.

### The StereoWidener example plug-in

The StereoWidener plug-in decomposes a stereo input signal into mono and side components with variable delay and then recombines these to increase the perceived stereo effect.

### The Lofinator plug-in

The Lofinator plug-in does simple downsampling and quantization of the signal. You could use this in combination with the NoiseBox plug-in and animated parameters to simulate effects such as mobile phone sounds, poor radio reception, and broken loudspeakers.

## Example audio plug-ins with custom GUIs

This section provides advanced plug-in use cases, including effects for equalization and multiband **compression**. These plug-ins have a higher number of parameters than the plug-ins in the previous section. Some plug-in parameters are too complex for the default sliders.

To access the example custom GUI code directly, in the SDK folder, navigate to `NativeAudioPlugins-master\GUICode\AudioPluginDemoGUI`.

### The Equalizer example plug-in

The Equalizer plug-in includes a display graph showing the result curve and the individual filter contributions. To change one or multiple parameter values, click and drag on the parameter(s) on the control graph. This way, you don’t need to change each slider individually.

![The Inspector window displays the Equalizer plug-in with control graph and sliders for individual configurable parameters.](../uploads/Main/AudioMixerCustomGUI.png)

The Inspector window displays the Equalizer plug-in with control graph and sliders for individual configurable parameters.

To change the gains and frequencies of the filter curve, drag the three bands on the control graph. To change the shape of each band, hold Shift and click and drag the band.

You can use Unity’s internal API functions to plot the curves and draw antialiased curves for the frequency response.

### Code to analyze plug-in data and GUI performance

Both Equalizer and Multiband plug-ins also provide code to overlay the input and output spectra for visualizing the effect of the plug-ins. The GUI code runs at much lower update rate than the audio processing and doesn’t have access to the audio streams. Therefore, to read this data, the native code provides the following function:

```
UNITY_AUDIODSP_RESULT UNITY_AUDIODSP_CALLBACK GetFloatParameterCallback(
    UnityAudioEffectState* state,
    int index,
    float* value,
    char *valuestr)
{
    EffectData::Data* data = &state->GetEffectData<EffectData>()->data;
    if(index >= P_NUM)
        return UNITY_AUDIODSP_ERR_UNSUPPORTED;
    if(value != NULL)
        *value = data->p[index];
    if(valuestr != NULL)
        valuestr[0] = 0;
  return UNITY_AUDIODSP_OK;
}
```

This function enables the GUI to read floating-point data from the **native plug-in**. There is also an array variant called `GetFloatBufferCallback`. The plug-in system accepts any data as long as the request doesn’t affect the performance of the UI or the native code.

For the Equalizer and Multiband code, a utility class called `FFTAnalyzer` makes it easy to feed in input and output data from the plug-in and get a spectrum back. Then, `GetFloatBufferCallback` resamples this spectrum data and sends it to the C# GUI code.

The data must be resampled so that the `FFTAnalyzer` runs the analysis at a fixed frequency resolution. `GetFloatBufferCallback` just returns the number of samples requested, which is determined by the width of the view that’s displaying the data.

### The CorrelationMeter example plug-in

The CorrelationMeter plug-in has a minimal amount of DSP code. This plug-in plots the amplitude of the left channel against the amplitude of the right channel to show the stereo effects of the signal.

![Custom GUI of the CorrelationMeter plugin.](../uploads/Main/AudioMixerDemoCorrelationMeter.png)

Custom GUI of the CorrelationMeter plugin.

![Equalizer GUI with overlaid spectrum analysis (green curve is source, red is processed).](../uploads/Main/AudioMixerDemoMultiband.png)

Equalizer GUI with overlaid spectrum analysis (green curve is source, red is processed).

Both the Equalizer and Multiband effects are kept intentionally simple and unoptimized, but they serve as good examples of more complex UIs that are supported by the plug-in system.

### Loudness Monitoring tool

This is an example of a loudness monitoring tool measuring levels at 3 different time-scales.
This is just for demonstration purposes, but lets you build a monitor tool that
conforms to modern loudness standardizations. The curve rendering code is built into Unity.

![The Inspector window displays the loudness monitoring tool plug-in with graph and sliders for individual configurable parameters.](../uploads/Main/AudioMixerDemoLoudnessMeter.png)

The Inspector window displays the loudness monitoring tool plug-in with graph and sliders for individual configurable parameters.

### Synchronizing to the DSP clock

You can also use the plug-in system to generate sound instead of processing it.

`Plugin_TeeBee.cpp` and `Plugin_TeeDee.cpp` are simple synths that:
\* generate patterns with random notes
\* have parameters for tweaking the filters and envelopes in the synthesis engine

`ProcessCallback` reads the `state->dsptick` parameter to determine the position in a song. This counter is a global sample position. To make all plug-in effects stay in sync to the same sample-based clock:

1. Divide the counter by the length of each note specified in the sample.
2. When this division has a zero remainder, fire a note event to the synthesis engine.

If you want to play a pre-recorded piece of music with a known tempo through such an effect, use the timing information to apply tempo-synchronized filter effects or delays on the music.

## Spatialization

The Spatialization SDK is an example of a plug-in that uses the native audio plug-in SDK as a foundation. The spatialization SDK lets you develop custom spatialization effects to instantiate per **audio source**.
For more information, refer to [Audio Spatializer SDK](AudioSpatializerSDK.md).

## Additional resources

* [Audio Mixer](AudioMixer.md)
* [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md)
* [Audio Spatializer](AudioSpatializerSDK.md)