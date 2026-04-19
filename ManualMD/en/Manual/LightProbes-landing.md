# Precalculating indirect light with Light Probes

Resources and techniques for using **Light Probes** to store the light at specific points in a **scene**, so that Unity can calculate indirect lighting for **GameObjects** more quickly.

| **Page** | **Description** |
| --- | --- |
| [Introduction to Light Probes](LightProbes.md) | Learn about using Light Probes to store the light passing through specific points in a scene. |
| [Light Probes and moving GameObjects](LightProbes-MovingObjects.md) | Understand how dynamic GameObjects sample the light from Light Probes. |
| [Place Light Probes with the Editor](class-LightProbeGroup.md) | Choose where to place Light Probes, and choose the right amount of probes if you use **Enlighten** Realtime **Global Illumination**. |
| [Place Light Probes with a script](LightProbes-Placing-Scripting.md) | An example of forming a ring of Light Probes with a script. |
| [Set a GameObject to use light from Light Probes](LightProbes-MeshRenderer.md) | Use a **Mesh** Renderer component to set a GameObject correctly to receive light from Light Probes. |
| [Load Light Probes in multiple scenes](light-probes-and-scene-loading.md) | Use a script to control when Unity updates Light Probes if you load multiple scenes. |
| [Move Light Probes at runtime](LightProbes-Moving.md) | Use the `LightProbes` API to move Light Probes, for example if you create a world by loading multiple scenes additively and moving each scene to a different position. |
| [Troubleshooting Light Probes](light-probes-troubleshooting.md) | Solve common issues with Light Probes, such as light bleeding and ringing. |
| [Troubleshooting objects appearing unlit by Light Probes](ts-objects-missing-lighting.md) | Troubleshoot issues causing objects to appear unlit by Light Probes. |
| [Light Probes reference](LightProbes-Reference.md) | Explore the properties and settings in the Light Probe component **Inspector** window reference and the Edit **Light Probe Group** tool. |

# Additional resources

* [Adaptive Probe Volumes (APV) in URP](urp/probevolumes.md)
* [Configure an object to sample more light probes in the Built-In Render Pipeline](LightProbeProxyVolume-landing.md)