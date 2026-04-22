# Removing unused override data

An [override](PrefabInstanceOverrides.md) on a **prefab** might become unused in the following situations:

* The override’s target object is invalid.
* The override’s [property path](../ScriptReference/PropertyModification-propertyPath.md) is unknown. This might happen if you delete a public field definition or rename it. To preserve overrides when renaming fields, you can use [`FormerlySerializedAsAttribute`](../ScriptReference/Serialization.FormerlySerializedAsAttribute.md).

Unity stores the unused override data in the **scene** file, which means that if you restore the deleted script or field definition, Unity reapplies the override data. Unity doesn’t automatically clean up unused data because you might have moved the object or property the data refers to temporarily or in error. It’s best practice to clean up unused override data to ensure that scene files only contain relevant data, which makes **version control** and collaboration easier.

## Remove unused overrides

To check for and remove unused overrides, perform the following steps:

1. In the prefab instance **Inspector**, select [**Overrides**](prefab-instance-inspector-reference.md).
2. Select **Unused overrides**. If the prefab instance doesn’t have any unused overrides, this section is unavailable.
3. The **Unused overrides** panel displays a list of the unused overrides. Select **Remove** to remove them.

Unity writes details of the removed overrides to the [Editor log](log-files.md).

![Unused overrides panel, displaying unused overrides on the selected prefab.](../uploads/Main/prefabs-unused-overrides.png)

Unused overrides panel, displaying unused overrides on the selected prefab.

You can also remove unused overrides from the Hierarchy:

* Right-click on a prefab instance and select **Prefab > Remove Unused Overrides**.
* Right-click on the scene name and select **Prefab > Remove Unused Overrides** to remove all unused overrides in the entire scene.

A dialog appears to confirm the removal of the overrides.

## Additional resources

* [Override prefab instances](PrefabInstanceOverrides.md)
* [Prefab instance reference](prefab-instance-inspector-reference.md)