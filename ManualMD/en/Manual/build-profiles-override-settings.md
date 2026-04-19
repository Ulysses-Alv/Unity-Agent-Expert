# Override settings with build profiles

You can override specific Unity Editor settings for each **build profile** in your project, allowing you to create unique settings profiles for all your target platforms.

**Note**: Build profile overrides are used in Play mode and within the Unity Player.

## Diagnostics

To override the default **Diagnostics** settings established in **Project Settings**, follow these steps:

1. Go to **File** > **Build Profiles**.
2. Select or create a build profile for your target platform. For information on how to create a build profile, refer to [Create a build profile](create-build-profile.md).
3. Under **Diagnostics**, select the required option from the **Diagnostic Data** dropdown. For more information, refer to the [Build Profile window reference](build-profiles-reference.md#diagnostics).

Your build profile now uses the assigned value and overrides the **Project Settings**. To revert to using the value from **Project Settings**, set **Diagnostic Data** to **Use Project Settings → Diagnostics**.

## Player settings

To override Player settings, follow these steps:

1. Navigate to **File** > **Build Profiles**.
2. Select or create a build profile for your target platform. For information on how to create a build profile, refer to [Create a build profile](create-build-profile.md).
3. In the **Player Settings Overrides** section, select **Customize player settings**.

The [Player settings](class-PlayerSettings.md) for the build profile’s target platform will appear, allowing you to customize them as required. The Player settings inherit their initial values from your build targets global Player settings. To update these global settings, navigate to **Edit** > **Project Settings** > **Player**.

**Note**: For an active build profile, the Player Settings Overrides are linked to the [Player settings APIs](../ScriptReference/PlayerSettings.md). If you use the Player Settings APIs to modify a Player setting for an active build profile, the change will update the corresponding override value.

### Remove and reset Player settings overrides

You can remove and reset any overridden Player settings for your build profile using the available options from the **More** (**⋮**) menu. The available options are as follows:

| **Property** | **Description** |
| --- | --- |
| **Remove Overrides** | Remove and reset all overridden Player settings. To add new overrides, you must select **Customize player settings**. |
| **Reset To Globals** | Reset overridden Player setting values to the global values for the target platform. Set the global values from **Edit** > **Project Settings** > **Player**. |

## Graphics settings

Use the **Graphics Settings** section in the **Build Profiles** window to override specific [Graphics](class-GraphicsSettings.md) and [Quality](class-QualitySettings.md) project settings.

### Override graphics settings

To override graphics settings, follow these steps:

1. Navigate to **File** > **Build Profiles**.
2. Select or create a build profile for your target platform. For information on how to create a build profile, refer to [Create a build profile](create-build-profile.md).
3. In the **Graphics Settings** section, enable **Override Global Graphics Settings.**
4. Open the **Override Options** dropdown.
5. Override the available [graphics](class-GraphicsSettings.md) settings as required. The settings inherit their initial values from the global settings located in **Edit** > **Project Settings** > **Graphics**.

You can reset any adjusted settings back to the global settings using the **More** (**⋮**) > **Reset to Globals** option. To remove all graphics overrides, disable **Override Global Graphics Settings.**

### Override quality settings

To override quality settings, follow these steps:

1. Navigate to **File** > **Build Profiles**.
2. Select or create a build profile for your target platform. For information on how to create a build profile, refer to [Create a build profile](create-build-profile.md).
3. In the **Graphics Settings** section, enable **Override Global Quality Settings.**
4. Open the **Included Quality Levels** dropdown.

![A list of available quality levels in the Graphics Settings section of the Build Profiles window.](../uploads/Main/build-profiles-quality-levels.png)

A list of available quality levels in the Graphics Settings section of the Build Profiles window.

The list inherits quality levels for your target platform from the global [Quality](class-QualitySettings.md) levels located in **Edit** > **Project Settings** > **Quality**.

**Note**: When in Play mode, the quality level used is the selected level in the global quality settings.

To add a new quality level to your build profile, you must first add a new quality level to the global quality settings. Once created, add the new level to your build profile with the **Add** (**+**) button in the **Included Quality Levels** section. Right-click (macOS: **Ctrl**+ click) on any quality level in the build profile to set it as the default.

**Note**: Renaming or deleting a quality level in the global quality settings updates any build profile that contains them.

You can reset any adjusted levels back to the global settings using the **More** (**⋮**) > **Reset to Globals** option. To remove all quality settings overrides, disable **Override Global Quality Settings.**

## Additional resources

* [Create a build profile](create-build-profile.md)
* [Project Settings](comp-ManagerGroup.md)