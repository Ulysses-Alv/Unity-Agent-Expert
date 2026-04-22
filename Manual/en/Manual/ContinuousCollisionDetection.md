# Continuous collision detection (CCD)

Continuous **collision** detection (CCD) modes use predictive algorithms to calculate collisions that happen between physics timesteps. They are more accurate, but usually require more computational resources than discrete **collision detection**.

CCD is supported for Box, Sphere and Capsule **colliders**. It is intended as a safety net to catch collisions in cases where colliders would otherwise pass through each other. However, it does not always deliver physically accurate collision results, so you might still consider decreasing the physics timestep frequency to make the simulation more precise.

In Unity, there are two CCD algorithms, represented by three **Collision Detection** mode options.

| **Topic** | **Description** |
| --- | --- |
| [Speculative CCD](speculative-ccd.md) | Learn about speculative collision detection. **Continuous Speculative** uses speculative collision detection. |
| [Sweep-based CCD](sweep-based-ccd.md) | Learn about sweep-based collision detection. Both **Continuous** and **Continuous Dynamic** modes use sweep-based collision detection. |