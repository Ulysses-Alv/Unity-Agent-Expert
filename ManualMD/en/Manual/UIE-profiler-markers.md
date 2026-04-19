# UI Toolkit profiler markers

The UI Toolkit **profiler markers** help you monitor your UI performance. The markers provide detailed information about the time spent to render and update your UI. You can use the markers to identify performance bottlenecks and optimize your UI elements.

To open the Profiler window, go to **Window** > **Analysis**, and select a desired option. For more information, refer to [Unity Profiler](Profiler.md).

UI Toolkit uses the same [core profiler markers](#core-markers) for both runtime and Editor UI, though [some markers](#editor-only-markers) are specific to Editor UI.

When you profile in Play mode, the profiler displays these markers as described in [Runtime loop](#runtime-loop), for each Panel Settings used in your game. In Editor mode, the profiler displays these markers as outlined in [Editor loop](#editor-loop), for each Editor window currently open.

**Note:** UI Toolkit markers display varying time across frames in the Profiler. This is because UI Toolkit caches previous computations and tries to process only what has changed.

## High-level markers

High-level markers help you locate your specific uses of UI Toolkit, with each representing a high-level phase of UI Toolkit’s internal update loop. Each marker is prefixed with the panel name.

In the following table, `<panel-name>` refers to the Editor window name or **Panel Settings** asset name.

| **Marker** | **High-level phase that** |
| --- | --- |
| **`<panel-name>`.TickScheduledActions** (Editor-only) | Schedules actions to run at 10-millisecond interval. |
| **`<panel-name>`.TickScheduledActionsPreLayout** (runtime-only) | Schedules actions to run at a regular interval before validating the layout. |
| **`<panel-name>`.TickScheduledActionsPostLayout** (runtime-only) | Schedules actions to run at a regular interval after validating the layout. |
| **`<panel-name>`.ValidateLayout** | Computes the position of visual elements within the panel.   For Editor UI, this occurs only if a visual element was modified since the last frame, an event was processed, or an EditorWindow repaint was requested.   For runtime UI, this occurs every frame. |
| **`<panel-name>`.PrepareRepaint** | Computes the rendering data of visual elements within the panel.   If properties of the element that affect their positions were modified during the `ValidateLayout` phase, the `PrepareRepaint` phase computes the new positions.  For Editor UI, this occurs only if a visual element was modified since the last frame or an EditorWindow repaint was requested.   For runtime UI, this occurs every frame. |
| **`<panel-name>`.Render** | Draws the panel. When no changes are made to visual elements, the cost of this marker remains stable across frames.  For Editor UI, this occurs only if a visual element was modified since the last frame or an EditorWindow repaint was requested.   For runtime UI, this occurs every frame. |

## Core markers

Core markers represent the internal update loop of UI Toolkit. They are the same for both runtime and Editor UI.

| **Marker** | **Time spent to** |
| --- | --- |
| **UIElements.PickAll** | Identify the element under the pointer or being hovered. |
| **UIElements.UpdateRuntimeBindings** | Update the runtime bindings. |
| **UIElements.UpdateStyle** | Apply style sheets and compute the resolved style of the element. |
| **UIElements.UpdateLayout** | Compute and update the layout position of every element. This includes invoking the [`GeometryChangedEvent`](../ScriptReference/UIElements.GeometryChangedEvent.md) callback. |
| **UIElements.UpdateAnimation** | Update the animations. |
| **UIElements.UpdateRenderData** | Prepare graphics data for rendering the panel. This includes the time spent to invoke the [`VisualElement.generateVisualContents`](ScripteRef:UIElements.VisualElement-generateVisualContent) callbacks. |
| **UIR.DrawChain** | Draw the panel. When no changes are made to elements, the cost of this marker remains stable across frames. |
| **`<class-name>`.ImmediateRepaint** | Draw custom immediate mode elements. The cost remains stable across frames, though it depends on the implementation of the custom element. |

## Editor-only markers

The following markers appear only when profiling UI that runs inside the Editor (either Editor UI or runtime UI in the Game view).

| **Marker** | **Time spent to** |
| --- | --- |
| **UIElements.UpdateEditorBindings** | Process Serialized Object binding requests. |
| **UIElements.UpdateViewData** | Apply properties to elements that use ViewData. |
| **UIElements.UpdateAssetsInEditor** | Check if assets used in this panel have changed. This happens for both runtime and Editor UI but only performs work if Live Reload is enabled for the panel. |
| **IMGUIContainer / OnGUI** | Update IMGUIContainer elements. OnGUI is the time spent to invoke the IMGUI code, the difference is overhead. The cost remains stable across frames, though it depends on the actual implementation of the IMGUI callback. |

## Runtime loop

UI Toolkit’s update loop for runtime UI performs the same actions every frame. Each active panel, driven by Panel Settings with one or more UI Documents, has its own high-level markers.

### PreLateUpdate.UIElementsUpdatePanels

A panel’s pre late update loop occurs before [`MonoBehaviour.LateUpdate()`](../ScriptReference/MonoBehaviour.LateUpdate.md) callbacks and is divided into three main phases: `TickScheduledActionsPreLayout`, `ValidateLayout`, and `TickScheduledActionsPostLayout`.

The following screenshot shows the hierarchy of the `PreLateUpdate.UIElementsUpdatePanels` marker:

![PreLateUpdate.UIElementsUpdatePanels marker in the Profiler](../uploads/Main/uitk/U6_RuntimeLoop_PreLateUpdate-UIElementsUpdatePanels.png)

PreLateUpdate.UIElementsUpdatePanels marker in the Profiler

The following table lists the sequence of markers called for each phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.TickScheduledActionsPreLayout** | 1. Update scheduled items 2. Update focus 3. UIElements.UpdateAssetsInEditor |
| **`<panel-name>`.ValidateLayout** | 1. UIElements.UpdateStyle 2. UIElements.UpdateLayout 3. UIElements.UpdateBounds UIElements.PickAll (if pointer moved) |
| **`<panel-name>`.TickScheduledActionsPostLayout** | 1. UIElements.UpdateAnimation 2. UIElements.UpdateEditorBindings |

### PostLateUpdate.UIElementsRepaintPanels

A panel’s post late update loop occurs after [`MonoBehaviour.LateUpdate()`](../ScriptReference/MonoBehaviour.LateUpdate.md) callbacks and only has the `PrepareRepaint` phase.

The following screenshot shows the hierarchy of the `PostLateUpdate.UIElementsRepaintPanels` marker:

![PostLateUpdate.UIElementsRepaintPanels marker in the Profiler](../uploads/Main/uitk/U6_RuntimeLoop_PostLateUpdate-UIElementsRepaintPanels.png)

PostLateUpdate.UIElementsRepaintPanels marker in the Profiler

The following table lists the sequence of markers called for the `PrepareRepaint` phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.PrepareRepaint** | 1. UIElements.UpdateViewData 2. UIElements.UpdateRuntimeBindings 3. UIElements.UpdateStyle 4. UIElements.UpdateLayout 5. UIElements.UpdateBounds UIElements.PickAll (if pointer moved) 6. UIElements.UpdateRenderData |

### UI.RenderOverlays

The `UI.RenderOverlays` marker is called after `PostLateUpdate.UIElementsRepaintPanels`. It draws the overlay elements on top of the main panel content.

The following screenshot shows the hierarchy of the `UI.RenderOverlays` marker:

![UI.RenderOverlays marker in the Profiler](../uploads/Main/uitk/U6_RuntimeLoop_UI-RenderOverlays.png)

UI.RenderOverlays marker in the Profiler

**Note:** This capture is from a project using the Universal **Render Pipeline**. If you use the built-in Render Pipeline, `UI.RenderOverlays` appears under different markers.

The following table lists the sequence of markers called for the `Render` phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.Render** | * UIR.DrawChain draws the UI Toolkit content. * If you use the Immediate mode elements, `<class-name>`.ImmediateRepaint handles the drawing. |

## Editor loop

UI Toolkit’s update loop for Editor UI does not run every frame. Certain actions are scheduled to run every 10 milliseconds, while input event processing is driven by user interactions. Window repaints occur only when necessary or when explicitly requested using `EditorWindow.Repaint()`. Each `EditorWindow` has a panel with its own high-level markers.

### Scheduler.Tick

Scheduled actions are processed every 10 milliseconds. The following screenshot shows the hierarchy of the `Scheduler.Tick` marker:

![Scheduler.Tick marker in the Profiler](../uploads/Main/uitk/U6_EditorLoop_Scheduler-Tick.png)

Scheduler.Tick marker in the Profiler

The following table lists the sequence of markers called for the `TickScheduledActions` phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.TickScheduledActions** | 1. UIElements.UpdateAssetsInEditor 2. Update scheduled items 3. Update focus 4. UIElements.UpdateRuntimeBindings 5. UIElements.UpdateAnimation 6. UIElements.UpdateEditorBindings |

### `<window-name>.<mouse-event-name>`

When a mouse event is processed, the **UIElementsUtility.DoDispatch** marker is called to process the event. The event is divided into two main phases: `UIElementsUtility.DoDispatch` and `ValidateLayout`.

The following screenshot shows the hierarchy of an example `<window-name>.<mouse-event-name>` marker:

![InspectorWindow.MouseMove marker in the Profiler](../uploads/Main/uitk/U6_EditorLoop-window-event.png)

`InspectorWindow.MouseMove` marker in the Profiler

The following table lists the sequence of markers called for each phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.ValidateLayout** | 1. UIElements.UpdateStyle 2. UIElements.UpdateLayout 3. UIElements.UpdateBounds UIElements.PickAll (if pointer moved) UIElements.UpdateStyle (if hover element changed) |
| **UIElementsUtility.DoDispatch** | UIElements.PickAll |

### `<window-name>`.Repaint

When an EditorWindow repaint is requested, the **UIElementsUtility.DoDispatch** marker is called to process the repaint. The repaint is divided into two main phases: `Render` and `PrepareRepaint`.

The following screenshot shows the hierarchy of an example `<window-name>.Repaint` marker:

![InspectorWindow.Repaint marker in the Profiler](../uploads/Main/uitk/U6_EditorLoop-window-repaint.png)

`InspectorWindow.Repaint` marker in the Profiler

The following table lists the sequence of markers called for each phase:

| **Markers** | **Sequence** |
| --- | --- |
| **`<panel-name>`.PrepareRepaint** | 1. UIElements.UpdateViewData 2. UIElements.UpdateStyle 3. UIElements.UpdateLayout 4. UIElements.UpdateBounds UIElements.PickAll (if pointer moved) UIElements.UpdateStyle (if hover element changed) 5. UIElements.UpdateRenderData |
| **`<panel-name>`.Render** | * UIR.DrawChain draws the UI Toolkit content. * If you use `IMGUIContainers`, IMGUIContainer markers are displayed, each with two OnGUI markers underneath. * If you use Immediate mode elements, `<class-name>`.ImmediateRepaint handles the drawing. |

## Additional resources

* [Unity Profiler](Profiler.md)
* [Profiler markers reference](profiler-markers.md)