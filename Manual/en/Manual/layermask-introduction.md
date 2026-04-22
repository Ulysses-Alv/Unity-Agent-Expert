# Introduction to layerMasks

Every **GameObject** exists on a single layer, but Unity APIs that let you set which layers the API affect don’t directly use layers. Instead, they use layerMasks.

A [layer](../ScriptReference/GameObject-layer.md) is a standard integer, but a layerMask is an integer formatted as a bitmask where every `1` represents a layer to include and every `0` represents a layer to exclude. This means that you can pass a layer to an API that expects a layerMasks and the script will still compile because layers and layerMasks use the same underlying type. However, the API call won’t produce the behavior you expect.

For example, if you want to perform a [RayCast](../ScriptReference/Physics.Raycast.md) against GameObjects on layer 9, if you pass `9` into the Physics.Raycast call as the layerMask, Unity actually performs the ray cast against GameObjects on layers `3` and `0`. This is because the binary representation of 9 is `00001001` and if you interpret this as a mask, the `1`s are in the place of layers `3` and `0`.