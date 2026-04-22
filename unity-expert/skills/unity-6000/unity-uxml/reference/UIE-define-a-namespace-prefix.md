# Define a namespace prefix

After you have defined a custom control element, you can use it in your UXML files. To categorize elements, create the class in a namespace. When you define a new namespace, you can define a prefix for the namespace. You must define namespace prefixes as attributes of the root `<UXML>` element and replace the full namespace name when scoping elements.

To define a namespace prefix, add a `UxmlNamespacePrefix` attribute to your assembly for each namespace prefix. For example:

```
[assembly: UxmlNamespacePrefix("My.First.Namespace", "first")]
[assembly: UxmlNamespacePrefix("My.Second.Namespace", "second")]
```

You can do this at the root level (outside any namespace) of any C# file of the assembly.

The schema generation system does the following:

* Checks for these attributes and uses them to generate the schema.
* Adds the namespace prefix definition as an attribute of the `<UXML>` element in newly created UXML files.
* Includes the schema file location for the namespace in its `xsi:schemaLocation` attribute.

To ensure that your text editor recognizes the new element, select **Assets** > **Update UXML Schema** to update the schema definition.

To create a new UXML document with the prefix, select **Assets** > **Create** > **UI Toolkit** > **UI Document**.

## Additional resources

* [Create custom controls](UIE-create-custom-controls.md)