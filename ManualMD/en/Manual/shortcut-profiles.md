# Manage shortcut profiles

The Unity Editor saves keyboard shortcuts to profiles. You can create multiple profiles and move between them without restarting the Editor. With this, you can customize your shortcuts for different workflows, package sets, or custom tools.

Any changes you make to shortcuts are automatically saved to the active profile, except for the default profile.

## The default profile

The default profile is the profile that the Editor uses when you first install it. It’s also the profile that forms the basis for any new profiles you create, so your don’t have to assign every shortcut in a new profile - only those you want to change.

You can’t modify the default profile. If you try to modify it, the Editor automatically creates a copy of the default profile, names it **Default Copy**, makes that copy the active profile, and applies your change to that copy.

## Manage profiles

To manage profiles, from the **Shortcuts** window, select **Profile**. You can create, rename, and delete profiles and move between them.

## Reuse profiles

The Editor supports reusing profiles across projects and versions:

* Your shortcut profiles are available in all Unity projects because they’re saved to the Unity preferences directory, not to the Project directory.
* When you upgrade from an older version of the Editor, your saved shortcuts are available in the new version.

## Additional sources

* [Keyboard shortcut concepts](keyboard-shortcut-concepts.md)
* [Shortcuts window reference](shortcuts-view.md)
* [Manage keyboard shortcuts](shortcuts-manage.md)