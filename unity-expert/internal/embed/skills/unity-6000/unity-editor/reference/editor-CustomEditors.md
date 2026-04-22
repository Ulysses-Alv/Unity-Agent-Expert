# Create custom Editors with IMGUI

**Note**: It’s strongly recommended to use the [UI Toolkit](UIElements.md) to extend the [Unity Editor](UIE-support-for-editor-ui.md), as it provides a more modern, flexible, and scalable solution than IMGUI.

To speed up application development, create custom editors for components you commonly use. This page shows you how to create a simple script to make **GameObjects** always look at a point.

1. Create a C# script and name it “LookAtPoint”.
2. Open the script and replace its contents with the code below.
3. Attach the script to a GameObject in your **Scene**.

```
using UnityEngine;
public class LookAtPoint : MonoBehaviour
{
    public Vector3 lookAtPoint = Vector3.zero;

    void Update()
    {
        transform.LookAt(lookAtPoint);
    }
}
```

When you enter Play mode, the GameObject that you attached the script to now orientates itself towards the coordinates you set to the “Look At Point” property. When writing Editor **scripts**, it’s often useful to make certain scripts execute during **Edit mode**, while your application is not running. To do this, add the `ExecuteInEditMode` attribute to the class, like this:

```
using UnityEngine;
[ExecuteInEditMode]
public class LookAtPoint : MonoBehaviour
{
    public Vector3 lookAtPoint = Vector3.zero;

    void Update()
    {
        transform.LookAt(lookAtPoint);
    }
}
```

Now if you move the GameObject in the Editor, or change the values of “Look At Point” in the **Inspector**, the GameObject updates its rotation so that it looks at the target point in world space.

### Making a Custom Editor

The above demonstrates how you can get simple scripts running during edit-time, however this alone does not allow you to create your own Editor tools. The next step is to create a **Custom Editor** for the script you just created.

When you create a script in Unity, by default it inherits from MonoBehaviour, and therefore is a component that you can attach to a GameObject. When you place a component on a GameObject, the Inspector displays a default interface that you can use to view and edit every public variable, for example: an integer, a float, or a string.

This is how the Inspector for the LookAtPoint component looks by default:

![A default Inspector with a public Vector3 field](../uploads/Main/NoCustomInspector.png)

A default Inspector with a public Vector3 field

**A custom editor is a separate script which *replaces* this default layout with any editor controls that you choose.**

To create the custom editor for the LookAtPoint script:

1. Create a new C# script and name it “LookAtPointEditor”.
2. Open the script and replace its contents with the code below.

```
using UnityEngine;
using UnityEditor;

[CustomEditor(typeof(LookAtPoint))]
[CanEditMultipleObjects]
public class LookAtPointEditor : Editor 
{
    SerializedProperty lookAtPoint;
    
    void OnEnable()
    {
        lookAtPoint = serializedObject.FindProperty("lookAtPoint");
    }

    public override void OnInspectorGUI()
    {
        serializedObject.Update();
        EditorGUILayout.PropertyField(lookAtPoint);
        serializedObject.ApplyModifiedProperties();
    }
}
```

This class must inherit from **Editor**. The **CustomEditor** attribute informs Unity which component it should act as an editor for. The **CanEditMultipleObjects** attribute tells Unity that you can select multiple objects with this editor and change them all at the same time.

Unity executes the code in OnInspectorGUI it displays the editor in the Inspector. You can put any GUI code in here and it works in the same way as OnGUI does, but runs inside the Inspector. Editor defines the target property that you can use to access the GameObject you are inspecting.

This is how the Inspector for the LookAtPoint component looks with the new editor:

![The Look At Point component in the new editor](../uploads/Main/CustomInspector.png)

The Look At Point component in the new editor

This looks very similar (although the “Script” field is now not present, because the editor script does not add any Inspector code to show it).

However now that you have control over how the Inspector displays in an Editor script, you can use any code you like to lay out the Inspector fields, allow the user to adjust the values, and even display graphics or other **visual elements**. In fact all of the Inspectors you see within the Unity Editor including the more complex Inspectors such as the **terrain** system and animation import settings, are all made using the same API that you have access to when creating your own custom Editors.

Here is a simple example which extends your editor script to display a message that indicates whether the target point is above or below the GameObject:

```
using UnityEngine;
using UnityEditor;

[CustomEditor(typeof(LookAtPoint))]
[CanEditMultipleObjects]
public class LookAtPointEditor : Editor
{
    SerializedProperty lookAtPoint;

    void OnEnable()
    {
        lookAtPoint = serializedObject.FindProperty("lookAtPoint");
    }

    public override void OnInspectorGUI()
    {
        serializedObject.Update();
        EditorGUILayout.PropertyField(lookAtPoint);
        serializedObject.ApplyModifiedProperties();
        if (lookAtPoint.vector3Value.y > (target as LookAtPoint).transform.position.y)
        {
            EditorGUILayout.LabelField("(Above this object)");
        }
        if (lookAtPoint.vector3Value.y < (target as LookAtPoint).transform.position.y)
        {
            EditorGUILayout.LabelField("(Below this object)");
        }
    }
}
```

This is how the Inspector for the LookAtPoint component looks with the message showing if the target point is above or below the GameObject.

![The component in the Inspector now includes information about the targets location relative to the GameObject](../uploads/Main/CustomInspector2.png)

The component in the Inspector now includes information about the target’s location relative to the GameObject

You have full access to all the IMGUI commands to draw any type of interface, including rendering Scenes using a **Camera** within Editor windows.

### Scene View Additions

You can add extra code to the **Scene View**. To do this, implement OnSceneGUI in your custom editor.

OnSceneGUI works just like OnInspectorGUI except it runs in the Scene view. To help you make your own editing controls, you can use the functions defined in the [Handles](../ScriptReference/Handles.md) class. All functions in there are designed for working in 3D Scene views.

```
using UnityEngine;
using UnityEditor;

[CustomEditor(typeof(LookAtPoint))]
[CanEditMultipleObjects]
public class LookAtPointEditor : Editor
{
    SerializedProperty lookAtPoint;

    void OnEnable()
    {
        lookAtPoint = serializedObject.FindProperty("lookAtPoint");
    }

    public override void OnInspectorGUI()
    {
        serializedObject.Update();
        EditorGUILayout.PropertyField(lookAtPoint);
        if (lookAtPoint.vector3Value.y > (target as LookAtPoint).transform.position.y)
        {
            EditorGUILayout.LabelField("(Above this object)");
        }
        if (lookAtPoint.vector3Value.y < (target as LookAtPoint).transform.position.y)
        {
            EditorGUILayout.LabelField("(Below this object)");
        }
        
            
        serializedObject.ApplyModifiedProperties();
    }

    public void OnSceneGUI()
    {
        var t = (target as LookAtPoint);

        EditorGUI.BeginChangeCheck();
        Vector3 pos = Handles.PositionHandle(t.lookAtPoint, Quaternion.identity);
        if (EditorGUI.EndChangeCheck())
        {
            Undo.RecordObject(target, "Move point");
            t.lookAtPoint = pos;
            t.Update();
        }
    }
}
```

If you want to add 2D GUI objects (for example: GUI or EditorGUI), you need to wrap them in calls to Handles.BeginGUI() and Handles.EndGUI().