# Configuration files

You can use a configuration file to store custom settings for the Package Manager.

For example, you can use configuration files to:

* Override the location or size of the global cache.
* Store your token information for each scoped registry that requires authentication.

Package Manager supports two configuration files: a global configuration file and a user configuration file. Both of these files use the [TOML](https://toml.io) format and they appear in different locations:

* **Global** configuration files apply to all users on the same computer. For example, you can define additional [SSL certificate authorities](upm-config-network.md#SSL) when setting up a proxy server for the entire computer.
* **User** configuration files apply to a single user. For example, you can set up [authentication tokens](upm-config-scoped.md) to use for custom package registry servers that you access with scoped registries. These tokens authenticate a specific user account.

## Global configuration file location

Package Manager uses a global configuration file named `upmconfig.toml`. This file isn’t automatically created when you install the Unity Hub or a Unity Editor. However, you can create the file in the following location if you need to customize your configuration:

| **Environment** | **Location** |
| --- | --- |
| Windows | `%ALLUSERSPROFILE%\Unity\config\upmconfig.toml` (for example, `C:\ProgramData\Unity\config\upmconfig.toml`) |
| macOS and Linux | `/etc/upmconfig.toml` |

You can define a custom location that overrides the default location for your configuration file. To do this, create a `UPM_GLOBAL_CONFIG_FILE` environment variable and set its value to the absolute path of your configuration file, including the file name.

## User configuration file location

Package Manager uses a user configuration file named `.upmconfig.toml`. This file isn’t created when you install the Unity Hub or Editor, but you can create it in the following location if you need to customize your configuration:

| **Environment** | **Location** |
| --- | --- |
| Windows (user account) | `%USERPROFILE%\.upmconfig.toml` (for example, `C:\Users\myusername\.upmconfig.toml`) |
| Windows ([system user account](https://docs.microsoft.com/en-us/windows/security/identity-protection/access-control/local-accounts#default-local-system-accounts)) | `%ALLUSERSPROFILE%\Unity\config\ServiceAccounts\.upmconfig.toml` (for example, `C:\Users\Public\Unity\config\ServiceAccounts\.upmconfig.toml`) |
| macOS and Linux | `~/.upmconfig.toml` (for example, `/Users/myusername/.upmconfig.toml`) |

You can define a custom location that overrides the default location for your configuration file. To do this, create a `UPM_USER_CONFIG_FILE` environment variable and set its value to the absolute path of your configuration file, including the file name.

## Additional resources

* [Package Manager configuration](upm-config.md)