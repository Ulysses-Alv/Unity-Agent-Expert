# Adaptive Performance provider settings reference

Explore the settings available for Adaptive Performance providers.

Each [Adaptive Performance provider](providers.md) supplies several settings for controlling behavior at runtime and during development. The following settings are available for all providers unless otherwise specified. Values are provider-specific, meaning the actual impact of the value depends on the provider you’re using.

Access these settings at **Edit** > **Project Settings** > **Adaptive Performance**.

![Adaptive Performance provider settings in the Project Settings window.](../../uploads/Main/ap/ap-provider-settings.png)

Adaptive Performance provider settings in the Project Settings window.

## Runtime Settings

The **Runtime Settings** section contains the following settings:

| **Property** | **Description** |
| --- | --- |
| **Auto Performance Mode** | Automatically adjusts CPU and GPU levels to balance performance and battery life. This setting is available only for the Simulator. |
| **Boost mode on startup** | Boosts CPU and GPU performance at application startup before anything loads. Learn more about [temporary performance boosts](temporary-performance-boosts.md). This setting is available only for the Simulator. |

### Indexer Settings

The Indexer monitors the device’s thermal and performance state to decide when to adjust quality levels using [scalers](scalers.md).

The **Indexer Settings** section contains the following settings:

| **Property** | **Description** |
| --- | --- |
| **Active** | Enables the Indexer system. |
| **Thermal Action Delay** | Sets the delay in seconds after any scaler is applied or unapplied because of thermal state. This option isn’t available for the [Basic provider](basic-provider.md). |
| **Performance Action Delay** | Sets the delay in seconds after any scaler is applied or unapplied because of performance state. |

### Scaler Profiles

Every scaler shares a group of settings that determine the scaler’s impact on performance. The Indexer prioritizes scalers with a higher impact, adjusting their levels more frequently to change the application’s quality. Adaptive Performance calculates this impact score at runtime using the scaler’s pre-configured settings and the current device bottleneck.

The **Scaler Profiles** section contains the active [scaler profiles](scalers-introduction.md#scaler-profiles). Each individual scaler has the following settings:

| **Property** | **Description** |
| --- | --- |
| **Enabled** | Enables or disables the scaler. You can change this setting during runtime. |
| **Min Scale** | Defines the minimum scale of the scaler. If the minimum scale is 0.5, the minimum quality is 50% of the original value. |
| **Max Scale** | Defines the maximum scale of the scaler. If the maximum scale is 1.5, the maximum quality is 150% of the original value. Often, the maximum scale is set to 1, which is then used as the default for a high-end device. |
| **Max Level** | Defines the number of discrete quality levels between the minimum and maximum scale. A maximum level of 20 means that a scaler has 20 levels available between maximum and minimum quality. A binary scaler has a maximum level of 1. |
| **Visual Impact** | Describes how noticeable the quality change from this scaler is to the player (**Low**, **Medium**, or **High**). Adaptive Performance prioritizes reducing scalers with low visual impact first. |
| **Target** | Identifies which performance bottleneck this scaler addresses. This read-only setting helps the system apply the most effective scaler when it detects a bottleneck.  * **CPU**: The bottleneck is CPU. * **GPU**: The bottleneck is the GPU. * **Fillrate**: The bottleneck is the target frame rate. |

#### Add New Scaler Profile

Add or remove scaler profiles to quickly move between different performance configurations.

To create scaler profiles and load different profiles at runtime, refer to [Use scaler profiles](scaler-profiles.md).

## Development Settings

You can only use these settings in **development builds**. Release builds automatically disable them.

The **Development Settings** section contains the following settings:

| **Property** | **Description** |
| --- | --- |
| **Logging** | Lets the Adaptive Performance subsystem log messages to the **player log**. |
| **Logging Frequency** | Sets how frequently the system logs messages. Specified in frames. |

## Additional resources

* [Set up Adaptive Performance](initial-setup.md)
* [Scalers](scalers.md)
* [Providers](providers.md)