# Create a ragdoll

You can create a skinned **mesh** character and rig it up with bones in a 3D modeling application.

Unity can then generate all **colliders**, **rigidbody** components, and joints that make up a ragdoll from your skinned mesh character.

1. Create and rig a skinned mesh character in your 3D modeling application.
2. Save the skinned mesh character in your project folder.
3. Launch Unity, and select the skinned mesh file. The **Import Settings** dialog appears in the **Inspector** window. If needed, disable **Generate Colliders** in the **Import Settings** dialog.
4. Create an instance of the character by dragging it from the **Project** window
   to the **Hierarchy** window..
5. Select **GameObject > 3D Object > Ragdoll…** from the menu bar.
6. Drag the different limbs of your character instance from the **Hierarchy** window to the matching wizard area and create your ragdoll. The inspector displays the following components:

   * Skinned mesh renderer
   * Box collider
   * Rigidbody

   You can check how your ragdoll moves in play mode.
7. To save the ragdoll as a prefab, select **Assets -> Create -> Prefab** from the menu bar.
8. Drag the ragdoll character instance from the Hierarchy on top of the ragdoll **prefab**.

You can now reuse your ragdoll character prefab in your project.

For **character joints** made with the Ragdoll wizard, the following naming scheme applies:

| **Character joint axis** | **Description** |
| --- | --- |
| **Twist** | Twists the limb. |
| **Swing 1** | Limb’s largest **swing axis**. |
| **Swing 2** | Limb’s smaller swing axis. |

## Additional resources

* [Joints](joints-section.md)
* [Character Joint component reference](class-CharacterJoint.md)