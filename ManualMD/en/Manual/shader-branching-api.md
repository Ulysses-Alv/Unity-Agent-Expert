# Branch based on platform or graphics API

To create optimized **shader** code for a platform, branch in your shader code using one of the preprocessor macros Unity defines.

```
#if defined(SHADER_API_D3D11)
    // Shader code for when the target graphics API is DirectX 11
#endif
```

## Branch based on the build target

Unity defines the following macros when you make the corresponding platform the target for your build.

| **Macro** | **Platform** |
| --- | --- |
| `UNITY_PLATFORM_ANDROID` | Android devices |
| `UNITY_PLATFORM_EMBEDDED_LINUX` | Embedded Linux devices |
| `UNITY_PLATFORM_GAMECORE_XBOXONE` | Xbox One Game Core |
| `UNITY_PLATFORM_IOS` | iOS devices |
| `UNITY_PLATFORM_LINUX` | Linux devices |
| `UNITY_PLATFORM_OSX` | macOS devices |
| `UNITY_PLATFORM_PS4` | PlayStation®4 |
| `UNITY_PLATFORM_PS5` | PlayStation®5 |
| `UNITY_PLATFORM_QNX` | QNX devices |
| `UNITY_PLATFORM_SWITCH` | Nintendo Switch™ |
| `UNITY_PLATFORM_SWITCH2` | Nintendo Switch™ 2 |
| `UNITY_PLATFORM_TVOS` | tvOS devices |
| `UNITY_PLATFORM_UWP` | Universal Windows Platform (UWP) devices |
| `UNITY_PLATFORM_VISIONOS` | visionOS devices |
| `UNITY_PLATFORM_WEBGL` | **WebGL** devices |
| `UNITY_PLATFORM_WINDOWS` | Windows devices |
| `UNITY_PLATFORM_XBOXONE` | Xbox One |

### Branch within a platform branch

To branch again within platform-specific branching code, declare a shader keyword after you test for the platform, then branch using the keyword. For example:

```
// Test if the build target is Android
#if defined(UNITY_PLATFORM_ANDROID)

    // Add a keyword only if the build target is Android
    #pragma shader_feature _ ANDROID_REFLECTIONS

#endif

...

fixed4 frag (v2f i) : SV_Target
    {
        // Test if the build target is Android
        #if defined(UNITY_PLATFORM_ANDROID)

            // Test if your custom keyword is enabled
            #if ANDROID_REFLECTIONS
                // Shader code for Android with reflections
            #else
                // Shader code for Android without reflections
            #endif

        #endif
    ...
    }
```

For more information, refer to [Changing how shaders work using keywords](SL-MultipleProgramVariants.md).

## Branch based on the target graphics API

Unity defines the following macros when Unity compiles shaders for the corresponding graphics API.

| **Macro** | **Graphics API** |
| --- | --- |
| `SHADER_API_D3D11_9X` | DirectX 11 feature level 9, for Universal Windows Platform (UWP) devices |
| `SHADER_API_D3D11` | DirectX 11 |
| `SHADER_API_DESKTOP` | Graphics APIs used by Windows devices, macOS devices, Linux desktop platforms, and web devices |
| `SHADER_API_GLCORE` | OpenGL 3.x, OpenGL 4.x |
| `SHADER_API_GLES3` | OpenGL ES 3.x |
| `SHADER_API_METAL` | Metal on iOS or macOS |
| `SHADER_API_MOBILE` | Graphics APIs used by iOS and Android mobile devices, and tvOS devices |
| `SHADER_API_VULKAN` | Vulkan |
| `SHADER_API_WEBGPU` | WebGPU |
| `SHADER_TARGET_GLSL` | OpenGL and GLES platforms where Unity compiles HLSL to GLSL |

## Additional resources

* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL pragma target command reference](SL-Pragma-target.md)
* [HLSL pragma require command reference](SL-Pragma-require.md)