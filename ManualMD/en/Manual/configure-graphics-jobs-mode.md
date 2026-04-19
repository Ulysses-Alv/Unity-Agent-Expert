# Configure graphics jobs mode on Vulkan

Unity supports multithreaded rendering through the [Graphics Jobs](class-PlayerSettingsAndroid.md#GraphicsJobs) Player setting. For applications using Vulkan API, multithreaded rendering can be implemented with three different graphics jobs modes. The graphics jobs mode set through the [Graphics Jobs Mode](class-PlayerSettingsAndroid.md#GraphicsJobsMode) Player setting can’t be changed at runtime. Moreover, your application might perform sub-optimally on certain Android devices that don’t support specific graphics jobs modes with Vulkan API. Using **Preferred Graphics Jobs Filter List**, you can configure specific graphics jobs modes for various devices to ensure your application performs optimally across a range of device specifications.

To configure graphics jobs modes, use the following steps:

1. [Create a Vulkan Device Filtering Asset](create-vulkan-device-filtering-asset.md). The asset displays three filter lists in the ****Inspector**** window.
2. Use the foldout (triangle) to expand the **Preferred Graphics Jobs Filter List**.
3. Select the **Add** (**+**) button to enter the specifications of the Android device for which you want to use a specific graphics jobs mode.
   A set of parameters is displayed.
4. Enter your preferred graphics jobs mode and the device specifications in the available parameters. All the parameters are optional.

   **Note**: Set **Preferred Graphics Jobs Mode** to **none** to restrict certain devices from using graphics jobs.

Android devices that meet the specifications defined in the parameter values will always use the specified graphics jobs mode at runtime. Unity processes the **Preferred Graphics Jobs Filter List** entries in the order of priority with the first entry in the list having highest priority. If a device matches multiple entries in the **Preferred Graphics Jobs Filter List**, the graphics jobs mode specified in the first match applies.

The devices that don’t meet the specifications defined in the **Preferred Graphics Jobs Filter List** parameter values, will use the graphics jobs mode set through the [Graphics Jobs Mode](class-PlayerSettingsAndroid.md#GraphicsJobsMode) Player setting. For example, if the Graphics Jobs Mode Player setting is set to **Native** but the preferred graphics jobs mode for certain Android devices is set to **Split**, all the devices that don’t match the specifications mentioned in the filter list will use **Native** jobs mode.

## Additional resources

* [Vulkan Device Filtering Asset reference](vulkan-device-filter-list-asset-reference.md)