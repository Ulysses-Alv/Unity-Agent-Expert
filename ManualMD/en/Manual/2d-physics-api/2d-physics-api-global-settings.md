# Configure global LowLevelPhysics2D API settings

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

To configure global settings for the `LowLevelPhysics2D` API, create and assign a Physics Low Level Settings 2D asset. Follow these steps:

1. In the **Project** window, select **Assets** > **Create** > **2D** > **Physics LowLevel Settings**.
2. From the main menu, select **Edit** > **Project Settings**.
3. In the ****Project Settings**** window, select **Physics 2D**, then the **Low Level** tab.
4. Drag the **Physics LowLevel Settings 2D** asset from the **Project** window to the **Physics Low Level Settings** setting.

The 2D physics system now uses the settings in the asset.

## Configure the settings

For more information about configuring each section of the settings, refer to the following:

* [Configure collisions between LowLevelPhysics2D API objects](2d-physics-api-collisions-enable.md) for using and editing the **Layers** setting.
* [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md) for editing the **Default Definitions** settings.
* [Physics Low Level Settings window reference](2d-physics-api-reference-low-level-settings) for information about the **Global** settings.

## Additional resources

* [Configure LowLevelPhysics2D API objects with definitions](2d-physics-api-definitions.md)
* [Physics Low Level Settings 2D asset Inspector window reference](../class-PhysicsLowLevelSettings2D.md)