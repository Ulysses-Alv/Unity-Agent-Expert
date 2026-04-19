# Character Controller component reference

[Switch to Scripting](../ScriptReference/CharacterController.md "Go to CharacterController page in the Scripting Reference")

The **Character Controller** is mainly used for third-person or first-person player control that does not make use of **Rigidbody** physics.

![The Character Controller inspector](../uploads/Main/Inspector-CharacterController.png)

The Character Controller inspector

## Properties

| **Property** | **Description** |
| --- | --- |
| **Slope Limit** | Limits the collider to only climb slopes that are less steep (in degrees) than the indicated value. |
| **Step Offset** | The character will step up a stair only if it is closer to the ground than the indicated value. This should not be greater than the Character Controller’s height or it will generate an error. |
| **Skin width** | Two colliders can penetrate each other as deep as their Skin Width. Larger Skin Widths reduce jitter. Low Skin Width can cause the character to get stuck. A good setting is to make this value 10% of the Radius. |
| **Min Move Distance** | If the character tries to move below the indicated value, it will not move at all. This can be used to reduce jitter. In most situations this value should be left at 0. |
| **Center** | This will offset the Capsule Collider in world space, and won’t affect how the Character pivots. |
| **Radius** | Length of the Capsule Collider’s radius. This is essentially the width of the collider. |
| **Height** | The Character’s **Capsule Collider** height. Changing this will scale the collider along the Y axis in both positive and negative directions. |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for this collider.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of this collider override. When two colliders have conflicting overrides, the settings of the collider with the higher value priority are taken.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in **collisions** with this collider. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with this collider. |

![The Character Controller](../uploads/Main/CharacterControllerWindow.png)

The Character Controller

## Details

The traditional Doom-style first person controls are not physically realistic. The character runs 90 miles per hour, comes to a halt immediately and turns on a dime. Because it is so unrealistic, use of Rigidbodies and physics to create this behavior is impractical and will feel wrong. The solution is the specialized Character Controller. It is simply a capsule shaped **Collider** which can be told to move in some direction from a script. The Controller will then carry out the movement but be constrained by collisions. It will slide along walls, walk up stairs (if they are lower than the **Step Offset**) and walk on slopes within the **Slope Limit**.

The Controller does not react to forces on its own and it does not automatically push Rigidbodies away.

If you want to push Rigidbodies or objects with the Character Controller, you can apply forces to any object that it collides with via the **OnControllerColliderHit()** function through scripting.

On the other hand, if you want your player character to be affected by physics then you might be better off using a [Rigidbody](class-Rigidbody.md) instead of the Character Controller.

### Fine-tuning your character

You can modify the **Height** and **Radius** to fit your Character’s **mesh**. It is recommended to always use around 2 meters for a human-like character. You can also modify the **Center** of the capsule in case your pivot point is not at the exact center of the Character.

**Step Offset** can affect this too, make sure that this value is between 0.1 and 0.4 for a 2 meter sized human.

**Slope Limit** should not be too small. Often using a value of 90 degrees works best. The Character Controller will not be able to climb up walls due to the capsule shape.

### Don’t get stuck

The **Skin Width** is one of the most critical properties to get right when tuning your Character Controller.
If your character gets stuck it is most likely because your **Skin Width** is too small. The **Skin Width** will let objects slightly penetrate the Controller but it removes jitter and prevents it from getting stuck.

It’s good practice to keep your **Skin Width** at least greater than 0.01 and more than 10% of the **Radius**.

We recommend keeping **Min Move Distance** at 0.

See the Character Controller script reference [here](../ScriptReference/CharacterController.md)

## Hints

* Try adjusting your **Skin Width** if you find your character getting stuck frequently.
* The Character Controller can affect objects using physics if you write your own **scripts**.
* The Character Controller can not be affected by objects through physics.
* Note that changing Character Controller properties in the **inspector** will recreate the controller in the **scene**, so any existing Trigger contacts will get lost, and you will not get any OnTriggerEntered messages until the controller is moved again.
* The Character Controller’s capsule used in queries such as raycast might shrink by a small factor. Queries therefore could miss in some corner cases, even when they appear to hit the Character Controller’s **gizmo**.