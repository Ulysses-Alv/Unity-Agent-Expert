# Optimize Metal graphics

Metal supports multiple optimizations that you can use to increase the performance of your application.

## Use memoryless render targets

Metal supports memoryless render targets which enables you to render to a [RenderTexture](../ScriptReference/RenderTexture.md) without backing it up in system memory. Unity only stores the contents temporarily in the on-tile memory during rendering.

Metal supports memoryless render targets on mobile devices from iOS 10.0, tvOS 10.0, visionOS 1.0, and on desktop devices (including Apple silicon) from macOS 11.

## Additional resources

* [RenderTexture.memorylessMode](../ScriptReference/RenderTexture-memorylessMode.md)