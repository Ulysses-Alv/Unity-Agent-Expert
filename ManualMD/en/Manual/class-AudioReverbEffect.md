# Audio SFX Reverb Effect

The **SFX Reverb Effect** takes the output of an [Audio Mixer](class-AudioMixer.md) group and distorts it to create a custom reverb effect.

## Properties

![The AudioGroup Inspector displays the configurable properties of an SFX Reverb Effect.](../uploads/Main/AudioSFXReverbEffect.png)

The AudioGroup Inspector displays the configurable properties of an SFX Reverb Effect.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Dry Level** | Mix level of dry signal in output in mB. Ranges from –10000.0 to 0.0. Default is 0 mB. |
| **Room** | Room effect level at low frequencies in mB. Ranges from –10000.0 to 0.0. Default is –10000.0 mB. |
| **Room HF** | Room effect high-frequency level in mB. Ranges from –10000.0 to 0.0. Default is 0.0 mB. |
| **Decay Time** | Reverberation decay time at low-frequencies in seconds. Ranges from 0.1 to 20.0. Default is 1.0. |
| **Decay HF Ratio** | Decay HF Ratio : High-frequency to low-frequency decay time ratio. Ranges from 0.1 to 2.0. Default is 0.5. |
| **Reflections** | Early reflections level relative to room effect in mB. Ranges from –10000.0 to 1000.0. Default is –10000.0 mB. |
| **Reflect Delay** | Early reflections delay time relative to room effect in mB. Ranges from –10000.0 to 2000.0. Default is 0.02. |
| **Reverb** | Late reverberation level relative to room effect in mB. Ranges from –10000.0 to 2000.0. Default is 0.0 mB. |
| **Reverb Delay** | Late reverberation delay time relative to first reflection in seconds. Ranges from 0.0 to 0.1. Default is 0.04 s. |
| **Diffusion** | Reverberation diffusion (echo density) in percent. Ranges from 0.0 to 100.0. Default is 100.0%. |
| **Density** | Reverberation density (modal density) in percent. Ranges from 0.0 to 100.0. Default is 100.0%. |
| **HFReference** | Reference high frequency in Hz. Ranges from 20.0 to 20000.0. Default is 5000.0 Hz. |
| **Room LF** | Room effect low-frequency level in mB. Ranges from –10000.0 to 0.0. Default is 0.0 mB. |
| **LFReference** | Reference low-frequency in Hz. Ranges from 20.0 to 1000.0. Default is 250.0 Hz. |

AudioReverbEffect