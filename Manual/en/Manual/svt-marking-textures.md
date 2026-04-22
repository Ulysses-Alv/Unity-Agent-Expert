# Marking textures as “Virtual Texturing Only”

In the Unity Editor, you can mark textures that only Virtual Texturing uses. This optimizes memory usage and **Scene** loading times, and reduces the size of the build. To mark a texture as one that only Virtual Texturing uses, open its Texture Importer and enable **Virtual Texture Only**.

In practice, when you use Virtual Texturing, you should mark as many textures as possible as **Virtual Texture Only** to maximize these benefits. Note that if you mark a texture as **Virtual Texture Only**, you cannot use it with a regular Texture Sampler in your Project. In the Editor, it might seem like you can use it with a regular Texture Sampler because a low resolution preview appears, but an error occurs if you then build a Player.

## How it works

By default, when you load a Scene in the Unity Editor, the Editor loads all referenced textures into both CPU and GPU memory. The Unity Editor cannot stream textures into GPU memory from disk, but SVT can stream them from CPU memory. When you mark a texture as **Virtual Texture Only**, this means that on scene load the Editor loads it into CPU memory only, and then extracts tiles from that texture and streams them into GPU memory on demand.

In a Project with many high resolution textures, this significantly reduces both GPU memory usage and Scene loading times in the Unity Editor. For this reason, it’s good practice to mark as many textures as possible as **Virtual Texture Only**, to ensure that the Streaming Virtual Texturing (SVT) system only streams them from the CPU.

If you don’t mark a texture as **Virtual Texture Only**, and then use it with Virtual Texturing in the Player, the SVT system keeps a copy of the texture in both CPU and GPU memory.

## Standard texture artifacts

By default, Unity adds all textures sampled in **Shader** Graph to the build as standard texture artifacts. Additionally, Virtual Texturing imports all Texture Stack textures into a tiled streaming file, which Unity also includes in the build.

* If you mark a texture as **Virtual Texture Only**, Unity doesn’t include it as a standard texture artifact, and includes it only in the tiled streaming file.
* If you don’t mark a texture as **Virtual Texture Only**, Unity includes it twice in the build; once as a standard texture artifact, and once in the tiled streaming file.

Note that if you mark a texture as **Virtual Texture Only**, but use it as a standard texture in the Player, the Player build fails.