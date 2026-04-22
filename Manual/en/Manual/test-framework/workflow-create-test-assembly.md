# Create a test assembly

Unity Test Framework tests must be in a test assembly, which is any assembly that references NUnit. For more information on assemblies, refer to [Organizing scripts into assemblies](../assembly-definition-files.md).

To create a new test assembly in the ****Test Runner**** window:

1. Select the `Assets` folder in your **Project window**.
2. Open the **Test Runner** window (menu: **Window** > **General** > **Test Runner**).
3. In the **Test Runner** window, select **Create a new Test Assembly Folder in the active path**.

![The Test Runner window with the EditMode tab selected.](../../uploads/Main/test-framework/create-test-assembly-folder.png)

The Test Runner window with the EditMode tab selected.

Alternatively, you can create a test assembly directly from the **Assets** menu:

1. Select the `Assets` folder in your Project window.
2. Create a new Test Assembly Folder (menu: **Assets** > **Create** > **Testing** > **Test Assembly Folder**).

This creates an `Tests` subfolder inside the `Assets` folder with a corresponding `.asmdef` file including the required references. You can change the name of the new assembly definition and press Enter to accept it.

![The Project window displays a new subfolder of the Assets folder called Tests, which contains an assembly file called Tests.asmdef.](../../uploads/Main/test-framework/test-assembly-folder.png)

The Project window displays a new subfolder of the Assets folder called Tests, which contains an assembly file called Tests.asmdef.

Click on the assembly definition file to inspect it in the ****Inspector**** window. You’ll see that it has references to `nunit.framework.dll`,`UnityEngine.TestRunner`, and `UnityEditor.TestRunner` assemblies. This combination of references is what identifies an assembly as a test assembly. The `UnityEditor.TestRunner` reference is only available for Edit mode tests.

The checkbox selections under **Platforms** determine which platforms the test assembly can run on. By default, assemblies created through the **Test Runner** target the **Editor** only. Selecting **Any Platform** or a specific platform other than **Editor** makes it possible to run any Play mode tests in the assembly on standalone Players for the additional platforms.

![The Inspector window displays the configurable properties of a test assembly definition asset.](../../uploads/Main/test-framework/import-settings.png)

The Inspector window displays the configurable properties of a test assembly definition asset.

You can repeat the steps above as many times as you like to create additional test assemblies. The first test assembly folder you create is named `Tests` by default and subsequent ones are named `Tests 1`, `Tests 2`, and so on. Remember that you can always rename the assemblies but each assembly name must be unique.

**Note**: Changing the file name of the assembly definition file doesn’t affect the value of the **Name** property in the file. Use the **Inspector** window or edit the .asmdef file directly in a text editor to make sure the name property is properly changed.

## Additional resources

* [Edit mode and Play mode tests](edit-mode-vs-play-mode-tests.md)
* [Create a test](workflow-create-test.md)