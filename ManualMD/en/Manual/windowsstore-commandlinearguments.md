# Command line arguments for UWP

You can launch Unity Players from the command line and pass in arguments to change how the Player executes. Universal Windows Platform (UWP) apps don’t accept command line arguments by default, so you have to pass them to an AppCallbacks constructor in `App.xaml.cpp` or `App.cpp` to specify them.

This code example demonstrates how to do this:

```
m_AppCallbacks = 
    ref new AppCallbacks
    (
        ref new Platform::Array<Platform::String\^> 
        {
            L"-force-gfx-direct" 
        }
    );
```

For more information on command line arguments for UWP, refer to [Unity Standalone Player command line arguments](PlayerCommandLineArguments.md).