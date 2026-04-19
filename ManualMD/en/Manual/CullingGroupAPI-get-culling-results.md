# Get culling results

### Receiving results via the onStateChanged callback

The most efficient way to respond to spheres changing their visibility or distance state is to use the onStateChanged callback field. Set this to a function which takes a CullingGroupEvent structure as an argument; it will then be called after culling is complete, for each sphere that has changed state. The members of the CullingGroupEvent structure tell you about the previous and new states of the sphere.

```
group.onStateChanged = StateChangedMethod;

private void StateChangedMethod(CullingGroupEvent evt)
{
    if(evt.hasBecomeVisible)
        Debug.LogFormat("Sphere {0} has become visible!", evt.index);
    if(evt.hasBecomeInvisible)
        Debug.LogFormat("Sphere {0} has become invisible!", evt.index);
}
```

### Receiving results via the CullingGroup Query API

In addition to the onStateChanged delegate, the CullingGroup provides an API for retrieving the latest visibility and distance results of any sphere in the bounding spheres array. To check the states of a single sphere, use the IsVisible and GetDistance methods:

```
bool sphereIsVisible = group.IsVisible(0);
int sphereDistanceBand = group.GetDistance(0);
```

To check the states of multiple spheres, you can use the QueryIndices method. This method scans a continuous range of spheres to find ones that match a given visibility or distance state.

```
// Allocate an array to hold the resulting sphere indices - the size of the array determines the maximum spheres checked per call
int[] resultIndices = new int[1000];
// Also set up an int for storing the actual number of results that have been placed into the array
int numResults = 0;

// Find all spheres that are visible
numResults = group.QueryIndices(true, resultIndices, 0);
// Find all spheres that are in distance band 1
numResults = group.QueryIndices(1, resultIndices, 0);
// Find all spheres that are hidden in distance band 2, skipping the first 100
numResults = group.QueryIndices(false, 2, resultIndices, 100);
```

Remember that the information retrieved by the query API is only updated when the **camera** used by the CullingGroup actually performs its culling.