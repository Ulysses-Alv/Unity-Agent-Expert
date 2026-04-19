# Validate and upload assets to your package

You can add assets to any package draft on the Publisher Portal or from the Unity Editor. To upload assets, use the [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745) package.

**Note**: To organize your package’s assets in an effective way, move all the assets to a top-level folder that has the same name as your package.

## Install the Asset Store Publishing Tools package

Use the **Asset Store** Publishing Tools package to validate, upload, and preview your assets. To add it to your project:

1. Go to the [Asset Store Publishing Tools](https://assetstore.unity.com/packages/package/5368745).
2. Select **Open in Unity**.
3. In the Unity Editor go to **Window** > **Package Management** > **Package Manager**.
4. Select **My Assets**.
5. Search for and select the Asset Store Publishing Tools package.
6. Select **Download**.
7. Select **Import [Version number] to project**.
8. In the **Import Unity Package** window that appears, select **Import**.

## Validate assets

The Asset Store Publishing Tools validator checks your assets to make sure they follow the [Submission Guidelines](https://unity3d.com/asset-store/sell-assets/submission-guidelines). Use it to identify errors that you can fix before you submit a package for approval.

To validate your assets:

1. Open the Unity Project that contains the content you want to submit to the Asset Store.
2. In the Main Menu, select **Tools** > **Asset Store** > **Validator**.
3. In the Asset Store Validator window select **Add a path**.
4. Select the folder that contains the assets you want to submit to the Asset Store.
5. Select **Validate**.
6. If any warnings or errors appear select them to learn more about how to fix each one.

**Note**: If your package contains audio or video files you must upload samples of your media assets in the Media section of your package draft.

## Upload assets to the Publisher Portal

Use the Asset Store Uploader to create and upload a `.unitypackage` file that includes your assets to the Publisher Portal. To do this:

1. In the Main Menu, select **Tools** > **Asset Store** > **Uploader**.
2. Enter your Unity ID login details.
3. Select the package draft to add these assets to.
4. Select **Export and Upload**.
5. Open the Publisher Portal and select the package draft you uploaded the assets to.
6. Go to the **Package upload** section. If your package doesn’t appear, select **Refresh Package**.

## Additional resources

* [Submit a package for approval](AssetStoreSubmit.md)
* [Check the status of an Asset Store package](AssetStoreStatus.md)
* [The Asset Store publishing workflow](asset-store-workflow.md)
* [Asset Store packages](AssetStorePackages.md)
* [Asset Store Submission Guidelines](https://unity3d.com/asset-store/sell-assets/submission-guidelines)
* [Asset Store Provider Agreement](https://unity3d.com/legal/as_provider)