# Web template build configuration and interaction

Configure and interact with Web builds using Web templates.

For information about template variables, JavaScript macros, and conditional directives, refer to [Web template variables](web-templates-variables.md).

## Build configuration

The configuration object contains the build configuration, which consists of code and data URLs, product name, company name, and version. You can define it using the following code:

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
  streamingAssetsUrl: "StreamingAssets",
  companyName: "{{{ COMPANY_NAME }}}",
  productName: "{{{ PRODUCT_NAME }}}",
  productVersion: "{{{ PRODUCT_VERSION }}}",
  // additional options can go here, see below
};
```

In the example above, the build folder URL is stored as a separate variable called `buildUrl`. This is useful in cases where you don’t know the relationship between the embedding page and build folder on the hosting server. It enables you to reuse the embedding code in other HTML documents. An example of when to use this is if you move the `Build` folder to another location on your server. You can adjust the value of the `buildUrl` variable on the embedding page, and you can use the same embedding code. This also applies to the `StreamingAssets` folder (`streamingAssetsUrl`).

In addition to setting the URL and field names on the build, you can set the following properties on the `config` object.

| Property | Use |
| --- | --- |
| `matchWebGLToCanvasSize: false` | By default (if set to true or unset), Unity synchronizes the **WebGL** canvas render target size with the Document Object Model (DOM) size of the canvas element (scaled by `window.devicePixelRatio`). Set this to false if you want to set the canvas DOM size and WebGL render target sizes manually. |
| `devicePixelRatio: <number>` | This field enables forcing the DPI scaling ratio for the rendered page. Set to 1 to force rendering to “standard DPI” (or non-Retina DPI), which can help performance on lower-end mobile devices. By default, this field is unset, meaning the rendered page uses the browser DPR scaling ratio, resulting in High DPI rendering. |
| `autoSyncPersistentDataPath: true` | If set to true, all file writes inside the Unity `Application.persistentDataPath` directory automatically persist so that the contents are remembered when the user revisits the website the next time. If unset (or set to false), you must manually sync file modifications inside the `Application.persistentDataPath` directory by calling the `JS_FileSystem_Sync()` JavaScript function. |

## Build interaction

After the build has successfully instantiated, the fulfillment handler callback of the Promise object receives the newly created Unity instance object as an argument. To interact with the build, call the following methods of the Unity instance:

| Method | Use |
| --- | --- |
| `unityInstance.SetFullscreen(fullscreen)` | The `SetFullscreen` method toggles full screen mode. This method doesn’t return a value. **Note**: You need to use a user interaction like a button or key press to activate fullscreen mode. You can’t activate fullscreen mode on startup.   - To activate Full screen mode, set the `fullscreen` argument to 1.   - To exit Full screen mode, set the `fullscreen` argument to 0. |
| `unityInstance.SendMessage(objectName, methodName, value)` | The `SendMessage` method sends messages to the GameObjects. This method doesn’t return a value.  - `objectName` is the name of an object in your Scene.   - `methodName` is the name of a method in the script, currently attached to that object.  - `value` can be a string, a number, or it can be empty. |
| `unityInstance.Quit().then(onQuit)` | Use the `Quit()` method to quit the runtime and clean up the memory used by the Unity instance. This method returns a Promise object. **Note**: Some browser add-ons interfere with the garbage collector’s ability to reclaim Unity content from the web page after calling `unityInstance.Quit()`. Disable browser add-ons to troubleshoot issues or refer to the Mozilla documentation for [`<iframe>`: The Inline Frame element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe) to quit and unload Unity applications with add-ons enabled.  - `onQuit` callback is called after the build runtime has quit. |

## Additional resources

* [Using Web templates](web-templates-intro.md)
* [Add a Web template](web-templates-add.md)
* [Web template structure and instantiation](web-templates-structure.md)
* [Web template variables](web-templates-variables.md)