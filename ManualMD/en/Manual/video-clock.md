# Clock management with the Video Player component

This section describes the different time update modes available for the Video Player component and how it interacts with them.

You can use the Video Player component to control how to time video playback relative to other interactive behaviors. For example, you can synchronize video playback with animation or with audio. You can do this through the following time update modes:

* Audio Digital Signal Processing (DSP) clock
* Game time
* Unscaled game time

The Video Player follows Unity’s capture frame rate [(`Time.captureFramerate`)](../ScriptReference/Time-captureFramerate.md) and capture delta time [(`Time.captureDeltaTime`)](../ScriptReference/Time-captureDeltaTime.md).

## Capture frame rate

If you set [capture frame rate](time-capture-frame-rate.md) to a value above 0, the Video Player’s playback becomes synchronous. The Video Player displays all frames at the expected time stamp even if it needs to delay the entire game’s execution.

With a non-zero capture frame rate, Video Player operations - including preparation and seek - run synchronously to ensure consistent results. However, this might slow down game execution especially when the video must remain synchronized with the rest of the **scene**. For example, the Unity Recorder uses this special behavior of time to produce smooth recordings.

The capture frame rate only affects the Video Player when you set Video Player’s **Time Update Mode** (`timeUpdateMode`) to **Game Time** (`VideoTimeUpdateMode.GameTime`).

If you set **Time Update Mode** to a different value, the Video Player ignores capture frame rate and runs asynchronously. The Video Player might skip or repeat frames to stay in sync, and it performs preparation and seek operations asynchronously so playback isn’t blocked.

**Note:** On the Web platform, Unity’s Video Player doesn’t support synchronous playback with `captureFramerate`. By default, it uses the normal asynchronous playback that’s described with the **Game Time** update mode.

## Audio DSP clock

The audio DSP clock comes from the Audio module. You can access it through [`AudioSettings.dspTime`](../ScriptReference/AudioSettings-dspTime.md).

## Game time

Game time is set in the Time module. You can access it using C# through [`Time.time`](../ScriptReference/Time-time.md). When you use game time, also known as scaled game time, Unity follows the Time Scale value set in the Time window. You can also modify the Time Scale value using C# through [`Time.timeScale`](../ScriptReference/Time-timeScale.md). However, if the capture frame rate or capture delta time have non-zero values, game time changes according to the rules that govern the **time manager**’s capture frame rate feature.

Note: Capture delta time is the same as 1 divided by the capture frame rate. It’s another view of the same information.

## Unscaled game time

When you use Unscaled game time, the Video Player ignores the Time Scale value. Unscaled game time comes from the Time module. You can access it using C# through [`Time.unscaledTime`](../ScriptReference/Time-unscaledTime.md).

## Additional resources

* [Important classes - Time](managing-time-and-frame-rate.md)
* [TimeManager class description](class-TimeManager.md)
* [Time scripting API](../ScriptReference/Time.md)