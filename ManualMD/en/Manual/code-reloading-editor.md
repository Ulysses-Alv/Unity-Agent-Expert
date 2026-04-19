# Code and scene reload on entering Play mode

Authoring your application in Edit mode and then switching to Play mode to preview its runtime behavior is a core feature of iterative development in the Unity Editor. By default the Editor reloads both your code and **scene** assets as part of the transition from Edit mode to Play mode. It’s important to understand what Unity reloads, why and when it does so, and how you can configure the reloading behavior. This helps you make informed trade-offs between development iteration time and the degree to which Play mode accurately reflects your built application’s performance.

| **Topic** | **Description** |
| --- | --- |
| [Configuring how Unity enters Play mode](configurable-enter-play-mode.md) | Configure the Unity Editor to enter Play mode more quickly and improve your development iteration times. |
| [Enter Play mode with domain reload disabled](domain-reloading.md) | Understand how disabling domain reload on enter Play mode affects your application state and how you can compensate for these effects in your code. |
| [Enter Play mode with scene reload disabled](scene-reloading.md) | Understand how disabling scene reload on enter Play mode affects your application and how you can compensate for these effects in your code. |
| [Details of disabling domain and scene reload](configurable-enter-play-mode-details.md) | Understand the work Unity performs during domain and scene reload and what’s skipped when they’re disabled. |

## Additional resources

* [Script compilation](script-compilation.md)
* [Scripting back ends](scripting-backends.md)