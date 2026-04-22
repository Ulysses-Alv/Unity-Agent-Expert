# Utility nodes | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Utility-Nodes.html)

# Utility nodes

Explore nodes that enable essential logic operations, customization with HLSL, complex math expressions, and previews.

| **Topic** | **Description** |
| --- | --- |
| [Custom Function](Custom-Function-Node.html) | Returns the results of a custom HLSL function specified as a string or in a referenced file. |
| [Expression](Expression-Node.html) | Returns the result of a custom mathematical expression specified as a string. |
| [Preview](Preview-Node.html) | Provides a preview window and passes the input value through without modification. |

## Logic

| **Topic** | **Description** |
| --- | --- |
| [All](All-Node.html) | Returns true if all components of the input In are non-zero. |
| [And](And-Node.html) | Returns true if both the inputs A and B are true. |
| [Any](Any-Node.html) | Returns true if any of the components of the input In are non-zero. |
| [Branch](Branch-Node.html) | Provides a dynamic branch to the shader. |
| [Comparison](Comparison-Node.html) | Compares the two input values A and B based on the condition selected on the dropdown. |
| [Is Infinite](Is-Infinite-Node.html) | Returns true if any of the components of the input In is an infinite value. |
| [Is NaN](Is-NaN-Node.html) | Returns true if any of the components of the input In is not a number (NaN). |
| [Nand](Nand-Node.html) | Returns true if both the inputs A and B are false. |
| [Not](Not-Node.html) | Returns the opposite of input In. If In is true, the output is false. Otherwise, it returns true. |
| [Or](Or-Node.html) | Returns true if either input A or input B is true. |
| [Switch](Switch-Node.html) | Enables branching based on different cases with value comparisons. |