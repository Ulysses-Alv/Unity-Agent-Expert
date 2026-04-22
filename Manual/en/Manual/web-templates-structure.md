# Web template structure and instantiation

Use custom user variables in a template’s `index.html` file and create a new instance of your content with `createUnityInstance()`.

For information about template variables, JavaScript macros, and conditional directives, refer to [Web template variables](web-templates-variables.md).

To add a custom Web template, refer to [Add a custom Web template](web-templates-add.md).

## Custom user variables

When you select a Web template, Unity parses the template and looks for [JavaScript macros and conditional directives](web-templates-variables.md).

Unity treats JavaScript variables as custom user variables if they’re:

* Used inside JavaScript macros and conditional directives.
* Not declared in the template code.
* Not internal pre-processor variables.

Unity automatically adds these custom user variables to the **Resolution and Presentation** section in the ****Player Settings**** window.

As an example, if you want to control the title of the generated `index.html` page directly from the **Player Settings** window, you first need to modify the `<title>` line of the `index.html` in your custom template, like this:
`lang-html
<title>{{{ PAGE_TITLE }}}</title>`
After you have done this, re-select your custom template. This parses the template again, and a **Page Title** field appears in the **Resolution and Presentation** > **Web Template** section of the **Player Settings** window.

![Image of Resolution and Presentation window with custom template](../uploads/Main/WebGLResolutionandPresentationCustomTemplate.png)

Image of Resolution and Presentation window with custom template

When you enter text into this field and build your project, the custom variable `PAGE_TITLE` in the template macro automatically becomes the text in the **Page Title** field.

If you would like to use custom integer or float variables in your macros, use `parseInt()` or `parseFloat()` JavaScript functions in your macros to preprocess string values supplied by the Unity Editor. This is because custom user variables are always assigned a string value.

**Note**: Underscore characters in variable names display as spaces inside the field to improve readability.

## Structure of the index.html file

The `index.html` contains the code necessary to load the build and should include the following:

* A `<canvas>` element. Unity runtime uses the `<canvas>` element to render the application.
* JavaScript code to download the build loader. For example:

```
var buildUrl = "Build";
var loaderUrl = buildUrl + "/{{{ LOADER_FILENAME }}}";
var script = document.createElement("script");
script.src = loaderUrl;
script.onload = () => {
  // code for instantiating the build
};
document.body.appendChild(script);`
```

In this example, `{{{ LOADER_FILENAME }}}` is automatically resolved by the template preprocessor when the build is generated.

Alternatively, you can download the build loader using a script tag, for example:
`lang-js
<script src="Build/{{{ LOADER_FILENAME }}}"></script>`

* JavaScript code for instantiating a build. Unity builds are instantiated using the `createUnityInstance()` function, which is defined in the build loader script.

## The instantiation function: createUnityInstance()

The `createUnityInstance()` function creates a new instance of your content. You can use it like this:
`lang-js
createUnityInstance(canvas, config, onProgress).then(onSuccess).catch(onError);`

This function returns a `Promise` object, where:

| **Object** | **Use** |
| --- | --- |
| `canvas` | Unity runtime uses the `canvas` object to render the game. |
| `config` | The `config` object contains the build configuration, such as the code and data URLs, product and company name, and version. For more information on config definition, refer to [Web template build configuration and interaction](web-templates-build-configuration.md). |
| `onProgress(progress) {...}` | The Web loader calls the `onProgress` callback object every time the download progress updates. The `progress` argument that comes with the `onProgress` callback determines the loading progress as a value between 0.0 and 1.0. |
| `onSuccess(unityInstance) {...}` | The `onSuccess` callback is called after the build has successfully instantiated. The created Unity instance object is provided as an argument. This object can be used for interaction with the build. |
| `onError(message) {...}` | The `onError` callback is called if an error occurs during build instantiation. An error message is provided as an argument. |

The `createUnityInstance()` function is defined in the build loader script and is specific to the instantiated build. Therefore, if you’re embedding two or more builds in the same HTML document, make sure that the `createUnityInstance()` function is called from an `onload` callback of the corresponding build loader script. For more information about the Unity Web Loader, refer to [Web Build folder](webgl-building.md) and [Web template structure and instantiation](web-templates-structure.md).

## Additional resources

* [Using Web templates](web-templates-intro.md)
* [Add a custom Web template](web-templates-add.md)
* [Web template variables](web-templates-variables.md)
* [Web template build configuration and interaction](web-templates-build-configuration.md)