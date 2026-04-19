# Scriptable processors concepts

A scriptable processor lets you add your own audio processing to Unity at specific integration points in the audio system to customize its behavior. Currently, Unity supports two integration points, each with its own processor type:

| **Scriptable processor type** | **Integration point** | **Function** | **Typical usage** |
| --- | --- | --- | --- |
| [Generator](audio-scriptable-processors-generators.md) | **Audio sources** | Produces audio samples | Custom sound generators that emit audio into a **scene**, typically via an `AudioSource`. |
| [Root output](audio-scriptable-processors-root-outputs.md) | Audio engine main output | Produces audio samples | Custom audio systems with job system access, integration of middlewares, or sound generators that emit audio directly to the main output mix. |

## General architecture

All scriptable processors share the same high‑level structure. Each processor is split into the following two parts:

* **Control part**: Creates, configures, and manages the processor from game logic and typically runs on the main thread.
* **Real‑time part**: Generates and processes audio samples. It typically runs on the audio thread or in the Job System.

Each part represents a separate interface: `Processor.IControl` for control logic and `Processor.IRealtime` for real‑time processing. You must implement both interfaces to create a working processor. Note that each processor type provides specific interfaces (derived from `Processor.IControl`/`Processor.IRealtime`). For instance, when implementing a generator, use `GeneratorInstance.IControl` and `GeneratorInstance.IRealtime`.

## Control and real‑time contexts

The scriptable audio processor framework is controlled through two contexts:

* **Control context**: Lets you allocate and destroy processor instances, and interact with them from the game logic. The instance itself acts as a handle, and you access its capabilities through the control context. Typically, you use the built‑in control context, which connects and renders directly to the Unity audio system. You can also create a manual control context for offline rendering. Manual control contexts give you full control over allocation, rendering, and disposal, which is effective for offline rendering, audio analysis, visualization, and automated tests.
* **Real‑time context** : The real-time counterpart to a control context. Use it to perform real‑time operations on a processor instance, such as generating or processing audio samples during a mix cycle. When using the built‑in control context, Unity creates and destroys real‑time contexts automatically. If you use a manual control context, you’re responsible for managing the lifecycle of real‑time contexts yourself.

## Mix cycles

A real‑time context always belongs to a specific mix cycle. When you use the built‑in control context, Unity’s audio system manages mix cycles for you. However, if you use a manual control context, you are responsible for starting and ending each cycle. To manage a mix cycle, perform the following steps:

1. Start a new mix cycle by calling `ControlContext.Manual.BeginMix()`, which returns a temporary `RealtimeContext`.
2. Use the `RealtimeContext` to perform real‑time functionality on processor instances.
3. End the mix cycle by calling `ControlContext.Manual.EndMix()`.

Following is a code example:

```
var realtimeContext = ControlContext.Manual.BeginMix();

foreach (var processor in processors)
{
    // Use realtimeContext to drive real-time functionality on the processor.
}

ControlContext.Manual.EndMix();
```

## Communication mechanisms

Processors use two main mechanisms for data communication: pipes and messages.

### Pipes

Pipes are the primary communication channel between a processor’s control part and real‑time part. Pipes are efficient, real‑time safe (no locks or blocking system calls), and don’t allocate memory dynamically. They carry arbitrary byte‑serializable payloads, such as small structs with parameter changes, or pointers to managed data.

Both parts (control/real-time) share a pipe for data exchange, accessible from the `IRealtime`/`IControl` callbacks. Data can flow in both directions. All data sent from the control part within the same frame arrives through the shared pipe in the real‑time part in the same mix cycle. Similarly, all data sent from the real‑time part within the same mix-cycle is visible to the control part in the same frame. Data sent to the real‑time part is available for only one mix cycle, and data sent back from the real‑time part is available on the control part for only one update frame.

### Messages

Messages provide a structured and efficient way to send commands, events, or notifications from a processor instance to its control part. Messages are blocking and passed by reference. Although messages are primarily one‑way (processor instance to control part), the control part can modify message fields before returning them, to answer a query, for example. Use messages for tasks like parameter updates, data conversion (typically from managed Unity Objects to a lower, unmanaged representation), event notifications, or gameplay‑driven triggers that require coordination between game logic and real-time audio processing.

## Updates

Update hooks let a processor respond to control‑side changes, external input, and time‑based behavior. There are two update surfaces:

* **Control side** — typically called once per rendered video frame. Use it to:
  + Advance game logic related to the processor.
  + Read editor/UI parameters.
  + Send/receive data to/from the real‑time part (via the pipe).
  + Poll external sources (for example, timeline state or gameplay triggers).
  + Schedule one‑shot events (for example, retrigger an envelope or switch clips).
* **Real‑time side** — called once per mix cycle. Use it to:
  + Update audio logic and real-time‑dependent state.
  + Send/receive data to/from the control part (via the pipe).
  + Prepare state needed by the upcoming `Process` call(s).

You can throttle the update frequency on each side by setting `controlUpdateSetting` and `realtimeUpdateSetting` in the `CreationParameters` used when creating the processor. The following flags are supported:

* `Default` — Equivalent to `UpdateIfDataIsAvailable`.
* `NeverUpdate` — The update hook is disabled; it never runs.
* `UpdateIfDataIsAvailable` — The hook runs only when new data is available on the pipe.
* `UpdateAlways` — The hook runs every frame (control) or every mix cycle (real‑time), regardless of data availability.

## Channel buffers and audio formats

A channel buffer (`ChannelBuffer`) holds audio sample data as it moves through the audio processing pipeline. It provides a two‑dimensional view, channels × frames, over an underlying block of samples, and is agnostic to the sample layout (planar or interleaved). This allows processors to work with audio data efficiently, regardless of how the samples are organized in memory. The backing buffer is transient and can be based on native/unmanaged memory, either persistent or temporarily allocated using the `stackalloc` construct in C# - or via a managed allocation.

Channel buffers are typically used together with an audio format (`AudioFormat`). An audio format describes the properties of an audio stream, such as the sample rate, speaker configuration, and buffer size. When creating nested processors manually, you typically supply the audio format. Otherwise, it inherits the common audio format for the system.