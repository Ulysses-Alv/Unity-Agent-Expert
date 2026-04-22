# Supporting presets for custom types

Use the [`ObjectFactory`](../ScriptReference/ObjectFactory.md) class in your Editor **scripts** to create new **GameObjects**, components, and assets that support and inherit presets by default. The `ObjectFactory` class automatically applies default presets to these items.

To support and enable presets by default for your custom classes, they must inherit from one of the following:

* [`UnityEngine.MonoBehaviour`](../ScriptReference/MonoBehaviour.md)
* [`UnityEngine.ScriptableObject`](../ScriptReference/ScriptableObject.md)
* [`UnityEngine.ScriptedImporter`](../ScriptReference/AssetImporters.ScriptedImporter.md)

The preset **Inspector** window creates a temporary instance of your class, so that users can modify its values. Therefore, make sure your class doesn’t affect or rely on other objects such as static values, project assets, or **scene** instances.

## Example: Preset settings in a custom Editor window

The following series of examples demonstrate how to add preset settings to a simple `EditorWindow`. When creating a custom [`EditorWindow`](../ScriptReference/EditorWindow.md) class with settings that can be saved to presets:

* Use a [`ScriptableObject`](../ScriptReference/ScriptableObject.md) to store a copy of your settings. Optionally, add a [`CustomEditor`](../ScriptReference/CustomEditor.md) attribute. The preset system handles this object.
* Always use this temporary `ScriptableObject` Inspector to show the preset settings in your UI. This allows your users to have the same UI in your `EditorWindow` as when editing saved presets.
* Expose a Preset button and use your own [`PresetSelectorReceiver`](../ScriptReference/Presets.PresetSelectorReceiver.md) implementation to keep your `EditorWindow` settings up-to-date when a Preset is selected in the **Select Preset** window.

The following example demonstrates a `ScriptableObject` that keeps and shows settings in a custom window:

```
using UnityEngine;

// Temporary ScriptableObject used by the Preset system

public class MyWindowSettings : ScriptableObject
{
    [SerializeField]
    string m_SomeSettings;
    
    public void Init(MyEditorWindow window)
    {
        m_SomeSettings = window.someSettings;
    }
    
    public void ApplySettings(MyEditorWindow window)
    {
        window.someSettings = m_SomeSettings;
        window.Repaint();
    }
}
```

The following example uses a `PresetSelectorReceiver` to update the `ScriptableObject` used in the custom window:

```
using UnityEditor.Presets;

// PresetSelector receiver to update the EditorWindow with the selected values.

public class MySettingsReceiver : PresetSelectorReceiver
{
    Preset initialValues;
    MyWindowSettings currentSettings;
    MyEditorWindow currentWindow;
    
    public void Init(MyWindowSettings settings, MyEditorWindow window)
    {
        currentWindow = window;
        currentSettings = settings;
        initialValues = new Preset(currentSettings);
    }
    
    public override void OnSelectionChanged(Preset selection)
    {
        if (selection != null)
        {
            // Apply the selection to the temporary settings
            selection.ApplyTo(currentSettings);
        }
        else
        {
            // None have been selected. Apply the Initial values back to the temporary selection.
            initialValues.ApplyTo(currentSettings);
        }
        
        // Apply the new temporary settings to our manager instance
        currentSettings.ApplySettings(currentWindow);
    }
    
    public override void OnSelectionClosed(Preset selection)
    {
        // Call selection change one last time to make sure you have the last selection values.
        OnSelectionChanged(selection);
        // Destroy the receiver here, so you don't need to keep a reference to it.
        DestroyImmediate(this);
    }
}
```

The following example creates an [`EditorWindow`](../ScriptReference/EditorWindow.md) that shows custom settings using a temporary ScriptableObject Inspector and its Preset button:

```
using UnityEngine;
using UnityEditor;
using UnityEditor.Presets;

public class MyEditorWindow : EditorWindow

{
    // get the Preset icon and a style to display it
    private static class Styles
    {
        public static GUIContent presetIcon = EditorGUIUtility.IconContent("Preset.Context");
        public static GUIStyle iconButton = new GUIStyle("IconButton");

    }

    Editor m_SettingsEditor;
    MyWindowSettings m_SerializedSettings;
    
    public string someSettings
    {
        get { return EditorPrefs.GetString("MyEditorWindow_SomeSettings"); }
        set { EditorPrefs.SetString("MyEditorWindow_SomeSettings", value); }
    }
   
    // Method to open the window
    [MenuItem("Window/MyEditorWindow")]
    static void OpenWindow()
    {
        GetWindow<MyEditorWindow>();
    }

    void OnEnable()
    {
        // Create your settings now and its associated Inspector
        // that allows to create only one custom Inspector for the settings in the window and the Preset.
        m_SerializedSettings = ScriptableObject.CreateInstance<MyWindowSettings>();
        m_SerializedSettings.Init(this);
        m_SerializedSettings.hideFlags = HideFlags.DontSave;
        m_SettingsEditor = Editor.CreateEditor(m_SerializedSettings);
        m_SettingsEditor.hideFlags = HideFlags.DontSave;
    }

    void OnDisable()
    {
        Object.DestroyImmediate(m_SerializedSettings);
        Object.DestroyImmediate(m_SettingsEditor);
    }

    void OnGUI()
    {
        EditorGUILayout.BeginHorizontal();
        EditorGUILayout.LabelField("My custom settings", EditorStyles.boldLabel);
        GUILayout.FlexibleSpace();
        // create the Preset button at the end of the "MyManager Settings" line.
        var buttonPosition = EditorGUILayout.GetControlRect(false, EditorGUIUtility.singleLineHeight, Styles.iconButton);

        if (EditorGUI.DropdownButton(buttonPosition, Styles.presetIcon, FocusType.Passive, Styles.iconButton))
        {
            // Create a receiver instance. This destroys itself when the window appears, so you don't need to keep a reference to it.
            var presetReceiver = ScriptableObject.CreateInstance<MySettingsReceiver>();
            presetReceiver.Init(m_SerializedSettings, this);
            // Show the PresetSelector modal window. The presetReceiver updates your data.
            PresetSelector.ShowSelector(m_SerializedSettings, null, true, presetReceiver);
        }
        EditorGUILayout.EndHorizontal();
        
        // Draw the settings default Inspector and catch any change made to it.
        EditorGUI.BeginChangeCheck();
        m_SettingsEditor.OnInspectorGUI();

        if (EditorGUI.EndChangeCheck())
        {
            // Apply changes made in the settings editor to our instance.
            m_SerializedSettings.ApplySettings(this);
        }
    }
}
```

## Additional resources

* [Creating and using presets](presets-creating-using.md)
* [Applying default presets to assets by folder](DefaultPresetsByFolder.md)