# Reference other files from UXML

In a UXML file, you can use the `<Template>` and the `<Style>` elements to reference other UXML or USS files. The two elements both accept either an `src` attribute or a `path` attribute.

## The `src` attribute

This is the recommended way to reference other files.

Use the following syntax for the `src` attribute:

```
src="<path-to-file>/<file-name-with-extension>"
```

Any errors during import, such as missing files, trigger an error message.

You can use a relative or an absolute path:

* Absolute paths start from the project’s `Assets` folder and begin with a `/` or `project://database/`. For example, `/Assets/myFolder/myFile.uss` or `project://database/Assets/myFolder/myFile.uss`.
* Relative paths start from the current file and exclude the `/`. For example, `../myFolder/myFile.uss`.

**Note**: To reference a file from packages, use the absolute path of the package file that starts from the `Packages` folder. For example, `/Packages/com.unity.package.name/file-name.uss` or `project://database/Packages/com.unity.package.name/file-name.uss`. you must use the format of `com.unity.package.name` rather than `package name` for the package name.

## The `path` attribute

The `path` attribute uses the Unity Resources mechanisms, but doesn’t offer error reporting at import time and doesn’t allow relative paths. This is a legacy way to reference other files and is not recommended. Use it only if you need to reference files that can only be loaded from the `Resources` folder.

The `path` attribute accepts files located in either the `Resources` folder or the `Editor Default Resources` folder, with the following rules:

* If the file is in the `Resources` folder, don’t include the file extension. For example, write `path="template"` for a file located at `Assets/Resources/template.uxml`.
* If the file is in the `Editor Default Resources` folder, you must include the file extension. For example, write `path="template.uxml"` for a file located at `Assets/Editor Default Resources/template.uxml`.

## Additional resources

* [Add styles to UXML](UIE-add-style-to-uxml.md)
* [Reuse UXML files](UIE-reuse-uxml-files.md)