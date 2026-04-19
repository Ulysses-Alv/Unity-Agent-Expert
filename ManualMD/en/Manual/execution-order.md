# Event function execution order

The following diagram provides a high-level overview of the execution sequence for event functions that run during the lifecycle of a MonoBehaviour script component. For readability, the scope of the chart is limited to key parts of the script lifecycle, with some extra internal subsystem updates provided for context. The full Player loop is a longer and more complex sequence of updates for specific systems and subsystems, which run one after the other in a defined default order. To retrieve the full Player loop and all of its systems, you can use the [PlayerLoop API](../ScriptReference/LowLevel.PlayerLoop.md). You can also use this API to [customize the Player loop](player-loop-customizing.md) sequence by removing systems, adding your own, and changing the update order.

For more information on each individual callback’s meaning and limitations, refer to the **Messages** section of the [MonoBehaviour](../ScriptReference/MonoBehaviour.md) API reference.

![Order of execution for event functions during the lifecycle of a MonoBehaviour script.](../uploads/Main/monobehaviour_flowchart.svg)

Order of execution for event functions during the lifecycle of a MonoBehaviour script.

## Before scene load and unload

Not shown in the previous diagram are the [`SceneManager.sceneLoaded`](../ScriptReference/SceneManagement.SceneManager-sceneLoaded.md) and [`SceneManager.sceneUnloaded`](../ScriptReference/SceneManagement.SceneManager-sceneUnloaded.md) events which allow you to receive callbacks when a **scene** has loaded and unloaded respectively. Unity raises the `sceneLoaded` event after `OnEnable` but before `Start` for all objects in the scene. For details and example usage, refer to the relevant API reference pages.

For a diagram that includes scene load as part of the execution flow, refer to [Details of disabling Domain and Scene reload](configurable-enter-play-mode-details.md)

## Run code on Editor launch

Sometimes it can be useful to make parts of your code run immediately on launch of the Unity Editor or runtime, without any additional user action and without the code needing to be part of a MonoBehaviour script. You can run code on Editor launch without requiring any user action by applying the [`[InitializeOnLoad]`](../ScriptReference/InitializeOnLoadAttribute.md) attribute to a class that has a [static constructor](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/static-constructors). Alternatively, you can apply the [`[InitializeOnLoadMethod]`](../ScriptReference/InitializeOnLoadMethodAttribute.md) attribute to individual methods. For more information and usage examples, refer to the API references for these attributes.

## Run code on runtime intialization

You can run code on initialization of the runtime application by applying the [`[RuntimeInitializeOnLoadMethodAttribute]`](../ScriptReference/RuntimeInitializeOnLoadMethodAttribute.md) to methods. You can also specify a [`RunTimeInitializeLoadType`](../ScriptReference/RuntimeInitializeLoadType.md) attribute parameter to control where in the Player loop the attributed code executes. For more information on the execution order of methods marked with this attribute, refer to the API reference for [`RuntimeInitializeOnLoadMethodAttribute`](../ScriptReference/RuntimeInitializeOnLoadMethodAttribute.md).

## Internal animation update

The following diagram shows the order of execution for the regular Animation update loop, and expands the nodes labelled **Internal animation update** in the previous diagram:

![Order of execution for the regular Animation update loop.](../uploads/Main/animation-update-sequence.svg)

Order of execution for the regular Animation update loop.

The following Animation loop callbacks shown in the diagram are called on **scripts** that derive from [MonoBehaviour](../ScriptReference/MonoBehaviour.md):

* [`MonoBehaviour.OnAnimatorMove`](../ScriptReference/MonoBehaviour.OnAnimatorMove.md)
* [`MonoBehaviour.OnAnimatorIK`](../ScriptReference/MonoBehaviour.OnAnimatorIK.md)

Additional animation-related event functions are called on scripts that derive from [`StateMachineBehaviour`](../ScriptReference/StateMachineBehaviour.md):

* [`StateMachineBehaviour.OnStateMachineEnter`](../ScriptReference/StateMachineBehaviour.OnStateMachineEnter.md)
* [`StateMachineBehaviour.OnStateMachineExit`](../ScriptReference/StateMachineBehaviour.OnStateMachineExit.md)
* [`StateMachineBehaviour.OnStateEnter`](../ScriptReference/StateMachineBehaviour.OnStateEnter.md)
* [`StateMachineBehaviour.OnStateUpdate`](../ScriptReference/StateMachineBehaviour.OnStateUpdate.md)
* [`StateMachineBehaviour.OnStateExit`](../ScriptReference/StateMachineBehaviour.OnStateExit.md)
* [`StateMachineBehaviour.OnStateMove`](../ScriptReference/StateMachineBehaviour.OnStateMove.md)
* [`StateMachineBehaviour.OnStateIK`](../ScriptReference/StateMachineBehaviour.OnStateIK.md)

