# Example of changing per-instance data at runtime in the Built-In Render Pipeline

The following example demonstrates how to use MaterialPropertyBlock objects to set per-instance data for a group of **GameObjects** at runtime. It sets the `_Color` property from the above **shader** examples to a random color.

**Important**: MaterialPropertyBlocks break SRP Batcher compatibility. For more information, see [GPU instancing: Requirements and Compatibility](GPUInstancing.md#requirements-and-compatibility).

```
using UnityEngine;

public class MaterialPropertyBlockExample : MonoBehaviour
{
    public GameObject[] objects;

    void Start()
    {
        MaterialPropertyBlock props = new MaterialPropertyBlock();
        MeshRenderer renderer;

        foreach (GameObject obj in objects)
        {
            float r = Random.Range(0.0f, 1.0f);
            float g = Random.Range(0.0f, 1.0f);
            float b = Random.Range(0.0f, 1.0f);
            props.SetColor("_Color", new Color(r, g, b));

            renderer = obj.GetComponent<MeshRenderer>();
            renderer.SetPropertyBlock(props);
        }
    }
}
```