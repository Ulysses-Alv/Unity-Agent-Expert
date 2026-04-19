# Global cache

When the Unity Package Manager fetches **UPM packages**, it stores the package contents and metadata in a global cache. This makes reusing and sharing packages more efficient, and allows you to install and update stored packages even when offline.

**Note**: Storing packages in the global cache applies to UPM packages fetched from a registry. Packages in the UPM format that come from the **Asset Store** are also stored in the global cache. However, [asset packages](AssetPackages.md) in the `.unitypackage` format that come from the Asset Store aren’t stored in the global cache. The Package Manager stores asset packages in a separate cache. For more information, refer to [Asset Store packages](AssetStorePackages.md).

## Location

By default, Unity stores the global cache in a root directory that depends on the operating system (and the user account type on Windows):

| **Operating system** | **Default root directory** | **Example** |
| --- | --- | --- |
| Windows (user account) | `%LOCALAPPDATA%\Unity\cache\upm` | `C:\Users\yourname\AppData\Local\Unity\cache\upm` |
| Windows ([system user account](https://docs.microsoft.com/en-us/windows/security/identity-protection/access-control/local-accounts#default-local-system-accounts)) | `%ALLUSERSPROFILE%\Unity\cache\upm` | `C:\ProgramData\Unity\cache\upm` |
| macOS | `$HOME/Library/Caches/Unity/upm` | `/Users/yourname/Library/Caches/Unity/upm` |
| Linux | `$HOME/.cache/Unity/upm` | `/home/yourname/.cache/Unity/upm` |

**Tip**: You can override the location of this root directory. Refer to [Customize the global cache](upm-config-cache.md) for more information.

## Structure

The Package Manager global cache uses subsidiary caches, each serving a different purpose. The Package Manager stores these subsidiary caches in subfolders under the folder of the global cache:

| **Subfolder** | **Description** |
| --- | --- |
| `db` | Registry data cache used for storing package content and metadata. |
| `git-lfs` | Contains downloaded Git Large File Storage (LFS) files, if you’ve enabled Git LFS. |
| `packages` (obsolete) | If you’ve created projects with Unity Editor 2023.2, this subfolder might exist. However, starting with Unity 6.0, the Package Manager doesn’t use this subfolder. |

Inside each of these subfolders, each registry has its own path so that packages hosted on different registries aren’t mixed up.

**Tip**: You can override the location of these folders. Refer to [Customize the global cache](upm-config-cache.md) for more information.

## Size

Starting with version 2023.2.0f1 of the Unity Editor, the size of the registry data cache (the `db` subfolder) is limited to 10 gigabytes (GB). When that limit is reached, the Package Manager prunes the registry data cache by evicting the least recently used content (based on oldest date a package was installed into a project). In other words, the first packages the Package Manager evicts are the ones you added to a project the longest time ago. However, although the Package Manager evicts these packages from the cache, they remain in any projects where you installed them. The next time you add these packages to a different project, the Package Manager will fetch them from the appropriate registry, rather than from your cache.

You can override the size of the registry data cache limit. For information, refer to [Customize the global cache](upm-config-cache.md).

## Requirements

The user account running the Unity Editor process must have full write permissions on the root directory and its contents. Without these permissions, the Package Manager can’t download and save the package metadata and contents in the cache.

## Additional resources

* [Package types](upm-package-types.md)