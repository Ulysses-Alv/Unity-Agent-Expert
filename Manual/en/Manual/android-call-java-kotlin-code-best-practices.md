# Best practices for calling Java/Kotlin code

Learn about the best practices for calling Java and Kotlin **plug-in** code from C# **scripts** to improve performance of your application.

## Minimize JNI calls

Using the Java Native Interface (JNI), through either the high-level or low-level C# API is resource intensive and can be slow. To improve performance, and also code clarity, it’s best practice to keep the number of JNI calls low.

To avoid unnecessary JNI calls, the high-level C# API caches the [ID](../ScriptReference/AndroidJNI.GetMethodID.md) of each Java method that you call. This means that subsequent calls to the same method aren’t as resource intensive as the first call. The calls don’t need to be during the same frame or even from the same `AndroidJavaObject`/`AndroidJavaClass` instance. If you use the low-level API and want this performance benefit, you must manually cache method ID. Otherwise, it’s best practice to use the high-level API.

**Note**: Unity maintains the cache until the application [closes](https://support.google.com/android/answer/9079646?hl=en-GB). This includes while the application is in the background.

## Manage garbage collection

It’s best practice to wrap any instance of `AndroidJavaObject` or `AndroidJavaClass` with a `using` statement to ensure Unity destroys them as soon as possible. If you don’t use `using`, Unity’s [garbage collector](performance-garbage-collector.md) will still release all created instances, but you lose control over when this happens.

The following code example demonstrates how to use `using` statements to access the system language in an optimal way:

```
using UnityEngine;

public class LocaleExample : MonoBehaviour
{
    void Start()
    {
        using (AndroidJavaClass cls = new AndroidJavaClass("java.util.Locale"))
        using (AndroidJavaObject locale = cls.CallStatic<AndroidJavaObject>("getDefault"))
        {
            if (locale != null)
            {
                Debug.Log("current lang = " + locale.Call<string>("getDisplayLanguage"));
            }
        }
    }
}
```

## Additional resources

* [Code examples: Call Java/Kotlin code from C# scripts](android-high-level-api-code-examples.md)
* [Supported data types for Java/Kotlin and C# code](android-datatypes-java-kotlin-csharp.md)
* [Example: Create and use Kotlin source plug-in](AndroidJavaSourcePlugins.md#CreateUseKotlinSourcePlugin)