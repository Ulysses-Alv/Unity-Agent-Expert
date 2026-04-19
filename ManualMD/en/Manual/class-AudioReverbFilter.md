# Audio Reverb Filter

[Switch to Scripting](../ScriptReference/AudioReverbFilter.md "Go to AudioReverbFilter page in the Scripting Reference")

The **Audio Reverb Filter** takes an [Audio Clip](class-AudioClip.md) and distorts it to create a custom reverb effect.

## Properties

![The AudioGroup Inspector displays the configurable properties of an Audio Reverb Filter.](../uploads/Main/AudioReverbFilter.png)

The AudioGroup Inspector displays the configurable properties of an Audio Reverb Filter.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Reverb Preset** | Custom reverb presets, select user to create your own customized reverbs. |
| **Dry Level** | Mix level of dry signal in output in mB. Ranges from –10000.0 to 0.0. Default is 0. |
| **Room** | Room effect level at low frequencies in mB. Ranges from –10000.0 to 0.0. Default is 0.0. |
| **Room HF** | Room effect high-frequency level in mB. Ranges from –10000.0 to 0.0. Default is 0.0. |
| **Room LF** | Room effect low-frequency level in mB. Ranges from –10000.0 to 0.0. Default is 0.0. |
| **Decay Time** | Reverberation decay time at low-frequencies in seconds. Ranges from 0.1 to 20.0. Default is 1.0. |
| **Decay HFRatio** | Decay HF Ratio : High-frequency to low-frequency decay time ratio. Ranges from 0.1 to 2.0. Default is 0.5. |
| **Reflections Level** | Early reflections level relative to room effect in mB. Ranges from –10000.0 to 1000.0. Default is –10000.0. |
| **Reflections Delay** | Early reflections delay time relative to room effect in seconds. Ranges from 0 to 0.3. Default is 0.0. |
| **Reverb Level** | Late reverberation level relative to room effect in mB. Ranges from –10000.0 to 2000.0. Default is 0.0. |
| **Reverb Delay** | Late reverberation delay time relative to first reflection in seconds. Ranges from 0.0 to 0.1. Default is 0.04. |
| **HFReference** | Reference high frequency in Hz. Ranges from 1000.0 to 20000.0. Default is 5000.0. |
| **LFReference** | Reference low-frequency in Hz. Ranges from 20.0 to 1000.0. Default is 250.0. |
| **Diffusion** | Reverberation diffusion (echo density) in percent. Ranges from 0.0 to 100.0. Default is 100.0. |
| **Density** | Reverberation density (modal density) in percent. Ranges from 0.0 to 100.0. Default is 100.0. |

**Note:** These values can only be modified if your **Reverb Preset** is set to **User**, else these values will be grayed out and they will have default values for each preset.

AudioReverbFilter