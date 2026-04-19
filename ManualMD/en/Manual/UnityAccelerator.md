# Introduction to Unity Accelerator

Unity **Accelerator** is a caching proxy agent that keeps copies of a team’s imported assets to help reduce development iteration time. Unity Accelerator uses the Asset Import Pipeline to coordinate asset sharing when your team is working on the same local network so that you don’t need to reimport portions of your project.

When you install Unity Accelerator on your local network, the Unity Editor communicates with the cache server to store and retrieve assets that other team members are using or might have modified.

## Installing Unity Accelerator

To install Unity Accelerator, you can either install it [with the installer](accelerator-install-installer.md) or [with Docker Hub](accelerator-install-docker.md). Make sure that the system you install it on meets the minimum [requirements](accelerator-requirements.md).

After installation, you can [verify the Unity Accelerator version](accelerator-verify-version.md) and [configure Unity Accelerator in the Editor](accelerator-configure.md).

## Using Unity Accelerator

After you’ve installed Unity Accelerator on your network and set it up in your project, it will work in the background, caching your assets. You [can monitor its usage](accelerator-monitor.md) via the dashboard, at `http://<accelerator-host-ip>/dashboard`.

## Additional resources

* [Unity Accelerator requirements](accelerator-requirements.md)
* [Importing Assets](ImportingAssets.md)
* [The Asset Database](AssetDatabase.md)