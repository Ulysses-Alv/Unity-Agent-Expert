# Audio High Pass Filter

[Switch to Scripting](../ScriptReference/AudioHighPassFilter.md "Go to AudioHighPassFilter page in the Scripting Reference")

The **Audio High Pass Filter** passes high frequencies of an AudioSource and cuts off signals with frequencies lower than the **Cutoff Frequency**.

## Properties

![The AudioGroup Inspector displays the configurable properties of an Audio High Pass Filter.](../uploads/Main/AudioHighPassFilter.png)

The AudioGroup Inspector displays the configurable properties of an Audio High Pass Filter.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Cutoff Frequency** | Highpass cutoff frequency in Hertz (range 10.0 to 22000.0, default = 5000.0). |
| **Highpass Resonance Q** | Highpass resonance quality value (range 1.0 to 10.0, default = 1.0). |

## Details

The **Highpass resonance Q** (short for Highpass Resonance Quality Factor) determines how much the filter’s self-resonance is dampened. Higher highpass resonance quality indicates a lower rate of energy loss, that is the oscillations die out more slowly.

AudioHighPassFilter