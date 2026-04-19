# Create a Scriptable Brush

Create a new class inheriting from `GridBrushBase` (or any useful subclasses of `GridBrushBase` like `GridBrush`). Override any required methods for your new Brush class. The following are the usual methods you would override:

* `Paint` allows the Brush to add items onto the target Grid.
* `Erase` allows the Brush to remove items from the target Grid.
* `FloodFill` allows the Brush to fill items onto the target Grid.
* `Rotate` rotates the items set in the Brush.
* `Flip` flips the items set in the Brush.

Create instances of your new class using `ScriptableObject.CreateInstance<(Your Brush Class>()`. You may convert this new instance to an Asset in the Editor in order to use it repeatedly by calling `AssetDatabase.CreateAsset()`.

You can also make a custom editor for your brush. This works the same way as custom editors for scriptable objects. The following are the main methods you would want to override when creating a custom editor:

* You can override `OnPaintInspectorGUI` to have an **Inspector** window appear on the Palette when the Brush is selected to provide additional behaviour when painting.
* You can also override `OnPaintSceneGUI` to add additional behaviours when painting on the `SceneView`.
* You can also override `validTargets` to have a custom list of targets which the Brush can interact with. This list of targets is shown as a dropdown list in the Palette window.

When created, the Scriptable Brush is listed in the **Brushes** dropdown menu in the Palette window. By default, an instance of the Scriptable Brush script is instantiated and stored in the *Library* folder of your project. Any modifications to the brush properties are stored in that instance. If you want to have multiple copies of that Brush with different properties, you can instantiate the Brush as Assets in your project. These Brush Assets are listed separately in the Brush dropdown menu.

You can add a `CustomGridBrush` attribute to your Scriptable Brush class. This allows you to configure the behavior of the Brush in the Palette window. The `CustomGridBrush` attribute has the following properties:

* `HideAssetInstances` - Set this to true to hide all copies of created Brush Assets in the Palette window. This is useful when you want only the default instance to show up in the Brush dropdown menu in the Tile Palette window.
* `HideDefaultInstances` - Set this to true to hide the default instance of the Brush in the Palette window. This is useful when you want only created Assets to show up in the Brush dropdown menu in the Tile Palette window.
* `DefaultBrush` - Set this to true to set the default instance of the Brush as the default Brush in the project. This makes this Brush the default selected Brush whenever the project starts. **Note:** Only set one Scriptable Brush as the Default Brush. Setting more than one Default Brush may cause your Scriptable Brushes to behave incorrectly.
* `DefaultName` - Set a name to this to have the Brush dropdown menu use the set name as the name for the Brush instead of the name of the Brush’s class.

If you want your Scriptable Brush class to use only certain tools, you can add a `BrushTools` attribute to your class with a list of compatible `TilemapEditorTools` types. This ensures that your Scriptable Brush only activates with these specific tools from the Tile Palette **toolbar**.

## Scriptable Brush example

![The Scene view with a result of the scriptable LineBrush example in context. Use the brush to draw tiles along a line across the tilemap.](../../../../uploads/Main/Tilemap-ScriptableBrushes-Example-32.png)

The Scene view with a result of the scriptable LineBrush example in context. Use the brush to draw tiles along a line across the tilemap.

`LineBrush` provides the ability to easily draw lines of Tiles onto the **Tilemap** by specifying the start point and the end point. The Paint method for the LineBrush is overridden to allow the user to specify the start of a line with the first mouse click in Paint mode and draw the line with the second mouse click in Paint mode. The `OnPaintSceneGUI` method is overridden to produce the preview of the line to be drawn between the first and second mouse clicks. The following is a script used to create the Brush.

