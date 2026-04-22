# Web template variables

During the build process, Unity pre-processes template files and evaluates all macros and conditional directives included in those files. As part of this process, Unity finds and replaces all macro declarations with the values the Unity Editor supplies. Unity automatically pre-processes all .html, .php, .css, .js and .json files in the template folder.

## Internal preprocessor variables

The following internal preprocessor variables refer to data in the Project, and Unity assigns value to them at build time according to the values the Editor supplies. JavaScript macros and conditional directives can use these internal preprocessor variables.

| **Variable** | **Type** | **Description** |
| --- | --- | --- |
| COMPANY\_NAME | String | The Company Name defined in the **Player Settings**. |
| PRODUCT\_NAME | String | The Product Name defined in the Player Settings. |
| PRODUCT\_VERSION | String | The Version defined in the Player Settings. |
| WIDTH | Integer | The Default Canvas Width defined in the Player Settings > Resolution and Presentation. |
| HEIGHT | Integer | The Default Canvas Height in the Player Settings > Resolution and Presentation. |
| SPLASH\_SCREEN\_STYLE | String | This is set to the “Dark” value when Splash Style Player Settings > Splash Image is set to Light on Dark, otherwise it’s set to the “Light” value. |
| BACKGROUND\_COLOR | String | Represents the Background Color defined in a form of a hex triplet. |
| UNITY\_VERSION | String | The Unity version. |
| DEVELOPMENT\_PLAYER | Boolean | This is set to true if the **Development Build** option is enabled. |
| DECOMPRESSION\_FALLBACK | String | This is set to `Gzip` or `Brotli`, depending on the **compression** method you use and the type of decompressor included in the build. If neither is included, the variable is set to an empty string. |
| INITIAL\_MEMORY | Integer | The initial size of the WASM memory heap in megabytes (MB). |
| USE\_WASM | Boolean | This is set to true if the current build is a WebAssembly build. |
| USE\_THREADS | Boolean | This is set to true if the current build uses threads. |
| USE\_WEBGL\_1\_0 | Boolean | This is set to true if the current build supports the WebGL1.0 graphics API. |
| USE\_WEBGL\_2\_0 | Boolean | This is set to true if the current build supports the WebGL2.0 graphics API. |
| USE\_DATA\_CACHING | Boolean | This is set to true if the current build uses indexedDB caching for the downloaded files. |
| LOADER\_FILENAME | String | This is set to the file name of the build loader script. |
| DATA\_FILENAME | String | This is set to the file name of the main data file. |
| FRAMEWORK\_FILENAME | String | This is set to the file name of the build framework script. |
| CODE\_FILENAME | String | This is set to the file name of the WebAssembly module when the current build is a WebAssembly build, otherwise it’s set to the file name of the `asm.js` module. |
| MEMORY\_FILENAME | String | This is set to the file name of the memory file when memory is stored in an external file, otherwise it’s set to an empty string. |
| SYMBOLS\_FILENAME | String | This is set to the file name of the JSON file containing debug symbols when the current build is using debug symbols, otherwise it’s set to an empty string. |
| BACKGROUND\_FILENAME | String | This is set to the file name of the background image when the background image is selected in the **Player Settings** > **Splash Image**, otherwise it’s set to an empty string. |

## JavaScript macros

JavaScript macros are blocks of JavaScript code in template files, surrounded by three sets of curly brackets. This JavaScript code can use the internal preprocessor variables listed above. These variables are assigned at build time according to the values the Editor supplies. During the build, the preprocessor evaluates all macros and replaces them with the output of the variable.

You can use JavaScript macros to preprocess values supplied by the Editor. These macros can be as complex as you like. They can include multiple operators, loops, functions, and any other JavaScript constructs.

The following example line is from the `index.html` file used in the Default template:

`<div id="unity-build-title">{{{ PRODUCT_NAME }}}</div>`

If the value of the **Product Name** in the Player Settings is set to `My WebGL Game`, the internal preprocessor variable `PRODUCT_NAME` has the value `My WebGL Game`. In the output `index.html` file, the line appears as:

`<div id="unity-build-title">My WebGL Game</div>`

Below is a more complex example from the same `index.html` template file:

`canvas.style.background = "url('" + buildUrl + "/{{{ BACKGROUND_FILENAME.replace(/'/g, '%27') }}}') center / cover";`

If the target build folder is called `Let’s try WebGL`, and if you select a background image in the Player Settings, the internal preprocessor variable `BACKGROUND_FILENAME` has the value `Let’s try WebGL.jpg`. In the output `index.html` file, the line changes into:

`canvas.style.background = "url('" + buildUrl + "/Let%27s try WebGL.jpg') center / cover";`

## Conditional directives

Conditional directives **#if**, **#else**, and **#endif** control whether Unity includes a specific portion of the preprocessed file in the output file, or discards it for the current build.

Code that starts with an **#if** directive and ends with an **#endif** directive is called a conditional group. Conditional groups can also include **#else** directives. Unity evaluates the expression written after the **#if** as a JavaScript expression. If this expression has a truthy value that translates to true when evaluated in a Boolean context, then Unity keeps the line group immediately following the **#if** directive in the output file. If the **#if** expression is false, and an **#else** directive is included in the conditional group, Unity keeps the line group immediately following the **#else** directive in the output.
An example of a conditional group is as follows:

```
#if EXPRESSION
  // this block is included in the output if EXPRESSION has a truthy value
#else
  // this block is included in the output otherwise
#endif
```

Evaluated JavaScript expressions can include brackets, logical operators and other JavaScript constructs. Conditional directives can be nested.

## Additional resources

* [Using Web templates](web-templates-intro.md)
* [Add a custom Web template](web-templates-add.md)
* [Web template structure and instantiation](web-templates-structure.md)
* [Web template build configuration and interaction](web-templates-build-configuration.md)