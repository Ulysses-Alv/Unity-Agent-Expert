# World Space UI size and position examples

In World Space UI, you can position, rotate, and scale UI elements like any other 2D or **3D object** in the **Scene**.

Key considerations:

* How is the layout or panel size determined?
* How does the chosen size mode affect pivot placement?
* How does the pivot reference size impact positioning?

The following examples demonstrate how these settings impact UI behavior in real-world scenarios.

## Example: Tutorial and hint notifications

As players progress through the game, they might receive tutorial and hint notifications, such as [this example](https://www.gameuidatabase.com/index.php?&set=1&tag=122&scrn=904&autoload=66060). These notifications are typically appear in front of the player, and their size might vary based on the content.

The following table describes the suggested settings:

| Setting | Description |
| --- | --- |
| **Size Mode**: **Dynamic** | The menu adjusts its height based on the number of available options. |
| **Pivot Reference Size**: **Layout** | Uses only the child document’s layout, ignoring positioned elements like pips or badges. This prevents occasional UI elements (for example, a badge in the menu’s corner) from shifting the entire layout. To include these elements in the size calculation, use **Bounding Box** instead. |
| **Pivot**: **Left Center** | Keeps the UI positioned to the right of the player and ensures vertical expansion occurs from the center, avoiding unexpected shifts. |

## Example: Interactive terminal

Another example is an interactive terminal, like [this one](https://www.gameuidatabase.com/index.php?&set=1&tag=122&scrn=904&autoload=13411). These terminals typically appear in front of the player, and their size remains fixed, regardless of the content.

The following table describes the suggested settings:

| Setting | Description |
| --- | --- |
| **Size Mode**: **Fixed** | Keeps the UI within a predefined screen size, ensuring consistency regardless of content length or complexity. |
| **Pivot Reference Size**: **Layout** | Uses only the child document’s layout for size calculations, ignoring elements outside the layout to maintain consistent alignment. |
| **Pivot**: **Center** | Centers the UI element based on its in-world position. Alternatively, you can use a top-left pivot, but with a fixed size, the result remains consistent. |

### Example: Map marker

A map marker, [like this one](https://www.gameuidatabase.com/index.php?&set=1&tag=136&scrn=904&autoload=42856), is another example of a World Space UI element. Its size and position might change based on its content.

The following table describes the suggested settings:

| Setting | Description |
| --- | --- |
| **Size Mode**: **Dynamic** | The tooltip or marker resizes dynamically to fit varying lengths of text, additional information, or other content. |
| **Pivot Reference Size**: **Bounding Box** | Ensures that **visual elements** extending beyond the regular layout, such as badges, background glows, drop shadows, or special effects, are included in the container size calculation. This maintains consistent alignment even when elements like duration badges affect the container size. |
| **Pivot**: **Left Center** | Ensures the map tooltip appears from a fixed point while allowing content to expand upward and away from the in-world elements the player interacts with. |

## Additional resources

* [World Space UI](world-space-ui.md)
* [Panel Input Configuration properties](world-space-panel-input-configuration.md)
* [Create a World Space UI](create-world-space-ui.md)