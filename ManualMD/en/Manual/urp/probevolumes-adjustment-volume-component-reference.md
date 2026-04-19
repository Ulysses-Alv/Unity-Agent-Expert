# Probe Adjustment Volume component reference

Select a [Probe Adjustment Volume Component](probevolumes-troubleshoot-light-leaks.md#probevolumeadjustment) and open the **Inspector** to view its properties.

Refer to the following for more information about using the Probe Adjustment Volume component:

* [Fix issues with Adaptive Probe Volumes](probevolumes-fixissues.md)
* [Configure the size and density of Adaptive Probe Volumes](probevolumes-changedensity.md)

| **Property** | **Description** |
| --- | --- |
| **Influence Volume** | Sets the shape, size and radius of the Adjustment Volume. The options are:  * **Shape:** Set the shape of the Adjustment Volume to either **Box** or **Sphere**. * **Size:** Set the size of the Adjustment Volume. This property only appears if you set **Shape** to **Box**. * **Radius:** Set the radius of the Adjustment Volume. This property only appears if you set **Shape** to **Sphere**. |
| **Mode** | Sets how Unity overrides probes inside the Adjustment Volume. |

### Mode

| **Value** | **Description** |
| --- | --- |
| **Invalidate Probes** | Mark selected probes as invalid. Refer to [How light probe validity works](probevolumes-fixissues.md#how-light-probe-validity-works) for more information. |
| **Override Validity Threshold** | Override the threshold URP uses to determine whether **Light Probes** are marked as invalid. Refer to [Adjust Dilation](probevolumes-fixissues.md#adjust-dilation) for more information. |
| **Apply Virtual Offset** | Change the position Light Probes use when sampling the lighting in the **scene** during baking. Refer to [Adjust Virtual Offset](probevolumes-fixissues.md#adjust-virtual-offset) for more information. |
| **Override Virtual Offset Settings** | Override the biases URP uses during baking to determine when Light Probes use Virtual Offset, and calculate sampling positions. Refer to [Adjust Virtual Offset](probevolumes-fixissues.md#adjust-virtual-offset) for more information. |
| **Intensity Scale** | Override the intensity of probes to brighten or darken affected areas. |
| **Override Sky Direction** | Override the directions Unity uses to sample the ambient probe, if you enable [sky occlusion](probevolumes-skyocclusion.md). |
| **Override Sample Count** | Override the number of samples Unity uses for Adaptive Probe Volumes. |
| **Override Rendering **Layer Mask**** | Control the **Rendering Layer Mask** for probes inside the Adjustment Volume. This option only appears if Rendering Layer Masks are enabled in the Baking Set. Refer to [Use Rendering Layer Masks](probevolumes-fixissues.md#layers). |

### Mode settings

| **Property** | **Description** |
| --- | --- |
| **Dilation Validity Threshold** | Override the ratio of backfaces a probe samples before URP considers it invalid. This option only appears if you set **Mode** to **Override Validity Threshold**, and you enable [Advanced Properties](urp-asset-additional-settings.md). |
| **Virtual Offset Rotation** | Set the rotation angle for the Virtual Offset vector on all probes in the Adjustment Volume. This option only appears if you set **Mode** to **Apply Virtual Offset**. |
| **Virtual Offset Distance** | Set how far URP pushes probes along the Virtual Offset Rotation vector. This option only appears if you set **Mode** to **Apply Virtual Offset**. |
| **Geometry Bias** | Sets how far URP pushes a probe’s capture point out of geometry after one of its sampling rays hits geometry. This option only appears if you set **Mode** to **Override Virtual Offset Settings**. |
| **Ray Origin Bias** | Override the distance between a probe’s center and the point URP uses to determine the origin of that probe’s sampling ray. This can be used to push rays beyond nearby geometry if the geometry causes issues. This option appears only if you set **Mode** to **Override Virtual Offset Settings**. |
| **Intensity Scale** | Change the brightness of all probes covered by the Probe Volumes Adjustment Volume component. Use this sparingly, because changing the intensity of probe data can lead to inconsistencies in the lighting. This option only appears if you set **Mode** to **Intensity Scale**. |
| **Sky Direction** | Set the direction Unity uses to sample the ambient probe. This option only appears if you set **Mode** to **Override Sky Direction**. |
| **Rendering Layer Mask Operation** | Choose how to combine the volume’s mask with the probes’ mask: **Override** replaces the probes’ mask, **Add** adds the selected mask, **Remove** removes the selected mask. This option only appears if you set **Mode** to **Override Rendering Layer Mask**. |
| **Rendering Layer Mask** | Select the mask defined in the Baking Set. Up to four masks are supported and displayed by name. This option only appears if you set **Mode** to **Override Rendering Layer Mask**. |

#### Probes

| **Property** | **Description** |
| --- | --- |
| **Direct Sample Count** | Set the number of samples Unity uses to calculate the direct light each probe stores. This option only appears if you set **Mode** to **Override Sample Count**. |
| **Indirect Sample Count** | Set the number of samples Unity uses to calculate the indirect light each probe stores, including environment samples. This option only appears if you set **Mode** to **Override Sample Count**. |
| **Sample Count Multiplier** | Set the value Unity multiplies **Direct Sample Count** and **Indirect Sample Count** by. This option only appears if you set **Mode** to **Override Sample Count**. |
| **Max Bounces** | Set the number of times Unity bounces light off objects. This option only appears if you set **Mode** to **Override Sample Count**. |

#### Sky Occlusion

| **Property** | **Description** |
| --- | --- |
| **Sample Count** | Set the number of samples Unity uses to calculate the amount of light each probe receives from the sky, if you enable [sky occlusion](probevolumes-skyocclusion.md). This option only appears if you set **Mode** to **Override Sample Count**. |
| **Max Bounces** | Set the number of times Unity bounces light from the sky off objects to calculate the sky occlusion data, if you enable [sky occlusion](probevolumes-skyocclusion.md). This option only appears if you set **Mode** to **Override Sample Count**. |