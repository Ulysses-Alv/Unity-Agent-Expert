# Create and use a source generator

You can use source generators as an additional step in your script compilation process, to add new code while you compile your existing code. As with code analyzers, you can use existing source generators or create your own.

**Note**: Unity only supports version 6.0.0-preview of the `System.Text.Json` namespace. If you want to use this namespace in your application, ensure you use version `6.0.0-preview`. For more information about `System.Text.Json`, refer to Microsoft’s [System.Text.Json Namespace documentation](https://docs.microsoft.com/en-us/dotnet/api/system.text.json?view=net-6.0).

To create a source generator in Visual Studio and then apply it for use in your Unity project:

1. In Visual Studio, create a C# class library project that targets .NET Standard 2.0 and name the project `ExampleSourceGenerator`.
2. Install the `Microsoft.CodeAnalysis.Csharp` NuGet package for the project. Your source generator must use [Microsoft.CodeAnalysis.Csharp 4.3](https://www.nuget.org/packages/Microsoft.CodeAnalysis.CSharp/4.3.0) to work with Unity.
3. In your Visual Studio project, create a new C# file and add the following code:

   ```
   using Microsoft.CodeAnalysis;
   using Microsoft.CodeAnalysis.Text;
   using System.Text;

   namespace ExampleSourceGenerator
   {
       [Generator]
       public class ExampleSourceGenerator : ISourceGenerator
       {
           public void Execute(GeneratorExecutionContext context)
           {
               System.Console.WriteLine(System.DateTime.Now.ToString());

               var sourceBuilder = new StringBuilder(
               @"
               using System;
               namespace ExampleSourceGenerated
               {
                   public static class ExampleSourceGenerated
                   {
                       public static string GetTestText()
                       {
                           return ""This is from source generator ");

               sourceBuilder.Append(System.DateTime.Now.ToString());

               sourceBuilder.Append(
                   @""";
                       }
       }
   }
   ");

               context.AddSource("exampleSourceGenerator", SourceText.From(sourceBuilder.ToString(), Encoding.UTF8));
           }

           public void Initialize(GeneratorInitializationContext context) { }
       }
   }
   ```
4. Build your source generator for **release**. To do this, go to **Build** and select the **Batch Build** option, then select the **Release** option and **Build**.
5. In your source generator’s project folder, find the `bin/Release/netstandard2.0/ExampleSourceGenerator.dll` file.
6. Copy this file into your Unity project, inside the Assets folder.
7. Inside the Asset Browser, click on the .dll file to open the [Plugin Inspector](plug-in-inspector.md) window.
8. Under **Select platforms for plugin**, disable **Any Platform**.
9. Under **Include Platforms**, disable **Editor** and **Standalone**.
10. Under **Asset Labels**, click on the label icon to open the Asset labels sub-menu.
11. Create and assign a new label called **RoslynAnalyzer**. To do this, type `RoslynAnalyzer` in the **text input field** of the Asset labels sub-menu and press enter. This label must match exactly and is case sensitive. Once created, the label appears in the Asset labels sub-menu from then on. You can click on the name of the label in the menu to assign it to other analyzers.
12. To test the source generator is working, [create a new MonoBehaviour script](creating-scripts.md) in the Editor with the following code:

    ```
    using UnityEngine;

    public class HelloFromSourceGenerator : MonoBehaviour
    {
        static string GetStringFromSourceGenerator()
        {
            return ExampleSourceGenerated.ExampleSourceGenerated.GetTestText();
        }

        // Start is called before the first frame update
        void Start()
        {
            var output = "Test";
            output = GetStringFromSourceGenerator();
            Debug.Log(output);
        }
    }
    ```
13. If the source generator is injected into more than one assembly, your IDE displays a compiler warning about conflicts (CS0436). To remove this warning, you can do one of the following:
    * Make the `ExampleSourceGenerated` class in the `ExampleSourceGenerator` example `internal`. **Important**: in this case you can only access the generated class from within the same assembly.
    * Change the name of the `ExampleSourceGenerated` class so that it’s unique to a given assembly, for example `ExampleSourceGenerated_MyAssembly`.
14. Add this script to a **GameObject** in the **scene** and enter Play mode. A message from the source generator appears in the **Console window**, including the time stamp.

For more information about source generators, refer to Microsoft’s [Source Generators documentation](https://docs.microsoft.com/en-us/dotnet/csharp/roslyn-sdk/source-generators-overview).

## Additional resources

* [Install and use an existing analyzer or source generator](install-existing-analyzer.md)