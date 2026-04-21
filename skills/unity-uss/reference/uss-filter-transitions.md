# USS filter transitions

USS filter transitions follow the CSS filter transition syntax, enabling smooth visual effects based on filter properties.

## Matching filter functions

For a smooth transition, the filter functions in both states must match exactly, including the order and type of functions.

The following example creates a smooth transition between two states with matching filter functions. When you hover over the element, the blur and invert functions transition smoothly.

```
.effect {
    filter: blur(0px) invert(0%);
}

.effect:hover {
    filter: blur(10px) invert(100%);
}
```

## Partially matching filter functions

When filter functions partially match, the system interpolates common functions and applies new ones. For example, the following example transitions between a blur and invert function.

```
.effect {
    filter: blur(0px); /* invert(0%) */
}

.effect:hover {
    filter: blur(10px) invert(100%);
}
```

In this example, the system pads the initial filter with an invert function, using a default interpolation value (`0%`). When you create a custom filter, you can specify the interpolation value in the asset definition.

## Mismatched filter functions

If the filter functions don’t match in type or order, such as shown in the following example, transitioning from `invert()` and `blur()` to a `blur()` and `invert()`, interpolation doesn’t happen. Instead, the transition abruptly changes to the new filter state.

```
.effect {
    filter: invert(0%) blur(0px);
}

.effect:hover {
    filter: blur(10px) invert(100%);
}
```

## Additional resources

* [Create a custom swirl filter](create-custom-swirl-filter.md)
* [Set a project’s color space](../set-project-color-space.md)