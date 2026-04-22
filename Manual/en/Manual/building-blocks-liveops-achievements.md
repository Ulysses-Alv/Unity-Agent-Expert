# Achievements Building Block

The Achievements Unity Building Block demonstrates an achievement system that uses Unity Gaming Services (UGS).

The achievements are synchronized and persist in the Unity cloud, which means you can access them on any other device or platform that runs the same project.

The Achievements Building Block uses the following Unity Gaming Services:

| **Unity Gaming Service** | **Description** |
| --- | --- |
| **Cloud Save** | Stores data about a player’s achievement status as [Player Data](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual/concepts/player-data) and [Game Data](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual/concepts/game-data). For more information, refer to [Cloud Save overview](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual). To learn how the Achievements Building Block uses the Cloud Save service, refer to [How the Achievements Building Block saves data](#how-the-achievements-building-block-saves-data). |
| **Cloud Code** | Performs actions with server authority to claim achievements. For more information about Cloud Code, refer to [Cloud Code overview](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual). This Building Block uses [Cloud Code modules](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to interact with the Cloud Save service. To learn more, refer to [Deploy the Cloud Code modules](#deploy-the-cloud-code-modules). |
| **Authentication** | Assigns each player a unique identity. This Building Block uses Authentication to sign in anonymously. To use it, [set up authentication](https://docs.unity.com/ugs/en-us/manual/authentication/manual/get-started). |
| **Remote Config** | Stores values as key-value pairs on the cloud as a JSON file. For more information, refer to [Remote Config overview](https://docs.unity.com/ugs/en-us/manual/remote-config/manual/WhatsRemoteConfig). This Building Block uses a remote config file (.rc) to define the achievements on the cloud. To learn more, refer to [Customize the Achievements Building Block](#customize-the-achievements-building-block). |
| **Access Control** | Sets authentication requirements for interacting with other services to prevent cheating. This Building Block uses Access Control to restrict players from writing to Cloud Save.  To write to Cloud Save, use the Trusted client which uses Cloud Code Modules to authenticate calls as a trusted source. |

## Prerequisites

To set up a project for the Achievements Building Block:

* Install the [Unity Building Block - Achievements](https://assetstore.unity.com/packages/essentials/tutorial-projects/unity-building-block-achievements-341918) from the Unity **Asset Store**.
* Follow the instructions in the [LiveOps Unity Building Blocks prerequisites](building-blocks-liveops-prerequisites.md).

## Initial setup

To use all features of the [Achievements test scene](#use-the-achievements-building-block-test-scene), perform the following actions.

### Deploy the configuration file

The test **scene** includes the `Achievement.ach` configuration file to define the achievements in the scene.

1. [Set up deployment in your project](building-blocks-liveops-prerequisites.md).
2. Go to **Services** > **Deployment**.
3. In the **Deployment** window, expand the **Blocks.Achievements** section.
4. Select the **Achievement.ach** checkbox.
5. Select **Deploy Selected.**

### Deploy the cloud code modules

The test scene uses the BlocksGameModule.ccmr [Cloud Code module](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to update the global score. To use this module:

1. [Set up deployment in your project](building-blocks-liveops-prerequisites.md).
2. Go to **Services** > **Deployment**.
3. In the **Deployment** window, expand the **Blocks.CloudCode** section.
4. Select the **BlocksGameModule.ccmr** and **BlocksAdminModule.ccmr** checkboxes.
5. Select **Deploy Selected**.

## Use the Achievements Building Block test scene

The Achievements Building Block test scene demonstrates in-game achievements that you can claim, rename, and reset. Each achievement exists in one of the following states:

* Unlocked
* Locked
* Hidden
* In progress

To open the test scene:

1. In the Project window, go to **Assets** > **Blocks** > **Achievements** > **TestScenes.**
2. Right-click the `AchievementsTestScene.unity` asset and select **Open**.

To use the test scene:

1. Enter Play Mode. During the loading screen the scene retrieves achievements defined in the [configuration file](#modify-the-configuration-file) from Cloud Save.
2. The scene loads and displays the achievements.
3. Select an achievement’s **Unlock** button to claim it, or use the **+** and **-** buttons to change its progress.
4. Select **Reset Achievements** to restore achievements to their default state.

## How the Achievements Building Block saves data

The Achievements Building Block uses the [Cloud Save](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual) service to store and track achievement progress for each player.

Each player is assigned a key and value pair saved to a protected store on the Cloud Save service. Player data is stored in the achievements key.

The AchievementDefinition data model is defined in the following locations:

* To find the local `AchievementDefinition.cs` file, go to **Assets** > **Blocks** > **Achievements** > **Scripts** > **Runtime** >
* To find the Cloud Code file:
  1. In the **Assets** folder, right-click **CloudCode.**
  2. Select **View in Explorer** on Windows or **View in Finder** on macOS.
  3. Go to **BlocksGameModule~** > **Project** > **Achievement** > `AchievementDefinition.cs`.

**Note**: The actual data stored is an array of objects of type `AchievementRecord`.

The following script is an example of the player data that the Achievements Building Block stores on the Cloud Save service:

```
[
{
"Id": "first_achievement", // the ID of the achievement
"Unlocked": true, // the unlocked status of the achievement for the current player
"ProgressCount": // current progress relative to the ProgressTarget of the definition
}
]
```

## Add the Achievements prefab to a different scene

To use the Achievements Building Block in a new scene:

1. Deploy the [Achievements Configuration file](#deploy-the-configuration-file).
2. Go to **Assets** > **Blocks** > **Achievements** > **Prefabs.**
3. Select the `AchievementsPrefab` to open it.

Any changes made to `Achievements.ach` affect the `AchievementsPrefab` in Play Mode. To learn how to modify `Achievements.ach`, refer to [Customize the Achievements Building Block](#customize-the-achievements-building-block).

**Note:** To create a new Achievements Configuration file, go to **Create** > **Blocks** > **Achievements Configuration**.

## Customize the Achievements Building Block

You can customize the achievements building block in the following ways:

* Modify the Achievements configuration file.
* Modify the Achievements **prefab**.
* Set an access control policy.
* Change achievement data.

### Modify the configuration file

The `Achievements.ach` file is a type of [Remote Config](https://docs.unity.com/ugs/en-us/manual/remote-config/manual/remote-config-files) (.rc) file that exposes achievement configuration properties in the **Inspector** window. The `Achievements.ach` file must match the data defined in the `AchievementDefinition` file. For more information, refer to [Change achievement data](#change-achievement-data).

To change descriptions or appearance of each achievement:

1. In the Project window, go to **Assets** > **Blocks** > **Achievements** > **Deployment** > **Achievements.ach**.
2. Select **Achievements.ach** to display its properties in the Inspector.
3. To apply changes, [deploy Achievements.ach to the current environment](#deploy-the-cloud-code-modules).

### Modify the Achievements prefab

To make changes to the Achievements prefab:

1. Go to the **Project window**.
2. Navigate to **Assets** > **Blocks** > **Achievements** > **Prefabs.**
3. Select `AchievementsPrefab` to view its properties in the Inspector.

You can also enable **Developer Mode** to change the progress of incremental achievements in the **Scene view**.

The following classes control the data displayed by `Achievements.ach`:

* `AchievementAsset`: Scriptable object in the project.
* `AchievementFile`: Defines how the contents of `Achievements.ach` are displayed.
* `AchievementAssetImporter`: Custom importer class enabling the custom inspector.
* `AchievementAssetInspector`: The custom inspector itself.

### Set an Access Control policy

An achievement client submits scores to the Achievement service. The block includes the following clients:

* **Local** (`LocalAchievementsClient`): Submits scores directly to Cloud Save without validation. You can use an [Access Control](https://docs.unity.com/ugs/manual/overview/manual/access-control) policy to prevent cheating.
* **Trusted** (`TrustedAchievementsClient`): Calls the [Cloud Code C# module](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to submit scores securely.

To change the client in the test scene:

1. Open the [Achievement test scene](#use-the-achievements-building-block-test-scene).
2. In the Hierarchy window, select the `AchievementsPreview` **GameObject**.
3. In the Inspector, enable **Use Trusted Client**.

### Change Achievement data

The `AchievementDefinition` files define achievement properties used in the Achievements.ach Inspector window. To learn more about the `Achievements.ach` file, refer to [Modify the configuration file](#modify-the-configuration-file).

Locations:

* Local: **Assets** > **Blocks** > **Achievements** > **Scripts** > **Runtime** > `AchievementDefinition.cs`
* Cloud Code: **Assets** > **Blocks** > **Achievements** > **CloudCode~** > **Project** > **Achievement** > `AchievementDefinition.cs`

To change the data in this file:

* Update the `AchievementDefinition.cs` file in each location so that their content is exactly the same.
* Deploy both files in the **Deployment** window. To learn more, refer to [Deploy the configuration file](#deploy-the-configuration-file).

## Troubleshooting

If scores aren’t visible and a warning message appears, it might be because the `AchievementsAccessControl.ac` file denies write access for the local client’s write access. To fix this issue:

1. Go to **Services** > **Deployment**.
2. Right-click the `AchievementsAccessControl.ac` file.
3. Select **Delete Remote**.

## Additional resources

* [Server Authority](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/server-authority)
* [Deployment](https://docs.unity3d.com/Packages/com.unity.services.deployment@1.6/manual/getting_started.html)