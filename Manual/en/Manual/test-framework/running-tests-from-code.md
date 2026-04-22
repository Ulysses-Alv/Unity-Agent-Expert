# Running tests from code

You can invoke Unity Test Framework’s **test runner** from code with the [`TestRunnerAPI`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.TestRunnerApi.html) class, which lets you control which tests are run and retrieve test data and results.

| **Topic** | **Description** |
| --- | --- |
| **[Specify which tests to run](extension-run-tests.md)** | Compose [`Filter`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.Filter.html) objects for the [`TestRunnerAPI`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.TestRunnerApi.html) to determine which tests to include in a test run. |
| **[Retrieving test results](extension-get-test-results.md)** | Register to receive callbacks from the test runner at the start or finish of an active test run or individual test. |
| **[Retrieve the list of tests](extension-retrieve-test-list.md)** | Retrieve the test tree for a given test mode (Edit mode or Play mode) from the test runner. |

## Additional resources

* [Run tests in the Test Runner window](workflow-run-test.md)
* [Run tests from the command line](run-tests-from-command-line.md)