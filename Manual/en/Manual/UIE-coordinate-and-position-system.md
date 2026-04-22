# Coordinate and position systems

UI Toolkit uses a powerful layout system that automatically calculates the position and size of individual elements based on the layout parameters in their style properties. This is based on [Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox), a web layout model. For more information, refer to [Layout Engine](UIE-LayoutEngine.md).

## Relative and absolute positions

In UI Toolkit, coordinates refer to the position of an element within a layout. You can configure each element in the layout system to use either relative or absolute coordinates:

* **Relative**: Coordinates relative to the element’s calculated position. The layout system calculates the position of the element, then adds the coordinates as an offset. Parent elements can influence the size and position of the child elements, as the layout engine takes them into account when it calculates the element position. Child elements can only influence the size of the parent element.
* **Absolute**: Coordinates relative to the parent element. This circumvents the automatic layout calculation and directly sets the position of the element. Sibling elements under the same parent have no influence on the element’s position. Similarly, the element doesn’t influence the position and size of other siblings under the same parent.

Each **visual element** determines the coordinate system used to calculate its position. You can configure which coordinate system to use in the element stylesheet.

The following C# code snippets set the coordinate space and the position of a visual element:

```
var newElement = new VisualElement();
newElement.style.position = Position.Relative;
newElement.style.left = 15;
newElement.style.top = 35;
```

The origin of an element is its top left corner.

The layout system computes the [`VisualElement.layout`](../ScriptReference/UIElements.VisualElement-layout.md) property (type `Rect`) for each element, which includes the final position of the element. This takes the relative or absolute position of the element into account.

The `layout.position` is expressed in points, relative to the coordinate space of its parent.

Each `VisualElement` has a transform property (`ITransform`) you can use to add an additional local offset to the position and rotation of an element. The offset isn’t represented in the calculated layout property. By default, the `transform` is the identity.

Use the [`worldBound`](../ScriptReference/UIElements.VisualElement-worldBound.md) property to retrieve the final window space coordinates of the `VisualElement`, taking into account both the layout position and the transform. This position includes the height of the header of the window.

## Transformation between coordinate systems

The `VisualElement.layout.position` and `VisualElement.transform` properties define how to transform between the local coordinate system and the parent coordinate system.

The [`VisualElementExtensions`](../ScriptReference/UIElements.VisualElementExtensions.md) static class provides the following extension methods that transform points and rectangles between coordinate systems:

* [`WorldToLocal`](../ScriptReference/UIElements.VisualElementExtensions.WorldToLocal.md) transforms a `Vector2` or `Rect` from `Panel` space to the referential within the element.
* [`LocalToWorld`](../ScriptReference/UIElements.VisualElementExtensions.LocalToWorld.md) transforms a `Vector2` or `Rect` to `Panel` space referential.
* [`ChangeCoordinatesTo`](../ScriptReference/UIElements.VisualElementExtensions.ChangeCoordinatesTo.md) transforms `Vector2` or `Rect` from the local space of one element to the local space of another.

## Additional resource

* [Relative and absolute positioning C# example](UIE-relative-absolute-positioning-example.md)