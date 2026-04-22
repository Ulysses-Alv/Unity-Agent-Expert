# Set up Play Asset Delivery

[Play Asset Delivery](play-asset-delivery.md) is the asset splitting solution for [Android App Bundles](https://developer.android.com/guide/app-bundle) (AAB). To use Play Asset Delivery, set up your project to:

1. Use the AAB publishing format. For information on how to do this, refer to [Publishing format](android-BuildProcess.md#publishing-format).
2. Split the application binary. For information on how to do this, refer to [Splitting the application binary](android-optimize-distribution-size.md#splitting-the-application-binary). If **Split Application Binary** is grayed out, it means your current Unity Editor version doesn’t support Play Asset Delivery. To resolve this, update the Unity Editor.

When you build your application, Unity creates an AAB that includes your application split into a [base module](https://developer.android.com/guide/app-bundle/configure-base) and asset packs. For more information, refer to [Asset packs in Unity](android-asset-packs-in-unity.md).

**Important**: Unity uses `PLAY_ASSET_PACKS` [Gradle template variable](android-gradle-template-variables.md) to specify which asset packs to include in the Android App Bundle. When you build a Unity project with Play Asset Delivery support, Unity automatically generates values for this variable. If you’re upgrading from a Unity version without Play Asset Delivery support and using a custom **Gradle** template, it’s recommended to recreate the template file and reapply the modifications. Otherwise, your project build will not include the asset packs.

## Additional resources

* [Asset packs in Unity](android-asset-packs-in-unity.md)
* [Create a custom asset pack](android-asset-packs-create-custom.md)
* [Manage asset packs at runtime](android-asset-packs-manage.md)