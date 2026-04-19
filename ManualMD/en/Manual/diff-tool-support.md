# Diff tool support

You can use the [**Revision Control Diff/Merge** setting](Preferences.md#external-tools) to set an installed diff tool as the default revision tool. You can also use this setting to define a custom revision tool with specific layouts.

If you want to change the diff tool that Unity uses, open the [Preferences](Preferences.md) window, and navigate to the **External Tools** section. Select your preferred tool from the **Revision Control Diff/Merge** dropdown list.

## Set up a custom revision tool

To set up a custom revision tool, follow these steps:

1. Open the [Preferences](Preferences.md) window, and navigate to the **External Tools** section.
2. In the **Revision Control Diff/Merge** dropdown list, select **Custom Tool**.
3. Enter the path to the custom tool’s installation folder. On macOS, this should point to the *Contents / MacOS* folder in the tool’s installation folder.
4. Enter the arguments for two-way diffs, three-way diffs, and merges.

To specify file layout in the revision tool, use these arguments:

| **Property** | **Function** |
| --- | --- |
| `#LTITLE` | Left title |
| `#RTITLE` | Right title |
| `#ATITLE` | Ancestor title |
| `#LEFT` | Left file |
| `#RIGHT` | Right file |
| `#ANCESTOR` | Ancestor file |
| `#OUTPUT` | Output file |
| `#ABSLEFT` | Absolute path to the left file |
| `#ABSRIGHT` | Absolute path to the right file |
| `#ABSANCESTOR` | Absolute path to the ancestor file |
| `#ABSOUTPUT` | Absolute path to the output file |

Examples:

![SourceGear DiffMerge](../uploads/Main/SourceGearDiffMerge.png)

SourceGear DiffMerge

![P4Merge](../uploads/Main/P4DiffMerge.png)

P4Merge

## Additional resources

* [Preferences reference](Preferences.md)
* [Version control integrations](Versioncontrolintegration.md)