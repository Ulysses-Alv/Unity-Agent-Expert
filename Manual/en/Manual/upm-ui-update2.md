# Update an asset package

You can update an asset package that you got from the **Asset Store** by using the **Package Manager** window.

**Note:** Follow this task to update an asset package from the Asset Store that’s already in your project. Refer to the following resources for other related tasks:

* For information about importing an asset package from the Asset Store that isn’t in your project yet, refer to [Download and import an asset package](upm-ui-import.md).
* For information about updating **UPM packages** that came from the Asset Store, refer to [Switch to another version of a UPM package](upm-ui-update.md).
* For information about the different package types, refer to [Differences between package types](upm-package-types.md).

To update an asset package:

1. Open the **Package Manager** window and select **My Assets** from the [navigation panel](upm-ui-nav.md).

   An upward-pointing arrow indicator appears next to any packages that have updates available.
2. Optional: To view your list with the most recently updated Asset Store packages first, [sort the list](upm-ui-sort.md) by **Recently updated**.
3. Select the Asset Store package you want to update from the [list of packages](upm-ui-list.md). The package information appears in the [details panel](upm-ui-details.md).
4. Optional: Select the **Releases** tab to review the changes in the latest versions of the package.
5. Select the **Download update #.#** button. **Note**:

   * If the download button appears as **Download update #.#\***, the asterisk (\*) means that you have an older version of that package with the same version number in your package cache. The publisher has subsequently published an updated version of the package without changing its version number. Click that download button to get the latest version of the package.
   * If the downloaded version of your package isn’t compatible with your version of the Unity Editor, a message appears in the details panel and an alternate **Update** button displays, which lets you download a compatible version of the package.
6. To update your project with the package you downloaded in the previous step, select **Import update #.# to project**.

You can update multiple Asset Store packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages or feature sets](upm-ui-multi.md).

## Additional resources

* [Differences between package types](upm-package-types.md)
* [Asset packages](AssetPackages.md)
* [Switch to another version of a UPM package](upm-ui-update.md)