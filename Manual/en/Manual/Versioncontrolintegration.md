# Version control integrations

Unity has integrations with two version control systems: [Perforce](perForceIntegration.md) and [Plastic SCM](https://docs.unity.com/ugs/en-us/manual/devops/manual/unity-version-control). To use the [version control](https://unity.com/products/plastic-scm/source-code-management#what-version-control) integrations in Unity, you need to have either a **Perforce** or **Plastic SCM** server set up for your Unity Project.

## Setting up version control in Unity

Open a Unity Project, go to **Edit** > **Project Settings**, and then select the **Version Control** category.

| **Property** |  | **Description** |
| --- | --- | --- |
| **Mode** |  | Select the version control mode. |
|  | **Hidden meta files** | Hide the .meta files in your operating system’s file explorer. Unity does not show .meta files in the Project view, no matter which mode you choose. |
|  | **Visible meta files** | Select this option to work with a version control system that Unity doesn’t support. This is the default setting. You can then manage the source Assets and metadata for those Assets with a version control system of your choice. |
|  | **Perforce** | Select this option if you use [Perforce](https://www.perforce.com/) as your version control system. |
| **Username** |  | Enter the username associated with your Perforce account. This property is only visible when **Mode** is set to **Perforce**. |
| **Password** |  | Enter the password associated with your Perforce account. This property is only visible when **Mode** is set to **Perforce**. |
| **Workspace** |  | Enter your workspace. For example, `Example**Workspace**1`. This property is only visible when **Mode** is set to **Perforce**. |
| **Server** |  | Enter the server your Unity Project is on. For example, localhost:1666. This property is only visible when **Mode** is set to **Perforce**. |
| **Host** |  | Enter the hostname that you want your computer to impersonate. For example, workstation123.perforce.com. This property is only visible when **Mode** is set to **Perforce**. |
| **Log Level** |  | Select how much version control information to receive in Unity’s console log. |
|  | **Verbose** | Log every operation related to version control. This option provides very detailed logging, and is useful if you want to debug your version control setup. This property is only visible when **Mode** is set to **Perforce**. |
|  | **Info** | Log errors, warnings, and information related to version control. |
|  | **Notice** | Log errors and warnings. |
|  | **Fatal** | Unity prints only fatal errors to the console. |
| **Status** |  | Display information on the status of the connection to the version control system. If you are not connected, select **Connect** to connect to the system you have configured. This property is only visible when **Mode** is set to **Perforce**. |
| **Automatic Add** |  | When this setting is enabled, automatically add your files to the version control system when you add them to the Project, either via the Editor or the folder on disk. When this setting is disabled, you need to add files manually to the version control system. This setting is enabled by default. This property is only visible when **Mode** is set to **Perforce**. |
| **Work Offline** |  | Enable this setting to work offline. When this setting is enabled, you need to reconcile offline work in P4V or use the reconcile command in P4 to bring the Perforce server depot up to date with the work you did while offline. For more information, refer to [Working offline with Perforce](perForceIntegration.md). This property is only visible when **Mode** is set to **Perforce**. |
| **Async Update** |  | Enable this setting to use asynchronous version control status queries. When enabled, Perforce updates the version control status of files without stalling the Unity Editor. Use this setting when the connection to your version control server has high latency.    **Note:** Only status queries are asynchronous. Unity synchronously performs operations that change the state of files, or require up-to-date knowledge of a file status. This property is only visible when **Mode** is set to **Perforce**. |
| **Show Failed Checkouts** |  | Enable this property to display a dialog when Perforce can’t perform the check out operation. This might happen if you lose connection, or if another user has exclusively checked out the Asset you want to edit. This property is only visible when **Mode** is set to **Perforce**. |
| **Overwrite Failed Checkout Assets** |  | When you enable this setting, Unity saves any Assets that can not be checked out. This means Unity forces through a save to a file, even if Perforce cannot check out the file. This is enabled by default. If you disable it, Unity doesn’t force your files to save if Perforce can’t check them out. This property is only visible when **Mode** is set to **Perforce**. |
| **Smart Merge** |  | [Smart Merge](SmartMerge.md) makes it easier for Unity to merge files that have changes on the same line. This is useful if several users are working on the same Project at the same time. This property is only visible when **Mode** is set to **Perforce**. |
|  | **Off** | Disable Smart Merge. |
|  | **Ask** | Enable Smart Merge, but receive a notification before you merge, if a conflict happens. This is the default setting. |
|  | **Premerge** | Automatically use Smart Merge. |
| **Version Packages Outside Project** |  | Tracks changes to packages that reside on disk outside of the Unity project’s root folder while still in the local workspace. This property is only visible when **Mode** is set to **Perforce**. |
| **Overlay Icons** |  | Enable this setting to display version control status icons in the Editor. This property is only visible when **Mode** is set to **Perforce**. |

Configure the **Version Control** settings, then select the **Connect** button next to the status area to connect to the version control system. When Unity connects to the system, **Connected** displays in the status area.

To adjust your revision control tool, open **Preferences** and then select the **External Tools** section. You can choose a new tool under **Revision Control Diff/Merge**.

## Using version control in Unity

When you set the Editor up to work with your version control system, you can perform version control operations via the Editor, instead of in the version control client. To do this, right-click on the Asset in the Project view.

The version control operations vary depending on which version control you use. The following table displays what actions are available for each version control:

| **Version control operation** | **Description** | **Perforce** | **Plastic SCM** |
| --- | --- | --- | --- |
| **Get Latest** | This updates the files on your machine to match those in the version control system. | Yes | No. To get the latest changes and update the file, you need to use the [version control window](#window) (**Window** > **Version Control** > **External Provider**). |
| **Submit** | Submits the current state of the files to the version control system. | Yes | Yes |
| **Check Out** | Allows changes to be made to the files. | Yes | Yes |
| **Check Out (Other)** | Select whether to check out both the Asset and its .meta file, or only the Asset, or only the .meta file. | Yes | No |
| **Mark Add** | Adds the files into version control. | Yes | Yes |
| **Revert** | Discards changes to open changed files. | Yes | Yes |
| **Revert Unchanged** | Removes the checked out status from files that have previously been checked out but you haven’t modified. | Yes | Yes |
| **Resolve Conflicts** | Resolves conflicts on files that have been changed by multiple users. | Yes | No. Conflicts appear in the version control menu, but you need to resolve them in the Plastic SCM GUI. |
| **Lock** | Prevents other users from submitting changes to the files. | Yes | No. To lock or unlock files in Plastic SCM, you must edit a specific Plastic SCM lock file externally. For more information, see the page on [Plastic SCM Integration](https://docs.unity.com/ugs/en-us/manual/devops/manual/unity-version-control). |
| **Unlock** | Releases the lock and allows anyone to submit changes. | Yes | No. To lock or unlock files in Plastic SCM, you must edit a specific Plastic SCM lock file externally. For more information, see the page on [Plastic SCM Integration](https://docs.unity.com/ugs/en-us/manual/devops/manual/unity-version-control). |
| **Diff** | Compares the differences between the local files on your computer and the files on the server. You can choose to diff only the Asset file, or the Asset file and its .meta file. | Yes | Yes |

### Checking out files

In some version control systems, such as Perforce, versioned files are read-only by default,
and require you to check them out before you edit them (unless you have enabled the **Work offline** setting). When you work with versioned Assets from the Editor, the **Inspector** displays a **Check Out** button that enables file editing. Additionally, **Project Settings** inspectors have a **Checkout** button that you can use to check out specific Project settings.

If you have custom Editor script code that disables parts of another custom editing tool
for read-only Assets, or if you are writing to versioned files manually, use the
[AssetDatabase.IsOpenForEdit](../ScriptReference/AssetDatabase.IsOpenForEdit.md) and
[AssetDatabase.MakeEditable](../ScriptReference/AssetDatabase.MakeEditable.md) APIs to
check for file editability and to perform check out operations.

The Version Control integration also exposes
[Provider.PreCheckoutCallback](../ScriptReference/VersionControl.Provider.PreCheckoutCallback.md)
and [Provider.PreSubmitCallback](../ScriptReference/VersionControl.Provider.PreSubmitCallback.md)
C# callbacks for custom version control operation validation logic.

Unity writes any Assets you modify or mark as modified in the Editor to the disk when it performs a **Save Project** operation. The Assets are then checked out in version control if needed. This might lead to Assets getting checked out even if no actual changes to the file happens. This most often happens when an Editor script calls [EditorUtility.SetDirty](../ScriptReference/EditorUtility.SetDirty.md) on an Asset, without checking if it was modified.

**Note:** If Unity cannot commit your changes to your version control client (for example, if the server is down or if license issues occur), it stores your changes in a separate changeset.

When you save your changes to a .**scene** file, Unity automatically checks it out. If you use Plastic SCM, it also automatically checks out automatically generated Assets, such as light maps.

## Version Control window

You can use the Version Control window to view the files in your changelist. To access the window navigate to **Window** > **Version Control** > **External Provider**.

The **Outgoing** tab lists all of the local changes that are pending a commit into version control. The **Incoming** tab lists all of the changes that need to be pulled from version control.

Right-click Assets or changelists in the window to perform operations on them. To move Assets between changelists, drag them from one changelist to the header of the target changelist.

### Icons

The Editor displays the following icons to visualize the version control status for files and Assets:

| **Icon** | **Purpose** |
| --- | --- |
|  | File added locally, and pending addition into version control. |
|  | File added to version control by another user, and pending addition into version control. |
|  | File is checked out locally. |
|  | File is checked out by another user remotely. |
|  | There has been a conflict merging this file and it needs to be resolved. |
|  | File has been deleted locally, and is pending deletion in version control. |
|  | File has been deleted by another user and is pending deletion in version control. |
|  | File is not yet under version control. You can use the **Mark Add** operation to add the file manually. |
|  | File is locked by you and other users cannot modify it. |
|  | File is locked by another user and you cannot modify it. |
|  | Another user has checked in a new version of this file. Use the **Apply Incoming Changes** operation to get the latest version. |
|  | The server is requesting the version control status of this file, or is waiting for a response. This only appears if you use a centralized version control system like Perforce. |