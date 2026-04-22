# Search performance trackers

The Performance Trackers search provider searches the Profiler data of applications and the Unity Editor. For more information, refer to [Collect performance data introduction](profiler-collect-data).

You can use the Performance Trackers search provider in two windows:

* In a dedicated window. From the main menu, select **Window** > **Analysis** > **Performance Markers**.
* In the **Search** window, using a visual or textual query with the provider token. However, the recommended best practice is using the dedicated window for performance trackers, as it provides a better experience.

## Query syntax

Provider token: `perf:`

Query example: Searches the performance trackers for an `Add Component Window` action with a sample count of over `1000`.

```
perf: count>1000 AddComponent
```

## Provider filters

The Performance Tracker provider has additional filters in the [visual query builder](search-visual-query.md), using standard [search query operators](search-query-operators.md).

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Average (avg)** | `perf: avg=<value>` | `perf: avg<512kb` | Searches for performance trackers with an average value smaller than `512kb`. |
| **Last time** | `perf: lastTime=<value>` | `perf: lastTime<1ms` | Searches for performance trackers that have been updated in the last millisecond. |
| **Peak** | `perf: peak=<value>` | `perf: peak=0.5` | Searches for performance trackers with a peak value of `0.5`. |
| **Sample count** | `perf: count=<value>` | `perf: count=0.5` | Searches for performance trackers with a sample count of `0.5`. This filter can’t use a unit. |
| **Total** | `perf: total=<value>` | `perf: total=0.5` | Searches for performance trackers with a total value of `0.5`. |

To match the unit the performance tracker logs, use different unit types with the filters. For example:

* To search all trackers that report an average time greater than `1ms`, use `avg>1ms`.
* To search for trackers that report an average memory consumption smaller than `512kb`, use `avg<512kb`.

If you don’t specify the unit, the search returns trackers that report hits.

**Note:** The Sample Count filter never uses units.

## Results

**Search** window tab: **Performance Trackers**.

The recommended best practice is to use the table view for these results.

## Actions

The context menu for the Performance Trackers search provider includes the following actions:

| Action | Description |
| --- | --- |
| **Profile** | Runs the profiler on the selected item. This is the default double-click action. To change the default action, refer to [Preferences](Preferences.md#search). |
| **Enable Deep Profiling** or **Disable Deep Profiling** | Toggles deep profiling for the selected item. |
| **Reset** | Resets all results to `0` and starts profiling again. |
| **Callstack** | Open the item’s call stack in the **Console** window. |

## Additional resources

* [Collect performance data introduction](profiler-collect-data)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)