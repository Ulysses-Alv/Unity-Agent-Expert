# Install Unity on offline computers

The Unity Download Assistant supports offline deployment.

## Before you begin

Make sure you’ve run the Download Assistant on an online computer. For more information, refer to [Install Unity using Download Assistant](install-unity-with-download-assistant.md). This process generates the installer files and the script that the offline computers require for the installation process.

## Unity installation on offline computers

Follow these steps to run the script that the Download Assistant generates to install Unity Editor and components on offline computers.

1. Transfer the folder that includes all the downloaded installation files and the `install.sh` or `install.bat` file to the offline computer.
2. On the offline computer, run the `install.sh` or `install.bat` file:

   * On macOS:
     1. Open Terminal.
     2. Go to the folder that contains the `install.sh` file.
     3. Run `sudo ./install.sh`. This command prompts you for your password, which doesn’t appear on the screen as you type it.
   * On Windows:
     1. In the **Start** menu, search for `cmd.exe`.
     2. Right-click **Command Prompt**.
     3. Select **Run as administrator**.
     4. Go to the folder that contains the `install.bat` file.
     5. Run `install.bat`.
3. Confirm that the script installed the Unity Editor on this computer.

   * On macOS, the default location is `Applications/Unity`.
   * On Windows, the default location is `C:\Program Files\Unity <version>`.
4. Repeat these steps on other offline computers if needed.

## Activate a license on the offline computer

After you install the Unity Editor on the offline computer, its user needs to activate their license on that computer by using the [manual activation method](ManualActivationGuide.md).

## Additional resources

* [Install Unity using Download Assistant](install-unity-with-download-assistant.md)
* [Install Unity from the command line](InstallingUnity.md)