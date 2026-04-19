# Troubleshooting Skinned Mesh Renderer visibility

Unity uses the **mesh**’s bounds to determine whether to render it. If the entire **bounding volume** is outside the view of any active **Camera**, Unity does not render the mesh.

In the case of a Skinned **Mesh Renderer**, its mesh bounds change as it deforms. Unity accounts for all animations present at import time when it calculates the maximum bounding volume, and uses this value to determine visibility; however, the following situations that occur after import might push vertices or bones outside of the maximum known bounds:

* Additive animations
* Changing the positions of bones from a script at runtime
* Using vertex **shaders** that change vertex positions
* Using ragdolls

If this happens, Unity might not correctly determine the visibility of the mesh, and might fail to show it when expected.

In these cases, you can try either of the following solutions to fix the problem:

* Modify the bounds to match the maximum potential bounding volume of your mesh. Use this option if possible, because it is better for performance.
* Enable **Update When Offscreen**, so that Unity continues to calculate the mesh bounds at all times, even when the mesh is not visible. Use this option if performance is not a major concern, or if you can’t predict the size of your bounding volume (for example, if you use ragdolls).