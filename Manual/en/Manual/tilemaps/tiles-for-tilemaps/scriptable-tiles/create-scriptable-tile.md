# Create a scriptable tile

The code example on this page demonstrates how to create your own scriptable tile, and how to use it in your project. The `PipelineExampleTile` scriptable tile is an example of a tile that can be used to layout linear segments onto the tilemap that automatically join up as you paint. This is very useful for designing tiles that are meant to be roads or pipes.

To create a scriptable tile for your project:

1. [Use a script to create the scriptable tile and its behavior](#script-tile)
2. [Paint with the scriptable tile](#paint-script)

![The Scene view with a result of the scriptable tile example in context, with pipes joined together.](../../../../uploads/Main/scriptable-tilemap-example-6000-1.png)

The Scene view with a result of the scriptable tile example in context, with pipes joined together.

## Prerequisite

You must have the [2D Tilemap Editor](com.unity.2d.tilemap.md) package installed before proceeding with the task. This package is part of the [2D feature set](../../../2DFeature.md) and is automatically installed if you select one of the [2D templates](https://docs.unity3d.com/hub/manual/project-create.html) when creating a new project. You can also manually install this package via [Unity’s Package Manager](../../../Packages.md).

## Use a script to create a scriptable tile

To create the `PipelineExampleTile` scriptable tile and have it as an available option in the UnityEditor’s **Asset** menu:

1. Create a blank MonoBehaviour script by going to **Assets** > **Create** > **MonoBehaviour Script**.
2. Name the file to `PipelineExampleTile.cs`.
3. Open the file in a text editor.
4. Replace the existing code with the following, and then save the file:

```
using System;

#if UNITY_EDITOR
using UnityEditor;
#endif

namespace UnityEngine.Tilemaps
{
   /// <summary>
   /// Pipeline Tiles are tiles which take into consideration its orthogonal neighboring tiles and displays a sprite depending on whether the neighboring tile is the same tile.
   /// </summary>
   [Serializable]
   public class PipelineExampleTile : TileBase
   {
       /// <summary>
       /// The Sprites used for defining the Pipeline.
       /// </summary>
       [SerializeField]
       public Sprite[] m_Sprites;

       /// <summary>
       /// This method is called when the tile is refreshed. The PipelineExampleTile will refresh all neighboring tiles to update their rendering data if they are the same tile.
       /// </summary>
       /// <param name="position">Position of the tile on the Tilemap.</param>
       /// <param name="tilemap">The Tilemap the tile is present on.</param>
       public override void RefreshTile(Vector3Int position, ITilemap tilemap)
       {
           for (int yd = -1; yd <= 1; yd++)
               for (int xd = -1; xd <= 1; xd++)
               {
                   Vector3Int pos = new Vector3Int(position.x + xd, position.y + yd, position.z);
                   if (TileValue(tilemap, pos))
                       tilemap.RefreshTile(pos);
               }
       }

       /// <summary>
       /// Retrieves any tile rendering data from the scripted tile.
       /// </summary>
       /// <param name="position">Position of the tile on the Tilemap.</param>
       /// <param name="tilemap">The Tilemap the tile is present on.</param>
       /// <param name="tileData">Data to render the tile.</param>
       public override void GetTileData(Vector3Int position, ITilemap tilemap, ref TileData tileData)
       {
           UpdateTile(position, tilemap, ref tileData);
       }

       /// <summary>
       /// Checks the orthogonal neighbouring positions of the tile and generates a mask based on whether the neighboring tiles are the same. The mask will determine the according Sprite and transform to be rendered at the given position. The Sprite and Transform is then filled into TileData for the Tilemap to use. The Flags lock the color and transform to the data provided by the tile. The ColliderType is set to the shape of the Sprite used.
       /// </summary>
       private void UpdateTile(Vector3Int position, ITilemap tilemap, ref TileData tileData)
       {
           tileData.transform = Matrix4x4.identity;
           tileData.color = Color.white;

           int mask = TileValue(tilemap, position + new Vector3Int(0, 1, 0)) ? 1 : 0;
           mask += TileValue(tilemap, position + new Vector3Int(1, 0, 0)) ? 2 : 0;
           mask += TileValue(tilemap, position + new Vector3Int(0, -1, 0)) ? 4 : 0;
           mask += TileValue(tilemap, position + new Vector3Int(-1, 0, 0)) ? 8 : 0;

           int index = GetIndex((byte)mask);
           if (index >= 0 && index < m_Sprites.Length && TileValue(tilemap, position))
           {
               tileData.sprite = m_Sprites[index];
               tileData.transform = GetTransform((byte)mask);
               tileData.flags = TileFlags.LockTransform | TileFlags.LockColor;
               tileData.colliderType = Tile.ColliderType.Sprite;
           }
       }

       /// <summary>
       /// Determines if the tile at the given position is the same tile as this.
       /// </summary>
       private bool TileValue(ITilemap tileMap, Vector3Int position)
       {
           TileBase tile = tileMap.GetTile(position);
           return (tile != null && tile == this);
       }

       /// <summary>
       /// Determines the index of the Sprite to be used based on the neighbour mask.
       /// </summary>
       private int GetIndex(byte mask)
       {
           switch (mask)
           {
               case 0: return 0;
               case 3:
               case 6:
               case 9:
               case 12: return 1;
               case 1:
               case 2:
               case 4:
               case 5:
               case 10:
               case 8: return 2;
               case 7:
               case 11:
               case 13:
               case 14: return 3;
               case 15: return 4;
           }
           return -1;
       }

       /// <summary>
       /// Determines the Transform to be used based on the neighbour mask.
       /// </summary>
       private Matrix4x4 GetTransform(byte mask)
       {
           switch (mask)
           {
               case 9:
               case 10:
               case 7:
               case 2:
               case 8:
                   return Matrix4x4.TRS(Vector3.zero, Quaternion.Euler(0f, 0f, -90f), Vector3.one);
               case 3:
               case 14:
                   return Matrix4x4.TRS(Vector3.zero, Quaternion.Euler(0f, 0f, -180f), Vector3.one);
               case 6:
               case 13:
                   return Matrix4x4.TRS(Vector3.zero, Quaternion.Euler(0f, 0f, -270f), Vector3.one);
           }
           return Matrix4x4.identity;
       }
   }
  
#if UNITY_EDITOR
   /// <summary>
   /// Custom Editor for a PipelineExampleTile. This is shown in the Inspector window when a PipelineExampleTile asset is selected.
   /// </summary>
   [CustomEditor(typeof(PipelineExampleTile))]
   public class PipelineExampleTileEditor : Editor
   {
       private PipelineExampleTile tile { get { return (target as PipelineExampleTile); } }

       public void OnEnable()
       {
           if (tile.m_Sprites == null || tile.m_Sprites.Length != 5)
               tile.m_Sprites = new Sprite[5];
       }

       /// <summary>
       /// Draws an Inspector for the PipelineExampleTile.
       /// </summary>
       public override void OnInspectorGUI()
       {
           EditorGUILayout.LabelField("Place sprites shown based on the number of tiles bordering it.");
           EditorGUILayout.Space();
          
           EditorGUI.BeginChangeCheck();
           tile.m_Sprites[0] = (Sprite) EditorGUILayout.ObjectField("None", tile.m_Sprites[0], typeof(Sprite), false, null);
           tile.m_Sprites[2] = (Sprite) EditorGUILayout.ObjectField("One", tile.m_Sprites[2], typeof(Sprite), false, null);
           tile.m_Sprites[1] = (Sprite) EditorGUILayout.ObjectField("Two", tile.m_Sprites[1], typeof(Sprite), false, null);
           tile.m_Sprites[3] = (Sprite) EditorGUILayout.ObjectField("Three", tile.m_Sprites[3], typeof(Sprite), false, null);
           tile.m_Sprites[4] = (Sprite) EditorGUILayout.ObjectField("Four", tile.m_Sprites[4], typeof(Sprite), false, null);
           if (EditorGUI.EndChangeCheck())
               EditorUtility.SetDirty(tile);
       }

       /// <summary>
       /// The following is a helper that adds a menu item to create a PipelineExampleTile Asset in the project.
       /// </summary>
       [MenuItem("Assets/Create/PipelineExampleTile")]
       public static void CreatePipelineExampleTile()
       {
           string path = EditorUtility.SaveFilePanelInProject("Save Pipeline Example Tile", "New Pipeline Example Tile", "Asset", "Save Pipeline Example Tile", "Assets");
           if (path == "")
               return;                           
           AssetDatabase.CreateAsset(ScriptableObject.CreateInstance<PipelineExampleTile>(), path);
        }
   }
#endif
}
```

Now wherever you need to use the scriptable tile you can create instances of your new class using `ScriptableObject.CreateInstance<YOUR_TILE_CLASS>()`. You can also convert this new instance to an Asset in the Editor in order to use it repeatedly by calling `AssetDatabase.CreateAsset()`.

## Paint with the scriptable tile

After [importing or saving](#script-tile) the `PipelineExampleTile.cs` script into your project, you will be able to create the **PipelineExampleTile** tile asset.

To paint with the **PipelineExampleTile** tile:

1. Create a **PipelineExampleTile** tile asset (menu: **Assets > Create > PipelineExampleTile**).
2. Select the created tile asset and go to its ****Inspector**** window.
3. Fill in the `PipelineExampleTile` with sprites according to the number of tiles bordering it. For example, the **sprite** for **One** has a single opening while the sprite for **Three** has three openings along the edges of the sprite. **Note:** When using your own sprites, it’s recommended to match the position and orientation shown in the following example:

   ![The scripted Pipeline Example Tile asset properties in the Inspector, with assigned sprites.](../../../../uploads/Main/scriptable-tilemap-example-6000-2.png)

   The scripted Pipeline Example Tile asset properties in the Inspector, with assigned sprites.
4. Save your project to save the changes made to the tile.
5. Add the tile to a [Tile Palette](../../tile-palettes/create-tile-palette.md) by dragging the tile asset from the **Project** window onto the Tile Palette in the [Tile Palette editor window](../../tile-palettes/tile-palette-editor-reference.md).
6. Use the [**Paint** tool](../../tile-palettes/tools/paint-tiles-with-paint-tool.md) with the scriptable tile to paint onto your [tilemap](../../work-with-tilemaps/create-tilemap.md).

## Additional resources

* [Tile Palette editor tools](../../tile-palettes/tools/tile-palette-tools-landing.md)