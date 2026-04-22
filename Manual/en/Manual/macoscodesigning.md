# Code sign your application

Code signing is the process of creating a code signature for an application. This signature guarantees the integrity of applications and safeguards from any tampering. Apple devices use an application’s code signature to detect changes made after the developer created the code signature. If an application doesn’t have a code signature, the device warns the user before they open it.

**Note**: You must code sign your application to notarize it with the [Xcode command-line](macosnotarizationxcode.md) or [Unity Build Automation](https://docs.unity.com/devops/en/manual/unity-build-automation).

## Create a signing identity

Unity adds a code signature to every macOS build it produces, known as a Signing Identity. To [notarize](macosnotarizationxcode.md) an application, Apple requires the code signature to include a cryptographic signature, such as a Developer ID certificate, that identifies the developer.

To create a new Developer ID certificate, use the following steps:

1. Go to [Create a New Certificate](https://developer.apple.com/account/resources/certificates/add).
2. Follow the prompts to create a private key and upload a **Certificate Signing Request**.
3. Download your Developer ID certificate. The file format for this certificate is `.cer`.
4. Click on the certificate to add it to your [Keychain](https://support.apple.com/en-gb/guide/keychain-access/kyca1083/11.0/mac/13.0). Within the Keychain, it appears with a name formatted like: **Developer ID Application : XXX (YYY)**.

## Locate the application identifier

To notarize your application, Apple needs to identify it using an application identifier. There are two ways to get an application identifier: in Unity, or in the application’s [information property list file](macos-building.md#information-property-list-file).

* In Unity:
  1. Go to **Edit** > **Project Settings** > **Player**.
  2. Expand the **Other Settings** section
  3. Navigate to **Identification**.
  4. The **Bundle Identifier** property is your application identifier.
* In the information property list file:
  1. In the Finder, go to your built application.
  2. Right-click on the application and select **Show Package Contents**.
  3. Go to **Contents** and
  4. Open **Info.plist**.
  5. The key that represents the application identifier is [**CFBundleIdentifier**](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-102070).

When you have your application identifier, you can register it with Apple. To do this, use the following steps:

1. Sign in to your [Apple Developer](https://developer.apple.com/) account.
2. Select **Certificates, IDs & Profiles**.
3. Select **Identifiers**
4. Enter your bundle ID.

## Create entitlements

Entitlements are permissions or restrictions your code signature includes that control the actions your application can take.

To set entitlements for your application, use the following steps:

1. In the same directory as your built application, create a file with the same name as your application with the `.entitlements` file extension. For example, if you name your application MyProject, create a file called `MyProject.entitlements`.
2. Open the file with a text editor and copy in the following content:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
       <key>com.apple.security.cs.disable-library-validation</key>
       <true/>
       <key>com.apple.security.cs.disable-executable-page-protection</key>
       <true/>
    </dict>
</plist>
```

These entitlements are the minimum entitlements a macOS application requires to have a [Hardened Runtime](https://developer.apple.com/documentation/security/hardened_runtime). If your application requires any more entitlements, add them to this list.

## Code sign the application

To code sign your application you need to use the command line. On your machine, open Terminal and navigate to the directory that the application is in.

To ensure that you have the necessary read permissions to process code signing, run the following command where `"application_name.app"` is the name of your application:

```
chmod -R a+xr "application_name.app"
```

To code sign your application, run the following command where:

* `"application_name.app"` is your built application.
* `"application_name.entitlements"` is the name of the [entitlements file](#entitlements).
* `"Developer ID Application : XXX (YYY)"` is your [signing identity](#signing-identity).

```
codesign 
    --deep 
    --force 
    --verify 
    --verbose 
    --timestamp 
    --options runtime 
    --entitlements "application_name.entitlements" 
    --sign "Developer ID Application : XXX (YYY)" "application_name.app"
```

This command works through the application bundle folder, signs all files, adds a secure timestamp, and embeds the entitlements you’ve set into the signature.

Using the `--deep` option might cause issues with your code signature. This is because:

* It applies the same code signing options and entitlements to all the code that it signs.
* It only signs code files that it finds. If there are code files in a place where the system expects to find data, using `--deep` doesn’t sign these code files.

For more information about the `--deep` option and how to resolve issues with it, refer to [Sign your code](https://developer.apple.com/forums/thread/128166).

## Additional resources

* [Notarize with Xcode and command-line tools](macosnotarizationxcode.md)
* [Build a macOS application](macos-building.md)