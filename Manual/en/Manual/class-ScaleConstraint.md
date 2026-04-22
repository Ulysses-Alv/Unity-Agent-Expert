# Scale Constraint component

A Scale Constraint component resizes a **GameObject** to match the scale of its source GameObjects.

![Scale Constraint component](../uploads/Main/ScaleConstraint.png)

Scale Constraint component

## Properties

| **Property** | **Description** |
| --- | --- |
| **Activate** | After you resize the constrained GameObject and its source GameObjects, click **Activate** to save this information. **Activate** saves the current offset from the source GameObjects in **Scale At Rest** and **Scale Offset** then enables **Is Active** and **Lock**. |
| **Zero** | Sets the rotation of the constrained GameObject to the source GameObjects. **Zero** resets the **Scale At Rest** and **Scale Offset** fields then enables **Is Active** and **Lock**. |
| **Is Active** | Enable to evaluate the Constraint. To also apply the Constraint, make sure **Lock** is enabled. |
| **Weight** | The strength of the Constraint. A weight of 1 causes the Constraint to resize this GameObject at the same rate as its source GameObjects. A weight of 0 removes the effect of the Constraint completely. This weight affects all source GameObjects. Each GameObject in the **Sources** list also has a weight. |

### Constraint Settings

| **Property** | **Description** |
| --- | --- |
| **Lock** | Enable to let the Constraint resize the GameObject. Disable this property to edit the scale of this GameObject. You can also edit the **Scale At Rest** and **Scale Offset** properties. If **Is Active** is enabled, the Constraint updates the **Scale At Rest** or **Scale Offset** properties for you as you resize the GameObject or its Source GameObjects. When you are satisfied with your changes, enable **Lock** to let the Constraint to control this GameObject. This property has no effect in Play Mode. |
| **Scale At Rest** | The X, Y, and Z values to use when **Weight** is 0 or when the corresponding **Freeze Scale Axes** is not enabled. To edit these fields, disable **Lock**. |
| **Scale Offset** | The X, Y, and Z offset from the transform that is imposed by the Constraint. To edit these fields, disable Lock. |
| **Freeze Scale Axes** | Enable X, Y, or Z to allow the Constraint to control the corresponding axes. Disable an axis to stop the Constraint from controlling it. This allows you to edit, animate, or script the unfrozen axis. |

### Sources

The list of GameObjects that constrain this GameObject. Each source has a weight from 0 to 1.