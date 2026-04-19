# Recommended Quality settings to optimize your Web build

Use the following recommended **Quality** settings to optimize your builds for the Unity Web platform.

Configure the following recommended settings in **Edit** > **Project Settings** > **Quality**. For more information on these settings, refer to [Quality](class-QualitySettings.md).

## Quality Level

The **Quality** settings contain a group of possible **Quality Levels** for your build. If you set your **Quality Level** to a lower setting (**Very Low** or **Low**), it results in faster load times and better performance which is recommended for a Web application. However, lower quality can mean the visuals don’t look as good.

For more information about this setting, refer to [Quality](class-QualitySettings.md).

### Change the Quality Level via C# code

To change the **Quality Level** via script instead, add this code to one of your **scripts**:

```
QualitySettings.SetQualityLevel(0, true); // Use the index of your quality settings matrix (0 = Very Low, 1 =  Low).
```

## Additional resources

* [Optimize your Web build](web-optimization.md)
* [Recommended Graphics settings to optimize your Web build](web-optimization-graphics.md)
* [Recommended Player settings to optimize your Web build](web-optimization-player.md)
* [Remove unused resources from your Web build](web-optimization-remove-resources.md)
* [Optimize Web platform for mobile](web-optimization-mobile.md)