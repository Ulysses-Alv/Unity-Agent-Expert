# Remove unused resources from your Web build

Remove unused packages, **shaders**, and engine code from a project to make the build smaller. A smaller build reduces application load times and improves the performance of your application.

## Remove or disable unused packages

Unused packages might affect the final output size of the build. Remove or disable all packages the project doesn’t use. This includes [built-in packages](pack-build.md) and packages that might be enabled by default, such as the [Input System package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest?subfolder=/manual/index.html), which can significantly increase the build size if it’s not removed.

You can use the Package Manager window to determine which packages are installed or enabled for your project:

* Select the **In Project** view and check the [list panel](upm-ui-list.md) for packages that are installed as direct and indirect [dependencies](upm-dependencies.md). You can also use the [details panel](upm-ui-details.md) to get more information about packages dependencies.
* Select the **Built-in** view for built-in packages that are enabled in your project.

To learn which packages your project needs, review the [project manifest file](upm-manifestPrj.md) located at `Packages/manifest.json`.

To remove a package, refer to [Remove a UPM package from a project](upm-ui-remove.md).

To disable a built-in package, refer to [Disable a built-in package](upm-ui-disable.md).

## Remove unused shader variants

Enable shader stripping to remove shader variants that aren’t used in a **scene**. Shader stripping can reduce your build file size.

**Important**: Test your application to make sure your unused shaders aren’t referenced by other shaders.

To enable shader stripping, refer to [Graphics](class-GraphicsSettings.md#stripping).

## Remove unused engine code

Use the [Web Stripping Tool package](https://docs.unity3d.com/Packages/com.unity.web.stripping-tool@latest?subfolder=/manual/index.html) to identify and remove unused engine code from your build. This tool is useful for projects that don’t use certain engine features, such as 3D graphics in a 2D game, allowing for a significant reduction in the final build size.

## Additional resources

* [Optimize your Web build](web-optimization.md)
* [Recommended Graphics settings to optimize your Web build](web-optimization-graphics.md)
* [Recommended Quality settings to optimize your Web build](web-optimization-quality.md)
* [Recommended Player settings to optimize your Web build](web-optimization-player.md)
* [Optimize Web platform for mobile](web-optimization-mobile.md)