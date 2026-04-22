# Create a slide toggle custom control

**Version**: 2023.2+

This example demonstrates how to create a “switch-like” variation of a toggle.

## Example overview

The example creates a custom control: a toggle that users can flip with the mouse, the keyboard, a gamepad, and other devices. It comes with a label that can describe what the toggle represents.

![A toggle button which a mouse cursor selects, and the toggle slides to the right.](../uploads/Main/slide-toggle-finished.gif)

A toggle button which a mouse cursor selects, and the toggle slides to the right.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023.2/slide-toggle).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [Handling events](UIE-Events-Handling.md)

## Create the slide toggle class

Create the slide toggle class with a C# script.

1. Create a Unity project with the 3D template. A 3D template has a better visual effect for this example. However, you can use any template.
2. Create a folder named `slide-toggle` to store your files.
3. In the `slide-toggle` folder, create a C# script named `SlideToggle.cs`.
4. Open `SlideToggle.cs` in a text editor and replace its content with the following:

```
using UnityEngine;
using UnityEngine.UIElements;

namespace MyUILibrary
{
    // Derives from BaseField<bool> base class. Represents a container for its input part.
    [UxmlElement]
    public partial class SlideToggle : BaseField<bool>
    {
        // In the spirit of the BEM standard, the SlideToggle has its own block class and two element classes. It also
        // has a class that represents the enabled state of the toggle.
        public static readonly new string ussClassName = "slide-toggle";
        public static readonly new string inputUssClassName = "slide-toggle__input";
        public static readonly string inputKnobUssClassName = "slide-toggle__input-knob";
        public static readonly string inputCheckedUssClassName = "slide-toggle__input--checked";

        VisualElement m_Input;
        VisualElement m_Knob;

        // Custom controls need a default constructor. This default constructor calls the other constructor in this
        // class.
        public SlideToggle() : this(null) { }

        // This constructor allows users to set the contents of the label.
        public SlideToggle(string label) : base(label, null)
        {
            // Style the control overall.
            AddToClassList(ussClassName);

            focusable = false;

            // Get the BaseField's visual input element and use it as the background of the slide.
            m_Input = this.Q(className: BaseField<bool>.inputUssClassName);
            m_Input.AddToClassList(inputUssClassName);

            // Create a "knob" child element for the background to represent the actual slide of the toggle.
            m_Knob = new();
            m_Knob.AddToClassList(inputKnobUssClassName);
            m_Input.Add(m_Knob);

            // There are three main ways to activate or deactivate the SlideToggle. All three event handlers use the
            // static function pattern described in the Custom control best practices.

            // ClickEvent fires when a sequence of pointer down and pointer up actions occurs.
            RegisterCallback<ClickEvent>(evt => OnClick(evt));
            // KeydownEvent fires when the field has focus and a user presses a key.
            RegisterCallback<KeyDownEvent>(evt => OnKeydownEvent(evt));
            // NavigationSubmitEvent detects input from keyboards, gamepads, or other devices at runtime.
            RegisterCallback<NavigationSubmitEvent>(evt => OnSubmit(evt));
        }

        static void OnClick(ClickEvent evt)
        {
            var slideToggle = evt.currentTarget as SlideToggle;
            slideToggle.ToggleValue();

            evt.StopPropagation();
        }

        static void OnSubmit(NavigationSubmitEvent evt)
        {
            var slideToggle = evt.currentTarget as SlideToggle;
            slideToggle.ToggleValue();

            evt.StopPropagation();
        }

        static void OnKeydownEvent(KeyDownEvent evt)
        {
            var slideToggle = evt.currentTarget as SlideToggle;

            // NavigationSubmitEvent event already covers keydown events at runtime, so this method shouldn't handle
            // them.
            if (slideToggle.panel?.contextType == ContextType.Player)
                return;

            // Toggle the value only when the user presses Enter, Return, or Space.
            if (evt.keyCode == KeyCode.KeypadEnter || evt.keyCode == KeyCode.Return || evt.keyCode == KeyCode.Space)
            {
                slideToggle.ToggleValue();
                evt.StopPropagation();
            }
        }

        // All three callbacks call this method.
        void ToggleValue()
        {
            value = !value;
        }

        // Because ToggleValue() sets the value property, the BaseField class dispatches a ChangeEvent. This results in a
        // call to SetValueWithoutNotify(). This example uses it to style the toggle based on whether it's currently
        // enabled.
        public override void SetValueWithoutNotify(bool newValue)
        {
            base.SetValueWithoutNotify(newValue);

            //This line of code styles the input element to look enabled or disabled.
            m_Input.EnableInClassList(inputCheckedUssClassName, newValue);
        }
    }
}
```

