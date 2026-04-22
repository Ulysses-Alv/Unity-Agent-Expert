# Install Unity Accelerator with the installer

To install Unity **Accelerator** through its installer:

1. Go to [Unity Accelerator downloads](accelerator-download.md).
2. Select the correct installer for your operating system.

   Refer to [Verify the Unity Accelerator version](accelerator-verify-version.md) to verify the installation on a Linux machine.
3. Open the installer, and follow the steps to complete installation.

On the final screen, the installer displays the server address of the installed accelerator, which you can use to [configure the accelerator in the Unity Editor](accelerator-configure.md).

## Run the installer on the command line

To run the installer in unattended mode from the command line, use the `--mode unattended` argument. Execute the installer with an argument of `--help` to display the options available.

| **Command line argument** | **Description** |
| --- | --- |
| `--install-dir` | Sets the directory where the binaries are copied. |
| `--storage-dir` | Sets the directory for the Accelerator to store files and configurations. |
| `--admin-username` | Admin username required to log into the dashboard. |
| `--admin-password` | Admin password required to log into the dashboard. |
| `--tls-config` | Chooses the security configuration method (allows: selfsigned, custom, none; default: selfsigned). |
| `--tls-selfsigned-host` | Hostname to use for the self-signed certificate. |
| `--tls-custom-host` | Hostname to use for the custom certificate. |
| `--tls-custom-certificate` | Path to a PEM encoded certificate. |
| `--tls-custom-private-key` | Path to a PEM encoded private key. |
| `--action-on-existing-install-found` | Chooses how to handle existing configuration (allows: reconfigure, replace; default: reconfigure). |
| `--linux-service-user` | User account to run the accelerator service on Linux (default: root). The user needs to have permissions to access the involved directories. |

Examples:

```
$ unity-accelerator-installer.run --mode unattended --admin-username mark --admin-password xxxxx

$ unity-accelerator-installer.run --mode unattended --admin-username mark --admin-password xxxxx --install-dir /opt/Unity/accelerator/ --storage-dir /var/lib/unity-accelerator --tls-config none
```

You can also run the installer in interactive text mode, available only for Linux and macOS.

```
$ sudo ./unity-accelerator-installer.run --mode text
```

On macOS, you need to mount the disk image (.dmg) and run the binary located in the installer app’s directory located at `Contents/macOS/installbuilder.sh`.

After you’ve installed the accelerator, follow the information in [Configure the Unity Editor to use an accelerator](accelerator-configure.md).

## Additional resources

* [Unity Accelerator requirements](accelerator-requirements.md)
* [Install Unity Accelerator with Docker Hub](accelerator-install-docker.md)
* [Verify the Unity Accelerator version](accelerator-verify-version.md)
* [Configure an accelerator in the Editor](accelerator-configure.md)