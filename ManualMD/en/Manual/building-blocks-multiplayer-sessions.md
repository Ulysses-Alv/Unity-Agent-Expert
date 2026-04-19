# Multiplayer Sessions Building Block

Implement Unity’s [multiplayer sessions](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual) in your project with this Unity Building Block’s pre-made UI elements. Learn what sessions are, how to set them up, how to debug them, and how to connect to Unity’s Netcode packages.

## Introduction

The Multiplayer Session Building Block helps you integrate Unity’s multiplayer sessions into your project to connect players. This Building Block can serve as a starting point for your multiplayer projects, or as a quick integration of sessions throughout development.

This Building Block demonstrates how to do the following:

* Create and join a session through the Multiplayer Services SDK.
* Use the Quick Join API to quickly get players into any existing session.
* Use the Session Browsing API to list existing sessions and join a specific one.
* Use a JoinCode to share and join a specific session.

The Multiplayer sessions Building Block includes example **scenes**, UIToolkit UI assets, and runtime components to create, browse, and join sessions, and debug them with a SessionInfo panel. If you need gameplay synchronization, you can optionally add Netcode using Unity [Netcode for GameObjects](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-your-first-session) or [Netcode for Entities](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-with-netcode-for-entities).

## Prerequisites

* Install the [Unity Building Block - Multiplayer Session](https://assetstore.unity.com/packages/essentials/tutorial-projects/unity-building-block-multiplayer-session-341930) from the Unity **Asset Store**:.
* [Link your Unity Editor project to a cloud project](https://docs.unity.com/en-us/cloud/projects/configure-project-for-unity-cloud).

## Initial setup

You can validate that the Building Block is installed correctly by opening one of the [included example scenes](#included-resources).

1. Open one of the Multiplayer scenes by navigating to **Assets** > **Blocks** > **MultiplayerSession** > **Scenes**.
2. Select **Play** to enter Play mode.

You can now interact with the UI elements in the scene to create a new session.

## Multiplayer sessions

A [multiplayer session](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual) represents a group of connected players and provides an abstraction layer for managing multiplayer game states, including initial player connections, host election, players joining and leaving, and network connection establishment.
Multiplayer sessions don’t facilitate gameplay synchronization, but are compatible with either [Netcode for GameObjects](https://docs.unity3d.com/Packages/com.unity.netcode.gameobjects@latest) or [Network for Entities](https://docs.unity3d.com/Packages/com.unity.netcode@latest), which can handle authoritative game state replication after players connect. For example, a typical connection flow might look like the following:

1. Initialize Unity Services
2. Authenticate player
3. Create or join a session
4. Connect to the right network server or host
5. Load gameplay scene

## Included resources

The Multiplayer Session Building Block contains the following scenes:

| Scene | Description |
| --- | --- |
| JoinByBrowsing | Let players browse available sessions to join or create their own new session. Ideal for finding lobbies with specific settings or players. |
| JoinByCode | Let players create a session and share a simple join code that other players can use to join the session. Perfect for joining and inviting players to a specific session through copying a code. |
| QuickJoin | Let players join the first available session. Perfect for quickly joining any existing session when matching with specific rules doesn’t matter. |
| QuickJoinDebug | Similar to QuickJoin but with an additional SessionInfo window with information about the session. Useful for debugging a session. |

### JoinByBrowsing

`SessionBrowser` UXML:

* `CreateSessionElement`: Specify a session name and create a session.
* `SessionBrowserElement`: Displays the available sessions. Click to select a session.
* `RefreshListButton`: Refreshes the list of sessions from the `SessionBrowserElement`.
* `JoinButton`: Join a selected session from the list.

`CurrentSession` UXML:

* `SessionNameLabel`: Displays the name of the watched session.
* `PlayerCountLabel`: Displays the number of players in the watched session.
* `PlayerNameLabel`: Displays the player’s own name.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `PlayerListView`: Displays the list of players that have joined the session.
* `LeaveSessionButton`: Select to leave the joined session.

### JoinByCode

`JoinSessionByCode` UXML:

* `CreateSessionElement`: Specify a session name and create a session.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `JoinSessionByCode`: Enter a session code to join a session.

`CurrentSession` UXML:

* `SessionNameLabel`: Displays the name of the watched session.
* `PlayerCountLabel`: Displays the number of players in the watched session.
* `PlayerNameLabel`: Displays the player’s own name.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `PlayerListView`: Displays the list of players that have joined the session.
* `LeaveSessionButton`: Select to leave the joined session.

### QuickJoin

`JoinSessionByQuickJoin` UXML:

* `QuickJoinButton`: Join a session with the [Quick Join API](https://docs.unity.com/ugs/manual/lobby/manual/quick-join).

`CurrentSession` UXML:

* `SessionNameLabel`: Displays the name of the watched session.
* `PlayerCountLabel`: Displays the number of players in the watched session.
* `PlayerNameLabel`: Displays the player’s own name.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `PlayerListView`: Displays the list of players that have joined the session.
* `LeaveSessionButton`: Select to leave the joined session.

### QuickJoinDebug

`JoinSessionByQuickJoin` UXML:

* `QuickJoinButton`: Join a session with the [Quick Join API](https://docs.unity.com/ugs/manual/lobby/manual/quick-join).

`SessionInfo` UXML:

* `SessionInfo`: Displays debugging information for an active session, similar to [the Sessions Viewer](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/session-viewer).

`CurrentSession` UXML:

* `SessionNameLabel`: Displays the name of the watched session.
* `PlayerCountLabel`: Displays the number of players in the watched session.
* `PlayerNameLabel`: Displays the player’s own name.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `PlayerListView`: Displays the list of players that have joined the session.
* `LeaveSessionButton`: Select to leave the joined session.

## Joining with multiple clients locally

Use [Multiplayer Play Mode](https://docs.unity3d.com/Packages/com.unity.multiplayer.playmode@latest) to test the setup with multiple players directly from the Unity Editor.

## Add Netcode to a session

To integrate either Netcode for **GameObjects** or Netcode for Entities with your sessions, do the following:

1. Install the Netcode package of your choice using the Package Manager (**Window** > **Package Management** > **Package Manager**).
   * Install either [Unity Netcode for GameObjects](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-your-first-session#build-a-session-with-netcode-for-gameobjects) or [Netcode for Entities](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-with-netcode-for-entities#build-a-session-with-netcode-for-entities).
2. Enable **Create Network Session** in `SessionSettings`.
   * In your `SessionSettings` ScriptableObject, locate and enable the **Create Network Session** option, which allows your sessions to be recognized and managed by your chosen Netcode solution.
3. Change the **Network Type** in `SessionSettings`.
   * Adjust the **Network Type** within your `SessionSettings` to match your intended network architecture. You can choose from the following:
     + Relay: For cloud-based relay services.
     + Distributed authority: For using Netcode for GameObjects’ distributed authority.
     + Direct: For direct connections using IP Address and Port.

To understand more about the different NetworkTypes, refer to the [Network connection management](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/manage-session-network-connection#types-of-network).

## Troubleshooting

### Multiplayer Services not available

* Multiplayer Services have to be initialized. You can initialize your own scenes with default settings by adding a ServicesInitialization and PlayerAuthentication components.

### Authentication failure

* Ensure you’re signed into the Editor, and the Project ID is correct.
* Check service enablement, for example, Authentication and Lobby.

### Can’t list or join sessions

* Verify network connectivity and firewall rules.
* Check Lobby service status and quotas.

#### Join code invalid or expired

* Confirm the affected session is active and not at capacity.
* Regenerate the join code by recreating the session.
* Ensure Relay is enabled and regions are available.
* Validate that endpoints are set before starting netcode.

### UI not showing data

* Check that each `VisualElement` in your UXML has its `SessionType` set properly.

## Additional resources

* [Unity Services: Multiplayer SDK](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual)
* [Unity Netcode for GameObjects](https://docs.unity3d.com/Packages/com.unity.netcode.gameobjects@latest)
* [Unity Netcode for Entities](https://docs.unity3d.com/Packages/com.unity.netcode@latest)