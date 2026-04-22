# Create a custom swirl filter

You can create custom filters to apply any effect you want to a **visual element**. To create a custom filter, define the filter function in a custom filter asset. The custom filter asset defines the filter parameters, passes, and parameter bindings.

This example demonstrates a custom filter, created in the custom filter asset’s ****Inspector**** window. You can also [create custom filters in C# code](uss-filter#custom-filters#create-a-custom-filter-in-c).

## Example overview

This example creates a custom filter with a swirl effect, which distorts the visual element by rotating it around a center point.

![Example swirl filter effect](../../uploads/Main/uss/swirl.png)

Example swirl filter effect

You can provide custom filter functions with parameters, set in style or directly in the UI Builder. For example, the built-in blur filter function has only one parameter: the blur radius. This custom swirl filter function has two parameters, the angle and the radius of the swirl:

* **Angle**: Sets the angle of the swirl effect.
* **Radius**: Sets the radius of the swirl effect.

This example uses a sample [swirl shader](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-custom-swirl-filter/SwirlFilter/Swirl.shader) and [material](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-custom-swirl-filter/SwirlFilter/Swirl.mat) to create the swirl filter. The **shader**, which defines the effect, includes two properties—`_Angle`and `_Radius`—to control the swirl’s angle and radius. The material file, `Swirl.mat`, is configured to use the swirl shader. To apply the effect, bind the material properties to the filter parameters.

You can find the completed files for the example in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-custom-swirl-filter).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UXML](../UIE-UXML.md)
* [USS](../UIE-about-uss.md)
* [USS filter](uss-filter.md)

## Create a custom filter asset

Create the filter asset and define the filter parameters, passes, and parameter bindings in the asset’s **Inspector** window.

1. Create a project in Unity with any template.
2. In the `Assets` folder, create a folder named `SwirlFilter`.
3. Download and import the following files into the `SwirlFilter` folder:
   * [Swirl.shader](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-custom-swirl-filter/SwirlFilter/Swirl.shader)
   * [Swirl.mat](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-custom-swirl-filter/SwirlFilter/Swirl.mat)
4. To create a [`FilterFunctionDefinition`](../../ScriptReference/UIElements.FilterFunctionDefinition.md) asset, right-click in the `SwirlFilter` folder and select **Create** > **UI Toolkit** > **Filter Function Definition**.
5. Rename the asset to `SwirlFilter`.
6. In the **Inspector** window for the `SwirlFilter` asset, set **Filter Name** to `swirl`.
7. In the **Parameters** section, add the two parameters and set the **Type** of both parameters to `float`.
8. Set the name of the first parameter to `Angle`. This controls the angle of the swirl.
9. Set the name of the second parameter to `Radius`. This controls the radius of the swirl.
10. In the **Passes** section, add a new element and set the **Material** to `Swirl.shader`.
11. In the **Parameter Bindings** section, add two bindings.
12. For the first binding, do the following:

    * Set the **Index** to `0`.
    * Set the **Property** to `_Angle`.
13. For the second binding, do the following:

    * Set the **Index** to `1`.
    * Set the **Property** to `_Radius`.
    ![Define swirl filter in the Inspector window](../../uploads/Main/uss/swirl-filter-function.png)

    Define swirl filter in the Inspector window

## Create the UI Document and the style sheet

To get the swirl effect, create a UXML file with two VisualElements with different background colors. The first VisualElement is the parent element, and the second VisualElement is the child element. Apply the custom filter to the parent element so that the child element is affected by the filter.

1. Create a new UI Toolkit USS file named `SwirlFilterExample.uss` with the following content:

   ```
   .outside {
       flex-grow: 1;
       position: absolute;
       height: 207px;
       width: 234px;
       top: 46px;
       left: 27px;
       background-color: rgb(255, 0, 0);
   }
       
   .inside {
       flex-grow: 1;
       position: absolute;
       height: 75px;
       width: 100px;
       top: 46px;
       left: 27px;
       background-color: rgb(0, 255, 247);
   }
   ```
2. Create a new UI Toolkit UXML file named `SwirlFilterExample.uxml` with the following content:

   ```
   <engine:UXML xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:engine="UnityEngine.UIElements" xmlns:editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
       <Style src="SwirlFilterExample.uss" />
       <engine:VisualElement class="outside">
           <engine:VisualElement class="inside" />
       </engine:VisualElement>
   </engine:UXML>
   ```

## Apply the custom filter in the UI Builder

To apply the custom filter in the UI Builder, add the custom filter function to the UI Builder and set the parameters.

1. Double-click the `SwirlFilterExample.uxml` file to open it in the UI Builder.
2. In the **StyleSheets** panel, select **+** > **Add Existing USS** and select `SwirlFilterExample.uss`.
3. In the **Hierarchy** panel, select the parent VisualElement.
4. In the **Inspector** panel, select **Inline Styles** > **Filter**.
5. In the **Filter** section, click the **Add(+)** button.
6. In the **Function** dropdown, select `Custom`.
7. Set **Definition** to `SwirlFilter`.
8. In the **Angle** field, set the angle value, such as `58.9`.
9. In the **Radius** field, set the radius value, such as `2.3`.

   ![Apply swirl filter in UI Builder](../../uploads/Main/uss/apply-swirl-filter.png)

   Apply swirl filter in UI Builder
10. You can adjust the angle and radius values to check the effect of the swirl filter in the ****ViewPort**** of the UI Builder.

## Save the filter function to USS

You can save the filter function to the USS file instead of applying it as the inline style.

1. In the **Inspector** panel of the parent VisualElement, click in the **Style Class List** field and enter `.filter-effect`.
2. Select **Extract Inlined Style to New Class**. This adds the following USS class to the `SwirlFilterExample.uss` file and applies the class to the parent VisualElement:

   ```
   .filterEffect {
       filter: filter("SwirlFilter/SwirlFilterFunction.asset" 58.9 2.3);
   }
   ```

## Addition resources

* [USS properties](../UIE-uss-properties.md)