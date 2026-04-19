# Set a layerMask

This page explains how to set up a layerMask correctly so you can use it in API calls that use a serialized layerMask property.

## Use a serialized layerMask property

The simplest way to set a layermask in the Unity Editor is to create a property that uses Unity’s [LayerMask](../ScriptReference/LayerMask.md) class. If the property is `public` or uses the [SerializeField](../ScriptReference/SerializeField.md) attribute, Unity provides an interface in the **Inspector** that you can use to select which layers the layermask represents.

```
using UnityEngine;

public class LayerMaskExample : MonoBehaviour
{
    [SerializeField] private LayerMask layermask;
}
```

## Convert from a layer

If you want to convert a layer to a layermask in a script at runtime, use the [binary left-shift operator](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/bitwise-and-shift-operators#left-shift-operator-) to left-shift `1` by the layer. The result is a layermask that represents the single layer.

```
using UnityEngine;

public class LayerExample : MonoBehaviour
{
    [SerializeField] private int layer = 10;
    private int layerAsLayerMask;

    private void Start()
    {
        layerAsLayerMask = (1 << layer);
    }
}
```

## Additional resources

* [Add a layer to a layermask](layermask-add.md)
* [Remove a layer from a layermask](layermask-remove.md)