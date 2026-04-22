# Parallel tessellation

To improve performance when you tessellate a large **mesh**, use the job system. This prevents the main thread from being blocked for an extended period. When you create a tessellation job, use the [MeshGenerationContext.AddMeshGenerationJob](../ScriptReference/UIElements.MeshGenerationContext.AddMeshGenerationJob.md) API to provide the job handle to UI Toolkit. This ensures that UI Toolkit will wait for the job to finish before accessing the mesh.

## Jobified tessellation with allocation on the main thread

If you can quickly calculate the number of vertices or indices, use the main thread to allocate the mesh and add it to the draw sequence. You can do so with the following APIs:

* [MeshGenerationContext.AllocateTempMesh](../ScriptReference/UIElements.MeshGenerationContext.AllocateTempMesh.md)
* [MeshGenerationContext.DrawMesh](../ScriptReference/UIElements.MeshGenerationContext.DrawMesh.md)

**Note**: You can add the mesh to the draw sequence at any point during the call to [generateVisualContent](../ScriptReference/UIElements.VisualElement-generateVisualContent.md).

The following example demonstrates a jobified tessellation with allocation on the main thread, where the job performs the actual tessellation.

```
using System;
using Unity.Collections;
using Unity.Jobs;
using UnityEngine;
using UnityEngine.UIElements;

class CheckerboardElement : VisualElement
{
    struct CheckerboardJob : IJob
    {
        [WriteOnly] public NativeSlice<Vertex> vertices;
        [WriteOnly] public NativeSlice<ushort> indices;

        public int horizontalCount;
        public int verticalCount;
        public float divisions;
        public Color32 color1;
        public Color32 color2;

        public void Execute()
        {
            Span<Color32> colors = stackalloc Color32[2];
            colors[0] = color1;
            colors[1] = color2;
            int colorIndex = 0;

            int vCount = 0;
            int iCount = 0;

            float left = 0;
            float right = divisions;
            for (int i = 0; i < horizontalCount; ++i)
            {
                float top = 0;
                float bottom = divisions;
                for (int j = 0; j < verticalCount; ++j)
                {
                    colorIndex = (i ^ j) & 1;

                    Color32 color = colors[colorIndex];

                    vertices[vCount + 0] = new Vertex { position = new Vector3(left, bottom), tint = color };
                    vertices[vCount + 1] = new Vertex { position = new Vector3(left, top), tint = color };
                    vertices[vCount + 2] = new Vertex { position = new Vector3(right, top), tint = color };
                    vertices[vCount + 3] = new Vertex { position = new Vector3(right, bottom), tint = color };

                    indices[iCount + 0] = (ushort)(vCount + 0);
                    indices[iCount + 1] = (ushort)(vCount + 1);
                    indices[iCount + 2] = (ushort)(vCount + 2);
                    indices[iCount + 3] = (ushort)(vCount + 2);
                    indices[iCount + 4] = (ushort)(vCount + 3);
                    indices[iCount + 5] = (ushort)(vCount + 0);

                    vCount += 4;
                    iCount += 6;
                    top += divisions;
                    bottom += divisions;
                }

                left += divisions;
                right += divisions;
            }
        }
    }

    int m_Divisions;
    Color32 m_Color1;
    Color32 m_Color2;

    public CheckerboardElement(int divisions, Color32 color1, Color32 color2)
    {
        generateVisualContent = OnGenerateVisualContent;
        m_Divisions = divisions;
        m_Color1 = color1;
        m_Color2 = color2;
    }

    void OnGenerateVisualContent(MeshGenerationContext mgc)
    {
        int horizontalCount = Mathf.FloorToInt(layout.width / m_Divisions);
        int verticalCount = Mathf.FloorToInt(layout.height / m_Divisions);

        if (horizontalCount == 0 || verticalCount == 0)
            return;

        var job = new CheckerboardJob
        {
            horizontalCount = horizontalCount,
            verticalCount = verticalCount,
            divisions = m_Divisions,
            color1 = m_Color1,
            color2 = m_Color2
        };

        int quads = horizontalCount * verticalCount;

        mgc.AllocateTempMesh(quads * 4, quads * 6, out job.vertices, out job.indices);
        mgc.DrawMesh(job.vertices, job.indices);

        JobHandle jobHandle = job.Schedule();
        mgc.AddMeshGenerationJob(jobHandle);
    }
}
```

## Jobified tessellation with allocation in the job

If it’s computationally expensive to calculate the number of vertices or indices, insert a [MeshGenerationNode](../ScriptReference/UIElements.MeshGenerationNode.md) into the draw sequence. Then provide the node to your job, along with a [TempMeshAllocator](../ScriptReference/UIElements.TempMeshAllocator.md). You can do so with the following APIs:

