# Shader Graph Asset reference | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Shader-Graph-Asset.html)

# Shader Graph Asset reference

A Shader Graph asset is a file that contains a graph you create and edit in the [**Shader Graph** window](Shader-Graph-Window.html). It is also a shader that you can select from a material's shader dropdown, as any other shader.

To access the properties of a Shader Graph asset in the **Inspector** window, select the asset in your project.

## Action buttons

Manage Shader Graph assets code.

| Property | Description |
| --- | --- |
| **Open Shader Editor** | Opens the selected asset in the [Shader Graph window](Shader-Graph-Window.html) so you can edit the graph. |
| **View Generated Shader** | Opens the shader code that the shader graph generates in a text editor or an IDE, such as Visual Studio. The code includes all possible passes and targets. |
| **Regenerate** | Updates the code you edited in your text editor or IDE. This button appears only when you select **View Generated Shader**. |
| **Copy Shader** | Copies the shader code to the clipboard. |

## Properties

Manage Shader Graph assets templates.

| Property | Description |
| --- | --- |
| **Use As Template** | Marks the selected Shader Graph asset as a template. When enabled, the asset appears in the [Shader Graph template browser](template-browser.html), but is no longer listed in any material's **Shader** dropdown by default. |
| **Expose As Shader** | Keeps the asset listed in a material's **Shader** dropdown so you can still use it as a shader when you also use it as a template. This is available only when **Use As Template** is enabled. |
| **Name** | Sets the name of the template in the Shader Graph template browser. |
| **Category** | Sets the category of the template in the Shader Graph template browser. |
| **Description** | Sets the description of the template in the Shader Graph template browser. |
| **Icon** | Sets the icon that represents the template in the Shader Graph template browser. |
| **Thumbnail** | Sets the image that represents the template in the Shader Graph template browser. |

## Additional resources

* [Creating a new shader graph asset](Create-Shader-Graph.html)
* [Shader Graph Window](Shader-Graph-Window.html)
* [Shader Graph template browser](template-browser.html)