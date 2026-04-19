# ShaderLab legacy functionality reference

ShaderLab includes a number of “fixed-function style” commands that allow you to write **shaders** without using any HLSL code.

**Note**: These commands are legacy, and are documented for backwards compatibility only. If your **Shader object** source file includes HLSL code, Unity ignores these commands completely. If your shader source file does not include HLSL code, Unity compiles these commands into regular shader programs on import.

* [Legacy fog](SL-Fog.md)
* [Legacy lighting](SL-Material.md)
* [Legacy alpha testing](SL-AlphaTest.md)
* [Legacy texture combining](SL-SetTexture.md)
* [Legacy vertex data channel mapping](SL-BindChannels.md)