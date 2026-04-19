# Finding package documentation

The Unity Manual provides documentation for a specific version of Unity. For package documentation, each package has its own version, so each package provides documentation for a specific version of that package. For this reason, the package documentation isn’t part of the main Unity Manual documentation. Instead, documentation for each package lives on its own website on Unity’s documentation server.

To access the documentation for a specific package, you have two options:

* [Get documentation for the latest version](#Manual) (from the Unity Manual).
* [Get documentation for a specific package version](#Editor) (from the Package Manager window).

When the package page opens, the top of the page has four links **(A)** and a drop-down menu **(B)** where you can select a different version:

![Unitys package documentation](../uploads/Main/upm-docs.png)

Unity’s package documentation

Use the links at the top of the page to view the documentation for the package’s core usage (**Manual**), the package’s **Scripting API** documentation, the **Changelog** for the package, and the **License** information for the package.

## Getting documentation for the latest version

Each version of the Unity Manual documentation provides a list of [released](pack-safe.md) and [pre-release](pack-preview.md) packages, a list of [built-in packages](pack-build.md), and a [list of packages by keyword](pack-keys.md). Each list has links to the matching package documentation which you can use with that version of Unity. **Note**: Only packages that are compatible with the version of Unity matching the documentation appear in these lists.

![List of Unity packages available in the manual](../uploads/Main/upm-docs-list.png)

List of Unity packages available in the manual

You can expand the entries in the sidebar navigation **(A)** to access specific information about each package. Information includes the list of compatible versions, a brief description of the package, and any associated keywords. Or you can click the link on the list pages **(B)** to directly open the most compatible version of that package documentation that matches the documentation version.

If you can’t find the package you want in this list, there might be several reasons:

* It isn’t compatible with this version of Unity.
* It’s an experimental or private package. Some packages aren’t available to everyone because someone outside of Unity is developing them or because they’re under a special license.

You might be able to access the documentation [through the Package Manager window](#Editor) for packages released in another version of Unity. For private packages, try to contact the developer directly to request access.

## Getting documentation for a specific package version

You can find out more about a package that comes from a registry (not Asset Store) by viewing its documentation, changelog, or license information.

To access any of these pages, you can click the **Documentation**, **Changelog**, or **Licenses** links from inside the Package Manager window.

To access the documentation for a specific package version:

1. Open the [Package Manager window](upm-ui-access.md) in Unity.
2. Follow the guidelines in [Finding packages](upm-ui-find.md) to locate the specific package and version number you want.
3. Click the **Documentation** link to open the documentation for the selected version of the package.

   **Note**: If the package is installed, you can right-click the **Documentation**, **Changelog**, or **Licenses** links and select **Open Locally**. The Package Manager opens your file browser at the appropriate folder.

You can follow this procedure for any version of a package that comes from a registry (not Asset Store). Note that the documentation isn’t necessarily different for each package version release, because some version updates (patches) involve only bug fixes or trivial changes.