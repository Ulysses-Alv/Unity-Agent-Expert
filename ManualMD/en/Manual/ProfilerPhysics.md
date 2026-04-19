# Physics Profiler module

The Physics Profiler module displays information about the physics that the physics system has processed in your project’s **scene**. This information can help you diagnose and resolve performance issues or unexpected discrepancies related to the physics in your project’s scene. You can also use the [Physics Debug Visualization](PhysicsDebugVisualization.md) to further debug and understand issues with physics in your application.

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**. For more information on how to use the Profiler window, refer to [Getting started with the Profiler window](ProfilerWindow.md).

![Physics Profiler module](../uploads/Main/profiler-physics-module.png)

Physics Profiler module

## Chart categories

The Physics Profiler module’s chart tracks the time your application spends on physics. The timings are divided into different chart categories:

| **Chart** | **Function** |
| --- | --- |
| **Physics Used Memory** | The total amount of memory that the physics module has used. |
| **Active Dynamic Bodies** | The number of [Rigidbody](../ScriptReference/Rigidbody.md) components and [ArticulationBody](../ScriptReference/ArticulationBody.md) components that aren’t in a [sleep](RigidbodiesOverview.md) state. |
| **Active Kinematic Bodies** | The number of active Kinematic Rigidbody components. A Kinematic Rigidbody is active when [MovePosition](../ScriptReference/Rigidbody.MovePosition.md) or [MoveRotation](../ScriptReference/Rigidbody.MoveRotation.md) is called in a frame, and remains active in the next frame.  **Note**: Unity might process Kinematic Rigidbody components that have joints attached multiple times per frame, and this contributes to the value displayed. |
| **Dynamic Bodies** | The number of Rigidbody components and ArticulationBody components. |
| **Overlaps** | The number of overlap events. An overlapping event is when [colliders](../ScriptReference/Collider.md) overlap with each other. |
| **Trigger Overlaps** | The number of overlap events with trigger colliders (counted in pairs). |
| **Discreet Overlaps** | The number of overlap events which Unity used discrete **collision** detection to solve. |
| **Continuous Overlaps** | The number of overlap events which Unity used continuous **collision detection** to solve. |
| **Physics Queries** | The total amount of physics queries, such as [Raycasts](../ScriptReference/Physics.Raycast.md) and shapecasts. |

To change the order of the categories in the chart, you can drag them in the chart’s legend. You can also click a category’s colored legend to toggle its display. When you click on the chart, you can view the exact numerical values of each chart category in the module details pane.

## Module details pane

When you click on the Physics Profiler module, the module details pane displays further information about the physics in your project’s scene.

