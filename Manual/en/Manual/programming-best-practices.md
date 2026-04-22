# Unity programming best practices

There are some unique features of Unity’s programming environment that require extra consideration when writing code compared to standard C#/.NET projects. The following is a summary of key issues to be aware of when writing code for Unity applications, along with best practices to help you avoid common pitfalls.

## Unity Object lifecycle and references

When writing C# in Unity, be careful when comparing objects for equality with other objects or with null. For types that inherit from [UnityEngine.Object](class-Object.md), Unity uses a custom version of the C# [equality and inequality operators](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/operators/equality-operators). This means the null check `myGameObject == null` can evaluate `true` (and conversely `myGameObject != null` can evaluate `false`) even if `myGameObject` technically holds a valid C# object reference. For more information on the specifics of this behavior, refer to [Custom equality operators](class-Object.md#custom-equality-comparers).

Unity’s custom equality behavior and object lifecycle have a few implications for your code:

* In circumstances where you want to be sure of excluding destroyed objects in your check, make sure to use `if (obj == null)` and not `ReferenceEquals` for Unity objects.
* In circumstances where you want to check for actual C# null references, use `ReferenceEquals` or cast to `System.Object` first.
* When comparing two Unity objects for equality, be aware that `obj1 == obj2` may return `true` even if both references are different C# objects, if one or both have been destroyed and recreated, for example through [`Undo`](../ScriptReference/Undo.md).
* The custom equality operator is slower than the standard C# one. This is usually not a problem, but be mindful of this at scale and in hot paths.
* Don’t cache components across **scene** unloads without guarding because they can be destroyed but remain as C# wrapper objects with no unmanaged counterpart.
* Don’t hold strong references to large assets in static fields because they persist across scenes and prevent unloading.
* When destroying objects, [`Object.DestroyImmediate`](../ScriptReference/Object.DestroyImmediate.md) is Editor-only so use [`Object.Destroy`](../ScriptReference/Object.Destroy.md) at runtime and let Unity schedule destruction.

## Avoid C# finalizers

