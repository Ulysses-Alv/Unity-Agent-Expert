# Roslyn analyzers and source generators

Use Roslyn analyzers, source generators and ruleset files in Unity projects to inspect your code for style, quality, and other issues.

You can use existing analyzer libraries to inspect your code, and write your own analyzers to promote the best practices or conventions within your organization.

**Note**: Roslyn analyzers are only compatible with the IDEs that Unity publicly supports, which are Visual Studio and JetBrains Rider.

For more information about how to write and use Roslyn analyzers, refer to Microsoft’s [Analyzer Configuration](https://docs.microsoft.com/en-us/visualstudio/code-quality/use-roslyn-analyzers?view=vs-2019) and [Get started with Roslyn analyzers](https://docs.microsoft.com/en-us/visualstudio/extensibility/getting-started-with-roslyn-analyzers?view=vs-2019) documentation.

| **Topic** | **Description** |
| --- | --- |
| [Create and use a source generator](create-source-generator.md) | Create a simple source generator and configure the Unity Editor to use it for your Unity project code. |
| [Install and use an existing analyzer or source generator](install-existing-analyzer.md) | Download an existing code analyzer or source generator and configure the Unity Editor to use it for your Unity project code. |
| [Additional files for Roslyn analyzers and source generators](roslyn-analyzers-additional-files.md) | Define additional text files for Roslyn analyzers or source generators to use. |
| [Create rule set files](ruleset-files.md) | Create rule sets to customize the standards your code analysis tools enforce. |
| [Code analyzer scope and diagnostics](analyzer-scope-and-diagnostics.md) | Control which parts of your code are subject to code analysis and get insights into how long analysis is taking. |

## Additional resources

* [C# compiler](csharp-compiler.md)
* [Debugging with Roslyn analyzers](https://unity.com/how-to/debugging-with-rosyln-analyzers)