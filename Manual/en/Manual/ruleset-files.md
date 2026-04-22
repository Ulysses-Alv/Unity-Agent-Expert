# Create rule set files

To define your own rules on how to handle the various warnings and errors that the analyzers in your project raise, you can create a ruleset file. For more information on how to create a custom ruleset, refer to Microsoft’s Visual Studio documentation on [how to create a custom rule set](https://docs.microsoft.com/en-us/visualstudio/code-quality/how-to-create-a-custom-rule-set?view=vs-2019).

In the Assets root folder, place a ruleset file named Default.ruleset. The rules you define in Default.ruleset apply to all predefined assemblies (for example `Assembly-CSharp.dll`), and all assemblies that are built using .asmdef files.

To override the rules in Default.ruleset for a predefined assembly, create a .ruleset file in the root folder with the name `[PredefinedAssemblyName].ruleset`. For example, the rules in `Assembly-CSharp.ruleset` apply to the code in `Assembly-CSharp.dll`. Only these .ruleset files are allowed inside the root folder:

* `Default.ruleset`
* `Assembly-CSharp.ruleset`
* `Assembly-CSharp-firstpass.ruleset`
* `Assembly-CSharp-Editor.ruleset`
* `Assembly-CSharp-Editor-firstpass.ruleset`

### Workflow: Testing ruleset files in Unity

To test ruleset files in Unity, follow these steps:

#### Step 1: Set up the ruleset file

1. Create a subfolder named “Subfolder” inside your project’s Assets folder.
2. Inside **Subfolder**:
   1. Create a new .asmdef file.
   2. Save a duplicate copy of `RethrowError.cs`.
3. Create a Default.ruleset file inside **Assets** with the following code:

```
<?xml version="1.0" encoding="utf-8"?>
<RuleSet Name="New Rule Set" Description=" " ToolsVersion="10.0">
  <Rules AnalyzerId="ErrorProne.NET.CodeAnalyzers" RuleNamespace="ErrorProne.NET.CodeAnalyzers">
    <Rule Id="ERP021" Action="Error" />
  <Rule Id="EPC12" Action="None" />
  </Rules>
</RuleSet>
```

The `Default.ruleset` file defines the following rules:

* Suppress EPC12, the warning about suspicious exception handling.
* Elevate ERP021, the warning about incorrect exception propagation, to an error.

#### Step 2: Reload the project

After you add the ruleset files to your project, reimport any script that lives in an assembly where the rules should apply. This forces Unity to recompile the assembly using the new ruleset files. After recompilation, two messages will appear in the **Console window**:

`Assets\Subfolder\RethrowError.cs(15,19): error ERP021: Incorrect exception propagation. Use throw; instead.`

`Assets\RethrowError.cs(15,19): error ERP021: Incorrect exception propagation. Use throw; instead.`

Notice that Unity applies the rules defined in Default.ruleset to both `Assets/RethrowError.cs` and `Assets/Subfolder/RethrowError.cs`.

#### Step 3: Add a custom ruleset

In Assets/Subfolder, create a .ruleset file, and give it any name you like (in this `exampleHello.ruleset`):

```
<?xml version="1.0" encoding="utf-8"?>
<RuleSet Name="New Rule Set" Description=" " ToolsVersion="10.0">
  <Rules AnalyzerId="ErrorProne.NET.CodeAnalyzers" RuleNamespace="ErrorProne.NET.CodeAnalyzers">
    <Rule Id="ERP021" Action="Info" />
    <Rule Id="EPC12" Action="Info" />
  </Rules>
</RuleSet>
```

This new `Hello.ruleset` file tells Unity to print both EPC12 and ERP021 to the Console, without treating them as warnings or errors.

After Unity compiles the project again, the following messages will appear in the Console window:

`Assets\Subfolder\RethrowError.cs(14,23): info EPC12: Suspicious exception handling: only e.Message is observed in exception block.`

`Assets\Subfolder\RethrowError.cs(15,19): info ERP021: Incorrect exception propagation. Use throw; instead.`

`Assets\RethrowError.cs(15,19): error ERP021: Incorrect exception propagation. Use throw; instead.`

The rules in `Default.ruleset` still apply to `Assets\RethrowError.cs`, but they no longer apply to `Assets\Subfolder\RethrowError.cs`, because the rules in `Hello.ruleset` override them.

For more information on all the allowed ruleset action files, refer to the Visual Studio documentation on [Using the code analysis rule set editor](https://docs.microsoft.com/en-us/visualstudio/code-quality/working-in-the-code-analysis-rule-set-editor?view=vs-2019).

## Additional resources

* [Special folders and script compilation order](script-compile-order-folders.md)