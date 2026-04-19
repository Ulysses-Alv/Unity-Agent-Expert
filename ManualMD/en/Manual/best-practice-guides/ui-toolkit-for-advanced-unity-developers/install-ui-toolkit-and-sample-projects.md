# Install UI Toolkit and sample projects

Table of Contents

* [UI Toolkit Sample – Dragon Crashers](#ui-toolkit-sample-dragon-crashers)
* [QuizU](#quizu)

UI Toolkit is integrated into the core Unity 6 platform, which means that you don’t need to install a separate package to use it with version Unity 6 and later. Starting a new project from one of the templates available will be sufficient to be able to follow the content of this guide.

![UIElements is the namespace for UI Toolkit, UI Builder and their features, all of which are now included in Unity 6.](../../../uploads/Main/bpg/uitk/image289.png)

UIElements is the namespace for UI Toolkit, UI Builder and their features, all of which are now included in Unity 6.

This e-book primarily uses the following samples to show and explain UI Toolkit capabilities in Unity projects. Each sample is available to download for free from the Unity **Asset Store**.

## UI Toolkit Sample – Dragon Crashers

[This demo](https://assetstore.unity.com/packages/essentials/tutorial-projects/dragon-crashers-ui-toolkit-sample-project-231178) uses the latest UI Toolkit workflow at runtime for a full-featured interface, including a front-end menu system, over a slice of the 2D project [*Dragon Crashers*](https://assetstore.unity.com/packages/essentials/tutorial-projects/dragon-crashers-urp-2d-sample-project-190721), a mini-RPG.

**This demo is not meant for beginners.** It was created for experienced Unity developers who have the capabilities to look at the UI structure and navigate the demo to observe specific implementations. This demo was originally released for Unity 2021 LTS and has since been [updated to Unity 6](https://www.youtube.com/watch?v=RVp3-D2nEEg). Here are some of the topics you can learn more about in the demo:

* Project structure and naming conventions
* Use of themes to create UI variations and add support for both portrait and landscape orientations
* Complex elements like tabbed menus, inventories, messages, or custom controls
* Use of the **SafeAreaAPI** to ensure content on mobile screens is displayed within the safe area
* Use of Localization for multiple language support
* Data Binding for simplifying the synchronization of data with UI components
* Examples of how to implement common casual game interfaces

You can find a [video walkthrough](https://www.youtube.com/watch?v=XtQist-I3Xo) of the sample and [download](https://assetstore.unity.com/packages/essentials/tutorial-projects/dragon-crashers-ui-toolkit-sample-project-231178) it from the Asset Store.

![The home screen can be displayed in landscape and portrait](../../../uploads/Main/bpg/uitk/image27.png)

The home screen can be displayed in landscape and portrait

## QuizU

The [QuizU demo](https://assetstore.unity.com/packages/essentials/tutorial-projects/quizu-a-ui-toolkit-sample-268492) showcases an interactive quiz game built with Unity’s UI Toolkit. Aimed at UI developers, this project highlights UI Toolkit workflows, event-driven architecture, and reusable design patterns for building modern game user interfaces. This demo shows how to:

* Structure UI elements efficiently using UXML files and nested **visual trees**.
* Apply styling rules using USS selectors and pseudo-classes to make your interactive elements react to user input.
* Use the FlexBox feature for flex-based layouts for responsive UI behavior.
* Query and modify UI elements dynamically using selectors.
* Encapsulate event handling in reusable classes, enabling custom interactions like dragging or multi-touch gestures.
* Use Event Dispatch to process events in phases and how to manage propagation.
* Use USS Transitions to add smooth animations and effects to UI elements with properties like duration and easing.

![A screen shot from the QuizU UI Toolkit demo](../../../uploads/Main/bpg/uitk/image4.png)

A screen shot from the QuizU UI Toolkit demo

Originally released for Unity 2022 LTS, QuizU has been updated with new features in Unity 6. The project now includes how-to demos on creating custom controls, setting up data binding, and implementing localization.

You can [download](https://assetstore.unity.com/packages/essentials/tutorial-projects/quizu-a-ui-toolkit-sample-268492) the project from the Asset Store.