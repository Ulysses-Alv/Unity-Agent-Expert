# Use the Multiplayer Center

The Multiplayer Center provides a starting point to create a multiplayer game. It recommends Unity multiplayer packages based on the needs of your game, and gives you access to samples and tutorials to help you use them.

To use the Multiplayer Center:

1. Open the Unity Multiplayer Center window **Window** > **Multiplayer** > **Multiplayer Center**.
2. Select **Game Specifications** to generate recommended packages and solutions in the **Recommendation** tab.
3. Install the packages that the Multiplayer Center recommends.
4. Follow the examples and resources in the **Quickstart** tab to use the packages you installed.

## Generate a list of recommended packages for a multiplayer game

To generate a list of packages that meet the needs of the multiplayer game you want to make:

1. Select a value in each **Game Specification** field.
2. Examine the packages that the Multiplayer Center suggests are the best solutions for your game.

Hover over each package name in the **Recommended Multiplayer Packages** section to learn why each package is useful for the game specifications you selected.

The recommendations that appear in the **Recommendation** tab change immediately when you change any of the **Game Specifications** properties.

You can select which packages to install or ignore. However, you can’t deselect packages that are essential for the netcode solution or hosting model you selected.

## Install recommended packages

To install the packages that the Multiplayer Center recommends, select **Install Packages**. Unity then installs these packages and their dependencies in your project and automatically opens the **Quickstart** tab.
To remove a package from the install process, deselect the checkbox next to the package name before you select **Install Packages**.

When a tick appears next to a package name, hover over it to learn which version of this package exists in your project.

To install an individual package:

1. Select the **Open Package Manager** icon next to the package you want to install.
2. In the Package Manager window that appears, select **Install**.

**Important**: If the Multiplayer Center detects that any other multiplayer package is already installed, a warning dialog appears to cancel the installation of those packages. To avoid breaking changes, back up your project before you continue. If you continue with the installation, Unity installs or upgrades the selected packages but doesn’t remove existing packages. This can cause compatibility issues. To fix them, remove the conflicting packages from your project in the [Package Manager window](https://docs.unity3d.com/Manual/upm-ui.html).

## Get started with the recommended packages

The **Quickstart** tab includes guidance, examples, and links to resources about the packages that the Multiplayer Center installs. Follow this guidance to set up and use these packages. The content in the **Quickstart** tab might change based on the recommendations you selected in the **Recommendation** tab.

The Quickstart tab organizes this guidance into categories. Follow the instructions in each category from top to bottom to set up each package in your project.