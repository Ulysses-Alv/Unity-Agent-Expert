# Optimize Enlighten Realtime Global Illumination

Enlighten Realtime **Global Illumination** uses a set of **lightmaps** to store real-time indirect bounces. For this reason, enabling it increases memory requirements, even if you are using it along with Baked Global Illumination.

The number of **shader** calculations needed to generate lighting also increases when you use **Enlighten** Realtime Global Illumination because it samples an additional set of lightmaps and **Light Probes**.

If Enlighten Realtime Global Illumination doesn’t respond quickly enough to changes in your **Scene** lighting, there are several ways to address this problem:

* Reduce the real-time lightmap resolution to speed up calculation at runtime.
* Increase the CPU Usage setting for **Realtime GI** in the Quality Settings window. The tradeoff is that other systems receive less CPU time to do their work. Whether this is acceptable depends on your Project. This is a per-Scene setting, so you can dedicate more or less CPU time based on the complexity of each individual Scene in your Project.