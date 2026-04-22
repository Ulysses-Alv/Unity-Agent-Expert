# Submit a license request from the Hub

Submit a license request from the Unity Hub to manually activate your license if you’re unable to use the [other activation methods](LicenseActivationMethods.md).

**Note**: The manual activation method doesn’t work with floating license subscriptions or Unity Personal. To activate a license for Unity Personal, log in to the Unity Hub. To return a Personal license, log out of the Unity Hub.

Although you can submit a license request from the computer where you plan to use Unity, the following procedure shows a workaround if that computer can’t connect to the internet. The following procedure assumes you have access to two computers, as follows:

* The first computer is the computer where you installed the Unity Hub and the Unity Editor, but the computer can’t connect to the internet.
* The second computer is any computer that can connect to the internet.

## Procedure

1. On the first computer, open the Unity Hub.
2. Select **Licenses** from the side menu.
3. Select **Add license**.
4. Select **Activate with license request**.
5. Follow the **Download license request** step on-screen. This step creates an `.alf` file on the first computer. Keep Unity Hub open for a later step.
6. Copy the `.alf` file from the previous step to a second computer that has internet access.
7. Use the second computer to go to [license.unity3d.com/manual](https://license.unity3d.com/manual) and follow the instructions there. This step generates a `.ulf` file. If you experience issues accessing the website:

   * Go to [id.unity.com](https://id.unity.com/) and make sure you’re signed in to your Unity ID. You need a valid login to request a Unity license file.
   * Try pasting `https://license.unity3d.com/manual` into your browser’s address bar.
8. Copy the `.ulf` license file from the second computer to the first computer.
9. From the first computer, follow Unity Hub’s **Activate with license file** on-screen. This step doesn’t need internet access.

After completing the process, you can view your active license in the Hub by selecting **Licenses** from the side menu.

## Additional resources

* [License activation methods](LicenseActivationMethods.md)
* [Activate a license for a Unity Personal plan](https://docs.unity3d.com/hub/manual/ManageLicense.html#personal) (Unity Hub documentation)
* [License troubleshooting](ActivationFAQ.md)