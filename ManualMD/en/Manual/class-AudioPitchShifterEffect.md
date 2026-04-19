# Audio Pitch Shifter Effect

The **Audio Pitch Shifter Effect** is used to shift a signal up or down in pitch.

## Properties

![The AudioGroup Inspector displays the configurable properties of an Audio Pitch Shifter Effect.](../uploads/Main/AudioPitchShifterEffect.png)

The AudioGroup Inspector displays the configurable properties of an Audio Pitch Shifter Effect.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Pitch** | The pitch multiplier (range 0.5 x to 2.0 x, default 1.0 x). |
| **FFT Size** | The size of the FFT window used to analyze the audio signal during pitch shifting (range 256.0 to 4096.0, default = 1024.0). Higher values reduce smearing but require more processing power. |
| **Overlap** | How much each successive FFT window overlaps (range 1 to 32, default = 4). Higher values mean smoother transitions, but increasing this property by a value of 2 doubles the CPU usage. |
| **Max channels** | The maximum number of channels (range 0 to 16, default = 0 channels). |

AudioPitchShifterEffect