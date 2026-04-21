# Custom Port Menu | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Custom-Port-Menu.html)

# Custom Port Menu

## Description

The Custom [Port](Port.html) Menu is displayed in the **Node Settings** tab of the [Graph Inspector](Internal-Inspector.html) by clicking on the [Custom Function Node](Custom-Function-Node.html) and [Sub Graph](Sub-graph.html) output node. This menu allows you to add, remove, rename, reorder, and define the types of your custom input and output ports.

## How to Use

Select the [Custom Function Node](Custom-Function-Node.html) or the [Sub Graph](Sub-graph.html) output node to view the Custom Port Menu in the Inspector. To close the menu, click anywhere in the graph or on another graph-element.

![01](images/Custom-Port-Menu-Empty.png)

### Adding and Removing Ports

To add ports, click the `+` icon at the bottom right corner of the port list.

![02](images/Custom-Port-Menu-Add.png)

To remove ports, select a port using the hamburger icon on the left, and click the `-` icon at the bottom right corner of the port list.

![03](images/Custom-Port-Menu-Remove.png)

### Renaming Ports

To rename a port, double-click its text field and enter the new name. Currently, only the following characters are valid for port names: A-Z, a-z, 0-9, \_, ( ), and whitespace. If the name contains an invalid character, an error badge appears.

![04](images/Custom-Port-Menu-Rename.png)

### Reordering Ports

To reorder ports, click and hold the hamburger icon on the left, and drag the port to your desired place in the list.

### Changing Port Types

To change a port type, use the Type drop-down menu on the right. See the [Data Types](Data-Types.html) page for a list of currently valid port types.

![05](images/Custom-Port-Menu-Type.png)