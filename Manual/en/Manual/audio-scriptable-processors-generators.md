# Generators

A generator is a custom sound source that produces audio and injects it directly into a **scene**, typically through an `AudioSource`. Generators are a core extension point of the scriptable audio pipeline and let you synthesize or stream real-time audio.

You can attach a generator to an `AudioSource` in the following two ways: Asset-based workflow and component-based workflow.

Both workflows rely on implementing the `IAudioGenerator` interface, which acts as a factory for creating instances of your generator.

## Asset-based Workflow

Create a `ScriptableObject` that implements `IAudioGenerator`.

```
using UnityEngine;

[CreateAssetMenu(fileName = "MyAssetBasedGenerator", menuName = "Audio/Generators/My Asset-Based Generator", order = 1)]
public class MyAssetBasedGenerator : ScriptableObject, IAudioGenerator
{
    public bool isFinite => ...;
    public bool isRealtime => ...;
    public DiscreteTime? length => ...;

    public GeneratorInstance CreateInstance(
        ControlContext context,
        AudioFormat? nestedConfiguration,
        ProcessorInstance.CreationParameters creationParameters)
    {
        return context.AllocateGenerator(new MyRealtime(...), new MyControl(...));
    }
}
```

After you’ve created the asset, you can assign it to an `AudioSource` by dragging and dropping to the Generator field in the **Inspector**.

## Component-based workflow

To use the Component-based workflow, define a `MonoBehaviour` that implements `IAudioGenerator` and then add it to a **GameObject** in your scene. In the Inspector for your `AudioSource`, assign this component to the Generator field.

```
using UnityEngine;

public class MyComponentBasedGenerator : MonoBehaviour, IAudioGenerator
{
    public bool isFinite => ...;
    public bool isRealtime => ...;
    public DiscreteTime? length => ...;

    public GeneratorInstance CreateInstance(
        ControlContext context,
        AudioFormat? nestedConfiguration,
        ProcessorInstance.CreationParameters creationParameters)
    {
        return context.AllocateGenerator(new MyRealtime(...), new MyControl(...));
    }
}
```

## The `IAudioGenerator` interface

An `IAudioGenerator` includes the following:

* **Capabilities:** Properties such as `GeneratorInstance.ICapabilities` that describe how the generator behaves in advance.
* **Factory method:** The `CreateInstance` method, which the audio system calls to create a `GeneratorInstance`.

Following is a list of the capabilities:

* **Finite**: When true, indicates the generator will eventually terminate. The total length does not need to be known up front.
* **Real-time**: When true, indicates the generator must render in real time at the system’s buffer size and sample rate. For example, hardware-driven streams, graphs that cannot run ahead, or anything with an internal timeline that is synchronized against the system dsp clock. If unsure, set this to `false`.
* **Length**: If the generator is finite and you know the total length, report it here. While optional, exposing this property is valuable for editor tooling and scheduling.

You must expose matching capabilities in your `GeneratorInstance.IRealtime` implementation. A mismatch between `IAudioGenerator` and `IRealtime` may cause unexpected behavior and will produce warnings in the Console.

## Configuring a generator

Generator configuration happens initially during construction, and additionally when the system changes configuration. This typically happens in response to a user action, such as changing the audio output device in the OS system settings or when connecting a pair of headphones. It is a two-step negotiation between the host (the owner of the generator) and the generator:

1. The host suggests a preferred `AudioFormat` for the generator (sample rate, speaker layout, buffer size).
2. The generator reports the `Setup` it will actually use.

Some generators (e.g. when synthesizing audio) can adapt to many formats, while others (e.g. when playing back audio files) may only support the file’s native format. If a generator selects a different format than suggested by the host, it is the host’s responsibility to convert between formats when needed.

When creating child generators via `ControlContext.AllocateGenerator`, a parent can provide an optional `AudioFormat` as a suggestion to the child. Children should follow the suggestion when possible to minimize conversions.

## Processing a generator

Generators are always processed within a mix cycle. For a description of mix cycles see the mix cycles subsection in the [Scriptable processors concepts](audio-scriptable-processors-concepts.md) section.

## Nested generators

Generators can be organized into hierarchical trees. In each tree, the root generator sits at the top and is responsible for invoking `Configure`, `Update`, and `Process` on its immediate child generators. Each child, in turn, forwards those calls to its own children. This pattern enables complex structures such as mixers, blend containers, and randomized sequencers.

When a parent generator creates a child using `ControlContext.AllocateGenerator`, it might pass an optional suggested `AudioFormat`. The child generator should match the suggestion if possible to minimize format conversions, but can select a different format, if required.

## Best practices

Consider the following:

* Keep `Process` real‑time safe. Avoid allocations, locks, blocking I/O, logging, and throwing exceptions on the audio thread. Use stack-allocated or pooled buffers, and struct-based states for predictable performance.
* Beware of threading. Never call `UnityEngine` APIs from within `Process`. Use pipes to synchronize state between the control and real-time part, and avoid any multithreaded communication that might yield indeterministic results.
* Use Burst-compiled, struct-based code. Use simple value types and tight loops for improved performance. Leverage `Unity.Mathematics` and organize data for optimal Burst auto-vectorization.
* Match capabilities consistently. Ensure that the values for `isFinite`, `isRealtime`, and `length` in your `IAudioGenerator` implementation match with the values in `GeneratorInstance.IRealtime`.
* Use the host format when you can. Adopting the host’s suggested `AudioFormat` minimizes additional resampling or channel mixing.
* Report the length when known. Providing an accurate `length` property enables better waveform previews, scheduling, and progress indicators in the Editor.
* Handle completion gracefully. When a finite generator reaches the end, return silence and signal completion so hosts can deallocate or transition.
* Validate externally received inputs. Clamp out-of-range parameters to valid ranges and handle mismatches in channel counts to prevent errors.

## Additional resources:

* [Scriptable processors concepts](audio-scriptable-processors-concepts.md)
* [Example: Creating a simple generator](audio-scriptable-processors-example-creating-a-generator.md)