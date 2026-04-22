# Panel Settings properties reference

The following tables describe the Panel Settings asset properties:

| **Property** | **Description** |
| --- | --- |
| **Theme Style Sheet** | Apply a default [TSS file](UIE-tss.md) to every UI Document that the panel renders. |
| **Text Settings** | Set the [UITK Text Settings](UIE-text-setting-asset.md) asset for this panel. If this asset isn’t set, UI Toolkit automatically creates one with the default settings. |
| **Binding Console Logs** | Set the [binding log levels](UIE-runtime-binding-logging-levels.md). |
| **Render Mode** | Set the render mode for the panel:    * **World Space**: Position and render UI elements alongside 2D or 3D objects in the Scene. * **Screen Space Overlay**: Render UI elements on the screen. |
| **Scale Mode Parameters** | Display [different properties](#scale-mode-parameters) depending on the **Scale Mode** setting. |
| **Dynamic Atlas Settings** | Set the [dynamic atlas settings](#dynamic-atlas-settings) for the panel. |

## Collider Update Mode

The ****Collider** Update Mode** property appears only when you set the **Render Mode** to **World Space**. It determines how the collider updates when the panel changes size or position.

The following table describes the **Collider Update Mode** properties:

| **Property** | **Description** |
| --- | --- |
| **Match 3D bounding box** | The collider matches the 3D bounding box of the panel. |
| **Match 2D document rect** | The collider matches the 2D document rectangle of the panel. |
| **Keep existing collider** | The collider doesn’t change. |

## Collider is Trigger

The **Collider is Trigger** checkbox appears only when you set the **Render Mode** to **World Space**. It specifies whether the collider is a trigger.

The [`Collider.isTrigger`](../ScriptReference/Collider-isTrigger.md) property tells the [Physics engine](PhysicsSection.md) whether the collider participates in physical **collisions**. Trigger colliders detect overlaps but don’t block movement or react to forces.

UI interactions work the same way regardless of whether the collider is a trigger. In general, use trigger colliders for UI elements because trigger colliders don’t interact with the physical world. Characters don’t collide with them, and objects like rocks don’t fall on buttons and get stuck.

However, if you use non-trigger colliders and don’t want them to affect gameplay physics, use the Physics layer interaction matrix to disable collisions with the UI layer. For more information, refer to [Create and configure a trigger collider](collider-interactions-create-trigger.md).

## Screen Space Overlay parameters

The following table describes the parameters that are only available for **Screen Space Overlay**:

| **Property** | **Description** |
| --- | --- |
| **Target Texture** | Set the [render texture](class-RenderTexture.md) used to render the panel. |
| **Target Display** | Set the display for the panel. Only set this if the target texture isn’t set, because the target texture takes precedence over the target display. |
| **Sort Order** | Set the order in which the UI System draws panels when the **Scene** uses more than one panel. Panels with higher **Sort Order** values are drawn on top of panels with lower values. |
| **Scale Mode** | Set [how the panel’s UI scales](#scale-mode-parameters) when the screen size changes. |

## Scale Mode parameters

The following tables describe the parameters for each scale mode:

### Constant Pixel Size

Set elements to stay the same size, in **pixels**, regardless of screen size.

| **Property** | **Description** |
| --- | --- |
| **Scale** | Multiply element sizes by this value. Must be greater than 0. |
| **Reference **Sprite** Pixel per Unit** | Set the number of pixels per unit for the reference sprite. This value is used to calculate the size of the element in pixels. |

### Constant Physical Size

Set elements stay the same physical size regardless of screen size and resolution.|

| **Parameters** | **Description** |
| --- | --- |
| **Reference DPI** | Set the **Reference DPI** value to the screen density that your UI was designed for. When the system renders your UI, it tries to find the actual DPI value of the screen, and compares it to the **Reference DPI**. If they’re different, the system scales the UI accordingly. |
| **Fallback DPI** | Use this value if the UI system can’t determine the screen DPI. |
| **Reference Sprite Pixel per Unit** | Set the number of pixels per unit for the reference sprite. This value is used to calculate the size of the element in pixels. |

### Scale with Screen Size

Set elements to grow or shrink depending on the screen size.

| **Property** | **Description** |
| --- | --- |
| **Screen Match Mode** | Set how to scale elements when the screen resolution is different from the **Reference Resolution**:    * **Match Width or Height**: Scale the canvas area with the width or the height as reference, or a linear interpolation between the width and the height. * **Shrink**: Crop the canvas area either horizontally or vertically, so the size of the canvas is smaller than the reference. * **Expand**: Expand the canvas area either horizontally or vertically, so the size of the canvas is larger than the reference. |
| **Reference Resolution** | Set the resolution that this panel’s UI is designed for. If the screen resolution is larger than the reference resolution, the UI scales up. If it’s smaller, the UI scales down. How the UI scales depends on the **Screen Match Mode**. |
| **Screen Match Mode Parameters** | When **Screen Match Mode** is set to **Match Width or Height**, the **Match** value controls whether the UI system scales the UI to match the screen width, the screen height, or a mix of the two.   For example, if the value is `0`, it matches the width; if the value is `1`, it matches the height; if the value is `0.4`, it linearly interpolates between width and height by `40%`. |

## Dynamic Atlas Settings

Specify the settings that the dynamic atlas system uses.

The following table describes the **Dynamic Atlas Settings** properties:

| **Property** | **Description** |
| --- | --- |
| **Min Atlas Size** | Set the minimum size (width/height) of the atlas texture, in pixels. |
| **Max Atlas Size** | Set the maximum size (width/height) of the atlas texture, in pixels. |
| **Max Sub Texture Size** | Set the maximum width and height of a texture that can be added to the atlas. |
| **Active Filters** | Set the [filters](#active-filters) that the dynamic atlas system uses to exclude textures from the texture atlas. |

### Active Filters

You can apply more than one filter at a time. The following table describes each active filter.

| **Property** | **Description** |
| --- | --- |
| **Nothing** | Disable all the filters. |
| **Everything** | Apply all the filters. |
| **Readability** | Exclude textures that [`Texture2D.isReadable`](../ScriptReference/Texture-isReadable.md) is set to `true`.   Unity can’t automatically update the dynamic atlas when you add a texture through APIs like [`Texture2D.SetPixels`](../ScriptReference/Texture2D.SetPixels.md). You can use [`RuntimePanelUtils.SetTextureDirty`](../ScriptReference/UIElements.RuntimePanelUtils.SetTextureDirty.md) to force the atlas to update its content for a given texture. |
| **Size** | Exclude textures that are larger than the **Max Sub Texture Size** setting.  Large textures can saturate the atlas quickly. If you don’t want to add large textures to the atlas, select this filter and adjust **Max Sub Texture Size** to fit your needs. By default, textures larger than `64x64` aren’t allowed into the atlas. |
| **Format** | Store sRGB-encoded data with 8-bits per channel precision, and excludes sub-textures that would lose precision or be truncated when added to the atlas, such as an R16G16B16A16\_FLOAT texture. |
| **Color Space** | Exclude R8G8B8A8\_UNORM content when the project is in a linear color space.  In a linear color space, the format of the [`RenderTexture`](../ScriptReference/RenderTexture.md) of the dynamic atlas is R8G8B8A8\_SRGB. The data stored in the [`RenderTexture`](../ScriptReference/RenderTexture.md) is sRGB-encoded. When read from, it’s linearized, and when written to, it’s encoded to sRGB. Because of the limited precision of the format, R8G8B8A8\_UNORM content stored in the atlas could cause banding to occur. |
| **Filter Mode** | Exclude textures for which the sub-texture filter mode, such as Point or Bilinear, doesn’t match the atlas filter mode. This prevents the mismatch between the sub-texture filter mode and the atlas filter mode. The mismatch could cause unexpected blurriness or blockiness. |

## Clear Settings

Set how the panel clears before rendering. The **Clear Settings** properties are only available when the **Render Mode** is set to **Screen Space Overlay**.

The following table describes the **Clear Settings** properties:

| **Property** | **Description** |
| --- | --- |
| **Clear Color** | Set whether the panel clears the color buffer before rendering. |
| **Clear Color Value** | Set the color that the panel uses to clear the color buffer. |
| **Clear Depth Stencil** | Set whether the panel clears the depth stencil before rendering. |

## Buffer Management

Set how the panel manages its buffers.

The following table describes the **Buffer Management** properties:

| **Property** | **Description** |
| --- | --- |
| **Vertex budget** | Set the approximate number of vertices the panel uses. This helps initialize vertex buffers with the right size, which can reduce draw calls. A value of `0` uses the default value for UI rendering. |

## Additional resources

* [Get started with runtime UI](UIE-get-started-with-runtime-ui.md)
* [Render UI in the Game view](UIE-render-runtime-ui.md)
* [Runtime event system](UIE-Runtime-Event-System.md)
* [FAQ for input and event systems with UI Toolkit](UIE-faq-event-and-input-system.md)
* [Performance consideration for runtime UI](UIE-performance-consideration-runtime.md)
* [`PanelSettings`](../ScriptReference/UIElements.PanelSettings.md)