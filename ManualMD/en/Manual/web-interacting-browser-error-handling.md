# Customize error handling

Choose how to handle errors in your Web builds.

Unity writes some JavaScript code to builds by default, which handles errors. This code displays warning messages when running on unsupported platforms and parses error and exception strings from the browser to check for known errors and display error dialogs with more comprehensive error messages. You can customize this handling, for instance, if you want to suppress the warning messages.

To customize error handling:

1. Open the `index.html` file in your [Web template](webgl-templates.md).
2. Add or uncomment the following code as a parameter in the `var config = {}` object:

   ```
     errorHandler: function(err, url, line) {
       alert("error " + err + " occurred at line " + line);
       // Return 'true' if you handled this error and don't want Unity
       // to process it further, 'false' otherwise.
       return true;
     },
   ```

The `errorhandler` function is called when the page invokes its `window.onerror` event handler and uses the same parameters. Return `true` from this function to signal that your error handler has processed the error, in order to suppress Unity’s default error handler from executing. Return `false` to pass the error to Unity’s default error handler to process.

## Additional resources

* [Interaction with browser scripting](webgl-interactingwithbrowserscripting.md)
* [Set up your JavaScript plug-in](web-interacting-browser-js.md)
* [Call JavaScript functions from Unity C# scripts](web-interacting-browser-js-to-unity.md)
* [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md)
* [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md)
* [Create callbacks between Unity C#, JavaScript, and C/C++/C# code](web-interacting-browser-example.md)