## Style the slide toggle

Style the slide toggle with a USS file.

1. In the `slide-toggle` folder, create a USS file named `SlideToggle.uss`.
2. Open `SlideToggle.uss` in a text editor and replace its contents with the following:

```
.slide-toggle__input {
    background-color: var(--unity-colors-slider_groove-background);
    max-width: 25px;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;
    border-top-right-radius: 8px;
    border-bottom-right-radius: 8px;
    overflow: visible;
    border-left-width: 1px;
    border-right-width: 1px;
    border-top-width: 1px;
    border-bottom-width: 1px;
    border-right-color: var(--unity-colors-slider_thumb-border);
    border-top-color: var(--unity-colors-slider_thumb-border);
    border-bottom-color: var(--unity-colors-slider_thumb-border);
    max-height: 16px;
    margin-top: 10px;
    border-left-color: var(--unity-colors-slider_thumb-border);
    transition-property: background-color;
    transition-duration: 0.5s;
}

.slide-toggle__input-knob {
    height: 16px;
    width: 16px;
    background-color: var(--unity-colors-slider_thumb-background);
    position: absolute;
    border-top-left-radius: 25px;
    border-bottom-left-radius: 25px;
    border-top-right-radius: 25px;
    border-bottom-right-radius: 25px;
    top: -1px;
    transition-property: translate, background-color;
    transition-duration: 0.5s, 0.5s;
    translate: -1px 0;
    border-left-width: 1px;
    border-right-width: 1px;
    border-top-width: 1px;
    border-bottom-width: 1px;
    border-left-color: var(--unity-colors-slider_thumb-border);
    border-right-color: var(--unity-colors-slider_thumb-border);
    border-top-color: var(--unity-colors-slider_thumb-border);
    border-bottom-color: var(--unity-colors-slider_thumb-border);
}

.slide-toggle__input--checked {
    background-color: rgb(0, 156, 10);
}

.slide-toggle__input--checked > .slide-toggle__input-knob {
    translate: 8px 0;
}

.slide-toggle:focus .slide-toggle__input-knob {
    border-left-width: 1px;
    border-right-width: 1px;
    border-top-width: 1px;
    border-bottom-width: 1px;
    border-left-color: var(--unity-colors-input_field-border-focus);
    border-right-color: var(--unity-colors-input_field-border-focus);
    border-top-color: var(--unity-colors-input_field-border-focus);
    border-bottom-color: var(--unity-colors-input_field-border-focus);
}
```

## Add a slide toggle to a UXML file with the UI Builder

1. In the `slide-toggle` folder, create a UI Document file named `SlideToggleUsage.uxml`.
2. Open `SlideToggleUsage.uxml` in the UI Builder.
3. In the UI Builder, select **Library** > **Project** > **MyUILibrary**.
4. Drag a **SlideToggle** into the **Hierarchy** window.
5. In the ****Inspector****, enter a label text for your toggle in the **Label** field.
6. In the **StyleSheets** section, add `SlideToggle.uss` as the existing USS.
7. Save and close the UI Builder.

## Test the slide toggle

1. In the SampleScene, create a UI Document **GameObject**.
2. Select the UI Document, and drag `SlideToggleUsage.uxml` into the **Source Asset** field in the **Inspector**.
3. Enter Play mode. You can click on the toggle, press Enter, press Return, or press Space and get it to flip back and forth.

## Additional resources

* [Create custom controls](UIE-create-custom-controls.md)
* [Built-in Unity variables](UIE-USS-UnityVariables.md)
* [BEM standard](http://getbem.com/)