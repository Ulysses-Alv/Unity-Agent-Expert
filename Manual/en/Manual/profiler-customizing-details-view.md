# Create a Profiler module detail panel

The module details panel appears in the bottom of the [Profiler window](ProfilerWindow.md) when you select a module. You can customize this section to display additional details relevant to your module or to display a custom visualization of your performance data.

To create a module details panel for your Profiler module:

* [Create a module details panel Controller script](#create-details-view-controller) to draw a custom module details panel.
* [Create a Profiler module script](#connect-details-view-controller) to connect the controller to your profiler module .

## Create a script to control the module details panel

Use the [`ProfilerModuleViewController`](../ScriptReference/Unity.Profiling.Editor.ProfilerModuleViewController.md) base class to customize the module details panel in the Profiler window. To do this, create a script that controls what appears in the module details panel when you select a specific module.

The script must do the following:

* Define a public constructor for the view controller that calls the base constructor `base(profilerWindow)`.
* Override `CreateView` to build the custom module details panel.

For example:

```
 public class CustomDetailsViewController : ProfilerModuleViewController
 {   
    public CustomDetailsViewController(ProfilerWindow profilerWindow) : base(profilerWindow) { }

    protected override VisualElement CreateView()
    {
        // Create your UI.
    }
}
```

### Module details panel controller script example

The following example script creates a module details panel controller that draws a single label in the module details panel that displays text:

![A custom Profiler module with a custom message in the module details panel](../uploads/Main/Profiler_tank_details_basic.png)

A custom Profiler module with a custom message in the module details panel

The script example does the following:

* Defines and creates a label to display the value you want to capture, and adds this label to the module details panel.
* Defines a constructor to control the module details panel and uses `CreateView` to build the custom module details panel.
* Populates the label with data from the current frame, and updates the label after each frame.
* Fetches a counter value as a string that you can display in the module details panel.
* Specifies the text to display in the module details panel, and tells the Profiler to automatically update it each frame.

```
 using UnityEditor;
 using UnityEditorInternal;
 using Unity.Profiling.Editor;
 using UnityEngine.UIElements;
 
 public class TankEffectsDetailsViewController : ProfilerModuleViewController
 {
    // Define a label, which will display the total particle count for tank trails in the selected frame.
    Label m_TankTrailParticleCountLabel;

    // Define a constructor for the view controller, which calls the base constructor with the Profiler Window passed from the module.
    public TankEffectsDetailsViewController(ProfilerWindow profilerWindow) : base(profilerWindow)
    {
        // Be notified when the selected frame index in the Profiler Window changes, so we can update the label.
        ProfilerWindow.SelectedFrameIndexChanged += OnSelectedFrameIndexChanged;
    }

    protected override void ViewLoaded()
    { 
        base.ViewLoaded();
        
        // Populate the view with the current data for the selected frame.
        ReloadData(); 
    }
    
    protected override VisualElement CreateView()
    {
        var view = new VisualElement();
        
        // Create the label and add it to the view.
        m_TankTrailParticleCountLabel = new Label() { style = { paddingTop = 8, paddingLeft = 8 } };
        view.Add(m_TankTrailParticleCountLabel);
        
        return view;
    }
    

    // Override Dispose to do any cleanup of the view when it is destroyed. This is a standard C# Dispose pattern.
    protected override void Dispose(bool disposing)
    {
        if (!disposing)
            return;

        // Unsubscribe from the Profiler window event that we previously subscribed to.
        ProfilerWindow.SelectedFrameIndexChanged -= OnSelectedFrameIndexChanged;

        base.Dispose(disposing);
    }

    protected virtual void ReloadData()
    {
        // Retrieve the TankTrailParticleCount counter value from the Profiler as a formatted string.
        var selectedFrameIndexInt32 = System.Convert.ToInt32(ProfilerWindow.selectedFrameIndex);
        var value = ProfilerDriver.GetFormattedCounterValue(selectedFrameIndexInt32, GameStatistics.TanksCategory.Name, GameStatistics.TankTrailParticleCountName);

        // Update the label's text with the value.
        m_TankTrailParticleCountLabel.text = $"The value of '{GameStatistics.TankTrailParticleCountName}' in the selected frame is {value}.";
    }

    void OnSelectedFrameIndexChanged(long selectedFrameIndex)
    {
        // Update the label with the current data for the newly selected frame.
        ReloadData();
    }
}
```

**Tip:** You can use Unity’s UI Toolkit to build a custom UI for the module details panel. For more information, refer to [UI Toolkit](UIElements.md).

The following example image shows a custom module details panel that belongs to a custom [Adaptive Performance](adaptive-performance/adaptive-performance.md) module:

![A custom Profiler module with a custom UI visualization.](../uploads/Main/Adaptive_performance_Profiler.png)

A custom Profiler module with a custom UI visualization.

## Connect a custom module details panel to a Profiler module

To display a custom module details panel, you need to instantiate the module details panel controller when you select your Profiler module. To do this, override `CreateDetailsViewController` to create and draw a new module details panel controller. Unity then invokes this method when it displays your module’s details panel.

The following code example instantiates a custom module details panel for a module called `TankEffectsProfilerModule`:

```
 using Unity.Profiling.Editor;

 [System.Serializable]
 [ProfilerModuleMetadata("Tank Effects")]
 public class TankEffectsProfilerModule : ProfilerModule
 {
    static readonly ProfilerCounterDescriptor[] k_Counters = new ProfilerCounterDescriptor[]
    {
        new ProfilerCounterDescriptor(GameStatistics.TankTrailParticleCountName, GameStatistics.TanksCategory),
        new ProfilerCounterDescriptor(GameStatistics.ShellExplosionParticleCountName, GameStatistics.TanksCategory),
        new ProfilerCounterDescriptor(GameStatistics.TankExplosionParticleCountName, GameStatistics.TanksCategory),
    };

    public TankEffectsProfilerModule() : base(k_Counters) { }

    public override ProfilerModuleViewController CreateDetailsViewController()
    {
        return new TankEffectsDetailsViewController(ProfilerWindow);
    }
}
```

## Visualize additional counters in the module details panel

You can display additional [profiler counters](profiler-adding-information-code-intro.md) that aren’t included in your module’s chart view. This is useful when you want to display additional data for a selected frame.

The Profiler automatically captures the categories of all counters that belong to a module’s chart view when that module is active. To capture additional counters, write a script to tell the Profiler to capture specific categories when your module is active.

For example, the following script uses the `autoEnabledCategoryNames` constructor argument to specify the `Scripts` and `Memory` categories. The script enables these categories when the module is active:

```
using Unity.Profiling;
using Unity.Profiling.Editor;

[System.Serializable]
[ProfilerModuleMetadata("Tank Effects & Memory")]
public class TankEffectsAndMemoryProfilerModule : ProfilerModule
{
   static readonly ProfilerCounterDescriptor[] k_Counters = new ProfilerCounterDescriptor[]
   {
       new ProfilerCounterDescriptor(GameStatistics.TankTrailParticleCountName, ProfilerCategory.Scripts),
       new ProfilerCounterDescriptor(GameStatistics.ShellExplosionParticleCountName, ProfilerCategory.Scripts),
       new ProfilerCounterDescriptor(GameStatistics.TankExplosionParticleCountName, ProfilerCategory.Scripts),
   };

   // Enable the ProfilerCategory.Scripts and ProfilerCategory.Memory categories when the module is active.
   static readonly string[] k_AutoEnabledCategoryNames = new string[]
   {
       ProfilerCategory.Scripts.Name,
       ProfilerCategory.Memory.Name
   };

   public override ProfilerModuleViewController CreateDetailsViewController()
   {
       return new TankEffectsAndMemoryDetailsViewController(ProfilerWindow);
   }
   
   // Pass the auto-enabled category names to the base constructor.
   public TankEffectsProfilerModule() : base(k_Counters, autoEnabledCategoryNames: k_AutoEnabledCategoryNames) { }
}
```

The following example code displays the built-in [Mesh Memory profiler counter](profiler-counters-reference.md#memory) alongside `TankTrailParticleCount`:

```
 using UnityEditor;
 using UnityEditorInternal;
 using Unity.Profiling.Editor;
 using UnityEngine.UIElements;
 
 public class TankEffectsAndMemoryDetailsViewController : TankEffectsDetailsViewController
 {
    // Define a label, which will display the total mesh memory in the selected frame
    Label m_MeshMemoryLabel;
    
    protected override VisualElement CreateView()
    {
        var view = base.CreateView();
        
        // Create the label and add it to the view.
        m_MeshMemoryLabel = new Label() { style = { paddingTop = 8, paddingLeft = 8 } };
        view.Add(m_MeshMemoryLabel);

        return view;
    }
    
    protected override void ReloadData()
    {
        base.ReloadData();
        
        // Retrieve the Mesh Memory counter value from the Profiler as a formatted string.
        var selectedFrameIndexInt32 = System.Convert.ToInt32(ProfilerWindow.selectedFrameIndex);
        var value = ProfilerDriver.GetFormattedCounterValue(selectedFrameIndexInt32, ProfilerArea.Memory, "Mesh Memory");

        // Update the label's text with the value.
        m_MeshMemoryLabel.text = $"The value of 'Mesh Memory' in the selected frame is {value}.";
    }
}
```

## Additional resources

* [Profiler modules introduction](profiler-modules-introduction.md)
* [Profiler Module Editor reference](profiler-module-editor.md)
* [Creating Profiler module detail panels](profiler-customizing-details-view.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)