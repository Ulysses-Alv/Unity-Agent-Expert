# Matchmaker Session Building Block

Implement Unity’s multiplayer Matchmaker in your project with this Unity Building Block’s pre-made UI elements. Learn what Matchmaker is, how to set it up, what Matchmaker queues are, and how to connect to Unity’s netcode packages.

## Introduction

The Matchmaker Session Building Block helps you integrate [Unity’s Matchmaker](https://docs.unity.com/ugs/manual/matchmaker/manual/matchmaker-overview) into your project to connect players together. This Building Block can serve as a starting point for your multiplayer projects, or as a quick integration of Matchmaker throughout development.

This Building Block demonstrates how to do the following:

* Manage a Matchmaker Queue configuration from your Unity project.
* Deploy such Queue asset to the cloud services to use it during Play mode.
* Connect players together through the Matchmaker service using a list of rules.

The Matchmaker Session Building Block includes an example **Scene**, a UIToolkit UI, and a preset of Matchmaker queue configuration with basic rules to connect players together. If you need gameplay synchronization, you can optionally add Netcode using Unity [Netcode for GameObjects](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-your-first-session) or [Netcode for Entities](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-with-netcode-for-entities).

## Prerequisites

* Install the [Unity Building Block - Matchmaker Session](https://assetstore.unity.com/packages/essentials/tutorial-projects/unity-building-block-matchmaker-session-341932) from the Unity **Asset Store**.
* [Link your Unity Editor project to a cloud project](https://docs.unity.com/en-us/cloud/projects/configure-project-for-unity-cloud).

## Initial setup

[Deploy the MatchmakerQueue](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/deployment/deploy-matchmaker-queue) asset to use the functionality in this Building Block.

1. In the Editor, under **Assets** > **Blocks** > **MatchmakerSession** > **Settings**, select the `MatchmakerQueue` asset file.
2. In the ****Inspector**** window, select **View in Deployment Window**.
3. In the **Deployment** window, select the `MatchmakerQueue` file under `Blocks.MatchmakerSession`, then right-click and select **Deploy** to set up the queue.

You can validate that the Building Block is installed correctly by opening the included sample scene:

1. Open the example scene located in the folder **Assets** > **Blocks** > **MatchmakerSession** > **Scenes**.
2. Enter Play mode and confirm that sign-in succeeds and service endpoints are reachable.

## Unity Matchmaker

[Matchmaking](https://docs.unity.com/ugs/manual/matchmaker/manual/matchmaker-overview) in Unity automatically groups players into [Multiplayer sessions](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual) based on rules involving skill, region, game mode, and party size.

### Common concepts

* Tickets/Queues: Players submit a request and wait to be matched.
* Attributes: Skill, region, input device, party size, game mode.
* Team balancing: Distributing players to achieve fair teams.
* Latency or region routing: Placing players on nearby servers.

### Typical flow

1. Player creates a join request.
2. The service places the request in a queue and searches by rules.
3. On a match, a session is created.
4. Players receive connection info and join the session.

## Included scenes

| Asset Name | Description |
| --- | --- |
| Matchmake | Connects users with compatible players based on set criteria. Settings can be added to optimize matches. |

### Matchmake

`JoinSessionByMatchmaking` UXML:

* `MatchmakeButton`: Starts a Matchmake with the given queue name.
* `CancelButton`: Cancels the current Matchmake request.
* `SessionConnectingTimerLabel`: Displays how long a Matchmake attempt has been active.

`CurrentSession` UXML:

* `SessionNameLabel`: Displays the name of the watched session.
* `PlayerCountLabel`: Displays the number of players in the watched session.
* `PlayerNameLabel`: Displays the player’s own name.
* `CopySessionCodeElement`: Displays the join code for the session so it can be copied.
* `PlayerListView`: Displays the list of players that have joined the session.
* `LeaveSessionButton`: Select to leave the joined session.

## Matchmaker queue

A [Matchmaker queue](https://docs.unity.com/ugs/manual/matchmaker/manual/advanced-topics-queues-pools) is a rules-driven waiting line where player tickets are collected and evaluated to form matches. The queue enforces your configuration: required attributes, regions, modes, team sizes, roles, and priorities, for example, fairness versus speed. The queue assembles compatible tickets into matches and handles rate limiting, returning join data when a valid match is created.

This Building Block contains a simple Matchmaker queue configuration that matches two teams from one to two players in each team and no other limitations. It also has a [relaxation rule](https://docs.unity.com/ugs/en-us/manual/matchmaker/manual/relaxations) after 20 seconds that reduces the number of teams required to one, so that a single joining player successfully joins a session after 20 seconds.

For more refined rules you can modify the already deployed Matchmaker queue on the Dashboard and have access to the full set of [adjustable Matchmaking rules](https://docs.unity.com/ugs/en-us/manual/matchmaker/manual/matchmaking-rules-rules).

## Matchmaking and Quick Join differences

[The Multiplayer Services Matchmake API](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/matchmake-session) has both Quick Join and Matchmake options. Quick Join is a simple action that finds any available lobby, while Matchmake is a rules-driven process that targets the best possible match, not just the first available one.

### Quick Join

* Fastest way to enter any open, compatible lobby.
* Minimal rules; prioritizes speed over quality.
* Best for casual, low-stakes play or small lobbies.

### Matchmaking

* Queue-based, rule-driven placement, for example, skill, region, mode, party size.
* Aims for fair teams, low latency, and proper server allocation.
* Supports backfill, cross-play filters, and dedicated servers.

Use Quick Join for immediacy. Use Matchmaking when balance, latency, role constraints, or server orchestration matter. Refer to the [Multiplayer session Building Block](building-blocks-multiplayer-sessions.md) for a Quick Join example.

## Testing matchmaking locally

To test your setup locally you can use [Multiplayer Play Mode](https://docs.unity3d.com/Packages/com.unity.multiplayer.playmode@latest).

## Add Netcode to a session

To integrate either Netcode for **GameObjects** or Netcode for Entities with your sessions, do the following:

1. Install the Netcode package of your choice using the Package Manager (**Window** > **Package Management** > **Package Manager**).
   * Install either [Unity Netcode for GameObjects](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-your-first-session#build-a-session-with-netcode-for-gameobjects) or [Netcode for Entities](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/build-with-netcode-for-entities#build-a-session-with-netcode-for-entities)
2. Enable **Create Network Session** in `SessionSettings`.
   * In your `SessionSettings` ScriptableObject, locate and enable the **Create Network Session** option, which allows your sessions to be recognized and managed by your chosen Netcode solution.
3. Change the **Network Type** in `SessionSettings`.
   * Adjust the **Network Type** within your `SessionSettings` to match your intended network architecture. You can choose from the following:
     + Relay: For cloud-based relay services.
     + Distributed authority: For using Netcode for GameObjects distributed authority.
     + Direct: For direct connections using IP Address and Port.

To understand more about the different NetworkTypes, refer to the [Network connection management](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual/manage-session-network-connection#types-of-network).

## Troubleshooting

### Multiplayer Services not available

* Multiplayer Services have to be initialized. You can initialize your own scenes with default settings by adding a **ServicesInitialization** and **PlayerAuthentication** components.

### Authentication failure

* Ensure Matchmaker is enabled in Unity Dashboard for the correct Project ID.
* Confirm you’re signed into the Editor with the right organization and environment.

### Matchmaker errors (404 or unable to match)

* Check queue deployment status. In the Editor, select the `MatchmakerQueue` asset inside **Assets** > **Blocks** > **MatchmakerSession** > **Settings** and confirm that the Inspector displays **Up to Date** with a green icon in the file header.
* Validate that the `MatchmakeButton` in your UXML references the right `MatchmakerQueue` asset.
* Verify the queue configuration by making sure that in the `MatchmakerQueue` asset, the minimum number of teams and players aligns with the number of clients you’re trying to connect. When making changes to the queue configuration in your project, you have to deploy it again through the Deployment Window for it to work during Play mode.

### UI not showing data

* Check that each `VisualElement` in your UXML has its `SessionType` set properly.

# Additional resources

* [Unity Services: Multiplayer SDK](https://docs.unity.com/ugs/en-us/manual/mps-sdk/manual)
* [Unity Netcode for GameObjects](https://docs.unity3d.com/Packages/com.unity.netcode.gameobjects@latest)
* [Unity Netcode for Entities](https://docs.unity3d.com/Packages/com.unity.netcode@latest)
* [Matchmaker Rules](https://docs.unity.com/ugs/en-us/manual/matchmaker/manual/matchmaking-rules-rules)