# Write and run coroutines

A coroutine is a method that can suspend execution and resume at a later time. In Unity applications, this means coroutines can start running in one frame and then resume in another, allowing you to spread tasks across several frames.

Regular, non-coroutine methods run to completion before returning control to the caller, which in the Unity runtime means their action completes within a single frame update. In situations where you want the work of a method to take effect over several frames, such as a gradual fade-out effect, you can use a coroutine. Coroutines are also useful for handling long asynchronous operations, such as waiting for HTTP transfers, asset loads, or file I/O to complete.

**Important**: Don’t confuse coroutines with threads. Synchronous operations that run within a coroutine still execute on the main thread. If you want to reduce the amount of CPU time spent on the main thread, it’s just as important to avoid blocking operations in coroutines as in any other script code. If you want to use multi-threaded code in Unity, your options are:

* The [job system](job-system.md)
* The .NET [async and await](async-await-support.md) and Unity’s custom `Awaitable` support

## Writing coroutines

Consider the task of gradually reducing an object’s alpha (opacity) value until it becomes invisible. For the fading effect to be visible, the opacity must reduce over a sequence of frames. If you tried to write a `Fade` method, you might write something like the following:

```
void Fade()
{
    Color c = renderer.material.color;
    for (float alpha = 1f; alpha >= 0; alpha -= 0.1f)
    {
        c.a = alpha;
        renderer.material.color = c;
    }
}
```

This method is not a coroutine, so it executes every iteration of its `for` loop within a single frame update and the object disappears instantly instead of appearing to fade out. One posible solution is to add code to the `Update` function that executes the fade on a frame-by-frame basis. However, it can be more convenient to use a coroutine.

Coroutines are methods with an [IEnumerator](https://docs.microsoft.com/en-us/dotnet/api/system.collections.ienumerator) return type and a [yield](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/yield) return statement included somewhere in the body. The `yield return` statement is the point at which execution is suspended. The previous `Fade` method can be rewritten as a coroutine as follows:

```
IEnumerator Fade()
{
    Color c = renderer.material.color;
    for (float alpha = 1f; alpha >= 0; alpha -= 0.1f)
    {
        c.a = alpha;
        renderer.material.color = c;
        yield return null;
    }
}
```

This version of the method executes one iteration of its `for` loop before suspending execution at the `yield return null` statement. It resumes and executes another iteration of the loop in the next frame, and so on, making the gradual fade effect visible. The loop counter in the `Fade` method maintains its correct value over the lifetime of the coroutine, and any variable or parameter is preserved between `yield` statements.

## Starting and stopping coroutines

To set a coroutine running, use the [StartCoroutine](../ScriptReference/MonoBehaviour.StartCoroutine.md) method:

```
void Update()
{
    if (Input.GetKeyDown("f"))
    {
        StartCoroutine(Fade());
    }
}
```

To stop a coroutine, use [StopCoroutine](../ScriptReference/MonoBehaviour.StopCoroutine.md) and [StopAllCoroutines](../ScriptReference/MonoBehaviour.StopAllCoroutines.md). A coroutine also stops if:

* The value of [`GameObject.activeSelf`](../ScriptReference/GameObject-activeSelf.md) becomes `false` for the **GameObject** the script is attached to.
* The MonoBehaviour script is destroyed with a call to [Destroy](../ScriptReference/Object.Destroy.md).

**Note:** Disabling the MonoBehaviour script by setting [enabled](../ScriptReference/Behaviour-enabled.md) to `false` doesn’t stop coroutines.

## Resuming coroutines

When a suspended coroutine resumes execution depends on the yield instruction provided in the `yield return` statement. A `yield return null` resumes on the next frame. Unity has a set of custom yield instructions that you can use to resume after a specified time, when a specified conditions is met, or at specific points in the Player loop. For more information, refer to [Yield instruction reference](coroutines-yield-instructions.md).

In the case of fade effect example, you might want the fade effect to happen at a lower and more consistent rate than the frame rate. You can `yield return` the [`WaitForSeconds`](../ScriptReference/WaitForSeconds.md) instruction to introduce a fixed time delay between iterations of the `Fade` method as follows:

```
IEnumerator Fade()
{
    Color c = renderer.material.color;
    for (float alpha = 1f; alpha >= 0; alpha -= 0.1f)
    {
        c.a = alpha;
        renderer.material.color = c;
        // Wait for 0.1 seconds before the next iteration
        yield return new WaitForSeconds(.1f);
    }
}
```

It’s also possible to `yield return` a Unity [`Awaitable`](../ScriptReference/Awaitable.md) from within a coroutine. This can be useful if you want to integrate coroutines with asynchronous code that uses `async` and `await`. For example, in the previous example you could `yield return Awaitable.WaitForSecondsAsync(.1f)` instead of `yield return new WaitForSeconds(.1f)` to achieve the same effect.

**Important**: It’s not supported to `yield return` the generic [`Awaitable<T0>`](../ScriptReference/Awaitable_1.md) from a coroutine.

## Coroutines in Edit mode

Coroutines are primarily a runtime feature. The associated [runtime yield instructions](coroutines-yield-instructions.md) are in the `UnityEngine` namespace and run in the Editor’s Play mode or in a standalone platform Player. They can also run in Edit mode if your **scripts** use the [`[ExecuteInEditMode]`](../ScriptReference/ExecuteInEditMode.md) or [`[ExecuteAlways]`](../ScriptReference/ExecuteAlways.md) attributes, but the update loop in Edit mode is not as fixed and regular as the Player loop.

For coroutines designed specifically to run in Edit mode, use the [Editor coroutines package](https://docs.unity3d.com/Packages/com.unity.editorcoroutines@latest).

## Coroutines in tests

Unity Test Framework Play mode tests marked with the `[UnityTest]` attribute run as coroutines and allow you to yield custom instructions for the Unity Editor from tests. For more information, refer to [Yield instructions for the Editor](test-framework/reference-custom-yield-instructions.md).

## Coroutine performance

Coroutines can cause hidden allocations and garbage collector spikes if misused. Each coroutine creates an `IEnumerator` **state machine**. Starting them frequently (for example, per frame) allocates and adds overhead. A `yield return null` does not allocate but yield instructions like `new WaitForSeconds` do. Cache commonly reused ones and avoid lambdas in [`WaitUntil`](../ScriptReference/WaitUntil.md) and [`WaitWhile`](../ScriptReference/WaitWhile.md) to prevent delegate and capture allocations.

Prefer long-lived coroutines that loop with `yield return null` instead of repeatedly starting new ones. Cache or pool [`WaitForSeconds`](../ScriptReference/WaitForSeconds.md) with fixed durations. Coroutines retain references to their owner and captured variables. Ensure they end or are stopped with [`MonoBehaviour.StopCoroutine`](../ScriptReference/MonoBehaviour.StopCoroutine.md) to avoid leaks.

Always profile, especially on constrained platforms, to confirm and locate allocations. For more information, refer to [Analyzing coroutines](coroutines-analyzing.md).

## Additional resources

* [Coroutine API reference](../ScriptReference/Coroutine.md)
* [MonoBehaviour.StartCoroutine](../ScriptReference/MonoBehaviour.StartCoroutine.md)