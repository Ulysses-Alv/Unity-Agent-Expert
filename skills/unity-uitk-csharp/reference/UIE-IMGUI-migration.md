# Migrate from Immediate Mode GUI (IMGUI) to UI Toolkit

This guide is for developers experienced with [Immediate Mode GUI (IMGUI)](GUIScriptingGuide.md) to migrate to UI Toolkit. This guide focuses on the Editor UI, but its information can also apply to the runtime UI as well.

## Key differences

### Code-driven versus UI-driven

IMGUI is code-driven by calls to the `OnGUI` function in a C# script that implements it. UI Toolkit provides more options for Editor UI creation. With UI Toolkit, you define the behaviors in C# **scripts**. However, when defining UI elements and styles, in addition to C#, you can visually define UI controls in [UI Builder](UIBuilder.md) or write in [an XML-like text file (called UXML)](UIE-UXML.md) directly. For more information, refer to [Get started with UI Toolkit](UIE-simple-ui-toolkit-workflow.md).

### Immediate versus retained mode

With IMGUI, you describe the UI tree when the UI is repainted within the [`OnGUI()`](../ScriptReference/EditorWindow.OnGUI.md) function. You must call this function when an event enters the UI or when you repaint the UI. There is no persistent information pertaining to the UI tree between different events. Whereas, you create **visual elements** with UI Toolkit in a tree structure called [Visual Tree](UIE-VisualTree.md). Information in the Visual Trees is retained persistently.

### Constant versus state changes

IMGUI is based on the `OnGUI()` function that runs at least once every frame. You define the look and the behaviors of the UI for every possible frame. The body of `OnGUI()` might contain many conditions and different states.

UI Toolkit operates in an event-driven system. You define the look of the UI in its default state and define the behaviors of the UI in response to events. Any changes you make in UI Toolkit cause persistent changes to the state of your UI.

For example, the declaration of a button in IMGUI looks like the following:

```
if (GUILayout.Button("Click me!"))
{
    //Code runs here in frames where the user clicks the button.

    //Code makes changes to the UI for frames where the user has just clicked the button.
}
else
{
    //Code specifies what happens in all other frames.
}
```

The example above looks like the following in UI Toolkit:

```
UIDocument document = GetComponent<UIDocument>();

//Create button.
Button button = new Button();
button.text = "Click me!";

//Set up event handler.
button.RegisterCallback<ClickEvent>((ClickEvent evt) =>
{
    //Code runs here after button receives ClickEvent.
});

//Add button to UI.
document.rootVisualElement.Add(button);
```

For a completed example of how to create a custom Editor window with UI Toolkit, refer to [Get started with UI Toolkit](UIE-simple-ui-toolkit-workflow.md).

## IMGUI support

Use the `IMGUIContainer` to place IMGUI code inside of a `VisualElement`. Everything you can do inside of `OnGUI()` is supported.

You can arrange multiple `IMGUIContainer`s and lay them out by mixing `GUILayout` and UI Toolkit layouts. Note that it’s not possible to add `VisualElement` instances inside of an `IMGUIContainer`.

![IMGUIContainer example](../uploads/Main/UIElementsIMGUI.png)

`IMGUIContainer` example

## From IMGUI to UI Toolkit conversion

The following table lists the equivalent functions between IMGUI and UI Toolkit:

| **Action** | **IMGUI** | **UI Toolkit** |
| --- | --- | --- |
| Create an [Editor Window](editor-EditorWindows.md) | [`EditorWindow.OnGUI()`](../ScriptReference/EditorWindow.OnGUI.md) | [`EditorWindow.CreateGUI()`](../ScriptReference/EditorWindow.CreateGUI.md) |
| Create a [Property Drawer](editor-PropertyDrawers.md) or a Property Attribute | [`PropertyDrawer.OnGUI()`](../ScriptReference/PropertyDrawer.OnGUI.md) | [`PropertyDrawer.CreatePropertyGUI()`](../ScriptReference/PropertyDrawer.CreatePropertyGUI.md) |
| Create a custom Editor for the ****Inspector**** | [`Editor.OnInspectorGUI()`](../ScriptReference/Editor.OnInspectorGUI.md) | [`Editor.CreateInspectorGUI()`](../ScriptReference/Editor.CreateInspectorGUI.md) |

The following table lists the equivalent methods, classes, and attributes between IMGUI and UI Toolkit:

