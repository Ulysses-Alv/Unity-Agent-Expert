# Aim Constraint component

An Aim Constraint rotates a **GameObject** to face its source GameObjects. It can also maintain a consistent orientation for another axis. For example, you can add an Aim Constraint to a **Camera**. To keep the Camera upright while the Constraint aims it, specify the up axis of the Camera and an up direction to align it to.

Use the **Up Vector** to specify the up axis of a constrained GameObject. Use the **World Up Vector** to specify the upward direction. As the Aim Constraint rotates the GameObject to face its source GameObjects, the Constraint also aligns the up axis of the constrained GameObject with the upward direction.

![Aim Constraint component](../uploads/Main/AimConstraint.png)

Aim Constraint component

### Properties

| **Property** | **Description** |
| --- | --- |
| **Activate** | After you rotate the constrained GameObject and move its source GameObjects, click **Activate** to save this information. **Activate** saves the current offset from the source GameObjects in **Rotation At Rest** and **Rotation Offset**, then enables **Is Active** and **Lock**. |
| **Zero** | Sets the rotation of the constrained GameObject to the source GameObjects. Zero resets the **Rotation At Rest** and **Rotation Offset** fields, then enables **Is Active** and **Lock**. |
| **Is Active** | Enable to evaluate the Constraint. To also apply the Constraint, make sure **Lock** is enabled. |
| **Weight** | The strength of the Constraint. A weight of 1 causes the Constraint to rotate the GameObject at the same rate that its source GameObjects move. A weight of 0 removes the effect of the Constraint completely. This weight affects all source GameObjects. Each GameObject in the **Sources** list also has a weight. |
| **Aim Vector** | Specifies the axis which faces the direction of its source GameObjects. For example, to orient the GameObject only on its positive z-axis to face the source GameObjects, enter an **Aim Vector** of 0, 0, 1 for the x, y, and z axes, respectively. |
| **Up Vector** | Specifies the up axis of this GameObject. For example, to keep the GameObject’s positive y-axis pointing upward, enter an **Up Vector** of 0, 1, 0 for the x, y, and z axes, respectively. |
| **World Up Type** | Specifies the axis for the upward direction. The Aim Constraint uses this vector to align the up axis of the GameObject the upward direction. Choose from the following:  * **Scene Up** Uses the y-axis of the scene as the upwards direction. * **Object Up**: Uses the y-axis of the GameObject referred to by **World Up Object** as the up direction. * **Object Up Rotation**: Uses the axis specified by **World Up Vector** of the GameObject referred to by **World Up Object** as the up direction. * **Vector**: Uses the World Up Vector as the up direction. * **None**: Do not use a World Up vector. |
| **World Up Vector** | Specifies the vector to use for the **Object Up Rotation** and **Vector** choices in **World Up Type**. |
| **World Up Object** | Specifies the GameObject to use for the **Object Up** and **Object Up Rotation** choices in **World Up Type**. |

### Constraint settings

Contains additional settings for the Aim Constraint.

| **Property** | **Description** |
| --- | --- |
| **Lock** | Enable this setting to let the Constraint rotate the GameObject. Disable this property to edit the rotation of this GameObject. You can also edit the Rotation At Rest and Rotation Offset properties. If Is Active is enabled, the Constraint updates the Rotation At Rest or Rotation Offset properties for you as you rotate the GameObject or its source GameObjects. When you are satisfied with your changes, enable Lock to let the Constraint control this GameObject. This property has no effect in Play mode. |
| **Rotation At Rest** | The X, Y, and Z values to use when Weight is 0 or when the corresponding Freeze Rotation Axes are not enabled. To edit these fields, disable Lock. |
| **Rotation Offset** | The X, Y, and Z offset from the rotation that is calculated by the Constraint. To edit these fields, disable Lock. |
| **Freeze Rotation Axes** | Enable X, Y, or Z to allow the Constraint to control the corresponding axes. Disable an axis to stop the Constraint from controlling it. This allows you to edit, animate, or script the unfrozen axis. |

### Sources

A list of GameObjects that constrain this GameObject. Unity evaluates source GameObjects in the order that they appear in this list. This order affects how this Constraint rotates the constrained GameObject. To get the result you want, drag the items in this list. Each source has a weight from 0 to 1.