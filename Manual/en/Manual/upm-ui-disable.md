# Disable a built-in package

You can disable a **Built-in package** if you don’t need some modules and you want to save resources. However, when you disable a built-in package, the corresponding Unity functionality is no longer available.

Disabling a built-in package results in the following:

* If you use a Scripting API implemented by a disabled package, you get compiler errors.
* Components implemented by the disabled built-in package are also disabled, which means you can’t add them to any GameObjects. If you have a **GameObject** that already has one of these components, Unity ignores them in Play mode. You can view them in the ****Inspector**** window but they’re grayed out to indicate that they’re not available.
* When building a game, Unity strips all disabled components. For build targets which support engine code stripping (like Web, iOS, and Android), Unity doesn’t add any code from a disabled built-in package.

To disable a built-in package:

1. Open the **Package Manager** window select **Built-in packages** from the [navigation panel](upm-ui-nav.md).

   ![Switch the list context to Built-in packages](../uploads/Main/upm-ui-builtin.png)

   Switch the list context to **Built-in** packages
2. Select the built-in package you want to disable. Its information appears in the [details panel](upm-ui-details.md).
3. Click **Disable**.

   ![Disable button in the corner of the details panel](../uploads/Main/upm-ui-disable.png)

   **Disable** button in the corner of the details panel

   When the progress bar finishes, the check mark no longer appears next to the built-in package and the **Disable** button becomes an **Enable** button.
4. If you want to re-enable a disabled built-in package, click the **Enable** button.

**Note**: You can disable multiple built-in packages with one click by using the multiple select feature. For more information, refer [Perform an action on multiple packages or feature sets](upm-ui-multi.md).