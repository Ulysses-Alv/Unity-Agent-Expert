# Primitive and placeholder objects

In the Unity Editor, you can work with 3D models of any shape that you can create with modeling software. You can create several object types in the Editor, for example, the:

* Cube
* Sphere
* Capsule
* Cylinder
* Plane
* Quad

To add a primitive to a scene, select **GameObject** > **3D Object** and select the required primitive. This adds the default primitive to your local coordinate space, and you can change the size and position of the object in the Inspector window with [transforms](class-Transform.md).

These objects can be useful to model specific items or you can use them to create placeholders and prototypes for testing. This page details common uses of each primitive object, but you can use the objects in the way that suits your project best.

## Cube

![The cube primitive.](../uploads/Main/PrimitiveCube.png)

The cube primitive.

The default Cube primitive is a white cube with six sides, which are 1 x 1 units. The cube is textured so the image repeats on each side. Cubes are often used to construct walls, posts, boxes, steps and other similar objects or as a useful placeholder during development. For example, because the cube’s edges are one unit long, you can check the proportions of a **mesh** imported into your **scene** if you add a cube close to it and compare the size.

## Sphere

![The sphere primitive.](../uploads/Main/PrimitiveSphere.png)

The sphere primitive.

The default sphere primitive has a diameter of one unit (that is, 0.5 unit radius). The sphere is textured with standard spherical [UV mapping](https://en.wikipedia.org/wiki/UV_mapping) so that the image wraps around the sphere and the top and bottom edges pinch together at each pole. Spheres are useful for representing balls, planets, and projectiles. You can make GUI device for representing the radius of an effect using a semi-transparent sphere.

## Cylinder

![The cylinder primitive](../uploads/Main/PrimitiveCylinder.png)

The cylinder primitive

The default cylinder primitive is two units high and one unit in diameter. The cylinder is textured so that the image wraps once around the tube shape of the body and repeats separately on the two flat ends. Cylinders are useful if you want to create posts, rods, and wheels. The default shape of a **Collider** is a capsule (there is no primitive cylinder collider in Unity). To get an accurate cylindrical physics Collider, create a mesh of a cylinder shape in a modeling program and attach a **mesh Collider**.

## Capsule

![The capsule primitive.](../uploads/Main/PrimitiveCapsule.png)

The capsule primitive.

The capsule primitive is a cylinder with hemispherical caps at each end. The default capsule is one unit in diameter and two units high (the body is one unit, and each cap is 0.5 units). It is textured so the image wraps around once and the edges pinch together at the apex of each hemisphere. The capsule is a useful placeholder for prototypes because the physics of a rounded object are more suited to certain tasks than those of a box.

## Quad

![The quad primitive.](../uploads/Main/PrimitiveQuad.png)

The quad primitive.

The default **quad** primitive is a square with edges one unit long, divided into two triangles, and is oriented in the xy plane of the local coordinate space. You can use a quad as a display screen for an image or movie. You can also use quads to implement simple GUI and information displays, **particles**, **sprites**, and imposter images as substitutes for solid objects views at a distance.

## Plane

![The plane primitive.](../uploads/Main/PrimitivePlane.png)

The plane primitive.

The default plane primitive is a flat square with edges ten units long, divided into 200 triangles, and is oriented in the xz plane of the local coordinate space. The square is textured so that the complete image appears once within the square.

Planes are useful for most flat surfaces, such as floors and walls. You can use a plane to show special effects and images or movies in GUI, but it might be easier to use a quad primitive.

A plane’s texture only renders from above. If you are below the plane, the plane’s texture is transparent.