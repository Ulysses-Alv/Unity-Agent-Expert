# UnityYAML

Unity uses a custom-optimized YAML library called UnityYAML. The UnityYAML library does not support the [full YAML specification](https://yaml.org/spec/). This documentation outlines which parts of the YAML spec UnityYAML supports.

You cannot externally produce or edit UnityYAML files.

## Supported features

| **Feature** | **Support** |
| --- | --- |
| **Mappings** | UnityYAML supports both flow and block styles. |
| **Scalars** | UnityYAML supports double and single quoted scalars as well as plain scalars. You can split them onto multiple lines. Be aware that multi-line scalars can create performance and memory overheads during parsing.  Plain scalars split onto multiple lines must be indented more than the previous line. See below this table for an example.  You can use UTF–8 characters in scalars, but UnityYAML only decodes them when they are part of a double quoted scalar. |
| **Sequences** | UnityYAML supports mapping, block styles, and block sequences that contain block mappings. |

### Example of indentation on multi-line plain scalars:

```
parent: This is a
  multi-line scalar
^
|
```

If there is no indentation, the scalar returns `This is a` and might trigger an Asset into further parsing.

## Unsupported features

| **Feature** | **Support** |
| --- | --- |
| **Chomping indicators** | UnityYAML does not support using `+` and `|` characters to indicate how it should treat new lines within a multi-line string. If you use these characters, UnityYAML adds them to the scalar value. |
| **Comments** | UnityYAML does not support comments. |
| **Complex mapping keys** | UnityYAML does not support complex mapping keys. |
| **Multiple documents** | The reader skips document and tag prefixes at the top of files, but does not handle YAML input that consists of multiple documents. |
| **Raw block sequences** | Nearly all nodes are part of a mapping in UnityYAML, so all sequences must be values of a mapping to work correctly. See below this table for an example.  Anonymous sequences increase the parser complexity. You cannot use indentation as a way of determining if a sequence element has finished in UnityYAML. |
| **Tags** | UnityYAML does not support tags. |

### Example of a raw block sequence

```
var:
  - 1
  - 2
  - 3
```

The sequence is designed for lookups upon `var`, so the following does not work:

```
- 1
- 2
- 3
```