# Submit a license request from a command line and browser (Windows)

Submit a license request from the command line to manually activate your license if you’re unable to use the [other activation methods](LicenseActivationMethods.md).

**Note**: The manual activation method doesn’t work with floating license subscriptions or Unity Personal. To activate a license for Unity Personal, log in to the Unity Hub. To return a Personal license, log out of the Unity Hub.

## Before you begin

* See [Manual license activation](ManualActivationGuide.md) to make sure you understand the scenario for using this procedure, its limitations, and its internet connectivity requirements.
* Make sure you know the path where you installed the Unity Editor. Use the Unity Hub to determine the path. Open the Unity Hub and select **Installs** from the side menu. The list shows the path for each installed Editor. For more information, see [Locate the Editor](https://docs.unity3d.com/hub/manual/AddEditor.html#locate-the-editor-program-file).

The default installation path for the Unity Editor is `"C:\Program Files\Unity\Hub\Editor\<version>\Editor\Unity.exe"`.

## Procedure

To manually activate your Unity license, follow these steps:

1. [Create a license request file](#alf-commandline) (`.alf`) from the command line. You must do this step from the computer where you installed Unity.
2. Use that `.alf` file to [generate a Unity license file](#request-ulf) (`.ulf`) from Unity. You must do this step from any computer that has internet access.
3. Use that `.ulf` file to [activate your license from the command line](#ulf-commandline). You must do this step from the computer where you installed Unity.

### 1) Create a license request file from the command line

**Important**: You must run this command from the computer where you installed Unity, but the computer doesn’t need internet access for this step to work.

1. Open a Command Prompt as an administrator. If you lack administrator privilege, make sure you run this procedure’s command from within your user profile folder.
2. Make sure you’re aware which directory you’re in. The directory you run the command from is the same directory where the `.alf` file is output.
3. Enter the following command, replacing `<editor-installation-location>` with the actual full path to `Unity.exe`:

   ```
   "<editor-installation-location>" -batchmode -createManualActivationFile -logfile
   ```

   **Note:** This command doesn’t return output to the Command Prompt. When the command completes, it returns control to the command prompt.
4. When the command completes, check the directory where you ran the command to make sure it created an activation license file, such as `Unity_v2022.2.0b4.alf`. If you experience issues, see [Troubleshooting](#troubleshooting).

Now you are ready to [generate a Unity license file](#request-ulf).

##### Example

Assuming you installed a `2022.2.0b4` Editor at `"C:\Program Files\Unity\Hub\Editor\2022.2.0b4\Editor\Unity.exe"`, the command for this step is:

```
"C:\Program Files\Unity\Hub\Editor\2022.2.0b4\Editor\Unity.exe" -batchmode -createManualActivationFile -logfile
```

### 2) Generate a Unity license file

The next step is to use the license activation file (`.alf`) that you just created to request a Unity license file (`.ulf`).

**Important**: You can generate a Unity license file from any computer; it doesn’t necessarily need to be the same computer where you installed Unity. However, the computer you use in this step must have an internet connection.

1. Go to [id.unity.com](https://id.unity.com/) and make sure you’re signed in to your Unity ID. You need a valid login to request a Unity license file.
2. Use the same browser session to go to [license.unity3d.com/manual](https://license.unity3d.com/manual). The **Manual activation** page appears.

   **Note:** If you experience issues accessing this page, try pasting `https://license.unity3d.com/manual` into your browser’s address bar.

   ![Upload license request page.](../uploads/Main/license-request-ulf.png)

   Upload license request page.
3. Select the **Browse** button and use your computer’s file browser to select the license activation (`.alf`) file from the [create a license request file](#alf-commandline) step.

   If the upload was successful and Unity detected that the file you chose was a license activation file, the filename appears in the text box with a green check mark. If you see a red **X** instead, try again.
4. Select the **Next** button. The **Activate your license** page appears.

   ![Activate your license page.](../uploads/Main/license-request-activate.png)

   Activate your license page.
5. Enter the serial key you received in an email when you purchased a single license. You need to enter the key in the exact format that it appears in the email. For example, `PS/SC/E3-XXXX-XXXX-XXXX-XXXX`.
6. Select the **Next** button. The **Download license file** page appears.
7. Select the **Download license file** button. If prompted, allow downloads for the [license.unity3d.com/manual](https://license.unity3d.com/manual) page.
8. Open your browser’s downloads location and confirm the new Unity license file, which has a `.ulf` extension. For example, `Unity_v2017.x.ulf`.

Now you are ready to [activate your Unity license from the command line](#ulf-commandline).

### 3) Activate your license from the command line

**Important**: You must run this command from the computer where you installed Unity, but the computer doesn’t need internet access for this step to work.

1. Make sure you know the following paths:

   * `<editor-installation-location>` is the actual full path to `Unity.exe`.
   * `<yourUlfFile>` is the full path of the `.ulf` file you generated in the [generate a Unity license file](#request-ulf) step.
2. Open a Command Prompt as an administrator. If you lack administrator privilege, make sure you run this procedure’s command from within your user profile folder.
3. Run the following command, replacing `<editor-installation-location>` and `<yourUlfFile>` with the actual full paths:

   ```
   "<editor-installation-location>" -batchmode -manualLicenseFile <yourUlfFile> -logfile
   ```

   **Note:** This command doesn’t return output to the Command Prompt. When the command completes, it returns control to the command prompt.
4. When the command completes, you can view your active license in the Hub by selecting **Licenses** from the side menu. If you experience issues, refer to [Troubleshooting](#troubleshooting).

##### Example

This example assumes that you:

* Installed a `2022.2.0b4` Editor at `"C:\Program Files\Unity\Hub\Editor\2022.2.0b4\Editor\Unity.exe"`
* Generated a `.ulf` file and stored it at `"C:\Users\myAccount\Downloads\Unity_v2017.x.ulf"`

```
"C:\Program Files\Unity\Hub\Editor\2022.2.0b4\Editor\Unity.exe" -batchmode -manualLicenseFile "C:\Users\myAccount\Downloads\Unity_v2017.x.ulf" -logfile
```

## Troubleshooting

If the command line operations don’t yield the expected results, view the `Editor.log` file. To find the the location of this file, refer to [Log files reference](log-files.md).

## Additional resources

* [License activation methods](LicenseActivationMethods.md)
* [Activate a license for a Unity Personal plan](https://docs.unity3d.com/hub/manual/ManageLicense.html#personal) (Unity Hub documentation)
* [License troubleshooting](ActivationFAQ.md)