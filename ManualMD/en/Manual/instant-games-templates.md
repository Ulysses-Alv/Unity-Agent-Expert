# Facebook Instant Games templates

Load required settings for Facebook Instant Games and customize the appearance of the HTML page that hosts your Web player.

The default templates for Facebook Instant Games have the same structure and functionality as the [Web templates](webgl-templates.md) but are tailored for the Facebook Instant Games environment. There are separate templates for portrait and landscape orientations. They also include the required `fbapp-config.json` file necessary for uploading to Facebook Instant Games.

**Note:** The HTML templates include a Facebook SDK script tag comment. If you edit the templates, leave in the comment. While the [Facebook Instant Games SDK for Unity package](https://docs.unity3d.com/Packages/com.unity.meta-instant-games-sdk@latest?subfolder=/manual/index.html) automatically downloads and initializes the Facebook Instant Games SDK, the comment is still required for a successful upload.

## Facebook Instant Games configuration file

The `fbapp-config.json` file stores Facebook Instant Games configuration values including orientation, play configuration, and custom template variables.

When you [make custom templates](web-templates-add.md), you can modify settings in the configuration file to customize your game. For example, you can configure the template to send out custom updates.

Refer to [Facebook’s Bundle Configuration Reference documentation](https://developers.facebook.com/docs/games/build/instant-games/reference/bundle-config) for examples of customizations and more information about the configuration file.

## Additional resources

* [Using Web templates](web-templates-intro.md)
* [Web template variables](web-templates-variables.md)
* [Web template structure and instantiation](web-templates-structure.md)
* [Web template build configuration and interaction](web-templates-build-configuration.md)
* [Upload to Facebook Instant Games](instant-games-upload.md)