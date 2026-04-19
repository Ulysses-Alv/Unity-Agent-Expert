# AppCallbacks class reference

You can use the `AppCallbacks` class to connect your main application to the Unity engine.

## Example: How to use the `AppCallbacks` class

### `App.xaml.cpp` file

```
App::App()
{
    InitializeComponent();
    SetupOrientation();
    m_AppCallbacks = ref new AppCallbacks();
}

void App::OnLaunched(LaunchActivatedEventArgs^ e)
{
    m_SplashScreen = e->SplashScreen;
    InitializeUnity(e->Arguments);
}

void App::InitializeUnity(String^ args)
{
    ApplicationView::GetForCurrentView()->SuppressSystemOverlays = true;

    m_AppCallbacks->SetAppArguments(args);
    auto rootFrame = safe_cast<Frame^>(Window::Current->Content);

    // Do not repeat app initialization when the Window already has content,
    // just ensure that the window is active
    if (rootFrame == nullptr && !m_AppCallbacks->IsInitialized())
    {
        rootFrame = ref new Frame();
        Window::Current->Content = rootFrame;
#if !UNITY_HOLOGRAPHIC
        Window::Current->Activate();
#endif

        rootFrame->Navigate(TypeName(MainPage::typeid ));
    }

    Window::Current->Activate();
}
```

### `MainPage.xaml.cpp` file

```
MainPage::MainPage()
{
    m_SplashScreenRemovalEventToken.Value = 0;
    m_OnResizeRegistrationToken.Value = 0;

    InitializeComponent();
    NavigationCacheMode = ::NavigationCacheMode::Required;

    auto appCallbacks = AppCallbacks::Instance;

    bool isWindowsHolographic = false;

#if UNITY_HOLOGRAPHIC
    // If application was exported as Holographic check if the device actually supports it
    // Otherwise, we treat this as a normal XAML application
    isWindowsHolographic = AppCallbacks::IsMixedRealitySupported();
#endif

    if (isWindowsHolographic)
    {
        appCallbacks->InitializeViewManager(Window::Current->CoreWindow);
    }
    else
    {
        m_SplashScreenRemovalEventToken = appCallbacks->RenderingStarted += ref new RenderingStartedHandler(this, &MainPage::RemoveSplashScreen);

        appCallbacks->SetSwapChainPanel(m_DXSwapChainPanel);
        // Subscribes to all needed system events
        appCallbacks->SetCoreWindowEvents(Window::Current->CoreWindow);

        // This is the main initialization function for Unity
        // Initializes engine graphics, DirectX, and gamepad and joystick input
        // Loads IL2CPP and all engine subsystems except graphics
        appCallbacks->InitializeD3DXAML();
        // At this point, when Unity finishes loading the first level, it enters the main loop.

        m_SplashScreen = safe_cast<App^>(App::Current)->GetSplashScreen();

        auto dispatcher = CoreWindow::GetForCurrentThread()->Dispatcher;
        ThreadPool::RunAsync(ref new WorkItemHandler([this, dispatcher](IAsyncAction^)
        {
            GetSplashBackgroundColor(dispatcher);
        }));

        OnResize();

        m_OnResizeRegistrationToken = Window::Current->SizeChanged += ref new WindowSizeChangedEventHandler([this](Object^, WindowSizeChangedEventArgs^)
        {
            OnResize();
        });
    }
}
```

## Create an app thread

Unity doesn’t run your app on the UI thread because the UI could become unresponsive when loading large apps. For more information on UI threads, refer to Microsoft documentation on [Keeping the UI thread responsive](https://learn.microsoft.com/en-gb/windows/uwp/debug-test-perf/keep-the-ui-thread-responsive).

When you create the `AppCallbacks` class using `m_AppCallbacks = ref new AppCallbacks();`, Unity creates a new thread called `App Thread`. Unity creates this new thread due to a Microsoft restriction: if your application doesn’t become responsive after 5 seconds, you’ll fail to pass the Windows App Certification Kit tests. For more information, refer to Microsoft documentation on the [Windows App Certification Kit](https://learn.microsoft.com/en-us/windows/uwp/debug-test-perf/windows-app-certification-kit).

**Note:** Code located in the `App.xaml.cpp` and `MainPage.xaml.cpp` files always runs on the UI thread, unless called from the `InvokeOnAppThread` function.

## Command line arguments

You can pass custom command line arguments as string arrays into the AppCallbacks constructor. For more information, refer to [UWP Command line arguments](windowsstore-commandlinearguments.md).

## AppCallbacks functions

| **Function** | **Description** |
| --- | --- |
| `appCallbacks->InitializeD3DXAML();` | Initializes your DirectX 11 device and loads the first level. |
| `appCallbacks->SetCoreWindowEvents(Window::Current->CoreWindow);` | Sets the core window for Unity. Unity subscribes to the following system events:  - VisibilityChanged - Closed - PointerCursor - SizeChanged - Activated - CharacterReceived - PointerPressed - PointerReleased - PointerMoved - PointerCaptureLost - PointerWheelChanged  - AcceleratorKeyActivated |
| `appCallbacks->SetSwapChainPanel(m_DXSwapChainPanel);` | Passes a XAML control to Unity which is used as a render target for DirectX 11. |
| `void GetSwapChainPanel()` | Returns the SwapChainPanel object, which you can set via the SetSwapChainPanel method. |
| `void Initialized()` | Returns whether the engine is initialized enough to run the main game loop. |
| `void InitializeD3DWindow()` | Initializes engine graphics, DirectX, and gamepad and joystick input for D3D applications. |
| `void Instance()` | Retrieves a singleton instance of a previously created AppCallbacks object. |
| `void InvokeOnAppThread(AppCallbackItem item, bool waitUntilDone)` | Invokes a delegate on the application thread. This function is useful when you want to execute your script function from a UI thread. |
| `void InvokeOnUIThread(AppCallbackItem item, bool waitUntilDone)` | Invokes a delegate on the UI thread. This function is useful when you want to invoke an XAML-specific API from your **scripts**. |
| `bool IsInitialized()` | Returns true when the first level of your application is fully loaded. |
| `void RenderingStarted()` | Starts after Unity renders its first frame. |
| `void Run()` | Enables D3D applications to enter the main loop. |
| `bool RunningOnAppThread()` | Returns true if you’re currently running in an application thread. |
| `bool RunningOnUIThread()` | Returns true if you’re currently running in a UI thread. |
| `void SetAppArguments(string arg)` / `string GetAppArguments()` | Sets your application arguments, which you can then access from [UnityEngine.WSA.Application.arguments](../ScriptReference/WSA.Application-arguments.md). |
| `void SetCoreApplicationViewEvents()` | Subscribes to the CoreApplicationView::Activated event and loads the **IL2CPP** **scripting backend** and all engine subsystems except graphics. |
| `bool UnityGetInput()` | Returns true if Unity processes incoming input. |
| `void UnitySetInput(bool enabled)` | Enables or disables input processing. |
| `bool UnityPause(int pause)` | Pauses Unity if you pass 1 and unpauses if you pass 0. This function is useful if you want to temporarily freeze your game. |

## Additional resources

* [Microsoft documentation on Keeping the UI thread responsive](https://learn.microsoft.com/en-gb/windows/uwp/debug-test-perf/keep-the-ui-thread-responsive)
* [UWP Command line arguments](windowsstore-commandlinearguments.md)
* [Command line arguments](PlayerCommandLineArguments.md)