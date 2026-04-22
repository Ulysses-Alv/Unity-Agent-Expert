# Add and remove UPM packages or feature sets

You can perform a variety of package management tasks by using the Package Manager window.

Follow these tasks to manage **UPM packages** or **feature sets**. For information about managing **asset packages**, refer to [Add and remove asset packages](upm-ui-actions-ap.md).

* [Install a feature set from the Unity registry](fs-install.md)
* [Install a UPM package from a registry](upm-ui-install.md)
* [Install a UPM package from Asset Store](upm-ui-install2.md)
* [Install a UPM package from a local folder](upm-ui-local.md)
* [Install a UPM package from a local tarball file](upm-ui-tarball.md)
* [Install a UPM package from a Git URL](upm-ui-giturl.md)
* [Install a UPM package from a registry by name](upm-ui-quick.md)
* [Remove a UPM package from a project](upm-ui-remove.md)
* [Switch to another version of a UPM package](upm-ui-update.md)

The procedures described in these sections obscure a lot of the details of what the Package Manager is actually doing behind the scenes. The Package Manager window provides a visual interface to install and uninstall packages. The Package Manager installs and uninstalls packages by adding and removing packages as project dependencies in your [project’s manifest](upm-manifestPrj.md). When it adds a package, it selects the correct version to install. The version that Package Manager selects doesn’t always match the version you indicated. For more information, refer to [Package dependency and resolution](upm-dependencies.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove asset packages](upm-ui-actions-ap.md)
* [Disable a built-in package](upm-ui-disable.md)
* [Perform an action on multiple packages or feature sets](upm-ui-multi.md)