# Keyword parameter reference | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Keywords-reference.html)

# Keyword parameter reference

There are two types of keywords: Boolean and Enum. Each keyword type has a few specific parameters in addition to the many parameters that all keyword types have in common.

## Common parameters

Parameters that all keyword types have in common.

| **Name** | **Description** |
| --- | --- |
| **Name** | The display name of the keyword. Unity shows this name in the title bar of nodes that reference the corresponding keyword, and also in the Material Inspector if you expose that keyword. |
| **Reference** | The internal name for the keyword in the shader. Use this **Reference** name instead of the display **Name** when you reference the keyword in a script.  If you overwrite this parameter, be aware of the following:  * A Keyword **Reference** has to be in full capitals. Unity converts any lowercase letters to uppercase. * If the string contains any characters that HLSL does not support, Unity replaces those characters with underscores. * You can revert to the default value: right-click on the **Reference** field label, and select **Reset Reference**. |
| **Promote to final Shader** | Makes the keyword available across the entire shader rather than only within the subgraph. |
| **Definition** | Sets the keyword declaration type, which determines how Unity compiles the shader code. This allows you to [optimize the balance between build time, runtime, and file sizes](Keywords-concepts.html#keyword-impact-optimization).  The options are:  * **Shader Feature**: Unity only compiles shader variants for keyword combinations used by materials in your build, and removes other shader variants. * **Multi Compile**: Unity compiles shader variants for all keyword combinations regardless of whether the build uses these variants. * **Predefined**: Specifies that the target/sub-target already defines the keyword and you just want to reuse it. Predefined Keywords can either use a [built-in macro](https://docs.unity3d.com/Manual/shader-branching-built-in-macros.html), which results in static branching at build time, or any of the keywords already defined by the Shader Graph Target (for example, [URP](https://docs.unity3d.com/Manual/urp/urp-shaders/shader-keywords-macros.html)), including [Built-In keyword sets](https://docs.unity3d.com/Manual/SL-MultipleProgramVariants-shortcuts.html), and where the branching depends on that definition. * **Dynamic Branch**: Unity keeps branching code in one compiled shader program. |
| **Is Overridable** | Indicates whether the keyword's state can be overridden. For more information, refer to [Toggle shader keywords in a script](https://docs.unity3d.com/Manual/shader-keywords-scripts.html). |
| **Generate Material Property** | Generates a material property declaration to display the keyword as a property in the material inspector. This adds a `[Toggle(_KEYWORD)]` attribute to the material property. For more information, refer to [`MaterialPropertyDrawer`](https://docs.unity3d.com/ScriptReference/MaterialPropertyDrawer.html). |
| **Stages** | Set the stage the keyword applies to.  The following options are available:  * **All** - Applies this keyword to all shader stages.* **Vertex** - Applies this keyword to the vertex stage.* **Fragment** - Applies this keyword to the fragment stage. |

## Boolean keywords

Parameter specific to Boolean keywords in addition to the [common parameters](#common-parameters).

| **Name** | **Description** |
| --- | --- |
| **Default Value** | Enable this parameter to set the keyword's default state to on, and disable it to set the keyword's default state to off.  This parameter determines the value to use for the keyword when Shader Graph generates previews. It also defines the keyword's default value when you use this shader to create a new Material. |

## Enum keywords

Parameters specific to Enum keywords in addition to the [common parameters](#common-parameters).

| **Name** | **Description** |
| --- | --- |
| **Default Value** | Select an entry from the drop-down menu to determine which value to use for the keyword when Shader Graph generates previews. This also defines the keyword's default value when you use this shader to create a new Material. When you edit the Entries list, Shader Graph automatically updates the options in this control. |
| **Entries** | This list defines all the states for the keyword. Each state has a separate **Entry Name** and **Reference Suffix**.  * **Entry Name**: The name displayed in drop-down menus for the keyword on the [Internal Inspector](Internal-Inspector.html) and the Material Inspector. Shader Graph also uses this name for port labels on nodes that reference the keyword.* **Reference Suffix**: This identifies the final keyword, presented in the format `Reference_ReferenceSuffix`. |

When you define an Enum keyword, Shader Graph displays labels for each state consisting of a version of the Enum's **Entry Name** appended to the main **Reference** name.

##### Note

Special characters such as `(`, `)`, `!`, or `@` are not valid in the **Entry Name** of an Enum keyword. Shader Graph converts invalid characters to underscores (`_`).

## Built-in keywords

The parameters of built-in keywords depend on their type, which is always of either the Boolean or Enum type, and you cannot edit their values.

All Built-in keyword fields in the **Node Settings** tab of the [Graph Inspector](Internal-Inspector.html) are grayed out except for the **Default** field, which you can enable or disable to show the differences in Shader Graph previews. You also cannot expose Built-in keywords in the Material Inspector.

## Additional resources

* [Introduction to keywords in Shader Graph](Keywords-concepts.html)
* [Manage keywords in Shader Graph](Keywords-manage.html)