# Activating and enabling Profiler modules

Profiler modules collect performance data about your application. For information about Profiler modules, refer to [Profiler modules introduction](profiler-modules-introduction.md).

Some Profiler modules have a large data collection overhead, such as the GPU, UI, and audio Profiler module. To prevent these modules from affecting your application’s performance, you can deactivate them.

## Activating and deactivating Profiler modules

To activate or deactivate a module:

1. Open the Profiler window (**Window** > **Analysis** > **Profiler**)
2. Select the **Profiler Modules** dropdown
3. Enable or disable the modules you want to display

![Profiler windows Profiler Modules dropdown menu](../uploads/Main/profiler-modules-dropdown.png)

Profiler window’s Profiler Modules dropdown menu

Disabling a module removes the module from the window, stops the Profiler from collecting that module’s data, and lowers the overhead of the Profiler. When you enable a Profiler module, it starts collecting data, but shows no data for the period in which it wasn’t active.

**Important:** The [CPU Usage module](ProfilerCPU.md) always collects data even when it’s not active, because other modules rely on it.

## Reordering Profiler modules

To change the order that the Profiler modules appear in the window, use the [Profiler Module Editor](profiler-module-editor.md):

1. Open the Profiler window (**Window** > **Analysis** > **Profiler**)
2. Select the **Profiler Modules** dropdown
3. Select the cog icon (⚙). The Profiler Module Editor opens in its own window.
4. Click and drag the handle (≡) next to the module’s name to the desired order.
5. Select **Save Changes** and close the window

![Profiler Module Editor window](../uploads/Main/profiler-module-editor-empty.png)

Profiler Module Editor window

The Profiler window now displays the new order of Profiler modules.

## Additional resources

* [Profiler modules introduction](profiler-modules-introduction.md)
* [Navigating the Profiler window](profiler-window-navigating.md)
* [Profiler window reference](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)