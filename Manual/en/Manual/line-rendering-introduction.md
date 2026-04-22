# Rendering lines

The ****Line Renderer**** component takes an array of two or more points in 3D space, and draws a straight line between each one. You can use a Line Renderer to draw anything from a simple straight line to a complex spiral.

The line is always continuous; if you need to draw two or more completely separate lines, you should use multiple **GameObjects**, each with its own Line Renderer.

The Line Renderer does not render lines that have a width in **pixels**. It renders polygons that have a width in world units. The Line Renderer uses the same algorithm for line rendering as the [Trail Renderer](class-TrailRenderer.md).