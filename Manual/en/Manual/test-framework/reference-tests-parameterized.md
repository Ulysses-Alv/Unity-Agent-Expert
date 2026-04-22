# Parameterized tests

Unity Test Framework supports parameterized tests, which can be useful for data-driven testing.

Regular NUnit tests support both the [`[TestCase]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/testcase.html) and [`[ValueSource]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/valuesource.html) attributes for parameterized tests. Unity tests only support `ValueSource`

The following example demonstrates use of `ValueSource` with a Unity test:

```
static int[] values = new int[] { 1, 5, 6 };

[UnityTest]
public IEnumerator MyTestWithMultipleValues([ValueSource("values")] int value)
{
    yield return null;
}
```

## Ignore based on parameters

You can selectively ignore tests based on the parameters supplied to the test method by using the [ParameterizedIgnoreAttribute](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.ParametrizedIgnoreAttribute.html).

## Additional resources

* [UnityTestAttribute](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityTestAttribute.html)