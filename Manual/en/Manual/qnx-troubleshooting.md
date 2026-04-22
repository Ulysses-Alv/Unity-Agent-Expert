# Troubleshooting the QNX Player

This page lists the common problems and limitations that might occur when using the QNX Player.

## The QNX Player fails to enumerate all display resolutions

Due to an ABI breakage in the QNX 7.1 screen system headers, the Player might not be able to query all available display modes. For more information, refer to the [QNX Article - Ref# J2941150](https://www.qnx.com/developers/articles/rel_6934_0.html).

## Rendering issues with MSAA on i.MX8 QNX devices

Multisample anti-aliasing (MSAA) might cause rendering issues on i.MX8 QNX devices.

### Solution

To avoid rendering issues on i.MX8 QNX devices, disable MSAA in your project’s quality settings. Go to **Edit** > **Project Settings** > **Quality**, then set **Anti Aliasing** to **Disabled**.

## Additional resources

* [Install the platform package for QNX](qnx-install-editor.md)
* [Set up your environment for QNX](qnx-environment-setup.md)
* [QNX Build Profiles reference](qnx-build-settings.md)