# Deep linking on iOS

Deep links are hyperlinks outside of your application that take a user to a specific location within the application rather than a website. When a user clicks a deep link, the application opens from the designated location, such as a specific **scene** in a Unity application. For more information about deep links and how to use them, refer to [Deep linking](deep-linking.md).

There are two ways to enable deep links for iOS applications: URL schemes and universal links.

For information on how to use deep links and handle them when your application opens, refer to [Using deep links](deep-linking.md#using-deep-links).

## URL schemes

A URL scheme specifies a link structure that your iOS application refers to. The device opens the application when the user clicks a deep link that matches the URL scheme structure. To add a URL scheme, use the following steps:

1. Go to **Edit** > **Project Settings** > **Player** > **Other Settings** > **Configuration**.
2. Expand **Supported URL schemes** to set the following properties:
   * **Size** property to `1`.
   * **Element 0** property to the URL scheme to use with your application. For example, use `unitydl` to open your application when the device processes a link that starts with `unitydl://`.

![Supported URL schemes settings for iOS.](../uploads/Main/ios-supported-url.png)

Supported URL schemes settings for iOS.

Your iOS application now opens when the device processes any link that starts with `unitydl://`.

**Note**: To use multiple URL schemes in your project, increase the value of the **Size** property.

## Universal links

Universal links are similar to deep links because they open a specific location within an application. However, if the application isn’t installed, a universal link opens a URL in Safari. To enable universal links, refer to Apple’s documentation on [Allowing Apps and Websites to Link to Your Content](https://developer.apple.com/documentation/xcode/allowing-apps-and-websites-to-link-to-your-content).

## Additional resources

* [Deep linking](deep-linking.md)