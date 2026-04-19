# Scripting API for packages

You can use the Package Manager scripting API to interact with the Package Manager using C# **scripts**. For example, you might want to install a specific package or version depending on the platform of the target machine.

Using the scripting API for package management relies on the [PackageManager.Client](../ScriptReference/PackageManager.Client.md) class. Use this class to find packages, browse the list of packages, and install and uninstall packages through scripting.

Another important class is [PackageManager.PackageInfo](../ScriptReference/PackageManager.PackageInfo.md), which contains the state of a package, including metadata obtained from the **package manifest** and the registry. For example, you can get a [list of versions](../ScriptReference/PackageManager.VersionsInfo.md) available for the package, or the [list of any errors](../ScriptReference/PackageManager.PackageInfo-errors.md) that might occur while locating or installing the package.

## Adding a package to the project

This example demonstrates how to use the [Client](../ScriptReference/PackageManager.Client.md) class to install or add a package to the project.

You can use [Client.Add](../ScriptReference/PackageManager.Client.Add.md) to add a package. When you call the `Client.Add` method, you can specify just the package name, or the name with a specific version. For example, using `Client.Add("com.unity.textmeshpro")` installs (or updates to) the latest version of the TextMesh Pro package. Using `Client.Add("com.unity.textmeshpro@1.3.0")` installs version 1.3.0 of the TextMesh Pro package.

The `Client.Add` method returns an [AddRequest](../ScriptReference/PackageManager.Requests.AddRequest.md) instance, which you can use to get the status, any errors, or a [Request](../ScriptReference/PackageManager.Requests.Request_1.Result.md) response that contains the [PackageInfo](../ScriptReference/PackageManager.PackageInfo.md) information for the newly added package.

```
using System;
using UnityEditor;
using UnityEditor.PackageManager.Requests;
using UnityEditor.PackageManager;
using UnityEngine;

namespace Unity.Editor.Example {
   static class AddPackageExample
   {
       static AddRequest Request;

       [MenuItem("Window/Add Package Example")]
       static void Add()
       {
           // Add a package to the project
           Request = Client.Add("com.unity.textmeshpro");
           EditorApplication.update += Progress;
       }

       static void Progress()
       {
           if (Request.IsCompleted)
           {
               if (Request.Status == StatusCode.Success)
                   Debug.Log("Installed: " + Request.Result.packageId);
               else if (Request.Status >= StatusCode.Failure)
                   Debug.Log(Request.Error.message);

               EditorApplication.update -= Progress;
           }
       }
   }
}
```

## Browsing the list of packages in a project

This example demonstrates how to use the [Client](../ScriptReference/PackageManager.Client.md) class to iterate over the packages in the project.

The [Client.List](../ScriptReference/PackageManager.Client.List.md) method returns a [ListRequest](../ScriptReference/PackageManager.Requests.ListRequest.md) instance. Use the `ListRequest` instance to get the status of the List operation, any errors, or a [Request](../ScriptReference/PackageManager.Requests.Request_1.Result.md) response that contains the [PackageCollection](../ScriptReference/PackageManager.PackageCollection.md) which you can iterate.

```
using System;
using UnityEditor;
using UnityEditor.PackageManager.Requests;
using UnityEditor.PackageManager;
using UnityEngine;

namespace Unity.Editor.Example {
   static class ListPackageExample
   {
       static ListRequest Request;

       [MenuItem("Window/List Package Example")]
       static void List()
       {
           Request = Client.List();    // List packages installed for the project
           EditorApplication.update += Progress;
       }

       static void Progress()
       {
           if (Request.IsCompleted)
           {
               if (Request.Status == StatusCode.Success)
                   foreach (var package in Request.Result)
                       Debug.Log("Package name: " + package.name);
               else if (Request.Status >= StatusCode.Failure)
                   Debug.Log(Request.Error.message);

               EditorApplication.update -= Progress;
           }
       }
   }
}
```

## Embedding a package in the project

This example demonstrates how to use the [Client](../ScriptReference/PackageManager.Client.md) class to [embed](upm-embed.md) one of the packages already installed in your project. The main method is the [Client.Embed](../ScriptReference/PackageManager.Client.Embed.md) method, which makes a copy of the package and stores it under the `Packages` folder of your project.

The [Client.Embed](../ScriptReference/PackageManager.Client.Embed.md) method returns an [EmbedRequest](../ScriptReference/PackageManager.Requests.EmbedRequest.md) instance. Use the `EmbedRequest` instance to get the status of the Embed operation, any errors, or a [Request](../ScriptReference/PackageManager.Requests.Request_1.Result.md) response that contains the [PackageInfo](../ScriptReference/PackageManager.PackageInfo.md) information for the newly **embedded package**.

This example also uses the [Client.List](../ScriptReference/PackageManager.Client.List.md) method to access the collection of packages currently installed in your project. The method in the example picks out the first one that’s neither embedded nor built-in.

