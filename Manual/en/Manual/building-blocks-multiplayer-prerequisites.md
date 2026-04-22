# Multiplayer Services Building Blocks prerequisites

To use the Multiplayer Unity Building Blocks, set up the following in your project:

* A [Unity project](https://docs.unity3d.com/hub/manual/Projects.html) (6.0 LTS or later recommended).
* Access to [Unity Dashboard](https://cloud.unity.com/) and a [Unity organization](https://docs.unity3d.com/Documentation/Manual/OrgsManagingyourOrganization.html).
* Editor sign-in with a user that has [access to the organization and project](https://support.unity.com/hc/en-us/articles/29991845748884-How-can-I-manage-users-access-to-projects-in-my-Unity-organization#:~:text=Cause:,that%20require%20this%20specific%20role.).
* Network access to [Unity Services](https://services.docs.unity.com/docs/) endpoints.

## How Multiplayer Building Blocks work

The Multiplayer Building Blocks are designed for flexible and modular integration of session management into your Unity projects. The `SessionSettings` ScriptableObject acts as a central configuration hub for the core functionality of these Building Blocks.

### SessionSettings ScriptableObject

The `SessionSettings` ScriptableObject stores parameters and settings related to your sessions. Crucially, `SessionSettings` includes the `SessionType` string, which is the primary identifier used for managing different types of sessions within your game.

### Session management via SessionType

The `SessionType` string drives session management within the Building Block. When you create, join, or browse for sessions, each `VisualElement` uses the `SessionType` specified in the active `SessionSettings` to filter and interact with the appropriate session. This approach enables dynamic configuration without requiring code changes for each session type.

### Implementation: C# scripts, visual elements, and data separation

The Building Blocks implementation is primarily done through C# **scripts**. A key architectural principle is the clear separation between the visual representation (UI) and the underlying data and logic.

The details of this architecture are as follows:

* `VisualElement` (UI) scripts: These scripts, extending `VisualElement` or its derivatives, are responsible for rendering the UI components that players interact with, handle user input, and display session-related information. These scripts always expose a `UXMLAttribute` for a `SessionSettings` reference or a `SessionType` string value that’s propagated to the `ViewModel` script so that it can work on the expected session.
* `ViewModel` scripts: These scripts manage the actual session data, interact with Unity Multiplayer Services, and handle the game logic related to session states via subscribing to session events tracked by the `SessionObserver` and the `ISession` API.

This separation ensures that UI changes can be made independently of the core session logic, promoting maintainability and scalability.

#### UXML file composition and binding

The user interfaces in the Building Blocks are built using Unity Extensible Markup Language (UXML) files, which are Unity’s equivalent of HTML for UI Toolkit. To learn about UXML scripts, refer to [Introduction to UXML](https://docs.unity3d.com/Documentation/Manual/UIE-WritingUXMLTemplate.html).

The details of this UI are as follows:

* UXML composition:
  + UXML files define the structure and hierarchy of the UI. These files reference various `VisualElement` types and can include custom `VisualElement` implementations. These files specify visual properties like styles, layout, and content.
* Setting up the `SessionType`:
  + To simplify the setup of `SessionSettings` and `SessionType` on the Building Block elements, the UXML files are setup with UIToolkit’s DataSource bindings to propagate a specific `SessionSettings` and associated `SessionType` to all child `VisualElement` from the root element of the UXML.

To learn more about runtime bindings in UIToolkit, refer to [Get started with runtime binding](https://docs.unity3d.com/Documentation/Manual/UIE-get-started-runtime-binding.html).

## Use Multiplayer Play Mode with your Building Blocks

You can use the [Multiplayer Play Mode package](https://docs.unity3d.com/Packages/com.unity.multiplayer.playmode@latest) to test your Multiplayer Building Blocks locally with multiple simulated players. With this package, you can run multiple instances of your game within the Unity Editor, simulating a multiplayer environment without needing to build and deploy to separate devices.

# Additional resources

* [Unity Services: Multiplayer SDK](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual)
* [Unity Netcode for GameObjects](https://docs.unity3d.com/Packages/com.unity.netcode.gameobjects@2.5/manual/index.html)
* [Unity Netcode for Entities](https://docs.unity3d.com/Packages/com.unity.netcode@1.9/manual/index.html)
* [UI Toolkit](https://docs.unity3d.com/6000.2/Documentation/Manual/UIElements.html)