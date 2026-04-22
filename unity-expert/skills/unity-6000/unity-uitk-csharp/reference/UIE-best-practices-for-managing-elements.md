# Best practices for managing elements

This page describes the best practices for managing elements in the **visual tree**.

## Pool recurring elements

Elements pooling is to keep hold of elements that you might recreate later on, rather than creating elements with `new()` every time and letting go of them.

When you implement a custom object pool for VisualElements, take full control of all pooled elements and reset them before returning them to the pool. UI Toolkit provides internal pooling for systems like events ([EventBase](../ScriptReference/UIElements.EventBase_1.GetPooled.md)) to improve performance, but it doesn’t include a built-in pool for VisualElements. If you create your own pool, you must handle cleanup. For example, if you register an event callback on an element and return it to your pool without un-registering it, the callback remains. This can cause instability, unexpected behavior, and memory leaks when the element is reused.

## Keep the number of visible elements low

To keep the number of **visual elements** low, use [ListView](UIE-uxml-element-ListView.md) when possible. [ListView](UIE-uxml-element-ListView.md) pools elements and recycles elements as the user scrolls.

Alternatively, you can implement your own pool and recycle mechanism similar to the [ListView](UIE-uxml-element-ListView.md), and use the following to manage the visible area:

* To observe the size of a container, use [`GeometryChangedEvent`](../ScriptReference/UIElements.GeometryChangedEvent.md)
* To compute the size of children, use the [`VisualElement.layout`](../ScriptReference/UIElements.VisualElement-layout.md) property

## Different approaches to hiding an element

When you use [`VisualElement.RemoveFromHierarchy()`](../ScriptReference/UIElements.VisualElement.RemoveFromHierarchy.md) to remove an element from the hierarchy and eliminate references to it, the element is garbage collected. This reduces the CPU and GPU cost to zero and frees a significant amount of memory. However, it’s a slow and costly operation to recreate elements and reload them in a hierarchy. To avoid this, you can pre-create the elements in a hierarchy, use [USS style](UIE-USS-Properties-Reference.md) properties to hide them, and only display them when necessary. While applying styles is generally faster, it can lead to increased memory usage if you create a large number of elements simultaneously.

The following describes the different approaches to hiding elements and the consequences on processors and memory usage.

### Hide with USS style `visibility: hidden;`

With this approach, the descendants can override the `visibility` style.

#### Single-frame cost

The following table describes the different aspects of single-frame cost when you hide or display a visual element with the `visibility` style:

| Aspect | `visibility: hidden;` | `visibility: visible;` |
| --- | --- | --- |
| **Styles** | Evaluated for the element and descendants to propagate the visibility. | Evaluated for it and the descendants, to propagate the visibility. |
| **Layout data** | Preserved | None |
| **Rendering commands** | Removed and deallocated | Recreated and reinserted into the chain of commands. |
| **Meshes** | Scheduled for deallocation. | Re-tessellated |

#### Per-frame behavior

The following table describes the per-frame behavior for CPU and GPU when you hide a visual element with the `visibility` style:

| Processor | Aspect | Per-frame behavior |
| --- | --- | --- |
| **CPU** | **Styles** | Fully evaluated to the element and its descendants. |
|  | **Layout data** | Updated |
|  | **Tessellation** | Minimal impact that only involves **stencil masking meshes**, if applicable. |
|  | **Rendering commands** | No commands to draw regular visible geometry. However, stencil masking meshes are still rendered to push or pop from the stencil, ensuring potential visible descendants are masked. |
| **GPU** | **Meshes** | Vertex and fragment shading on stencil masking meshes. |

### Hide with USS style `opacity: 0;`

With this approach, some GPU impact can be expected, as the vertex **shader** processes all elements.

#### Single-frame cost

The following table describes the single-frame cost when you hide or display a visual element with the `opacity` style:

| Action | Single-frame cost |
| --- | --- |
| `opacity: 0;` | The first time when you set the `opacity` to a value other than `1`, the UI Toolkit renderer modifies the vertices to accelerate the application of the opacity on GPU. This triggers a one-time, minimal CPU cost. While typically negligible, this cost might become noticeable if the element has a large number of descendants or if there are many vertices to modify. This cost is not incurred again unless the element is removed from the visual tree and re-added. |
| `opacity: 1;` | None |

