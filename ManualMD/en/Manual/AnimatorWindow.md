# Animator window

Use the **Animator window** to create, view, and modify Animator Controller assets.

![The Animator window](../uploads/Main/MecanimAnimatorControllerWindow.png)

The Animator window

The Animator window displays the **state machine** from the most recently selected `.controller` asset, regardless of which **scene** is loaded.

The Animator window contains:

* [Animation Layers](AnimationLayers.md)
* [Animation Parameters](AnimationParameters.md)
* The Animator Controller view where you create, arrange, and connect states for the [Animator Controller](Animator.md).

You can right-click (macOS: **Ctrl\_\_+click) on the grid to create a new state node. Use the middle mouse button or press** Alt\_\_ (macOS: **Option**) and drag to pan the Animator Controller view. Click to select and edit a state node. Click and drag a state node to rearrange your state machine.

![The Parameters view with two parameters](../uploads/Main/AnimatorWindowParametersPane.png)

The Parameters view with two parameters

Use the Parameters view to create, view, and edit [Animator Controller Parameters](AnimationParameters.md). These are variables you define that act as inputs for the state machine. To add a parameter, select the Plus icon, then select the parameter type from the context menu. To delete a parameter, select the parameter in the list and press **Delete** (macOS: **Ctrl\_\_+click and select** Delete\_\_).

![The Layers view](../uploads/Main/AnimatorWindowLayersPane.png)

The Layers view

Use the Layers view to create, view, and edit [layers](AnimationLayers.md) for your Animator Controller. You can control each layer with a different state machine. For example, you can create a base layer that controls the full body animation of your character and a second layer that controls the upper-body animation.

![The Eye icon](../uploads/Main/AnimatorWindowEyeIcon.png)

The Eye icon

Enable or disable the Eye icon to display or hide the side pane. Hide the side pane to have more room to edit your state machine.

![The breadcrumb list](../uploads/Main/AnimatorWindowBreadcrumbLocation.png)

The breadcrumb list

States can contain [sub-states](NestedStateMachines.md) and [blend trees](class-BlendTree.md). You can nest these structures repeatedly. When you view a sub-state or blend tree within another state, the breadcrumb list displays the nested hierarchy. Select an item in the breadcrumb list to display the state, sub-state, or blend tree.

![The lock icon](../uploads/Main/AnimatorWindowLockIcon.png)

The lock icon

Select the Lock icon to focus the Animator window on the selected Animator Controller asset. The Animator window doesn’t change focus when you select a different Animator Controller asset. If you disable the Lock icon, the Animator window changes focus when you select a different Animator Controller asset.