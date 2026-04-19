# LiveOps Building Blocks prerequisites

To use the LiveOps Unity Building Blocks, set up the following in your project:

* Link your project to the Unity Dashboard
* Set a target environment
* Set up deployment

## Link your project to the Unity Dashboard

To use Unity Cloud, follow the instructions to [Configure a project for Unity Cloud](https://docs.unity.com/en-us/cloud/projects/configure-project-for-unity-cloud).

## Set a target environment

A project that uses Unity Cloud targets an [environment](https://docs.unity.com/ugs/en-us/manual/overview/manual/service-environments) in the Cloud project. You can create and delete environments in the Unity Dashboard. To select a target environment, do the following:

1. Link your project to the Unity dashboard.
2. In the Editor, go to **Edit** > **Project Settings.**
3. Go to the **Services** > **Environments** section.
4. Follow the steps that appear to create or select an existing environment.
5. Expand the **Environment** dropdown to select an environment.

## Set up deployment

To deploy assets in the LiveOps Building Blocks to the cloud, set up deployment in your **scene**:

1. Go to **Services** > **General Settings** > **Services**.
2. In the **Project Settings** window that appears, follow the instructions to select an organization and a cloud project.

## Additional resources

* [Leaderboards](https://docs.unity.com/ugs/en-us/manual/leaderboards/manual/leaderboards)
* [Get started with Unity Player Accounts](https://docs.unity.com/ugs/en-us/manual/authentication/manual/unity-player-accounts#get-started)
* [Cloud Code Authentication.](https://docs.unity.com/ugs/en-us/manual/cloud-code/manual/modules/how-to-guides/authentication)