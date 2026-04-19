# Create a Vulkan Device Filtering Asset

To create a Vulkan Device Filtering asset:

1. In your Unity project, navigate to the main menu.
2. Select **Assets** > **Create** > **Vulkan Settings** > **Device Filter**.

Unity creates the Vulkan Device Filtering Asset in the **Assets** folder with the file extension `.vulkandevicefilter`.

Each asset displays three filter lists in the ****Inspector**** window: **Allow Filter List**, **Deny Filter List**, and **Preferred Graphics Jobs Filter List**. The **Deny Filter List** includes Android device specifications that Unity restricts by default from using the Vulkan API. If your application performs better with these device specifications, you can override those by removing from the list.

## Additional resources

* [Introduction to Vulkan Device Filtering Asset](introduction-vulkan-device-filtering-asset.md)
* [Configure Vulkan API usage](allow-deny-vulkan-usage.md)
* [Configure graphics jobs mode on Vulkan](configure-graphics-jobs-mode.md)