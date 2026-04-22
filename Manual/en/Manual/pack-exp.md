# Experimental packages

Experimental packages are new packages or experimental modifications made to mature packages. Unity doesn’t support Experimental packages because they’re in the early stages of development.

**Note:** Before Unity Editor version 2021.1, the Package Manager used the “Preview” state to describe packages that are experimental or risky but otherwise mature. The Package Manager used the “Preview” state to describe packages that had not yet been fully validated as safe to use in production. Starting with 2021.1, the “Preview” state no longer exists, and packages can either be “Experimental” or “Pre-release.” This provides a clearer distinction between packages that are mature but risky to use, and packages that are almost fully mature.

Experimental packages can go through many changes before they’re ready for release in a specific version of Unity. At some point in the future, they might pass [the verification requirements](pack-safe.md) but they might also become deprecated instead. Because there is no guarantee for future support, you shouldn’t use experimental packages in production.

Packages in experimental state don’t usually appear in the [Unity Registry](upm-ui-nav.md#contexts) context of the Package Manager, even though they’re on Unity’s official package registry server. These packages aren’t discoverable in the Package Manager window because:

* They’re too risky to use in production. Some of these packages require a lot of training and expertise and are recommended only in specific circumstances.
* They provide shared or additional functionality to existing packages. You shouldn’t use them on their own because they’re “support” packages only.

Experimental packages that aren’t discoverable can still appear in the Package Manager window if you already installed them in your project or installed them as dependencies of supported packages. Otherwise, they’re hidden so that you don’t discover them by accident and use them without realizing the risks. If they do appear in the Editor, they’re always marked in the Package Manager window with the ![Experimental](../uploads/Main/iconExperimental.png) label (details view) and the ![Exp](../uploads/Main/iconExp.png) label (list view). Also, the Editor toolbar displays the following warning state:

![The Experimental Packages In Use tooltip appears as a warning in the Editor toolbar](../uploads/Main/upm-lifecycle.png)

The **Experimental Packages In Use** tooltip appears as a warning in the Editor toolbar

If you have experimental packages in your project, the Package Manager icon in the Editor toolbar displays in a warning state, with an **Experimental Packages In Use** tooltip. Click the icon to open the Package Manager with a filter applied, which lists the packages in experimental state.

For a list of stable packages verified for this release, refer to [Released packages](pack-safe.md).

For more information about package states, refer to [Package states and lifecycle](upm-lifecycle.md).