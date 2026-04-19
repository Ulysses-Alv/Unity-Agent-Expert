# Validate and upload a UPM package for the Asset Store (Beta)

**Important:** This is an early access workflow, and is available only to selected publishers to publish free Unity Package Manager (UPM) packages. To register your interest and apply for early access to UPM publishing, go to [UPM Publishing on the Asset Store](https://assetstore.unity.com/publishing/upm-publishing).

To validate a **UPM package** for the **Asset Store**, perform the following steps.

## Prerequisites

* [Set up a product draft in the Publisher Portal with a reserved package technical name](asset-store-upm-create-draft.md).
* Follow the instructions in [Package development workflow](CustomPackages.md) to create a UPM package. Your package must follow the same naming convention that you set in the draft package, for example `com.company.product.tools`.

## Validate and upload to the Editor

1. Install [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745).
2. Open the Unity Editor and go to **Window** > **Tools** > **Asset Store** > **Validator**.
3. Select **UPM** as the **Validation Type** and run the Validator.
4. After you’ve fixed any issues the Validator highlights, open the **Asset Store Uploader** (**Window** > **Tools** > **Asset Store** > **Uploader**).
5. Select the **UPM Packages** tab, and then select **Upload** on the package you want to upload to the Asset Store.

Once the upload is successful, a link to the Publisher Portal is available, where you can continue to configure product listing. In the Product Portal you can fill in the package’s metadata details, product details, media, and submit the product for approval. For more information, refer to [Submit an UPM package for approval to the Asset Store (Beta)](asset-store-upm-submit.md).

## Additional resources

* [Asset Store UPM Publishing Tools](https://assetstore.unity.com/packages/package/5368745)
* [Submit an UPM package for approval to the Asset Store (Beta)](asset-store-upm-submit.md)
* [Check the status of a UPM package for the Asset Store (Beta)](asset-store-upm-check-status.md)