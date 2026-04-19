# USS properties reference

The following table lists the USS properties in UI Toolkit.

The table indicates which properties are inherited and their [animatability](UIE-Transitions.md#discrete). For an inherited property, the element gets the value from its parent element if you don’t specify a value for it. The following example uses an inherited property to set the font for all elements.

```
:root {
    -unity-font: resource("Font/consola.ttf");
}
```

Click the property name to see detailed information about that property.

**Note**: The usage of a USS property that’s the same as the CSS property links to the Mozilla developer documentation.

| **Property** | **Inherited** | **Animatability** | **Description** |
| --- | --- | --- | --- |
| **[`align-content`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Alignment of the whole area of children on the cross axis if they span over multiple lines in this container. |
| **[`align-items`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Alignment of children on the cross axis of this container. |
| **[`align-self`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Similar to align-items, but only for this specific element. |
| **[`all`](UIE-USS-SupportedProperties.md#all)** | No | Fully animatable | Allows resetting all properties with the initial keyword. Does not apply to custom USS properties. |
| **[`aspect-ratio`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Sets a preferred **aspect ratio** for the box, which will be used in the calculation of auto sizes and some other layout functions. |
| **[`background-color`](UIE-USS-SupportedProperties.md#unity-background)** | No | Fully animatable | Background color to paint in the element’s box. |
| **[`background-image`](UIE-USS-SupportedProperties.md#unity-background)** | No | Discrete | Background image to paint in the element’s box. |
| **[`background-position`](https://developer.mozilla.org/en-US/docs/Web/CSS/background-position)** | No | Fully animatable | Background image position value. |
| **[`background-position-x`](https://developer.mozilla.org/en-US/docs/Web/CSS/background-position-x)** | No | Discrete | Background image x position value. |
| **[`background-position-y`](https://developer.mozilla.org/en-US/docs/Web/CSS/background-position-y)** | No | Discrete | Background image y position value. |
| **[`background-repeat`](https://developer.mozilla.org/en-US/docs/Web/CSS/background-repeat)** | No | Discrete | Background image repeat value. |
| **[`background-size`](https://developer.mozilla.org/en-US/docs/Web/CSS/background-size)** | No | Fully animatable | Background image size value. Transitions are fully supported only when using size in **pixels** or percentages, such as pixel-to-pixel or percentage-to-percentage transitions. |
| **[`border-bottom-color`](UIE-USS-SupportedProperties.md#border-color)** | No | Fully animatable | Color of the element’s bottom border. |
| **[`border-bottom-left-radius`](UIE-USS-SupportedProperties.md#drawing-borders)** | No | Fully animatable | The radius of the bottom-left corner when a rounded rectangle is drawn in the element’s box. |
| **[`border-bottom-right-radius`](UIE-USS-SupportedProperties.md#drawing-borders)** | No | Fully animatable | The radius of the bottom-right corner when a rounded rectangle is drawn in the element’s box. |
| **[`border-bottom-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the bottom edge of the border during the layout phase. |
| **[`border-color`](UIE-USS-SupportedProperties.md#border-color)** | No | Fully animatable | Shorthand for border-top-color, border-right-color, border-bottom-color, border-left-color |
| **[`border-left-color`](UIE-USS-SupportedProperties.md#border-color)** | No | Fully animatable | Color of the element’s left border. |
| **[`border-left-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the left edge of the border during the layout phase. |
| **[`border-radius`](UIE-USS-SupportedProperties.md#drawing-borders)** | No | Fully animatable | Shorthand for border-top-left-radius, border-top-right-radius, border-bottom-right-radius, border-bottom-left-radius |
| **[`border-right-color`](UIE-USS-SupportedProperties.md#border-color)** | No | Fully animatable | Color of the element’s right border. |
| **[`border-right-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the right edge of the border during the layout phase. |
| **[`border-top-color`](UIE-USS-SupportedProperties.md#border-color)** | No | Fully animatable | Color of the element’s top border. |
| **[`border-top-left-radius`](UIE-USS-SupportedProperties.md#drawing-borders)** | No | Fully animatable | The radius of the top-left corner when a rounded rectangle is drawn in the element’s box. |
| **[`border-top-right-radius`](UIE-USS-SupportedProperties.md#drawing-borders)** | No | Fully animatable | The radius of the top-right corner when a rounded rectangle is drawn in the element’s box. |
| **[`border-top-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the top edge of the border during the layout phase. |
| **[`border-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Shorthand for border-top-width, border-right-width, border-bottom-width, border-left-width |
| **[`bottom`](UIE-USS-SupportedProperties.md#positioning)** | No | Fully animatable | Bottom distance from the element’s box during layout. |
| **[`color`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Color to use when drawing the text of an element. |
| **[`cursor`](UIE-USS-SupportedProperties.md#cursor)** | No | Non-animatable | Mouse cursor to display when the mouse pointer is over an element. |
| **[`display`](UIE-USS-SupportedProperties.md#appearance)** | No | Non-animatable | Defines how an element is displayed in the layout. |
| **[`filter`](uss-filter)** | No | Fully animatable | Filter effects to apply to the element. |
| **[`flex`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Fully animatable | Shorthand for flex-grow, flex-shrink, flex-basis |
| **[`flex-basis`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Fully animatable | Initial main size of a flex item, on the main flex axis. The final layout might be smaller or larger, according to the flex shrinking and growing determined by the other flex properties. |
| **[`flex-direction`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Direction of the main axis to layout children in a container. |
| **[`flex-grow`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Fully animatable | Specifies how the item will grow relative to the rest of the flexible items inside the same container. |
| **[`flex-shrink`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Fully animatable | Specifies how the item will shrink relative to the rest of the flexible items inside the same container. |
| **[`flex-wrap`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Placement of children over multiple lines if not enough space is available in this container. |
| **[`font-size`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Font size to draw the element’s text, specified in point size. |
| **[`height`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Fixed height of an element for the layout. |
| **[`justify-content`](UIE-USS-SupportedProperties.md#flex-layout)** | No | Discrete | Justification of children on the main axis of this container. |
| **[`left`](UIE-USS-SupportedProperties.md#positioning)** | No | Fully animatable | Left distance from the element’s box during layout. |
| **[`letter-spacing`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Increases or decreases the space between characters. |
| **[`margin`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Shorthand for margin-top, margin-right, margin-bottom, margin-left |
| **[`margin-bottom`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the bottom edge of the margin during the layout phase. |
| **[`margin-left`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the left edge of the margin during the layout phase. |
| **[`margin-right`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the right edge of the margin during the layout phase. |
| **[`margin-top`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the top edge of the margin during the layout phase. |
| **[`max-height`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Maximum height for an element, when it is flexible or measures its own size. |
| **[`max-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Maximum width for an element, when it is flexible or measures its own size. |
| **[`min-height`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Minimum height for an element, when it is flexible or measures its own size. |
| **[`min-width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Minimum width for an element, when it is flexible or measures its own size. |
| **[`opacity`](UIE-USS-SupportedProperties.md#opacity)** | No | Fully animatable | Specifies the transparency of an element and of its children. |
| **[`overflow`](UIE-USS-SupportedProperties.md#appearance)** | No | Discrete | How a container behaves if its content overflows its own box. |
| **[`padding`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Shorthand for padding-top, padding-right, padding-bottom, padding-left |
| **[`padding-bottom`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the bottom edge of the padding during the layout phase. |
| **[`padding-left`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the left edge of the padding during the layout phase. |
| **[`padding-right`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the right edge of the padding during the layout phase. |
| **[`padding-top`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Space reserved for the top edge of the padding during the layout phase. |
| **[`position`](UIE-USS-SupportedProperties.md#positioning)** | No | Discrete | Element’s positioning in its parent container. |
| **[`right`](UIE-USS-SupportedProperties.md#positioning)** | No | Fully animatable | Right distance from the element’s box during layout. |
| **[`rotate`](UIE-Transform.md)** | No | Fully animatable | A rotation transformation. |
| **[`scale`](UIE-Transform.md)** | No | Fully animatable | A scaling transformation. |
| **[`text-overflow`](UIE-USS-SupportedProperties.md#unity-text)** | No | Discrete | The element’s text overflow mode. |
| **[`text-shadow`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Drop shadow of the text. |
| **[`top`](UIE-USS-SupportedProperties.md#positioning)** | No | Fully animatable | Top distance from the element’s box during layout. |
| **[`transform-origin`](UIE-Transform.md)** | No | Fully animatable | The transformation origin is the point around which a transformation is applied. |
| **[`transition`](UIE-Transitions.md)** | No | Non-animatable | Shorthand for transition-delay, transition-duration, transition-property, transition-timing-function |
| **[`transition-delay`](UIE-Transitions.md)** | No | Non-animatable | Duration to wait before starting a property’s transition effect when its value changes. |
| **[`transition-duration`](UIE-Transitions.md)** | No | Non-animatable | Time a transition animation should take to complete. |
| **[`transition-property`](UIE-Transitions.md)** | No | Non-animatable | Properties to which a transition effect should be applied. |
| **[`transition-timing-function`](UIE-Transitions.md)** | No | Non-animatable | Determines how intermediate values are calculated for properties modified by a transition effect. |
| **[`translate`](UIE-Transform.md)** | No | Fully animatable | A translate transformation. |
| **[`-unity-background-image-tint-color`](UIE-USS-SupportedProperties.md#unity-background)** | No | Fully animatable | Tinting color for the element’s backgroundImage. |
| **[`-unity-background-scale-mode`](UIE-USS-SupportedProperties.md#unity-background)** | No | Discrete | Background image scaling in the element’s box. |
| **[`-unity-editor-text-rendering-mode`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Non-animatable | TextElement editor rendering mode. |
| **[`-unity-font`](UIE-USS-SupportedProperties.md#unity-font)** | Yes | Discrete | Font to draw the element’s text, defined as a Font object. |
| **[`-unity-font-definition`](UIE-USS-SupportedProperties.md#unity-font)** | Yes | Discrete | Font to draw the element’s text, defined as a FontDefinition structure. It takes precedence over `-unity-font`. |
| **[`-unity-font-style`](UIE-USS-SupportedProperties.md#unity-font)** | Yes | Discrete | Font style and weight (normal, bold, italic) to draw the element’s text. |
| **[`-unity-material`](UIE-USS-SupportedProperties.md#unity-material)** | Yes | Fully animatable | Custom material to use on the element. |
| **[`-unity-overflow-clip-box`](UIE-USS-SupportedProperties.md#appearance)** | No | Discrete | Specifies which box the element content is clipped against. |
| **[`-unity-paragraph-spacing`](UIE-USS-SupportedProperties.md#appearance)** | Yes | Fully animatable | Increases or decreases the space between paragraphs. |
| **[`-unity-slice-bottom`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Fully animatable | Size of the 9-slice’s bottom edge when painting an element’s background image. |
| **[`-unity-slice-left`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Fully animatable | Size of the 9-slice’s left edge when painting an element’s background image. |
| **[`-unity-slice-right`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Fully animatable | Size of the 9-slice’s right edge when painting an element’s background image. |
| **[`-unity-slice-scale`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Fully animatable | Scale applied to an element’s slices. |
| **[`-unity-slice-top`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Fully animatable | Size of the 9-slice’s top edge when painting an element’s background image. |
| **[`-unity-slice-type`](UIE-USS-SupportedProperties.md#unity-slice)** | No | Discrete | Specifies the type of sclicing. |
| **[`-unity-text-align`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Discrete | Horizontal and vertical text alignment in the element’s box. |
| **[`-unity-text-auto-size`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Non-animatable | Overrides any explicit font-size to scale text within the defined minimum and maximum bounds, recalculating as needed to fit its container. |
| **[`-unity-text-generator`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Non-animatable | Switches between Unity’s standard and advanced text generator |
| **[`-unity-text-outline`](UIE-USS-SupportedProperties.md#unity-text)** | No | Fully animatable | Outline width and color of the text. |
| **[`-unity-text-outline-color`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Outline color of the text. |
| **[`-unity-text-outline-width`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Outline width of the text. |
| **[`-unity-text-overflow-position`](UIE-USS-SupportedProperties.md#unity-text)** | No | Discrete | The element’s text overflow position. |
| **[`visibility`](UIE-USS-SupportedProperties.md#appearance)** | Yes | Discrete | Specifies whether or not an element is visible. |
| **[`white-space`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Discrete | Word wrap over multiple lines if not enough space is available to draw the text of an element. |
| **[`width`](UIE-USS-SupportedProperties.md#box-model)** | No | Fully animatable | Fixed width of an element for the layout. |
| **[`word-spacing`](UIE-USS-SupportedProperties.md#unity-text)** | Yes | Fully animatable | Increases or decreases the space between words. |