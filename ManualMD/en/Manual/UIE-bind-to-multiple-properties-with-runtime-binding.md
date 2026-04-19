# Bind to multiple properties with runtime binding

**Version**: 6000.0+

This example demonstrates how to bind multiple properties of a data source asset to a UI control in UI Builder.

## Example overview

This example creates a data source asset that contains a `Vector3` property and a `float` property. The `float` property is a read-only property that returns the sum of the `x`, `y`, and `z` values of the `Vector3` property. The example binds the `Vector3` property to a `Vector3Field` and the `float` property to a `FloatField`. When you change the `Vector3` property in the UI, `FloatField` displays the sum of the `x`, `y`, and `z` values of the `Vector3` property.

This example also demonstrates how to use the To Source [binding mode](UIE-runtime-binding-mode-update.md) to update the data source when the UI changes.

![UI Builder preview mode](../uploads/Main/uitk/runtime-binding-to-source-result.png)

UI Builder preview mode

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/runtime-data-binding-multiple-properties).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [UXML](UIE-UXML.md)

## Create a data source asset

Create a data source asset that contains the properties you want to bind to.

1. Create a project in Unity with any template.
2. In the `Assets` folder of your project, create a C# script named `ExampleMultiPropertiesObject.cs` with the following content:

```
using Unity.Properties;
using UnityEngine;
using UnityEngine.UIElements;

#if UNITY_EDITOR
using UnityEditor;
#endif

[CreateAssetMenu]
public class ExampleMultiPropertiesObject : ScriptableObject
{
    [Header("Bind to multiple properties")]

    [CreateProperty]
    public Vector3 vector3Value;
    
    [CreateProperty]
    public float sumOfVector3Properties => vector3Value.x + vector3Value.y + vector3Value.z;
}
```

## Create the example object asset

1. Select **Assets** > **Create** > **Example MultiProperties Object**. This creates a data source asset in the `Assets` folder of your project.
2. Rename the asset as `ExampleMultiPropertiesObject.asset`.

## Create the UI

Create a UI that binds to the data source asset you created in the previous step. In this example, you create a UXML file that contains a VisualElement with a Vector3Field and a FloatField as its child elements.

1. In the `Assets` folder of your project, create a UXML file named `ExampleMultiPropertiesObject.uxml`.
2. Double-click the `ExampleMultiPropertiesObject.uxml` file to open it in UI Builder.
3. In the **Hierarchy** panel, add a **VisualElement**.
4. Add a **Vector3Field** and a **FloatField** as the child elements of the **VisualElement**.

## Bind the Vector3Field to the data source

Bind the UI to the data source asset you created in the previous step.

1. In the **Inspector** panel of the **VisualElement**, from the **Bindings** > **Data Source** > **Object** list, select **ExampleMultiPropertiesObject**.
2. In the ****Inspector**** panel of the **Vector3Field**, right-click the **Value** property and select **Add binding**.
3. In the **Add Binding** window, select **vector3Value** from the **Data Source Path** list.
4. From the **Binding Mode** list, select **To Source**. This updates the data source when UI changes.

   ![Set the VisualElement binding data source](../uploads/Main/uitk/runtime-binding-to-source.png)
5. Select **Add binding** to apply your changes.

## Bind the FloatField to the data source

Bind the `value` property of the FloatField to the `sumOfVector3Properties` property of the data source asset.

1. In the **Inspector** panel of the **Float**, right-click the **Value** property and select **Add binding**.
2. In the **Add Binding** window, from the **Data Source Path** list, select **sumOfVector3Properties**.
3. From the **Binding Mode** list, select **To Target**. This updates the UI when the data source changes.
4. Select **Add binding** to apply your changes.
5. Save and close UI Builder. Your `ExampleMultiPropertiesObject.uxml` file looks like the following:

```
<engine:UXML xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:engine="UnityEngine.UIElements" xmlns:editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
    <engine:VisualElement data-source="MulPropertyObject.asset" name="VisualElement" style="flex-grow: 1;">
        <engine:Vector3Field label="Vec3 Field">
            <Bindings>
                <engine:DataBinding property="value" data-source-path="vector3Value" binding-mode="ToSource"/>
            </Bindings>
        </engine:Vector3Field>
        <engine:FloatField label="Float Field" name="FloatField">
            <Bindings>
                <engine:DataBinding property="value" data-source-path="sumOfVector3Properties" binding-mode="ToTarget"/>
            </Bindings>
        </engine:FloatField>
    </engine:VisualElement>
</engine:UXML>
```

## Test the binding in UI Builder preview mode

To test the binding in the UI Builder, update the value of the Vector3 property. If the binding is set up correctly, the values replicate in the data source asset.

1. In the UI Builder **Viewport**, select **Preview**.
2. Enter random numbers in the **X**, **Y**, and **Z** fields of the **Vector3Field**. The **FloatField** displays the sum of the numbers you entered.
3. In the `Assets` folder of your project, select the `ExampleMultiPropertiesObject.asset` file. The **Vector3 Value** in the Inspector window of the `ExampleObject.asset` changes to the numbers you entered.

**Note**:

* If you can’t enter numbers in the **Vector3Field**, make the field wider by dragging the right edge of the canvas.
* To troubleshoot binding issues, enable [runtime data binding logging](UIE-runtime-binding-logging-levels.md).

## Additional resources

* [Get started with runtime data binding](UIE-get-started-runtime-binding.md)
* [Create a runtime binding type converter](UIE-create-runtime-binding-type-converter.md)
* [Create runtime data binding](UIE-runtime-binding-types.md)
* [Define a data source](UIE-runtime-binding-define-data-source.md)