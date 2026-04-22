# Create a lens flare

Create a flare and apply it to a **lens flare** component, then configure it to be visible by cameras.

**Note:** This workflow is compatible only with the Built-in **Render Pipeline**. For similar functionality in other render pipelines, see [Lens flares and halos](visual-effects-lens-flares.md).

## Apply a Flare asset

1. Assign the Flare asset to a [Light component](class-Light.md) or a [Lens flare component](class-LensFlare.md).
   * If you assign it to the **Flare** property of a [Light component](class-Light.md), Unity automatically tracks the position and direction of the Light and uses those values to configure the appearance of the lens flare.
   * If you assign it to the **Flare** property of a [Lens flare component](class-LensFlare.md), you can use the Lens Flare component to configure additional values for more precise control.
2. If you want a [Camera](class-Camera.md) to see lens flares, select the camera then select **Add Component** > **Flare Layer**.

   **Note:** The Flare Layer component has no properties.
3. To see the lens flare effect in the **Scene View**, enable the Effect toggle in the **Scene** View **toolbar** and, in the drop-down, enable **Flares**.

## Create a lens flare with the Lens Flare component

1. Create a new **GameObject** (menu bar: **GameObject > Create Empty**).
2. In the Inspector, click **Add Component > Effects > Lens Flare**.
3. Assign a [Flare Asset](class-Flare.md) to the **Flare** property.
4. If you want a [Camera](class-Camera.md) to see lens flares, select the camera then select **Add Component** > **Flare Layer**.

   **Note:** The Flare Layer component has no properties.
5. To see the lens flare effect in the **Scene View**, enable the Effect toggle in the Scene View toolbar and, in the drop-down, enable **Flares**.

![Enable the Effect toggle to view lens flares in the Scene view](../uploads/Main/LensFlare-FXButton.png)

Enable the Effect toggle to view lens flares in the Scene view