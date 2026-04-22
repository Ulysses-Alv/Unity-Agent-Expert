# Tags and Layers

The **Tags and Layers** settings (main menu: **Edit** > **Project Settings**, then select the **Tags and Layers** category) allows you to set up [Tags](#Tags), [Sorting Layers](#SortingLayers) and [Layers](#Layers).

![The Tags and Layers Manager, before any custom tags or layers have been defined](../uploads/Main/TagManager55.png)

The Tags and Layers Manager, before any custom tags or layers have been defined

## Tags

**Tags** are marker values that you can use to identify objects in your Project (see documentation on [Tags](Tags.md) for further details). To add a new Tag, click the plus button (+) at the bottom-right of the list, and name your new Tag.

![Adding a new Tag](../uploads/Main/TagManagerAddNew.png)

Adding a new Tag

Note that once you have named a Tag, you cannot rename it. To remove a Tag, click on it and then click the minus (-) button at the bottom-right of the list.

![The tags list showing four custom tags](../uploads/Main/TagManagerAddedNew.png)

The tags list showing four custom tags

## Sorting Layers

Sorting Layers are used in conjunction with [Sprite](sprite/sprite-landing.md) graphics in the 2D system. *Sorting* refers to the overlay order of different Sprites.

![Adding a new Sorting Layer](../uploads/Main/SortingLayerManagerAddNew.png)

Adding a new Sorting Layer

To add and remove Sorting Layers, use the plus and minus (+/-) buttons at the bottom-right of the list. To change their order, drag the handle at the left-hand side of each Layer item.

![The Sorting Layers list, showing four custom sorting layers](../uploads/Main/SortingLayerManagerAddedNew.png)

The Sorting Layers list, showing four custom sorting layers

## Layers

Use Layers throughout the Unity Editor as a way to create groups of objects that share particular characteristics (see documentation on [Layers](Layers.md) for further details). Use Layers primarily to restrict operations such as raycasting or rendering, so that they are only applied to the relevant groups of objects. Layers marked as **Builtin Layer** are default layers used by Unity, which you cannot edit. You can customise layers marked as **User Layer**.

![Adding a new Layer](../uploads/Main/LayerManagerAddNew.png)

Adding a new Layer

To customise **User Layers**, type a custom name into the text field for each one you wish to use. Note that you can’t add to the number of Layers but, unlike Tags, you can rename Layers.

## Rendering Layers

If your project uses the Universal **Render Pipeline** (URP) or the High Definition Render Pipeline (HDRP), this section lists the names of Rendering Layers. Use Rendering Layers to configure which lights or decals affect which **GameObjects**. Refer to the following for more information:

* [Rendering Layers in URP](urp/features/rendering-layers.md)
* [Rendering Layers in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/Rendering-Layers.html)

The Built-In Rendering Pipeline doesn’t support Rendering Layers.

TagManager