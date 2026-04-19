# PackageRequirements block in ShaderLab reference

To specify package requirements for a SubShader or a Pass, you use the `PackageRequirements` block. **ShaderLab** supports a single `PackageRequirements` block per SubShader or Pass, but each block can specify multiple package requirements.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: PackageRequirements block** | Yes | Yes | Yes | Yes |

## Syntax

**Note**: if you provide a `PackageRequirements` block, it must come before all other declarations inside the SubShader or Pass.

| **Signature** | **Function** |
| --- | --- |
| `PackageRequirements{ [requirement definition]}` | Defines the package requirements for the Pass or SubShader. |

There are multiple ways to declare package requirements. Each one provides a different behavior. They are:

* **“<package name>”**: Specifies that the SubShader or Pass works with any version of the package.
* **“<package name>”: “<version restrictions>”**: Specifies that the SubShader or Pass only works with a subset of package versions.
* **“<package name>”: “unity=<version restrictions>”**: Specifies that the SubShader or Pass only works with a subset of Unity versions and requires a package with the given name.
* **“unity”:“<version restrictions>”**: Specifies that the SubShader or Pass only works with a subset of Unity versions.

### Version syntax

In ShaderLab’s package requirements, a version uses the `major.minor` or `major.minor.patch` format. If you only use `major.minor`, Unity uses `0` for the `patch`. Package versions can also include a `-preview` or a `-preview.n` postfix where `-preview` is equivalent to `-preview.0`. Preview versions come before non-preview versions, so `1.2.3-preview.4` comes after `1.2.2` but before `1.2.3`.

There are multiple ways to specify a version range. Each one provides a different behavior. They are:

* **<version>**: Includes the version you enter and all versions after that. For example, `1.2.3` includes all versions starting with `1.2.3`;
* **[<version>]**: Specifies the exact version. For example, `[1.2.3]` only includes version `1.2.3`;
* **[<version1>,<version2>]**: Specifies a range between <version1> and <version2>. Using square brackets and round brackets causes the range to include or exclude to version respectively. The opening bracket affects <version1> and the closing bracket affects <version2>. Example: `[1.2.3,2.3.4)` includes all versions from 1.2.3 to 2.3.3.

You can also specify sets of version ranges for a single package. To create a set of version ranges from individual ranges, use a semicolon as a separator. For example, `[2.0,3.4.5];[3.7];4.0` includes versions from 2.0.0 to 3.4.5, version 3.7.0, and version 4.0.0 and above.

When you set the versions for a package, be aware of the following:

* Versions, version ranges, and sets of version ranges cannot contain any extra characters.
* Version ranges cannot be empty.
* Sets of version ranges cannot have intersections.
* Package requirements are sensitive to whitespace.

If the syntax does not adhere to the above, the version restriction is invalid. For more information, refer to [Set a shader to require a package](writing-shader-tags-require-package.md).

## Additional resources

* [Set a shader to require a package](writing-shader-tags-require-package.md)