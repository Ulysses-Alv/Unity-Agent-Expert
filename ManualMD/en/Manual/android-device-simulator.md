# Simulate an Android device

Unity’s [Device Simulator](device-simulator.md) can simulate the appearance and behavior of a variety of Android devices. This helps you test basic interactions and view the layout of your application on Android devices. This reduces costs, because you don’t need to own a device to simulate it, and helps you to improve device coverage. Additionally, the Device Simulator doesn’t require you to build your application which means you can debug layout issues and iterate quickly.

The Device Simulator doesn’t simulate the graphics backend of the target device and renders your application in the same way that the Editor does. This means that it doesn’t impose limitations that the graphics backend of the target device might have. The Device Simulator also doesn’t simulate the performance characteristics of the device, so to profile your application, you must create a build and run it directly on the device. For more information, refer to [Collecting performance data on an Android device](android-profile-on-an-android-device.md). For more information on the capabilities and limitations of the Device Simulator, refer to [Device Simulator introduction](device-simulator-introduction.md).

The device Simulator includes a variety of Android devices, but you can also add other devices if necessary. For information on how to do this, refer to [Add a device](device-simulator-adding-a-device.md).

## Additional resources

* [Debug on an Android device](android-debugging-on-an-android-device.md)
* [Unity Remote](UnityRemote5.md)