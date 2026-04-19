# Create a skybox

To create a new **skybox** material:

1. From the main menu, select **Assets > Create > Material**.
2. In the new material’s ****Inspector**** window, search for `skybox` in the ****Shader**** dropdown, then select a type of skybox shader.
3. Configure the properties of the material to set up the skybox.

   The available properties depend on the type of skybox shader you selected.

   **Note**: Each skybox shader has unique prerequisite textures that differ in type, number, and format. For details on the required Textures for a specific skybox shader, refer to its documentation. You can find the list of skybox shaders and their respective documentation on the [Skybox shaders](skybox-shaders.md) page.

## Drawing a skybox in your scene

Once you have created a skybox material, you can apply it to your **scene**’s background. To draw the skybox in your scene:

1. From the main menu, select **Window** > **Rendering** > **Lighting**.
2. In the **Lighting** window, select the **Environment** tab.
3. Drag your skybox material to the **Skybox Material** property.

This draws the skybox in the background of every **camera** in your scene. To draw the skybox for a specific camera only, refer to [Rendering a skybox for a specific Camera](#DrawingForCamera).

## Drawing a skybox for a particular camera

To draw a skybox for a particular camera, use the [Skybox component](sky-landing.md). When attached to a **GameObject** with a camera, the Skybox component overrides the global skybox for that camera. To add and configure the Skybox component:

1. Select a camera in your scene and open the **Inspector** window.
2. Select **Add Component**, then add a new **Skybox** component.
3. Drag your skybox material to the **Custom Skybox** property.

## Best practices for skyboxes

### Synchronizing light sources

If your skybox contains a visible sun, moon, or other light source, align a **Directional Light** in your scene to match the light’s direction in the skybox. This creates the illusion that the light in the skybox is casting shadows. For scenes with multiple directional lights, specify the one to use:

1. From the main menu, select **Window** > **Rendering** > **Lighting**.
2. Select the **Environment** tab.
3. Drag the appropriate directional light to the **Sun Source** property.

### Blending fog with the skybox

To create a seamless blend between fog and the skybox, match the fog color to the skybox hues. To blend fog with the skybox:

1. From the main menu, select **Window** > **Rendering** > **Lighting**.
2. Select the **Environment** tab.
3. Under **Other Settings**, activate the **Fog** checkbox.
4. Adjust the **Color** property to complement your skybox.

   Use the eyedropper tool for precise color selection from the scene.