Don’t use [C# finalizers](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/finalizers) in runtime code, for the following reasons:

* They run on separate finalizer threads and Unity APIs usually require the main thread.
* They run non-determinisitically, leading to unpredictable behavior.
* They might not run at all, unless the application garbage collects and waits.
* Exceptions thrown from finalizers can cause an application halt unless specially handled.
* By their very existence they increase garbage collector overhead.

## Garbage collector overhead and allocations

One of the most important performance risks to be mindful of in Unity applications is runtime code that allocates memory and increases garbage collector overhead, especially in hot paths.

To avoid this, apply the following coding practices:

* Avoid per-frame allocations by [caching and reusing lists](performance-reusable-code.md) and using non-allocating versions of methods where available. This includes caching [`WaitForSeconds`](../ScriptReference/WaitForSeconds.md) and other [yield instructions](../ScriptReference/YieldInstruction.md) when using [coroutines](coroutines.md).
* Where possible, use non-allocating versions of methods, and perform expensive operations like [`GameObject.GetComponent`](../ScriptReference/GameObject.GetComponent.md) in [`Awake`](../ScriptReference/MonoBehaviour.Awake.md) and cache references to returned objects, rather than calling repeatedly from [`Update`](../ScriptReference/MonoBehaviour.Update.md).
* Avoid use of [LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/linq/) in runtime code, and especially in the context of the per-frame `Update` or `FixedUpdate` and other hot paths. Methods from the `System.Linq` namespace can create unnecessary allocations and involve boxing and closures.
* Avoid [repeated string operations](performance-reference-types.md) like concatenation.
* Identify and avoid instances of reflection. For more information, refer to [Avoid C# reflection overhead](performance-gc-avoid-reflection.md).

For more detailed guidance and examples of these issues, refer to [Optimizing your code for managed memory](performance-optimizing-code-managed-memory.md).

For information on tracking and reducing garbage collector overhead, refer to [Managed memory](performance-managed-memory.md).

## MonoBehaviour Update loop optimization

The traditional pattern for many Unity projects involves using MonoBehaviour script components to regularly update the game state through built-in callbacks such as [`MonoBehaviour.Update`](../ScriptReference/MonoBehaviour.Update.md), [`MonoBehaviour.FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md), and [`MonoBehaviour.LateUpdate`](../ScriptReference/MonoBehaviour.LateUpdate.md) that typically run many times per second.

This is a simple model that can still work well when used appropriately, but it has some key performance risks that commonly catch inexperienced developers out:

* The default implementation of Unity’s per-frame or other regular event functions can scale poorly. Each of these `Update` functions incurs a small overhead from Unity’s internal management and interaction with the native layer. When you have many such MonoBehaviour **scripts**, the cumulative overhead can have significant performance impact.
* The fact that built-in updates run very often makes them a hot code path and magnifies the effect of any inefficient, memory-intensive operations you place in them. A common bad pattern from inexperienced users is to have many MonoBehaviour scripts with `Update` functions that run unnecessarily most of the time, or that are unnecessarily memory-intensive when they do run.

To mitigate these risks, consider the following options:

* Consider converting your project to a data-oriented architecture using Unity’s [Entity Component System (ECS)](https://docs.unity3d.com/Packages/com.unity.entities@latest) for better scalability with many entities.
* If you use a MonoBehaviour-based architecture:
  + To ensure you’re minimizing managed memory impact in your hot paths, refer to [Optimizing your code for managed memory](performance-optimizing-code-managed-memory.md).
  + Minimize the number of active `Update` functions by using a centralized update manager or customizing the Player loop. For more information, refer to [Using a custom update manager](events-per-frame-optimization.md) and [Customizing the Player loop](player-loop-customizing.md).
  + Bear in mind that even in a MonoBehaviour-based project, you can often use specific features of Unity’s data-oriented systems. You can use [Jobs](job-system.md), [Burst-compile](script-compilation-burst.md) sections of your code, use more efficient data structures like [`NativeArray`](../ScriptReference/Unity.Collections.NativeArray_1.md) in performance-critical sections of your code, and choose unmanaged alternatives to managed APIs such as [those for transform operations](transformhandle-landing.md).

## Thread safety

While Unity has multithreaded capabilities, the core runtime is single-threaded and most APIs in the `UnityEngine` and `UnityEditor` namespaces can only be called from the main thread. Don’t reference **GameObjects**, Transforms, Components, or asset APIs from background threads. Never `await` with a `Task.Result` or `Task.Wait` on the main thread as this leads to deadlocks.

When dealing with inherently asynchronous and long-running operations, Unity provides the [`Awaitable`](../ScriptReference/Awaitable.md) class as a Unity-specific alternative to .NET `Task`. `Awaitable` uses object pooling to reduce allocations and is aware of Unity-specific concepts like `Update` and `FixedUpdate`, which allows you to `await` tasks and schedule them to resume at specific points in the Player loop. For more information refer to [Asynchronous programming with the Awaitable class](async-await-support.md).

For shorter-lived but more computationally-intensive parallelized work, Unity provides the [job system](job-system.md), which can be [Burst compiled](script-compilation-burst.md). For more information, refer to [Write multithreaded code with the job system](job-system.md).

## Compilation considerations

For optimal performance it’s important to think not just about how you write code but how it’s compiled. Compiling code naively rather than actively defining which contexts certain source files or regions of your code are relevant for imposes the following costs:

* Increased build size if unnecessary code is included.
* Time spent compiling and recompiling to apply changes. This can especially affect your iteration time in the Editor.
* Potential runtime errors if inappropriate code is included for a given platform or context.

Unity provides several mechanisms to help you control which parts of your code are compiled for different platforms and contexts:

* You can use Assembly Definitions to group source files into assemblies that can be compiled separately. This allows you to isolate Editor-only code from runtime code, and platform-specific code from cross-platform code. For more information, refer to [Organizing scripts into assemblies](assembly-definition-files.md).
* You can use `#if` directives with [scripting symbols](scripting-symbol-reference.md) to exclude specific regions of code from compilation based on the target platform or context. For example, guard Editor-only code behind `#if UNITY_EDITOR` directives to exclude it from runtime builds. For more information on the various methods Unity offers for conditionally including or excluding code, refer to [Conditional compilation](platform-dependent-compilation.md).
* A key Unity-specific concept is [domain reload](configurable-enter-play-mode.md) in the Editor on entering and exiting Play mode or recompiling scripts. Domain reloads are time consuming and can affect iteration times when writing and testing in the Editor. You can disable domain reloads on entering Play mode to improve iteration times, but you must then reset static state manually. For more information, refer to [Code and scene reload on entering Play mode](code-reloading-editor.md).

## Unity’s analysis tools

Unity provides a variety of tools to help you identify bottlenecks and write more performant code. The [Project Auditor](project-auditor/project-auditor) tool can analyze your project code to identify common performance issues and suggest fixes. The [Profiler](Profiler.md) can help you identify runtime performance bottlenecks in your code by providing detailed information about CPU and GPU usage, memory allocation, and more. You can also create [Roslyn analyzers](https://learn.microsoft.com/en-us/visualstudio/code-quality/roslyn-analyzers-overview?view=visualstudio) to enforce coding standards and identify performance issues specific to your project.

For more information on creating custom Roslyn analyzers and source generators, refer to [Roslyn analyzers and source generators](roslyn-analyzers.md).

For more information on Unity’s suite of analysis tools, refer to [Optimization](analysis.md).

## Additional resources

* [Memory in Unity](performance-memory.md)
* [Optimizing your code for managed memory](performance-optimizing-code-managed-memory.md)
* [Debugging and diagnostics](debugging-and-diagnostics.md)