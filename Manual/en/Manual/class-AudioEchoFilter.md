# Audio Echo Filter

[Switch to Scripting](../ScriptReference/AudioEchoFilter.md "Go to AudioEchoFilter page in the Scripting Reference")

The **Audio Echo Filter** repeats a sound after a given **Delay**, attenuating the repetitions based on the **Decay Ratio**.

## Properties

![The AudioGroup Inspector displays the configurable properties of an Audio Echo Filter.](../uploads/Main/AudioEchoFilter.png)

The AudioGroup Inspector displays the configurable properties of an Audio Echo Filter.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Delay** | Echo delay in ms. 10 to 5000. Default = 500. |
| **Decay Ratio** | Echo decay per delay. 0 to 1. 1.0 = No decay, 0.0 = total decay (ie simple 1 line delay). Default = 0.5.L |
| **Wet Mix** | Volume of echo signal to pass to output. 0.0 to 1.0. Default = 1.0. |
| **Dry Mix** | Volume of original signal to pass to output. 0.0 to 1.0. Default = 1.0. |

## Details

The **Wet Mix** value determines the amplitude of the filtered signal, where the **Dry Mix** determines the amplitude of the unfiltered sound output.

Hard surfaces reflects the propagation of sound. For example a large canyon can be made more convincing with the Audio Echo Filter.

Sound propagates slower than light - we all know that from lightning and thunder. To simulate this, add an Audio Echo Filter to an event sound, set the Wet Mix to 0.0 and modulate the Delay to the distance between [AudioSource](class-AudioSource.md) and [AudioListener](class-AudioListener.md).

AudioEchoFilter