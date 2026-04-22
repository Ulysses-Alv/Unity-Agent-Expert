# Debug Metal graphics

Unity’s [Frame Debugger](FrameDebugger.md) supports Metal and should be the first tool you use to debug graphics issues in a Unity application that uses Metal. For more in-depth, Metal-specific graphics issues, [Xcode](https://developer.apple.com/xcode/) provides graphics debugging tools such as API validation and **shader** validation.

**Note**: Enabling validation increases CPU resource intensity, so only enable it for debugging purposes.

## Additional resources

* [Metal developer workflows](https://developer.apple.com/documentation/metal/diagnosing_metal_programming_issues_early)
* [Analyzing your Metal workload](https://developer.apple.com/documentation/metal/debugging_tools/viewing_your_gpu_workload_with_the_metal_debugger)