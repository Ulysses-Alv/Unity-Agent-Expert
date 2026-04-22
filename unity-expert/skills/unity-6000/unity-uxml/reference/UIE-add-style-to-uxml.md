# Add styles to UXML

In UI Toolkit, you can use [USS](UIE-USS.md) to customize the appearance of **visual elements**. The suggested workflow for USS is that you visually style an element in UI Builder, extract the style to a USS file, and then reference it in UXML.

If you style an element in UI Builder, the style is added as an inline style to the `style` attribute of UXML elements:

```
<ui:UXML ...>
    <ui:VisualElement style="width: 200px; height: 200px; background-color: red;" />
</ui:UXML>
```

To reference a stylesheet file, add the `<Style>` element under the root element of a UXML file.

For example, if you have a USS file named `styles.uss` with the following content:

```
#root {
    width: 200px;
    height: 200px;
    background-color: red;
}
```

You can reference it like this:

```
<ui:UXML ...>
    <Style src="<path-to-file>/styles.uss" />
    <ui:VisualElement name="root" />
</ui:UXML>
```

You can use a relative or an absolute path:

* Absolute paths start from the project’s `Assets` folder and begin with a `/` or `project://database/`. For example, `/Assets/myFolder/myFile.uss` or `project://database/Assets/myFolder/myFile.uss`.
* Relative paths start from the current file and exclude the `/`. For example, `../myFolder/myFile.uss`.

**Note**: To reference a file from packages, use the absolute path of the package file that starts from the `Packages` folder. For example, `/Packages/com.unity.package.name/file-name.uss` or `project://database/Packages/com.unity.package.name/file-name.uss`. you must use the format of `com.unity.package.name` rather than `package name` for the package name.

## Additional resources

* [Style UI](UIE-USS.md)
* [Best practices for managing elements](UIE-best-practices-for-managing-elements.md)