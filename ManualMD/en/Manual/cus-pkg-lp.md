# Package creation

Create packages that deliver Unity Editor or runtime tools to support game development workflows.

This information focuses on creating Unity Package Manager (UPM) packages. **UPM packages** typically include functional content, such as custom Editor windows, runtime tools, or utilities designed for use during development, rather than content-only assets like models or textures. These tools are often designed to streamline workflows for game designers, developers, or technical artists within a Unity project.

UPM packages offer modularity, versioning, and reuse across projects. UPM packages support a wide range of content, including C# **scripts**, managed assemblies, and built-in **plug-ins**. While the primary purpose is tool delivery, packages can also include supporting assets like animations, textures, **audio clips**, or models when necessary.

For more information about the differences between UPM packages and **asset packages** (`.unitypackage` format), refer to [package types](upm-package-types.md). If you want to create an asset package rather than a UPM package, refer to [create and export asset packages](AssetPackagesCreate.md).

| **Topic** | **Description** |
| --- | --- |
| [Get started with package creation](cus-pkg-intro.md) | Get a high-level overview of creating UPM packages. |
| [Comparison of package creation locations](cus-location.md) | Choose whether to create your package inside or outside of the project’s `Packages` folder. |
| [Package development workflow](CustomPackages.md) | Develop your package by supplying the required and recommended components. |
| [Export and sign your UPM package](cus-export.md) | Export and sign your package so you can share it with others. |
| [Share your package](cus-share.md) | Learn the different ways you can share your UPM package when it’s ready for consumption. |

## Additional resources

* [Package types](upm-package-types.md)
* [Create and export asset packages](AssetPackagesCreate.md)