# UPM package Asset Store publishing workflow (Beta)

**Important:** This is an early access workflow, and is available only to selected publishers to publish free Unity Package Manager (UPM) packages. To register your interest and apply for early access to UPM publishing, go to [UPM Publishing on the Asset Store](https://assetstore.unity.com/publishing/upm-publishing).

You can create a [Unity Package Manager (UPM) package](upm-package-types.md) from assets that you own and make it available to download on Unity’s [Asset Store](https://assetstore.unity.com/). Manage **UPM packages** through the UPM publishing portal, which differs from the [asset package publishing workflow](asset-store-workflow.md).

**Important:** An Asset Store package must meet certain legal requirements. For more information, refer to the [Asset Store Provider Agreement](https://unity.com/legal/provider) and the [Submission Guidelines](https://unity3d.com/asset-store/sell-assets/submission-guidelines).

## UPM package and products

The early access UPM publishing workflow introduces the concept of a product. A product is different from a package in the following ways:

* A product is a container that can include one or more UPM packages.
* Customers can get a product from the Asset Store.
* A package stores **scripts**, features, and assets for Unity, including Editor or Runtime tools and libraries, asset collections, and project templates.

You create a [product draft](asset-store-upm-create-draft.md) in the Publisher Portal to manage and configure the structure of your packages. A product can contain one or multiple packages. For more information, refer to [Create a Publisher Portal package draft (Beta)](asset-store-upm-create-draft.md).

## UPM package namespaces

UPM packages require consistent namespaces. The following table shows each of the namespaces you need to set for the different stages of the UPM publishing workflow:

| **UPM workflow step** | **Namespace action** | **Description** |
| --- | --- | --- |
| **Sign up to the early access UPM workflow** | Set a publisher namespace. | Establish and claim a namespace for your packages during the early access enrollment process. |
| **Create a UPM product draft for a single package** | Set a product namespace. | Assign a product namespace. The product namespace becomes part of the technical name. |
| **Create a UPM product draft for multiple packages** | Set a product namespace and a package namespace for each package. | Assign a package namespace for each package you add. The package namespace becomes part of the technical name. |
| **Create and validate UPM packages** | Set the UPM package folder and `package.json` namespaces in the `name` field. | The UPM package folder and the `package.json` name must match the package’s technical name. |
| **Upload UPM packages** | The Asset Store Publishing Tools checks and matches the package technical name to the one set on the Publisher Portal. | If the publisher space doesn’t match the product namespace, a button appears that you can select to create a new ID. If the technical name doesn’t match the product namespace, a button appears that you can select to create a new product draft. |

For more information about setting namespaces, refer to [Create a Publisher Portal package draft (Beta)](asset-store-upm-create-draft.md).

## Prerequisites

To publish assets to the UPM publishing portal, you must first create the following:

1. [Create a Unity ID](https://id.unity.com/account/new).
2. [Set up an Asset Store publisher profile](AssetStoreCreateAcct.md).

## Publish a UPM package to the Asset Store

To publish a UPM package to the Asset Store, you need an invitation to the closed beta. To apply to join the UPM publishing early access program, go to [UPM Publishing on the Asset Store](https://assetstore.unity.com/publishing/upm-publishing).

If accepted, you receive an email invite with a link to enroll in early access, and then must perform the following steps:

1. [Create a product draft](asset-store-upm-create-draft.md).
2. [Validate and upload your package](asset-store-upm-validate.md).
3. [Submit the package for approval](asset-store-upm-submit.md).

You can then [check the status of the package](asset-store-upm-check-status.md), and once published, you can make [further updates in the Publisher Portal](asset-store-upm-update.md).

## Additional resources

* [Apply for early access to UPM publishing (Beta)](asset-store-upm-apply.md)
* [Create a Publisher Portal product draft (Beta)](asset-store-upm-create-draft.md)
* [Managing your organization](https://docs.unity.com/en-us/cloud/organizations)
* [Asset Store packages](AssetStorePackages.md)