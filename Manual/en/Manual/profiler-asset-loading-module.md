# Asset Loading Profiler module reference

The Asset Loading Profiler module displays information about how your application loads assets, including a breakdown of read operations by area. The details pane provides an insight into each of the asset loading markers captured during the profile. You can use this information to find out how efficiently your application loads assets, and to identify any specific issues.

It uses the information from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) to display information. For information about `AsyncReadManager`, refer to [Collect asset loading metrics](AsyncReadManagerMetrics.md).

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**.

This module isn’t enabled by default. To enable the Asset Loading Profiler module, open the Profiler window, select the **Profiler Modules** dropdown menu, and toggle the **Asset Loading** checkbox.

![The Asset Loading Profiler module](../uploads/Main/profiler-asset-loading.png)

The Asset Loading Profiler module

## Chart categories

The Asset Loading Profiler module’s chart is divided into different categories. To change the order of the categories in the chart, drag them in the chart’s legend. You can also click a category’s colored legend to toggle its display. For more information on how to use the Profiler window, refer to [Getting started with the Profiler window](ProfilerWindow.md).

| **Chart** | **Description** |
| --- | --- |
| **Other Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) for an unspecified subsystem. |
| **Texture Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) for a texture load. |
| **Virtual Texture Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) for Virtual Texturing. |
| **Mesh Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) for a **mesh** load. |
| **Audio Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) for an audio load. |
| **Scripting Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) via the scripting API. |
| **Entities Reads** | The number of bytes requested from [`AsyncReadManager`](../ScriptReference/Unity.IO.LowLevel.Unsafe.AsyncReadManager.md) by **scripts** in the [Entities package](https://docs.unity3d.com/Packages/com.unity.entities@latest/). |

## Module details pane

The Asset Loading Profiler module doesn’t populate the details pane automatically. To populate the view with data, select the **Analyze Markers** button at the top of the details pane, then select a frame in the Asset Loading Profiler module.

The module details pane displays detailed information about how and when Unity loaded assets during the capture. This module can display information for either a single selected frame, or for all frames in the capture. Use the dropdown menu at the top-left of the details pane to choose between a single frame or all frames.

The **Asset Markers** view displays selected markers related to asset loading during the capture. Some rows have child markers that create a hierarchy, which is initially contained within a single expandable row. Select the dropdown arrow next to the marker’s name to expand the row and display all child markers.

The following table explains all the possible columns that the **Asset Markers** view can display. Some of these columns are disabled by default. To enable or disable a column, right-click anywhere in the header row of the table. Use the context menu to enable or disable the desired columns. You can enter text into the search bar next to the **Analyze Markers** button to filter the results based on the **Source** column.

| **Heading** | **Description** | **Enabled by default** |
| --- | --- | --- |
| **Index** | An ID number that represents the order of execution of the asset loading markers. | Yes |
| **Source** | The origin of the asset. This can be either a file path, an AssetBundle name, or a **Scene**. To display the full source file path for an entry, hover over the entry name. | Yes |
| **Asset Name** | The name of the asset, if applicable and available in the Profiler’s metadata for the marker. | Yes |
| **Thread** | The thread that this marker occurred on. | Yes |
| **Type** | The asset type. | Yes |
| **Size** | The size of the load recorded by this marker.   **Note:** This size might not include child markers, depending on the marker type. | Yes |
| **Marker Type** | The marker type recorded. | Yes |
| **Marker Length (ms)** | Total duration of this marker across all frames it spans in milliseconds. | Yes |
| **Frame Count** | Displays the index of the selected frame from the span of frames that the marker occupies. Also displays the total number of frames the marker spans. | Yes |
| **Timestamp (ms)** | The time this marker started, relative to the start of the Profiler capture. | No |
| **First Frame Index** | The frame this marker started on. | No |
| **Last Frame Index** | The frame this marker ended on. | No |

Right-click on any of the asset loading markers to open a menu with additional options to visualize that data differently. The available options are as follows:

| **Property** | **Function** |
| --- | --- |
| **Show marker in timeline view** | Opens this marker in the [CPU Usage Profiler](ProfilerCPU.md) module. |
| **Go to the first frame of this marker** | Selects the frame in which the marker first appears in the chart. This option is disabled if the marker only appears on one frame. |
| **Go to the last frame of this marker** | Selects the frame in which this marker last appears in the chart. This option is disabled if the marker only appears on one frame. |
| **Filter the view to this source** | Populates the search bar with the contents of the **Source** column to filter the results. |
| **Filter the view to this asset name** | Populates the search bar with the contents of the **Asset Name** column to filter the results. This option is disabled if the **Asset Name** column is empty. |

## Additional resources

* [Collect asset loading metrics](AsyncReadManagerMetrics.md)
* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)