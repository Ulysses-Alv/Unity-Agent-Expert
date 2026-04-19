# Testing Lost Crypt

Welcome to the this training material for the Unity Test Framework (UTF).

The training is structured with a selection of exercises, starting with more basic topics and then expanding on that knowledge.

Each section has a Learning Objectives section, which can help you pick what exercises will teach you new things. The exercises are grouped thematically and their difficulty varies.

**This course focus on testing an actual game. Our candidate is the [LostCrypt](https://assetstore.unity.com/packages/essentials/tutorial-projects/lost-crypt-2d-sample-project-158673) example project.**

## Course outline

| **Topic** | **Description** |
| --- | --- |
| **[Setting up](setting-up.md)** | Set up a simple Unity 2D project and import a sample project (LostCrypt). |
| **[Running a test in LostCrypt](first-test.md)** | Set up a simple Play mode test for LostCrypt. |
| **[Moving character](moving-character.md)** | Use the Unity InputSystem package to have a generic way of moving your character programmatically in tests. |
| **[Reach wand test](reach-wand-test.md)** | Perform assertions on your character position and behavior. |
| **[Collision test](collision-test.md)** | Check for **collisions** and make sure that LostCrypt does not have bugs that allow your character to move outside the map. |
| **[Asset change test](asset-change-test.md)** | Use a common pattern in game testing to verify if assets change over time. |
| **[Scene validation test](scene-validation-test.md)** | Test the **scene** for the presence of specific game objects and make this test use all scenes as fixtures. |
| **[Performance tests](performance-tests.md)** | Extend Unity Test Framework with performance tests. |

## Additional resources

[General Introduction course](../welcome)