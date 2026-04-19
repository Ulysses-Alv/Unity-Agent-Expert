# Place Light components

Lights are easy to use in Unity - you need to create a light of the desired type (for example, from the menu **GameObject > Light > Point Light**) and place it where you want it in the scene. If you enable scene view lighting (the “lightbulb” button on the toolbar) then you can see a preview of how the lighting will look as you move light objects and set their parameters.

A directional light can be placed anywhere in the **scene** (unless it’s using a Cookie) with the forward/Z axis indicating the direction. A Spot Light also has a direction but since it has a limited range, its position *does* matter. The shape parameters of spot, point, and area lights can be adjusted from the **inspector** or by using the lights’ **Gizmos** directly in the **scene view**.

![A Spot Light with Gizmos visible](../uploads/Main/LightUsageSpotGizmo.png)

A Spot Light with Gizmos visible

## Guidelines for Placing Lights

A directional light often represents the sun and has a significant effect on the look of a scene. The direction of the light should point slightly downwards but you will usually want to make sure that it also makes a slight angle with major objects in the scene. For example, a roughly cubic object will be more interestingly shaded and appear to “pop” out in 3D much more if the light isn’t coming head-on to one of the faces.

Spot lights and point lights usually represent artificial light sources and so their positions are usually determined by scene objects. One common pitfall with these lights is that they appear to have no effect at all when you first add them to the scene. This happens when you adjust the range of the light to fit neatly within the scene. The range of a light is the limit at which the light’s brightness dims to zero. If you set, say, a spot light so the base of the cone neatly lands on the floor then the light will have little or no effect unless another object passes underneath it. If you want the level geometry to be illuminated then you should expand point and spot lights so they pass through the walls and floors.

## Color and Intensity

A light’s color and intensity (brightness) are properties you can set from the inspector. The default intensity and white color are fine for “ordinary” lighting that you use to apply shading to objects but you might want to vary the properties to produce special effects. For example, a glowing green forcefield might be bright enough to bathe surrounding objects in intense green light; car headlights (especially on older cars) typically have a slight yellow color rather than brilliant white. These effects are most often used with point and spot lights but you might change the color of a directional light if, say, your game is set on a distant planet with a red sun.