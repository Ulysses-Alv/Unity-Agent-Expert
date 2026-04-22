# Command-line reference

Unity Test Framework adds the following additional arguments to the regular [Unity Editor command-line arguments](../EditorCommandLineArguments.md):

| **Command** | **Description** |
| --- | --- |
| `-assemblyNames` | A semicolon-separated list of test assemblies to include in the run. A semi-colon separated list should be formatted as a string enclosed in quotation marks, e.g. `-assemblyNames "firstAssembly;secondAssembly"`. |
| `-forgetProjectPath` | Prevents Unity saving your current project into the Unity launcher/hub history. |
| `-orderedTestListFile` | Path to a .txt file which contains a list of full test names you want to run in the specified order. There are no restrictions on the file name as long as the format is text. Separate tests by new lines and specify parameters if the tests have them. The following is an example of the content of such a file:  `Unity.Framework.Tests.OrderedTests.NoParameters` `Unity.Framework.Tests.OrderedTests.ParametrizedTestA(3,2)` `Unity.Framework.Tests.OrderedTests.ParametrizedTestB("Assets/file.fbx")` `Unity.Framework.Tests.OrderedTests.ParametrizedTestD(1.0f)` `Unity.Framework.Tests.OrderedTests.ParametrizedTestE(null)` `Unity.Framework.Tests.OrderedTests.ParametrizedTestF(False, 1)` `Unity.Framework.Tests.OrderedTests.ParametrizedTestG(float.NaN)` `Unity.Framework.Tests.OrderedTests.ParametrizedTestH(SomeEnum)` |
| `-playerHeartbeatTimeout` | Defines the length of time, in seconds, for which the Editor waits for heartbeats after starting a test run on a Player. This defaults to 10 minutes. |
| `-randomOrderSeed` | A non-zero integer representing the seed value to use for randomizing tests in the project, independently from the fixture. `# normal order` `Test_1` `Test_2` `Test_3` `Test_4` `# randomized with seed x` `Test_3`  `Test_1` `Test_4` `Test_2` `# randomized with same seed x and a new test` `Test_3` `Test_5` `Test_1` `Test_4` `Test_2` |
| `-repeat` | An integer that sets the repeat count. Successful tests are repeated up to this number of times, or until they fail, whichever happens first. |
| `-retry` | An integer that sets the retry count. Failing tests are retried up to this number of times, or until they succeed, whichever happens first. |
| `-runSynchronously` | Supply this argument to run tests synchronously, guaranteeing that all tests run in one Editor update call. This is only supported for Edit mode tests. Tests that take multiple frames, such as `[UnityTest]` tests, or tests with `[UnitySetUp]` or `[UnityTearDown]` scaffolding, are filtered out. |
| `-runTests` | Runs tests in the Project. |
| `-testCategory` | A semicolon-separated list of test categories to include in the run, or a regular expression pattern to match category names. A semi-colon separated list should be formatted as a string enclosed in quotation marks, for example `-testCategory "firstCategory;secondCategory"`. If both `-testFilter` and `-testCategory` are provided, only tests that match both are run. This argument supports negation using `!`. Using `!MyCategory` excludes tests with the `MyCategory` category from the run. |
| `-testfilter` | A semicolon-separated list of test names to run, or a regular expression pattern to match tests by their full name. A semi-colon separated list should be formatted as a string enclosed in quotation marks, for example `testFilter "Low;Medium"`. This argument supports negation using `!`. The test filter `!MyNamespace.Something.MyTest` runs all tests except the one called `MyNamespace.Something.MyTest`. Run a specific variation of a parameterized test as follows: `"ClassName\.MethodName\(Param1,Param2\)"` |
| `-testPlatform` | The platform to run tests on. Accepted values:  * `EditMode`: Edit mode tests. Equivalent to running tests from the EditMode tab of the **Test Runner** window. * `PlayMode`: Play mode tests that run in the Editor. Equivalent to running tests from the PlayMode tab of the **Test Runner** window. * Any value from the [BuildTarget](https://docs.unity3d.com/ScriptReference/BuildTarget.html) enum. Play mode tests that run on a Player built for the specified platform. Equivalent to using the Run All option in the Player tab of the Test Runner window.  **Note**: If no value is specified for this argument, tests run in Edit mode. |
| `-testResults` | Specifies a path to which Unity saves the test result file. By default, Unity saves the file in the project’s root folder. Test results follow the XML format as defined by NUnit, for more information refer to [Test Result XML Format](https://docs.nunit.org/articles/nunit/technical-notes/usage/Test-Result-XML-Format.html). There is currently no common definition for exit codes reported by individual Unity components under test. The best way to understand the source of a problem is the content of error messages and stack traces. |
| `-testSettingsFile` | Path to a [TestSettings.json](#test-settings) file. |

**Note**: The Editor’s regular `-quit` command-line argument is not supported while tests are running.

## Test settings reference

You can define options for a test run in a `TestSettings.json` file. You can specify the location of this file using the `-testSettingsFile` command line argument.

The following example illustrates the content of a typical `TestSettings.json` file:

```
{
  "scriptingBackend":"WinRTDotNET",
  "Architecture":null,
  "apiProfile":0,
  "featureFlags": { "requiresSplashScreen": true }
}
```

The following table lists the valid options and values for a `TestSettings.json` file:

| **Option** | **Description** |
| --- | --- |
| `apiProfile` | The [.NET API compatability level](https://docs.unity3d.com/ScriptReference/ApiCompatibilityLevel.html). Set to one of the following values:  * 1 - .Net 2.0 * 2 - .Net 2.0 Subset * 3 - .Net 4.6 * 5 - .Net micro profile (used by Mono scripting backend if **Stripping Level** is set to **Use micro mscorlib**) * 6 - .Net Standard 2.0 |
| `appleEnableAutomaticSigning` | Enables automatic signing of Apple devices. For more information, refer to [PlayerSettings.iOS.appleEnableAutomaticSigning](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-appleEnableAutomaticSigning.html). |
| `appleDeveloperTeamID` | Sets the team ID for the apple developer account. For more information, refer to [PlayerSettings.iOS.appleDeveloperTeamID](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-appleDeveloperTeamID.html). |
| `architecture` | Target [Android architecture](https://docs.unity3d.com/ScriptReference/AndroidArchitecture.html). Set to one of the following values:  * None = 0 * ARMv7 = 1 * ARM64 = 2 * All = 4294967295 |
| `iOSManualProvisioningProfileType` | For more information, refer to [PlayerSettings.iOS.iOSManualProvisioningProfileType](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-iOSManualProvisioningProfileType.html). Set to one of the following values:  * 0 - Automatic * 1 - Development * 2 - Distribution [iOSManualProvisioningProfileID](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-iOSManualProvisioningProfileID.html) |
| `iOSTargetSDK` | Target SDK for iOS. Set to one of the following values, which should be given as a string literal enclosed in quotes:  * DeviceSDK * SimulatorSDK |
| `tvOSManualProvisioningProfileType` | For more information, refer to [PlayerSettings.iOS.tvOSManualProvisioningProfileType](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-tvOSManualProvisioningProfileType.html). Set to one of the following values:  * 0 - Automatic * 1 - Development * 2 - Distribution [tvOSManualProvisioningProfileID](https://docs.unity3d.com/ScriptReference/PlayerSettings.iOS-tvOSManualProvisioningProfileID.html) |
| `tvOSTargetSDK` | Target SDK for tvOS. Set to one of the following values, which should be given as a string literal enclosed in quotes:  * DeviceSDK * SimulatorSDK |
| `scriptingBackend` | Set to one of the following values, which should be given as a string literal enclosed in quotes:  * Mono2x * IL2CPP |
| `playerGraphicsAPI` | Set graphics API that will be used during test execution in the player. Value can be any [GraphicsDeviceType](https://docs.unity3d.com/ScriptReference/Rendering.GraphicsDeviceType.html) as a string literal enclosed in quotes. Value will only be set if it is supported on the target platform. |
| `webGLClientBrowserType` | A browser to be used when running test using WebGL platform. Accepted browser types:  * Safari * Firefox * Chrome * Chromium |
| `webGLClientBrowserPath` | An absolute path to the browser’s location on your device. If not defined, path from `UNITY_AUTOMATION_DEFAULT_BROWSER_PATH` enviromental variable will be used. |
| `androidBuildAppBundle` | A boolean setting that allows to build an Android App Bundle (AAB) instead of **APK** for tests. |
| `featureFlags` | Map of strings and boolean values which can switch Unity Test Framework features on or off. The currently supported features are:  * `fileCleanUpCheck` - Throws an error message (instead of warning) if tests generate files which are not cleaned up. False (off) by default. * `requiresSplashScreen` - By default Unity Test Framework disables the Made with Unity splash screen to speed up building the Player and running tests. Set this flag to `true` to override the default and always include a Made with Unity splash screen in the Player build. |

## Additional resources

* [Run tests from the command line](run-tests-from-command-line.md)
* [Unity Editor command-line arguments](../EditorCommandLineArguments.md)