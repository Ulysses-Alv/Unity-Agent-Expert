# Search box

You can use the search box at the top of the list panel in the Package Manager window to find any packages or **feature sets** that match the text you enter.

The Package Manager window updates the list as you enter search text, displaying every matching package or feature set for the context you’ve selected in the [navigation panel](upm-ui-nav.md).

**Note**: You can search for either Unity packages and feature sets or **Asset Store** packages. The list context you choose defines where the Package Manager searches:

| **List context** | **Where the Package Manager searches** |
| --- | --- |
| **In Project** | The package cache and the `Packages` folder under your current project. |
| **Unity Registry** | The list of packages and feature sets on the Unity [registry](upm-concepts.md#Registry). |
| **My Assets** | The list of [Asset Store packages](AssetPackagesOrganize.md) available with your Unity account. |
| **Built-in** | The list of [built-in](upm-concepts.md#BuiltIn) packages on the Unity registry. |
| **Services** | The list of integrated [services](UnityServices.md) you can install. |
| **My Registries** | Any [scoped registry](upm-scoped.md) defined for your project. |

**Important**: The Package Manager window recalls your search terms on a per-context basis. For example, if you set a search term in **In Project**, the Package Manager window filters that list. If you then select any other context, the Package Manager window resets the search and lists all packages for that context. But if you then return to **In Project**, the Package Manager window restores your search term and filters the **In Project** list again.

To illustrate this per-context search scope further, if you enter the same search term in separate contexts, you’ll get different results. For example, if you choose the **Unity Registry** context then type `web` in the search box, the results will be different if you then choose the **Built-in** context and search for the same term:

![Searching for the same term in separate contexts yields different results](../uploads/Main/upm-ui-search.png)

Searching for the same term in separate contexts yields different results

You can search for a package or a feature set either by:

* Its [display name](upm-manifestPkg.md#displayName) (as it appears at the top of the [details panel](upm-ui-details.md)).
* Its [name (ID)](upm-manifestPkg.md#name) (Unity packages and feature sets only), which is the registered package name that begins with **com.unity**.

Your search text can match the beginning, middle, or end of the name. The Package Manager window updates the list dynamically.

## Additional resources

* [list panel](upm-ui-list.md)