# Troubleshoot mismatched Cell Layouts

When you create a **tilemap** in an existing Grid **GameObject**, it is possible to try to create a tilemap that uses a different **Cell Layout** to the Grid. If this happens, Unity detects the mismatch in the **Cell Layouts** and a dialog appears. The dialog prompts you to decide if you want to continue creating the child tilemap or stop the action.

Select **Continue** to continue creating the child tilemap, and Unity changes the parent Grid’s **Cell Layout** to match the child’s. Alternately you can select **Cancel** to stop creating the new child tilemap.

**Note:** If you are running the Unity Editor in [headless mode](../../EditorCommandLineArguments.md), this dialog won’t appear and the Editor creates the tilemaps automatically as if you have selected **Continue**.