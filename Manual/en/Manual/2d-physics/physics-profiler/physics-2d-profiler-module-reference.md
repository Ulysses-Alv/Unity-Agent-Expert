# Physics 2D Profiler module reference

The Physics 2D Profiler module displays information about the physics that the physics system has processed in your project’s **scene**. This information can help you diagnose and resolve performance issues or unexpected discrepancies related to the physics in your project’s scene.

![Profiler window with the Physics 2D module selected.](../../../uploads/Main/profiler-physics-2d-module.png)

Profiler window with the Physics 2D module selected.

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**. For more information on how to use the Profiler window, refer to [Getting started with the Profiler window](../../ProfilerWindow.md).

## Chart categories

The Physics 2D Profiler module’s chart is divided into the following categories:

| **Chart** | **Description** |
| --- | --- |
| **Total Contacts** | The total number of contacts that were present in this frame. This includes both **collision** and trigger contacts. |
| **Total Shapes** | The total number of [physics shapes](../../../ScriptReference/PhysicsShape2D.md) that were present in this frame. Different [Collider2D](../../../ScriptReference/Collider2D.md) produce different amounts of physics shapes ranging from one to unlimited. You can get the [shape count](../../../ScriptReference/Collider2D-shapeCount.md) and [retrieve the physics shapes](../../../ScriptReference/Collider2D.GetShapes.md) to determine this for any Collider2D. |
| **Total Queries** | The total number of queries that were called this frame. This includes queries such as [Physics2D.Raycast](../../../ScriptReference/Physics2D.Raycast.md), [Physics2D.OverlapPoint](../../../ScriptReference/Physics2D.OverlapPoint.md) etc. |
| **Total Callbacks** | The total number of [OnCollisionEnter2D](../../../ScriptReference/Collider2D.OnCollisionEnter2D.md), [OnCollisionStay2D](../../../ScriptReference/Collider2D.OnCollisionStay2D.md), [OnCollisionExit2D](../../../ScriptReference/Collider2D.OnCollisionExit2D.md), [OnTriggerEnter2D](../../../ScriptReference/Collider2D.OnTriggerEnter2D.md), [OnTriggerStay2D](../../../ScriptReference/Collider2D.OnTriggerStay2D.md) and [OnTriggerExit2D](../../../ScriptReference/Collider2D.OnTriggerExit2D.md) callbacks that were called in this frame. |
| **Total **Joints**** | The total number of any [Joint2D](../../../ScriptReference/Joint2D.md) that were present in this frame. |
| **Total Bodies** | The total number of [Rigidbody2D](../../../ScriptReference/Rigidbody2D.md) that were present in this frame. |
| **Awake Bodies** | The total number of Rigidbody2D that were both [awake](../../../ScriptReference/Rigidbody2D.IsAwake.md) (not [sleeping](../../../ScriptReference/Rigidbody2D.IsSleeping.md)) and were present in this frame. |
| **Dynamic Bodies** | The total number of Rigidbody2D with a [Dynamic](../../../ScriptReference/RigidbodyType2D.Dynamic.md) **body type** that were present in this frame. |
| **Continuous Bodies** | The total number of Rigidbody2D with a [Continuous](../../../ScriptReference/CollisionDetectionMode2D.Continuous.md) **collision detection** mode that were present in this frame. |
| **Physics Used Memory** | The total amount of persistent memory used exclusively by the 2D physics system. This includes both the core engine and the memory used by each physics component, but doesn’t include the temporary memory used in this frame. |

