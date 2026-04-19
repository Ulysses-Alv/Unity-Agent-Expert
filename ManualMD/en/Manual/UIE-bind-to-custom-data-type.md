# Bind a custom control to custom data type

**Version**: 2021.3+

This example demonstrates how to bind custom controls to custom data types.

## Example overview

This example creates a custom data type and a custom control based on three built-in controls. It binds the custom control to the custom data type. The drawer converts between Celsius and Fahrenheit.

![Script with a custom control that converts temperature, displayed in the Unity Editor.](../uploads/Main/uie_bind_custom_data_type.png)

Script with a custom control that converts temperature, displayed in the Unity Editor.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-custom-data-type).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UXML](UIE-UXML.md)
* [`PropertyField`](../ScriptReference/UIElements.PropertyField.md)

## Create a custom data type

Create a custom data type `Temperature`, and use it as a serialized property.

1. Create a Unity project with any template.
2. Create a folder named `bind-custom-data-type` to store all the files.
3. Create a C# script named `Temperature.cs` and replace its contents with the following:

   ```
   using System;

   namespace UIToolkitExamples
   {
       public enum TemperatureUnit
       {
           Celsius,
           Farenheit
       }

       [Serializable]
       public struct Temperature
       {
           public double value;
           public TemperatureUnit unit;
       }
   }
   ```
4. Create a C# script named `PlanetScript.cs` and replace its contents with the following:

   ```
   using UnityEngine;

   namespace UIToolkitExamples
   {
       public class PlanetScript : MonoBehaviour
       {
           public Temperature coreTemperature;
       }
   }
   ```

## Create a custom control

Create a custom Editor for `Planet` and a custom **Property Drawer** for `Temperature`.

In the custom Property Drawer, implement a button that converts temperature between Fahrenheit and Celsius by writing to properties of [`SerializedProperty`](../ScriptReference/SerializedProperty.md) (use [`doubleValue`](../ScriptReference/SerializedProperty-doubleValue.md) and [`enumValueIndex`](../ScriptReference/SerializedProperty-enumValueIndex.md)) and then calling [`SerializedObject.ApplyModifiedProperties()`](../ScriptReference/SerializedObject.ApplyModifiedProperties.md).

The custom Property Draweris is considered as a custom control. It’s a built-in control that behaves in a custom way.

1. Create a folder named `Editor`.
2. In the **Editor** folder, create a C# script named `PlanetEditor.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine.UIElements;
   using UnityEditor.UIElements;

   namespace UIToolkitExamples
   {
       [CustomEditor(typeof(PlanetScript))]
       public class PlanetEditor : Editor
       {
           public override VisualElement CreateInspectorGUI()
           {
               return new PropertyField(serializedObject.FindProperty("coreTemperature"));
           }
       }
   }
   ```
3. In the **Editor** folder, create a C# script named `TemperatureDrawer.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;

   namespace UIToolkitExamples
   {
       [CustomPropertyDrawer(typeof(Temperature))]
       public class TemperatureDrawer : PropertyDrawer
       {
           public override VisualElement CreatePropertyGUI(SerializedProperty property)
           {
               var asset = Resources.Load<VisualTreeAsset>("temperature_drawer");
               var drawer = asset.Instantiate(property.propertyPath);

               drawer.Q<Label>().text = property.displayName;

               // Don't allow conversion when you've selected multiple objects in the Inspector
               if (!property.serializedObject.isEditingMultipleObjects)
               {
                   drawer.Q<Button>().RegisterCallback<ClickEvent, SerializedProperty>(Convert, property);
               }

               return drawer;
           }

           static void Convert(ClickEvent evt, SerializedProperty property)
           {
               var valueProperty = property.FindPropertyRelative("value");
               var unitProperty = property.FindPropertyRelative("unit");

               // F -> C
               if (unitProperty.enumValueIndex == (int)TemperatureUnit.Farenheit)
               {
                   valueProperty.doubleValue -= 32;
                   valueProperty.doubleValue *= 5.0d / 9.0d;
                   unitProperty.enumValueIndex = (int)TemperatureUnit.Celsius;
               }
               else // C -> F
               {
                   valueProperty.doubleValue *= 9.0d / 5.0d;
                   valueProperty.doubleValue += 32;
                   unitProperty.enumValueIndex = (int)TemperatureUnit.Farenheit;
               }

               // Important: Because you are bypassing the binding system, you must save the modified SerializedObject
               property.serializedObject.ApplyModifiedProperties();
           }
       }
   }
   ```

## Set the binding

Create a UXML file with the following:

* A DoubleField
* An EnumField
* A Button

Set the `binding-path` of the two fields to `value` and `unit` of the `Temperature` property.

1. In the **Editor** folder, create a folder named `Resources`.
2. In the **Resources** folder, create a UI Document named `temperature_drawer.uxml` and replace its contents with the following:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
       <ui:VisualElement class="unity-base-field">
           <ui:Label class="unity-base-field__label" />
           <ui:VisualElement class="unity-base-field__input" style="flex-direction: row;">
               <uie:DoubleField binding-path="value" />
               <uie:EnumField binding-path="unit" />
               <ui:Button text="Convert" />
           </ui:VisualElement>
       </ui:VisualElement>
   </ui:UXML>
   ```

## Test the binding

1. Create an empty **GameObject** in a **scene**.
2. In the Hierarchy, select the GameObject.
3. Drag **PlanetScript.cs** to the **Inspector**. This adds a **Planet Script** component to the GameObject.
4. Enter a number in the **Temperature** field and select a unit from the dropdown.
5. Select the **Convert** button to convert between units. If you make changes in the Inspector UI, the `Temperature` property of the custom contol changes.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)