# Player Accounts Building Block

Use the Player Accounts Unity Building Block to add sign-in functionality to your project, for example, to track and preserve users’ progress in your application. To learn more about PLayer Accounts, refer to the [Player Accounts](https://docs.unity.com/ugs/en-us/manual/authentication/manual/unity-player-accounts) documentation.

This Building Block includes a test **scene** that demonstrates the player authentication process. You can use the **scripts** in this Building Block to authenticate players in your own project. For more information, refer to [Add the Player Accounts Building Block to a different scene](#add-the-player-accounts-building-block-to-a-different-scene).

The Player Accounts Building Block uses the following Unity Gaming Services (UGS):

| **Unity Gaming ServiceTopic** | **Description** |
| --- | --- |
| Authentication | Assigns each player a unique identity. To use these authentication flows, (set up an identity provider)[set-up-an-identity-provider].  To learn more, refer to [Authentication](https://docs.unity.com/ugs/en-us/manual/authentication/manual/get-started). |
| Cloud Save | Stores Player Data and Game Data for your games, and provides access classes to control read and write access to Player Data and protected Game Data.  To learn more, refer to [Cloud Save](https://docs.unity.com/ugs/en-us/manual/cloud-save/manual). |
| Cloud Code | Provides server authority that you can use to write to protected or private access classes. You can use Cloud Code to write to game data or protected player data, or to read private game data.  To learn more, refer to [Cloud Code Authentication](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules/how-to-guides/authentication). |

## Prerequisites

To set up a project for the Player Accounts Building Block, follow the instructions in:

* Follow the instructions in the [LiveOps Building Blocks prerequisites](building-blocks-liveops-prerequisites.md).
* Install the [Unity Building Block - Player Account](https://assetstore.unity.com/packages/essentials/tutorial-projects/unity-building-block-player-account-341928) from the Unity **Asset Store**.
* Follow the instructions in [Get started with Unity Player Accounts](https://docs.unity.com/ugs/en-us/manual/authentication/manual/unity-player-accounts#get-started).

## Initial setup

To use all features of the [Player Accounts test scene](#use-the-player-accounts-building-block-test-scene), follow the steps in this section.

### Set up an identity provider

The Player Accounts Building Block test scene demonstrates multiple sign-in methods. For your project to recognize these authentication methods, do the following:

1. [Set up the Username and Password provider.](https://docs.unity.com/ugs/manual/authentication/manual/platform-signin-username-password#set-up-the-username--password-provider)
2. [Set up a Unity Player Accounts ID provider](https://docs.unity.com/ugs/en-us/manual/authentication/manual/unity-player-accounts#set-up-the-unity-player-accounts-id-provider).

### Deploy the Cloud Code module

The test scene uses the `BlocksGameModule.ccmr` [Cloud Code module](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules) to update the global score which is the result of the values of all authenticated players added together. To use this Cloud Code module, deploy it:

1. Follow the instructions in [LiveOps Building Blocks prerequisites](building-blocks-liveops-prerequisites.md) to set up deployment in your project.
2. Go to **Services** > **Deployment**.
3. In the **Deployment** window, in the Blocks.CloudCode section, select the **BlocksGameModule.ccmr** checkbox.
4. Select **Deploy Selected**.

## Use the Player Accounts Building Block test scene

The Player Accounts Building Block test scene demonstrates the following steps in a player account workflow:

* Sign-in: Uses the [Unity Extensible Markup Language (UXML)](https://docs.unity3d.com/6000.2/Documentation/Manual/UIE-WritingUXMLTemplate.html) to display a sign-in interface that authenticates players. Authenticated players can update or view the global score. For more information, refer to [The Player Accounts test scene interface](#customize-the-player-accounts-user-interface).
* Data visualization: Displays the player data that the labels in the `PlayerAccountScene.uxml` asset declare.
* Cloud integration: Displays the shared global score that the `BlocksGameModule.ccmr` Cloud Code module manages.

To open the Player Accounts Building Block test scene:

1. In the **Project** window, go to **Assets**> **Blocks** > **PlayerAccount** > **TestScenes**.
2. Right-click the `PlayerAccountScene` asset and select **Open**.

To use the Player Accounts Building Block test scene:

1. Enter Play Mode.
2. In Play Mode, select one of the following options to sign in:
   * **Sign in Anonymously**: Creates anonymous player data to authenticate the user. To learn more, refer to [Anonymous authentication.](https://docs.unity.com/ugs/en-us/manual/authentication/manual/approaches-to-authentication#anonymous-authentication)
   * **Sign in with Unity**: Uses Unity account details to authenticate the player. To learn more about this authentication method, refer to [External authentication](https://docs.unity.com/ugs/en-us/manual/authentication/manual/approaches-to-authentication#external-authentication). To use this sign-in method, [set up a Unity Player Accounts ID provider](https://docs.unity.com/ugs/en-us/manual/authentication/manual/unity-player-accounts#set-up-the-unity-player-accounts-id-provider).
   * **Sign in with Password**: Accepts a custom username and password that meets the [requirements](https://docs.unity.com/ugs/manual/authentication/manual/platform-signin-username-password#sign-up-or-sign-in-a-returning-player). To use this sign-in method, [set up the Username & Password provider](https://docs.unity.com/ugs/manual/authentication/manual/platform-signin-username-password#set-up-the-username--password-provider).
3. In the new scene that loads, select the **Increment Global Score** button to increase the global score by a value of one. The global score is the result of the values of all authenticated players added together.

**Note**: The Authentication service stores player identity on the cloud.

## Add the Player Accounts Building Block to a different scene

To add player authentication from the Player Accounts building block to a new scene, do the following:
1. [Deploy the Cloud Code module](#deploy-the-cloud-code-module).
2. Go to **Assets** > **Blocks** > **Achievements** > **Prefabs**.
3. Select the `AchievementsPrefab` to open it.

Any changes you make to `Achievements.ach` affect the `AchievementsPrefab` when you enter Play mode.

You can also add the Player Accounts user interface assets to a different scene. To learn more, refer to [Customize the Player Accounts user interface](#customize-the-player-accounts-user-interface).

**Note**: To create a new Achievements Configuration file, go to **Create** > **Blocks** > **Achievements Configuration**.

## Customize the Player Accounts starter scene

The player accounts starter scene contains the following scripts that you can modify to change how player authentication works in the scene:

* The Player Accounts user interface scripts.
* The Player Accounts player data scripts.

### Customize the Player Accounts user interface

The Player Accounts Building Block contains Unity Extensible Markup Language (UXML) assets to display the test scene user interface (UI). To learn about UXML scripts, refer to [Introduction to UXML](https://docs.unity3d.com/6000.2/Documentation/Manual/UIE-WritingUXMLTemplate.html).

To customize the Player Accounts user interface:

1. Go to **Window** > **UI Toolkit** > **UI builder**.
2. In the ****Viewport**** **toolbar**, select the **Blocks Runtime Theme**.
3. In the **Library** panel go to **Blocks** > **PlayerAccount**.
4. Select **PlayerDataLabel** and drag it to the Hierarchy panel.
5. Modify the exposed properties in the **Attributes** panel.

To learn more about UI builder, refer to [UI Builder](UIBuilder.md).

### Player Accounts user interface assets

The Player Accounts Building Block contains the following Unity Extensible Markup Language (UXML) assets that you can modify or add to a different scene.

| **Asset name** | **Description** |
| --- | --- |
| `PlayerSignIn` | Displays authentication options that you can use to sign in anonymously, with Unity’s Player Account web portal, or with a username and password. To learn more, refer to Platform-agnostic authentication. |
| `PlayerIdLabel` | A reusable label that displays and tracks the player ID with clipboard integration. |
| `PlayerNameLabel` | A reusable, data-bound label that tracks the player name. You can edit the name in the UXML field in the test scene. |
| `PlayerDataLabel` | A reusable label that loads Public, Default, or Protected player data according to the assigned data-key attribute. |
| `PlayerEditableDataLabel` | Specialized `PlayerDataLabel` that allows editing player data (except for Protected). |
| `PlayerLoadDataControl` | Debug control that displays all available default access class keys, and loads the value of one on demand for the player. |

The Player Accounts UXML scripts use the following classes:

* [`AuthenticationService`](https://docs.unity3d.com/Packages/com.unity.services.authentication@3.5/api/Unity.Services.Authentication.AuthenticationService.html) from the [Authentication](https://docs.unity.com/ugs/manual/authentication/manual/overview) package.
* [`ServiceObserver&T`](https://docs.unity3d.com/Packages/com.unity.services.core@1.15/api/Unity.Services.Core.ServiceObserver-1.html) from the [Services Core](https://docs.unity.com/ugs/manual/overview/manual/services-core-api#servicesinitializationexception) package.

## Player Accounts player data scripts

The Player Accounts Building Block contains scripts that manage player data, such as sign-in information, and game data, for example, scores. These scripts use the [`ServiceObserver<T>`](https://docs.unity3d.com/Packages/com.unity.services.core@1.15/api/Unity.Services.Core.ServiceObserver-1.html) base class, which you can use from the Services Core package, to manage the lifetime of services without the need to track them manually.

Find the following script assets in the **Project** window: **Assets** > **Blocks** > **PlayerAccount** > **Scripts** > **Runtime**.

| **Asset name** | **Description** |
| --- | --- |
| `AuthenticationObserver` | Handles sign-in events and authentication state, and allows the project to register callbacks. |
| `CloudDataContainer` | Displays the information that `AuthenticationObserver` stores in the test scene. This container also manages asynchronous player data loading and binding that you can instantiate and access in a script. |
| `GlobalScoreClient` | A Cloud Code module file that uses a write lock mechanism to manage game-level shared data samples of a global score. |

## Troubleshooting

If the **Increment Global Score** button in the [Player Accounts test scene](#use-the-player-accounts-building-block-test-scene) doesn’t update the Global Score value, follow the instructions to [deploy the Cloud Code module](#deploy-the-cloud-code-module).

To learn more about read and write restrictions for Cloud Save Data, refer to the [Services API](https://services.docs.unity.com/cloud-save/v1/#section/General-Information) documentation.

## Additional resources

* [Username and password setup](https://docs.unity.com/ugs/manual/authentication/manual/platform-signin-username-password)
* [Deployment](https://docs.unity3d.com/Packages/com.unity.services.deployment@1.6/manual/getting_started.html)