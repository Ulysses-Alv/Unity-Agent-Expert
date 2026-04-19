# Calling Java and Kotlin plug-in code from C# scripts

Unity allows you to call Java and Kotlin code directly from C# scripts and integrate Android-specific features in your Unity project. To achieve this, Unity provides C# APIs that communicate with the Android [Java Native Interface](https://developer.android.com/training/articles/perf-jni) (JNI) through C++. You can choose between low level and high level APIs to interact with the Java code using JNI.

Refer to the following sections for detailed information on how to call Java and Kotlin **plug-in** code from C# **scripts**.

| **Topic** | **Description** |
| --- | --- |
| [Java Native Interface APIs in Unity](android-jni-apis-unity.md) | Learn about the Java Native Interface APIs in Unity that you can use to call Java/Kotlin code from C# scripts. |
| [Code examples: Call Java/Kotlin code from C# scripts](android-high-level-api-code-examples.md) | Refer to the code examples for integrating Java/Kotlin functionality in your Unity project. |
| [Supported data types for Java/Kotlin and C# code](android-datatypes-java-kotlin-csharp.md) | Understand the data types that you can transfer between Java/Kotlin code and C# scripts. |
| [Best practices for calling Java/Kotlin code](android-call-java-kotlin-code-best-practices.md) | Learn about the best practices for calling Java/Kotlin code from C# scripts. |

## Additional resources

* [Example: Create and use Kotlin source plug-in](AndroidJavaSourcePlugins.md#CreateUseKotlinSourcePlugin)