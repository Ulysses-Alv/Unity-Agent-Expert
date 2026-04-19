# Text-based scene files

Unity provides a text-based format for **scene** data, in addition to the default binary format. This can be useful when working with **version control** software, because text files generated separately can be merged more easily than binary files. Also, the text data can be generated and parsed by tools, making it possible to create and analyze scenes automatically. The pages in this section provide some reference material for working with the format.

| **Topic** | **Description** |
| --- | --- |
| **[Format of text serialized files](FormatDescription.md)** | Understand the structure and syntax of Unity’s text-based serialization format for scenes and assets. |
| **[UnityYAML](UnityYAML.md)** | Learn about Unity’s implementation of YAML and the specific conventions used in scene files. |
| **[An example of a YAML scene file](YAMLSceneExample.md)** | An example of a scene file in YAML format with explanations of key elements. |
| **[YAML class ID reference](ClassIDReference.md)** | Reference list of class IDs used in YAML files to identify different Unity object types. |

## Additional resources

* [Creating, loading, and saving scenes](scenes-working-with.md)
* [Script serialization](script-serialization.md)
* [Content output of a build](build-content-output.md)