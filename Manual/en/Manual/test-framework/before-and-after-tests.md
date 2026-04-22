# Performing actions before and after tests

Write reliable tests by performing the appropriate pre-test setup and post-test cleanup work to ensure your tests run in the right state. Unity Test Framework has a range of APIs that allow you to perform work at defined stages before and after a test run.

| **Topic** | **Description** |
| --- | --- |
| **[Setting up and tearing down tests](reference-unitysetup-and-unityteardown.md)** | Use Unity Test Framework APIs for test setup and tear down. |
| **[Performing actions before setup or after tear down](reference-outerunitytestaction.md)** | Use the [IOuterUnityTestAction](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.IOuterUnityTestAction.html) API to do work outside the test before setup and after tear down. |
| **[Setting up and cleaning up at build time](reference-setup-and-cleanup.md)** | Make changes to Unity or the file system before building tests and then clean up these changes after the test run. |
| **[Execution order of test actions](reference-actions-outside-tests.md)** | Understand the order in which test actions run relative to one another. |

## Additional resources

* [Setup and teardown](course/setup-teardown.md)
* [Setup and cleanup at build time](course/build-setup-cleanup.md)