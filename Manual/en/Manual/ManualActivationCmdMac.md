# Submit a license request from a command line and browser (macOS, Linux)

Submit a license request from the command line to manually activate your license if you’re unable to use the [other activation methods](LicenseActivationMethods.md).

**Note**: The manual activation method doesn’t work with floating license subscriptions or Unity Personal. To activate a license for Unity Personal, log in to the Unity Hub. To return a Personal license, log out of the Unity Hub.

## Before you begin

* See [Manual license activation](ManualActivationGuide.md) to make sure you understand the scenario for using this procedure, its limitations, and its internet connectivity requirements.
* Make sure you know the path where you installed the Unity Editor. Use the Unity Hub to determine the path. Open the Unity Hub and select **Installs** from the side menu. The list shows the path for each installed Editor. For more information, see [Locate the Editor](https://docs.unity3d.com/hub/manual/AddEditor.html#locate-the-editor-program-file).

The default installation paths for the Unity Editor are:

* macOS: `/Applications/Unity/Hub/Editor/<version>/Unity.app`
* Linux: `/Applications/Unity/Hub/Editor/<version>/Unity.app`

## Procedure

To manually activate your Unity license, follow these steps:

1. [Create a license request file](#alf-commandline) (`.alf`) from the command line. You must do this step from the computer where you installed Unity.
2. Use that `.alf` file to [generate a Unity license file](#request-ulf) (`.ulf`) from Unity. You must do this step from any computer that has internet access.
3. Use that `.ulf` file to [activate your license from the command line](#ulf-commandline). You must do this step from the computer where you installed Unity.

### 1) Create a license request file from the command line

**Important**: You must run this command from the computer where you installed Unity, but the computer doesn’t need internet access for this step to work.

1. Make sure you know the installation location of your Unity Editor (ending with `Unity.app`), described earlier. On macOS, you’ll add `/Contents/MacOS/Unity` to this path for this procedure.

   **Note**: On macOS, although you don’t need Finder for this step, this latter subdirectory structure isn’t visible in Finder, unless you right-click the Unity application (`Unity.app`) and select **Show Package Contents**.

   For example, for a `2022.2.0b4` Editor, the complete path to the `Unity` command on macOS is:

   ```
   /Applications/Unity/Hub/Editor/2022.2.0b4/Unity.app/Contents/MacOS/Unity
   ```
2. Open a Terminal session (macOS) or the command line (Linux), and make sure you’re aware which directory you’re in. The directory you run the command from is the same directory where the `.alf` file is output. You can use the `pwd` command to determine your directory.
3. Run the following command, replacing `<editor-command-location>` with the complete path to the `Unity` command from Step 1:

   ```
   <editor-command-location> -batchmode -createManualActivationFile -logfile
   ```

   **Note:** This command doesn’t return output to the Terminal or command line window. When the command completes, it returns control to the command prompt.
4. When the command completes, check the directory where you ran the command to make sure it created an activation license file, such as `Unity_v2022.2.0b4.alf`. If you experience issues, see [Troubleshooting](#troubleshooting).

Now you are ready to [generate a Unity license file](#request-ulf).

##### Example (macOS)

Assuming you installed a `2022.2.0b4` Editor at `/Applications/Unity/Hub/Editor/2022.2.0b4/Unity.app`, the command for this step is:

```
/Applications/Unity/Hub/Editor/2022.2.0b4/Unity.app/Contents/MacOS/Unity -batchmode -createManualActivationFile -logfile
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

   * `<editor-command-location>` is the same value as the path described in the [create a license request file](#alf-commandline) step.
   * `<yourUlfFile>` is the full path of the `.ulf` file you generated in the [generate a Unity license file](#request-ulf) step.
2. Open a Terminal session (macOS) or the command line (Linux), and run the following command, replacing `<editor-command-location>` and `<yourUlfFile>` with the actual full paths:

   ```
   <editor-command-location> -batchmode -manualLicenseFile <yourUlfFile> -logfile
   ```

   **Note:** This command doesn’t return output to the Terminal or command line window. When the command completes, it returns control to the command prompt.
3. When the command completes, you can view your active license in the Hub by selecting **Licenses** from the side menu. If you experience issues, refer to [Troubleshooting](#troubleshooting).

##### Example (macOS)

This example assumes that you:

* Installed a `2022.2.0b4` Editor at `/Applications/Unity/Hub/Editor/2022.2.0b4/Unity.app`
* Generated a `.ulf` file and stored it at `/Users/myAccount/Downloads/Unity_v2017.x.ulf`

```
/Applications/Unity/Hub/Editor/2022.2.0b4/Unity.app/Contents/MacOS/Unity -batchmode -manualLicenseFile /Users/myAccount/Downloads/Unity_v2017.x.ulf -logfile
```

## Troubleshooting

If the command line operations don’t yield the expected results, view the `Editor.log` file. To find the the location of this file, refer to [Log files reference](log-files.md).

## Additional resources

* [License activation methods](LicenseActivationMethods.md)
* [Activate a license for a Unity Personal plan](https://docs.unity3d.com/hub/manual/ManageLicense.html#personal) (Unity Hub documentation)
* [License troubleshooting](ActivationFAQ.md)