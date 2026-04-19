# Introduction to normal maps (bump mapping)

Normal maps are a type of ****Bump Map****. They are a special kind of texture that allow you to add surface detail such as bumps, grooves, and scratches to a model which catch the light as if they are represented by real geometry.

Unity uses Y+ **normal maps**, sometimes known as OpenGL format.

For example, you might want to show a surface which has grooves and screws or rivets across the surface, like an aircraft hull. One way to do this would be to model these details as geometry, as shown below.

![A sheet of aircraft metal with details modeled as real geometry.](../uploads/Main/StandardShaderNormalMapBadGeometry.jpg)

A sheet of aircraft metal with details modeled as real geometry.

Depending on the situation it is not normally a good idea to have such tiny details modelled as “real” geometry. On the right you can see the polygons required to make up the detail of a single screwhead. Over a large model with lots of fine surface detail this would require a very high number of polygons to be drawn. To avoid this, we should use a normal map to represent the fine surface detail, and a lower resolution polygonal surface for the larger shape of the model.

If we instead represent this detail with a bump map, the surface geometry can become much simpler, and the detail is represented as a texture which modulates how light reflects off the surface. This is something modern graphics hardware can do extremely fast. Your metal surface can now be a low-poly flat plane, and the screws, rivets, grooves and scratches will catch the light and appear to have depth because of the texture.

![The screws, grooves and scratches are defined in a normalmap, which modifies how light reflects off the surface of this low-poly plane, giving the impression of 3D detail. As well as the rivets and screws, a texture allows us to include far more detail like subtle bumps and scratches.](../uploads/Main/StandardShaderNormalMapAircraftSurface.jpg)

The screws, grooves and scratches are defined in a normalmap, which modifies how light reflects off the surface of this low-poly plane, giving the impression of 3D detail. As well as the rivets and screws, a texture allows us to include far more detail like subtle bumps and scratches.

In modern game development art pipelines, artists will use their 3D modelling applications to generate normal maps based on very high resolution source models. The normal maps are then mapped onto a lower-resolution game-ready version of the model, so that the original high-resolution detail is rendered using the normalmap.

## How normal mapping works

Normal mapping takes this modification of [surface normals](StandardShaderMaterialParameterNormalMapSurfaceNormals.md) one step further, by using a texture to store information about how to modify the surface normals across the model. A normal map is an image texture mapped to the surface of a model, similar to regular colour textures, however each **pixel** in the texture of the normal map (called a **texel**) represents a deviation in surface normal direction away from the “true” surface normal of the flat (or smooth interpolated) polygon.

![Normal mapping across three polygons, viewed as a 2D diagram](../uploads/Main/BumpMapBumpShadingDiagram.svg)

Normal mapping across three polygons, viewed as a 2D diagram

In this diagram, which is again a 2D representation of three polygons on the surface of a 3D model, each orange arrow corresponds to a pixel in the normalmap texture. Below, is a single-pixel slice of a normalmap texture. In the centre, you can see the normals have been modified, giving the *appearance* of a couple of bumps on the surface of the polygon. These bumps would only be apparent due to the way lighting appears on the surface, because these modified normals are used in the lighting calculations.

The colours visible in a raw normal map file typically have a blueish hue, and don’t contain any actual light or dark shading - this is because the colours themselves are not intended to be displayed as they are. Instead, the RGB values of each texel represent the X,Y & Z values of a direction vector, and are applied as a modification to the basic interpolated smooth normals of the polygon surfaces.

![An example normal map texture](../uploads/Main/BumpMapTexturePreview.png)

An example normal map texture

This is a simple normal map, containing the bump information for some raised rectangles and text. This normal map can be imported into Unity and placed into Normal Map slot of the Standard **Shader**. When combined in a material with a colour map (the Albedo map) and applied to the surface of the cylinder **mesh** above, the result looks like this:

![The example normal map applied to the surface of the cylinder mesh used above](../uploads/Main/BumpMapLitExample.jpg)

The example normal map applied to the surface of the cylinder mesh used above

Again, this does not affect the actual polygonal nature of the mesh, only how the lighting is calculated on the surfaces. This apparent raised lettering and shapes on the surface are not really present, and viewing the faces at glancing angles will reveal the true nature of the flat surface, however from most viewing angles the cylinder now appears to have embossed detail raised off the surface.

## Difference between bump maps, normal maps and height maps

**Normal Maps** and **Height Maps** are both *types* of Bump Map. They both contain data for representing apparent detail on the surface of simpler polygonal meshes, but they each store that data in a different way.

