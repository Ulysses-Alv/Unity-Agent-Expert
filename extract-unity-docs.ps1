$files = @(
'SL-Shader.html','SL-Properties.html','SL-SubShader.html','SL-Pass.html','SL-Commands.html',
'SL-SurfaceShaders.html','SL-SurfaceShaderExamples.html','SL-SurfaceShaderLighting.html',
'SL-SurfaceShaderOptimize.html','SL-SurfaceShaderTessellation.html','SL-SurfaceShader-create.html',
'SL-BuiltinFunctions.html','SL-BuiltinIncludes.html','SL-ShadingLanguage.html',
'SL-PropertiesInPrograms.html','SL-VertexProgramInputs.html','SL-HLSLSemantics.html',
'SL-PragmaDirectives.html','SL-MultipleProgramVariants.html','SL-ShaderCompileTargets.html',
'SL-ShaderCompilationAPIs.html','SL-RenderPipeline.html','SL-UnityShaderVariables.html',
'SL-SamplerStates.html','SL-SetTexture.html','SL-Cull.html','SL-ZTest.html','SL-ZWrite.html',
'SL-Blend.html','SL-BlendOp.html','SL-ColorMask.html','SL-Stencil.html','SL-Fog.html',
'SL-AlphaTest.html','SL-Offset.html','SL-Fallback.html','SL-GrabPass.html','SL-UsePass.html',
'SL-ShaderLOD.html','SL-Use16BitPrecisionInShaders.html','SL-GLSLShaderPrograms.html',
'SL-PlatformDifferences.html','shader-keywords.html','shader-keywords-default.html',
'shader-keywords-scripts.html','shader-keywords-material-inspector.html',
'shader-variant-collections.html','shader-variant-stripping.html','shader-variants.html',
'shader-branching-api.html','shader-branching-pass.html','shader-how-many-variants.html',
'shader-compilation.html','shader-prewarm.html','shader-pso-introduction.html',
'shader-pso-example.html','shader-pso-trace.html','shader-objects.html','shader-writing.html',
'shader-introduction.html','MaterialsAccessingViaScript.html','create-material.html',
'material-properties-texture-properties.html','renderer-shader-user-value.html',
'renderer-shader-user-value-set-and-use.html','renderer-shader-user-value-sample.html',
'shader-StandardShader.html','StandardShaderMaterialParameters.html',
'StandardShaderMaterialParameterAlbedoColor.html','StandardShaderMaterialParameterMetallic.html',
'StandardShaderMaterialParameterNormalMap.html','StandardShaderMaterialParameterEmission.html',
'StandardShaderMaterialParameterOcclusionMap.html','StandardShaderMaterialParameterRenderingMode.html',
'StandardShaderTransparency.html','StandardShaderFresnel.html','StandardShaderMakeYourOwn.html',
'StandardShaderMaterialCharts.html','shader-graph.html','render-pipelines.html',
'render-pipelines-overview.html','render-pipelines-set-up.html',
'render-pipelines-feature-comparison.html','choose-a-render-pipeline.html',
'scriptable-render-pipeline-introduction.html','srp-setting-render-pipeline-asset.html',
'built-in-render-pipeline.html','built-in-rendering-order.html','built-in-rendering-paths.html',
'RenderingPaths.html','built-in-materials-and-shaders.html','DrawCallBatching.html',
'DrawCallBatching-Enable.html','DrawCallBatching-Properties.html','DrawCallBatching-SetUp.html',
'SRPBatcher.html','SRPBatcher-Enable.html','SRPBatcher-Incompatible.html','SRPBatcher-Materials.html',
'SRPBatcher-Profile.html','graphics-performance-profiling.html','optimizing-draw-calls.html',
'graphics-tiers.html','graphics-tiers-customize.html','color-spaces.html','gamma-color-space.html',
'linear-color-space.html','differences-linear-gamma-color-space.html','linear-textures.html',
'lighting-birp.html','lighting-light-sources.html','lighting-light-components.html',
'lighting-ambient-light.html','lighting-emissive-materials.html','LightingOverview.html',
'direct-and-indirect-lighting.html','hdr-birp.html','shadows-in-birp.html',
'shadow-configuration.html','shadow-mapping.html','shadow-realtime.html','shadow-cascades.html',
'shadows-optimization.html','ShadowPerformance.html','PerPixelLights.html','Lightmapping.html',
'Lightmapping-bake.html','Lightmapping-configure.html','Lightmapping-preview.html',
'Lightmapping-troubleshooting.html','LightmappingDirectional.html','Lightmappers.html',
'Lightmaps-reference.html','LightmapParameters.html','LightingGiUvs.html','LightProbes.html',
'LightProbes-Reference.html','LightProbeProxyVolume-reference.html','ReflectionProbes.html',
'ReflectionProbes-set-gameobjects.html','RefProbeTypes.html','RefProbePerformance.html',
'sky.html','skyboxes-using.html','skybox-shaders.html','skybox-material-reference.html',
'ImportingTextures.html','class-Texture2DArray.html','class-Texture3D.html','class-Cubemap.html',
'class-CubemapArray.html','class-RenderTexture.html','class-CustomRenderTexture.html',
'SparseTextures.html','class-TextureImporter.html','mesh.html','mesh-vertex-data.html',
'mesh-index-data.html','mesh-topology-data.html','AnatomyofaMesh.html','creating-meshes.html',
'create-mesh.html','combining-meshes.html','gpu-instancing-birp.html',
'gpu-instancing-birp-shader-modifications.html','gpu-instancing-enable.html',
'gpu-instancing-per-instance-properties.html','gpu-instancing-samples.html',
'gpu-instancing-shader.html','gpu-instancing-strip-variants.html',
'gpu-instancing-surface-shader-example.html','gpu-instancing-vertex-fragment-shader-example.html',
'GPUInstancing.html','dots-instancing-shaders.html','dots-instancing-shaders-access.html',
'dots-instancing-shaders-declare.html','dots-instancing-shaders-macros.html',
'dots-instancing-shaders-per-instance.html','dots-instancing-shaders-samples.html',
'SinglePassInstancing.html','batch-renderer-group.html','batch-renderer-group-getting-started.html',
'batch-renderer-group-writing-shaders.html','Cameras.html','CamerasOverview.html','cameras-birp.html',
'cameras-overlay.html','MultipleCameras.html','CameraView.html','CameraRays.html','CameraOutput.html',
'PhysicalCameras.html','PhysicalCameras-GateFit.html','control-camera.html','class-Camera.html',
'LevelOfDetail.html','lod-group-configure.html','GraphicsAPIs.html','configure-graphicsAPIs.html',
'OpenGLCoreDetails.html','Metal.html','metal-optimize.html','metal-debug.html',
'AsynchronousShaderCompilation.html','AsynchronousShaderCompilation-introduction.html',
'DebuggingShadersWithPIX.html','RenderDocIntegration.html',
'post-processing-and-full-screen-effects.html','PostProcessingOverview.html','decals-birp.html',
'Cookies.html','lens-flare-introduction.html','material-validator-introduction.html',
'material-validator-validate.html','MaterialValidator.html','materialvariant-concept.html',
'materialvariant-properties.html','materialvariant-tasks.html','materials-introduction.html',
'materials-troubleshooting.html','shader-troubleshooting.html','shader-debugging.html',
'shader-memory.html','shader-Performance.html','shader-prevent-stutter.html',
'shader-reduce-stalling.html','shader-loading.html','shader-error.html','shader-variant-limit.html',
'shader-conditionals-choose-a-type.html','shader-branching-built-in-macros.html',
'shader-branching-platform.html','shader-branching-shader-model.html',
'shader-branching-unity-version.html','shader-pso-trace-warming.html','shader-prewarm-other.html',
'shader-reducing.html','shader-built-in.html','shader-built-in-birp.html',
'shader-built-in-configure-properties.html','SL-Reference.html','SL-AlphaToMask.html',
'SL-ZClip.html','SL-Name.html','SL-Material.html','SL-ShaderPerformance.html',
'SL-ShaderReplacement.html','SL-CustomEditor.html','SL-SubShaderTags.html','SL-PassTags.html',
'SL-BindChannels.html','SL-Conservative.html','SL-Other.html','SL-SurfaceShaders-output.html',
'SL-SurfaceShaderExamples-CustomData.html','SL-SurfaceShaderExamples-Decals.html',
'SL-SurfaceShaderExamples-FinalColor.html','SL-SurfaceShaderExamples-GlobalIllumination.html',
'SL-SurfaceShaderExamples-NormalMapping.html','SL-SurfaceShaderExamples-Reflection.html',
'SL-SurfaceShaderExamples-ToonRamp.html','SL-SurfaceShaderExamples-VertexModifier.html',
'SL-SurfaceShaderExamples-WrappedDiffuse.html','gpu-instancing-troubleshoot.html',
'dots-instancing-shaders-best-practice.html','dots-instancing-shaders-constant.html',
'dots-instancing-shaders-functions.html','dots-instancing-shaders-support.html',
'dots-instancing-shaders-unity-dots-instanced-prop.html',
'batch-renderer-group-creating-a-renderer.html','batch-renderer-group-creating-batches.html',
'batch-renderer-group-creating-draw-commands.html','batch-renderer-group-how.html',
'batch-renderer-group-initializing.html','batch-renderer-group-registering-meshes-and-materials.html',
'loading-texture-mesh-data-asynchronously.html','LoadingTextureandMeshData.html',
'LoadingTextureandMeshData-introduction.html','LoadingTextureandMeshData-make-compatible.html',
'identify-mesh-upload-pipeline.html','configure-asynchronous-upload-pipeline.html',
'configure-mesh-compression.html','configure-vertex-compression.html',
'compressing-mesh-data-optimization.html','mesh-data-deformable-meshes.html',
'mesh-colliders.html','mesh-colliders-introduction.html','prepare-mesh-for-mesh-collider.html',
'get-started-with-meshes.html','mesh-components-reference.html','mesh-select-mesh-asset.html',
'mesh-introduction.html','Lightmapping-baked-tags.html','Lightmapping-baking-before-runtime.html',
'Lightmapping-landing.html','LightmappingDirectional.html','Lightmap-data-landing.html',
'LightmapSnapshot.html','Lightmaps-TechnicalInformation.html','LightingGiUvs-GeneratingLightmappingUVs.html',
'LightingGiUvs-landing.html','LightingGiUvs-Reference.html','LightingGiUvs-visualizing.html',
'LightProbeProxyVolume-landing.html','LightProbes-landing.html','LightProbes-MeshRenderer.html',
'LightProbes-Moving.html','LightProbes-MovingObjects.html','LightProbes-Placing-Scripting.html',
'LightProbes-TechnicalInformation.html','light-probes-and-scene-loading.html',
'light-probes-troubleshooting.html','LightProbeProxyVolume-add.html',
'LightProbeProxyVolume-configure.html','LightProbeProxyVolume-Shader.html',
'realtime-gi-using-enlighten.html','realtime-gi-using-enlighten-landing.html',
'realtime-gi-using-enlighten-optimize.html','realtime-gi-using-enlighten-use.html',
'LODRealtimeGI.html','global-illumination-configure.html','GICache.html','GIVis.html',
'bakingmultiplescenes.html','configure-lightmapping-settings.html',
'configure-with-lightmap-parameters-asset.html','lighting-window.html','LightingExplorer.html',
'LightingExplorer-landing.html','LightExplorerExtension.html','GPUProgressiveLightmapper.html',
'progressive-lightmapper.html','ReflectionProbes-EnableReflectionsOfReflections.html',
'ReflectionProbes-set-gameobjects-use.html','reflections-landing.html','AdvancedRefProbe.html',
'blend-reflection-probes-birp.html','sky-landing.html','2d-images-lod.html',
'class-Texture2DArray-create.html','class-Texture2DArray-import.html',
'class-Texture2DArray-render-target.html','class-Texture2DArray-use-in-shader.html',
'class-Texture3D-create.html','class-Texture3D-import.html','class-Texture3D-introduction.html',
'class-Texture3D-preview.html','class-Texture3D-reference.html','class-Texture3D-use-in-shader.html',
'class-Cubemap-create.html','class-Cubemap-introduction.html','class-CubemapArray-create.html',
'class-CubemapArray-use-in-shader.html','class-RenderTexture-create.html',
'class-CustomRenderTexture-configure.html','class-CustomRenderTexture-create.html',
'class-CustomRenderTexture-introduction.html','class-CustomRenderTexture-landing.html',
'class-CustomRenderTexture-write-shader.html','render-texture-landing.html',
'class-TextureImporter-type-and-shape.html','class-TextureImporter-type-specific.html',
'camera-background-birp.html','multiple-cameras-birp.html','MultipleCameras-landing.html',
'CameraTroubleshooting.html','CameraRays-cast.html','CameraRays-introduction.html',
'CameraRays-move.html','CameraOutput-introduction.html','CameraOutput-shader.html',
'CameraOutput-troubleshoot.html','PhysicalCameras-introduction.html','PhysicalCameras-LensShift.html',
'PhysicalCameras-GateFit-Configure.html','PhysicalCameras-GateFit-Landing.html',
'FrustumSizeAtDistance.html','ObliqueFrustum.html','MultiDisplay.html','2DRendererSorting.html',
'2d-urp-landing.html','lod-landing.html','allow-deny-d3d12-usage.html','allow-deny-vulkan-usage.html',
'configure-d3d12-graphics-jobs-mode.html','configure-graphics-jobs-mode.html',
'd3d12-device-filter-list-asset-reference.html','d3d12-graphics-jobs-configuration.html',
'create-d3d12-device-filtering-asset.html','create-vulkan-device-filtering-asset.html',
'introduction-d3d12-device-filtering-asset.html','introduction-vulkan-device-filtering-asset.html',
'AsynchronousShaderCompilation-detect.html','AsynchronousShaderCompilation-enable-or-disable.html',
'AsynchronousShaderCompilation-avoid-cyan-placeholder-shaders.html',
'post-processing-effect-availability-reference.html','introduction-decals-projection.html',
'Cookies-introduction.html','Cookies-apply.html','creating-cookies-built-in-render-pipeline.html',
'lens-flare-birp.html','configuring-flare-elements.html','create-configure-halo.html',
'Decals-apply.html','materialvariant-landingpage.html','materials-and-shaders.html',
'Materials.html','shader-StandardShader-landing.html','shader-variants-landing.html',
'shader-StandardParticleShaders.html','shader-StandardParticleShadersCreate.html',
'shader-StandardParticleShadersLanding.html','shader-shaderlab-code-blocks.html',
'shader-shaderlab-legacy.html','SL-landing.html','SL-Shader-object.html',
'SL-SubShader-object.html','SL-SubShader-pass.html','SL-MultipleProgramVariants-declare.html',
'SL-MultipleProgramVariants-make-conditionals.html','SL-MultipleProgramVariants-shortcuts.html',
'SL-CameraDepthTexture.html','SL-CameraDepthTexture-motionvectors.html',
'shader-NormalFamily.html','shader-NormalDiffuse.html','shader-NormalBumpedDiffuse.html',
'shader-NormalBumpedSpecular.html','shader-NormalSpecular.html','shader-NormalParallaxDiffuse.html',
'shader-NormalParallaxSpecular.html','shader-NormalVertexLit.html','shader-NormalDecal.html',
'shader-ReflectiveFamily.html','shader-ReflectiveDiffuse.html','shader-ReflectiveBumpedDiffuse.html',
'shader-ReflectiveBumpedSpecular.html','shader-ReflectiveSpecular.html',
'shader-ReflectiveParallaxDiffuse.html','shader-ReflectiveParallaxSpecular.html',
'shader-ReflectiveVertexLit.html','shader-ReflectiveBumpedVertexLit.html',
'shader-ReflectiveBumpedUnlit.html','shader-SelfIllumFamily.html','shader-SelfIllumDiffuse.html',
'shader-SelfIllumBumpedDiffuse.html','shader-SelfIllumBumpedSpecular.html',
'shader-SelfIllumSpecular.html','shader-SelfIllumParallaxDiffuse.html',
'shader-SelfIllumParallaxSpecular.html','shader-SelfIllumVertexLit.html',
'shader-TransparentFamily.html','shader-TransDiffuse.html','shader-TransBumpedDiffuse.html',
'shader-TransBumpedSpecular.html','shader-TransSpecular.html','shader-TransParallaxDiffuse.html',
'shader-TransParallaxSpecular.html','shader-TransVertexLit.html',
'shader-TransparentCutoutFamily.html','shader-TransCutDiffuse.html','shader-TransCutBumpedDiffuse.html',
'shader-TransCutBumpedSpecular.html','shader-TransCutSpecular.html','shader-TransCutVertexLit.html',
'shader-skybox-6sided.html','shader-skybox-cubemap.html','shader-skybox-panoramic.html',
'shader-skybox-procedural.html','shader-Autodesk-Interactive.html',
'built-in-shader-examples.html','built-in-shader-examples-unlit.html',
'built-in-shader-examples-single-color.html','built-in-shader-examples-vertex-data.html',
'built-in-shader-examples-mesh-normals.html','built-in-shader-examples-simple-diffuse-lighting.html',
'built-in-shader-examples-diffuse-lighting-with-ambient-light.html',
'built-in-shader-examples-environment-reflections.html','built-in-shader-examples-fog.html',
'built-in-shader-examples-receive-shadows.html','built-in-shader-examples-shadow-casting.html',
'built-in-shader-examples-reflections.html','built-in-shader-examples-tri-planar-texturing.html',
'built-in-shader-examples-checkerboard.html','example-particle-system-custom-vertex-streams-standard-shaders.html',
'example-particle-system-custom-vertex-streams-surface-shaders.html',
'example-particle-system-gpu-instancing-custom-shader.html',
'example-particle-system-gpu-instancing-custom-vertex-streams.html',
'example-particle-system-gpu-instancing-surface-shader.html',
'MetaPass.html','StandardShaderChangeProperties.html','StandardShaderMetallicVsSpecular.html',
'StandardShaderTextureMaps.html','StandardShaderMaterialParameterDetail.html',
'StandardShaderMaterialParameterHeightMap.html',
'StandardShaderMaterialParameterNormalMapImport.html',
'StandardShaderMaterialParameterNormalMapLanding.html',
'StandardShaderMaterialParameterNormalMapSurfaceNormals.html',
'renderer-shader-user-value-intro.html','writing-custom-shaders.html',
'writing-shader-change-properties.html','writing-shader-tags.html',
'writing-shader-create-subshader-object.html','writing-shader-create-shader-pass.html',
'writing-shaders-birp.html','Shaders.html','shaders-reference.html',
'Lighting.html','LightingInUnity.html','LightingGiUvs.html',
'ShadowPerformance.html','shadow-distance.html','shadow-resolution-birp.html',
'shadow-cascades-landing.html','shadow-cascades-use.html','shadow-cascades-performance.html',
'shadow-cascades-implementation-details.html','shadows-optimization.html',
'hdr-landing.html','hdr-color-picker-reference.html','introduction-hdr.html',
'lighting-mode.html','lighting-mode-landing.html','lighting-mode-runtime.html',
'lighting-reference.html','lighting-optimize-builtin.html','lighting-precomputed-data.html',
'lighting-configuration-workflow.html','choose-a-lighting-setup.html',
'built-in-graphics-quality-settings.html','graphics-color.html',
'gamma-textures-linear-color-space.html','disable-srgb-sampling-textures.html',
'linear-color-space-landing.html','color-spaces-landing.html','set-project-color-space.html',
'reduce-draw-calls-landing.html','reduce-draw-calls-landing-urp.html',
'optimizing-draw-calls-choose-method.html','graphics-performance-birp.html',
'graphics-performance-and-profiling-in-urp.html','profile-rendering.html',
'avoid-shader-duplication.html','SRPBatcher-landing.html','DrawCallBatching-landing.html',
'LightingGiUvs.html'
)

