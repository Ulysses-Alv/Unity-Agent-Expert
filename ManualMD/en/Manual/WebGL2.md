# WebGL2

WebGL is an API for rendering graphics in web browsers, which is based on the functionality of the OpenGL ES graphics library. **WebGL** 2.0 (WebGL2) is the version of WebGL that Unity supports. WebGL2 almost matches with the OpenGL ES 3.0 functionality.

WebGL2 is widely accepted by most browsers and is the default Web graphics API in Unity.

## Recommendations for WebGL2

There are best practice recommendations and restrictions for working with WebGL2:

* [Recommendation for camera clear](#recommendation-for-camera-clear).
* [Enable anti-aliasing](#enable-anti-aliasing).
* [Restrictions on Web shader code](#restrictions-on-web-shader-code).
* [Recommendations for both WebGL2 and WebGPU graphics APIs](web-graphics-apis-intro.md).

### Recommendation for camera clear

By default, Unity Web clears the drawing buffer after each frame, which means the content of the frame buffer clears regardless of the [Camera.clearFlags](https://docs.unity3d.com/ScriptReference/Camera-clearFlags.html) setting. However, you can change this behavior at instantiation time. To do this, set `webglContextAttributes.preserveDrawingBuffer` to `true` in the `index.html` file of your Web template.

**Note** : If you set any [WebGL context attributes](https://developer.mozilla.org/en-US/docs/Web/API/WebGLRenderingContext/getContextAttributes), you must also add a line to preserve the **Power Preference** [Player setting](class-PlayerSettingsWebGL.md#Publishing).

```
script.onload = () => {
  config['webglContextAttributes'] = {
    preserveDrawingBuffer: true, //Add this line to preserve the Camera.clearFlags setting
    powerPreference: {{{ WEBGL_POWER_PREFERENCE }}} //Add this line to preserve the Power Preference Player setting
  };
  createUnityInstance(canvas, config, (progress) => {
```

### Enable anti-aliasing

WebGL supports anti-aliasing on most (but not on all) combinations of browsers and GPUs. Anti-aliasing softens jagged edges in your **scene**.

To enable anti-aliasing:

1. Go to the Quality Settings (menu: **Edit** > **Project Settings**, then select the **Quality** category).
2. In the **Rendering** section, make sure **Anti Aliasing** isn’t set to **Disabled**.

For more information about anti-aliasing, refer to the [Anti-aliasing documentation](https://docs.unity3d.com/Packages/com.unity.postprocessing@3.4/manual/Anti-aliasing.html) and [Quality Settings](class-QualitySettings.md).

### Restrictions on Web shader code

The WebGL 2.0 specification imposes some limitations on OpenGL Shading Language (GLSL) **shader** code. This is mainly relevant if you write your own shaders.
Precision qualifiers
WebGL 2.0 requires you to specify the precision of all variables in the shader. You can use `highp`, `mediump`, or `lowp` to specify the precision of the variable. If you don’t specify the precision, the shader will use the default precision, which is `mediump`. You can also use precision to specify the precision of a block of variables.

## Recommendations for both WebGL2 and WebGPU

There are recommendations relevant to both web graphics APIs. For more information, refer to [Recommendations for both WebGL2 and WebGPU graphics APIs](web-graphics-apis-intro.md).

## Additional resources

* [Recommendations for both WebGL2 and WebGPU APIs](web-graphics-apis-intro.md)
* [WebGPU (Experimental)](WebGPU.md)