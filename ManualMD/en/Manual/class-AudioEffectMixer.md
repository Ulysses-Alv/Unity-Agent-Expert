# Audio Effects

You can modify the output of [Audio Mixer](class-AudioMixer.md) components by applying **Audio Effects**. These can filter the frequency ranges of the sound or apply reverb and other effects.

The effects are applied by adding effect components to the relevant section of the Audio Mixer. The ordering of the components is important, since it reflects the order in which the effects will be applied to the source audio. For example, in the image below, the Music section of an Audio Mixer is modified first by a Lowpass effect and then a compressor Effect, Flange Effect and so on.

![The Audio Mixer window displays Audio Groups called Master, Music, Reverb, and Effects. The Music group has multiple effect components, including a Compressor, Low Pass, High Pass, and Flange effects.](../uploads/Main/AudioMixer1.png)

The Audio Mixer window displays Audio Groups called Master, Music, Reverb, and Effects. The Music group has multiple effect components, including a Compressor, Low Pass, High Pass, and Flange effects.

To change the ordering of these and any other components, open a context menu in the **inspector** and select the *Move Up* or *Move Down* commands. Enabling or disabling an effect component determines whether it will be applied or not.

Though highly optimized, some filters are still CPU intensive. Audio CPU usage can monitored in the [profiler](Profiler.md) under the Audio Tab.

See the other pages in this section for further information about the specific effect types available.

AudioEffectMixer