# Notarize with Xcode and command-line tools

You must notarize your application to distribute it outside of the Mac App Store. The process verifies your application, ensuring it has a Developer ID code signature and doesn’t contain malicious content. You can notarize your application with Xcode, Xcode command-line tools, or [Unity Build Automation](https://docs.unity.com/devops/en/manual/build-automation/macos-notarization).

Unity can create an Xcode project that represents your Unity project during the build process. You can use this Xcode project to notarize your application. Follow the [macOS application build steps](macos-building.md#building) to create an Xcode project from your Unity project. For information on notarizing the Xcode project, refer to [Notarizing macOS Software Before Distribution](https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution).

**Note**: Notarization isn’t required to distribute your application through the Mac App Store. The Mac App Store’s upload process includes content validation similar to notarization. Refer to [Delivering to the Mac App Store](macos-distribution-mac-app-store.md) for more information.

## Use Xcode command-line tools

If you can’t notarize your application in the Xcode environment or want to control notarization manually, you can use Xcode’s command-line tools.

To do this, you need the following:

* An [Apple ID](https://appleid.apple.com/).
* An Apple device that runs macOS 10.15 Catalina or later.
* [Xcode 11 or later](https://developer.apple.com/xcode/). This installs the required command-line tools.
* An [Apple Developer Program membership](https://developer.apple.com/programs/). If you don’t have an Apple Developer membership, sign up at [Apple Developer](https://developer.apple.com/).

### Compress the application

Apple requires that you compress your application before you upload it for notarization. To do this, use the following steps:

1. Open Terminal and navigate to the directory the application is in.
2. Run the following command where:
   * `"application_name.app"` is your built application.
   * `"application_name.zip"` is the name of the compressed file to generate.

```
ditto -c -k 
    --sequesterRsrc 
    --keepParent "application_name.app" "application_name.zip"
```

This compresses your application and outputs the compressed file to the same directory as your application.

### Generate an application password

To notarize an application, Apple requires a unique password in a particular format. For information on how to generate an application password, refer to [How to generate an app-specific password](https://support.apple.com/en-gb/HT204397). The password you generate uses the following format: `xxxx-xxxx-xxxx-xxxx`.

## Notarize your application

If you are using Xcode 13 or later, use Apple’s notarytool to upload and notarize your application. For more information, refer to [Customizing the Notarization Workflow](https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution/customizing_the_notarization_workflow?language=objc) and [Enable notarization on an older version of macOS](https://developer.apple.com/documentation/technotes/tn3147-migrating-to-the-latest-notarization-tool#Enable-notarization-on-an-older-version-of-macOS).

**Note**: As of November 1st 2023, Apple have [deprecated the use of altool](https://developer.apple.com/news/?id=y5mjxqmn) for notarization. You must update your workflow to use notarytool to notarize your macOS applications. For more information, refer to [Migrating to the latest notarization tool](https://developer.apple.com/documentation/technotes/tn3147-migrating-to-the-latest-notarization-tool).

## Staple the application

After notarizing your application, any device that runs it can verify that it has a code signature and has no malicious content. However, the device can only perform this verification online. To let a device verify your application without an internet connection, you must staple the application. For information about stapling, refer to [Staple the Ticket to Your Distribution](https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution/customizing_the_notarization_workflow#3087720).

To staple your application,

1. Open Terminal and navigate to the directory the application is in.
2. Run the following command where `"ApplicationName.app"` is the name of your application:

```
xcrun stapler staple "ApplicationName.app"
```

## Additional resources

* [Code sign your application](macoscodesigning.md)