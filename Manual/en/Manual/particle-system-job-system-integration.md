# Optimize the Particle System with the C# Job System

A **Particle** System can use Unity’s [C# Job System](job-system.md) to apply custom behaviors to particles.

Unity distributes work from the C# Job System across worker threads, and can make use of the Burst Compiler. The [GetParticles()](../ScriptReference/ParticleSystem.GetParticles.md) and [SetParticles()](../ScriptReference/ParticleSystem.SetParticles.md) methods offer similar functionality, but run on the main thread and cannot make use of Unity’s Burst Compiler.

By default, a **Particle System** job only has access to one or more particles belonging to that Particle System. Unity passes this data to the job using a [ParticleSystemJobData](../ScriptReference/ParticleSystemJobs.ParticleSystemJobData.md) struct. You must pass any other data that the job requires as additional parameters.

To access particle data, Unity supports the following job types:

## [IJobParticleSystem](../ScriptReference/ParticleSystemJobs.IJobParticleSystem.md)

This job type executes a single job on a single worker thread. The job has access to every particle belonging to the Particle System. For example code on this job type, see the [IJobParticleSystem.Execute()](../ScriptReference/ParticleSystemJobs.IJobParticleSystem.Execute.md) Scripting reference.

## [IJobParticleSystemParallelFor](../ScriptReference/ParticleSystemJobs.IJobParticleSystemParallelFor.md)

This job type executes multiple jobs across multiple worker threads. Each job can only access the particle at the index specified by the job’s Execute() function. For example code on this job type, see the [IJobParticleSystemParallelFor.Execute()](../ScriptReference/ParticleSystemJobs.IJobParticleSystemParallelFor.Execute.md) Scripting reference.

## [IJobParticleSystemParallelForBatch](../ScriptReference/ParticleSystemJobs.IJobParticleSystemParallelForBatch.md)

This job type executes multiple jobs across multiple worker threads. Each job can only access the particles within the range specified by the job’s Execute() function. For example code on this job type, see the [IJobParticleSystemParallelForBatch.Execute()](../ScriptReference/ParticleSystemJobs.IJobParticleSystemParallelForBatch.Execute.md) Scripting reference.

## Burst

As with any other C# job, you can use the Burst Compiler to compile your particle jobs into highly optimized Burst jobs. For more information, see the [Burst Compiler documentation](https://docs.unity3d.com/Packages/com.unity.burst@latest/index.html).

New feature in Unity 2019.3