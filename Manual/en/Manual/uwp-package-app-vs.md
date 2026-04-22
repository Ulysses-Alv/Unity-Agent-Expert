# Package a UWP app in Visual Studio

Before you distribute your Universal Windows Platform (UWP) application, you need to package it in Visual Studio. For more information, refer to Microsoft documentation on [packaging a desktop or UWP app in Visual Studio](https://learn.microsoft.com/en-us/windows/msix/package/packaging-uwp-apps).

## Create an app package in Visual Studio

To create an app package in Visual Studio:

1. Open your built UWP project in Visual Studio.
2. In the **Solution Explorer**, right-click on your main project.
3. Go to **Publish** > **Create App Packages**. The **Create App Packages** wizard appears.
4. Select **Microsoft Store using a new app name** and then click **Next**.   
   **Note:** If you choose **Sideloading**, Visual Studio will not generate the app package upload file for Partner Center submissions. You can select the **Sideloading** option if you only want to create a MSIX package for non-Store distribution.
5. On the next page, sign in to the Partner Center with your developer account. If you don’t have a developer account yet, the wizard will show you how to create one.
6. Select the app name for your package from the list of apps currently registered to your account, or reserve a new name in the Partner Center.
7. In the **Select and Configure Packages** dialog, select the architectures you want to target based on the devices you want to deploy your application to.
8. In the **Generate app bundle** listbox, select **Always**.
9. Click **Create** to generate the app package.

Your app is now successfully packaged.

## Install your package on your machine

After you create your app package, use the following steps to install the package on your machine:

1. Open the folder which contains your package.
2. Right-click on the `Add-AppxPackage <yourappx>.appx` file. Choose **Run with PowerShell** and follow the prompts.
3. If you [signed your file](https://learn.microsoft.com/en-us/windows/msix/package/sign-app-package-using-signtool), the file will now install on your machine.  
   **Note:** If you’re re-installing an .appx file, you must uninstall the file that was previously installed by right-clicking the file icon, and then clicking **Uninstall**.