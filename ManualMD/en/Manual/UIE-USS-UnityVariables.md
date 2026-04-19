# Introduction to USS built-in variables

USS built-in [variables](UIE-USS-CustomProperties.md) specify default values for Editor and runtime UI. You can use these variables in your own USS to match your custom interfaces with Unity style.

The name of each built-in variable indicates how and where the variable is used. A variable name consists of one or more parts, separated by hyphens. Each part consists of one or more words separated by underscores.

`--unity-{group}-{role_and_control}-{sub_element}-{pseudo_state_sequence}`

Each part of the name shows the types of USS rules that use the variable.

* [**Group**](#group): The kind of data the variable represents.
* [**Role/Control**](#role-and-control): A conceptual grouping for the elements the variable affects.
* [**Sub-Element**](#sub-element): An element or control the variable affects.
* [**Pseudo States**](#pseudo-state-sequence): Lists the states Unity uses the variable for.

For example, the following variable name:

`--unity-colors-button-text-hover`

Provides the following information about how Unity uses the variable:

|  | **Value** | **Meaning** |
| --- | --- | --- |
| **Group** | `colors` | Represents color data. |
| **Role/Control** | `button` | Affects buttons. |
| **Sub-Element** | `text` | Affects text. Its `group` is `colors`, so it affects text color. |
| **Pseudo-States** | `hover` | Applies to elements when the mouse pointer hovers over them. |

This USS built-in variable changes the color of button text when a user hovers over the button.

## Group

The group specifies what kind of data the variable represents. Each group has several possible [sub-elements](#sub-element).

The variable names have the following groups:

| **Group** | **Used for** |
| --- | --- |
| `colors` | Color properties, such as `background-color` and `border-color` |
| `metrics` | Properties that control dimensions and shapes. For example, `border-radius`, `border-width`, `margin`, and `padding` |
| `icons` | Standard Unity icon images |

## Role and Control

Roles and controls are two ways of grouping elements conceptually.

* **Role** refers to a group of elements that have the same purpose, regardless of what type each element is. For example, the `error` role includes all elements that display error messages to users.
* **Control** refers to a group of elements of the same type, regardless of what they do. For example, `buttons` includes all buttons in the Editor.

Each variable has either a role or a control.

The variable names have the following roles and controls:

### Roles

| **Role** | **Description** |
| --- | --- |
| `default` | Default style settings such as text color, background, and margin. |
| `alternated_rows` | Elements that display tabular data with alternating row colors. For example, list items in a list view. |
| `error` | Elements that communicate error states to users. |
| `highlight` | Highlighted parts of the UI. For example, text selections, or selected items in tree views. |
| `link` | Parts of the UI (typically text) that are clickable links. This is the un-clicked state. |
| `play_mode` | Elements displayed when the Editor is in Play mode. |
| `visited_link` | Parts of the UI, typically text, that are clickable links that a user has already clicked. |
| `warning` | Elements that communicate a warning message to users. |

### Controls

| **Control** | **Description** |
| --- | --- |
| `app_toolbar` | The main Unity **toolbar** |
| `app_toolbar_button` | Buttons in the main Unity toolbar |
| `box` | Boxes used to group elements in the Editor UI |
| `button` | Buttons in the UI, except for toolbars |
| `dropdown` | Dropdown lists or menus |
| `helpbox` | Boxes used to display help information |
| `input_field` | Fields used to input text or numeric values |
| `label` | Text labels in the Editor UI |
| `object_field` | Fields used for object values. For example, property values for a **GameObject** or Asset. |
| `popup` | Popup menu and other popup controls |
| `preview` | Views used to display previews. For example, of Assets such as meshes and textures. |
| `scrollbar_groove` | The background element of a scrollbar in which users drag the scrollbar thumb |
| `scrollbar_thumb` | The draggable handle element in a scrollbar |
| `slider_groove` | The background element of a slider in which users drag the slider thumb |
| `slider_thumb` | The draggable handle element in a slider |
| `slider_thumb_halo` | An overlay displayed around the slider thumb when the user drags it |
| `tab` | Tab items in tab controls |
| `toolbar` | Any Editor toolbar except for the main Unity toolbar (`app_toolbar`) |
| `toolbar_button` | Buttons in an Editor toolbar |
| `window` | Editor windows |

## Sub-Element

The sub-element is the part of an element that the variable affects. Together with a variable’s `group`, the sub-element shows what kind of data the variable represents.

For example, when you see a variable name with the `colors` group and the `text` sub-element, it means Unity uses the variable in styles that affect text color.

The variable names have the following sub-elements:

### Colors group

| **Sub-element** | **Description** |
| --- | --- |
| `background` | An element’s background color |
| `border` | An element’s outer border color |
| `border_accent` | An element’s inner border color. For example, **Inspector** windows have a two-tone border |
| `text` | Text color for elements that display text |

### Metrics group

| **Sub-element** | **Description** |
| --- | --- |
| `margin_{left, top, right, bottom}` | An element’s margin values |
| `padding_{left, top, right, bottom}` | An element’s padding values |
| `border_{left, top, right, bottom}_width` | An element’s border width values |
| `border_{left_top, left_bottom, right_top, right_bottom}_radius` | An element’s border radius values values |
| `width, height` | An element’s width and height values |

## Pseudo-States

The pseudo-state sequence is a list of UI states that Unity uses a variable for.

For example, when you see a variable name with the `hover` pseudo state, it means Unity uses the variable in styles that affect elements when a user hovers the pointer over them.

For example: `--unity-colors-toolbar_button-text-hover`.

A variable name can have more than one pseudo-state. Multiple pseudo states appear in alphabetical order, separated by underscores `_`.

For example: `--unity-colors-toolbar_button-text-focus_selected`.

Unity variable names can have any combination of the following pseudo-states:

| **Pseudo-state** | **Description** |
| --- | --- |
| (none) | Normal state |
| `checked` | A checkbox-type control is checked |
| `disabled` | A control is disabled |
| `focus` | A control has focus |
| `hover` | A user hovers over a control |
| `inactive` | A control doesn’t have focus |
| `pressed` | A control is pressed |
| `selected` | A control is selected |

## Additional resources

* [Create USS variables](UIE-USS-CustomProperties.md)
* [USS built-in variable reference](UIE-uss-built-in-variable-reference.md)