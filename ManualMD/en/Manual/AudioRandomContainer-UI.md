# Audio Random Container reference

Use an Audio Random Container to set up, play, and randomize a list of audio clips. Unless otherwise stated, the following table describes each Audio Random Container property and the effect of modifying its value in Edit mode.

| **Property** | **Description** |  |
| --- | --- | --- |
| **Play** | Start or stop the playlist and preview the selected Playback mode. The default shortcut for previewing an Audio Random Container is `Alt+P` (`Option+P` on macOS). This shortcut can be configured in the [Shortcuts Manager](ShortcutsManager.md). |  |
| **Skip** | Skip between the different audio clips in the playlist. |  |
| **Volume** | Set the volume of the Audio Random Container. The volume is expressed as a number between 0–1 and is displayed as a scaled db (decibel) value. Note: The playback volume of each audio clip is additive based on the volume of the Audio Source that references the Audio Random Container, and the value of this property. You can also adjust the value of this property while in Play mode. |  |
|  | **Randomization** | Randomize the volume, with a minimum and a maximum distribution. Set both the minimum and maximum to zero to deactivate randomization. |
| **Pitch** | Set the pitch on the Audio Random Container. You can set in Edit and Play modes, and it keeps the values set in both modes. If multiple Audio Sources reference the Audio Random Container, all of the setups use this pitch. The Pitch on the Audio Random Container and Audio Source is additive and combined. The Pitch on the Audio Random Container is equal to the Audio Source range, but is logarithm defined in a cent scale. |  |
|  | **Randomization** | Randomize the pitch, with a minimum and a maximum distribution. Set both the minimum and maximum to zero to deactivate randomization. |
| **Audio Clips** | Add the audio files in the Audio Random Container. You can reorder the list, and you can activate and deactivate each clip. The changes only take effect from the next audio clip used. Each clip can also have a unique volume, which is useful when one clip is louder or quieter than the rest. |  |
| **Trigger** |  |  |
|  | **Manual** | Deactivate the Automatic trigger mode. When Trigger is set to Manual, the Automatic Trigger Mode and related properties are hidden. To advance the play cycle, you have to manually call the `AudioSource.Play` method from an external script or through Visual Scripting. |
|  | **Automatic** | Activate the Automatic trigger mode, which activates an internal trigger that can loop playlist cycles. |
| **Playback Mode** | Set up the playlist to accommodate different use cases. |  |
|  | **Sequential** | Play the files in the order the audio clip list is in, so it plays the list from the top. Triggers and offset won’t work in this mode in this scope. |
|  | **Shuffle** | Remove audio clips from the pool after they’ve played. This option avoids repetition until all audio clips have played. |
|  | **Random** | Keep the pool of audio clips intact. After an object plays, it remains on the list of audio clips that can play and can therefore repeat. |
| **Avoid Repeating Last** | Define how many items can play before an object plays again. |  |
| **Automatic Trigger Mode** |  |  |
|  | **Pulse** | Trigger an audio clip from the list in steady pulse defined by the Time and the randomized values. Contrary to Offset, the Pulse intervals start from the beginning of the Audio Random Container, so the clips play at the same time interval. For example, you can use this mode for an automatic rifle. |
|  | **Offset** | Trigger an audio clip from the list based on when the last audio clip was done. Offset is the pause between two audio clips. |
| **Time** | Determine the time interval for both Pulse and Offset. |  |
|  | **Randomization** | Randomize the time, with a minimum and a maximum distribution. Set both the minimum and maximum to zero to deactivate randomization. |
| **Loop** | Define different ways to loop the Audio Random Container. |  |
|  | **Infinite** | Ensure the Audio Random Container never stops playing. |
|  | **Clips** | Set the number of audio clips played that the count refers to. |
|  | **Cycles** | Set the number of audio cycles played that the count refers to. |
| **Count** | Set the number of clips and cycles. |  |
|  | **Randomization** | Randomize the count, with a minimum and a maximum distribution. Set both the minimum and maximum to zero to deactivate randomization. |

## Additional resources

* [Audio Clips](class-AudioClip.md)
* [Audio Source](class-AudioSource.md)