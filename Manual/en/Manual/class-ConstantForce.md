# Constant Force component reference

[Switch to Scripting](../ScriptReference/ConstantForce.md "Go to ConstantForce page in the Scripting Reference")

****Constant Force**** adds constant forces to a [Rigidbody](rigidbody-physics-section.md). This is useful for **GameObject** movement that accelerates over time.

If you add a Constant Force component to a GameObject that does not have a Rigidbody, Unity automatically creates and adds a Rigidbody to the same GameObject.

For more details, see [Apply constant force to a Rigidbody](rigidbody-constant-force.md).

## Properties

For all values, a higher value produces a stronger force, which in turn produces a faster initial velocity.

| **Property:** | **Function:** |
| --- | --- |
| **Force** | Define the direction of the linear force. The XYZ vectors refer to the **scene**’s global axes. |
| **Relative Force** | Define the direction of the linear force. The XYZ vectors refer to the Rigidbody’s local axes. |
| **Torque** | Define the global axes that the Rigidbody rotates around. |
| **Relative Torque** | Define the local axes that the Rigidbody rotates around. |

ConstantForce