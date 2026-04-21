# Expression node | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Expression-Node.html)

# Expression node

Use the Expression node to specify a complex mathematical expression as a string instead of using multiple [Math nodes](Math-Nodes.html).

## Input ports

The number and type of input ports automatically adjust to the values of the [node controls](#controls). Unity adds an input port for each variable in the expression in the text field.

## Output port

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| **Out** | Depends on the selected **Type** in the [node controls](#controls). | The value resulting from the expression specified in the text field. |

## Controls

| **Name** | **Description** |
| --- | --- |
| **Type** | Specifies the data type to use for all input and output ports of the node. The options are:  * **Vector 1** * **Vector 2** * **Vector 3** * **Vector 4** |
| (Text field) | Use this field to specify the mathematical expression. The field supports the following:  * Math operators: `+`, `-`, `*`, `/`, and parentheses. * HLSL intrinsic functions: for example, `frac` and `saturate`. * Vector swizzling: for example, `a.xw`. * Vector construction: `float2(x, y)`. |

## Example

If you specify the expression:

`a + b * c / 2`

* Unity adds three input ports to the node: **A**, **B**, and **C**.
* If you input 1, 2, and 10, the node calculates `1 + 2 * 10 / 2` and outputs the result to the **Out** port.

**Note**: If you use Shader Graph math nodes instead, the example requires three separate nodes: [Divide](Divide-Node.html), [Multiply](Multiply-Node.html), and [Add](Add-Node.html).