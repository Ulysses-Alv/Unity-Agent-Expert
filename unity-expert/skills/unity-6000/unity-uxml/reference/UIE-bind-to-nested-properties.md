# Bind to nested properties

**Version**: 2021.3+

This example demonstrates how to use the `binding-path` attribute of a BindableElement in UXML to bind fields to nested properties of a SerializedObject.

## Example overview

This example creates a custom **Inspector** UI with the following:

* Two fields that bind to a **GameObject**’s name and the scale of a USS transform
* Two fields that bind to nested properties of a SerializedObject

![A custom Inspector for a script, which has fields that map to GameObjects and SerializedObjects.](../uploads/Main/uie_bind_nested_properties.png)

A custom Inspector for a script, which has fields that map to GameObjects and SerializedObjects.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-nested-properties).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [`BindableElement`](../ScriptReference/UIElements.BindableElement.md)
* [`PropertyField`](../ScriptReference/UIElements.PropertyField.md)

## Create the destructible tank object

Create a C# script to define a class for a tank that has health values to make it destructible.

1. Create a project in Unity with any template.
2. In your **Project window**, create a folder named `bind-nested-properties` to store all the files.
3. Create a C# script named `DestructibleTankScript.cs` and replace its content with the following:

   ```
   using System;
   using UnityEngine;
   using UnityEngine.Serialization;

   [Serializable]
   public struct Health
   {
       public int armor;
       public int life;
   }

   [ExecuteInEditMode]
   public class DestructibleTankScript : MonoBehaviour
   {
       public string tankName = "Tank";
       public float tankSize = 1;

       public Health health;

       private void Update()
       {
           gameObject.name = tankName;
           gameObject.transform.localScale = new Vector3(tankSize, tankSize, tankSize);
       }

       public void Reset()
       {
           health.armor = 100;
           health.life = 10;
       }
   }
   ```

## Create the UXML and the Inspector UI

Create a UXML file with a BindableElement. Set the BindableElement’s `binding-path` to the `health` property and set each child element’s `binding-path` of the BindableElement to the `armor` and `life` properties of `health`.

1. In the **bind-nested-properties** folder, create a folder named `Editor`.
2. In the **Editor** folder, create a USS file named `tank_inspector_styles.uss` and replace its contents with the following:

   ```
   .container {
       background-color: rgb(80, 80, 80);
       flex-direction: column;
   }
   Label {
       background-color: rgb(80, 80, 80);
   }
   TextField:hover {
       background-color: rgb(255, 255, 0);
   }
   FloatField {
       background-color: rgb(0, 0, 255);
   }
   ```
3. Create a UI Document named `destructible_tank_editor.uxml` and replace its contents with the following:

   ```
   <UXML xmlns="UnityEngine.UIElements" xmlns:ue="UnityEditor.UIElements">
       <Style src="tank_inspector_styles.uss"/>
       <VisualElement name="row" class="container">
           <Label text="Tank Script - Custom Inspector" />
           <ue:PropertyField binding-path="tankName" name="tank-name-field" />
           <ue:PropertyField binding-path="tankSize" name="tank-size-field" />
           <BindableElement binding-path="health">
               <ue:PropertyField binding-path="armor"/>
               <ue:PropertyField binding-path="life"/>
           </BindableElement>
       </VisualElement>
   </UXML>
   ```

## Create the custom Editor

Create a C# script that registers a custom Editor for the `DestructibleTankScript`.

1. Create a C# script named `DestructibleTankEditor.cs` and replace its content with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;

   [CustomEditor(typeof(DestructibleTankScript))]
   public class DestructibleTankEditor : Editor
   {
       [SerializeField]
       VisualTreeAsset visualTreeAsset;

       public override VisualElement CreateInspectorGUI()
       {
           return visualTreeAsset.CloneTree();
       }
   }
   ```
2. Select `DestructibleTankEditor.cs` in the Project window.
3. Drag `destructible_tank_editor.uxml` to **Visual Tree Asset** in the Inspector.

## Test the binding

1. In Unity, add an empty GameObject to a **scene**.
2. Select the GameObject.
3. Add the **Destructible Tank Script** component in the Inspector. The **Armor** and **Life** fields are bound to the `health.armor` and `health.life` properties. If you change the values in the Inspector, the values of the underlying properties change.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)