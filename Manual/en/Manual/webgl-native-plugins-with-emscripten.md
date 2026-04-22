# Web native plug-ins for Emscripten

Unity uses the [Emscripten compiler](https://emscripten.org/) toolchain to cross-compile the Unity runtime code written in C and C++ into [WebAssembly](https://webassembly.org/) (also called Wasm). Emscripten generates small code, is load-time and memory efficient, and aims to deliver close to native speed execution.

Unity uses [IL2CPP](./scripting-backends-il2cpp.md) to convert the .NET game code in your C# **scripts** to Wasm. IL2CPP takes .NET bytecode and converts it to corresponding C++ source files, which is then compiled using Emscripten to convert your scripts to Wasm.

## Targeted versions

Unity bundles the Emscripten compiler with the Web platform package, so that you don’t have to manually download it. The version of Emscripten that the Editor uses varies by Editor version:

| **Unity Version** | **Emscripten version** | **Supported native **plug-in** file formats** |
| --- | --- | --- |
| Unity 2023.2 and later | Emscripten 3.1.38-unity | .a, .bc |
| Unity 2022.2 and later | Emscripten 3.1.8-unity | .a, .bc |
| Unity 2021.2 and later | Emscripten 2.0.19.6-unity | .a, .bc |
| Unity 2019.2 to 2021.1 | 1.38.11-unity | .bc |
| Unity 2018.4 up until Unity 2019.2 | 1.37.40-unity | .bc |
| Unity 2018.2 up until Unity 2018.4 | 1.37.33-unity | .bc |

**Notes:**

* The following list provides expansion for the acronyms of file formats listed in the above table:

  + GNU archive file format (`.a`)
  + WebAssembly Object Files (`.o`)
  + LLVM Bitcode file format (`.bc`)
* With Emscripten versions before Emscripten 2.0, Unity recommends building `Bitcode` plug-ins (files of type `.bc`). From Emscripten 2.0 onwards, Unity recommends to build `Wasm Object File` plug-ins that are Wasm object files of type `.o`, bundled into GNU archive files of type `.a`.

You can determine the path to the exact version of Emscripten toolchain from the `emscripten-version.txt` located at:

```
C:\Program Files\Unity\Hub\Editor\<Editor version>\Editor\Data\PlaybackEngines\WebGLSupport\BuildTools\Emscripten\emscripten\emscripten-version.txt
```

## Compile native plug-ins with Emscripten

If you’re creating native C/C++ plug-ins for your Unity application, you can either bundle the C/C++ code with your project (see [Native plug-ins](plug-ins-native.md)) or build the native code to a plug-in archive ahead of time with the Emscripten compiler toolchain.

If you choose to build plug-in code in advance, you should use the Emscripten compiler toolchain. To ensure [LLVM binary format](https://llvm.org/) compatibility, the version of Emscripten that’s used to compile the plug-ins must match the version of Emscripten that Unity uses.

Typically, plug-ins are compiled to a static library archive format. For versions before Unity 2021.2 (Emscripten 2.0), LLVM Bitcode file format (`.bc`) is the preferred plug-in format.

From Unity 2021.2, the preferred plug-in format is GNU archive file format (`.a`) that contain WebAssembly Object Files (`.o`). Although, earlier LLVM Bitcode .bc plug-in files are still supported, using them might have a negative impact to compilation times.

**Note:** When migrating a native plug-in from one Unity version to a version that uses a different version of Emscripten, Unity recommends to recompile the Unity plug-in from source, because the LLVM compiler project doesn’t guarantee binary compatibility of build artifact files across compiler versions.

### Emscripten build flags

If compiling WebAssembly Object Files to use with WebAssembly 2023 or C/C++ multithreading, include the following Emscripten build flags. Without these flags, the code might not link properly with code built with Unity.

| **Target** | **Required Emscripten build flags** |
| --- | --- |
| WebAssembly 2023 | `-msimd128` `-fwasm-exceptions` `-mbulk-memory` `-mnontrapping-fptoint` `-msse4.2` `-sWASM_BIGINT` `-sSUPPORT_LONGJMP=wasm` |
| C/C++ native multithreading | `-pthread` |

## Additional resources:

* [Native plug-ins](plug-ins-native.md)