# CameraEvent and LightEvent events order reference for the Built-In Render Pipeline

## CameraEvent

The order of execution for CameraEvents depends on the [rendering path](RenderingPaths.md) that your Project uses.

### Deferred rendering path

* [BeforeGBuffer](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeGBuffer.html): Unity renders opaque geometry
* [AfterGBuffer](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterGBuffer.html): Unity resolves depth.
* [BeforeReflections](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeReflections.html): Unity renders default reflections, and **Reflection Probe** reflections.
* [AfterReflections](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterReflections.html): Unity copies reflections to the Emissive channel of the G-buffer.
* [BeforeLighting](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeLighting.html): Unity renders shadows. See [LightEvent order of execution](#LightEvent).
* [AfterLighting](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterLighting.html)
* [BeforeFinalPass](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeFinalPass.html): Unity processes the final pass.
* [AfterFinalPass](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterFinalPass.html)
* [BeforeForwardOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeForwardOpaque.html) (only called if there is opaque geometry that cannot be rendered using deferred): Unity renders opaque geometry that cannot be rendered with deferred rendering.
* [AfterForwardOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterForwardOpaque.html) (only called if there is opaque geometry that cannot be rendered using deferred)
* [BeforeSkybox](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeSkybox.html): Unity renders the **skybox**.
* [AfterSkybox](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterSkybox.html): Unity renders halos.
* [BeforeImageEffectsOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeImageEffectsOpaque.html): Unity applies opaque-only **post-processing** effects.
* [AfterImageEffectsOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterImageEffectsOpaque.html)
* [BeforeForwardAlpha](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeForwardAlpha.html): Unity renders transparent geometry, and UI Canvases with a Rendering Mode of **Screen Space - **Camera****.
* [AfterForwardAlpha](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterForwardAlpha.html): [BeforeHaloAndLensFlares](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeHaloAndLensFlares.html): Unity renders **lens flares**.
* [AfterHaloAndLensFlares](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterHaloAndLensFlares.html)
* [BeforeImageEffects](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeImageEffects.html): Unity applies post-processing effects.
* [AfterImageEffects](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterImageEffects.html)
* [AfterEverything](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterEverything.html): Unity renders UI Canvases with a Rendering Mode that is not **Screen Space - Camera**.

### Forward rendering path

* [BeforeDepthTexture](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeDepthTexture.html): Unity renders depth for opaque geometry.
* [AfterDepthTexture](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterDepthTexture.html)
* [BeforeDepthNormalsTexture](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeDepthNormalsTexture.html): Unity renders depth normals for opaque geometry.
* [AfterDepthNormalsTexture](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterDepthNormalsTexture.html): Unity renders shadows. See [LightEvent order of execution](#LightEvent).
* [BeforeForwardOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeForwardOpaque.html): Unity renders opaque geometry.
* [AfterForwardOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterForwardOpaque.html)
* [BeforeSkybox](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeSkybox.html): Unity renders the skybox.
* [AfterSkybox](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterSkybox.html): Unity renders halos.
* [BeforeImageEffectsOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeImageEffectsOpaque.html): Unity applies opaque-only post-processing effects.
* [AfterImageEffectsOpaque](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterImageEffectsOpaque.html)
* [BeforeForwardAlpha](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeForwardAlpha.html): Unity renders transparent geometry, and UI Canvases with a Rendering Mode of **Screen Space - Camera**.
* [AfterForwardAlpha](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterForwardAlpha.html)
* [BeforeHaloAndLensFlares](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeHaloAndLensFlares.html): Unity renders lens flares.
* [AfterHaloAndLensFlares](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterHaloAndLensFlares.html)
* [BeforeImageEffects](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.BeforeImageEffects.html): Unity applies post-processing effects.
* [AfterImageEffects](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterImageEffects.html)
* [AfterEverything](https://docs.unity3d.com/ScriptReference/Rendering.CameraEvent.AfterEverything.html): Unity renders UI Canvases with a Rendering Mode that is not **Screen Space - Camera**.

## LightEvent order of execution

During the “render shadows” stage above, for each shadow-casting Light, Unity performs these steps:

* [BeforeShadowMap](../ScriptReference/Rendering.LightEvent.BeforeShadowMap.md)
* [BeforeShadowMapPass](../ScriptReference/Rendering.LightEvent.BeforeShadowMapPass.md): Unity renders all shadow casters for the current Pass
* [AfterShadowMapPass](../ScriptReference/Rendering.LightEvent.AfterShadowMapPass.md): Unity repeats the last three steps, for each Pass
* [AfterShadowMap](../ScriptReference/Rendering.LightEvent.AfterShadowMap.md)
* [BeforeScreenSpaceMask](../ScriptReference/Rendering.LightEvent.BeforeScreenSpaceMask.md): Unity gathers the shadow map into a screen space buffer and performs filtering
* [AfterScreenSpaceMask](../ScriptReference/Rendering.LightEvent.AfterScreenSpaceMask.md)