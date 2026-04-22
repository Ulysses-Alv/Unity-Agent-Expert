# Web development and publishing process

Build and deploy a Web application in mobile and desktop environments.

Before you get started, check the [browser compatibility](webgl-browsercompatibility.md) and [technical limitations](webgl-technical-overview.md) for Web. Make sure you’re aware of any limitations for developing a project for this platform.

For a guided tutorial of the Web platform, refer to [Getting started with Unity Web](https://learn.unity.com/tutorial/getting-started-with-unity-web).

To build and deploy a Web application, complete the following tasks:

* [Add the Web module.](#add-module)
* [Set up a testing environment.](#test)
* [Profile and optimize.](#profile-optimize)
* [Host your build.](#host)

## Add the Web module

To get started with Web, [add the Web module](https://docs.unity3d.com/hub/manual/AddModules.html) to your project. Once you add the module, you can access the Web **Player** and **Build** settings.

## Set up a testing environment

Browser security policies restrict loading **scripts** from `file://` URLs, so you can’t open a Web build directly from your file system by double-clicking `index.html`.

To test your project, do either of the following:

* Use the **Build and Run** option in the Unity Editor’s [Build Profiles](build-profiles-reference.md) window.
* Run your own web server.

If you run your own web server, Unity doesn’t need to rebuild the project for every test. For a simple local server, try running one of these options:

* Python: `python -m http.server`
* Node.js: `npx http-server`

**Note**: WebGPU builds require a secure context. They only run on `localhost` or sites served over HTTPS.

For more advanced server configurations, refer to [Server configuration code samples](webgl-server-configuration-code-samples.md) and [Deploy a web application](webgl-deploying.md) for information about running your own server.

## Profile and optimize

Before you publish, gather performance metrics and reduce the build size to make a project that runs smoothly with the best quality possible.

[Profile your Web build](web-profile.md) to get performance data and refer to [Optimize your Web build](web-optimization.md) and [Optimize Web platform for mobile](web-optimization-mobile.md).

## Host your build

You need to host projects online to make them accessible to users. Use web hosting services such as [Unity Play](https://play.unity.com/en) depending on your needs.

## Additional resources

* [Unity Learn tutorial: Getting started with Unity Web](https://learn.unity.com/tutorial/getting-started-with-unity-web)
* [Web Player settings](class-PlayerSettingsWebGL.md)
* [Web build settings](web-build-settings.md)
* [Memory in Web](webgl-memory.md)
* [Web graphics](webgl-graphics.md)