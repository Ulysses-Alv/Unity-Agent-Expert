# Prepare your sprites for the 2D Pixel Perfect Camera in URP

Set up your **sprites** for compatibility with the **Pixel** Perfect **Camera** component.

1. After importing your asset into the project as sprites, set all sprites to the same **Pixels Per Unit** value.  
   ![Setting PPU value](../../uploads/urp/2D/2D_Pix_image_1.png)
2. In the sprites’ **Inspector** window, open the **Filter Mode** dropdown and select **Point**.  
   ![Set Point mode](../../uploads/urp/2D/2D_Pix_image_2.png)
3. Open the ****Compression**** dropdown and select **None**.  
   ![Set None compression](../../uploads/urp/2D/2D_Pix_image_3.png)

Follow the steps below to set the pivot for a sprite:

1. Open the **Sprite Editor** window for the selected sprite.
2. If the sprite’s **Sprite Mode** is set to **Multiple** and there are multiple individual sprite elements in the imported texture, then you need to set a pivot point for each individual sprite element.
3. In the Sprite panel at the lower left of the **Sprite Editor** window, open the **Pivot** dropdown and select **Custom**. Then open the **Pivot Unit Mode** and select **Pixels**. This allows you to set the pivot point’s coordinates in pixels, or drag the pivot point around in the **Sprite Editor** and have it automatically snap to pixel corners.  
   ![Setting the Sprite’s Pivot](../../uploads/urp/2D/2D_Pix_image_4.png)
4. Repeat step 3 for each sprite element as necessary.

## Set up snap settings

Follow the steps below to set the snap settings for your project to ensure that the movement of pixelated sprites are consistent with each other:

1. Open the Snapping menu from the [**Grid and Snap** toolbar overlay](../GridAndSnapOverlay).
2. Set the **Grid Size** properties to 1 divided by the **Asset Pixels Per Unit** value of the Pixel Perfect Camera component. For example, if the **Asset Pixels Per Unit** is 100, set the *X*, *Y*, and *Z* values of the **Grid Size** property to to 0.01 (1 divided by 100).
3. Select the **Grid Snapping** icon to enable grid snapping.

Unity doesn’t automatically snap existing **GameObjects** to the grid. To snap existing GameObjects, select them in the **Hierarchy** window, then in the **Align Selected** property select **All Axes**.

To snap sprites to smaller increments of the grid, press **Ctrl + [** (macOS: **Cmd + [**) to decrease the grid size, or** Ctrl + ]** (macOS: **Cmd + ]**) to increase the grid size.