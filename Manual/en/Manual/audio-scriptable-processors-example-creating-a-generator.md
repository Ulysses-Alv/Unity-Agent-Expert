# Example: Create a generator

This example walks through setting up a sine tone generator and attaching it to an `AudioSource` using the component-based workflow. It’s organized into two parts:

* **Part I (Minimal):** Uses a fixed-frequency sine wave.
* **Part II (Parameterized):** Extends the minimal example with a frequency control.

## Part I — Minimal example

Implement `GeneratorInstance.IRealtime` to generate audio in `Process`. The real‑time struct holds the phase accumulator and precomputed increment.

```
[BurstCompile(CompileSynchronously = true)]
struct Realtime : GeneratorInstance.IRealtime
{
    private const float k_TwoPi = 2.0f * Mathf.PI;

    // Normalized phase accumulator in [0, 1).
    private float phase;

    // Precomputed phase step per sample (cycles per sample).
    internal float phaseIncrement;

    // Capabilities must match those reported from IAudioGenerator and IRealtime.
    public bool isFinite => false;
    public bool isRealtime => false;
    public DiscreteTime? length => null;

    // Called when the real-time side of the graph updates (e.g., new control data available).
    // Keep this method allocation-free and exception-free.
    public void Update(UpdatedDataContext context, Pipe pipe) { }

    public GeneratorInstance.Result Process(
        in RealtimeContext context, Pipe pipe,
        ChannelBuffer buffer,
        GeneratorInstance.Arguments args)
    {
        // Compute a mono sine and copy it to all channels.
        for (int frame = 0; frame < buffer.frameCount; frame++)
        {
            float s = Mathf.Sin(phase * k_TwoPi);

            for (int ch = 0; ch < buffer.channelCount; ch++)
                buffer[ch, frame] = s;

            // Advance and wrap the phase into [0, 1).
            phase += phaseIncrement;
            if (phase >= 1.0f) phase -= 1.0f;
        }

        // Return the number of frames written.
        return buffer.frameCount;
    }
}
```

Next, implement `GeneratorInstance.IControl<Realtime>`. Use `Configure` to set any real‑time fields that depend on the format.

```
struct Control : GeneratorInstance.IControl<Realtime>
{
    private const float k_Frequency = 440.0f; // A4

    // Dispose is called when the generator instance is destroyed.
    public void Dispose(ControlContext context, ref Realtime realtime) { }

    // Control-side tick; e.g., poll external state or schedule events.
    public void Update(ControlContext context, Pipe pipe) { }

    // Optional message hook; return `Unhandled` for messages you don't consume.
    public Response OnMessage(ControlContext context, Pipe pipe, Message message) => Response.Unhandled;

    // Called initially when constructed and additionally when the system changes configuration.
    public void Configure(
        ControlContext context,
        ref Realtime realtime,
        in AudioFormat format,
        out GeneratorInstance.Setup setup,
        ref GeneratorInstance.Properties properties)
    {
        // Configure real-time fields that depend on the audio format.
        realtime.phaseIncrement = k_Frequency / format.sampleRate; // cycles/sample

        // Prefer matching the host's format when possible to avoid conversion.
        setup = new GeneratorInstance.Setup(format.speakerMode, format.sampleRate);
    }
}
```

Finally, tie everything together in a `MonoBehaviour` that implements `IAudioGenerator`:

```
public class Driver : MonoBehaviour, IAudioGenerator
{
    public bool isFinite => false;
    public bool isRealtime => false;
    public DiscreteTime? length => null;

    public GeneratorInstance CreateInstance(
        ControlContext context,
        AudioFormat? nestedConfiguration,
        CreationParameters creationParameters)
    {
        // Allocate a new generator instance pairing the realtime and control structs.
        return context.AllocateGenerator(new Realtime(), new Control(), nestedConfiguration, creationParameters);
    }
}
```

Finally, add the `SineGeneratorDriver` component to a `GameObject`, assign it as the `Generator` in the `AudioSource`, and enter Play Mode to hear the sine tone.

## Part II — Configure the frequency

To parameterize the frequency of the sine generator, start by adding a `frequency` field to the driver. Then, send frequency updates to the real-time part using the pipe in the `Update` method, using a lightweight value type to represent frequency change messages:

```
// Small value-type message for the pipe.
readonly struct FrequencyEvent
{
    public readonly float value;
    public FrequencyEvent(float value) => this.value = value;
}
```

