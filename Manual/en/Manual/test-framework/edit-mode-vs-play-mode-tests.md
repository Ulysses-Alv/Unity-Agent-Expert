# Edit mode and Play mode tests

Unity Test Framework identifies tests as either Edit mode tests or Play mode tests depending on the references of their parent assembly.

## Edit mode tests

Edit mode tests (also known as Editor tests) only run in the Unity Editor and have access to Editor code and runtime application code. As such, Edit mode test assemblies can reference code in both the `UnityEditor` and `UnityEngine` namespaces.

With Edit mode tests you can test any of your [Editor extensions](../ExtendingTheEditor.md) using the [`[UnityTest]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityTestAttribute.html) attribute. Edit mode tests run in the [EditorApplication.update](https://docs.unity3d.com/ScriptReference/EditorApplication-update.html) callback loop. You can’t run [coroutines](https://docs.unity3d.com/ScriptReference/Coroutine.html) in Edit mode tests.

You can also control entering and exiting Play mode from an Edit mode test, allowing your test to make changes before entering Play mode.

Edit mode tests must have an [assembly definition](workflow-create-test-assembly.md) that references `nunit.framework.dll` and have the Editor as their only target platform:

```
assembly
    "includePlatforms": [
        "Editor"
    ],
```

## Play mode tests

You can run Play mode tests [in a Player](workflow-run-playmode-test-standalone.md) or inside the Editor. Play mode tests allow you to test your runtime application code, and the tests run as [coroutines](https://docs.unity3d.com/ScriptReference/Coroutine.html) if marked with the `[UnityTest]` attribute.

Play mode tests must fulfill the following conditions:

* Tests must have their own [assembly definition](workflow-create-test-assembly.md) with a reference to `nunit.framework.dll`.
* Test **scripts** must be in a folder alongside the .asmdef file.
* The test assembly must reference any additional assemblies containing code you want to test.

```
assembly
    "references": [
        "NewAssembly"
    ],
    "optionalUnityReferences": [
        "TestAssemblies"
   ],
    "includePlatforms": [],
```

**Note**: Your test assembly can’t reference the predefined `Assembly-Csharp.dll` assembly. You must move code you want to test into a custom assembly that your test assembly references. For more information, refer to [Creating assembly assets](../assembly-definitions-creating.md).

## Recommendations

Use the NUnit [`[Test]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/test.html) attribute instead of the [`[UnityTest]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityTestAttribute.html) unless:

* You need to [yield instructions for the Editor](reference-custom-yield-instructions.md) in Edit mode tests
* You need to skip a frame or wait for a certain amount of time in Play mode tests.

## Additional resources

* [Create a test assembly](workflow-create-test-assembly.md)
* [Create a test](workflow-create-test.md)