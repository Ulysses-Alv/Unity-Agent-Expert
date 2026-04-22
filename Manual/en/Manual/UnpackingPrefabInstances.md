# Revert a prefab instance to a GameObject

If you want to convert a **prefab** instance into a regular [GameObject](GameObjects.md), you can use the **Unpack** command. Unity then unpacks the prefab instance into GameObjects but retains the original prefab asset.

To use the **Unpack** command to convert a prefab instance into a GameObject, perform the following steps:

1. In the [Hierarchy](Hierarchy.md), right-click on the prefab instance.
2. Select **Prefab** > **Unpack**.

Unity converts the prefab instance into a GameObject which no longer has any links to the original prefab asset. Unity also saves any [overrides](PrefabInstanceOverrides.md) on the prefab instance into the GameObject.

The **Unpack** command only unpacks the selected prefab instance, and has the following effects:

* If you unpack a prefab instance that has [nested prefabs](NestedPrefabs.md), the nested prefabs remain as prefab instances inside the new GameObject.
* If you unpack a [prefab variant](PrefabVariants.md), the root becomes a new prefab instance of the base prefab asset of the variants. Unity converts all properties of the variant which differ from the original prefab asset into overrides.

To completely unpack a prefab instance and its nested prefabs, or to completely unpack a variant to its base set of GameObjects, select **Prefab** > **Unpack Completely**.

## Additional resources

* [Create prefabs](CreatingPrefabs.md)
* [Nest prefab instances in other prefabs](NestedPrefabs.md)
* [Create variations of prefabs](PrefabVariants.md)