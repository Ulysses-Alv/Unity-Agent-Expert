# Searching in the Unity Editor introduction

There are two types of search functionality available:

* [Unity Search](search-overview.md) (formerly QuickSearch)- more advanced search features including performing actions on the results. To open choose **Edit** > **Search All**, or in **Preferences** [set the Search Engines](Preferences.md#search) to **Default**.
* Legacy Search - the original search functionality for Unity, described below.

Legacy search functionality allows you to search the **Scene** view and Hierarchy window for **GameObjects**, and the **Project window** for assets.

Each search bar has the same basic workflow. To enter a search term, select the textbox and type a word or phrase. While you type, Unity automatically searches for those terms and provides the results.

Some search bars allow you to search by different criteria, such as name, type, or label. To select the criteria you want to search, select the magnifying glass icon on the bar to open the Search drop-down menu, and choose which criteria to use.

To clear your search and return your GameObject or Asset list view to normal, empty the search bar or select the small x icon in the search bar.

## Search for GameObjects in the Scene view and Hierarchy window

Both the **Scene view** and the Hierarchy window have a search bar that allows you to search for GameObjects by name or type. The Scene view and Hierarchy window represent the same information, and when you use one search bar, Unity automatically populates the other search bar with the same text.

A search term acts as a filter for the GameObjects in the Scene view and Hierarchy Window.

* In the Scene view, GameObjects that match the search term appear in color. All other GameObjects appear in gray.
* In the Hierarchy window, GameObjects that match the search term appear in alphabetical order. All the other GameObjects disappear from view.

![Scene and Hierarchy views with search filtering applied](../uploads/Main/SearchingInEditor-Scene.png)

Scene and Hierarchy views with search filtering applied

## Search for assets in the Project view

The Project window has a search bar that allows you to search for assets by name, type, or label. Select the magnifying glass in the search box to define which criteria to search by.

A search term acts as a filter for the assets in the Project window. Assets that match the search term appear in alphabetical order. Assets that do not match the search term do not appear in the results.

### Asset labels

A label is a short piece of text that you can use to group particular assets. To view the menu of existing labels, navigate to the Project window header and select the **Search by Label** button (the label icon). Select a label from the menu to search the project for any assets that have that label assigned to them.

To add a label to an asset, follow these steps:

1. Open the asset in the ****Inspector**** and find the preview panel at the bottom. This window might be minimized; if this is the case, click and drag the title bar upwards to expand it.

   ![The label icon found in the preview panel](../uploads/Main/SearchingInEditor-Inspector.png)

   The label icon found in the preview panel
2. To open the **label menu**, click the **label icon** shown on the bottom right of this panel.
3. To add an existing label to the asset, select a label from this menu. Labels currently applied to the asset have a check mark next to them.

   ![Adding a new label to an asset](../uploads/Main/SearchingInEditor-Labels.png)

   Adding a new label to an asset
4. To add a new label, you first need to create one. Type your new label name into the text box shown at the top of the menu and press either Space or Enter. This creates a label which you can then add to any asset, and use to search in the Project window. Assets can have as many labels as desired and belong to several different label groups at once.

To remove a label from an asset, open the label menu and select it to remove the tick icon.