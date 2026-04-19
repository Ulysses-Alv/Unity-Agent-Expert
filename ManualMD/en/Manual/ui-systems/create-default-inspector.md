# Create a default Inspector

Default **Inspector** is a standard way to display serialized object properties without any custom modifications. When you use [FillDefaultInspector](../../ScriptReference/UIElements.InspectorElement.FillDefaultInspector.md) method of the [InspectorElement](../../ScriptReference/UIElements.InspectorElement.md) class, it automatically creates the default hierarchy with the default property fields.

## Example overview

This example extend the `CreateInspectorGUI()` method in the [Create a Custom Inspector](../UIE-HowTo-CreateCustomInspector.md) example to create a default Inspector for the `Car` component. The example creates a `Foldout` control in the `Car_Inspector_UXML.uxml` file and attaches the default Inspector UI to it.

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Foldout](../UIE-uxml-element-Foldout.md)
* [InspectorElement](../../ScriptReference/UIElements.InspectorElement.md)

## Create the default Inspector container

Create a Foldout control to display the default Inspector UI.

1. Double-click the `Car_Inspector_UXML.uxml` file to open it in UI Builder.
2. Add a Foldout control to your UI, name it `Default_Inspector`, and set a label text:

   ![Foldout for the default Inspector](../../uploads/Main/uie-howto-customInspector-uibuilder-Foldout.png)

   Foldout for the default Inspector

## Attach the default Inspector UI

To attach the default Inspector UI to the Foldout, you must obtain a reference to it. You can retrieve the **visual element** of the Foldout from the **visual tree** of your Inspector using [UQuery](../UIE-UQuery.md), and use the [FillDefaultInspector](../../ScriptReference/UIElements.InspectorElement.FillDefaultInspector.md) method of the [InspectorElement](../../ScriptReference/UIElements.InspectorElement.md) class to attach the default Inspector UI to the Foldout control.

1. In the `Car_Inspector.cs` file, update the `CreateInspectorGUI()` method to get a reference to the `Default_Inspector` Foldout and attach the default Inspector UI to it:

   ```
   public override VisualElement CreateInspectorGUI()
   {
      ...
          
      // Get a reference to the default Inspector Foldout control.
       VisualElement InspectorFoldout = myInspector.Q("Default_Inspector");
               
       // Attach a default Inspector to the Foldout.
       InspectorElement.FillDefaultInspector(InspectorFoldout, serializedObject, this);

       // Return the finished Inspector UI.
       return myInspector;
   }
   ```
2. Select the **GameObject** that has the `Car` component to it. The Inspector for the `car` component now displays the `Default Inspector` Foldout with the default Inspector UI inside.

   ![Inspector with a default Inspector](../../uploads/Main/uie-howto-custominspector-foldoutinspector.png)

   Inspector with a default Inspector

## Additional resources

* [Create a custom Inspector](../UIE-HowTo-CreateCustomInspector.md)