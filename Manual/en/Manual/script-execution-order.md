# Script execution order

In addition to the [execution order of functions within scripts](execution-order.md), another important consideration for your application is the order in which Unity executes different script components. For example, if you have two **scripts**, `EngineBehaviour` and `SteeringBehaviour`, you might want to ensure that `EngineBehaviour` always updates before `SteeringBehaviour`.

You can configure script execution order in the following ways:

* In the Unity Editor, by using the [**Script Execution Order**](class-MonoManager.md) settings in the ****Project Settings**** window.
* In code, by applying the [`[DefaultExecutionOrder]`](../ScriptReference/DefaultExecutionOrder.md) attribute to your MonoBehaviour-derived classes.

## Important considerations and limitations

* Execution order defined in code with [`[DefaultExecutionOrder]`](../ScriptReference/DefaultExecutionOrder.md) does not show in the [**Script Execution Order**](class-MonoManager.md) settings window. If you define an execution order for a MonoBehaviour-derived type in code with `[DefaultExecutionOrder]` but define a different value for the same type in the Editor’s Project settings window, Unity uses the value defined in the Editor UI.
* If you assign multiple instances of the same scripts to different GameObjects, all instances of a script with a lower execution order value are executed before any instances of a script with a higher execution order value, regardless of which **GameObject** they’re attached to.
* When multiple **scenes** are [loaded additively](../ScriptReference/SceneManagement.LoadSceneMode.Additive.md), the configured script execution order is applied in full for one scene at a time rather than partially across scenes. In the previously cited example of an `EngineBehaviour` script configured to execute before a `SteeringBehaviour` script, both would update on one scene before they updated on the next one.
* When multiple scripts have either the same configured execution order or the default execution order, the order of execution between them is not deterministic. While the order might appear consistent during testing, you should never rely on this behavior, because it isn’t guaranteed across builds, machines, or Unity versions. For more information, refer to [Event function execution order](execution-order.md).
* The execution order specified in the **Script Execution Order** settings window doesn’t affect:
  + The order of functions marked with the [`[RuntimeInitializeOnLoadMethod]`](../ScriptReference/RuntimeInitializeOnLoadMethodAttribute.md) attribute.
  + [`OnDisable`](../ScriptReference/MonoBehaviour.OnDisable.md) and [`OnDestroy`](../ScriptReference/MonoBehaviour.OnDestroy.md) functions.

## Additional resources

* [Event function execution order](execution-order.md)
* [`[DefaultExecutionOrder]` API reference](../ScriptReference/DefaultExecutionOrder.md)