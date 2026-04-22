# Run tests from the command line

You can run a test project from the Unity Editor command line. The following example demonstrates this on Windows:

```
Unity.exe -runTests -batchmode -projectPath PATH_TO_YOUR_PROJECT -testResults C:\temp\results.xml -testPlatform PS4
```

**Note**: Use the `-batchmode` option when running tests on the command line to remove the need for manual user inputs. For more information, refer to [Unity Editor command-line arguments](../EditorCommandLineArguments.md).

For the full list of Editor command line arguments provided by Unity Test Framework, refer to [Command-line reference](reference-command-line.md).

## Additional resources

* [Command-line reference](reference-command-line.md)
* [Unity Editor command-line arguments](../EditorCommandLineArguments.md)