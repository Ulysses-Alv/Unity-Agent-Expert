# Diagnostics

Use the **Diagnostics** settings to enable or disable the collection of diagnostic data for your project. When enabled, you can further control what data is collected and how it’s used in the Unity Dashboard. For more information, refer to [Configure Developer Data settings in the Dashboard](https://docs.unity.com/en-us/cloud/developer-data/configure-diagnostics-dashboard).

| **Property** | **Description** |
| --- | --- |
| **Diagnostic Data** | Enable or disable collection of diagnostic data for your Unity project.  This value is used as the default setting for all your build profiles. You can override per build via Build Profiles.  The following options are available:  * **Enabled**: Enables collection of diagnostic data for the project. * **Disabled**: Disables collection of diagnostic data for the project. |

Note: Diagnostic Data is set to Disabled if your project isn’t connected to Unity Cloud. To collect diagnostic data, you must link your project to Unity Cloud via Project Settings. For more information, refer to [Configure a project for Unity Cloud](https://docs.unity.com/cloud/en-us/projects/configure-project-for-unity-cloud).   
If you disable diagnostic data collection, this can impact the performance and behavior of services that rely on Developer Data.