Other animation functions shown in the diagram are internal to the animation system and are provided for context. These functions have associated Profiler markers so you can use the [Profiler](Profiler.md) to see when in the frame Unity calls them. Knowing when Unity calls these functions can help you understand exactly when the event functions you do call are executed. For a full execution order of animation functions and profiler markers, refer to [Profiler markers](profiler-markers.md#animation).

## Rendering

This execution order applies for the [Built-in Render Pipeline](built-in-render-pipeline.md) only. For details of execution order in **render pipelines** based on the [Scriptable Render Pipeline](scriptable-render-pipeline-introduction.md), refer to the relevant sections of the documentation for the [Universal Render Pipeline](urp/customize/custom-pass-injection-points.md) or the [High Definition Render Pipeline](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/index.html?subfolder=/manual/rendering-execution-order.html). If you want to do work immediately prior to rendering, refer to [Application.onBeforeRender](../ScriptReference/Application-onBeforeRender.md).

* [`OnPreCull`](../ScriptReference/MonoBehaviour.OnPreCull.md): Called before the **camera** culls the scene. Culling determines which objects are visible to the camera. `OnPreCull` is called just before culling takes place.
* [`OnBecameVisible`](../ScriptReference/MonoBehaviour.OnBecameVisible.md)/[`OnBecameInvisible`](../ScriptReference/MonoBehaviour.OnBecameInvisible.md): Called when an object becomes visible/invisible to any camera. `OnBecameInvisible` is not shown in the flow chart above since an object may become invisible at any time.
* [`OnWillRenderObject`](../ScriptReference/MonoBehaviour.OnWillRenderObject.md): Called **once** for each camera if the object is visible.
* [`OnPreRender`](../ScriptReference/MonoBehaviour.OnPreRender.md): Called before the camera starts rendering the scene.
* [`OnRenderObject`](../ScriptReference/MonoBehaviour.OnRenderObject.md): Called after all regular scene rendering is done. You can use [GL](../ScriptReference/GL.md) class or [Graphics.DrawMeshNow](../ScriptReference/Graphics.DrawMeshNow.md) to draw custom geometry at this point.
* [`OnPostRender`](../ScriptReference/MonoBehaviour.OnPostRender.md): Called after a camera finishes rendering the scene.
* [`OnRenderImage`](../ScriptReference/MonoBehaviour.OnRenderImage.md): Called after scene rendering is complete to allow **post-processing** of the image, see [Post-processing Effects](PostProcessingOverview.md).
* [`OnGUI`](../ScriptReference/MonoBehaviour.OnGUI.md): Called multiple times per frame in response to GUI events. The Layout and Repaint events are processed first, followed by a Layout and keyboard/mouse event for each input event.
* [`OnDrawGizmos`](../ScriptReference/MonoBehaviour.OnDrawGizmos.md) Used for drawing **Gizmos** in the **scene view** for visualisation purposes.

**Note**: [OnPreCull](../ScriptReference/Camera.OnPreCull.md), [OnPreRender](../ScriptReference/Camera.OnPreRender.md), [OnPostRender](../ScriptReference/Camera.OnPostRender.md), and [OnRenderImage](../ScriptReference/Camera.OnRenderImage.md) are built-in Unity event functions that are called on MonoBehaviour scripts but **only if those scripts are attached to the same object as an enabled Camera component**. If you want to receive the equivalent callbacks for `OnPreCull`, `OnPreRender`, and `OnPostRender` on a MonoBehaviour attached to a **different** object, you must use the equivalent delegates (note the lowercase `on` in the names) [Camera.onPreCull](../ScriptReference/Camera-onPreCull.md), [Camera.onPreRender](../ScriptReference/Camera-onPreRender.md), and [Camera.onPostRender](../ScriptReference/Camera-onPostRender.md) as shown in the code examples in the relevant pages of the scripting reference.

## Resumption of coroutines and asynchronous tasks

Suspended coroutines can resume at different points in the execution sequence depending on the yield instruction used. For example, coroutines that use [`WaitForEndOfFrame`](../ScriptReference/WaitForEndOfFrame.md) resume at the end of the frame, while those that use [`WaitForFixedUpdate`](../ScriptReference/WaitForFixedUpdate.md) resume at the end of the fixed update step. For more information, refer to [Coroutines](coroutines.md).

Regular .NET Tasks and asynchronous methods resume in the `Update` phase. Similarly to coroutines, Unity’s custom [`Awaitable`](../ScriptReference/Awaitable.md) class can resume at different points depending on the method you use when awaiting. For more information, refer to [Asynchronous programming with the Awaitable class](async-await-support.md).

**Note**: The exact order of execution between resuming coroutines and asynchronous tasks is not guaranteed. Awaitables are grouped together and executed in the order they were awaited.

## Combining MonoBehaviours with Entities

When using the [Entity Component System](https://docs.unity3d.com/Packages/com.unity.entities@latest/index.html) (ECS), Unity merges ECS system group updates into the Player update loop.

You can use the Entities Systems window to view the update order of ECS system groups relative to the full Player loop. For more information, refer to [Update order of systems](https://docs.unity3d.com/Packages/com.unity.entities@1.4/manual/systems-update-order.html#update-order-of-systems) in the Entities package documentation.

## Limitations

In general, you can’t rely on the order in which the same event function is invoked for different **GameObjects**, except when the order is explicitly documented or settable.

You can’t specify the order in which an event function is called for different instances of the same MonoBehaviour script. For example, the `Update` function of one MonoBehaviour might be called before or after the `Update` function for the same MonoBehaviour on another GameObject, including its own parent or child GameObjects.

To configure the execution order between different MonoBehaviour scripts, refer to [Script execution order](script-execution-order.md).

## Additional resources

* [Event functions](event-functions.md)
* [MonoBehaviour](class-MonoBehaviour.md)
* [PlayerLoop API reference](../ScriptReference/LowLevel.PlayerLoop.md)
* [Script execution order](script-execution-order.md)