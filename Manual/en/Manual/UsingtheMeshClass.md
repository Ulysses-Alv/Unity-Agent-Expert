# Access meshes via the Mesh API

The [Mesh](../ScriptReference/Mesh.md) class is the basic script interface to an object’s mesh geometry. It uses arrays to represent the triangles, vertex positions, normals and texture coordinates and also supplies a number of other useful properties and functions to assist mesh generation.

## Access an object’s Mesh

The mesh data is attached to an object using the [Mesh Filter](class-MeshFilter.md) component (and the object will also need a [MeshRenderer](class-MeshRenderer.md) to make the geometry visible). This component is accessed using the familiar GetComponent function:

```
using UnityEngine;
public class ExampleScript : MonoBehaviour
{
    MeshFilter mf;
    void Start()
    {
        //if this gameObject has a MeshFilter, mf will reference the component
        mf = GetComponent<MeshFilter>();    
    }
}
```

## Add mesh data

The Mesh object has properties for the vertices and their associated data (normals and UV coordinates) and also for the triangle data. The vertices may be supplied in any order but the arrays of normals and UVs must be ordered so that the indices all correspond with the vertices (ie, element 0 of the normals array supplies the normal for vertex 0, etc). The vertices are [Vector3s](../ScriptReference/Vector3.md) representing points in the object’s local space. The normals are normalised Vector3s representing the directions, again in local coordinates. The UVs are specified as [Vector2s](../ScriptReference/Vector2.md), but since the Vector2 type doesn’t have fields called U and V, you must mentally convert them to X and Y respectively.

The triangles are specified as triples of integers that act as indices into the vertex array. Rather than use a special class to represent a triangle the array is just a simple list of integer indices. These are taken in groups of three for each triangle, so the first three elements define the first triangle, the next three define the second triangle, and so on. An important detail of the triangles is the ordering of the corner vertices. They should be arranged so that the corners go around clockwise as you look down on the visible outer surface of the triangle, although it doesn’t matter which corner you start with.

### Work with raw mesh data

The [Mesh](../ScriptReference/Mesh.md) class also has a lower-level advanced API
that enables you to work with raw mesh vertex and index buffer data. This is useful in situations that require maximum performance or the lowest amount of memory allocations.

* [Mesh.SetIndexBufferParams](../ScriptReference/Mesh.SetIndexBufferParams.md) and [Mesh.SetIndexBufferData](../ScriptReference/Mesh.SetIndexBufferData.md) for setting up size, format of the index buffer, and updating data inside it.
* [Mesh.SetVertexBufferParams](../ScriptReference/Mesh.SetVertexBufferParams.md) and [Mesh.SetVertexBufferData](../ScriptReference/Mesh.SetVertexBufferData.md) for setting up size, format of the vertex buffer(s), and updating data inside them.
* [Mesh.SetSubMesh](../ScriptReference/Mesh.SetSubMesh.md) for setting up index buffer topology and ranges.

## Additional resources

* [Mesh data](AnatomyofaMesh.md) documentation.
* [Mesh](../ScriptReference/Mesh.md) API reference.