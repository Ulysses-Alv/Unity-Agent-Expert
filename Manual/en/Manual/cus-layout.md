# Package layout for UPM packages

Explore the folder and file layout for Unity Package Manager (UPM) packages.

When you [create the package structure](cus-create.md), the Package Manager creates the following folder and file layout, which is consistent with **UPM packages** from the Unity Registry:

```
<package-root>
  ├── package.json
  ├── README.md
  ├── CHANGELOG.md
  ├── Third Party Notices.md
  ├── Editor
  │   ├── <organization-name>.<package-name>.Editor.asmdef
  │   └── EditorExample.cs
  ├── Runtime
  │   ├── <organization-name>.<package-name>.asmdef
  │   └── RuntimeExample.cs
  ├── Tests
  │   ├── Editor
  │   │   ├── <organization-name>.<package-name>.Editor.Tests.asmdef
  │   │   └── EditorExampleTest.cs
  │   └── Runtime
  │        ├── <organization-name>.<package-name>.Tests.asmdef
  │        └── RuntimeExampleTest.cs
  ├── Samples
  │        ├── Example
  │        └── ...
  └── Documentation
       └── <package-name>.md
```

| **File or folder** | **Description** |
| --- | --- |
| `package.json` | The [package manifest](upm-manifestPkg.md), which defines the package dependencies and other metadata. |
| `README.md` | Developer package documentation. This documentation helps developers who want to change the package or push a new change on the package’s main branch. |
| `CHANGELOG.md` | Description of package changes in reverse chronological order. It’s good practice to use a standard format, like [Keep a Changelog](http://keepachangelog.com/en/1.0.0/). |
| `Third Party Notices.md` | Contains information that’s required to meet [legal requirements](cus-legal.md) for third-party code or dynamic link libraries (DLL) used in your package. |
| `Editor/` | Contains assets (which can include scripts and libraries) for tools and utilities used in the Unity Editor. Unlike folders under the project’s `Assets` folder, this is only a convention and doesn’t affect the asset import pipeline. Refer to [Create or edit the assembly definitions](cus-asmdef.md) to configure Editor-specific assemblies in this folder. |
| `Runtime/` | Contains assets (which can include scripts and libraries) for tools and utilities used at runtime. This is only a convention and doesn’t affect the asset import pipeline. Refer to [Assembly definition and packages](cus-asmdef.md) to configure runtime assemblies in this folder. |
| `Tests/` | Folder to store any [tests included](cus-tests.md) in the package. |
| `Tests/Editor/` | Editor platform-specific tests folder. Refer to [Create or edit the assembly definitions](cus-asmdef.md) to configure Editor-specific test assemblies in this folder. |
| `Tests/Runtime/` | Runtime platform-specific tests. Refer to [Create or edit the assembly definitions](cus-asmdef.md) to configure runtime test assemblies in this folder. |
| `Samples/` | Folder to store any [samples included](cus-samples.md) in the package. |
| `Documentation/` | Folder to store any [documentation included](cus-document.md) in the package. |

**Note**: The **Create package** function names the `Documentation` and `Samples` folders without a trailing tilde (`~`). This differs from packages installed from the Unity registry, where the included tilde hides these folders in the **Project** window. When you [export your UPM package](cus-export.md), Unity automatically adds the tilde to the `Samples` folder to match the behavior of registry packages. For more information about hidden assets, refer to [Reserved folder name reference](SpecialFolders.md).

## Additional resources

* [Package creation](cus-pkg-lp.md)
* [Package development workflow](CustomPackages.md)
* [Reserved folder name reference](SpecialFolders.md)