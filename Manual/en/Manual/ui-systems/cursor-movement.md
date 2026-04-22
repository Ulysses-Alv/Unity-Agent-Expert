# Cursor movement

This page explains how cursor movement behaves in the input fields when handling bidirectional (BIDI) text.

## Logical Cursor Movement

Unity currently follows a **Logical Cursor Movement** approach. This means the cursor moves through BIDI text based on the direction of the segment of text. For example, using the left arrow key in a sentence with Arabic and English text, it moves right-to-left through Arabic, then jumps at the leftmost character in the English segment and continues left-to-right until it reaches the end of the segment.

![Logical Cursor Movement Example](../../uploads/Main/font/UIE-CursorMovement.gif)

Logical Cursor Movement Example

## Visual Cursor Movement

Some applications follow a **Visual Cursor Movement** approach. This means the cursor moves to the next visual character regardless of the direction of the text, which can sometimes feel more intuitive for users.

## Additional resources

* [Set language direction](ui-systems/set-language-direction)