# Configure managed code stripping

The **Managed Stripping Level** property determines which [rules](managed-code-stripping-marking-rules.md) the Unity linker follows when it analyzes and strips your application’s code. Increasing the managed stripping level expands the scope of the linker’s search for unused code but also increases build time.

To change the **Managed Stripping Level** property:

1. Go to **Edit** > **Project Settings** > **Player**.
2. In **Other Settings**, navigate to the Optimization heading.
3. Set the **Managed Stripping Level** property to the desired value.

| **Property** | **Function** |
| --- | --- |
| **Disabled** | Unity doesn’t remove any code.  This setting is only available for the [Mono scripting backend](scripting-backends-mono.md) and is the default setting in that case. |
| **Minimal** | Unity searches only the `UnityEngine` and the .NET class libraries for unused code. Unity doesn’t remove any user-written code. This setting is the least likely to cause unexpected runtime behavior.  This setting is useful for projects where usability is of higher priority than build size. This is the default setting if you use the [IL2CPP scripting backend](scripting-backends-il2cpp.md). |
| **Low** | Unity searches for unused code in all `UnityEngine` and .NET class libraries. It also searches user-written assemblies, but only if none of their types are referenced in scenes included in the Player build. This setting applies a set of rules that removes some unused code but minimizes the likelihood of unintended consequences, such as changes in behavior of runtime code that uses reflection.  **Note**: The **Low** managed stripping level is marked for future deprecation and using it is not recommended. Use **Minimal** or **Medium** instead. |
| **Medium** | Unity partially searches all assemblies to find unused code. This setting applies a set of rules that strips more types of code patterns to reduce the build size. Although Unity doesn’t strip all possible unused code, this setting does increase the risk of undesirable or unexpected behavior changes. |
| **High** | Unity performs an extensive search of all assemblies to find unused code. At this setting, Unity prioritizes size reduction more than code stability and removes as much code as possible.  This search can take much longer than for lower stripping levels. Use this setting only for projects where a compact build size is extremely important. Test your application thoroughly and make careful use of `[Preserve]` attributes and `link.xml` files to ensure that the Unity linker doesn’t strip vital code. |

## Additional resources

* [Managed code stripping and the Unity linker](unity-linker.md)
* [Preserving code using annotations](managed-code-stripping-preserving.md)
* [Link XML formatting reference](managed-code-stripping-xml-formatting.md)
* [Unity linker marking rules reference](managed-code-stripping-marking-rules.md)
* [IUnityLinkerProcessor.GenerateAdditionalLinkXmlFile](../ScriptReference/Build.IUnityLinkerProcessor.GenerateAdditionalLinkXmlFile.md)