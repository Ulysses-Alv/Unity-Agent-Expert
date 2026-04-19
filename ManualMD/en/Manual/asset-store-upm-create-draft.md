# Create a Publisher Portal product draft (Beta)

**Important:** This is an early access workflow, and is available only to selected publishers to publish free Unity Package Manager (UPM) packages. To register your interest and apply for early access to UPM publishing, go to [UPM Publishing on the Asset Store](https://assetstore.unity.com/publishing/upm-publishing).

Before you can validate and publish your package, you must create a product draft in the Publisher Portal.

To access the Publisher Portal for UPM publishing, you must select the organization that has been granted early access as follows:

1. Open the **Asset Store**, and select your user icon.
2. Under **Switch Organization**, select your publisher organization.

You can then access the Publisher Portal in the following ways:

* Use the [direct link to the early access Publisher Portal](https://cloud.unity.com/assetstore/)
* Access the UPM publishing Publisher Portal through the Cloud Dashboard as follows:
  1. Go to the [Cloud Dashboard](https://cloud.unity.com/home/).
  2. Select the **Products** tab.
  3. In the search bar, enter Publisher Portal.
  4. Select **Launch**.

You can configure a product to handle single or multiple packages, as follows:

* **Single package**: Only one package uploaded and linked to this product draft. New versions of the package continuously share the same package technical names. Ideal for a unified tool that supports a wide array of **project settings**.
* **Multiple packages:** Various packages uploaded and linked to this product draft. The packages can be dependencies, dependents, or packages intended for different project settings or use cases. Ideal for support for multiple use cases, or different project compatibility settings, such as different Unity Editor versions.

During the early access period, you must choose this configuration upfront, but for general availability, you can start with a single package and add more later.

When planning dependencies, use multiple packages if dependencies are only needed within one product. Dependent packages automatically install their dependencies. However, if a dependency is shared across multiple products, list it as its own separate product to avoid conflicts and project breaks. For more information about package dependencies, refer to [Package dependency and resolution](upm-dependencies.md).

The following table compares the properties of each product structure:

| **Property** | **Single package product** | **Multiple package product** |
| --- | --- | --- |
| **Number of packages** | One | Two or more |
| **Compatibility** | Compatible with multiple Unity versions, render pipelines, or build targets. | Optionally, each package can be compatible with a different Unity version, **render pipeline**, or build target. |
| **Version history** | Single version history. | Separate version history for each package. |
| **Support dependency** | No | Yes |
| **Architecture** | Monolithic. All features bundled together. | Can include core and optional packages, and package dependencies. |
| **Dependencies** | Publish a dependency as a single-package product when two or more different products require this dependency. | Publish a dependency as a multi-package product when any package in one product requires this dependency. |

By default the uploaded Unity Editor version is checked for every package uploaded. Additionally, you can list all supported Unity versions that it’s compatible with (major.minor). For build target and render pipeline, setting a specific compatibility is optional.

## Create a product draft

To create a product draft:

1. Use the unique link in your confirmation email to open the Publisher Portal.
2. Go to the **Products** tab.
3. Select **Create product**.
4. Enter a product name.

### Product and package naming

When a product has only one package, the product namespace and package namespace are the same, which follows the naming convention of `com.publisher.product`. When a product has multiple packages, the product namespace follows the naming convention of `com.publisher.product` (for example `com.unity.assetstore`), and the package namespaces follow the naming convention of `com.publisher.product.package1` (for example `com.unity.assetstore.tools`).

The following table describes the information each namespace identifies and its format in a complete UPM namespace:

| **Namespace type** | **Description** | **Example** |
| --- | --- | --- |
| **Publisher namespace** | The publisher organization. This namespace uses the reverse domain name notation (reverse-DNS) naming convention. For example, the domain `example.com` has the publisher namespace `com.example`. | `com.publisher` |
| **Product namespace** | The product. This namespace isn’t visible in the product’s Asset Store listing. | `.product` |
| **Package namespace** | The packages that the product contains. Set a package’s namespace only for multi-package products. | * `product.package1` * `product.package2` |
| **[Package technical name](upm-manifestPkg.md#required)** | The package technical name combines the publisher namespace, the product namespace, and for multi-package products, the package namespace. The Package’s technical name must match the following:  * The package folder name. * The `name` field [in the package manifest file](upm-manifestPkg.md). | * **Single package**:   + `com.publisher.product` * **Multiple packages**:   + `com.publisher.product.package1`   + `com.publisher.product.package2` |

**Note:** The publisher namespace (`com.publisher`) is set during enrolment to the development membership program. If you need to change it, you must [contact Unity support](https://support.unity.com/hc/en-us/requests/new?ticket_form_id=65905).

## Additional resources

* [Validate and upload a UPM package for the Asset Store (Beta)](asset-store-upm-validate.md)
* [Submit an UPM package for approval to the Asset Store (Beta)](asset-store-upm-submit.md)