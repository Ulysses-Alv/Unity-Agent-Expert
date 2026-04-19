# Update the package changelog

Update the changelog file (`CHANGELOG.md`) each time you publish a new version of your package.

Use the changelog to summarize new features, major fixes, and significant changes in each release. The changelog explains to users what’s changed between versions to help them decide whether to upgrade.

Maintaining a changelog is optional for packages that you don’t share, but strongly recommended for packages you share or distribute. Users rely on this information to know which version best suits their needs.

When you [create your package](cus-create.md) using the Package Manager window, the Unity Editor creates the changelog file for you.

To update the changelog:

1. Open the `CHANGELOG.md` file in the base folder of your package.
2. Document key additions, fixes, and breaking changes. Aim for clarity so users understand what changed and why.
3. Save the file.
4. (Optional) Link to an external changelog. If you host the changelog on a website or repository, [edit the package manifest](cus-edit-manifest.md#edit-manifest) by adding a `changelogUrl` property. For example:

   ```
   "changelogUrl": "https://example.com/your-package/changelog"
   ```

For more information about changelogs, refer to the [Keep a Changelog](http://keepachangelog.com) documentation.

## Additional resources

* [Package development workflow](CustomPackages.md)