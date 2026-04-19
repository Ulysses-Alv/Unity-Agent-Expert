# Materials tab

You can use this tab to change how Unity deals with materials and textures when importing your model.

When Unity imports a model without any material assigned, it uses the Unity diffuse material. If the model has materials, Unity imports them as subassets.

![The Materials tab defines how Unity imports materials and textures](../uploads/Main/FBXImporter-Materials-1.png)

The Materials tab defines how Unity imports materials and textures

If your model has textures, you can also extract them into your project using the [Extract Textures](#textures) button.

| **Property** | **Function** |
| --- | --- |
| **Material Creation Mode** | Define how you want Unity to generate or import the materials for your model. When you choose **None** from this drop-down menu, the Inspector hides the rest of the settings on this tab. The following options are available:  * **None** - Do not use any materials embedded within this model. Use Unity’s default diffuse material instead. * **Standard (Legacy)** - On import, Unity applies a set of default rules to generate materials. If you want to customize how Unity generates material via scripting, choose the **Import via MaterialDescription** mode instead. * **Import via MaterialDescription** - On import, Unity uses the material description embedded within the FBX file to generate the materials. This method provides more accurate results than previous import methods, and supports a wider range of material types, such as [Arnold](https://www.arnoldrenderer.com/home/) and [Physical](https://knowledge.autodesk.com/support/3ds-max/learn-explore/caas/CloudHelp/cloudhelp/2020/ENU/3DSMax-Lighting-Shading/files/GUID-809B9123-21A2-443E-A7A4-0DAB70410B8D-htm.html?st=Physical%20Material) from Autodesk, as well as Unity’s [HDRP Materials](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/Material-Type.html). For more information, see the [Material description](#material_description) section below. |
| **sRGB Albedo Colors** | Enable this option to use Albedo colors in gamma space. This is enabled by default for legacy import methods.   Disable this for Projects using [linear color space](color-spaces-landing.md).  This property is not available if you choose **Import via MaterialDescription** from the **Material Creation Mode** drop-down menu. |
| **Location** | Define how to access the materials and textures. Different properties are available depending on which of these options you choose. The following options are available:  * **Use Embedded Materials** - [Keep the imported materials inside the imported asset](#Embedded). This is the default option from Unity 2017.2 onwards. * **Use External Materials (Legacy)** - [Extract imported materials as external assets](#Legacy). This is a Legacy way of handling materials, and is intended for Projects created with 2017.1 or previous versions of Unity. |

## Use Embedded Materials

When you choose **Use Embedded Materials** for the **Location** option, the following import options appear:

![Import settings for materials](../uploads/Main/FBXImporter-Materials-2.png)

Import settings for materials

**(A)** Click the **Extract Materials** and **Extract textures** buttons to extract all materials and textures that are embedded in your imported asset. These are greyed out if there are no subassets to extract. Below these buttons, Unity displays any messages about the import process.

**(B)** The [On Demand Remap](#remapped) section provides the [Naming](#naming) and [Search](#search) properties which allow you to customize how Unity maps imported materials to the model. Click the **Search and Remap** button to remap your imported materials to existing material assets. Nothing changes if Unity can’t find any materials with the correct name.

**(C)** Unity displays all imported materials found in the asset in the [Remapped Materials](#remapped) list. If Unity can’t automatically match each material to an existing material asset in your Project, you can set references to the materials yourself in this list.

### Remapped Materials

New imports or changes to the original asset do not affect extracted materials. If you want to re-import the materials from the source asset, you need to remove the references to the extracted materials in the **Remapped Materials** list. To remove an item from the list, select it and press the Backspace key on your keyboard.

### Naming

Define a naming strategy for the materials.

| **Property** | **Function** |
| --- | --- |
| **By Base Texture Name** | Use the name of the diffuse texture of the imported material to name the material. When you don’t assign a diffuse texture to the material, Unity uses the name of the imported material. |
| **From Model’s Material** | Use the name of the imported material to name the material. |
| **Model Name + Model’s Material** | Use the name of the **model file** in combination with the name of the imported material to name the material. |

### Search

Define where Unity tries to locate existing materials when it uses the name defined by the **Naming** option.

| **Property** | **Function** |
| --- | --- |
| **Local Materials Folder** | Find existing materials in the local `Materials` subfolder, which is in the same folder as the model file. |
| **Recursive-Up** | Find existing materials in all materials subfolders in all parent folders up to the `Assets` folder. |
| **Project-Wide** | Find existing materials in all Unity Project folders. |

### Material description

Starting with version 2019.3, Unity introduced the ability to modify the material mapping during import via scripting. Users can modify how Unity maps the imported material properties from the data embedded in an FBX file to Unity material properties. The material description defines a name and several sets of values that describe the material and any textures it references. For more information about the structure of this description, see the [MaterialDescription](../ScriptReference/AssetImporters.MaterialDescription.md) class reference page.

When in [ImportViaMaterialDescription](../ScriptReference/ModelImporterMaterialImportMode.ImportViaMaterialDescription.md) mode, the model importer delegates the creation of materials to the [AssetPostProcessor.OnPreprocessMaterialDescription](../ScriptReference/AssetPostprocessor.OnPreprocessMaterialDescription.md) callback.

Unity provides default implementations of this Post Processor that handle the following materials:

* FBX Standard Material, Arnold Standard, Autodesk Interactive, and 3ds Physical Material from FBX files
* Sketchup, Collada, and 3ds Materials

These default implementations handle importing materials differently from the [ImportStandard](../ScriptReference/ModelImporterMaterialImportMode.ImportStandard.md) mode, including the following improvements:

* It supports more material types, such as Autodesk’s [Arnold](https://www.arnoldrenderer.com/home/) and Interactive, or [Physical](https://knowledge.autodesk.com/support/3ds-max/learn-explore/caas/CloudHelp/cloudhelp/2020/ENU/3DSMax-Lighting-Shading/files/GUID-809B9123-21A2-443E-A7A4-0DAB70410B8D-htm.html?st=Physical%20Material), as well as Unity’s [HDRP Materials](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/manual/Material-Type.html).
* It supports [Emissive Materials](lighting-emissive-materials.md).
* If a diffuse texture is set, it ignores the diffuse color (this matches how it works in Autodesk® Maya® and Autodesk® 3ds Max®).
* It takes the bump factor, the emissive color, and emissive factor into account.
* It imports emissive color animation when defined in the FBX file.

  **Note**: 3ds Max does not export emissive color animation, so Unity cannot import it.
* It imports transparent materials as fully transparent. The legacy system imports them as fully opaque.

In addition, it imports all [Autodesk Interactive](https://knowledge.autodesk.com/support/3ds-max/learn-explore/caas/CloudHelp/cloudhelp/2020/ENU/3DSMax-Lighting-Shading/files/GUID-7EEAC650-7D26-40AE-AC14-577F7A2EF2B3-htm.html) material property animations and no longer ignores the opacity when importing materials from 3ds files.

## Use External Materials (Legacy)

When you choose **Use External Materials (Legacy)** for the **Location** option, the following import options appear:

![Import settings for Use External Materials (Legacy)](../uploads/Main/FBXImporter-Materials-3.png)

Import settings for Use External Materials (Legacy)

This option extracts materials and saves them externally instead of saving them inside your model asset. The [Naming](#naming) and [Search](#search) properties help Unity find imported materials to map to the model.

Before Unity version 2017.2, this was the default way of handling materials.