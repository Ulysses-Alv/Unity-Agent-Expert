# Introduction to TransformHandle API

[Switch to Scripting](../ScriptReference/TransformHandle.md "Go to TransformHandle page in the Scripting Reference")

The [`TransformHandle`](../ScriptReference/TransformHandle.md) API is an alternative to the [`Transform`](../ScriptReference/Transform.md) API. Unlike the `Transform` component (a managed class), `TransformHandle` is an unmanaged `struct`, which makes it fully compatible with the [Burst](http://docs.unity3d.com/Packages/com.unity.burst@latest) compiler.

While the `TransformHandle` API covers the same core operations as the `Transform` API, its design introduces several key differences described on this page.

## Differences with Transform API

The `TransformHandle` API functionality is equivalent to the `Transform` API, the key differences are in the way you access it, and in the implementation details.

**Transform API (Unity 6.2 and earlier):**

* You access the transform using the `GameObject.transform` property.
* The `transform` property returns a `Transform` component reference (managed type).
* The API is managed and not compatible with the Burst compiler.
* The API works alongside `TransformHandle` API in Unity 6.

**TransformHandle API (Unity 6.3 and later):**

* You access the transform using `gameObject.transformHandle` or `transform.GetTransformHandle` (where `transform` is the `Transform` component).
* Returns a `TransformHandle` struct (unmanaged type).
* Is compatible with the Burst compiler and can be used in Burst-compiled jobs.
* Ensures compatibility with future entity and **GameObject** interactions.
* TransformHandle API is runtime-only.

### Transform APIs that are not available in the TransformHandle API

The following methods and properties from `Transform` APIs are not available in `TransformHandle` API:

* `GetSiblingIndex` / `SetSiblingIndex`
* `Find`
* `hasChanged` property

### TransformHandle API without an equivalent in the Transform API

The following `TransformHandle` API elements don’t have an equivalent in the `Transform` API.

* [`TransformHandle.DirectChildrenEnumerable`](../ScriptReference/TransformHandle.DirectChildrenEnumerable.md) represents an enumerable of the direct children of a transformHandle, and this enumerable returns a DirectChildrenEnumerator when enumerated.
* [`TransformHandle.DirectChildrenEnumerator`](../ScriptReference/TransformHandle.DirectChildrenEnumerator.md) is an enumerator that goes through the direct children of a transform.
* [`TransformHandle.DirectChildren`](../ScriptReference/TransformHandle.DirectChildren.md) returns a DirectChildrenEnumerable of this transform, which you can use to iterate over direct children with a foreach loop.
* [`TransformHandle.GetDirectChildrenEnumerator`](../ScriptReference/TransformHandle.GetDirectChildrenEnumerator.md) returns a DirectChildrenEnumerator of this transform, which you can use to manually iterate over direct children.

### Differences in common operations

The following common operations have a different implementation in `TransformHandle` API.

Check if a transform is valid:

* **Transform**: `transform != null`
* **TransformHandle**: `TransformHandle.IsValid`

Iterate the direct children of the transform:

* **Transform**: `foreach (Transform t in transform)`
* **TransformHandle**: `foreach (TransformHandle t in handle.DirectChildren)`

Set a parent of a transform to `None`:

* **Transform**: `transform.SetParent(null)`
* **TransformHandle**:

  ```
  transformHandle.SetParent(TransformHandle.None)
  ```

  or

  ```
  TransformHandle h = transformHandle;
  h.parent = TransformHandle.None;
  ```

For specific examples on how to use the API, refer to [TransfromHandle API examples](transformhandle-examples.md).

### Methods with different argument or return types

The following methods are functionally equivalent, but have different argument types or return types (`Transform`, or `TransformHandle` depending on the API they are called from).

Different argument types:

* `IsChildOf`
* `LookAt`
* `SetParent`
* `Translate` (overload)

Different return types:

* `GetChild`
* `parent`
* `root`

## Additional resources

* [TransfromHandle API examples](transformhandle-examples.md)