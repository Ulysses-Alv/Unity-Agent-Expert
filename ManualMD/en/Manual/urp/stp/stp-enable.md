# Enable Spatial-Temporal Post-processing in URP

To enable STP in the Universal **Render Pipeline** (URP):

1. Select the active URP Asset in the **Project window**.
2. In the Inspector, go to **Quality** > **Upscaling Filter**, and select **Spatial-Temporal Post-Processing**.

STP remains active when **Render Scale** is set to **1.0** as it applies temporal anti-aliasing (TAA) to the final rendered output.

**Note**: STP is not compatible with ****Dynamic Resolution**** in URP, only with **Render Scale**.