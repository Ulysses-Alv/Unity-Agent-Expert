# Scriptable audio pipeline

The scriptable audio pipeline framework lets you extend the Unity audio engine at specific integration points using Burst-compatible HPC#. The below sections explain what the scriptable audio pipeline is and how you can use it to customize and extend the audio system.

| **Topic** | **Description** |
| --- | --- |
| **[Scriptable processors concepts](audio-scriptable-processors-concepts.md)** | Learn the core concepts required to work with scriptable processors. |
| **[Generators](audio-scriptable-processors-generators.md)** | Learn about generators. |
| **[Root outputs](audio-scriptable-processors-root-outputs.md)** | Learn about root outputs. |
| **[Example: create a generator](audio-scriptable-processors-example-creating-a-generator.md)** | Create a generator and attach it to an **audio source**. |
| **[Example: create a root output](audio-scriptable-processors-example-creating-a-root-output.md)** | Create a root output and connect it to the Unity audio engine. |

**Note**: The scriptable audio pipeline is not supported on the Web Platform. Attempting to use it in a Web build will trigger a warning.