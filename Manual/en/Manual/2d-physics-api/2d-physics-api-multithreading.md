# Optimize the LowLevelPhysics2D API with multithreading

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To improve the performance of the `LowLevelPhysics2D` API, use your available CPU cores to run the physics simulation and your own 2D physics code on multiple threads.

## Run the simulation on multiple threads

To increase the number of worker threads the 2D physics system uses to calculate the simulation, set the **Simulation Workers** property after you [create a world](2d-physics-api-world.md). The number of CPU cores on your device might mean the system uses fewer threads than you request.

To avoid performance issues, avoid changing the number of simulation workers at runtime.

## Run your own code on multiple threads

To run `LowLevelPhysics2D` API methods on all available CPU cores, use the [job system](../job-system.md) to split the calculations across worker threads. For example, use multiple threads to check for overlaps between a very large number of physics objects.

Most of the API is thread-safe, which means multiple worker threads can read and write the same physics data without interfering with each other. Inside a physics world, any number of threads can read data simultaneously, but only one thread can write data at a time. This is called a Write Once, Read Many (WORM) locking mechanism.

Each physics world is isolated from the others, so you can write data to different physics worlds on separate threads simultaneously.

The following methods aren’t guaranteed to be thread-safe, so you can’t use them in the job system:

* `Create`, for example `PhysicsBody.Create`
* `CreateBatch`, for example `PhysicsBody.CreateBatch`
* `Destroy`, for example `PhysicsBody.Destroy`
* `DestroyBatch`, for example `PhysicsBody.DestroyBatch`

For more information about creating a multithreaded job, refer to [Create and run a job](../job-system-creating-jobs.md).

## Additional resources

* [Write multithreaded code with the job system](../job-system.md)
* [`PhysicsWorld.concurrentSimulations`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld-concurrentSimulations.md)
* [Physics Low Level Settings 2D asset Inspector window reference](../class-PhysicsLowLevelSettings2D.md)