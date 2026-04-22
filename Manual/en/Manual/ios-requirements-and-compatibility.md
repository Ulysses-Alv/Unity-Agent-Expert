# iOS requirements and compatibility

Learn about the system requirements and compatibility information for iOS to make sure you’re aware of any limitations for developing a Unity application for this platform.

## iOS version support

Unity supports iOS 15 and above. For more information, refer to [`PlayerSettings.iOS-targetOSVersionString`](../ScriptReference/PlayerSettings.iOS-targetOSVersionString.md).

## Xcode version support

When developing for iOS, it’s recommended to use Xcode version 16 or later. For more information, refer to [Xcode](https://developer.apple.com/support/xcode/).

## Graphics API support

iOS devices support [Metal](https://developer.apple.com/documentation/metal). For more information, refer to [Metal Graphics API](Metal.md).

## Audio compression

Unity supports importing a variety of source format sound files. However, when importing these files (with the exception of tracker files), they are always re-encoded to the build target format. By default, this format is Vorbis, though this can be overridden per platform to other formats (ADPCM, MP3, etc.) if required. Vorbis playback provides better **compression** and quality for iOS compared with MP3 playback.

## ASTC instead of DXT texture compression

Unity iOS does not support DXT textures. Instead, ASTC, ETC2, and ETC **texture compression** are natively supported by iPhone/iPad devices. For more information about iOS **texture formats**, refer to [texture import settings](class-TextureImporter.md) and [texture compression format](texture-choose-format-by-platform.md#ios-and-tvos).

## iPadOS multitasking

iPadOS supports multitasking modes that allow users to run applications in split-screen or windowed layouts. The available modes vary by iPadOS version: Split View and Slide Over (iPadOS 9 to iPadOS 18), Stage Manager (iPadOS 16 and later), and Windowed Apps (iPadOS 26 and later). Unity applications support all multitasking modes and respond correctly when the available space changes. For more information, refer to [iPad multitasking modes](iphone-API.md#ipad-multitasking-modes).