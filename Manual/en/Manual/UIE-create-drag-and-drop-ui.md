# Create a drag-and-drop UI inside a custom Editor window

**Version**: 2021.3+

Drag-and-drop is a common feature in UI design. You can use UI Toolkit to create a drag-and-drop UI inside a custom Editor window or inside an application built by Unity. This example demonstrates how to create a drag-and-drop UI inside a custom Editor window.

## Example overview

The example adds several slots and one object in a custom Editor window. You can drag the object into any slot, as shown below:

![A preview of a drag-and-drop UI](../uploads/Main/UIBuilder/drag-drop.gif)

A preview of a drag-and-drop UI

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-drag-and-drop-window-inside-a-custom-editor-window).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [Pointer events](UIE-Pointer-Events.md)

## Create a custom Editor window

To start, create a default custom Editor window from the menu. Change the menu name and window title to `Drag And Drop`, and remove the code for the default labels, to make the UI more user-friendly.

1. Create a project in Unity with any template.
2. Right-click in the `Assets` folder and select **Create** > **UI Toolkit** > **Editor Window**.
3. In **UI Toolkit Editor Window Creator**, enter `DragAndDropWindow`.
4. Click **Confirm**. This automatically creates three files: `DragAndDropWindow.cs`, `DragAndDropWindow.uxml`, and `DragAndDropWindow.uss`.
5. Replace the content of `DragAndDropWindow.cs` with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;
   using UnityEditor.UIElements;

   public class DragAndDropWindow : EditorWindow
   {
       [MenuItem("Window/UI Toolkit/Drag And Drop")]
       public static void ShowExample()
       {
           DragAndDropWindow wnd = GetWindow<DragAndDropWindow>();
           wnd.titleContent = new GUIContent("Drag And Drop");
       }

       public void CreateGUI()
       {
           // Each editor window contains a root VisualElement object
           VisualElement root = rootVisualElement;

           // Import UXML
           var visualTree = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/Editor/DragAndDropWindow.uxml");
           VisualElement labelFromUXML = visualTree.Instantiate();
           root.Add(labelFromUXML);

           // A stylesheet can be added to a VisualElement.
           // The style will be applied to the VisualElement and all of its children.
           var styleSheet = AssetDatabase.LoadAssetAtPath<StyleSheet>("Assets/Editor/DragAndDropWindow.uss");
       }
   }
   ```

## Create slots and objects

Next, add UI controls to your custom window:

* One named `slots` with two children named `slot_row1` and `slot_row2`. Each row should have two children named `slot1` and `slot2`.
* One named `object` at the same level as `slots`. `object` must come after `slots` in the **Hierarchy**.

Style the UI controls as the following:

* For `slot1` and `slot2`, style them as 80px X 80px squares with white background color and rounded corners. Line up the slots as two rows with two slots in each row.
* For `object`, style it as a 50px X 50px round spot with a black background color.

**Tip**: To make your project more fun, you can use a background image for the object. You can find the image (Pouch.png) in the [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/blob/master/create-a-drag-and-drop-window-inside-a-custom-editor-window/Pouch.png).

1. Replace the content of `DragAndDropWindow.uxml` with the following:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
       <Style src="DragAndDropWindow.uss" />
       <ui:VisualElement name="slots">
           <ui:VisualElement name="slot_row1" class="slot_row">
               <ui:VisualElement name="slot1" class="slot" />
               <ui:VisualElement name="slot2" class="slot" />
           </ui:VisualElement>
           <ui:VisualElement name="slot_row2" class="slot_row">
               <ui:VisualElement name="slot1" class="slot" />
               <ui:VisualElement name="slot2" class="slot" />
           </ui:VisualElement>
       </ui:VisualElement>
       <ui:VisualElement name="object" class="object" />
   </ui:UXML>
   ```
2. Replace the content of `DragAndDropWindow.uss` with the following:

   ```
   .slot {
   width: 80px;
   height: 80px;
   margin: 5px;
   background-color: rgb(255, 255, 255);
   border-top-radius: 10px;
   border-top-left-radius: 10px;
   border-bottom-left-radius: 10px;
   border-top-right-radius: 10px;
   border-bottom-right-radius: 10px;
   }

   .slot_row {
       flex-direction: row;
   }

   .object {
       width: 50px;
       height: 50px;
       position: absolute;
       left: 20px;
       top: 20px;
       border-radius: 30px;
       background-color: rgb(0, 0, 0);
   }
   ```

## Define the drag-and-drop logic

To define the drag-and-drop behavior, extend the [`PointerManipulator`](../ScriptReference/UIElements.PointerManipulator.md) class and define the logic. Write a constructor to set [`target`](../ScriptReference/UIElements.Manipulator-target.md) and store a reference to the root of the visual tree. Write four methods that act as callbacks for [`PointerDownEvent`](../ScriptReference/UIElements.PointerDownEvent.md)s, [`PointerMoveEvent`](../ScriptReference/UIElements.PointerMoveEvent.md)s, [`PointerUpEvent`](../ScriptReference/UIElements.PointerUpEvent.md)s, and [`PointerCaptureOutEvent`](../ScriptReference/UIElements.PointerCaptureOutEvent.md)s. Implement [`RegisterCallbacksOnTarget()`](../ScriptReference/UIElements.Manipulator.RegisterCallbacksOnTarget.md) and [`UnregisterCallbacksOnTarget()`](../ScriptReference/UIElements.Manipulator.UnregisterCallbacksFromTarget.md) to register and un-register these four callbacks from `target`.

