# Deep linking on Android

Deep links are hyperlinks outside of your application that take a user to a specific location within the application rather than a website. When a user clicks a deep link, the application opens from the designated location, such as a specific **scene** in a Unity application. For more information about deep links and how to use them, refer to [Deep linking](deep-linking.md).

## Enable deep linking for Android applications

Use an [intent filter](https://developer.android.com/guide/components/intents-filters) to enable deep linking for Android applications. An intent filter overrides the standard [Android App Manifest](android-manifest.md) to include a specific intent filter section for [Activity](https://developer.android.com/reference/android/app/Activity).

To set up an intent filter, use the following steps:

1. Navigate to **Edit** > **Project Settings** > **Player**.
2. In the **Android settings** tab, expand the **Publishing Settings** section.
3. In the **Build** section, enable **Custom Main Manifest**. This creates a new file called `AndroidManifest.xml` in `Assets/Plugins/Android`.
4. In the Project window, go to **Assets** > **Plugins** > **Android** and open the `AndroidManifest.xml` file.
5. Add the following code sample inside the Unity`<activity>` element, named `com.unity3d.player.UnityPlayerGameActivity` or `com.unity3d.player.UnityPlayerActivity`, and save the file.

```
<intent-filter>
  <action android:name="android.intent.action.VIEW" />
  <category android:name="android.intent.category.DEFAULT" />
  <category android:name="android.intent.category.BROWSABLE" />
  <data android:scheme="unitydl" android:host="mylink" />
</intent-filter>
```

The built application now opens when the device processes any link that starts with `unitydl://`.

## Use deep linking on Android

After you enable deep links for Android, the way that you use them is platform-agnostic. For information on how to handle deep links when your application opens, refer to [Using deep links](deep-linking.md#using-deep-links).