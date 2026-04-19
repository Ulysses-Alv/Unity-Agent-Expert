# GLSL in Unity

As well as writing [HLSL shader programs](writing-shader-writing-shader-programs-hlsl.md), you can also write raw OpenGL Shading Language (GLSL) **shader** programs in Unity. This is only supported on OpenGL Core and OpenGL ES platforms.

**Note:** This is not recommended or needed as part of a normal workflow; Unity cross-compiles your HLSL into optimized GLSL when needed. The use of raw GLSL is only recommended for testing, or for some special use cases.

## GLSL snippets

GLSL program snippets are written between `GLSLPROGRAM` and `ENDGLSL` keywords.

In GLSL, all shader function entry points have to be called `main()`. When Unity loads the GLSL shader, it loads the source once for the vertex program, with the `VERTEX` preprocessor define, and once more for the fragment program, with the `FRAGMENT` preprocessor define. So the way to separate vertex and fragment program parts in GLSL snippet is to surround them with `#ifdef VERTEX .. #endif` and `#ifdef FRAGMENT .. #endif`. Each GLSL snippet must contain both a vertex program and a fragment program.

Standard include files match those provided for Cg/HLSL shaders; they just have a `.glslinc` extension:

```
UnityCG.glslinc
```

Vertex shader inputs come from predefined GLSL variables (`gl_Vertex`, `gl_MultiTexCoord0`, …) or are user defined attributes. Usually only the tangent vector needs a user defined attribute:

```
attribute vec4 Tangent;
```

Data from vertex to fragment programs is passed through *varying* variables, for example:

```
varying vec3 lightDir; // vertex shader computes this, fragment shader uses this
```

### External OES textures

Unity does some preprocessing during Shader compilation; for example, `texture2D/texture2DProj` functions may be replaced to `texture/textureProj`, based on graphics API (GlES3, GLCore). Some extensions don’t support new convention, most notably [`GL_OES_EGL_image_external`](https://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image_external.txt).

If you want to sample external textures in GLSL shaders, use `textureExternal/textureProjExternal` calls instead of `texture2D/texture2DProj`.

Example:

```
gl_FragData[0] = textureExternal(_MainTex, uv);
```