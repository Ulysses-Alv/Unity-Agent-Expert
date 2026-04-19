# Search profiler markers

The Profiler Markers search provider searches the Profiler data of **scripts** that use the `ProfilerMarker` API.

You can use the Profiler Markers search provider in two windows:

* In a dedicated window. From the main menu, select **Window** > **Analysis** > **Performance Markers**.
* In the **Search** window, using a visual or textual query with the provider token. However, the recommended best practice is using the dedicated window for profiler markers, as it provides a better experience.

## Query syntax

Provider token: `profile:`

Query example: Searches the profiler markers for the text `active` and a sample count of over `1000`.

```
profile: count>1000 active
```

## Provider filters

The Profiler Markers provider has additional filters in the [visual query builder](search-visual-query.md), using standard [search query operators](search-query-operators.md).

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

To learn more about these markers, refer to the [Profiler markers reference](profiler-markers.md).

## Results

**Search** window tab: **Profile Markers**.

The recommended best practice is to use the table view for these results.

## Actions

The context menu for the Profiler Markers search provider includes the following actions:

| Action | Description |
| --- | --- |
| **Profile** | Runs the profiler on the selected item. This is the default double-click action. To change the default action, refer to [Preferences](Preferences.md#search). |
| **Enable Deep Profiling** or **Disable Deep Profiling** | Toggles deep profiling for the selected item. |
| **Reset** | Resets all results to `0` and starts profiling again. |

## Additional resources

* [Profiler markers reference](profiler-markers.md)
* [Profiler markers](profiler-adding-information-code-intro.md#profiler-markers)
* [ProfilerMarker API reference](../ScriptReference/Unity.Profiling.ProfilerMarker.md)
* [Visual query builder reference](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)