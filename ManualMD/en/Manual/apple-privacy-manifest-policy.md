# Apple’s privacy manifest policy requirements

Apple requires new and updated applications targeted for iOS, iPadOS, macOS, tvOS, and visionOS platforms on the App Store to include [privacy manifest files](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files).

The privacy manifest file (`PrivacyInfo.xcprivacy`) lists the [types of data](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_data_use_in_privacy_manifests) your Unity applications, or any third-party SDKs, packages, and **plug-ins** collect, and the reasons for using certain [required reason API](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api) categories.

**Important:**

* If the use of the required reason APIs by you or third-party SDKs isn’t declared in the privacy manifest, your application might be rejected by the App Store.
* The applications targeted for macOS platforms don’t require to declare the use of the required reason APIs. For macOS applications, the privacy manifest requirement only applies if you plan to publish your application on the App Store. For more information, visit Apple’s documentation on [Describing use of required reason API](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api).

## Prepare your Unity applications for Apple’s privacy manifest policy

You must review your native code, C# code, and data collection and tracking practices to understand if Apple’s privacy manifest policy applies to you. Follow these guidelines to decide if you need to include a privacy manifest file in your product:

* If your application includes multiple third-party SDKs, packages, and plug-ins, then these third-party components (if applicable) must provision their own privacy manifest files separately.

  **Note:** It’s your responsibility however, to make sure that the owners of these third-party components include privacy manifest files. Unity isn’t responsible for any third-party privacy manifest, and their data collection and tracking practices.
