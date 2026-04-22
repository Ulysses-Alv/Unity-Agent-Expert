# UI Toolkit live reload

Use the UI Toolkit live reload feature to view changes to your UI immediately in the Unity Editor and runtime UI. When you make changes to your UI, such as modifying a UXML or USS file in UI Builder or a text editor, the changes are automatically applied without needing to enter Play mode or manually refresh the UI.

## Editor windows

For an Editor window, live reload is disabled by default. You can enable it through the Editor window’s **More** (⋮) menu. It applies to all windows of the same type.

![The menu to toggle live reload in an Editor window.](../../uploads/Main/uitk/live-reload.png)

The menu to toggle live reload in an Editor window.

**Note**: Live reload recreates the entire UI. If you need to preserve state or values, store them so the window restores them after recreation.

## Runtime

For the runtime, live reload is enabled by default. You can disable it through the Game view’s **More** (⋮) menu. It applies to all Game views.

**Note**: If you add MonoBehaviour components to the same **GameObject** as the [UI Document component](../UIE-create-ui-document-component.md), the reload disables them before the reload and re-enables them after. Place UI interaction code in `OnEnable` and `OnDisable` on those companions for initialization and teardown.

## Additional resources

* [UI debugger](../UIE-ui-debugger.md)