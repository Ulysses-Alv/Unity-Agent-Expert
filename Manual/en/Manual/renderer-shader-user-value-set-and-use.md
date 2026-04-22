# Set and use the RSUV

Learn how to set and use the [Renderer Shader User Value (RSUV)](renderer-shader-user-value-intro.md) through an example.

## Set the value

To set the value, you have to write a script.

The following example sets a random color to all MeshRenderer of the **scene**, encoded as a 24-bit integer:

```
        MeshRenderer[] allMeshRenderers = FindObjectsByType<MeshRenderer>(FindObjectsSortMode.InstanceID);
        foreach (MeshRenderer meshRenderer in allMeshRenderers)
        {
            // compute a random color
            Color32 c = Color.HSVToRGB(UnityEngine.Random.Range(0f, 1f), UnityEngine.Random.Range(0.6f, 1.0f), UnityEngine.Random.Range(0.6f, 1.0f));

            // set it as a LDR color in the RSUV value of the renderer
            uint cc = ((uint)c.b << 16) | ((uint)c.g << 8) | ((uint)c.r << 0);
            meshRenderer.SetShaderUserValue(cc);
        }
```

## Use the value

Once you’ve set the value, you can use it in **shader** code or in a shader graph. You can also use it with [Entities Graphics](https://docs.unity3d.com/Packages/com.unity.entities.graphics@latest).

### Use the value in shader code

In shader code, you can decode the 24-bit integer as a `float4` color. For example:

```
uint c = unity_RendererUserValue;
float4 mycolor = float4((float)((c >> 0) & 255) * (1.f / 255.f),
                        (float)((c >> 8) & 255) * (1.f / 255.f),
                        (float)((c >> 16) & 255) * (1.f / 255.f),
                        (float)1.f);
```

### Use the value in a shader graph

To retrieve the RSUV in a [shader graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest), use a [**Custom Function** node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Custom-Function-Node.html). For example:

1. In the Shader Graph workspace, add a Custom Function node and select it.
2. In the Graph **Inspector**, in the Custom Function node properties:
3. Add an **Output** of type **Vector 4** and call it **mycolor**.
4. Set the **Type** to **String**.
5. Set the **Name** to **getMyColor**.
6. In the **Body** field, paste the code provided in the previous section.
7. In the Shader Graph workspace, connect the Custom Function node’s **mycolor** output to the Fragment’s **Base Color** input.

### Use RSUV with Entities Graphics

You can also use the RSUV feature along with Unity’s [Entities Graphics package](https://docs.unity3d.com/Packages/com.unity.entities.graphics@latest).

The following code example assigns a distinct RSUV value to all entities that include a `MaterialMeshInfo` component.

```
using Unity.Entities;
using UnityEngine;
using Unity.Rendering;
using Unity.Burst;

[MaterialProperty("unity_RendererUserValuesPropertyEntry")]
public struct RendererUserValue_BaseValue : IComponentData
{
    public uint Value;
}

[BurstCompile]
public partial struct MaterialMeshInfoDebugSystem : ISystem
{

    private bool m_done;

    public void OnCreate(ref SystemState state)
    {
        m_done = false;
    }

    public void OnUpdate(ref SystemState state)
    {
        if ( !m_done )
        {
            var entityManager = state.EntityManager;
            var mQuery = entityManager.CreateEntityQuery(ComponentType.ReadOnly<MaterialMeshInfo>());
            // subscene may not be fully loaded, so wait for some entities to be here
            if (!mQuery.IsEmpty)
            {
                using (var entities = mQuery.ToEntityArray(Unity.Collections.Allocator.Temp))
                {
                    foreach (var entity in entities)
                    {
                        Color32 c = Color.HSVToRGB(UnityEngine.Random.Range(0f, 1f), UnityEngine.Random.Range(0.6f, 1.0f), UnityEngine.Random.Range(0.6f, 1.0f));
                        uint cc = ((uint)255<<24) | ((uint)c.b << 16) | ((uint)c.g << 8) | ((uint)c.r << 0);

                        entityManager.AddComponentData(entity, new RendererUserValue_BaseValue { Value = cc });
                    }
                }
                m_done = true;
            }
        }
    }
}
```

## Additional resources

* [Introduction to RSUV](renderer-shader-user-value-intro.md)
* [RSUV sample](renderer-shader-user-value-sample.md)