Click on the frame chart window or select a captured frame in the chart graph to track selected categories. To change the order of the categories in the chart, drag them in the chart’s legend. You can also click a category’s colored legend to toggle its display. Refer to the [module details pane](#module-details-pane) for more information about the selected statistics.

## Module details pane

When you select a frame in the Physics 2D Profiler module, the details pane displays detailed information about the physics in your project’s scene. The details pane is sorted by category, where each category exists on a single line.

The following reference table describes the statistics available, plus its corresponding [profiler counter](../../profiler-counters-reference.md), and availability in release builds. The profiler counters are always available in the Editor and in **development builds**. This information is also available via the [ProfilerRecorder API](../../../ScriptReference/Unity.Profiling.ProfilerRecorder.md) and in the [Profiler Module Editor](../../profiler-module-editor.md) so you can add them to a custom Profiler module.

### Physics Used Memory

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total amount of persistent memory used exclusively by the 2D physics system. This includes both the core engine and the memory used by each physics component, but doesn’t include temporary memory used in this frame. | Physics Used Memory 2D | No |
| **Relative** | The relative percentage of memory used by the 2D physics system compared to the overall memory usage of Unity. | N/A | N/A |

### Bodies

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of [Rigidbody2D](../../../ScriptReference/Rigidbody2D.md) that were present in this frame. | Total Bodies | No |
| **Awake** | The number of Rigidbody2D that were both [awake](../../../ScriptReference/Rigidbody2D.IsAwake.md) (not [sleeping](../../../ScriptReference/Rigidbody2D.IsSleeping.md)) and were present in this frame. Note that a Rigidbody2D with a [Static](../../../ScriptReference/RigidbodyType2D.Static.md) body type is always asleep. | Awake Bodies | No |
| **Asleep** | The number of Rigidbody2D that were both sleeping (not awake) and were present in this frame. Note that a Rigidbody2D with a Static body type is always asleep. | Asleep Bodies | No |
| **Dynamic** | The number of Rigidbody2D with a [Dynamic](../../../ScriptReference/RigidbodyType2D.Dynamic.md) body type that were present in this frame. Dynamic bodies take the most processing of all body types. | Dynamic Bodies | No |
| **Kinematic** | The number of Rigidbody2D with a [Kinematic](../../../ScriptReference/RigidbodyType2D.Kinematic.md) body type that were present in this frame. Kinematic bodies have minimal processing. | Kinematic Bodies | No |
| **Static** | The number of Rigidbody2D with a [Static](../../../ScriptReference/RigidbodyType2D.Static.md) body type that were present in this frame. Static bodies take the least processing of all body types. | Static Bodies | No |
| **Discrete** | The number of Rigidbody2D with a [Discrete](../../../ScriptReference/CollisionDetectionMode2D.Discrete.md) collision detection mode that were present in this frame. Discrete bodies are far less performance-intensive than when using [Continuous](../../../ScriptReference/CollisionDetectionMode2D.Continuous.md) collision detection mode. | Discrete Bodies | No |
| **Continuous** | The number of Rigidbody2D with a [Continuous](../../../ScriptReference/CollisionDetectionMode2D.Continuous.md) collision detection mode that were present in this frame. Continuous bodies are much more performance-intensive than when using **Discrete collision detection** mode.. | Continuous Bodies | No |

### Shapes

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of [physics shapes](../../../ScriptReference/PhysicsShape2D.md) that were present in this frame. Different [Collider2D](../../../ScriptReference/Collider2D.md) produce different amounts of physics shapes ranging from one to unlimited. You can get the [shape count](../../../ScriptReference/Collider2D-shapeCount.md) and [retrieve the physics shapes](../../../ScriptReference/Collider2D.GetShapes.md) to determine this for any Collider2D. | Total Shapes | No |
| **Awake** | A physics shape is awake if it is attached to a [Rigidbody2D](../../../ScriptReference/Rigidbody2D.md) that is [awake](../../../ScriptReference/Rigidbody2D.IsAwake.md). This is the number of physics shapes that were both awake (not [sleeping](../../../ScriptReference/Rigidbody2D.IsSleeping.md)) and were present in this frame. | Awake Shapes | No |
| **Asleep** | A physics shape is asleep if it is attached to a Rigidbody2D that is asleep. This is the number of physics shapes that were both sleeping (not awake) and were present in this frame. | Asleep Shapes | No |
| **Dynamic** | A physics shape is [Dynamic](../../../ScriptReference/RigidbodyType2D.Dynamic.md) if it is attached to a Rigidbody2D with a Dynamic body type. This is the number of physics shapes that were both Dynamic and were present in this frame. | Dynamic Shapes | No |
| **Kinematic** | A physics shape is Kinematic if it is attached to a Rigidbody2D with a [Kinematic](../../../ScriptReference/RigidbodyType2D.Kinematic.md) body type. This is the number of physics shapes that were both Kinematic and were present in this frame. | Kinematic Shapes | No |
| **Static** | A physics shape is Static if it is attached to a Rigidbody2D with a [Static](../../../ScriptReference/RigidbodyType2D.Static.md) body type. This is the number of physics shapes that were both Static and were present in this frame. | Static Shapes | No |

### Queries

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of queries that were called this frame. This includes queries such as [`Physics2D.Raycast`](../../../ScriptReference/Physics2D.Raycast.md), and [`Physics2D.OverlapPoint`](../../../ScriptReference/Physics2D.OverlapPoint.md). | Total Queries | No |
| **Raycast** | The number of ray or line based queries that were called this frame. This includes queries such as `Physics2D.Raycast` and [`Physics2D.Linecast`](../../../ScriptReference/Physics2D.Linecast.md). | Raycast Queries | No |
| **Shapecast** | The number of shape swept queries that were called this frame. This includes queries such as [`Physics2D.BoxCast`](../../../ScriptReference/Physics2D.BoxCast.md), [`Physics2D.CircleCast`](../../../ScriptReference/Physics2D.CircleCast.md), and [`Collider2D.Cast`](../../../ScriptReference/Collider2D.Cast.md). | Shapecast Queries | No |
| **Overlap** | The number of overlap queries that were called this frame. This includes queries such as [`Physics2D.OverlapPoint`](../../../ScriptReference/Physics2D.OverlapPoint.md), [`Physics2D.OverlapCircle`](../../../ScriptReference/Physics2D.OverlapCircle.md), and [`Collider2D.OverlapCollider`](../../../ScriptReference/Collider2D.OverlapCollider.md). | Overlap Queries | No |
| **IsTouching** | The number of contact touching queries that were called this frame. This includes queries such as [`Physics2D.IsTouching`](../../../ScriptReference/Physics2D.IsTouching.md), [`Collider2D.IsTouching`](../../../ScriptReference/Collider2D.IsTouching.md), [`Rigidbody2D.IsTouching`](../../../ScriptReference/Rigidbody2D.IsTouchingLayers.md) etc. | IsTouching Queries | No |
| **GetContacts** | The number of contact retrieval queries that were called this frame. This includes queries such as [`Physics2D.GetContacts`](../../../ScriptReference/Physics2D.GetContacts.md), [`Collider2D.GetContacts`](../../../ScriptReference/Collider2D.GetContacts.md), and [Rigidbody2D.GetContacts](../../../ScriptReference/Rigidbody2D.GetContacts.md). Note that this doesn’t include [`Collision2D.GetContacts`](../../../ScriptReference/Collision2D-contacts.md) which isn’t a physics query. | GetContacts Queries | No |
| ****Particle**** | The number of queries that the **Particle System** called this frame. This is used when the Particle System module is configured to contact 2D physics **colliders** and is entirely controlled by the Particle System. Note that this can have a high impact on performance, but is also very efficient to process. | Particle Queries | No |

### Contacts

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of contacts that were present in this frame. This includes both Collision and Trigger contacts. Processing and solving contacts can take a long time to process. | Total Contacts | No |
| **Added** | The number of contacts that were added in this frame. This includes both Collision and Trigger contacts. Adding too many contacts in a single frame can cause performance spikes. | Added Contacts | No |
| **Removed** | The number of contacts that were removed in this frame. This includes both Collision and Trigger contacts. Removing contacts is fast and has minimum impact on performance. | Removed Contacts | No |
| **Broadphase Updates** | The number of broadphase updates that were processed in this frame. A broadphase update occurs when physics shapes are added, removed or change in size. Broadphase updates are used to detect contact changes when two physics shapes potentially overlap and can result in a broadphase pair being created. | Broadphase Updates | No |
| **Broadphase Pairs** | The number of broadphase pairs that were processed in this frame. A broadphase pair is created when a broadphase update results in a potential overlap of two physics shapes. A broadphase pair is then processed and the result will be a new contact or it will be ignored if the physics shapes are not configured to contact each other. | Broadphase Pairs | No |

### Callbacks

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of [OnCollisionEnter2D](../../../ScriptReference/Collider2D.OnCollisionEnter2D.md), [OnCollisionStay2D](../../../ScriptReference/Collider2D.OnCollisionStay2D.md), [OnCollisionExit2D](../../../ScriptReference/Collider2D.OnCollisionExit2D.md), [OnTriggerEnter2D](../../../ScriptReference/Collider2D.OnTriggerEnter2D.md), [OnTriggerStay2D](../../../ScriptReference/Collider2D.OnTriggerStay2D.md) and [OnTriggerExit2D](../../../ScriptReference/Collider2D.OnTriggerExit2D.md) callbacks that were called in this frame. | Total Callbacks | No |
| **Collision Enter** | The number of OnCollisionEnter2D callbacks that were called in this frame. | Collision Enter | No |
| **Collision Stay** | The number of OnCollisionStay2D callbacks that were called in this frame. | Collision Stay | No |
| **Collision Exit** | The number of OnCollisionExit2D callbacks that were called in this frame. | Collision Exit | No |
| **Trigger Enter** | The number of OnTriggerEnter2D callbacks that were called in this frame. | Trigger Enter | No |
| **Trigger Stay** | The number of OnTriggerStay2D callbacks that were called in this frame. | Trigger Stay | No |
| **Trigger Exit** | The number of OnTriggerExit2D callbacks that were called in this frame. | Trigger Exit | No |

### Solver

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **World Count** | The total number of [physics scene](../../../ScriptReference/PhysicsScene2D.md) that were present in this frame. Each physics scene contains a physics world that can be simulated independently of any other physics world. Having a lot of worlds isn’t a performance issue because it only occupies memory and doesn’t perform any work unless it is simulated. | Solver World Count | No |
| **Simulation Count** | The number of times all physics scene were simulated either by Unity automatically, by calling [Physics2D.Simulate](../../../ScriptReference/Physics2D.Simulate.md) or by directly calling [PhysicsScene2D.Simulate](../../../ScriptReference/PhysicsScene2D.Simulate.md). | Solver Simulation Count | No |
| **Discrete Islands** | An island is a connected graph of bodies connected via mutual joints or mutual contacts. Note that [Static](../../../ScriptReference/RigidbodyType2D.Static.md) body types don’t connect islands. The number of contact islands solved when handling the discrete solving step. | Solver Discrete Islands | No |
| **Continuous Islands** | An island is a connected graph of bodies connected via mutual joints or mutual contacts. Note that Static body types don’t connect islands. This is the number of islands solved when handling the continuous solving step. Solving continuous islands is a performance-intensive process and involves multiple iterations that require islands to be regenerated and reprocessed. Only a [Rigidbody2D](../../../ScriptReference/Rigidbody2D.md) with a [Continuous](../../../ScriptReference/CollisionDetectionMode2D.Continuous.md) collision detection mode will result in this additional continuous island being formed and processed. | Solver Continuous Islands | No |

### Transform Sync

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Sync Calls** | The total number of Transform sync calls that were called in this frame. A Transform sync (known as a Transform Read) involves checking if any Transforms have changed and if so, the Transform poses are read and cause any [Rigidbody2D](../../../ScriptReference/Rigidbody2D.md) or [Collider2D](../../../ScriptReference/Collider2D.md) to be updated. Transforms should not be changed when using physics components however sometimes this is necessary but should be avoided due to potential performance issues if performing too many. Any movements should be performed by using the Rigidbody2D API.  The physics system will perform a single Transform sync as the first part of performing a simulation step so this will always be at least one if a simulation occurred (see [Simulation Count](#SimCount) above). The physics system will also perform a single Transform sync per-frame if it is handling any Rigidbody2D interpolation.  Additional calls are shown if either [Physics2D.AutoSyncTransforms](../../../ScriptReference/Physics2D-autoSyncTransforms.md) is active (inactive by default) or if [Physics2D.SyncTransforms](../../../ScriptReference/Physics2D.SyncTransforms.md) is called although both of these should be avoided as they can both have a severe impact on performance. | Total Transform Sync Calls | No |
| **Sync Bodies** | The number of Rigidbody2D that were affected by a Transform sync. This should be kept to a minimum, preferably zero. | Transform Sync Bodies | No |
| **Sync Colliders** | The number of Collider2D that were affected by a Transform sync. This should be kept to a minimum, preferably zero. | Transform Sync Colliders | No |
| **Parent Sync Bodies** | The number of Rigidbody2D that were affected by a Transform sync caused by reparenting a Transform.. This should be kept to a minimum, preferably zero. | Transform Parent Sync Bodies | No |
| **Parent Sync Colliders** | The number of Collider2D that were affected by a Transform sync caused by reparenting a Transform. This should be kept to a minimum, preferably zero. | Transform Parent Sync Colliders | No |

### Joints

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Total** | The total number of any [Joint2D](../../../ScriptReference/Joint2D.md) that were present in this frame. Solving joints can become expensive so these should be kept to a minimum. | Total Joints | No |

### Timings

**Note**: All timings are summed over all physics worlds (see [World Count](#WorldCount)). The number of times any timing was sampled is shown in square brackets after the timing itself. Effectively the timing can be divided by the **World Count** value to give an average time.

| **Statistic** | **Description** | **Corresponding Profiler Counter (exact name)** | **Available in Release Players?** |
| --- | --- | --- | --- |
| **Sim** | The total amount of time spent handling a full simulation step. This can be called by Unity automatically, by calling [Physics2D.Simulate](../../../ScriptReference/Physics2D.Simulate.md) or by directly calling [PhysicsScene2D.Simulate](../../../ScriptReference/PhysicsScene2D.Simulate.md). This time includes all the stages involved in completing a simulation step including Transform Sync (read), Calculating Contacts, Integration, Solving Contacts and Joints, Transform Write and Contact Callbacks. | N/A | N/A |
| **Sync** | The total amount of time spent processing Transform Sync (see [Sync Calls](#SyncCalls)). | N/A | N/A |
| **Step** | The total amount of time spent processing simulation steps. This time includes only the core stages involved in completing a simulation step including Calculating Contacts, Integration, Solving Contacts and Joints. | N/A | N/A |
| **Write** | The total amount of time spent processing Transform write. This happens during the end of the simulation step where body poses are read and written back to the Transform system. | N/A | N/A |
| **Callbacks** | The total amount of time spent processing all callbacks (see [Total Callbacks](#TotalCallbacks)). | N/A | N/A |

## Additional resources

* [Profiler window introduction](../../ProfilerWindow.md)
* [Profiling your application](../../profiler-profiling-applications.md)