```
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Tilemaps;

namespace UnityEditor.Tilemaps
{
    [CustomGridBrush(true, false, false, "Line Brush")]
    public class LineBrush : GridBrush {
        public bool lineStartActive = false;
        public Vector3Int lineStart = Vector3Int.zero;
        public override void Paint(GridLayout grid, GameObject brushTarget, Vector3Int position)
        {
            if (lineStartActive)
            {
                Vector2Int startPos = new Vector2Int(lineStart.x, lineStart.y);
                Vector2Int endPos = new Vector2Int(position.x, position.y);
                if (startPos == endPos)
                    base.Paint(grid, brushTarget, position);    
                else
                {
                    foreach (var point in GetPointsOnLine(startPos, endPos))
                    {
                        Vector3Int paintPos = new Vector3Int(point.x, point.y, position.z);
                        base.Paint(grid, brushTarget, paintPos);
                    }
                }
                lineStartActive = false;
            }
            else
            {
                lineStart = position;
                lineStartActive = true;
            }
        }
        [MenuItem("Assets/Create/Line Brush")]
        public static void CreateBrush()
        {
            string path = EditorUtility.SaveFilePanelInProject("Save Line Brush", "New Line Brush", "Asset", "Save Line Brush", "Assets");
            if (path == "")
                return;
AssetDatabase.CreateAsset(ScriptableObject.CreateInstance<LineBrush>(), path);
        }
        // http://ericw.ca/notes/bresenhams-line-algorithm-in-csharp.html
        public static IEnumerable<Vector2Int> GetPointsOnLine(Vector2Int p1, Vector2Int p2)
        {
            int x0 = p1.x;
            int y0 = p1.y;
            int x1 = p2.x;
            int y1 = p2.y;
            bool steep = Math.Abs(y1 - y0) > Math.Abs(x1 - x0);
            if (steep)
            {
                int t;
                t = x0; // swap x0 and y0
                x0 = y0;
                y0 = t;
                t = x1; // swap x1 and y1
                x1 = y1;
                y1 = t;
            }
            if (x0 > x1)
            {
                int t;
                t = x0; // swap x0 and x1
                x0 = x1;
                x1 = t;
                t = y0; // swap y0 and y1
                y0 = y1;
                y1 = t;
            }
            int dx = x1 - x0;
            int dy = Math.Abs(y1 - y0);
            int error = dx / 2;
            int ystep = (y0 < y1) ? 1 : -1;
            int y = y0;
            for (int x = x0; x <= x1; x++)
            {
                yield return new Vector2Int((steep ? y : x), (steep ? x : y));
                error = error - dy;
                if (error < 0)
                {
                    y += ystep;
                    error += dx;
                }
            }
            yield break;
        }
    }
    [CustomEditor(typeof(LineBrush))]
    public class LineBrushEditor : GridBrushEditor
    {
        private LineBrush lineBrush { get { return target as LineBrush; } }
        public override void OnPaintSceneGUI(GridLayout grid, GameObject brushTarget, BoundsInt position, GridBrushBase.Tool tool, bool executing)
        {
            base.OnPaintSceneGUI(grid, brushTarget, position, tool, executing);
            if (lineBrush.lineStartActive)
            {
                Tilemap tilemap = brushTarget.GetComponent<Tilemap>();
                if (tilemap != null)
                    tilemap.ClearAllEditorPreviewTiles();
                // Draw preview tiles for tilemap
                Vector2Int startPos = new Vector2Int(lineBrush.lineStart.x, lineBrush.lineStart.y);
                Vector2Int endPos = new Vector2Int(position.x, position.y);
                if (startPos == endPos)
                    PaintPreview(grid, brushTarget, position.min);
                else
                {
                    foreach (var point in LineBrush.GetPointsOnLine(startPos, endPos))
                    {
                        Vector3Int paintPos = new Vector3Int(point.x, point.y, position.z);
                        PaintPreview(grid, brushTarget, paintPos);
                    }
                }
                if (Event.current.type == EventType.Repaint)
                {
                    var min = lineBrush.lineStart;
                    var max = lineBrush.lineStart + position.size;
                    // Draws a box on the picked starting position
                    GL.PushMatrix();
                    GL.MultMatrix(GUI.matrix);
                    GL.Begin(GL.LINES);
                    Handles.color = Color.blue;
                    Handles.DrawLine(new Vector3(min.x, min.y, min.z), new Vector3(max.x, min.y, min.z));
                    Handles.DrawLine(new Vector3(max.x, min.y, min.z), new Vector3(max.x, max.y, min.z));
                    Handles.DrawLine(new Vector3(max.x, max.y, min.z), new Vector3(min.x, max.y, min.z));
                    Handles.DrawLine(new Vector3(min.x, max.y, min.z), new Vector3(min.x, min.y, min.z));
                    GL.End();
                    GL.PopMatrix();
                }
            }
        }
    }
}
```