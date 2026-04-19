# Configure mesh compression

To use the **Mesh** **Compression** setting:

1. Select a Model in your project’s Assets folder to open the Model tab of the Import Settings window.
2. Navigate to the Meshes heading and find the Mesh Compression setting.
3. Select the dropdown menu to choose a level for all meshes in that Model. You can also change this setting in code with the ModelImporterMeshCompression enumeration.

Available values are High, Medium, Low, or Off. The following table shows typical compression ratios for each of these settings:

| **Value** | **Vertices** | **Normals** | **Tangents** | **UVs** | **Color** |
| --- | --- | --- | --- | --- | --- |
| **Off** | 1.0 | 1.0 | 1.0 | 1.0 | 1.0 |
| **Low** | 1.6 | 4.6 | 4.4 | 2.0 | 1.0 |
| **Medium** | 2.0 | 5.6 | 5.3 | 3.2 | 1.3 |
| **High** | 3.2 | 7.4 | 6.7 | 4.0 | 2.0 |

Compression ratios for the mesh compression technique

Note: The “Color” column in the above table shows ratios for a mesh that uses the UNorm8 format. For a mesh that uses the FP32 format for vertex colors, the ratios are 4.0 on the Low setting, 5.3 on the Medium setting, and 8.0 on the High setting.