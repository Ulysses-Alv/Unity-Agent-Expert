# Vulkan Device Filtering Asset reference

The Vulkan Device Filtering Asset includes filter lists each consisting of the following set of parameters. Specify values for these parameters to define Vulkan API usage and preferred graphics jobs modes for a device or set of devices.

| **Property** | **Description** |
| --- | --- |
| **Preferred Graphics Jobs Mode** | The graphics jobs mode the specified device or set of devices should use at runtime for your application. The available modes are **Off**, **Native**, **Legacy**, and **Split**. For information on these graphics jobs modes, refer to [Graphics Jobs Mode](class-PlayerSettingsAndroid.md#GraphicsJobsMode) documentation.   **Notes**:  • You can restrict devices from using graphics jobs with **Preferred Graphics Jobs Mode** set to **Off**. • This parameter is only available for the **Preferred Graphics Jobs Filter List**. |
| **Vendor** | Vendor name of the GPU hardware used in the Android device. For example, Qualcomm, ARM. |
| **Device Name** | The name of the GPU model used in the Android device. For example, Adreno (TM) 740. |
| **Brand** | The manufacturer name of the Android device. For example, google, samsung. |
| **Product Name** | The product codename for the Android device. For example, barbet, mermaid. |
| **Android OS version** | The Android operating system version the device runs on. For example, `13`. |
| **Vulkan API version** | The maximum required Vulkan API version available on the Android device. Enter the version number in either of the following formats:  • `MajorVersion.MinorVersion.PatchVersion` where MinorVersion and PatchVersion are optional. For example, `1.1.128`.  • Hexadecimal format starting with `0x` with maximum 8 digits. For example, `0x00401080`   **Note**: Regular expressions aren’t allowed for this parameter. |
| **Driver version** | The maximum required Vulkan device driver version. Enter the version number in either of the following formats:  • `MajorVersion.MinorVersion.PatchVersion` where MinorVersion and PatchVersion are optional. For example, `0.676.0`.  •Hexadecimal format starting with `0x` with maximum 8 digits. For example, `0x002a4000`   **Note**: Regular expressions aren’t allowed for this parameter. |

**Note**: Regular expressions use **ECMAScript** format for the supported properties. However, the Unity Editor and runtime use different regex engines that both support **ECMAScript**. The Unity Editor uses the `.NET` regular expression engine with the [RegexOptions.ECMAScript](https://learn.microsoft.com/en-us/dotnet/api/system.text.regularexpressions.regexoptions?view=net-10.0#system-text-regularexpressions-regexoptions-ecmascript) flag, while the runtime uses the [C++ modified ECMAScript](https://en.cppreference.com/w/cpp/regex/ecmascript.html) by default. The Editor validation confirms valid regular expression syntax, but the expression might produce different filtering behavior between the Editor and runtime. To verify runtime behavior, test the device filtering on actual devices.

## Additional resources

* [Introduction to Vulkan Device Filtering Asset](introduction-vulkan-device-filtering-asset.md)
* [Create a Vulkan Device Filtering Asset](create-vulkan-device-filtering-asset.md)
* [Configure Vulkan API usage](allow-deny-vulkan-usage.md)
* [Configure graphics jobs mode on Vulkan](configure-graphics-jobs-mode.md)
* [Regular expressions](https://learn.microsoft.com/en-us/dotnet/standard/base-types/regular-expressions)