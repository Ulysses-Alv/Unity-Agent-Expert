# Bind a custom control

**Version**: 2021.3+

This example demonstrates how to bind a custom control to a native Unity type.

## Example overview

The example creates a custom control that displays a 2D image. It uses the custom control in a custom Editor for a new asset type, and it binds the custom control to a field in an asset of the new type.

![Custom control for a 2D image, displayed in the Unity Editor.](../uploads/Main/uie_bind_custom_control.png)

Custom control for a 2D image, displayed in the Unity Editor.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023.2/bind-custom-control).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [`BindableElement`](../ScriptReference/UIElements.BindableElement.md)
* [`INotifyValueChanged`](../ScriptReference/UIElements.INotifyValueChanged_1.md)

## Create a serialized object

Create a serialized object that contains a Texture2D.

1. Create a Unity project with any template.
2. In the **Project window**, create a folder named `bind-custom-control` to store the files.
3. Create a C# script named `TextureAsset.cs` and replace its contents with the following:

```
using UnityEngine;

namespace UIToolkitExamples
{
    [CreateAssetMenu(menuName = "UIToolkitExamples/TextureAsset")]
    public class TextureAsset : ScriptableObject
    {
        public Texture2D texture;

        public void Reset()
        {
            texture = null;
        }
    }
}
```

## Create and style a custom control

Create a custom control with C# that represents a reference to a 2D texture asset, and style it with USS.

1. In that folder, create a folder named `Editor`.
2. In the **Editor** folder, create a C# script named `TexturePreviewElement.cs`.
3. Replace the contents of `TexturePreviewElement.cs` with the following:

```
using System;
using UnityEditor.UIElements;
using UnityEngine;
using UnityEngine.UIElements;
using Object = UnityEngine.Object;

namespace UIToolkitExamples
{
    [UxmlElement]
    public partial class TexturePreviewElement : BindableElement, INotifyValueChanged<Object>
    {
        public static readonly string ussClassName = "texture-preview-element";

        Image m_Preview;
        ObjectField m_ObjectField;
        Texture2D m_Value;

        public TexturePreviewElement()
        {
            AddToClassList(ussClassName);

            // Create a preview image.
            m_Preview = new Image();
            Add(m_Preview);

            // Create an ObjectField, set its object type, and register a callback when its value changes.
            m_ObjectField = new ObjectField();
            m_ObjectField.objectType = typeof(Texture2D);
            m_ObjectField.RegisterValueChangedCallback(OnObjectFieldValueChanged);
            Add(m_ObjectField);

            styleSheets.Add(Resources.Load<StyleSheet>("texture_preview_element"));
        }
        
        void OnObjectFieldValueChanged(ChangeEvent<Object> evt)
        {
            value = evt.newValue;
        }

        public void SetValueWithoutNotify(Object newValue)
        {
            if (newValue == null || newValue is Texture2D)
            {
                // Update the preview Image and update the ObjectField.
                m_Value = newValue as Texture2D;
                m_Preview.image = m_Value;
                // Notice that this line calls the ObjectField's SetValueWithoutNotify() method instead of just setting
                // m_ObjectField.value. This is very important; you don't want m_ObjectField to send a ChangeEvent.
                m_ObjectField.SetValueWithoutNotify(m_Value);
            }
            else throw new ArgumentException($"Expected object of type {typeof(Texture2D)}");
        }

        public Object value
        {
            get => m_Value;
            // The setter is called when the user changes the value of the ObjectField, which calls
            // OnObjectFieldValueChanged(), which calls this.
            set
            {
                if (value == this.value)
                    return;

                var previous = this.value;
                SetValueWithoutNotify(value);

                using (var evt = ChangeEvent<Object>.GetPooled(previous, value))
                {
                    evt.target = this;
                    SendEvent(evt);
                }
            }
        }
    }
}
```

1. In the **Editor** folder, create a folder named `Resources`.
2. In the **Resources** folder, create a StyleSheet named `texture_preview_element.uss` and replace its contents with the following:

```
.texture-preview-element {
    width: 200px;
    height: 200px;
}

.texture-preview-element > .unity-image {
    flex-grow: 1;
}
```

## Create the Inspector UI

Create the **Inspector** UI, and bind the custom control to the Texture2D object.

1. In the **Editor** folder, create a C# script named `TextureAssetEditor.cs` and replace its contents with the following:

```
using UnityEditor;
using UnityEngine.UIElements;
using UnityEngine;

namespace UIToolkitExamples
{
    [CustomEditor(typeof(TextureAsset))]
    public class TextureAssetEditor : Editor
    {
        [SerializeField]
        VisualTreeAsset m_VisualTree;

        public override VisualElement CreateInspectorGUI()
        {
            return m_VisualTree.CloneTree();
        }
    }
}
```

## Set the binding with UXML

Create a UXML file with a TexturePreviewElement, and set the `binding-path` property to `texture`. This binds the TexturePreviewElement to `TextureAsset.texture`.

1. In the **Editor** folder, create a UI Document named `texture_asset_editor.uxml` and replace its contents with the following:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:example="UIToolkitExamples" editor-extension-mode="True">
    <example:TexturePreviewElement binding-path="texture" />
</ui:UXML>
```

1. In the Project window, select **TextureAssetEditor.cs**.
2. Drag **texture\_asset\_editor.uxml** to **Visual Tree** in the Inspector.

## Test the binding

1. From the menu, select **Assets** > **Create** > **UIToolkitExamples** > **TextureAsset**. This creates an instance of a `TextureAsset`.
2. Select the new TextureAsset object. In the **Inspector**, you can see the Texture Preview Element. If you put a texture asset reference into the field, you can see a preview above the field. If you change the asset reference in the Texture Preview Element in the UI, the `TextureAsset.texture` changes.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)