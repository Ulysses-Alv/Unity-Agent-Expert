# Use Mesh API to create a radial progress indicator

**Version**: 2023.2+

This example demonstrates how to use the **Mesh** API to draw visual content onto a **visual element**.

**Note**: The Mesh API is a tool for advanced users. In version 2022.1 and later, if you only want to generate simple geometry, use the Vector API instead. For more information, see [Use Vector API to create a radial progress indicator](UIE-radial-progress-use-vector-api.md).

## Example overview

This example creates a custom control that displays progress, as an alternative to a loading bar. The progress indicator displays a progress value in a partially filled ring around a label that displays the percentage. It supports a value between 0 and 100, which determines how much of the ring is filled.

![A radial control with progress highlighted in green, and the number 87% displayed in the center of the radial control.](../uploads/Main/radial-progress-finished.png)

A radial control with progress highlighted in green, and the number 87% displayed in the center of the radial control.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023.2/radial-progress).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [USS](UIE-USS.md)
* [Handling events](UIE-Events-Handling.md)
* [`MeshGenerationContext`](../ScriptReference/UIElements.MeshGenerationContext.md)

## Create the radial progress indicator and its custom mesh

Create a C# script to define the `RadialProgress` class, and expose the control to UXML and UI Builder.

1. Create a Unity project with any template.
2. Create a folder named `radial-progress` to store your files.
3. In the `radial-progress` folder, create a C# scrip named `RadialProgress.cs` with the following content:

```
using Unity.Collections;
using UnityEngine;
using UnityEngine.UIElements;

namespace MyUILibrary
{
    /// <summary>
    /// An element that displays progress inside a partially filled circle
    /// </summary>
    [UxmlElement]
    public partial class RadialProgress : VisualElement
    {
        // These are USS class names for the control overall and the label.
        public static readonly string ussClassName = "radial-progress";
        public static readonly string ussLabelClassName = "radial-progress__label";

        // These objects allow C# code to access custom USS properties.
        static CustomStyleProperty<Color> s_TrackColor = new CustomStyleProperty<Color>("--track-color");
        static CustomStyleProperty<Color> s_ProgressColor = new CustomStyleProperty<Color>("--progress-color");

        // These are the meshes this control uses.
        EllipseMesh m_TrackMesh;
        EllipseMesh m_ProgressMesh;

        // This is the label that displays the percentage.
        Label m_Label;

        // This is the number of outer vertices to generate the circle.
        const int k_NumSteps = 200;

        // This is the number that the Label displays as a percentage.
        float m_Progress;

        /// <summary>
        /// A value between 0 and 100
        /// </summary>
        [UxmlAttribute]
        public float progress
        {
            // The progress property is exposed in C#.
            get => m_Progress;
            set
            {
                // Whenever the progress property changes, MarkDirtyRepaint() is named. This causes a call to the
                // generateVisualContents callback.
                m_Progress = value;
                m_Label.text = Mathf.Clamp(Mathf.Round(value), 0, 100) + "%";
                MarkDirtyRepaint();
            }
        }

        // This default constructor is RadialProgress's only constructor.
        public RadialProgress()
        {
            // Create a Label, add a USS class name, and add it to this visual tree.
            m_Label = new Label();
            m_Label.AddToClassList(ussLabelClassName);
            Add(m_Label);

            // Create meshes for the track and the progress.
            m_ProgressMesh = new EllipseMesh(k_NumSteps);
            m_TrackMesh = new EllipseMesh(k_NumSteps);

            // Add the USS class name for the overall control.
            AddToClassList(ussClassName);

            // Register a callback after custom style resolution.
            RegisterCallback<CustomStyleResolvedEvent>(evt => CustomStylesResolved(evt));

            // Register a callback to generate the visual content of the control.
            generateVisualContent += context => GenerateVisualContent(context);

            progress = 0.0f;
        }

        static void CustomStylesResolved(CustomStyleResolvedEvent evt)
        {
            RadialProgress element = (RadialProgress)evt.currentTarget;
            element.UpdateCustomStyles();
        }

        // After the custom colors are resolved, this method uses them to color the meshes and (if necessary) repaint
        // the control.
        void UpdateCustomStyles()
        {
            if (customStyle.TryGetValue(s_ProgressColor, out var progressColor))
            {
                m_ProgressMesh.color = progressColor;
            }

            if (customStyle.TryGetValue(s_TrackColor, out var trackColor))
            {
                m_TrackMesh.color = trackColor;
            }

            if (m_ProgressMesh.isDirty || m_TrackMesh.isDirty)
                MarkDirtyRepaint();
        }

        // The GenerateVisualContent() callback method calls DrawMeshes().
        static void GenerateVisualContent(MeshGenerationContext context)
        {
            RadialProgress element = (RadialProgress)context.visualElement;
            element.DrawMeshes(context);
        }

        // DrawMeshes() uses the EllipseMesh utility class to generate an array of vertices and indices, for both the
        // "track" ring (in grey) and the progress ring (in green). It then passes the geometry to the MeshWriteData
        // object, as returned by the MeshGenerationContext.Allocate() method. For the "progress", only a slice of
        // the index arrays is used to progressively reveal parts of the mesh.
        void DrawMeshes(MeshGenerationContext context)
        {
            float halfWidth = contentRect.width * 0.5f;
            float halfHeight = contentRect.height * 0.5f;

            if (halfWidth < 2.0f || halfHeight < 2.0f)
                return;

            m_ProgressMesh.width = halfWidth;
            m_ProgressMesh.height = halfHeight;
            m_ProgressMesh.borderSize = 10;
            m_ProgressMesh.UpdateMesh();

            m_TrackMesh.width = halfWidth;
            m_TrackMesh.height = halfHeight;
            m_TrackMesh.borderSize = 10;
            m_TrackMesh.UpdateMesh();

            // Draw track mesh first
            var trackMeshWriteData = context.Allocate(m_TrackMesh.vertices.Length, m_TrackMesh.indices.Length);
            trackMeshWriteData.SetAllVertices(m_TrackMesh.vertices);
            trackMeshWriteData.SetAllIndices(m_TrackMesh.indices);

            // Keep progress between 0 and 100
            float clampedProgress = Mathf.Clamp(m_Progress, 0.0f, 100.0f);

            // Determine how many triangle are used to depending on progress, to achieve a partially filled circle
            int sliceSize = Mathf.FloorToInt((k_NumSteps * clampedProgress) / 100.0f);

            if (sliceSize == 0)
                return;

            // Every step is 6 indices in the corresponding array
            sliceSize *= 6;

            var progressMeshWriteData = context.Allocate(m_ProgressMesh.vertices.Length, sliceSize);
            progressMeshWriteData.SetAllVertices(m_ProgressMesh.vertices);

            var tempIndicesArray = new NativeArray<ushort>(m_ProgressMesh.indices, Allocator.Temp);
            progressMeshWriteData.SetAllIndices(tempIndicesArray.Slice(0, sliceSize));
            tempIndicesArray.Dispose();
        }

    }
}
```

