# Group commands with the Category block

Use the **Category** block to group commands that set the render state, so that you can “inherit” the grouped rendering state within the block.

For example, your **Shader** object might have multiple [SubShaders](SL-SubShader.md), each of which requires [blending](SL-Blend.md) set to additive. You can use the Category block for that:

```
Shader "example" {
Category {
    Blend One One
    SubShader {
        // ...
    }
    SubShader {
        // ...
    }
    // ...
}
}
```

The Category block does not have an impact on shader performance; it is essentially the same as copy-pasting the code.