# Set up Adaptive Performance

Set up Adaptive Performance in your Unity project.

To use Adaptive Performance in your project, you first need to enable it and then install a provider for your target platform. You can set up Adaptive Performance globally in the ****Project Settings**** window or for specific builds from a [build profile](../BuildSettings.md).

**Note**: The recommended best practice is to install only one provider per build on each platform. Otherwise, Unity can’t guarantee which provider takes priority at runtime.

## Set up Adaptive Performance for your project

To set up Adaptive Performance globally in your project:

1. Open the **Project Settings** window (menu: **Edit** > **Project Settings**).
2. Select **Adaptive Performance**.
3. Enable **Enable Adaptive Performance**. Unity creates standalone assets that manage the general settings of Adaptive Performance.
4. Select the checkbox next to the provider you want to install. This creates [provider settings](provider-settings-reference.md) for that provider.

Disabling Adaptive Performance deactivates the feature in both runtime and Play mode. However, the related assets stay in the project, allowing you to restore your previous settings if you re-enable.

## Set up Adaptive Performance for a specific build

You can override the global project settings for a specific build by enabling Adaptive Performance in a build profile.

To set up Adaptive Performance for a specific build:

1. Open the **Build Profiles** window (menu: **File** > **Build Profiles**).
2. Select the build profile you want to modify.
3. Select **Add Settings** > **Adaptive Performance Settings**. Unity enables Adaptive Performance and creates subassets of the build profile that manage the general settings of Adaptive Performance.

   ![The Build Profiles window with a build profile selected and Adaptive Performance settings added via the Add Settings option.](../../uploads/Main/ap/ap-buildprofile-enable.png)

   The Build Profiles window with a build profile selected and Adaptive Performance settings added via the **Add Settings** option.
4. In **Providers**, enable the checkbox next to the provider you want to install. This creates [provider settings](provider-settings-reference.md) for that provider.

These settings now apply to any build created using this build profile. Disabling the feature deactivates it and removes the subassets from the build profile.

## Additional resources

* [Introduction to Adaptive Performance](introduction.md)
* [Scripting with Adaptive Performance](scripting-api.md)
* [Providers](providers.md)
* [Check for feature support](check-feature-support.md)