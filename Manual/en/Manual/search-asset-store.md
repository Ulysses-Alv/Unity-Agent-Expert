# Search the Unity Asset Store

The **Asset Store** search provider finds items in the [Unity Asset Store](https://assetstore.unity.com/).

**Tip:** You can use the Asset Store provider in the general **Search** window, or in a specialised window. To open a specialized search window for the Asset Store provider, from the main menu, select **Window** > **Search** > **Asset Store**.

## Query syntax

Provider token: `store:`

**Note:** You can’t implicitly use this search provider [as part of a default search](search-filter-by-provider.md). You must explicitly use its search token or its [visual query](search-visual-query.md) filter.

Query example: Search for 3D character assets with a price of less than $100 and a minimum rating of 5.

```
store: price=100 min_rating=5 category=3d/characters
```

## Provider filters

The Asset Store provider uses the asset’s store information as filters in the [visual query builder](search-visual-query.md). These tables list the filters and their textual query equivalents.

### Types

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Category** | `store: category=<category>` | `store: category=audio/music` | Searches for music. |
| **Platform** | `store: platform=<platform>` | `store: platform=standaloneosxuniversal` | Searches for assets that support macOS. |

### Pricing

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Maximum price** | `store: price=<number>` | `store: price=20` | Searches for assets with a maximum price of $20. |
| **On sale** | `store: on_sale=<boolean>` | `store: on_sale=true` | Searches for assets that are on sale. |
| **Free** | `store: free=<boolean>` | `store: free=true asteroid` | Searches for assets that are free and that have the word `asteroid` in their name. |

### Release information

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Rating** | `store: rating=<number>` | `store: min_rating=3` | Searches for assets with a rating of three or higher. |
| **Release (days since)** | `store: release=<range>` | `store: release=OneYearAgo` | Searches for assets released in the past year. Possible values: `OneDayAgo`, `OneWeekAgo`, `OneMonthAgo`, `SixMonthsAgo`, and `OneYearAgo`. |
| **Version** | `store: version=<number>` | `store: version=2022` | Searches for assets that support Unity 2022 and newer. |
| **Publisher** | `store: publisher=<name>` | `store: publisher=Unity` | Searches for assets published by the company Unity. Note that the publisher name must be exact: `unity` finds Unity Technologies, but `uni` doesn’t. |

## Results and action

**Search** window tab: Asset Store.

### Sorting

You can add a sorting parameter to your query to sort the results.

Possible values: `price`, `rating`, `popularity`, `name`, and `relevance`.

Query example: Sort the results by price.

```
store: sort=Price
```

You can only sort by one parameter at a time; you can’t nest sorting parameters.

### Actions

Default action: Open the Unity Asset Store on the asset’s page.

## Additional resources

* [Unity Asset Store](https://assetstore.unity.com/)
* [Unity’s Asset Store documentation](AssetStore.md)
* [Visual query builder](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)