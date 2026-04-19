# Develop a native DSP audio plug-in

You can use the native Digital Signal Processing (DSP) **plug-in** to process audio and expose parameters for a user to experiment with **audio effects**. [The example plug-ins that Unity provides](https://github.com/Unity-Technologies/NativeAudioPlugins)
are a good place to experiment with plug-ins and to get ideas about parameters you need.

To develop a native audio plug-in for Unity:

1. [Create an audio plug-in file](#1-create-an-audio-plug-in-file).
2. [Define your parameters in your plug-in file](#2-define-your-parameters-in-your-plug-in-file).
3. [Instantiate your plug-in](#3-instantiate-your-native-audio-dsp-plug-in).
4. [Unload your plug-in](#4-unload-your-native-audio-dsp-plug-in).
5. [Handle the audio processing of your plug-in](#5-handle-audio-processing-in-your-plug-in).

## 1. Create an audio plug-in file

To create an audio plug-in file:

1. Download the [latest audio plug-in SDK](https://github.com/Unity-Technologies/NativeAudioPlugins).
2. In the folder, go to **NativeAudioPlugins** > **NativeCode**. The native example plug-in .cpp files are located here.
3. You can duplicate one of the plug-in .cpp files so you can use it as a template for your own plug-in, work directly in one of the example plug-in files (for example, `Plugin_Equalizer.cpp`), or create your own .cpp file.
4. Include `AudioPluginUtil.h` in your file if it’s not already there.

## 2. Define your parameters in your audio plug-in file

Create a list of parameters that would be useful for your user to interact with when they use your plug-in. To add your parameters to your plug-in:

1. In your plug-in .cpp file, define your parameters as enum values. For example:

   ```
   enum Param
   {
       P_FREQ, //Frequency parameter
       P_MIX,  //Mix parameter
       P_NUM   //An extra value to keep track of length of the enum
   };
   ```
2. Create an array of `UnityAudioParameterDefinitions` and set its size to be the number of parameters you have:

   ```
   int numparams = P_NUM;
   definition.paramdefs = new UnityAudioParameterDefinition [numparams];
   ```
3. Use the `RegisterParameter` function to register each of your enum values.

   ```
   int InternalRegisterEffectDefinition(UnityAudioEffectDefinition& definition)
   {
               int numparams = P_NUM;
           definition.paramdefs = new  UnityAudioParameterDefinition [numparams];
           RegisterParameter(definition, "Frequency", "Hz",
               0.0f, kMaxSampleRate, 1000.0f,
               1.0f, 3.0f,
               P_FREQ);
           RegisterParameter(definition, "Mix amount", "%",
               0.0f, 1.0f, 0.5f,
               100.0f, 1.0f,
               P_MIX);
           return numparams;
   }
   ```

The following table gives an overview of the `RegisterParameter` function, its parameters, and how it’s used in the code example above:

| **Parameter type and name** | **Variable in example code** | **Description** |
| --- | --- | --- |
| **`UnityAudioEffectDefinition definition`** | `definition` | The UnityAudioEffectDefinition structure contains an array of UnityAudioParameterDefinition. The RegisterParameter function inserts your parameter definition as an entry into this array. |
| **`char* name`** | `“Frequency”`, `“Mix Amount”` | The display name you want to give the parameter. |
| **`char* unit`** | `“Hz”`, `“%”` | The type of the value. |
| **`float minval`** | `0.0f` | The minimum value of the parameter. |
| **`float maxval`** | `kMaxSampleRate`, `1.0f` | The maximum value of the parameter. |
| **`float defaultval`** | `1000.0f`, `0.5f` | The default and initial value of the parameter. |
| **`float displayscale`** | `1.0f`, `100.0f` | Scale factor for the display of parameters only. For example, the percentage in the example code has a minimum value of 0, a maximum value of 1, and a scale factor 100.0f. This means that although the actual value is between 0 and 1, the value that is shown in the GUI in Unity is between 0% and 100%. |
| **`float displayexponent`** | `3.0f`, `1.0f` | Maps parameters to sliders. |
| **`int enumvalue`** | `P_FREQ` , `P_MIX` | Enum value to assign these values to. |

Unity generates a default GUI from these basic parameter definitions.

## 3. Instantiate your native audio DSP plug-in

To create the instance of your plug-in, use the `CreateCallback` function. Unity calls the `CreateCallback` function as soon as it creates your plug-in. It can be null.

```
struct EffectData
{
    struct Data
    {
        float p[P_NUM]; // Parameters
        float s;        // Sine output of oscillator
        float c;        // Cosine output of oscillator
    };
    union
    {
        Data data;
        unsigned char pad[(sizeof(Data) + 15) & ~15];
    };
};
```

```
UNITY_AUDIODSP_RESULT UNITY_AUDIODSP_CALLBACK CreateCallback(
    UnityAudioEffectState* state)
{
    EffectData* effectdata = new EffectData;
    memset(effectdata, 0, sizeof(EffectData));
    effectdata->data.c = 1.0f;
    state->effectdata = effectdata;
    InitParametersFromDefinitions(
        InternalRegisterEffectDefinition, effectdata->data.p);
    return UNITY_AUDIODSP_OK;
}
```

The `UnityAudioEffectState` object stores data it receives from the host and passes the data to all callback functions. The data it stores includes:

* the sampling rate
* the total number of samples processed (for timing)
* whether the plug-in is bypassed

## 4. Unload your native audio DSP plug-in

To free the plug-in instance, use the `ReleaseCallback` function.
Unity calls the `ReleaseCallback` function just before it frees the plug-in and also frees any data associated with this specific instance of the plug-in. After this, no further callbacks related to the instance happen.

```
UNITY_AUDIODSP_RESULT UNITY_AUDIODSP_CALLBACK ReleaseCallback(
    UnityAudioEffectState* state)
{
    EffectData::Data* data = &state->GetEffectData<EffectData>()->data;
    delete data;
    return UNITY_AUDIODSP_OK;
}
```

## 5. Handle audio processing in your plug-in

To handle the processing of audio, use the `ProcessCallback` function. Unity repeatedly calls the `ProcessCallback` function with a block of input audio to read from and an output block to write to.

The following code gives an example of a sine-wave being multiplied to all channels:

```
UNITY_AUDIODSP_RESULT UNITY_AUDIODSP_CALLBACK ProcessCallback(
    UnityAudioEffectState* state,
    float* inbuffer, float* outbuffer,
    unsigned int length,
    int inchannels, int outchannels)
{
    EffectData::Data* data = &state->GetEffectData<EffectData>()->data;

    float w = 2.0f * sinf(kPI * data->p[P_FREQ] / state->samplerate);
    for(unsigned int n = 0; n < length; n++)
    {
        for(int i = 0; i < outchannels; i++)
        {
            outbuffer[n * outchannels + i] =
                inbuffer[n * outchannels + i] *
                (1.0f - data->p[P_MIX] + data->p[P_MIX] * data->s);
        }
        data->s += data->c * w; // cheap way to calculate a sine-wave
        data->c -= data->s * w;
    }

    return UNITY_AUDIODSP_OK;
}
```

The `GetEffectData` function is a helper function that casts the `effectdata` field of the state variable to the `EffectData::Data` in the structure.

## Customize your GUI and import your audio plug-in into Unity

If you want to customize how Unity displays your plug-in parameters, refer to [Customize the GUI for your audio plug-in](AudioNativeCustomGUI.md).

To import your plug-ins into Unity, refer to [Use your native DSP plug-in and GUI in Unity](AudioNativePluginImport.md).

## Additional resources

* [Customize the GUI for your audio plug-in](AudioNativeCustomGUI.md)
* [Use your native DSP plug-in and GUI in Unity](AudioNativePluginImport.md).
* [Audio Mixer](AudioMixer.md)
* [Native audio plug-in SDK](AudioMixerNativeAudioPlugin.md)
* [Develop an audio spatializer plug-in](AudioSpatializerSDK.md)
* [Develop an ambisonic decoder plug-in](AudioDevelopAmbisonicDecoder.md)