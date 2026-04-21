# Set up your sprite for 9-slicing

Before you 9-slice a **Sprite**, you need to ensure the Sprite is set up properly.

1. Make sure the **Mesh Type** is set to **Full Rect**. To apply this, select the Sprite, then in the Inspector window click the **Mesh Type** drop-down and select **Full Rect**. If the **Mesh Type** is set to **Tight**, 9-slicing might not work correctly, because of how the **Sprite Renderer** generates and renders the Sprite when it is set up for 9-slicing.
2. Define the borders of the Sprite using the Sprite Editor. To do this, select the Sprite, then in the [Inspector](../../UsingTheInspector.md) window click the **Sprite Editor** button.
3. Use the Sprite Editor window to define the borders of the Sprite (that is, where you want to define the tiled areas, such as the walls of a floor tile). To do this, use the Sprite control panel’s **L**, **R**, **T**, and **B** fields (left, right, top, and bottom, respectively). Alternatively, click and drag the green dots at the top, bottom, and sides.
4. Click **Apply** in the Sprite Editor window’s top bar.
5. Close the Sprite Editor window, and drag the Sprite from the [Project window](../../ProjectView.md) into the **Scene** view to begin working on it.