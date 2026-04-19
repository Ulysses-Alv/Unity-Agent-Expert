# Deactivate GameObjects

To temporarily remove a **GameObject** from your **scene**, you can mark the GameObject as inactive.

To do this, navigate to the **Inspector** window and clear the checkbox to the left of the GameObject’s name. The names of deactivated GameObjects appear faded in the Hierarchy window.

To deactivate a GameObject through script, use the [SetActive](../ScriptReference/GameObject.SetActive.md) method. To see if an object is active or inactive, check the [activeSelf](../ScriptReference/GameObject-activeSelf.md) property.

If you deactivate a GameObject, coroutines attached to it are stopped.

## Deactivate a parent GameObject

If you deactivate a parent GameObject, you also deactivate all of its child GameObjects because the deactivation overrides the `activeSelf` setting on all child GameObjects. The child GameObjects return to their original state when you reactivate the parent.

To know if a child GameObject is active in your scene, use the [activeInHierarchy](../ScriptReference/GameObject-activeInHierarchy.md) property.

**Note:** The `activeSelf` property is not always accurate if you check a child GameObject because even if it is set to active, you might have set one of its parent GameObjects to inactive.

![The selected GameObject (Cube) is set as active, but remains inactive until you set its parent GameObject to active.](../uploads/Main/deactivating2.png)

The selected GameObject (Cube) is set as active, but remains inactive until you set its parent GameObject to active.