The [Client.List](../ScriptReference/PackageManager.Client.List.md) method returns a [ListRequest](../ScriptReference/PackageManager.Requests.ListRequest.md) instance. Use the `ListRequest` instance to get the status of the List operation, any errors, or a [Request](../ScriptReference/PackageManager.Requests.Request_1.Result.md) response that contains the [PackageCollection](../ScriptReference/PackageManager.PackageCollection.md) which you can iterate.

```
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEditor;
using UnityEditor.PackageManager.Requests;
using UnityEditor.PackageManager;
using UnityEngine;

namespace Unity.Editor.Example
{
    static class EmbedPackageExample
    {
        static String targetPackage;
        static EmbedRequest Request;
        static ListRequest LRequest;

        [MenuItem("Window/Embed Package Example")]
        static void GetPackageName()
        {
            // First get the name of an installed package
            LRequest = Client.List();
            EditorApplication.update += LProgress;
        }

        static void LProgress()
        {
            if (LRequest.IsCompleted)
            {
                if (LRequest.Status == StatusCode.Success)
                {
                    foreach (var package in LRequest.Result)
                    {
                        // Only retrieve packages that are currently installed in the
                        // project (and are neither Built-In nor already Embedded)
                        if (package.isDirectDependency && package.source
                            != PackageSource.BuiltIn && package.source
                            != PackageSource.Embedded)
                        {
                            targetPackage = package.name;
                            break;
                        }
                    }

                }
                else
                    Debug.Log(LRequest.Error.message);

                EditorApplication.update -= LProgress;

                Embed(targetPackage);

            }
        }

        static void Embed(string inTarget)
        {
            // Embed a package in the project
            Debug.Log("Embed('" + inTarget + "') called");
            Request = Client.Embed(inTarget);
            EditorApplication.update += Progress;

        }

        static void Progress()
        {
            if (Request.IsCompleted)
            {
                if (Request.Status == StatusCode.Success)
                    Debug.Log("Embedded: " + Request.Result.packageId);
                else if (Request.Status >= StatusCode.Failure)
                    Debug.Log(Request.Error.message);

                EditorApplication.update -= Progress;
            }
        }
    }
}
```

## Package Manager events

Use the [Events](../ScriptReference/PackageManager.Events.md) class to register an event handler with the Package Manager. The [Events](../ScriptReference/PackageManager.Events.md) class contains two events you can subscribe to, which the Package Manager raises at these points:

* Immediately before the Package Manager changes the list of dependencies ([registeringPackages](../ScriptReference/PackageManager.Events-registeringPackages.md))
* After the Package Manager imports and compiles the changed list of package dependencies ([registeredPackages](../ScriptReference/PackageManager.Events-registeredPackages.md))

The following examples demonstrate how to use both of these events.

### Example of using the registeringPackages event

```
using UnityEditor.PackageManager;
using UnityEngine;

namespace Unity.Editor.Example
{
    public class EventSubscribingExample_RegisteringPackages
    {
        public EventSubscribingExample_RegisteringPackages()
        {
            // Subscribe to the event using the addition assignment operator (+=).
            // This executes the code in the handler whenever the event is fired.
            Events.registeringPackages += RegisteringPackagesEventHandler;
        }

        // The method is expected to receive a PackageRegistrationEventArgs event argument.
        void RegisteringPackagesEventHandler(PackageRegistrationEventArgs packageRegistrationEventArgs)
        {
            Debug.Log("The list of registered packages is about to change!");

           foreach (var addedPackage in packageRegistrationEventArgs.added)
            {
                Debug.Log($"Adding {addedPackage.displayName}");
            }

            foreach (var removedPackage in packageRegistrationEventArgs.removed)
            {
                Debug.Log($"Removing {removedPackage.displayName}");
            }

            // The changedFrom and changedTo collections contain the packages that are about to be updated.
            // Both collections are guaranteed to be the same size with indices matching the same package name.
            for (int i = 0; i <= packageRegistrationEventArgs.changedFrom.Count; i++)
            {
                var oldPackage = packageRegistrationEventArgs.changedFrom[i];
                var newPackage = packageRegistrationEventArgs.changedTo[i];

                Debug.Log($"Changing ${oldPackage.displayName} version from ${oldPackage.version} to ${newPackage.version}");
            }
        }
    }
}
```

### Example of using the registeredPackages event

```
using UnityEditor;
using UnityEditor.PackageManager;
using UnityEngine;

namespace Unity.Editor.Example
{
    public class EventSubscribingExample_RegisteredPackages
    {
        // You must use '[InitializeOnLoadMethod]' or '[InitializeOnLoad]' to subscribe to this event.
        [InitializeOnLoadMethod]
        static void SubscribeToEvent()
        {
            // This causes the method to be invoked after the Editor registers the new list of packages.
            Events.registeredPackages += RegisteredPackagesEventHandler;
        }

        static void RegisteredPackagesEventHandler(PackageRegistrationEventArgs packageRegistrationEventArgs)
        {
            // Code executed here can safely assume that the Editor has finished compiling the new list of packages
            Debug.Log("The list of registered packages has changed!");
        }
    }
}
```