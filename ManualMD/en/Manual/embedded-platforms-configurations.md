# Configurations for embedded platforms

Unity supports a wide range of embedded platform configurations through [embedded support plans](https://unity.com/products/compare-plans/embedded).

When starting an embedded support plan, the chosen configuration will be fixed and supported throughout the length of the support plan. Unity only supports deployment to these embedded platforms with a valid embedded support plan. For more information, contact your Account Manager or the [Unity Sales](https://create.unity.com/unity-for-industries) team.

**Note:** For all configurations, usage of Unity for commercial embedded systems is subject to the [Embedded Software Restriction](https://unity.com/legal/terms-of-service#:~:text=Embedded%20Software%20Restriction) in the Unity terms of service.

## Configurations supported with Full Embedded Support plan

Unity is verified and supported to run on chipsets from AMD, Intel, Mediatek, Qualcomm, Renesas, Rockchip, Samsung, and others, depending on the requirements of each project. Support for these chipsets is currently limited to a [Full Embedded Support plan](https://unity.com/products/compare-plans/embedded#:~:text=Full).

## Verified configurations with Basic Embedded Support plan

Verified embedded platform configurations are continuously tested and confirmed to run Unity, using a publicly available board support package (BSP) provided by the chipset or OS vendor. These are the only configurations supported through the [Basic Embedded Support plan](https://unity.com/products/compare-plans/embedded#:~:text=Basic).

| **Vendor** | **Model** | **Verified OS/BSP** |
| --- | --- | --- |
| **NVIDIA** | [Jetson AGX Orin](https://www.nvidia.com/en-us/autonomous-machines/embedded-systems/jetson-orin/) | Ubuntu for Tegra |
| **NXP** | [I.MX8 QM](https://www.nxp.com/design/design-center/development-boards-and-designs/i-mx-evaluation-and-development-boards/i-mx-8quadmax-multisensory-enablement-kit-mek:MCIMX8QM-CPU) | * Yocto Linux * QNX 7.1 * QNX 8 * Android Automotive 13 * Android Automotive 14 |
| **Raspberry Pi** | * [Raspberry Pi 4b](https://www.raspberrypi.com/products/raspberry-pi-4-model-b/) * [Raspberry Pi 5](https://www.raspberrypi.com/products/raspberry-pi-5/) | **For Raspberry Pi 4b:**  * QNX 7.1 (Vulkan only) * Yocto Linux  **For Raspberry Pi 5:**  * Yocto Linux * Raspberry Pi OS |

**Notes:**

* The configurations listed above are verified to run Unity with the following graphics APIs:
  + Vulkan 1.0
  + OpenGL ES 3.0
* The performance of these configurations is hardware-dependent and might vary based on the specific use case. On some embedded configurations, the available **feature set** is limited.
* Multisample anti-aliasing (MSAA) might cause rendering issues on i.MX8 devices for both QNX and embedded Linux. If you’re experiencing rendering issues, disable MSAA in your **project settings**.

## Additional resources

* [Unity for embedded systems](https://unity.com/solutions/embedded)
* [System requirements](system-requirements.md)