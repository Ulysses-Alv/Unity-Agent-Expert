# Manually set physics simulation

Control the physics simulation manually to decide when physics calculations occur to align physics updates with your application’s actual performance and avoid redundant updates. This can help optimize performance during busy frames and is beneficial for scenarios that require tight synchronization with game logic or custom physics update loops.

Manually control the physics simulation to:

* Ensure stability and consistency: Physics systems are designed for small, consistent time steps. Variable time step values, such as [`Time.deltaTime`](../ScriptReference/Time-deltaTime.md), can cause unstable simulations, and lead to issues such as objects passing through each other, a phenomenon known as tunneling, at low frame rates or jittery behavior when frame rates fluctuate.
* Create reproducible behavior: Using a fixed step is crucial for more reproducible physics behavior (though true determinism with PhysX has other challenges).
* Avoid large, inaccurate steps: If frame rates drop, [`Time.deltaTime`](../ScriptReference/Time-deltaTime.md) increases. Simulating a physics step with a large, variable delta can be computationally resource intensive and less accurate, potentially worsening performance issues.

Use `Physics.Simulate(float stepTime)` (or `yourPhysicsScene.Simulate(float stepTime)`) to manually control the physics simulation and in specific, local `PhysicsScene` instances. Refer to the [Multi-Scene Physics learn tutorial](https://learn.unity.com/tutorial/multi-scene-physics) to learn more about multi-scene physics.

When you call [`Physics.Simulate`](../ScriptReference/Physics.Simulate.md), the recommended best practice is to use a fixed time step, rather than directly passing `Time.deltaTime`. For example, use the same value that you might use in the ****Project Settings****’ ****Fixed Timestep**** value, such as **0.02** for 50Hz.

If you manually simulate physics in a **scene** using `Physics.Simulate()`, ensure that automatic simulation is disabled. You can do this by setting the **Simulation Mode** to **Script**. Once in **Script** mode (either through **Project Settings** or script), you must call `Physics.Simulate` for the scene at a selected frequency with an appropriate, preferably fixed, time step.

To disable automatic simulation, in the Editor:

1. Select **Edit > Project Settings** to open the Project Settings window.
2. Select the **Physics > Settings** tab.
3. Select the ****GameObject**** tab.
4. Set **Simulation Mode** to **Script**.

To disable automatic simulation, in script, set `Physics.simulationMode = SimulationMode.Script;`. Typically, you do this at the start of your game, before any physics simulation occurs.

## Additional resources

* [Unity Profiler](Profiler.md)
* [Managing time and frame rate](managing-time-and-frame-rate.md)
* [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Physics Project Settings](class-PhysicsManager.md)
* [Time](class-TimeManager.md)
* [Optimize physics for query-only or non-simulating games](physics-optimization-cpu-query-only.md)
* [Physics.Simulate](../ScriptReference/Physics.Simulate.md)
* [Multi-Scene Physics learn tutorial](https://learn.unity.com/tutorial/multi-scene-physics)