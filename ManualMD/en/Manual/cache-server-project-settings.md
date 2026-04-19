# Cache Server Project Settings reference

You can adjust how Unity Accelerator works with the Cache Server Project Settings. To open the Cache Server settings go to **Edit > Project Settings > Editor > Accelerator Cache Server**

| **Setting** | **Description** |
| --- | --- |
| **Mode** | Choose whether to use the cache server. Choose from the following options:  * **Use global settings:** Uses the settings in the **Preferences** window (**Settings > Asset Pipeline > Unity Accelerator**). * **Enabled**: Enable the cache server. * **Disabled:** Disable the cache server. |
| **IP Address** | Input the **accelerator**’s IP address and port. |
| **Check Connection** | Select **Check Connection** to test the connectivity of the accelerator. |
| **Namespace prefix** | Set a custom namespace for the server. |
| **Download** | Enable downloads from the cache server. |
| **Upload** | Enable uploads from the cache server. |
| **TLS/SSL** | Enable encryption on the cache server. |
| **Authentication (Using Unity ID)** (Obsolete) | **Warning:** This setting is obsolete. Do not use it in your project.  Enable authentication for the cache server using Unity ID. |
| **Content Validation** | Select the level of content validation: Disabled, Upload Only, Enabled, Required. |
| **Download Batch Size** | Set the size of download. |

## Additional resources

* [Configure Unity Accelerator in the Editor](accelerator-configure.md)