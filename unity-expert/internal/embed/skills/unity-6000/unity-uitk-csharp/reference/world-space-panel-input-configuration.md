# Panel Input Configuration properties reference

The **Panel Input Configuration** component provides properties to configure the input settings for the World Space UI, and to specify how UI Toolkit interacts with uGUI (Unity UI).

## World Space Options properties

The following table describes the properties available for **World Space Options**:

| **Property** | **Description** |
| --- | --- |
| **Interaction Layers** | Select one or more Physics layers to cast World Space rays and determine pointer event targets. |
| **Max Interaction Distance** | Set how far away interactions with World Space UI are possible. Defaults to unlimited (infinity), but you can customize it for **XR** or performance needs. The distance uses **GameObject** units, consistent with transform positions and **Camera** **clipping planes**. |
| **Default Event Camera Is Main Camera** | Enable or disable the use of the main camera as the event camera to transform screen positions into World Space rays.  If enabled and the main camera changes, the event camera automatically updates to the new camera. If no main camera is active, Unity disables screen-based events until a main camera becomes active. |
| **Event Cameras** | Set one or more cameras to transform screen positions into World Space rays when the main camera isn’t used.   If you specify multiple cameras, Unity processes them sequentially until one produces a ray that hits a Collider. **Note**: If you don’t specify any event cameras, Unity disables screen-based events. |

## Event System Interaction properties

You can use the ****Event System** Interaction** properties to configure how the UI Toolkit interacts with the uGUI system.

### Panel Input Redirection

The **Panel Input Redirection** property determines whether the uGUI EventSystem redirects panel input. If enabled, you can choose from the following options:

* **Auto-switch (redirect from EventSystem if present)** (default): As long as an EventSystem is active, it serves as the source of input for all UI. Otherwise, UIToolkit’s built-in input handles the input.
* **No input redirection**: UI Toolkit’s built-in input handles the input, regardless of whether an EventSystem is active.
* **Always redirect from EventSystem (wait if unavailable)**: Input remains disabled until an EventSystem is active, at which point it becomes the input source. Use this mode to dynamically enable or disable all UI input by toggling the EventSystem component.

### Auto Create Panel Components

The **Auto Create Panel Components** property determines whether to automatically create UI Toolkit components under the EventSystem to handle input redirection between UI Toolkit and uGUI panels.

If you enable this option, Unity automatically creates the following components:

* **Panel Event Handler**
* **Panel Raycaster**
* **World Document Raycaster**

If you disable this option, ensure to assign a **World Document Raycaster** component for each World Space event camera. You must also ensure that each panel has an associated **[Panel Raycaster](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/script-PanelRaycaster.html)** and **[Panel Event Handler](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/script-PanelEventHandler.html)** component through your C# code. The following code example shows how to do this:

```
using UnityEngine.UIElements;

// Add this component to a GameObject to automatically initialize the following components
// - PanelRaycaster
// - PanelEventHandler
public class PanelComponentsInitializer : MonoBehaviour
{
    public PanelSettings panelSettings;

    void Start()
    {
        var panelRaycaster = GetComponent<PanelRaycaster>();
        if (panelRaycaster != null)
            panelRaycaster.panel = panelSettings.panel;

        var panelEventHandler = GetComponent<PanelEventHandler>();
        if (panelEventHandler != null)
            panelEventHandler.panel = panelSettings.panel;
    }
}
```

## Additional resources

* [Runtime Panel Settings](../UIE-Runtime-Panel-Settings.md)