# Specify which tests to run

To run tests from code, call `Execute` on the [TestRunnerApi](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.TestRunnerApi.html) and specify which tests to run with a [Filter](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEditor.TestTools.TestRunner.Api.Filter.html).

The following example runs all Play mode tests in a project:

```
var testRunnerApi = ScriptableObject.CreateInstance<TestRunnerApi>();
var filter = new Filter()
{
    testMode = TestMode.PlayMode
};
testRunnerApi.Execute(new ExecutionSettings(filter));
```

## Multiple filter values

You can define additional fields on the `Filter` class to create a more specific filter.

Many of the fields allow for multiple values. When you define multiple values for a field, the **test runner** runs any tests that match at least one of the values.

The following example runs tests with full names that match either of the two names provided:

```
var api = ScriptableObject.CreateInstance<TestRunnerApi>();
api.Execute(new ExecutionSettings(new Filter()
{
    testNames = new[] {"MyTestClass.NameOfMyTest", "SpecificTestFixture.NameOfAnotherTest"}
}));
```

## Multiple filter fields

If you define multiple fields on a filter, the runner only runs tests that match all the fields.

In this example, it runs any test in the specified assembly with a name that matches either of the two specified test names.

```
var api = ScriptableObject.CreateInstance<TestRunnerApi>();
api.Execute(new ExecutionSettings(new Filter()
{
    assemblyNames = new [] {"MyTestAssembly"},
    testNames = new [] {"MyTestClass.NameOfMyTest", "MyTestClass.AnotherNameOfATest"}
}));
```

## Multiple constructor filters

The execution settings take one or more filters in its constructor. If there is no filter provided, then it runs all Edit mode tests by default. If there are multiple filters provided, then a test runs if it matches any of the filters.

The following runs all tests in the assembly named `MyTestAssembly` any test with a name that matches either of the two specified test names:

```
var api = ScriptableObject.CreateInstance<TestRunnerApi>();
api.Execute(new ExecutionSettings(
    new Filter()
    {
        assemblyNames = new[] {"MyTestAssembly"},
    },
    new Filter()
    {
        testNames = new[] {"MyTestClass.NameOfMyTest", "MyTestClass.AnotherNameOfATest"}
    }
));
```

**Note**: Specifying different test modes or platforms in each `Filter` is not currently supported. The test mode and platform is from the first `Filter` only and defaults to Edit Mode, if not supplied.

## Additional resources

* [Retrieve test results](extension-get-test-results.md)
* [Retrieve the list of tests](extension-retrieve-test-list.md)