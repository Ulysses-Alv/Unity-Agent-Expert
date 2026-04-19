# Scripting Root Motion

Sometimes your animation comes as “in-place”, which means if you put it in a **scene**, it will not move the character that it’s on. In other words, the animation does not contain “**root motion**”. For this, you can modify root motion from a script. To put everything together follow the steps below (note there are many variations of achieving the same result, this is just one recipe).

1. Open the **inspector** for the FBX file that contains the in-place animation, and go to the **Animation** tab.
2. Make sure the **Muscle Definition** is set to the **Avatar** you intend to control (let’s say this avatar is called *Dude*, and it has already been added to the **Hierarchy View**).
3. Select the **animation clip** from the available clips.
4. Make sure **Loop Pose** is properly aligned (the light next to it is green), and that the checkbox for **Loop Pose** is clicked.

   ![The Animation tab displays the selected Run animation clip with the Loop Pose property enabled. The start and end of the loop match as indicated by Loop Match being green.](../uploads/Main/MecanimRootMotionChristmasTree.png)

   The Animation tab displays the selected Run animation clip with the Loop Pose property enabled. The start and end of the loop match as indicated by Loop Match being green.
5. Preview the animation in the animation viewer to make sure the beginning and the end of the animation align smoothly, and that the character is moving “in-place”.
6. On the animation clip [create a curve](animeditor-AnimationCurves.md) that will control the speed of the character (you can add a curve from the **Animation Import inspector** **Curves-> +**).
7. Name that curve something meaningful, like “Runspeed”.

   ![The Curves property expanded and displaying the animation curve for the Runspeed property.](../uploads/Main/MecanimRootMotionCurve.png)

   The Curves property expanded and displaying the animation curve for the Runspeed property.
8. Create a new **Animator Controller**, (let’s call it **RootMotionController**).
9. Drop the desired animation clip into it, this should create a state with the name of the animation (say **Run**).
10. Add a parameter to the Controller with the same name as the curve (in this case, “Runspeed”).

    ![The Runspeed parameter added to the BaseLayer of the RootMotionController.](../uploads/Main/MecanimRootMotionController.png)

    The Runspeed parameter added to the BaseLayer of the RootMotionController.
11. Select the character **Dude** in the **Hierarchy**, whose inspector should already have an **Animator** component.
12. Drag **RootMotionController** onto the **Controller** property of the Animator.
13. If you press play now, you should see the “Dude” running in place.
14. Finally, to control the motion, you will need to create a script (RootMotionScript.cs), that implements the [OnAnimatorMove](../ScriptReference/MonoBehaviour.OnAnimatorMove.md) callback:

    ```
    using UnityEngine;
    using System.Collections;

    [RequireComponent(typeof(Animator))]

    public class RootMotionScript : MonoBehaviour {

       void OnAnimatorMove()
       {
          Animator animator = GetComponent<Animator>();

          if (animator)
          {
             Vector3 newPosition = transform.position;
             newPosition.z += animator.GetFloat("Runspeed") * Time.deltaTime;
             transform.position = newPosition;
          }
       }
    }
    ```
15. You should attach RootMotionScript.cs to the “Dude” object. When you do this, the **Animator component** will detect that the script has an [OnAnimatorMove](../ScriptReference/MonoBehaviour.OnAnimatorMove.md) function and show the **Apply Root Motion** property as *Handled by Script*

    ![The Animator component with the Apply Root Motion property set to Handled by Script.](../uploads/Main/MecanimRootMotionDude.png)

    The Animator component with the Apply Root Motion property set to Handled by Script.