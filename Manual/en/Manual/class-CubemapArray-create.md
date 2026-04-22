# Create a cubemap array

A **cubemap** array is an array of cubemaps that are the same size and format, and that the GPU can access as a single texture resource. Cubemap arrays are often used for implementing efficient **reflection probe**, lighting and shadowing systems.

To create a cubemap array in your Project, you must use a script.

The following example is an Editor script that creates an instance of the `CubemapArray` class, populates it with color data, and then saves it to your Project as a Texture Asset.

```
using UnityEngine;
using UnityEditor;

public class CreateCubeArrayTexture : MonoBehaviour
{
    [UnityEditor.MenuItem("CreateExamples/CubemapArray")]
    static void CreateCubemapArray()
    {
        // Configure the cubemap array and color data
        int faceSize = 16;
        int arraySize = 4;
        int[] kCubeXRemap = new int[] { 2, 2, 0, 0, 0, 0 };
        int[] kCubeYRemap = new int[] { 1, 1, 2, 2, 1, 1 };
        int[] kCubeZRemap = new int[] { 0, 0, 1, 1, 2, 2 };
        float[] kCubeXSign = new float[] { -1.0F, 1.0F, 1.0F, 1.0F, 1.0F, -1.0F };
        float[] kCubeYSign = new float[] { -1.0F, -1.0F, 1.0F, -1.0F, -1.0F, -1.0F };
        float[] kCubeZSign = new float[] { 1.0F, -1.0F, 1.0F, -1.0F, 1.0F, -1.0F };
        var baseCols = new Color[] { Color.white, new Color(1, .5f, .5f, 1), new Color(.5f, 1, .5f, 1), new Color(.5f, .5f, 1, 1), Color.gray };
        
        // Create an instance of CubemapArray
        var tex = new CubemapArray(faceSize, arraySize, TextureFormat.ARGB32, true);
        tex.filterMode = FilterMode.Trilinear;
        
        // Iterate over each cubemap
        var col = new Color[tex.width * tex.width];
        float invSize = 1.0f / tex.width;
        for (var i = 0; i < tex.cubemapCount; ++i)
        {
            var baseCol = baseCols[i % baseCols.Length];

            // Iterate over each face of the current cubemap
            for (var face = 0; face < 6; ++face)
            {
                var idx = 0;
                Vector3 signScale = new Vector3(kCubeXSign[face], kCubeYSign[face], kCubeZSign[face]);
                
                // Iterate over each pixel of the current face
                for (int y = 0; y < tex.width; ++y)
                {
                    for (int x = 0; x < tex.width; ++x)
                    {
                        // Calculate a "normal direction" color for the current pixel
                        Vector3 uvDir = new Vector3(x * invSize * 2.0f - 1.0f, y * invSize * 2.0f - 1.0f, 1.0f);
                        uvDir = uvDir.normalized;
                        uvDir.Scale(signScale);
                        Vector3 dir = Vector3.zero;
                        dir[kCubeXRemap[face]] = uvDir[0];
                        dir[kCubeYRemap[face]] = uvDir[1];
                        dir[kCubeZRemap[face]] = uvDir[2];

                        // Shift the color into the 0.4..1.0 range
                        Color c = new Color(dir.x * 0.3f + 0.7f, dir.y * 0.3f + 0.7f, dir.z * 0.3f + 0.7f, 1.0f);
                        
                        // Add a pattern to some pixels, so that mipmaps are more clearly visible
                        if (((x ^ y) & 3) == 1)
                            c *= 0.5f;
                        
                        // Tint the color with the baseCol tint
                        col[idx] = baseCol * c;
                        ++idx;
                    }
                }

                // Copy the color values for this face to the texture
                tex.SetPixels(col, (CubemapFace)face, i);
            }
        }

        // Apply the changes to the texture and upload the updated texture to the GPU
        tex.Apply();        

        // Save the texture to your Unity Project
        AssetDatabase.CreateAsset(tex, "Assets/ExampleCubemapArray.asset");
        UnityEditor.AssetDatabase.SaveAssets();
    }
}
```