| **IMGUI** | **IMGUI namespaces** | **UI Toolkit** |
| --- | --- | --- |
| [`AddCursorRect()`](../ScriptReference/EditorGUIUtility.AddCursorRect.md) | EditorGUIUtility | Set [`VisualElement.style.cursor`](../ScriptReference/UIElements.IStyle-cursor.md), or set a visual element’s cursor texture in the UI Builder or USS. For more detailed interactivity, use C# events. |
| [`AreaScope`](../ScriptReference/GUILayout.AreaScope.md) | GUILayout | Scopes are generally not needed in UI Toolkit. See `BeginArea()`. |
| [`BeginArea()`](../ScriptReference/GUILayout.BeginArea.md) | GUILayout | To define the area itself, create a visual element and set [`style.position`](../ScriptReference/UIElements.IStyle-position.md) to [`Position.Absolute`](../ScriptReference/UIElements.Position.Absolute.md). To create children for the area, create child visual elements under it. |
| [`BeginBuildTargetSelectionGrouping()`](../ScriptReference/EditorGUILayout.BeginBuildTargetSelectionGrouping.md) | EditorGUILayout | No equivalent. |
| [`BeginChangeCheck()`](../ScriptReference/EditorGUI.BeginChangeCheck.md) | EditorGUI | Register callbacks on each element in the change check range. If using a [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) as a stand-in for a serialized field in a custom Inspector, use [`PropertyField.RegisterCallback<SerializedPropertyChangeEvent>()`](../ScriptReference/UIElements.CallbackEventHandler.RegisterCallback.md) or [`PropertyField.RegisterValueChangeCallback()`](../ScriptReference/UIElements.PropertyField.RegisterValueChangeCallback.md). In all other cases, use [`VisualElement.RegisterCallback<ChangeEvent<T>>()`](../ScriptReference/UIElements.CallbackEventHandler.RegisterCallback.md) or [`VisualElement.RegisterValueChangedCallback<T>()`](../ScriptReference/UIElements.INotifyValueChangedExtensions.RegisterValueChangedCallback.md). |
| [`BeginDisabledGroup()`](../ScriptReference/EditorGUI.BeginDisabledGroup.md) | EditorGUI | [`VisualElement.SetEnabled(false)`](../ScriptReference/UIElements.VisualElement.SetEnabled.md) |
| [`BeginFoldoutHeaderGroup()`](../ScriptReference/EditorGUI.BeginFoldoutHeaderGroup.md) | EditorGUI, EditorGUILayout | See `Foldout()`. |
| [`BeginGroup()`](../ScriptReference/GUI.BeginGroup.md) | GUI | See `BeginArea()`. |
| [`BeginHorizontal()`](../ScriptReference/EditorGUILayout.BeginHorizontal.md) | EditorGUILayout, GUILayout | See `BeginArea()`. |
| [`BeginProperty()`](../ScriptReference/EditorGUI.BeginProperty.md) | EditorGUI | If you use `BeginProperty()`/`EndProperty()` to bind a simple control to a serialized property, you can do that in UI Toolkit by calling [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md), by setting [`bindingPath`](../ScriptReference/UIElements.BindableElement-bindingPath.md), or by setting the `binding-path` UXML attribute. If you use `BeginProperty()`/`EndProperty()` to make a single property out of complex custom UI, that is not supported well in UI Toolkit. |
| [`BeginScrollView()`](../ScriptReference/GUI.BeginScrollView.md) | EditorGUILayout, GUI, GUILayout | [`UnityEngine.UIElements.ScrollView`](../ScriptReference/UIElements.ScrollView.md) |
| [`BeginToggleGroup()`](../ScriptReference/EditorGUILayout.BeginToggleGroup.md) | EditorGUILayout | No equivalent. |
| [`BeginVertical()`](../ScriptReference/EditorGUILayout.BeginVertical.md) | EditorGUILayout, GUILayout | See `BeginArea()`. |
| [`BoundsField()`](../ScriptReference/EditorGUI.BoundsField.md) | EditorGUI, EditorGUILayout | [`BoundsField`](../ScriptReference/UIElements.BoundsField.md) |
| [`BoundsIntField()`](../ScriptReference/EditorGUI.BoundsIntField.md) | EditorGUI, EditorGUILayout | [`BoundsIntField`](../ScriptReference/UIElements.BoundsIntField.md) |
| [`Box()`](../ScriptReference/GUI.Box.md) | GUI, GUILayout | [`Box`](../ScriptReference/UIElements.Box.md) |
| [`BringWindowToBack()`](../ScriptReference/GUI.BringWindowToBack.md) | GUI | See `Window()`. |
| [`BringWindowToFront()`](../ScriptReference/GUI.BringWindowToFront.md) | GUI | See `Window()`. |
| [`Button()`](../ScriptReference/GUI.Button.md) | GUI, GUILayout | [`Button`](../ScriptReference/UIElements.Button.md) |
| [`CanCacheInspectorGUI()`](../ScriptReference/EditorGUI.CanCacheInspectorGUI.md) | EditorGUI | Not needed in retained mode. |
| [`ChangeCheckScope`](../ScriptReference/EditorGUI.ChangeCheckScope.md) | EditorGUI | Scopes are generally not needed in UI Toolkit. See `BeginChangeCheck()`. |
| [`ColorField()`](../ScriptReference/EditorGUI.ColorField.md) | EditorGUI, EditorGUILayout | [`ColorField`](../ScriptReference/UIElements.ColorField.md) |
| [`CommandEvent()`](../ScriptReference/EditorGUIUtility.CommandEvent.md) | EditorGUIUtility | Generally not needed in retained mode. Use C# callbacks to handle events. |
| [`CurveField()`](../ScriptReference/EditorGUI.CurveField.md) | EditorGUI, EditorGUILayout | [`CurveField`](../ScriptReference/UIElements.CurveField.md) |
| [`DelayedDoubleField()`](../ScriptReference/EditorGUI.DelayedDoubleField.md) | EditorGUI, EditorGUILayout | [`DoubleField`](../ScriptReference/UIElements.DoubleField.md) with [`isDelayed`](../ScriptReference/UIElements.TextInputBaseField_1-isDelayed.md) set to true. |
| [`DelayedFloatField()`](../ScriptReference/EditorGUI.DelayedFloatField.md) | EditorGUI, EditorGUILayout | [`FloatField`](../ScriptReference/UIElements.FloatField.md) with [`isDelayed`](../ScriptReference/UIElements.TextInputBaseField_1-isDelayed.md) set to true. |
| [`DelayedIntField()`](../ScriptReference/EditorGUI.DelayedIntField.md) | EditorGUI, EditorGUILayout | [`IntegerField`](../ScriptReference/UIElements.IntegerField.md) with [`isDelayed`](../ScriptReference/UIElements.TextInputBaseField_1-isDelayed.md) set to true. |
| [`DelayedTextField()`](../ScriptReference/EditorGUI.DelayedTextField.md) | EditorGUI, EditorGUILayout | [`TextField`](../ScriptReference/UIElements.TextField.md) with [`isDelayed`](../ScriptReference/UIElements.TextInputBaseField_1-isDelayed.md) set to true. |
| [`DisabledScope`](../ScriptReference/EditorGUI.DisabledScope.md) | EditorGUI | Scopes are generally not needed in UI Toolkit. See `BeginDisabledGroup()`. |
| [`DoubleField()`](../ScriptReference/EditorGUI.DoubleField.md) | EditorGUI, EditorGUILayout | [`DoubleField`](../ScriptReference/UIElements.DoubleField.md) |
| [`DragWindow()`](../ScriptReference/GUI.DragWindow.md) | GUI | See `Window()`. |
| [`DrawPreviewTexture()`](../ScriptReference/EditorGUI.DrawPreviewTexture.md) | EditorGUI | No equivalent. |
| [`DrawRect()`](../ScriptReference/EditorGUI.DrawRect.md) | EditorGUI | Use [`VisualElement`](../ScriptReference/UIElements.VisualElement.md). Set [`style.position`](../ScriptReference/UIElements.IStyle-position.md) to [`Absolute`](../ScriptReference/UIElements.Position.Absolute.md). Set [`style.top`](../ScriptReference/UIElements.IStyle-top.md) and [`style.left`](../ScriptReference/UIElements.IStyle-left.md) to define the position. Set [`style.width`](../ScriptReference/UIElements.IStyle-width.md) and [`style.height`](../ScriptReference/UIElements.IStyle-height.md) to define the size. Set [`style.backgroundColor`](../ScriptReference/UIElements.IStyle-backgroundColor.md) to set the color. |
| [`DrawTexture()`](../ScriptReference/GUI.DrawTexture.md) | GUI | [`Image`](../ScriptReference/UIElements.Image.md). Set [`tintColor`](../ScriptReference/UIElements.Image-tintColor.md) in place of `color`. There is no equivalent for a false `alphaBlend`. There are no equivalents for `borderWidth`, `borderWidths`, `borderRadius`, or `borderRadiuses`. |
| [`DrawTextureAlpha()`](../ScriptReference/EditorGUI.DrawTextureAlpha.md) | EditorGUI | No equivalent. |
| [`DrawTextureWithTexCoords()`](../ScriptReference/GUI.DrawTextureWithTexCoords.md) | GUI | [`Image`](../ScriptReference/UIElements.Image.md). Set [`uv`](../ScriptReference/UIElements.Image-uv.md) in place of `texCoords`. There is no equivalent for a false `alphaBlend`. |
| [`DropdownButton()`](../ScriptReference/EditorGUI.DropdownButton.md) | EditorGUI, EditorGUILayout | No exact equivalent. Use fully-fledged [`DropdownField`](../ScriptReference/UIElements.DropdownField.md)s instead of just a `DropdownButton()`. |
| [`DropShadowLabel()`](../ScriptReference/EditorGUI.DropShadowLabel.md) | EditorGUI | [`Label`](../ScriptReference/UIElements.Label.md) with shadow values set in [`style.textShadow`](../ScriptReference/UIElements.IStyle-textShadow.md). |
| [`EditorToolbar()`](../ScriptReference/EditorGUILayout.EditorToolbar.md) | EditorGUILayout | Create a [`Toolbar`](../ScriptReference/UIElements.Toolbar.md) with one [`ToolbarButton`](../ScriptReference/UIElements.ToolbarButton.md) for each tool. For each `ToolbarButton`, register a callback when clicked to call either [`ToolManager.SetActiveTool()`](../ScriptReference/EditorTools.ToolManager.SetActiveTool.md) or [`ToolManager.RestorePreviousTool()`](../ScriptReference/EditorTools.ToolManager.RestorePreviousTool.md) to make that button activate the tool or deactivate it, respectively. |
| [`EndArea()`](../ScriptReference/GUILayout.EndArea.md) | GUILayout | See `BeginArea()`. |
| [`EndBuildTargetSelectionGrouping()`](../ScriptReference/EditorGUILayout.EndBuildTargetSelectionGrouping.md) | EditorGUILayout | See `BeginBuildTargetSelectionGrouping()`. |
| [`EndChangeCheck()`](../ScriptReference/EditorGUI.EndChangeCheck.md) | EditorGUI | See `BeginChangeCheck()`. |
| [`EndDisabledGroup()`](../ScriptReference/EditorGUI.EndDisabledGroup.md) | EditorGUI | See `BeginDisabledGroup()`. |
| [`EndFoldoutHeaderGroup()`](../ScriptReference/EditorGUI.EndFoldoutHeaderGroup.md) | EditorGUI, EditorGUILayout | See `Foldout()`. |
| [`EndGroup()`](../ScriptReference/GUI.EndGroup.md) | GUI | See `BeginArea()`. |
| [`EndHorizontal()`](../ScriptReference/EditorGUILayout.EndHorizontal.md) | EditorGUILayout, GUILayout | See `BeginArea()`. |
| [`EndProperty()`](../ScriptReference/EditorGUI.EndProperty.md) | EditorGUI | See `BeginProperty()`. |
| [`EndScrollView()`](../ScriptReference/GUI.EndScrollView.md) | EditorGUILayout, GUI, GUILayout | See `BeginScrollView()`. |
| [`EndToggleGroup()`](../ScriptReference/EditorGUILayout.EndToggleGroup.md) | EditorGUILayout | See `BeginToggleGroup()`. |
| [`EndVertical()`](../ScriptReference/EditorGUILayout.EndVertical.md) | EditorGUILayout, GUILayout | See `BeginArea()`. |
| [`EnumFlagsField()`](../ScriptReference/EditorGUI.EnumFlagsField.md) | EditorGUI, EditorGUILayout | [`EnumFlagsField`](../ScriptReference/UIElements.EnumFlagsField.md) |
| [`EnumPopup()`](../ScriptReference/EditorGUI.EnumPopup.md) | EditorGUI, EditorGUILayout | [`EnumField`](../ScriptReference/UIElements.EnumField.md) |
| [`ExpandHeight()`](../ScriptReference/GUILayout.ExpandHeight.md) | GUILayout | No equivalent. |
| [`ExpandWidth()`](../ScriptReference/GUILayout.ExpandWidth.md) | GUILayout | No equivalent. |
| [`FlexibleSpace()`](../ScriptReference/GUILayout.FlexibleSpace.md) | GUILayout | See `Space()`. |
| [`FloatField()`](../ScriptReference/EditorGUI.FloatField.md) | EditorGUI, EditorGUILayout | [`FloatField`](../ScriptReference/UIElements.FloatField.md) |
| [`FocusControl()`](../ScriptReference/GUI.FocusControl.md) | GUI | [`VisualElement.Focus()`](../ScriptReference/UIElements.Focusable.Focus.md) |
| [`FocusTextInControl()`](../ScriptReference/EditorGUI.FocusTextInControl.md) | EditorGUI | [`TextField.Focus()`](../ScriptReference/UIElements.Focusable.Focus.md) |
| [`FocusWindow()`](../ScriptReference/GUI.FocusWindow.md) | GUI | See `Window()`. |
| [`Foldout()`](../ScriptReference/EditorGUI.Foldout.md) | EditorGUI, EditorGUILayout | [`Foldout`](../ScriptReference/UIElements.Foldout.md) |
| [`GetControlRect()`](../ScriptReference/EditorGUILayout.GetControlRect.md) | EditorGUILayout | Only needed to convert from EditorGUILayout to EditorGUI. Not needed in UI Toolkit. |
| [`GetNameOfFocusedControl()`](../ScriptReference/GUI.GetNameOfFocusedControl.md) | GUI | [`VisualElement.focusController.focusedElement`](../ScriptReference/UIElements.FocusController-focusedElement.md) |
| [`GetPropertyHeight()`](../ScriptReference/EditorGUI.GetPropertyHeight.md) | EditorGUI | [`PropertyField.layout.height`](../ScriptReference/Rect-height.md) |
| [`GradientField()`](../ScriptReference/EditorGUI.GradientField.md) | EditorGUI, EditorGUILayout | [`GradientField`](../ScriptReference/UIElements.GradientField.md) |
| [`GroupScope`](../ScriptReference/GUI.GroupScope.md) | GUI | Scopes are generally not needed in UI Toolkit. See `BeginArea()`. |
| [`Height()`](../ScriptReference/GUILayout.Height.md) | GUILayout | [`VisualElement.style.height`](../ScriptReference/UIElements.IStyle-height.md) |
| [`HelpBox()`](../ScriptReference/EditorGUI.HelpBox.md) | EditorGUI, EditorGUILayout | [`HelpBox`](../ScriptReference/UIElements.HelpBox.md) |
| [`HorizontalScope`](../ScriptReference/EditorGUILayout.HorizontalScope.md) | EditorGUILayout, GUILayout | Scopes are generally not needed in UI Toolkit. See `BeginArea()`. |
| [`HorizontalScrollbar()`](../ScriptReference/GUI.HorizontalScrollbar.md) | GUI, GUILayout | [`Scroller`](../ScriptReference/UIElements.Scroller.md) with [`direction`](../ScriptReference/UIElements.Scroller-direction.md) set to [`Horizontal`](../ScriptReference/UIElements.SliderDirection.Horizontal.md). |
| [`HorizontalSlider()`](../ScriptReference/GUI.HorizontalSlider.md) | GUI, GUILayout | [`Slider`](../ScriptReference/UIElements.Slider.md) with [`direction`](../ScriptReference/UIElements.BaseSlider_1-direction.md) set to [`Horizontal`](../ScriptReference/UIElements.SliderDirection.Horizontal.md) |
| [`InspectorTitlebar()`](../ScriptReference/EditorGUI.InspectorTitlebar.md) | EditorGUI, EditorGUILayout | No equivalent. |
| [`IntField()`](../ScriptReference/EditorGUI.IntField.md) | EditorGUI, EditorGUILayout | [`IntegerField`](../ScriptReference/UIElements.IntegerField.md) |
| [`IntPopup()`](../ScriptReference/EditorGUI.IntPopup.md) | EditorGUI, EditorGUILayout | No equivalent. |
| [`IntSlider()`](../ScriptReference/EditorGUI.IntSlider.md) | EditorGUI, EditorGUILayout | [`SliderInt`](../ScriptReference/UIElements.SliderInt.md) |
| [`Label()`](../ScriptReference/GUI.Label.md) | GUI, GUILayout | [`Label`](../ScriptReference/UIElements.Label.md) |
| [`LabelField()`](../ScriptReference/EditorGUI.LabelField.md) | EditorGUI, EditorGUILayout | [`TextField`](../ScriptReference/UIElements.TextField.md) with [`isReadOnly`](../ScriptReference/UIElements.TextInputBaseField_1-isReadOnly.md) set to true. |
| [`LayerField()`](../ScriptReference/EditorGUI.LayerField.md) | EditorGUI, EditorGUILayout | [`LayerField`](../ScriptReference/EditorGUI.LayerField.md) |
| [`LinkButton()`](../ScriptReference/EditorGUI.LinkButton.md) | EditorGUI, EditorGUILayout | No equivalent. |
| [`Load()`](../ScriptReference/EditorGUIUtility.Load.md) | EditorGUIUtility | If using C#, you can use this function as is and assign its return value to the `VisualElement.style` property you want. If using USS, use function `resource()` with the same argument you would give to `Load()`. |
| [`LongField()`](../ScriptReference/EditorGUI.LongField.md) | EditorGUI, EditorGUILayout | [`LongField`](../ScriptReference/UIElements.LongField.md) |
| [`MaskField()`](../ScriptReference/EditorGUI.MaskField.md) | EditorGUI, EditorGUILayout | [`MaskField`](../ScriptReference/UIElements.MaskField.md) |
| [`MaxHeight()`](../ScriptReference/GUILayout.MaxHeight.md) | GUILayout | [`VisualElement.style.maxHeight`](../ScriptReference/UIElements.IStyle-maxHeight.md) |
| [`MaxWidth()`](../ScriptReference/GUILayout.MaxWidth.md) | GUILayout | [`VisualElement.style.maxWidth`](../ScriptReference/UIElements.IStyle-maxWidth.md) |
| [`MinHeight()`](../ScriptReference/GUILayout.MinHeight.md) | GUILayout | [`VisualElement.style.minHeight`](../ScriptReference/UIElements.IStyle-minHeight.md) |
| [`MinMaxSlider()`](../ScriptReference/EditorGUI.MinMaxSlider.md) | EditorGUI, EditorGUILayout | [`MinMaxSlider`](../ScriptReference/UIElements.MinMaxSlider.md) |
| [`MinWidth()`](../ScriptReference/GUILayout.MinWidth.md) | GUILayout | [`VisualElement.style.minWidth`](../ScriptReference/UIElements.IStyle-minWidth.md) |
| [`ModalWindow()`](../ScriptReference/GUI.ModalWindow.md) | GUI | See `Window()`. |
| [`[NonReorderable]` attribute](../ScriptReference/NonReorderableAttribute.md) |  | Ensure that [`ListView.reorderable`](../ScriptReference/UIElements.BaseVerticalCollectionView-reorderable.md) is false. |
| [`ObjectField()`](../ScriptReference/EditorGUI.ObjectField.md) | EditorGUI, EditorGUILayout | [`ObjectField`](../ScriptReference/UIElements.ObjectField.md) |
| [`PasswordField()`](../ScriptReference/EditorGUI.PasswordField.md) | EditorGUI, EditorGUILayout, GUI, GUILayout | [`TextField`](../ScriptReference/UIElements.TextField.md) with [`isPasswordField`](../ScriptReference/UIElements.TextInputBaseField_1-isPasswordField.md) set to true |
| [`PixelsToPoints()`](../ScriptReference/EditorGUIUtility.PixelsToPoints.md) | EditorGUIUtility | [`VisualElement.scaledPixelPerPoint`](../ScriptReference/UIElements.VisualElement-scaledPixelsPerPoint.md) |
| [`PointsToPixels()`](../ScriptReference/EditorGUIUtility.PointsToPixels.md) | EditorGUIUtility | use [`VisualElement.scaledPixelPerPoint`](../ScriptReference/UIElements.VisualElement-scaledPixelsPerPoint.md) to do the conversion. |
| [`pixelsPerPoint`](../ScriptReference/EditorGUIUtility.pixelsPerPoint.md) | EditorGUIUtility | use [`VisualElement.scaledPixelPerPoint`](../ScriptReference/UIElements.VisualElement-scaledPixelsPerPoint.md) to do the conversion. |
| [`Popup()`](../ScriptReference/EditorGUI.Popup.md) | EditorGUI, EditorGUILayout | [`PopupField<T0>`](../ScriptReference/UIElements.PopupField_1.md) |
| [`ProgressBar()`](../ScriptReference/EditorGUI.ProgressBar.md) | EditorGUI | [`ProgressBar`](../ScriptReference/UIElements.ProgressBar.md) |
| [`PropertyField()`](../ScriptReference/EditorGUI.PropertyField.md) | EditorGUI, EditorGUILayout | [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) |
| [`PropertyScope`](../ScriptReference/EditorGUI.PropertyScope.md) | EditorGUI | Scopes are generally not needed in UI Toolkit. See `BeginProperty()`. |
| [`RectField()`](../ScriptReference/EditorGUI.RectField.md) | EditorGUI, EditorGUILayout | [`RectField`](../ScriptReference/UIElements.RectField.md) |
| [`RectIntField()`](../ScriptReference/EditorGUI.RectIntField.md) | EditorGUI, EditorGUILayout | [`RectIntField`](../ScriptReference/UIElements.RectIntField.md) |
| [`RepeatButton()`](../ScriptReference/GUI.RepeatButton.md) | GUI, GUILayout | [`RepeatButton`](../ScriptReference/UIElements.RepeatButton.md) |
| [`ScrollTo()`](../ScriptReference/GUI.ScrollTo.md) | GUI | [`ScrollView.ScrollTo()`](../ScriptReference/UIElements.ScrollView.ScrollTo.md) or [`ScrollView.scrollOffset`](../ScriptReference/UIElements.ScrollView-scrollOffset.md) |
| [`ScrollViewScope`](../ScriptReference/EditorGUILayout.ScrollViewScope.md) | EditorGUILayout, GUI, GUILayout | Scopes are generally not needed in UI Toolkit. See `BeginScrollView()`. |
| [`SelectableLabel()`](../ScriptReference/EditorGUI.SelectableLabel.md) | EditorGUI, EditorGUILayout | [`Label`](../ScriptReference/UIElements.Label.md) with [`isSelectable`](../ScriptReference/UIElements.ITextSelection-isSelectable.md) and [`focusable`](../ScriptReference/UIElements.Focusable-focusable.md) set to true. |
| [`SelectionGrid()`](../ScriptReference/GUI.SelectionGrid.md) | GUI, GUILayout | [`RadioButton`](../ScriptReference/UIElements.RadioButton.md) |
| [`SetNextControlName()`](../ScriptReference/GUI.SetNextControlName.md) | GUI | [`VisualElement.name`](../ScriptReference/UIElements.VisualElement-name.md) |
| [`singleLineHeight`](../ScriptReference/EditorGUIUtility-singleLineHeight.md) | EditorGUIUtility | Use USS variable `--unity-metrics-single_line-height`. |
| [`Slider()`](../ScriptReference/EditorGUI.Slider.md) | EditorGUI, EditorGUILayout | [`Slider`](../ScriptReference/UIElements.Slider.md) |
| [`Space()`](../ScriptReference/EditorGUILayout.Space.md) | EditorGUILayout, GUILayout | Use `flex` properties to configure spacing between visual elements. |
| [`TagField()`](../ScriptReference/EditorGUI.TagField.md) | EditorGUI, EditorGUILayout | [`TagField`](../ScriptReference/UIElements.TagField.md) |
| [`TextArea()`](../ScriptReference/EditorGUI.TextArea.md) | EditorGUI, EditorGUILayout, GUI, GUILayout | [`TextField`](../ScriptReference/UIElements.TextField.md) with [`multiline`](../ScriptReference/UIElements.TextField-multiline.md) set to true, [`style.whiteSpace`](../ScriptReference/UIElements.IStyle-whiteSpace.md) set to [`Normal`](../ScriptReference/UIElements.WhiteSpace.Normal.md), and [`ScrollView.verticalScrollerVisibility`](../ScriptReference/UIElements.ScrollView-verticalScrollerVisibility.md) set to [`Auto`](../ScriptReference/UIElements.ScrollerVisibility.Auto.md). |
| [`TextField()`](../ScriptReference/EditorGUI.TextField.md) | EditorGUI, EditorGUILayout, GUI, GUILayout | [`TextField`](../ScriptReference/UIElements.TextField.md) with [`multiline`](../ScriptReference/UIElements.TextField-multiline.md) set to true and [`style.whiteSpace`](../ScriptReference/UIElements.IStyle-whiteSpace.md) set to [`NoWrap`](../ScriptReference/UIElements.WhiteSpace.NoWrap.md). |
| [`Toggle()`](../ScriptReference/EditorGUI.Toggle.md) | EditorGUI, EditorGUILayout, GUI, GUILayout | [`Toggle`](../ScriptReference/UIElements.Toggle.md) |
| [`ToggleGroupScope`](../ScriptReference/EditorGUILayout.ToggleGroupScope.md) | EditorGUILayout | Scopes are generally not needed in UI Toolkit. See `BeginToggleGroup()`. |
| [`ToggleLeft()`](../ScriptReference/EditorGUI.ToggleLeft.md) | EditorGUI, EditorGUILayout | [`Toggle`](../ScriptReference/UIElements.Toggle.md), but instead of setting [`label`](../ScriptReference/UIElements.BaseField_1-label.md), set [`text`](../ScriptReference/UIElements.BaseBoolField-text.md). |
| [`Toolbar()`](../ScriptReference/GUI.Toolbar.md) | GUI, GUILayout | No equivalent. |
| [`UnfocusWindow()`](../ScriptReference/GUI.UnfocusWindow.md) | GUI | See `Window()`. |
| [`Vector2Field()`](../ScriptReference/EditorGUI.Vector2Field.md) | EditorGUI, EditorGUILayout | [`Vector2Field`](../ScriptReference/UIElements.Vector2Field.md) |
| [`Vector2IntField()`](../ScriptReference/EditorGUI.Vector2IntField.md) | EditorGUI, EditorGUILayout | [`Vector2IntField`](../ScriptReference/UIElements.Vector2IntField.md) |
| [`Vector3Field()`](../ScriptReference/EditorGUI.Vector3Field.md) | EditorGUI, EditorGUILayout | [`Vector3Field`](../ScriptReference/UIElements.Vector3Field.md) |
| [`Vector3IntField()`](../ScriptReference/EditorGUI.Vector3IntField.md) | EditorGUI, EditorGUILayout | [`Vector3IntField`](../ScriptReference/UIElements.Vector3IntField.md) |
| [`Vector4Field()`](../ScriptReference/EditorGUI.Vector4Field.md) | EditorGUI, EditorGUILayout | [`Vector4Field`](../ScriptReference/UIElements.Vector4Field.md) |
| [`VerticalScope`](../ScriptReference/EditorGUILayout.VerticalScope.md) | EditorGUILayout, GUILayout | Scopes are generally not needed in UI Toolkit. See `BeginArea()`. |
| [`VerticalScrollbar()`](../ScriptReference/GUI.VerticalScrollbar.md) | GUI, GUILayout | [`Scroller`](../ScriptReference/UIElements.Scroller.md) with [`direction`](../ScriptReference/UIElements.Scroller-direction.md) set to [`Vertical`](../ScriptReference/UIElements.SliderDirection.Vertical.md). |
| [`VerticalSlider()`](../ScriptReference/GUI.VerticalSlider.md) | GUI, GUILayout | [`Slider`](../ScriptReference/UIElements.Slider.md) with [`direction`](../ScriptReference/UIElements.BaseSlider_1-direction.md) set to [`Vertical`](../ScriptReference/UIElements.SliderDirection.Vertical.md). |
| [`Width()`](../ScriptReference/GUILayout.Width.md) | GUILayout | [`VisualElement.style.width`](../ScriptReference/UIElements.IStyle-width.md) |
| [`Window()`](../ScriptReference/GUI.Window.md) | GUI, GUILayout | No equivalent. |

## Additional resources

* [Get started with UI Toolkit](UIE-simple-ui-toolkit-workflow.md)
* [Unity style sheets (USS)](UIE-USS.md)
* [Examples](UIE-examples.md)
* [IMGUI events](UIE-IMGUI-Events.md)