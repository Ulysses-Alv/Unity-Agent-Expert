# Pre-release packages

When a package is ready for testing, its developer makes sure there is some basic test coverage, and that the package passes internal testing. The package also needs at least preliminary documentation, an updated changelog, and licensing.

Pre-release packages are partially tested and verified to use safely with the current version of the Unity Editor. Unity supports pre-release packages and guarantees they will be fully validated for use in production by the time the Editor reaches the long term support (LTS) phase of the current cycle. For more information about the pre-release state and the lifecycle of packages, refer to [package states and lifecycle](upm-lifecycle.md).

Pre-release packages appear in the details view with the ![](../uploads/Main/iconPrerelease.png) label and in the list view with the ![](../uploads/Main/iconPre.png) label.

By default, pre-release packages don’t appear in the list of packages you can install. However, you can force the Package Manager to display them by enabling the [Show Pre-release Package Versions](class-PackageManager.md) project setting.

**Note:** Prior to Unity Editor version 2021.1, the Package Manager used the “Preview” state to describe packages that are experimental or risky but otherwise mature, as well as packages that haven’t yet been fully validated as safe to use in production. Starting with 2021.1, the “Preview” state no longer exists, and packages can either be “Experimental” or “Pre-release”. This provides a clearer distinction between packages that are mature but risky to use, and packages that are almost fully mature.

For a list of stable packages released for this version of the Editor, refer to [Released packages](pack-safe.md).

(There are no Pre-released packages for this release.)