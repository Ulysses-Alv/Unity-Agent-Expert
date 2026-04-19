# Modify Tilemap reference

The **gizmo** **toolbar** in this section has different gizmos you can use to change the **tilemap** and its contents. You can select different options and behaviors for inserting or removing rows and columns of blank cells into the tilemap from the [dropdown menu](#dropdown).

## Gizmo toolbar

Select a gizmo from the toolbar to activate as specific gizmo to change the selected contents in the tilemap. The following table describes each option, with links to examples showing how they affect the tilemap.

| **Gizmo** | **Function** |
| --- | --- |
| [None](#None) | No gizmo is active or shown in the **Scene** view. |
| [Move](#Move) | Activates and displays a **Move** gizmo in the **Scene view**. Use this to change the offset of the selected contents. |
| [Rotate](#Rotate) | Activates and displays a **Rotate** gizmo in the Scene view. Use this to change the rotation of the selected contents. |
| [Scale](#Scale) | Activates and displays a **Scale** gizmo in the Scene view. Use this to change the scale of the selected contents. |
| [Transform](#Transform) | Activates and displays a **Transform** gizmo in the Scene view. Use this to change the offset, rotation and scale of the selected contents all at once. |

### Gizmo function examples

**None**

![No gizmo](../../../../../uploads/Main/None_ex.png)  
Default Tilemap and selected cell location. No Gizmo is activated or visible.

**Move**

![Move](../../../../../uploads/Main/Move_ex.png)  
**Left:** Default Tilemap and selected cell location. **Right:** Offset changed for the selected cell locations.

**Rotate**

![Rotate](../../../../../uploads/Main/Rotate_ex.png)  
**Left:** Default Tilemap and selected cell location. **Right:** Rotation changed for the selected cell locations.

**Scale**

![Scale](../../../../../uploads/Main/Scale_ex.png)  
**Left**: Default Tilemap and selected cell location. **Right**: Scale changed for the selected cell locations.

**Transform**

![Transform](../../../../../uploads/Main/Transform_ex.png)  
**Left:** Default Tilemap and selected cell location. **Right:** Offset, rotation and scale of the selected cell locations are modified.

## Cell rows and columns options

The dropdown menu provides different options for inserting or removing rows and columns of blank cells onto the tilemap. After selecting one of the dropdown menu options, enter the number of rows or columns to insert or remove into the box and select **Modify**.

The following table describes each option, with links to examples showing how they affect the tilemap.

| **Property** | **Function** |
| --- | --- |
| [Insert Row](#InsertRow) | Inserts one or more rows of blank cells at the selected location. Existing cells are displaced upward along the positive y-axis. |
| [Insert Row Before](#InsertRowB4) | Inserts one or more rows of blank cells below the selected location. Existing cells are displaced downward along the negative y-axis. |
| [Delete Row](#DelRow) | Removes one or more rows of cells at the selected location and above. Existing cells above then collapse down to fill the space left by the deleted rows. |
| [Delete Row Before](#DelRowB4) | Removes one or more rows of cells below the selected location. Existing cells below then shift upward along the positive y-axis to fill the space left by the deleted rows. |
| [Insert Column](#InsertCol) | Inserts one or more columns of blank cells at the selected location. Existing cells are displaced to the right along the positive x-axis. |
| [Insert Column Before](#InsertColB4) | Inserts one or more columns of blank cells to the left of the selected cell. Existing cells are displaced to the left along the negative x-axis. |
| [Delete Column](#DelCol) | Removes one or more columns of cells at the selected location and to its right. Existing cells then shift to the left along the negative x-axis to fill the space left by the deleted columns. |
| [Delete Column Before](#DelCol) | Removes one or more columns of cells to the left of the selected cell. Existing cells shifted to the right along the positive x-axis to fill the space left by the deleted columns. |

### Examples of different dropdown menu options

**Insert Row**

![Insert Row](../../../../../uploads/Main/InsertRow_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Insert Row.*

**Insert Row Before**

![Insert Row Before](../../../../../uploads/Main/InsertRowB4_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Insert Row Before.*

**Delete Row**

![Delete Row](../../../../../uploads/Main/DelRow_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Delete Row.*

**Delete Row Before**

![Delete Row Before](../../../../../uploads/Main/DelRowB4_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Delete Row Before.*

**Insert Column**

![Insert Column](../../../../../uploads/Main/InsCol_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Insert Column.*

**Insert Column Before**

![Insert Column Before](../../../../../uploads/Main/InsColB4_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Insert Column Before.*

**Delete Column**

![Delete Column](../../../../../uploads/Main/DelCol_ex.png)**Left:** *Default Tilemap and selected cell location.* **Right:** *Delete Column.*

**Delete Column Before**

![Delete Column Before](../../../../../uploads/Main/DelColB4_ex.png)  
**Left:** *Default Tilemap and selected cell location.* **Right:** *Delete Column Before.*

### Multiple cell selection

When multiple cells are selected, the lower-leftmost cell is the main point of reference when applying the Modify Tilemap options. See the following examples of selecting multiple cells then modifying the Tilemap.

![Insert Row with multiple selected cells](../../../../../uploads/Main/InsertRow_mult_ex.png)  
**Left:** *Default Tilemap with multiple cells selected.* **Right:** *Insert Row.*

![Insert Column with multiple selected cells](../../../../../uploads/Main/InsertCol_mult_ex.png)  
**Left:** *Default Tilemap with multiple cells selected.* **Right:** *Insert Column.*

![Delete Row with multiple selected cells](../../../../../uploads/Main/DelRow_mult_ex.png)  
**Left:** *Default Tilemap with multiple cells selected.* **Right:** *Delete Row.*

![Delete Column with multiple selected cells](../../../../../uploads/Main/DelCol_mult_ex.png)  
**Left:** *Default Tilemap with multiple cells selected.* **Right:** *Delete Column.*

## Additional resources

* [Grid Selection properties](./grid-selection-properties-reference.md)
* [Tile Palette editor reference](../../tile-palette-editor-reference.md)