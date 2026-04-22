# Modify GameActivity bridge code

GameActivity interacts with Unity via a bridge that you can modify to make changes and implement additional features. The code that makes up the bridge is written in C++ and during the build process, GameActivity builds it into a shared library called `libgame.so`.

You can’t modify bridge code within Unity itself; you must first [export your project](android-export-process.md). After you export your project, you can find the files that comprise the bridge code in `<exported_project_directory>/unityLibrary/src/main/cpp/GameActivity/`. Most of the code files in this directory contain utility code. The following table shows you the purpose of the most important bridge code files.

| **File** | **Purpose** |
| --- | --- |
| `UGAInput.cpp` | Input events: Here you can adjust or transform input data before GameActivity passes it to Unity. |
| `UGAApplication.cpp` | Lifecycle events: Here you can change how to handle events such as pausing, resuming, focusing, and unfocusing. This is the core of the code bridge. |
| `UGASoftKeyboard.cpp` | Touchscreen keyboard: Here you can change the implementation of the on-screen keyboard. The default implementation uses [GameTextInput](https://developer.android.com/games/agdk/add-support-for-text-input). |

During the project export process, Unity’s [incremental build pipeline](building-introduction.md#incremental-build-pipeline) might overwrite any changes you make in the exported project. If you want your changes to persist:

1. [Export your project](android-export-process.md).
2. Create a new directory that’s outside your Unity project. This new directory is your modified bridge code directory.
3. Copy the code files that you want to modify from the `<exported_project_directory>/unityLibrary/src/main/cpp/GameActivity/` directory into your modified bridge code directory.
4. In Unity, create a new C# script that uses [Android.IPostGenerateGradleAndroidProject](../ScriptReference/Android.IPostGenerateGradleAndroidProject.md) to copy the code files from your modified bridge code directory back into the `<exported_project_directory>/unityLibrary/src/main/cpp/GameActivity/` directory. When Unity builds your application, this code will overwrite the default bridge code files with your modified versions.
5. Make any bridge code modifications in the files in your modified bridge code directory.

## Expand GameActivity bridge code

You can add extra source files to the existing GameActivity bridge files which are then compiled together.

For example (You can find an example project [here](../uploads/Examples/AndroidExpandingGameActivity.zip) ):

1. Create C# script in Unity and name it **SendMessageReceiver.cs**.

   ```
   using UnityEngine;

   public class SendMessageReceiver : MonoBehaviour
   {
       public void SendMessageFromCpp(string message)
       {
           Debug.LogFormat(LogType.Log, LogOption.NoStacktrace, null, message);
       }
   }
   ```
2. Create a new **GameObject** and name it **SendMessageReceiver**.
3. Attach a **SendMessageReceiver** script to **SendMessageReceiver** GameObject.
4. Create `MyFile.cpp` in `\<unityproject\>/ExtraSourceFiles` directory.

   **Note**: Don’t place `.cpp` files in the **Assets** folder, because they will link into IL2CPP’s `libil2cpp.so` shared library and won’t compile.

   The following code calls `SendMessageFromCpp` on a GameObject called `SendMessageReceiver` and, passes `HelloFromBridge` as an extra parameter whenever you touch the phone screen.

   ```
   #include "UGAApplication.h"
   #include "game-activity/native_app_glue/android_native_app_glue.h"

   void UnityGameActivityPluginLoad(Unity::UnityApplication& application)
   {
       application.GetEvents().Register<Unity::UnityEventProcessInput>([](const Unity::UnityEventProcessInput& e)
       {
           auto inputBuffer = e.GetInputBuffer();

           if (inputBuffer->motionEventsCount != 0) {
               for (uint64_t i = 0; i < inputBuffer->motionEventsCount; ++i) {
                   GameActivityMotionEvent* motionEvent = &inputBuffer->motionEvents[i];
                   if (motionEvent->action == AKEY_EVENT_ACTION_DOWN)
                    e.GetApplication().SendMessage("SendMessageReceiver", "SendMessageFromCpp", "HelloFromBridge");
                }
           }
        });
   }
   ```
5. Place the following editor script **PostProcessor.cs** in the **Assets/Editor** folder:

   (It ensures that ‘ExtraSourceFiles/MyFile.cpp’ is copied to ‘unityLibrary/src/main/cpp/GameActivity/CustomFolder/MyFile.cpp’ in an incremental build friendly way. )

   ```
   using System;
   using UnityEditor.Android;
   using UnityEditor;
   using UnityEngine;

   public class PostProcessor : AndroidProjectFilesModifier
   {
       const string CustomSourceFileSrc = "ExtraSourceFiles/MyFile.cpp";
       const string CustomSourceFileDst = "unityLibrary/src/main/cpp/GameActivity/CustomFolder/MyFile.cpp";

       public override AndroidProjectFilesModifierContext Setup()
       {
           var ctx = new AndroidProjectFilesModifierContext();
           ctx.Dependencies.DependencyFiles = new[]
           {
               CustomSourceFileSrc
           };
           ctx.AddFileToCopy(CustomSourceFileSrc, CustomSourceFileDst);

           return ctx;
        }

       public override void OnModifyAndroidProjectFiles(AndroidProjectFiles projectFiles)
       {
       }
   }
   ```
6. From the Android Player settings window, go to **Other Settings** > **Configuration** > **Application Entry Point** and select **GameActivity**.
7. Select **Build & Run**.
8. Touch the phone screen and check the logcat.

You can now check the `HelloFromBridge` log, sent from `MyFile.cpp` and printed by `SendMessageReceiver.cs` script.

**Notes**:

* `UnityGameActivityPluginLoad` in `MyFile.cpp` is weakly linked and is called when GameActivity bridge initializes. There’s also ShutdownUserCode if you need it.
* `MyFile.cpp` contains `UnityEventProcessInput` event. You can find more events in `UGAEvents.h` file.

## Additional resources

* [Update the GameActivity library](android-application-entries-game-activity-update-library.md)