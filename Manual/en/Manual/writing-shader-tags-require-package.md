# Set a shader to require a package

To set a **shader** to only compile if a specific package is installed, add a `PackageRequirements` block in the `SubShader` block or `Pass` block.

When Unity imports the shader, it checks if the specified package is installed in the project, and if the package version meets the requirements. If the requirements aren’t met, Unity skips further processing of the subshader or shader pass.

**Note:** Unity warns you in the **Console window** if the `PackageRequirements` blocks in a shader mean none of the subshaders match the installed packages.

For more information, refer to [PackageRequirements block in ShaderLab reference](SL-PackageRequirements.md).

## Require a package version

Inside the `PackageRequirements` block, use the following syntax:

* `"com.my.package": "MAJOR.MINOR"` to require package version `MAJOR.MINOR` or later.
* `"unity": "6000.MINOR"` to require Unity version `6000.MINOR` or later.

The `PackageRequirements` block must be the first block in the `SubShader` block or `Pass` block.

## Require a version within a range

Inside the `PackageRequirements` block, use `[MINIMUM,MAXIMUM]`, where `MININUM` is the minimum version and `MAXIMUM` is the maximum version. The range is inclusive. For example:

* `"com.my.package": "[17.2,17.5]"` to require package version `17.2` or later, up to and including `17.5`.
* `"unity": "[6000.2,6000.3]"` to require Unity version `6000.2` or later, up to and including `6000.3`.

To specify a non-inclusive range, use parentheses. For example `"com.my.package": "(17.2,17.5)"` to require package version `17.3` to `17.4`, excluding `17.2` and `17.5`.

To specify multiple ranges, use a semicolon to separate the ranges. For example `"unity": "[6000.0,6000.1];[6000.3,6000.4]"`.

### Example

The following shader compiles only if the project has the `com.my.package` package installed with a version of 2.2 or higher. The shader pass compiles only if the project has both the Universal **Render Pipeline** package with a version between 16.2.1 and 17.0, and the TextMesh Pro package with a version of 3.2 or higher.

```
Shader "Examples/ExampleShader"
{
    SubShader
    {
        PackageRequirements
        {
            "com.my.package": "2.2"
        }
        
        Pass
        {
            PackageRequirements
            {
                "com.unity.render-pipelines.universal": "[16.2.1,17.0]"
                "com.unity.textmeshpro": "3.2"
            }
            ...
        }
    }
}
```

## Additional resources

* [PackageRequirements block in ShaderLab reference](SL-PackageRequirements.md)