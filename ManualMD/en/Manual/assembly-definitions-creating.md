# Creating assembly assets

You can create and define assemblies using two special kinds of asset:

* Assembly Definition (`.asmdef`) assets are placed in the root of a folder containing **scripts** to define a new assembly.
* Assembly Definition Reference (`.asmref`) assets are placed in the root of a folder containing scripts to include those scripts in a preexisting assembly.

## Create an Assembly Definition

To create an Assembly Definition:

1. In the **Project** window, locate the folder containing the scripts you want to include in the assembly.
2. Create an Assembly Definition in the folder (menu: **Assets** > **Create** > **Scripting** > **Assembly Definition**).
3. Assign a name to the asset. By default, the assembly file uses the name you assign to the asset, but you can change the name in the ****Inspector**** window.

Unity recompiles the scripts in the project to create the new assembly. You can then edit the [properties](class-AssemblyDefinitionImporter.md) of the new Assembly Definition.

The new assembly includes all scripts in the same folder as the Assembly Definition plus those in any subfolders that don’t have their own Assembly Definition or Reference file. Unity removes scripts from their previous assembly where applicable.

## Create an Assembly Definition Reference

To create an Assembly Definition Reference:

1. In the **Project** window, locate the folder containing the scripts you want to include in the referenced assembly.
2. Create an Assembly Reference in the folder (menu: **Assets** > **Create** > **Scripting** > **Assembly Definition Reference**).
3. Assign a name to the asset.

   Unity recompiles the scripts in the project to create the new assembly. When it’s finished, you can edit the [properties](class-AssemblyDefinitionReferenceImporter.md) of the new Assembly Definition Reference.
4. Select the new Assembly Definition Reference asset to view its properties in the **Inspector**.
5. Set the Assembly Definition property to reference the target Assembly Definition asset.
6. Click **Apply**.

The referenced assembly now includes all scripts in the same folder as the Assembly Definition Reference, plus those in any subfolders that don’t have their own Assembly Definition or Reference file. Unity removes scripts from their previous assembly where applicable.

## Create a platform-specific assembly

To create an assembly for a specific platform:

1. [Create an Assembly Definition](#create-asmdef).
2. Select the new Assembly Definition to view its properties in the **Inspector**.
3. In the [**Platforms**](class-AssemblyDefinitionImporter.md#platforms) section, check the **Any Platform** option and choose specific platforms to exclude or uncheck **Any Platform** and choose specific platforms to include.
4. Click **Apply**.

The assembly is included or excluded according to this configuration when you build your project for a target platform.

## Create an assembly for Editor code

Editor assemblies allow you to put your **Editor scripts** anywhere in the project, not just in top-level folders named `Editor`.

To create an assembly that contains the Editor code in your project:

1. [Create a platform-specific assembly](#create-platform-specific) in a folder containing your Editor scripts.
2. Include only the Editor platform.
3. If you have additional folders containing Editor scripts, [create Assembly Definition Reference assets](#create-asmref) in those folders and set them to reference this Assembly Definition.

## Create a test assembly

Test assemblies are assemblies that Unity expects to contain tests. Putting your tests in test assemblies has the following benefits:

* They keep your test code separate from the application code you’ll ship to users, so the test code can be compiled only when needed.
* Any tests in a test assembly are automatically visible to the [Test Framework package](test-framework/test-framework-introduction.md), which makes them available to run from the ****Test Runner**** window.

Unity automatically identifies any assembly as a test assembly if it has an [assembly reference](class-AssemblyDefinitionImporter.md#assembly-references) to **nunit.framework.dll** and [assembly definition references](class-AssemblyDefinitionImporter.md#asmdef-references) to **UnityEngine.TestRunner** and **UnityEditor.TestRunner**.

For instructions on installing the Test Framework package and creating test assemblies, refer to the Unity [Test Framework documentation](test-framework/test-framework-introduction.md). You can use the Editor UI to create an assembly definition file with the relevant references predefined, or you can configure the references manually [through the Inspector window](class-AssemblyDefinitionImporter.md) or by [editing the JSON file directly](assembly-definition-file-format.md).

**Note**: Test assemblies are not compiled as part of the regular build pipeline, so any code placed in a test assembly is excluded from a standard [project build](BuildSettings.md). Your test assembly code is only included in a Player build when you [run Play Mode tests in a Player](test-framework/workflow-run-playmode-test-standalone.md) through the **Test Runner** window. If you have production code that unexpectedly doesn’t compile into your project build, double-check to make sure it’s not in a test assembly.

# Additional resources

* [Predefined assemblies reference](script-compile-order-folders.md)