# Audio Spatializer SDK

The **audio spatializer** SDK provides controls to change the way your application transmits audio from an **audio source** into the surrounding space. It is an extension of the native audio plugin SDK.

The built-in panning of audio sources is a simple form of spatialization. It takes the source and regulates the gains of the left and right ear contributions based on the distance and angle between the **Audio Listener** and the Audio Source. This provides simple directional cues for the player on the horizontal plane.

## The Unity Audio Spatializer SDK and example implementation

To provide flexibility and support for working with audio spatialization, Unity has an open interface, the Audio Spatializer SDK, as an extension on top of the Native Audio Plugin SDK. You can replace the standard panner in Unity with a more advanced one, and give it access to important meta-data about the source and listener needed for the computation.

For an example of a native spatializer audio plugin, see the [Unity Native Audio Plugin SDK](https://github.com/Unity-Technologies/NativeAudioPlugins). The plugin only supports a direct Head-Related Transfer Function (HRTF), and is intended for example purposes only.

You can use a simple reverb, included in the plugin, to route audio data from the spatializer plugin to the reverb plugin. The HRTF filtering is based on a modified version of the KEMAR data set. For more information on the KEMAR data set, see the [MIT Media Lab’s documentation and measurement files](http://sound.media.mit.edu/resources/KEMAR.html).

If you would like to explore a data set obtained from a human subject, refer to [IRCAM’s data sets](http://recherche.ircam.fr/equipes/salles/listen).

### Initializing the Unity Audio Spatializer

Unity applies spatialization effects directly after the Audio Source decodes audio data. This produces a stream of audio data in which each source has its own separate effect instance. Unity only processes the audio from that source with its corresponding effect instance.

To enable a plugin to operate as a spatializer, you need to set a flag in the description bit-field of the effect:

```
definition.flags |= UnityAudioEffectDefinitionFlags_IsSpatializer;
```

If you set the `UnityAudioEffectDefinitionFlags_IsSpatializer` flag, Unity recognizes the plugin as a spatializer during the plugin scanning phase. When Unity creates an instance of the plugin, it allocates the `UnityAudioSpatializerData` structure for the `spatializerdata` member of the `UnityAudioEffectState` structure.

To use the spatializer in a project, select it in your Project settings (menu: **Edit** > **Project Settings** > **Audio**):

![Spatializer plugin selector](../uploads/Main/AudioSpatializerProjectSettings.png)

Spatializer plugin selector

Then, in the **Inspector** window for an Audio Source that you want to use with the spatializer plugin, enable **Spatialize**:

![Spatializer checkbox on Audio Source](../uploads/Main/AudioSpatializerEnable.png)

Spatializer checkbox on Audio Source

You can also enable the spatializer for an Audio Source through a C# script, using the `AudioSource.spatialize` property.

In an application with a lot of sounds, you might want to only enable the spatializer on your nearby sounds, and use traditional panning on the distant ones, to reduce the CPU load on the mixing thread for the spatializer effects.

If you want Unity to pass spatial data to an audio mixer plugin that is not a spatializer, you can use the following flag in the description bit-field:

```
definition.flags |= UnityAudioEffectDefinitionFlags_NeedsSpatializerData;
```

If a plugin initializes with the `UnityAudioEffectDefinitionFlags_NeedsSpatializerData` flag, the plugin receives the `UnityAudioSpatializerData` structure, but only the `listenermatrix` field is valid. For more information about UnityAudioSpatializerData, see the [Spatializer effect meta-data section](#SpatializerMetaData).

To stop Unity from applying distance-attenuation on behalf of a spatializer plugin, use the following flag:

```
definition.flags |= UnityAudioEffectDefinitionFlags_AppliesDistanceAttenuation;
```

The `UnityAudioEffectDefinitionFlags_AppliesDistanceAttenuation` flag indicates to Unity that the spatializer handles the application of distance-attenuation. For more information on distance-attenuation, see the [Attenuation curves and audibility](#AttenuationCurves) section.

### Spatializer effect meta-data

Unlike other Unity **audio effects** that run on a mixture of sounds, Unity applies spatializers directly after the Audio Source decodes audio data. Each instance of the spatializer effect has its own instance of `UnityAudioSpatializerData`, mainly associated with data about the Audio Source.

```
struct UnityAudioSpatializerData
{
    float listenermatrix[16];   // Matrix that transforms sourcepos into the local space of the listener
    float sourcematrix[16];     // Transform matrix of the Audio Source
    float spatialblend;         // Distance-controlled spatial blend
    float reverbzonemix;        // Reverb zone mix level parameter (and curve) on 
                                // the Audio Source
    float spread;               // Spread parameter of the Audio Source (0..360 degrees)
    float stereopan;            // Stereo panning parameter of the Audio Source (-1: fully left, 1: fully right)
                                // The spatializer plugin may override the distance attenuation to
                                // influence the voice prioritization (leave this callback as NULL 
                                // to use the built-in Audio Source attenuation curve)
    UnityAudioEffect_DistanceAttenuationCallback distanceattenuationcallback;
    float minDistance;          // The minimum distance of the Audio Source. 
                                // This value may be useful for determining when to apply near-field effects. 
    float maxDistance;          // The maximum distance of the Audio Source, or the 
                                // distance where the audio becomes inaudible to the listener. 
    
};
```

The structure contains fields corresponding to the properties of the Audio Source component in the Inspector: **Spatial Blend**, **Reverb Zone Mix**, **Spread**, **Stereo Pan**, **Minimum Distance**, and **Maximum Distance**.

The `UnityAudioSpatializerData` structure contains the full 4x4 transform matrices for the Audio Listener and Audio Source. The listener matrix is inverted so that you can multiply the two matrices to get a relative direction-vector. The listener matrix is always orthonormal, so you can quickly calculate the inverse matrix.

Unity’s audio system only provides the raw source sound as a stereo signal. The signal is stereo even when the source is mono or multi-channel, and Unity uses up- or down-mixing, as required.

### Matrix conventions

The `sourcematrix` field contains a copy of the Audio Source’s transformation matrix. For a default Audio Source on a **GameObject** that is not rotated, the matrix is a translation matrix where the position is encoded in elements 12, 13 and 14.

The `listenermatrix` field contains the inverse of the AudioListener’s transform matrix.

You can determine the direction vector from the AudioListener to the Audio Source as shown below, where L is the `listenermatrix` and S is the `sourcematrix`:

```
float dir_x = L[0] * S[12] + L[4] * S[13] + L[ 8] * S[14] + L[12];
float dir_y = L[1] * S[12] + L[5] * S[13] + L[ 9] * S[14] + L[13];
float dir_z = L[2] * S[12] + L[6] * S[13] + L[10] * S[14] + L[14];
```

The position in (L[12], L[13], L[14]) is actually the negative value of what you see in Unity’s Inspector window for the **camera** matrix. If the camera had also been rotated, you would also have to undo the effect of the rotation first. To invert a Transformation-Rotation matrix, transpose the top-left 3x3 rotation matrix of L, and calculate the positions as shown below:

```
float listenerpos_x = -(L[0] * L[12] + L[ 1] * L[13] + L[ 2] * L[14]);
float listenerpos_y = -(L[4] * L[12] + L[ 5] * L[13] + L[ 6] * L[14]);
float listenerpos_z = -(L[8] * L[12] + L[ 9] * L[13] + L[10] * L[14]);
```

For an example in the code for the Audio Spatializer plugin, see line 215 in the [Plugin\_Spatializer.cpp file](https://github.com/Unity-Technologies/NativeAudioPlugins/blob/master/NativeCode/Plugin_Spatializer.cpp).

### Attenuation curves and audibility

Unless you specify the `UnityAudioEffectDefinitionFlags_AppliesDistanceAttenuation` flag, as specified in the [Initializing the Unity Audio Spatializer](#Initialization) section, the Unity audio system still controls the distance-attenuation. Unity applies distance-attenuation to the sound before it enters the spatialization stage, and allows the audio system to know the approximate audibility of the source. The audio system uses approximate audibility for dynamic virtualization of sounds based on importance to match the user-defined **Max Real Voices** limit.

Unity does not retrieve audibility information from actual signal level measurements, but uses the combination of the values that it reads from the distance-controlled attenuation curve, the **Volume** property, and the mixer’s applied attenuations.

You can directly override the attenuation curve, or you can use the value calculated by the Audio Source’s curve as a base for modification. To override or modify the value, use the callback in the `UnityAudioSpatializerData` structure as shown below:

```
typedef UNITY_AUDIODSP_RESULT (UNITY_AUDIODSP_CALLBACK* UnityAudioEffect_DistanceAttenuationCallback)(
    UnityAudioEffectState* state,
    float distanceIn,
    float attenuationIn,
    float* attenuationOut);
```

You can also use a simple custom logarithmic curve, as shown below:

```
UNITY_AUDIODSP_RESULT UNITY_AUDIODSP_CALLBACK SimpleLogAttenuation(
    UnityAudioEffectState* state,
    float distanceIn,
    float attenuationIn,
    float* attenuationOut)
{
    const float rollOffScale = 1.0f; // Similar to the one in the Audio Project Settings
    *attenuationOut = 1.0f / max(1.0f, rollOffScale * distanceIn);
    return UNITY_AUDIODSP_OK;
}
```

## Using C# scripts from the Unity API

There are two methods on the Audio Source that allow setting and getting parameters from the spatializer effect: `SetSpatializerFloat` and `GetSpatializerFloat`. These methods work similarly to the `SetFloatParameter` and `GetFloatParameter` methods in the generic native audio plugin interface. However, `SetSpatializerFloat` and `GetSpatializerFloat` take an index to the parameter they must set or read, while `SetFloatParameter` and `GetFloatParameter` refer to the parameters by name.

The boolean property `AudioSource.spatializer` is linked to the **Spatialize** option in Unity’s Inspector window for an Audio Source. The property controls how Unity instantiates and deallocates the spatializer effect, based on the selected plugin in your Audio **Project Settings**.

If instantiation of your spatializer effect is very resource-intensive, in terms of memory or other resources in your project, it might be effective to allocate your spatialization effects from a preset “pool,” so that Unity does not need to create a new instance of the spatializer every time you need to use it. If you keep your Unity plugin interface bindings very lightweight and dynamically allocate your audio effects, you can avoid frame drops or other performance issues in your project.

## Known limitations of the example plugin

Due to the fast convolution algorithm, quick movements cause some zipper artifacts, which you can remove through the use of overlap-save convolution or cross-fading buffers.

The code also does not support tilting the listener’s head, whether the listener is directly attached to the player character, or a camera located elsewhere.