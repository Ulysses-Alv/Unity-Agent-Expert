# Audio Profiler module reference

The Audio Profiler module monitors the performance of the audio system in your application, such as the total load and voice counts. To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**. For information on how to use the Profiler, refer to [Profiler introduction](profiler-introduction.md)

![The Audio Profiler in the Profiler Window](../uploads/Main/audio-profiler-module.png)

The Audio Profiler in the Profiler Window

The Audio Profiler module’s chart tracks the time spent on the audio in your application. The timings are divided into the following categories:

| **Chart** | **Description** |
| --- | --- |
| **Playing Audio Sources** | Total number of audio sources playing in the **scene** at the selected frame. This is useful to monitor if the audio is overloaded. |
| **Audio Voices** | The number of audio (FMOD channels) voices used in the selected frame. |
| **Total Audio CPU** | The amount of CPU usage the audio used in the selected frame. |
| **Total Audio Memory** | The amount of RAM the audio engine used in the selected frame. |

To change the order of the categories in the chart, you can drag them in the chart’s legend. You can also click a category’s colored legend to toggle its display.

## Module details pane

When you select the Audio Profiler module, the details pane displays a breakdown of the audio data for the selected frame. There are two views you can use to inspect the audio profiler data: **Simple** and **Detailed**. To change the display, use the top-left dropdown in the details pane (set to **Simple** by default).

### Simple view

The **Simple** view contains the following information about the audio data for the frame you selected.

| **Data** | **Description** |
| --- | --- |
| **Total Audio Sources** | Total number of [audio sources](class-AudioSource.md) in the scene. |
| **Playing Audio Sources** | Total number of audio sources playing in the scene |
| **Paused Audio Sources** | The total number of paused audio sources in the scene. |
| **Audio Clip Count** | Total number of [audio clips](class-AudioClip.md) in the scene. |
| **Audio Voices** | The total number of audio channels (FMOD channels) your Project uses. |
| **Total Audio CPU** | The total amount of CPU the audio uses. |
| **DSP CPU** | The amount of CPU your project uses by mixing, **audio effects**, and decompression of non-streamed sounds that have the [Compressed In Memory](class-AudioClip.md) load type. This doesn’t cover the CPU required for sounds that Unity decodes in the background whose load type is **Decompress On Load** and have the **Load In Background** flag checked. |
| **Streaming CPU** | The amount of CPU your project uses to stream the audio in your application. |
| **Other CPU** | General CPU overhead not covered by the other CPU categories. |
| **Total Audio Memory** | The amount of memory the audio uses in your project. |
| **Streaming File Memory** | The amount of memory that audio files with the load type of **Streaming** use for short-term buffering of compressed audio data when the files are progressively read from disk. |
| **Streaming Decode Memory** | The amount of memory that audio files with the load type of **Streaming** use for buffering the decoded sample stream. |
| **Sample Sound Memory** | The amount of memory that audio files with the load type of **Decompress On Load** use for the decompressed sample data.  **Note:** Unity pools the memory that the audio system allocates and it keeps growing until it reaches saturation over the runtime of your application. The audio system internally reuses the allocated memory, which can’t be compacted during runtime. |
| **Other Memory** | Overhead caused by various subsystems in the audio system. |

### Detailed view

The **Detailed** view contains all the information in the **Simple** view, and additionally contains detailed per-frame logging of audio events. You can view these either as **Channels**, **Groups**, or **Channels and Groups**. To open this view, open the dropdown at the top of the module details pane, select **Detailed**, and then select one of the views from the **toolbar** at the top of the pane.

The **Groups** view shows the hierarchy of the buses in the audio mixer. The **Channels and Groups** view displays this information alongside information about the playing sounds.

![The Audio Profiler Groups view](../uploads/Main/audio-profiler-module-groups.png)

The Audio Profiler Groups view

![The Audio Profiler Channel and Groups view](../uploads/Main/audio-profiler-module-channel-groups.png)

The Audio Profiler Channel and Groups view

Select the **Reset play count on play** to reset the number in the **Plays** column the next time you click Play in the Player window, or when you connect to a new target device.

The information in the detailed view is arranged into the following columns:

| **Column** | **Description** |
| --- | --- |
| **Object** | The **GameObject** that contains the Audio Source playing the audio. |
| **Asset** | The audio Asset that the corresponding GameObject Audio Source is playing. |
| **Volume** | The volume the Audio Source applies to the audio. This is a combination of its overall volume property, and the dynamic volume that the distance-dependent attenuation curve applies. |
| **Audibility** | The actual level at which the audio plays. This is the sum of the Audio Source’s volume and the other attenuations the mixer channels apply to it. |
| **Plays** | The number of times Unity has played the audio. This information is useful to debug logic errors where Unity might not use some audio files. |
| **3D** | Displays **YES** if the audio uses [dynamic distance-dependent attenuation](AudioMixerOverview.md) and directional panning. |
| **Paused** | Displays **YES** if the audio is paused in this frame. |
| **Muted** | Displays **YES** if the audio is muted in this frame. |
| **Virtual** | Displays **YES** if the audio is suspended because of the **Max Real Voice Count**, which you can set in the [Audio Project Settings](class-AudioManager.md). The **Max Real Voice Count** sets the maximum amount of Audio Sources that Unity simultaneously plays at once. If this displays true, Unity has prioritized other audio that has higher audibility or priority in this frame. |
| **OneShot** | Displays **YES** if [AudioSource.PlayOneShot()](../ScriptReference/AudioSource.PlayOneShot.md) played the audio. |
| **Looped** | Displays **YES** if [AudioSource.Play()](../ScriptReference/AudioSource.Play.md) played the audio. |
| **Distance** | The distance from the Audio Source to the [AudioListener](../ScriptReference/AudioListener.md). |
| **MinDist** | The minimum distance defined on the [AudioSource curve editor](class-AudioSource.md). This defines a spherical region around the audio where the volume stays at a constant level within it. |
| **MaxDist** | The maximum distance defined on the Audio Source curve editor. This defines a spherical region around the audio where the volume stays at a constant level outside of it. |
| **Time** | The current relative time in the audio’s playback. While the audio playback is paused, this time does not advance. |
| **Duration** | The length of the audio in seconds. |

## Additional resources

* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)