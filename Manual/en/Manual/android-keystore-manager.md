# Keystore Manager window reference

This page describes the Keystore Manager window interface.

To open the Keystore Manager window, open [Android Publishing Settings](class-PlayerSettingsAndroid.md#Publishing) and select **Keystore Manager**.

You can use this interface to:

* [Create a new keystore](android-keystore-create.md)
* [Add keys to a keystore](android-keystore-add-keys.md)
* [Load an existing keystore](android-keystore-load.md)

## Keystore

The properties in this section of the interface relate to the keystore that the Keystore Manager window currently manages.

| **Property** | **Description** |
| --- | --- |
| **Keystore dropdown** | Specifies which keystore to use in the Keystore Manager window. You can either create a new keystore, or load an existing one. For more information, refer to:  * [Create a new keystore](android-keystore-create.md) * [Load an existing keystore](android-keystore-load.md) |
| **Password** | The password for the keystore. If you create a new keystore, use this to create a password for the keystore. If you load an existing keystore, use this to enter the password for existing keystore. This property supports ASCII characters only. |
| **Confirm Password** | The confirmed password for the new keystore. Set the value of this property to the same as **Password**.  This property only appears if you create a new keystore. |

## Existing Keys

This section of the interface contains a read-only list of keys that the current keystore contains.

## New Key Values

The properties in this section of the interface describe a new key to add to the keystore. For more information, see [Add keys to a keystore](android-keystore-add-keys.md).

**Note**: The application doesn’t display the personal information in this section. Instead, the certificate that is part of your application includes it.

| **Property** | **Description** |
| --- | --- |
| **Alias** | An identifying name for the key. |
| **Password** | The password to use for the key. This property supports ASCII characters only. |
| **Confirm password** | The confirmed password for the key. Set the value of this property to the same as **Password**. |
| **Validity (years)** | The number of years that the key is valid for. This should exceed the amount of time you expect to manage the application for, so that you can use the same key to sign application updates. The default validity is 50 years. |
| **First and Last Name** | Your first and last name. |
| **Organizational Unit** | The division of your organization that you are part of. |
| **Organization** | The organization that manages your application. This is often your company name. |
| **City or Locality** | Your personal city or locality. |
| **State or Province** | Your personal state or province. |
| **Country Code** | Your personal country code. |

## Additional resources

* [Create a new keystore](android-keystore-create.md)
* [Add keys to a keystore](android-keystore-add-keys.md)
* [Load a keystore](android-keystore-load.md)