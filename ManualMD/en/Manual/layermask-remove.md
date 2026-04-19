# Remove a layer from a layerMask

To remove a layer from a layermask, use the [logical AND operator](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/bitwise-and-shift-operators#logical-and-operator-) on the original layermask and the [bitwise complement](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/bitwise-and-shift-operators#bitwise-complement-operator-) of the layer to remove it.

```
originalLayerMask &= ~(1 << layerToRemove);
```

## Additional resources

* [Set a layermask](layermask-set.md)
* [Add a layer to a layermask](layermask-add.md)