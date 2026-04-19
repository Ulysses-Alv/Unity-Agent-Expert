# Mesh topology data

Topology describes the type of face that a **mesh** has.

A mesh’s topology defines the structure of the index buffer, which in turn describes how the vertex positions combine into faces. Each type of topology uses a different number of elements in the index array to define a single face.

Unity supports the following mesh topologies:

* Triangle
* Quad
* Lines
* LineStrip
* Points

**Note:** The Points topology doesn’t create faces; instead, Unity renders a single point at each position. All other mesh topologies use more than one index to create either faces or edges.

In the `Mesh` class, you can get the topology with [Mesh.GetTopology](../ScriptReference/Mesh.GetTopology.md), and set it as a parameter of [Mesh.SetIndices](../ScriptReference/Mesh.SetIndices.md).

For more information on supported mesh topologies, see the documentation for the MeshTopology enum.

**Note:** You must convert any meshes that use other modelling techniques (such as NURBS or NURMS/Subdivision Surfaces modelling) into supported formats in your modelling software before you can use them in Unity.