# Configure Capsule Collider 2D

There are multiple properties for a Capsule **Collider** 2D that you must configure before it can be used correctly.

## Defining size and direction

The settings that define the shape of the **Capsule Collider** 2D are **Size** and **Direction**. Both the Size and Direction properties refer to **X** and **Y** (horizontal and vertical, respectively) in the local space of the Capsule Collider 2D, and not in world space.

A typical way to set up the Capsule Collider 2D is to set the **Size** to match the **Direction**. For example, if the Capsule Collider 2D’s **Direction** is **Vertical**, the **Size** of **X** is 0.5 and the **Size** of **Y** is 1, this makes the vertical direction capsule taller, rather than wider.

In the example below, the **X** and **Y** are represented by the yellow lines.

![An example of a Capsule Collider 2D set so the Size matches Direction](../../../../uploads/Main/CapsuleCollider2D-Example1.png)

An example of a Capsule Collider 2D set so the **Size** matches **Direction**

## Capsule configuration examples

You can change the Capsule Collider 2D with different configurations. Below are some examples.

Note that when the **X** and **Y** of the **Size** property are the same, the Capsule Collider 2D always approximates to a circle.

![Examples of Capsule Collider 2D configurations](../../../../uploads/Main/CapsuleCollider2D-Example2.png)

Examples of Capsule Collider 2D configurations

**Note:** A known issue in the 2D physics system is that when a **GameObject** moves across multiple Colliders, one or several of the Colliders may register a **collision** between the Colliders. This may occur even when the Colliders are perfectly aligned. This collision can cause the Collider to slow down or stop.

While constructing a surface with the Capsule Collider 2D can help reduce this problem, it is recommeneded to use a single Collider rather than multiple Colliders for a surface, such as the [Edge Collider 2D](../edge-collider-2d-reference.md).