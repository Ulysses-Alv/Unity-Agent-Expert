# Retrieve the list of tests

You can use the [`TestRunnerApi`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.TestRunnerApi.html) to retrieve the test tree for a given test mode (Edit mode or Play mode). You can retrieve the test tree by invoking `RetrieveTestList` with the desired `TestMode` and a callback action, with an [`ITestAdaptor`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.ITestAdaptor.html) representing the test tree.

## Example

The following example retrieves the test tree for Edit Mode tests and prints the number of total test cases:

```
var api = ScriptableObject.CreateInstance<TestRunnerApi>();
api.RetrieveTestList(TestMode.EditMode, (testRoot) =>
{
    Debug.Log(string.Format("Tree contains {0} tests.", testRoot.TestCaseCount));
});
```

## Additional resources

* [Specifying which tests to run](extension-run-tests.md)
* [Retrieve test results](extension-get-test-results.md)