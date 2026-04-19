# Retarget Humanoid animations

One of the most powerful features of Mecanim is the ability to retarget **humanoid animations**. This means that you can apply the same set of animations to different character models. **Retargeting** is only possible for humanoid models with a configured **Avatar**. A configured avatar provides the ability to correspondence between the models’ bone structure.

## Recommended Hierarchy structure

When working with Mecanim animations, you can expect your **scene** to contain the following elements:

* The Imported character model, which has an Avatar on it.
* The **Animator Component**, referencing an **Animator Controller** asset.
* A set of **animation clips**, referenced from the Animator Controller.
* **Scripts** for the character.
* Character-related components, such as the **Character Controller**.

Your project should also contain another character model with a valid Avatar.

To retarget between character models, follow these steps:

1. Create a **GameObject** in the Hierarchy that contains Character-related components.

   ![The MainChar GameObject with its Character Controller component.](../uploads/Main/MecanimRetargetingTopLevel.png)

   The MainChar GameObject with its Character Controller component.
2. Put the model as a child of the GameObject, together with the Animator component.

   ![The child Kyle_Robot GameObject with its Animator Controller component.](../uploads/Main/MecanimRetargetingModel.png)

   The child Kyle\_Robot GameObject with its Animator Controller component.
3. Make sure scripts referencing the Animator are looking for the animator in the children instead of the root. To do this, use `GetComponentInChildren<Animator>()` instead of `GetComponent<Animator>()`.

   ![The Kyle_Robot character.](../uploads/Main/MecanimRetargetingKyle.jpg)

   The Kyle\_Robot character.

Then, to reuse the same animations on another model, do the following:

1. Disable the original model.
2. Drop in the desired model as another child of GameObject.

   ![The child Teddy GameObject with its Animator Controller component set to no controller.](../uploads/Main/MecanimRetargetingOtherModel.png)

   The child Teddy GameObject with its Animator Controller component set to no controller.
3. Make sure the Animator Controller property for the new model is referencing the same controller asset.

   ![The Teddy GameObject with its Animator Controller set to the same controller as the Kyle GameObject.](../uploads/Main/MecanimRetargetingOtherModelCorrectController.png)

   The Teddy GameObject with its Animator Controller set to the same controller as the Kyle GameObject.
4. Tweak the character controller, the transform, and other properties on the top-level GameObject to make sure that the animations work smoothly with the new model.

   ![The Teddy character.](../uploads/Main/MecanimRetargetingTed.jpg)

   The Teddy character.