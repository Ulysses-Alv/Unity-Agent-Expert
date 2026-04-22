# Import legacy Allow and Deny Filter List values

The **Vulkan Device Filtering Asset** replaces the **Allow and Deny Filter List** properties in the Android **Player settings** that are now obsolete. These properties are available in Android Player settings only if you upgrade an existing project that previously used these properties. It is recommended to use the **Vulkan Device Filtering Asset** to configure Vulkan API usage for Android devices. The asset allows you to import the values configured in the legacy **Allow and Deny Filter Lists** of the existing project into the **Vulkan Device Filtering Asset**.

**Note**: The **Allow and Deny Filter List** properties aren’t available in the Android Player settings of new projects by default.

To import the legacy **Allow and Deny Filter List** values,

1. In your Unity project, select the Vulkan Device Filtering Asset you’ve created.
2. In the ****Inspector**** view of the asset, select **Import Legacy Player Settings Filter Lists** option.

   **Note**: This option is only available in the projects that previously used legacy **Allow and Deny Filter Lists**.

After the values are imported,

* The **Allow and Deny Filter List** properties no longer appear in the Android Player settings
* The **Import Legacy Player Settings Filter Lists** option is no longer available in the **Inspector** view of the asset.

## Additional resources

* [Configure Vulkan API usage](allow-deny-vulkan-usage.md)