* If your application includes Unity framework implementations of the [C# .Net framework APIs](#CSharpDotNetAPIs) that call certain APIs listed in the Apple’s [required reason API](https://developer.apple.com/documentation/bundleresources/app-privacy-configuration/nsprivacyaccessedapitypes/nsprivacyaccessedapitype#possibleValues) categories, then you must assess your product for the API usage. For assessing what constitutes as part of data collection and tracking practices, refer to Apple’s documentation on [privacy manifest files](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files).

  **Note:** It’s your responsibility to assess your use of each of these APIs and declare the applicable reasons for using them.

Depending on whether you’re using [Unity to develop an application](#UnityAppManifest) or providing your [SDKs, packages, and plug-ins](#ThirdPartyManifest) to use with Unity applications, the requirement for providing a privacy manifest file might differ.

**Note:** These guidelines are provided for your convenience. It’s important that you review Apple’s documentation on [privacy manifest files](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files) before creating a privacy manifest for your project.

## Privacy manifest for Unity applications

If you’re developing an application using Unity, consider the following steps:

1. Assess if your native application code uses any of the following APIs:
   * APIs listed under the [required reason API](https://developer.apple.com/documentation/bundleresources/app-privacy-configuration/nsprivacyaccessedapitypes/nsprivacyaccessedapitype?language=objc#possibleValues) category.
   * The [C# .Net framework APIs](#CSharpDotNetAPIs) in Unity framework.
2. If you meet one or both of the conditions from step 1, [create a privacy manifest file](https://developer.apple.com/documentation/bundleresources/privacy-manifest-files#Create-a-privacy-manifest) in Xcode and save it in the **Assets/Plugins** folder of your project.
3. In the privacy manifest file, declare the approved reasons for using the required reason APIs or C# .Net framework APIs, as applicable.

   **Important**: If you don’t declare the reasons for the use of APIs, your application might be rejected by the App Store. The applications targeted for macOS platforms don’t require to declare the use of the required reason APIs.
4. Verify if your native application code collects any type of data [categorized by Apple](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_data_use_in_privacy_manifests#4250555) and declare those data types in the privacy manifest file as applicable.
   Any third-party SDKs, packages, and plug-ins used in your application must include their own separate manifest files to declare data collection and the use of any [required reason APIs](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api) with approved reasons.

**Notes:**

* It’s your responsibility to check the accuracy of the privacy manifest within the Unity framework and if any third-party components included in your Unity project require any declarations in you privacy manifest. It’s recommended that you search these third-party components for any references to a privacy manifest declaration.
* If you’re developing an application using Unity as a library, check if your native application code collects any of the following information outside of the Unity project:
  + [Data collection](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_data_use_in_privacy_manifests)
  + [Required reason APIs](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api)
    If it does, then you must declare their usage in the your privacy manifest file.

## Privacy manifest for third-party SDKs, packages, and plug-ins

If you’re a plug-in, package, or SDK owner, and you provide a service or package to be included within a Unity project as a .framework, follow these steps:

1. Assess if your native code or native library includes any of the following categories of information:
   * Any type of data [categorized by Apple](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_data_use_in_privacy_manifests#4250555)
   * Any APIs listed under [required reason API](https://developer.apple.com/documentation/bundleresources/app-privacy-configuration/nsprivacyaccessedapitypes/nsprivacyaccessedapitype#possibleValues) categories
2. Declare the types of data being collected and the approved reasons for using the required reason APIs in the privacy manifest file of your framework (if applicable).

If your product isn’t used as a .framework within a Unity project, then consider delivering it as a .framework. If you’re unable to do so, then follow these steps:

1. Assess if your native code or native library uses the following APIs:
   * Any APIs listed under [required reason API](https://developer.apple.com/documentation/bundleresources/app-privacy-configuration/nsprivacyaccessedapitypes/nsprivacyaccessedapitype#possibleValues) categories.
   * Any [C# .Net framework APIs](#CSharpDotNetAPIs) in Unity framework.
2. If one or both of the API categories are used, then [create a privacy manifest](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files#4284009) file in Xcode and save it in your source code. Make sure to select the iOS platform in the plug-in **Inspector** window.
3. In the privacy manifest file, declare the use of required reason APIs or C# .Net framework APIs, as applicable with approved reasons.
4. Verify if your native code or native library collects any type of data [categorized by Apple](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_data_use_in_privacy_manifests#4250555) and declare those data types in the privacy manifest file as applicable.

## C# .Net framework APIs in Unity

The following table provides a list of C# .Net APIs that call the required reason API category named [File timestamp API](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api#4278393). If your application, SDK, package, or plug-in code calls any of the APIs from this list, declare the reasons for their use in your privacy manifest file following the guidelines specified in Apple’s documentation on [required reason API](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api).

**Note**: The following list is verified only for Unity versions 2021.3 LTS and later.

| **File APIs** | **Directory APIs** |
| --- | --- |
| [fileInfo.CreationTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.CreationTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [fileInfo.LastAccessTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.LastAccessTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [fileInfo.LastWriteTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.LastWriteTime;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [fileInfo.CreationTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.CreationTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [fileInfo.LastAccessTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.LastAccessTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [fileInfo.LastWriteTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=net-8.0#properties) | [directoryInfo.LastWriteTimeUtc;](https://learn.microsoft.com/en-us/dotnet/api/system.io.directoryinfo?view=net-8.0#properties) |
| [File.GetCreationTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getcreationtime?view=net-8.0) | [Directory.GetCreationTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getcreationtime?view=net-8.0) |
| [File.GetLastAccessTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getlastaccesstime?view=net-8.0) | [Directory.GetLastAccessTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getlastaccesstime?view=net-8.0) |
| [File.GetLastWriteTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getlastwritetime?view=net-8.0) | [Directory.GetLastWriteTime(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getlastwritetime?view=net-8.0) |
| [File.GetCreationTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getcreationtimeutc?view=net-8.0) | [Directory.GetCreationTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getcreationtimeutc?view=net-7.0) |
| [File.GetLastAccessTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getlastaccesstimeutc?view=net-8.0) | [Directory.GetLastAccessTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getlastaccesstimeutc?view=net-7.0) |
| [File.GetLastWriteTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.getlastwritetimeutc?view=net-8.0) | [Directory.GetLastWriteTimeUtc(filePath);](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getlastwritetimeutc?view=net-7.0) |

## Unity Engine privacy manifest

Unity Engine calls the required reason APIs from the following categories and declares these categories with the listed reasons for their use. For more information on the categories, refer to [Apple’s documentation](https://developer.apple.com/documentation/bundleresources/app-privacy-configuration/nsprivacyaccessedapitypes/nsprivacyaccessedapitype#possibleValues).

| **Category** | **Reason** | **Description** |
| --- | --- | --- |
| File timestamp APIs | 0A2A.1 | Used by C# timestamp APIs |
| File timestamp APIs | C617.1 | Used to check files |
| User defaults APIs | CA92.1 | Used for PlayerPrefs API |
| System boot time APIs | 35F9.1 | Used to calculate time between various events |
| Disk space APIs | E174.1 | Used to check disk space for AssetBundles |

**Note**: Starting with Unity Editor versions 2021.3.35f1, 2022.3.18f1, and 2023.2.7f1, Unity Engine automatically includes these reasons in the privacy manifest file. If you’re using an older Unity Editor version, you must always declare these reasons in the privacy manifest file of your project.