# Configure graphics jobs mode on D3D12

Unity supports multithreaded rendering through the [Graphics Jobs](playersettings-windows.md#GraphicsJobs) Player setting. For applications that use D3D12, multithreaded rendering can be implemented with three different graphics jobs modes. You can’t change the graphics jobs mode set through the [Graphics Jobs Mode](playersettings-windows.md#GraphicsJobsMode) Player setting at runtime. Moreover, your application might perform sub-optimally on certain Windows devices that don’t support specific graphics jobs modes with D3D12 API. To configure specific graphics jobs modes for various devices and ensure your application performs optimally across a range of device specifications, use the **Preferred Graphics Jobs Filter List**

To configure graphics jobs modes, use the following steps:

1. [Create a D3D12 Device Filtering Asset](create-d3d12-device-filtering-asset.md). The asset displays three filter lists in the ****Inspector**** window.
2. Expand the **Preferred Graphics Jobs Filter List**.
3. Select **Add** (**+**).
   A set of parameters is displayed.
4. Enter your preferred graphics jobs mode and the device specifications in the available parameters. All the parameters are optional.

   **Note**: Set **Preferred Graphics Jobs Mode** to **Off** to restrict certain devices from using graphics jobs.

Windows devices that meet the specifications defined in the parameter values always use the specified graphics jobs mode at runtime. Unity processes the **Preferred Graphics Jobs Filter List** entries in the order of priority with the first entry in the list having highest priority. If a device matches multiple entries in the **Preferred Graphics Jobs Filter List**, the graphics jobs mode specified in the first match applies.

The devices that don’t meet the specifications defined in the **Preferred Graphics Jobs Filter List** parameter values use the graphics jobs mode set through the [Graphics Jobs Mode](playersettings-windows.md#GraphicsJobsMode) Player setting. For example, if the Graphics Jobs Mode Player setting is set to **Native** but the preferred graphics jobs mode for certain Windows devices is set to **Split**, all the devices that don’t match the specifications mentioned in the filter list will use **Native** jobs mode.

## Additional resources

* [D3D12 Device Filtering Asset reference](d3d12-device-filter-list-asset-reference.md)