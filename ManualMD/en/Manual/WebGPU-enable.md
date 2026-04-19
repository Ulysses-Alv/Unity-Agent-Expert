# Enable the WebGPU graphics API

Enable WebGPU in the **Player** settings so you can access its modern graphics features.

Unity doesn’t automatically enable WebGPU, unlike [WebGL2](WebGL2.md). This information covers how to enable the appropriate setting in your project so you can use WebGPU features. For information about features and limitations of WebGPU, refer to [WebGPU](WebGPU.md).

To enable WebGPU in Unity:

1. [Add WebGPU to your project’s supported web graphics APIs](#add-webgpu-to-your-projects-supported-web-graphics-apis).
2. [Choose whether you want to keep WebGL2 as a fallback option](#choose-whether-you-want-to-keep-webgl2-as-a-fallback-option).

## Add WebGPU to your project’s supported web graphics APIs

To add the WebGPU web graphics API to your supported graphics list in your Unity project:

1. Open the Project Settings window (menu: **Edit** > **Project Settings**).
2. Select **Player**. This opens the **Player** settings.
3. Select **Web settings**.
4. Expand **Other Settings**.
5. Disable **Auto Graphics API**. Unity shows the **Graphics API** list, which is a list of enabled graphics APIs for Web.
6. Select **Plus (+)** and choose **WebGPU**.

WebGPU is now a supported web graphics API for your project. Continue to [Choose whether you want to keep WebGL2 as a fallback option](#choose-whether-you-want-to-keep-webgl2-as-a-fallback-option).

## Choose whether to keep WebGL2 as a fallback

Unity prioritizes the first API on the **Graphics API** list as the default web graphics API for your project. If a browser doesn’t support that API or there’s another issue, Unity tries the next API on the list (if one exists) as a fallback. Therefore, you have the option to keep support of both APIs available, or use WebGPU only.

It’s useful to keep both because WebGPU is still experimental and not all browsers support it, so users who have older browsers can still access your application.

However, there are reasons you might want to remove WebGL2:

* Your project relies on WebGPU features that WebGL2 doesn’t support.
* You need to create extra code to cater to 2 different graphics APIs, which makes it more complex to develop and test.

### Keep WebGL2 as a fallback

To use WebGPU as the priority web graphics API but keep WebGL2 as a fallback:

1. Make sure WebGL2 is in the **Graphics API** list. If it’s not, select **Add (+)** and choose **WebGL 2**.
2. Click and drag **WebGPU** to the top of the list.

Your project now supports both WebGPU and WebGL2, but WebGPU is the priority.

### Remove WebGL2 from the supported graphics API list

To remove WebGL2 from the supported web graphics API list:

1. Select the **WebGL 2** entry in the list.
2. Select **Remove (-)**.

Your project now only uses the WebGPU web graphics API.

## Additional resources

* [WebGPU (Experimental)](WebGPU.md)
* [Introduction to the WebGPU graphics API](WebGPU-features.md)
* [Limitations of the WebGPU graphics API](WebGPU-limitations.md)
* [Recommendations for both WebGL2 and WebGPU APIs](web-graphics-apis-intro.md)