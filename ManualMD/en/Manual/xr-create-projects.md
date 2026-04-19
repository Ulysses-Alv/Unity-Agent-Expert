# Create an XR project

To create an **XR** project, you can use an XR template from the Unity Hub or start with a non-XR project.

* [Create a new XR project](#create-a-new-xr-project)
* [Start from a non-XR project](#start-from-a-non-xr-project)

## Prerequisites

To create an XR project, you must first perform the following tasks:

* [Install the Unity Editor](https://docs.unity3d.com/hub/manual/AddEditor.html).
* [Add Editor modules](https://docs.unity3d.com/hub/manual/AddModules.html) to support the platform build targets on which the XR devices that you want to support run. For example, to support Android devices with ARCore or Meta Quest devices, you must add the Android module to your Editor using the Unity Hub.

**Note:** The makers of some XR devices might impose additional requirements, such as signing up for a developer account, in order to create applications for their platform. Such requirements are outside the scope of Unity documentation.

## XR templates

The quickest way to create a new XR project is with one of the XR templates in the Unity Hub. Projects you create with these templates are already configured with the XR **Plug-in** Management system, common XR plug-in and support packages, and a starting **scene** set up with the XR components.

The XR templates available for this version of Unity include:

| Template | Version | Description |
| --- | --- | --- |
| [AR Mobile template](https://docs.unity3d.com/Packages/com.unity.template.ar-mobile@2.0) | 2.0 | For mobile **augmented reality** projects. The template configures the project for URP, **AR** Foundation, iOS, Android and the XR Interaction Toolkit. It includes an example scene and assets to demonstrate how to set up a project that’s ready for AR development on mobile devices. |
| [Mixed Reality template](https://docs.unity3d.com/Packages/com.unity.template.mixed-reality@2.1) | 2.1 | For **mixed reality** projects. The template configures the project for URP, OpenXR, AR Foundation, and the XR Interaction Toolkit. It includes an example scene and assets to demonstrate how to set up a project that’s ready for Mixed Reality. |
| [Mixed Reality multiplayer tabletop template](https://docs.unity3d.com/Packages/com.unity.template.mr-multiplayer@1.0/manual/index.html) | 1.0 | For multiplayer mixed reality projects. This template combines XRI, ARF, NGO, and UGS to create a starting point for Tabletop **MR** Multiplayer apps. |
| [VR template](https://docs.unity3d.com/Packages/com.unity.template.vr@9.1) | 9.1 | For **virtual reality** projects. The template configures the project for URP, OpenXR, and the XR Interaction Toolkit. It includes an example scene and assets to demonstrate how to set up a project that’s ready for interactive **VR**. |
| [VR Multiplayer template](https://docs.unity3d.com/Packages/com.unity.template.vr-multiplayer@2.0) | 2.0 | For virtual reality projects with multiplayer functionality. The template configures the project for URP, Netcode for Game Objects, Unity Cloud Gaming Services, and XR Interaction Toolkit. It includes example scenes and assets to demonstrate how to set up a project that’s ready for multiplayer VR experiences. |

Choose a template in the Unity Hub when you create a new project.

## Create a new XR project

To create an XR project using a template:

1. Open the Unity Hub.
2. In the Hub, click the **New Project** button.
3. Select the desired template: **Mixed Reality**, **VR**, or **AR Mobile**.
4. If necessary, click **Download template**.
5. Set a project name and save location.
6. Click **Create project**. Refer to the [Unity Hub documentation](https://docs.unity3d.com/hub/manual/AddProject.html) for more information about creating projects in the Unity Hub.
7. After the project opens in the Editor, [configure the project’s XR plug-ins](configuring-project-for-xr.md) with the **XR Plug-in Management** system.
8. Add additional XR packages, such as [AR Foundation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/) and the [XR Interaction Toolkit](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/), using the Package Manager. (A template might include these and other packages already.)
9. [Set up an XR scene](xr-scene-setup.md).

**Tip:** After you create an XR project, go to the **XR Plug-in Management** section in your ****Project Settings**** and enable the plug-ins for the platforms you want to support. If the tab for a platform is missing on the XR Plug-in Management page, add the platform module to your Editor installation using the [Unity Hub](https://docs.unity3d.com/hub/manual/AddModules.html).

## Start from a non-XR project

You can always convert an existing non-XR project:

1. Use the Unity Hub to open the project.
2. In the Editor, [configure the project’s XR plug-ins](configuring-project-for-xr.md) with the **XR Plug-in Management** system.
3. Add additional XR packages, such as [AR Foundation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/) and the [XR Interaction Toolkit](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/), using the Package Manager.
4. [Set up an XR scene](xr-scene-setup.md).