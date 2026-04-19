# Adding a device

To add a new device to the Device Simulator, you create a device definition and a device overlay.

A device definition is a text file with the `.device` extension in your Unity project. It contains JSON that describes the properties of a device.

A device overlay is an image that contains the border of the device screen, together with notches, punchouts, and any other additions to the screen rectangle. You can optionally use it with a device definition to visualize how hardware elements obstruct the device screen, and to determine when touch inputs fail as a result.

**Note**: Unity includes built-in device definitions that represent common screen characteristics (such as notch styles, punch hole positions, and screen sizes). These are intended for testing layout and interaction patterns. For production testing, we recommend to always validate your application on actual target devices.

## Creating a device definition

A device definition is a JSON file that represents the device. It has both required properties and some optional properties. If a device definition file contains any errors, the errors appear in the **Inspector** when you select the file.

### Schema

| **Property** | **Required** | **Description** |
| --- | --- | --- |
| **friendlyName** | Yes | The name to display in the UI for this device. |
| **version** | Yes | Indicates the version of the device definition file. Currently, the version is `1`. |
| **screens** | Yes | A list of objects that each describe a screen to simulate the device for. This must contain at least one screen. For information about the schema of each screen object, see [screen](#screen). |
| **systemInfo** | Yes | An object that describes the capabilities of the device. The values in this object map to [SystemInfo](../ScriptReference/SystemInfo.md). For information about the schema of the systemInfo object, see [systemInfo](#systeminfo). |

#### screen

| **Property** | **Required** | **Description** |
| --- | --- | --- |
| **width** | Yes | The width, in **pixels**, of the screen. |
| **height** | Yes | The height, in pixels, of the screen. |
| **navigationBarHeight** | No | The height, in pixels, of the on-screen Android navigation bar that appears on some devices when not in fullscreen. |
| **dpi** | Yes | The dpi of the screen. |
| **orientations** | No | A list of objects that each describe an orientation the screen can simulate. If you don’t set a value for this property, the screen supports all orientations. For information about the schema of each orientation object, see [orientation](#orientation). |
| **presentation** | No | An object that describes the device overlay. For information about the schema of this object, see [presentation](#presentation). |

#### orientation

| **Properties** | **Required** | **Description** |
| --- | --- | --- |
| **orientation** | Yes | The screen orientation. The value of this property is a number that maps to the [ScreenOrientation](../ScriptReference/ScreenOrientation.md) enum. |
| **safeArea** | No | A [Rect](../ScriptReference/Rect.md) that determines the safe area of the screen. If you don’t set a value for this property, the simulator assumes the entire screen is safe. |
| **cutouts** | No | A list of [Rect](../ScriptReference/Rect.md)s that specify areas of the screen that aren’t functional for displaying content. |

#### presentation

| **Property** | **Required** | **Description** |
| --- | --- | --- |
| **overlayPath** | No | A relative path from the device definition file to an image to use as the device overlay. |
| **borderSize** | No | The distance, in pixels, from the overlay to where the screen begins. |

#### systeminfo

The properties in this object describe the capabilities and system information of the device. Since they describe the system information, many of them map to properties in [SystemInfo](../ScriptReference/Device.SystemInfo.md).

| **Property** | **Required** | **Description** |
| --- | --- | --- |
| **deviceModel** | No | See [Device.SystemInfo.deviceModel](../ScriptReference/Device.SystemInfo-deviceModel.md). |
| **deviceType** | No | See [Device.SystemInfo.deviceType](../ScriptReference/Device.SystemInfo-deviceType.md). |
| **operatingSystem** | Yes | See [Device.SystemInfo.operatingSystem](../ScriptReference/Device.SystemInfo-operatingSystem.md). |
| **operatingSystemFamily** | No | See [Device.SystemInfo.operatingSystemFamily](../ScriptReference/Device.SystemInfo-operatingSystemFamily.md). |
| **processorCount** | No | See [Device.SystemInfo.processorCount](../ScriptReference/Device.SystemInfo-processorCount.md). |
| **processorFrequency** | No | See [Device.SystemInfo.processorFrequency](../ScriptReference/Device.SystemInfo-processorFrequency.md). |
| **processorType** | No | See [Device.SystemInfo.processorType](../ScriptReference/Device.SystemInfo-processorType.md). |
| **supportsAccelerometer** | No | See [Device.SystemInfo.supportsAccelerometer](../ScriptReference/Device.SystemInfo-supportsAccelerometer.md). |
| **supportsAudio** | No | See [Device.SystemInfo.supportsAudio](../ScriptReference/Device.SystemInfo-supportsAudio.md). |
| **supportsGyroscope** | No | See [Device.SystemInfo.supportsGyroscope](../ScriptReference/Device.SystemInfo-supportsGyroscope.md). |
| **supportsLocationService** | No | See [Device.SystemInfo.supportsLocationService](../ScriptReference/Device.SystemInfo-supportsLocationService.md). |
| **supportsVibration** | No | See [Device.SystemInfo.supportsVibration](../ScriptReference/Device.SystemInfo-supportsVibration.md). |
| **systemMemorySize** | No | See [Device.SystemInfo.systemMemorySize](../ScriptReference/Device.SystemInfo-systemMemorySize.md). |
| **unsupportedIdentifier** | No | See [Device.SystemInfo.unsupportedIdentifier](../ScriptReference/Device.SystemInfo-unsupportedIdentifier.md). |
| **graphicsDependentData** | No | A list of objects that each describe graphics APIs that the device supports. For information about the schema of each object, see [graphicsDependentData](#graphicsdependentdata). |

#### graphicsDependentData

The properties in the object describe a graphics API that the device supports.

| **Property** | **Required** | **Description** |
| --- | --- | --- |
| **graphicsDeviceType** | Yes | See [Device.SystemInfo.graphicsDeviceType](../ScriptReference/Device.SystemInfo-graphicsDeviceType.md). |
| **graphicsMemorySize** | No | See [Device.SystemInfo.graphicsMemorySize](../ScriptReference/Device.SystemInfo-graphicsMemorySize.md). |
| **graphicsDeviceName** | No | See [Device.SystemInfo.graphicsDeviceName](../ScriptReference/Device.SystemInfo-graphicsDeviceName.md). |
| **graphicsDeviceVendor** | No | See [Device.SystemInfo.graphicsDeviceVendor](../ScriptReference/Device.SystemInfo-graphicsDeviceVendor.md). |
| **graphicsDeviceID** | No | See [Device.SystemInfo.graphicsDeviceID](../ScriptReference/Device.SystemInfo-graphicsDeviceID.md). |
| **graphicsDeviceVendorID** | No | See [Device.SystemInfo.graphicsDeviceVendorID](../ScriptReference/Device.SystemInfo-graphicsDeviceVendorID.md). |
| **graphicsUVStartsAtTop** | No | See [Device.SystemInfo.graphicsUVStartsAtTop](../ScriptReference/Device.SystemInfo-graphicsUVStartsAtTop.md). |
| **graphicsDeviceVersion** | No | See [Device.SystemInfo.graphicsDeviceVersion](../ScriptReference/Device.SystemInfo-graphicsDeviceVersion.md). |
| **graphicsShaderLevel** | No | See [Device.SystemInfo.graphicsShaderLevel](../ScriptReference/Device.SystemInfo-graphicsShaderLevel.md). |
| **graphicsMultiThreaded** | No | See [Device.SystemInfo.graphicsMultiThreaded](../ScriptReference/Device.SystemInfo-graphicsMultiThreaded.md). |
| **renderingThreadingMode** | No | See [Device.SystemInfo.renderingThreadingMode](../ScriptReference/Device.SystemInfo-renderingThreadingMode.md). |
| **hasHiddenSurfaceRemovalOnGPU** | No | See [Device.SystemInfo.hasHiddenSurfaceRemovalOnGPU](../ScriptReference/Device.SystemInfo-hasHiddenSurfaceRemovalOnGPU.md). |
| **hasDynamicUniformArrayIndexingInFragmentShaders** | No | See [Device.SystemInfo.hasDynamicUniformArrayIndexingInFragmentShaders](../ScriptReference/Device.SystemInfo-hasDynamicUniformArrayIndexingInFragmentShaders.md). |
| **supportsShadows** | No | See [Device.SystemInfo.supportsShadows](../ScriptReference/Device.SystemInfo-supportsShadows.md). |
| **supportsRawShadowDepthSampling** | No | See [Device.SystemInfo.supportsRawShadowDepthSampling](../ScriptReference/Device.SystemInfo-supportsRawShadowDepthSampling.md). |
| **supportsMotionVectors** | No | See [Device.SystemInfo.supportsMotionVectors](../ScriptReference/Device.SystemInfo-supportsMotionVectors.md). |
| **supports3DTextures** | No | See [Device.SystemInfo.supports3DTextures](../ScriptReference/Device.SystemInfo-supports3DTextures.md). |
| **supports2DArrayTextures** | No | See [Device.SystemInfo.supports2DArrayTextures](../ScriptReference/Device.SystemInfo-supports2DArrayTextures.md). |
| **supports3DRenderTextures** | No | See [Device.SystemInfo.supports3DRenderTextures](../ScriptReference/Device.SystemInfo-supports3DRenderTextures.md). |
| **supportsCubemapArrayTextures** | No | See [Device.SystemInfo.supportsCubemapArrayTextures](../ScriptReference/Device.SystemInfo-supportsCubemapArrayTextures.md). |
| **copyTextureSupport** | No | See [Device.SystemInfo.copyTextureSupport](../ScriptReference/Device.SystemInfo-copyTextureSupport.md). |
| **supportsComputeShaders** | No | See [Device.SystemInfo.supportsComputeShaders](../ScriptReference/Device.SystemInfo-supportsComputeShaders.md). |
| **supportsGeometryShaders** | No | See [Device.SystemInfo.supportsGeometryShaders](../ScriptReference/Device.SystemInfo-supportsGeometryShaders.md). |
| **supportsTessellationShaders** | No | See [Device.SystemInfo.supportsTessellationShaders](../ScriptReference/Device.SystemInfo-supportsTessellationShaders.md). |
| **supportsInstancing** | No | See [Device.SystemInfo.supportsInstancing](../ScriptReference/Device.SystemInfo-supportsInstancing.md). |
| **supportsHardwareQuadTopology** | No | See [Device.SystemInfo.supportsHardwareQuadTopology](../ScriptReference/Device.SystemInfo-supportsHardwareQuadTopology.md). |
| **supports32bitsIndexBuffer** | No | See [Device.SystemInfo.supports32bitsIndexBuffer](../ScriptReference/Device.SystemInfo-supports32bitsIndexBuffer.md). |
| **supportsSparseTextures** | No | See [Device.SystemInfo.supportsSparseTextures](../ScriptReference/Device.SystemInfo-supportsSparseTextures.md). |
| **supportedRenderTargetCount** | No | See [Device.SystemInfo.supportedRenderTargetCount](../ScriptReference/Device.SystemInfo-supportedRenderTargetCount.md). |
| **supportsSeparatedRenderTargetsBlend** | No | See [Device.SystemInfo.supportsSeparatedRenderTargetsBlend](../ScriptReference/Device.SystemInfo-supportsSeparatedRenderTargetsBlend.md). |
| **supportedRandomWriteTargetCount** | No | See [Device.SystemInfo.supportedRandomWriteTargetCount](../ScriptReference/Device.SystemInfo-supportedRandomWriteTargetCount.md). |
| **supportsMultisampledTextures** | No | See [Device.SystemInfo.supportsMultisampledTextures](../ScriptReference/Device.SystemInfo-supportsMultisampledTextures.md). |
| **supportsMultisampleAutoResolve** | No | See [Device.SystemInfo.supportsMultisampleAutoResolve](../ScriptReference/Device.SystemInfo-supportsMultisampleAutoResolve.md). |
| **supportsTextureWrapMirrorOnce** | No | See [Device.SystemInfo.supportsTextureWrapMirrorOnce](../ScriptReference/Device.SystemInfo-supportsTextureWrapMirrorOnce.md). |
| **usesReversedZBuffer** | No | See [Device.SystemInfo.usesReversedZBuffer](../ScriptReference/Device.SystemInfo-usesReversedZBuffer.md). |
| **npotSupport** | No | See [Device.SystemInfo.npotSupport](../ScriptReference/Device.SystemInfo-npotSupport.md). |
| **maxTextureSize** | No | See [Device.SystemInfo.maxTextureSize](../ScriptReference/Device.SystemInfo-maxTextureSize.md). |
| **maxCubemapSize** | No | See [Device.SystemInfo.maxCubemapSize](../ScriptReference/Device.SystemInfo-maxCubemapSize.md). |
| **maxComputeBufferInputsVertex** | No | See [Device.SystemInfo.maxComputeBufferInputsVertex](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsVertex.md). |
| **maxComputeBufferInputsFragment** | No | See [Device.SystemInfo.maxComputeBufferInputsFragment](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsFragment.md). |
| **maxComputeBufferInputsGeometry** | No | See [Device.SystemInfo.maxComputeBufferInputsGeometry](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsGeometry.md). |
| **maxComputeBufferInputsDomain** | No | See [Device.SystemInfo.maxComputeBufferInputsDomain](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsDomain.md). |
| **maxComputeBufferInputsHull** | No | See [Device.SystemInfo.maxComputeBufferInputsHull](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsHull.md). |
| **maxComputeBufferInputsCompute** | No | See [Device.SystemInfo.maxComputeBufferInputsCompute](../ScriptReference/Device.SystemInfo-maxComputeBufferInputsCompute.md). |
| **maxComputeWorkGroupSize** | No | See [Device.SystemInfo.maxComputeWorkGroupSize](../ScriptReference/Device.SystemInfo-maxComputeWorkGroupSize.md). |
| **maxComputeWorkGroupSizeX** | No | See [Device.SystemInfo.maxComputeWorkGroupSizeX](../ScriptReference/Device.SystemInfo-maxComputeWorkGroupSizeX.md). |
| **maxComputeWorkGroupSizeY** | No | See [Device.SystemInfo.maxComputeWorkGroupSizeY](../ScriptReference/Device.SystemInfo-maxComputeWorkGroupSizeY.md). |
| **maxComputeWorkGroupSizeZ** | No | See [Device.SystemInfo.maxComputeWorkGroupSizeZ](../ScriptReference/Device.SystemInfo-maxComputeWorkGroupSizeZ.md). |
| **supportsAsyncCompute** | No | See [Device.SystemInfo.supportsAsyncCompute](../ScriptReference/Device.SystemInfo-supportsAsyncCompute.md). |
| **supportsGraphicsFence** | No | See [Device.SystemInfo.supportsGraphicsFence](../ScriptReference/Device.SystemInfo-supportsGraphicsFence.md). |
| **supportsAsyncGPUReadback** | No | See [Device.SystemInfo.supportsAsyncGPUReadback](../ScriptReference/Device.SystemInfo-supportsAsyncGPUReadback.md). |
| **supportsParallelPSOCreation** | No | See [Device.SystemInfo.supportsParallelPSOCreation](../ScriptReference/Device.SystemInfo-supportsParallelPSOCreation.md). |
| **supportsRayTracing** | No | See [Device.SystemInfo.supportsRayTracing](../ScriptReference/Device.SystemInfo-supportsRayTracing.md). |
| **supportsRayTracingShaders** | No | See [Device.SystemInfo.supportsRayTracingShaders](../ScriptReference/Device.SystemInfo-supportsRayTracingShaders.md). |
| **supportsInlineRayTracing** | No | See [Device.SystemInfo.supportsInlineRayTracing](../ScriptReference/Device.SystemInfo-supportsInlineRayTracing.md). |
| **supportsSetConstantBuffer** | No | See [Device.SystemInfo.supportsSetConstantBuffer](../ScriptReference/Device.SystemInfo-supportsSetConstantBuffer.md). |
| **hasMipMaxLevel** | No | See [Device.SystemInfo.hasMipMaxLevel](../ScriptReference/Device.SystemInfo-hasMipMaxLevel.md). |
| **supportsMipStreaming** | No | See [Device.SystemInfo.supportsMipStreaming](../ScriptReference/Device.SystemInfo-supportsMipStreaming.md). |
| **usesLoadStoreActions** | No | See [Device.SystemInfo.usesLoadStoreActions](../ScriptReference/Device.SystemInfo-usesLoadStoreActions.md). |

#### Minimal device definition

The following device definition contains every required property and no optional properties. This is the minimum device definition you can have.

**Note**: This device definition doesn’t provide orientation data, so the simulator assumes the device supports all orientations and that the safe area covers the entire screen.

```
{
    "friendlyName": "Minimal Device",
    "version": 1,
    "screens": [
        {
            "width": 1080,
            "height": 1920,
            "dpi": 450.0
        }
    ],
    "systemInfo": {
        "operatingSystem": "Android"
    }
}
```

#### Complete device definition

The following device definition contains every required and optional property.

```
{
    "friendlyName": "Apple iPhone XR",
    "version": 1,
    "screens": [
        {
            "width": 828,
            "height": 1792,
            "navigationBarHeight": 0,
            "dpi": 326.0,
            "orientations": [
                {
                    "orientation": 1,
                    "safeArea": {
                        "serializedVersion": "2",
                        "x": 0.0,
                        "y": 68.0,
                        "width": 828.0,
                        "height": 1636.0
                    },
                    "cutouts": [
                        {
                            "serializedVersion": "2",
                            "x": 184.0,
                            "y": 1726.0,
                            "width": 460.0,
                            "height": 66.0
                        }
                    ]
                },
                {
                    "orientation": 3,
                    "safeArea": {
                        "serializedVersion": "2",
                        "x": 88.0,
                        "y": 42.0,
                        "width": 1616.0,
                        "height": 786.0
                    },
                    "cutouts": [
                        {
                            "serializedVersion": "2",
                            "x": 0.0,
                            "y": 184.0,
                            "width": 66.0,
                            "height": 460.0
                        }
                    ]
                },
                {
                    "orientation": 4,
                    "safeArea": {
                        "serializedVersion": "2",
                        "x": 88.0,
                        "y": 42.0,
                        "width": 1616.0,
                        "height": 786.0
                    },
                    "cutouts": [
                        {
                            "serializedVersion": "2",
                            "x": 1726.0,
                            "y": 184.0,
                            "width": 66.0,
                            "height": 460.0
                        }
                    ]
                }
            ],
            "presentation": {
                "overlayPath": "Apple iPhone 11_Overlay.png",
                "borderSize": {
                    "x": 51.0,
                    "y": 51.0,
                    "z": 51.0,
                    "w": 51.0
                }
            }
        }
    ],
    "systemInfo": {
        "deviceModel": "iPhone11,8",
        "deviceType": 1,
        "operatingSystem": "iOS 12.0",
        "operatingSystemFamily": 0,
        "processorCount": 6,
        "processorFrequency": 0,
        "processorType": "arm64e",
        "supportsAccelerometer": true,
        "supportsAudio": true,
        "supportsGyroscope": true,
        "supportsLocationService": true,
        "supportsVibration": true,
        "systemMemorySize": 2813,
        "unsupportedIdentifier": "n/a",
        "graphicsDependentData": [
            {
                "graphicsDeviceType": 16,
                "graphicsMemorySize": 1024,
                "graphicsDeviceName": "Apple A12 GPU",
                "graphicsDeviceVendor": "Apple",
                "graphicsDeviceID": 0,
                "graphicsDeviceVendorID": 0,
                "graphicsUVStartsAtTop": true,
                "graphicsDeviceVersion": "Metal",
                "graphicsShaderLevel": 50,
                "graphicsMultiThreaded": true,
                "renderingThreadingMode": 0,
                "hasHiddenSurfaceRemovalOnGPU": true,
                "hasDynamicUniformArrayIndexingInFragmentShaders": true,
                "supportsShadows": true,
                "supportsRawShadowDepthSampling": true,
                "supportsMotionVectors": true,
                "supports3DTextures": true,
                "supports2DArrayTextures": true,
                "supports3DRenderTextures": true,
                "supportsCubemapArrayTextures": true,
                "copyTextureSupport": 31,
                "supportsComputeShaders": true,
                "supportsGeometryShaders": false,
                "supportsTessellationShaders": true,
                "supportsInstancing": true,
                "supportsHardwareQuadTopology": false,
                "supports32bitsIndexBuffer": true,
                "supportsSparseTextures": false,
                "supportedRenderTargetCount": 8,
                "supportsSeparatedRenderTargetsBlend": true,
                "supportedRandomWriteTargetCount": 8,
                "supportsMultisampledTextures": 1,
                "supportsMultisampleAutoResolve": false,
                "supportsTextureWrapMirrorOnce": 0,
                "usesReversedZBuffer": true,
                "npotSupport": 2,
                "maxTextureSize": 16384,
                "maxCubemapSize": 16384,
                "maxComputeBufferInputsVertex": 8,
                "maxComputeBufferInputsFragment": 8,
                "maxComputeBufferInputsGeometry": 0,
                "maxComputeBufferInputsDomain": 8,
                "maxComputeBufferInputsHull": 8,
                "maxComputeBufferInputsCompute": 8,
                "maxComputeWorkGroupSize": 1024,
                "maxComputeWorkGroupSizeX": 1024,
                "maxComputeWorkGroupSizeY": 1024,
                "maxComputeWorkGroupSizeZ": 1024,
                "supportsAsyncCompute": false,
                "supportsGraphicsFence": true,
                "supportsAsyncGPUReadback": true,
                "supportsParallelPSOCreation": true,
                "supportsRayTracing": false,
                "supportsRayTracingShaders": false,
                "supportsInlineRayTracing": false,
                "supportsSetConstantBuffer": true,
                "hasMipMaxLevel": true,
                "supportsMipStreaming": true,
                "usesLoadStoreActions": true,
                "supportedTextureFormats": [1, 2, 3, 4, 5],
                "supportedRenderTextureFormats": [1, 2, 3, 4, 5],
                "ldrGraphicsFormat": 59,
                "hdrGraphicsFormat": 74
            }
        ]
    }
}
```

## Creating a device overlay

A device overlay is an image that contains the border of the device screen and other features such as notches, punchouts, and any other additions to the screen rectangle. You can optionally use it with a device definition to visualize how hardware elements obstruct the device screen, and to determine when touch inputs fail as a result.

The Device Simulator interprets transparent pixels as areas of the screen you can tap, and opaque pixels of any other color as areas that the hardware obstructs. The texture itself can be any shape.

The following examples show device overlays for two iPhone models.

![iPhone 8 overlay, displaying Unity’s default skybox filling the area of the iPhone 8 screen where you can tap](../uploads/Main/device-simulator-overlay-iphone8.png)

iPhone 8 overlay, displaying Unity’s default skybox filling the area of the iPhone 8 screen where you can tap

![iPhone XS overlay, displaying Unity’s default skybox filling the area of the iPhone XS screen where you can tap](../uploads/Main/device-simulator-overlay-iphonexs.png)

iPhone XS overlay, displaying Unity’s default skybox filling the area of the iPhone XS screen where you can tap

**Note**: To mimic what you see when you use a device overlay, these examples display Unity’s default **skybox** in the area of the screen where you can tap. In a real device overlay, these pixels should be transparent.

### Using a device overlay

After you create a device overlay texture, to use it with a device definition you must first import the device overlay texture file into your project.

**Note**: When the Device Simulator loads a device overlay texture, it attempts to enable **Read/Write** for it. If this isn’t possible, the Device Simulator displays the texture but can’t use the texture to mask input. This means that if you click on notches and other areas of the screen that the device overlay should mask, the Device Simulator detects input. To ensure this doesn’t happen, when you import the device overlay texture, enable **Read/Write** in the Texture Import Settings window.

When the device overlay texture is in your project, open the device definition file and, in the object that defines a screen the device supports, add the [presentation](#presentation) property. Here, set the path to the image file (`overlayPath`) and the size of the borders (`borderSize`). For an example of how to do this, see the following device definition file:

```
{
    "friendlyName": "Minimal Device with Overlay",
    "version": 1,
    "screens": [
        {
            "width": 1080,
            "height": 1920,
            "dpi": 450.0,
            "presentation": {
                "overlayPath": "Overlays/MinimalDeviceOverlay.png",
                "borderSize": {
                    "x": 51.0,
                    "y": 51.0,
                    "z": 51.0,
                    "w": 130.0
                }
            }
        }
    ],
    "systemInfo": {
        "operatingSystem": "Android"
    }
}
```

**Note**: The path to the device overlay texture file can be relative to the device definition file, or relative to the directory that contains the **Assets** or **Packages** directory in your Unity project. For example, if the device definition file is in the **Assets/Devices** directory and the device overlay file is in the **Assets/Devices/Overlays** directory, the following file paths are both valid:

* Relative to the device definition file: **Overlays/MinimalDeviceOverlay.png**
* Relative to the directory that contains the **Assets** directory: **Assets/Devices/Overlays/MinimalDeviceOverlay.png**