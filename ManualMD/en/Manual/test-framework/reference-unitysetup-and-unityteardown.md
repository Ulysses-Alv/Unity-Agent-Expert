# Setting up and tearing down tests

The [`[UnitySetUp]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnitySetUpAttribute.html) and [`[UnityTearDown]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityTearDownAttribute.html) attributes are equivalent to the NUnit [`[SetUp]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/setup.html) and [`[TearDown]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/teardown.html) attributes, with the exception that they allow for [yielding instructions](reference-custom-yield-instructions.md) for the Unity Editor. The `[UnitySetUp]` and `[UnityTearDown]` attributes expect a return type of [IEnumerator](https://docs.microsoft.com/en-us/dotnet/api/system.collections.ienumerator?view=netframework-4.8).

The [`[UnityOneTimeSetUp]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityOneTimeSetUpAttribute.html) and [`[UnityOneTimeTearDown]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityOneTimeTearDownAttribute.html) attributes are equivalent to the NUnit [`[OneTimeSetUp]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/onetimesetup.html) and [`[OneTimeTearDown]`](https://docs.nunit.org/articles/nunit/writing-tests/attributes/onetimeteardown.html) attributes, with the exception that they allow for [yielding instructions](reference-custom-yield-instructions.md) for the Unity Editor. The `[UnityOneTimeSetUp]` and `[UnityOneTimeTearDown]` attributes expect a return type of [IEnumerator](https://docs.microsoft.com/en-us/dotnet/api/system.collections.ienumerator?view=netframework-4.8).

For more information and usage examples, refer to the respective API references for [`[UnitySetUp]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnitySetUpAttribute.html), [`[UnityTearDown]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityTearDownAttribute.html), [`[UnityOneTimeSetUp]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityOneTimeSetUpAttribute.html), and [`[UnityOneTimeTearDown]`](https://docs.unity3d.com/Packages/com.unity.test-framework@latest/index.html?subfolder=/api/UnityEngine.TestTools.UnityOneTimeTearDownAttribute.html).

## Execution order

`[UnitySetUp]` and `[UnityTearDown]` can be used with either the `[Test]` or `[UnityTest]` test attributes. In both cases the relative execution order of Unity and non-Unity `[SetUp]` and `[TearDown]` attributes is the same. The only difference is that a `[UnityTest]` allows for yielding instructions during the test that can result in a domain reload, in which case the non-Unity `[SetUp]` methods are re-run before proceeding to the second part of the test.

![The order of execution for setup and teardown event callbacks, with those that re-run on domain reload differentiated from those that dont.](../../uploads/Main/test-framework/execution-order-unitysetup-teardown.png)

The order of execution for setup and teardown event callbacks, with those that re-run on domain reload differentiated from those that don’t.

## Base and derived classes

The term **base** in the execution order denotes a base class from which a test class inherits. `UnitySetUp` and `UnityTearDown` follow the same pattern as NUnit `SetUp` and `TearDown` attributes in determining execution order between base classes and their derivatives. `SetUp` methods are called on base classes first, and then on derived classes. `TearDown` methods are called on derived classes first, and then on the base class. For more information, refer to the [NUnit Documentation](https://docs.nunit.org/articles/nunit/technical-notes/usage/SetUp-and-TearDown.html).

The following example demonstrates a base and derived class. You can verify the execution order through the order of messages printed in the console:

```
    public class BaseClass
    {
        [OneTimeSetUp]
        public void OneTimeSetUp()
        {
            Debug.Log("OneTimeSetUp Base");
        }

        [SetUp]
        public void SetUp()
        {
            Debug.Log("SetUp Base");
        }

        [UnitySetUp]
        public IEnumerator UnitySetUp()
        {
            Debug.Log("UnitySetup Base");
            yield return null;
        }

        [TearDown]
        public void TearDown()
        {
            Debug.Log("TearDown Base");
        }

        [UnityTearDown]
        public IEnumerator UnityTearDown()
        {
            Debug.Log("UnityTearDown Base");
            yield return null;
        }
    }

    public class DerivedClass: BaseClass
    {
        [OneTimeSetUp]
        public new void OneTimeSetUp()
        {
            Debug.Log("OneTimeSetUp");
        }

        [SetUp]
        public new void SetUp()
        {
            Debug.Log("SetUp");
        }

        [UnitySetUp]
        public new IEnumerator UnitySetUp()
        {
            Debug.Log("UnitySetup");
            yield return null;
        }

        [Test]
        public void UnitTest()
        {
            Debug.Log("Test");
        }

        [UnityTest]
        public IEnumerator UnityTest()
        {
            Debug.Log("UnityTest before yield");
            yield return null;
            Debug.Log("UnityTest after yield");
        }

        [TearDown]
        public new void TearDown()
        {
            Debug.Log("TearDown");
        }

        [UnityTearDown]
        public new IEnumerator UnityTearDown()
        {
            Debug.Log("UnityTearDown");
            yield return null;
        }

        [OneTimeTearDown]
        public void OneTimeTearDown()
        {
            Debug.Log("OneTimeTearDown");
        }
    }
```

## Domain reload

Edit mode tests can [yield instructions](reference-custom-yield-instructions.md) that result in a domain reload. On domain reload, all non-Unity actions (such as `OneTimeSetup` and `Setup`) are rerun before the code that initiated the domain reload continues. Unity actions (such as `UnitySetup`) are not rerun. If the Unity action is the code that initiated the domain reload, then the rest of the code in the `UnitySetup` method runs after the domain reload.

The following example demonstrates base and derived classes with a domain reload:

```
    public class BaseClass
    {
        [OneTimeSetUp]
        public void OneTimeSetUp()
        {
            Debug.Log("OneTimeSetUp Base");
        }

        [SetUp]
        public void SetUp()
        {
            Debug.Log("SetUp Base");
        }

        [UnitySetUp]
        public IEnumerator UnitySetUp()
        {
            Debug.Log("UnitySetup Base");
            yield return null;
        }

        [TearDown]
        public void TearDown()
        {
            Debug.Log("TearDown Base");
        }

        [UnityTearDown]
        public IEnumerator UnityTearDown()
        {
            Debug.Log("UnityTearDown Base");
            yield return null;
        }
    }

    public class DerivedClass: BaseClass
    {
        [OneTimeSetUp]
        public new void OneTimeSetUp()
        {
            Debug.Log("OneTimeSetUp");
        }

        [SetUp]
        public new void SetUp()
        {
            Debug.Log("SetUp");
        }

        [UnitySetUp]
        public new IEnumerator UnitySetUp()
        {
            Debug.Log("UnitySetup");
            yield return null;
        }

        [Test]
        public void UnitTest()
        {
            Debug.Log("Test");
        }

        [UnityTest]
        public IEnumerator UnityTest()
        {
            Debug.Log("UnityTest before yield");
            yield return new EnterPlayMode(); 
            //Domain reload happening
            yield return new ExitPlayMode();
            Debug.Log("UnityTest after yield");
        }

        [TearDown]
        public new void TearDown()
        {
            Debug.Log("TearDown");
        }

        [UnityTearDown]
        public new IEnumerator UnityTearDown()
        {
            Debug.Log("UnityTearDown");
            yield return null;
        }

        [OneTimeTearDown]
        public void OneTimeTearDown()
        {
            Debug.Log("OneTimeTearDown");
        }
    }
```

## Additional resources

* [Setup and teardown](course/setup-teardown.md)
* [Execution order of test actions](reference-actions-outside-tests.md)