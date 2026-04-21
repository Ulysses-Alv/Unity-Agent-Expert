# Add a custom interpolator | Shader Graph | 17.6.0

Source: [Original](https://docs.unity3d.com/Packages/com.unity.shadergraph@17.6/manual/Custom-Interpolators.html)

# Add a custom interpolator

To pass custom data from the vertex shader to the fragment shader, add a custom interpolator to the vertex context of the Master Stack.

There are two target audiences for custom interpolators:

* Technical Directors and Lead Technical Artists setting up environments for their teams.
* Graphics programmers helping artists to optimize content performance.

**Note:** If you use the Built-In Render Pipeline, refer to [Input vertex data into a shader](https://docs.unity3d.com/Manual/SL-VertexProgramInputs.html) instead.

##### Important

The Built-In Render Pipeline is deprecated and will be made obsolete in a future release.  
It remains supported, including bug fixes and maintenance, through the full Unity 6.7 LTS lifecycle.  
For more information on migration, refer to [Migrating from the Built-In Render Pipeline to URP](https://docs.unity3d.com/6000.5/Documentation/Manual/urp/upgrading-from-birp.html) and [Render pipeline feature comparison](https://docs.unity3d.com/6000.5/Documentation/Manual/render-pipelines-feature-comparison.html).

## Supported data types

Custom interpolators support float, vector 2, vector 3, and vector 4 types.

## Channel limits

A custom interpolator supports a maximum of 32 channels. A channel is equivalent to four floats. Each float is an interpolator variable.

Different platforms and GPUs may have different limits, which might prevent your shaders compiling. Test your custom interpolators on your build targets to make sure your shaders compile properly. For more information, refer to the **Interpolator count limits** section in [Input vertex data into a shader](https://docs.unity3d.com/Manual/SL-VertexProgramInputs.html).

You can't limit the number of channels another user creates in a shader graph. However, to warn users about the limits, go to [Project Settings](https://docs.unity3d.com/Manual/comp-ManagerGroup.html) and set the following:

* **Warning Threshold** to tell users when they approach the channel limit. The range is 8 to 32 channels.
* **Error Threshold** to tell users when they reach or exceed the channel limit. The minimum value is 8 channels, and it must be higher than the **Warning Threshold**.

## Add a custom interpolator block to the Master Stack

1. Right-click in the **Vertex** context to create a block node.
2. Select **Custom Interpolator**.
3. In the **Node Settings** tab of the **Graph Inspector** window, select a data type, for example **Vector 4**.
4. Select an [interpolation method](Custom-Interpolators-reference.html).
5. In the same tab, enter a name for the interpolator.

## Write data to the interpolator

1. Right-click in your graph to create a node.
2. Select the type, for example **Vertex ID**.

   Custom interpolator blocks support many types of data, so you can connect the data from many other nodes including UV nodes and color nodes.
3. Connect the node to the custom interpolator block.

The graph now writes Vertex ID values into the custom interpolator.

## Read data from the interpolator

1. Right-click in your graph to create a node.
2. Select **Custom Interpolator**.
3. Connect the **Custom Interpolator** node to the relevant block in the **Fragment** context, for example **Base Color** to use the Vertex ID as color output.

## Delete a custom interpolator

If you delete a custom interpolator that's associated with nodes that are still in your graph, Unity displays an alert. If you want to keep using these nodes, you can create a new custom interpolator and associate the nodes with it. This prevents the alert from appearing.

## Additional resources

* [Built In Blocks](Built-In-Blocks.html)
* [Custom Interpolator reference](Custom-Interpolators-reference.html)