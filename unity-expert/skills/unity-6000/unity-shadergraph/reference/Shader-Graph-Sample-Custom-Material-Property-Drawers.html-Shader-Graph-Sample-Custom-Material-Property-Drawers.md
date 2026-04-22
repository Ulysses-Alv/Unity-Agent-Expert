# Custom Material Property Drawers sample | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Shader-Graph-Sample-Custom-Material-Property-Drawers.html)

## Custom Material Property Drawers sample

Learn how to use custom attributes in Shader Graph to improve the display and usability of property controls in the material's Inspector window.

The Custom Material Property Drawers sample includes a shader graph based material set up for adjustable color gradient and noise. This example shows how to achieve the following in the material Inspector UI with Shader Graph:

* Hide or display a basic material color property.
* Display a Vector2 property as a min/max slider.
* Display a help box for a property.

## Before you start

To access the Custom Material Property Drawers sample, follow these steps:

1. [Import](ShaderGraph-Samples-Import.html) the Custom Material Property Drawers sample in your project.
2. In the Project window, open the `Assets/Samples/Shader Graph/<your version>/Custom Material Property Drawers` folder.
3. Open the `RemapDrawerExample` shader graph asset.

## Hide or display a property in the material Inspector UI

To hide or display a material property from the shader graph in the material Inspector UI, follow these steps:

1. In the [Shader Graph window](Shader-Graph-Window.html), in the [Blackboard](Blackboard.html), select a property, for example **InnerColor**.
2. In the [Graph Inspector](Internal-Inspector.html), in the **Node Settings** tab, disable **Show in Inspector**.
3. Save the graph.
4. In the Project window, select the `RemapDrawerExample` material variant asset.
5. In the Inspector window, notice that the **InnerColor** property isn't displayed, contrary to the other properties of the graph.
6. In the [Shader Graph window](Shader-Graph-Window.html), re-enable **Show in Inspector** for the **InnerColor** property.
7. Save the graph.
8. In the Inspector window, notice that the **InnerColor** property is now displayed.

## Display a material property as a min/max range slider

To learn how to display a material property as a min/max range slider in the material Inspector UI, follow these steps:

1. In the [Shader Graph window](Shader-Graph-Window.html), in the [Blackboard](Blackboard.html), select the **Range** property.
2. In the [Graph Inspector](Internal-Inspector.html), in the **Node Settings** tab, notice the following:

   * The **Custom Attributes** include an entry named **Remap**, which modifies the **Range** property's appearance in the material UI based on a script provided with the sample.
   * The attribute includes no parameter values for the called function. For more information, refer to the sticky note besides the **Range** property instance in the graph.
3. In the Project window, open the `Editor` subfolder.
4. Select the `RemapDrawer` C# file.
5. In the Inspector window, notice the following:

   * The C# file includes a set of custom functions named `RemapDrawer` to display a Vector2 as a min/max slider.
   * The function names include a `Drawer` suffix while in the shader graph the **Range** property excludes this suffix to call the functions.

## Display a help box in the material Inspector UI

To learn how to display a help box in the material Inspector UI, follow these steps:

1. In the [Shader Graph window](Shader-Graph-Window.html), in the [Blackboard](Blackboard.html), select the **NoiseRange** property.
2. In the [Graph Inspector](Internal-Inspector.html), in the **Node Settings** tab, notice the following:

   * The **Custom Attributes** include an entry named **Remap** (previously described) and a second one named **HelpBox**, which adds a help box to the **NoiseRange** property in the material UI based on a second script provided with the sample.
   * The two attributes include specific parameter values for the called functions. For more information, refer to the sticky note under the **NoiseRange** property instance in the graph.
3. In the Project window, still in the `Editor` subfolder, select the `HelpBoxDecorator` C# file.
4. In the Inspector window, notice the following:

   * The C# file includes a set of custom functions named `HelpBoxDecorator` to draw a message box in the material UI.
   * The function names include a `Decorator` suffix while in the shader graph the **NoiseRange** property excludes this suffix to call the functions.

## Additional resources

* [Property attribute reference](Property-Types.html)