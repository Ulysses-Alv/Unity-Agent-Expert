# Audio High Pass Simple Effect

The **Highpass Simple Effect** passes high frequencies of an AudioMixer group and cuts off signals with frequencies lower than the **Cutoff Frequency**.

## Properties

![The AudioGroup Inspector displays the configurable property of an Audio High Pass Simple Effect.](../uploads/Main/AudioHighPassSimpleEffect.png)

The AudioGroup Inspector displays the configurable property of an Audio High Pass Simple Effect.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **Cutoff freq** | Highpass cutoff frequency in Hertz (range 10.0 to 22000.0, default = 5000.0). |

## Details

The **Resonance** (short for Highpass Resonance Quality Factor) determines how much the filter’s self-resonance is dampened. Higher highpass resonance quality indicates a lower rate of energy loss, that is the oscillations die out more slowly.

For additional control over the resonance value of the high pass filter use the [Audio High Pass effect](class-AudioHighPassEffect.md).

AudioHighPassSimpleEffect