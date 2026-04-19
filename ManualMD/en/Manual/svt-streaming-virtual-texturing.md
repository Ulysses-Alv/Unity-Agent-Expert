# Streaming Virtual Texturing

Streaming Virtual Texturing (SVT) is a feature that reduces GPU memory usage and texture loading times when you have many high resolution textures in your **Scene**. It splits textures into tiles, and progressively uploads these tiles to GPU memory when they are needed.

SVT lets you set a fixed memory cost. For full texture quality, the required GPU cache size depends mostly on the frame resolution, and not the number or resolution of textures in the Scene. The more high resolution textures you have in your Scene, the more GPU memory you can save with SVT.

The workflow requires no additional import time, no additional build step, and no additional streaming files. You work with regular Unity Textures in the Unity Editor.