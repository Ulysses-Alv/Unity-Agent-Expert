# Licensing overview

Unity provides a range of plans for all levels. Some plans are free, while others have a subscription fee. Check out all the plans at <https://unity.com/pricing>.

Depending on your plan and type, your license will be governed by one of the following models:

* Named user licensing.
* Serial-based licensing.
* Floating licensing.

## Named user licensing

Named user licensing is a seat-based licensing solution designed for individual users in a network-connected environment. Users can be members in one or more organizations, with one or more subscriptions. When a user signs in, they’re granted all the entitlements from their organizations’ subscriptions.

Named user licensing requires an internet connection and requires you to sign in to the Unity Hub or the Unity Editor with your Unity ID credentials. Signing in activates your license on the current computer. To move your license to another computer, sign out of Unity applications, then sign in on another computer to continue your work there.

The license provides a grace period, in case you experience temporary internet connection issues. During this time, you can continue to use the software offline. However, you must reconnect to the internet before the grace period ends, so your license can be re-authenticated automatically.

If you’re an organization administrator, named user licensing provides improved license visibility and compliance, because you don’t need to track licensed devices.

## Serial-based licensing

Serial-based licensing is a node-locked licensing solution that manages your access through a serial key that Unity sends you. You activate the key (or return it) by using the Unity Hub or the command line.

The activation for a serial-based license is associated with the computer where you activated it.

This license type lets you work offline, after you [activate the license online](LicenseActivationMethods.md).

## Floating licensing

Floating licensing is a concurrency-based model designed for complex projects or larger organizations that need to support multiple users at the same time.

With this model, your organization sets up an on-premises licensing server, and buys a pool of licenses large enough to accommodate its members who use Unity concurrently. Whenever an organization member starts the Unity Editor from a client computer connected to the licensing server, the server removes a license from the pool. The server assigns that license to the client computer and the license can be used there for a session. When the last instance of a Unity Editor is closed on the client computer, the license is released. That license goes back to the pool, for others to use.

To find out more about floating licensing, contact [Unity Sales](https://unity.com/support-services). For information about setting up a floating licensing server, refer to the [Unity Licensing Server documentation](https://docs.unity.com/licensing/manual).

## Additional resources

* [Licenses and activation](LicensesAndActivation.md)
* [License activation methods](LicenseActivationMethods.md)
* [Unity Floating Licensing Server](https://docs.unity.com/licensing/en-us/manual)