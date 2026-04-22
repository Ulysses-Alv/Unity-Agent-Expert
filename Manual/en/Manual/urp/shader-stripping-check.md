# Check how many shader variants your build has in URP

To log how many variants Unity compiles and strips in total in the Universal **Render Pipeline** (URP), follow these steps:

1. Open the [Graphics settings window](urp-global-settings.md).
2. Open the **URP** tab.
3. In the **Additional **Shader** Stripping Settings** section, select a **Shader Variant Log Level** other than **Disabled**.
4. From the main menu, select **File** > **Build Profiles**.
5. Enable ****Development Build****.
6. Build your project.

Unity logs the shader variant information to the following locations:

* The [Console window](../Console.md)
* The `Editor.log` log file. Search for the section that begins with `Shader Stripping`. For the location of `Editor.log`, refer to [log files](../log-files.md).

For each shader, Unity uses the format `Total=<Number of variants in build>/<Total number of variants>(<Percentage in build>)`. For example, the following log indicates that Unity kept 8 of the 39 variants of the [Lit shader](lit-shader.md), including 1 of the 6 vertex shader variants of the `ForwardLit` shader pass:

```
Universal Render Pipeline/Lit - Total=8/39(20.51%)
- ForwardLit (ScriptableRenderPipeline) (SubShader: 0) (ShaderType: Vertex) - Total=1/6(16.67%) - Time=0.0691ms
...
```

To export the information in JSON format, follow these steps:

1. In the Graphics settings window, open the **URP** tab.
2. In the **Additional Shader Stripping Settings** section, enable **Export Shader Variants**.
3. Build your project.
4. In your project folder, open the `Temp/graphics-settings-stripping.json` and `Temp/shader-stripping.json` files.

## Additional resources

* [Check how many shader variants you have](../shader-how-many-variants.md)
* [Shader variant stripping](../shader-variant-stripping.md)