# JavaScript interface in Unity Web builds

The JavaScript interface in Unity Web builds provides useful functions and variables you can use to configure your web applications. You can call these functions in your JavaScript library code (.jslib files).

## Display error, warning, or custom banners

Use the `unityShowBanner()` function to display a message banner, or override the function to customize it for your own purposes.

| **Parameter** | **Description** |
| --- | --- |
| `msg` | The message you want to display. |
| `type` | The type of message you want to display. For example:  * Use `type == error` to display a permanent error message on top of the canvas. * Use `type == warning` to show a yellow highlight color over the message. |

The following code from the default **WebGL** template sets the appearance of the error and warning banners. In your custom template, you can override this function to create your own banner types, or customize the way non-critical warnings and error messages display.

```
function unityShowBanner(msg, type) {
        var warningBanner = document.querySelector("#unity-warning");
        function updateBannerVisibility() {
          warningBanner.style.display = warningBanner.children.length ? 'block' : 'none';
        }
        var div = document.createElement('div');
        div.innerHTML = msg;
        warningBanner.appendChild(div);
        if (type == 'error') div.style = 'background: red; padding: 10px;';
        else {
          if (type == 'warning') div.style = 'background: yellow; padding: 10px;';
          setTimeout(function() {
            warningBanner.removeChild(div);
            updateBannerVisibility();
          }, 5000);
        }
        updateBannerVisibility();
      }
```

You can call the function in your .jslib file. The following banner permanently displays **The message you want to show** as an error message:

```
unityShowBanner('The message you want to show', 'error');
```

## Change WebGL context attributes

Use `webglContextAttributes` to customize the creation attributes that the WebGL context is initialized with. You can only change the following attributes:

* `powerPreference`
* `premultipliedAlpha`
* `preserveDrawingBuffer`

These options only apply when you use the WebGL rendering API. If you use the WebGPU rendering API, the `webglContextAttributes` field is ignored.

To change these attributes, change the configuration object in a custom WebGL template. For example:

```
script.onload = () => {
  config['webglContextAttributes'] = {

    powerPreference: 'high-performance',
    premultipliedAlpha: false,
    preserveDrawingBuffer: true, 
  };
  createUnityInstance(canvas, config, (progress) => {
```

The `powerPreference` attribute has the following options:

| **Option** | **Description** |
| --- | --- |
| default | Let the user agent decide which GPU configuration is most suitable. This is the default value. |
| high-performance | Prioritizes rendering performance over power consumption. |
| low-power | Prioritizes power saving over rendering performance. |

For more information about each attribute, refer to the Mozilla documentation on [HTMLCanvasElement: getContext()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/getContext).

For information about Unity WebGL graphics, refer to [WebGL graphics](webgl-graphics.md).

## Direct your project to your streaming assets

If you need to host your streaming assets on a different CDN, you can adjust the `streamingAssetsUrl` property on the `createUnityInstance()` configuration object to point to the full URL where your streaming assets are.

```
var buildUrl = "Build";
var config = {
  dataUrl: buildUrl + "/{{{ DATA_FILENAME }}}",
  frameworkUrl: buildUrl + "/{{{ FRAMEWORK_FILENAME }}}",
  codeUrl: buildUrl + "/{{{ CODE_FILENAME }}}",
#if MEMORY_FILENAME
  memoryUrl: buildUrl + "/{{{ MEMORY_FILENAME }}}",
#endif
#if SYMBOLS_FILENAME
  symbolsUrl: buildUrl + "/{{{ SYMBOLS_FILENAME }}}",
#endif
   streamingAssetsUrl: "https://mygameserver.com/StreamingAssets/",  // Add this line to override the default streaming assets location
  companyName: "{{{ COMPANY_NAME }}}",
  productName: "{{{ PRODUCT_NAME }}}",
  productVersion: "{{{ PRODUCT_VERSION }}}",
};
```

## Control cache of web pages

Use the `cacheControl()` function to control which URLs are cached in the `UnityCache`. The function accepts the URL as a parameter and returns either `must-revalidate`, `immutable`, or `no-store`. You can also override the function for your own purposes.

For a code example and more information, refer to [Cache behavior in Web](webgl-caching.md).

## Additional resources

* [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md)
* [Web graphics](webgl-graphics.md)
* [Cache behavior in Web](webgl-caching.md)