# Enable or disable asynchronous shader compilation

Asynchronous **shader** compilation is enabled by default.

To enable or disable asynchronous shader compilation:

1. Go to **Edit > Project Settings > Editor**.
2. At the bottom of the Editor settings, under **Shader Compilation**, check or uncheck the **Asynchronous Shader Compilation** checkbox.

**Note:** Enabling and disabling asynchronous shader compilation in this way affects only the **Scene** and Game views by default. If you want to use it in other parts of the Editor, see [Custom Editor tools and asynchronous shader compilation](#using-it-in-custom-editor-tools).

## Enable or disable asynchronous shader compilation for specific rendering calls

You can enable or disable asynchronous shader compilation for specific rendering commands in your C# **scripts**.

The following instructions show you how to enable or disable the feature in an immediate scope, and a [CommandBuffer](../ScriptReference/Rendering.CommandBuffer.md) scope.

#### In an immediate scope

In an immediate scope, you can use [`ShaderUtil.allowAsyncCompilation`](../ScriptReference/ShaderUtil-allowAsyncCompilation.md).

To do this:

1. Store the current state of `ShaderUtil.allowAsyncCompilation` in a variable.
2. Before you call the rendering commands, set `ShaderUtil.allowAsyncCompilation` to `false`.
3. Call the rendering commands.
4. After calling the rendering commands, restore `ShaderUtil.allowAsyncCompilation` back to its previous state.

Here is a pseudo-code example:

```
// Store the current state
bool oldState = ShaderUtil.allowAsyncCompilation;

// Disable async compilation
ShaderUtil.allowAsyncCompilation = false;

// Enter your rendering code that should never use the placeholder shader, for example UI elements or characters.
Graphics.RenderMesh(...);

// Restore the old state
ShaderUtil.allowAsyncCompilation = oldState;
```

#### In a CommandBuffer scope

In a `CommandBuffer` scope, you can use [`ShaderUtil.SetAsyncCompilation`](../ScriptReference/ShaderUtil.SetAsyncCompilation.md) and [`ShaderUtil.RestoreAsyncCompilation`](../ScriptReference/ShaderUtil.RestoreAsyncCompilation.md).

1. Immediately before you call the rendering commands, call `ShaderUtil.SetAsyncCompilation`, and set it to `false`. Subsequent commands in the CommandBuffer won’t allow asynchronous compilation.
2. Add the rendering commands to the CommandBuffer.
3. After the rendering commands, call [`Shader.Util.RestoreAsyncCompilation`](../ScriptReference/ShaderUtil.RestoreAsyncCompilation.md) to restore the state of asynchronous shader compilation.

Here is an example:

```
// Create the CommandBuffer
CommandBuffer cmd = new CommandBuffer();

// Disable async compilation for subsequent commands
ShaderUtil.SetAsyncCompilation(cmd, false);

/// Enter your rendering commands that should never use the placeholder shader, for example UI elements or characters.
cmd.DrawMesh(...);

// Restore the old state
ShaderUtil.RestoreAsyncCompilation(cmd);
```

## Disable asynchronous compilation for specific Shader objects

You can disable asynchronous shader compilation for specific **Shader objects** by forcing the Editor to always compile them synchronously. This is a good option for data generating Shader objects that are always present at the start of your rendering, and which are relatively quick to compile. You would most likely need this if you are performing [advanced rendering](AsynchronousShaderCompilation-avoid-cyan-placeholder-shaders.md).

To force synchronous compilation for a Shader object, add the `#pragma editor_sync_compilation` [directive](writing-shader-writing-shader-programs-hlsl.md) to your shader source code.

**Note:** You should not force synchronous compilation for complex Shader objects that encounter new shader variants during rendering; this can stall rendering in the Editor.

## Custom Editor tools and asynchronous shader compilation

By default, asynchronous Shader compilation works in the Game and **Scene views**. If you want to use it in custom Editor tools, you can enable it via C# for your custom tool.

To do this, you can [enable asynchronous shader compilation for specific rendering calls](#enabling-disabling-calls).