# Static GameObjects

If a **GameObject** does not move at runtime, it is known as a **static GameObject**. If a GameObject moves at runtime, it is known as a **dynamic GameObject**.

Many systems in Unity can precompute information about static GameObjects in the Editor. Because the GameObjects do not move, the results of these calculations are still valid at runtime. This means that Unity can save on runtime calculations, and potentially improve performance.

## The Static Editor Flags property

![The Static Editor Flags checkbox and drop-down menu, as seen when viewing a GameObject in the Inspector](../uploads/Main/GameObjectStaticDropDownMenu1.png)

The Static Editor Flags checkbox and drop-down menu, as seen when viewing a GameObject in the Inspector

The **Static Editor Flags** property lists a number of Unity systems which can include a static GameObject in their precomputations. Use the drop-down to define which systems include the GameObject in their precomputations. Setting Static Editor Flags at runtime has no effect on these systems.

Only include a GameObject in the precomputations for systems that need to know about that GameObject. Including a GameObject in the precomputations for a system that does not need to know about that GameObject can result in wasted calculations, unnecessarily large data files, or unexpected behavior.

The **Static Editor Flags** property is located in the **Inspector** for a GameObject, in the extreme top-right. It includes a checkbox, which sets the value to **Everything** or **Nothing**, and a dropdown menu that lets you choose which values to include.

To set the Static Editor Flags property in code, use the [GameObjectUtility.SetStaticEditorFlags](../ScriptReference/GameObjectUtility.SetStaticEditorFlags.md) API and the [GameObject.isStatic](../ScriptReference/GameObject-isStatic.md). To get the status of the Static Editor Flags property in code, use the [GameObjectUtility.GetStaticEditorFlags](../ScriptReference/GameObjectUtility.GetStaticEditorFlags.md) API.

The following values are available:

| **Property** | **Description** |
| --- | --- |
| **Nothing** | Do not include the GameObject in precomputations for any systems. |
| **Everything** | Include the GameObject in precomputations for all systems below. |
| **Contribute GI** | When you enable this property, Unity includes the target [Mesh Renderer](class-MeshRenderer.md) in **global illumination** calculations. These calculations take place while precomputing lighting data at bake time. The ContributeGI property exposes the [ReceiveGI](../ScriptReference/ReceiveGI.md) property. The ContributeGI property only takes effect if you enable a global illumination setting such as [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or [Enlighten Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for the target **Scene**. For additional context, refer to [this tutorial for setting up the Built-in Render Pipeline and lighting](choose-a-lighting-setup.md) in Unity. |
| **Occluder Static** | Mark the GameObject as a Static Occluder in the occlusion culling system. For more information, refer to [the Occlusion Culling system](OcclusionCulling.md). |
| **Occludee Static** | Mark the GameObject as a Static Occludee in the occlusion culling system. For more information, refer to [the Occlusion Culling system](OcclusionCulling.md). |
| **Batching Static** | Combine the GameObject’s **Mesh** with other eligible Meshes, to potentially reduce runtime rendering costs. For more information, refer to the documentation on [Static Batching](DrawCallBatching.md). |
| **Navigation Static** | Include the GameObject when precomputing navigation data. This workflow is deprecated and you cannot set **Navigation Static** here. However, you can still set this option in code with [`StaticEditorFlags.NavigationStatic`](../ScriptReference/StaticEditorFlags.NavigationStatic.md).  Instead of **Navigation Static** flags, use the [**NavMesh Modifier**](https://docs.unity3d.com/Packages/com.unity.ai.navigation@latest/index.html?subfolder=/manual/NavMeshModifier.html) component together with [**NavMesh Surfaces**](https://docs.unity3d.com/Packages/com.unity.ai.navigation@latest/index.html?subfolder=/manual/NavMeshSurface.html). |
| **Off Mesh Link Generation** | Attempt to generate an OffMesh Link that starts from this GameObject when precomputing navigation data. This workflow is deprecated and you cannot set **Off Mesh Link Generation** from this menu. However, you can still set this option in code with [`StaticEditorFlags.OffMeshLinkGeneration`](../ScriptReference/StaticEditorFlags.OffMeshLinkGeneration.md).  Instead of **Off Mesh Link Generation** flags, use the [**NavMesh Modifier**](https://docs.unity3d.com/Packages/com.unity.ai.navigation@latest/index.html?subfolder=/manual/NavMeshModifier.html) component together with [**NavMesh Surfaces**](https://docs.unity3d.com/Packages/com.unity.ai.navigation@latest/index.html?subfolder=/manual/NavMeshSurface.html). |
| **Reflection Probe** | Include this GameObject when precomputing data for **Reflection Probes** whose **Type** property is set to **Baked**. For more information, refer to [Reflection Probes](ReflectionProbes.md). |

## Additional resources

* [GameObjectUtility.SetStaticEditorFlags](../ScriptReference/GameObjectUtility.SetStaticEditorFlags.md)
* [GameObjectUtility.GetStaticEditorFlags](../ScriptReference/GameObjectUtility.GetStaticEditorFlags.md)
* [StaticEditorFlags.NavigationStatic](../ScriptReference/StaticEditorFlags.NavigationStatic.md)
* [StaticEditorFlags.OffMeshLinkGeneration](../ScriptReference/StaticEditorFlags.OffMeshLinkGeneration.md)
* [StaticEditorFlags](../ScriptReference/StaticEditorFlags.md)
* [Draw call batching](DrawCallBatching.md)
* [Occlusion Culling](OcclusionCulling.md)
* [NavMesh Modifier component](https://docs.unity3d.com/Packages/com.unity.ai.navigation@2.0/manual/NavMeshModifier.html)
* [NavMesh Surface component](https://docs.unity3d.com/Packages/com.unity.ai.navigation@2.0/manual/NavMeshSurface.html)
* [Reflection Probes](ReflectionProbes.md)