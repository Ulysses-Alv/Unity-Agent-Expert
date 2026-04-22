# Example: Create a root output

This example walks through setting up a simple sine‑wave root output and attaching it to the main audio output.

First, implement `RootOutputInstance.IRealtime`. For better scalability, move intensive calculations to a job that is kicked off in the processing stage. Use a temporary `NativeArray<float>` as an intermediate buffer. In the end stage, fence on the job and copy the result to the output.

```
[BurstCompile(CompileSynchronously = true)]
struct Realtime : RootOutputInstance.IRealtime
{
    internal NativeArray<float> phase; // Normalized phase accumulator in the range [0,1).
    internal float phaseIncrement; // Precomputed phase step.
    internal AudioFormat format; // Captured format from Configure.
    internal NativeArray<float> nativeBuffer; // Native backing buffer.
    JobHandle m_Job;

    struct Job : IJob
    {
        const float k_TwoPi = 2.0f * Mathf.PI;

        public NativeArray<float> phase;
        public float phaseIncrement;
        public AudioFormat format;
        public NativeArray<float> nativeBuffer;

        public void Execute()
        {
            // Use `nativeBuffer` as a backing buffer for the temporary channel buffer.
            var buffer = new ChannelBuffer(nativeBuffer, format.channelCount);

            for (var frame = 0; frame < buffer.frameCount; frame++)
            {
                var s = Mathf.Sin(phase[0] * k_TwoPi);

                for (var channel = 0; channel < buffer.channelCount; channel++)
                {
                    buffer[channel, frame] = s;
                }

                phase[0] += phaseIncrement;

                if (phase[0] >= 1.0f) phase[0] -= 1f;
            }
        }
    }

    public void Update(UpdatedDataContext context, Pipe pipe) { }

    public JobHandle EarlyProcessing(in RealtimeContext context, Pipe pipe) { return default; }

    public void Process(in RealtimeContext context, Pipe pipe, JobHandle input)
    {
        m_Job = new Job
        {
            phase = phase,
            phaseIncrement = phaseIncrement,
            format = format,
            nativeBuffer = nativeBuffer
        }.Schedule(input);
    }

    public void EndProcessing(in RealtimeContext context, Pipe pipe, ChannelBuffer output)
    {
        // Wait for the job to finish.
        m_Job.Complete();

        // Copy from the temp buffer to the output buffer.
        var buffer = new ChannelBuffer(nativeBuffer, format.channelCount);

        // Assumes format/channel layout matches `output`. If not, convert/mix here.
        for (var frame = 0; frame < output.frameCount; frame++)
        {
            for (var channel = 0; channel < output.channelCount; channel++)
            {
                output[channel, frame] = buffer[channel, frame];
            }
        }
    }

    public void RemovedFromProcessing()
    {
        // We'll dispose `nativeBuffer` in `Control.Dispose` or when we reconfigure.
    }
}
```

Next, implement `RootOutputInstance.IControl<Realtime>` to configure the sample on the `Realtime` part. You can manage the lifetime of the `NativeArray` from the control part. This ensures that it is properly allocated when configuring and disposed of when no longer needed, preventing memory leaks and ensuring efficient resource usage.

```
struct Control : RootOutputInstance.IControl<Realtime>
{
    const float k_Frequency = 440.0f;

    public void Dispose(ControlContext context, ref Realtime realtime)
    {
        realtime.phase.Dispose();
        realtime.nativeBuffer.Dispose();
    }

    public void Update(ControlContext context, Pipe pipe) { }

    public Response OnMessage(ControlContext context, Pipe pipe, Message message)
    {
        return Response.Unhandled;
    }

    public JobHandle Configure(ControlContext context, ref Realtime realtime, in AudioFormat format)
    {
        realtime.format = format;
        realtime.phaseIncrement = k_Frequency / format.sampleRate;

        // (Re)allocate the temp buffer.

        if (realtime.nativeBuffer.IsCreated)
        {
            realtime.phase.Dispose();
            realtime.nativeBuffer.Dispose();
        }

        realtime.phase = new NativeArray<float>(1, Allocator.Persistent);
        realtime.nativeBuffer = new NativeArray<float>(format.bufferFrameCount * format.channelCount, Allocator.Persistent);

        return default;
    }
}
```

Finally, add a `MonoBehaviour` to handle allocation in `Start` and cleanup in `OnDestroy`.

```
public class Driver : MonoBehaviour
{
    RootOutputInstance m_RootOutputInstance;

    void Start()
    {
        // Allocate the root output and attach it to the main audio output.
        m_RootOutputInstance = ControlContext.builtIn.AllocateRootOutput(new Realtime(), new Control());
    }

    void OnDestroy()
    {
        // Detach and destroy the root output.
        ControlContext.builtIn.Destroy(m_RootOutputInstance);
    }
}
```