#### Per-frame behavior

The following table describes the per-frame behavior for CPU and GPU when you hide a visual element with the `opacity` style:

| Processor | Aspect | Per-frame behavior |
| --- | --- | --- |
| **CPU** | **Styles** | Fully evaluated to the element and its descendants. |
|  | **Tessellation** | Operates normally and responds to changes. |
|  | **Rendering commands** | Executed |
| **GPU** | **Meshes** | The **vertex shader** operates as though the `visibility` is set to `1`, which can be detrimental in GPU-bound projects if there are many vertices to process. Triangles are moved offscreen so there is no fragment shader cost. |

### Hide with USS style `display: none;`

With this approach, the element behaves like it is removed from the layout tree, which might potentially affect the layout of other elements.

#### Single-frame cost

The following table describes the different aspects of single-frame cost when you hide or display a visual element with the `display` style:

| Aspect | `display: none;` | `display: flex;` |
| --- | --- | --- |
| **Layout data** | Might recompute the layout of other elements. | Pending layout changes are processed. |
| **Rendering commands/Meshes** | Regenerated for elements affected by the layout change. | * Regenerated for elements affected by the layout change. * Pending changes are processed. |

#### Per-frame behavior

The following table describes the per-frame behavior of the CPU when you hide a visual element with the `display` style. Note that there’s no GPU cost.

| Aspect | Per-frame behavior |
| --- | --- |
| **Layout data** | Preserved but can become invalidated and not being updated. |
| **Rendering commands** | Although retained, they can be skipped during execution. The way they are skipped is very cheap but not entirely free. The cost is proportional to the number of commands. |
| **Meshes** | Retained but can become invalidated and not being updated. |

### Translate outside of the Viewport

You can use `translate: -5000px -5000px;` combined with [`DynamicTransform`](../ScriptReference/UIElements.UsageHints.DynamicTransform.md) usage hints to move the elements out of the **Viewport**. The geometry remains fully active, resulting in minimal CPU usage when you bring back the element on the screen. However, the GPU continues to process the vertices, which might be acceptable depending on the scenario.

#### Single-frame cost

The transform is computed and uploaded into GPU memory, which is generally fast.

#### Per-frame behavior

The following table describes the per-frame behavior for CPU and GPU when you hide a visual element by translating it outside of the Viewport:

| Processor | Aspect | Per-frame behavior |
| --- | --- | --- |
| **CPU** | **Styles** | Updated |
|  | **Layout date** | Updated |
|  | **Draw calls** | Executed |
| **GPU** | **Meshes** | Vertices are shaded. |

### Remove from the hierarchy

When you use the [`VisualElement.RemoveFromHierarchy()`](../ScriptReference/UIElements.VisualElement.RemoveFromHierarchy.md) method to remove the element from the hierarchy, you free up CPU and GPU memories, thereby eliminating any computing costs.

#### Single-frame cost

The following table describes the single-frame cost when you hide or display a visual element by removing it from the hierarchy:

| Aspect | Remove | Add |
| --- | --- | --- |
| **Styles** | None | Updated for the subtree. |
| **Layout** | None | Recomputed for the subtree and possibly other elements. |
| **Rendering commands/meshes** | Regenerated for elements affected by the layout change. | * Regenerated for elements affected by the layout change. * Pending changes are processed. |

### Memory usage after being hidden

The following table summarizes the memory usage after the element is hidden with different approaches:

| Processor | Aspect | `visibility:hidden;` | `opacity:0;` | `display:None;` | Translated out of Viewport | Removed from the hierarchy |
| --- | --- | --- | --- | --- | --- | --- |
| **CPU** | **Styles** | Retained | Retained | Retained | Retained | Freed |
|  | **Layout** | Retained | Retained | Retained | Retained | Retained[[1]](#fn:1 "see footnote") |
|  | **Rendering commands/meshes** | Freed | Retained | Retained | Retained | Freed |
| **GPU** | **Meshes** | Freed | Retained | Retained | Retained | Freed |

---

1. The layout memory is retained because it remains reserved for the element. When the VisualElement is garbage collected, the layout memory is returned to the pool, making it available for use by other elements. [↩](#fnref:1 "return to article")