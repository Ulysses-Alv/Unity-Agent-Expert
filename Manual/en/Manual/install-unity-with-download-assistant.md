# Install Unity using Download Assistant

Unity Download Assistant is a small-sized executable file provided as `.exe` for Windows and `.dmg` for macOS. The Download Assistant guides you through downloading and installing Unity Editor and other components together.

To install Unity using the Download Assistant, follow these steps:

1. Access the [Unity download archive](https://unity.com/releases/editor/archive) website to select the release versions of the Unity Editor.
2. Use the filter controls to list the versions of Unity you’re interested in.
3. In the **Downloads** column, select **See all** for the version of Unity you want to install.  
   This opens a new page that displays the **Manual installs** and **Component installers** sections for the selected Unity Editor version.
4. In the **Manual installs** section, expand **Other installs**.
5. Select **Download Assistant** for the operating system of your computer.  
   This starts the download process of the **Download Assistant** executable file. If prompted, save the file, and note the download location.
6. Use your file manager application to go to the location where you downloaded the Download Assistant (`.exe` for Windows and `.dmg` for macOS) and open it.

   **Note**: On **macOS**, the `.dmg` file opens a new Finder window. From that window, open `Unity Download Assistant.app`.
7. Follow the **Introduction** and **License** prompts, then select the components you want to install.  
   The **Download Assistant** will install all the components you select on the current computer. If you want to install on other computers later, you have an option to save the installer files to a folder.
8. To save the installer files on macOS, select **Advanced** on the **Destination Select** panel (**Select a Destination**).
9. Without closing the **Download Assistant**, go to your file manager application, and create a folder to store the installer files you’ll download in the next step.
10. Return to the Download Assistant and set a location for storing installer files.

    **Important**: Don’t save the installation files at the default location. Selecting the default location might result in incomplete file sets, or complete removal of the files you’ll need to run the setup on other computers.

    * On macOS, change **Default** to **Custom**. Select the folder icon and choose the folder you explicitly created in the previous step.
    ![Change the default location to Custom for macOS](../uploads/Main/download-assistant-location-mac.png)

    Change the default location to Custom for macOS

    * On Windows, select **Download to** and choose the folder you explicitly created in the previous step.
    ![Change the default location to Download to for Windows](../uploads/Main/download-assistant-location-windows.png)

    Change the default location to Download to for Windows
11. Proceed through the remaining steps to complete the installation process. When the process completes, the installer deploys the Unity Editor to the following default locations:

    * On macOS, `Applications/Unity`.
    * On Windows, `C:\Program Files\Unity <version>`.

The Download Assistant also places the setup files you need to install Unity on offline computers. To confirm that the download of these files was successful, do the following:

1. Open your file manager application.
2. Go to the folder you specified when selecting the location for storing installation files during the Download Assistant setup.
3. Confirm that an `install.sh` or `install.bat` file is inside that folder.

   If this folder is empty, or missing an `install.sh` or `install.bat` file, repeat the main procedure. Make sure you specify an explicit location for installation files, rather than accepting the default location.
4. Check the contents of the `install.sh` or `install.bat` file to make sure they resemble the following example:

   ![Sample contents of an install.bat file (Windows)](../uploads/Main/install-bat.png)

   Sample contents of an install.bat file (Windows)

## Additional resources

* [Install Unity from the command line](InstallingUnity.md)
* [Install Unity on offline computers](DeployingUnityOffline.md)