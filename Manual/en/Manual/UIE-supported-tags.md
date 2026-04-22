# Supported rich text tags

The following table lists all the supported rich text tags. [Advance Text Generator(ATG)](UIE-advanced-text-generator.md) doesn’t support tags marked with an asterisk (\*).

| **Tag** | **Description** | **Example** |
| --- | --- | --- |
| `<a>` | Define a hyperlink. Use the `href` attribute to define the hyperlink URL. **Note**: This tag has a fixed limit of 256 characters. | `<a href="https://www.unity.com">Visit Unity!</a>` |
| `<align>` | Change the text’s horizontal alignment. The supported values are:  * `left` * `center` * `right` * `justified` * `flush`(not supported by ATG)    If you put multiple alignment tags on the same line, the last one overrides the others. | `<align="left">Left-aligned</align>` |
| `<allcaps>`\* | Convert text to uppercase. | `<allcaps>Alice and Bob watched TV.</allcaps>` |
| `<alpha>`\* | Change text opacity. It works with hexadecimal values. | `<alpha=#FF>FF <alpha=#CC>CC <alpha=#AA>AA <alpha=#88>88 <alpha=#66>66 <alpha=#44>44 <alpha=#22>22 <alpha=#00>00` |
| `<b>` | Render text in boldface. | `The fox jumps over the <b>lazy dog</b>` |
| `<br>` | Forces a line break in text. | `Break the line here <br> New line starts` |
| `<color>` | Change text color or color and opacity. It supports color names and hexadecimal values. If you apply successive tags in the same text, the last one takes precedence over the others until you either add another tag or use a closing tag to end the current color’s scope. | `<color="red">Red <color=#005500>Dark Green <#0000FF>Blue <color=#FF000088>Semitransparent Red` |
| `<cspace>` | Change spacing between characters, either absolute or relative to the original font Asset. Use pixels or font units. Positive adjustments push the characters apart, negative adjustments pull them together. | `<cspace=1em>Spacing</cspace> is just as important as <cspace=-0.5em>timing.` |
| `<font>*` | Change text font. **Note**: You must place the font asset in the path defined by **Default Font Asset** in the [UITK Text Setting](UIE-text-setting-asset.md) asset. The default path is `Resources/Fonts & Materials/`. | `Would you like <font="Impact SDF">a different font?</font>` |
| `<font-weight>`\* | Change the text’s font weight to any of the weights defined in the [font Asset](UIE-font-asset-properties.md). If you haven’t defined any font weights, you can still use `400` for normal, `700` for bold. | `<font-weight="100">Thin</font-weight>` |
| `<gradient>*` | Apply a [color gradient](UIE-color-gradient.md) to text. | `<gradient="Light to Dark Green - Vertical">gradient` |
| `<i>` | Render text in italics. | `The <i>quick brown fox</i>` |
| `<indent>`\* | Set an absolute horizontal position for the text that follows. When encountered, the text cursor jumps to the specified position regardless of where the preceding text ended. The indentation applies to all text between the tag and the next hard line break. Use this tag to create text patterns, such as bullet points, that work with word-wrapping. Specify indentation in pixels, font units, or percentages. | `<indent=15%>It is useful for things like bullet points.</indent>` |
| `<line-height>`\* | Modify the line height relative to the default line height specified in the font Asset. Specify the line height in pixels, font units, or percentages. | `<line-height=50%>Line height at 50%` |
| `<line-indent>`\* | Indent the first line after every hard line break. New lines created by word-wrapping are not indented. | `<line-indent=15%>This line is indented. <br>This line is also indented.` |
| `<lowercase>`\* | Convert text to lowercase. | `<lowercase>Alice and Bob watched TV.</lowercase>` |
| `<margin>`\* | Set the text horizontal margins. If you only want to adjust the left or right margin, you can use the `<margin-left>` or `<margin-right>` tag. Specify the margins in pixels, font units, and percentages. Negative values have no effect. | `<margin=5em>` |
| `<mark>` | Highlight the text with a colored overlay. The overlay must be translucent (alpha less than 1) for the text to show through. | `Text <mark=#ffff00aa>can be marked with</mark> an overlay.` |
| `<mspace>` | Override a font’s character spacing and turn it into a monospace font. | `Any font can become <mspace=2.75em>monospace, if you really want it.` |
| `<nobr>`\* | Keep a segment of text together. | `You don't want <nobr>I M P O R T A N T</nobr> things to be broken up.` |
| `<noparse>` | Prevent parsing of rich text tags. | `Use <noparse><b></noparse> for <b>bold</b> text.` |
| `<pos>`\* | Set the horizontal caret position on the current line. Specify the horizontal position in pixels, font units, or percentages. | `at <pos=75%>75%` |
| `<rotate>`\* | Rotates each character about its center. Specify the amount of rotation in degrees. Positive values rotate characters counter-clockwise. Negative values rotate them clockwise.  Rotation affects the spacing between characters, and might cause characters to overlap in some cases. Use the `<cspace>` tag to correct character spacing as needed. | `Rotate text <rotate="45">counter-clockwise</rotate>` |
| `<s>` | Render a line across the text. | `The <s>quick brown</s> fox` |
| `<size>`\* | Adjusts the font size. Specify the new size in pixels, font units, or percentage. Pixel adjustments can be absolute (such as `5px`) or relative (such as `+1` or `-1`). Relative sizes are based on the original font size, so they’re not cumulative. | `<size=100%>Echo <size=80%>Echo <size=60%>Echo <size=40%>Echo <size=20%>Echo` |
| `<smallcaps>`\* | Convert text to uppercase characters in a small caps style. | `<smallcaps>Alice and Bob watched TV.` |
| `<space>`\* | Add a horizontal offset between itself and the rest of the text. Specify the offset in pixels or font units. | `Give me some <space=5em> space` |
| `<sprite>` | Add a [sprite](UIE-sprite.md) from a sprite asset into your text. | `<sprite name="spriteName">` |
| `<style>`\* | Apply a [custom style](UIE-style-sheet.md) to the text. | `<style="H1">Styles</style>` |
| `<sub>`\* | Convert the text to subscript. | `We have 1m<sup>3</sup> of H<sub>2</sub>O.` |
| `<sup>`\* | Convert the text to superscript. | `We have 1m<sup>3</sup> of H<sub>2</sub>O.` |
| `<u>` | Underline the text. | `<u>The lazy dog</u>` |
| `<uppercase>`\* | Convert text to uppercase. | `<uppercase>Alice and Bob watched TV.</uppercase>` |
| `<voffset>`\* | Give the baseline a vertical offset. Specify the offset in pixels or font units. The offset is always relative to the original baseline. | `Up <voffset=1em>up <voffset=2em>UP</voffset> and <voffset=-0.5em>down</voffset> we go again.` |
| `<width>`\* | Change the horizontal size of text area. | `<width=60%>Those days are long gone</width>` |
| `<link>` | Specify a link ID for a text segment. **Note**: This tag has a fixed limit of 256 characters. | `<link="ID">my link</link>` |

## Additional resources

* [Style text with rich text tags](UIE-rich-text-tags.md)
* [Add hyperlinks in text](ui-systems/add-hyperlinks-in-text.md)