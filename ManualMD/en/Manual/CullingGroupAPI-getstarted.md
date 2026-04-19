# Create a Culling Group

There are no components or visual tools for working with CullingGroups; they are purely accessible via script.

A CullingGroup can be constructed using the ‘new’ operator:

```
CullingGroup group = new CullingGroup();
```

To have the CullingGroup perform visibility and/or distance calculations, specify the **camera** it should use:

```
group.targetCamera = Camera.main;
```

Create and populate an array of BoundingSphere structures with the positions and radii of your spheres, and pass it to SetBoundingSpheres along with the number of spheres that are actually in the array. The number of spheres does not need to be the same as the length of the array. Unity recommends that you create an array that is big enough to hold the most spheres you will ever have at one time, even if the initial number of spheres you actually have in the array is very low. Using a larger array allows you to add or remove spheres as needed without the computationally expensive process of resizing the array at runtime.

```
BoundingSphere[] spheres = new BoundingSphere[1000];
spheres[0] = new BoundingSphere(Vector3.zero, 1f);
group.SetBoundingSpheres(spheres);
group.SetBoundingSphereCount(1);
```

At this point, the CullingGroup will begin computing the visibility of the single sphere each frame.

To clean up the CullingGroup and free all memory it uses, dispose of the CullingGroup via the standard .NET IDisposable mechanism:

```
group.Dispose();
group = null;
```