# Tab

A Tab element represents a single tab within a [TabView](UIE-uxml-element-TabView.md). In a window or menu, you can use a tab to group related content.

## Make a Tab closable

To make it so a Tab can be closed, set the [`closable`](../ScriptReference/UIElements.Tab-closeable.md) property to `true`. When a Tab is closable, a close icon appears on the Tab. If a user selects the close icon, the Tab closes.

## Add an icon to a Tab

You can add an icon to your Tab to make it more visually appealing. Icons can be image assets that exist in your project such as a Texture, **Render Texture**, **Sprite**, or Vector. For information on how to reference an image asset, refer to [Assets](UIE-USS-PropertyTypes.md#assets).

Do one of the following to add an icon to a Tab with UI Builder:

* From the Tab’s **Inspector** panel, select an icon from the **Icon Image** dropdown list.
* Drag an icon from the Assets window to the **Icon Image** field located in the Tab’s Inspector panel.

To add an icon for a Tab with UXML, add the source of the image to the `icon-image` attribute:

```
<ui:Tab name="Tab" label="Tab text" icon-image="/path/to/image-file.png" />
```

To add an icon for a Tab with C#, assign the image source to the [`iconImage`](../ScriptReference/UIElements.Tab-iconImage.md) property:

```
Tab myTab = new Tab("Tab text");
var TabIconImage = Resources.Load<Texture2D>("image-file");

myTab.iconImage = TabIconImage;
```

## Examples

The following UXML example creates a TabView with Tabs:

```
<UXML xmlns="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <TabView>
       <Tab label="UXML Tab A">
           <Label text="UXML tab: This is some content for the first Tab." />
       </Tab>
       <Tab label="UXML Tab B">
           <Label text="UXML tab: This is some content for the second Tab." />
       </Tab>
    </TabView>
</UXML>
```

The following C# example illustrates some of the customizable functionalities of the TabView and its Tabs:

```
/// <sample>
// Create a TabView with Tabs that only contains a label.
var csharpTabViewWithLabels = new TabView() { style = { marginTop = 15 } }; // marginTop not required, only for demonstration purposes.
var tabOne = new Tab("One");
tabOne.Add(new Label("Tab with labels only: This is some content for the first Tab.") { style = { marginTop = 10 } });
csharpTabViewWithLabels.Add(tabOne);
var tabTwo = new Tab("Two");
tabTwo.Add(new Label("Tab with labels only: This is some content for the second Tab.") { style = { marginTop = 10 } });
csharpTabViewWithLabels.Add(tabTwo);
container.Add(csharpTabViewWithLabels);

// Create a TabView with Tabs that only contains an icon.
var csharpTabViewWithIcons = new TabView() { style = { marginTop = 15 } }; // marginTop not required, only for demonstration purposes.
var tabIconConnect = new Tab(EditorGUIUtility.FindTexture("CloudConnect"));
tabIconConnect.Add(new Label("Tab with icons only: This is some content for the first Tab.") { style = { marginTop = 10 } });
csharpTabViewWithIcons.Add(tabIconConnect);
var tabIconStore = new Tab(EditorGUIUtility.FindTexture("Asset Store"));
tabIconStore.Add(new Label("Tab with icons only: This is some content for the second Tab.") { style = { marginTop = 10 } });
csharpTabViewWithIcons.Add(tabIconStore);
container.Add(csharpTabViewWithIcons);

// Create a TabView with Tabs that only contains an icon and a label.
var csharpTabViewWithIconsAndLabels = new TabView() { style = { marginTop = 15 } }; // marginTop not required, only for demonstration purposes.
var tabConnect = new Tab("Connect", EditorGUIUtility.FindTexture("CloudConnect"));
tabConnect.Add(new Label("Tab with an icon and a labels: This is some content for the first Tab.") { style = { marginTop = 10 } });
csharpTabViewWithIconsAndLabels.Add(tabConnect);
var tabStore = new Tab("Store", EditorGUIUtility.FindTexture("Asset Store"));
tabStore.Add(new Label("Tab with an icon and a labels: This is some content for the second Tab.") { style = { marginTop = 10 } });
csharpTabViewWithIconsAndLabels.Add(tabStore);
container.Add(csharpTabViewWithIconsAndLabels);

// Create a TabView that allows re-ordering of the tabs.
var csharpReorderableTabView = new TabView() { reorderable = true, style = { marginTop = 10 } }; // marginTop not required, only for demonstration purposes.
var tabA = new Tab("Tab A");
tabA.Add(new Label("Reorderable tabs: This is some content for Tab A") { style = { marginTop = 10 } });
csharpReorderableTabView.Add(tabA);
var tabB = new Tab("Tab B");
tabB.Add(new Label("Reorderable tabs: This is some content for Tab B") { style = { marginTop = 10 } });
csharpReorderableTabView.Add(tabB);
var tabC = new Tab("Tab C");
tabC.Add(new Label("Reorderable tabs: This is some content for Tab C") { style = { marginTop = 10 } });
csharpReorderableTabView.Add(tabC);
container.Add(csharpReorderableTabView);

// Create a TabView with closeable tabs.
var closeTabInfoLabel = new Label($"Last tab closed: None");
void UpdateLabel(string newLabel) => closeTabInfoLabel.text = $"Last tab closed: {newLabel}";
var cSharpCloseableTabs = new TabView() { style = { marginTop = 10 } }; // marginTop not required, only for demonstration purposes.
var closeableTabA = new Tab("Title A") { closeable = true };
closeableTabA.closed += (tab) => { UpdateLabel(tab.label); };
closeableTabA.Add(new Label("Closeable tabs: This is some content for Tab A") { style = { marginTop = 10 } });
cSharpCloseableTabs.Add(closeableTabA);
var closeableTabB = new Tab("Title B") { closeable = true };
closeableTabB.closed += (tab) => { UpdateLabel(tab.label); };
closeableTabB.Add(new Label("Closeable tabs: This is some content for Tab B") { style = { marginTop = 10 } });
cSharpCloseableTabs.Add(closeableTabB);
var closeableTabC = new Tab("Title C") { closeable = true };
closeableTabC.closed += (tab) => { UpdateLabel(tab.label); };
closeableTabC.Add(new Label("Closeable tabs: This is some content for Tab C") { style = { marginTop = 10 } });
cSharpCloseableTabs.Add(closeableTabC);
container.Add(cSharpCloseableTabs);
container.Add(closeTabInfoLabel);

// Create a TabView and apply custom styling to specific areas of their tabs.
var csharpCustomStyledTabView = new TabView() { style = { marginTop = 15 }, classList = { "some-styled-class" }}; // marginTop not required, only for demonstration purposes.
var customStyledTabOne = new Tab("One");
customStyledTabOne.Add(new Label("Custom styled tabs: This is some content for the first Tab."));
csharpCustomStyledTabView.Add(customStyledTabOne);
var customStyledTabTwo = new Tab("Two");
customStyledTabTwo.Add(new Label("Custom styled tabs: This is some content for the second Tab."));
csharpCustomStyledTabView.Add(customStyledTabTwo);
container.Add(csharpCustomStyledTabView);
/// </sample>
```

To try this example live in Unity, go to **Window** > **UI Toolkit** > **Samples**.

For more examples, refer to the following:

-[Create a tabbed menu](UIE-create-tabbed-menu-for-runtime.md).

## C# base class and namespace

**C# class**: [`Tab`](../ScriptReference/UIElements.Tab.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`VisualElement`](../ScriptReference/UIElements.VisualElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `closeable` | `boolean` | A property that adds the ability to close tabs.  The default value is `false`. Set this value to `true` to allow the user to close tabs in the tab view. |
| `icon-image` | `Object` | Sets the icon for the Tab’s header. |
| `label` | `string` | Sets the label of the Tab’s header. |

## Inherited UXML attributes

This element inherits the following attributes from its base class:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `focusable` | `boolean` | If false, this prevents the element from being focused.  The element can only be focused if its canGrabFocus property is true. |
| `tabindex` | `int` | An integer used to sort focusable elements in the focus ring. Must be greater than or equal to zero.  Setting the `tabIndex` value to less than `0` (for example, `−1`) removes the element from the focus ring and tab navigation. |

This element also inherits the following attributes from [`VisualElement`](UIE-uxml-element-VisualElement.md):

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `content-container` | `string` | Logical container where child elements are added. If a child is added to this element, the child is added to this element’s content container instead.  When iterating over the `VisualElement.Children` of an element, the element’s content container hierarchy is used instead of the element itself. This can lead to unexpected results, such as elements being ignored by the navigation events if they are not directly in the content container’s hierarchy. Refer to `IFocusRing` for more information.    If the content container is the same as the element itself, child elements are added directly to the element. This is true for most elements but can be overridden by more complex types.   The `ScrollView`, for example, has a content container that is different from itself. In that case, child elements added to the scroll view are added to its content container element instead. While the physical parent (`VisualElement.Hierarchy.parent`) of the child elements is the scroll view’s content container element, their logical parent (`VisualElement.parent`) still refers to the scroll view itself. Since some of the scroll view’s focusable children are not part of its logical hierarchy, like its `Scroller` elements, these focusable children are not considered by default when using sequential navigation. Refer to [How can I change what element is focused next](UIE-faq-event-and-input-system.md) for an example of a workaround solution if the default navigation rules don’t correspond to your needs. |
| `data-source` | `Object` | Assigns a data source to this VisualElement which overrides any inherited data source. This data source is inherited by all children. |
| `data-source-path` | `string` | Path from the data source to the value. |
| `data-source-type` | `System.Type` | The possible type of data source assignable to this VisualElement.    This information is only used by the UI Builder as a hint to provide some completion to the data source path field when the effective data source cannot be specified at design time. |
| `enabled` | `boolean` | Returns true if the `VisualElement` is enabled locally.  This flag isn’t changed if the VisualElement is disabled implicitly by one of its parents. To verify this, use `enabledInHierarchy`. |
| `language-direction` | [`UIElements.LanguageDirection`](../ScriptReference/UIElements.LanguageDirection.md) | Indicates the directionality of the element’s text. The value will propagate to the element’s children.  Setting `languageDirection` to `RTL` can only get the basic RTL support like text reversal. To get more comprehensive RTL support, such as line breaking, word wrapping, or text shaping, you must enable [Advance Text Generator](UIE-advanced-text-generator.md). |
| `name` | `string` | The name of this VisualElement.  Use this property to write USS selectors that target a specific element. The standard practice is to give an element a unique name. |
| `picking-mode` | [`UIElements.PickingMode`](../ScriptReference/UIElements.PickingMode.md) | Determines if this element can be the target of pointer events or picked by `IPanel.Pick` queries.  Elements can not be picked if:  - They are invisible- Their `style.display` is set to `DisplayStyle.None`   Elements with a picking mode of `PickingMode.Ignore` never receive the hover pseudo-state. |
| `style` | `string` | Sets the style values on a `VisualElement`.  The returned style data, computed from USS files or inline styles written to this object in C#, doesn’t represent the fully resolved styles, such as the final height and width of a VisualElement. To access these fully resolved styles, use `resolvedStyle`.      For information about how to use this property and all the supported USS properties, refer to the [Apply styles in C# scripts](UIE-apply-styles-with-csharp.md) and [USS properties reference](UIE-USS-Properties-Reference.md) manual pages. |
| `tooltip` | `string` | Text to display inside an information box after the user hovers the element for a small amount of time. This is only supported in the Editor UI. |
| `usage-hints` | [`UIElements.UsageHints`](../ScriptReference/UIElements.UsageHints.md) | A combination of hint values that specify high-level intended usage patterns for the `VisualElement`. This property can only be set when the `VisualElement` is not yet part of a `Panel`. Once part of a `Panel`, this property becomes effectively read-only, and attempts to change it will throw an exception. The specification of proper `UsageHints` drives the system to make better decisions on how to process or accelerate certain operations based on the anticipated usage pattern. Note that those hints do not affect behavioral or visual results, but only affect the overall performance of the panel and the elements within. It’s advised to always consider specifying the proper `UsageHints`, but keep in mind that some `UsageHints` might be internally ignored under certain conditions (e.g. due to hardware limitations on the target platform). |
| `view-data-key` | `string` | Used for view data persistence, such as tree expanded states, scroll position, or zoom level.  This key is used to save and load the view data from the view data store. If you don’t set this key, the persistence is disabled for the associated `VisualElement`. For more information, refer to [View data persistence in the Unity Manual](UIE-ViewData.md). |

## USS classes

The following table lists all the C# public property names and their related USS selector.

| **C# property** | **USS selector** | **Description** |
| --- | --- | --- |
| `ussClassName` | `.unity-tab` | USS class name of elements of this type. |
| `tabHeaderUssClassName` | `.unity-tab__header` | USS class name for the header of this type. |
| `tabHeaderImageUssClassName` | `.unity-tab__header-image` | USS class name for the icon inside the header. |
| `tabHeaderEmptyImageUssClassName` | `.unity-tab__header-image--empty` | USS class name for the icon inside the header when the value is null. |
| `tabHeaderStandaloneImageUssClassName` | `.unity-tab__header-image--standalone` | USS class name for the icon inside the header when the label is empty or null. |
| `tabHeaderLabelUssClassName` | `.unity-tab__header-label` | USS class name for the label of the header. |
| `tabHeaderEmptyLabeUssClassName` | `.unity-tab__header-label--empty` | USS class name for the label of the header when the value is empty or null. |
| `tabHeaderUnderlineUssClassName` | `.unity-tab__header-underline` | USS class name for the active state underline of the header. |
| `contentUssClassName` | `.unity-tab__content-container` | USS class name of container element of this type. |
| `draggingUssClassName` | `.unity-tab--dragging` | USS class name for the dragging state of this type. |
| `reorderableUssClassName` | `.unity-tab__reorderable` | USS class name for reorderable tab elements. |
| `reorderableItemHandleUssClassName` | `.unity-tab__reorderable-handle` | USS class name for drag handle in reorderable tabs. |
| `reorderableItemHandleBarUssClassName` | `.unity-tab__reorderable-handle-bar` | USS class name for drag handlebar in reorderable tabs. |
| `closeableUssClassName` | `.unity-tab__header__closeable` | The USS class name for a closeable tab. |
| `closeButtonUssClassName` | `.unity-tab__close-button` | The USS class name for close button in closable tabs. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

You can also use the [Matching Selectors section in the Inspector or the UI Toolkit Debugger](UIB-testing-ui.md#matching-selectors) to see which USS selectors affect the components of the `VisualElement` at every level of its hierarchy.

## Additional resources

* [TabView](UIE-uxml-element-TabView.md)
* 📺 **Video**: [UI Toolkit Tutorial Series: Tabbed Menu](https://www.youtube.com/watch?v=sF33-iNy0rU)