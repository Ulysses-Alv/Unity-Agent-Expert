# Update the GameActivity library

The GameActivity application entry point is implemented as a library separate from the Unity Editor which means that you can update the library independently. This is useful if Google provides bug fixes that your project requires because you can acquire the fixes via a GameActivity library version update.

**Note**: Unity doesn’t test all combinations of Unity version and GameActivity library version. If you update to a newer GameActivity library version, test your application thoroughly. For more information about the recommended GameActivity library version per Unity version, refer to [Unity and GameActivity version compatibility](android-application-entries-game-activity-requirements.md).

To update the GameActivity library version, change the value of the `androidx.games:games-activity` dependency in `build.gradle`. For information on the possible methods to do this, refer to [Modify Gradle project files](android-modify-gradle-project-files.md).

**Note**: Make sure that the other AndroidX dependencies support the GameActivity version that you want to use. If they don’t you must update them too. For more information, refer to [Declaring dependencies](https://developer.android.com/jetpack/androidx/releases/games#declaring-dependencies).

## Additional resources

* [Modify GameActivity bridge code](android-application-entries-game-activity-modify-bridge.md)