# Image

Use the Image element to add graphical assets to your application. It’s a versatile and powerful tool for displaying visual content in your UI, such as image galleries, product listings, a rendering preview, and user profiles.

## Create an image

You can create an Image element in UI Builder, UXML, or C#.

The following example creates an Image element in UXML:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" >
    <ui:Image source="Resources/sample-image.png" scale-mode="StretchToFill" tint-color="rgb(49, 49, 49)"/>
</ui:UXML>
```

The following example creates an Image element in C#:

```
var image = new Image();
image.image = Resources.Load<Texture2D>("sample-image");
// Alternative if using sprites.
// image.sprite = Resources.Load<Sprite>("sample-sprite");
image.scaleMode = ScaleMode.StretchToFill;
image.tintColor = Color.white;
```

## Image versus `VisualElement.style.backgroundImage`

You can use both the Image element and the `VisualElement.style.backgroundImage` property to add visual content to your UI. The choice between the two depends on the specific usage and requirements of your application.

The Image element is typically used to display an image within a specific area of a UI, such as an image gallery or an **avatar** in a user profile. It provides more fine-grained control over the appearance of the image, including options for resizing, scaling, and cropping. Use Image when you want the size of an image to drive the size of an element in the layout. You can also add additional elements to the Image element to create overlays or add interactivity. However, for performance reasons, specify a fixed width and height if you use many instances of Image.

The `VisualElement.style.backgroundImage` property is used to [set an image as the background of a UI element](UIB-styling-ui-backgrounds.md). This is useful when you want to add visual interest to an element without detracting from the main content or functionality of the UI. The `VisualElement.style.backgroundImage` property also allows the convenience of using USS to change the source image for many different elements.

In general, it’s recommended to use the Image element when you need more control over the appearance and layout of an image; and to use the `VisualElement.style.backgroundImage` property when you simply want to add a background image to a UI element. However, both approaches are valid and you can use them together in the same UI if necessary.

**Note**:

* The Image element also supports the `style.backgroundImage` property. The background image won’t be visible if the Image element has an image source set and it’s not transparent.
* UI Builder partially supports the Image element. It reads from and writes to UXML, and allows basic authoring as a standard **visual element**. Most of the Image-specific settings are defined as USS custom properties and aren’t visible in the **Attributes** section of the UI Builder **Inspector**.

## C# base class and namespace

**C# class**: [`Image`](../ScriptReference/UIElements.Image.md)  
**Namespace**: `UnityEngine.UIElements`  
**Base class**: [`VisualElement`](../ScriptReference/UIElements.VisualElement.md)

## Member UXML attributes

This element has the following member attributes:

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| `scale-mode` | `ScaleMode` | ScaleMode used to display the Image. |
| `tint-color` | `Color` | Tinting color for this Image. |
| `uv` | `Rect` | The base texture coordinates of the Image relative to the bottom left corner. |

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
| `ussClassName` | `.unity-image` | USS class name of elements of this type. |
| `disabledUssClassName` | `.unity-disabled` | USS class name of local disabled elements. |

## USS custom properties

The following table outlines the custom properties that are available exclusively for the Image element in USS:

| Property | Type | Description |
| --- | --- | --- |
| `--unity-image` | string | The source of the image file. For information on how to reference an image asset, refer to [Assets](UIE-USS-PropertyTypes.md#assets). |
| `--unity-image-size` | string | The image scale mode. Available values are `stretch-to-fill`, `scale-and-crop` and `scale-to-fit`. |
| `--unity-image-tint-color` | string | The image tint color. |

## Additional resources

* [Set background images](UIB-styling-ui-backgrounds.md)
* [Button](UIE-uxml-element-Button.md)