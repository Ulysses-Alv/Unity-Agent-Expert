# Simulate hitches for testing

Introducing deliberate hitches is a common technique to test how your game handles performance issues like lag spikes, frame drops, stuttering, or degraded UI responsiveness under stress.

The following is a simple MonoBehaviour script that you can attach to any **GameObject** in your **scene**. It causes the main thread to hang at regular intervals, simulating a hitch or spike in frame time. You can easily adjust the frequency and duration of the hitches for your testing.

```
using UnityEngine;
using Unity.Profiling;

public class HitchSimulator : MonoBehaviour
{
    static readonly ProfilerMarker k_HitchSimulatorMarker = new ProfilerMarker("[HitchSimulator] CauseHitch() introducing artificial hitches");
    [Tooltip("How often (in seconds) to cause a hitch.")]
    public float hitchInterval = 5f;

    [Tooltip("How long (in milliseconds) each hitch should last.")]
    public int hitchDurationMs = 200;

    private float _nextHitchTime = 0f;

    void Update()
    {
        if (Time.time >= _nextHitchTime)
        {
            Debug.LogWarning($"[HitchSimulator] Causing a hitch for {hitchDurationMs} ms at {Time.time:F2}s");
            CauseHitch(hitchDurationMs);
            _nextHitchTime = Time.time + hitchInterval;
        }
    }

    void CauseHitch(int durationMs)
    {
        // Adding a clarifying profiler marker so no one accidentally confuses this for an actual issue.
        k_HitchSimulatorMarker.Auto();
        float start = Time.realtimeSinceStartup;
        // Busy wait for the specified duration
        while ((Time.realtimeSinceStartup - start) < durationMs / 1000f)
        {
            // Just spin
        }
    }
}
```

To use the script:

1. Copy the script into a new C# file called `HitchSimulator.cs`.
2. Attach it to any active GameObject in your scene.
3. Set parameters in the **Inspector**:
   * `hitchInterval`: How often to hitch. Default is 5 seconds.
   * `hitchDurationMs`: How long each hitch lasts. Default is 200 millisconds.
4. Enter Play mode and observe the deliberate frame hitches in the Profiler, or visually as stutter.

**Note**: This script blocks the main thread. It simulates a real lag spike from heavy computation, not just an artificial slowdown. This can make the Editor unresponsive if you use very large values. You can also trigger hitches via keypresses or random intervals for more robust testing.

## Additional resources

* [Unity Profiler](profiler-overview): Use the Profiler to monitor how your game behaves during hitches.
* [Performance Optimization](performance-optimization): General tips for optimizing performance in Unity.