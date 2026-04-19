## Troubleshooting materials missing specular highlights

Fix issues causing specular highlights to be missing after baking **lightmaps**.

![Two images of a rendered room with a sphere in the middle. The first image shows the sphere has several baked point lights that reflect on its surface. The second image shows the same scene after the lightmap is baked and the specular highlights in the first image are missing.](../uploads/Main/ex-missing-spec-response.png)

Two images of a rendered room with a sphere in the middle. The first image shows the sphere has several baked point lights that reflect on its surface. The second image shows the same scene after the lightmap is baked and the specular highlights in the first image are missing.

## Symptoms

Materials don’t appear glossy or have specular highlights when the lighting is baked.

## Cause

[Baked lights](LightModes-introduction.md#baked-light-behavior) don’t provide real-time specular response to materials. This means that glossy materials lack specular highlights after you generate lighting.

## Resolution - Use mixed lights

Use [mixed lights](LightModes-introduction.md#mixed) instead of baked lights, as mixed lights provide real-time direct specular response to materials. To use mixed lights, in the [Light](class-Light.md) component, set the ****Light Mode**** property to **Mixed**.

## Resolution - Use emissive proxies

![Two images of a rendered room with a sphere in the middle. The first image shows the sphere with bake point lights, and emissive proxies captured by a reflection probe. The second image is a zoomed out perspective of the first, showing the emissive proxies are placed above the sphere.](../uploads/Main/ex-emissive-proxies.png)

Two images of a rendered room with a sphere in the middle. The first image shows the sphere with bake point lights, and emissive proxies captured by a reflection probe. The second image is a zoomed out perspective of the first, showing the emissive proxies are placed above the sphere.

Use emissive proxies to imitate the specular response from using emissive objects. To use emissive proxies, follow these steps:

1. Place a [reflection probe](class-ReflectionProbe.md) in your **scene**.
2. Right-click the **Hierarchy** panel and select **3D Object** > **Sphere**.
3. Select the newly created object and set its [Static Editor Flag](StaticObjects.md) to **Reflection Probe Static**.
4. Right-click the **Project** panel and select **Create** > **Material**.
5. Select the material and enable the **Emission** checkbox. Set the ****Global Illumination**** property to **None**.
6. Drag the material onto the sphere to assign it.
7. Place the sphere in the same position as your light.
8. Go to the **Baked Lightmaps** tab in the Lighting window (**Window** > **Rendering** > **Lighting**).
9. Select the **Generate Lighting** button to generate lighting based on these settings.

After Unity generates the lighting, the emissive objects are captured in the reflection probe **cubemap**. You can hide the emissive proxies after baking or use a ****Culling Mask**** in the [Camera](class-Camera.md) component to hide the emissive objects.