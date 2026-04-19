# Unity Accelerator downloads

Welcome to the Unity **Accelerator** downloads page. Here you can find the latest release and previous versions, along with their release notes and installers for all supported platforms.

---

## Current version — **v1.5.137**

**Release date:** March 10, 2026

| Platform | Installer |
| --- | --- |
| **Windows** | [unity-accelerator-v1.5.137+g4eddaf6a-windows-x64-installer.exe](https://download-accelerator.unity3d.com/v1.5.137+g4eddaf6a/windows/unity-accelerator-v1.5.137+g4eddaf6a-windows-x64-installer.exe) |
| **macOS** | [unity-accelerator-v1.5.137+g4eddaf6a-osx-installer.dmg](https://download-accelerator.unity3d.com/v1.5.137+g4eddaf6a/osx/unity-accelerator-v1.5.137+g4eddaf6a-osx-installer.dmg) |
| **Linux** | [unity-accelerator-v1.5.137+g4eddaf6a-linux-x64-installer.run](https://download-accelerator.unity3d.com/v1.5.137+g4eddaf6a/linux/unity-accelerator-v1.5.137+g4eddaf6a-linux-x64-installer.run) |
| **Linux Signature** | [unity-accelerator-v1.5.137+g4eddaf6a-linux-x64-installer.run.sig](https://download-accelerator.unity3d.com/v1.5.137+g4eddaf6a/linux/unity-accelerator-v1.5.137+g4eddaf6a-linux-x64-installer.run.sig) |

### Release notes

#### Installer: Fixed upgrade failure when previous installation is tampered or partially deleted

The installer could fail during upgrade if a previous Unity Accelerator installation was partially deleted. For example, if the `unity-accelerator.cfg` configuration file was manually removed while the system service registration remained in place.

Now, the scenario is detected, and the installation proceeds as a fresh installation instead of failing. A warning message is added to the installation log file.
This ensures that users with corrupted or partially uninstalled environments can successfully reinstall or upgrade without manual cleanup.

#### Fixed startup failure after upgrade due to leftover `cachedbpendingprevious` directory

An upgrade to the Unity Accelerator could result in a failure during the removal of the `cachedbpending` directory. If a previous stop left a dangling `cachedbpendingprevious` directory in the storage, subsequent restarts would fail with an error, so the Editor wouldn’t connect and the dashboard wouldn’t load.

The application now detects this condition. The service starts normally after upgrades or restarts and doesn’t require manual cleanup.

#### macOS: Notarize the accelerator binary to reduce Gatekeeper warnings

On macOS, the standalone `unity-accelerator` binary was code-signed, but not notarized. Because Apple had no record of the binary in their notarization database, macOS Gatekeeper displayed a warning dialog stating the software might be malware.

Note: macOS Gatekeeper might still show a warning. For example, if the binary is downloaded using a method that doesn’t set the quarantine attribute, or the user’s system can’t reach Apple’s notarization servers.
However, notarizing the binary significantly reduces the likelihood of these warnings for typical download scenarios.

## Older versions

### **v1.5.99**

**Release date:** December 17, 2025

| Platform | Installer |
| --- | --- |
| **Windows** | [unity-accelerator-v1.5.99+g262ce04b-windows-x64-installer.exe](https://download-accelerator.unity3d.com/v1.5.99+g262ce04b/windows/unity-accelerator-v1.5.99+g262ce04b-windows-x64-installer.exe) |
| **macOS** | [unity-accelerator-v1.5.99+g262ce04b-osx-installer.dmg](https://download-accelerator.unity3d.com/v1.5.99+g262ce04b/osx/unity-accelerator-v1.5.99+g262ce04b-osx-installer.dmg) |
| **Linux** | [unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run](https://download-accelerator.unity3d.com/v1.5.99+g262ce04b/linux/unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run) |
| **Linux Signature** | [unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run.sig](https://download-accelerator.unity3d.com/v1.5.99+g262ce04b/linux/unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run.sig) |

In this release, the Accelerator is cleaner and easier to install and configure. Here’s a list of improvements and fixes in this update:

#### Installation and cleanup

* Installer size reduction: Dramatically reduced the installer size for each platform.
  + Windows: from 143MB → 19MB
  + macOS: from 319MB → 16MB
  + Linux: from 85MB → 19MB
* Unified installer for all needs: A single installer now works for both GUI and headless servers.
* Cleaner uninstall (Windows): The uninstall process is now thorough, leaving no residual services, binaries, cache folders or apps entries.
* Reliable installation protection: Safeguards prevent multiple automatic installations running simultaneously.
* Enhanced installation process: More reliable upgrades and reinstalls — no more stalling due to previous configurations.
* Proactive admin privileges request: Installer now immediately requests admin rights to avoid mid-install failures.
* No more JavaScript errors: All underlying JavaScript errors during installation have been eliminated.
* Consistent SSL configuration: HTTPS settings are now correctly applied, defaulting to ports **443/10443**.
* Simplified installer screens: Streamlined the installation experience by removing the cache selection and Collab setup screens, as caching is now automatically managed.
* Agent cleanup complete: Unsupported features like Collab, auto-upgrade, and legacy mode have been removed for a leaner, more focused agent.
* Install from external disk (macOS): Installer can run from an external disk.
* Uninstallation improvements: Enhanced the process to handle cases where the config folder and storage folder differ.
* App entries cleanup: Ensured proper clean-up of installed app entries during uninstallation or upgrade.
* Unattended mode added: All installers now support unattended mode installation.

  ```
  unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run --mode unattended --admin-username admin --admin-password pwd --tls-config none
  ```
* Text mode added: Linux and macOS installers now support text mode installation.

  ```
  unity-accelerator-v1.5.99+g262ce04b-linux-x64-installer.run --mode text
  ```

#### Dashboard improvements

* Secure dashboard access: After a re-installation, the dashboard now enforces the use of new credentials.
* Accurate disk usage dialog: The “Disk Space Usage” dialog now displays and saves correct information.
* Improved certificate config: Certificate configuration works as intended.
* Clearer metrics: Removed misleading dashboard metrics (for example, Time Saved by Accelerator).
* Editable protobuf port: The protobuf port (Import Pipeline Port) is now editable from the dashboard.
* Removed restart service option from dashboard: Configuration changes now require a manual service restart. The previous restart option was removed because it caused issues with certificate updates and unpredictable stalls.

### **v1.0.941**

This is the last officially published release before the installer revamp.

**Release date:** October 4, 2021

| Platform | Installer |
| --- | --- |
| **Windows** | [unity-accelerator-app-v1.0.941+g6b39b61.exe](https://download-accelerator.unity3d.com/unity-accelerator-app-v1.0.941+g6b39b61.exe) |
| **macOS** | [unity-accelerator-app-v1.0.941+g6b39b61.dmg](https://download-accelerator.unity3d.com/latest-release/mac/unity-accelerator-app-v1.0.941+g6b39b61.dmg) |
| **Linux** | [unity-accelerator-app-v1.0.941+g6b39b61.AppImage](https://download-accelerator.unity3d.com/latest-release/linux/unity-accelerator-app-v1.0.941+g6b39b61.AppImage) |