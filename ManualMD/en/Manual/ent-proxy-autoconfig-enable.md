# Enable automatic proxy configuration

If your environment is compatible with Unity’s automatic proxy configuration feature, client computers must be configured to use it.

Review the solutions at [Using Unity through web proxies](ent-proxy-autoconfig.md) to check if your environment supports Unity’s automatic proxy configuration feature.

Whether you enable Unity’s automatic proxy configuration feature manually or by using a more scalable method, configure client computers as follows:

1. Check if the following directory exists on the client computer, and create it if it’s missing:

   * Windows: `C:\ProgramData\Unity\config`
   * macOS: `/Library/Application\ Support/Unity/config`
2. Check if the directory in the previous step has a `services-config.json` file, and create it if it’s missing.
3. Add the `"enableProxyAutoconfig"` key to the file and assign a value of `true`. If you’re creating the file, make sure it resembles the following example:

   ```
   {
       "enableProxyAutoconfig": true
   }
   ```
4. Save the `services-config.json` file.

With automatic proxy configuration enabled, Unity applications will look for stored user credentials on [Windows](ent-proxy-autoconfig-store.md) or [macOS](ent-proxy-autoconfig-store-mac.md), if your organization’s web proxy requires authentication. If the credentials aren’t stored or aren’t correct, you’ll experience issues such as:

* Project templates won’t be available in the Unity Hub when you create a new project.
* Package Manager operations like searching or downloading packages will fail. In this case, the Editor displays error messages in the **Console window**:

![Console errors after a failed proxy authentication](../uploads/Main/ent-proxy-autoconfig-00.png)

Console errors after a failed proxy authentication

## Next steps

* If your organization set up its web proxy to require authentication, you need to store your credentials to authenticate with the web proxy. Refer to the following instructions:
  + Windows: [Store credentials for automatic proxy configuration (Windows)](ent-proxy-autoconfig-store.md)
  + macOS: [Store credentials for automatic proxy configuration (macOS)](ent-proxy-autoconfig-store-mac.md)
* Review the solutions in [Using Unity through web proxies](ent-proxy-autoconfig.md) to check if your environment requires any additional actions.

## Additional resources

* [Solving network issues](upm-config-network.md) (Unity Package Manager)