# UI systems

Unity provides three UI systems: UI Toolkit, uGUI, and IMGUI. You can use these systems to create user interfaces (UI) for the Unity Editor and applications made in the Unity Editor.

| **Topic** | **Description** |
| --- | --- |
| [Comparison of UI systems in Unity](UI-system-compare.md) | Unity intends for [UI Toolkit](UIToolkits.md) to become the recommended UI system for new UI development projects, but it still lacks some features available in uGUI and IMGUI. These legacy systems are better suited for certain use cases and are required to support legacy projects. Choose a UI system based on the type of UI you need to develop and the features it requires. |
| [UI Toolkit](UIElements.md) | The latest UI system, designed for optimized performance across platforms and inspired by standard web technologies. Use UI Toolkit to create extensions for the Unity Editor, and to create runtime UI for games and applications. |
| [Create vector graphics](ui-systems/create-vector-graphics) | Use the vector graphics API to create and manipulate vector graphics in your UI. |
| [uGUI (Unity UI)](https://docs.unity3d.com/Packages/com.unity.ugui@latest) | A legacy, GameObject-based UI system for developing runtime UI for games and applications. uGUI uses components and the Game view to arrange, position, and style the UI. |
| [IMGUI (Immediate Mode GUI)](GUIScriptingGuide.md) | A legacy, code-driven UI system that uses the `OnGUI` function to draw and manage UI. Use IMGUI to create custom **Inspectors** for script components, extensions for the Unity Editor, and in-game debugging displays. It’s not recommended for runtime UI. |

## Additional resources

* 📚 **Documentation**: [Get started with UI Toolkit](UIE-simple-ui-toolkit-workflow.md)
* 📚 **Documentation**: [uGUI API reference](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/api/index.html)
* 👥 **Community**: [Join the conversation on Unity Discussions](https://discussions.unity.com/lists/ui)