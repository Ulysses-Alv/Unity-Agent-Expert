# Probe Volumes Options Override reference

Refer to [Fix issues with Adaptive Probe Volumes](probevolumes-fixissues.md) for more information about using the Probe Volumes Options Override.

| **Property** | **Description** |
| --- | --- |
| **Normal Bias** | Enable to move the position used by shaded pixels when sampling **Light Probes**. The value is in meters. This affects how sampling is moved along the **pixel**’s surface normal. |
| **View Bias** | Enable to move the sampling position towards the **camera** when sampling Light Probes. The results of **View Bias** vary depending on the camera position. The value is in meters. |
| **Scale Bias with Min Probe Distance** | Scale the **Normal Bias** or **View Bias** so it’s proportional to the spacing between Light Probes in a [brick](probevolumes-concept.md#how-probe-volumes-work). |
| **Sampling Noise** | Enable to increase or decrease the amount of noise URP adds to the position used by shaded pixels when sampling Light Probes. This can help [fix seams](probevolumes-troubleshoot-light-leaks.md#fix-seams) between bricks. |
| **Animate Sampling Noise** | Enable to animate sampling noise when Temporal Anti-Aliasing (TAA) is enabled. This can make noise patterns less visible. |
| **Leak Reduction Mode** | Enable to choose the method Unity uses to reduce leaks. Refer to [Fix light leaks](probevolumes-troubleshoot-light-leaks.md). Options: • **None**: No leak reduction. • **Performance**: The uvw used to sample APV data are warped to try to have invalid probe not contributing to lighting. This samples APV a single time so it’s a cheap option but will only work in the simplest cases. • **Quality**: This option samples APV between 1 and 3 times to provide the smoothest result without introducing artifacts. This is as expensive as Performance mode in the simplest cases, and is better and more expensive in the most complex cases. |
| **Min Valid Dot Product Value** | Enable to make URP reduce a Light Probe’s influence on a **GameObject** if the direction towards the Light Probe is too different to the GameObject’s surface normal direction. The value is the minimum [dot product](https://docs.unity3d.com/ScriptReference/Vector3.Dot.html) between the two directions where URP will reduce the Light Probe’s influence. |
| **Intensity Multiplier** | Set the strength of the light contribution from Adaptive Probe Volumes. A value of 0 means Unity doesn’t use the Adaptive Probe Volume data. |
| **Sky Occlusion Intensity Multiplier** | Set the strength of the light contribution from sky occlusion data in Adaptive Probe Volumes, if you enable [sky occlusion](probevolumes-skyocclusion.md). |