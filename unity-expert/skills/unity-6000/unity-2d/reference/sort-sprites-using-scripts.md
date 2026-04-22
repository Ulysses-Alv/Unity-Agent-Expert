# Sort sprites using scripts

You can sort **sprites** per **camera** through **scripts**, by modifying the following properties in Camera:

* [TransparencySortMode](../../../ScriptReference/Camera-transparencySortMode.md) (corresponds with **Transparency Sort Mode**)
* [TransparencySortAxis](../../../ScriptReference/Camera-transparencySortAxis.md) (corresponds with **Transparency Sort Axis**)

For example:

```
var camera = GetComponent<Camera>();

camera.transparencySortMode = TransparencySortMode.CustomAxis;

camera.transparencySortAxis = new Vector3(0.0f, 1.0f, 0.0f);
```