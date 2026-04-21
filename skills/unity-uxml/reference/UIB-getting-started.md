# Get started with UI Builder

Want to learn how to create UI with UI Builder? Use this example to get started.

To create UI in the UI Builder:

1. Create a new UI Document (UXML).
2. Add elements to create your UI hierarchy.
3. Set up attributes and style properties in the **Inspector**.
4. When more than one element starts to need the same style properties, [create USS style sheets and selectors](#create-a-uss-class-for-the-border-style).
5. Test your UI and if you are satisfied with the results, [extract inline style properties to USS classes](#extract-inline-styles).
6. Save the UI Document (UXML).

## Prerequisites

Before you start, get familiar with the following:

* [Introduction to visual elements and visual tree](UIE-VisualTree.md)
* [Introduction to UXML](UIE-WritingUXMLTemplate.md)
* [Introduction to USS](UIE-about-uss.md)
* [UI Builder interface overview](UIB-interface-overview.md)

## Example overview

This example creates the main view for the [Create a list view runtime UI](UIE-HowTo-CreateRuntimeUI.md) example. It creates a root element as the background, with two containers. One container holds the character name list and another holds the character details. In the character details container, you add background and foreground frames. Finally, you add two labels for the character name. This example won’t create the character name list entry UI.

![Final main view layout](../uploads/Main/uie-howto-runtimeUI-finalMainView.png)

Final main view layout

## Create the root element

Create a new project and then create a root **visual element** that covers the entire screen. Set your root element to have a background color and center all child elements in the middle of the screen.

1. Create a project in Unity with any template.
2. Select **Window** > **UI Toolkit** > **UI Builder**.
3. In the UI Builder window, at the top left of the **Viewport** window, select **File** > **New** to create a new UXML document.

   ![UI Builder file menu](../uploads/Main/UIBuilder/uie-howto-runtimeUI-fileMenu.png)

   UI Builder file menu
4. Name it as `MainView.uxml` and save.
5. Drag **VisualElement** from **Library** into the **Hierarchy** panel.

   ![Create new elements by dragging from the Library](../uploads/Main/UIBuilder/uie-howto-runtimeUI-newVisualElement.png)

   Create new elements by dragging from the Library

   **Tip**: You can also double-click a control to add it to the **Hierarchy** panel.
6. Select the element from the **Hierarchy** panel.
7. In the **Inspector** panel, make sure **Flex** > **Grow** is `1`. This sets the `flex-grow` property to `1`, making it cover the entire screen.

   ![Set the Flex property](../uploads/Main/UIBuilder/uie-howto-runtimeUI-flexGrow.png)

   Set the Flex property
8. To center all child elements, set both **Align Items** and **Justify Content** to `Center`.

   ![Center children](../uploads/Main/UIBuilder/uie-howto-runtimeUI-alignCenter.png)

   Center children
9. Set **Background** > **Color** to `#732526`.

   **Note:** By default, the Alpha value for colors is `0`, making them fully transparent. To make color opaque, set the Alpha value to `255`.

   ![Root element background color](../uploads/Main/UIBuilder/uie-howto-runtimeUI-backgroundColor.png)

   Root element background color

## Create the parent container

Create a new VisualElement underneath the root element. This element becomes the parent container for the left and right sections of the UI.

1. Drag **VisualElement** from **Library** to the root VisualElement in the **Hierarchy** panel.

   ![Add a child VisualElement](../uploads/Main/UIBuilder/uie-howto-runtimeUI-secondaryContainer.png)

   Add a child VisualElement
2. Select the element from the **Hierarchy** panel.
3. In the **Inspector** panel, set **Flex** > **Direction** to `row`.
4. Set **Flex** > **Grow** to `0`.
5. Set **Size** > **Height** to `350` pixels.

   ![Center container properties](../uploads/Main/UIBuilder/uie-howto-runtimeUI-centerContainer.png)

   Center container properties

## Create the character names list container

Add a ListView as the child element of the container to hold the character names.

1. Drag a **ListView** from the **Library** to the container VisualElement in the **Hierarchy** panel.
2. Select the element from the **Hierarchy** panel.
3. In the **Inspector** panel, set **Name** to `CharacterList`.

   ![Background container with the empty element inside](../uploads/Main/UIBuilder/uie-howto-runtimeUI-setElementName.png)

   Background container with the empty element inside
4. Set **Size** > **Width** to `230` pixels.
5. Set **Spacing** > **Margin** > **Right** to `6` pixels.

   ![Size and Margin foldouts for the character list](../uploads/Main/UIBuilder/uie-howto-runtimeUI-sizeAndMargin.png)

   Size and Margin foldouts for the character list
6. Set **Background** > **Color** to `#6E3925`.
7. Set **Border** > **Color** to `#311A11`.
8. Set **Border** > **Width** to `4` pixels.
9. Set **Border** > **Radius** to `15` pixels.

   ![Styled character list](../uploads/Main/UIBuilder/uie-howto-runtimeUI-characterListStyling.png)

   Styled character list

## Create the character details container

Add a new VisualElement under the same parent as the `#CharacterList` to hold the character details container. The purpose is that when the user selects a character from the list on the left, it displays the character’s portrait, name, and class.

1. Drag a **VisualElement** from the **Library** to the container element in the **Hierarchy** panel. This is the container to hold all the elements on the right.
2. Select the element from the **Hierarchy** panel.
3. In the **Inspector** panel, set **Align** > **Align Items** to `flex-end`.
4. Set **Align** > **Justify Content** to `space-between`.

   ![Justify content property](../uploads/Main/UIBuilder/uie-howto-runtimeUI-justifySpaceBetween.png)

   Justify content property
5. Add another **VisualElement** as the child of the right container.
6. Select the element from the **Hierarchy** panel.
7. Set **Flex** > **Grow** to `0`.
8. Set **Size** > **Width** to `276` pixels.
9. In the **Align** section, set both **Align Items** and **Justify Content** to `center`.
10. Set **Spacing** > **Padding** to `8` pixels.

    ![Properties of the character details container](../uploads/Main/UIBuilder/uie-howto-runtimeUI-characterDetailsProperties.png)

    Properties of the character details container
11. Set **Background** > **Color** to `#AA5939`.

## Create a USS class for the border style

The character details container will use the same border styles as the character names list container. Create a USS class to apply to both containers.

1. In the **StyleSheet** panel, select **+** > **Create New USS**.
2. Name it `MainView.uss` and select **Save**.
3. Click in the **Add new selector…** field and enter `.border`. A **`.border`** selector appears in the **StyleSheet** panel.
4. In the **StyleSheet** panel, select **`.border`**.
5. In the **Inspector** panel, set **Border** > **Color** to `#311A11`.
6. Set **Border** > **Width** to `4` pixels.
7. Set **Border** > **Radius** to `15` pixels.
8. Drag **`.border`** from the **StyleSheet** panel to the character details container **VisualElement**.
9. Drag **`.border`** from the **StyleSheet** panel to **#CharacterList**.
10. Select **#CharacterList**.
11. In the **Inspector** panel, right-click **Border** and select **Unset** to remove the inline style you set earlier.

Your UI layout now looks like the following:

![Empty character details panel](../uploads/Main/UIBuilder/uie-howto-runtimeUI-intermediate02.png)

Empty character details panel

## Create the character portrait background

Add the individual UI controls to the character details container. The first step is to add the character portrait background.

1. Drag **VisualElement** from **Library** to the character details container.
2. Select the element from the **Hierarchy** panel.
3. In the **Size** section, set both **Width** and **Height** to `120` **pixels**.
4. Set **Spacing** > **Padding** to `4` pixels.
5. Set **Background** > **Color** to `#FF8554`.
6. Set **Border** > **Color** to `#311A11`.
7. Set **Border** > **Width** to `2` pixels.
8. Set **Border** > **Radius** to `13` pixels.

   ![Background frame for the character portrait](../uploads/Main/UIBuilder/uie-howto-runtimeUI-pictureFrame.png)

   Background frame for the character portrait

## Create the character portrait foreground

Next in the character details container, add the foreground for an actual image.

1. Drag **VisualElement** from **Library** to the character details container.
2. Select the element from the **Hierarchy** panel.
3. Set **Name** to `CharacterPortrait`.
4. Set **Flex** > **Grow** to `1`, so that the image can use all the available space.
5. Set **Background > Scale Mode** to `scale-to-fit`, so that you can scale the image to match the element size while keeping the correct aspect ratio.

   ![VisualElement for the portrait image](../uploads/Main/UIBuilder/uie-howto-runtimeUI-portraitElement.png)

   VisualElement for the portrait image

## Create labels

Add two label controls to the character details container to display the selected character’s name and class.

1. Drag **Label** from **Library** to the character details container in the **Hierarchy** panel.
2. Set `Name` to `CharacterName`.
3. Drag **Label** from **Library** to the character details container in the **Hierarchy** panel.
4. Set `Name` to `CharacterClass`.

   ![Add labels for name and class](../uploads/Main/UIBuilder/uie-howto-runtimeUI-characterLabels.png)

   Add labels for name and class
5. Select the `#CharacterName` element.
6. Set **Text** > **Font Style** to `B`.
7. Set **Text** > **Size** to `18` pixels.

   ![Change font settings](../uploads/Main/UIBuilder/uie-howto-runtimeUI-labelSettings.png)

   Change font settings
8. In the **Viewport** window, select **File** > **Save** to save the changes to `MainView.uxml`.

## Extract inline styles

In UI Builder, you can create elements and use inline styles only to experiment while the number of elements is still small. As you build a more complex UI, it’s easier to manage styles using style sheets. You can extract inline styles to a style sheet in UI Builder.

1. Select the root visual element.
2. In the **Inspector** panel, in the **Style Class List** field, enter `.background` as the class name.
3. Select **Extract Inlined Styles to New Class**. This creates a `.background` class selector with the inline styles you set for the root element and updates the UI Document (UXML) for the root visual element to use the class selector instead of the inline styles.

   ![Extract inline styles to a new class](../uploads/Main/UIBuilder/extract-inline-styles.png)

   Extract inline styles to a new class
4. In the **Viewport** window, select **File** > **Save**.

If you want to continue to work on the [Create a list view runtime UI](UIE-HowTo-CreateRuntimeUI.md) example, you can repeat the steps to extract styles for all the other elements, and follow the instructions to create the example.

## Additional resources

* [Style elements with UI Builder](UIB-styling-ui-using-uss-selectors.md)