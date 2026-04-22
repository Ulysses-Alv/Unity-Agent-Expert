# USS transform

The transform properties apply a 2D transformation to a **visual element**. You can use them to rotate, scale, or move a visual element.

If you change the layout of an element, Unity recalculates the layout of other elements in the same hierarchy. This recalculation might reduce an animation’s frame rate. Applying transform to an element reduces recalculations because it doesn’t change the layout of other elements in the hierarchy.

It’s possible to use transform to define the static appearance of a visual element. However, transform is best used for changes and animations. For example, if you want to make a visual element shake when an event happens in an application, set the position of the visual element using the regular layout properties such as `top` and `left`, and then use `translate` to align an offset relative to the initial position.

Transform includes the following properties:

| **Property** | **USS syntax** | **Description** |
| --- | --- | --- |
| **Transform Origin** | `transform-origin` | Represents the point of origin where rotation, scaling, and translation occur. |
| **Translate** | `translate` | Repositions the visual element in horizontal or vertical directions. |
| **Scale** | `scale` | Changes the apparent size, padding, border, and margins of a visual element. Negative values flip visual elements along the scale axis. |
| **Rotate** | `rotate` | Rotates a visual element. Positive values represent clockwise rotation and negative values represent counterclockwise rotation. You can set rotation with `deg`, `grad`, `rad`, or `turn` units. For more information on these units, see [MDN Web Docs’s page on the `<angle>` CSS data type](https://developer.mozilla.org/en-US/docs/Web/CSS/angle). |

**Note**: All transformations are performed in the following order:

1. Scale
2. Rotate
3. Translate

You can set transform properties for a visual element using the controls in the [UI Builder](UIBuilder.md), within a [USS](UIE-USS.md) file, or using a C# script.

## Transform controls in the UI Builder

You can use the controls in the **Transform** section of the ****Inspector**** in the UI Builder to set the transform properties of a visual element.

### Pivot Origin

The **Pivot Origin** widget sets the transform origin property. To use it, do one of the following:

* Click a point in the widget to set the origin to a corner, the center of an edge, or the center. You can also define the values using the keyboard. When the widget is in focus, use the arrow keys to specify a point in the widget.
* Enter values for **X** and **Y** and specify the unit.

**Tip**: You can enter `%` or `px` after values. This automatically changes the displayed unit in the unit selector. You can also drag to define the values in the **X** and **Y** boxes.

**Note**: The default value for the transform origin is center.

If you use percentages for both the X and Y values, the widget shows the new origin point when you edit the X and Y text boxes.

If you specify a transform origin point outside the element, such as having a value less than 0% or greater than 100%, the widget shows the directions of the X and Y axes.

![A transform origin with an X value less than 0, and a Y value greater than 100. The pivot origin is highlighted in the bottom left corner.](../uploads/Main/TransformOriginWgt.png)

A transform origin with an X value less than 0, and a Y value greater than 100. The pivot origin is highlighted in the bottom left corner.

### Translate

The **Translate** control sets the translate property. To use it, enter values in the **X** and **Y** boxes and specify the unit.

**Tip**: You can enter `%` or `px` after values. This automatically changes the displayed unit in the unit selector. You can also drag to define the values in the **X** and **Y** boxes.

### Scale

The **Scale** control sets the scale property. To use it, enter values in the **X** and **Y** boxes and specify the unit.

**Tip**: You can enter `%` or `px` after values. This automatically changes the displayed unit in the unit selector. You can also drag to define the values in the **X** and **Y** boxes.

### Rotate

The **Rotate** control sets the rotate property. To use it, set the **Axis** values in the **X**, **Y**, and **Z** boxes, and set the angle value in the **Angle** box.

## USS transform properties

You can use styling rules to set the transform properties for a visual element. You can set the styling rules within a USS file or inline in a UXML file.

### `transform-origin`

The `transform-origin` property sets the transform origin along the X and Y axes in **pixels** or percentages.

You can also use keywords to set the value of the `transform-origin` attribute. These keywords match the dots in the widget in the UI Builder. The following keywords are supported:

| **Pivot point** | **Keywords** |
| --- | --- |
| **Center** | * `center` * `center center` (This is the default value) |
| **Center of left edge** | * `left` * `left center` * `center left` |
| **Center of right edge** | * `right` * `right center` * `center right` |
| **Center of top edge** | * `top` * `top center` * `center top` |
| **Center of bottom edge** | * `bottom` * `bottom center` * `center bottom` |
| **Top-left corner** | * `top left` * `left top` |
| **Top-right corner** | * `top right` * `right top` |
| **Bottom-left corner** | * `bottom left` * `left bottom` |
| **Bottom-right corner** | * `bottom right` * `right bottom` |

**Examples**

```
transform-origin: 0% 100%;
transform-origin: 20px 10px;
transform-origin: 0px 100%;
transform-origin: 60% 10px;
```

### `translate`

The `translate` property sets the translation along the X and Y axes in pixels or percentages relative to the size of this visual element. You can omit Y if it equals X.

**Examples**

```
translate: 80%;
translate: 35px;
translate: 5% 10px;
translate: 24px 0%;
```

### `scale`

The `scale` property sets the scale along the X and Y axes in pixels or percentages. You can omit Y if it equals X.

The keyword `none` sets no scale.

**Examples**

```
scale: 2.5;
scale: -1 1;
scale: none;
```

### `rotate`

The `rotate` property sets the rotation using a number or a unit.

The keyword `none` sets no rotation.

**Examples**

```
rotate: 45deg;
rotate: -100grad;
rotate: -3.14rad;
rotate: 0.75turn;
rotate: none;
```

## Transform C# properties

You can set the transform properties for a visual element in a C# script.

### `IStyle.transformOrigin`

The [`IStyle.transformOrigin`](../ScriptReference/UIElements.IStyle-transformOrigin.md) property sets the transform origin.

The `transformOrigin` property of the [`style`](../ScriptReference/UIElements.VisualElement-style.md) is of type [`StyleTransformOrigin`](../ScriptReference/UIElements.StyleTransformOrigin.md). Its [constructor](../ScriptReference/UIElements.StyleTransformOrigin-ctor.md) takes a [`TransformOrigin`](../ScriptReference/UIElements.TransformOrigin.md) as an argument. You can construct a new [`TransformOrigin`](../ScriptReference/UIElements.TransformOrigin.md) using an X value and a Y value. The X value and the Y value are of type [`Length`](../ScriptReference/UIElements.Length.md).

**Examples**

```
//This example sets the transform origin of the element to be 100 pixels from the left edge and 50% of the way down from the top edge.
element.style.transformOrigin = new StyleTransformOrigin(new TransformOrigin(new Length(100f, LengthUnit.Pixel), new Length(50f, LengthUnit.Percent)));
```

You can simplify the above code as follows using implicit conversions:

```
element.style.transformOrigin = new TransformOrigin(100, Length.Percent(50));
```

### `IStyle.translate`

The [`IStyle.translate`](../ScriptReference/UIElements.IStyle-translate.md) property sets the translation.

[`IStyle.translate`](../ScriptReference/UIElements.IStyle-translate.md) is of type [`StyleTranslate`](../ScriptReference/UIElements.StyleTranslate.md). Its constructor takes a [`Translate`](../ScriptReference/UIElements.Translate.md) as an argument. You can construct a new [`Translate`](../ScriptReference/UIElements.Translate.md) using an X value and a Y value. The X value and the Y value are of type [`Length`](../ScriptReference/UIElements.Length.md).

**Examples**

```
//This example sets the translation of the element. The X-axis is 10% and the Y-axis is 50 pixels.
element.style.translate = new StyleTranslate(new Translate(new Length(10f, LengthUnit.Percent), new Length(50f, LengthUnit.Pixel)));
```

You can simplify the above code as follows using implicit conversions:

```
element.style.translate = new Translate(Length.Percent(10), 50);
```

### `IStyle.scale`

The [`IStyle.scale`](../ScriptReference/UIElements.IStyle-scale.md) property sets the scale.

[`IStyle.scale`](../ScriptReference/UIElements.IStyle-scale.md) is of type [`StyleScale`](../ScriptReference/UIElements.StyleScale.md). [`StyleScale`](../ScriptReference/UIElements.StyleScale.md)’s [constructor](../ScriptReference/UIElements.StyleScale-ctor.md) takes a [Scale](../ScriptReference/UIElements.Scale.md) as an argument. You can [construct](../ScriptReference/UIElements.Scale-ctor.md) this `Scale` with a `Vector2`.

**Examples**

```
element.style.scale = new StyleScale(new Scale(new Vector2(0.5f, -1f)));
```

You can simplify the code above as follows using implicit conversions:

```
element.style.scale = new Scale(new Vector2(0.5f, -1));
```

### `IStyle.rotate`

The [`IStyle.rotate`](../ScriptReference/UIElements.IStyle-rotate.md) property sets the rotation.

The [`IStyle.rotate`](../ScriptReference/UIElements.IStyle-rotate.md) property is of type [`StyleRotate`](../ScriptReference/UIElements.StyleRotate.md). The [`StyleRotate`](../ScriptReference/UIElements.StyleRotate.md)’s [constructor](../ScriptReference/UIElements.StyleRotate-ctor.md) takes a [`Rotate`](../ScriptReference/UIElements.Rotate.md) as an argument. You can [construct](../ScriptReference/UIElements.Rotate-ctor.md) this `Rotate` with an [`Angle`](../ScriptReference/UIElements.Angle.md). You can [construct](../ScriptReference/UIElements.Angle-ctor.md) an `Angle` with a float and an optional [`AngleUnit`](../ScriptReference/UIElements.AngleUnit.md) enum, or you can use static methods [`Degrees()`](../ScriptReference/UIElements.Angle.Degrees.md), [`Gradians()`](../ScriptReference/UIElements.Angle.Gradians.md), [`Radians()`](../ScriptReference/UIElements.Angle.Radians.md), and [`Turns()`](../ScriptReference/UIElements.Angle.Turns.md).

**Examples**

```
//Rotate by 180 degrees
elements[0].style.rotate = new StyleRotate(new Rotate(new Angle(180f, AngleUnit.Degree)));
//Rotate by 200 gradians
elements[1].style.rotate = new StyleRotate(new Rotate(new Angle(200f, AngleUnit.Gradian)));
//Rotate by pi radians
elements[2].style.rotate = new StyleRotate(new Rotate(new Angle(Mathf.PI, AngleUnit.Radian)));
//Rotate by half a turn
elements[3].style.rotate = new StyleRotate(new Rotate(new Angle(0.5f, AngleUnit.Turn)));
```

You can simplify the above code as follows:

```
//Rotate by 180 degrees
elements[0].style.rotate = new Rotate(180);
//Rotate by 200 gradians
elements[1].style.rotate = new Rotate(Angle.Gradians(200));
//Rotate by pi radians
elements[2].style.rotate = new Rotate(Angle.Radians(Mathf.PI));
//Rotate by half a turn
elements[3].style.rotate = new Rotate(Angle.Turns(0.5f));
```

## Additional resources

* [Layout engine](UIE-LayoutEngine.md)
* [CSS transform-origin property](https://developer.mozilla.org/en-US/docs/Web/CSS/transform-origin) (Mozilla developer documentation)
* [CSS translate property](https://developer.mozilla.org/en-US/docs/Web/CSS/translate) (Mozilla developer documentation)
* [CSS rotate property](https://developer.mozilla.org/en-US/docs/Web/CSS/rotate) (Mozilla developer documentation)
* [CSS scale property](https://developer.mozilla.org/en-US/docs/Web/CSS/scale) (Mozilla developer documentation)