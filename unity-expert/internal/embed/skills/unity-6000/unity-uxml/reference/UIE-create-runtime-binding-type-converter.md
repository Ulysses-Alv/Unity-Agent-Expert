# Create a runtime binding with a type converter

**Version**: 6000.0+

This example demonstrates how to create a [type converter](UIE-runtime-binding-data-type-conversion.md) to convert data types between a data source and the UI.

## Example overview

This example creates a Label control in UI Builder that binds to a data source asset. The data source asset contains a float property named `dangerLevel`. The label displays the `dangerLevel` as a string and changes its background color based on the value of the `dangerLevel`.

![Runtime binding type converter example](../uploads/Main/uitk/runtime-binding-converter.gif)

Runtime binding type converter example

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023/runtime-data-binding-converter).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [UXML](UIE-UXML.md)

## Create a data source asset

Create a data source asset that contains the properties you want to bind to. In this example, you create a `ScriptableObject` asset named `ExampleConverterObject` that contains a local converter that converts a `float` to a `Color` and a `string`.

1. Create a project in Unity with any template.
2. In the `Assets` folder of your project, create a C# script named `ExampleConverterObject.cs` with the following content:

```
using Unity.Properties;
using UnityEngine;
using UnityEngine.UIElements;

#if UNITY_EDITOR
using UnityEditor;
#endif

[CreateAssetMenu]
public class ExampleConverterObject : ScriptableObject
{
#if UNITY_EDITOR
    [InitializeOnLoadMethod]
#else
    [RuntimeInitializeOnLoadMethod(RuntimeInitializeLoadType.SubsystemRegistration)]
#endif
    public static void RegisterConverters()
    {
    
        // Create local Converters.
        var group = new ConverterGroup("Value To Progress");
    
        // Converter groups can have multiple converters. This example converts a float to both a color and a string.
        group.AddConverter((ref float v) => new StyleColor(Color.Lerp(Color.red, Color.green, v)));
        group.AddConverter((ref float value) =>
        {
            return value switch
            {
                >= 0 and < 1.0f/3.0f => "Danger",
                >= 1.0f/3.0f and < 2.0f/3.0f => "Neutral",
                _ => "Good"
            };
        });
    
        // Register the converter group in InitializeOnLoadMethod to make it accessible from the UI Builder.
        ConverterGroups.RegisterConverterGroup(group);
    }
    
    [Header("Binding using a converter group")]
    [Range(0, 1)] public float dangerLevel;
}
```

## Create the example object asset

1. Select **Assets** > **Create** > **Example Object**. This creates a data source asset in the `Assets` folder of your project.
2. Rename the asset as `ExampleConverterObject.asset`.

## Create the UI

Create a UI that binds to the data source asset you created in the previous step. In this example, you create a UXML file that contains a Label control.

1. In the `Assets` folder of your project, create a UXML file named `ExampleConverterObject.uxml`.
2. Double-click the `ExampleConverterObject.uxml` file to open it in UI Builder.
3. In the **Hierarchy** panel, add a **Label**.

## Bind the Label to the data source

Bind the properties of the Label to the properties of the data source asset and apply the type converter for the Label.

1. In the **Hierarchy** panel, select the **Label**.
2. In the **Inspector** panel, from the **Bindings** > **Data Source** > **Object** list, select **ExampleConverterObject**.
3. From the **Data Source Path** list, select **dangerLevel**.
4. In the ****Inspector**** panel, right-click the **Text** field and select **Add binding**.
5. In the **Add Binding** window, do the following:

   * From the **Binding Mode** list, select **To Target**. This updates the UI when the data source changes.
   * From the **Advanced Settings** > **Local converters** > **To target property (UI)** list, select **Value To Progress**. This converts a float to a string.
     ![Set the binding for the label text](../uploads/Main/uitk/runtime-binding-value-to-process.png)
6. Select **Add binding** to apply your changes.
7. In the **Inspector** panel, expand the **Background** field.
8. Right-click the **Color** property and select **Add binding**.
9. In the **Add Binding** window, do the following:

   * From the **Binding Mode** list, select **To Target**. This updates the UI when the data source changes.
   * From the **Advanced Settings** > **Local Converters** > **To target property (UI)** list, select **Value To Progress**. This converts a float to a color.
10. Select **Add binding** to apply your changes.
11. Save and close UI Builder. Your `ExampleConverterObject.uxml` file looks like the following:

    ```
    <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" 
    engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
        <ui:Label text="Label" data-source="ExampleConverterObject.asset" data-source-path="dangerLevel">
            <Bindings>
                <ui:DataBinding property="style.backgroundColor" binding-mode="ToTarget" source-to-ui-converters="Value To Progress" />
                <ui:DataBinding property="text" binding-mode="ToTarget" source-to-ui-converters="Value To Progress" />
            </Bindings>
        </ui:Label>
    </ui:UXML>
    ```

## Test the binding

In the Inspector window of the `ExampleObject.asset`, move the **Danger Level** slider. The text and the background color of the **Label** in the UI Builder **Viewport** change according to the value of the **Danger Level**.

## Additional resources

* [Runtime data binding](UIE-runtime-binding.md)
* [Create runtime data binding](UIE-runtime-binding-types.md)
* [Create custom binding types](UIE-runtime-binding-custom-types.md)
* [Define a data source](UIE-runtime-binding-define-data-source.md)
* [Define logging levels](UIE-runtime-binding-logging-levels.md)