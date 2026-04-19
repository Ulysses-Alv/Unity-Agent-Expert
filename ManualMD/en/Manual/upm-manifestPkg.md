# Package manifest reference

The Unity Editor uses a **package manifest** file (`package.json`) to manage information about a specific version of a specific package.

The package manifest uses JSON (JavaScript Object Notation) syntax to describe what’s inside the package. The file’s format is similar to [npm](https://www.npmjs.com/)’s `package.json` format, but uses different semantics for some of its properties. Refer to the [example](#example) for a sample package manifest file.

The Package Manager reads this manifest to find out the package’s contents, how to unpack its contents, and what information to display in the [Package Manager window](upm-ui.md). The manifest stores this information in a series of [required](#required), [recommended](#mandatory), and [optional](#optional) properties.

## Required properties

If these properties aren’t present, either the registry refuses the package when it’s published, or the Package Manager can’t fetch or load the package.

| **Property** | **JSON Type** | **Description** |
| --- | --- | --- |
| **name** | String | A unique multi-part identifier that conforms to the Unity Package Manager naming convention, which uses reverse domain name notation. For more information about the naming convention, refer to [Package name guidelines](cus-naming.md).  **Note**: The `name` property is different from the [user-friendly display name](#displayName) that appears in the **Package Manager** window [list panel](upm-ui-list.md) and the top of the [details panel](upm-ui-details.md). The `name` property displays fully in the [details panel](upm-ui-details.md) as the **Technical Name** value, in the **Technical Details** tab. In the **Inspector** window, only the final part of a package’s `name` property displays in the **Name** field. |
| **version** | String | The package version number, which uses the format `"major.minor.patch"`.   For example, `"3.2.1"` indicates that this is the third major release, the second minor release, and the first patch.   This value must respect [Semantic Versioning](http://semver.org/). For more information, refer to [Package versioning](upm-semver.md). |

## Recommended properties

The Package Manager can install packages in a project even if the recommended properties are missing or have invalid values.

However, the recommended best practice is to assign values for these properties to make sure that your package is discoverable and to offer a better experience for users.

| **Property** | **JSON Type** | **Description** |
| --- | --- | --- |
| **description** | String | A brief description of the package. This is the text that appears in the [details panel](upm-ui-details.md) of the **Package Manager** window. This field supports [UTF–8](https://en.wikipedia.org/wiki/UTF-8) character codes. This means that you can use special formatting character codes, such as line breaks (`\n`) and bullets (`\u25AA`). |
| **displayName** | String | A user-friendly name that appears in the Unity Editor, such as in the **Project** window and the **Package Manager** window.   Examples of `displayName` values are **Unity Timeline**, **ProBuilder**, and **In App Purchasing**. |
| **unity** | String | Indicates the lowest Unity version the package is compatible with. If omitted, the Package Manager considers the package compatible with all Unity versions.   The expected format is `"major.minor"` (for example, `"2018.3"`). Unity 6.0 and later versions follow the same `"####.#"` format. For example, the technical version number for Unity 6.0 is `"6000.0"`.  To point to a specific patch, also include the `unityRelease` property, described in [Optional properties](#optional).  **Note**: Installing a package in a project using an Editor lower than the value specified in `unity` might cause compilation or compatibility issues. |

## Optional properties

These properties are optional, meaning that you can omit them. However, if they’re present, they must have a valid value.

| **Property** | **JSON Type** | **Description** |
| --- | --- | --- |
| **author** | Object or string | The author of the package. This property supports only one author.   This property has one required field, **name**, and two optional fields, **email** and **url**.   You can specify these fields as a JSON object, or collapse them into a single string whose key is **author**.  Object example:  `{`    `"name" : "John Doe",`    `"email" : "john.doe@example.com",`    `"url" : "http://john.doe.example.com/"` `}`   String example:  `"John Doe <john.doe@example.com> (http://john.doe.example.com/)"` |
| **changelogUrl** | String | Custom location for this package’s changelog specified as a URL. For example: `"changelogUrl": "https://example.com/changelog.html"`  **Note**: When the Package Manager can’t reach the URL location (for example, if there’s a network issue), it does the following:  - If the package is installed, the Package Manager opens a file browser displaying the `CHANGELOG.md` file from the package cache.  - If the package isn’t installed, the Package Manager displays a warning that an offline changelog isn’t available. |
| **dependencies** | Object | A map of package dependencies. Keys are package names, and values are specific versions. They identify other packages that this package depends on.   **Note**: The Package Manager doesn’t support range syntax, only **SemVer** versions. |
| **documentationUrl** | String | Custom location for this package’s documentation, specified as a URL. For example: `"documentationUrl": "https://example.com/"`  **Note**: When the Package Manager can’t reach the URL location (for example, if there’s a network issue), it does the following:  - If the package is installed, the Package Manager opens a file browser displaying the `Documentation~` folder from the package cache.  - If the package isn’t installed, the Package Manager displays a warning that offline documentation isn’t available. |
| **hideInEditor** | Boolean | By default, the **Project** window hides a package’s assets and omits them from the results when you use the [Object Picker](InspectorReferences.md) in the **Inspector** window. Set this property to `"false"` to make sure that this package’s assets are always visible. |
| **keywords** | Array of Strings | An array of keywords used by the Package Manager search APIs. This helps users find relevant packages. |
| **license** | String | Identifier for an OSS license using the [SPDX identifier format](https://spdx.org/licenses/), or a string such as ‘Refer to LICENSE.md file’.  **Note**: If you omit this property in your package manifest, your package must contain a `LICENSE.md` file. |
| **licensesUrl** | String | Custom location for this package’s license information, specified as a URL. For example: `"licensesUrl": "https://example.com/licensing.html"`  **Note**: If the Package Manager can’t reach the URL location (for example, if there’s a network issue), it does the following:  - If the package is installed, it opens a file browser displaying the `LICENSE.md` file from the package cache.  - If the package isn’t installed, the Package Manager displays a warning that offline license information isn’t available. |
| **samples** | Array of Objects | List of samples included in the package. Each sample has a display name, a description, and the path to the sample folder starting at the `Samples~` folder:  `{`    `"displayName": "<name-to-appear-in-the-UI>",`    `"description": "<brief-description>",`    `"path": "Samples~/<sample-subfolder>"` `}`   For more information, refer to [Creating samples for packages](cus-samples.md). |
| **type** | String | **Reserved for internal use.** |
| **unityRelease** | String | Part of a Unity version indicating the specific release of Unity that the package is compatible with. You can use this property when an updated package requires changes made during the Unity alpha/beta development cycle. This might be the case if the package needs newly introduced APIs, or uses existing APIs that have changed in a non-backward-compatible way without [API updater](APIUpdater.md) rules.   The expected format is `"<update><release>` (for example, `"0b4"`).   **Note**: If you omit the recommended [unity](#unity) property, this property has no effect.   A package that’s not compatible with Unity doesn’t appear in the [Package Manager window](upm-ui.md). |

## Package manifest example

```
{
  "name": "com.[company-name].[package-name]",
  "version": "1.2.3",
  "displayName": "Package Example",
  "description": "This is an example package",
  "unity": "2019.1",
  "unityRelease": "0b5",
  "documentationUrl": "https://example.com/",
  "changelogUrl": "https://example.com/changelog.html",
  "licensesUrl": "https://example.com/licensing.html",
  "dependencies": {
    "com.[company-name].some-package": "1.0.0",
    "com.[company-name].other-package": "2.0.0"
 },
 "keywords": [
    "keyword1",
    "keyword2",
    "keyword3"
  ],
  "author": {
    "name": "Unity",
    "email": "unity@example.com",
    "url": "https://www.unity3d.com"
  }
}
```

## Additional resources

* [Edit the package manifest](cus-edit-manifest.md)
* [Project manifest file](upm-manifestPrj.md)