1. In the `Editor` folder, create another C# file named `DragAndDropManipulator.cs`.
2. Replace the content of `DragAndDropManipulator.cs` with the following:

   ```
   using System.Collections;
   using System.Collections.Generic;
   using UnityEngine;
   using UnityEngine.UIElements;

   public class DragAndDropManipulator : PointerManipulator
   {
       // Write a constructor to set target and store a reference to the
       // root of the visual tree.
       public DragAndDropManipulator(VisualElement target)
       {
           this.target = target;
           root = target.parent;
       }

       protected override void RegisterCallbacksOnTarget()
       {
           // Register the four callbacks on target.
           target.RegisterCallback<PointerDownEvent>(PointerDownHandler);
           target.RegisterCallback<PointerMoveEvent>(PointerMoveHandler);
           target.RegisterCallback<PointerUpEvent>(PointerUpHandler);
           target.RegisterCallback<PointerCaptureOutEvent>(PointerCaptureOutHandler);
       }

       protected override void UnregisterCallbacksFromTarget()
       {
           // Un-register the four callbacks from target.
           target.UnregisterCallback<PointerDownEvent>(PointerDownHandler);
           target.UnregisterCallback<PointerMoveEvent>(PointerMoveHandler);
           target.UnregisterCallback<PointerUpEvent>(PointerUpHandler);
           target.UnregisterCallback<PointerCaptureOutEvent>(PointerCaptureOutHandler);
       }

       private Vector2 targetStartPosition { get; set; }

       private Vector3 pointerStartPosition { get; set; }

       private bool enabled { get; set; }

       private VisualElement root { get; }

       // This method stores the starting position of target and the pointer,
       // makes target capture the pointer, and denotes that a drag is now in progress.
       private void PointerDownHandler(PointerDownEvent evt)
       {
           targetStartPosition = target.transform.position;
           pointerStartPosition = evt.position;
           target.CapturePointer(evt.pointerId);
           enabled = true;
       }

       // This method checks whether a drag is in progress and whether target has captured the pointer.
       // If both are true, calculates a new position for target within the bounds of the window.
       private void PointerMoveHandler(PointerMoveEvent evt)
       {
           if (enabled && target.HasPointerCapture(evt.pointerId))
           {
               Vector3 pointerDelta = evt.position - pointerStartPosition;

               target.transform.position = new Vector2(
                   Mathf.Clamp(targetStartPosition.x + pointerDelta.x, 0, target.panel.visualTree.worldBound.width),
                   Mathf.Clamp(targetStartPosition.y + pointerDelta.y, 0, target.panel.visualTree.worldBound.height));
           }
       }

       // This method checks whether a drag is in progress and whether target has captured the pointer.
       // If both are true, makes target release the pointer.
       private void PointerUpHandler(PointerUpEvent evt)
       {
           if (enabled && target.HasPointerCapture(evt.pointerId))
           {
               target.ReleasePointer(evt.pointerId);
           }
       }

       // This method checks whether a drag is in progress. If true, queries the root
       // of the visual tree to find all slots, decides which slot is the closest one
       // that overlaps target, and sets the position of target so that it rests on top
       // of that slot. Sets the position of target back to its original position
       // if there is no overlapping slot.
       private void PointerCaptureOutHandler(PointerCaptureOutEvent evt)
       {
           if (enabled)
           {
               VisualElement slotsContainer = root.Q<VisualElement>("slots");
               UQueryBuilder<VisualElement> allSlots =
                   slotsContainer.Query<VisualElement>(className: "slot");
               UQueryBuilder<VisualElement> overlappingSlots =
                   allSlots.Where(OverlapsTarget);
               VisualElement closestOverlappingSlot =
                   FindClosestSlot(overlappingSlots);
               Vector3 closestPos = Vector3.zero;
               if (closestOverlappingSlot != null)
               {
                   closestPos = RootSpaceOfSlot(closestOverlappingSlot);
                   closestPos = new Vector2(closestPos.x - 5, closestPos.y - 5);
               }
               target.transform.position =
                   closestOverlappingSlot != null ?
                   closestPos :
                   targetStartPosition;

               enabled = false;
           }
       }

       private bool OverlapsTarget(VisualElement slot)
       {
           return target.worldBound.Overlaps(slot.worldBound);
       }

       private VisualElement FindClosestSlot(UQueryBuilder<VisualElement> slots)
       {
           List<VisualElement> slotsList = slots.ToList();
           float bestDistanceSq = float.MaxValue;
           VisualElement closest = null;
           foreach (VisualElement slot in slotsList)
           {
               Vector3 displacement =
                   RootSpaceOfSlot(slot) - target.transform.position;
               float distanceSq = displacement.sqrMagnitude;
               if (distanceSq < bestDistanceSq)
               {
                   bestDistanceSq = distanceSq;
                   closest = slot;
               }
           }
           return closest;
       }

       private Vector3 RootSpaceOfSlot(VisualElement slot)
       {
           Vector2 slotWorldSpace = slot.parent.LocalToWorld(slot.layout.position);
           return root.WorldToLocal(slotWorldSpace);
       }
   }
   ```

## Instantiate the drag-and-drop behavior

To enable drag-and-drop in the custom window, instantiate it when the window opens.

1. In `DragAndDropWindow.cs`, add the following to the `CreateGUI()` method to instantiate the `DragAndDropManipulator` class:

   ```
   DragAndDropManipulator manipulator =
       new(rootVisualElement.Q<VisualElement>("object"));
   ```
2. From the menu bar, select **Window** > **UI Toolkit** > **Drag And Drop**. In the opened custom Editor window, you can drag the object into any slot.

## Additional resources

* [Create a drag-and-drop UI to drag between Editor windows](UIE-drag-across-windows.md)
* [Manipulators](UIE-manipulators.md)