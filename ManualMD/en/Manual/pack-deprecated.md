# Deprecated packages

Packages that reach their end of life are no longer supported in Unity Editors where they’re labeled as deprecated. You shouldn’t use deprecated packages, because they might be nonfunctional or unsafe to use.

There are two deprecated labels, to distinguish two variants:

| **Deprecation variant** | **Label** | **Description** |
| --- | --- | --- |
| Unity Package Lifecycle deprecation | Deprecated - yellow | This label identifies a package that’s no longer under development for that Unity Editor version. Unity will no longer offer support for packages that reach this state. |
| Package version deprecation | Deprecated - red | This label identifies a specific version of the package that’s deprecated. Unity usually applies this designation when they discover a critical issue in a specific version of a package that hasn’t yet reached its end of life. Custom package authors can also apply this designation based on their own deprecation flow. |

If your project has deprecated packages of either variant, a **Deprecated packages** dialog appears and asks if you want to open Package Manager. If you open the Package Manager, the Package Manager windows loads with the deprecated package displayed. At this point, you can:

* Remove a package that’s reached lifecycle deprecation. Although you might be able to install a different version of the package, the deprecation state still applies. Consider finding another package to use instead of this deprecated package.
* Remove a package whose specific version is deprecated. Look for another version of the same package that’s compatible with your Editor but not deprecated.

## Additional resources

* [Package states and lifecycle](upm-lifecycle.md)