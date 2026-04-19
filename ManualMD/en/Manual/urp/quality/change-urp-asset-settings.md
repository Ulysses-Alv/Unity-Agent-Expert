# Change URP Asset settings at runtime

You can change some properties of the URP Asset at runtime with C# **scripts**. This can help fine tune performance on devices with hardware that doesn’t perfectly match any of the quality levels in your project.

> **Note**: To change a property of the URP Asset with a C# script, the property must have a `set` method. For more information on these properties, refer to [Universal Render Pipeline Asset API](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.2/api/UnityEngine.Rendering.Universal.UniversalRenderPipelineAsset.html#properties).

The following example uses the QualityControls script and QualityController object from the [Change Quality Level](quality-settings-through-code.md) section, and extends the functionality to locate the active URP Asset and change some of its properties to fit the performance level of the hardware.

1. Open the QualityControls script.
2. At the top of the script add `using UnityEngine.Rendering` and `using UnityEngine.Rendering.Universal`.
3. Add a method with the name `ChangeAssetProperties` and the type `void` to the `QualityControls` class as shown below.

   ```
   using System.Collections;
   using System.Collections.Generic;
   using UnityEngine;
   using UnityEngine.Rendering;
   using UnityEngine.Rendering.Universal;

   public class QualityController : MonoBehaviour
   {
       void Start()
       {
           // Select the appropriate Quality Level first
           SwitchQualityLevel();
       }

       private void SwitchQualityLevel()
       {
           // Code from previous example
       }

       private void ChangeAssetProperties()
       {
           // New code is added to this method
       }
   }
   ```
4. Retrieve the active **Render Pipeline** Asset with `GraphicsSettings.currentRenderPipeline` as shown below.

   > **Note**: You must use the `as` keyword to cast the Render Pipeline Asset as the `UniversalRenderPipelineAsset` type for the script to work correctly.

   ```
   using System.Collections;
   using System.Collections.Generic;
   using UnityEngine;
   using UnityEngine.Rendering;
   using UnityEngine.Rendering.Universal;

   public class QualityController : MonoBehaviour
   {
       void Start()
       {
           // Select the appropriate Quality Level first
           SwitchQualityLevel();
       }

       private void SwitchQualityLevel()
       {
           // Code from previous example
       }

       private void ChangeAssetProperties()
       {
           // Locate the current URP Asset
           UniversalRenderPipelineAsset data = GraphicsSettings.currentRenderPipeline as UniversalRenderPipelineAsset;

           // Do nothing if Unity can't locate the URP Asset
           if (!data) return;
       }
   }
   ```
5. Add a `switch` statement in the ChangeAssetProperties method to set the value of the URP Asset properties.

   ```
   using System.Collections;
   using System.Collections.Generic;
   using UnityEngine;
   using UnityEngine.Rendering;
   using UnityEngine.Rendering.Universal;

   public class QualityController : MonoBehaviour
   {
       void Start()
       {
           // Select the appropriate Quality Level first
           SwitchQualityLevel();
       }

       private void SwitchQualityLevel()
       {
           // Code from previous example
       }

       private void ChangeAssetProperties()
       {
           // Locate the current URP Asset
           UniversalRenderPipelineAsset data = GraphicsSettings.currentRenderPipeline as UniversalRenderPipelineAsset;

           // Do nothing if Unity can't locate the URP Asset
           if (!data) return;

           // Change URP Asset settings based on the size of the device's graphics memory
           switch (SystemInfo.graphicsMemorySize)
           {
               case <= 1024:
                   data.renderScale = 0.7f;
                   data.shadowDistance = 50.0f;
                   break;
               case <= 3072:
                   data.renderScale = 0.9f;
                   data.shadowDistance = 150.0f;
                   break;
               default:
                   data.renderScale = 0.7f;
                   data.shadowDistance = 25.0f;
                   break;
           }
       }
   }
   ```
6. Add a call to the `ChangeAssetProperties` method in the `Start` method. This ensures that the URP Asset only changes when the **scene** first loads.

   ```
   using System.Collections;
   using System.Collections.Generic;
   using UnityEngine;
   using UnityEngine.Rendering;
   using UnityEngine.Rendering.Universal;

   public class QualityController : MonoBehaviour
   {
       void Start()
       {
           // Select the appropriate Quality Level first
           SwitchQualityLevel();

           // Fine tune performance with specific URP Asset properties
           ChangeAssetProperties();
       }

       private void SwitchQualityLevel()
       {
           // Code from previous example
       }

       private void ChangeAssetProperties()
       {
           // Locate the current URP Asset
           UniversalRenderPipelineAsset data = GraphicsSettings.currentRenderPipeline as UniversalRenderPipelineAsset;

           // Do nothing if Unity can't locate the URP Asset
           if (!data) return;

           // Change URP Asset settings based on the size of the device's graphics memory
           switch (SystemInfo.graphicsMemorySize)
           {
               case <= 1024:
                   data.renderScale = 0.7f;
                   data.shadowDistance = 50.0f;
                   break;
               case <= 3072:
                   data.renderScale = 0.9f;
                   data.shadowDistance = 150.0f;
                   break;
               default:
                   data.renderScale = 0.7f;
                   data.shadowDistance = 25.0f;
                   break;
           }
       }
   }
   ```

Now when this scene loads, Unity detects the system’s total graphics memory and sets the URP Asset properties accordingly.

You can use this method of changing particular URP Asset properties in conjunction with changing quality levels to fine tune the performance of your project for different systems without the need to create a quality level for every target hardware configuration.