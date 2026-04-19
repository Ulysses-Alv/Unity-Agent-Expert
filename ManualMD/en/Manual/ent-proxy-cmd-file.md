# Create a command file to set environment variables and open applications

When you define environment variables from a command line, the values don’t persist across sessions. A process can use those environment variables if it’s launched from the same session, but after you end that session, you must define the values again.

Administrators can set environment variables at the operating system level. If that’s not an option, you can create a command file to set environment variables and launch a Unity application, such as the Hub, from the same session. This executable file becomes a reusable method for opening the Unity Hub with the environment variables required by your environment.

**Important**: After the command file is created, be sure to always launch the Unity Hub using that file, instead of standard methods, such as the **Start** menu (Windows) or the **Applications** folder (macOS).

## Before you begin

Decide which environment variables you need to set, and include only the required environment variables:

* If your environment doesn’t support automatic proxy configuration, set the `HTTP_PROXY` and `HTTPS_PROXY` environment variables to identify your web proxy. For more information, refer to [Use environment variables to identify your web proxy](ent-proxy-env-vars.md).
* If your environment uses a Unity Licensing Server for floating licenses, set the `NO_PROXY` environment variable.
* If your environment uses a web proxy with SSL inspection, set the `NODE_EXTRA_CA_CERTS` environment variable. For more information, refer to [Trusting the web proxy security certificate](ent-proxy-cert-trust.md).

If you’re setting the `HTTP_PROXY` and `HTTPS_PROXY` environment variables, make sure you use the proper value to [include](ent-proxy-env-vars.md#include-auth) or [exclude](ent-proxy-env-vars.md#exclude-auth) authentication information.

## Create a command file (Windows)

These instructions create an executable file named `launchUnityHub.cmd`.

1. Close the Unity Hub, if it’s running.
2. Open a text editor such as Notepad.
3. Enter the following text, adjusting as necessary:

   * Include only the required environment variables, and set their values properly. Refer to [Before you begin](#prereqs).
   * Set the Hub path to match the location of your Unity Hub program file.

   ```
   @echo off
   REM *** NOTE: Add the next 3 lines only if you’re not using Automatic Proxy Configuration
   set HTTP_PROXY=http://<username>:<password>@<proxy_name_or_IP_address>:8080
   set HTTPS_PROXY=http://<username>:<password>@<proxy_name_or_IP_address>:8080
   set NO_PROXY=<licensing_server_name_or_IP_address>
   REM *** NOTE: Add the following line only if your web proxy uses SSL inspection
   set NODE_EXTRA_CA_CERTS=<path_to_pem_file>
   start "" "C:\Program Files\Unity Hub\Unity Hub.exe"
   ```

   **Note:** If there are spaces in the path, you must use double quotes around the path to the program.
4. Save the file to a location where you can find it (such as the `Desktop`), and make sure the file has the `.cmd` extension (for example, `launchUnityHub.cmd`), not `launchUnityHub.txt` or `launchUnityHub.cmd.txt`.
5. Double-click `launchUnityHub.cmd` to launch the Unity Hub with the required environment variables intact.

The Unity Hub passes these environment variables to any process it spawns, such as the Unity Editor.

## Create a command file (macOS and Linux)

These instructions create an executable file named `launchUnityHub.command`.

1. Close the Unity Hub, if it’s running.
2. Open a Terminal window.
3. Run the following command, adjusting as necessary:

   * Include only the required environment variables, and set their values properly. Refer to [Before you begin](#prereqs).
   * Set the Hub path to match the location of your Unity Hub application.

   ```
   echo '#!/usr/bin/env bash
   # *** NOTE: Add the next 3 lines only if you’re not using Automatic Proxy Configuration
   export HTTP_PROXY=http://<username>:<password>@<proxy_name_or_IP_address>:8080
   export HTTPS_PROXY=http://<username>:<password>@<proxy_name_or_IP_address>:8080
   export NO_PROXY=<licensing_server_name_or_IP_address>
   # *** NOTE: Add the following line only if your web proxy uses SSL inspection
   export NODE_EXTRA_CA_CERTS=<path_to_pem_file>
   nohup "/Applications/Unity Hub.app/Contents/MacOS/Unity Hub" &>/dev/null &' > launchUnityHub.command
   ```

   **Note:** If there are spaces in the path, you must use double quotes around the path to the application.
4. Run the following command to make `launchUnityHub.command` executable:

   ```
   chmod +x launchUnityHub.command
   ```
5. Move the `launchUnityHub.command` file to a convenient location (for example, the `Desktop`), if you prefer.
6. Double-click `launchUnityHub.command` to launch the Unity Hub with the required environment variables intact.

The Unity Hub passes these environment variables to any process it spawns, such as the Unity Editor.

## Additional resources

* [Using Unity through web proxies](ent-proxy-autoconfig.md)
* [Define exceptions on your web proxy](ent-proxy-exception-list.md)
* [Unity Licensing Server](https://docs.unity3d.com/licensing/manual/index.html)