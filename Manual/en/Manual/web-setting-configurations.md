# Web setting configurations

Choose a configuration when you create a **build profile** to apply preconfigured settings to the new profile. These presets optimize your build for development or release on mobile or desktop devices.

For descriptions of all available Web settings, refer to [Web build settings](web-build-settings.md) and [Web Player settings](class-PlayerSettingsWebGL.md). For the build profile settings shared across platforms, refer to [Build Profiles window reference](build-profiles-reference.md).

**Note**: The [Facebook Instant Games platform](instant-games.md) has the same configurations as the Web platform, with any exceptions noted on this page.

## Configurations

Configurations are groups of default settings that tailor builds for specific delivery environments. Use configurations to create multiple optimized build profiles and simplify the process of switching between different environments.

Select a configuration in the **Platform Browser** when you [create a build profile](create-build-profile.md). If you don’t choose a configuration, the build profile inherits the values from the global settings.

## Configuration defaults

The Web platform has the following configurations:

* Mobile - Development
* Mobile - Release
* Desktop - Development
* Desktop - Release

### Shared defaults for configurations

All configurations have the following default settings.

#### Player Settings Overrides

| **Property** | **Default** |
| --- | --- |
| **IL2CPP** Code Generation | Optimize for code size and build time |
| Target WebAssembly 2023 | Enabled |
| Use WebAssembly.Table | Enabled |
| Enable BigInt | Enabled |

### Mobile and desktop defaults

The mobile and desktop configurations differ as follows:

#### Platform Settings

| **Property** | **Mobile** | **Desktop** |
| --- | --- | --- |
| Texture **compression** format | ASTC | DXT |

#### Player Settings Overrides

**Note**: These overrides apply only to Facebook Instant Games build profiles.

| **Property** | **Mobile** | **Desktop** |
| --- | --- | --- |
| **WebGL** template | Default (Portrait) | Landscape |

### Release and development defaults

The release and development configurations differ as follows:

#### Platform Settings

| **Property** | **Development** | **Release** |
| --- | --- | --- |
| **Development Build** | Enabled | Disabled |
| Code Optimization | N/A | Disk Size with LTO |

#### Player Settings Overrides

| **Property** | **Development** | **Release** |
| --- | --- | --- |
| C++ Compiler Configuration | Debug | Master |
| Compression Format | Disabled | Brotli |
| Data Caching | Disabled | Enabled |

## Additional resources

* [Introduction to build profiles](build-profiles.md)
* [Create and manage build profiles](create-build-profile.md)
* [Override settings with build profiles](build-profiles-override-settings.md)