* [MeshGenerationContext.InsertMeshGenerationNode](../ScriptReference/UIElements.MeshGenerationContext.InsertMeshGenerationNode.md)
* [MeshGenerationContext.GetTempMeshAllocator](../ScriptReference/UIElements.MeshGenerationContext.GetTempMeshAllocator.md)

In this approach, perform the allocation from the job. Use the [TempMeshAllocator.AllocateTempMesh](../ScriptReference/UIElements.TempMeshAllocator.AllocateTempMesh.md) to safely allocate a temporary mesh from a job. The [MeshGenerationNode.DrawMesh](../ScriptReference/UIElements.MeshGenerationNode.DrawMesh.md) acts like a container for a nested draw sequence that’s populated by the job.

Building on the prior example, the following demonstrates how to perform both the allocation and the tessellation from the job:

```
using System;
using Unity.Collections;
using Unity.Jobs;
using UnityEngine;
using UnityEngine.UIElements;

class CheckerboardElement_AllocFromJob : VisualElement
{
    struct CheckerboardJob : IJob
    {
        [ReadOnly]
        public TempMeshAllocator allocator;
        public MeshGenerationNode node;

        public int horizontalCount;
        public int verticalCount;
        public float divisions;
        public Color32 color1;
        public Color32 color2;

        public void Execute()
        {
            // Let's pretend that the number of vertices/indices is difficult to compute...
            int quads = horizontalCount * verticalCount;
            int totalVertexCount = quads * 4;
            int totalIndexCount = quads * 6;

            // Allocate the mesh from the job
            allocator.AllocateTempMesh(totalVertexCount, totalIndexCount, out NativeSlice<Vertex> vertices, out NativeSlice<ushort> indices);

            Span<Color32> colors = stackalloc Color32[2];
            colors[0] = color1;
            colors[1] = color2;
            int colorIndex = 0;

            int vCount = 0;
            int iCount = 0;

            float left = 0;
            float right = divisions;
            for (int i = 0; i < horizontalCount; ++i)
            {
                float top = 0;
                float bottom = divisions;
                for (int j = 0; j < verticalCount; ++j)
                {
                    colorIndex = (i ^ j) & 1;

                    Color32 color = colors[colorIndex];

                    vertices[vCount + 0] = new Vertex { position = new Vector3(left, bottom), tint = color };
                    vertices[vCount + 1] = new Vertex { position = new Vector3(left, top), tint = color };
                    vertices[vCount + 2] = new Vertex { position = new Vector3(right, top), tint = color };
                    vertices[vCount + 3] = new Vertex { position = new Vector3(right, bottom), tint = color };

                    indices[iCount + 0] = (ushort)(vCount + 0);
                    indices[iCount + 1] = (ushort)(vCount + 1);
                    indices[iCount + 2] = (ushort)(vCount + 2);
                    indices[iCount + 3] = (ushort)(vCount + 2);
                    indices[iCount + 4] = (ushort)(vCount + 3);
                    indices[iCount + 5] = (ushort)(vCount + 0);

                    vCount += 4;
                    iCount += 6;
                    top += divisions;
                    bottom += divisions;
                }

                left += divisions;
                right += divisions;
            }

            // Add the mesh to the draw sequence of the node
            node.DrawMesh(vertices, indices);
        }
    }

    int m_Divisions;
    Color32 m_Color1;
    Color32 m_Color2;

    public CheckerboardElement_AllocFromJob(int divisions, Color32 color1, Color32 color2)
    {
        generateVisualContent = OnGenerateVisualContent;
        m_Divisions = divisions;
        m_Color1 = color1;
        m_Color2 = color2;
    }

    void OnGenerateVisualContent(MeshGenerationContext mgc)
    {
        int horizontalCount = Mathf.FloorToInt(layout.width / m_Divisions);
        int verticalCount = Mathf.FloorToInt(layout.height / m_Divisions);

        if (horizontalCount == 0 || verticalCount == 0)
            return;

        var job = new CheckerboardJob
        {
            horizontalCount = horizontalCount,
            verticalCount = verticalCount,
            divisions = m_Divisions,
            color1 = m_Color1,
            color2 = m_Color2
        };

        mgc.GetTempMeshAllocator(out job.allocator); // Provide the job with the allocator
        mgc.InsertMeshGenerationNode(out job.node); // Insert a node in the element draw sequence

        JobHandle jobHandle = job.Schedule();
        mgc.AddMeshGenerationJob(jobHandle);
    }
}
```

## Additional resources

* [UI Renderer](UIE-ui-renderer.md)