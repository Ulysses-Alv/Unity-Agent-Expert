# Graph Inspector | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Internal-Inspector.html)

# Graph Inspector

## Description

The **Graph Inspector** makes it possible for you to interact with any selectable graph elements and graph-wide settings for a [Shader Graph Asset](Shader-Graph-Asset.html). You can use the **Graph Inspector** to edit attributes and default values.

When you open a Shader Graph, the **Graph Inspector** displays the **[Graph Settings](Graph-Settings-Tab.html)** tab by default. Graph-wide settings for that specific Shader Graph appear in this tab.

## How to use

Select a node in the graph to display settings available for that node in the **Graph Inspector**. Settings available for that node appear in the **Node Settings** tab of the Graph Inspector. For example, if you select a Property node either in the graph or the [Blackboard](Blackboard.html), the **Node Settings** tab displays attributes of the Property that you can edit.

![The Blackboard with a property selected, and the Graph Inspector showing the property settings.](images/InternalInspectorBlackboardProperty.png)

Graph elements that currently work with the Graph Inspector:

* [Properties](https://docs.unity3d.com/Manual/SL-Properties.html)

  ![A property selected in the workspace, and the Graph Inspector showing the property settings.](images/InternalInspectorGraphProperty.png)
* [Keywords](Keywords.html)

  ![The Blackboard with a keyword selected, and the Graph Inspector showing the keyword settings.](images/keywords_enum.png)
* [Custom Function nodes](Custom-Function-Node.html)

  ![A Custom Function node selected in the workspace, and the Graph Inspector showing the node settings.](images/Custom-Function-Node-File.png)
* [Subgraph Output nodes](Sub-graph.html)

  ![A subgraph output node selected in the workspace, and the Graph Inspector showing the node settings.](images/Inspector-SubgraphOutput.png)
* [Per-node precision](Precision-Modes.html)

  ![A node selected in the workspace, and the Graph Inspector showing the Precision setting.](images/Inspector-PerNodePrecision.png)

Graph elements that currently do not work with the Graph Inspector:

* Edges
* [Sticky Notes](Sticky-Notes.html)
* Groups

## Material Override

Enabling the [Allow Material Override](surface-options.html) option in the Graph Settings makes it possible for you to override certain graph properties via the Material Inspector.