Secondly, update the real‑time struct to store `frequency` and `sampleRate`. In the `Update` method, read any pending `FrequencyEvent` message from the pipe, and update the frequency. This approach ensures the audio thread safely handles any number of pending events without throwing errors on the audio thread. The phase increment is then calculated per audio block based on the current frequency and sample rate.

```
[BurstCompile(CompileSynchronously = true)]
struct Realtime : GeneratorInstance.IRealtime
{
    private const float k_TwoPi = 2.0f * Mathf.PI;

    private float phase;       // [0, 1)
    internal float frequency;  // Hz, set from control messages
    internal float sampleRate; // Hz, set from Configure

    public bool isFinite => false;
    public bool isRealtime => false;
    public DiscreteTime? length => null;

    public void Update(UpdatedDataContext context, Pipe pipe)
    {
        // Iterate over all available events (newer overwrite older).
        foreach (var element in pipe.GetAvailableData(context))
        {
            if (element.TryGetData(out FrequencyEvent evt))
            {
                frequency = evt.value;
            }

            // Ignore other message types gracefully.
        }
    }

    public GeneratorInstance.Result Process(
        in RealtimeContext context,
        Pipe pipe,
        ChannelBuffer buffer,
        GeneratorInstance.Arguments args)
    {
        // Compute increment locally from current control values.
        float phaseIncrement = frequency / sampleRate;

        for (int frame = 0; frame < buffer.frameCount; frame++)
        {
            float s = Mathf.Sin(phase * k_TwoPi);

            for (int ch = 0; ch < buffer.channelCount; ch++)
                buffer[ch, frame] = s;

            phase += phaseIncrement;
            if (phase >= 1.0f) phase -= 1.0f;
        }

        return buffer.frameCount;
    }
}
```

In the control struct, capture the runtime `sampleRate` in `Configure` and set up a mono output. Intercept messages sent via the instance and forward them to the real-time part.

```
struct Control : GeneratorInstance.IControl<Realtime>
{
    public void Dispose(ControlContext context, ref Realtime realtime) { }

    public void Update(ControlContext context, Pipe pipe) { }

    public Response OnMessage(ControlContext context, Pipe pipe, Message message)
    {
        if (message.Is<FrequencyEvent>())
        {
            pipe.SendData(context, message.Get<FrequencyEvent>());

            return Response.Handled;
        }

        return Response.Unhandled;
    }

    public void Configure(
        ControlContext context,
        ref Realtime realtime,
        in AudioFormat format,
        out GeneratorInstance.Setup setup,
        ref GeneratorInstance.Properties properties)
    {
        realtime.sampleRate = format.sampleRate;

        setup = new GeneratorInstance.Setup(AudioSpeakerMode.Mono, format.sampleRate);
    }
}
```

Finally, add a frequency slider and update the audio instance only when the value changes to avoid spamming the control part.

**Note**: Always guard accesses to the instance. It may be missing or have been destroyed, if the **audio source** was stopped elsewhere.

```
public class Driver : MonoBehaviour, IAudioGenerator
{
    private AudioSource m_AudioSource;

    [Range(100f, 10000f)]
    public float frequency = 440.0f; // A4

    private float m_PreviousFrequency;

    public bool isFinite => false;
    public bool isRealtime => false;
    public DiscreteTime? length => null;

    public GeneratorInstance CreateInstance(
        ControlContext context,
        AudioFormat? nestedConfiguration,
        CreationParameters creationParameters)
        => context.AllocateGenerator(new Realtime(), new Control());

    private void Awake()
    {
        // Expects an AudioSource on the same GameObject.
        m_AudioSource = GetComponent<AudioSource>();
    }

    private void Update()
    {
        // Early out if unchanged (use Approximate to avoid redundant updates).
        if (Mathf.Approximately(frequency, m_PreviousFrequency))
            return;

        // Access the instance via the AudioSource.
        var instance = m_AudioSource.generatorInstance;

        // Guard the handle: instance may be missing or have been destroyed, if the audio source was stopped elsewhere.
        if (!ControlContext.builtIn.Exists(instance))
            return;

        var message = new FrequencyEvent(frequency);

        // Send frequency change to the control side.
        ControlContext.builtIn.SendMessage(instance, ref message);
        m_PreviousFrequency = frequency;
    }
}
```