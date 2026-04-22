# D3D12 Device Filtering Asset reference

The D3D12 Device Filtering Asset includes filter lists each consisting of the following set of parameters. Specify values for these parameters to define D3D12 API usage and preferred graphics jobs modes for a device or set of devices.

| **Property** | **Description** |
| --- | --- |
| **Preferred Graphics Jobs Mode** | The graphics jobs mode the specified device or set of devices should use at runtime for your application. The available modes are **Off**, **Native**, **Legacy**, and **Split**. For information on these graphics jobs modes, refer to [Graphics Jobs Mode](class-PlayerSettingsAndroid.md#GraphicsJobsMode) documentation.   **Notes**:  • To restrict devices from using graphics jobs, set **Preferred Graphics Jobs Mode** to **Off**. • This parameter is only available for the **Preferred Graphics Jobs Filter List**. |
| **Vendor** | Vendor name of the GPU hardware used in the Windows device. For example, `NVIDIA`. |
| **Device Name** | The name of the GPU model used in the Windows device. For example, `NVIDIA RTX 3080`. |
| **Driver Version** | The GPU device driver version. Enter the version number in the following format:  • `MajorVersion.MinorVersion.PatchVersion.MinorPatchVersion` where MinorVersion, PatchVersion and MinorPatchVersion are optional. For example, `30.0.13044.17006`.  **Note**: Regular expressions aren’t allowed for this parameter. |
| **D3D12 Feature Level** | The required D3D12 feature level available on the Windows device. Enter the version number in the following format:  • `MajorVersion.MinorVersion` where MinorVersion is optional. For example, `12.1`.  **Note**: Regular expressions aren’t allowed for this parameter. |
| **Graphics Memory** | The amount of video memory available on the GPU in megabytes. For example, `8192`. |
| **Processor Count** | The number of CPU processor cores on the Windows device. For example, `8`. |
| **Device Type** | Whether the GPU is discrete or integrated. Select **Discrete**, **Integrated**, or **Do Not Care**. The **Do Not Care** option means that the **Device Type** parameter is not applied to the filter. |

**Note**: Regular expressions use **ECMAScript** format for the supported properties. However, the Unity Editor and runtime use different regex engines that both support **ECMAScript**. The Unity Editor uses the `.NET` regular expression engine with the [RegexOptions.ECMAScript](https://learn.microsoft.com/en-us/dotnet/api/system.text.regularexpressions.regexoptions?view=net-10.0#system-text-regularexpressions-regexoptions-ecmascript) flag, while the runtime uses the [C++ modified ECMAScript](https://en.cppreference.com/w/cpp/regex/ecmascript.html) by default. The Editor validation confirms valid regular expression syntax, but the expression might produce different filtering behavior between the Editor and runtime. To verify runtime behavior, test the device filtering on actual devices.

## Additional resources

* [Introduction to D3D12 Device Filtering Asset](introduction-d3d12-device-filtering-asset.md)
* [Create a D3D12 Device Filtering Asset](create-d3d12-device-filtering-asset.md)
* [Configure D3D12 API usage](allow-deny-d3d12-usage.md)
* [Configure graphics jobs mode on D3D12](configure-d3d12-graphics-jobs-mode.md)
* [Regular expressions](https://learn.microsoft.com/en-us/dotnet/standard/base-types/regular-expressions)