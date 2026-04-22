# Leaderboards Building Block

The Leaderboards Unity Building Block demonstrates score-keeping in a game and generates a leaderboard to display player scores. Each score is stored as player data on the cloud.

The Leaderboards Building Block uses the following Unity Gaming Services (UGS):

| **Unity Gaming Service** | **Description** |
| --- | --- |
| [Leaderboards](https://docs.unity.com/ugs/en-us/manual/leaderboards/manual/leaderboards) | Stores, retrieves, and displays your player scores. |
| [Authentication](https://docs.unity.com/ugs/en-us/manual/authentication/manual/get-started) | Provides a unique identity to the player. In the [Leaderboards test scene](#use-the-leaderboards-building-block-test-scene), this service allows users to sign in anonymously. |
| [Access Control](https://docs.unity.com/ugs/en-us/manual/overview/manual/access-control) | Determines the authentication requirements for interacting with other services. This Building Block uses Access Control to stop players from writing data to Leaderboards. To write to Cloud Save, use the Trusted client which uses Cloud Code Modules to authenticate the calls as a trusted source. For more information, refer to [Change the client in the Leaderboards test scene](#test-scene-client). |
| [Cloud Save](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual) | Stores each score as [Player Data](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual/concepts/player-data) and [Game Data](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual/concepts/game-data). |
| [Cloud Code](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual) | The Cloud Code service allows you to create code that runs in the cloud. This code can authenticate a player as a trusted server or a client depending on your needs and interact with other services. This Building Block uses Cloud Code to demonstrate server-authoritative actions to store and retrieve player scores. The Leaderboards Building Block uses the following [Cloud Code C# modules](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules):  `Assets/CloudCode/BlocksAdminModule.ccmr` `Assets/CloudCode/BlocksGameModule.ccmr` |

## Prerequisites

To set up a project for the Achievements Building Block:

* Install the [Unity Building Block - Leaderboards](https://assetstore.unity.com/packages/essentials/tutorial-projects/unity-building-block-leaderboards-341926) from the Unity **Asset Store**.
* Follow the instructions in the [LiveOps Unity Building Blocks prerequisites](building-blocks-liveops-prerequisites.md).

## Initial setup

To use all features of the [Leaderboards test scene](#use-the-leaderboards-building-block-test-scene), perform the following actions.

### Deploy the Cloud Code modules

The test **scene** uses the `BlocksGameModule.ccmr` [Cloud Code module](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to update the global score. To use this Cloud Code module, deploy it in the scene:

1. [Set up deployment in your project](building-blocks-liveops-prerequisites.md).
2. Go to **Services** > **Deployment**.
3. In the **Deployment** window, expand the **Blocks.CloudCode** section.
4. Select the **BlocksGameModule.ccmr** and **BlocksAdminModule.ccmr** checkboxes.
5. Select **Deploy Selected**.

### Deploy the Leaderboards configuration file

The test scene uses a configuration file to add data to a leaderboard. The Leaderboards test scene includes the `simple_leaderboards.lb` configuration file. To display a populated leaderboard in the test scene, deploy this file:

1. Go to **Services** > **Deployment**.
2. In the **Deployment** window, select the `simple_leaderboard.lb` checkbox.
   If you want to use a **Trusted** game client, also select `LeaderboardsAccessControl.ac`.
3. Select **Deploy Selected**.

To learn about Leaderboard configuration on the Unity Dashboard, refer to [Create a leaderboard](https://docs.unity.com/ugs/manual/leaderboards/manual/configuration/unity-dashboard#create-a-leaderboard).

## Use the Leaderboards Building Block test scene

The Leaderboards Building Block test scene contains a guessing game that generates a score based on how close your guess is to a randomly generated number. The result of this game appears on a leaderboard. You can also generate other random player scores.

To open the Leaderboards test scene:

1. In the **Project** window, go to **Assets** > **Blocks** > **Leaderboards** > **TestScenes**.
2. Select the **Leaderboards (Scene Asset)** asset.
3. In the ****Inspector**** window, select **Open**.

To use each element in the Leaderboards test scene:

1. Enter Play mode.
2. Enter a number in the **Your Guess** field.
3. Select **Guess** to display your score and submit it to the leaderboard.
4. Select **Populate Leaderboard** to add other random user scores to the leaderboard and rank them.
5. Select **Show Leaderboard** to display the leaderboard.
6. Select one of the leaderboard tabs to choose which scores to view:
7. **Global:** Displays the top 50 scores in order.
8. **Self:** Displays the current player score and the 10 scores above and below it.
9. Select **Hide Leaderboard** to return to the game.

To reset the Leaderboard scores:

1. Open the Deployment window (**Services** > **Deployment**).
2. Right-click the `leaderboard.lb` and select **Reset Leaderboard**.

## Add the Leaderboard Prefab to a new scene

To use the Leaderboards Building Block in a new scene, add the following **prefabs** from the Leaderboards Building Block to your scene:

* `LeaderboardPrefab.prefab` adds an empty leaderboard, and displays scores that the client fetches from the Leaderboard service. Find it in the **Project** window: **Assets** > **Blocks** > **Leaderboards** > **Prefabs**
* `UnityServices.prefab` in **Assets** > **Blocks** > **Common** > **Prefabs** to initiate Unity Gaming Services.

To use the Leaderboard prefab:

1. [Set up your project for Leaderboards](#initial-setup).
2. Open the Unity Services prefab.
3. Right-click in the **Project** window and select **Create** > **Services** > **Leaderboard configuration**.
4. [Deploy the configuration file](#deploy-the-leaderboards-configuration-file).

## Customize the Leaderboards Building Block

You can change the content, description, and appearance of the Achievements Building Block, and the type of client the Leaderboard uses.

### Modify the Leaderboard configuration file

You can change the rules for the leaderboard in the `simple_leaderboard.lb` configuration file:

1. In the Project window, go to **Assets** > **Blocks** > **Leaderboards** > **Deployment**.
2. Select the `simple_leaderboard.lb` file to display its properties in the Inspector window.

Toggle each **Advanced Setting** to apply the following leaderboards concepts:

* [Buckets](https://docs.unity.com/ugs/en-us/manual/leaderboards/manual/concepts/buckets)
* [Tiers](https://docs.unity.com/ugs/en-us/manual/leaderboards/manual/concepts/tiers)
* [Resets](https://docs.unity.com/ugs/en-us/manual/leaderboards/manual/concepts/resets)

To synchronize changes to the configuration file on the cloud, [deploy it](#deploy-the-leaderboards-configuration-file).

### Change the client in the Leaderboards test scene

A leaderboard client submits scores to the Leaderboard service. The Leaderboards Building Block includes the following leaderboard clients:

* **Local** (`LocalNumberGameClient`): The local client submits scores directly to the Leaderboards service and doesn’t validate each score. You can use an [Access Control](https://docs.unity.com/ugs/manual/overview/manual/access-control) policy to stop players from cheating by submitting their own scores.
* **Trusted** (`TrustedNumberGameClient`): The remote client calls the [Cloud Code C# module](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to submit the score.

To change the client the leaderboard uses:

1. Open the [Leaderboard test scene](#use-the-leaderboards-building-block-test-scene).
2. In the Hierarchy window, select the **NumberGame** **GameObject**.
3. Expand the **Game Client Type** dropdown.
4. Select **Local** or **Trusted**.

### Set an access control policy

The `LeaderboardsAccessControl.ac` file is an access control asset that prevents users from cheating by adding their own scores to a leaderboard. It uses player authentication to deny unauthorized access to the leaderboard service.
For more information, refer to [Access Control for Unity Gaming Services (UGS)](https://docs.unity.com/ugs/manual/overview/manual/access-control).

To use an access control asset, set the leaderboard client to **Trusted**.
For more information about access control, refer to [How to use Access Control](https://docs.unity.com/ugs/manual/overview/manual/access-control#how-to-use-access-control).

## Troubleshooting

If the score doesn’t appear on the client and a warning message appears, this indicates that the local client isn’t working correctly. This is because the <LeaderboardsAccessControl.ac> file denies write access to players.

To fix this issue:

1. Go to **Services** > **Deployment**.
2. Right-click the `LeaderboardsAccessControl.ac` file.
3. Select **Delete Remote**.

## Additional resources

* [Deployment](https://docs.unity3d.com/Packages/com.unity.services.deployment@1.6/manual/getting_started.html)
* [UGS Leaderboards](https://docs.unity.com/ugs/manual/leaderboards/manual/configuration/unity-dashboard)