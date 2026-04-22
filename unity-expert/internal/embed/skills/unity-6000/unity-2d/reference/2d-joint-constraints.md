# 2D joint constraints

A constraint is a rule which a **joint** will try to ensure isn’t permanently broken. There are different types of constraints, and all joints provide at least one constraint that apply to and govern the **Rigidbody** 2D behavior. Some constraints limit behavior such as ensuring a Rigidbody stays on a line, or in a certain position. Some are ‘driving’ constraints which actively compel a Rigidbody **GameObject** to behave in a certain way, such as trying to make a GameObject maintain a certain speed.

## Temporarily breaking constraints

Unity’s physics system expects that constraints can be temporarily broken, such as when the objects move further apart than their set distance constraint or move faster than their set speed constraint. When a constraint isn’t broken, the joint doesn’t apply any forces. When a constraint is broken, the joint applies forces to fix the constraint.

For example, with ‘driving’ constraints, the joint applies forces to maintain the distance or ensure the speed set by the constraint. While this application of force is typically performed quickly, it doesn’t always instantly fix the constraint and instead it fixes the constraint gradually over time. This can lead to joints appearing to stretch or appear less rigid. The lag happens because the physics system is trying to apply joint forces to fix constraints, while other physics forces are simultaneously still acting to break those same constraints. In addition to the conflicting forces acting on GameObjects, some joints are more stable and react faster than others.

Whatever constraints the joint provides, the joint only uses forces to fix the constraint. These are either a linear (straight line) force or angular (torque) force.

**Note**: It’s recommended to be cautious when applying large forces to Rigidbody objects that have joints attached, especially those with large masses, due to the conflicting forces acting on joints.

## Permanently breaking joints

All joints are able to monitor the force or torque that they’re applying to stay within its own constraints. Some joints monitor both force and torque while others monitor only force. This informs you of when a joint exceeds a specific force or torque in trying to maintain its constraints, and you can specify these thresholds as [`Joint2D.breakForce`](../../../ScriptReference/Joint2D-breakForce.md) and [`Joint2D.breakTorque`](../../../ScriptReference/Joint2D-breakTorque.md). When a joint exceeds these thresholds, it’s known as joint breaking.

You can specify the action to be taken when a joint breaks with [`Joint2D.breakAction`](../../../ScriptReference/Joint2D-breakAction.md). The default break action is to destroy the Joint2D component, and you can refer to [`JointBreakAction2D`](../../../ScriptReference/JointBreakAction2D.md) for other available fixed actions.

## Additional resources

* [API documentation on Joint2D](../../../ScriptReference/Joint2D.md)