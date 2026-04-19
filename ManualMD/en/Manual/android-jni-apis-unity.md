# Java Native Interface APIs in Unity

Unity provides low-level and high-level [Java Native Interface](https://developer.android.com/training/articles/perf-jni) (JNI) APIs that allow you to interact with Java code from C# **scripts**.

## Low-level API

The low-level [AndroidJNI](../ScriptReference/AndroidJNI.md) class wraps JNI calls and provides static methods that directly map to JNI methods. The [AndroidJNIHelper](../ScriptReference/AndroidJNIHelper.md) API provides helper functionality that’s primarily used by the high-level API, but they can be useful in certain situations.

## High-level API

The high-level [AndroidJavaObject](../ScriptReference/AndroidJavaObject.md), [AndroidJavaClass](../ScriptReference/AndroidJavaClass.md), and [AndroidJavaProxy](../ScriptReference/AndroidJavaProxy.md) classes automate a lot of tasks required for JNI calls. They also use caching to make calls to Java faster. The combination of `AndroidJavaObject` and `AndroidJavaClass` is built on top of `AndroidJNI` and `AndroidJNIHelper`, but they also contain additional functionality such as static methods that you can use to access static members of Java classes.

Additionally, Unity provides the [AndroidApplication](../ScriptReference/Android.AndroidApplication.md) class to simplify access to instances of `currentActivity`, `currentContext`, and `currentConfiguration` for your application. This class also allows you to delegate code execution on the UI or main thread based on your application’s requirement.

Instances of `AndroidJavaObject` and `AndroidJavaClass` have a one-to-one mapping to an instance of [java.lang.Object](https://docs.oracle.com/javase/7/docs/api/java/lang/Object.html) and [java.lang.Class](https://docs.oracle.com/javase/7/docs/api/java/lang/Class.html) respectively. They provide three types of interactions with Java/Kotlin code:

* [Call](../ScriptReference/AndroidJavaObject.Call.md) a method.
* [Get](../ScriptReference/AndroidJavaObject.Get.md) the value of a field.
* [Set](../ScriptReference/AndroidJavaObject.Set.md) the value of a field.

Each interaction also has a static version:

* [CallStatic](../ScriptReference/AndroidJavaObject.CallStatic.md) to call a static method.
* [GetStatic](../ScriptReference/AndroidJavaObject.GetStatic.md) to get the value of a static field.
* [SetStatic](../ScriptReference/AndroidJavaObject.SetStatic.md) to set the value of a static field.

When you get the value of a field or call a method that returns a value, you use [generics](https://docs.microsoft.com/en-us/dotnet/csharp/fundamentals/types/generics) to specify the return type. When you set the value of a field, you also use generics to specify the type of the field that you’re setting. For methods that don’t return a value, there’s a regular, non-generic version of [Call](../ScriptReference/AndroidJavaObject.Call.md).

**Important**: You must access any [non-primitive type](http://developer.android.com/reference/java/lang/Class.html) as an `AndroidJavaObject`. The only exception is a string which you access directly, even though they don’t represent a primitive type in Java.

## Additional resources

* [Code examples: Call Java/Kotlin code from C# scripts](android-high-level-api-code-examples.md)
* [Best practices for calling Java/Kotlin code](android-call-java-kotlin-code-best-practices.md)