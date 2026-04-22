# Animation Parameters

Animation Parameters are variables that are defined within an **Animator Controller** that can be accessed and assigned values from **scripts**. This is how a script can control or affect the flow of the **state machine**.

![Animation Parameters in the Animator window.](../uploads/Main/AnimationEditorParametersSection.png)

Animation Parameters in the Animator window.

For example, the value of a parameter can be updated by an [animation curve](animeditor-AnimationCurves.md) and then accessed from a script so that, say, the pitch of a sound effect can be varied as if it were a piece of animation. Likewise, a script can set parameter values to be picked up by Mecanim. For example, a script can set a parameter to control a [Blend Tree](class-BlendTree.md).

Default parameter values can be set up using the Parameters section of the **Animator window**, selectable in the top right corner of the Animator window. They can be of four basic types:

* *Integer* - a whole number
* *Float* - a number with a fractional part
* *Bool* - true or false value (represented by a checkbox)
* *Trigger* - a boolean parameter that is reset by the controller when consumed by a transition (represented by a circle button)

Parameters can be assigned values from a script using functions in the Animator class: [SetFloat](../ScriptReference/Animator.SetFloat.md), [SetInteger](../ScriptReference/Animator.SetInteger.md), [SetBool](../ScriptReference/Animator.SetBool.md), [SetTrigger](../ScriptReference/Animator.SetTrigger.md), and [ResetTrigger](../ScriptReference/Animator.ResetTrigger.md).

Here’s an example of a script that modifies parameters based on user input and **collision** detection:

```
using UnityEngine;
using System.Collections;

public class SimplePlayer : MonoBehaviour {

    Animator animator;
    
    // Use this for initialization
    void Start () {
        animator = GetComponent<Animator>();
    }

    // Update is called once per frame
    void Update () {
        float h = Input.GetAxis("Horizontal");
        float v = Input.GetAxis("Vertical");
        bool fire = Input.GetButtonDown("Fire1");

        animator.SetFloat("Forward",v);
        animator.SetFloat("Strafe",h);
        animator.SetBool("Fire", fire);
    }

    void OnCollisionEnter(Collision col) {
        if (col.gameObject.CompareTag("Enemy"))
        {
            animator.SetTrigger("Die");
        }
    }
}
```