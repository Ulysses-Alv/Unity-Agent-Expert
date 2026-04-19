# Introduction to rich text tags

Rich text tags are tags that you can place inside a text string to style the text between the tags.

**Note**: In the current release, rich text tags aren’t supported for [TextField](../UIE-uxml-element-TextField.md).

## Rich text syntax

Rich text tags are similar to HTML or XML tags, but have less strict syntax.

A simple tag can have be just its name and have no additional values or attributes. For example, the `<b>` tag makes text bold.

Some tags have additional values or attributes like this:

* `<tag="value">`
* `<tag attribute="value">`

For example:

* `<color=”red”>`: Makes text red
* `<sprite index=3>`: Inserts the fourth sprite from the default Sprite Asset.

**Note**: In a UXML file, you must use HTML code for the following characters:

* `<`: `(&lt;)`
* `>`: `(&gt;)`
* `"`: `(&quot;)`

The following table lists possible attribute value types and example values.

| **Value type** | **Example value** |
| --- | --- |
| Decimals | `0.5` |
| Percentages | `25%` |
| **Pixel** values | `5px` |
| Font units | `1.5em` |
| Hex color values | `#FFFFFF` (RGB) `#FFFFFFFF` (RGBA) `#FF` (A) |
| Names | Both `<link=”ID”>` and `<link=ID>` are valid. |

## Tag scope and nested tags

Tags have a scope that defines how much of the text they affect. Most tags affect all text from the point they appear onward.

For example, if you add the tag `<color="red">` to the beginning of the text, it affects the entire text block: `<color="red">This text is red`.

If you add the same tag in the middle of the text block, it affects only the text between the tag and the end of the block : `This text turns<color="red"> red`.

If you use the same tag more than once in a text block, the last tag supersedes all previous tags of the same type: `<color="red">This text goes from red<color="green"> to green`.

You can also use a closing tag to limit the scope of a tag, and use nested tag within another tag: `<color=red>This text is <color=green>mostly </color>red`.

The first `<color>` tag’s scope is the entire text block. The the second `<color>` tag has a closing tag that limits its scope to one word.

When you nest tags, you don’t need to close their scopes in the same order that you started them.

## Enable and disable rich text tags

Rich text tags are enabled by default.

To disable the rich text tag, do one of the following:

* In UI Builder, select the control and clear the **Enable Rich Text** checkbox in the **Inspector** window.
* In UXML, set the `enable-rich-text` attribute to `false`.

## Additional resources

* [Supported rich text tags](../UIE-supported-tags.md)
* [Add hyperlinks in text](add-hyperlinks-in-text.md)