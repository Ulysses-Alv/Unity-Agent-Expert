# Prepare your application for in-app purchases

In-app purchases (IAP) allow you to offer additional downloadable content in your application, such as new levels or character cosmetics. You must integrate with the StoreKit API within your application using a [native code plug-in](./plug-ins.md) before you can set up in-app purchases. For more information, refer to [StoreKit](https://developer.apple.com/documentation/storekit) (Apple).

**Note**: The [Unity IAP package](https://docs.unity3d.com/Packages/com.unity.purchasing@latest) can be used to implement in-app purchases for iOS and other platforms you might want to develop for.

## Organize your assets

The [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest/) package provides a ready-made system to manage and organize [AssetBundles](AssetBundlesIntro.md) in your project. It’s recommended to use the Addressables package rather than manage AssetBundles yourself.

## Download your assets

If you are managing AssetBundles yourself, it’s recommended to use [UnityWebRequest](./web-request.md) to access any remote assets. If using the [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest/) package, this will handle asset downloads for you.

**Note**: Apple might change the folder locations where you’re permitted to write data. Make sure to check the latest Apple guidelines for the most up-to-date information.

## Additional resources

* [StoreKit](https://developer.apple.com/documentation/storekit) (Apple)
* [Unity IAP package](https://docs.unity3d.com/Packages/com.unity.purchasing@latest)
* [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest/)
* [AssetBundles](AssetBundlesIntro.md)