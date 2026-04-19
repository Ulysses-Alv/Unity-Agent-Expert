# Install Unity using installer files

The installer files are executable programs that you can use to install the Unity Editor and other components. The components provide target platform build support for the installed Unity Editor version. You must download separate installer files for the Unity Editor and other components.

The installation process involves the following steps:

1. **Download the Unity Editor installer file**: The installer files for Windows and macOS are available in `.exe` and `.pkg` formats respectively. On Linux, the installer file is an archive file in `.tar.xz` format.
2. **Install the Unity Editor**: The installation process varies per operating system, with Windows and macOS using a similar process.
3. **Download and install components**: If you want to enable build support for any additional target platforms, you can download and install components for the installed Editor version.

## Download the Unity Editor installer file

To download the Unity Editor for the operating system of your development computer, follow these steps:

1. Access the [Unity download archive](https://unity.com/releases/editor/archive) website to select the release versions of the Unity Editor.  
   **Note**: To get early access to features in the upcoming release, check the [Unity Beta Program](https://unity.com/releases/editor/beta).
2. Use the filter controls to list the versions of Unity you’re interested in.
3. In the **Downloads** column, select **See all** for the version of Unity you want to install.  
   This opens a new page that displays the **Manual installs** and **Component installers** sections for the selected Unity Editor version.
4. In the **Manual installs** > **Operating Systems** section, select the Unity Editor installer file for the operating system of your development computer.

The download process of the installer file starts. If prompted, save the file, and note the download location.

## Install the Unity Editor

After you download the installer files for Windows and macOS, or the archive file for Linux, follow these steps to install the Unity Editor.

### Installation on Windows and macOS

1. Use your file manager application to go to the location where you downloaded the installer file.

   **Note**: You can use the downloaded installer file to install the Unity Editor on other computers using the command line.
2. Run the installer file and follow the **Introduction** and **License** prompts.
3. Choose the folder location where you want to install the Unity Editor. You can change the default location. Note the selected location.
4. Proceed through the remaining steps to complete the installation process. When the process completes, the installer deploys the Unity Editor to the following default locations:

   * On macOS, `Applications/Unity`.
   * On Windows, `C:\Program Files\Unity <version>`.

### Installation on Linux

1. Use your file manager application to go to the location where you downloaded the archive file.
2. Right-click the downloaded archive file and select **Extract** from the drop-down menu.
3. Create a directory named `bin` within the `~/.local` directory, if not already present.
4. If installing for a single user, move the directory containing the extracted Linux Editor files to the `~/.local/bin/` directory.

   **Note**: You can also use the command line to move the files: `mv [Extracted-Linux-Editor-Directory-Name] ~/.local/bin`.
5. Add the Linux Editor to the **Installs** section of Unity Hub.
   1. Sign in to Unity Hub.
   2. Select **Installs** on the side panel of the Hub.
   3. Select the **Locate** button.
   4. Navigate to the directory containing the extracted Linux Editor files: `~/.local/bin/[Extracted-Linux-Editor-Directory-Name]`
   5. In this directory, select the executable named `Unity`.
   6. Select the **Select Editor** button.

The Linux Unity Editor now appears in the **Installs** section of Unity Hub.

## Download and install components

After you install the Unity Editor, you can install components for that Editor version.

### Component installation on Windows and macOS

To download and install components for the installed Windows and macOS Editor, follow these steps:

1. Return to the [Unity download archive](https://unity.com/releases/editor/archive) website and navigate to the page for the Unity Editor version you installed.
2. In the **Component installers** section, expand the relevant operating system section and select the required component for platform build support.
   The download process of the installer file starts.
3. Run the installer file and proceed through the installation steps similar to those for the Editor installation. Select the root folder of the Unity Editor installation you want to add the components to.

This installs the selected component for the corresponding Windows or macOS Unity Editor version.

### Component installation on Linux

To download and install components for the installed Linux Editor, follow these steps:

1. Return to the [Unity download archive](https://unity.com/releases/editor/archive) website and navigate to the page for the Unity Editor version you installed.
2. In the **Component installers** section, expand the section relevant to the Linux operating system and select the required component.
   The download process of the component installer archive file starts.
3. Use your file manager application to go to the location where you downloaded the installer archive file.
4. Extract the downloaded archive file.
5. Navigate to the directory containing the extracted files.
6. In this directory, go to `Editor/Data/PlaybackEngines`.
7. Copy the contents of the `PlaybackEngines` directory to the `~/.local/bin/[Unity-Editor-Version]/Editor/Data/PlaybackEngines` directory.

This installs the selected component for the corresponding Linux Unity Editor version.

**Note**: Some platforms might require you to install additional dependencies alongside the platform build support component. Refer to [Platform development](PlatformSpecific.md) to understand the dependency requirement of the supported platforms.

## Additional resources

* [System requirements](system-requirements.md)
* [Install Unity from the command line](InstallingUnity.md)
* [Install Unity using Download Assistant](install-unity-with-download-assistant.md)
* [Install Unity on offline computers](DeployingUnityOffline.md)