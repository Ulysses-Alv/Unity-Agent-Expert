# Update a UPM package for the Asset Store (Beta)

**Important:** This is an early access workflow, and is available only to selected publishers to publish free Unity Package Manager (UPM) packages. To register your interest and apply for early access to UPM publishing, go to [UPM Publishing on the Asset Store](https://assetstore.unity.com/publishing/upm-publishing).

You can update a published package to release a new version of the package with bug fixes or new features, or to change product details such as pricing and media. When you publish an update, users can access the new version in the [Package Manager window](upm-ui.md).

You can make two types of updates, and each has different required information as follows:

| **Update type** | **Required actions** | **Optional actions** |
| --- | --- | --- |
| **New package version** | * Upload a new package version. * Edit or fill in package metadata. | Edit product metadata or media. |
| **Update product information metadata or media** | Edit product metadata or media. | None. |

## Prerequisites

* Install the [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745) in the Unity Editor.
* Have a published product ready on Publisher Portal.

## Create a package draft

1. Use the unique link in your confirmation email to open the Publisher Portal, or return to the Publisher Portal’s Product Overview page.
2. Select the published package.
3. Select **Create Draft**.

## Validate and update the package

1. [Edit the package manifest file](cus-edit-manifest.md#edit-manifest) by updating the [version property](upm-manifestPkg.md#required). The version number of the package must be unique, and the technical naming consistent.
2. Use the ****Asset Store** Validator** window in the [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745) validate the package.
3. Use the **Asset Store Uploader** window to upload a new version of the package.

## Update details of the package

1. In the Publisher Portal, select **Fill out your package metadata** for the updated package, and add the compatibility information, release notes, and package information.
2. In the **Product Information** tab, review the existing information.
3. In the **Media** tab, review the existing uploaded media.
4. Select **Submit**.

## Additional resources

* [Package versioning](upm-semver.md)
* [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745)
* [Validate and upload a UPM package for the Asset Store (Beta)](asset-store-upm-validate.md)