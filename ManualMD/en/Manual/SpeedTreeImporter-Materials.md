# Materials tab

Use the **Materials** tab to change how Unity deals with materials for an imported SpeedTree model. If the model has materials, Unity imports them as sub-assets.

To change other import settings for the [SpeedTree](SpeedTree.md) model, see the [Model tab](SpeedTreeImporter-Model.md).

![SpeedTree Importer Materials tab](../uploads/Main/class-SpeedTreeImporter-Materials.png)

SpeedTree Importer Materials tab

The **Materials** tab has these sections:

**(A)** The [Location and Materials](#location) properties allow you to control how Unity accesses the materials inside the imported asset.

**(B)** The [Remapped Materials](#remapped) section allows you to remap the imported materials to existing materials in your project.

**(C)** The **Revert** and **Apply** options always appear, but you can only select them after you make changes to the import settings. If you change settings in the Remapped Materials section of the **Materials** tab, the **Apply & Generate Materials** option appears.

**(D)** The properties for the **GameObject** (read-only) appear at the bottom of the ****Inspector**** view, including a preview of the SpeedTree model with its material applied.

## Location and Materials

| **Property** | **Description** |
| --- | --- |
| **Location** | Define how the materials are accessed. These options aren’t available for .st9 files, which only use embedded materials. The following options are available:  * **Use Embedded Materials** - Import materials as sub-assets inside the SpeedTree model. This is the default and recommended option. * **Use External Materials (Legacy)** - Extract materials to a folder with the same name and location as the model. This is a deprecated option for projects created with Unity 2017.1 or earlier and isn’t recommended for use. |
| **Materials** | For .st and .spm files, this property is visible only when **Location** is set to **Use Embedded Materials**. Select **Extract Materials** to extract the embedded materials from the imported asset and save them to a folder of your choice. Extracted materials populate the Remapped Materials list. Use this option to modify the materials or share them with other SpeedTree assets. This option isn’t available if there are no sub-asset materials to extract from the model. |

## Remapped Materials

The Remapped Materials section lists the embedded materials and allows you to remap them to other materials that are already in the project.

If you already have a set of materials you want to use for this asset, expand the **On Demand Remap** group, select **Search and Remap** , and give the location of the materials. Unity remaps the imported materials to existing materials with the same name. Nothing changes if Unity can’t find any materials with the same name. You can manually remap materials from the Remapped Materials list.

![The Remapped Materials section with the On Demand Remap group expanded](../uploads/Main/class-SpeedTreeImporter-Materials_SearchAndMap.png)

The Remapped Materials section with the On Demand Remap group expanded

To save the imported SpeedTree model with the remapped materials, select **Apply**.

If you make changes to the material properties through the Model tab’s import settings or an **LOD** setting override, select **Regenerate Materials** to re-import the materials in the Remapped Materials list with the updated settings.

> **Warning:** Re-generating the external materials removes any manual modifications you made to the material files.

New imports or changes to the original asset don’t affect extracted materials. If you want to re-import the materials from the source asset, remove the references to the extracted materials in the Remapped Materials list. To remove an item from the list, select it and press the Backspace key on your keyboard.