| **Statistic** | **Description** |
| --- | --- |
| **Physics Used Memory** | The total amount of memory that the physics module has used. |
| **Dynamic Bodies** | The number of [Rigidbody](../ScriptReference/Rigidbody.md) components and [ArticulationBody](../ScriptReference/ArticulationBody.md) components. |
| **Articulation Bodies** | The number of ArticulationBody components in the scene |
| **Active Dynamic Bodies** | The number of Rigidbody components and ArticulationBody components that aren’t in a [sleep](RigidbodiesOverview.md) state.. |
| **Active Kinematic Bodies** | The number of active Kinematic Rigidbody components. A Kinematic Rigidbody is active when [MovePosition](../ScriptReference/Rigidbody.MovePosition.md) or [MoveRotation](../ScriptReference/Rigidbody.MoveRotation.md) is called in a frame, and remains active in the next frame.  **Note**: Unity might process Kinematic Rigidbody components that have joints attached multiple times per frame, and this contributes to the value displayed. |
| **Static Colliders** | The number of [colliders](../ScriptReference/Collider.md) that don’t have a Rigidbody or ArticulationBody component |
| **Colliders Synced** | The amount of colliders synced with [Transforms](../ScriptReference/Transform.md). |
| **Rigidbodies Synced** | The amount of Rigidbody components synced with Transforms. |
| **Physics Queries** | The total amount of physics queries, such as [Raycasts](../ScriptReference/Physics.Raycast.md) and shapecasts. |
| **Total Overlaps** | The number of overlap events. An overlapping event is when colliders overlap with each other. The events are organized into the following categories:   * **Discreet**: The number of overlap events which Unity used discrete collision detection to solve. * **Continuous**: The number of overlap events which Unity used continuous collision detection to solve. * **Trigger**: The number of overlap events with trigger colliders (counted in pairs). * **Modified**: The number of overlap events which Unity used the [Contact Modification](../ScriptReference/Physics.ContactModifyEvent.md) API to modify. |
| **Broadphase Adds/Removes** | The total number of colliders that the [broadphase algorithm](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/guide/3.3.4/Manual/RigidBodyCollision.html#broad-phase-algorithms) either added or removed. |
| **Narrowphase Touches** | The total amount of collision events that were either lost or appeared as new since the previous frame. |

The numbers displayed in the Profiler might not correspond to the exact number of **GameObjects** with physics components in your Scene. This is because Unity processes some physics components at a different rate depending on which other components affect it (for example, an attached Joint component). To calculate the exact number of GameObjects with specific physics components attached, you must write a custom script with the [FindObjectsOfType](../ScriptReference/Object.FindObjectsOfType.md) function.

The Physics Profiler module doesn’t display the number of sleeping Rigidbody components. These are components which don’t engage with the physics system, so the Profiler doesn’t process them. For more information on sleeping Rigidbody components, see the documentation on [Rigidbody](RigidbodiesOverview.md).

## Legacy Physics Profiler module

You can switch to the **Legacy** mode to see the older version of the Physics Profiler module, which was the default module in older versions of Unity. To do this, select **Legacy** from the dropdown menu in the top right of the Physics Profiler module’s details pane.

In this mode, you can load and inspect Profiler data that was saved in an older version of Unity. If you switch to this mode to inspect data captured in a newer version of Unity, then the data displayed is unreliable and inaccurate. You should always use the **Current** mode to inspect new Physics Profiler data.

| **Statistic** | **Description** |
| --- | --- |
| **Active Dynamic** | The number of active non-Kinematic Rigidbody components. An active Rigidbody is one that isn’t [sleeping](RigidbodiesOverview.md). |
| **Active Kinematic** | The number of active Kinematic Rigidbody components. A Kinematic Rigidbody is active when [MovePosition](../ScriptReference/Rigidbody.MovePosition.md) or [MoveRotation](../ScriptReference/Rigidbody.MoveRotation.md) is called in a frame, and remains active in the next frame.  **Note:** Unity might process Kinematic Rigidbody components that have joints attached multiple times per frame, and this contributes to the value displayed. |
| **Static Colliders** | The number of Collider components on GameObjects that don’t have Rigidbody components attached to the GameObjects or their parent GameObjects.  Collider components on GameObjects or parent GameObjects that have Rigidbody components do not count as Static Colliders. These are called Compound Colliders. Compound Colliders arrange multiple Colliders of a body in a convenient way, rather than having all of the Colliders on the same GameObject as the Rigidbody component. |
| **Rigidbody** | The number of Rigidbody components processed by the physics system, irrespective of the components’ sleeping state. |
| **Trigger Overlaps** | The number of overlapping triggers (counted in pairs). |
| **Active Constraints** | The number of primitive constraints the physics system has processed. Constraints are used as a building block of **Joints** as well as collision response. For example, restricting a linear or rotational degree of freedom of a [ConfigurableJoint](../ScriptReference/ConfigurableJoint.md) involves a primitive constraint per each restriction. |
| **Contacts** | The total number of contact pairs between all Colliders in the Scene, including the amount of trigger overlap pairs. A contact is a pair of Colliders that either touch or overlap. Note: Unity creates contact pairs per Collider pair once the distance between them is below a certain user configurable limit. As such, you might see contacts generated for Rigidbody components that are not yet touching or overlapping. See documentation on [Collider.contactOffset](../ScriptReference/Collider-contactOffset.md) and [ContactPoint.separation](../ScriptReference/ContactPoint-separation.md) for more details. |

## Additional resources

* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)