![On the left, a height map for bump mapping a stone wall. On the right, a normal map for bump mapping a stone wall.](../uploads/Main/BumpMapHeightMapNormalMapComparison.jpg)

On the left, a height map for bump mapping a stone wall. On the right, a normal map for bump mapping a stone wall.

Above, on the left, you can see a height map used for bump mapping a stone wall. A height map is a simple black and white texture, where each pixel represents the amount that point on the surface should appear to be raised. The whiter the pixel colour, the higher the area appears to be raised.

A normal map is an RGB texture, where each pixel represents the *difference in direction* the surface should appear to be facing, relative to its un-modified surface normal. These textures tend to have a bluey-purple tinge, because of the way the vector is stored in the RGB values.

Modern real-time 3D graphics hardware rely on Normal Maps, because they contain the vectors required to modify how light should appear to bounce of the surface. Unity can also accept Height Maps for bump mapping, but they must be converted to Normal Maps on import in order to use them.

## Colors in a normal map

Understanding this is not vital for using normal maps! It’s ok to skip this paragraph. However if you really want to know: The RGB colour values are used to store the X,Y,Z direction of the vector, with Z being “up” (contrary to Unity’s usual convention of using Y as “up”). In addition, the values in the texture are treated as having been halved, with 0.5 added. This allows vectors of all directions to be stored. Therefore to convert an RGB colour to a vector direction, you must multiply by two, then subtract 1. For example, an RGB value of (0.5, 0.5, 1) or #8080FF in hex results in a vector of (0,0,1) which is “up” for the purposes of normal-mapping - and represents no change to the surface of the model. This is the colour you see in the flat areas of the “example” normal map earlier on this page.

![A normal map using only #8080FF, which translates to a normal vector of 0,0,1 or straight up. This applies no modification to the surface normal of the polygon, and therefore produces no change to the lighting. Any pixels which are different to this colour results in a vectors that point in a different direction - which therefore modify the angle that is used to calculate light bounce at that point.](../uploads/Main/BumpMapFlatColour.png)

A normal map using only #8080FF, which translates to a normal vector of 0,0,1 or “straight up”. This applies no modification to the surface normal of the polygon, and therefore produces no change to the lighting. Any pixels which are different to this colour results in a vectors that point in a different direction - which therefore modify the angle that is used to calculate light bounce at that point.

A value of (0.43, 0.91, 0.80) gives a vector of (–0.14, 0.82, 0.6), which is quite a steep modification to the surface. Colours like this can be seen in the bright cyan areas of the stone wall normal map at the top of some of the stone edges. The result is that these edges catch the light at a very different angle to the flatter faces of the stones.

![The bright cyan areas in the normalmap for these stones show a steep modification to the polygons surface normals at the top edge of each stone, causing them to catch the light at the correct angle.](../uploads/Main/BumpMapColourMapStoneWallExample.jpg)

The bright cyan areas in the normalmap for these stones show a steep modification to the polygon’s surface normals at the top edge of each stone, causing them to catch the light at the correct angle.

Normal maps

![A stone wall with no bumpmap effect. The edges and facets of the rock do not catch the directional sun light in the scene.](../uploads/Main/BumpMapStoneExampleNoBumps.jpg)

A stone wall with no bumpmap effect. The edges and facets of the rock do not catch the directional sun light in the scene.

![The same stone wall with bumpmapping applied. The edges of the stones facing the sun reflect the directional sun light very differently to the faces of the stones, and the edges facing away.](../uploads/Main/BumpMapStoneExampleDay.jpg)

The same stone wall with bumpmapping applied. The edges of the stones facing the sun reflect the directional sun light very differently to the faces of the stones, and the edges facing away.

![The same bumpmapped stone wall, in a different lighting scenario. A point light torch illuminates the stones. Each pixel of the stone wall is lit according to how the light hits the angle of the base model (the polygon), adjusted by the vectors in the normal maps. Therefore pixels facing the light are bright, and pixels facing away from the light are darker, or in shadow.](../uploads/Main/BumpMapStoneExampleNightTorch.jpg)

The same bumpmapped stone wall, in a different lighting scenario. A point light torch illuminates the stones. Each pixel of the stone wall is lit according to how the light hits the angle of the base model (the polygon), adjusted by the vectors in the normal maps. Therefore pixels facing the light are bright, and pixels facing away from the light are darker, or in shadow.

## Limitations

* Unity does not support normal map scaling on mobile platforms.