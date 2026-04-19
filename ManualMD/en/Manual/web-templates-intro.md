# Using Web templates

When you build a Web project, Unity embeds the Player in a HTML page such that a browser can open it. A Web template is a configuration setting that lets you control the appearance of this HTML page, so that you can test, demonstrate, and preview your Web application in a browser.

To access Web templates:

1. Go to [Player settings](class-PlayerSettings.md) (menu: **Edit** > **Project Settings** > **Player**).
2. Set the platform-specific settings to Web.
3. Open **Resolution and Presentation**.

![Image of Player Resolution and Presentation window. ](../uploads/Main/WebGLResolutionandPresentationWindow.png)

Image of Player Resolution and Presentation window.

## Built-in templates

By default, the **Web Template** setting has the following options:

* **Default**: A white page with a loading bar on a grey canvas.
* **Minimal**: A minimal Web template that includes the necessary boilerplate code to run the Web content.
* **PWA:** A **Progressive Web App** including a web manifest file and service worker code.

These built-in HTML pages are useful for testing and demonstrating a minimal Player.

## Custom templates

You can also use JavaScript to build and supply your own Web templates to host the Player. This is useful for production purposes, so that you can preview the Player hosted in the page where it’s deployed. For example, if the Unity Player content interacts with other elements in the page via the external call interface, test it with a page that contains those interacting elements.

## Additional resources

* [Add a custom Web template](web-templates-add.md)
* [Web template structure and instantiation](web-templates-structure.md)
* [Web template variables](web-templates-variables.md)
* [Web template build configuration and interaction](web-templates-build-configuration.md)