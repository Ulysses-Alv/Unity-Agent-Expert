# Create a D3D12 Device Filtering asset

To create a D3D12 Device Filtering asset:

1. In your Unity project, navigate to the main menu.
2. Select **Assets** > **Create** > **D3D12 Settings** > **Device Filter**.

Unity creates the D3D12 Device Filtering Asset in the **Assets** folder with the file extension `.d3d12devicefilter`.

Each asset displays three filter lists in the ****Inspector**** window: **Allow Filter List**, **Deny Filter List**, and **Preferred Graphics Jobs Filter List**. The **Deny Filter List** includes Windows device specifications that Unity restricts by default from using the D3D12 API. If your application performs better with these device specifications, you can remove an item from the list.

## Additional resources

* [Introduction to D3D12 Device Filtering Asset](introduction-d3d12-device-filtering-asset.md)
* [Configure D3D12 API usage](allow-deny-d3d12-usage.md)
* [Configure graphics jobs mode on D3D12](configure-d3d12-graphics-jobs-mode.md)