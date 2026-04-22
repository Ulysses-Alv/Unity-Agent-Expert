# Creating scripts

Scripts allow you to customize and extend the capabilities of your applicaton with C# code. With **scripts** that derive from Unity’s built-in [MonoBehaviour](class-MonoBehaviour.md) class you can create your own custom Components to control the behavior of **GameObjects**. With scripts that derive from [ScriptableObject](class-ScriptableObject.md) you can store large amounts of data efficiently in your application. Alternatively, you can start with an empty C# script to develop your own non-Unity classes.

Unlike most other assets, scripts are usually created within Unity directly. To create a new script:

* From the main menu: go to **Assets > Create > Scripting** and select the type of script you want to create.

Or:

* From the Create menu (plus sign) in the [Project window toolbar](ProjectView.md): go to **Scripting** and select the type of script you want to create.

This creates a new script in whichever folder you have selected in the Project panel. It also selects the script’s file name for editing, prompting you to change the name. For things you should take into account when naming your scripts, refer to [Naming considerations for scripts](naming-scripts.md).

## Additional resources

* [Naming scripts](naming-scripts.md)
* [MonoBehaviour](class-MonoBehaviour.md)
* [ScriptableObject](class-ScriptableObject.md)