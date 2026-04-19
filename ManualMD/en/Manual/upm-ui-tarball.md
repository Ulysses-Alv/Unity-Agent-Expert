# Install a UPM package from a local tarball file

The Package Manager can load a **UPM package** from a tarball file stored locally. This is for advanced scenarios where you break your package publishing workflow into parts and your users can use the intermediate product of one of those parts.

For example, if you’ve set up continuous integration (CI) on your custom package repository, you can create a [Gzip](https://www.gnu.org/software/gzip/) tarball file from a package folder. You can create a `Gzip` tarball file by using the [npm pack](https://docs.npmjs.com/cli/pack.html) CLI or Unity Package Manager’s [Pack](../ScriptReference/PackageManager.Client.Pack.md) API. If you create a `Gzip` tarball file, test it before you publish it to a custom registry.

To install a UPM package from a local tarball file:

1. Open the [Package Manager window](upm-ui-access.md), if it’s not already open.
2. Click the **install (+)** button in the toolbar. The options for installing packages appear.
3. Select **Install package from tarball** from the install menu to bring up a file browser.
4. Go to the folder where you saved your tarball.

   **Note**: The Package Manager’s file selection dialog recognizes tarballs only if they have a `.tgz` extension.
5. Double-click the tarball file in the file selection dialog.

The file selection dialog closes, and the package now appears in the [list panel](upm-ui-list.md) with the **Local** label.

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove UPM packages or feature sets](upm-ui-actions.md)
* [Add and remove asset packages](upm-ui-actions-ap.md)