## Create the custom mesh

Create a C# script named `EllipseMesh.cs` with the following content:

```
using UnityEngine;
using UnityEngine.UIElements;

namespace MyUILibrary
{
    public class EllipseMesh
    {
        int m_NumSteps;
        float m_Width;
        float m_Height;
        Color m_Color;
        float m_BorderSize;
        bool m_IsDirty;
        public Vertex[] vertices { get; private set; }
        public ushort[] indices { get; private set; }

        public EllipseMesh(int numSteps)
        {
            m_NumSteps = numSteps;
            m_IsDirty = true;
        }

        public void UpdateMesh()
        {
            if (!m_IsDirty)
                return;

            int numVertices = numSteps * 2;
            int numIndices = numVertices * 6;

            if (vertices == null || vertices.Length != numVertices)
                vertices = new Vertex[numVertices];

            if (indices == null || indices.Length != numIndices)
                indices = new ushort[numIndices];

            float stepSize = 360.0f / (float)numSteps;
            float angle = -180.0f;

            for (int i = 0; i < numSteps; ++i)
            {
                angle -= stepSize;
                float radians = Mathf.Deg2Rad * angle;

                float outerX = Mathf.Sin(radians) * width;
                float outerY = Mathf.Cos(radians) * height;
                Vertex outerVertex = new Vertex();
                outerVertex.position = new Vector3(width + outerX, height + outerY, Vertex.nearZ);
                outerVertex.tint = color;
                vertices[i * 2] = outerVertex;

                float innerX = Mathf.Sin(radians) * (width - borderSize);
                float innerY = Mathf.Cos(radians) * (height - borderSize);
                Vertex innerVertex = new Vertex();
                innerVertex.position = new Vector3(width + innerX, height + innerY, Vertex.nearZ);
                innerVertex.tint = color;
                vertices[i * 2 + 1] = innerVertex;

                indices[i * 6] = (ushort)((i == 0) ? vertices.Length - 2 : (i - 1) * 2); // previous outer vertex
                indices[i * 6 + 1] = (ushort)(i * 2); // current outer vertex
                indices[i * 6 + 2] = (ushort)(i * 2 + 1); // current inner vertex

                indices[i * 6 + 3] = (ushort)((i == 0) ? vertices.Length - 2 : (i - 1) * 2); // previous outer vertex
                indices[i * 6 + 4] = (ushort)(i * 2 + 1); // current inner vertex
                indices[i * 6 + 5] = (ushort)((i == 0) ? vertices.Length - 1 : (i - 1) * 2 + 1); // previous inner vertex
            }

            m_IsDirty = false;
        }

        public bool isDirty => m_IsDirty;

        void CompareAndWrite(ref float field, float newValue)
        {
            if (Mathf.Abs(field - newValue) > float.Epsilon)
            {
                m_IsDirty = true;
                field = newValue;
            }
        }

        public int numSteps
        {
            get => m_NumSteps;
            set
            {
                m_IsDirty = value != m_NumSteps;
                m_NumSteps = value;
            }
        }

        public float width
        {
            get => m_Width;
            set => CompareAndWrite(ref m_Width, value);
        }

        public float height
        {
            get => m_Height;
            set => CompareAndWrite(ref m_Height, value);
        }

        public Color color
        {
            get => m_Color;
            set
            {
                m_IsDirty = value != m_Color;
                m_Color = value;
            }
        }

        public float borderSize
        {
            get => m_BorderSize;
            set => CompareAndWrite(ref m_BorderSize, value);
        }

    }
}
```

