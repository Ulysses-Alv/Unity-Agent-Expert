# Introduction to 2D light batching

Use the **Light Batching Debugger** to visualize how Unity batches [2D Lights](2DLightProperties.md) and [Shadow Casters](2DShadows.md) according to the [Sorting Layers](../class-TagManager.md#SortingLayers) they target in the **scene**.

For Unity to batch Sorting Layers, the layers need to fulfill the following conditions:

* The layers share the same sets of Lights; that is, the 2D Lights target the same Sorting Layers.
* The layers share the same sets of Shadow Casters; that is, the Shadow Casters target the same Sorting Layers.

The debugger compares adjacent batches and highlights the Lights or Shadow Casters that target each Sorting Layer, and displays which Lights or Shadow Casters you need to add or remove for Unity to be able to batch the Sorting Layers.

Check how Unity batches 2D lights in your project with the [Light Batching Debugger window.](2d-light-batching-debugger.md).

## Examples of different batching scenarios

The following are examples of how Unity batches Lights and Shadow Casters under different conditions. Each example consists of two Sorting Layers named **BG** and **Default**, and two Lights named **A** and **B**.

### Scenario 1

**Conditions:**

* **Lights A** and **B** target both the **BG** and **Default** Sorting Layers.
* Shadows are disabled for both Lights; that is, there are no Shadow Casters.

![Batch Case 1](../../uploads/urp/2D/light-batching-debugger-4.png)

Batch Case 1

**Result:** Unity batches both Lights together as they target the same layers.

![Batch Case 1](../../uploads/urp/2D/light-batching-debugger-5.png)

Batch Case 1

### Scenario 2

**Conditions:**

* **Light A** targets **BG**, while **Light B** targets **Default**.
* Shadows are disabled for both Lights.

![Batch Case 2](../../uploads/urp/2D/light-batching-debugger-6.png)

Batch Case 2

**Result:** Unity doesn’t batch the layers as both Lights target different Sorting Layers.

![Batch Case 2](../../uploads/urp/2D/light-batching-debugger-7.png)

Batch Case 2

### Scenario 3

**Conditions:**

* Both **Lights A** and **B** target both **BG** and **Default** Sorting Layers.
* Shadows are enabled for both Lights and the Shadow Casters target both the **BG** and **Default** layers.

![Batch Case 3](../../uploads/urp/2D/light-batching-debugger-8.png)

Batch Case 3

**Result:** Unity batches the layers as both Lights and sets of Shadow Casters target the same layers.

![Batch Case 3](../../uploads/urp/2D/light-batching-debugger-9.png)

Batch Case 3

### Scenario 4

**Conditions:**
\* Both **Lights A** and **B** target both **BG** and **Default**.
\* Shadows are only enabled for **Light A**, and the Shadow Caster targets both **BG** and **Default**.

![Batch Case 4](../../uploads/urp/2D/light-batching-debugger-10.png)

Batch Case 4

**Result:** Unity batches the layers as the Shadow Caster targets both Sorting Layers so that both layers share the same shadow settings, making the light texture the same for both layers.

![Batch Case 4](../../uploads/urp/2D/light-batching-debugger-11.png)

Batch Case 4

### Scenario 5

**Conditions:**

* Both **Lights A** and **B** target both the **BG** and **Default** Sorting Layers.
* Shadows are enabled for both Lights, and the Shadow Caster only targets the **BG** layer.

![Batch Case 5](../../uploads/urp/2D/light-batching-debugger-12.png)

Batch Case 5

**Result:** Unity doesn’t batch the layers as the Shadow Caster targets one layer and not the other; this results in the light textures of both layers not being the same and unable to be batched.

![Batch Case 5](../../uploads/urp/2D/light-batching-debugger-13.png)