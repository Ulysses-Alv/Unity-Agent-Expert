# Deliver applications to the Mac App Store

This page has information about Mac App Store delivery requirements.

## Delivery requirements

If you want to distribute your application to the Mac App Store, then the upload process includes a similar content validation to [notarization](macos-building-notarization.md), which means that you don’t have to notarize your application. However, when you [build your macOS application](macos-building.md), the Mac App Store has additional requirements.

### Validation

Before you build your application, enable the **Mac App Store Validation** property. To do this:

1. Open Project Settings (menu: **Edit** > **Project Settings**).
2. Go to **Player** > **Other Settings** > **Identification**.
3. Enable **Mac App Store Validation**.

Enabling the **Mac App Store Validation** option means that your application only runs when it holds a valid receipt from the Mac App Store. It prevents people from running the application on a different computer than the one they purchased.

### Signing identity

To distribute your application via the Mac App Store, Apple requires you to use one of the following signing identities when you [code sign](macos-building.md#code-signing) your application:

| **Signing Identity** | **Format** |
| --- | --- |
| Mac App Distribution | `3rd Party Mac Developer Application: TTT` |
| Apple Distribution | `Apple Distribution: TTT` |

**Note**: Replace `TTT` with your team identity.

### Submit a macOS application

For information on how to submit your application to the Mac App Store, refer to the
Submit and promote section in [Submitting apps to the Mac App Store](https://developer.apple.com/macos/submit/).

## Additional resources

* [Build a macOS application](macos-building.md)