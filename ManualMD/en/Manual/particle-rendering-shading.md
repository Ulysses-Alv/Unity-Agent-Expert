# Particle rendering and shading

Understand how the properties on the [Renderer](PartSysRendererModule.md) module interact to determine how a **particle**’s image or [Mesh](class-Mesh.md) is transformed, shaded and overdrawn by other particles.

## Render Mode

Use **Render Mode** to choose between several 2D **Billboard** graphic modes and a Mesh mode. 3D Meshes give particles extra authenticity when they represent solid **GameObjects**, such as rocks, and can also improve the sense of volume for clouds, fireballs and liquids.
Meshes must be read/write enabled to work in the **Particle System**’s **Mesh** Render Mode. When you assign meshes to a Particle System (using the **Meshes** list in the Inspector window), Unity automatically enables the read/write enabled setting for those meshes.

### Billboard render modes

When you use 2D billboard graphics, the different render modes can produce a variety of results that make them suitable for specific uses:

* Billboard mode is useful for particles that represent volumes that look similar from any direction (such as clouds).
* Horizontal Billboard mode is useful when the particles cover the ground (such as target indicators and magic spell effects) or when they are flat objects that fly or float parallel to the ground.
* Vertical Billboard mode keeps each particle upright and perpendicular to the XZ plane, but allows it to rotate around its y-axis. This can be helpful when you are using an orthographic **Camera** and want particle sizes to stay consistent.
* Stretched Billboard mode accentuates the apparent speed of particles in a similar way to the “stretch and squash” techniques used by traditional animators. Note that in Stretched Billboard mode, particles are aligned to face the Camera and also aligned to their velocity. This alignment occurs regardless of the Velocity Scale value - even if Velocity Scale is set to 0, particles in this mode still align to the velocity.

When you use Billboard render modes, you can use the Normal Direction to create spherical shading on the flat rectangular billboards. This can help create the illusion of 3D particles if you use a Material that applies lighting to your particles.

## Mesh Distribution: Non-uniform Random

![The Renderer module Inspector window, with a Mesh Weightings input highlighted](../uploads/Main/mesh-render-mode.png)

The Renderer module Inspector window, with a **Mesh Weightings** input highlighted

When **Mesh Distribution** is set to **Non-uniform Random**, you can customize how often Unity randomly assigns specific meshes to particles. To do this, you use the **Meshes** list and the **Mesh Weighting** field.

In the **Meshes** list, the field on the left contains the mesh you want the Particle System to use, and the field on the right (highlighted in the image above) contains the Mesh Weighting for that mesh. Use the Mesh Weighting field for each mesh to control how often Unity assigns that mesh to a particle, relative to every other mesh. This is set to 1 by default, which results in the likelihood that Unity assigns a mesh being equal.

To add a mesh to the Meshes list, select the Add (+) icon. To remove a mesh from the Meshes list, select the mesh, then select the Remove (-) icon.