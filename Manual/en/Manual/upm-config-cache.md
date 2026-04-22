# Customize the global cache

The Package Manager maintains a [global cache](upm-cache.md) for compressed package contents and registry metadata.

The global cache is separate from the cache that the Package Manager maintains for asset packages. For information about the **asset package** cache, refer to [Customize the asset package cache location](upm-config-cache-as.md).

By default, the Package Manager’s global cache root has the following subdirectories, with each subsidiary cache serving a different purpose:

* The `db` subdirectory stores registry data (package metadata and tarballs).
* If you’ve enabled Git LFS, the `git-lfs` subdirectory stores downloaded Git Large File Storage (LFS) files.

**Note**: If you’ve created projects with Unity Editor 2023.2, your global cache might also have a `packages` subdirectory. Starting with Unity 6.0 Beta, the Package Manager doesn’t use this subdirectory. If you no longer maintain projects created with older versions of the Editor, you can delete this subdirectory.

The Package Manager stores the global cache in a default location, which you can override. For information on the default location of the global cache root, refer to [global cache](upm-cache.md).

You might want to override the global cache location for several reasons:

* You want to save space on your internal drive.
* You want to store the cache on a shared drive, so others can access it.
* You want to store the cache in a folder that’s not in a system directory.

Starting with 2023.2.0f1, the maximum size of the `db` subdirectory defaults to 10 gigabytes (GB), which you can override. You might want to override the maximum size of the `db` subdirectory for several reasons:

* If you have a small amount of available disk space, you can lower the limit to save space.
* If you have a large amount of available disk space, you can increase the limit. This lets the Package Manager store more data in the cache, so it can fetch packages from the cache instead of downloading them again. Cached packages install faster and reduce network usage.

The following table illustrates the different methods you can use to override the default properties:

| **Method** | **Global cache location** | **`db` subdirectory max size** | **Additional notes** |
| --- | --- | --- | --- |
| [Preferences window](#pref-method) | Yes | No | This is the recommended method because of its ease and permanence. When you change the package cache location using this method, Unity stores the preference for the current and future Unity Editor sessions. |
| [User configuration file](#cacheconfigFile) | Yes | Yes | Package cache property overrides in a user configuration file persist beyond the current Unity Editor session. |
| [Environment variables](#cacheconfigVar) | Yes | Yes | By default, this method sets the configuration for the current command prompt (or terminal) session, but you can make it permanent in your operating system settings. Advanced users might prefer session-based variables in specific use cases. |

**Note**: Using the **Preferences** window to set the global cache location also sets the `cacheRoot` property in a user configuration file. If you want to set the user configuration file manually, refer to [Using the user configuration file](#cacheconfigFile).

## Important notes before you begin

* Regardless of the method you choose, changing the cache location tells the Package Manager to use that location going forward. Existing packages that Package Manager stored in the original location remain in that folder structure.
* If you set the maximum size of the `db` subdirectory and the new size is less than the current size of that subdirectory, the Package Manager will evict content from the cache until the new limit is reached. The Package Manager evicts the least recently used content first (based on the installation date into a project).
* When you use a user configuration file, your change applies to the current user account. Although a global configuration file exists, you can’t set cache overrides in the global configuration file.
* There are several factors that affect how setting several values with different methods work together:
  + If you used the environment variable method to change the cache location, then you can’t use the **Preferences** window to change the location.
  + Setting an environment variable takes precedence over applying the same setting in the user configuration file or the **Preferences** window.
  + If you set an environment variable for the Git LFS cache (`UPM_GIT_LFS_CACHE_PATH`), Package Manager uses the location specified by the environment variable, regardless of where the global cache root is.
  + If you enable the Git LFS cache (`UPM_ENABLE_GIT_LFS_CACHE`) without explicitly setting its path (`UPM_GIT_LFS_CACHE_PATH`), Package Manager uses the `git-lfs` folder under the global cache root as its location.

## Using the Preferences window

To use the **Preferences** window to override the default location of the global cache, follow these steps.

1. Use one of the following methods to open the **Preferences** window:

   * Use the Unity Editor’s menus, as described in [Preferences](Preferences.md).
   * Open the Package Manager window, open the **More** (⋮) menu, and select **Preferences**.
2. Select the **Package Manager** category.
3. Under **Packages**, open the folder menu beside **Cache Location**.
4. Choose **Change Location**.
5. Choose a new location for the global cache.

![The Preferences window with the Package Manager category selected](../uploads/Main/preferences-package-manager.png)

The Preferences window with the Package Manager category selected

## Using the user configuration file

You can manually edit the user configuration file to override the location of the global cache root, or the maximum size of the `db` subdirectory, or both.

Follow these steps to edit the user configuration file:

1. Locate the `.upmconfig.toml` user configuration file. To find the file location, refer to [Configuration files](upm-config-files.md). If the file doesn’t exist, create an empty text file.
2. To override the location of the global cache root, add the `cacheRoot` key and set its value as an absolute path to a custom location. **Important**: When setting Windows paths in TOML files, use forward slashes (`/`) or double backslashes (`\\`). Don’t use single backslashes (`\`) because they’re special characters which mark escape sequences and can cause TOML parsing errors.
3. To override the maximum size of the `db` subdirectory, add the `maxCacheSize` key and set its value to a positive integer representing the number of bytes. For example, to set the size to 5 GB, assign a value of `5000000000`.

You don’t need to restart the Unity Editor or the Hub after setting either property for the change to take effect.

**Windows example**

```
cacheRoot = "E:\\Unity\\cache"
maxCacheSize = 5000000000
```

**macOS and Linux example**

```
cacheRoot = "/dev/external/shared/Unity/cache"
maxCacheSize = 5000000000
```

## Using environment variables

In scenarios that involve automation or continuous integration, it’s less practical and more error prone to configure settings in the user configuration file or the **Preferences** window. In such scenarios, you might consider using environment variables to configure settings for the global cache and its subsidiary caches. The following environment variables are available:

| **Environment variable** | **Description** |
| --- | --- |
| `UPM_CACHE_ROOT` | Specify the location of the [global cache](upm-cache.md) (the root folder containing the registry data and uncompressed packages caches). |
| `UPM_NPM_CACHE_PATH` | Obsolete. Use `UPM_CACHE_ROOT` instead. |
| `UPM_ENABLE_GIT_LFS_CACHE` | Specify any value except empty ("") to enable the Git LFS cache. By default, Package Manager uses the `git-lfs` subdirectory under the global cache root, unless you specify a different path using the `UPM_GIT_LFS_CACHE_PATH` environment variable. |
| `UPM_GIT_LFS_CACHE_PATH` | Specify the absolute path to the location where you want Package Manager to store [the Git LFS cache](upm-git.md#git-lfs). By setting this path, you automatically enable the Git LFS cache. |
| `UPM_MAX_CACHE_SIZE` | Specify the maximum size of the registry data cache (the `db` subdirectory) as a positive integer, in bytes. For example, to set the size to 5 GB, assign a value of `5000000000`. |

**Important**: Follow these guidelines when you use this method:

* Close the Unity Editor and Unity Hub if they’re already running before setting environment variables.
* Launch the Unity Editor or Unity Hub from the same command prompt or terminal session where you set the environment variables.
* If you set any environment variables in the preceding table, you must set them every time you launch Unity.

For information on setting environment variables, refer to the documentation for your operating system. For an introduction to environment variables, refer to <https://en.wikipedia.org/wiki/Environment_variable>.

## Additional resources

* [Configuration files](upm-config-files.md)
* [Customize the Asset Store cache location](upm-config-cache-as.md)