## Style the custom control

Create a USS file named `RadialProgress.uss` with the following content:

```
.radial-progress {
        min-width: 26px;
        min-height: 20px;
        --track-color: rgb(130, 130, 130);
        --progress-color: rgb(46, 132, 24);
        --percentage-color: white;
        margin-left: 5px;
        margin-right: 5px;
        margin-top: 5px;
        margin-bottom: 5px;
        flex-direction: row;
        justify-content: center;
        width: 100px; 
        height: 100px;
    }

    .radial-progress__label {
        -unity-text-align: middle-left;
        color: var(--percentage-color);
    }
```

## Use the custom control in a UI Document and test

Use UI Builder to add the control and apply the USS stylesheet. Test the control with different `Progress` values.

1. Create a UI Document named `RadialProgressExample.uxml`.
2. Double-click `RadialProgressExample.uxml` to open it in the UI Builder.
3. In the Library window, select **Project** > **Custom Controls** > **UxmlSerializedData** > **MyUILibrary**.
4. Drag **RadialProgress** to the Hierarchy window.
5. In the **StyleSheets** section of the UI Builder, add `RadialProgress.uss` as the existing USS.
6. In the Hierarchy window, select **RadialProgress**.
7. In the **Inspector** window, enter `radial-progress` in the **Name** box.
8. In the Inspector window, enter different values in the **Progress** box. The percentage in the ****Viewport**** changes, and the green progress ring resizes.

## Create logic to update the progress with dynamic values

Create a C# MonoBehaviour script to update the `Progress` property of the control with dynamic values for demo purposes. In the `radial-progress` folder, create a C# MonoBehaviour named `RadialProgressComponent.cs` with the following content:

```
using MyUILibrary;
using UnityEngine;
using UnityEngine.UIElements;

[RequireComponent(typeof(UIDocument))]
public class RadialProgressComponent : MonoBehaviour
{
    RadialProgress m_RadialProgress;

    void OnEnable()
    {
        var root = GetComponent<UIDocument>().rootVisualElement;

        m_RadialProgress = new RadialProgress() {
            style = {
                position = Position.Absolute,
                left = 20, top = 20, width = 200, height = 200
            }
        };

        root.Add(m_RadialProgress);
    }

    void Update()
    {
        m_RadialProgress.progress = ((Mathf.Sin(Time.time) + 1.0f) / 2.0f) * 60.0f + 10.0f;
    }
}
```

## Test the progress indicator in runtime

1. In Unity, select **GameObject** > **UI Toolkit** > **UI Document**.
2. Select the **UIDocument** in the Hierarchy window.
3. Add **RadialProgressComponent.cs** as a component of the UIDocument **GameObject**.
4. Enter play mode. The progress indicator appears in the **scene**, and the progress ring and value change dynamically.

## Additional resources

* [Create custom controls](UIE-create-custom-controls.md)