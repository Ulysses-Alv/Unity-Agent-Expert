# Managed code stripping

During the build process, Unity removes unused or unreachable code through a process called managed code stripping, which can significantly decrease your application’s final size. Managed code stripping removes code from managed assemblies, including assemblies built from the C# **scripts** in your project, assemblies that are part of packages and plugins, and assemblies in .NET Framework.

Unity uses a tool called the Unity linker to perform a static analysis of the code in your project’s assemblies. The static analysis identifies any classes, portions of classes, functions, or portions of functions that can’t be reached during execution. This analysis only includes code that exists at build time because code generated at runtime doesn’t exist when Unity performs the static analysis.

You can configure the level of code stripping Unity performs for your project in the Unity Editor. You can also provide annotations in the form of attributes or a special XML configuration file to instruct the Unity linker about which parts of your code base to preserve.

| **Topic** | **Description** |
| --- | --- |
| **[Managed code stripping and the Unity linker](unity-linker.md)** | Understand how the Unity linker analyzes and marks elements of your code before stripping. |
| **[Configure managed code stripping](managed-code-stripping-configure.md)** | Configure managed code stripping in the Unity Editor and understand what Unity does at each stripping level. |
| **[How code stripping affects content](managed-code-stripping-content.md)** | Understand how your managed code stripping configuration can affect content builds and how to preserve code your content depends on. |
| **[Preserving code using annotations](managed-code-stripping-preserving.md)** | Use annotations in the form of C# attributes or a `link.xml` file to specify which parts of your code the Unity linker should not strip. |
| **[Link XML formatting reference](managed-code-stripping-xml-formatting.md)** | Understand how to format a `link.xml` file, including supported XML elements and attributes and examples of their usage. |
| **[Unity linker marking rules reference](managed-code-stripping-marking-rules.md)** | Understand the marking rules the Unity linker applies at each managed stripping level to determine which parts of your code assemblies to strip. |

## Additional resources

* [Scripting backends](scripting-backends.md)
* [Organizing scripts into assemblies](assembly-definition-files.md)