$baseDir = "C:\Program Files\Unity\Hub\Editor\6000.3.9f1\Editor\Data\Documentation\en\Manual"
$outFile = "D:\Projects\Skill creation\unity-docs-extracted.txt"

# Clear output file
Set-Content -Path $outFile -Value "" -Encoding UTF8

$count = 0
foreach ($f in $files) {
    $path = Join-Path $baseDir $f
    if (Test-Path $path) {
        $html = Get-Content $path -Raw -Encoding UTF8
        
        # Remove head section
        $html = $html -replace '(?s)<head>.*?</head>', ''
        # Remove script tags
        $html = $html -replace '(?s)<script[^>]*>.*?</script>', ''
        # Remove style tags  
        $html = $html -replace '(?s)<style[^>]*>.*?</style>', ''
        # Remove nav
        $html = $html -replace '(?s)<nav[^>]*>.*?</nav>', ''
        # Remove header wrapper
        $html = $html -replace '(?s)<div class="header-wrapper"[^>]*>.*?</div>\s*</div>\s*</div>', ''
        # Remove sidebar
        $html = $html -replace '(?s)<div id="sidebar"[^>]*>.*?</div>\s*</div>\s*</div>\s*</div>\s*</div>', ''
        # Remove footer
        $html = $html -replace '(?s)<div class="footer-wrapper"[^>]*>.*?</div>\s*</div>\s*</div>\s*</div>', ''
        # Remove nextprev
        $html = $html -replace '(?s)<div class="nextprev"[^>]*>.*?</div>', ''
        # Remove mb20
        $html = $html -replace '(?s)<div class="mb20"[^>]*>.*?</div>', ''
        # Remove tooltips spans
        $html = $html -replace '(?s)<span class="tooltip"[^>]*>(.*?)<span class="tooltiptext">.*?</span></span>', '$1'
        # Remove all remaining HTML tags
        $html = $html -replace '<[^>]+>', "`n"
        # Decode HTML entities
        $html = $html -replace '&nbsp;', ' '
        $html = $html -replace '&lt;', '<'
        $html = $html -replace '&gt;', '>'
        $html = $html -replace '&amp;', '&'
        $html = $html -replace '&quot;', '"'
        $html = $html -replace '&#39;', "'"
        # Clean up whitespace
        $html = $html -replace '\r?\n\s*\r?\n', "`n`n"
        $html = $html.Trim()
        
        if ($html.Length -gt 50) {
            Add-Content -Path $outFile -Value "`n========== FILE: $f ==========`n" -Encoding UTF8
            Add-Content -Path $outFile -Value $html -Encoding UTF8
            $count++
        }
    }
}

Write-Host "Done. Extracted $count out of $($files.Count) files."
Write-Host "Output: $outFile"
