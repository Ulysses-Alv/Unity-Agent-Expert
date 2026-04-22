# Add a layer to a layerMask

To add a layer to a layermask, use the [logical OR operator](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/bitwise-and-shift-operators#logical-or-operator-) on the original layermask and the layer to add.

```
originalLayerMask |= (1 << layerToAdd);
```

## Additional resources

* [Set a layermask](layermask-set.md)
* [Remove a layer from a layermask](layermask-remove.md)