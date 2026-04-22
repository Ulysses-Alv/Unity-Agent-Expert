# Particle System custom vertex streams in the Built-In Render Pipeline

If you are comfortable writing your own **Shaders**, you can use the [Renderer Module](../ScriptReference/ParticleSystemRenderer.md)’s **Custom Vertex Streams** feature to configure your **Particle** Systems to pass a wider range of data into your custom Shaders.

There are a number of built-in [data streams](../ScriptReference/ParticleSystemVertexStream.md) to choose from, such as velocity, size and center position. Aside from the ability to create powerful custom Shaders, these streams allow a number of more general benefits:

* Use the [Tangent](../ScriptReference/ParticleSystemVertexStream.Tangent.md) stream to support normal mapped particles.
* You can remove Color and then add the Tangent [UV2](../ScriptReference/ParticleSystemVertexStream.UV2.md) and [AnimBlend](../ScriptReference/ParticleSystemVertexStream.AnimBlend.md) streams to use the Standard Shader on particles.
* To easily perform linear texture blending of flipbooks, add the UV2 and AnimBlend streams, and attach the Particles/Anim Alpha Blended Shader.

There are also two completely custom per-particle data streams ([ParticleSystemVertexStreams.Custom1](../ScriptReference/ParticleSystemVertexStreams.Custom1.md) and [ParticleSystemVertexStreams.Custom2](../ScriptReference/ParticleSystemVertexStreams.Custom2.md)), which can be populated from script. Call [SetCustomParticleData](../ScriptReference/ParticleSystem.SetCustomParticleData.md) and [GetCustomParticleData](../ScriptReference/ParticleSystem.GetCustomParticleData.md) with your array of data to use them. There are two ways of using this:

* To drive custom behavior in **scripts** by attaching your own data to particles; for example, attaching a “health” value to each particle.
* To pass this data into a Shader by adding one of the two custom streams, in the same way you would send any other stream to your Shader (see [ParticleSystemRenderer.EnableVertexStreams](../ScriptReference/ParticleSystemRenderer.EnableVertexStreams.md)). To elaborate on the first example, maybe your custom health attribute could now also drive some kind of visual effect, as well as driving script-based game logic.

When adding vertex streams, Unity will provide you with some information in brackets, next to each item, to help you read the correct data in your shader:

* Position (POSITION.xyz)
* Normal (NORMAL.xyz)
* Color (COLOR.xyzw)
* UV (TEXCOORD0.xy)
* UV2 (TEXCOORD0.zw)
* AnimBlend (TEXCOORD1.x)

Each item in brackets corresponds to a **Vertex Shader** input, which you should specify in your Shader. Here is the correct input structure for this configuration.

```
            struct appdata_t {
                float4 vertex : POSITION;
                float3 normal : NORMAL;
                fixed4 color : COLOR;
                float4 texcoords : TEXCOORD0;
                float texcoordBlend : TEXCOORD1;
            };
```

Notice that both UV and UV2 are passed in different parts of TEXCOORD0, so we use a single declaration for both. To access each one in your shader, you would use the xy and zw swizzles. This allows you to pack your vertex data efficiently.

## Additional resources

* [Example of Particle System Standard Shader custom vertex streams](example-particle-system-custom-vertex-streams-standard-shaders.md)
* [Example of Particle System Surface Shader custom vertex streams](example-particle-system-custom-vertex-streams-surface-shaders.md)