# Audio Settings

[Switch to Scripting](../ScriptReference/AudioSettings.md "Go to AudioSettings page in the Scripting Reference")

The AudioSettings class contains various bits of global information relating to the sound system, but most importantly it contains API that allows resetting the audio system at runtime in order to change settings such as speaker mode, sample rate (if supported by the platform), DSP buffer size and real/virtual voice counts.

Many of these settings can also be configured in the Audio section of the **project settings**. When changed these settings will apply both to the editor and define the initial state of the game while changes performed using the AudioSettings API only apply to the runtime of the game and are reset back to the state defined in the project settings when stopping the game in the editor.

The game may provide a sound options menu in which the user can change the sound settings or changes may come from outside in response to a device change such as plugging in an external audio input/output device or even an HDMI monitor which may also act as an audio device. The AudioConfiguration AudioSettings.GetConfiguration() / bool AudioSettings.Reset(AudioConfiguration config) API can read and apply global changes to the current sound system configuration and essentially replaces the AudioSettings.SetDSPBufferSize(…) function and the AudioSettings.outputSampleRate, AudioSettings.speakerMode which had the side effect of reinitializing the whole audio system when the properties were modified.

The API defines the AudioSettings.OnAudioConfigurationChanged(bool device) to set up a callback through which **scripts** can be notified about audio configuration changes. These can either be caused by actual device changes or by a configuration initiated by script.

It is important to note that whenever runtime modifications of the global audio system configuration are performed all audio objects have to be reloaded. This works for disk-based AudioClip assets and audio mixers, but any AudioClips generated or modified by scripts are lost and have to be recreated. Likewise any play state is lost too, so these need to be regenerated in the AudioSettings.OnAudioConfigurationChanged(…) callback.

For details and examples see the scripting API reference.

AudioSettings