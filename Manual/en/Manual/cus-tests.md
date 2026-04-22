# Add tests to your package

Add tests to your Unity Package Manager (UPM) package to verify its behavior and identify issues during development.

Adding tests is relevant if you’re creating a package that contains **scripts** or logic you want to validate as your package evolves. Including tests helps you catch bugs, improve reliability, and support safe reuse of the package across projects. Unity’s test framework lets you organize and run these tests from within the Unity Editor or as part of your continuous integration (CI) pipeline.

When you [create your package](cus-create.md) with the Package Manager window, Package Manager creates `Tests/Editor` and `Tests/Runtime` folders with C# and assembly definition (`.asmdef`) files in each, to help you get started.

To complete the test files for your package:

1. Go to your package folder, either by using the **Project** window or your file management application.
2. Go to the `Tests/Editor` or `Tests/Runtime` folder for the tests you want to complete.
3. Open the sample C# script.
4. Add test cases as needed. A good practice is to create at least one test file for each C# script. Refer to [Writing tests](test-framework/writing-tests.md).
5. Save your file.
6. If you create a more elaborate folder structure in your `Tests` folder, you can optionally create additional `asmdef` files for the code you want to put in separate assemblies. If you choose to create additional assembly definitions, refer to [Assembly definition files for tests](#files).
7. [Enable tests](#tests) for your package.
8. Run your tests. For more information, refer to [Running tests](test-framework/running-tests.md).

## Assembly definition files for tests

Use the [Test Framework](test-framework/test-framework-introduction.md) to create or edit your assembly definition files. For more information, refer to [Create a test assembly](test-framework/workflow-create-test-assembly.md).

When you [create your package](cus-create.md) with the Package Manager window, Package Manager creates assembly definition files (`.asmdef`) in `Tests/Editor` and `Tests/Runtime` to accompany the test scripts it created in those folders.

The sample assembly definition files have the required properties and values. If you want to add optional properties and values, use the ****Inspector**** window to edit the assembly definitions. For more information, refer to [Assembly Definition properties reference](class-AssemblyDefinitionImporter.md).

You can also edit `.asmdef` files directly, but this method is more error prone. For more information, refer to [Assembly Definition File Format reference](assembly-definition-file-format.md).

### Editor file example

The `.asmdef` file in the `Tests/Editor` folder looks like this:

```
{
  "name": "MyCompany.MyPackage.Editor.Tests",
  "references": [
    "MyPackage.Editor",
    "MyPackage"
  ],
  "optionalUnityReferences": [
    "TestAssemblies"
  ],
  "includePlatforms": [
    "Editor"
  ],
  "excludePlatforms": []
}
```

### Runtime file example

The `.asmdef` file in the `Tests/Runtime` folder looks like this:

```
{
  "name": "MyCompany.MyPackage.Tests",
  "references": [
    "MyPackage"
  ],
  "optionalUnityReferences": [
    "TestAssemblies"
  ],
  "includePlatforms": [],
  "excludePlatforms": []
}
```

## Enable tests for a package

You might need to enable tests depending on the folder where you’re developing your package.

If you’re developing your package in the `Packages` folder of your project, you don’t need to explicitly enable tests.

However, if you’re developing a package outside the project’s `Packages` folder, you need to manually enable its tests. To do this, add the [testables](upm-manifestPrj.md#testables) property to your **project manifest** and list the packages whose tests you want to run. If you also want to run tests from packages they depend on, include those **indirect dependencies** as well. For example:

```
{
  "dependencies": {
    "com.unity.some-package": "1.0.0",
    "com.unity.other-package": "2.0.0",
    "com.unity.yet-another-package": "3.0.0"
  },
  "testables": ["com.unity.some-package", "com.unity.other-package"]
}
```

This example adds tests for the `com.unity.some-package` and `com.unity.other-package` packages in Unity’s [Test Framework](test-framework/test-framework-introduction.md) package.

## Additional resources

* [Testing your code](test-framework/test-framework-introduction.md)
* [Organizing scripts into assemblies](assembly-definition-files.md)
* [Assembly Definition properties reference](class-AssemblyDefinitionImporter.md)
* [Comparison of package creation locations](cus-location.md)
* [Package dependency and resolution](upm-dependencies.md)