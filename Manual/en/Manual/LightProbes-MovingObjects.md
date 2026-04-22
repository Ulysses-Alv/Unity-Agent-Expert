# Light Probes and moving GameObjects

[Lightmapping](Lightmapping.md) adds greatly to the realism of a scene by capturing realistic bounced light as textures which are “baked” onto the surface of **static** objects. However, due to the nature of lightmapping, it can only be applied to non-moving objects marked as [Contribute GI](StaticObjects.md).

While realtime and mixed mode lights can cast *direct* light on moving objects, moving objects do not receive bounced light from your static environment unless you use **light probes**. Light probes store information about how light is bouncing around in your **scene**. Therefore as objects move through the spaces in your game environment, they can use the information stored in your light probes to show an approximation of the bounced light at their current position.

![A simple scene showing bounced light from static scenery.](../uploads/Main/LightProbes-MovingObjects-1.jpg)

A simple scene showing bounced light from static scenery.

In the above scene, as the directional light hits the red and green buildings, which are static scenery, *bounced light* is cast into the scene. The bounced light is visible as a red and green tint on the ground directly in front of each building. Because all these models are **static**, all this lighting is stored in **lightmaps**.

When you introduce moving objects into your scene, they do not automatically receive bounced light. In the below image, you can see the ambulance (a dynamic moving object) is not affected by the bounced red light coming off the building. Instead, its side is a flat grey color. This is because the ambulance is a dynamic object which can move around in the game, and therefore cannot use lightmaps, because they are static by nature. The scene needs Light Probes so that the moving ambulance can receive bounced light.

![The side of the ambulance is a flat grey color, even though it should be receiving some bounced red light from the front of the building.](../uploads/Main/LightProbes-MovingObjects-2.png)

The side of the ambulance is a flat grey color, even though it should be receiving some bounced red light from the front of the building.

To use the light probe feature to cast bounced light onto dynamic moving objects, you must position light probes throughout your scene, so that they cover the areas of space that moving objects in your game might pass through.

The probes you place in your scene define a 3D volume. The lighting at any position within this volume is then approximated on moving objects by interpolating between the information baked into the nearest probes.

![Light probes placed around the static scenery in a simple scene. The light probes are shown as yellow dots. They are shown connected by magenta lines, to visualise the volume that they define.](../uploads/Main/LightProbes-MovingObjects-3.png)

Light probes placed around the static scenery in a simple scene. The light probes are shown as yellow dots. They are shown connected by magenta lines, to visualise the volume that they define.

Once you have added probes, and baked the light in your scene, your dynamic moving objects will receive bounced light based on the nearest probes in the scene. Using the same example as above, the dynamic object (the ambulance) now receives bounced light from the static scenery, giving the side of the vehicle a red tint, because it is in front of the red building which is casting bounced light.

![The side of the ambulance now has a red tint because it is receiving bounced red light from the front of the building, via the light probes in the scene.](../uploads/Main/LightProbes-MovingObjects-4.png)

The side of the ambulance now has a red tint because it is receiving bounced red light from the front of the building, via the light probes in the scene.

When a dynamic object is selected, the **Scene view** will draw a visualisation of which light probes are being used for the interpolated bounced light. The nearest probes to the dynamic object are used to form a tetrahedral volume, and the dynamic object’s light is interpolated from the four points of this tetrahedron.

![The light probes that are being used to light a dynamic object are revealed in the scene view when the object is selected, connected by yellow lines to show the tetrahedral volume.](../uploads/Main/LightProbes-MovingObjects-5.jpg)

The light probes that are being used to light a dynamic object are revealed in the scene view when the object is selected, connected by yellow lines to show the tetrahedral volume.

As an object passes through the scene, it moves from one tetrahedral volume to another, and the lighting is calculated based on its position within the current tetrahedron.

![A dynamic object moving through a scene with light probes, showing how it passes from one tetrahedral light probe volume to another.](../uploads/Main/LightProbes-MovingObjects-6.gif)

A dynamic object moving through a scene with light probes, showing how it passes from one tetrahedral light probe volume to another.

---

* 2017–06–08 Page published
* Light Probes updated in 5.6