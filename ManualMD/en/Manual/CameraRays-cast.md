# Cast a ray from a camera

The most common use of a Ray from the **camera** is to perform a [raycast](../ScriptReference/Physics.Raycast.md) out into the **scene**. A raycast sends an imaginary “laser beam” along the ray from its origin until it hits a **collider** in the scene. Information is then returned about the object and the point that was hit in a [RaycastHit](../ScriptReference/RaycastHit.md) object. This is a very useful way to locate an object based on its onscreen image. For example, the object at the mouse position can be determined with the following code:

```
using UnityEngine;
using System.Collections;

public class ExampleScript : MonoBehaviour {
    public Camera camera;

    void Start(){
        RaycastHit hit;
        Ray ray = camera.ScreenPointToRay(Input.mousePosition);
        
        if (Physics.Raycast(ray, out hit)) {
            Transform objectHit = hit.transform;
            
            // Do something with the object that was hit by the raycast.
        }
    }
}
```