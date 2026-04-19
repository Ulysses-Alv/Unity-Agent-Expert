# Create USS variables

Variables make it easier to manage styles for complex UI, where multiple rules, sometimes in different style sheets, use the same values.

You can create a USS variable and use it in other USS properties. When you update a USS variable, all of the USS properties that use that variable update. You can also specify default values for USS variables.

## Create and assign USS variables

You can create and assign a USS variable in [UI Builder](UIB-styling-ui-using-uss-variables.md) or directly in a USS file.

To create a USS variable in a USS file, prefix its name with a double-hyphen (`--`).

```
--color-1: red;
```

To use a USS variable value in another USS rule, use the `var()` function to call it.

```
var(--color-1);
```

When you update a variable, it updates all the USS properties that use it.

For example, the following USS example defines one style rule that declares two color variables, and two style rules that use those variables. To update the color scheme, you can change the two variable values instead of changing the four color values.

```
:root {
  --color-1: blue;
  --color-2: yellow;
}

.paragraph-regular {
  color: var(--color-1);
  background: var(--color-2);
  padding: 2px;
  }

.paragraph-reverse {
  color: var(--color-2);
  background: var(--color-1);
  padding: 2px;
  }
```

## Specify default values for USS variables

The `var()` function accepts an optional default value. The UI system uses the default value when it can’t resolve the variable. For example, the UI system uses the default value if you remove a variable from a style sheet but forget to remove a reference to it.

To specify a default value for a variable, add it after the variable value, separated by a comma `,`.

The following USS snippet calls the `--color-1` variable. If the UI system can’t resolve the variable, it uses the hex value for red (`#FF0000`).

```
var(--color-1, #FF0000);
```

## Differences from CSS variables

Variables work mostly the same way in USS as they do in CSS. For detailed information about CSS variables, refer to the [MDN documentation](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties). However, USS doesn’t support some CSS functionality:

* USS doesn’t support the `var()` function inside of other functions, such as the following:

```
  background-color: rgb(var(--red), 0, 0);
```

* USS doesn’t support mathematical operations on variables.

## Additional resources

* [Introduction to USS built-in variables](UIE-USS-UnityVariables.md)
* [USS built-in variable reference](UIE-uss-built-in-variable-reference.md)