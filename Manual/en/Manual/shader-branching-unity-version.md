# Branch based on Unity version

`UNITY_VERSION` contains the numeric value of the Unity version.

This can be used for version comparisons if you need to write shaders that use different built-in shader functionality. For example, use `#if UNITY_VERSION >= 60000000` if you want the preprocessor check to pass only on Unity version 6.0 or later.

## Unity 6

Use the format `6MMMPPPP`, where:

* `6` is Unity 6.
* `MMM` is the minor version with leading zeroes, for example `002` for Unity 6.2.
* `PPPP` is the patch version with leading zeroes, for example `1234` for Unity 6.2.1234.

For example, for Unity 6000.2.3, use `60020003`.

## Unity 2023 or earlier

Use the format `YYYYMP`, where:

* `YYYY` is the major version.
* `M` is the minor version.
* `P` is the patch version.

For example, for Unity 2022.3.0, use `202230`.

You can use only values up to 9 for the minor and patch versions. This means you can’t check for a Unity version with a minor version larger than 9 or a patch version larger than 9.