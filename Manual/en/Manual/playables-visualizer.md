## PlayableGraph Visualizer

The PlayableGraph Visualizer displays a graph of your PlayableGraph at runtime. Unity must be in Play mode to use the Playable Visualizer to display PlayableGraphs. When in Edit mode, the Playable Visualizer displays graphs for Mecanim features and packages that use **Playables** such as **Animator Controllers** and the Unity Timeline package.

**Note**: The PlayableGraph Visualizer is a discontinued experimental package that might not work with your version of Unity.

To install and use the PlayableGraph Visualizer:

1. Download the PlayableGraph Visualizer from its [GitHub repository](https://github.com/UnityTech/graph-visualizer).
2. Add the package to your project.
3. Select **Window** > **Analysis** > **PlayableGraph Visualizer**.
4. Select **Play** to switch the Editor to Play mode.
5. In the PlayableGraph Visualizer window, select the PlayableGraph to display its graph.

![The PlayableGraph Visualizer window with the CustomPlayableExample PlayableGraph selected](../uploads/Main/playables-vis-custom-playqueue.png)

The PlayableGraph Visualizer window with the `CustomPlayableExample` PlayableGraph selected

Colored nodes represent Playables in the graph. Wire color intensity indicates blending weight. Consult [GitHub](https://github.com/UnityTech/graph-visualizer) for more information on the PlayableGraph Visualizer.

## Additional resources

* [The PlayableGraph, nodes, and output](Playables-Graph.md)
* [Install a UPM package from a local folder](upm-ui-local.md)
* [Playable struct](../ScriptReference/Playables.Playable.md)