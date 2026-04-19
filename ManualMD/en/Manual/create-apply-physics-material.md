# Create and apply a custom Physics Material

You can create as many custom [Physics Material](class-PhysicsMaterial.md) assets as you need, and apply them to **colliders** in your **scene**.

Several colliders can have the same Physics Material asset assigned to them, and you can set a project-wide default to apply new default settings to all colliders in the project.

## Create a custom Physics Material asset

To create a Physics Material asset, go to **Assets** > **Create** > **Physics Material**. By default, Unity places new Physics Material assets in your `Assets` directory.

You can create as many custom Physics Material assets as you need. Physics Material assets have the file extension `.physicMaterial`.

## Apply a custom Physics Material asset to a collider

To apply a Physics Material asset to a collider:

1. Navigate to the target collider’s **Material** property.
2. In the **Material** property field, select the picker icon.
3. Select the Physics Material asset you want to use.

Alternatively, you can click and drag the Physics Material asset file directly from the **Project window** onto the target collider’s **Material** field.

Each collider can only have one assigned Physics Material asset at a time. You can assign the same Physics Material asset to multiple colliders.

## Set a custom Physics Material asset as the project default

You can use a custom Physics Material asset to replace the project-wide default settings. Unity applies the project-wide default settings to any collider that does not have an assigned Physics Material asset.

To change the default Physics Material values:

1. Create a Physics Material asset and configure it to the default settings you want for the project.
2. Go to the Physics Settings (**Edit** > **Project Settings** > **Physics**).
3. In the **Default Material** property field, select the picker icon.
4. Select the Physics Material you want to use.

## Additional resources

* [Physics Material](class-PhysicsMaterial.md)
* [Physics Settings](class-PhysicsManager.md)