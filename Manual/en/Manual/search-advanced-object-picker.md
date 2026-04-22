# Use the Advanced Object Picker

The **Advanced Object Picker** window is similar to the [**Search** window](search-overview.md) in the following ways:

* It supports visual queries.
* It lists results in four tabs: All, Project, Resources, and Hierarchy. To learn more, refer to [Search Providers](search-filter-by-provider.md).
* It has three views: grid, list, and table.
* It can display an ****Inspector**** panel for an item, which allows you to edit the object’s properties before selecting it as a reference.

![The Advanced Object Picker window, showing a visual query with two parameters, and an Inspector panel for the selected item.](../uploads/Main/search-advanced-object-picker.png)

The Advanced Object Picker window, showing a visual query with two parameters, and an Inspector panel for the selected item.

## Use the Advanced Object Picker in all reference fields

To use the **Advanced Object Picker** in all reference fields in the Unity Editor:

1. From the main menu, go to **Edit** > **Preferences** (macOS: **Unity** > **Settings**).
2. In the **Preferences** window, go to the **Search** tab.
3. Under **Search Engines**, from the **Object Selector** dropdown, select **Advanced**.

   To revert to the default object selector, select **Classic** from the dropdown.
4. Close the **Preferences** window.

## Assign the Advanced Object Picker to individual property

To add the `SearchContextAttribute` to specific object properties, refer to [SearchContextAttribute](../ScriptReference/Search.SearchContextAttribute.md).

## Additional resources

* [Search window](search-overview.md)
* [Search Providers](search-filter-by-provider.md)
* [SearchContextAttribute](../ScriptReference/Search.SearchContextAttribute.md) reference