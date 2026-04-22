# Add required legal files to your package

Add licensing and attribution information to your package to meet any legal requirements.

Your package might require one or both of the following:

* License information to reflect any licenses that govern your package.
* Third-party notices (separately from licensing information), if your package has third-party elements.

## License information

Use `LICENSE.md` to explain how others can use your package. For example, specify whether it’s for internal use or public distribution. Include rules around modification, redistribution, and any restrictions.

**Important**: Package Manager’s **Create package** operation doesn’t create this file for you.

Follow these steps if you want to include licensing information:

1. Use your computer’s file management application or Unity’s **Project** window to locate your package folder.
2. Create a file named `LICENSE.md` in the base folder of your package.
3. Open `LICENSE.md` using your preferred script editor.
4. Enter the licensing information about your package, optionally formatting it with Markdown. Refer to the [License file example](#license-file-sample).
5. Save the file.

### License file sample

Here’s a sample license file from the Unity Timeline package:

```
Unity Timeline copyright © 2017-2025 Unity Technologies ApS

Licensed under the Unity Companion License for Unity-dependent projects.
Refer to [Unity Companion License](https://unity.com/legal/licenses/unity-companion-license).

Unless expressly provided otherwise, the Software under this license is made available 
strictly on an “AS IS” BASIS WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED.
Please review the license for details on these and other terms and conditions.
```

## Third party notices

When you [create your package](cus-create.md) with the Package Manager window, the operation creates a `Third Party Notices.md` file in base folder of your package. Update this file if your package has third-party elements.

You can include the following information for each license you want to include:

* **Component Name**
* **License Type**
* **Version Number**: Specify the version number if the license governs a specific version of the component.
* **License Details** section. Refer to the [SPDX License List](https://spdx.org/licenses/) website. It has many common software components’ license text, which you might be able to copy into your `Third Party Notices.md` file. As an alternative, you can also point this section’s URL to a location that contains the reproduced license and the copyright information (if applicable). If the license governs a specific version of the component, give the URL of that version, if possible. Many of the entries from the SPDX website also include a canonical permanent URL you can link to.

If your package doesn’t contain third-party elements, delete `Third Party Notices.md` from the base folder of your package.

### Third party notices sample

```
This package contains third-party software components governed by the license(s) indicated below:

Component Name: Semver
License Type: "MIT"
[SemVer License](https://github.com/myusername/semver/blob/master/License.txt)

Component Name: MyComponent
License Type: "MyLicense"
[MyComponent License](https://www.mycompany.com/licenses/License.txt)
```

## Additional resources

* [Package development workflow](CustomPackages.md)
* [Markdown syntax guide](https://confluence.atlassian.com/bitbucketserver/markdown-syntax-guide-776639995.html) (Bitbucket)